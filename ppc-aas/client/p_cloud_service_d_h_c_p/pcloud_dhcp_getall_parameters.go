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

// NewPcloudDhcpGetallParams creates a new PcloudDhcpGetallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudDhcpGetallParams() *PcloudDhcpGetallParams {
	return &PcloudDhcpGetallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudDhcpGetallParamsWithTimeout creates a new PcloudDhcpGetallParams object
// with the ability to set a timeout on a request.
func NewPcloudDhcpGetallParamsWithTimeout(timeout time.Duration) *PcloudDhcpGetallParams {
	return &PcloudDhcpGetallParams{
		timeout: timeout,
	}
}

// NewPcloudDhcpGetallParamsWithContext creates a new PcloudDhcpGetallParams object
// with the ability to set a context for a request.
func NewPcloudDhcpGetallParamsWithContext(ctx context.Context) *PcloudDhcpGetallParams {
	return &PcloudDhcpGetallParams{
		Context: ctx,
	}
}

// NewPcloudDhcpGetallParamsWithHTTPClient creates a new PcloudDhcpGetallParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudDhcpGetallParamsWithHTTPClient(client *http.Client) *PcloudDhcpGetallParams {
	return &PcloudDhcpGetallParams{
		HTTPClient: client,
	}
}

/*
PcloudDhcpGetallParams contains all the parameters to send to the API endpoint

	for the pcloud dhcp getall operation.

	Typically these are written to a http.Request.
*/
type PcloudDhcpGetallParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud dhcp getall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudDhcpGetallParams) WithDefaults() *PcloudDhcpGetallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud dhcp getall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudDhcpGetallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud dhcp getall params
func (o *PcloudDhcpGetallParams) WithTimeout(timeout time.Duration) *PcloudDhcpGetallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud dhcp getall params
func (o *PcloudDhcpGetallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud dhcp getall params
func (o *PcloudDhcpGetallParams) WithContext(ctx context.Context) *PcloudDhcpGetallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud dhcp getall params
func (o *PcloudDhcpGetallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud dhcp getall params
func (o *PcloudDhcpGetallParams) WithHTTPClient(client *http.Client) *PcloudDhcpGetallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud dhcp getall params
func (o *PcloudDhcpGetallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud dhcp getall params
func (o *PcloudDhcpGetallParams) WithCloudInstanceID(cloudInstanceID string) *PcloudDhcpGetallParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud dhcp getall params
func (o *PcloudDhcpGetallParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudDhcpGetallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
