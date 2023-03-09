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

// SetMachineStateReader is a Reader for the SetMachineState structure.
type SetMachineStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SetMachineStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSetMachineStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSetMachineStateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSetMachineStateOK creates a SetMachineStateOK with default headers values
func NewSetMachineStateOK() *SetMachineStateOK {
	return &SetMachineStateOK{}
}

/*
SetMachineStateOK describes a response with status code 200, with default header values.

OK
*/
type SetMachineStateOK struct {
	Payload *models.V1MachineResponse
}

// IsSuccess returns true when this set machine state o k response has a 2xx status code
func (o *SetMachineStateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this set machine state o k response has a 3xx status code
func (o *SetMachineStateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this set machine state o k response has a 4xx status code
func (o *SetMachineStateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this set machine state o k response has a 5xx status code
func (o *SetMachineStateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this set machine state o k response a status code equal to that given
func (o *SetMachineStateOK) IsCode(code int) bool {
	return code == 200
}

func (o *SetMachineStateOK) Error() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/state][%d] setMachineStateOK  %+v", 200, o.Payload)
}

func (o *SetMachineStateOK) String() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/state][%d] setMachineStateOK  %+v", 200, o.Payload)
}

func (o *SetMachineStateOK) GetPayload() *models.V1MachineResponse {
	return o.Payload
}

func (o *SetMachineStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSetMachineStateDefault creates a SetMachineStateDefault with default headers values
func NewSetMachineStateDefault(code int) *SetMachineStateDefault {
	return &SetMachineStateDefault{
		_statusCode: code,
	}
}

/*
SetMachineStateDefault describes a response with status code -1, with default header values.

Error
*/
type SetMachineStateDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the set machine state default response
func (o *SetMachineStateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this set machine state default response has a 2xx status code
func (o *SetMachineStateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this set machine state default response has a 3xx status code
func (o *SetMachineStateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this set machine state default response has a 4xx status code
func (o *SetMachineStateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this set machine state default response has a 5xx status code
func (o *SetMachineStateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this set machine state default response a status code equal to that given
func (o *SetMachineStateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SetMachineStateDefault) Error() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/state][%d] setMachineState default  %+v", o._statusCode, o.Payload)
}

func (o *SetMachineStateDefault) String() string {
	return fmt.Sprintf("[POST /v1/machine/{id}/state][%d] setMachineState default  %+v", o._statusCode, o.Payload)
}

func (o *SetMachineStateDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *SetMachineStateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
