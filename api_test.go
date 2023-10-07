package zabbixgosdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

var testingClient *ZabbixService

func init() {
	testingClient = newTestingService()
}

func newTestingService() *ZabbixService {
	client := NewZabbixService()
	client.SetUrl("http://localhost:4444/api_jsonrpc.php")
	client.SetUser(&ApiUser{
		User: "Admin",
		Pwd:  "zabbix",
	})

	err := client.Authenticate()
	if err != nil {
		log.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	return client
}

// generateId is used to generate a random id to prevent duplicate entry during benchmarks.
func generateId() int {
	rand.NewSource(time.Now().UnixNano())
	value := rand.Intn(rand.Intn(9999))

	return value
}

func TestNewRequest(t *testing.T) {
	req := testingClient.Auth.Client.NewRequest("method.test", []string{
		"test-params",
	})

	if req.Method != "method.test" {
		t.Fatalf("Wrong method returned.\nExpected : 'method.test'\nReturned : %s", req.Method)
	}

	if req.Params.([]string)[0] != "test-params" {
		t.Fatalf("Wrong params returned.\nExpected : 'test-params'\nReturned : %v", req.Params)
	}
}

func BenchmarkNewRequest(b *testing.B) {
	params := []string{
		"test-params",
	}

	for i := 0; i < b.N; i++ {
		testingClient.Auth.Client.NewRequest("method.test", params)
	}
}

func TestExecuteRequest(t *testing.T) {
	params := testingClient.UserGroup.Client.NewRequest("usergroup.get", UserGroupGetParameters{
		Filter: map[string]string{},
	})

	b, err := json.Marshal(params)
	if err != nil {
		t.Fatalf("Error when convertig request to json.\nReason : %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, testingClient.UserGroup.Client.Url, bytes.NewReader(b))
	if err != nil {
		t.Fatalf("Error when creating a new POST request.\nReason : %v", err)
	}

	req.Header.Set("Content-Type", "application/json-rpc")

	res, err := testingClient.UserGroup.Client.ExecuteRequest(req)
	if err != nil {
		t.Fatalf("Error when executing POST request.\nReason %v", err)
	}

	if res == nil {
		t.Fatal("The POST request returned an empty slice of byte.")
	}
}

func BenchmarkExecuteRequest(b *testing.B) {
	params := testingClient.UserGroup.Client.NewRequest("usergroup.get", UserGroupGetParameters{
		Filter: map[string]string{},
	})

	byte, err := json.Marshal(params)
	if err != nil {
		b.Fatalf("error while converting request to json\nReason : %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, testingClient.UserGroup.Client.Url, bytes.NewReader(byte))
	if err != nil {
		b.Fatalf("error while creating a new POST request\nReason : %v", err)
	}

	req.Header.Set("Content-Type", "application/json-rpc")

	for i := 0; i < b.N; i++ {
		_, err = testingClient.UserGroup.Client.ExecuteRequest(req)
		if err != nil {
			b.Fatalf("error while executing ExecuteRequest function\nReason : %v", err)
		}
	}
}

func TestPost(t *testing.T) {
	params := testingClient.UserGroup.Client.NewRequest("usergroup.get", UserGroupGetParameters{
		Filter: map[string]string{},
	})

	res, err := testingClient.UserGroup.Client.Post(params)
	if err != nil {
		t.Fatalf("Error when executing POST request.\nReason : %v", err)
	}

	if res.Result == nil {
		t.Fatalf("The POST request returned an empty result.\nResponse error : %v", res.Error)
	}
}

func BenchmarkPost(b *testing.B) {
	params := testingClient.UserGroup.Client.NewRequest("usergroup.get", UserGroupGetParameters{
		Filter: map[string]string{},
	})

	var err error
	for i := 0; i < b.N; i++ {
		_, err = testingClient.UserGroup.Client.Post(params)
		if err != nil {
			b.Fatalf("error while executing Post function\nReason : %v", err)
		}
	}
}

func TestConvertResponse(t *testing.T) {
	params := testingClient.UserGroup.Client.NewRequest("usergroup.get", UserGroupGetParameters{
		Filter: map[string]string{},
	})

	res, err := testingClient.UserGroup.Client.Post(params)
	if err != nil {
		t.Fatalf("Error when executing POST request.\nReason : %v", err)
	}

	if res.Result == nil {
		t.Fatalf("The POST request returned an empty result.\nResponse error : %v", res.Error)
	}

	groups := make([]UserGroup, 0)
	err = testingClient.UserGroup.Client.ConvertResponse(*res, &groups)
	if err != nil {
		t.Fatalf("Error when converting the response.\nReason : %v", err)
	}

	if groups[0].UsrGrpId == "" {
		t.Fatalf("Converted response is empty.\nResponse : %v", res)
	}
}

func BenchmarkConvertResponse(b *testing.B) {
	params := testingClient.UserGroup.Client.NewRequest("usergroup.get", UserGroupGetParameters{
		Filter: map[string]string{},
	})

	res, err := testingClient.UserGroup.Client.Post(params)
	if err != nil {
		b.Fatalf("error when executing POST request\nReason : %v", err)
	}

	if res.Result == nil {
		b.Fatalf("the request returned an empty result\nResponse error : %v", res.Error)
	}

	groups := make([]UserGroup, 0)
	for i := 0; i < b.N; i++ {
		groups = []UserGroup{}

		err = testingClient.UserGroup.Client.ConvertResponse(*res, &groups)
		if err != nil {
			b.Fatalf("error while executing ConvertResponse function\nReason : %v", err)
		}
	}
}

func TestError(t *testing.T) {
	params := testingClient.UserGroup.Client.NewRequest("test.error", map[string]string{})

	res, err := testingClient.UserGroup.Client.Post(params)
	if err != nil {
		t.Fatalf("Error when executing POST request.\nReason : %v", err)
	}

	err = testingClient.Auth.Client.Error(*res)
	if err == nil {
		t.Fatalf("Error when formatting error message.\nResponse : %v", res)
	}
}

func BenchmarkError(b *testing.B) {
	params := testingClient.UserGroup.Client.NewRequest("test.error", map[string]string{})

	res, err := testingClient.UserGroup.Client.Post(params)
	if err != nil {
		b.Fatalf("error while executing POST request\nReason : %v", err)
	}

	for i := 0; i < b.N; i++ {
		err = testingClient.Auth.Client.Error(*res)
		if err == nil {
			b.Fatalf("Error while executing Error function\nResponse : %v\nError : %v", res, err)
		}
	}
}

func TestResourceAlreadyExistSuccess(t *testing.T) {
	resource := "ExampleResource"
	value := "Test"
	err := ResponseError{
		Data: fmt.Sprintf("%s \"%s\" already exists", resource, value),
	}

	exist := testingClient.Auth.Client.ResourceAlreadyExist(resource, value, err)
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

	exist := testingClient.Auth.Client.ResourceAlreadyExist(resource, "WrongValue", err)

	if exist {
		t.Fatalf("Resource do not match the exist.\nString : %s\nResource : %s\nValue : %s", err.Data, resource, value)
	}
}

func BenchmarkResourceAlreadyExist(b *testing.B) {
	resource := "ExampleResource"
	value := "Test"
	err := ResponseError{
		Data: fmt.Sprintf("%s \"%s\" already exists", resource, value),
	}

	for i := 0; i < b.N; i++ {
		testingClient.Auth.Client.ResourceAlreadyExist(resource, value, err)
	}

}

func TestCheckConnectivitySuccess(t *testing.T) {
	err := testingClient.Auth.Client.CheckConnectivity()
	if err != nil {
		t.Fatalf("Error when executing CheckConnectivity.\nReason : %v", err)
	}
}

func TestCheckConnectivityFail(t *testing.T) {
	client := newTestingService()
	client.Auth.Client.Url = "http://localhost:7777"

	err := client.Auth.Client.CheckConnectivity()
	if err == nil {
		t.Fatalf("CheckConnectivity should returned an error when Zabbix server is unreachable.")

	}

	expected_err := fmt.Sprintf("connectivity check failed for Zabbix server : %s", client.Auth.Client.Url)
	if err.Error() != expected_err {
		t.Fatalf("Expected error : %s\nError returned : %s", expected_err, err.Error())
	}
}

func TestCheckConnectivityMissingUrl(t *testing.T) {
	client := newTestingService()
	client.Auth.Client.Url = ""

	err := client.Auth.Client.CheckConnectivity()
	if err == nil {
		t.Fatalf("CheckConnectivity should returned an error when the 'Url' property is not set.")
	}
}

func BenchmarkCheckConnectivity(b *testing.B) {
	var err error

	for i := 0; i < b.N; i++ {
		err = testingClient.Auth.Client.CheckConnectivity()
		if err != nil {
			b.Fatalf("error while executing CheckConnectivity function\nReason : %v", err)
		}
	}
}
