// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/metal-pod/metal-go/api/models"
)

// FindProjectReader is a Reader for the FindProject structure.
type FindProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFindProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewFindProjectDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindProjectOK creates a FindProjectOK with default headers values
func NewFindProjectOK() *FindProjectOK {
	return &FindProjectOK{}
}

/*FindProjectOK handles this case with default header values.

OK
*/
type FindProjectOK struct {
	Payload *models.V1ProjectResponse
}

func (o *FindProjectOK) Error() string {
	return fmt.Sprintf("[GET /v1/project/{id}][%d] findProjectOK  %+v", 200, o.Payload)
}

func (o *FindProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ProjectResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindProjectDefault creates a FindProjectDefault with default headers values
func NewFindProjectDefault(code int) *FindProjectDefault {
	return &FindProjectDefault{
		_statusCode: code,
	}
}

/*FindProjectDefault handles this case with default header values.

Error
*/
type FindProjectDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the find project default response
func (o *FindProjectDefault) Code() int {
	return o._statusCode
}

func (o *FindProjectDefault) Error() string {
	return fmt.Sprintf("[GET /v1/project/{id}][%d] findProject default  %+v", o._statusCode, o.Payload)
}

func (o *FindProjectDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
