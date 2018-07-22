// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetDogmaAttributesAttributeIDNotFoundBody get_dogma_attributes_attribute_id_not_found
//
// Not found
// swagger:model getDogmaAttributesAttributeIdNotFoundBody
type GetDogmaAttributesAttributeIDNotFoundBody struct {

	// get_dogma_attributes_attribute_id_404_not_found
	//
	// Not found message
	Error string `json:"error,omitempty"`
}

// Validate validates this get dogma attributes attribute Id not found body
func (m *GetDogmaAttributesAttributeIDNotFoundBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetDogmaAttributesAttributeIDNotFoundBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetDogmaAttributesAttributeIDNotFoundBody) UnmarshalBinary(b []byte) error {
	var res GetDogmaAttributesAttributeIDNotFoundBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
