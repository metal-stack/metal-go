// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/metal-stack/metal-go/api/models"
)

// FreeNetworkReader is a Reader for the FreeNetwork structure.
type FreeNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FreeNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFreeNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewFreeNetworkConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewFreeNetworkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFreeNetworkOK creates a FreeNetworkOK with default headers values
func NewFreeNetworkOK() *FreeNetworkOK {
	return &FreeNetworkOK{}
}

/*FreeNetworkOK handles this case with default header values.

OK
*/
type FreeNetworkOK struct {
	Payload *models.V1NetworkResponse
}

func (o *FreeNetworkOK) Error() string {
	return fmt.Sprintf("[POST /v1/network/free/{id}][%d] freeNetworkOK  %+v", 200, o.Payload)
}

func (o *FreeNetworkOK) GetPayload() *models.V1NetworkResponse {
	return o.Payload
}

func (o *FreeNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1NetworkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFreeNetworkConflict creates a FreeNetworkConflict with default headers values
func NewFreeNetworkConflict() *FreeNetworkConflict {
	return &FreeNetworkConflict{}
}

/*FreeNetworkConflict handles this case with default header values.

Conflict
*/
type FreeNetworkConflict struct {
	Payload *models.HttperrorsHTTPErrorResponse
}

func (o *FreeNetworkConflict) Error() string {
	return fmt.Sprintf("[POST /v1/network/free/{id}][%d] freeNetworkConflict  %+v", 409, o.Payload)
}

func (o *FreeNetworkConflict) GetPayload() *models.HttperrorsHTTPErrorResponse {
	return o.Payload
}

func (o *FreeNetworkConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFreeNetworkDefault creates a FreeNetworkDefault with default headers values
func NewFreeNetworkDefault(code int) *FreeNetworkDefault {
	return &FreeNetworkDefault{
		_statusCode: code,
	}
}

/*FreeNetworkDefault handles this case with default header values.

Error
*/
type FreeNetworkDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the free network default response
func (o *FreeNetworkDefault) Code() int {
	return o._statusCode
}

func (o *FreeNetworkDefault) Error() string {
	return fmt.Sprintf("[POST /v1/network/free/{id}][%d] freeNetwork default  %+v", o._statusCode, o.Payload)
}

func (o *FreeNetworkDefault) GetPayload() *models.HttperrorsHTTPErrorResponse {
	return o.Payload
}

func (o *FreeNetworkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
