// Code generated by go-swagger; DO NOT EDIT.

package invoice_payment

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

// GetInvoicePaymentAuditLogsWithHistoryReader is a Reader for the GetInvoicePaymentAuditLogsWithHistory structure.
type GetInvoicePaymentAuditLogsWithHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInvoicePaymentAuditLogsWithHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvoicePaymentAuditLogsWithHistoryOK()
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

// NewGetInvoicePaymentAuditLogsWithHistoryOK creates a GetInvoicePaymentAuditLogsWithHistoryOK with default headers values
func NewGetInvoicePaymentAuditLogsWithHistoryOK() *GetInvoicePaymentAuditLogsWithHistoryOK {
	return &GetInvoicePaymentAuditLogsWithHistoryOK{}
}

/*
GetInvoicePaymentAuditLogsWithHistoryOK describes a response with status code 200, with default header values.

successful operation
*/
type GetInvoicePaymentAuditLogsWithHistoryOK struct {
	Payload      []*kbmodel.AuditLog
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get invoice payment audit logs with history o k response
func (o *GetInvoicePaymentAuditLogsWithHistoryOK) Code() int {
	return 200
}

// IsSuccess returns true when this get invoice payment audit logs with history o k response has a 2xx status code
func (o *GetInvoicePaymentAuditLogsWithHistoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get invoice payment audit logs with history o k response has a 3xx status code
func (o *GetInvoicePaymentAuditLogsWithHistoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoice payment audit logs with history o k response has a 4xx status code
func (o *GetInvoicePaymentAuditLogsWithHistoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get invoice payment audit logs with history o k response has a 5xx status code
func (o *GetInvoicePaymentAuditLogsWithHistoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoice payment audit logs with history o k response a status code equal to that given
func (o *GetInvoicePaymentAuditLogsWithHistoryOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetInvoicePaymentAuditLogsWithHistoryOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoicePayments/{invoicePaymentId}/auditLogsWithHistory][%d] getInvoicePaymentAuditLogsWithHistoryOK  %+v", 200, o.Payload)
}

func (o *GetInvoicePaymentAuditLogsWithHistoryOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/invoicePayments/{invoicePaymentId}/auditLogsWithHistory][%d] getInvoicePaymentAuditLogsWithHistoryOK  %+v", 200, o.Payload)
}

func (o *GetInvoicePaymentAuditLogsWithHistoryOK) GetPayload() []*kbmodel.AuditLog {
	return o.Payload
}

func (o *GetInvoicePaymentAuditLogsWithHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInvoicePaymentAuditLogsWithHistoryNotFound creates a GetInvoicePaymentAuditLogsWithHistoryNotFound with default headers values
func NewGetInvoicePaymentAuditLogsWithHistoryNotFound() *GetInvoicePaymentAuditLogsWithHistoryNotFound {
	return &GetInvoicePaymentAuditLogsWithHistoryNotFound{}
}

/*
GetInvoicePaymentAuditLogsWithHistoryNotFound describes a response with status code 404, with default header values.

Invoice payment not found
*/
type GetInvoicePaymentAuditLogsWithHistoryNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get invoice payment audit logs with history not found response
func (o *GetInvoicePaymentAuditLogsWithHistoryNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this get invoice payment audit logs with history not found response has a 2xx status code
func (o *GetInvoicePaymentAuditLogsWithHistoryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get invoice payment audit logs with history not found response has a 3xx status code
func (o *GetInvoicePaymentAuditLogsWithHistoryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get invoice payment audit logs with history not found response has a 4xx status code
func (o *GetInvoicePaymentAuditLogsWithHistoryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get invoice payment audit logs with history not found response has a 5xx status code
func (o *GetInvoicePaymentAuditLogsWithHistoryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get invoice payment audit logs with history not found response a status code equal to that given
func (o *GetInvoicePaymentAuditLogsWithHistoryNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetInvoicePaymentAuditLogsWithHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/invoicePayments/{invoicePaymentId}/auditLogsWithHistory][%d] getInvoicePaymentAuditLogsWithHistoryNotFound ", 404)
}

func (o *GetInvoicePaymentAuditLogsWithHistoryNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/invoicePayments/{invoicePaymentId}/auditLogsWithHistory][%d] getInvoicePaymentAuditLogsWithHistoryNotFound ", 404)
}

func (o *GetInvoicePaymentAuditLogsWithHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
