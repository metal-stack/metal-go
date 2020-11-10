// Code generated by go-swagger; DO NOT EDIT.

package machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/metal-stack/metal-go/api/models"
	"github.com/metal-stack/metal-lib/httperrors"
)

// GetProvisioningEventContainerReader is a Reader for the GetProvisioningEventContainer structure.
type GetProvisioningEventContainerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProvisioningEventContainerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProvisioningEventContainerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetProvisioningEventContainerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetProvisioningEventContainerOK creates a GetProvisioningEventContainerOK with default headers values
func NewGetProvisioningEventContainerOK() *GetProvisioningEventContainerOK {
	return &GetProvisioningEventContainerOK{}
}

/*GetProvisioningEventContainerOK handles this case with default header values.

OK
*/
type GetProvisioningEventContainerOK struct {
	Payload *models.V1MachineRecentProvisioningEvents
}

func (o *GetProvisioningEventContainerOK) Error() string {
	return fmt.Sprintf("[GET /v1/machine/{id}/event][%d] getProvisioningEventContainerOK  %+v", 200, o.Payload)
}

func (o *GetProvisioningEventContainerOK) GetPayload() *models.V1MachineRecentProvisioningEvents {
	return o.Payload
}

func (o *GetProvisioningEventContainerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineRecentProvisioningEvents)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProvisioningEventContainerDefault creates a GetProvisioningEventContainerDefault with default headers values
func NewGetProvisioningEventContainerDefault(code int) *GetProvisioningEventContainerDefault {
	return &GetProvisioningEventContainerDefault{
		_statusCode: code,
	}
}

/*GetProvisioningEventContainerDefault handles this case with default header values.

Error
*/
type GetProvisioningEventContainerDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the get provisioning event container default response
func (o *GetProvisioningEventContainerDefault) Code() int {
	return o._statusCode
}

func (o *GetProvisioningEventContainerDefault) Error() string {
	return fmt.Sprintf("[GET /v1/machine/{id}/event][%d] getProvisioningEventContainer default  %+v", o._statusCode, o.Payload)
}

func (o *GetProvisioningEventContainerDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *GetProvisioningEventContainerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
