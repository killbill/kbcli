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
)

// DeleteInvoicePaymentCustomFieldsReader is a Reader for the DeleteInvoicePaymentCustomFields structure.
type DeleteInvoicePaymentCustomFieldsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteInvoicePaymentCustomFieldsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteInvoicePaymentCustomFieldsNoContent()
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

// NewDeleteInvoicePaymentCustomFieldsNoContent creates a DeleteInvoicePaymentCustomFieldsNoContent with default headers values
func NewDeleteInvoicePaymentCustomFieldsNoContent() *DeleteInvoicePaymentCustomFieldsNoContent {
	return &DeleteInvoicePaymentCustomFieldsNoContent{}
}

/*
DeleteInvoicePaymentCustomFieldsNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteInvoicePaymentCustomFieldsNoContent struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the delete invoice payment custom fields no content response
func (o *DeleteInvoicePaymentCustomFieldsNoContent) Code() int {
	return 204
}

// IsSuccess returns true when this delete invoice payment custom fields no content response has a 2xx status code
func (o *DeleteInvoicePaymentCustomFieldsNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete invoice payment custom fields no content response has a 3xx status code
func (o *DeleteInvoicePaymentCustomFieldsNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete invoice payment custom fields no content response has a 4xx status code
func (o *DeleteInvoicePaymentCustomFieldsNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete invoice payment custom fields no content response has a 5xx status code
func (o *DeleteInvoicePaymentCustomFieldsNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete invoice payment custom fields no content response a status code equal to that given
func (o *DeleteInvoicePaymentCustomFieldsNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteInvoicePaymentCustomFieldsNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/invoicePayments/{paymentId}/customFields][%d] deleteInvoicePaymentCustomFieldsNoContent ", 204)
}

func (o *DeleteInvoicePaymentCustomFieldsNoContent) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/invoicePayments/{paymentId}/customFields][%d] deleteInvoicePaymentCustomFieldsNoContent ", 204)
}

func (o *DeleteInvoicePaymentCustomFieldsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteInvoicePaymentCustomFieldsBadRequest creates a DeleteInvoicePaymentCustomFieldsBadRequest with default headers values
func NewDeleteInvoicePaymentCustomFieldsBadRequest() *DeleteInvoicePaymentCustomFieldsBadRequest {
	return &DeleteInvoicePaymentCustomFieldsBadRequest{}
}

/*
DeleteInvoicePaymentCustomFieldsBadRequest describes a response with status code 400, with default header values.

Invalid payment id supplied
*/
type DeleteInvoicePaymentCustomFieldsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the delete invoice payment custom fields bad request response
func (o *DeleteInvoicePaymentCustomFieldsBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this delete invoice payment custom fields bad request response has a 2xx status code
func (o *DeleteInvoicePaymentCustomFieldsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete invoice payment custom fields bad request response has a 3xx status code
func (o *DeleteInvoicePaymentCustomFieldsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete invoice payment custom fields bad request response has a 4xx status code
func (o *DeleteInvoicePaymentCustomFieldsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete invoice payment custom fields bad request response has a 5xx status code
func (o *DeleteInvoicePaymentCustomFieldsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete invoice payment custom fields bad request response a status code equal to that given
func (o *DeleteInvoicePaymentCustomFieldsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteInvoicePaymentCustomFieldsBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/invoicePayments/{paymentId}/customFields][%d] deleteInvoicePaymentCustomFieldsBadRequest ", 400)
}

func (o *DeleteInvoicePaymentCustomFieldsBadRequest) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/invoicePayments/{paymentId}/customFields][%d] deleteInvoicePaymentCustomFieldsBadRequest ", 400)
}

func (o *DeleteInvoicePaymentCustomFieldsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
