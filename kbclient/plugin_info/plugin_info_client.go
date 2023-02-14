// Code generated by go-swagger; DO NOT EDIT.

package plugin_info

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new plugin info API client.
func New(transport runtime.ClientTransport,
	formats strfmt.Registry,
	authInfo runtime.ClientAuthInfoWriter,
	defaults KillbillDefaults) *Client {

	return &Client{transport: transport, formats: formats, authInfo: authInfo, defaults: defaults}
}

// killbill default values. When a call is made to an operation, these values are used
// if params doesn't specify them.
type KillbillDefaults interface {
	// Default CreatedBy. If not set explicitly in params, this will be used.
	XKillbillCreatedBy() *string
	// Default Comment. If not set explicitly in params, this will be used.
	XKillbillComment() *string
	// Default Reason. If not set explicitly in params, this will be used.
	XKillbillReason() *string
	// Default WithWithProfilingInfo. If not set explicitly in params, this will be used.
	KillbillWithProfilingInfo() *string
	// Default WithStackTrace. If not set explicitly in params, this will be used.
	KillbillWithStackTrace() *bool
}

/*
Client for plugin info API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
	defaults  KillbillDefaults
}

// IPluginInfo - interface for PluginInfo client.
type IPluginInfo interface {
	/*
		GetPluginsInfo retrieves the list of registered plugins
	*/
	GetPluginsInfo(ctx context.Context, params *GetPluginsInfoParams) (*GetPluginsInfoOK, error)
}

/*
GetPluginsInfo retrieves the list of registered plugins
*/
func (a *Client) GetPluginsInfo(ctx context.Context, params *GetPluginsInfoParams) (*GetPluginsInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPluginsInfoParams()
	}
	params.Context = ctx
	if params.WithProfilingInfo == nil && a.defaults.KillbillWithProfilingInfo() != nil {
		params.WithProfilingInfo = a.defaults.KillbillWithProfilingInfo()
	}

	if params.WithStackTrace == nil && a.defaults.KillbillWithStackTrace() != nil {
		params.WithStackTrace = a.defaults.KillbillWithStackTrace()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPluginsInfo",
		Method:             "GET",
		PathPattern:        "/1.0/kb/pluginsInfo",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPluginsInfoReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetPluginsInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getPluginsInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
