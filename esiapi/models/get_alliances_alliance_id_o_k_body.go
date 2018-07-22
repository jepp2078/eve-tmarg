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

// GetAlliancesAllianceIDOKBody get_alliances_alliance_id_ok
//
// 200 ok object
// swagger:model getAlliancesAllianceIdOKBody
type GetAlliancesAllianceIDOKBody struct {

	// get_alliances_alliance_id_creator_corporation_id
	//
	// ID of the corporation that created the alliance
	// Required: true
	CreatorCorporationID *int32 `json:"creator_corporation_id"`

	// get_alliances_alliance_id_creator_id
	//
	// ID of the character that created the alliance
	// Required: true
	CreatorID *int32 `json:"creator_id"`

	// get_alliances_alliance_id_date_founded
	//
	// date_founded string
	// Required: true
	// Format: date-time
	DateFounded *strfmt.DateTime `json:"date_founded"`

	// get_alliances_alliance_id_executor_corporation_id
	//
	// the executor corporation ID, if this alliance is not closed
	ExecutorCorporationID int32 `json:"executor_corporation_id,omitempty"`

	// get_alliances_alliance_id_faction_id
	//
	// Faction ID this alliance is fighting for, if this alliance is enlisted in factional warfare
	FactionID int32 `json:"faction_id,omitempty"`

	// get_alliances_alliance_id_name
	//
	// the full name of the alliance
	// Required: true
	Name *string `json:"name"`

	// get_alliances_alliance_id_ticker
	//
	// the short name of the alliance
	// Required: true
	Ticker *string `json:"ticker"`
}

// Validate validates this get alliances alliance Id o k body
func (m *GetAlliancesAllianceIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreatorCorporationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateFounded(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTicker(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetAlliancesAllianceIDOKBody) validateCreatorCorporationID(formats strfmt.Registry) error {

	if err := validate.Required("creator_corporation_id", "body", m.CreatorCorporationID); err != nil {
		return err
	}

	return nil
}

func (m *GetAlliancesAllianceIDOKBody) validateCreatorID(formats strfmt.Registry) error {

	if err := validate.Required("creator_id", "body", m.CreatorID); err != nil {
		return err
	}

	return nil
}

func (m *GetAlliancesAllianceIDOKBody) validateDateFounded(formats strfmt.Registry) error {

	if err := validate.Required("date_founded", "body", m.DateFounded); err != nil {
		return err
	}

	if err := validate.FormatOf("date_founded", "body", "date-time", m.DateFounded.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetAlliancesAllianceIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *GetAlliancesAllianceIDOKBody) validateTicker(formats strfmt.Registry) error {

	if err := validate.Required("ticker", "body", m.Ticker); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetAlliancesAllianceIDOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetAlliancesAllianceIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetAlliancesAllianceIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}