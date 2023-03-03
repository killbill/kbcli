// Code generated by go-swagger; DO NOT EDIT.

package account

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

// GetPaymentMethodsForAccountReader is a Reader for the GetPaymentMethodsForAccount structure.
type GetPaymentMethodsForAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentMethodsForAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentMethodsForAccountOK()
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

// NewGetPaymentMethodsForAccountOK creates a GetPaymentMethodsForAccountOK with default headers values
func NewGetPaymentMethodsForAccountOK() *GetPaymentMethodsForAccountOK {
	return &GetPaymentMethodsForAccountOK{}
}

/*
GetPaymentMethodsForAccountOK describes a response with status code 200, with default header values.

successful operation
*/
type GetPaymentMethodsForAccountOK struct {
	Payload      []*kbmodel.PaymentMethod
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get payment methods for account o k response
func (o *GetPaymentMethodsForAccountOK) Code() int {
	return 200
}

// IsSuccess returns true when this get payment methods for account o k response has a 2xx status code
func (o *GetPaymentMethodsForAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get payment methods for account o k response has a 3xx status code
func (o *GetPaymentMethodsForAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get payment methods for account o k response has a 4xx status code
func (o *GetPaymentMethodsForAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get payment methods for account o k response has a 5xx status code
func (o *GetPaymentMethodsForAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get payment methods for account o k response a status code equal to that given
func (o *GetPaymentMethodsForAccountOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetPaymentMethodsForAccountOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/paymentMethods][%d] getPaymentMethodsForAccountOK  %+v", 200, o.Payload)
}

func (o *GetPaymentMethodsForAccountOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/paymentMethods][%d] getPaymentMethodsForAccountOK  %+v", 200, o.Payload)
}

func (o *GetPaymentMethodsForAccountOK) GetPayload() []*kbmodel.PaymentMethod {
	return o.Payload
}

func (o *GetPaymentMethodsForAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentMethodsForAccountBadRequest creates a GetPaymentMethodsForAccountBadRequest with default headers values
func NewGetPaymentMethodsForAccountBadRequest() *GetPaymentMethodsForAccountBadRequest {
	return &GetPaymentMethodsForAccountBadRequest{}
}

/*
GetPaymentMethodsForAccountBadRequest describes a response with status code 400, with default header values.

Invalid account id supplied
*/
type GetPaymentMethodsForAccountBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get payment methods for account bad request response
func (o *GetPaymentMethodsForAccountBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this get payment methods for account bad request response has a 2xx status code
func (o *GetPaymentMethodsForAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get payment methods for account bad request response has a 3xx status code
func (o *GetPaymentMethodsForAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get payment methods for account bad request response has a 4xx status code
func (o *GetPaymentMethodsForAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get payment methods for account bad request response has a 5xx status code
func (o *GetPaymentMethodsForAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get payment methods for account bad request response a status code equal to that given
func (o *GetPaymentMethodsForAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetPaymentMethodsForAccountBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/paymentMethods][%d] getPaymentMethodsForAccountBadRequest ", 400)
}

func (o *GetPaymentMethodsForAccountBadRequest) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/paymentMethods][%d] getPaymentMethodsForAccountBadRequest ", 400)
}

func (o *GetPaymentMethodsForAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPaymentMethodsForAccountNotFound creates a GetPaymentMethodsForAccountNotFound with default headers values
func NewGetPaymentMethodsForAccountNotFound() *GetPaymentMethodsForAccountNotFound {
	return &GetPaymentMethodsForAccountNotFound{}
}

/*
GetPaymentMethodsForAccountNotFound describes a response with status code 404, with default header values.

Account not found
*/
type GetPaymentMethodsForAccountNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get payment methods for account not found response
func (o *GetPaymentMethodsForAccountNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this get payment methods for account not found response has a 2xx status code
func (o *GetPaymentMethodsForAccountNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get payment methods for account not found response has a 3xx status code
func (o *GetPaymentMethodsForAccountNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get payment methods for account not found response has a 4xx status code
func (o *GetPaymentMethodsForAccountNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get payment methods for account not found response has a 5xx status code
func (o *GetPaymentMethodsForAccountNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get payment methods for account not found response a status code equal to that given
func (o *GetPaymentMethodsForAccountNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetPaymentMethodsForAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/paymentMethods][%d] getPaymentMethodsForAccountNotFound ", 404)
}

func (o *GetPaymentMethodsForAccountNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/paymentMethods][%d] getPaymentMethodsForAccountNotFound ", 404)
}

func (o *GetPaymentMethodsForAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
