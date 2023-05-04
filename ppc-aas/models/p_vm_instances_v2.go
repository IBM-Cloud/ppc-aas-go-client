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

// PVMInstancesV2 p VM instances v2
//
// swagger:model PVMInstancesV2
type PVMInstancesV2 struct {

	// PVM Instance References
	// Required: true
	PvmInstances []*PVMInstanceReferenceV2 `json:"pvmInstances"`
}

// Validate validates this p VM instances v2
func (m *PVMInstancesV2) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePvmInstances(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PVMInstancesV2) validatePvmInstances(formats strfmt.Registry) error {

	if err := validate.Required("pvmInstances", "body", m.PvmInstances); err != nil {
		return err
	}

	for i := 0; i < len(m.PvmInstances); i++ {
		if swag.IsZero(m.PvmInstances[i]) { // not required
			continue
		}

		if m.PvmInstances[i] != nil {
			if err := m.PvmInstances[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pvmInstances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pvmInstances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this p VM instances v2 based on the context it is used
func (m *PVMInstancesV2) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePvmInstances(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PVMInstancesV2) contextValidatePvmInstances(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PvmInstances); i++ {

		if m.PvmInstances[i] != nil {
			if err := m.PvmInstances[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("pvmInstances" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("pvmInstances" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PVMInstancesV2) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PVMInstancesV2) UnmarshalBinary(b []byte) error {
	var res PVMInstancesV2
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
