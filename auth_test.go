package zabbixgosdk

import (
	"testing"
)

func TestNewAuthRequest(t *testing.T) {
	user := "testing-user"
	pwd := "testing-pwd"

	client := NewZabbixService()

	body := client.Auth.NewAuthRequest(user, pwd)

	if body.Params["user"] != user {
		t.Fatalf("Request param 'user' was not set correctly\n.Value : %s", body.Params["user"])
	}

	if body.Params["password"] != pwd {
		t.Fatalf("Request param 'password' was not set correctly\n.Value : %s", body.Params["password"])
	}
}

func TestGetCredentials(t *testing.T) {
	user := "Admin"
	pwd := "zabbix"

	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	res, err := client.Auth.GetCredentials(user, pwd)
	if err != nil {
		t.Fatalf("Error when retrieving API credentials.\nReason : %v\nThe following username / password were used :\n- %s / %s", err, user, pwd)
	}

	var token string
	err = client.Auth.Client.ConvertResponse(*res, &token)
	if err != nil {
		t.Fatalf("Error when converting response.\nReason : %v", err)
	}

	if token == "" {
		t.Fatalf("An empty token was returned.\nResponse error : %v", res.Error)
	}
}

func TestAuthenticate(t *testing.T) {
	client, err := NewTestingService()
	if err != nil {
		t.Fatalf("Error when creating new testing service.\nReason : %v", err)
	}

	err = client.Authenticate()
	if err != nil {
		t.Fatalf("Error when retrieving API credentials.\nReason : %v", err)
	}

	if client.Auth.Client.Token == "" {
		t.Fatalf("Auth client 'token property was not set correctly\n.Value : %s", client.Auth.Client.Token)
	}
}
