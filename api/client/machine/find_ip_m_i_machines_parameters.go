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

// NewFindIPMIMachinesParams creates a new FindIPMIMachinesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindIPMIMachinesParams() *FindIPMIMachinesParams {
	return &FindIPMIMachinesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindIPMIMachinesParamsWithTimeout creates a new FindIPMIMachinesParams object
// with the ability to set a timeout on a request.
func NewFindIPMIMachinesParamsWithTimeout(timeout time.Duration) *FindIPMIMachinesParams {
	return &FindIPMIMachinesParams{
		timeout: timeout,
	}
}

// NewFindIPMIMachinesParamsWithContext creates a new FindIPMIMachinesParams object
// with the ability to set a context for a request.
func NewFindIPMIMachinesParamsWithContext(ctx context.Context) *FindIPMIMachinesParams {
	return &FindIPMIMachinesParams{
		Context: ctx,
	}
}

// NewFindIPMIMachinesParamsWithHTTPClient creates a new FindIPMIMachinesParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindIPMIMachinesParamsWithHTTPClient(client *http.Client) *FindIPMIMachinesParams {
	return &FindIPMIMachinesParams{
		HTTPClient: client,
	}
}

/* FindIPMIMachinesParams contains all the parameters to send to the API endpoint
   for the find IP m i machines operation.

   Typically these are written to a http.Request.
*/
type FindIPMIMachinesParams struct {

	// Body.
	Body *models.V1MachineFindRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find IP m i machines params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindIPMIMachinesParams) WithDefaults() *FindIPMIMachinesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find IP m i machines params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindIPMIMachinesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find IP m i machines params
func (o *FindIPMIMachinesParams) WithTimeout(timeout time.Duration) *FindIPMIMachinesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find IP m i machines params
func (o *FindIPMIMachinesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find IP m i machines params
func (o *FindIPMIMachinesParams) WithContext(ctx context.Context) *FindIPMIMachinesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find IP m i machines params
func (o *FindIPMIMachinesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find IP m i machines params
func (o *FindIPMIMachinesParams) WithHTTPClient(client *http.Client) *FindIPMIMachinesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find IP m i machines params
func (o *FindIPMIMachinesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the find IP m i machines params
func (o *FindIPMIMachinesParams) WithBody(body *models.V1MachineFindRequest) *FindIPMIMachinesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the find IP m i machines params
func (o *FindIPMIMachinesParams) SetBody(body *models.V1MachineFindRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *FindIPMIMachinesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
