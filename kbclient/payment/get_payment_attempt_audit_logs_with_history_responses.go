// Code generated by go-swagger; DO NOT EDIT.

package payment

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

// GetPaymentAttemptAuditLogsWithHistoryReader is a Reader for the GetPaymentAttemptAuditLogsWithHistory structure.
type GetPaymentAttemptAuditLogsWithHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPaymentAttemptAuditLogsWithHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPaymentAttemptAuditLogsWithHistoryOK()
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

// NewGetPaymentAttemptAuditLogsWithHistoryOK creates a GetPaymentAttemptAuditLogsWithHistoryOK with default headers values
func NewGetPaymentAttemptAuditLogsWithHistoryOK() *GetPaymentAttemptAuditLogsWithHistoryOK {
	return &GetPaymentAttemptAuditLogsWithHistoryOK{}
}

/*GetPaymentAttemptAuditLogsWithHistoryOK handles this case with default header values.

successful operation
*/
type GetPaymentAttemptAuditLogsWithHistoryOK struct {
	Payload []*kbmodel.AuditLog

	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentAttemptAuditLogsWithHistoryOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/payments/attempts/{paymentAttemptId}/auditLogsWithHistory][%d] getPaymentAttemptAuditLogsWithHistoryOK  %+v", 200, o.Payload)
}

func (o *GetPaymentAttemptAuditLogsWithHistoryOK) GetPayload() []*kbmodel.AuditLog {
	return o.Payload
}

func (o *GetPaymentAttemptAuditLogsWithHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPaymentAttemptAuditLogsWithHistoryNotFound creates a GetPaymentAttemptAuditLogsWithHistoryNotFound with default headers values
func NewGetPaymentAttemptAuditLogsWithHistoryNotFound() *GetPaymentAttemptAuditLogsWithHistoryNotFound {
	return &GetPaymentAttemptAuditLogsWithHistoryNotFound{}
}

/*GetPaymentAttemptAuditLogsWithHistoryNotFound handles this case with default header values.

Account not found
*/
type GetPaymentAttemptAuditLogsWithHistoryNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetPaymentAttemptAuditLogsWithHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/payments/attempts/{paymentAttemptId}/auditLogsWithHistory][%d] getPaymentAttemptAuditLogsWithHistoryNotFound ", 404)
}

func (o *GetPaymentAttemptAuditLogsWithHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
