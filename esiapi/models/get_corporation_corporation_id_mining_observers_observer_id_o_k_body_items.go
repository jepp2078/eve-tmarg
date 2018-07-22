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

// GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems get_corporation_corporation_id_mining_observers_observer_id_200_ok
//
// 200 ok object
// swagger:model getCorporationCorporationIdMiningObserversObserverIdOKBodyItems
type GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems struct {

	// get_corporation_corporation_id_mining_observers_observer_id_character_id
	//
	// The character that did the mining
	//
	// Required: true
	CharacterID *int32 `json:"character_id"`

	// get_corporation_corporation_id_mining_observers_observer_id_last_updated
	//
	// last_updated string
	// Required: true
	// Format: date
	LastUpdated *strfmt.Date `json:"last_updated"`

	// get_corporation_corporation_id_mining_observers_observer_id_quantity
	//
	// quantity integer
	// Required: true
	Quantity *int64 `json:"quantity"`

	// get_corporation_corporation_id_mining_observers_observer_id_recorded_corporation_id
	//
	// The corporation id of the character at the time data was recorded.
	//
	// Required: true
	RecordedCorporationID *int32 `json:"recorded_corporation_id"`

	// get_corporation_corporation_id_mining_observers_observer_id_type_id
	//
	// type_id integer
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get corporation corporation Id mining observers observer Id o k body items
func (m *GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCharacterID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecordedCorporationID(formats); err != nil {
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

func (m *GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems) validateCharacterID(formats strfmt.Registry) error {

	if err := validate.Required("character_id", "body", m.CharacterID); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems) validateLastUpdated(formats strfmt.Registry) error {

	if err := validate.Required("last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	if err := validate.FormatOf("last_updated", "body", "date", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems) validateQuantity(formats strfmt.Registry) error {

	if err := validate.Required("quantity", "body", m.Quantity); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems) validateRecordedCorporationID(formats strfmt.Registry) error {

	if err := validate.Required("recorded_corporation_id", "body", m.RecordedCorporationID); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", m.TypeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetCorporationCorporationIDMiningObserversObserverIDOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}