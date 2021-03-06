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

// GetCharactersCharacterIDPlanetsOKBodyItems get_characters_character_id_planets_200_ok
//
// 200 ok object
// swagger:model getCharactersCharacterIdPlanetsOKBodyItems
type GetCharactersCharacterIDPlanetsOKBodyItems struct {

	// get_characters_character_id_planets_last_update
	//
	// last_update string
	// Required: true
	// Format: date-time
	LastUpdate *strfmt.DateTime `json:"last_update"`

	// get_characters_character_id_planets_num_pins
	//
	// num_pins integer
	// Required: true
	// Minimum: 1
	NumPins *int32 `json:"num_pins"`

	// get_characters_character_id_planets_owner_id
	//
	// owner_id integer
	// Required: true
	OwnerID *int32 `json:"owner_id"`

	// get_characters_character_id_planets_planet_id
	//
	// planet_id integer
	// Required: true
	PlanetID *int32 `json:"planet_id"`

	// get_characters_character_id_planets_planet_type
	//
	// planet_type string
	// Required: true
	// Enum: [temperate barren oceanic ice gas lava storm plasma]
	PlanetType *string `json:"planet_type"`

	// get_characters_character_id_planets_solar_system_id
	//
	// solar_system_id integer
	// Required: true
	SolarSystemID *int32 `json:"solar_system_id"`

	// get_characters_character_id_planets_upgrade_level
	//
	// upgrade_level integer
	// Required: true
	// Maximum: 5
	// Minimum: 0
	UpgradeLevel *int32 `json:"upgrade_level"`
}

// Validate validates this get characters character Id planets o k body items
func (m *GetCharactersCharacterIDPlanetsOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastUpdate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNumPins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanetID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanetType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSolarSystemID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpgradeLevel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDPlanetsOKBodyItems) validateLastUpdate(formats strfmt.Registry) error {

	if err := validate.Required("last_update", "body", m.LastUpdate); err != nil {
		return err
	}

	if err := validate.FormatOf("last_update", "body", "date-time", m.LastUpdate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDPlanetsOKBodyItems) validateNumPins(formats strfmt.Registry) error {

	if err := validate.Required("num_pins", "body", m.NumPins); err != nil {
		return err
	}

	if err := validate.MinimumInt("num_pins", "body", int64(*m.NumPins), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDPlanetsOKBodyItems) validateOwnerID(formats strfmt.Registry) error {

	if err := validate.Required("owner_id", "body", m.OwnerID); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDPlanetsOKBodyItems) validatePlanetID(formats strfmt.Registry) error {

	if err := validate.Required("planet_id", "body", m.PlanetID); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdPlanetsOKBodyItemsTypePlanetTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["temperate","barren","oceanic","ice","gas","lava","storm","plasma"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdPlanetsOKBodyItemsTypePlanetTypePropEnum = append(getCharactersCharacterIdPlanetsOKBodyItemsTypePlanetTypePropEnum, v)
	}
}

const (

	// GetCharactersCharacterIDPlanetsOKBodyItemsPlanetTypeTemperate captures enum value "temperate"
	GetCharactersCharacterIDPlanetsOKBodyItemsPlanetTypeTemperate string = "temperate"

	// GetCharactersCharacterIDPlanetsOKBodyItemsPlanetTypeBarren captures enum value "barren"
	GetCharactersCharacterIDPlanetsOKBodyItemsPlanetTypeBarren string = "barren"

	// GetCharactersCharacterIDPlanetsOKBodyItemsPlanetTypeOceanic captures enum value "oceanic"
	GetCharactersCharacterIDPlanetsOKBodyItemsPlanetTypeOceanic string = "oceanic"

	// GetCharactersCharacterIDPlanetsOKBodyItemsPlanetTypeIce captures enum value "ice"
	GetCharactersCharacterIDPlanetsOKBodyItemsPlanetTypeIce string = "ice"

	// GetCharactersCharacterIDPlanetsOKBodyItemsPlanetTypeGas captures enum value "gas"
	GetCharactersCharacterIDPlanetsOKBodyItemsPlanetTypeGas string = "gas"

	// GetCharactersCharacterIDPlanetsOKBodyItemsPlanetTypeLava captures enum value "lava"
	GetCharactersCharacterIDPlanetsOKBodyItemsPlanetTypeLava string = "lava"

	// GetCharactersCharacterIDPlanetsOKBodyItemsPlanetTypeStorm captures enum value "storm"
	GetCharactersCharacterIDPlanetsOKBodyItemsPlanetTypeStorm string = "storm"

	// GetCharactersCharacterIDPlanetsOKBodyItemsPlanetTypePlasma captures enum value "plasma"
	GetCharactersCharacterIDPlanetsOKBodyItemsPlanetTypePlasma string = "plasma"
)

// prop value enum
func (m *GetCharactersCharacterIDPlanetsOKBodyItems) validatePlanetTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdPlanetsOKBodyItemsTypePlanetTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCharactersCharacterIDPlanetsOKBodyItems) validatePlanetType(formats strfmt.Registry) error {

	if err := validate.Required("planet_type", "body", m.PlanetType); err != nil {
		return err
	}

	// value enum
	if err := m.validatePlanetTypeEnum("planet_type", "body", *m.PlanetType); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDPlanetsOKBodyItems) validateSolarSystemID(formats strfmt.Registry) error {

	if err := validate.Required("solar_system_id", "body", m.SolarSystemID); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDPlanetsOKBodyItems) validateUpgradeLevel(formats strfmt.Registry) error {

	if err := validate.Required("upgrade_level", "body", m.UpgradeLevel); err != nil {
		return err
	}

	if err := validate.MinimumInt("upgrade_level", "body", int64(*m.UpgradeLevel), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("upgrade_level", "body", int64(*m.UpgradeLevel), 5, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDPlanetsOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDPlanetsOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDPlanetsOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
