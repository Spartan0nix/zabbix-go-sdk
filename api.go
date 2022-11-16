package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type ApiClient struct {
	client *http.Client
	Url    string
	Token  string
}

type Request struct {
	Jsonrpc string      `json:"jsonrpc"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	Id      int         `json:"id"`
	Auth    string      `json:"auth"`
}

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type Response struct {
	Jsonrpc string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   ResponseError   `json:"error,omitempty"`
	Id      int             `json:"id"`
}

func (a *ApiClient) NewRequest(method string, params interface{}) Request {
	return Request{
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
		Id:      1,
		Auth:    a.Token,
	}
}

func (a *ApiClient) ExecuteRequest(r *http.Request) ([]byte, error) {
	resp, err := a.client.Do(r)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (a *ApiClient) Post(body interface{}) (*Response, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, a.Url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json-rpc")

	b, err = a.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	res := Response{}
	err = json.Unmarshal(b, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (a *ApiClient) ConvertResponse(r Response, v interface{}) error {
	err := json.Unmarshal(r.Result, &v)
	if err != nil {
		// Check if an error was returned by the Zabbix Server
		// If an error occure, the result field is empty
		// This while leads to Unmarshal error
		if r.Error.Code != 0 {
			err = a.Error(r)
			return err
		}

		return err
	}

	return nil
}

func (a *ApiClient) Error(r Response) error {
	return fmt.Errorf("Code : %d\nMessage : %s\nData : %s", r.Error.Code, r.Error.Message, r.Error.Data)
}

func (a *ApiClient) ResourceAlreadyExist(resource string, value string, err ResponseError) bool {
	re := regexp.MustCompile(fmt.Sprintf("%s \"%s\" already exists", resource, value) + `.*`)
	if re.Match([]byte(err.Data)) {
		return true
	} else {
		return false
	}
}
