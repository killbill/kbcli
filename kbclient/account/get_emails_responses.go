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

// GetEmailsReader is a Reader for the GetEmails structure.
type GetEmailsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEmailsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetEmailsOK()
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

// NewGetEmailsOK creates a GetEmailsOK with default headers values
func NewGetEmailsOK() *GetEmailsOK {
	return &GetEmailsOK{}
}

/*
GetEmailsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetEmailsOK struct {
	Payload      []*kbmodel.AccountEmail
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get emails o k response
func (o *GetEmailsOK) Code() int {
	return 200
}

// IsSuccess returns true when this get emails o k response has a 2xx status code
func (o *GetEmailsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get emails o k response has a 3xx status code
func (o *GetEmailsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get emails o k response has a 4xx status code
func (o *GetEmailsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get emails o k response has a 5xx status code
func (o *GetEmailsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get emails o k response a status code equal to that given
func (o *GetEmailsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetEmailsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/emails][%d] getEmailsOK  %+v", 200, o.Payload)
}

func (o *GetEmailsOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/emails][%d] getEmailsOK  %+v", 200, o.Payload)
}

func (o *GetEmailsOK) GetPayload() []*kbmodel.AccountEmail {
	return o.Payload
}

func (o *GetEmailsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEmailsBadRequest creates a GetEmailsBadRequest with default headers values
func NewGetEmailsBadRequest() *GetEmailsBadRequest {
	return &GetEmailsBadRequest{}
}

/*
GetEmailsBadRequest describes a response with status code 400, with default header values.

Invalid account id supplied
*/
type GetEmailsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get emails bad request response
func (o *GetEmailsBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this get emails bad request response has a 2xx status code
func (o *GetEmailsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get emails bad request response has a 3xx status code
func (o *GetEmailsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get emails bad request response has a 4xx status code
func (o *GetEmailsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get emails bad request response has a 5xx status code
func (o *GetEmailsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get emails bad request response a status code equal to that given
func (o *GetEmailsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetEmailsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/emails][%d] getEmailsBadRequest ", 400)
}

func (o *GetEmailsBadRequest) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/emails][%d] getEmailsBadRequest ", 400)
}

func (o *GetEmailsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
