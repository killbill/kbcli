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

// GetAccountEmailAuditLogsWithHistoryReader is a Reader for the GetAccountEmailAuditLogsWithHistory structure.
type GetAccountEmailAuditLogsWithHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAccountEmailAuditLogsWithHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAccountEmailAuditLogsWithHistoryOK()
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

// NewGetAccountEmailAuditLogsWithHistoryOK creates a GetAccountEmailAuditLogsWithHistoryOK with default headers values
func NewGetAccountEmailAuditLogsWithHistoryOK() *GetAccountEmailAuditLogsWithHistoryOK {
	return &GetAccountEmailAuditLogsWithHistoryOK{}
}

/*GetAccountEmailAuditLogsWithHistoryOK handles this case with default header values.

successful operation
*/
type GetAccountEmailAuditLogsWithHistoryOK struct {
	Payload []*kbmodel.AuditLog

	HttpResponse runtime.ClientResponse
}

func (o *GetAccountEmailAuditLogsWithHistoryOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/emails/{accountEmailId}/auditLogsWithHistory][%d] getAccountEmailAuditLogsWithHistoryOK  %+v", 200, o.Payload)
}

func (o *GetAccountEmailAuditLogsWithHistoryOK) GetPayload() []*kbmodel.AuditLog {
	return o.Payload
}

func (o *GetAccountEmailAuditLogsWithHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAccountEmailAuditLogsWithHistoryNotFound creates a GetAccountEmailAuditLogsWithHistoryNotFound with default headers values
func NewGetAccountEmailAuditLogsWithHistoryNotFound() *GetAccountEmailAuditLogsWithHistoryNotFound {
	return &GetAccountEmailAuditLogsWithHistoryNotFound{}
}

/*GetAccountEmailAuditLogsWithHistoryNotFound handles this case with default header values.

Account not found
*/
type GetAccountEmailAuditLogsWithHistoryNotFound struct {
	HttpResponse runtime.ClientResponse
}

func (o *GetAccountEmailAuditLogsWithHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/{accountId}/emails/{accountEmailId}/auditLogsWithHistory][%d] getAccountEmailAuditLogsWithHistoryNotFound ", 404)
}

func (o *GetAccountEmailAuditLogsWithHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
