// Code generated by go-swagger; DO NOT EDIT.

package sizeimageconstraint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new sizeimageconstraint API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sizeimageconstraint API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CreateSizeImageConstraint(params *CreateSizeImageConstraintParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSizeImageConstraintCreated, error)

	DeleteSizeImageConstraint(params *DeleteSizeImageConstraintParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSizeImageConstraintOK, error)

	FindSizeImageConstraint(params *FindSizeImageConstraintParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindSizeImageConstraintOK, error)

	ListSizeImageConstraints(params *ListSizeImageConstraintsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSizeImageConstraintsOK, error)

	TrySizeImageConstraint(params *TrySizeImageConstraintParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TrySizeImageConstraintOK, error)

	UpdateSizeImageConstraint(params *UpdateSizeImageConstraintParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSizeImageConstraintOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateSizeImageConstraint creates a sizeimageconstraint if the given ID already exists a conflict is returned
*/
func (a *Client) CreateSizeImageConstraint(params *CreateSizeImageConstraintParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateSizeImageConstraintCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateSizeImageConstraintParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createSizeImageConstraint",
		Method:             "PUT",
		PathPattern:        "/v1/size-image-constraint",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateSizeImageConstraintReader{formats: a.formats},
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
	success, ok := result.(*CreateSizeImageConstraintCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateSizeImageConstraintDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  DeleteSizeImageConstraint deletes an sizeimageconstraint and returns the deleted entity
*/
func (a *Client) DeleteSizeImageConstraint(params *DeleteSizeImageConstraintParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteSizeImageConstraintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteSizeImageConstraintParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "deleteSizeImageConstraint",
		Method:             "DELETE",
		PathPattern:        "/v1/size-image-constraint/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteSizeImageConstraintReader{formats: a.formats},
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
	success, ok := result.(*DeleteSizeImageConstraintOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteSizeImageConstraintDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  FindSizeImageConstraint gets sizeimageconstraint by id
*/
func (a *Client) FindSizeImageConstraint(params *FindSizeImageConstraintParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*FindSizeImageConstraintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFindSizeImageConstraintParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "findSizeImageConstraint",
		Method:             "GET",
		PathPattern:        "/v1/size-image-constraint/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FindSizeImageConstraintReader{formats: a.formats},
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
	success, ok := result.(*FindSizeImageConstraintOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*FindSizeImageConstraintDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  ListSizeImageConstraints gets all sizeimageconstraints
*/
func (a *Client) ListSizeImageConstraints(params *ListSizeImageConstraintsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ListSizeImageConstraintsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSizeImageConstraintsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "listSizeImageConstraints",
		Method:             "GET",
		PathPattern:        "/v1/size-image-constraint",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListSizeImageConstraintsReader{formats: a.formats},
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
	success, ok := result.(*ListSizeImageConstraintsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ListSizeImageConstraintsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  TrySizeImageConstraint tries if the given combination of image and size is supported and possible to allocate
*/
func (a *Client) TrySizeImageConstraint(params *TrySizeImageConstraintParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*TrySizeImageConstraintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTrySizeImageConstraintParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "trySizeImageConstraint",
		Method:             "POST",
		PathPattern:        "/v1/size-image-constraint/try",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TrySizeImageConstraintReader{formats: a.formats},
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
	success, ok := result.(*TrySizeImageConstraintOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*TrySizeImageConstraintDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
  UpdateSizeImageConstraint updates a sizeimageconstraint if the sizeimageconstraint was changed since this one was read a conflict is returned
*/
func (a *Client) UpdateSizeImageConstraint(params *UpdateSizeImageConstraintParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateSizeImageConstraintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateSizeImageConstraintParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateSizeImageConstraint",
		Method:             "POST",
		PathPattern:        "/v1/size-image-constraint",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateSizeImageConstraintReader{formats: a.formats},
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
	success, ok := result.(*UpdateSizeImageConstraintOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UpdateSizeImageConstraintDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
