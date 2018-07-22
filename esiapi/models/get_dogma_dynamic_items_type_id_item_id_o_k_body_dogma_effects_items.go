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

// GetDogmaDynamicItemsTypeIDItemIDOKBodyDogmaEffectsItems get_dogma_dynamic_items_type_id_item_id_dogma_effect
//
// dogma_effect object
// swagger:model getDogmaDynamicItemsTypeIdItemIdOKBodyDogmaEffectsItems
type GetDogmaDynamicItemsTypeIDItemIDOKBodyDogmaEffectsItems struct {

	// get_dogma_dynamic_items_type_id_item_id_effect_id
	//
	// effect_id integer
	// Required: true
	EffectID *int32 `json:"effect_id"`

	// get_dogma_dynamic_items_type_id_item_id_is_default
	//
	// is_default boolean
	// Required: true
	IsDefault *bool `json:"is_default"`
}

// Validate validates this get dogma dynamic items type Id item Id o k body dogma effects items
func (m *GetDogmaDynamicItemsTypeIDItemIDOKBodyDogmaEffectsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEffectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsDefault(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetDogmaDynamicItemsTypeIDItemIDOKBodyDogmaEffectsItems) validateEffectID(formats strfmt.Registry) error {

	if err := validate.Required("effect_id", "body", m.EffectID); err != nil {
		return err
	}

	return nil
}

func (m *GetDogmaDynamicItemsTypeIDItemIDOKBodyDogmaEffectsItems) validateIsDefault(formats strfmt.Registry) error {

	if err := validate.Required("is_default", "body", m.IsDefault); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetDogmaDynamicItemsTypeIDItemIDOKBodyDogmaEffectsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetDogmaDynamicItemsTypeIDItemIDOKBodyDogmaEffectsItems) UnmarshalBinary(b []byte) error {
	var res GetDogmaDynamicItemsTypeIDItemIDOKBodyDogmaEffectsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}