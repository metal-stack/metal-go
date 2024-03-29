// Code generated by go-swagger; DO NOT EDIT.

package filesystemlayout

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

// NewCreateFilesystemLayoutParams creates a new CreateFilesystemLayoutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateFilesystemLayoutParams() *CreateFilesystemLayoutParams {
	return &CreateFilesystemLayoutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateFilesystemLayoutParamsWithTimeout creates a new CreateFilesystemLayoutParams object
// with the ability to set a timeout on a request.
func NewCreateFilesystemLayoutParamsWithTimeout(timeout time.Duration) *CreateFilesystemLayoutParams {
	return &CreateFilesystemLayoutParams{
		timeout: timeout,
	}
}

// NewCreateFilesystemLayoutParamsWithContext creates a new CreateFilesystemLayoutParams object
// with the ability to set a context for a request.
func NewCreateFilesystemLayoutParamsWithContext(ctx context.Context) *CreateFilesystemLayoutParams {
	return &CreateFilesystemLayoutParams{
		Context: ctx,
	}
}

// NewCreateFilesystemLayoutParamsWithHTTPClient creates a new CreateFilesystemLayoutParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateFilesystemLayoutParamsWithHTTPClient(client *http.Client) *CreateFilesystemLayoutParams {
	return &CreateFilesystemLayoutParams{
		HTTPClient: client,
	}
}

/*
CreateFilesystemLayoutParams contains all the parameters to send to the API endpoint

	for the create filesystem layout operation.

	Typically these are written to a http.Request.
*/
type CreateFilesystemLayoutParams struct {

	// Body.
	Body *models.V1FilesystemLayoutCreateRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create filesystem layout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateFilesystemLayoutParams) WithDefaults() *CreateFilesystemLayoutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create filesystem layout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateFilesystemLayoutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create filesystem layout params
func (o *CreateFilesystemLayoutParams) WithTimeout(timeout time.Duration) *CreateFilesystemLayoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create filesystem layout params
func (o *CreateFilesystemLayoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create filesystem layout params
func (o *CreateFilesystemLayoutParams) WithContext(ctx context.Context) *CreateFilesystemLayoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create filesystem layout params
func (o *CreateFilesystemLayoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create filesystem layout params
func (o *CreateFilesystemLayoutParams) WithHTTPClient(client *http.Client) *CreateFilesystemLayoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create filesystem layout params
func (o *CreateFilesystemLayoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create filesystem layout params
func (o *CreateFilesystemLayoutParams) WithBody(body *models.V1FilesystemLayoutCreateRequest) *CreateFilesystemLayoutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create filesystem layout params
func (o *CreateFilesystemLayoutParams) SetBody(body *models.V1FilesystemLayoutCreateRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateFilesystemLayoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
