// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new network API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new network API client with basic auth credentials.
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

// New creates a new network API client with a bearer token for authentication.
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
Client for network API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// This client is generated with a few options you might find useful for your swagger spec.
//
// Feel free to add you own set of options.

// WithContentType allows the client to force the Content-Type header
// to negotiate a specific Consumer from the server.
//
// You may use this option to set arbitrary extensions to your MIME media type.
func WithContentType(mime string) ClientOption {
	return func(r *runtime.ClientOperation) {
		r.ConsumesMediaTypes = []string{mime}
	}
}

// WithContentTypeStarStar sets the Content-Type header to "*/*".
func WithContentTypeStarStar(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"*/*"}
}

// WithContentTypeApplicationJSON sets the Content-Type header to "application/json".
func WithContentTypeApplicationJSON(r *runtime.ClientOperation) {
	r.ConsumesMediaTypes = []string{"application/json"}
}

// ClientService is the interface for Client methods
type ClientService interface {
	AllocateNetwork(params *AllocateNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AllocateNetworkCreated, error)

	CreateNetwork(params *CreateNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateNetworkCreated, error)

	DeleteNetwork(params *DeleteNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteNetworkOK, error)

	FindNetwork(params *FindNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindNetworkOK, error)

	FindNetworks(params *FindNetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindNetworksOK, error)

	FreeNetwork(params *FreeNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FreeNetworkOK, error)

	FreeNetworkDeprecated(params *FreeNetworkDeprecatedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FreeNetworkDeprecatedOK, error)

	ListNetworks(params *ListNetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNetworksOK, error)

	UpdateNetwork(params *UpdateNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateNetworkOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AllocateNetwork allocates a child network from a partition s private super network
*/
func (a *Client) AllocateNetwork(params *AllocateNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*AllocateNetworkCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAllocateNetworkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "allocateNetwork",
		Method:             "POST",
		PathPattern:        "/v1/network/allocate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AllocateNetworkReader{formats: a.formats},
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
	success, ok := result.(*AllocateNetworkCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AllocateNetworkDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CreateNetwork creates a network if the given ID already exists a conflict is returned
*/
func (a *Client) CreateNetwork(params *CreateNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateNetworkCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateNetworkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createNetwork",
		Method:             "PUT",
		PathPattern:        "/v1/network",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateNetworkReader{formats: a.formats},
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
	success, ok := result.(*CreateNetworkCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateNetworkDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
DeleteNetwork deletes a network and returns the deleted entity
*/
func (a *Client) DeleteNetwork(params *DeleteNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNetworkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteNetwork",
		Method:             "DELETE",
		PathPattern:        "/v1/network/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteNetworkReader{formats: a.formats},
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
	success, ok := result.(*DeleteNetworkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteNetworkDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
FindNetwork gets network by id
*/
func (a *Client) FindNetwork(params *FindNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindNetworkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findNetwork",
		Method:             "GET",
		PathPattern:        "/v1/network/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindNetworkReader{formats: a.formats},
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
	success, ok := result.(*FindNetworkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FindNetworkDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
FindNetworks gets all networks that match given properties
*/
func (a *Client) FindNetworks(params *FindNetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindNetworksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindNetworksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findNetworks",
		Method:             "POST",
		PathPattern:        "/v1/network/find",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindNetworksReader{formats: a.formats},
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
	success, ok := result.(*FindNetworksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FindNetworksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
FreeNetwork frees a network
*/
func (a *Client) FreeNetwork(params *FreeNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FreeNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFreeNetworkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "freeNetwork",
		Method:             "DELETE",
		PathPattern:        "/v1/network/free/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FreeNetworkReader{formats: a.formats},
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
	success, ok := result.(*FreeNetworkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FreeNetworkDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
FreeNetworkDeprecated frees a network
*/
func (a *Client) FreeNetworkDeprecated(params *FreeNetworkDeprecatedParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FreeNetworkDeprecatedOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFreeNetworkDeprecatedParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "freeNetworkDeprecated",
		Method:             "POST",
		PathPattern:        "/v1/network/free/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"*/*", "application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FreeNetworkDeprecatedReader{formats: a.formats},
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
	success, ok := result.(*FreeNetworkDeprecatedOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FreeNetworkDeprecatedDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ListNetworks gets all networks
*/
func (a *Client) ListNetworks(params *ListNetworksParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListNetworksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListNetworksParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listNetworks",
		Method:             "GET",
		PathPattern:        "/v1/network",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListNetworksReader{formats: a.formats},
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
	success, ok := result.(*ListNetworksOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListNetworksDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UpdateNetwork updates a network if the network was changed since this one was read a conflict is returned
*/
func (a *Client) UpdateNetwork(params *UpdateNetworkParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateNetworkOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNetworkParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateNetwork",
		Method:             "POST",
		PathPattern:        "/v1/network",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateNetworkReader{formats: a.formats},
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
	success, ok := result.(*UpdateNetworkOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateNetworkDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
