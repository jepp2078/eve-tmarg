// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetKillmailsKillmailIDKillmailHashOKBodyVictim get_killmails_killmail_id_killmail_hash_victim
//
// victim object
// swagger:model getKillmailsKillmailIdKillmailHashOKBodyVictim
type GetKillmailsKillmailIDKillmailHashOKBodyVictim struct {

	// get_killmails_killmail_id_killmail_hash_victim_alliance_id
	//
	// alliance_id integer
	AllianceID int32 `json:"alliance_id,omitempty"`

	// get_killmails_killmail_id_killmail_hash_victim_character_id
	//
	// character_id integer
	CharacterID int32 `json:"character_id,omitempty"`

	// get_killmails_killmail_id_killmail_hash_victim_corporation_id
	//
	// corporation_id integer
	CorporationID int32 `json:"corporation_id,omitempty"`

	// get_killmails_killmail_id_killmail_hash_damage_taken
	//
	// How much total damage was taken by the victim
	//
	// Required: true
	DamageTaken *int32 `json:"damage_taken"`

	// get_killmails_killmail_id_killmail_hash_victim_faction_id
	//
	// faction_id integer
	FactionID int32 `json:"faction_id,omitempty"`

	// get_killmails_killmail_id_killmail_hash_items
	//
	// items array
	// Max Items: 10000
	Items []*GetKillmailsKillmailIDKillmailHashOKBodyVictimItemsItems `json:"items"`

	// position
	Position *GetKillmailsKillmailIDKillmailHashOKBodyVictimPosition `json:"position,omitempty"`

	// get_killmails_killmail_id_killmail_hash_victim_ship_type_id
	//
	// The ship that the victim was piloting and was destroyed
	//
	// Required: true
	ShipTypeID *int32 `json:"ship_type_id"`
}

// Validate validates this get killmails killmail Id killmail hash o k body victim
func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictim) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDamageTaken(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShipTypeID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictim) validateDamageTaken(formats strfmt.Registry) error {

	if err := validate.Required("damage_taken", "body", m.DamageTaken); err != nil {
		return err
	}

	return nil
}

func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictim) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.Items) { // not required
		return nil
	}

	iItemsSize := int64(len(m.Items))

	if err := validate.MaxItems("items", "body", iItemsSize, 10000); err != nil {
		return err
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictim) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if m.Position != nil {
		if err := m.Position.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("position")
			}
			return err
		}
	}

	return nil
}

func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictim) validateShipTypeID(formats strfmt.Registry) error {

	if err := validate.Required("ship_type_id", "body", m.ShipTypeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictim) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetKillmailsKillmailIDKillmailHashOKBodyVictim) UnmarshalBinary(b []byte) error {
	var res GetKillmailsKillmailIDKillmailHashOKBodyVictim
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}