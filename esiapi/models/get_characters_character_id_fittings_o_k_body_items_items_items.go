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

// GetCharactersCharacterIDFittingsOKBodyItemsItemsItems get_characters_character_id_fittings_item
//
// item object
// swagger:model getCharactersCharacterIdFittingsOKBodyItemsItemsItems
type GetCharactersCharacterIDFittingsOKBodyItemsItemsItems struct {

	// get_characters_character_id_fittings_flag
	//
	// flag integer
	// Required: true
	Flag *int32 `json:"flag"`

	// get_characters_character_id_fittings_quantity
	//
	// quantity integer
	// Required: true
	Quantity *int32 `json:"quantity"`

	// get_characters_character_id_fittings_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get characters character Id fittings o k body items items items
func (m *GetCharactersCharacterIDFittingsOKBodyItemsItemsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFlag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDFittingsOKBodyItemsItemsItems) validateFlag(formats strfmt.Registry) error {

	if err := validate.Required("flag", "body", m.Flag); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDFittingsOKBodyItemsItemsItems) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDFittingsOKBodyItemsItemsItems) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", m.TypeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDFittingsOKBodyItemsItemsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDFittingsOKBodyItemsItemsItems) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDFittingsOKBodyItemsItemsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
