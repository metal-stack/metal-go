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
)

// NewFindSizeImageConstraintsParams creates a new FindSizeImageConstraintsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindSizeImageConstraintsParams() *FindSizeImageConstraintsParams {
	return &FindSizeImageConstraintsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindSizeImageConstraintsParamsWithTimeout creates a new FindSizeImageConstraintsParams object
// with the ability to set a timeout on a request.
func NewFindSizeImageConstraintsParamsWithTimeout(timeout time.Duration) *FindSizeImageConstraintsParams {
	return &FindSizeImageConstraintsParams{
		timeout: timeout,
	}
}

// NewFindSizeImageConstraintsParamsWithContext creates a new FindSizeImageConstraintsParams object
// with the ability to set a context for a request.
func NewFindSizeImageConstraintsParamsWithContext(ctx context.Context) *FindSizeImageConstraintsParams {
	return &FindSizeImageConstraintsParams{
		Context: ctx,
	}
}

// NewFindSizeImageConstraintsParamsWithHTTPClient creates a new FindSizeImageConstraintsParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindSizeImageConstraintsParamsWithHTTPClient(client *http.Client) *FindSizeImageConstraintsParams {
	return &FindSizeImageConstraintsParams{
		HTTPClient: client,
	}
}

/* FindSizeImageConstraintsParams contains all the parameters to send to the API endpoint
   for the find size image constraints operation.

   Typically these are written to a http.Request.
*/
type FindSizeImageConstraintsParams struct {

	/* ID.

	   identifier of the sizeimageconstraint
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find size image constraints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindSizeImageConstraintsParams) WithDefaults() *FindSizeImageConstraintsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find size image constraints params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindSizeImageConstraintsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find size image constraints params
func (o *FindSizeImageConstraintsParams) WithTimeout(timeout time.Duration) *FindSizeImageConstraintsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find size image constraints params
func (o *FindSizeImageConstraintsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find size image constraints params
func (o *FindSizeImageConstraintsParams) WithContext(ctx context.Context) *FindSizeImageConstraintsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find size image constraints params
func (o *FindSizeImageConstraintsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find size image constraints params
func (o *FindSizeImageConstraintsParams) WithHTTPClient(client *http.Client) *FindSizeImageConstraintsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find size image constraints params
func (o *FindSizeImageConstraintsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the find size image constraints params
func (o *FindSizeImageConstraintsParams) WithID(id string) *FindSizeImageConstraintsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the find size image constraints params
func (o *FindSizeImageConstraintsParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FindSizeImageConstraintsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
