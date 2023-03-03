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

// CreateChargebackReader is a Reader for the CreateChargeback structure.
type CreateChargebackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateChargebackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreateChargebackCreated()
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

// NewCreateChargebackCreated creates a CreateChargebackCreated with default headers values
func NewCreateChargebackCreated() *CreateChargebackCreated {
	return &CreateChargebackCreated{}
}

/*
CreateChargebackCreated describes a response with status code 201, with default header values.

Created chargeback successfully
*/
type CreateChargebackCreated struct {
	Payload      *kbmodel.InvoicePayment
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the create chargeback created response
func (o *CreateChargebackCreated) Code() int {
	return 201
}

// IsSuccess returns true when this create chargeback created response has a 2xx status code
func (o *CreateChargebackCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create chargeback created response has a 3xx status code
func (o *CreateChargebackCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create chargeback created response has a 4xx status code
func (o *CreateChargebackCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create chargeback created response has a 5xx status code
func (o *CreateChargebackCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create chargeback created response a status code equal to that given
func (o *CreateChargebackCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateChargebackCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/chargebacks][%d] createChargebackCreated  %+v", 201, o.Payload)
}

func (o *CreateChargebackCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/chargebacks][%d] createChargebackCreated  %+v", 201, o.Payload)
}

func (o *CreateChargebackCreated) GetPayload() *kbmodel.InvoicePayment {
	return o.Payload
}

func (o *CreateChargebackCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.InvoicePayment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateChargebackBadRequest creates a CreateChargebackBadRequest with default headers values
func NewCreateChargebackBadRequest() *CreateChargebackBadRequest {
	return &CreateChargebackBadRequest{}
}

/*
CreateChargebackBadRequest describes a response with status code 400, with default header values.

Invalid payment id supplied
*/
type CreateChargebackBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the create chargeback bad request response
func (o *CreateChargebackBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this create chargeback bad request response has a 2xx status code
func (o *CreateChargebackBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create chargeback bad request response has a 3xx status code
func (o *CreateChargebackBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create chargeback bad request response has a 4xx status code
func (o *CreateChargebackBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create chargeback bad request response has a 5xx status code
func (o *CreateChargebackBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create chargeback bad request response a status code equal to that given
func (o *CreateChargebackBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateChargebackBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/chargebacks][%d] createChargebackBadRequest ", 400)
}

func (o *CreateChargebackBadRequest) String() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/chargebacks][%d] createChargebackBadRequest ", 400)
}

func (o *CreateChargebackBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCreateChargebackNotFound creates a CreateChargebackNotFound with default headers values
func NewCreateChargebackNotFound() *CreateChargebackNotFound {
	return &CreateChargebackNotFound{}
}

/*
CreateChargebackNotFound describes a response with status code 404, with default header values.

Account or payment not found
*/
type CreateChargebackNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the create chargeback not found response
func (o *CreateChargebackNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this create chargeback not found response has a 2xx status code
func (o *CreateChargebackNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create chargeback not found response has a 3xx status code
func (o *CreateChargebackNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create chargeback not found response has a 4xx status code
func (o *CreateChargebackNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this create chargeback not found response has a 5xx status code
func (o *CreateChargebackNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this create chargeback not found response a status code equal to that given
func (o *CreateChargebackNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *CreateChargebackNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/chargebacks][%d] createChargebackNotFound ", 404)
}

func (o *CreateChargebackNotFound) String() string {
	return fmt.Sprintf("[POST /1.0/kb/invoicePayments/{paymentId}/chargebacks][%d] createChargebackNotFound ", 404)
}

func (o *CreateChargebackNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
