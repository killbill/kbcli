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
	"github.com/go-openapi/validate"
)

// BundleTimeline bundle timeline
//
// swagger:model BundleTimeline
type BundleTimeline struct {

	// account Id
	// Format: uuid
	AccountID strfmt.UUID `json:"accountId,omitempty"`

	// audit logs
	AuditLogs []*AuditLog `json:"auditLogs"`

	// bundle Id
	// Format: uuid
	BundleID strfmt.UUID `json:"bundleId,omitempty"`

	// events
	Events []*EventSubscription `json:"events"`

	// external key
	ExternalKey string `json:"externalKey,omitempty"`
}

// Validate validates this bundle timeline
func (m *BundleTimeline) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuditLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBundleID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BundleTimeline) validateAccountID(formats strfmt.Registry) error {
	if swag.IsZero(m.AccountID) { // not required
		return nil
	}

	if err := validate.FormatOf("accountId", "body", "uuid", m.AccountID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BundleTimeline) validateAuditLogs(formats strfmt.Registry) error {
	if swag.IsZero(m.AuditLogs) { // not required
		return nil
	}

	for i := 0; i < len(m.AuditLogs); i++ {
		if swag.IsZero(m.AuditLogs[i]) { // not required
			continue
		}

		if m.AuditLogs[i] != nil {
			if err := m.AuditLogs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BundleTimeline) validateBundleID(formats strfmt.Registry) error {
	if swag.IsZero(m.BundleID) { // not required
		return nil
	}

	if err := validate.FormatOf("bundleId", "body", "uuid", m.BundleID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BundleTimeline) validateEvents(formats strfmt.Registry) error {
	if swag.IsZero(m.Events) { // not required
		return nil
	}

	for i := 0; i < len(m.Events); i++ {
		if swag.IsZero(m.Events[i]) { // not required
			continue
		}

		if m.Events[i] != nil {
			if err := m.Events[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("events" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this bundle timeline based on the context it is used
func (m *BundleTimeline) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAuditLogs(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BundleTimeline) contextValidateAuditLogs(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.AuditLogs); i++ {

		if m.AuditLogs[i] != nil {
			if err := m.AuditLogs[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("auditLogs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BundleTimeline) contextValidateEvents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Events); i++ {

		if m.Events[i] != nil {
			if err := m.Events[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("events" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *BundleTimeline) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BundleTimeline) UnmarshalBinary(b []byte) error {
	var res BundleTimeline
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
