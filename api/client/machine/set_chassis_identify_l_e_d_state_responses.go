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

// SetChassisIdentifyLEDStateReader is a Reader for the SetChassisIdentifyLEDState structure.
type SetChassisIdentifyLEDStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetChassisIdentifyLEDStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetChassisIdentifyLEDStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSetChassisIdentifyLEDStateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSetChassisIdentifyLEDStateOK creates a SetChassisIdentifyLEDStateOK with default headers values
func NewSetChassisIdentifyLEDStateOK() *SetChassisIdentifyLEDStateOK {
	return &SetChassisIdentifyLEDStateOK{}
}

/*
	SetChassisIdentifyLEDStateOK describes a response with status code 200, with default header values.

OK
*/
type SetChassisIdentifyLEDStateOK struct {
	Payload *models.V1MachineResponse
}

func (o *SetChassisIdentifyLEDStateOK) Error() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/chassis-identify-led-state][%d] setChassisIdentifyLEDStateOK  %+v", 200, o.Payload)
}
func (o *SetChassisIdentifyLEDStateOK) GetPayload() *models.V1MachineResponse {
	return o.Payload
}

func (o *SetChassisIdentifyLEDStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetChassisIdentifyLEDStateDefault creates a SetChassisIdentifyLEDStateDefault with default headers values
func NewSetChassisIdentifyLEDStateDefault(code int) *SetChassisIdentifyLEDStateDefault {
	return &SetChassisIdentifyLEDStateDefault{
		_statusCode: code,
	}
}

/*
	SetChassisIdentifyLEDStateDefault describes a response with status code -1, with default header values.

Error
*/
type SetChassisIdentifyLEDStateDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the set chassis identify l e d state default response
func (o *SetChassisIdentifyLEDStateDefault) Code() int {
	return o._statusCode
}

func (o *SetChassisIdentifyLEDStateDefault) Error() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/chassis-identify-led-state][%d] setChassisIdentifyLEDState default  %+v", o._statusCode, o.Payload)
}
func (o *SetChassisIdentifyLEDStateDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *SetChassisIdentifyLEDStateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
