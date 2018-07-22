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

// GetCharactersCharacterIDMailMailIDOKBody get_characters_character_id_mail_mail_id_ok
//
// 200 ok object
// swagger:model getCharactersCharacterIdMailMailIdOKBody
type GetCharactersCharacterIDMailMailIDOKBody struct {

	// get_characters_character_id_mail_mail_id_body
	//
	// Mail's body
	Body string `json:"body,omitempty"`

	// get_characters_character_id_mail_mail_id_from
	//
	// From whom the mail was sent
	From int32 `json:"from,omitempty"`

	// get_characters_character_id_mail_mail_id_labels
	//
	// Labels attached to the mail
	// Max Items: 25
	Labels []*int32 `json:"labels"`

	// get_characters_character_id_mail_mail_id_read
	//
	// Whether the mail is flagged as read
	Read bool `json:"read,omitempty"`

	// get_characters_character_id_mail_mail_id_recipients
	//
	// Recipients of the mail
	// Max Items: 52
	// Min Items: 1
	// Unique: true
	Recipients []*GetCharactersCharacterIDMailMailIDOKBodyRecipientsItems `json:"recipients"`

	// get_characters_character_id_mail_mail_id_subject
	//
	// Mail subject
	Subject string `json:"subject,omitempty"`

	// get_characters_character_id_mail_mail_id_timestamp
	//
	// When the mail was sent
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this get characters character Id mail mail Id o k body
func (m *GetCharactersCharacterIDMailMailIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabels(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecipients(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDMailMailIDOKBody) validateLabels(formats strfmt.Registry) error {

	if swag.IsZero(m.Labels) { // not required
		return nil
	}

	iLabelsSize := int64(len(m.Labels))

	if err := validate.MaxItems("labels", "body", iLabelsSize, 25); err != nil {
		return err
	}

	for i := 0; i < len(m.Labels); i++ {
		if swag.IsZero(m.Labels[i]) { // not required
			continue
		}

		if err := validate.MinimumInt("labels"+"."+strconv.Itoa(i), "body", int64(*m.Labels[i]), 0, false); err != nil {
			return err
		}

	}

	return nil
}

func (m *GetCharactersCharacterIDMailMailIDOKBody) validateRecipients(formats strfmt.Registry) error {

	if swag.IsZero(m.Recipients) { // not required
		return nil
	}

	iRecipientsSize := int64(len(m.Recipients))

	if err := validate.MinItems("recipients", "body", iRecipientsSize, 1); err != nil {
		return err
	}

	if err := validate.MaxItems("recipients", "body", iRecipientsSize, 52); err != nil {
		return err
	}

	if err := validate.UniqueItems("recipients", "body", m.Recipients); err != nil {
		return err
	}

	for i := 0; i < len(m.Recipients); i++ {
		if swag.IsZero(m.Recipients[i]) { // not required
			continue
		}

		if m.Recipients[i] != nil {
			if err := m.Recipients[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("recipients" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GetCharactersCharacterIDMailMailIDOKBody) validateTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("timestamp", "body", "date-time", m.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDMailMailIDOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDMailMailIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDMailMailIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
