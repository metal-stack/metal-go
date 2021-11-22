// Code generated by go-swagger; DO NOT EDIT.

package firewall

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/metal-stack/metal-lib/httperrors"
)

// TryAllocateFirewallReader is a Reader for the TryAllocateFirewall structure.
type TryAllocateFirewallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TryAllocateFirewallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTryAllocateFirewallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTryAllocateFirewallDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTryAllocateFirewallOK creates a TryAllocateFirewallOK with default headers values
func NewTryAllocateFirewallOK() *TryAllocateFirewallOK {
	return &TryAllocateFirewallOK{}
}

/* TryAllocateFirewallOK describes a response with status code 200, with default header values.

OK
*/
type TryAllocateFirewallOK struct {
}

func (o *TryAllocateFirewallOK) Error() string {
	return fmt.Sprintf("[POST /v1/firewall/tryallocate][%d] tryAllocateFirewallOK ", 200)
}

func (o *TryAllocateFirewallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTryAllocateFirewallDefault creates a TryAllocateFirewallDefault with default headers values
func NewTryAllocateFirewallDefault(code int) *TryAllocateFirewallDefault {
	return &TryAllocateFirewallDefault{
		_statusCode: code,
	}
}

/* TryAllocateFirewallDefault describes a response with status code -1, with default header values.

Error
*/
type TryAllocateFirewallDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the try allocate firewall default response
func (o *TryAllocateFirewallDefault) Code() int {
	return o._statusCode
}

func (o *TryAllocateFirewallDefault) Error() string {
	return fmt.Sprintf("[POST /v1/firewall/tryallocate][%d] tryAllocateFirewall default  %+v", o._statusCode, o.Payload)
}
func (o *TryAllocateFirewallDefault) GetPayload() *httperrors.HTTPErrorResponse {
	return o.Payload
}

func (o *TryAllocateFirewallDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
