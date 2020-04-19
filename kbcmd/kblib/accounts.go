package kblib

import (
	"context"

	"github.com/killbill/kbcli/v2/kbclient"

	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v2/kbclient/account"
	"github.com/killbill/kbcli/v2/kbmodel"
)

// GetAccountByKeyOrID - get account information command
func GetAccountByKeyOrID(ctx context.Context, c *kbclient.KillBill, keyOrID string) (*kbmodel.Account, error) {
	return GetAccountByKeyOrIDWithMaybeBalance(ctx, c, keyOrID, false)
}

func GetAccountByKeyOrIDWithBalanceAndCBA(ctx context.Context, c *kbclient.KillBill, keyOrID string) (*kbmodel.Account, error) {
	return GetAccountByKeyOrIDWithMaybeBalance(ctx, c, keyOrID, true)
}

func GetAccountByKeyOrIDWithMaybeBalance(ctx context.Context, c *kbclient.KillBill, keyOrID string, withBalanceAndCBA bool) (*kbmodel.Account, error) {
	keyOrID, isID := ParseKeyOrID(keyOrID)
	if isID {
		resp, err := c.Account.GetAccount(ctx, &account.GetAccountParams{
			AccountID:                strfmt.UUID(keyOrID),
			AccountWithBalanceAndCBA: &withBalanceAndCBA,
		})
		if err != nil {
			return nil, err
		}
		return resp.Payload, nil
	}
	resp, err := c.Account.GetAccountByKey(ctx, &account.GetAccountByKeyParams{
		ExternalKey:              keyOrID,
		AccountWithBalanceAndCBA: &withBalanceAndCBA,
	})
	if err != nil {
		return nil, err
	}

	return resp.Payload, nil
}
