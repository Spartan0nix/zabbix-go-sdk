package main

import (
	"reflect"
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

func TestNewApiClient(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	if reflect.TypeOf(client) != reflect.TypeOf(&ZabbixService{}) {
		t.Fail()
	}
}
