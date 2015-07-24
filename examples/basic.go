package main

import (
	"fmt"
	"github.com/twinj/uuid"
	"github.com/killbill/kbcli"
	"github.com/killbill/kbcli/models/gen"
	"os"
)

func generateRandomTestKey() string {
	return uuid.NewV1().String()
}

func displaySuccessMsgOrAbort(msg string, errorMsg string, err error, args ...interface{}) {
	if err == nil {
		fmt.Println(msg, args)
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
	apiKey  := generateRandomTestKey()

	// Query params reused across the test
	var params kbcli.QueryParams;

	s := kbcli.CreateSession(ipAddr, ipPort, username, password, apiKey, apiSecret, createdBy, false)

	createdTenant, err := kbcli.CreateTenant(s)

	displaySuccessMsgOrAbort("Created tenant", "Failed to create tenant", err, createdTenant.ApiKey)

	accountData := gen.AccountAttributes{Name: apiKey, ExternalKey: apiKey, Email: apiKey, Currency: "USD" }
	createdAccount, err := kbcli.CreateAccount(s, &accountData, nil)

	displaySuccessMsgOrAbort("Created account id = ", "Failed to create account", err, createdAccount.AccountId)

	params = make(kbcli.QueryParams)
	params[kbcli.QUERY_EXTERNAL_KEY] = createdAccount.ExternalKey
	params[kbcli.QUERY_AUDIT] = "FULL"

	accountByKey, err := kbcli.GetAccount(s, "", &params)
	displaySuccessMsgOrAbort("Retrieved account by external key", "Failed to retrieve account by external key", err, accountByKey.ExternalKey)


	paymentMethodData := gen.PaymentMethodAttributes{AccountId: createdAccount.AccountId, PluginName: "__EXTERNAL_PAYMENT__", PluginInfo: gen.PaymentMethodPluginDetailAttributes{}}
	params = make(kbcli.QueryParams)
	params[kbcli.QUERY_PAYMENT_METHOD_IS_DEFAULT] = "true"
	createdPaymentMethod, err := kbcli.CreatePaymentMethod(s, &paymentMethodData, &params)
	displaySuccessMsgOrAbort("Created payment method id = ", "Failed to create payment method", err, createdPaymentMethod.PaymentMethodId)


	authorizationData := gen.PaymentTransactionAttributes{TransactionType: "AUTHORIZE", Amount: 12.5, Currency: createdAccount.Currency}
	createdPyament, err := kbcli.CreatePaymentAuthorization(s, createdAccount.AccountId, &authorizationData, nil)
	displaySuccessMsgOrAbort("Created payment transaction (AUTH)", "Failed to create payment transaction (AUTH)", err, createdPyament.PaymentId)

	captureData := gen.PaymentTransactionAttributes{PaymentId: createdPyament.PaymentId, TransactionType: "CAPTURE", Amount: 12.5, Currency: createdAccount.Currency}
	createdPyament, err = kbcli.CreatePaymentCapture(s, createdAccount.AccountId, &captureData, nil)
	displaySuccessMsgOrAbort("Created payment transaction (CAPTURE)", "Failed to create payment transaction (AUTH)", err, createdPyament.PaymentId)

}
