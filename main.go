package main

import (
	"fmt"
	"log"
)

const (
	URL      = "http://localhost:8080/api_jsonrpc.php"
	USER     = "Admin"
	USER_PWD = "zabbix"
	// USER_PWD = "zabb"
)

func main() {
	client := NewZabbixService()
	client.SetUrl(URL)
	client.SetUser(&ApiUser{
		User: USER,
		Pwd:  USER_PWD,
	})

	err := client.Authenticate()
	if err != nil {
		log.Fatalf("Error when retrieving API credentials.\nReason : %v", err)
	}

	fmt.Println(client.Auth.Client.Token)
	// fmt.Println(client.UserGroup.Client.Token)

}
