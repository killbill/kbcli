package commands

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/killbill/kbcli/kbclient/catalog"
	"github.com/killbill/kbcli/kbcmd/cmdlib"
	"github.com/urfave/cli"
)

func getCatalog(ctx context.Context, o *cmdlib.Options) error {
	if len(o.Args) > 1 {
		return cmdlib.ErrorInvalidArgs
	}
	var outputFile string

	if len(o.Args) == 1 {
		outputFile = o.Args[0]
		_, err := os.Stat(outputFile)
		if err == nil || !os.IsNotExist(err) {
			return fmt.Errorf("output file already exists %v", outputFile)
		}
	}

	res, err := o.Client().Catalog.GetCatalogXML(ctx, &catalog.GetCatalogXMLParams{})
	if err != nil {
		return err
	}

	if outputFile != "" {
		if err = ioutil.WriteFile(outputFile, []byte(res.Payload), 0644); err != nil {
			return fmt.Errorf("unable to write output: %v", err)
		}
		o.Outputln("catalog successfully written to %v", outputFile)
	} else {
		o.Outputln("%s", res.Payload)
	}
	return nil
}

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

	_, err = o.Client().Catalog.UploadCatalogXML(ctx, &catalog.UploadCatalogXMLParams{
		Body: &strContents,
	})
	if err != nil {
		return err
	}
	o.Outputln("catalog successfully uploaded")
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
		Name:  "get",
		Usage: "get catalog for tenant",
		ArgsUsage: `[OUTPUT_FILE]
		
		For ex.,
		
		# To write the catalog to a file
		kbcmd catalog get /tmp/catalog.json
		
		# To write to stdout
		kbcmd catalog get
`,
		Description: "Retrieves the catalog and stores in the file",
	}, getCatalog)

	r.Register("catalog", cli.Command{
		Name:  "upload",
		Usage: "upload given catalog",
		ArgsUsage: `CATALOG_FILE
		
For ex.,
		kbcmd catalog upload /tmp/catalog.json
`,
		Description: "Uploads the catalog",
	}, uploadCatalog)
}
