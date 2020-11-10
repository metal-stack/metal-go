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
)

// FreeIPReader is a Reader for the FreeIP structure.
type FreeIPReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FreeIPReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFreeIPOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFreeIPDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFreeIPOK creates a FreeIPOK with default headers values
func NewFreeIPOK() *FreeIPOK {
	return &FreeIPOK{}
}

/*FreeIPOK handles this case with default header values.

OK
*/
type FreeIPOK struct {
	Payload *models.V1IPResponse
}

func (o *FreeIPOK) Error() string {
	return fmt.Sprintf("[POST /v1/ip/free/{id}][%d] freeIpOK  %+v", 200, o.Payload)
}

func (o *FreeIPOK) GetPayload() *models.V1IPResponse {
	return o.Payload
}

func (o *FreeIPOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1IPResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFreeIPDefault creates a FreeIPDefault with default headers values
func NewFreeIPDefault(code int) *FreeIPDefault {
	return &FreeIPDefault{
		_statusCode: code,
	}
}

/*FreeIPDefault handles this case with default header values.

Error
*/
type FreeIPDefault struct {
	_statusCode int

	Payload *models.HttperrorsHTTPErrorResponse
}

// Code gets the status code for the free IP default response
func (o *FreeIPDefault) Code() int {
	return o._statusCode
}

func (o *FreeIPDefault) Error() string {
	return fmt.Sprintf("[POST /v1/ip/free/{id}][%d] freeIP default  %+v", o._statusCode, o.Payload)
}

func (o *FreeIPDefault) GetPayload() *models.HttperrorsHTTPErrorResponse {
	return o.Payload
}

func (o *FreeIPDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.HttperrorsHTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
