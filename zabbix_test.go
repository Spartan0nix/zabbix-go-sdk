package main

import (
	"testing"
)

func TestNewZabbixService(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	if client == nil {
		t.Fatalf("Error when creating new testing service.")
	}
}

func TestSetUrl(t *testing.T) {
	url := "localhost:7777"
	client := NewZabbixService()

	client.SetUrl(url)

	if client.Auth.Client.Url != url {
		t.Fatalf("Auth client 'url' property was not set correctly\n.Value : %s", client.Auth.Client.Url)
	}

	if client.UserGroup.Client.Url != url {
		t.Fatalf("UserGroup client 'url' property was not set correctly\n.Value : %s", client.UserGroup.Client.Url)
	}

	if client.UserMacro.Client.Url != url {
		t.Fatalf("UserMacro client 'url' property was not set correctly\n.Value : %s", client.UserGroup.Client.Url)
	}
}

func TestSetUser(t *testing.T) {
	user := "testing-user"
	pwd := "testing-password"

	client := NewZabbixService()

	client.SetUser(&ApiUser{
		User: user,
		Pwd:  pwd,
	})

	if client.Auth.User.User != user {
		t.Fatal("Auth client 'user' property was not set correctly.")
	}

	if client.Auth.User.Pwd != pwd {
		t.Fatal("Auth client 'pwd' property was not set correctly.")
	}
}

func TestToken(t *testing.T) {
	token := "random-complex-token"
	client := NewZabbixService()

	client.SetToken(token)

	if client.Auth.Client.Token != token {
		t.Fatalf("Auth client 'token property was not set correctly\n.Value : %s", client.Auth.Client.Token)
	}

	if client.UserGroup.Client.Token != token {
		t.Fatalf("UserGroup client 'token' property was not set correctly\n.Value : %s", client.UserGroup.Client.Token)
	}

	if client.UserMacro.Client.Token != token {
		t.Fatalf("UserMacro client 'token' property was not set correctly\n.Value : %s", client.UserGroup.Client.Token)
	}
}
