// Code generated by go-swagger; DO NOT EDIT.

package subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewCancelSubscriptionPlanParams creates a new CancelSubscriptionPlanParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCancelSubscriptionPlanParams() *CancelSubscriptionPlanParams {
	return &CancelSubscriptionPlanParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCancelSubscriptionPlanParamsWithTimeout creates a new CancelSubscriptionPlanParams object
// with the ability to set a timeout on a request.
func NewCancelSubscriptionPlanParamsWithTimeout(timeout time.Duration) *CancelSubscriptionPlanParams {
	return &CancelSubscriptionPlanParams{
		timeout: timeout,
	}
}

// NewCancelSubscriptionPlanParamsWithContext creates a new CancelSubscriptionPlanParams object
// with the ability to set a context for a request.
func NewCancelSubscriptionPlanParamsWithContext(ctx context.Context) *CancelSubscriptionPlanParams {
	return &CancelSubscriptionPlanParams{
		Context: ctx,
	}
}

// NewCancelSubscriptionPlanParamsWithHTTPClient creates a new CancelSubscriptionPlanParams object
// with the ability to set a custom HTTPClient for a request.
func NewCancelSubscriptionPlanParamsWithHTTPClient(client *http.Client) *CancelSubscriptionPlanParams {
	return &CancelSubscriptionPlanParams{
		HTTPClient: client,
	}
}

/*
CancelSubscriptionPlanParams contains all the parameters to send to the API endpoint

	for the cancel subscription plan operation.

	Typically these are written to a http.Request.
*/
type CancelSubscriptionPlanParams struct {

	// XKillbillComment.
	XKillbillComment *string

	// XKillbillCreatedBy.
	XKillbillCreatedBy string

	// XKillbillReason.
	XKillbillReason *string

	// BillingPolicy.
	BillingPolicy *string

	// CallCompletion.
	CallCompletion *bool

	// CallTimeoutSec.
	//
	// Format: int64
	// Default: 5
	CallTimeoutSec *int64

	// EntitlementPolicy.
	EntitlementPolicy *string

	// PluginProperty.
	PluginProperty []string

	// RequestedDate.
	//
	// Format: date
	RequestedDate *strfmt.Date

	// SubscriptionID.
	//
	// Format: uuid
	SubscriptionID strfmt.UUID

	// UseRequestedDateForBilling.
	UseRequestedDateForBilling *bool

	WithProfilingInfo     *string // If set, return KB hprof headers
	WithStackTrace        *bool   // If set, returns full stack trace with error message
	timeout               time.Duration
	Context               context.Context
	HTTPClient            *http.Client
	ProcessLocationHeader bool // For create APIs that return 201, send another request and retrieve the resource.
}

// WithDefaults hydrates default values in the cancel subscription plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelSubscriptionPlanParams) WithDefaults() *CancelSubscriptionPlanParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cancel subscription plan params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CancelSubscriptionPlanParams) SetDefaults() {
	var (
		callCompletionDefault = bool(false)

		callTimeoutSecDefault = int64(5)

		useRequestedDateForBillingDefault = bool(false)
	)

	val := CancelSubscriptionPlanParams{
		CallCompletion:             &callCompletionDefault,
		CallTimeoutSec:             &callTimeoutSecDefault,
		UseRequestedDateForBilling: &useRequestedDateForBillingDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) WithTimeout(timeout time.Duration) *CancelSubscriptionPlanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) WithContext(ctx context.Context) *CancelSubscriptionPlanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) WithHTTPClient(client *http.Client) *CancelSubscriptionPlanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXKillbillComment adds the xKillbillComment to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) WithXKillbillComment(xKillbillComment *string) *CancelSubscriptionPlanParams {
	o.SetXKillbillComment(xKillbillComment)
	return o
}

// SetXKillbillComment adds the xKillbillComment to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) SetXKillbillComment(xKillbillComment *string) {
	o.XKillbillComment = xKillbillComment
}

// WithXKillbillCreatedBy adds the xKillbillCreatedBy to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) WithXKillbillCreatedBy(xKillbillCreatedBy string) *CancelSubscriptionPlanParams {
	o.SetXKillbillCreatedBy(xKillbillCreatedBy)
	return o
}

// SetXKillbillCreatedBy adds the xKillbillCreatedBy to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) SetXKillbillCreatedBy(xKillbillCreatedBy string) {
	o.XKillbillCreatedBy = xKillbillCreatedBy
}

// WithXKillbillReason adds the xKillbillReason to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) WithXKillbillReason(xKillbillReason *string) *CancelSubscriptionPlanParams {
	o.SetXKillbillReason(xKillbillReason)
	return o
}

// SetXKillbillReason adds the xKillbillReason to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) SetXKillbillReason(xKillbillReason *string) {
	o.XKillbillReason = xKillbillReason
}

// WithBillingPolicy adds the billingPolicy to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) WithBillingPolicy(billingPolicy *string) *CancelSubscriptionPlanParams {
	o.SetBillingPolicy(billingPolicy)
	return o
}

// SetBillingPolicy adds the billingPolicy to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) SetBillingPolicy(billingPolicy *string) {
	o.BillingPolicy = billingPolicy
}

// WithCallCompletion adds the callCompletion to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) WithCallCompletion(callCompletion *bool) *CancelSubscriptionPlanParams {
	o.SetCallCompletion(callCompletion)
	return o
}

// SetCallCompletion adds the callCompletion to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) SetCallCompletion(callCompletion *bool) {
	o.CallCompletion = callCompletion
}

// WithCallTimeoutSec adds the callTimeoutSec to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) WithCallTimeoutSec(callTimeoutSec *int64) *CancelSubscriptionPlanParams {
	o.SetCallTimeoutSec(callTimeoutSec)
	return o
}

// SetCallTimeoutSec adds the callTimeoutSec to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) SetCallTimeoutSec(callTimeoutSec *int64) {
	o.CallTimeoutSec = callTimeoutSec
}

// WithEntitlementPolicy adds the entitlementPolicy to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) WithEntitlementPolicy(entitlementPolicy *string) *CancelSubscriptionPlanParams {
	o.SetEntitlementPolicy(entitlementPolicy)
	return o
}

// SetEntitlementPolicy adds the entitlementPolicy to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) SetEntitlementPolicy(entitlementPolicy *string) {
	o.EntitlementPolicy = entitlementPolicy
}

// WithPluginProperty adds the pluginProperty to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) WithPluginProperty(pluginProperty []string) *CancelSubscriptionPlanParams {
	o.SetPluginProperty(pluginProperty)
	return o
}

// SetPluginProperty adds the pluginProperty to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) SetPluginProperty(pluginProperty []string) {
	o.PluginProperty = pluginProperty
}

// WithRequestedDate adds the requestedDate to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) WithRequestedDate(requestedDate *strfmt.Date) *CancelSubscriptionPlanParams {
	o.SetRequestedDate(requestedDate)
	return o
}

// SetRequestedDate adds the requestedDate to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) SetRequestedDate(requestedDate *strfmt.Date) {
	o.RequestedDate = requestedDate
}

// WithSubscriptionID adds the subscriptionID to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) WithSubscriptionID(subscriptionID strfmt.UUID) *CancelSubscriptionPlanParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) SetSubscriptionID(subscriptionID strfmt.UUID) {
	o.SubscriptionID = subscriptionID
}

// WithUseRequestedDateForBilling adds the useRequestedDateForBilling to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) WithUseRequestedDateForBilling(useRequestedDateForBilling *bool) *CancelSubscriptionPlanParams {
	o.SetUseRequestedDateForBilling(useRequestedDateForBilling)
	return o
}

// SetUseRequestedDateForBilling adds the useRequestedDateForBilling to the cancel subscription plan params
func (o *CancelSubscriptionPlanParams) SetUseRequestedDateForBilling(useRequestedDateForBilling *bool) {
	o.UseRequestedDateForBilling = useRequestedDateForBilling
}

// WriteToRequest writes these params to a swagger request
func (o *CancelSubscriptionPlanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.BillingPolicy != nil {

		// query param billingPolicy
		var qrBillingPolicy string

		if o.BillingPolicy != nil {
			qrBillingPolicy = *o.BillingPolicy
		}
		qBillingPolicy := qrBillingPolicy
		if qBillingPolicy != "" {

			if err := r.SetQueryParam("billingPolicy", qBillingPolicy); err != nil {
				return err
			}
		}
	}

	if o.CallCompletion != nil {

		// query param callCompletion
		var qrCallCompletion bool

		if o.CallCompletion != nil {
			qrCallCompletion = *o.CallCompletion
		}
		qCallCompletion := swag.FormatBool(qrCallCompletion)
		if qCallCompletion != "" {

			if err := r.SetQueryParam("callCompletion", qCallCompletion); err != nil {
				return err
			}
		}
	}

	if o.CallTimeoutSec != nil {

		// query param callTimeoutSec
		var qrCallTimeoutSec int64

		if o.CallTimeoutSec != nil {
			qrCallTimeoutSec = *o.CallTimeoutSec
		}
		qCallTimeoutSec := swag.FormatInt64(qrCallTimeoutSec)
		if qCallTimeoutSec != "" {

			if err := r.SetQueryParam("callTimeoutSec", qCallTimeoutSec); err != nil {
				return err
			}
		}
	}

	if o.EntitlementPolicy != nil {

		// query param entitlementPolicy
		var qrEntitlementPolicy string

		if o.EntitlementPolicy != nil {
			qrEntitlementPolicy = *o.EntitlementPolicy
		}
		qEntitlementPolicy := qrEntitlementPolicy
		if qEntitlementPolicy != "" {

			if err := r.SetQueryParam("entitlementPolicy", qEntitlementPolicy); err != nil {
				return err
			}
		}
	}

	if o.PluginProperty != nil {

		// binding items for pluginProperty
		joinedPluginProperty := o.bindParamPluginProperty(reg)

		// query array param pluginProperty
		if err := r.SetQueryParam("pluginProperty", joinedPluginProperty...); err != nil {
			return err
		}
	}

	if o.RequestedDate != nil {

		// query param requestedDate
		var qrRequestedDate strfmt.Date

		if o.RequestedDate != nil {
			qrRequestedDate = *o.RequestedDate
		}
		qRequestedDate := qrRequestedDate.String()
		if qRequestedDate != "" {

			if err := r.SetQueryParam("requestedDate", qRequestedDate); err != nil {
				return err
			}
		}
	}

	// path param subscriptionId
	if err := r.SetPathParam("subscriptionId", o.SubscriptionID.String()); err != nil {
		return err
	}

	if o.UseRequestedDateForBilling != nil {

		// query param useRequestedDateForBilling
		var qrUseRequestedDateForBilling bool

		if o.UseRequestedDateForBilling != nil {
			qrUseRequestedDateForBilling = *o.UseRequestedDateForBilling
		}
		qUseRequestedDateForBilling := swag.FormatBool(qrUseRequestedDateForBilling)
		if qUseRequestedDateForBilling != "" {

			if err := r.SetQueryParam("useRequestedDateForBilling", qUseRequestedDateForBilling); err != nil {
				return err
			}
		}
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

// bindParamCancelSubscriptionPlan binds the parameter pluginProperty
func (o *CancelSubscriptionPlanParams) bindParamPluginProperty(formats strfmt.Registry) []string {
	pluginPropertyIR := o.PluginProperty

	var pluginPropertyIC []string
	for _, pluginPropertyIIR := range pluginPropertyIR { // explode []string

		pluginPropertyIIV := pluginPropertyIIR // string as string
		pluginPropertyIC = append(pluginPropertyIC, pluginPropertyIIV)
	}

	// items.CollectionFormat: "multi"
	pluginPropertyIS := swag.JoinByFormat(pluginPropertyIC, "multi")

	return pluginPropertyIS
}
