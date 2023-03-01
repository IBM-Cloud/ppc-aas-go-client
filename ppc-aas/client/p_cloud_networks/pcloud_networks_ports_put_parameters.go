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

	"github.com/IBM-Cloud/ppc-aas-go-client/ppc-aas/models"
)

// NewPcloudNetworksPortsPutParams creates a new PcloudNetworksPortsPutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudNetworksPortsPutParams() *PcloudNetworksPortsPutParams {
	return &PcloudNetworksPortsPutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudNetworksPortsPutParamsWithTimeout creates a new PcloudNetworksPortsPutParams object
// with the ability to set a timeout on a request.
func NewPcloudNetworksPortsPutParamsWithTimeout(timeout time.Duration) *PcloudNetworksPortsPutParams {
	return &PcloudNetworksPortsPutParams{
		timeout: timeout,
	}
}

// NewPcloudNetworksPortsPutParamsWithContext creates a new PcloudNetworksPortsPutParams object
// with the ability to set a context for a request.
func NewPcloudNetworksPortsPutParamsWithContext(ctx context.Context) *PcloudNetworksPortsPutParams {
	return &PcloudNetworksPortsPutParams{
		Context: ctx,
	}
}

// NewPcloudNetworksPortsPutParamsWithHTTPClient creates a new PcloudNetworksPortsPutParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudNetworksPortsPutParamsWithHTTPClient(client *http.Client) *PcloudNetworksPortsPutParams {
	return &PcloudNetworksPortsPutParams{
		HTTPClient: client,
	}
}

/*
PcloudNetworksPortsPutParams contains all the parameters to send to the API endpoint

	for the pcloud networks ports put operation.

	Typically these are written to a http.Request.
*/
type PcloudNetworksPortsPutParams struct {

	/* Body.

	   Parameters for updating a Port
	*/
	Body *models.NetworkPortUpdate

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

// WithDefaults hydrates default values in the pcloud networks ports put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudNetworksPortsPutParams) WithDefaults() *PcloudNetworksPortsPutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud networks ports put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudNetworksPortsPutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud networks ports put params
func (o *PcloudNetworksPortsPutParams) WithTimeout(timeout time.Duration) *PcloudNetworksPortsPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud networks ports put params
func (o *PcloudNetworksPortsPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud networks ports put params
func (o *PcloudNetworksPortsPutParams) WithContext(ctx context.Context) *PcloudNetworksPortsPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud networks ports put params
func (o *PcloudNetworksPortsPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud networks ports put params
func (o *PcloudNetworksPortsPutParams) WithHTTPClient(client *http.Client) *PcloudNetworksPortsPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud networks ports put params
func (o *PcloudNetworksPortsPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud networks ports put params
func (o *PcloudNetworksPortsPutParams) WithBody(body *models.NetworkPortUpdate) *PcloudNetworksPortsPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud networks ports put params
func (o *PcloudNetworksPortsPutParams) SetBody(body *models.NetworkPortUpdate) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud networks ports put params
func (o *PcloudNetworksPortsPutParams) WithCloudInstanceID(cloudInstanceID string) *PcloudNetworksPortsPutParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud networks ports put params
func (o *PcloudNetworksPortsPutParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithNetworkID adds the networkID to the pcloud networks ports put params
func (o *PcloudNetworksPortsPutParams) WithNetworkID(networkID string) *PcloudNetworksPortsPutParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the pcloud networks ports put params
func (o *PcloudNetworksPortsPutParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithPortID adds the portID to the pcloud networks ports put params
func (o *PcloudNetworksPortsPutParams) WithPortID(portID string) *PcloudNetworksPortsPutParams {
	o.SetPortID(portID)
	return o
}

// SetPortID adds the portId to the pcloud networks ports put params
func (o *PcloudNetworksPortsPutParams) SetPortID(portID string) {
	o.PortID = portID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudNetworksPortsPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
