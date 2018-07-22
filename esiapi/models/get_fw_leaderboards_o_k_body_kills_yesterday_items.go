// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetFwLeaderboardsOKBodyKillsYesterdayItems get_fw_leaderboards_yesterday_yesterday
//
// yesterday object
// swagger:model getFwLeaderboardsOKBodyKillsYesterdayItems
type GetFwLeaderboardsOKBodyKillsYesterdayItems struct {

	// get_fw_leaderboards_yesterday_amount
	//
	// Amount of kills
	Amount int32 `json:"amount,omitempty"`

	// get_fw_leaderboards_yesterday_faction_id
	//
	// faction_id integer
	FactionID int32 `json:"faction_id,omitempty"`
}

// Validate validates this get fw leaderboards o k body kills yesterday items
func (m *GetFwLeaderboardsOKBodyKillsYesterdayItems) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetFwLeaderboardsOKBodyKillsYesterdayItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetFwLeaderboardsOKBodyKillsYesterdayItems) UnmarshalBinary(b []byte) error {
	var res GetFwLeaderboardsOKBodyKillsYesterdayItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
