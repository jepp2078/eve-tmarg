// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetAlliancesAllianceIDNotFoundBody get_alliances_alliance_id_not_found
//
// Not found
// swagger:model getAlliancesAllianceIdNotFoundBody
type GetAlliancesAllianceIDNotFoundBody struct {

	// get_alliances_alliance_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get alliances alliance Id not found body
func (m *GetAlliancesAllianceIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetAlliancesAllianceIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetAlliancesAllianceIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetAlliancesAllianceIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
