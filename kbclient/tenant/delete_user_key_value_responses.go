// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/killbill/kbcli/v2/kbcommon"
)

// DeleteUserKeyValueReader is a Reader for the DeleteUserKeyValue structure.
type DeleteUserKeyValueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteUserKeyValueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteUserKeyValueNoContent()
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

// NewDeleteUserKeyValueNoContent creates a DeleteUserKeyValueNoContent with default headers values
func NewDeleteUserKeyValueNoContent() *DeleteUserKeyValueNoContent {
	return &DeleteUserKeyValueNoContent{}
}

/*
DeleteUserKeyValueNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteUserKeyValueNoContent struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the delete user key value no content response
func (o *DeleteUserKeyValueNoContent) Code() int {
	return 204
}

// IsSuccess returns true when this delete user key value no content response has a 2xx status code
func (o *DeleteUserKeyValueNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete user key value no content response has a 3xx status code
func (o *DeleteUserKeyValueNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user key value no content response has a 4xx status code
func (o *DeleteUserKeyValueNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete user key value no content response has a 5xx status code
func (o *DeleteUserKeyValueNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user key value no content response a status code equal to that given
func (o *DeleteUserKeyValueNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DeleteUserKeyValueNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tenants/userKeyValue/{keyName}][%d] deleteUserKeyValueNoContent ", 204)
}

func (o *DeleteUserKeyValueNoContent) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tenants/userKeyValue/{keyName}][%d] deleteUserKeyValueNoContent ", 204)
}

func (o *DeleteUserKeyValueNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteUserKeyValueBadRequest creates a DeleteUserKeyValueBadRequest with default headers values
func NewDeleteUserKeyValueBadRequest() *DeleteUserKeyValueBadRequest {
	return &DeleteUserKeyValueBadRequest{}
}

/*
DeleteUserKeyValueBadRequest describes a response with status code 400, with default header values.

Invalid tenantId supplied
*/
type DeleteUserKeyValueBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the delete user key value bad request response
func (o *DeleteUserKeyValueBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this delete user key value bad request response has a 2xx status code
func (o *DeleteUserKeyValueBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete user key value bad request response has a 3xx status code
func (o *DeleteUserKeyValueBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete user key value bad request response has a 4xx status code
func (o *DeleteUserKeyValueBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete user key value bad request response has a 5xx status code
func (o *DeleteUserKeyValueBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete user key value bad request response a status code equal to that given
func (o *DeleteUserKeyValueBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *DeleteUserKeyValueBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tenants/userKeyValue/{keyName}][%d] deleteUserKeyValueBadRequest ", 400)
}

func (o *DeleteUserKeyValueBadRequest) String() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tenants/userKeyValue/{keyName}][%d] deleteUserKeyValueBadRequest ", 400)
}

func (o *DeleteUserKeyValueBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
