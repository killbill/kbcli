// Code generated by go-swagger; DO NOT EDIT.

package tenant

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/killbill/kbcli/v3/kbcommon"

	strfmt "github.com/go-openapi/strfmt"
)

// DeletePerTenantConfigurationReader is a Reader for the DeletePerTenantConfiguration structure.
type DeletePerTenantConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeletePerTenantConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeletePerTenantConfigurationNoContent()
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

// NewDeletePerTenantConfigurationNoContent creates a DeletePerTenantConfigurationNoContent with default headers values
func NewDeletePerTenantConfigurationNoContent() *DeletePerTenantConfigurationNoContent {
	return &DeletePerTenantConfigurationNoContent{}
}

/*DeletePerTenantConfigurationNoContent handles this case with default header values.

Successful operation
*/
type DeletePerTenantConfigurationNoContent struct {
	HttpResponse runtime.ClientResponse
}

func (o *DeletePerTenantConfigurationNoContent) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tenants/uploadPerTenantConfig][%d] deletePerTenantConfigurationNoContent ", 204)
}

func (o *DeletePerTenantConfigurationNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeletePerTenantConfigurationBadRequest creates a DeletePerTenantConfigurationBadRequest with default headers values
func NewDeletePerTenantConfigurationBadRequest() *DeletePerTenantConfigurationBadRequest {
	return &DeletePerTenantConfigurationBadRequest{}
}

/*DeletePerTenantConfigurationBadRequest handles this case with default header values.

Invalid tenantId supplied
*/
type DeletePerTenantConfigurationBadRequest struct {
	HttpResponse runtime.ClientResponse
}

func (o *DeletePerTenantConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /1.0/kb/tenants/uploadPerTenantConfig][%d] deletePerTenantConfigurationBadRequest ", 400)
}

func (o *DeletePerTenantConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
