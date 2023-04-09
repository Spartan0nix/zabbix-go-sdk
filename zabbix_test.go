package zabbixgosdk

import (
	"testing"
)

func TestNewZabbixService(t *testing.T) {
	client := NewZabbixService()
	if client == nil {
		t.Fatalf("A nol pointer was returned instead of *ZabbixService.")
	}
}

func TestSetUrl(t *testing.T) {
	url := "localhost:7777"
	client := NewZabbixService()
	client.SetUrl(url)

	if client.Auth.Client.Url != url {
		t.Fatalf("Auth client 'url' property was not set correctly\n.Expected : %s\nReturned : %s", url, client.Auth.Client.Url)
	}

	if client.UserGroup.Client.Url != url {
		t.Fatalf("UserGroup client 'url' property was not set correctly\n.Expected : %s\nReturned : %s", url, client.UserGroup.Client.Url)
	}

	if client.UserMacro.Client.Url != url {
		t.Fatalf("UserMacro client 'url' property was not set correctly\n.Expected : %s\nReturned : %s", url, client.UserMacro.Client.Url)
	}

	if client.HostGroup.Client.Url != url {
		t.Fatalf("HostGroup client 'url' property was not set correctly\n.Expected : %s\nReturned : %s", url, client.HostGroup.Client.Url)
	}

	if client.Template.Client.Url != url {
		t.Fatalf("Template client 'url' property was not set correctly\n.Expected : %s\nReturned : %s", url, client.Template.Client.Url)
	}

	if client.HostInterface.Client.Url != url {
		t.Fatalf("HostInterface client 'url' property was not set correctly\n.Expected : %s\nReturned : %s", url, client.HostInterface.Client.Url)
	}

	if client.Host.Client.Url != url {
		t.Fatalf("Host client 'url' property was not set correctly\n.Expected : %s\nReturned : %s", url, client.Host.Client.Url)
	}

	if client.Service.Client.Url != url {
		t.Fatalf("Service client 'url' property was not set correctly\n.Expected : %s\nReturned : %s", url, client.Service.Client.Url)
	}

	if client.IconMap.Client.Url != url {
		t.Fatalf("IconMap client 'url' property was not set correctly\n.Expected : %s\nReturned : %s", url, client.IconMap.Client.Url)
	}

	if client.Image.Client.Url != url {
		t.Fatalf("Image client 'url' property was not set correctly\n.Expected : %s\nReturned : %s", url, client.Image.Client.Url)
	}

	if client.Trigger.Client.Url != url {
		t.Fatalf("Trigger client 'url' property was not set correctly\n.Expected : %s\nReturned : %s", url, client.Trigger.Client.Url)
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
		t.Fatalf("Auth client 'user' property was not set correctly.\n.Expected : %s\nReturned : %s", user, client.Auth.User.User)
	}

	if client.Auth.User.Pwd != pwd {
		t.Fatalf("Auth client 'pwd' property was not set correctly.\n.Expected : %s\nReturned : %s", pwd, client.Auth.User.Pwd)
	}
}

func TestToken(t *testing.T) {
	token := "random-complex-token"
	client := NewZabbixService()

	client.SetToken(token)

	if client.Auth.Client.Token != token {
		t.Fatalf("Auth client 'token' property was not set correctly\n.Expected : %s\nReturned : %s", token, client.Auth.Client.Token)
	}

	if client.UserGroup.Client.Token != token {
		t.Fatalf("UserGroup client 'token' property was not set correctly\n.Expected : %s\nReturned : %s", token, client.UserGroup.Client.Token)
	}

	if client.UserMacro.Client.Token != token {
		t.Fatalf("UserMacro client 'token' property was not set correctly\n.Expected : %s\nReturned : %s", token, client.UserMacro.Client.Token)
	}

	if client.HostGroup.Client.Token != token {
		t.Fatalf("HostGroup client 'token' property was not set correctly\n.Expected : %s\nReturned : %s", token, client.HostGroup.Client.Token)
	}

	if client.Template.Client.Token != token {
		t.Fatalf("Template client 'token' property was not set correctly\n.Expected : %s\nReturned : %s", token, client.Template.Client.Token)
	}

	if client.HostInterface.Client.Token != token {
		t.Fatalf("HostInterface client 'token' property was not set correctly\n.Expected : %s\nReturned : %s", token, client.HostInterface.Client.Token)
	}

	if client.Host.Client.Token != token {
		t.Fatalf("Host client 'token' property was not set correctly\n.Expected : %s\nReturned : %s", token, client.Host.Client.Token)
	}

	if client.Service.Client.Token != token {
		t.Fatalf("Service client 'token' property was not set correctly\n.Expected : %s\nReturned : %s", token, client.Service.Client.Token)
	}

	if client.IconMap.Client.Token != token {
		t.Fatalf("IconMap client 'token' property was not set correctly\n.Expected : %s\nReturned : %s", token, client.IconMap.Client.Token)
	}

	if client.Image.Client.Token != token {
		t.Fatalf("Image client 'token' property was not set correctly\n.Expected : %s\nReturned : %s", token, client.Image.Client.Token)
	}

	if client.Trigger.Client.Token != token {
		t.Fatalf("Trigger client 'token' property was not set correctly\n.Expected : %s\nReturned : %s", token, client.Trigger.Client.Token)
	}
}
