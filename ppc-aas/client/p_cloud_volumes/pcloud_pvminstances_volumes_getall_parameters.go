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

// NewPcloudPvminstancesVolumesGetallParams creates a new PcloudPvminstancesVolumesGetallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudPvminstancesVolumesGetallParams() *PcloudPvminstancesVolumesGetallParams {
	return &PcloudPvminstancesVolumesGetallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPvminstancesVolumesGetallParamsWithTimeout creates a new PcloudPvminstancesVolumesGetallParams object
// with the ability to set a timeout on a request.
func NewPcloudPvminstancesVolumesGetallParamsWithTimeout(timeout time.Duration) *PcloudPvminstancesVolumesGetallParams {
	return &PcloudPvminstancesVolumesGetallParams{
		timeout: timeout,
	}
}

// NewPcloudPvminstancesVolumesGetallParamsWithContext creates a new PcloudPvminstancesVolumesGetallParams object
// with the ability to set a context for a request.
func NewPcloudPvminstancesVolumesGetallParamsWithContext(ctx context.Context) *PcloudPvminstancesVolumesGetallParams {
	return &PcloudPvminstancesVolumesGetallParams{
		Context: ctx,
	}
}

// NewPcloudPvminstancesVolumesGetallParamsWithHTTPClient creates a new PcloudPvminstancesVolumesGetallParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudPvminstancesVolumesGetallParamsWithHTTPClient(client *http.Client) *PcloudPvminstancesVolumesGetallParams {
	return &PcloudPvminstancesVolumesGetallParams{
		HTTPClient: client,
	}
}

/*
PcloudPvminstancesVolumesGetallParams contains all the parameters to send to the API endpoint

	for the pcloud pvminstances volumes getall operation.

	Typically these are written to a http.Request.
*/
type PcloudPvminstancesVolumesGetallParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* PvmInstanceID.

	   PCloud PVM Instance ID
	*/
	PvmInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud pvminstances volumes getall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesVolumesGetallParams) WithDefaults() *PcloudPvminstancesVolumesGetallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud pvminstances volumes getall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPvminstancesVolumesGetallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud pvminstances volumes getall params
func (o *PcloudPvminstancesVolumesGetallParams) WithTimeout(timeout time.Duration) *PcloudPvminstancesVolumesGetallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud pvminstances volumes getall params
func (o *PcloudPvminstancesVolumesGetallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud pvminstances volumes getall params
func (o *PcloudPvminstancesVolumesGetallParams) WithContext(ctx context.Context) *PcloudPvminstancesVolumesGetallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud pvminstances volumes getall params
func (o *PcloudPvminstancesVolumesGetallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud pvminstances volumes getall params
func (o *PcloudPvminstancesVolumesGetallParams) WithHTTPClient(client *http.Client) *PcloudPvminstancesVolumesGetallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud pvminstances volumes getall params
func (o *PcloudPvminstancesVolumesGetallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud pvminstances volumes getall params
func (o *PcloudPvminstancesVolumesGetallParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPvminstancesVolumesGetallParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud pvminstances volumes getall params
func (o *PcloudPvminstancesVolumesGetallParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithPvmInstanceID adds the pvmInstanceID to the pcloud pvminstances volumes getall params
func (o *PcloudPvminstancesVolumesGetallParams) WithPvmInstanceID(pvmInstanceID string) *PcloudPvminstancesVolumesGetallParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the pcloud pvminstances volumes getall params
func (o *PcloudPvminstancesVolumesGetallParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPvminstancesVolumesGetallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param pvm_instance_id
	if err := r.SetPathParam("pvm_instance_id", o.PvmInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
