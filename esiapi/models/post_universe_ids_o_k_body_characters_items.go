// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PostUniverseIdsOKBodyCharactersItems post_universe_ids_character
//
// character object
// swagger:model postUniverseIdsOKBodyCharactersItems
type PostUniverseIdsOKBodyCharactersItems struct {

	// post_universe_ids_character_id
	//
	// id integer
	ID int32 `json:"id,omitempty"`

	// post_universe_ids_character_name
	//
	// name string
	Name string `json:"name,omitempty"`
}

// Validate validates this post universe ids o k body characters items
func (m *PostUniverseIdsOKBodyCharactersItems) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostUniverseIdsOKBodyCharactersItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostUniverseIdsOKBodyCharactersItems) UnmarshalBinary(b []byte) error {
	var res PostUniverseIdsOKBodyCharactersItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
