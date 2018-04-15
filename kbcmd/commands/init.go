package commands

import (
	"github.com/killbill/kbcli/kbcmd/cmdlib"
	"github.com/killbill/kbcli/kbcmd/commands/accounts"
)

// RegisterAll registers all commands
func RegisterAll(r *cmdlib.App) {
	accounts.RegisterAccountCommands(r)

	registerInvoicesCommands(r)
	registerTenantCommands(r)
	registerCatalogCommands(r)
	registerSubscriptionCommands(r)
	registerBundleCommands(r)

	// Misc
	registerStripeCommands(r)
}
