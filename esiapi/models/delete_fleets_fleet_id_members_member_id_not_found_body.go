// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DeleteFleetsFleetIDMembersMemberIDNotFoundBody delete_fleets_fleet_id_members_member_id_not_found
//
// Not found
// swagger:model deleteFleetsFleetIdMembersMemberIdNotFoundBody
type DeleteFleetsFleetIDMembersMemberIDNotFoundBody struct {

	// delete_fleets_fleet_id_members_member_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this delete fleets fleet Id members member Id not found body
func (m *DeleteFleetsFleetIDMembersMemberIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DeleteFleetsFleetIDMembersMemberIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DeleteFleetsFleetIDMembersMemberIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res DeleteFleetsFleetIDMembersMemberIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
