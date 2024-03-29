// Code generated by go-swagger; DO NOT EDIT.

package partition

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

// NewPartitionCapacityParams creates a new PartitionCapacityParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPartitionCapacityParams() *PartitionCapacityParams {
	return &PartitionCapacityParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPartitionCapacityParamsWithTimeout creates a new PartitionCapacityParams object
// with the ability to set a timeout on a request.
func NewPartitionCapacityParamsWithTimeout(timeout time.Duration) *PartitionCapacityParams {
	return &PartitionCapacityParams{
		timeout: timeout,
	}
}

// NewPartitionCapacityParamsWithContext creates a new PartitionCapacityParams object
// with the ability to set a context for a request.
func NewPartitionCapacityParamsWithContext(ctx context.Context) *PartitionCapacityParams {
	return &PartitionCapacityParams{
		Context: ctx,
	}
}

// NewPartitionCapacityParamsWithHTTPClient creates a new PartitionCapacityParams object
// with the ability to set a custom HTTPClient for a request.
func NewPartitionCapacityParamsWithHTTPClient(client *http.Client) *PartitionCapacityParams {
	return &PartitionCapacityParams{
		HTTPClient: client,
	}
}

/*
PartitionCapacityParams contains all the parameters to send to the API endpoint

	for the partition capacity operation.

	Typically these are written to a http.Request.
*/
type PartitionCapacityParams struct {

	// Body.
	Body *models.V1PartitionCapacityRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the partition capacity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartitionCapacityParams) WithDefaults() *PartitionCapacityParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the partition capacity params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PartitionCapacityParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the partition capacity params
func (o *PartitionCapacityParams) WithTimeout(timeout time.Duration) *PartitionCapacityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the partition capacity params
func (o *PartitionCapacityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the partition capacity params
func (o *PartitionCapacityParams) WithContext(ctx context.Context) *PartitionCapacityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the partition capacity params
func (o *PartitionCapacityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the partition capacity params
func (o *PartitionCapacityParams) WithHTTPClient(client *http.Client) *PartitionCapacityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the partition capacity params
func (o *PartitionCapacityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the partition capacity params
func (o *PartitionCapacityParams) WithBody(body *models.V1PartitionCapacityRequest) *PartitionCapacityParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the partition capacity params
func (o *PartitionCapacityParams) SetBody(body *models.V1PartitionCapacityRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PartitionCapacityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
