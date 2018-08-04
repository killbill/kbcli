package cmdlib

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"strings"

	"github.com/killbill/kbcli/kbcommon"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/kbclient"
	"github.com/urfave/cli"
)

// App provides interface for rest
type App struct {
	app *cli.App
	o   *Options
	ctx context.Context
}

var formatStr string

// NewApp creates new command registry
func NewApp() *App {
	r := &App{
		app: cli.NewApp(),
		o:   &Options{},
		ctx: context.Background(),
	}
	r.o.Log = NewLogger()
	r.o.out = os.Stdout
	r.o.FO = &FormatOptions{
		Type: FormatTypeDefault,
	}
	r.init()

	return r
}

// Run the program
func (r *App) Run(args []string) error {
	return r.app.Run(args)
}

// Register new command
func (r *App) Register(parentCmd string, command cli.Command, fn HandlerFn) {
	var parent *cli.Commands
	// Locate parent command
	if parentCmd != "" {
		root := r.app.Commands
		for _, cmdName := range strings.Split(parentCmd, ".") {
			for i := range root {
				c := &root[i]
				if c.Name == cmdName {
					if c.Subcommands == nil {
						c.Subcommands = cli.Commands{}
					}
					parent = &c.Subcommands
					root = c.Subcommands
					break
				}
			}
		}
		if parent == nil {
			panic(fmt.Sprintf("registry error: parent command '%s' not found", parentCmd))
		}
	} else {
		parent = (*cli.Commands)(&r.app.Commands)
	}
	if fn != nil {
		command.Action = r.toAction(fn)
	}
	*parent = append(*parent, command)
}

// init initializes the registry
func (r *App) init() {
	r.app.Name = "kbcmd"
	r.app.Usage = "Invoke kill bill commands"
	r.app.Version = "0.1"
	r.app.EnableBashCompletion = true
	r.app.Before = func(c *cli.Context) error {
		r.ctx = context.Background()
		r.o.FO.Type.Scan(formatStr)
		return nil
	}

	r.app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "host",
			Value:       "127.0.0.1:8080",
			Usage:       "Value to use in X-Killbill-Host",
			Destination: &r.o.Host,
			EnvVar:      "KB_HOST",
		},
		cli.StringFlag{
			Name:        "user",
			Value:       "admin",
			Usage:       "killbill username",
			Destination: &r.o.Username,
			EnvVar:      "KB_USER",
		},
		cli.StringFlag{
			Name:        "password",
			Value:       "password",
			Usage:       "killbill password",
			Destination: &r.o.Password,
			EnvVar:      "KB_PASSWORD",
		},
		cli.StringFlag{
			Name:        "created_by",
			Value:       os.Getenv("USER"),
			Usage:       "Value to use in X-Killbill-CreatedBy",
			Destination: &r.o.CreatedBy,
			EnvVar:      "KB_API_CREATED_BY",
		},
		cli.StringFlag{
			Name:        "api_key",
			Usage:       "Killbill API key X-Killbill-ApiKey",
			Destination: &r.o.APIKey,
			EnvVar:      "KB_API_KEY",
		},
		cli.StringFlag{
			Name:        "api_secret",
			Usage:       "Killbill API secret X-Killbill-ApiSecret",
			Destination: &r.o.APISecret,
			EnvVar:      "KB_API_SECRET",
		},
		cli.StringFlag{
			Name:  "format, f",
			Value: "default",
			Usage: `Output format. (One of table, short, default, list, json)

     table   - tabular format
     short   - short tabular format. child items are not printed.
	 default - use short for collections, and table for single item.
	 list    - list of key value pairs.
     json    - print json.
`,
			Destination: &formatStr,
		},
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "Print debug information",
			Destination: &r.o.PrintDebug,
			EnvVar:      "KB_DEBUG",
		},
		cli.BoolFlag{
			Name:        "no_header",
			Usage:       "Don't print header in csv/table format",
			Destination: &r.o.FO.NoHeader,
		},
	}
	r.app.Commands = []cli.Command{}
}

// toAction converts handler function to action handler to be usable by cli.
func (r *App) toAction(fn HandlerFn) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		o := *r.o
		o.Args = c.Args()

		trp := httptransport.New(o.Host, "", nil)
		// Add text/xml producer which is not handled by openapi runtime.
		trp.Producers["text/xml"] = runtime.TextProducer()
		trp.Debug = o.PrintDebug
		authWriter := runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
			encoded := base64.StdEncoding.EncodeToString([]byte(o.Username + ":" + o.Password))
			if err := r.SetHeaderParam("Authorization", "Basic "+encoded); err != nil {
				return err
			}
			if err := r.SetHeaderParam("X-KillBill-ApiKey", o.APIKey); err != nil {
				return err
			}
			if err := r.SetHeaderParam("X-KillBill-ApiSecret", o.APISecret); err != nil {
				return err
			}
			return nil
		})

		o.client = kbclient.New(trp, strfmt.Default, authWriter, kbclient.KillbillDefaults{})

		// Set defaults

		createdBy := os.Getenv("USER") + "-kbcmd"
		comment := "Created by kbcmd tool"
		reason := ""

		o.client.SetDefaults(kbclient.KillbillDefaults{
			CreatedBy:      &createdBy,
			Comment:        &comment,
			Reason:         &reason,
			WithStackTrace: &o.PrintDebug,
		})

		err := fn(r.ctx, &o)
		if err == nil {
			return err
		}
		if o.PrintDebug {
			if kberr, ok := err.(*kbcommon.KillbillError); ok {
				o.Outputln("%s", kberr.FormatStackTrace())
			}
		}
		return err
	}
}
