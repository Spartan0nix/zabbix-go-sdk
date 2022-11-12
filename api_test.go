package main

import (
	"reflect"
	"testing"
)

func newTestingClient() *ApiClient {
	client := NewApiClient()
	client.Url = "http://localhost:4444/"

	return client
}

func TestNewApiClient(t *testing.T) {
	client := NewApiClient()

	if reflect.TypeOf(client) != reflect.TypeOf(&ApiClient{}) {
		t.Fail()
	}
}
