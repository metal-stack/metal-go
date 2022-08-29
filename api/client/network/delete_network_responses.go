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
	"github.com/metal-stack/metal-lib/httperrors"
)

// DeleteNetworkReader is a Reader for the DeleteNetwork structure.
type DeleteNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDeleteNetworkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteNetworkOK creates a DeleteNetworkOK with default headers values
func NewDeleteNetworkOK() *DeleteNetworkOK {
	return &DeleteNetworkOK{}
}

/*
	DeleteNetworkOK describes a response with status code 200, with default header values.

OK
*/
type DeleteNetworkOK struct {
	Payload *models.V1NetworkResponse
}

func (o *DeleteNetworkOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/network/{id}][%d] deleteNetworkOK  %+v", 200, o.Payload)
}
func (o *DeleteNetworkOK) GetPayload() *models.V1NetworkResponse {
	return o.Payload
}

func (o *DeleteNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1NetworkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteNetworkDefault creates a DeleteNetworkDefault with default headers values
func NewDeleteNetworkDefault(code int) *DeleteNetworkDefault {
	return &DeleteNetworkDefault{
		_statusCode: code,
	}
}

/*
	DeleteNetworkDefault describes a response with status code -1, with default header values.

Error
*/
type DeleteNetworkDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the delete network default response
func (o *DeleteNetworkDefault) Code() int {
	return o._statusCode
}

func (o *DeleteNetworkDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/network/{id}][%d] deleteNetwork default  %+v", o._statusCode, o.Payload)
}
func (o *DeleteNetworkDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *DeleteNetworkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
