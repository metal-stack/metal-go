// Code generated by go-swagger; DO NOT EDIT.

package image

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
	"github.com/go-openapi/swag"

	"github.com/metal-stack/metal-go/api/models"
)

// NewFindImagesParams creates a new FindImagesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFindImagesParams() *FindImagesParams {
	return &FindImagesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFindImagesParamsWithTimeout creates a new FindImagesParams object
// with the ability to set a timeout on a request.
func NewFindImagesParamsWithTimeout(timeout time.Duration) *FindImagesParams {
	return &FindImagesParams{
		timeout: timeout,
	}
}

// NewFindImagesParamsWithContext creates a new FindImagesParams object
// with the ability to set a context for a request.
func NewFindImagesParamsWithContext(ctx context.Context) *FindImagesParams {
	return &FindImagesParams{
		Context: ctx,
	}
}

// NewFindImagesParamsWithHTTPClient creates a new FindImagesParams object
// with the ability to set a custom HTTPClient for a request.
func NewFindImagesParamsWithHTTPClient(client *http.Client) *FindImagesParams {
	return &FindImagesParams{
		HTTPClient: client,
	}
}

/*
FindImagesParams contains all the parameters to send to the API endpoint

	for the find images operation.

	Typically these are written to a http.Request.
*/
type FindImagesParams struct {

	// Body.
	Body *models.V1ImageFindRequest

	/* ShowUsage.

	   include image usage into response
	*/
	ShowUsage *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the find images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindImagesParams) WithDefaults() *FindImagesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the find images params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FindImagesParams) SetDefaults() {
	var (
		showUsageDefault = bool(false)
	)

	val := FindImagesParams{
		ShowUsage: &showUsageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the find images params
func (o *FindImagesParams) WithTimeout(timeout time.Duration) *FindImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the find images params
func (o *FindImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the find images params
func (o *FindImagesParams) WithContext(ctx context.Context) *FindImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the find images params
func (o *FindImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the find images params
func (o *FindImagesParams) WithHTTPClient(client *http.Client) *FindImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the find images params
func (o *FindImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the find images params
func (o *FindImagesParams) WithBody(body *models.V1ImageFindRequest) *FindImagesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the find images params
func (o *FindImagesParams) SetBody(body *models.V1ImageFindRequest) {
	o.Body = body
}

// WithShowUsage adds the showUsage to the find images params
func (o *FindImagesParams) WithShowUsage(showUsage *bool) *FindImagesParams {
	o.SetShowUsage(showUsage)
	return o
}

// SetShowUsage adds the showUsage to the find images params
func (o *FindImagesParams) SetShowUsage(showUsage *bool) {
	o.ShowUsage = showUsage
}

// WriteToRequest writes these params to a swagger request
func (o *FindImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.ShowUsage != nil {

		// query param show-usage
		var qrShowUsage bool

		if o.ShowUsage != nil {
			qrShowUsage = *o.ShowUsage
		}
		qShowUsage := swag.FormatBool(qrShowUsage)
		if qShowUsage != "" {

			if err := r.SetQueryParam("show-usage", qShowUsage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
