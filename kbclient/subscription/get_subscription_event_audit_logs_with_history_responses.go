// Code generated by go-swagger; DO NOT EDIT.

package subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v2/kbcommon"

	"github.com/killbill/kbcli/v2/kbmodel"
)

// GetSubscriptionEventAuditLogsWithHistoryReader is a Reader for the GetSubscriptionEventAuditLogsWithHistory structure.
type GetSubscriptionEventAuditLogsWithHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSubscriptionEventAuditLogsWithHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSubscriptionEventAuditLogsWithHistoryOK()
		result.HttpResponse = response
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		errorResult := kbcommon.NewKillbillError(response.Code())
		if err := consumer.Consume(response.Body(), &errorResult); err != nil && err != io.EOF {
			return nil, err
		}
		return nil, errorResult
	}
}

// NewGetSubscriptionEventAuditLogsWithHistoryOK creates a GetSubscriptionEventAuditLogsWithHistoryOK with default headers values
func NewGetSubscriptionEventAuditLogsWithHistoryOK() *GetSubscriptionEventAuditLogsWithHistoryOK {
	return &GetSubscriptionEventAuditLogsWithHistoryOK{}
}

/*
GetSubscriptionEventAuditLogsWithHistoryOK describes a response with status code 200, with default header values.

successful operation
*/
type GetSubscriptionEventAuditLogsWithHistoryOK struct {
	Payload      []*kbmodel.AuditLog
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get subscription event audit logs with history o k response
func (o *GetSubscriptionEventAuditLogsWithHistoryOK) Code() int {
	return 200
}

// IsSuccess returns true when this get subscription event audit logs with history o k response has a 2xx status code
func (o *GetSubscriptionEventAuditLogsWithHistoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get subscription event audit logs with history o k response has a 3xx status code
func (o *GetSubscriptionEventAuditLogsWithHistoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get subscription event audit logs with history o k response has a 4xx status code
func (o *GetSubscriptionEventAuditLogsWithHistoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get subscription event audit logs with history o k response has a 5xx status code
func (o *GetSubscriptionEventAuditLogsWithHistoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get subscription event audit logs with history o k response a status code equal to that given
func (o *GetSubscriptionEventAuditLogsWithHistoryOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetSubscriptionEventAuditLogsWithHistoryOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions/events/{eventId}/auditLogsWithHistory][%d] getSubscriptionEventAuditLogsWithHistoryOK  %+v", 200, o.Payload)
}

func (o *GetSubscriptionEventAuditLogsWithHistoryOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions/events/{eventId}/auditLogsWithHistory][%d] getSubscriptionEventAuditLogsWithHistoryOK  %+v", 200, o.Payload)
}

func (o *GetSubscriptionEventAuditLogsWithHistoryOK) GetPayload() []*kbmodel.AuditLog {
	return o.Payload
}

func (o *GetSubscriptionEventAuditLogsWithHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSubscriptionEventAuditLogsWithHistoryNotFound creates a GetSubscriptionEventAuditLogsWithHistoryNotFound with default headers values
func NewGetSubscriptionEventAuditLogsWithHistoryNotFound() *GetSubscriptionEventAuditLogsWithHistoryNotFound {
	return &GetSubscriptionEventAuditLogsWithHistoryNotFound{}
}

/*
GetSubscriptionEventAuditLogsWithHistoryNotFound describes a response with status code 404, with default header values.

Subscription event not found
*/
type GetSubscriptionEventAuditLogsWithHistoryNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get subscription event audit logs with history not found response
func (o *GetSubscriptionEventAuditLogsWithHistoryNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this get subscription event audit logs with history not found response has a 2xx status code
func (o *GetSubscriptionEventAuditLogsWithHistoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get subscription event audit logs with history not found response has a 3xx status code
func (o *GetSubscriptionEventAuditLogsWithHistoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get subscription event audit logs with history not found response has a 4xx status code
func (o *GetSubscriptionEventAuditLogsWithHistoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get subscription event audit logs with history not found response has a 5xx status code
func (o *GetSubscriptionEventAuditLogsWithHistoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get subscription event audit logs with history not found response a status code equal to that given
func (o *GetSubscriptionEventAuditLogsWithHistoryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetSubscriptionEventAuditLogsWithHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions/events/{eventId}/auditLogsWithHistory][%d] getSubscriptionEventAuditLogsWithHistoryNotFound ", 404)
}

func (o *GetSubscriptionEventAuditLogsWithHistoryNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/subscriptions/events/{eventId}/auditLogsWithHistory][%d] getSubscriptionEventAuditLogsWithHistoryNotFound ", 404)
}

func (o *GetSubscriptionEventAuditLogsWithHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
