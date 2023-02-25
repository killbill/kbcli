// Code generated by go-swagger; DO NOT EDIT.

package payment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v3/kbcommon"

	strfmt "github.com/go-openapi/strfmt"

	kbmodel "github.com/killbill/kbcli/v3/kbmodel"
)

// GetPaymentReader is a Reader for the GetPayment structure.
type GetPaymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentOK()
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

// NewGetPaymentOK creates a GetPaymentOK with default headers values
func NewGetPaymentOK() *GetPaymentOK {
	return &GetPaymentOK{}
}

/*GetPaymentOK handles this case with default header values.

successful operation
*/
type GetPaymentOK struct {
	Payload *kbmodel.Payment

	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/payments/{paymentId}][%d] getPaymentOK  %+v", 200, o.Payload)
}

func (o *GetPaymentOK) GetPayload() *kbmodel.Payment {
	return o.Payload
}

func (o *GetPaymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Payment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentBadRequest creates a GetPaymentBadRequest with default headers values
func NewGetPaymentBadRequest() *GetPaymentBadRequest {
	return &GetPaymentBadRequest{}
}

/*GetPaymentBadRequest handles this case with default header values.

Invalid paymentId supplied
*/
type GetPaymentBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/payments/{paymentId}][%d] getPaymentBadRequest ", 400)
}

func (o *GetPaymentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetPaymentNotFound creates a GetPaymentNotFound with default headers values
func NewGetPaymentNotFound() *GetPaymentNotFound {
	return &GetPaymentNotFound{}
}

/*GetPaymentNotFound handles this case with default header values.

Payment not found
*/
type GetPaymentNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/payments/{paymentId}][%d] getPaymentNotFound ", 404)
}

func (o *GetPaymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
