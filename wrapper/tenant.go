package killbill

import (
	"context"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/killbill/kbcli/v2/kbclient/tenant"
	"github.com/killbill/kbcli/v2/kbmodel"
)

func (cli *RawClient) CreateTenant(ctx context.Context, apiKey string, apiSecret string, extKey string) (*strfmt.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	res, err := cli.CrossTenantClient.Tenant.CreateTenant(ctx, &tenant.CreateTenantParams{
		Body: &kbmodel.Tenant{
			APIKey:      swag.String(apiKey),
			APISecret:   swag.String(apiSecret),
			ExternalKey: extKey,
		},
		// TODO Should it be set to false instead?
		UseGlobalDefault: swag.Bool(true),
		// Setting ProcessLocationHeader does not work so we cannot get out tenant back
		ProcessLocationHeader: false,
	})
	if err != nil {
		return nil, err
	}

	// This returns an empty ID
	return &res.Payload.TenantID, nil
}

func (cli *RawClient) GetTenant(ctx context.Context, apiKey string) (*strfmt.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	res, err := cli.CrossTenantClient.Tenant.GetTenantByAPIKey(ctx, &tenant.GetTenantByAPIKeyParams{
		APIKey:                swag.String(apiKey),
		ProcessLocationHeader: false,
	})
	if err != nil {
		return nil, err
	}
	return &res.Payload.TenantID, nil
}
