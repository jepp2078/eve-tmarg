// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetCharactersCharacterIDMedalsOKBodyItems get_characters_character_id_medals_200_ok
//
// 200 ok object
// swagger:model getCharactersCharacterIdMedalsOKBodyItems
type GetCharactersCharacterIDMedalsOKBodyItems struct {

	// get_characters_character_id_medals_corporation_id
	//
	// corporation_id integer
	// Required: true
	CorporationID *int32 `json:"corporation_id"`

	// get_characters_character_id_medals_date
	//
	// date string
	// Required: true
	// Format: date-time
	Date *strfmt.DateTime `json:"date"`

	// get_characters_character_id_medals_description
	//
	// description string
	// Required: true
	Description *string `json:"description"`

	// get_characters_character_id_medals_graphics
	//
	// graphics array
	// Required: true
	// Max Items: 9
	// Min Items: 3
	Graphics []*GetCharactersCharacterIDMedalsOKBodyItemsGraphicsItems `json:"graphics"`

	// get_characters_character_id_medals_issuer_id
	//
	// issuer_id integer
	// Required: true
	IssuerID *int32 `json:"issuer_id"`

	// get_characters_character_id_medals_medal_id
	//
	// medal_id integer
	// Required: true
	MedalID *int32 `json:"medal_id"`

	// get_characters_character_id_medals_reason
	//
	// reason string
	// Required: true
	Reason *string `json:"reason"`

	// get_characters_character_id_medals_status
	//
	// status string
	// Required: true
	// Enum: [public private]
	Status *string `json:"status"`

	// get_characters_character_id_medals_title
	//
	// title string
	// Required: true
	Title *string `json:"title"`
}

// Validate validates this get characters character Id medals o k body items
func (m *GetCharactersCharacterIDMedalsOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCorporationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGraphics(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssuerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMedalID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReason(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDMedalsOKBodyItems) validateCorporationID(formats strfmt.Registry) error {

	if err := validate.Required("corporation_id", "body", m.CorporationID); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDMedalsOKBodyItems) validateDate(formats strfmt.Registry) error {

	if err := validate.Required("date", "body", m.Date); err != nil {
		return err
	}

	if err := validate.FormatOf("date", "body", "date-time", m.Date.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDMedalsOKBodyItems) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDMedalsOKBodyItems) validateGraphics(formats strfmt.Registry) error {

	if err := validate.Required("graphics", "body", m.Graphics); err != nil {
		return err
	}

	iGraphicsSize := int64(len(m.Graphics))

	if err := validate.MinItems("graphics", "body", iGraphicsSize, 3); err != nil {
		return err
	}

	if err := validate.MaxItems("graphics", "body", iGraphicsSize, 9); err != nil {
		return err
	}

	for i := 0; i < len(m.Graphics); i++ {
		if swag.IsZero(m.Graphics[i]) { // not required
			continue
		}

		if m.Graphics[i] != nil {
			if err := m.Graphics[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("graphics" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetCharactersCharacterIDMedalsOKBodyItems) validateIssuerID(formats strfmt.Registry) error {

	if err := validate.Required("issuer_id", "body", m.IssuerID); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDMedalsOKBodyItems) validateMedalID(formats strfmt.Registry) error {

	if err := validate.Required("medal_id", "body", m.MedalID); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDMedalsOKBodyItems) validateReason(formats strfmt.Registry) error {

	if err := validate.Required("reason", "body", m.Reason); err != nil {
		return err
	}

	return nil
}

var getCharactersCharacterIdMedalsOKBodyItemsTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["public","private"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCharactersCharacterIdMedalsOKBodyItemsTypeStatusPropEnum = append(getCharactersCharacterIdMedalsOKBodyItemsTypeStatusPropEnum, v)
	}
}

const (

	// GetCharactersCharacterIDMedalsOKBodyItemsStatusPublic captures enum value "public"
	GetCharactersCharacterIDMedalsOKBodyItemsStatusPublic string = "public"

	// GetCharactersCharacterIDMedalsOKBodyItemsStatusPrivate captures enum value "private"
	GetCharactersCharacterIDMedalsOKBodyItemsStatusPrivate string = "private"
)

// prop value enum
func (m *GetCharactersCharacterIDMedalsOKBodyItems) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCharactersCharacterIdMedalsOKBodyItemsTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCharactersCharacterIDMedalsOKBodyItems) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDMedalsOKBodyItems) validateTitle(formats strfmt.Registry) error {

	if err := validate.Required("title", "body", m.Title); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDMedalsOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDMedalsOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDMedalsOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}