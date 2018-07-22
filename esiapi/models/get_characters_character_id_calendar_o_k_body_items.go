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

// GetCharactersCharacterIDCalendarOKBodyItems get_characters_character_id_calendar_200_ok
//
// event
// swagger:model getCharactersCharacterIdCalendarOKBodyItems
type GetCharactersCharacterIDCalendarOKBodyItems struct {

	// get_characters_character_id_calendar_event_date
	//
	// event_date string
	// Format: date-time
	EventDate strfmt.DateTime `json:"event_date,omitempty"`

	// get_characters_character_id_calendar_event_id
	//
	// event_id integer
	EventID int32 `json:"event_id,omitempty"`

	// get_characters_character_id_calendar_event_response
	//
	// event_response string
	// Enum: [declined not_responded accepted tentative]
	EventResponse string `json:"event_response,omitempty"`

	// get_characters_character_id_calendar_importance
	//
	// importance integer
	Importance int32 `json:"importance,omitempty"`

	// get_characters_character_id_calendar_title
	//
	// title string
	Title string `json:"title,omitempty"`
}

// Validate validates this get characters character Id calendar o k body items
func (m *GetCharactersCharacterIDCalendarOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEventDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEventResponse(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDCalendarOKBodyItems) validateEventDate(formats strfmt.Registry) error {

	if swag.IsZero(m.EventDate) { // not required
		return nil
	}

	if err := validate.FormatOf("event_date", "body", "date-time", m.EventDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdCalendarOKBodyItemsTypeEventResponsePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["declined","not_responded","accepted","tentative"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdCalendarOKBodyItemsTypeEventResponsePropEnum = append(getCharactersCharacterIdCalendarOKBodyItemsTypeEventResponsePropEnum, v)
	}
}

const (

	// GetCharactersCharacterIDCalendarOKBodyItemsEventResponseDeclined captures enum value "declined"
	GetCharactersCharacterIDCalendarOKBodyItemsEventResponseDeclined string = "declined"

	// GetCharactersCharacterIDCalendarOKBodyItemsEventResponseNotResponded captures enum value "not_responded"
	GetCharactersCharacterIDCalendarOKBodyItemsEventResponseNotResponded string = "not_responded"

	// GetCharactersCharacterIDCalendarOKBodyItemsEventResponseAccepted captures enum value "accepted"
	GetCharactersCharacterIDCalendarOKBodyItemsEventResponseAccepted string = "accepted"

	// GetCharactersCharacterIDCalendarOKBodyItemsEventResponseTentative captures enum value "tentative"
	GetCharactersCharacterIDCalendarOKBodyItemsEventResponseTentative string = "tentative"
)

// prop value enum
func (m *GetCharactersCharacterIDCalendarOKBodyItems) validateEventResponseEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdCalendarOKBodyItemsTypeEventResponsePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCharactersCharacterIDCalendarOKBodyItems) validateEventResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.EventResponse) { // not required
		return nil
	}

	// value enum
	if err := m.validateEventResponseEnum("event_response", "body", m.EventResponse); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDCalendarOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDCalendarOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDCalendarOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
