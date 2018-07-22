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

// GetSovereigntyCampaignsOKBodyItemsParticipantsItems get_sovereignty_campaigns_participant
//
// participant object
// swagger:model getSovereigntyCampaignsOKBodyItemsParticipantsItems
type GetSovereigntyCampaignsOKBodyItemsParticipantsItems struct {

	// get_sovereignty_campaigns_alliance_id
	//
	// alliance_id integer
	// Required: true
	AllianceID *int32 `json:"alliance_id"`

	// get_sovereignty_campaigns_score
	//
	// score number
	// Required: true
	Score *float32 `json:"score"`
}

// Validate validates this get sovereignty campaigns o k body items participants items
func (m *GetSovereigntyCampaignsOKBodyItemsParticipantsItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAllianceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScore(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSovereigntyCampaignsOKBodyItemsParticipantsItems) validateAllianceID(formats strfmt.Registry) error {

	if err := validate.Required("alliance_id", "body", m.AllianceID); err != nil {
		return err
	}

	return nil
}

func (m *GetSovereigntyCampaignsOKBodyItemsParticipantsItems) validateScore(formats strfmt.Registry) error {

	if err := validate.Required("score", "body", m.Score); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetSovereigntyCampaignsOKBodyItemsParticipantsItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSovereigntyCampaignsOKBodyItemsParticipantsItems) UnmarshalBinary(b []byte) error {
	var res GetSovereigntyCampaignsOKBodyItemsParticipantsItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}