package main

import (
	"fmt"
	"github.com/twinj/uuid"
	"github.com/killbill/kbcli"
	"github.com/killbill/kbcli/models/gen"
	"os"
	//"runtime"
	"time"
	"strconv"
	"strings"
)

type Empty interface{}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

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

// Just to check how many ping calls we could make
func doPing(s *kbcli.Session) {
	kbcli.Ping(s)
}


// Do account/pm, auth, capture NOT using combo
func doAuthCapture(s *kbcli.Session, apiKey string) {

	// Query params reused across the test
	var params kbcli.QueryParams;

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


// Do account/pm, auth, capture  using combo

func doComboAuthCapture(s *kbcli.Session, key string, isCapture bool) {

	accountData := gen.AccountAttributes{Name: key, Email: key, Currency: "USD" }
	paymentMethodData := gen.PaymentMethodAttributes{PluginName: "__EXTERNAL_PAYMENT__", PluginInfo: gen.PaymentMethodPluginDetailAttributes{}}
	authorizationData := gen.PaymentTransactionAttributes{TransactionType: "AUTHORIZE", Amount: 12.5, Currency: accountData.Currency}

	var pluginProperties []gen.PluginPropertyAttributes;
	comboPaymentData := gen.ComboPaymentTransactionAttributes{Account:accountData, PaymentMethod:paymentMethodData, Transaction:authorizationData, PaymentMethodPluginProperties:pluginProperties, TransactionPluginProperties:pluginProperties}
	createdPayment, _ := kbcli.CreateComboAuthorization(s, &comboPaymentData, nil)
	//displaySuccessMsgOrAbort("Created COMBO payment transaction (AUTH)", "Failed to create payment transaction (AUTH)", err, createdPayment.PaymentId)

	if isCapture {
		captureData := gen.PaymentTransactionAttributes{PaymentId: createdPayment.PaymentId, TransactionType: "CAPTURE", Amount: 12.5, Currency: createdPayment.Currency}
		createdPayment, err := kbcli.CreatePaymentCapture(s, createdPayment.AccountId, &captureData, nil)
		displaySuccessMsgOrAbort("Created payment transaction (CAPTURE)", "Failed to create payment transaction (AUTH)", err, createdPayment.PaymentId)
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


	s := kbcli.CreateSession(ipAddr, ipPort, username, password, apiKey, apiSecret, createdBy, false)
	createdTenant, err := kbcli.CreateTenant(s)
	displaySuccessMsgOrAbort("Created tenant", "Failed to create tenant", err, createdTenant.ApiKey)

	var concurrency int = 10
	var sequentialIterations int = 1000

	sem := make(chan int, concurrency);

	ini := makeTimestamp()
	for i := 0; i < concurrency; i++ {
		go func(j int) {
			for k := 0; k < sequentialIterations; k++ {
				keyParts := []string{apiKey, "-", strconv.Itoa(i), "-", strconv.Itoa(k)}
				key := strings.Join(keyParts, "")
				doComboAuthCapture(s, key, false)
			}
			sem <- i
		}(i);
	}
	for i := 0; i < concurrency; i++ {
		<-sem
	}

	totalRequests := concurrency * sequentialIterations
	elapsed := makeTimestamp() - ini
	fmt.Printf("Processed %d requests in %d millis, rate = %v",
		totalRequests, elapsed, (float64(totalRequests) / float64(elapsed) * 1000))
}
