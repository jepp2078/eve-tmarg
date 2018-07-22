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

// GetCharactersCharacterIDFwStatsOKBodyKills get_characters_character_id_fw_stats_kills
//
// Summary of kills done by the given character against enemy factions
// swagger:model getCharactersCharacterIdFwStatsOKBodyKills
type GetCharactersCharacterIDFwStatsOKBodyKills struct {

	// get_characters_character_id_fw_stats_last_week
	//
	// Last week's total number of kills by a given character against enemy factions
	// Required: true
	LastWeek *int32 `json:"last_week"`

	// get_characters_character_id_fw_stats_total
	//
	// Total number of kills by a given character against enemy factions since the character enlisted
	// Required: true
	Total *int32 `json:"total"`

	// get_characters_character_id_fw_stats_yesterday
	//
	// Yesterday's total number of kills by a given character against enemy factions
	// Required: true
	Yesterday *int32 `json:"yesterday"`
}

// Validate validates this get characters character Id fw stats o k body kills
func (m *GetCharactersCharacterIDFwStatsOKBodyKills) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLastWeek(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateYesterday(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDFwStatsOKBodyKills) validateLastWeek(formats strfmt.Registry) error {

	if err := validate.Required("last_week", "body", m.LastWeek); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDFwStatsOKBodyKills) validateTotal(formats strfmt.Registry) error {

	if err := validate.Required("total", "body", m.Total); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDFwStatsOKBodyKills) validateYesterday(formats strfmt.Registry) error {

	if err := validate.Required("yesterday", "body", m.Yesterday); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDFwStatsOKBodyKills) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDFwStatsOKBodyKills) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDFwStatsOKBodyKills
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}