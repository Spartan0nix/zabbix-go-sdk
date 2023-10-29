package zabbixgosdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// ApiClient is the default client structure to interact with the Zabbix API
type apiClient struct {
	// Handle HTTP request
	client *http.Client
	// URL of the Zabbix API
	// Format : http[s]://<zabbix-ip-or-dns>[/zabbix]/api_jsonrpc.php
	// [] are optional fields
	// <> are required fields
	url string
	// API token use to authenticate when calling the differents methods
	token string
}

// Request define the default body format to interact with the API
type request struct {
	// Jsonrpc define the version used of the json-rpc implementation
	// By default "2.0"
	JsonRpc string `json:"jsonrpc"`
	// Method is the name of the method called
	Method string `json:"method"`
	// Params contains the differents parameters to pass to the method
	Params interface{} `json:"params"`
	// Id is an arbitrary identifier of the request
	Id int `json:"id"`
	// Auth is the API token used to authenticate
	Auth string `json:"auth"`
}

// ResponseError define the error format returned in the Response body when a server error occurs.
type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

// responseResults define the available data type support for the Result field of response.
type ResponseResults interface {
	any
}

// Response define the base format for each API response.
// As the Result field can vary from one request to another, it must be defined for each method response.
type Response[T ResponseResults] struct {
	// Jsonrpc define the version used of the json-rpc implementation
	// By default "2.0"
	JsonRpc string `json:"jsonrpc"`
	// Result define the data returned by the method called.
	// Since the response differs from one method to another, the Result field need to be defined for each method.
	Result T `json:"result,omitempty"`
	// Error contains the more details about the server side error
	Error ResponseError `json:"error,omitempty"`
	// Id is an identifier for the corresponding request.
	Id int `json:"id"`
}

// CommonGetParameters define parameters supported by all get methods.
type CommonGetParameters struct {
	CountOutput            bool        `json:"countOutput,omitempty"`
	Editable               bool        `json:"editable,omitempty"`
	ExcludeSearch          bool        `json:"excludeSearch,omitempty"`
	Filter                 interface{} `json:"filter,omitempty"`
	Limit                  string      `json:"limit,omitempty"`
	Output                 interface{} `json:"output,omitempty"`
	PreserveKeys           bool        `json:"preservekeys,omitempty"`
	Search                 interface{} `json:"search,omitempty"`
	SearchByAny            bool        `json:"searchByAny,omitempty"`
	SearchWildcardsEnabled bool        `json:"searchWildcardsEnabled,omitempty"`
	SortField              []string    `json:"sortfield,omitempty"`
	SortOrder              []string    `json:"sortorder,omitempty"`
	StartSearch            bool        `json:"startSearch,omitempty"`
}

// Error return a more user-friendly error using the Error property in the given Response.
func (r ResponseError) Error() string {
	return "an error was returned by the Zabbix API: '" + r.Message + "' (data: '" + r.Data + "') (code: '" + strconv.Itoa(r.Code) + "')"
}

// Validate is used to check if an error was returned by the API in the current Response.
func (r *Response[T]) Validate() error {
	if r.Error.Code != 0 {
		return r.Error
	}

	return nil
}

func (a *apiClient) buildRequest(body interface{}) (*http.Request, error) {
	b, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, a.url, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json-rpc")

	return req, nil
}

// NewRequest is used to build a new HTTP Request with the given method and params.
// The request (r) struct is used as the base while method is used a r.Method and params as r.Params.
func (a *apiClient) NewRequest(method string, params interface{}) (*http.Request, error) {
	body := request{
		JsonRpc: "2.0",
		Method:  method,
		Params:  params,
		Id:      1,
		Auth:    a.token,
	}

	return a.buildRequest(body)
}

// NewCustomRequest is a similar function to NewRequest but uses a fully customisable body.
func (a *apiClient) NewCustomRequest(body interface{}) (*http.Request, error) {
	return a.buildRequest(body)
}

// ExecuteRequest is used to execute an HTTP Request.
// The HTTP response is returned a pointer.
func (a *apiClient) ExecuteRequest(r *http.Request) (*http.Response, error) {
	return a.client.Do(r)
}

// ConvertResponse is used to convert an HTTP response body to the value pointed by v.
// The HTTP body is closed a the end of the function.
func (a *apiClient) ConvertResponse(r *http.Response, v interface{}) error {
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(&v)
}

// Post is wrapper method to handle the differents actions required to execute an API request.
// The given method and body are used to build a new request.
// The API response is then converted and stored to the value pointed by v.
func (a *apiClient) Post(method string, params interface{}, v interface{}) error {
	// Create a new HTTP request
	req, err := a.NewRequest(method, params)
	if err != nil {
		return err
	}

	// Execute the HTTP request
	res, err := a.ExecuteRequest(req)
	if err != nil {
		return err
	}

	// Convert the HTTP response body
	err = a.ConvertResponse(res, &v)
	return err
}

// ResourceAlreadyExist check it the Data field in the given ResponseError indicate a 'resource already exists' type error.
func (a *apiClient) ResourceAlreadyExist(resource string, value string, err ResponseError) bool {
	substring := fmt.Sprintf("%s \"%s\" already exists", resource, value)

	if strings.Contains(err.Data, substring) {
		return true
	} else {
		return false
	}
}
