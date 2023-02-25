// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/v3/kbmodel"
)

// NewChargebackReversalPaymentByExternalKeyParams creates a new ChargebackReversalPaymentByExternalKeyParams object
// with the default values initialized.
func NewChargebackReversalPaymentByExternalKeyParams() *ChargebackReversalPaymentByExternalKeyParams {
	var ()
	return &ChargebackReversalPaymentByExternalKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewChargebackReversalPaymentByExternalKeyParamsWithTimeout creates a new ChargebackReversalPaymentByExternalKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewChargebackReversalPaymentByExternalKeyParamsWithTimeout(timeout time.Duration) *ChargebackReversalPaymentByExternalKeyParams {
	var ()
	return &ChargebackReversalPaymentByExternalKeyParams{

		timeout: timeout,
	}
}

// NewChargebackReversalPaymentByExternalKeyParamsWithContext creates a new ChargebackReversalPaymentByExternalKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewChargebackReversalPaymentByExternalKeyParamsWithContext(ctx context.Context) *ChargebackReversalPaymentByExternalKeyParams {
	var ()
	return &ChargebackReversalPaymentByExternalKeyParams{

		Context: ctx,
	}
}

// NewChargebackReversalPaymentByExternalKeyParamsWithHTTPClient creates a new ChargebackReversalPaymentByExternalKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewChargebackReversalPaymentByExternalKeyParamsWithHTTPClient(client *http.Client) *ChargebackReversalPaymentByExternalKeyParams {
	var ()
	return &ChargebackReversalPaymentByExternalKeyParams{
		HTTPClient: client,
	}
}

/*ChargebackReversalPaymentByExternalKeyParams contains all the parameters to send to the API endpoint
for the chargeback reversal payment by external key operation typically these are written to a http.Request
*/
type ChargebackReversalPaymentByExternalKeyParams struct {

	/*XKillbillComment*/
	XKillbillComment *string
	/*XKillbillCreatedBy*/
	XKillbillCreatedBy string
	/*XKillbillReason*/
	XKillbillReason *string
	/*Body*/
	Body *kbmodel.PaymentTransaction
	/*ControlPluginName*/
	ControlPluginName []string
	/*PluginProperty*/
	PluginProperty []string

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithTimeout adds the timeout to the chargeback reversal payment by external key params
func (o *ChargebackReversalPaymentByExternalKeyParams) WithTimeout(timeout time.Duration) *ChargebackReversalPaymentByExternalKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the chargeback reversal payment by external key params
func (o *ChargebackReversalPaymentByExternalKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the chargeback reversal payment by external key params
func (o *ChargebackReversalPaymentByExternalKeyParams) WithContext(ctx context.Context) *ChargebackReversalPaymentByExternalKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the chargeback reversal payment by external key params
func (o *ChargebackReversalPaymentByExternalKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the chargeback reversal payment by external key params
func (o *ChargebackReversalPaymentByExternalKeyParams) WithHTTPClient(client *http.Client) *ChargebackReversalPaymentByExternalKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the chargeback reversal payment by external key params
func (o *ChargebackReversalPaymentByExternalKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the chargeback reversal payment by external key params
func (o *ChargebackReversalPaymentByExternalKeyParams) WithXKillbillComment(xKillbillComment *string) *ChargebackReversalPaymentByExternalKeyParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the chargeback reversal payment by external key params
func (o *ChargebackReversalPaymentByExternalKeyParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the chargeback reversal payment by external key params
func (o *ChargebackReversalPaymentByExternalKeyParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *ChargebackReversalPaymentByExternalKeyParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the chargeback reversal payment by external key params
func (o *ChargebackReversalPaymentByExternalKeyParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the chargeback reversal payment by external key params
func (o *ChargebackReversalPaymentByExternalKeyParams) WithXKillbillReason(xKillbillReason *string) *ChargebackReversalPaymentByExternalKeyParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the chargeback reversal payment by external key params
func (o *ChargebackReversalPaymentByExternalKeyParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBody adds the body to the chargeback reversal payment by external key params
func (o *ChargebackReversalPaymentByExternalKeyParams) WithBody(body *kbmodel.PaymentTransaction) *ChargebackReversalPaymentByExternalKeyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the chargeback reversal payment by external key params
func (o *ChargebackReversalPaymentByExternalKeyParams) SetBody(body *kbmodel.PaymentTransaction) {
	o.Body = body
}

// WithControlPluginName adds the controlPluginName to the chargeback reversal payment by external key params
func (o *ChargebackReversalPaymentByExternalKeyParams) WithControlPluginName(controlPluginName []string) *ChargebackReversalPaymentByExternalKeyParams {
	o.SetControlPluginName(controlPluginName)
	return o
}

// SetControlPluginName adds the controlPluginName to the chargeback reversal payment by external key params
func (o *ChargebackReversalPaymentByExternalKeyParams) SetControlPluginName(controlPluginName []string) {
	o.ControlPluginName = controlPluginName
}

// WithPluginProperty adds the pluginProperty to the chargeback reversal payment by external key params
func (o *ChargebackReversalPaymentByExternalKeyParams) WithPluginProperty(pluginProperty []string) *ChargebackReversalPaymentByExternalKeyParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the chargeback reversal payment by external key params
func (o *ChargebackReversalPaymentByExternalKeyParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WriteToRequest writes these params to a swagger request
func (o *ChargebackReversalPaymentByExternalKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XKillbillComment != nil {

		// header param X-Killbill-Comment
		if err := r.SetHeaderParam("X-Killbill-Comment", *o.XKillbillComment); err != nil {
			return err
		}

	}

	// header param X-Killbill-CreatedBy
	if err := r.SetHeaderParam("X-Killbill-CreatedBy", o.XKillbillCreatedBy); err != nil {
		return err
	}

	if o.XKillbillReason != nil {

		// header param X-Killbill-Reason
		if err := r.SetHeaderParam("X-Killbill-Reason", *o.XKillbillReason); err != nil {
			return err
		}

	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	valuesControlPluginName := o.ControlPluginName

	joinedControlPluginName := swag.JoinByFormat(valuesControlPluginName, "multi")
	// query array param controlPluginName
	if err := r.SetQueryParam("controlPluginName", joinedControlPluginName...); err != nil {
		return err
	}

	valuesPluginProperty := o.PluginProperty

	joinedPluginProperty := swag.JoinByFormat(valuesPluginProperty, "multi")
	// query array param pluginProperty
	if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
		return err
	}

	// header param WithProfilingInfo
	if o.WithProfilingInfo != nil && len(*o.WithProfilingInfo) > 0 {
		if err := r.SetHeaderParam("X-Killbill-Profiling-Req", *o.WithProfilingInfo); err != nil {
			return err
		}
	}

	// header param withStackTrace
	if o.WithStackTrace != nil && *o.WithStackTrace {
		if err := r.SetQueryParam("withStackTrace", "true"); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
