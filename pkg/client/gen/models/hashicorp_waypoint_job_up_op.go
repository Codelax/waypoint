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

// HashicorpWaypointJobUpOp UpOp runs the "waypoint up" logic: it does a build (with push),
// deploy, and release all in one. The results for each child operation
// will be set directly on the Release message (i.e. "build" will be
// populated in addition to "up").
//
// swagger:model hashicorp.waypoint.Job.UpOp
type HashicorpWaypointJobUpOp struct {

	// Options for the release stage. The "deployment" field in this will
	// be ignored since we'll always use the deployment from the deploy
	// step in Up.
	Release *HashicorpWaypointJobReleaseOp `json:"release,omitempty"`
}

// Validate validates this hashicorp waypoint job up op
func (m *HashicorpWaypointJobUpOp) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRelease(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointJobUpOp) validateRelease(formats strfmt.Registry) error {
	if swag.IsZero(m.Release) { // not required
		return nil
	}

	if m.Release != nil {
		if err := m.Release.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("release")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("release")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this hashicorp waypoint job up op based on the context it is used
func (m *HashicorpWaypointJobUpOp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateRelease(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *HashicorpWaypointJobUpOp) contextValidateRelease(ctx context.Context, formats strfmt.Registry) error {

	if m.Release != nil {
		if err := m.Release.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("release")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("release")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *HashicorpWaypointJobUpOp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *HashicorpWaypointJobUpOp) UnmarshalBinary(b []byte) error {
	var res HashicorpWaypointJobUpOp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}