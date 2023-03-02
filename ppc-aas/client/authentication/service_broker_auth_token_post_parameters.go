// Code generated by go-swagger; DO NOT EDIT.

package authentication

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

// NewServiceBrokerAuthTokenPostParams creates a new ServiceBrokerAuthTokenPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewServiceBrokerAuthTokenPostParams() *ServiceBrokerAuthTokenPostParams {
	return &ServiceBrokerAuthTokenPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewServiceBrokerAuthTokenPostParamsWithTimeout creates a new ServiceBrokerAuthTokenPostParams object
// with the ability to set a timeout on a request.
func NewServiceBrokerAuthTokenPostParamsWithTimeout(timeout time.Duration) *ServiceBrokerAuthTokenPostParams {
	return &ServiceBrokerAuthTokenPostParams{
		timeout: timeout,
	}
}

// NewServiceBrokerAuthTokenPostParamsWithContext creates a new ServiceBrokerAuthTokenPostParams object
// with the ability to set a context for a request.
func NewServiceBrokerAuthTokenPostParamsWithContext(ctx context.Context) *ServiceBrokerAuthTokenPostParams {
	return &ServiceBrokerAuthTokenPostParams{
		Context: ctx,
	}
}

// NewServiceBrokerAuthTokenPostParamsWithHTTPClient creates a new ServiceBrokerAuthTokenPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewServiceBrokerAuthTokenPostParamsWithHTTPClient(client *http.Client) *ServiceBrokerAuthTokenPostParams {
	return &ServiceBrokerAuthTokenPostParams{
		HTTPClient: client,
	}
}

/*
ServiceBrokerAuthTokenPostParams contains all the parameters to send to the API endpoint

	for the service broker auth token post operation.

	Typically these are written to a http.Request.
*/
type ServiceBrokerAuthTokenPostParams struct {

	/* Body.

	   Parameters for requesting a new Token from a Refresh Token
	*/
	Body *models.TokenRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the service broker auth token post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceBrokerAuthTokenPostParams) WithDefaults() *ServiceBrokerAuthTokenPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the service broker auth token post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ServiceBrokerAuthTokenPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the service broker auth token post params
func (o *ServiceBrokerAuthTokenPostParams) WithTimeout(timeout time.Duration) *ServiceBrokerAuthTokenPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the service broker auth token post params
func (o *ServiceBrokerAuthTokenPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the service broker auth token post params
func (o *ServiceBrokerAuthTokenPostParams) WithContext(ctx context.Context) *ServiceBrokerAuthTokenPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the service broker auth token post params
func (o *ServiceBrokerAuthTokenPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the service broker auth token post params
func (o *ServiceBrokerAuthTokenPostParams) WithHTTPClient(client *http.Client) *ServiceBrokerAuthTokenPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the service broker auth token post params
func (o *ServiceBrokerAuthTokenPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the service broker auth token post params
func (o *ServiceBrokerAuthTokenPostParams) WithBody(body *models.TokenRequest) *ServiceBrokerAuthTokenPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the service broker auth token post params
func (o *ServiceBrokerAuthTokenPostParams) SetBody(body *models.TokenRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ServiceBrokerAuthTokenPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
