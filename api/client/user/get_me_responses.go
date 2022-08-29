// Code generated by go-swagger; DO NOT EDIT.

package user

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

// GetMeReader is a Reader for the GetMe structure.
type GetMeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetMeDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetMeOK creates a GetMeOK with default headers values
func NewGetMeOK() *GetMeOK {
	return &GetMeOK{}
}

/*
	GetMeOK describes a response with status code 200, with default header values.

OK
*/
type GetMeOK struct {
	Payload *models.V1User
}

func (o *GetMeOK) Error() string {
	return fmt.Sprintf("[GET /v1/user/me][%d] getMeOK  %+v", 200, o.Payload)
}
func (o *GetMeOK) GetPayload() *models.V1User {
	return o.Payload
}

func (o *GetMeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMeDefault creates a GetMeDefault with default headers values
func NewGetMeDefault(code int) *GetMeDefault {
	return &GetMeDefault{
		_statusCode: code,
	}
}

/*
	GetMeDefault describes a response with status code -1, with default header values.

Error
*/
type GetMeDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the get me default response
func (o *GetMeDefault) Code() int {
	return o._statusCode
}

func (o *GetMeDefault) Error() string {
	return fmt.Sprintf("[GET /v1/user/me][%d] getMe default  %+v", o._statusCode, o.Payload)
}
func (o *GetMeDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *GetMeDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
