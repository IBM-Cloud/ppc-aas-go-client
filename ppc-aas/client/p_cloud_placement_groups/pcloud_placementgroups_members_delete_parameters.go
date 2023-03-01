// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_placement_groups

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

	"github.com/IBM-Cloud/ppc-aas-go-sdk/ppc-aas/models"
)

// NewPcloudPlacementgroupsMembersDeleteParams creates a new PcloudPlacementgroupsMembersDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudPlacementgroupsMembersDeleteParams() *PcloudPlacementgroupsMembersDeleteParams {
	return &PcloudPlacementgroupsMembersDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPlacementgroupsMembersDeleteParamsWithTimeout creates a new PcloudPlacementgroupsMembersDeleteParams object
// with the ability to set a timeout on a request.
func NewPcloudPlacementgroupsMembersDeleteParamsWithTimeout(timeout time.Duration) *PcloudPlacementgroupsMembersDeleteParams {
	return &PcloudPlacementgroupsMembersDeleteParams{
		timeout: timeout,
	}
}

// NewPcloudPlacementgroupsMembersDeleteParamsWithContext creates a new PcloudPlacementgroupsMembersDeleteParams object
// with the ability to set a context for a request.
func NewPcloudPlacementgroupsMembersDeleteParamsWithContext(ctx context.Context) *PcloudPlacementgroupsMembersDeleteParams {
	return &PcloudPlacementgroupsMembersDeleteParams{
		Context: ctx,
	}
}

// NewPcloudPlacementgroupsMembersDeleteParamsWithHTTPClient creates a new PcloudPlacementgroupsMembersDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudPlacementgroupsMembersDeleteParamsWithHTTPClient(client *http.Client) *PcloudPlacementgroupsMembersDeleteParams {
	return &PcloudPlacementgroupsMembersDeleteParams{
		HTTPClient: client,
	}
}

/*
PcloudPlacementgroupsMembersDeleteParams contains all the parameters to send to the API endpoint

	for the pcloud placementgroups members delete operation.

	Typically these are written to a http.Request.
*/
type PcloudPlacementgroupsMembersDeleteParams struct {

	/* Body.

	   Parameters for removing a Server in a Placement Group
	*/
	Body *models.PlacementGroupServer

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* PlacementGroupID.

	   Placement Group ID
	*/
	PlacementGroupID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud placementgroups members delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPlacementgroupsMembersDeleteParams) WithDefaults() *PcloudPlacementgroupsMembersDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud placementgroups members delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudPlacementgroupsMembersDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud placementgroups members delete params
func (o *PcloudPlacementgroupsMembersDeleteParams) WithTimeout(timeout time.Duration) *PcloudPlacementgroupsMembersDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud placementgroups members delete params
func (o *PcloudPlacementgroupsMembersDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud placementgroups members delete params
func (o *PcloudPlacementgroupsMembersDeleteParams) WithContext(ctx context.Context) *PcloudPlacementgroupsMembersDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud placementgroups members delete params
func (o *PcloudPlacementgroupsMembersDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud placementgroups members delete params
func (o *PcloudPlacementgroupsMembersDeleteParams) WithHTTPClient(client *http.Client) *PcloudPlacementgroupsMembersDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud placementgroups members delete params
func (o *PcloudPlacementgroupsMembersDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud placementgroups members delete params
func (o *PcloudPlacementgroupsMembersDeleteParams) WithBody(body *models.PlacementGroupServer) *PcloudPlacementgroupsMembersDeleteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud placementgroups members delete params
func (o *PcloudPlacementgroupsMembersDeleteParams) SetBody(body *models.PlacementGroupServer) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud placementgroups members delete params
func (o *PcloudPlacementgroupsMembersDeleteParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPlacementgroupsMembersDeleteParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud placementgroups members delete params
func (o *PcloudPlacementgroupsMembersDeleteParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithPlacementGroupID adds the placementGroupID to the pcloud placementgroups members delete params
func (o *PcloudPlacementgroupsMembersDeleteParams) WithPlacementGroupID(placementGroupID string) *PcloudPlacementgroupsMembersDeleteParams {
	o.SetPlacementGroupID(placementGroupID)
	return o
}

// SetPlacementGroupID adds the placementGroupId to the pcloud placementgroups members delete params
func (o *PcloudPlacementgroupsMembersDeleteParams) SetPlacementGroupID(placementGroupID string) {
	o.PlacementGroupID = placementGroupID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPlacementgroupsMembersDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param placement_group_id
	if err := r.SetPathParam("placement_group_id", o.PlacementGroupID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}