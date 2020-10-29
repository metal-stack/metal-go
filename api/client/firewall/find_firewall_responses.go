// Code generated by go-swagger; DO NOT EDIT.

package firewall

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

// FindFirewallReader is a Reader for the FindFirewall structure.
type FindFirewallReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FindFirewallReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewFindFirewallOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewFindFirewallDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFindFirewallOK creates a FindFirewallOK with default headers values
func NewFindFirewallOK() *FindFirewallOK {
	return &FindFirewallOK{}
}

/*FindFirewallOK handles this case with default header values.

OK
*/
type FindFirewallOK struct {
	Payload *models.V1FirewallResponse
}

func (o *FindFirewallOK) Error() string {
	return fmt.Sprintf("[GET /v1/firewall/{id}][%d] findFirewallOK  %+v", 200, o.Payload)
}

func (o *FindFirewallOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1FirewallResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFindFirewallDefault creates a FindFirewallDefault with default headers values
func NewFindFirewallDefault(code int) *FindFirewallDefault {
	return &FindFirewallDefault{
		_statusCode: code,
	}
}

/*FindFirewallDefault handles this case with default header values.

Error
*/
type FindFirewallDefault struct {
	_statusCode int

	Payload *httperrors.HTTPErrorResponse
}

// Code gets the status code for the find firewall default response
func (o *FindFirewallDefault) Code() int {
	return o._statusCode
}

func (o *FindFirewallDefault) Error() string {
	return fmt.Sprintf("[GET /v1/firewall/{id}][%d] findFirewall default  %+v", o._statusCode, o.Payload)
}

func (o *FindFirewallDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(httperrors.HTTPErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
