// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_networks

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

// NewPcloudNetworksPortsGetParams creates a new PcloudNetworksPortsGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudNetworksPortsGetParams() *PcloudNetworksPortsGetParams {
	return &PcloudNetworksPortsGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudNetworksPortsGetParamsWithTimeout creates a new PcloudNetworksPortsGetParams object
// with the ability to set a timeout on a request.
func NewPcloudNetworksPortsGetParamsWithTimeout(timeout time.Duration) *PcloudNetworksPortsGetParams {
	return &PcloudNetworksPortsGetParams{
		timeout: timeout,
	}
}

// NewPcloudNetworksPortsGetParamsWithContext creates a new PcloudNetworksPortsGetParams object
// with the ability to set a context for a request.
func NewPcloudNetworksPortsGetParamsWithContext(ctx context.Context) *PcloudNetworksPortsGetParams {
	return &PcloudNetworksPortsGetParams{
		Context: ctx,
	}
}

// NewPcloudNetworksPortsGetParamsWithHTTPClient creates a new PcloudNetworksPortsGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudNetworksPortsGetParamsWithHTTPClient(client *http.Client) *PcloudNetworksPortsGetParams {
	return &PcloudNetworksPortsGetParams{
		HTTPClient: client,
	}
}

/*
PcloudNetworksPortsGetParams contains all the parameters to send to the API endpoint

	for the pcloud networks ports get operation.

	Typically these are written to a http.Request.
*/
type PcloudNetworksPortsGetParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	/* PortID.

	   Port ID
	*/
	PortID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud networks ports get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudNetworksPortsGetParams) WithDefaults() *PcloudNetworksPortsGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud networks ports get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudNetworksPortsGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud networks ports get params
func (o *PcloudNetworksPortsGetParams) WithTimeout(timeout time.Duration) *PcloudNetworksPortsGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud networks ports get params
func (o *PcloudNetworksPortsGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud networks ports get params
func (o *PcloudNetworksPortsGetParams) WithContext(ctx context.Context) *PcloudNetworksPortsGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud networks ports get params
func (o *PcloudNetworksPortsGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud networks ports get params
func (o *PcloudNetworksPortsGetParams) WithHTTPClient(client *http.Client) *PcloudNetworksPortsGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud networks ports get params
func (o *PcloudNetworksPortsGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud networks ports get params
func (o *PcloudNetworksPortsGetParams) WithCloudInstanceID(cloudInstanceID string) *PcloudNetworksPortsGetParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud networks ports get params
func (o *PcloudNetworksPortsGetParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithNetworkID adds the networkID to the pcloud networks ports get params
func (o *PcloudNetworksPortsGetParams) WithNetworkID(networkID string) *PcloudNetworksPortsGetParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the pcloud networks ports get params
func (o *PcloudNetworksPortsGetParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithPortID adds the portID to the pcloud networks ports get params
func (o *PcloudNetworksPortsGetParams) WithPortID(portID string) *PcloudNetworksPortsGetParams {
	o.SetPortID(portID)
	return o
}

// SetPortID adds the portId to the pcloud networks ports get params
func (o *PcloudNetworksPortsGetParams) SetPortID(portID string) {
	o.PortID = portID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudNetworksPortsGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	// path param port_id
	if err := r.SetPathParam("port_id", o.PortID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
