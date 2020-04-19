package main

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v2/kbclient"
	"github.com/killbill/kbcli/v2/kbclient/account"
)

// NewClient creates new kill bill client
func NewClient() *kbclient.KillBill {
	trp := httptransport.New("127.0.0.1:8080", "", nil)
	// Add text/xml producer which is not handled by openapi runtime.
	trp.Producers["text/xml"] = runtime.TextProducer()
	// Set this to true to dump http messages
	trp.Debug = false
	apiKey := "bob"
	apiSecret := "lazar"
	authWriter := runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		encoded := base64.StdEncoding.EncodeToString([]byte("admin" /*username*/ + ":" + "password" /**password*/))
		if err := r.SetHeaderParam("Authorization", "Basic "+encoded); err != nil {
			return err
		}
		if err := r.SetHeaderParam("X-KillBill-ApiKey", apiKey); err != nil {
			return err
		}
		if err := r.SetHeaderParam("X-KillBill-ApiSecret", apiSecret); err != nil {
			return err
		}
		return nil
	})
	client := kbclient.New(trp, strfmt.Default, authWriter, kbclient.KillbillDefaults{})

	// Set defaults. You can override them in each API call.
	createdBy := "John Doe"
	comment := "Created by John Doe"
	reason := ""

	client.SetDefaults(kbclient.KillbillDefaults{
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
