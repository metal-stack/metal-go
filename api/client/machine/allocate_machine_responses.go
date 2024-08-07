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

// AllocateMachineReader is a Reader for the AllocateMachine structure.
type AllocateMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllocateMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAllocateMachineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAllocateMachineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAllocateMachineOK creates a AllocateMachineOK with default headers values
func NewAllocateMachineOK() *AllocateMachineOK {
	return &AllocateMachineOK{}
}

/*
AllocateMachineOK describes a response with status code 200, with default header values.

OK
*/
type AllocateMachineOK struct {
	Payload *models.V1MachineResponse
}

// IsSuccess returns true when this allocate machine o k response has a 2xx status code
func (o *AllocateMachineOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this allocate machine o k response has a 3xx status code
func (o *AllocateMachineOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this allocate machine o k response has a 4xx status code
func (o *AllocateMachineOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this allocate machine o k response has a 5xx status code
func (o *AllocateMachineOK) IsServerError() bool {
	return false
}

// IsCode returns true when this allocate machine o k response a status code equal to that given
func (o *AllocateMachineOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the allocate machine o k response
func (o *AllocateMachineOK) Code() int {
	return 200
}

func (o *AllocateMachineOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/machine/allocate][%d] allocateMachineOK %s", 200, payload)
}

func (o *AllocateMachineOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/machine/allocate][%d] allocateMachineOK %s", 200, payload)
}

func (o *AllocateMachineOK) GetPayload() *models.V1MachineResponse {
	return o.Payload
}

func (o *AllocateMachineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllocateMachineDefault creates a AllocateMachineDefault with default headers values
func NewAllocateMachineDefault(code int) *AllocateMachineDefault {
	return &AllocateMachineDefault{
		_statusCode: code,
	}
}

/*
AllocateMachineDefault describes a response with status code -1, with default header values.

Error
*/
type AllocateMachineDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// IsSuccess returns true when this allocate machine default response has a 2xx status code
func (o *AllocateMachineDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this allocate machine default response has a 3xx status code
func (o *AllocateMachineDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this allocate machine default response has a 4xx status code
func (o *AllocateMachineDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this allocate machine default response has a 5xx status code
func (o *AllocateMachineDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this allocate machine default response a status code equal to that given
func (o *AllocateMachineDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the allocate machine default response
func (o *AllocateMachineDefault) Code() int {
	return o._statusCode
}

func (o *AllocateMachineDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/machine/allocate][%d] allocateMachine default %s", o._statusCode, payload)
}

func (o *AllocateMachineDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /v1/machine/allocate][%d] allocateMachine default %s", o._statusCode, payload)
}

func (o *AllocateMachineDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *AllocateMachineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
