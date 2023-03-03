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

/*
GetOverdueAccountOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOverdueAccountOK struct {
	Payload      *kbmodel.OverdueState
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get overdue account o k response
func (o *GetOverdueAccountOK) Code() int {
	return 200
}

// IsSuccess returns true when this get overdue account o k response has a 2xx status code
func (o *GetOverdueAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get overdue account o k response has a 3xx status code
func (o *GetOverdueAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get overdue account o k response has a 4xx status code
func (o *GetOverdueAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get overdue account o k response has a 5xx status code
func (o *GetOverdueAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get overdue account o k response a status code equal to that given
func (o *GetOverdueAccountOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOverdueAccountOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/overdue][%d] getOverdueAccountOK  %+v", 200, o.Payload)
}

func (o *GetOverdueAccountOK) String() string {
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

/*
GetOverdueAccountBadRequest describes a response with status code 400, with default header values.

Invalid account id supplied
*/
type GetOverdueAccountBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get overdue account bad request response
func (o *GetOverdueAccountBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this get overdue account bad request response has a 2xx status code
func (o *GetOverdueAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get overdue account bad request response has a 3xx status code
func (o *GetOverdueAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get overdue account bad request response has a 4xx status code
func (o *GetOverdueAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get overdue account bad request response has a 5xx status code
func (o *GetOverdueAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get overdue account bad request response a status code equal to that given
func (o *GetOverdueAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOverdueAccountBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/overdue][%d] getOverdueAccountBadRequest ", 400)
}

func (o *GetOverdueAccountBadRequest) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/overdue][%d] getOverdueAccountBadRequest ", 400)
}

func (o *GetOverdueAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOverdueAccountNotFound creates a GetOverdueAccountNotFound with default headers values
func NewGetOverdueAccountNotFound() *GetOverdueAccountNotFound {
	return &GetOverdueAccountNotFound{}
}

/*
GetOverdueAccountNotFound describes a response with status code 404, with default header values.

Account not found
*/
type GetOverdueAccountNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get overdue account not found response
func (o *GetOverdueAccountNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this get overdue account not found response has a 2xx status code
func (o *GetOverdueAccountNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get overdue account not found response has a 3xx status code
func (o *GetOverdueAccountNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get overdue account not found response has a 4xx status code
func (o *GetOverdueAccountNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get overdue account not found response has a 5xx status code
func (o *GetOverdueAccountNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get overdue account not found response a status code equal to that given
func (o *GetOverdueAccountNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOverdueAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/overdue][%d] getOverdueAccountNotFound ", 404)
}

func (o *GetOverdueAccountNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/overdue][%d] getOverdueAccountNotFound ", 404)
}

func (o *GetOverdueAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
