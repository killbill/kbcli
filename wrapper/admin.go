package killbill

import (
	"context"
	"github.com/killbill/kbcli/v3/kbclient/admin"
)

func (cli *RawClient) InvalidateTenantCache(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	_, err := cli.TenantClient.Admin.InvalidatesCacheByTenant(ctx, &admin.InvalidatesCacheByTenantParams{})
	return err
}
