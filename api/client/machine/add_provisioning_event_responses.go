// Code generated by go-swagger; DO NOT EDIT.

package machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/metal-stack/metal-go/api/models"
	httperrors "github.com/metal-stack/metal-lib/httperrors"
)

// AddProvisioningEventReader is a Reader for the AddProvisioningEvent structure.
type AddProvisioningEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddProvisioningEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddProvisioningEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewAddProvisioningEventDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddProvisioningEventOK creates a AddProvisioningEventOK with default headers values
func NewAddProvisioningEventOK() *AddProvisioningEventOK {
	return &AddProvisioningEventOK{}
}

/*AddProvisioningEventOK handles this case with default header values.

OK
*/
type AddProvisioningEventOK struct {
	Payload *models.V1MachineRecentProvisioningEvents
}

func (o *AddProvisioningEventOK) Error() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/event][%d] addProvisioningEventOK  %+v", 200, o.Payload)
}

func (o *AddProvisioningEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineRecentProvisioningEvents)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProvisioningEventDefault creates a AddProvisioningEventDefault with default headers values
func NewAddProvisioningEventDefault(code int) *AddProvisioningEventDefault {
	return &AddProvisioningEventDefault{
		_statusCode: code,
	}
}

/*AddProvisioningEventDefault handles this case with default header values.

Error
*/
type AddProvisioningEventDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the add provisioning event default response
func (o *AddProvisioningEventDefault) Code() int {
	return o._statusCode
}

func (o *AddProvisioningEventDefault) Error() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/event][%d] addProvisioningEvent default  %+v", o._statusCode, o.Payload)
}

func (o *AddProvisioningEventDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
