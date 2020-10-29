// Code generated by go-swagger; DO NOT EDIT.

package network

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

// AllocateNetworkReader is a Reader for the AllocateNetwork structure.
type AllocateNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AllocateNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewAllocateNetworkCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 409:
		result := NewAllocateNetworkConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewAllocateNetworkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAllocateNetworkCreated creates a AllocateNetworkCreated with default headers values
func NewAllocateNetworkCreated() *AllocateNetworkCreated {
	return &AllocateNetworkCreated{}
}

/*AllocateNetworkCreated handles this case with default header values.

Created
*/
type AllocateNetworkCreated struct {
	Payload *models.V1NetworkResponse
}

func (o *AllocateNetworkCreated) Error() string {
	return fmt.Sprintf("[POST /v1/network/allocate][%d] allocateNetworkCreated  %+v", 201, o.Payload)
}

func (o *AllocateNetworkCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1NetworkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllocateNetworkConflict creates a AllocateNetworkConflict with default headers values
func NewAllocateNetworkConflict() *AllocateNetworkConflict {
	return &AllocateNetworkConflict{}
}

/*AllocateNetworkConflict handles this case with default header values.

Conflict
*/
type AllocateNetworkConflict struct {
	Payload *httperrors.HTTPErrorResponse
}

func (o *AllocateNetworkConflict) Error() string {
	return fmt.Sprintf("[POST /v1/network/allocate][%d] allocateNetworkConflict  %+v", 409, o.Payload)
}

func (o *AllocateNetworkConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAllocateNetworkDefault creates a AllocateNetworkDefault with default headers values
func NewAllocateNetworkDefault(code int) *AllocateNetworkDefault {
	return &AllocateNetworkDefault{
		_statusCode: code,
	}
}

/*AllocateNetworkDefault handles this case with default header values.

Error
*/
type AllocateNetworkDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the allocate network default response
func (o *AllocateNetworkDefault) Code() int {
	return o._statusCode
}

func (o *AllocateNetworkDefault) Error() string {
	return fmt.Sprintf("[POST /v1/network/allocate][%d] allocateNetwork default  %+v", o._statusCode, o.Payload)
}

func (o *AllocateNetworkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
