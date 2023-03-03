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

// NodeCommand node command
//
// swagger:model NodeCommand
type NodeCommand struct {

	// is system command type
	IsSystemCommandType bool `json:"isSystemCommandType,omitempty"`

	// node command properties
	NodeCommandProperties []*NodeCommandProperty `json:"nodeCommandProperties"`

	// node command type
	NodeCommandType string `json:"nodeCommandType,omitempty"`
}

// Validate validates this node command
func (m *NodeCommand) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNodeCommandProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeCommand) validateNodeCommandProperties(formats strfmt.Registry) error {
	if swag.IsZero(m.NodeCommandProperties) { // not required
		return nil
	}

	for i := 0; i < len(m.NodeCommandProperties); i++ {
		if swag.IsZero(m.NodeCommandProperties[i]) { // not required
			continue
		}

		if m.NodeCommandProperties[i] != nil {
			if err := m.NodeCommandProperties[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodeCommandProperties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodeCommandProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this node command based on the context it is used
func (m *NodeCommand) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNodeCommandProperties(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NodeCommand) contextValidateNodeCommandProperties(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.NodeCommandProperties); i++ {

		if m.NodeCommandProperties[i] != nil {
			if err := m.NodeCommandProperties[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nodeCommandProperties" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("nodeCommandProperties" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NodeCommand) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeCommand) UnmarshalBinary(b []byte) error {
	var res NodeCommand
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
