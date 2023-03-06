package killbill

import (
	"context"
	"github.com/killbill/kbcli/v2/kbclient/catalog"
	"os"
)

func (cli *RawClient) UploadCatalogXML(ctx context.Context, path string) error {
	ctx, cancel := context.WithTimeout(ctx, cli.Timeout)
	defer cancel()

	buf, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	_, err = cli.TenantClient.Catalog.UploadCatalogXML(ctx, &catalog.UploadCatalogXMLParams{
		Body:                  string(buf),
		ProcessLocationHeader: false,
	})
	if err != nil {
		return err
	}
	return nil
}
