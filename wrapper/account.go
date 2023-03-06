package killbill

import (
	"context"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v2/kbclient/account"
	"github.com/killbill/kbcli/v2/kbmodel"
	"time"
)

func (cli *RawClient) CreateAccountFromBody(ctx context.Context, body *kbmodel.Account) (strfmt.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	res, err := cli.TenantClient.Account.CreateAccount(ctx, &account.CreateAccountParams{
		Body:                  body,
		ProcessLocationHeader: true,
	})
	if err != nil {
		return "", err
	}
	return res.Payload.AccountID, nil
}

func (cli *RawClient) CreateAccount(ctx context.Context, extKey string, currency string, locale string, bcd int32) (strfmt.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	body := &kbmodel.Account{
		ExternalKey:   extKey,
		Currency:      kbmodel.AccountCurrencyEnum(currency),
		Locale:        locale,
		ReferenceTime: strfmt.DateTime(time.Now()),
	}

	if bcd > 0 {
		body.BillCycleDayLocal = bcd
	}

	return cli.CreateAccountFromBody(ctx, body)
}

func (cli *RawClient) GetAccount(ctx context.Context, accountId strfmt.UUID, withBalance bool) (*kbmodel.Account, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	res, err := cli.TenantClient.Account.GetAccount(ctx, &account.GetAccountParams{
		AccountID:                accountId,
		AccountWithBalanceAndCBA: &withBalance,
	})
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (cli *RawClient) GetAccountByKey(ctx context.Context, externalKey string) (*kbmodel.Account, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	res, err := cli.TenantClient.Account.GetAccountByKey(ctx, &account.GetAccountByKeyParams{
		ExternalKey: externalKey,
	})
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (cli *RawClient) AddAccountCustomFields(ctx context.Context, accountId strfmt.UUID, fields map[string]string) error {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	body := make([]*kbmodel.CustomField, len(fields))
	idx := 0
	for name, value := range fields {
		body[idx] = &kbmodel.CustomField{
			Name:       &name,
			ObjectID:   accountId,
			ObjectType: kbmodel.CustomFieldObjectTypeACCOUNT,
			Value:      &value,
		}
		idx++
	}

	_, err := cli.TenantClient.Account.CreateAccountCustomFields(ctx, &account.CreateAccountCustomFieldsParams{
		Body:                  body,
		AccountID:             accountId,
		ProcessLocationHeader: true,
	})
	return err
}

func (cli *RawClient) GetAccountTags(ctx context.Context, accountId strfmt.UUID) ([]*kbmodel.Tag, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	res, err := cli.TenantClient.Account.GetAccountTags(ctx, &account.GetAccountTagsParams{
		AccountID: accountId,
	})
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (cli *RawClient) AddAccountTags(ctx context.Context, accountId strfmt.UUID, tags ...string) error {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	fmtTags := make([]strfmt.UUID, len(tags))
	for i := 0; i < len(tags); i++ {
		fmtTags[i] = strfmt.UUID(tags[i])
	}

	_, err := cli.TenantClient.Account.CreateAccountTags(ctx, &account.CreateAccountTagsParams{
		AccountID: accountId,
		Body:      fmtTags,
	})
	return err
}

func (cli *RawClient) DeleteAccountTags(ctx context.Context, accountId strfmt.UUID, tags ...string) error {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	fmtTags := make([]strfmt.UUID, len(tags))
	for i := 0; i < len(tags); i++ {
		fmtTags[i] = strfmt.UUID(tags[i])
	}

	_, err := cli.TenantClient.Account.DeleteAccountTags(ctx, &account.DeleteAccountTagsParams{
		AccountID: accountId,
		TagDef:    fmtTags,
	})
	return err
}

// GetInvoicesBetweenDates returns all invoices belonging to an
// account, whose invoice date is in between the given arguments.
//
// In case of connection errors, the underlying error is returned.
// For other cases (HTTP status >= 400), a kbcommon.KillbillError is returned.
func (cli *RawClient) GetInvoicesBetweenDates(ctx context.Context, accountID strfmt.UUID, from *strfmt.Date, to *strfmt.Date) ([]*kbmodel.Invoice, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()
	b := true

	res, err := cli.TenantClient.Account.GetInvoicesForAccount(ctx, &account.GetInvoicesForAccountParams{
		AccountID:             accountID,
		StartDate:             from,
		EndDate:               to,
		IncludeVoidedInvoices: &b,
	})
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

// GetMonthlyInvoice returns the invoice matching the given target
//
// In case of connection errors, the underlying error is returned.
// For other cases (HTTP status >= 400), a kbcommon.KillbillError is returned.
func (cli *RawClient) GetMonthlyInvoice(ctx context.Context, accountID strfmt.UUID, targetDt strfmt.Date) (*kbmodel.Invoice, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	var from, to strfmt.Date

	res, err := cli.TenantClient.Account.GetInvoicesForAccount(ctx, &account.GetInvoicesForAccountParams{
		AccountID: accountID,
		StartDate: &from,
		EndDate:   &to,
	})
	if err != nil {
		return nil, err
	}
	// Kill Bill could return a successful status but without any invoices matching the query,
	// so we transform this as an error to make it easier on the consumer
	if len(res.Payload) == 0 {
		return nil, &account.GetInvoicesForAccountNotFound{
			HttpResponse: res.HttpResponse,
		}
	}
	return res.Payload[0], nil
}

func (cli *RawClient) addPaymentMethod(ctx context.Context, accountId strfmt.UUID, isDefault bool) error {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	_, err := cli.TenantClient.Account.CreatePaymentMethod(ctx, &account.CreatePaymentMethodParams{
		AccountID: accountId,
		Body: &kbmodel.PaymentMethod{
			AccountID:  accountId,
			IsDefault:  isDefault,
			PluginName: "__EXTERNAL_PAYMENT__",
		},
		IsDefault: &isDefault,
	})
	return err
}
