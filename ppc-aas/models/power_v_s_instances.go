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

// PowerVSInstances The List of Power Private Cloud Instances for a specific IBM Cloud Account
//
// swagger:model PowerVSInstances
type PowerVSInstances struct {

	// power vs instances
	// Required: true
	PowerVsInstances []*PowerVSInstance `json:"powerVsInstances"`
}

// Validate validates this power v s instances
func (m *PowerVSInstances) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePowerVsInstances(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PowerVSInstances) validatePowerVsInstances(formats strfmt.Registry) error {

	if err := validate.Required("powerVsInstances", "body", m.PowerVsInstances); err != nil {
		return err
	}

	for i := 0; i < len(m.PowerVsInstances); i++ {
		if swag.IsZero(m.PowerVsInstances[i]) { // not required
			continue
		}

		if m.PowerVsInstances[i] != nil {
			if err := m.PowerVsInstances[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("powerVsInstances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("powerVsInstances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this power v s instances based on the context it is used
func (m *PowerVSInstances) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePowerVsInstances(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PowerVSInstances) contextValidatePowerVsInstances(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PowerVsInstances); i++ {

		if m.PowerVsInstances[i] != nil {

			if swag.IsZero(m.PowerVsInstances[i]) { // not required
				return nil
			}

			if err := m.PowerVsInstances[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("powerVsInstances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("powerVsInstances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PowerVSInstances) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PowerVSInstances) UnmarshalBinary(b []byte) error {
	var res PowerVSInstances
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
