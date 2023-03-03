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

// AddAccountBlockingStateReader is a Reader for the AddAccountBlockingState structure.
type AddAccountBlockingStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddAccountBlockingStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewAddAccountBlockingStateCreated()
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

// NewAddAccountBlockingStateCreated creates a AddAccountBlockingStateCreated with default headers values
func NewAddAccountBlockingStateCreated() *AddAccountBlockingStateCreated {
	return &AddAccountBlockingStateCreated{}
}

/*
AddAccountBlockingStateCreated describes a response with status code 201, with default header values.

Blocking state created successfully
*/
type AddAccountBlockingStateCreated struct {
	Payload      []*kbmodel.BlockingState
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the add account blocking state created response
func (o *AddAccountBlockingStateCreated) Code() int {
	return 201
}

// IsSuccess returns true when this add account blocking state created response has a 2xx status code
func (o *AddAccountBlockingStateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add account blocking state created response has a 3xx status code
func (o *AddAccountBlockingStateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add account blocking state created response has a 4xx status code
func (o *AddAccountBlockingStateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add account blocking state created response has a 5xx status code
func (o *AddAccountBlockingStateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add account blocking state created response a status code equal to that given
func (o *AddAccountBlockingStateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *AddAccountBlockingStateCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/block][%d] addAccountBlockingStateCreated  %+v", 201, o.Payload)
}

func (o *AddAccountBlockingStateCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/block][%d] addAccountBlockingStateCreated  %+v", 201, o.Payload)
}

func (o *AddAccountBlockingStateCreated) GetPayload() []*kbmodel.BlockingState {
	return o.Payload
}

func (o *AddAccountBlockingStateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddAccountBlockingStateBadRequest creates a AddAccountBlockingStateBadRequest with default headers values
func NewAddAccountBlockingStateBadRequest() *AddAccountBlockingStateBadRequest {
	return &AddAccountBlockingStateBadRequest{}
}

/*
AddAccountBlockingStateBadRequest describes a response with status code 400, with default header values.

Invalid account id supplied
*/
type AddAccountBlockingStateBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the add account blocking state bad request response
func (o *AddAccountBlockingStateBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this add account blocking state bad request response has a 2xx status code
func (o *AddAccountBlockingStateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add account blocking state bad request response has a 3xx status code
func (o *AddAccountBlockingStateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add account blocking state bad request response has a 4xx status code
func (o *AddAccountBlockingStateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add account blocking state bad request response has a 5xx status code
func (o *AddAccountBlockingStateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add account blocking state bad request response a status code equal to that given
func (o *AddAccountBlockingStateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *AddAccountBlockingStateBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/block][%d] addAccountBlockingStateBadRequest ", 400)
}

func (o *AddAccountBlockingStateBadRequest) String() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/block][%d] addAccountBlockingStateBadRequest ", 400)
}

func (o *AddAccountBlockingStateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddAccountBlockingStateNotFound creates a AddAccountBlockingStateNotFound with default headers values
func NewAddAccountBlockingStateNotFound() *AddAccountBlockingStateNotFound {
	return &AddAccountBlockingStateNotFound{}
}

/*
AddAccountBlockingStateNotFound describes a response with status code 404, with default header values.

Account not found
*/
type AddAccountBlockingStateNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the add account blocking state not found response
func (o *AddAccountBlockingStateNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this add account blocking state not found response has a 2xx status code
func (o *AddAccountBlockingStateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add account blocking state not found response has a 3xx status code
func (o *AddAccountBlockingStateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add account blocking state not found response has a 4xx status code
func (o *AddAccountBlockingStateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add account blocking state not found response has a 5xx status code
func (o *AddAccountBlockingStateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add account blocking state not found response a status code equal to that given
func (o *AddAccountBlockingStateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *AddAccountBlockingStateNotFound) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/block][%d] addAccountBlockingStateNotFound ", 404)
}

func (o *AddAccountBlockingStateNotFound) String() string {
	return fmt.Sprintf("[POST /1.0/kb/accounts/{accountId}/block][%d] addAccountBlockingStateNotFound ", 404)
}

func (o *AddAccountBlockingStateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
