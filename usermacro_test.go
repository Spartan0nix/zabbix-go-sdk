package main

import "testing"

func TestUserMacroCreate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	h, err := client.UserMacro.Create(HostMacroCreateParameters{
		Hostid: 10084,
		Macro:  "{$TEST}",
		Value:  "test",
	})

	if err != nil {
		t.Fatalf("Error when creating user macro.\nReason : %v", err)
	}

	if h.Hostmacroids[0] == "" {
		t.Fatalf("Create request returned an empty response.")
	}
}

func TestUserMacroCreateWrongFormat(t *testing.T) {
	client := NewZabbixService()

	_, err := client.UserMacro.Create(HostMacroCreateParameters{
		Hostid: 10084,
		Macro:  "TEST",
		Value:  "test",
	})

	if err == nil {
		t.Fatal("No error returned when trying to create a user macro without following Zabbix required format")
	}
}
