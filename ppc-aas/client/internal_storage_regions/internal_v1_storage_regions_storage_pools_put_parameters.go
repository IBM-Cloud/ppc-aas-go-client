// Code generated by go-swagger; DO NOT EDIT.

package internal_storage_regions

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

// NewInternalV1StorageRegionsStoragePoolsPutParams creates a new InternalV1StorageRegionsStoragePoolsPutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInternalV1StorageRegionsStoragePoolsPutParams() *InternalV1StorageRegionsStoragePoolsPutParams {
	return &InternalV1StorageRegionsStoragePoolsPutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInternalV1StorageRegionsStoragePoolsPutParamsWithTimeout creates a new InternalV1StorageRegionsStoragePoolsPutParams object
// with the ability to set a timeout on a request.
func NewInternalV1StorageRegionsStoragePoolsPutParamsWithTimeout(timeout time.Duration) *InternalV1StorageRegionsStoragePoolsPutParams {
	return &InternalV1StorageRegionsStoragePoolsPutParams{
		timeout: timeout,
	}
}

// NewInternalV1StorageRegionsStoragePoolsPutParamsWithContext creates a new InternalV1StorageRegionsStoragePoolsPutParams object
// with the ability to set a context for a request.
func NewInternalV1StorageRegionsStoragePoolsPutParamsWithContext(ctx context.Context) *InternalV1StorageRegionsStoragePoolsPutParams {
	return &InternalV1StorageRegionsStoragePoolsPutParams{
		Context: ctx,
	}
}

// NewInternalV1StorageRegionsStoragePoolsPutParamsWithHTTPClient creates a new InternalV1StorageRegionsStoragePoolsPutParams object
// with the ability to set a custom HTTPClient for a request.
func NewInternalV1StorageRegionsStoragePoolsPutParamsWithHTTPClient(client *http.Client) *InternalV1StorageRegionsStoragePoolsPutParams {
	return &InternalV1StorageRegionsStoragePoolsPutParams{
		HTTPClient: client,
	}
}

/*
InternalV1StorageRegionsStoragePoolsPutParams contains all the parameters to send to the API endpoint

	for the internal v1 storage regions storage pools put operation.

	Typically these are written to a http.Request.
*/
type InternalV1StorageRegionsStoragePoolsPutParams struct {

	/* Body.

	   Parameters for updating a storage pool
	*/
	Body *models.UpdateStoragePool

	/* RegionZoneID.

	   ID of a Power Cloud Region Zone
	*/
	RegionZoneID string

	/* StoragePoolName.

	   Storage pool name
	*/
	StoragePoolName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the internal v1 storage regions storage pools put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InternalV1StorageRegionsStoragePoolsPutParams) WithDefaults() *InternalV1StorageRegionsStoragePoolsPutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the internal v1 storage regions storage pools put params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InternalV1StorageRegionsStoragePoolsPutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the internal v1 storage regions storage pools put params
func (o *InternalV1StorageRegionsStoragePoolsPutParams) WithTimeout(timeout time.Duration) *InternalV1StorageRegionsStoragePoolsPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the internal v1 storage regions storage pools put params
func (o *InternalV1StorageRegionsStoragePoolsPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the internal v1 storage regions storage pools put params
func (o *InternalV1StorageRegionsStoragePoolsPutParams) WithContext(ctx context.Context) *InternalV1StorageRegionsStoragePoolsPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the internal v1 storage regions storage pools put params
func (o *InternalV1StorageRegionsStoragePoolsPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the internal v1 storage regions storage pools put params
func (o *InternalV1StorageRegionsStoragePoolsPutParams) WithHTTPClient(client *http.Client) *InternalV1StorageRegionsStoragePoolsPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the internal v1 storage regions storage pools put params
func (o *InternalV1StorageRegionsStoragePoolsPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the internal v1 storage regions storage pools put params
func (o *InternalV1StorageRegionsStoragePoolsPutParams) WithBody(body *models.UpdateStoragePool) *InternalV1StorageRegionsStoragePoolsPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the internal v1 storage regions storage pools put params
func (o *InternalV1StorageRegionsStoragePoolsPutParams) SetBody(body *models.UpdateStoragePool) {
	o.Body = body
}

// WithRegionZoneID adds the regionZoneID to the internal v1 storage regions storage pools put params
func (o *InternalV1StorageRegionsStoragePoolsPutParams) WithRegionZoneID(regionZoneID string) *InternalV1StorageRegionsStoragePoolsPutParams {
	o.SetRegionZoneID(regionZoneID)
	return o
}

// SetRegionZoneID adds the regionZoneId to the internal v1 storage regions storage pools put params
func (o *InternalV1StorageRegionsStoragePoolsPutParams) SetRegionZoneID(regionZoneID string) {
	o.RegionZoneID = regionZoneID
}

// WithStoragePoolName adds the storagePoolName to the internal v1 storage regions storage pools put params
func (o *InternalV1StorageRegionsStoragePoolsPutParams) WithStoragePoolName(storagePoolName string) *InternalV1StorageRegionsStoragePoolsPutParams {
	o.SetStoragePoolName(storagePoolName)
	return o
}

// SetStoragePoolName adds the storagePoolName to the internal v1 storage regions storage pools put params
func (o *InternalV1StorageRegionsStoragePoolsPutParams) SetStoragePoolName(storagePoolName string) {
	o.StoragePoolName = storagePoolName
}

// WriteToRequest writes these params to a swagger request
func (o *InternalV1StorageRegionsStoragePoolsPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param region_zone_id
	if err := r.SetPathParam("region_zone_id", o.RegionZoneID); err != nil {
		return err
	}

	// path param storage_pool_name
	if err := r.SetPathParam("storage_pool_name", o.StoragePoolName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
