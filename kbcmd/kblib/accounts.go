package kblib

import (
	"context"

	"github.com/killbill/kbcli/kbclient"

	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/kbclient/account"
	"github.com/killbill/kbcli/kbmodel"
)

// GetAccountByKeyOrID - get account information command
func GetAccountByKeyOrID(ctx context.Context, c *kbclient.KillBill, keyOrID string) (*kbmodel.Account, error) {
	keyOrID, isID := ParseKeyOrID(keyOrID)
	if isID {
		resp, err := c.Account.GetAccount(ctx, &account.GetAccountParams{
			AccountID: strfmt.UUID(keyOrID),
		})
		if err != nil {
			return nil, err
		}
		return resp.Payload, nil
	}
	resp, err := c.Account.GetAccountByKey(ctx, &account.GetAccountByKeyParams{
		ExternalKey: keyOrID,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
