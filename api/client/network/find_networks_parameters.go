// Code generated by go-swagger; DO NOT EDIT.

package network

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

// NewFindNetworksParams creates a new FindNetworksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindNetworksParams() *FindNetworksParams {
	return &FindNetworksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindNetworksParamsWithTimeout creates a new FindNetworksParams object
// with the ability to set a timeout on a request.
func NewFindNetworksParamsWithTimeout(timeout time.Duration) *FindNetworksParams {
	return &FindNetworksParams{
		timeout: timeout,
	}
}

// NewFindNetworksParamsWithContext creates a new FindNetworksParams object
// with the ability to set a context for a request.
func NewFindNetworksParamsWithContext(ctx context.Context) *FindNetworksParams {
	return &FindNetworksParams{
		Context: ctx,
	}
}

// NewFindNetworksParamsWithHTTPClient creates a new FindNetworksParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindNetworksParamsWithHTTPClient(client *http.Client) *FindNetworksParams {
	return &FindNetworksParams{
		HTTPClient: client,
	}
}

/* FindNetworksParams contains all the parameters to send to the API endpoint
   for the find networks operation.

   Typically these are written to a http.Request.
*/
type FindNetworksParams struct {

	// Body.
	Body *models.V1NetworkFindRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindNetworksParams) WithDefaults() *FindNetworksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindNetworksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the find networks params
func (o *FindNetworksParams) WithTimeout(timeout time.Duration) *FindNetworksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find networks params
func (o *FindNetworksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find networks params
func (o *FindNetworksParams) WithContext(ctx context.Context) *FindNetworksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find networks params
func (o *FindNetworksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find networks params
func (o *FindNetworksParams) WithHTTPClient(client *http.Client) *FindNetworksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find networks params
func (o *FindNetworksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the find networks params
func (o *FindNetworksParams) WithBody(body *models.V1NetworkFindRequest) *FindNetworksParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the find networks params
func (o *FindNetworksParams) SetBody(body *models.V1NetworkFindRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *FindNetworksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
