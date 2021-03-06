// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PostUniverseIdsOKBodyStationsItems post_universe_ids_station
//
// station object
// swagger:model postUniverseIdsOKBodyStationsItems
type PostUniverseIdsOKBodyStationsItems struct {

	// post_universe_ids_station_id
	//
	// id integer
	ID int32 `json:"id,omitempty"`

	// post_universe_ids_station_name
	//
	// name string
	Name string `json:"name,omitempty"`
}

// Validate validates this post universe ids o k body stations items
func (m *PostUniverseIdsOKBodyStationsItems) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostUniverseIdsOKBodyStationsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostUniverseIdsOKBodyStationsItems) UnmarshalBinary(b []byte) error {
	var res PostUniverseIdsOKBodyStationsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
