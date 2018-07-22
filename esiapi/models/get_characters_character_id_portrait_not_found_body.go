// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetCharactersCharacterIDPortraitNotFoundBody get_characters_character_id_portrait_not_found
//
// No image server for this datasource
// swagger:model getCharactersCharacterIdPortraitNotFoundBody
type GetCharactersCharacterIDPortraitNotFoundBody struct {

	// get_characters_character_id_portrait_error
	//
	// error message
	Error string `json:"error,omitempty"`
}

// Validate validates this get characters character Id portrait not found body
func (m *GetCharactersCharacterIDPortraitNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDPortraitNotFoundBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDPortraitNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDPortraitNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
