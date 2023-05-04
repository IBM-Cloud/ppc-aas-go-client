// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TransitGatewayInstance transit gateway instance
//
// swagger:model TransitGatewayInstance
type TransitGatewayInstance struct {

	// errors
	Errors []*TransitConnectionErrorItem `json:"errors,omitempty"`

	// IBM Resource Group ID associated with the Power Private Cloud Service Instance
	// Example: 2bf1887bf5c947b1966de2bd88220489
	ResourceGroupID string `json:"resourceGroupId,omitempty"`

	// The route distinguisher for a network
	// Example: 47902:3255120092
	RouteDistinguisher string `json:"routeDistinguisher,omitempty"`

	// The route target for a network
	// Example: 47902:3255120092
	RouteTarget string `json:"routeTarget,omitempty"`

	// The Power Private Cloud Service Instance CRN
	// Example: crn:v1:bluemix:public:ppc-aas:dal12:a/2bc3df23c0d14ebe921397bd8aa2573a:3a5798f1-4d2b-4e0a-9311-9b0fd6b94698::
	ServiceCrn string `json:"serviceCrn,omitempty"`

	// The Power Private Cloud Service Instance ID
	// Example: 3a5798f1-4d2b-4e0a-9311-9b0fd6b94698
	ServiceID string `json:"serviceId,omitempty"`

	// The trace id for debugging purposes
	Trace string `json:"trace,omitempty"`
}

// Validate validates this transit gateway instance
func (m *TransitGatewayInstance) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransitGatewayInstance) validateErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.Errors) { // not required
		return nil
	}

	for i := 0; i < len(m.Errors); i++ {
		if swag.IsZero(m.Errors[i]) { // not required
			continue
		}

		if m.Errors[i] != nil {
			if err := m.Errors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this transit gateway instance based on the context it is used
func (m *TransitGatewayInstance) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransitGatewayInstance) contextValidateErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Errors); i++ {

		if m.Errors[i] != nil {
			if err := m.Errors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("errors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("errors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransitGatewayInstance) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransitGatewayInstance) UnmarshalBinary(b []byte) error {
	var res TransitGatewayInstance
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
