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

// MachineOnReader is a Reader for the MachineOn structure.
type MachineOnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MachineOnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMachineOnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMachineOnDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMachineOnOK creates a MachineOnOK with default headers values
func NewMachineOnOK() *MachineOnOK {
	return &MachineOnOK{}
}

/*
MachineOnOK describes a response with status code 200, with default header values.

OK
*/
type MachineOnOK struct {
	Payload *models.V1MachineResponse
}

// IsSuccess returns true when this machine on o k response has a 2xx status code
func (o *MachineOnOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this machine on o k response has a 3xx status code
func (o *MachineOnOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this machine on o k response has a 4xx status code
func (o *MachineOnOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this machine on o k response has a 5xx status code
func (o *MachineOnOK) IsServerError() bool {
	return false
}

// IsCode returns true when this machine on o k response a status code equal to that given
func (o *MachineOnOK) IsCode(code int) bool {
	return code == 200
}

func (o *MachineOnOK) Error() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/power/on][%d] machineOnOK  %+v", 200, o.Payload)
}

func (o *MachineOnOK) String() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/power/on][%d] machineOnOK  %+v", 200, o.Payload)
}

func (o *MachineOnOK) GetPayload() *models.V1MachineResponse {
	return o.Payload
}

func (o *MachineOnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMachineOnDefault creates a MachineOnDefault with default headers values
func NewMachineOnDefault(code int) *MachineOnDefault {
	return &MachineOnDefault{
		_statusCode: code,
	}
}

/*
MachineOnDefault describes a response with status code -1, with default header values.

Error
*/
type MachineOnDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the machine on default response
func (o *MachineOnDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this machine on default response has a 2xx status code
func (o *MachineOnDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this machine on default response has a 3xx status code
func (o *MachineOnDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this machine on default response has a 4xx status code
func (o *MachineOnDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this machine on default response has a 5xx status code
func (o *MachineOnDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this machine on default response a status code equal to that given
func (o *MachineOnDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *MachineOnDefault) Error() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/power/on][%d] machineOn default  %+v", o._statusCode, o.Payload)
}

func (o *MachineOnDefault) String() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/power/on][%d] machineOn default  %+v", o._statusCode, o.Payload)
}

func (o *MachineOnDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *MachineOnDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
