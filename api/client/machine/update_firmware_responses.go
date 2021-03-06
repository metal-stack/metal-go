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

// UpdateFirmwareReader is a Reader for the UpdateFirmware structure.
type UpdateFirmwareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateFirmwareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateFirmwareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUpdateFirmwareDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateFirmwareOK creates a UpdateFirmwareOK with default headers values
func NewUpdateFirmwareOK() *UpdateFirmwareOK {
	return &UpdateFirmwareOK{}
}

/* UpdateFirmwareOK describes a response with status code 200, with default header values.

OK
*/
type UpdateFirmwareOK struct {
	Payload *models.V1MachineResponse
}

func (o *UpdateFirmwareOK) Error() string {
	return fmt.Sprintf("[POST /v1/machine/update-firmware/{id}][%d] updateFirmwareOK  %+v", 200, o.Payload)
}
func (o *UpdateFirmwareOK) GetPayload() *models.V1MachineResponse {
	return o.Payload
}

func (o *UpdateFirmwareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateFirmwareDefault creates a UpdateFirmwareDefault with default headers values
func NewUpdateFirmwareDefault(code int) *UpdateFirmwareDefault {
	return &UpdateFirmwareDefault{
		_statusCode: code,
	}
}

/* UpdateFirmwareDefault describes a response with status code -1, with default header values.

Error
*/
type UpdateFirmwareDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the update firmware default response
func (o *UpdateFirmwareDefault) Code() int {
	return o._statusCode
}

func (o *UpdateFirmwareDefault) Error() string {
	return fmt.Sprintf("[POST /v1/machine/update-firmware/{id}][%d] updateFirmware default  %+v", o._statusCode, o.Payload)
}
func (o *UpdateFirmwareDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *UpdateFirmwareDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
