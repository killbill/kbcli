// Code generated by go-swagger; DO NOT EDIT.

package kbclient

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/killbill/kbcli/v2/kbclient/account"
	"github.com/killbill/kbcli/v2/kbclient/admin"
	"github.com/killbill/kbcli/v2/kbclient/bundle"
	"github.com/killbill/kbcli/v2/kbclient/catalog"
	"github.com/killbill/kbcli/v2/kbclient/credit"
	"github.com/killbill/kbcli/v2/kbclient/custom_field"
	"github.com/killbill/kbcli/v2/kbclient/export"
	"github.com/killbill/kbcli/v2/kbclient/invoice"
	"github.com/killbill/kbcli/v2/kbclient/invoice_item"
	"github.com/killbill/kbcli/v2/kbclient/invoice_payment"
	"github.com/killbill/kbcli/v2/kbclient/nodes_info"
	"github.com/killbill/kbcli/v2/kbclient/overdue"
	"github.com/killbill/kbcli/v2/kbclient/payment"
	"github.com/killbill/kbcli/v2/kbclient/payment_gateway"
	"github.com/killbill/kbcli/v2/kbclient/payment_method"
	"github.com/killbill/kbcli/v2/kbclient/payment_transaction"
	"github.com/killbill/kbcli/v2/kbclient/plugin_info"
	"github.com/killbill/kbcli/v2/kbclient/security"
	securityops "github.com/killbill/kbcli/v2/kbclient/security"
	"github.com/killbill/kbcli/v2/kbclient/subscription"
	"github.com/killbill/kbcli/v2/kbclient/tag"
	"github.com/killbill/kbcli/v2/kbclient/tag_definition"
	"github.com/killbill/kbcli/v2/kbclient/tenant"
	"github.com/killbill/kbcli/v2/kbclient/usage"
)

// New creates a new kill bill client
// The following snippet provides creating killbill client with basic auth.
//
//		   trp := httptransport.New("127.0.0.1:8080" /*host*/, "" /*basePath*/, nil /*schemes*/)
//		   // Add missing handler. OpenAPI runtime doesn't have this by default
//		   trp.Producers["text/xml"] = runtime.TextProducer()
//		   // Set tro true to print http/debug logs
//		   trp.Debug = enableDebug
//		   // Setup basic auth
//		   authWriter := httptransport.BasicAuth("admin", "password")
//		   client := kbclient.New(trp, strfmt.Default)
//	    // Use client
//	    client.Accounts.GetAccount(...)
func New(transport runtime.ClientTransport,
	formats strfmt.Registry,
	authInfo runtime.ClientAuthInfoWriter,
	defaults KillbillDefaults) *KillBill {

	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := &KillBill{
		Transport: transport,
		defaults:  defaults,
	}

	cli.Account = account.New(transport, formats, authInfo, &cli.defaults)

	cli.Admin = admin.New(transport, formats, authInfo, &cli.defaults)

	cli.Bundle = bundle.New(transport, formats, authInfo, &cli.defaults)

	cli.Catalog = catalog.New(transport, formats, authInfo, &cli.defaults)

	cli.Credit = credit.New(transport, formats, authInfo, &cli.defaults)

	cli.CustomField = custom_field.New(transport, formats, authInfo, &cli.defaults)

	cli.Export = export.New(transport, formats, authInfo, &cli.defaults)

	cli.Invoice = invoice.New(transport, formats, authInfo, &cli.defaults)

	cli.InvoiceItem = invoice_item.New(transport, formats, authInfo, &cli.defaults)

	cli.InvoicePayment = invoice_payment.New(transport, formats, authInfo, &cli.defaults)

	cli.NodesInfo = nodes_info.New(transport, formats, authInfo, &cli.defaults)

	cli.Overdue = overdue.New(transport, formats, authInfo, &cli.defaults)

	cli.Payment = payment.New(transport, formats, authInfo, &cli.defaults)

	cli.PaymentGateway = payment_gateway.New(transport, formats, authInfo, &cli.defaults)

	cli.PaymentMethod = payment_method.New(transport, formats, authInfo, &cli.defaults)

	cli.PaymentTransaction = payment_transaction.New(transport, formats, authInfo, &cli.defaults)

	cli.PluginInfo = plugin_info.New(transport, formats, authInfo, &cli.defaults)

	cli.Security = security.New(transport, formats, authInfo, &cli.defaults)

	cli.Subscription = subscription.New(transport, formats, authInfo, &cli.defaults)

	cli.Tag = tag.New(transport, formats, authInfo, &cli.defaults)

	cli.TagDefinition = tag_definition.New(transport, formats, authInfo, &cli.defaults)

	cli.Tenant = tenant.New(transport, formats, authInfo, &cli.defaults)

	cli.Usage = usage.New(transport, formats, authInfo, &cli.defaults)

	return cli
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// KillBill is a client for kill bill
type KillBill struct {
	Account account.ClientService

	Admin admin.ClientService

	Bundle bundle.ClientService

	Catalog catalog.ClientService

	Credit credit.ClientService

	CustomField custom_field.ClientService

	Export export.ClientService

	Invoice invoice.ClientService

	InvoiceItem invoice_item.ClientService

	InvoicePayment invoice_payment.ClientService

	NodesInfo nodes_info.ClientService

	Overdue overdue.ClientService

	Payment payment.ClientService

	PaymentGateway payment_gateway.ClientService

	PaymentMethod payment_method.ClientService

	PaymentTransaction payment_transaction.ClientService

	PluginInfo plugin_info.ClientService

	Security securityops.ClientService

	Subscription subscription.ClientService

	Tag tag.ClientService

	TagDefinition tag_definition.ClientService

	Tenant tenant.ClientService

	Usage usage.ClientService

	Transport runtime.ClientTransport
	defaults  KillbillDefaults
}

// SetTransport changes the transport on the client and all its subresources
func (c *KillBill) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Account.SetTransport(transport)
	c.Admin.SetTransport(transport)
	c.Bundle.SetTransport(transport)
	c.Catalog.SetTransport(transport)
	c.Credit.SetTransport(transport)
	c.CustomField.SetTransport(transport)
	c.Export.SetTransport(transport)
	c.Invoice.SetTransport(transport)
	c.InvoiceItem.SetTransport(transport)
	c.InvoicePayment.SetTransport(transport)
	c.NodesInfo.SetTransport(transport)
	c.Overdue.SetTransport(transport)
	c.Payment.SetTransport(transport)
	c.PaymentGateway.SetTransport(transport)
	c.PaymentMethod.SetTransport(transport)
	c.PaymentTransaction.SetTransport(transport)
	c.PluginInfo.SetTransport(transport)
	c.Security.SetTransport(transport)
	c.Subscription.SetTransport(transport)
	c.Tag.SetTransport(transport)
	c.TagDefinition.SetTransport(transport)
	c.Tenant.SetTransport(transport)
	c.Usage.SetTransport(transport)
}

// Defaults returns killbill defaults
func (c *KillBill) Defaults() KillbillDefaults {
	return c.defaults
}

// SetDefaults sets killbill defaults
func (c *KillBill) SetDefaults(defaults KillbillDefaults) {
	c.defaults = defaults
}

// Implements killbill default values.
type KillbillDefaults struct {
	// XKillbillCreatedBy property
	CreatedBy *string
	// XKillbillComment property
	Comment *string
	// XKillbillReason property
	Reason *string
	// withProfilingInfo property
	WithProfilingInfo *string
	// withStackTrace property
	WithStackTrace *bool
}

// Default CreatedBy. If not set explicitly in params, this will be used.
func (d KillbillDefaults) XKillbillCreatedBy() *string {
	return d.CreatedBy
}

// Default Comment. If not set explicitly in params, this will be used.
func (d KillbillDefaults) XKillbillComment() *string {
	return d.Comment
}

// Default Reason. If not set explicitly in params, this will be used.
func (d KillbillDefaults) XKillbillReason() *string {
	return d.Reason
}

// Default WithProfilingInfo. If not set explicitly in params, this will be used.
func (d KillbillDefaults) KillbillWithProfilingInfo() *string {
	return d.WithProfilingInfo
}

// Default WithStackTrace. If not set explicitly in params, this will be used.
func (d KillbillDefaults) KillbillWithStackTrace() *bool {
	return d.WithStackTrace
}
