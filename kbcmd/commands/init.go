package commands

import "github.com/killbill/kbcli/kbcmd/cmdlib"

// RegisterAll registers all commands
func RegisterAll(r *cmdlib.App) {
	registerAccountCommands(r)
	registerInvoicesCommands(r)
	registerTenantCommands(r)
	registerCatalogCommands(r)
	registerSubscriptionCommands(r)
	registerBundleCommands(r)
}
