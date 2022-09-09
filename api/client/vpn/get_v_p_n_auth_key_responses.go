// Code generated by go-swagger; DO NOT EDIT.

package vpn

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

// GetVPNAuthKeyReader is a Reader for the GetVPNAuthKey structure.
type GetVPNAuthKeyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVPNAuthKeyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVPNAuthKeyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetVPNAuthKeyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetVPNAuthKeyOK creates a GetVPNAuthKeyOK with default headers values
func NewGetVPNAuthKeyOK() *GetVPNAuthKeyOK {
	return &GetVPNAuthKeyOK{}
}

/*
GetVPNAuthKeyOK describes a response with status code 200, with default header values.

OK
*/
type GetVPNAuthKeyOK struct {
	Payload *models.V1VPNResponse
}

// IsSuccess returns true when this get v p n auth key o k response has a 2xx status code
func (o *GetVPNAuthKeyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get v p n auth key o k response has a 3xx status code
func (o *GetVPNAuthKeyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get v p n auth key o k response has a 4xx status code
func (o *GetVPNAuthKeyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get v p n auth key o k response has a 5xx status code
func (o *GetVPNAuthKeyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get v p n auth key o k response a status code equal to that given
func (o *GetVPNAuthKeyOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetVPNAuthKeyOK) Error() string {
	return fmt.Sprintf("[POST /v1/vpn/authkey][%d] getVPNAuthKeyOK  %+v", 200, o.Payload)
}

func (o *GetVPNAuthKeyOK) String() string {
	return fmt.Sprintf("[POST /v1/vpn/authkey][%d] getVPNAuthKeyOK  %+v", 200, o.Payload)
}

func (o *GetVPNAuthKeyOK) GetPayload() *models.V1VPNResponse {
	return o.Payload
}

func (o *GetVPNAuthKeyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1VPNResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVPNAuthKeyDefault creates a GetVPNAuthKeyDefault with default headers values
func NewGetVPNAuthKeyDefault(code int) *GetVPNAuthKeyDefault {
	return &GetVPNAuthKeyDefault{
		_statusCode: code,
	}
}

/*
GetVPNAuthKeyDefault describes a response with status code -1, with default header values.

Error
*/
type GetVPNAuthKeyDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the get v p n auth key default response
func (o *GetVPNAuthKeyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this get v p n auth key default response has a 2xx status code
func (o *GetVPNAuthKeyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this get v p n auth key default response has a 3xx status code
func (o *GetVPNAuthKeyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this get v p n auth key default response has a 4xx status code
func (o *GetVPNAuthKeyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this get v p n auth key default response has a 5xx status code
func (o *GetVPNAuthKeyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this get v p n auth key default response a status code equal to that given
func (o *GetVPNAuthKeyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *GetVPNAuthKeyDefault) Error() string {
	return fmt.Sprintf("[POST /v1/vpn/authkey][%d] getVPNAuthKey default  %+v", o._statusCode, o.Payload)
}

func (o *GetVPNAuthKeyDefault) String() string {
	return fmt.Sprintf("[POST /v1/vpn/authkey][%d] getVPNAuthKey default  %+v", o._statusCode, o.Payload)
}

func (o *GetVPNAuthKeyDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *GetVPNAuthKeyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
