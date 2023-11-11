package zabbixgosdk

import (
	"log"
	"math/rand"
	"testing"
	"time"
)

var testingClient *ZabbixService

func init() {
	testingClient = newTestingService()
}

func newTestingService() *ZabbixService {
	client := NewZabbixService("http://localhost:4444/api_jsonrpc.php")
	err := client.Authenticate("Admin", "zabbix")
	if err != nil {
		log.Fatalf("Error when executing Authenticate function on the ZabbixService\nReason: %v", err)
	}

	return client
}

// generateId is used to generate a random id to prevent duplicate entry during benchmarks.
func generateId() int {
	rand.NewSource(time.Now().UnixNano())
	value := rand.Intn(9999)

	return value
}

func TestNewZabbixService(t *testing.T) {
	client := NewZabbixService("http://localhost:2222")
	if client == nil {
		t.Fatalf("A nil pointer was returned instead of *ZabbixService.")
	}
}

func BenchmarkNewZabbixService(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewZabbixService("http://localhost:2222")
	}
}

func TestSetUrl(t *testing.T) {
	client := NewZabbixService("http://localhost:2222")
	url := "localhost:7777"
	client.SetUrl(url)

	if client.User.client.url != url {
		t.Fatalf("User client 'url' property was not set correctly\nExpected : %s\nReturned : %s", url, client.User.client.url)
	}

	// if client.UserGroup.client.url != url {
	// 	t.Fatalf("UserGroup client 'url' property was not set correctly\nExpected : %s\nReturned : %s", url, client.UserGroup.client.url)
	// }

	// if client.UserMacro.client.url != url {
	// 	t.Fatalf("UserMacro client 'url' property was not set correctly\nExpected : %s\nReturned : %s", url, client.UserMacro.client.url)
	// }

	if client.HostGroup.client.url != url {
		t.Fatalf("HostGroup client 'url' property was not set correctly\nExpected : %s\nReturned : %s", url, client.HostGroup.client.url)
	}

	if client.Template.client.url != url {
		t.Fatalf("Template client 'url' property was not set correctly\nExpected : %s\nReturned : %s", url, client.Template.client.url)
	}

	// if client.HostInterface.client.url != url {
	// 	t.Fatalf("HostInterface client 'url' property was not set correctly\nExpected : %s\nReturned : %s", url, client.HostInterface.client.url)
	// }

	if client.Host.client.url != url {
		t.Fatalf("Host client 'url' property was not set correctly\nExpected : %s\nReturned : %s", url, client.Host.client.url)
	}

	// if client.Service.client.url != url {
	// 	t.Fatalf("Service client 'url' property was not set correctly\nExpected : %s\nReturned : %s", url, client.Service.client.url)
	// }

	// if client.IconMap.client.url != url {
	// 	t.Fatalf("IconMap client 'url' property was not set correctly\nExpected : %s\nReturned : %s", url, client.IconMap.client.url)
	// }

	// if client.Image.client.url != url {
	// 	t.Fatalf("Image client 'url' property was not set correctly\nExpected : %s\nReturned : %s", url, client.Image.client.url)
	// }

	// if client.Trigger.client.url != url {
	// 	t.Fatalf("Trigger client 'url' property was not set correctly\nExpected : %s\nReturned : %s", url, client.Trigger.client.url)
	// }

	// if client.Map.client.url != url {
	// 	t.Fatalf("Map client 'url' property was not set correctly\nExpected : %s\nReturned : %s", url, client.Map.client.url)
	// }
}

func BenchmarkSetUrl(b *testing.B) {
	client := NewZabbixService("http://localhost:2222")
	url := "http://localhost:4444"

	for i := 0; i < b.N; i++ {
		client.SetUrl(url)
	}
}

func TestToken(t *testing.T) {
	token := "random-complex-token"
	client := NewZabbixService("http://localhost:2222")

	client.SetToken(token)

	if client.User.client.token != token {
		t.Fatalf("User client 'token' property was not set correctly\nExpected : %s\nReturned : %s", token, client.User.client.token)
	}

	// if client.UserGroup.client.token != token {
	// 	t.Fatalf("UserGroup client 'token' property was not set correctly\nExpected : %s\nReturned : %s", token, client.UserGroup.client.token)
	// }

	// if client.UserMacro.client.token != token {
	// 	t.Fatalf("UserMacro client 'token' property was not set correctly\nExpected : %s\nReturned : %s", token, client.UserMacro.client.token)
	// }

	if client.HostGroup.client.token != token {
		t.Fatalf("HostGroup client 'token' property was not set correctly\nExpected : %s\nReturned : %s", token, client.HostGroup.client.token)
	}

	if client.Template.client.token != token {
		t.Fatalf("Template client 'token' property was not set correctly\nExpected : %s\nReturned : %s", token, client.Template.client.token)
	}

	// if client.HostInterface.client.token != token {
	// 	t.Fatalf("HostInterface client 'token' property was not set correctly\nExpected : %s\nReturned : %s", token, client.HostInterface.client.token)
	// }

	if client.Host.client.token != token {
		t.Fatalf("Host client 'token' property was not set correctly\nExpected : %s\nReturned : %s", token, client.Host.client.token)
	}

	// if client.Service.client.token != token {
	// 	t.Fatalf("Service client 'token' property was not set correctly\nExpected : %s\nReturned : %s", token, client.Service.client.token)
	// }

	// if client.IconMap.client.token != token {
	// 	t.Fatalf("IconMap client 'token' property was not set correctly\nExpected : %s\nReturned : %s", token, client.IconMap.client.token)
	// }

	// if client.Image.client.token != token {
	// 	t.Fatalf("Image client 'token' property was not set correctly\nExpected : %s\nReturned : %s", token, client.Image.client.token)
	// }

	// if client.Trigger.client.token != token {
	// 	t.Fatalf("Trigger client 'token' property was not set correctly\nExpected : %s\nReturned : %s", token, client.Trigger.client.token)
	// }

	// if client.Map.client.token != token {
	// 	t.Fatalf("Map client 'token' property was not set correctly\nExpected : %s\nReturned : %s", token, client.Trigger.client.token)
	// }
}

func BenchmarkSetToken(b *testing.B) {
	client := NewZabbixService("http://localhost:2222")

	for i := 0; i < b.N; i++ {
		client.SetToken("testing-token")
	}
}

func TestAuthenticate(t *testing.T) {
	service := NewZabbixService("http://localhost:4444/api_jsonrpc.php")
	err := service.Authenticate("Admin", "zabbix")
	if err != nil {
		t.Fatalf("error while executing Authenticate function\nReason: %v", err)
	}
}

func BenchmarkAuthenticate(b *testing.B) {
	var err error
	service := NewZabbixService("http://localhost:4444/api_jsonrpc.php")

	for i := 0; i < b.N; i++ {
		err = service.Authenticate("Admin", "zabbix")
		if err != nil {
			b.Fatalf("error while executing Authenticate function\nReason: %v", err)
		}
	}
}

func TestLogout(t *testing.T) {
	service := NewZabbixService("http://localhost:4444/api_jsonrpc.php")
	err := service.Authenticate("Admin", "zabbix")
	if err != nil {
		t.Fatalf("error while executing Authenticate function\nReason: %v", err)
	}

	err = service.Logout()
	if err != nil {
		t.Fatalf("error while executing Logout function\nReason: %v", err)
	}
}

func BenchmarkLogout(b *testing.B) {
	var err error
	service := NewZabbixService("http://localhost:4444/api_jsonrpc.php")

	for i := 0; i < b.N; i++ {
		err = service.Authenticate("Admin", "zabbix")
		if err != nil {
			b.Fatalf("error while executing Authenticate function\nReason: %v", err)
		}

		err = service.Logout()
		if err != nil {
			b.Fatalf("error while executing Logout function\nReason: %v", err)
		}
	}
}

func TestGetApiVersion(t *testing.T) {
	v, err := testingClient.GetApiVersion()
	if err != nil {
		t.Fatalf("error while executing GetApiVersion\nReason: %v", err)
	}

	if v == "" {
		t.Fatal("an empty string was returned instead of the Zabbix API version")
	}
}

func BenchmarkGetApiVersion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, err := testingClient.GetApiVersion()
		if err != nil {
			b.Fatalf("error while executing GetApiVersion\nReason: %v", err)
		}
	}
}

func TestCheckConnectivity(t *testing.T) {
	err := testingClient.CheckConnectivity()
	if err != nil {
		t.Fatalf("error while executing CheckConnectivity\nReason: %v", err)
	}

}

func BenchmarkCheckConnectivity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		err := testingClient.CheckConnectivity()
		if err != nil {
			b.Fatalf("error while executing CheckConnectivity\nReason: %v", err)
		}
	}
}
