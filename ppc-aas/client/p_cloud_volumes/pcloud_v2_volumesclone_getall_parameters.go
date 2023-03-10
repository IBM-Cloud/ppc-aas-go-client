// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

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

// NewPcloudV2VolumescloneGetallParams creates a new PcloudV2VolumescloneGetallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudV2VolumescloneGetallParams() *PcloudV2VolumescloneGetallParams {
	return &PcloudV2VolumescloneGetallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudV2VolumescloneGetallParamsWithTimeout creates a new PcloudV2VolumescloneGetallParams object
// with the ability to set a timeout on a request.
func NewPcloudV2VolumescloneGetallParamsWithTimeout(timeout time.Duration) *PcloudV2VolumescloneGetallParams {
	return &PcloudV2VolumescloneGetallParams{
		timeout: timeout,
	}
}

// NewPcloudV2VolumescloneGetallParamsWithContext creates a new PcloudV2VolumescloneGetallParams object
// with the ability to set a context for a request.
func NewPcloudV2VolumescloneGetallParamsWithContext(ctx context.Context) *PcloudV2VolumescloneGetallParams {
	return &PcloudV2VolumescloneGetallParams{
		Context: ctx,
	}
}

// NewPcloudV2VolumescloneGetallParamsWithHTTPClient creates a new PcloudV2VolumescloneGetallParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudV2VolumescloneGetallParamsWithHTTPClient(client *http.Client) *PcloudV2VolumescloneGetallParams {
	return &PcloudV2VolumescloneGetallParams{
		HTTPClient: client,
	}
}

/*
PcloudV2VolumescloneGetallParams contains all the parameters to send to the API endpoint

	for the pcloud v2 volumesclone getall operation.

	Typically these are written to a http.Request.
*/
type PcloudV2VolumescloneGetallParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* Filter.

	   volumes-clone filter to limit list items:
	prepare - includes status values (preparing, prepared)
	start   - includes status values (starting, available)
	execute - includes status values (executing, available-rollback)
	cancel  - includes status values (cancelling)
	completed - includes status values (completed)
	failed - includes status values (failed)
	cancelled - includes status values (cancelled)
	finalized - included status values (completed, failed, cancelled)

	*/
	Filter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud v2 volumesclone getall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudV2VolumescloneGetallParams) WithDefaults() *PcloudV2VolumescloneGetallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud v2 volumesclone getall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudV2VolumescloneGetallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud v2 volumesclone getall params
func (o *PcloudV2VolumescloneGetallParams) WithTimeout(timeout time.Duration) *PcloudV2VolumescloneGetallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud v2 volumesclone getall params
func (o *PcloudV2VolumescloneGetallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud v2 volumesclone getall params
func (o *PcloudV2VolumescloneGetallParams) WithContext(ctx context.Context) *PcloudV2VolumescloneGetallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud v2 volumesclone getall params
func (o *PcloudV2VolumescloneGetallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud v2 volumesclone getall params
func (o *PcloudV2VolumescloneGetallParams) WithHTTPClient(client *http.Client) *PcloudV2VolumescloneGetallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud v2 volumesclone getall params
func (o *PcloudV2VolumescloneGetallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud v2 volumesclone getall params
func (o *PcloudV2VolumescloneGetallParams) WithCloudInstanceID(cloudInstanceID string) *PcloudV2VolumescloneGetallParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud v2 volumesclone getall params
func (o *PcloudV2VolumescloneGetallParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithFilter adds the filter to the pcloud v2 volumesclone getall params
func (o *PcloudV2VolumescloneGetallParams) WithFilter(filter *string) *PcloudV2VolumescloneGetallParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the pcloud v2 volumesclone getall params
func (o *PcloudV2VolumescloneGetallParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudV2VolumescloneGetallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
