package commands

import (
	"context"
	"fmt"
	"reflect"

	"github.com/killbill/kbcli/kbcmd/cmdlib/args"

	"github.com/killbill/kbcli/kbcmd/cmdlib"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/client"
	"github.com/urfave/cli"
)

var stripeSecret string
var createCardProperties = []args.Property{
	{Name: "Name", Required: true},
	{Name: "Number", Required: true},
	{Name: "Month", Required: true},
	{Name: "Year", Required: true},
	{Name: "CVC"},
	{Name: "Country"},
	{Name: "Currency"},
	{Name: "State"},
	{Name: "Zip"},
}

// createStripeCardToken - new stripe credit card token
func createStripeCardToken(ctx context.Context, o *cmdlib.Options) error {
	sc := &client.API{}
	if stripeSecret == "" {
		return fmt.Errorf("stripe_key flag must be specified")
	}
	stripe.LogLevel = 1
	if o.PrintDebug {
		stripe.LogLevel = 3
	}

	sc.Init(stripeSecret, nil)
	params := &stripe.TokenParams{
		Card: &stripe.CardParams{},
	}

	err := args.LoadProperties(params.Card, createCardProperties, o.Args)
	if err != nil {
		return err
	}

	t, err := sc.Tokens.New(params)
	if err != nil {
		return err
	}

	o.Print(t)
	return nil
}

func registerStripeCommands(r *cmdlib.App) {
	cmdlib.AddFormatter(reflect.TypeOf(&stripe.Token{}), cmdlib.Formatter{
		Columns: []cmdlib.Column{
			{
				Name: "TOKEN",
				Path: "$.id",
			},
			{
				Name: "Name",
				Path: "$.card.name",
			},
			{
				Name: "MONTH",
				Path: "$.card.exp_month",
			},
			{
				Name: "YEAR",
				Path: "$.card.exp_year",
			},
		},
	})
	// Register top level command
	r.Register("", cli.Command{
		Name:  "stripe",
		Usage: "Stripe utility methods",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "stripe_key",
				Usage:       "Stripe private key",
				Destination: &stripeSecret,
				EnvVar:      "KB_STRIPE_KEY",
			},
		},
	}, nil)

	createTokenUsage := args.GenerateUsageString(&stripe.CardParams{}, createCardProperties)
	r.Register("stripe", cli.Command{
		Name:  "new-card-token",
		Usage: "create new stripe credit card token",
		ArgsUsage: fmt.Sprintf(`%s

			For e.g.,
				stripe --stripe_key=sk_test_aT1UY1u211iUYI4LEL2tLFay new-card-token Name="Prem Ramanathan" Number=4242424242424242 Month=12 Year=2019 CVC=123
		`, createTokenUsage),
		Description: "Creates new stripe card token",
	}, createStripeCardToken)
}
