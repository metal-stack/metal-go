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

// NewIssuesParams creates a new IssuesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIssuesParams() *IssuesParams {
	return &IssuesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIssuesParamsWithTimeout creates a new IssuesParams object
// with the ability to set a timeout on a request.
func NewIssuesParamsWithTimeout(timeout time.Duration) *IssuesParams {
	return &IssuesParams{
		timeout: timeout,
	}
}

// NewIssuesParamsWithContext creates a new IssuesParams object
// with the ability to set a context for a request.
func NewIssuesParamsWithContext(ctx context.Context) *IssuesParams {
	return &IssuesParams{
		Context: ctx,
	}
}

// NewIssuesParamsWithHTTPClient creates a new IssuesParams object
// with the ability to set a custom HTTPClient for a request.
func NewIssuesParamsWithHTTPClient(client *http.Client) *IssuesParams {
	return &IssuesParams{
		HTTPClient: client,
	}
}

/*
IssuesParams contains all the parameters to send to the API endpoint

	for the issues operation.

	Typically these are written to a http.Request.
*/
type IssuesParams struct {

	// Body.
	Body *models.V1MachineIssuesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the issues params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssuesParams) WithDefaults() *IssuesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the issues params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IssuesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the issues params
func (o *IssuesParams) WithTimeout(timeout time.Duration) *IssuesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the issues params
func (o *IssuesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the issues params
func (o *IssuesParams) WithContext(ctx context.Context) *IssuesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the issues params
func (o *IssuesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the issues params
func (o *IssuesParams) WithHTTPClient(client *http.Client) *IssuesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the issues params
func (o *IssuesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the issues params
func (o *IssuesParams) WithBody(body *models.V1MachineIssuesRequest) *IssuesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the issues params
func (o *IssuesParams) SetBody(body *models.V1MachineIssuesRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IssuesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
