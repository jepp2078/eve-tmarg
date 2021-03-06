// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetFwLeaderboardsCharactersOKBodyKillsActiveTotalItems get_fw_leaderboards_characters_active_total_active_total
//
// active_total object
// swagger:model getFwLeaderboardsCharactersOKBodyKillsActiveTotalItems
type GetFwLeaderboardsCharactersOKBodyKillsActiveTotalItems struct {

	// get_fw_leaderboards_characters_amount
	//
	// Amount of kills
	Amount int32 `json:"amount,omitempty"`

	// get_fw_leaderboards_characters_character_id
	//
	// character_id integer
	CharacterID int32 `json:"character_id,omitempty"`
}

// Validate validates this get fw leaderboards characters o k body kills active total items
func (m *GetFwLeaderboardsCharactersOKBodyKillsActiveTotalItems) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetFwLeaderboardsCharactersOKBodyKillsActiveTotalItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetFwLeaderboardsCharactersOKBodyKillsActiveTotalItems) UnmarshalBinary(b []byte) error {
	var res GetFwLeaderboardsCharactersOKBodyKillsActiveTotalItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
