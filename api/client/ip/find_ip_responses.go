// Code generated by go-swagger; DO NOT EDIT.

package ip

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

// FindIPReader is a Reader for the FindIP structure.
type FindIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFindIPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFindIPDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindIPOK creates a FindIPOK with default headers values
func NewFindIPOK() *FindIPOK {
	return &FindIPOK{}
}

/*
	FindIPOK describes a response with status code 200, with default header values.

OK
*/
type FindIPOK struct {
	Payload *models.V1IPResponse
}

func (o *FindIPOK) Error() string {
	return fmt.Sprintf("[GET /v1/ip/{id}][%d] findIpOK  %+v", 200, o.Payload)
}
func (o *FindIPOK) GetPayload() *models.V1IPResponse {
	return o.Payload
}

func (o *FindIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1IPResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindIPDefault creates a FindIPDefault with default headers values
func NewFindIPDefault(code int) *FindIPDefault {
	return &FindIPDefault{
		_statusCode: code,
	}
}

/*
	FindIPDefault describes a response with status code -1, with default header values.

Error
*/
type FindIPDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the find IP default response
func (o *FindIPDefault) Code() int {
	return o._statusCode
}

func (o *FindIPDefault) Error() string {
	return fmt.Sprintf("[GET /v1/ip/{id}][%d] findIP default  %+v", o._statusCode, o.Payload)
}
func (o *FindIPDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *FindIPDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
