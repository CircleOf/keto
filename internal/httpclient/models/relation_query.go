// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// RelationQuery relation query
//
// swagger:model relationQuery
type RelationQuery struct {

	// Namespace to query
	Namespace string `json:"namespace,omitempty"`

	// Object to query
	Object string `json:"object,omitempty"`

	// Relation to query
	Relation string `json:"relation,omitempty"`

	// SubjectID to query
	//
	// Either SubjectSet or SubjectID can be provided.
	SubjectID string `json:"subject_id,omitempty"`

	// subject set
	SubjectSet *SubjectSet `json:"subject_set,omitempty"`
}

// Validate validates this relation query
func (m *RelationQuery) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSubjectSet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelationQuery) validateSubjectSet(formats strfmt.Registry) error {
	if swag.IsZero(m.SubjectSet) { // not required
		return nil
	}

	if m.SubjectSet != nil {
		if err := m.SubjectSet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subject_set")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subject_set")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this relation query based on the context it is used
func (m *RelationQuery) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSubjectSet(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RelationQuery) contextValidateSubjectSet(ctx context.Context, formats strfmt.Registry) error {

	if m.SubjectSet != nil {
		if err := m.SubjectSet.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subject_set")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("subject_set")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RelationQuery) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RelationQuery) UnmarshalBinary(b []byte) error {
	var res RelationQuery
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
