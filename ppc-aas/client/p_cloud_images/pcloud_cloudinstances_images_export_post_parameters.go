// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_images

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

// NewPcloudCloudinstancesImagesExportPostParams creates a new PcloudCloudinstancesImagesExportPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudCloudinstancesImagesExportPostParams() *PcloudCloudinstancesImagesExportPostParams {
	return &PcloudCloudinstancesImagesExportPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudCloudinstancesImagesExportPostParamsWithTimeout creates a new PcloudCloudinstancesImagesExportPostParams object
// with the ability to set a timeout on a request.
func NewPcloudCloudinstancesImagesExportPostParamsWithTimeout(timeout time.Duration) *PcloudCloudinstancesImagesExportPostParams {
	return &PcloudCloudinstancesImagesExportPostParams{
		timeout: timeout,
	}
}

// NewPcloudCloudinstancesImagesExportPostParamsWithContext creates a new PcloudCloudinstancesImagesExportPostParams object
// with the ability to set a context for a request.
func NewPcloudCloudinstancesImagesExportPostParamsWithContext(ctx context.Context) *PcloudCloudinstancesImagesExportPostParams {
	return &PcloudCloudinstancesImagesExportPostParams{
		Context: ctx,
	}
}

// NewPcloudCloudinstancesImagesExportPostParamsWithHTTPClient creates a new PcloudCloudinstancesImagesExportPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudCloudinstancesImagesExportPostParamsWithHTTPClient(client *http.Client) *PcloudCloudinstancesImagesExportPostParams {
	return &PcloudCloudinstancesImagesExportPostParams{
		HTTPClient: client,
	}
}

/*
PcloudCloudinstancesImagesExportPostParams contains all the parameters to send to the API endpoint

	for the pcloud cloudinstances images export post operation.

	Typically these are written to a http.Request.
*/
type PcloudCloudinstancesImagesExportPostParams struct {

	/* Body.

	   Parameters for exporting an image
	*/
	Body *models.ExportImage

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* ImageID.

	   Image ID of a image
	*/
	ImageID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud cloudinstances images export post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudCloudinstancesImagesExportPostParams) WithDefaults() *PcloudCloudinstancesImagesExportPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud cloudinstances images export post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudCloudinstancesImagesExportPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud cloudinstances images export post params
func (o *PcloudCloudinstancesImagesExportPostParams) WithTimeout(timeout time.Duration) *PcloudCloudinstancesImagesExportPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud cloudinstances images export post params
func (o *PcloudCloudinstancesImagesExportPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud cloudinstances images export post params
func (o *PcloudCloudinstancesImagesExportPostParams) WithContext(ctx context.Context) *PcloudCloudinstancesImagesExportPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud cloudinstances images export post params
func (o *PcloudCloudinstancesImagesExportPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud cloudinstances images export post params
func (o *PcloudCloudinstancesImagesExportPostParams) WithHTTPClient(client *http.Client) *PcloudCloudinstancesImagesExportPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud cloudinstances images export post params
func (o *PcloudCloudinstancesImagesExportPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud cloudinstances images export post params
func (o *PcloudCloudinstancesImagesExportPostParams) WithBody(body *models.ExportImage) *PcloudCloudinstancesImagesExportPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud cloudinstances images export post params
func (o *PcloudCloudinstancesImagesExportPostParams) SetBody(body *models.ExportImage) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud cloudinstances images export post params
func (o *PcloudCloudinstancesImagesExportPostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudCloudinstancesImagesExportPostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud cloudinstances images export post params
func (o *PcloudCloudinstancesImagesExportPostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithImageID adds the imageID to the pcloud cloudinstances images export post params
func (o *PcloudCloudinstancesImagesExportPostParams) WithImageID(imageID string) *PcloudCloudinstancesImagesExportPostParams {
	o.SetImageID(imageID)
	return o
}

// SetImageID adds the imageId to the pcloud cloudinstances images export post params
func (o *PcloudCloudinstancesImagesExportPostParams) SetImageID(imageID string) {
	o.ImageID = imageID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudCloudinstancesImagesExportPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param image_id
	if err := r.SetPathParam("image_id", o.ImageID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
