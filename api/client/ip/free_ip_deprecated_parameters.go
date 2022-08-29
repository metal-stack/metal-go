// Code generated by go-swagger; DO NOT EDIT.

package ip

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
)

// NewFreeIPDeprecatedParams creates a new FreeIPDeprecatedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFreeIPDeprecatedParams() *FreeIPDeprecatedParams {
	return &FreeIPDeprecatedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFreeIPDeprecatedParamsWithTimeout creates a new FreeIPDeprecatedParams object
// with the ability to set a timeout on a request.
func NewFreeIPDeprecatedParamsWithTimeout(timeout time.Duration) *FreeIPDeprecatedParams {
	return &FreeIPDeprecatedParams{
		timeout: timeout,
	}
}

// NewFreeIPDeprecatedParamsWithContext creates a new FreeIPDeprecatedParams object
// with the ability to set a context for a request.
func NewFreeIPDeprecatedParamsWithContext(ctx context.Context) *FreeIPDeprecatedParams {
	return &FreeIPDeprecatedParams{
		Context: ctx,
	}
}

// NewFreeIPDeprecatedParamsWithHTTPClient creates a new FreeIPDeprecatedParams object
// with the ability to set a custom HTTPClient for a request.
func NewFreeIPDeprecatedParamsWithHTTPClient(client *http.Client) *FreeIPDeprecatedParams {
	return &FreeIPDeprecatedParams{
		HTTPClient: client,
	}
}

/*
FreeIPDeprecatedParams contains all the parameters to send to the API endpoint

	for the free IP deprecated operation.

	Typically these are written to a http.Request.
*/
type FreeIPDeprecatedParams struct {

	/* ID.

	   identifier of the ip
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the free IP deprecated params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FreeIPDeprecatedParams) WithDefaults() *FreeIPDeprecatedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the free IP deprecated params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FreeIPDeprecatedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the free IP deprecated params
func (o *FreeIPDeprecatedParams) WithTimeout(timeout time.Duration) *FreeIPDeprecatedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the free IP deprecated params
func (o *FreeIPDeprecatedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the free IP deprecated params
func (o *FreeIPDeprecatedParams) WithContext(ctx context.Context) *FreeIPDeprecatedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the free IP deprecated params
func (o *FreeIPDeprecatedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the free IP deprecated params
func (o *FreeIPDeprecatedParams) WithHTTPClient(client *http.Client) *FreeIPDeprecatedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the free IP deprecated params
func (o *FreeIPDeprecatedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the free IP deprecated params
func (o *FreeIPDeprecatedParams) WithID(id string) *FreeIPDeprecatedParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the free IP deprecated params
func (o *FreeIPDeprecatedParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FreeIPDeprecatedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
