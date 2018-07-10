package main

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/kbclient"
	"github.com/killbill/kbcli/kbclient/account"
)

// NewClient creates new kill bill client
func NewClient() *kbclient.KillBill {
	trp := httptransport.New("127.0.0.1:8080", "", nil)
	// Add text/xml producer which is not handled by openapi runtime.
	trp.Producers["text/xml"] = runtime.TextProducer()
	// Set this to true to dump http messages
	trp.Debug = false
	authWriter := httptransport.BasicAuth("admin" /*username*/, "password" /**password*/)
	client := kbclient.New(trp, strfmt.Default, authWriter, kbclient.KillbillDefaults{})

	// Set defaults. You can override them in each API call.
	apiKey := "bob"
	apiSecret := "lazar"
	createdBy := "John Doe"
	comment := "Created by John Doe"
	reason := ""

	client.SetDefaults(kbclient.KillbillDefaults{
		APIKey:    &apiKey,
		APISecret: &apiSecret,
		CreatedBy: &createdBy,
		Comment:   &comment,
		Reason:    &reason,
	})
	return client
}

func main() {
	c := NewClient()
	resp, err := c.Account.GetAccounts(context.Background(), &account.GetAccountsParams{})
	if err != nil {
		panic(err)
	}
	for _, acc := range resp.Payload {
		fmt.Printf("%s %s %s\n", acc.AccountID, acc.ExternalKey, acc.Email)
	}
}
