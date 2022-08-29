// Code generated by go-swagger; DO NOT EDIT.

package firmware

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/metal-stack/metal-lib/httperrors"
)

// RemoveFirmwareReader is a Reader for the RemoveFirmware structure.
type RemoveFirmwareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveFirmwareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveFirmwareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRemoveFirmwareDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRemoveFirmwareOK creates a RemoveFirmwareOK with default headers values
func NewRemoveFirmwareOK() *RemoveFirmwareOK {
	return &RemoveFirmwareOK{}
}

/*
	RemoveFirmwareOK describes a response with status code 200, with default header values.

OK
*/
type RemoveFirmwareOK struct {
}

func (o *RemoveFirmwareOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/firmware/{kind}/{vendor}/{board}/{revision}][%d] removeFirmwareOK ", 200)
}

func (o *RemoveFirmwareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRemoveFirmwareDefault creates a RemoveFirmwareDefault with default headers values
func NewRemoveFirmwareDefault(code int) *RemoveFirmwareDefault {
	return &RemoveFirmwareDefault{
		_statusCode: code,
	}
}

/*
	RemoveFirmwareDefault describes a response with status code -1, with default header values.

Error
*/
type RemoveFirmwareDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the remove firmware default response
func (o *RemoveFirmwareDefault) Code() int {
	return o._statusCode
}

func (o *RemoveFirmwareDefault) Error() string {
	return fmt.Sprintf("[DELETE /v1/firmware/{kind}/{vendor}/{board}/{revision}][%d] removeFirmware default  %+v", o._statusCode, o.Payload)
}
func (o *RemoveFirmwareDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *RemoveFirmwareDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
