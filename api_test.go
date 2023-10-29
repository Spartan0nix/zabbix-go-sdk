package zabbixgosdk

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestError(t *testing.T) {
	testCode := 1
	testMessage := "testing error message"
	testData := "testing error data"
	errExpected := fmt.Sprintf("an error was returned by the Zabbix API: '%s' (data: '%s') (code: '%d')", testMessage, testData, testCode)

	err := ResponseError{
		Code:    testCode,
		Message: testMessage,
		Data:    testData,
	}

	errReturned := err.Error()
	if errReturned != errExpected {
		t.Fatalf("wrong format returned\nExpected: '%s'\nReturned: '%s'", errExpected, errReturned)
	}
}

func BenchmarkError(b *testing.B) {
	err := ResponseError{
		Code:    1,
		Message: "testing error message",
		Data:    "testing error data",
	}

	for i := 0; i < b.N; i++ {
		_ = err.Error()
	}
}

func TestValidate(t *testing.T) {
	testCode := 1
	testMessage := "testing error message"
	testData := "testing error data"
	errExpected := fmt.Sprintf("an error was returned by the Zabbix API: '%s' (data: '%s') (code: '%d')", testMessage, testData, testCode)

	res := Response[string]{
		JsonRpc: "2.0",
		Result:  "",
		Error: ResponseError{
			Code:    testCode,
			Message: testMessage,
			Data:    testData,
		},
		Id: 1,
	}

	errReturned := res.Validate()

	if errReturned == nil {
		t.Fatalf("expected an error while executing Validate, a nil value was returned")
	}

	if errReturned.Error() != errExpected {
		t.Fatalf("wrong format returned\nExpected: '%s'\nReturned: '%s'", errExpected, errReturned.Error())
	}
}

func TestValidateNoError(t *testing.T) {
	res := Response[string]{
		JsonRpc: "2.0",
		Result:  "my-response",
		Id:      1,
	}

	err := res.Validate()
	if err != nil {
		t.Fatalf("expected an nil error while executing Validate\nValue returned: %v", err)
	}

}

func BenchmarkValidate(b *testing.B) {
	res := Response[string]{
		JsonRpc: "2.0",
		Result:  "",
		Error: ResponseError{
			Code:    1,
			Message: "testing error message",
			Data:    "testing error data",
		},
		Id: 1,
	}

	var err error
	for i := 0; i < b.N; i++ {
		err = res.Validate()
		if err == nil {
			b.Fatalf("expected an error while executing Validate, a nil value was returned")
		}
	}
}

func BenchmarkValidateNoError(b *testing.B) {
	res := Response[string]{
		JsonRpc: "2.0",
		Result:  "my-response",
		Id:      1,
	}

	var err error
	for i := 0; i < b.N; i++ {
		err = res.Validate()
		if err != nil {
			b.Fatalf("expected an nil error while executing Validate\nValue returned: %v", err)
		}
	}
}

func TestBuildRequest(t *testing.T) {
	req, err := testingClient.User.client.buildRequest(map[string]string{
		"key": "value",
	})

	if err != nil {
		t.Fatalf("error while executing buildRequest function\nReason: %v", err)
	}

	if req == nil {
		t.Fatalf("expected an *http.Request to be returned\nelement returned: %v", reflect.TypeOf(req))
	}
}

func BenchmarkBuildRequest(b *testing.B) {
	params := map[string]string{
		"key": "value",
	}

	for i := 0; i < b.N; i++ {
		req, err := testingClient.User.client.buildRequest(params)
		if err != nil {
			b.Fatalf("error while executing buildRequest function\nReason: %v", err)
		}

		if req == nil {
			b.Fatalf("expected an *http.Request to be returned\nelement returned: %v", reflect.TypeOf(req))
		}
	}
}

func TestNewRequest(t *testing.T) {
	req, err := testingClient.User.client.NewRequest("method.test", []string{
		"test-params",
	})

	if err != nil {
		t.Fatalf("error while executing NewRequest function\nReason: %v", err)
	}

	if req == nil {
		t.Fatalf("expected an *http.Request to be returned\nelement returned: %v", reflect.TypeOf(req))
	}
}

func BenchmarkNewRequest(b *testing.B) {
	params := []string{
		"test-params",
	}

	for i := 0; i < b.N; i++ {
		testingClient.User.client.NewRequest("method.test", params)
	}
}

func TestNewCustomRequest(t *testing.T) {
	type testingRequest struct {
		JsonRpc string            `json:"jsonrpc"`
		Method  string            `json:"method"`
		Params  map[string]string `json:"params"`
		Id      int               `json:"id"`
	}

	body := testingRequest{
		JsonRpc: "2.0",
		Method:  "method.test",
		Params: map[string]string{
			"field1": "value1",
			"field2": "value2",
		},
		Id: 1,
	}

	req, err := testingClient.User.client.NewCustomRequest(body)

	if err != nil {
		t.Fatalf("error while executing NewCustomRequest function\nReason: %v", err)
	}

	if req == nil {
		t.Fatalf("expected an *http.Request to be returned\nelement returned: %v", reflect.TypeOf(req))
	}
}

func BenchmarkNewCustomRequest(b *testing.B) {
	type testingRequest struct {
		JsonRpc string            `json:"jsonrpc"`
		Method  string            `json:"method"`
		Params  map[string]string `json:"params"`
		Id      int               `json:"id"`
	}

	body := testingRequest{
		JsonRpc: "2.0",
		Method:  "method.test",
		Params: map[string]string{
			"field1": "value1",
			"field2": "value2",
		},
		Id: 1,
	}

	for i := 0; i < b.N; i++ {
		testingClient.User.client.NewCustomRequest(body)
	}
}

func TestExecuteRequest(t *testing.T) {
	req, err := testingClient.HostGroup.client.NewRequest("usergroup.get", HostGroupGetParameters{
		CommonGetParameters: CommonGetParameters{
			Filter: map[string]string{},
		},
	})

	if err != nil {
		t.Fatalf("error while executing NewRequest function\nReason: %v", err)
	}

	res, err := testingClient.HostGroup.client.ExecuteRequest(req)
	if err != nil {
		t.Fatalf("Error when executing ExecuteRequest function\nReason: %v", err)
	}

	if res == nil {
		t.Fatal("A nil value was returned instead of *http.Response")
	}
}

func BenchmarkExecuteRequest(b *testing.B) {
	body := request{
		JsonRpc: "2.0",
		Method:  "usergroup.get",
		Params: &HostGroupGetParameters{
			CommonGetParameters: CommonGetParameters{
				Filter: map[string]string{},
			},
		},
		Id:   1,
		Auth: testingClient.HostGroup.client.token,
	}

	baseReq, err := testingClient.HostGroup.client.NewCustomRequest(body)
	if err != nil {
		b.Fatalf("error while executing NewRequest function\nReason: %v", err)
	}

	var buf bytes.Buffer
	buf.ReadFrom(baseReq.Body)

	for i := 0; i < b.N; i++ {
		req := baseReq.Clone(context.Background())
		req.Body = io.NopCloser(bytes.NewReader(buf.Bytes()))

		res, err := testingClient.HostGroup.client.ExecuteRequest(req)
		if err != nil {
			b.Fatalf("Error when executing ExecuteRequest function\nReason: %v", err)
		}

		if res == nil {
			b.Fatal("A nil value was returned instead of *http.Response")
		}
	}
}

func TestConvertResponse(t *testing.T) {
	resBody := Response[string]{
		JsonRpc: "2.0",
		Result:  "my-response",
		Id:      1,
	}

	body, err := json.Marshal(resBody)
	if err != nil {
		t.Fatalf("error while converting body to json\nReason: %v", err)
	}

	res := http.Response{
		Status: http.StatusText(200),
		Body:   io.NopCloser(bytes.NewReader(body)),
	}

	expectedResponse := Response[string]{}
	err = testingClient.HostGroup.client.ConvertResponse(&res, &expectedResponse)
	if err != nil {
		t.Fatalf("error while executing ConvertResponse function\nReason: %v", err)
	}

	if expectedResponse.JsonRpc != "2.0" {
		t.Fatalf("expected '%s' for the 'JsonRpc' field\nValue returned: '%s'", "2.0", expectedResponse.JsonRpc)
	}

	if expectedResponse.Result != "my-response" {
		t.Fatalf("expected '%s' for the 'Result' field\nValue returned: '%s'", "my-response", expectedResponse.Result)
	}

	if expectedResponse.Id != 1 {
		t.Fatalf("expected '%d' for the 'Id' field\nValue returned: '%d'", 1, expectedResponse.Id)
	}
}

func BenchmarkConvertResponse(b *testing.B) {
	resBody := Response[string]{
		JsonRpc: "2.0",
		Result:  "my-response",
		Id:      1,
	}

	body, err := json.Marshal(resBody)
	if err != nil {
		b.Fatalf("error while converting body to json\nReason: %v", err)
	}

	for i := 0; i < b.N; i++ {
		res := http.Response{
			Status: http.StatusText(200),
			Body:   io.NopCloser(bytes.NewReader(body)),
		}

		expectedResponse := Response[string]{}
		err = testingClient.HostGroup.client.ConvertResponse(&res, &expectedResponse)
		if err != nil {
			b.Fatalf("error while executing ConvertResponse function\nReason: %v", err)
		}

	}
}

func TestPost(t *testing.T) {
	res := Response[[]*HostGroup]{}
	params := HostGroupGetParameters{
		CommonGetParameters: CommonGetParameters{
			Filter: map[string]string{},
		},
	}

	err := testingClient.HostGroup.client.Post("hostgroup.get", &params, &res)
	if err != nil {
		t.Fatalf("error while executing Post function\nReason: %v", err)
	}

	if len(res.Result) == 0 {
		t.Fatalf("an empty list of HostGroup was returned, expected at least the default groups")
	}
}

func BenchmarkPost(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := Response[[]*HostGroup]{}
		params := HostGroupGetParameters{
			CommonGetParameters: CommonGetParameters{
				Filter: map[string]string{},
			},
		}

		err := testingClient.HostGroup.client.Post("hostgroup.get", &params, &res)
		if err != nil {
			b.Fatalf("error while executing Post function\nReason: %v", err)
		}

		if len(res.Result) == 0 {
			b.Fatalf("an empty list of HostGroup was returned, expected at least the default groups")
		}
	}
}

func TestResourceAlreadyExist(t *testing.T) {
	resource := "host-group"
	value := "my-host-group"
	err := ResponseError{
		Code:    1,
		Data:    "host-group \"my-host-group\" already exists",
		Message: "",
	}

	b := testingClient.HostGroup.client.ResourceAlreadyExist(resource, value, err)
	if !b {
		t.Fatalf("expected a positive boolean to be returned, '%t' was returned instead", b)
	}
}

func TestResourceAlreadyExistNoMatch(t *testing.T) {
	resource := "host-group"
	value := "my-host-group"
	err := ResponseError{
		Code:    1,
		Data:    "host-group \"my-host-group-2\" already exists",
		Message: "",
	}

	b := testingClient.HostGroup.client.ResourceAlreadyExist(resource, value, err)
	if b {
		t.Fatalf("expected a negative boolean to be returned, '%t' was returned instead", b)
	}
}

func BenchmarkResourceAlreadyExist(b *testing.B) {
	resource := "host-group"
	value := "my-host-group"
	err := ResponseError{
		Code:    1,
		Data:    "host-group \"my-host-group\" already exists",
		Message: "",
	}

	for i := 0; i < b.N; i++ {
		testingClient.HostGroup.client.ResourceAlreadyExist(resource, value, err)
	}
}

func BenchmarkResourceAlreadyExistNoMatch(b *testing.B) {
	resource := "host-group"
	value := "my-host-group"
	err := ResponseError{
		Code:    1,
		Data:    "host-group \"my-host-group-2\" already exists",
		Message: "",
	}

	for i := 0; i < b.N; i++ {
		testingClient.HostGroup.client.ResourceAlreadyExist(resource, value, err)
	}
}
