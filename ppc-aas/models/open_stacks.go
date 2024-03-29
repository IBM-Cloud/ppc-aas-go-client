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

// OpenStacks open stacks
//
// swagger:model OpenStacks
type OpenStacks struct {

	// OpenStacks managed by Power Private Cloud
	// Required: true
	OpenStacks []*OpenStack `json:"openStacks"`
}

// Validate validates this open stacks
func (m *OpenStacks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOpenStacks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenStacks) validateOpenStacks(formats strfmt.Registry) error {

	if err := validate.Required("openStacks", "body", m.OpenStacks); err != nil {
		return err
	}

	for i := 0; i < len(m.OpenStacks); i++ {
		if swag.IsZero(m.OpenStacks[i]) { // not required
			continue
		}

		if m.OpenStacks[i] != nil {
			if err := m.OpenStacks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("openStacks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("openStacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this open stacks based on the context it is used
func (m *OpenStacks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOpenStacks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OpenStacks) contextValidateOpenStacks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.OpenStacks); i++ {

		if m.OpenStacks[i] != nil {

			if swag.IsZero(m.OpenStacks[i]) { // not required
				return nil
			}

			if err := m.OpenStacks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("openStacks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("openStacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *OpenStacks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenStacks) UnmarshalBinary(b []byte) error {
	var res OpenStacks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
