package main

import (
	"fmt"
	"github.com/killbill/kbcli"
)

func main() {

	const username string = "admin"
	const password string = "password"

	const createdBy string = "admin"

	const ipAddr string= "127.0.0.1"
	const ipPort string= "8080"

	const apiKey string = "bob45"
	const apiSecret string = "foo"

	s := kbcli.CreateSession(ipAddr, ipPort, username, password, apiKey, apiSecret, createdBy, false)

	resp, err := kbcli.CreateTenant(s)

	fmt.Println("response err:", err)
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Result:", resp.Result)
}
