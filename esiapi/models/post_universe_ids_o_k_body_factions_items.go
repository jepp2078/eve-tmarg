// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PostUniverseIdsOKBodyFactionsItems post_universe_ids_faction
//
// faction object
// swagger:model postUniverseIdsOKBodyFactionsItems
type PostUniverseIdsOKBodyFactionsItems struct {

	// post_universe_ids_faction_id
	//
	// id integer
	ID int32 `json:"id,omitempty"`

	// post_universe_ids_faction_name
	//
	// name string
	Name string `json:"name,omitempty"`
}

// Validate validates this post universe ids o k body factions items
func (m *PostUniverseIdsOKBodyFactionsItems) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostUniverseIdsOKBodyFactionsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostUniverseIdsOKBodyFactionsItems) UnmarshalBinary(b []byte) error {
	var res PostUniverseIdsOKBodyFactionsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
