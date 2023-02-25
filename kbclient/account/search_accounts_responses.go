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

// SearchAccountsReader is a Reader for the SearchAccounts structure.
type SearchAccountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchAccountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSearchAccountsOK()
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

// NewSearchAccountsOK creates a SearchAccountsOK with default headers values
func NewSearchAccountsOK() *SearchAccountsOK {
	return &SearchAccountsOK{}
}

/*SearchAccountsOK handles this case with default header values.

successful operation
*/
type SearchAccountsOK struct {
	Payload []*kbmodel.Account

	HttpResponse runtime.ClientResponse
}

func (o *SearchAccountsOK) Error() string {
	return fmt.Sprintf("[GET /1.0/kb/accounts/search/{searchKey}][%d] searchAccountsOK  %+v", 200, o.Payload)
}

func (o *SearchAccountsOK) GetPayload() []*kbmodel.Account {
	return o.Payload
}

func (o *SearchAccountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
