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

	registerAuditLogCommands(r)
	registerBundleCommands(r)
	registerCatalogCommands(r)
	registerCustomFieldFunctions(r)
	registerInvoicesCommands(r)
	registerStripeCommands(r)
	registerTagDefinitionCommands(r)
	registerTenantCommands(r)

	// Dev
	registerDevCommands(r)
}
