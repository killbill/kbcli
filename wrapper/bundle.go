package killbill

import (
	"context"
	"fmt"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v2/kbclient/bundle"
	"github.com/killbill/kbcli/v2/kbmodel"
	"strings"
)

func (cli *RawClient) GetBundle(ctx context.Context, bundleId strfmt.UUID) (*kbmodel.Bundle, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	res, err := cli.TenantClient.Bundle.GetBundle(ctx, &bundle.GetBundleParams{
		BundleID: bundleId,
	})
	if err != nil {
		return nil, err
	}
	return res.Payload, nil
}

func (cli *RawClient) GetBundleByKey(ctx context.Context, extKey string) (*kbmodel.Bundle, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	res, err := cli.TenantClient.Bundle.GetBundleByKey(ctx, &bundle.GetBundleByKeyParams{
		ExternalKey: extKey,
	})
	if err != nil {
		return nil, err
	}
	if len(res.Payload) > 1 {
		return nil, fmt.Errorf("Unexpected number of bundles for extKey=%s: Got %d bundles", extKey, len(res.Payload))
	}
	return res.Payload[0], nil
}

func (cli *RawClient) TransferBundle(ctx context.Context, bundleId strfmt.UUID, destAccountId strfmt.UUID, requestedDate *strfmt.Date, subExtkeys map[string]string, follow bool) (strfmt.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	// Serialize subscription external keys, rely on undocumented mechanism until https://github.com/killbill/killbill/issues/1627 is fixed.
	props := make([]string, 0)
	for k, v := range subExtkeys {
		prop := fmt.Sprintf("KB_SUB_ID_%s=%s", k, v)
		props = append(props, prop)
	}

	billingPolicy := "IMMEDIATE"
	res, err := cli.TenantClient.Bundle.TransferBundle(ctx, &bundle.TransferBundleParams{
		BillingPolicy: &billingPolicy,
		Body: &kbmodel.Bundle{
			AccountID: &destAccountId,
		},
		PluginProperty:        props,
		BundleID:              bundleId,
		RequestedDate:         requestedDate,
		ProcessLocationHeader: follow,
	})
	if err != nil {
		return "", err
	}
	var result strfmt.UUID
	if follow {
		result = res.Payload.BundleID
	} else {
		location := res.HttpResponse.GetHeader("Location")
		parts := strings.Split(location, "/")
		result = strfmt.UUID(parts[len(parts)-1])
	}
	return result, nil
}
