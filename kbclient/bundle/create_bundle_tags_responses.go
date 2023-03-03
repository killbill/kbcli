// Code generated by go-swagger; DO NOT EDIT.

package bundle

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

// CreateBundleTagsReader is a Reader for the CreateBundleTags structure.
type CreateBundleTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateBundleTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreateBundleTagsCreated()
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

// NewCreateBundleTagsCreated creates a CreateBundleTagsCreated with default headers values
func NewCreateBundleTagsCreated() *CreateBundleTagsCreated {
	return &CreateBundleTagsCreated{}
}

/*
CreateBundleTagsCreated describes a response with status code 201, with default header values.

Tag created successfully
*/
type CreateBundleTagsCreated struct {
	Payload      []*kbmodel.Tag
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the create bundle tags created response
func (o *CreateBundleTagsCreated) Code() int {
	return 201
}

// IsSuccess returns true when this create bundle tags created response has a 2xx status code
func (o *CreateBundleTagsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create bundle tags created response has a 3xx status code
func (o *CreateBundleTagsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create bundle tags created response has a 4xx status code
func (o *CreateBundleTagsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create bundle tags created response has a 5xx status code
func (o *CreateBundleTagsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create bundle tags created response a status code equal to that given
func (o *CreateBundleTagsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateBundleTagsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/bundles/{bundleId}/tags][%d] createBundleTagsCreated  %+v", 201, o.Payload)
}

func (o *CreateBundleTagsCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/bundles/{bundleId}/tags][%d] createBundleTagsCreated  %+v", 201, o.Payload)
}

func (o *CreateBundleTagsCreated) GetPayload() []*kbmodel.Tag {
	return o.Payload
}

func (o *CreateBundleTagsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateBundleTagsBadRequest creates a CreateBundleTagsBadRequest with default headers values
func NewCreateBundleTagsBadRequest() *CreateBundleTagsBadRequest {
	return &CreateBundleTagsBadRequest{}
}

/*
CreateBundleTagsBadRequest describes a response with status code 400, with default header values.

Invalid bundle id supplied
*/
type CreateBundleTagsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the create bundle tags bad request response
func (o *CreateBundleTagsBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this create bundle tags bad request response has a 2xx status code
func (o *CreateBundleTagsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create bundle tags bad request response has a 3xx status code
func (o *CreateBundleTagsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create bundle tags bad request response has a 4xx status code
func (o *CreateBundleTagsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create bundle tags bad request response has a 5xx status code
func (o *CreateBundleTagsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create bundle tags bad request response a status code equal to that given
func (o *CreateBundleTagsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateBundleTagsBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/bundles/{bundleId}/tags][%d] createBundleTagsBadRequest ", 400)
}

func (o *CreateBundleTagsBadRequest) String() string {
	return fmt.Sprintf("[POST /1.0/kb/bundles/{bundleId}/tags][%d] createBundleTagsBadRequest ", 400)
}

func (o *CreateBundleTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
