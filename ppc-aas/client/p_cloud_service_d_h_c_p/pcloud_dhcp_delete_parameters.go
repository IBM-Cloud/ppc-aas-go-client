// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_service_d_h_c_p

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

// NewPcloudDhcpDeleteParams creates a new PcloudDhcpDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudDhcpDeleteParams() *PcloudDhcpDeleteParams {
	return &PcloudDhcpDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudDhcpDeleteParamsWithTimeout creates a new PcloudDhcpDeleteParams object
// with the ability to set a timeout on a request.
func NewPcloudDhcpDeleteParamsWithTimeout(timeout time.Duration) *PcloudDhcpDeleteParams {
	return &PcloudDhcpDeleteParams{
		timeout: timeout,
	}
}

// NewPcloudDhcpDeleteParamsWithContext creates a new PcloudDhcpDeleteParams object
// with the ability to set a context for a request.
func NewPcloudDhcpDeleteParamsWithContext(ctx context.Context) *PcloudDhcpDeleteParams {
	return &PcloudDhcpDeleteParams{
		Context: ctx,
	}
}

// NewPcloudDhcpDeleteParamsWithHTTPClient creates a new PcloudDhcpDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudDhcpDeleteParamsWithHTTPClient(client *http.Client) *PcloudDhcpDeleteParams {
	return &PcloudDhcpDeleteParams{
		HTTPClient: client,
	}
}

/*
PcloudDhcpDeleteParams contains all the parameters to send to the API endpoint

	for the pcloud dhcp delete operation.

	Typically these are written to a http.Request.
*/
type PcloudDhcpDeleteParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* DhcpID.

	   The ID of the DHCP Server
	*/
	DhcpID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud dhcp delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudDhcpDeleteParams) WithDefaults() *PcloudDhcpDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud dhcp delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudDhcpDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud dhcp delete params
func (o *PcloudDhcpDeleteParams) WithTimeout(timeout time.Duration) *PcloudDhcpDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud dhcp delete params
func (o *PcloudDhcpDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud dhcp delete params
func (o *PcloudDhcpDeleteParams) WithContext(ctx context.Context) *PcloudDhcpDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud dhcp delete params
func (o *PcloudDhcpDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud dhcp delete params
func (o *PcloudDhcpDeleteParams) WithHTTPClient(client *http.Client) *PcloudDhcpDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud dhcp delete params
func (o *PcloudDhcpDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud dhcp delete params
func (o *PcloudDhcpDeleteParams) WithCloudInstanceID(cloudInstanceID string) *PcloudDhcpDeleteParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud dhcp delete params
func (o *PcloudDhcpDeleteParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithDhcpID adds the dhcpID to the pcloud dhcp delete params
func (o *PcloudDhcpDeleteParams) WithDhcpID(dhcpID string) *PcloudDhcpDeleteParams {
	o.SetDhcpID(dhcpID)
	return o
}

// SetDhcpID adds the dhcpId to the pcloud dhcp delete params
func (o *PcloudDhcpDeleteParams) SetDhcpID(dhcpID string) {
	o.DhcpID = dhcpID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudDhcpDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param dhcp_id
	if err := r.SetPathParam("dhcp_id", o.DhcpID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
