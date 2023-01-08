package zabbixgosdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// ApiClient is the default client structure to interact with the Zabbix API
type ApiClient struct {
	// Handle HTTP request
	client *http.Client
	// URL of the Zabbix API
	// Format : http[s]://<zabbix-ip-or-dns>[/zabbix]/api_jsonrpc.php
	// [] are optional fields
	// <> are required fields
	Url string
	// API token use to authenticate when calling the differents methods
	Token string
}

// Request define the default body format to interact with the API
type Request struct {
	// Jsonrpc defined the version used of the json-rpc implementation
	// By default "2.0"
	Jsonrpc string `json:"jsonrpc"`
	// Method is the name of the method called
	Method string `json:"method"`
	// Params contains the differents parameters to pass to the method
	Params interface{} `json:"params"`
	// Id is an arbitrary identifier of the request
	Id int `json:"id"`
	// Auth is the API token used to authenticate
	Auth string `json:"auth"`
}

// ResponseError define the error format returned in the Response body when a server error occured
type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// Response define the response format returned by the API
type Response struct {
	// Jsonrpc defined the version used of the json-rpc implementation
	// By default "2.0"
	Jsonrpc string `json:"jsonrpc"`
	// Result is the data returned by the method called
	// Since the response differs from one method to another,
	// the result is returned raw ([]byte) and will need more processing
	Result json.RawMessage `json:"result"`
	// Error contains the more details about the server side error
	Error ResponseError `json:"error,omitempty"`
	// Id is an identifier for the corresponding request.
	Id int `json:"id"`
}

// NewRequest build a new API Request with the given method and params
func (a *ApiClient) NewRequest(method string, params interface{}) Request {
	return Request{
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
		Id:      1,
		Auth:    a.Token,
	}
}

// ExecuteRequest execute the given Request and return an []byte or an error
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

// Post is a method that accept a body an wrap multiple methods.
//
// 1. Convert the body to a JSON encoding.
//
// 2. Generate a new HTTP POST Request with the ApiClient.Url property and the enconded JSON body.
//
// 3. Set the 'Content-type' header to 'application/json-rpc'.
//
// 4. Execute the HTTP Request.
//
// 5. Convert the HTTP response to a API Response object.
//
// Return a Response object or an error that occured during the 5 previous steps
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

// ConvertResponse convert the Result property in the given Response to the given 'v' type.
// 'v' must be a pointer.
func (a *ApiClient) ConvertResponse(r Response, v interface{}) error {
	err := json.Unmarshal(r.Result, &v)
	if err != nil {
		// Check if an error was returned by the Zabbix Server.
		// If an error occure, the result field is empty.
		// This while leads to Unmarshal error.
		if r.Error.Code != 0 {
			err = a.Error(r)
			return err
		}

		return err
	}

	return nil
}

// Error return a more user-friendly error using the Error property in the given Response.
func (a *ApiClient) Error(r Response) error {
	return fmt.Errorf("Code : %d\nMessage : %s\nData : %s", r.Error.Code, r.Error.Message, r.Error.Data)
}

// ResourceAlreadyExist check it the Data field in the given ResponseError indicate a 'resource already exists' type error.
func (a *ApiClient) ResourceAlreadyExist(resource string, value string, err ResponseError) bool {
	re := regexp.MustCompile(fmt.Sprintf("%s \"%s\" already exists", resource, value) + `.*`)
	if re.Match([]byte(err.Data)) {
		return true
	} else {
		return false
	}
}
