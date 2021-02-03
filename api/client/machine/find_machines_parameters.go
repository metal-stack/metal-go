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

// NewFindMachinesParams creates a new FindMachinesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindMachinesParams() *FindMachinesParams {
	return &FindMachinesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindMachinesParamsWithTimeout creates a new FindMachinesParams object
// with the ability to set a timeout on a request.
func NewFindMachinesParamsWithTimeout(timeout time.Duration) *FindMachinesParams {
	return &FindMachinesParams{
		timeout: timeout,
	}
}

// NewFindMachinesParamsWithContext creates a new FindMachinesParams object
// with the ability to set a context for a request.
func NewFindMachinesParamsWithContext(ctx context.Context) *FindMachinesParams {
	return &FindMachinesParams{
		Context: ctx,
	}
}

// NewFindMachinesParamsWithHTTPClient creates a new FindMachinesParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindMachinesParamsWithHTTPClient(client *http.Client) *FindMachinesParams {
	return &FindMachinesParams{
		HTTPClient: client,
	}
}

/* FindMachinesParams contains all the parameters to send to the API endpoint
   for the find machines operation.

   Typically these are written to a http.Request.
*/
type FindMachinesParams struct {

	// Body.
	Body *models.V1MachineFindRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find machines params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMachinesParams) WithDefaults() *FindMachinesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find machines params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindMachinesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find machines params
func (o *FindMachinesParams) WithTimeout(timeout time.Duration) *FindMachinesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find machines params
func (o *FindMachinesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find machines params
func (o *FindMachinesParams) WithContext(ctx context.Context) *FindMachinesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find machines params
func (o *FindMachinesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find machines params
func (o *FindMachinesParams) WithHTTPClient(client *http.Client) *FindMachinesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find machines params
func (o *FindMachinesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the find machines params
func (o *FindMachinesParams) WithBody(body *models.V1MachineFindRequest) *FindMachinesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the find machines params
func (o *FindMachinesParams) SetBody(body *models.V1MachineFindRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *FindMachinesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
