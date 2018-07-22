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

// GetCharactersCharacterIDBookmarksOKBodyItems get_characters_character_id_bookmarks_200_ok
//
// 200 ok object
// swagger:model getCharactersCharacterIdBookmarksOKBodyItems
type GetCharactersCharacterIDBookmarksOKBodyItems struct {

	// get_characters_character_id_bookmarks_bookmark_id
	//
	// bookmark_id integer
	// Required: true
	BookmarkID *int32 `json:"bookmark_id"`

	// coordinates
	Coordinates *GetCharactersCharacterIDBookmarksOKBodyItemsCoordinates `json:"coordinates,omitempty"`

	// get_characters_character_id_bookmarks_created
	//
	// created string
	// Required: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created"`

	// get_characters_character_id_bookmarks_creator_id
	//
	// creator_id integer
	// Required: true
	CreatorID *int32 `json:"creator_id"`

	// get_characters_character_id_bookmarks_folder_id
	//
	// folder_id integer
	FolderID int32 `json:"folder_id,omitempty"`

	// item
	Item *GetCharactersCharacterIDBookmarksOKBodyItemsItem `json:"item,omitempty"`

	// get_characters_character_id_bookmarks_label
	//
	// label string
	// Required: true
	Label *string `json:"label"`

	// get_characters_character_id_bookmarks_location_id
	//
	// location_id integer
	// Required: true
	LocationID *int32 `json:"location_id"`

	// get_characters_character_id_bookmarks_notes
	//
	// notes string
	// Required: true
	Notes *string `json:"notes"`
}

// Validate validates this get characters character Id bookmarks o k body items
func (m *GetCharactersCharacterIDBookmarksOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBookmarkID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCoordinates(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatorID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateItem(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCharactersCharacterIDBookmarksOKBodyItems) validateBookmarkID(formats strfmt.Registry) error {

	if err := validate.Required("bookmark_id", "body", m.BookmarkID); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDBookmarksOKBodyItems) validateCoordinates(formats strfmt.Registry) error {

	if swag.IsZero(m.Coordinates) { // not required
		return nil
	}

	if m.Coordinates != nil {
		if err := m.Coordinates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("coordinates")
			}
			return err
		}
	}

	return nil
}

func (m *GetCharactersCharacterIDBookmarksOKBodyItems) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("created", "body", m.Created); err != nil {
		return err
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDBookmarksOKBodyItems) validateCreatorID(formats strfmt.Registry) error {

	if err := validate.Required("creator_id", "body", m.CreatorID); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDBookmarksOKBodyItems) validateItem(formats strfmt.Registry) error {

	if swag.IsZero(m.Item) { // not required
		return nil
	}

	if m.Item != nil {
		if err := m.Item.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("item")
			}
			return err
		}
	}

	return nil
}

func (m *GetCharactersCharacterIDBookmarksOKBodyItems) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDBookmarksOKBodyItems) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("location_id", "body", m.LocationID); err != nil {
		return err
	}

	return nil
}

func (m *GetCharactersCharacterIDBookmarksOKBodyItems) validateNotes(formats strfmt.Registry) error {

	if err := validate.Required("notes", "body", m.Notes); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCharactersCharacterIDBookmarksOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCharactersCharacterIDBookmarksOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetCharactersCharacterIDBookmarksOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}