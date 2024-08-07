// Code generated by go-swagger; DO NOT EDIT.

package machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/metal-stack/metal-go/api/models"
	"github.com/metal-stack/metal-lib/httperrors"
)

// MachineResetReader is a Reader for the MachineReset structure.
type MachineResetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MachineResetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMachineResetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMachineResetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMachineResetOK creates a MachineResetOK with default headers values
func NewMachineResetOK() *MachineResetOK {
	return &MachineResetOK{}
}

/*
MachineResetOK describes a response with status code 200, with default header values.

OK
*/
type MachineResetOK struct {
	Payload *models.V1MachineResponse
}

// IsSuccess returns true when this machine reset o k response has a 2xx status code
func (o *MachineResetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this machine reset o k response has a 3xx status code
func (o *MachineResetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this machine reset o k response has a 4xx status code
func (o *MachineResetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this machine reset o k response has a 5xx status code
func (o *MachineResetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this machine reset o k response a status code equal to that given
func (o *MachineResetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the machine reset o k response
func (o *MachineResetOK) Code() int {
	return 200
}

func (o *MachineResetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/machine/{id}/power/reset][%d] machineResetOK %s", 200, payload)
}

func (o *MachineResetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/machine/{id}/power/reset][%d] machineResetOK %s", 200, payload)
}

func (o *MachineResetOK) GetPayload() *models.V1MachineResponse {
	return o.Payload
}

func (o *MachineResetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMachineResetDefault creates a MachineResetDefault with default headers values
func NewMachineResetDefault(code int) *MachineResetDefault {
	return &MachineResetDefault{
		_statusCode: code,
	}
}

/*
MachineResetDefault describes a response with status code -1, with default header values.

Error
*/
type MachineResetDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this machine reset default response has a 2xx status code
func (o *MachineResetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this machine reset default response has a 3xx status code
func (o *MachineResetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this machine reset default response has a 4xx status code
func (o *MachineResetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this machine reset default response has a 5xx status code
func (o *MachineResetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this machine reset default response a status code equal to that given
func (o *MachineResetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the machine reset default response
func (o *MachineResetDefault) Code() int {
	return o._statusCode
}

func (o *MachineResetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/machine/{id}/power/reset][%d] machineReset default %s", o._statusCode, payload)
}

func (o *MachineResetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/machine/{id}/power/reset][%d] machineReset default %s", o._statusCode, payload)
}

func (o *MachineResetDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *MachineResetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
