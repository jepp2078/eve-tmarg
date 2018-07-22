// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetCharactersCharacterIDStatsOKBodyItemsTravel get_characters_character_id_stats_travel
//
// travel object
// swagger:model getCharactersCharacterIdStatsOKBodyItemsTravel
type GetCharactersCharacterIDStatsOKBodyItemsTravel struct {

	// get_characters_character_id_stats_acceleration_gate_activations
	//
	// acceleration_gate_activations integer
	AccelerationGateActivations int64 `json:"acceleration_gate_activations,omitempty"`

	// get_characters_character_id_stats_align_to
	//
	// align_to integer
	AlignTo int64 `json:"align_to,omitempty"`

	// get_characters_character_id_stats_distance_warped_high_sec
	//
	// distance_warped_high_sec integer
	DistanceWarpedHighSec int64 `json:"distance_warped_high_sec,omitempty"`

	// get_characters_character_id_stats_distance_warped_low_sec
	//
	// distance_warped_low_sec integer
	DistanceWarpedLowSec int64 `json:"distance_warped_low_sec,omitempty"`

	// get_characters_character_id_stats_distance_warped_null_sec
	//
	// distance_warped_null_sec integer
	DistanceWarpedNullSec int64 `json:"distance_warped_null_sec,omitempty"`

	// get_characters_character_id_stats_distance_warped_wormhole
	//
	// distance_warped_wormhole integer
	DistanceWarpedWormhole int64 `json:"distance_warped_wormhole,omitempty"`

	// get_characters_character_id_stats_docks_high_sec
	//
	// docks_high_sec integer
	DocksHighSec int64 `json:"docks_high_sec,omitempty"`

	// get_characters_character_id_stats_docks_low_sec
	//
	// docks_low_sec integer
	DocksLowSec int64 `json:"docks_low_sec,omitempty"`

	// get_characters_character_id_stats_docks_null_sec
	//
	// docks_null_sec integer
	DocksNullSec int64 `json:"docks_null_sec,omitempty"`

	// get_characters_character_id_stats_jumps_stargate_high_sec
	//
	// jumps_stargate_high_sec integer
	JumpsStargateHighSec int64 `json:"jumps_stargate_high_sec,omitempty"`

	// get_characters_character_id_stats_jumps_stargate_low_sec
	//
	// jumps_stargate_low_sec integer
	JumpsStargateLowSec int64 `json:"jumps_stargate_low_sec,omitempty"`

	// get_characters_character_id_stats_jumps_stargate_null_sec
	//
	// jumps_stargate_null_sec integer
	JumpsStargateNullSec int64 `json:"jumps_stargate_null_sec,omitempty"`

	// get_characters_character_id_stats_jumps_wormhole
	//
	// jumps_wormhole integer
	JumpsWormhole int64 `json:"jumps_wormhole,omitempty"`

	// get_characters_character_id_stats_warps_high_sec
	//
	// warps_high_sec integer
	WarpsHighSec int64 `json:"warps_high_sec,omitempty"`

	// get_characters_character_id_stats_warps_low_sec
	//
	// warps_low_sec integer
	WarpsLowSec int64 `json:"warps_low_sec,omitempty"`

	// get_characters_character_id_stats_warps_null_sec
	//
	// warps_null_sec integer
	WarpsNullSec int64 `json:"warps_null_sec,omitempty"`

	// get_characters_character_id_stats_warps_to_bookmark
	//
	// warps_to_bookmark integer
	WarpsToBookmark int64 `json:"warps_to_bookmark,omitempty"`

	// get_characters_character_id_stats_warps_to_celestial
	//
	// warps_to_celestial integer
	WarpsToCelestial int64 `json:"warps_to_celestial,omitempty"`

	// get_characters_character_id_stats_warps_to_fleet_member
	//
	// warps_to_fleet_member integer
	WarpsToFleetMember int64 `json:"warps_to_fleet_member,omitempty"`

	// get_characters_character_id_stats_warps_to_scan_result
	//
	// warps_to_scan_result integer
	WarpsToScanResult int64 `json:"warps_to_scan_result,omitempty"`

	// get_characters_character_id_stats_warps_wormhole
	//
	// warps_wormhole integer
	WarpsWormhole int64 `json:"warps_wormhole,omitempty"`
}

// Validate validates this get characters character Id stats o k body items travel
func (m *GetCharactersCharacterIDStatsOKBodyItemsTravel) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDStatsOKBodyItemsTravel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDStatsOKBodyItemsTravel) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDStatsOKBodyItemsTravel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}