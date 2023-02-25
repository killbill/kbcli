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

	kbmodel "github.com/killbill/kbcli/v3/kbmodel"
)

// GetOverdueAccountReader is a Reader for the GetOverdueAccount structure.
type GetOverdueAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOverdueAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOverdueAccountOK()
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

// NewGetOverdueAccountOK creates a GetOverdueAccountOK with default headers values
func NewGetOverdueAccountOK() *GetOverdueAccountOK {
	return &GetOverdueAccountOK{}
}

/*GetOverdueAccountOK handles this case with default header values.

successful operation
*/
type GetOverdueAccountOK struct {
	Payload *kbmodel.OverdueState

	HttpResponse runtime.ClientResponse
}

func (o *GetOverdueAccountOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/overdue][%d] getOverdueAccountOK  %+v", 200, o.Payload)
}

func (o *GetOverdueAccountOK) GetPayload() *kbmodel.OverdueState {
	return o.Payload
}

func (o *GetOverdueAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.OverdueState)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOverdueAccountBadRequest creates a GetOverdueAccountBadRequest with default headers values
func NewGetOverdueAccountBadRequest() *GetOverdueAccountBadRequest {
	return &GetOverdueAccountBadRequest{}
}

/*GetOverdueAccountBadRequest handles this case with default header values.

Invalid account id supplied
*/
type GetOverdueAccountBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetOverdueAccountBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/overdue][%d] getOverdueAccountBadRequest ", 400)
}

func (o *GetOverdueAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOverdueAccountNotFound creates a GetOverdueAccountNotFound with default headers values
func NewGetOverdueAccountNotFound() *GetOverdueAccountNotFound {
	return &GetOverdueAccountNotFound{}
}

/*GetOverdueAccountNotFound handles this case with default header values.

Account not found
*/
type GetOverdueAccountNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetOverdueAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/overdue][%d] getOverdueAccountNotFound ", 404)
}

func (o *GetOverdueAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
