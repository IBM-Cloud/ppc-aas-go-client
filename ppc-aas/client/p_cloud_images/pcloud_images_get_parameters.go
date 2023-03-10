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
)

// NewPcloudImagesGetParams creates a new PcloudImagesGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudImagesGetParams() *PcloudImagesGetParams {
	return &PcloudImagesGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudImagesGetParamsWithTimeout creates a new PcloudImagesGetParams object
// with the ability to set a timeout on a request.
func NewPcloudImagesGetParamsWithTimeout(timeout time.Duration) *PcloudImagesGetParams {
	return &PcloudImagesGetParams{
		timeout: timeout,
	}
}

// NewPcloudImagesGetParamsWithContext creates a new PcloudImagesGetParams object
// with the ability to set a context for a request.
func NewPcloudImagesGetParamsWithContext(ctx context.Context) *PcloudImagesGetParams {
	return &PcloudImagesGetParams{
		Context: ctx,
	}
}

// NewPcloudImagesGetParamsWithHTTPClient creates a new PcloudImagesGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudImagesGetParamsWithHTTPClient(client *http.Client) *PcloudImagesGetParams {
	return &PcloudImagesGetParams{
		HTTPClient: client,
	}
}

/*
PcloudImagesGetParams contains all the parameters to send to the API endpoint

	for the pcloud images get operation.

	Typically these are written to a http.Request.
*/
type PcloudImagesGetParams struct {

	/* ImageID.

	   Image ID of a image
	*/
	ImageID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud images get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudImagesGetParams) WithDefaults() *PcloudImagesGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud images get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudImagesGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud images get params
func (o *PcloudImagesGetParams) WithTimeout(timeout time.Duration) *PcloudImagesGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud images get params
func (o *PcloudImagesGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud images get params
func (o *PcloudImagesGetParams) WithContext(ctx context.Context) *PcloudImagesGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud images get params
func (o *PcloudImagesGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud images get params
func (o *PcloudImagesGetParams) WithHTTPClient(client *http.Client) *PcloudImagesGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud images get params
func (o *PcloudImagesGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImageID adds the imageID to the pcloud images get params
func (o *PcloudImagesGetParams) WithImageID(imageID string) *PcloudImagesGetParams {
	o.SetImageID(imageID)
	return o
}

// SetImageID adds the imageId to the pcloud images get params
func (o *PcloudImagesGetParams) SetImageID(imageID string) {
	o.ImageID = imageID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudImagesGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param image_id
	if err := r.SetPathParam("image_id", o.ImageID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
