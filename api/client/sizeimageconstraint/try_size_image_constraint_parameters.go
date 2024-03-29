// Code generated by go-swagger; DO NOT EDIT.

package sizeimageconstraint

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

// NewTrySizeImageConstraintParams creates a new TrySizeImageConstraintParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTrySizeImageConstraintParams() *TrySizeImageConstraintParams {
	return &TrySizeImageConstraintParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTrySizeImageConstraintParamsWithTimeout creates a new TrySizeImageConstraintParams object
// with the ability to set a timeout on a request.
func NewTrySizeImageConstraintParamsWithTimeout(timeout time.Duration) *TrySizeImageConstraintParams {
	return &TrySizeImageConstraintParams{
		timeout: timeout,
	}
}

// NewTrySizeImageConstraintParamsWithContext creates a new TrySizeImageConstraintParams object
// with the ability to set a context for a request.
func NewTrySizeImageConstraintParamsWithContext(ctx context.Context) *TrySizeImageConstraintParams {
	return &TrySizeImageConstraintParams{
		Context: ctx,
	}
}

// NewTrySizeImageConstraintParamsWithHTTPClient creates a new TrySizeImageConstraintParams object
// with the ability to set a custom HTTPClient for a request.
func NewTrySizeImageConstraintParamsWithHTTPClient(client *http.Client) *TrySizeImageConstraintParams {
	return &TrySizeImageConstraintParams{
		HTTPClient: client,
	}
}

/*
TrySizeImageConstraintParams contains all the parameters to send to the API endpoint

	for the try size image constraint operation.

	Typically these are written to a http.Request.
*/
type TrySizeImageConstraintParams struct {

	// Body.
	Body *models.V1SizeImageConstraintTryRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the try size image constraint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TrySizeImageConstraintParams) WithDefaults() *TrySizeImageConstraintParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the try size image constraint params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TrySizeImageConstraintParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the try size image constraint params
func (o *TrySizeImageConstraintParams) WithTimeout(timeout time.Duration) *TrySizeImageConstraintParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the try size image constraint params
func (o *TrySizeImageConstraintParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the try size image constraint params
func (o *TrySizeImageConstraintParams) WithContext(ctx context.Context) *TrySizeImageConstraintParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the try size image constraint params
func (o *TrySizeImageConstraintParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the try size image constraint params
func (o *TrySizeImageConstraintParams) WithHTTPClient(client *http.Client) *TrySizeImageConstraintParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the try size image constraint params
func (o *TrySizeImageConstraintParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the try size image constraint params
func (o *TrySizeImageConstraintParams) WithBody(body *models.V1SizeImageConstraintTryRequest) *TrySizeImageConstraintParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the try size image constraint params
func (o *TrySizeImageConstraintParams) SetBody(body *models.V1SizeImageConstraintTryRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TrySizeImageConstraintParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
