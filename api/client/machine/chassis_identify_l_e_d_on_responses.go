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

// ChassisIdentifyLEDOnReader is a Reader for the ChassisIdentifyLEDOn structure.
type ChassisIdentifyLEDOnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChassisIdentifyLEDOnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewChassisIdentifyLEDOnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewChassisIdentifyLEDOnDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChassisIdentifyLEDOnOK creates a ChassisIdentifyLEDOnOK with default headers values
func NewChassisIdentifyLEDOnOK() *ChassisIdentifyLEDOnOK {
	return &ChassisIdentifyLEDOnOK{}
}

/*
ChassisIdentifyLEDOnOK describes a response with status code 200, with default header values.

OK
*/
type ChassisIdentifyLEDOnOK struct {
	Payload *models.V1MachineResponse
}

// IsSuccess returns true when this chassis identify l e d on o k response has a 2xx status code
func (o *ChassisIdentifyLEDOnOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this chassis identify l e d on o k response has a 3xx status code
func (o *ChassisIdentifyLEDOnOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this chassis identify l e d on o k response has a 4xx status code
func (o *ChassisIdentifyLEDOnOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this chassis identify l e d on o k response has a 5xx status code
func (o *ChassisIdentifyLEDOnOK) IsServerError() bool {
	return false
}

// IsCode returns true when this chassis identify l e d on o k response a status code equal to that given
func (o *ChassisIdentifyLEDOnOK) IsCode(code int) bool {
	return code == 200
}

func (o *ChassisIdentifyLEDOnOK) Error() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/power/chassis-identify-led-on][%d] chassisIdentifyLEDOnOK  %+v", 200, o.Payload)
}

func (o *ChassisIdentifyLEDOnOK) String() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/power/chassis-identify-led-on][%d] chassisIdentifyLEDOnOK  %+v", 200, o.Payload)
}

func (o *ChassisIdentifyLEDOnOK) GetPayload() *models.V1MachineResponse {
	return o.Payload
}

func (o *ChassisIdentifyLEDOnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChassisIdentifyLEDOnDefault creates a ChassisIdentifyLEDOnDefault with default headers values
func NewChassisIdentifyLEDOnDefault(code int) *ChassisIdentifyLEDOnDefault {
	return &ChassisIdentifyLEDOnDefault{
		_statusCode: code,
	}
}

/*
ChassisIdentifyLEDOnDefault describes a response with status code -1, with default header values.

Error
*/
type ChassisIdentifyLEDOnDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the chassis identify l e d on default response
func (o *ChassisIdentifyLEDOnDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this chassis identify l e d on default response has a 2xx status code
func (o *ChassisIdentifyLEDOnDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this chassis identify l e d on default response has a 3xx status code
func (o *ChassisIdentifyLEDOnDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this chassis identify l e d on default response has a 4xx status code
func (o *ChassisIdentifyLEDOnDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this chassis identify l e d on default response has a 5xx status code
func (o *ChassisIdentifyLEDOnDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this chassis identify l e d on default response a status code equal to that given
func (o *ChassisIdentifyLEDOnDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ChassisIdentifyLEDOnDefault) Error() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/power/chassis-identify-led-on][%d] chassisIdentifyLEDOn default  %+v", o._statusCode, o.Payload)
}

func (o *ChassisIdentifyLEDOnDefault) String() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/power/chassis-identify-led-on][%d] chassisIdentifyLEDOn default  %+v", o._statusCode, o.Payload)
}

func (o *ChassisIdentifyLEDOnDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *ChassisIdentifyLEDOnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
