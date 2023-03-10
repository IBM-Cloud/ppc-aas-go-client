// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_events

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

// NewPcloudEventsGetqueryParams creates a new PcloudEventsGetqueryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudEventsGetqueryParams() *PcloudEventsGetqueryParams {
	return &PcloudEventsGetqueryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudEventsGetqueryParamsWithTimeout creates a new PcloudEventsGetqueryParams object
// with the ability to set a timeout on a request.
func NewPcloudEventsGetqueryParamsWithTimeout(timeout time.Duration) *PcloudEventsGetqueryParams {
	return &PcloudEventsGetqueryParams{
		timeout: timeout,
	}
}

// NewPcloudEventsGetqueryParamsWithContext creates a new PcloudEventsGetqueryParams object
// with the ability to set a context for a request.
func NewPcloudEventsGetqueryParamsWithContext(ctx context.Context) *PcloudEventsGetqueryParams {
	return &PcloudEventsGetqueryParams{
		Context: ctx,
	}
}

// NewPcloudEventsGetqueryParamsWithHTTPClient creates a new PcloudEventsGetqueryParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudEventsGetqueryParamsWithHTTPClient(client *http.Client) *PcloudEventsGetqueryParams {
	return &PcloudEventsGetqueryParams{
		HTTPClient: client,
	}
}

/*
PcloudEventsGetqueryParams contains all the parameters to send to the API endpoint

	for the pcloud events getquery operation.

	Typically these are written to a http.Request.
*/
type PcloudEventsGetqueryParams struct {

	/* AcceptLanguage.

	   The language requested for the return document
	*/
	AcceptLanguage *string

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* FromTime.

	   A from query time in either ISO 8601 or unix epoch format
	*/
	FromTime *string

	/* Time.

	   (deprecated - use from_time) A time in either ISO 8601 or unix epoch format
	*/
	Time *string

	/* ToTime.

	   A to query time in either ISO 8601 or unix epoch format
	*/
	ToTime *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud events getquery params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudEventsGetqueryParams) WithDefaults() *PcloudEventsGetqueryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud events getquery params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudEventsGetqueryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud events getquery params
func (o *PcloudEventsGetqueryParams) WithTimeout(timeout time.Duration) *PcloudEventsGetqueryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud events getquery params
func (o *PcloudEventsGetqueryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud events getquery params
func (o *PcloudEventsGetqueryParams) WithContext(ctx context.Context) *PcloudEventsGetqueryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud events getquery params
func (o *PcloudEventsGetqueryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud events getquery params
func (o *PcloudEventsGetqueryParams) WithHTTPClient(client *http.Client) *PcloudEventsGetqueryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud events getquery params
func (o *PcloudEventsGetqueryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAcceptLanguage adds the acceptLanguage to the pcloud events getquery params
func (o *PcloudEventsGetqueryParams) WithAcceptLanguage(acceptLanguage *string) *PcloudEventsGetqueryParams {
	o.SetAcceptLanguage(acceptLanguage)
	return o
}

// SetAcceptLanguage adds the acceptLanguage to the pcloud events getquery params
func (o *PcloudEventsGetqueryParams) SetAcceptLanguage(acceptLanguage *string) {
	o.AcceptLanguage = acceptLanguage
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud events getquery params
func (o *PcloudEventsGetqueryParams) WithCloudInstanceID(cloudInstanceID string) *PcloudEventsGetqueryParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud events getquery params
func (o *PcloudEventsGetqueryParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithFromTime adds the fromTime to the pcloud events getquery params
func (o *PcloudEventsGetqueryParams) WithFromTime(fromTime *string) *PcloudEventsGetqueryParams {
	o.SetFromTime(fromTime)
	return o
}

// SetFromTime adds the fromTime to the pcloud events getquery params
func (o *PcloudEventsGetqueryParams) SetFromTime(fromTime *string) {
	o.FromTime = fromTime
}

// WithTime adds the time to the pcloud events getquery params
func (o *PcloudEventsGetqueryParams) WithTime(time *string) *PcloudEventsGetqueryParams {
	o.SetTime(time)
	return o
}

// SetTime adds the time to the pcloud events getquery params
func (o *PcloudEventsGetqueryParams) SetTime(time *string) {
	o.Time = time
}

// WithToTime adds the toTime to the pcloud events getquery params
func (o *PcloudEventsGetqueryParams) WithToTime(toTime *string) *PcloudEventsGetqueryParams {
	o.SetToTime(toTime)
	return o
}

// SetToTime adds the toTime to the pcloud events getquery params
func (o *PcloudEventsGetqueryParams) SetToTime(toTime *string) {
	o.ToTime = toTime
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudEventsGetqueryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AcceptLanguage != nil {

		// header param Accept-Language
		if err := r.SetHeaderParam("Accept-Language", *o.AcceptLanguage); err != nil {
			return err
		}
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if o.FromTime != nil {

		// query param from_time
		var qrFromTime string

		if o.FromTime != nil {
			qrFromTime = *o.FromTime
		}
		qFromTime := qrFromTime
		if qFromTime != "" {

			if err := r.SetQueryParam("from_time", qFromTime); err != nil {
				return err
			}
		}
	}

	if o.Time != nil {

		// query param time
		var qrTime string

		if o.Time != nil {
			qrTime = *o.Time
		}
		qTime := qrTime
		if qTime != "" {

			if err := r.SetQueryParam("time", qTime); err != nil {
				return err
			}
		}
	}

	if o.ToTime != nil {

		// query param to_time
		var qrToTime string

		if o.ToTime != nil {
			qrToTime = *o.ToTime
		}
		qToTime := qrToTime
		if qToTime != "" {

			if err := r.SetQueryParam("to_time", qToTime); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
