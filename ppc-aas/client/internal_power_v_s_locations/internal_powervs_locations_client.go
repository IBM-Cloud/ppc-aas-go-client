// Code generated by go-swagger; DO NOT EDIT.

package internal_power_v_s_locations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new internal power v s locations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for internal power v s locations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	InternalV1PowervsLocationsTransitgatewayGet(params *InternalV1PowervsLocationsTransitgatewayGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1PowervsLocationsTransitgatewayGetOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
InternalV1PowervsLocationsTransitgatewayGet gets list of p e r enabled power private cloud locations
*/
func (a *Client) InternalV1PowervsLocationsTransitgatewayGet(params *InternalV1PowervsLocationsTransitgatewayGetParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*InternalV1PowervsLocationsTransitgatewayGetOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInternalV1PowervsLocationsTransitgatewayGetParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "internal.v1.powervs.locations.transitgateway.get",
		Method:             "GET",
		PathPattern:        "/internal/v1/powervs/locations/transit-gateway",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InternalV1PowervsLocationsTransitgatewayGetReader{formats: a.formats},
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
	success, ok := result.(*InternalV1PowervsLocationsTransitgatewayGetOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for internal.v1.powervs.locations.transitgateway.get: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
