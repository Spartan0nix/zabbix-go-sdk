package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ApiUser struct {
	User string
	Pwd  string
}

type ApiClient struct {
	http  *http.Client
	Url   string
	User  ApiUser
	Token string
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

func NewApiClient() *ApiClient {
	return &ApiClient{
		http: &http.Client{},
		User: ApiUser{},
	}
}

func (s *ApiClient) ExecuteRequest(r *http.Request) ([]byte, error) {
	resp, err := s.http.Do(r)
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

func (a *ApiClient) Post(params interface{}) ([]byte, error) {
	b, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, a.Url, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json-rpc")

	res, err := a.ExecuteRequest(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
