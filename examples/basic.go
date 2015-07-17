package main

import (
	"fmt"
	"github.com/killbill/kbcli"
	"github.com/killbill/kbcli/models/gen"
)

func main() {

	const username string = "admin"
	const password string = "password"

	const createdBy string = "admin"

	const ipAddr string= "127.0.0.1"
	const ipPort string= "8080"

	const apiKey string = "bob56"
	const apiSecret string = "foo"

	s := kbcli.CreateSession(ipAddr, ipPort, username, password, apiKey, apiSecret, createdBy, false)

	createdTenant, err := kbcli.CreateTenant(s)

	fmt.Println("tenant err:", err)
	fmt.Println("createdTenant:", createdTenant)

	accountData := gen.AccountAttributes{Name: apiKey, ExternalKey: apiKey, Email: apiKey, Currency: "USD" }
	createdAccount, err := kbcli.CreateAccount(s, &accountData)

	fmt.Println("account err:", err)
	fmt.Println("createdAccount:", createdAccount)

	accountByKey, err := kbcli.GetAccountByKey(s,createdAccount.ExternalKey)
	fmt.Println("account err:", err)
	fmt.Println("accountByKey:", accountByKey)

}
