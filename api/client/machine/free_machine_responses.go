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

// FreeMachineReader is a Reader for the FreeMachine structure.
type FreeMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FreeMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFreeMachineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFreeMachineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFreeMachineOK creates a FreeMachineOK with default headers values
func NewFreeMachineOK() *FreeMachineOK {
	return &FreeMachineOK{}
}

/*FreeMachineOK handles this case with default header values.

OK
*/
type FreeMachineOK struct {
	Payload *models.V1MachineResponse
}

func (o *FreeMachineOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/machine/{id}/free][%d] freeMachineOK  %+v", 200, o.Payload)
}

func (o *FreeMachineOK) GetPayload() *models.V1MachineResponse {
	return o.Payload
}

func (o *FreeMachineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFreeMachineDefault creates a FreeMachineDefault with default headers values
func NewFreeMachineDefault(code int) *FreeMachineDefault {
	return &FreeMachineDefault{
		_statusCode: code,
	}
}

/*FreeMachineDefault handles this case with default header values.

Error
*/
type FreeMachineDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the free machine default response
func (o *FreeMachineDefault) Code() int {
	return o._statusCode
}

func (o *FreeMachineDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/machine/{id}/free][%d] freeMachine default  %+v", o._statusCode, o.Payload)
}

func (o *FreeMachineDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *FreeMachineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
