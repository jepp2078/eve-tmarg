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

// GetUniverseSystemJumpsOKBodyItems get_universe_system_jumps_200_ok
//
// 200 ok object
// swagger:model getUniverseSystemJumpsOKBodyItems
type GetUniverseSystemJumpsOKBodyItems struct {

	// get_universe_system_jumps_ship_jumps
	//
	// ship_jumps integer
	// Required: true
	ShipJumps *int32 `json:"ship_jumps"`

	// get_universe_system_jumps_system_id
	//
	// system_id integer
	// Required: true
	SystemID *int32 `json:"system_id"`
}

// Validate validates this get universe system jumps o k body items
func (m *GetUniverseSystemJumpsOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateShipJumps(formats); err != nil {
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

func (m *GetUniverseSystemJumpsOKBodyItems) validateShipJumps(formats strfmt.Registry) error {

	if err := validate.Required("ship_jumps", "body", m.ShipJumps); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseSystemJumpsOKBodyItems) validateSystemID(formats strfmt.Registry) error {

	if err := validate.Required("system_id", "body", m.SystemID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetUniverseSystemJumpsOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUniverseSystemJumpsOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetUniverseSystemJumpsOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
