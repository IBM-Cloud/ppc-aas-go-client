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

// NewPcloudNetworksGetallParams creates a new PcloudNetworksGetallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudNetworksGetallParams() *PcloudNetworksGetallParams {
	return &PcloudNetworksGetallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudNetworksGetallParamsWithTimeout creates a new PcloudNetworksGetallParams object
// with the ability to set a timeout on a request.
func NewPcloudNetworksGetallParamsWithTimeout(timeout time.Duration) *PcloudNetworksGetallParams {
	return &PcloudNetworksGetallParams{
		timeout: timeout,
	}
}

// NewPcloudNetworksGetallParamsWithContext creates a new PcloudNetworksGetallParams object
// with the ability to set a context for a request.
func NewPcloudNetworksGetallParamsWithContext(ctx context.Context) *PcloudNetworksGetallParams {
	return &PcloudNetworksGetallParams{
		Context: ctx,
	}
}

// NewPcloudNetworksGetallParamsWithHTTPClient creates a new PcloudNetworksGetallParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudNetworksGetallParamsWithHTTPClient(client *http.Client) *PcloudNetworksGetallParams {
	return &PcloudNetworksGetallParams{
		HTTPClient: client,
	}
}

/*
PcloudNetworksGetallParams contains all the parameters to send to the API endpoint

	for the pcloud networks getall operation.

	Typically these are written to a http.Request.
*/
type PcloudNetworksGetallParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* Filter.

	   A filter expression that filters resources listed in the response
	*/
	Filter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud networks getall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudNetworksGetallParams) WithDefaults() *PcloudNetworksGetallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud networks getall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudNetworksGetallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud networks getall params
func (o *PcloudNetworksGetallParams) WithTimeout(timeout time.Duration) *PcloudNetworksGetallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud networks getall params
func (o *PcloudNetworksGetallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud networks getall params
func (o *PcloudNetworksGetallParams) WithContext(ctx context.Context) *PcloudNetworksGetallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud networks getall params
func (o *PcloudNetworksGetallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud networks getall params
func (o *PcloudNetworksGetallParams) WithHTTPClient(client *http.Client) *PcloudNetworksGetallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud networks getall params
func (o *PcloudNetworksGetallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud networks getall params
func (o *PcloudNetworksGetallParams) WithCloudInstanceID(cloudInstanceID string) *PcloudNetworksGetallParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud networks getall params
func (o *PcloudNetworksGetallParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithFilter adds the filter to the pcloud networks getall params
func (o *PcloudNetworksGetallParams) WithFilter(filter *string) *PcloudNetworksGetallParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the pcloud networks getall params
func (o *PcloudNetworksGetallParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudNetworksGetallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if o.Filter != nil {

		// query param filter
		var qrFilter string

		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {

			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
