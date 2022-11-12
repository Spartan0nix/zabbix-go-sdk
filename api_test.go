package main

import (
	"reflect"
	"testing"
)

func newTestingService() *ZabbixService {
	client := NewZabbixService()
	client.SetUrl("http://localhost:4444/")

	return client
}

func TestNewApiClient(t *testing.T) {
	client := newTestingService()

	if reflect.TypeOf(client) != reflect.TypeOf(&ApiClient{}) {
		t.Fail()
	}
}
