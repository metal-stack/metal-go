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
)

// NewListIssuesParams creates a new ListIssuesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListIssuesParams() *ListIssuesParams {
	return &ListIssuesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListIssuesParamsWithTimeout creates a new ListIssuesParams object
// with the ability to set a timeout on a request.
func NewListIssuesParamsWithTimeout(timeout time.Duration) *ListIssuesParams {
	return &ListIssuesParams{
		timeout: timeout,
	}
}

// NewListIssuesParamsWithContext creates a new ListIssuesParams object
// with the ability to set a context for a request.
func NewListIssuesParamsWithContext(ctx context.Context) *ListIssuesParams {
	return &ListIssuesParams{
		Context: ctx,
	}
}

// NewListIssuesParamsWithHTTPClient creates a new ListIssuesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListIssuesParamsWithHTTPClient(client *http.Client) *ListIssuesParams {
	return &ListIssuesParams{
		HTTPClient: client,
	}
}

/*
ListIssuesParams contains all the parameters to send to the API endpoint

	for the list issues operation.

	Typically these are written to a http.Request.
*/
type ListIssuesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list issues params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListIssuesParams) WithDefaults() *ListIssuesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list issues params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListIssuesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list issues params
func (o *ListIssuesParams) WithTimeout(timeout time.Duration) *ListIssuesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list issues params
func (o *ListIssuesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list issues params
func (o *ListIssuesParams) WithContext(ctx context.Context) *ListIssuesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list issues params
func (o *ListIssuesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list issues params
func (o *ListIssuesParams) WithHTTPClient(client *http.Client) *ListIssuesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list issues params
func (o *ListIssuesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListIssuesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
