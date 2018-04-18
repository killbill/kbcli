package commands

import (
	"github.com/killbill/kbcli/kbcmd/cmdlib"
	"github.com/killbill/kbcli/kbcmd/commands/accounts"
	"github.com/killbill/kbcli/kbcmd/commands/subscriptions"
)

// RegisterAll registers all commands
func RegisterAll(r *cmdlib.App) {
	accounts.RegisterAccountCommands(r)
	subscriptions.RegisterCommands(r)

	registerInvoicesCommands(r)
	registerTenantCommands(r)
	registerCatalogCommands(r)
	registerBundleCommands(r)

	// Misc
	registerStripeCommands(r)
}
