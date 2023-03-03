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

// GetAllTagsReader is a Reader for the GetAllTags structure.
type GetAllTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllTagsOK()
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

// NewGetAllTagsOK creates a GetAllTagsOK with default headers values
func NewGetAllTagsOK() *GetAllTagsOK {
	return &GetAllTagsOK{}
}

/*
GetAllTagsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetAllTagsOK struct {
	Payload      []*kbmodel.Tag
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get all tags o k response
func (o *GetAllTagsOK) Code() int {
	return 200
}

// IsSuccess returns true when this get all tags o k response has a 2xx status code
func (o *GetAllTagsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get all tags o k response has a 3xx status code
func (o *GetAllTagsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all tags o k response has a 4xx status code
func (o *GetAllTagsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get all tags o k response has a 5xx status code
func (o *GetAllTagsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get all tags o k response a status code equal to that given
func (o *GetAllTagsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetAllTagsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/allTags][%d] getAllTagsOK  %+v", 200, o.Payload)
}

func (o *GetAllTagsOK) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/allTags][%d] getAllTagsOK  %+v", 200, o.Payload)
}

func (o *GetAllTagsOK) GetPayload() []*kbmodel.Tag {
	return o.Payload
}

func (o *GetAllTagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllTagsBadRequest creates a GetAllTagsBadRequest with default headers values
func NewGetAllTagsBadRequest() *GetAllTagsBadRequest {
	return &GetAllTagsBadRequest{}
}

/*
GetAllTagsBadRequest describes a response with status code 400, with default header values.

Invalid account id supplied
*/
type GetAllTagsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get all tags bad request response
func (o *GetAllTagsBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this get all tags bad request response has a 2xx status code
func (o *GetAllTagsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all tags bad request response has a 3xx status code
func (o *GetAllTagsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all tags bad request response has a 4xx status code
func (o *GetAllTagsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all tags bad request response has a 5xx status code
func (o *GetAllTagsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get all tags bad request response a status code equal to that given
func (o *GetAllTagsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetAllTagsBadRequest) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/allTags][%d] getAllTagsBadRequest ", 400)
}

func (o *GetAllTagsBadRequest) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/allTags][%d] getAllTagsBadRequest ", 400)
}

func (o *GetAllTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllTagsNotFound creates a GetAllTagsNotFound with default headers values
func NewGetAllTagsNotFound() *GetAllTagsNotFound {
	return &GetAllTagsNotFound{}
}

/*
GetAllTagsNotFound describes a response with status code 404, with default header values.

Account not found
*/
type GetAllTagsNotFound struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the get all tags not found response
func (o *GetAllTagsNotFound) Code() int {
	return 404
}

// IsSuccess returns true when this get all tags not found response has a 2xx status code
func (o *GetAllTagsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get all tags not found response has a 3xx status code
func (o *GetAllTagsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get all tags not found response has a 4xx status code
func (o *GetAllTagsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get all tags not found response has a 5xx status code
func (o *GetAllTagsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get all tags not found response a status code equal to that given
func (o *GetAllTagsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetAllTagsNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/allTags][%d] getAllTagsNotFound ", 404)
}

func (o *GetAllTagsNotFound) String() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/allTags][%d] getAllTagsNotFound ", 404)
}

func (o *GetAllTagsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
