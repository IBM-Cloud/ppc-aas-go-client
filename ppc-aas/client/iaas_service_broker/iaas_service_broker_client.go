// Code generated by go-swagger; DO NOT EDIT.

package iaas_service_broker

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new iaas service broker API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for iaas service broker API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ServiceBrokerHealth(params *ServiceBrokerHealthParams, opts ...ClientOption) (*ServiceBrokerHealthOK, error)

	ServiceBrokerHealthHead(params *ServiceBrokerHealthHeadParams, opts ...ClientOption) (*ServiceBrokerHealthHeadOK, error)

	ServiceBrokerTestTimeout(params *ServiceBrokerTestTimeoutParams, opts ...ClientOption) (*ServiceBrokerTestTimeoutOK, error)

	ServiceBrokerVersion(params *ServiceBrokerVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceBrokerVersionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ServiceBrokerHealth gets current server health
*/
func (a *Client) ServiceBrokerHealth(params *ServiceBrokerHealthParams, opts ...ClientOption) (*ServiceBrokerHealthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBrokerHealthParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "serviceBroker.health",
		Method:             "GET",
		PathPattern:        "/broker/v1/health",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBrokerHealthReader{formats: a.formats},
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
	success, ok := result.(*ServiceBrokerHealthOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for serviceBroker.health: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ServiceBrokerHealthHead gets current server health
*/
func (a *Client) ServiceBrokerHealthHead(params *ServiceBrokerHealthHeadParams, opts ...ClientOption) (*ServiceBrokerHealthHeadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBrokerHealthHeadParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "serviceBroker.health.head",
		Method:             "HEAD",
		PathPattern:        "/broker/v1/health",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBrokerHealthHeadReader{formats: a.formats},
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
	success, ok := result.(*ServiceBrokerHealthHeadOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for serviceBroker.health.head: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ServiceBrokerTestTimeout gets current server version
*/
func (a *Client) ServiceBrokerTestTimeout(params *ServiceBrokerTestTimeoutParams, opts ...ClientOption) (*ServiceBrokerTestTimeoutOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBrokerTestTimeoutParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "serviceBroker.test.timeout",
		Method:             "GET",
		PathPattern:        "/broker/v1/test/timeout",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBrokerTestTimeoutReader{formats: a.formats},
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
	success, ok := result.(*ServiceBrokerTestTimeoutOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for serviceBroker.test.timeout: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ServiceBrokerVersion gets current server version
*/
func (a *Client) ServiceBrokerVersion(params *ServiceBrokerVersionParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*ServiceBrokerVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServiceBrokerVersionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "serviceBroker.version",
		Method:             "GET",
		PathPattern:        "/broker/v1/version",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ServiceBrokerVersionReader{formats: a.formats},
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
	success, ok := result.(*ServiceBrokerVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for serviceBroker.version: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
