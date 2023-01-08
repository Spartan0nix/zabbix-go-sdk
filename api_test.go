package zabbixgosdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func NewTestingService() (*ZabbixService, error) {
	client := NewZabbixService()
	client.SetUrl("http://localhost:4444/api_jsonrpc.php")
	client.SetUser(&ApiUser{
		User: "Admin",
		Pwd:  "zabbix",
	})

	err := client.Authenticate()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func TestNewRequest(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	req := client.Auth.Client.NewRequest("method.test", []string{
		"test-params",
	})

	if req.Method != "method.test" {
		t.Fatalf("Wrong method returned.\nValue returned : %s\nValue expected : 'method.test'", req.Method)
	}

	if req.Params.([]string)[0] != "test-params" {
		t.Fatalf("Wrong params returned.\nValue returned : %v\nValue expected : 'test-params'", req.Params)
	}
}

func TestExecuteRequest(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	params := client.UserGroup.Client.NewRequest("usergroup.get", UserGroupGetParameters{
		Filter: map[string]string{},
	})

	b, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("Error when convertig request to json.\nReason : %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, client.UserGroup.Client.Url, bytes.NewReader(b))
	if err != nil {
		t.Fatalf("Error when creating a new POST request.\nReason : %v", err)
	}

	req.Header.Set("Content-Type", "application/json-rpc")

	res, err := client.UserGroup.Client.ExecuteRequest(req)
	if err != nil {
		t.Fatalf("Error when executing POST request.\nReason %v", err)
	}

	if res == nil {
		t.Fatal("The POST request returned an empty []byte.")
	}
}

func TestPost(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	params := client.UserGroup.Client.NewRequest("usergroup.get", UserGroupGetParameters{
		Filter: map[string]string{},
	})

	res, err := client.UserGroup.Client.Post(params)
	if err != nil {
		t.Fatalf("Error when executing POST request.\nReason : %v", err)
	}

	if res.Result == nil {
		t.Fatalf("The POST request returned an empty result.\nResponse error : %v", res.Error)
	}
}

func TestConvertResponse(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	params := client.UserGroup.Client.NewRequest("usergroup.get", UserGroupGetParameters{
		Filter: map[string]string{},
	})

	res, err := client.UserGroup.Client.Post(params)
	if err != nil {
		t.Fatalf("Error when executing POST request.\nReason : %v", err)
	}

	if res.Result == nil {
		t.Fatalf("The POST request returned an empty result.\nResponse error : %v", res.Error)
	}

	groups := make([]UserGroup, 0)
	err = client.UserGroup.Client.ConvertResponse(*res, &groups)
	if err != nil {
		t.Fatalf("Error when converting the response.\nReason : %v", err)
	}

	if groups[0].Id == "" {
		t.Fatalf("Converted response is empty.\nResponse : %v", res)
	}
}

func TestError(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	params := client.UserGroup.Client.NewRequest("test.error", map[string]string{})

	res, err := client.UserGroup.Client.Post(params)
	if err != nil {
		t.Fatalf("Error when executing POST request.\nReason : %v", err)
	}

	err = client.Auth.Client.Error(*res)
	if err == nil {
		t.Fatalf("Error when formatting error message.\nResponse : %v", res)
	}
}

func TestResourceAlreadyExistSuccess(t *testing.T) {
	resource := "ExampleResource"
	value := "Test"
	err := ResponseError{
		Data: fmt.Sprintf("%s \"%s\" already exists", resource, value),
	}

	client := NewZabbixService()

	exist := client.Auth.Client.ResourceAlreadyExist(resource, value, err)

	if !exist {
		t.Fatalf("Resource do not match the exist.\nString : %s\nResource : %s\nValue : %s", err.Data, resource, value)
	}
}

func TestResourceAlreadyExistFail(t *testing.T) {
	resource := "ExampleResource"
	value := "Test"
	err := ResponseError{
		Data: fmt.Sprintf("%s \"%s\" already exists", resource, value),
	}

	client := NewZabbixService()

	exist := client.Auth.Client.ResourceAlreadyExist(resource, "WrongValue", err)

	if exist {
		t.Fatalf("Resource do not match the exist.\nString : %s\nResource : %s\nValue : %s", err.Data, resource, value)
	}
}
