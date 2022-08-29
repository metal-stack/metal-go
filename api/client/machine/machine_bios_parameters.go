// Code generated by go-swagger; DO NOT EDIT.

package machine

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

// NewMachineBiosParams creates a new MachineBiosParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMachineBiosParams() *MachineBiosParams {
	return &MachineBiosParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMachineBiosParamsWithTimeout creates a new MachineBiosParams object
// with the ability to set a timeout on a request.
func NewMachineBiosParamsWithTimeout(timeout time.Duration) *MachineBiosParams {
	return &MachineBiosParams{
		timeout: timeout,
	}
}

// NewMachineBiosParamsWithContext creates a new MachineBiosParams object
// with the ability to set a context for a request.
func NewMachineBiosParamsWithContext(ctx context.Context) *MachineBiosParams {
	return &MachineBiosParams{
		Context: ctx,
	}
}

// NewMachineBiosParamsWithHTTPClient creates a new MachineBiosParams object
// with the ability to set a custom HTTPClient for a request.
func NewMachineBiosParamsWithHTTPClient(client *http.Client) *MachineBiosParams {
	return &MachineBiosParams{
		HTTPClient: client,
	}
}

/*
MachineBiosParams contains all the parameters to send to the API endpoint

	for the machine bios operation.

	Typically these are written to a http.Request.
*/
type MachineBiosParams struct {

	// Body.
	Body models.V1EmptyBody

	/* ID.

	   identifier of the machine
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the machine bios params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MachineBiosParams) WithDefaults() *MachineBiosParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the machine bios params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MachineBiosParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the machine bios params
func (o *MachineBiosParams) WithTimeout(timeout time.Duration) *MachineBiosParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the machine bios params
func (o *MachineBiosParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the machine bios params
func (o *MachineBiosParams) WithContext(ctx context.Context) *MachineBiosParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the machine bios params
func (o *MachineBiosParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the machine bios params
func (o *MachineBiosParams) WithHTTPClient(client *http.Client) *MachineBiosParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the machine bios params
func (o *MachineBiosParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the machine bios params
func (o *MachineBiosParams) WithBody(body models.V1EmptyBody) *MachineBiosParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the machine bios params
func (o *MachineBiosParams) SetBody(body models.V1EmptyBody) {
	o.Body = body
}

// WithID adds the id to the machine bios params
func (o *MachineBiosParams) WithID(id string) *MachineBiosParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the machine bios params
func (o *MachineBiosParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *MachineBiosParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
