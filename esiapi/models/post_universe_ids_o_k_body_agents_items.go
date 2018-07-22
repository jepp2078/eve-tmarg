// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PostUniverseIdsOKBodyAgentsItems post_universe_ids_agent
//
// agent object
// swagger:model postUniverseIdsOKBodyAgentsItems
type PostUniverseIdsOKBodyAgentsItems struct {

	// post_universe_ids_id
	//
	// id integer
	ID int32 `json:"id,omitempty"`

	// post_universe_ids_name
	//
	// name string
	Name string `json:"name,omitempty"`
}

// Validate validates this post universe ids o k body agents items
func (m *PostUniverseIdsOKBodyAgentsItems) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PostUniverseIdsOKBodyAgentsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostUniverseIdsOKBodyAgentsItems) UnmarshalBinary(b []byte) error {
	var res PostUniverseIdsOKBodyAgentsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}