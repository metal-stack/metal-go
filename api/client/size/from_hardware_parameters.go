// Code generated by go-swagger; DO NOT EDIT.

package size

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/metal-stack/metal-go/api/models"
)

// NewFromHardwareParams creates a new FromHardwareParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFromHardwareParams() *FromHardwareParams {
	return &FromHardwareParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFromHardwareParamsWithTimeout creates a new FromHardwareParams object
// with the ability to set a timeout on a request.
func NewFromHardwareParamsWithTimeout(timeout time.Duration) *FromHardwareParams {
	return &FromHardwareParams{
		timeout: timeout,
	}
}

// NewFromHardwareParamsWithContext creates a new FromHardwareParams object
// with the ability to set a context for a request.
func NewFromHardwareParamsWithContext(ctx context.Context) *FromHardwareParams {
	return &FromHardwareParams{
		Context: ctx,
	}
}

// NewFromHardwareParamsWithHTTPClient creates a new FromHardwareParams object
// with the ability to set a custom HTTPClient for a request.
func NewFromHardwareParamsWithHTTPClient(client *http.Client) *FromHardwareParams {
	return &FromHardwareParams{
		HTTPClient: client,
	}
}

/*
FromHardwareParams contains all the parameters to send to the API endpoint

	for the from hardware operation.

	Typically these are written to a http.Request.
*/
type FromHardwareParams struct {

	// Body.
	Body *models.V1MachineHardware

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the from hardware params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FromHardwareParams) WithDefaults() *FromHardwareParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the from hardware params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FromHardwareParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the from hardware params
func (o *FromHardwareParams) WithTimeout(timeout time.Duration) *FromHardwareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the from hardware params
func (o *FromHardwareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the from hardware params
func (o *FromHardwareParams) WithContext(ctx context.Context) *FromHardwareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the from hardware params
func (o *FromHardwareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the from hardware params
func (o *FromHardwareParams) WithHTTPClient(client *http.Client) *FromHardwareParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the from hardware params
func (o *FromHardwareParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the from hardware params
func (o *FromHardwareParams) WithBody(body *models.V1MachineHardware) *FromHardwareParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the from hardware params
func (o *FromHardwareParams) SetBody(body *models.V1MachineHardware) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *FromHardwareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
