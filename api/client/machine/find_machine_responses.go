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

// FindMachineReader is a Reader for the FindMachine structure.
type FindMachineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindMachineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFindMachineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewFindMachineDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindMachineOK creates a FindMachineOK with default headers values
func NewFindMachineOK() *FindMachineOK {
	return &FindMachineOK{}
}

/*FindMachineOK handles this case with default header values.

OK
*/
type FindMachineOK struct {
	Payload *models.V1MachineResponse
}

func (o *FindMachineOK) Error() string {
	return fmt.Sprintf("[GET /v1/machine/{id}][%d] findMachineOK  %+v", 200, o.Payload)
}

func (o *FindMachineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1MachineResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindMachineDefault creates a FindMachineDefault with default headers values
func NewFindMachineDefault(code int) *FindMachineDefault {
	return &FindMachineDefault{
		_statusCode: code,
	}
}

/*FindMachineDefault handles this case with default header values.

Error
*/
type FindMachineDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the find machine default response
func (o *FindMachineDefault) Code() int {
	return o._statusCode
}

func (o *FindMachineDefault) Error() string {
	return fmt.Sprintf("[GET /v1/machine/{id}][%d] findMachine default  %+v", o._statusCode, o.Payload)
}

func (o *FindMachineDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
