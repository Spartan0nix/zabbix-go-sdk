package main

import (
	"fmt"
	"log"
)

const (
	URL      = "http://localhost:8080/api_jsonrpc.php"
	USER     = "zabbix-api"
	USER_PWD = "kRTALYSh"
	// USER_PWD = "kRTALY"
)

func main() {
	client := NewApiClient()
	client.Url = URL
	client.User.User = USER
	client.User.Pwd = USER_PWD

	err := client.Authenticate()
	if err != nil {
		log.Fatalf("Error when retrieving API credentials.\nReason : %v1", err)
	}

	fmt.Println(client.Token)
}
