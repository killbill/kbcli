package main

import (
	"fmt"
	"github.com/killbill/kbcli"
	"github.com/killbill/kbcli/models/gen"
	"math/rand"
	"strconv"
	"strings"
	"os"
)

func generateRandomTestKey(suffix string) string {
	id := rand.Int63()
	parts := []string{suffix, strconv.FormatInt(id, 10)}
	return strings.Join(parts, "-")
}

func displaySuccessMsgOrAbort(msg string, errorMsg string, err error) {
	if err == nil {
		fmt.Println(msg)
	} else {
		fmt.Println("Exit test:")
		fmt.Println(errorMsg)
		fmt.Println("err = ", err)
		os.Exit(1)
	}
}

func main() {

	const username string = "admin"
	const password string = "password"

	const createdBy string = "admin"

	const ipAddr string= "127.0.0.1"
	const ipPort string= "8080"


	const apiSecret string = "foo"
	// The one that changes for each test iterations
	apiKey  := generateRandomTestKey("bob7")


	// Query params reused across the test
	var params kbcli.QueryParams;

	s := kbcli.CreateSession(ipAddr, ipPort, username, password, apiKey, apiSecret, createdBy, false)

	_, err := kbcli.CreateTenant(s)

	displaySuccessMsgOrAbort("Created tenant", "Failed to create tenant", err)

	accountData := gen.AccountAttributes{Name: apiKey, ExternalKey: apiKey, Email: apiKey, Currency: "USD" }
	createdAccount, err := kbcli.CreateAccount(s, &accountData, nil)

	displaySuccessMsgOrAbort("Created account", "Failed to create account", err)

	params = make(kbcli.QueryParams)
	params[kbcli.QUERY_EXTERNAL_KEY] = createdAccount.ExternalKey
	params[kbcli.QUERY_AUDIT] = "FULL"

	accountByKey, err := kbcli.GetAccount(s, "", &params)
	displaySuccessMsgOrAbort("Retrieved account by external key", "Failed retrieve account by external key", err)
	fmt.Println("account err:", err)
	fmt.Println("accountByKey:", accountByKey)


	paymentMethodData := gen.PaymentMethodAttributes{AccountId: createdAccount.AccountId, PluginName: "__EXTERNAL_PAYMENT__", PluginInfo: gen.PaymentMethodPluginDetailAttributes{}}
	params = make(kbcli.QueryParams)
	params[kbcli.QUERY_PAYMENT_METHOD_IS_DEFAULT] = "true"
	_, err = kbcli.CreatePaymentMethod(s, &paymentMethodData, &params)
	displaySuccessMsgOrAbort("Created payment method", "Failed to create payment method", err)


	authorizationData := gen.PaymentTransactionAttributes{TransactionType: "AUTHORIZE", Amount: 12.5, Currency: createdAccount.Currency}
	// TODO
	//authorizationData.PaymentId = nil
	_, err = kbcli.CreatePaymentTransaction(s, createdAccount.AccountId, &authorizationData, nil)
	displaySuccessMsgOrAbort("Created payment transaction (AUTH)", "Failed to create payment transaction (AUTH)", err)

}
