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

// NewCreateSizeReservationParams creates a new CreateSizeReservationParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateSizeReservationParams() *CreateSizeReservationParams {
	return &CreateSizeReservationParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateSizeReservationParamsWithTimeout creates a new CreateSizeReservationParams object
// with the ability to set a timeout on a request.
func NewCreateSizeReservationParamsWithTimeout(timeout time.Duration) *CreateSizeReservationParams {
	return &CreateSizeReservationParams{
		timeout: timeout,
	}
}

// NewCreateSizeReservationParamsWithContext creates a new CreateSizeReservationParams object
// with the ability to set a context for a request.
func NewCreateSizeReservationParamsWithContext(ctx context.Context) *CreateSizeReservationParams {
	return &CreateSizeReservationParams{
		Context: ctx,
	}
}

// NewCreateSizeReservationParamsWithHTTPClient creates a new CreateSizeReservationParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateSizeReservationParamsWithHTTPClient(client *http.Client) *CreateSizeReservationParams {
	return &CreateSizeReservationParams{
		HTTPClient: client,
	}
}

/*
CreateSizeReservationParams contains all the parameters to send to the API endpoint

	for the create size reservation operation.

	Typically these are written to a http.Request.
*/
type CreateSizeReservationParams struct {

	// Body.
	Body *models.V1SizeReservationCreateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create size reservation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSizeReservationParams) WithDefaults() *CreateSizeReservationParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create size reservation params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateSizeReservationParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create size reservation params
func (o *CreateSizeReservationParams) WithTimeout(timeout time.Duration) *CreateSizeReservationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create size reservation params
func (o *CreateSizeReservationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create size reservation params
func (o *CreateSizeReservationParams) WithContext(ctx context.Context) *CreateSizeReservationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create size reservation params
func (o *CreateSizeReservationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create size reservation params
func (o *CreateSizeReservationParams) WithHTTPClient(client *http.Client) *CreateSizeReservationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create size reservation params
func (o *CreateSizeReservationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create size reservation params
func (o *CreateSizeReservationParams) WithBody(body *models.V1SizeReservationCreateRequest) *CreateSizeReservationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create size reservation params
func (o *CreateSizeReservationParams) SetBody(body *models.V1SizeReservationCreateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateSizeReservationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
