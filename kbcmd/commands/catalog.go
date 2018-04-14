package commands

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/killbill/kbcli/kbclient/catalog"
	"github.com/killbill/kbcli/kbcmd/cmdlib"
	"github.com/urfave/cli"
)

func uploadCatalog(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) != 1 {
		return cmdlib.ErrorInvalidArgs
	}
	catalogFile := o.Args[0]
	contents, err := ioutil.ReadFile(catalogFile)
	if err != nil {
		return fmt.Errorf("unable to read catalog file %s. %v", catalogFile, err)
	}
	strContents := string(contents)

	resp, err := o.Client().Catalog.UploadCatalogXML(ctx, &catalog.UploadCatalogXMLParams{
		Body: &strContents,
	})
	o.Output("%s\n", resp.Payload)
	return err
}

func registerCatalogCommands(r *cmdlib.App) {
	// Register top level command
	r.Register("", cli.Command{
		Name:    "catalog",
		Aliases: []string{"cat"},
		Usage:   "catalog related commands",
	}, nil)

	r.Register("catalog", cli.Command{
		Name:        "upload",
		Usage:       "upload given catalog",
		ArgsUsage:   "CATALOG_FILE",
		Description: "Uploads the catalog",
	}, uploadCatalog)
}
