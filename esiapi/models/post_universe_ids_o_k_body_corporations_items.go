// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PostUniverseIdsOKBodyCorporationsItems post_universe_ids_corporation
//
// corporation object
// swagger:model postUniverseIdsOKBodyCorporationsItems
type PostUniverseIdsOKBodyCorporationsItems struct {

	// post_universe_ids_corporation_id
	//
	// id integer
	ID int32 `json:"id,omitempty"`

	// post_universe_ids_corporation_name
	//
	// name string
	Name string `json:"name,omitempty"`
}

// Validate validates this post universe ids o k body corporations items
func (m *PostUniverseIdsOKBodyCorporationsItems) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostUniverseIdsOKBodyCorporationsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostUniverseIdsOKBodyCorporationsItems) UnmarshalBinary(b []byte) error {
	var res PostUniverseIdsOKBodyCorporationsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
