// Code generated by go-swagger; DO NOT EDIT.

package kbmodel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CatalogValidation catalog validation
//
// swagger:model CatalogValidation
type CatalogValidation struct {

	// catalog validation errors
	CatalogValidationErrors []*CatalogValidationError `json:"catalogValidationErrors"`
}

// Validate validates this catalog validation
func (m *CatalogValidation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCatalogValidationErrors(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogValidation) validateCatalogValidationErrors(formats strfmt.Registry) error {
	if swag.IsZero(m.CatalogValidationErrors) { // not required
		return nil
	}

	for i := 0; i < len(m.CatalogValidationErrors); i++ {
		if swag.IsZero(m.CatalogValidationErrors[i]) { // not required
			continue
		}

		if m.CatalogValidationErrors[i] != nil {
			if err := m.CatalogValidationErrors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("catalogValidationErrors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("catalogValidationErrors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this catalog validation based on the context it is used
func (m *CatalogValidation) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCatalogValidationErrors(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CatalogValidation) contextValidateCatalogValidationErrors(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.CatalogValidationErrors); i++ {

		if m.CatalogValidationErrors[i] != nil {
			if err := m.CatalogValidationErrors[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("catalogValidationErrors" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("catalogValidationErrors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CatalogValidation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CatalogValidation) UnmarshalBinary(b []byte) error {
	var res CatalogValidation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
