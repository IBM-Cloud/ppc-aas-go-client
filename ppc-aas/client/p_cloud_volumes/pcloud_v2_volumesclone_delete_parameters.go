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

// NewPcloudV2VolumescloneDeleteParams creates a new PcloudV2VolumescloneDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudV2VolumescloneDeleteParams() *PcloudV2VolumescloneDeleteParams {
	return &PcloudV2VolumescloneDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudV2VolumescloneDeleteParamsWithTimeout creates a new PcloudV2VolumescloneDeleteParams object
// with the ability to set a timeout on a request.
func NewPcloudV2VolumescloneDeleteParamsWithTimeout(timeout time.Duration) *PcloudV2VolumescloneDeleteParams {
	return &PcloudV2VolumescloneDeleteParams{
		timeout: timeout,
	}
}

// NewPcloudV2VolumescloneDeleteParamsWithContext creates a new PcloudV2VolumescloneDeleteParams object
// with the ability to set a context for a request.
func NewPcloudV2VolumescloneDeleteParamsWithContext(ctx context.Context) *PcloudV2VolumescloneDeleteParams {
	return &PcloudV2VolumescloneDeleteParams{
		Context: ctx,
	}
}

// NewPcloudV2VolumescloneDeleteParamsWithHTTPClient creates a new PcloudV2VolumescloneDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudV2VolumescloneDeleteParamsWithHTTPClient(client *http.Client) *PcloudV2VolumescloneDeleteParams {
	return &PcloudV2VolumescloneDeleteParams{
		HTTPClient: client,
	}
}

/*
PcloudV2VolumescloneDeleteParams contains all the parameters to send to the API endpoint

	for the pcloud v2 volumesclone delete operation.

	Typically these are written to a http.Request.
*/
type PcloudV2VolumescloneDeleteParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* VolumesCloneID.

	   Volumes Clone ID
	*/
	VolumesCloneID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud v2 volumesclone delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudV2VolumescloneDeleteParams) WithDefaults() *PcloudV2VolumescloneDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud v2 volumesclone delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudV2VolumescloneDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud v2 volumesclone delete params
func (o *PcloudV2VolumescloneDeleteParams) WithTimeout(timeout time.Duration) *PcloudV2VolumescloneDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud v2 volumesclone delete params
func (o *PcloudV2VolumescloneDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud v2 volumesclone delete params
func (o *PcloudV2VolumescloneDeleteParams) WithContext(ctx context.Context) *PcloudV2VolumescloneDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud v2 volumesclone delete params
func (o *PcloudV2VolumescloneDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud v2 volumesclone delete params
func (o *PcloudV2VolumescloneDeleteParams) WithHTTPClient(client *http.Client) *PcloudV2VolumescloneDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud v2 volumesclone delete params
func (o *PcloudV2VolumescloneDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud v2 volumesclone delete params
func (o *PcloudV2VolumescloneDeleteParams) WithCloudInstanceID(cloudInstanceID string) *PcloudV2VolumescloneDeleteParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud v2 volumesclone delete params
func (o *PcloudV2VolumescloneDeleteParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithVolumesCloneID adds the volumesCloneID to the pcloud v2 volumesclone delete params
func (o *PcloudV2VolumescloneDeleteParams) WithVolumesCloneID(volumesCloneID string) *PcloudV2VolumescloneDeleteParams {
	o.SetVolumesCloneID(volumesCloneID)
	return o
}

// SetVolumesCloneID adds the volumesCloneId to the pcloud v2 volumesclone delete params
func (o *PcloudV2VolumescloneDeleteParams) SetVolumesCloneID(volumesCloneID string) {
	o.VolumesCloneID = volumesCloneID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudV2VolumescloneDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param volumes_clone_id
	if err := r.SetPathParam("volumes_clone_id", o.VolumesCloneID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
