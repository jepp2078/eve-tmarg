// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostUniverseNamesOKBodyItems post_universe_names_200_ok
//
// 200 ok object
// swagger:model postUniverseNamesOKBodyItems
type PostUniverseNamesOKBodyItems struct {

	// post_universe_names_category
	//
	// category string
	// Required: true
	// Enum: [alliance character constellation corporation inventory_type region solar_system station]
	Category *string `json:"category"`

	// post_universe_names_id
	//
	// id integer
	// Required: true
	ID *int32 `json:"id"`

	// post_universe_names_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this post universe names o k body items
func (m *PostUniverseNamesOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var postUniverseNamesOKBodyItemsTypeCategoryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["alliance","character","constellation","corporation","inventory_type","region","solar_system","station"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postUniverseNamesOKBodyItemsTypeCategoryPropEnum = append(postUniverseNamesOKBodyItemsTypeCategoryPropEnum, v)
	}
}

const (

	// PostUniverseNamesOKBodyItemsCategoryAlliance captures enum value "alliance"
	PostUniverseNamesOKBodyItemsCategoryAlliance string = "alliance"

	// PostUniverseNamesOKBodyItemsCategoryCharacter captures enum value "character"
	PostUniverseNamesOKBodyItemsCategoryCharacter string = "character"

	// PostUniverseNamesOKBodyItemsCategoryConstellation captures enum value "constellation"
	PostUniverseNamesOKBodyItemsCategoryConstellation string = "constellation"

	// PostUniverseNamesOKBodyItemsCategoryCorporation captures enum value "corporation"
	PostUniverseNamesOKBodyItemsCategoryCorporation string = "corporation"

	// PostUniverseNamesOKBodyItemsCategoryInventoryType captures enum value "inventory_type"
	PostUniverseNamesOKBodyItemsCategoryInventoryType string = "inventory_type"

	// PostUniverseNamesOKBodyItemsCategoryRegion captures enum value "region"
	PostUniverseNamesOKBodyItemsCategoryRegion string = "region"

	// PostUniverseNamesOKBodyItemsCategorySolarSystem captures enum value "solar_system"
	PostUniverseNamesOKBodyItemsCategorySolarSystem string = "solar_system"

	// PostUniverseNamesOKBodyItemsCategoryStation captures enum value "station"
	PostUniverseNamesOKBodyItemsCategoryStation string = "station"
)

// prop value enum
func (m *PostUniverseNamesOKBodyItems) validateCategoryEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, postUniverseNamesOKBodyItemsTypeCategoryPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PostUniverseNamesOKBodyItems) validateCategory(formats strfmt.Registry) error {

	if err := validate.Required("category", "body", m.Category); err != nil {
		return err
	}

	// value enum
	if err := m.validateCategoryEnum("category", "body", *m.Category); err != nil {
		return err
	}

	return nil
}

func (m *PostUniverseNamesOKBodyItems) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *PostUniverseNamesOKBodyItems) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostUniverseNamesOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostUniverseNamesOKBodyItems) UnmarshalBinary(b []byte) error {
	var res PostUniverseNamesOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
