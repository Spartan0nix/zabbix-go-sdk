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
		t.Fatalf("Request param 'user' was not set correctly.\nExpected : %s\nReturned : %s", user, body.Params["user"])
	}

	if body.Params["password"] != pwd {
		t.Fatalf("Request param 'password' was not set correctly.\nExpected : %s\nReturned : %s", pwd, body.Params["password"])
	}
}

func TestGetCredentials(t *testing.T) {
	user := "Admin"
	pwd := "zabbix"
	client := newTestingService()

	res, err := client.Auth.GetCredentials(user, pwd)
	if err != nil {
		t.Fatalf("Error when retrieving API credentials.\nReason : %v", err)
	}

	var token string
	err = client.Auth.Client.ConvertResponse(*res, &token)
	if err != nil {
		t.Fatalf("Error when converting auth response.\nReason : %v", err)
	}

	if token == "" {
		t.Fatalf("An empty token was returned.\nResponse error : %v", res.Error)
	}
}

func TestAuthenticate(t *testing.T) {
	client := newTestingService()

	err := client.Authenticate()
	if err != nil {
		t.Fatalf("Error when retrieving API credentials.\nReason : %v", err)
	}

	if client.Auth.Client.Token == "" {
		t.Fatal("Api token for the Auth service is null.\nExpected a string")
	}
}

func TestLogout(t *testing.T) {
	client := newTestingService()

	err := client.Logout()
	if err != nil {
		t.Fatalf("Error when executing Logout function.\nReason : %v", err)
	}
}

func TestLogoutWithoutToken(t *testing.T) {
	client := newTestingService()
	client.Auth.Client.Token = ""

	err := client.Logout()
	if err != nil {
		t.Fatalf("Logout function should not returned an error when the Token property of the ApiClient for the AuthService is empty.\nError returned : %v", err)
	}
}
