// Code generated by go-swagger; DO NOT EDIT.

package account

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v3/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// RefreshPaymentMethodsReader is a Reader for the RefreshPaymentMethods structure.
type RefreshPaymentMethodsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RefreshPaymentMethodsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewRefreshPaymentMethodsNoContent()
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

// NewRefreshPaymentMethodsNoContent creates a RefreshPaymentMethodsNoContent with default headers values
func NewRefreshPaymentMethodsNoContent() *RefreshPaymentMethodsNoContent {
	return &RefreshPaymentMethodsNoContent{}
}

/*RefreshPaymentMethodsNoContent handles this case with default header values.

Successful operation
*/
type RefreshPaymentMethodsNoContent struct {
	HttpResponse runtime.ClientResponse
}

func (o *RefreshPaymentMethodsNoContent) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/accounts/{accountId}/paymentMethods/refresh][%d] refreshPaymentMethodsNoContent ", 204)
}

func (o *RefreshPaymentMethodsNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRefreshPaymentMethodsBadRequest creates a RefreshPaymentMethodsBadRequest with default headers values
func NewRefreshPaymentMethodsBadRequest() *RefreshPaymentMethodsBadRequest {
	return &RefreshPaymentMethodsBadRequest{}
}

/*RefreshPaymentMethodsBadRequest handles this case with default header values.

Invalid account id supplied
*/
type RefreshPaymentMethodsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *RefreshPaymentMethodsBadRequest) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/accounts/{accountId}/paymentMethods/refresh][%d] refreshPaymentMethodsBadRequest ", 400)
}

func (o *RefreshPaymentMethodsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRefreshPaymentMethodsNotFound creates a RefreshPaymentMethodsNotFound with default headers values
func NewRefreshPaymentMethodsNotFound() *RefreshPaymentMethodsNotFound {
	return &RefreshPaymentMethodsNotFound{}
}

/*RefreshPaymentMethodsNotFound handles this case with default header values.

Account not found
*/
type RefreshPaymentMethodsNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *RefreshPaymentMethodsNotFound) Error() string {
	return fmt.Sprintf("[PUT /1.0/kb/accounts/{accountId}/paymentMethods/refresh][%d] refreshPaymentMethodsNotFound ", 404)
}

func (o *RefreshPaymentMethodsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
