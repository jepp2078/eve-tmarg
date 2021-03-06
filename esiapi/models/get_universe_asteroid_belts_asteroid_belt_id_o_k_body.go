// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetUniverseAsteroidBeltsAsteroidBeltIDOKBody get_universe_asteroid_belts_asteroid_belt_id_ok
//
// 200 ok object
// swagger:model getUniverseAsteroidBeltsAsteroidBeltIdOKBody
type GetUniverseAsteroidBeltsAsteroidBeltIDOKBody struct {

	// get_universe_asteroid_belts_asteroid_belt_id_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// position
	// Required: true
	Position *GetUniverseAsteroidBeltsAsteroidBeltIDOKBodyPosition `json:"position"`

	// get_universe_asteroid_belts_asteroid_belt_id_system_id
	//
	// The solar system this asteroid belt is in
	// Required: true
	SystemID *int32 `json:"system_id"`
}

// Validate validates this get universe asteroid belts asteroid belt Id o k body
func (m *GetUniverseAsteroidBeltsAsteroidBeltIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSystemID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetUniverseAsteroidBeltsAsteroidBeltIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseAsteroidBeltsAsteroidBeltIDOKBody) validatePosition(formats strfmt.Registry) error {

	if err := validate.Required("position", "body", m.Position); err != nil {
		return err
	}

	if m.Position != nil {
		if err := m.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

func (m *GetUniverseAsteroidBeltsAsteroidBeltIDOKBody) validateSystemID(formats strfmt.Registry) error {

	if err := validate.Required("system_id", "body", m.SystemID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetUniverseAsteroidBeltsAsteroidBeltIDOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUniverseAsteroidBeltsAsteroidBeltIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseAsteroidBeltsAsteroidBeltIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
