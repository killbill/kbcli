// Code generated by go-swagger; DO NOT EDIT.

package payment_method

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v2/kbcommon"
)

// DeletePaymentMethodReader is a Reader for the DeletePaymentMethod structure.
type DeletePaymentMethodReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePaymentMethodReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeletePaymentMethodNoContent()
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

// NewDeletePaymentMethodNoContent creates a DeletePaymentMethodNoContent with default headers values
func NewDeletePaymentMethodNoContent() *DeletePaymentMethodNoContent {
	return &DeletePaymentMethodNoContent{}
}

/*
DeletePaymentMethodNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeletePaymentMethodNoContent struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the delete payment method no content response
func (o *DeletePaymentMethodNoContent) Code() int {
	return 204
}

// IsSuccess returns true when this delete payment method no content response has a 2xx status code
func (o *DeletePaymentMethodNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete payment method no content response has a 3xx status code
func (o *DeletePaymentMethodNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete payment method no content response has a 4xx status code
func (o *DeletePaymentMethodNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete payment method no content response has a 5xx status code
func (o *DeletePaymentMethodNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete payment method no content response a status code equal to that given
func (o *DeletePaymentMethodNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeletePaymentMethodNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/paymentMethods/{paymentMethodId}][%d] deletePaymentMethodNoContent ", 204)
}

func (o *DeletePaymentMethodNoContent) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/paymentMethods/{paymentMethodId}][%d] deletePaymentMethodNoContent ", 204)
}

func (o *DeletePaymentMethodNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePaymentMethodBadRequest creates a DeletePaymentMethodBadRequest with default headers values
func NewDeletePaymentMethodBadRequest() *DeletePaymentMethodBadRequest {
	return &DeletePaymentMethodBadRequest{}
}

/*
DeletePaymentMethodBadRequest describes a response with status code 400, with default header values.

Invalid paymentMethodId supplied
*/
type DeletePaymentMethodBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the delete payment method bad request response
func (o *DeletePaymentMethodBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this delete payment method bad request response has a 2xx status code
func (o *DeletePaymentMethodBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete payment method bad request response has a 3xx status code
func (o *DeletePaymentMethodBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete payment method bad request response has a 4xx status code
func (o *DeletePaymentMethodBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete payment method bad request response has a 5xx status code
func (o *DeletePaymentMethodBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete payment method bad request response a status code equal to that given
func (o *DeletePaymentMethodBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeletePaymentMethodBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/paymentMethods/{paymentMethodId}][%d] deletePaymentMethodBadRequest ", 400)
}

func (o *DeletePaymentMethodBadRequest) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/paymentMethods/{paymentMethodId}][%d] deletePaymentMethodBadRequest ", 400)
}

func (o *DeletePaymentMethodBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePaymentMethodNotFound creates a DeletePaymentMethodNotFound with default headers values
func NewDeletePaymentMethodNotFound() *DeletePaymentMethodNotFound {
	return &DeletePaymentMethodNotFound{}
}

/*
DeletePaymentMethodNotFound describes a response with status code 404, with default header values.

Account or payment method not found
*/
type DeletePaymentMethodNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the delete payment method not found response
func (o *DeletePaymentMethodNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this delete payment method not found response has a 2xx status code
func (o *DeletePaymentMethodNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete payment method not found response has a 3xx status code
func (o *DeletePaymentMethodNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete payment method not found response has a 4xx status code
func (o *DeletePaymentMethodNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete payment method not found response has a 5xx status code
func (o *DeletePaymentMethodNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete payment method not found response a status code equal to that given
func (o *DeletePaymentMethodNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *DeletePaymentMethodNotFound) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/paymentMethods/{paymentMethodId}][%d] deletePaymentMethodNotFound ", 404)
}

func (o *DeletePaymentMethodNotFound) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/paymentMethods/{paymentMethodId}][%d] deletePaymentMethodNotFound ", 404)
}

func (o *DeletePaymentMethodNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
