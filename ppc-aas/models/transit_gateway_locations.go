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
	"github.com/go-openapi/validate"
)

// TransitGatewayLocations transit gateway locations
//
// swagger:model TransitGatewayLocations
type TransitGatewayLocations struct {

	// The List of PER enabled Power Private Cloud Service Locations
	// Required: true
	TransitGatewayLocations []*TransitGatewayLocation `json:"transitGatewayLocations"`
}

// Validate validates this transit gateway locations
func (m *TransitGatewayLocations) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTransitGatewayLocations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransitGatewayLocations) validateTransitGatewayLocations(formats strfmt.Registry) error {

	if err := validate.Required("transitGatewayLocations", "body", m.TransitGatewayLocations); err != nil {
		return err
	}

	for i := 0; i < len(m.TransitGatewayLocations); i++ {
		if swag.IsZero(m.TransitGatewayLocations[i]) { // not required
			continue
		}

		if m.TransitGatewayLocations[i] != nil {
			if err := m.TransitGatewayLocations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("transitGatewayLocations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("transitGatewayLocations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this transit gateway locations based on the context it is used
func (m *TransitGatewayLocations) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTransitGatewayLocations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TransitGatewayLocations) contextValidateTransitGatewayLocations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.TransitGatewayLocations); i++ {

		if m.TransitGatewayLocations[i] != nil {

			if swag.IsZero(m.TransitGatewayLocations[i]) { // not required
				return nil
			}

			if err := m.TransitGatewayLocations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("transitGatewayLocations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("transitGatewayLocations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *TransitGatewayLocations) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransitGatewayLocations) UnmarshalBinary(b []byte) error {
	var res TransitGatewayLocations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
