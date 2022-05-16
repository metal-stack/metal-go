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

// AddProvisioningEventsReader is a Reader for the AddProvisioningEvents structure.
type AddProvisioningEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddProvisioningEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddProvisioningEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddProvisioningEventsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddProvisioningEventsOK creates a AddProvisioningEventsOK with default headers values
func NewAddProvisioningEventsOK() *AddProvisioningEventsOK {
	return &AddProvisioningEventsOK{}
}

/* AddProvisioningEventsOK describes a response with status code 200, with default header values.

OK
*/
type AddProvisioningEventsOK struct {
	Payload *models.V1MachineRecentProvisioningEventsResponse
}

func (o *AddProvisioningEventsOK) Error() string {
	return fmt.Sprintf("[POST /v1/machine/event][%d] addProvisioningEventsOK  %+v", 200, o.Payload)
}
func (o *AddProvisioningEventsOK) GetPayload() *models.V1MachineRecentProvisioningEventsResponse {
	return o.Payload
}

func (o *AddProvisioningEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineRecentProvisioningEventsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProvisioningEventsDefault creates a AddProvisioningEventsDefault with default headers values
func NewAddProvisioningEventsDefault(code int) *AddProvisioningEventsDefault {
	return &AddProvisioningEventsDefault{
		_statusCode: code,
	}
}

/* AddProvisioningEventsDefault describes a response with status code -1, with default header values.

Error
*/
type AddProvisioningEventsDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the add provisioning events default response
func (o *AddProvisioningEventsDefault) Code() int {
	return o._statusCode
}

func (o *AddProvisioningEventsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/machine/event][%d] addProvisioningEvents default  %+v", o._statusCode, o.Payload)
}
func (o *AddProvisioningEventsDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *AddProvisioningEventsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
