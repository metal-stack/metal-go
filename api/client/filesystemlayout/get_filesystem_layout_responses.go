// Code generated by go-swagger; DO NOT EDIT.

package filesystemlayout

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

// GetFilesystemLayoutReader is a Reader for the GetFilesystemLayout structure.
type GetFilesystemLayoutReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFilesystemLayoutReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFilesystemLayoutOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetFilesystemLayoutDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetFilesystemLayoutOK creates a GetFilesystemLayoutOK with default headers values
func NewGetFilesystemLayoutOK() *GetFilesystemLayoutOK {
	return &GetFilesystemLayoutOK{}
}

/*
	GetFilesystemLayoutOK describes a response with status code 200, with default header values.

OK
*/
type GetFilesystemLayoutOK struct {
	Payload *models.V1FilesystemLayoutResponse
}

func (o *GetFilesystemLayoutOK) Error() string {
	return fmt.Sprintf("[GET /v1/filesystemlayout/{id}][%d] getFilesystemLayoutOK  %+v", 200, o.Payload)
}
func (o *GetFilesystemLayoutOK) GetPayload() *models.V1FilesystemLayoutResponse {
	return o.Payload
}

func (o *GetFilesystemLayoutOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1FilesystemLayoutResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilesystemLayoutDefault creates a GetFilesystemLayoutDefault with default headers values
func NewGetFilesystemLayoutDefault(code int) *GetFilesystemLayoutDefault {
	return &GetFilesystemLayoutDefault{
		_statusCode: code,
	}
}

/*
	GetFilesystemLayoutDefault describes a response with status code -1, with default header values.

Error
*/
type GetFilesystemLayoutDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the get filesystem layout default response
func (o *GetFilesystemLayoutDefault) Code() int {
	return o._statusCode
}

func (o *GetFilesystemLayoutDefault) Error() string {
	return fmt.Sprintf("[GET /v1/filesystemlayout/{id}][%d] getFilesystemLayout default  %+v", o._statusCode, o.Payload)
}
func (o *GetFilesystemLayoutDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *GetFilesystemLayoutDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
