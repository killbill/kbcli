// Code generated by go-swagger; DO NOT EDIT.

package payment_transaction

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

// CreateTransactionTagsReader is a Reader for the CreateTransactionTags structure.
type CreateTransactionTagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateTransactionTagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201, 200:
		result := NewCreateTransactionTagsCreated()
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

// NewCreateTransactionTagsCreated creates a CreateTransactionTagsCreated with default headers values
func NewCreateTransactionTagsCreated() *CreateTransactionTagsCreated {
	return &CreateTransactionTagsCreated{}
}

/*
CreateTransactionTagsCreated describes a response with status code 201, with default header values.

Tag created successfully
*/
type CreateTransactionTagsCreated struct {
	Payload      []*kbmodel.Tag
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the create transaction tags created response
func (o *CreateTransactionTagsCreated) Code() int {
	return 201
}

// IsSuccess returns true when this create transaction tags created response has a 2xx status code
func (o *CreateTransactionTagsCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create transaction tags created response has a 3xx status code
func (o *CreateTransactionTagsCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create transaction tags created response has a 4xx status code
func (o *CreateTransactionTagsCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create transaction tags created response has a 5xx status code
func (o *CreateTransactionTagsCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create transaction tags created response a status code equal to that given
func (o *CreateTransactionTagsCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateTransactionTagsCreated) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/paymentTransactions/{transactionId}/tags][%d] createTransactionTagsCreated  %+v", 201, o.Payload)
}

func (o *CreateTransactionTagsCreated) String() string {
	return fmt.Sprintf("[POST /1.0/kb/paymentTransactions/{transactionId}/tags][%d] createTransactionTagsCreated  %+v", 201, o.Payload)
}

func (o *CreateTransactionTagsCreated) GetPayload() []*kbmodel.Tag {
	return o.Payload
}

func (o *CreateTransactionTagsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateTransactionTagsBadRequest creates a CreateTransactionTagsBadRequest with default headers values
func NewCreateTransactionTagsBadRequest() *CreateTransactionTagsBadRequest {
	return &CreateTransactionTagsBadRequest{}
}

/*
CreateTransactionTagsBadRequest describes a response with status code 400, with default header values.

Invalid transaction id supplied
*/
type CreateTransactionTagsBadRequest struct {
	HttpResponse runtime.ClientResponse
}

// Code gets the status code for the create transaction tags bad request response
func (o *CreateTransactionTagsBadRequest) Code() int {
	return 400
}

// IsSuccess returns true when this create transaction tags bad request response has a 2xx status code
func (o *CreateTransactionTagsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create transaction tags bad request response has a 3xx status code
func (o *CreateTransactionTagsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create transaction tags bad request response has a 4xx status code
func (o *CreateTransactionTagsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create transaction tags bad request response has a 5xx status code
func (o *CreateTransactionTagsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create transaction tags bad request response a status code equal to that given
func (o *CreateTransactionTagsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateTransactionTagsBadRequest) Error() string {
	return fmt.Sprintf("[POST /1.0/kb/paymentTransactions/{transactionId}/tags][%d] createTransactionTagsBadRequest ", 400)
}

func (o *CreateTransactionTagsBadRequest) String() string {
	return fmt.Sprintf("[POST /1.0/kb/paymentTransactions/{transactionId}/tags][%d] createTransactionTagsBadRequest ", 400)
}

func (o *CreateTransactionTagsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
