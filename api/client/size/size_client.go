// Code generated by go-swagger; DO NOT EDIT.

package size

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new size API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new size API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new size API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for size API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateSize(params *CreateSizeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSizeCreated, error)

	CreateSizeReservation(params *CreateSizeReservationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSizeReservationCreated, error)

	DeleteSize(params *DeleteSizeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSizeOK, error)

	DeleteSizeReservation(params *DeleteSizeReservationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSizeReservationOK, error)

	FindSize(params *FindSizeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindSizeOK, error)

	FindSizeReservations(params *FindSizeReservationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindSizeReservationsOK, error)

	GetSizeReservation(params *GetSizeReservationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSizeReservationOK, error)

	ListSizeReservations(params *ListSizeReservationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSizeReservationsOK, error)

	ListSizes(params *ListSizesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSizesOK, error)

	SizeReservationsUsage(params *SizeReservationsUsageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SizeReservationsUsageOK, error)

	Suggest(params *SuggestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SuggestOK, error)

	UpdateSize(params *UpdateSizeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSizeOK, error)

	UpdateSizeReservation(params *UpdateSizeReservationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSizeReservationOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CreateSize creates a size if the given ID already exists a conflict is returned
*/
func (a *Client) CreateSize(params *CreateSizeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSizeCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSizeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createSize",
		Method:             "PUT",
		PathPattern:        "/v1/size",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSizeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateSizeCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateSizeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreateSizeReservation creates a size reservation if the given ID already exists a conflict is returned
*/
func (a *Client) CreateSizeReservation(params *CreateSizeReservationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSizeReservationCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSizeReservationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createSizeReservation",
		Method:             "PUT",
		PathPattern:        "/v1/size/reservations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSizeReservationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateSizeReservationCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateSizeReservationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteSize deletes an size and returns the deleted entity
*/
func (a *Client) DeleteSize(params *DeleteSizeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSizeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSizeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteSize",
		Method:             "DELETE",
		PathPattern:        "/v1/size/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSizeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteSizeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteSizeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteSizeReservation deletes a size reservation and returns the deleted entity
*/
func (a *Client) DeleteSizeReservation(params *DeleteSizeReservationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSizeReservationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSizeReservationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteSizeReservation",
		Method:             "DELETE",
		PathPattern:        "/v1/size/reservations/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSizeReservationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteSizeReservationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteSizeReservationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
FindSize gets size by id
*/
func (a *Client) FindSize(params *FindSizeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindSizeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindSizeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findSize",
		Method:             "GET",
		PathPattern:        "/v1/size/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindSizeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindSizeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FindSizeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
FindSizeReservations gets all size reservations
*/
func (a *Client) FindSizeReservations(params *FindSizeReservationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindSizeReservationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindSizeReservationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findSizeReservations",
		Method:             "POST",
		PathPattern:        "/v1/size/reservations/find",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindSizeReservationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*FindSizeReservationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FindSizeReservationsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetSizeReservation gets size reservation by id
*/
func (a *Client) GetSizeReservation(params *GetSizeReservationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetSizeReservationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetSizeReservationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getSizeReservation",
		Method:             "GET",
		PathPattern:        "/v1/size/reservations/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetSizeReservationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetSizeReservationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetSizeReservationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListSizeReservations gets all size reservations
*/
func (a *Client) ListSizeReservations(params *ListSizeReservationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSizeReservationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSizeReservationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listSizeReservations",
		Method:             "GET",
		PathPattern:        "/v1/size/reservations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListSizeReservationsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListSizeReservationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListSizeReservationsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListSizes gets all sizes
*/
func (a *Client) ListSizes(params *ListSizesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSizesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSizesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listSizes",
		Method:             "GET",
		PathPattern:        "/v1/size",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListSizesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListSizesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListSizesDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SizeReservationsUsage gets all size reservations
*/
func (a *Client) SizeReservationsUsage(params *SizeReservationsUsageParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SizeReservationsUsageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSizeReservationsUsageParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "sizeReservationsUsage",
		Method:             "POST",
		PathPattern:        "/v1/size/reservations/usage",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SizeReservationsUsageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SizeReservationsUsageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SizeReservationsUsageDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
Suggest froms a given machine id returns the appropriate size
*/
func (a *Client) Suggest(params *SuggestParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*SuggestOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSuggestParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "suggest",
		Method:             "POST",
		PathPattern:        "/v1/size/suggest",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SuggestReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SuggestOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SuggestDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateSize updates a size if the size was changed since this one was read a conflict is returned
*/
func (a *Client) UpdateSize(params *UpdateSizeParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSizeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSizeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateSize",
		Method:             "POST",
		PathPattern:        "/v1/size",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSizeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateSizeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateSizeDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateSizeReservation updates a size reservation if the size reservation was changed since this one was read a conflict is returned
*/
func (a *Client) UpdateSizeReservation(params *UpdateSizeReservationParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSizeReservationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSizeReservationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateSizeReservation",
		Method:             "POST",
		PathPattern:        "/v1/size/reservations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSizeReservationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateSizeReservationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateSizeReservationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
