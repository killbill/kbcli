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

// GetAccountReader is a Reader for the GetAccount structure.
type GetAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountOK()
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

// NewGetAccountOK creates a GetAccountOK with default headers values
func NewGetAccountOK() *GetAccountOK {
	return &GetAccountOK{}
}

/*
GetAccountOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAccountOK struct {
	Payload      *kbmodel.Account
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get account o k response
func (o *GetAccountOK) Code() int {
	return 200
}

// IsSuccess returns true when this get account o k response has a 2xx status code
func (o *GetAccountOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get account o k response has a 3xx status code
func (o *GetAccountOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get account o k response has a 4xx status code
func (o *GetAccountOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get account o k response has a 5xx status code
func (o *GetAccountOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get account o k response a status code equal to that given
func (o *GetAccountOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAccountOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}][%d] getAccountOK  %+v", 200, o.Payload)
}

func (o *GetAccountOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}][%d] getAccountOK  %+v", 200, o.Payload)
}

func (o *GetAccountOK) GetPayload() *kbmodel.Account {
	return o.Payload
}

func (o *GetAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(kbmodel.Account)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountBadRequest creates a GetAccountBadRequest with default headers values
func NewGetAccountBadRequest() *GetAccountBadRequest {
	return &GetAccountBadRequest{}
}

/*
GetAccountBadRequest describes a response with status code 400, with default header values.

Invalid account id supplied
*/
type GetAccountBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get account bad request response
func (o *GetAccountBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this get account bad request response has a 2xx status code
func (o *GetAccountBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get account bad request response has a 3xx status code
func (o *GetAccountBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get account bad request response has a 4xx status code
func (o *GetAccountBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get account bad request response has a 5xx status code
func (o *GetAccountBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get account bad request response a status code equal to that given
func (o *GetAccountBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetAccountBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}][%d] getAccountBadRequest ", 400)
}

func (o *GetAccountBadRequest) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}][%d] getAccountBadRequest ", 400)
}

func (o *GetAccountBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAccountNotFound creates a GetAccountNotFound with default headers values
func NewGetAccountNotFound() *GetAccountNotFound {
	return &GetAccountNotFound{}
}

/*
GetAccountNotFound describes a response with status code 404, with default header values.

Account not found
*/
type GetAccountNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get account not found response
func (o *GetAccountNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this get account not found response has a 2xx status code
func (o *GetAccountNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get account not found response has a 3xx status code
func (o *GetAccountNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get account not found response has a 4xx status code
func (o *GetAccountNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get account not found response has a 5xx status code
func (o *GetAccountNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get account not found response a status code equal to that given
func (o *GetAccountNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAccountNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}][%d] getAccountNotFound ", 404)
}

func (o *GetAccountNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}][%d] getAccountNotFound ", 404)
}

func (o *GetAccountNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
