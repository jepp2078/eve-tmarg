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

// GetCorporationsCorporationIDShareholdersOKBodyItems get_corporations_corporation_id_shareholders_200_ok
//
// 200 ok object
// swagger:model getCorporationsCorporationIdShareholdersOKBodyItems
type GetCorporationsCorporationIDShareholdersOKBodyItems struct {

	// get_corporations_corporation_id_shareholders_share_count
	//
	// share_count integer
	// Required: true
	ShareCount *int64 `json:"share_count"`

	// get_corporations_corporation_id_shareholders_shareholder_id
	//
	// shareholder_id integer
	// Required: true
	ShareholderID *int32 `json:"shareholder_id"`

	// get_corporations_corporation_id_shareholders_shareholder_type
	//
	// shareholder_type string
	// Required: true
	// Enum: [character corporation]
	ShareholderType *string `json:"shareholder_type"`
}

// Validate validates this get corporations corporation Id shareholders o k body items
func (m *GetCorporationsCorporationIDShareholdersOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateShareCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShareholderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShareholderType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCorporationsCorporationIDShareholdersOKBodyItems) validateShareCount(formats strfmt.Registry) error {

	if err := validate.Required("share_count", "body", m.ShareCount); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDShareholdersOKBodyItems) validateShareholderID(formats strfmt.Registry) error {

	if err := validate.Required("shareholder_id", "body", m.ShareholderID); err != nil {
		return err
	}

	return nil
}

var getCorporationsCorporationIdShareholdersOKBodyItemsTypeShareholderTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["character","corporation"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdShareholdersOKBodyItemsTypeShareholderTypePropEnum = append(getCorporationsCorporationIdShareholdersOKBodyItemsTypeShareholderTypePropEnum, v)
	}
}

const (

	// GetCorporationsCorporationIDShareholdersOKBodyItemsShareholderTypeCharacter captures enum value "character"
	GetCorporationsCorporationIDShareholdersOKBodyItemsShareholderTypeCharacter string = "character"

	// GetCorporationsCorporationIDShareholdersOKBodyItemsShareholderTypeCorporation captures enum value "corporation"
	GetCorporationsCorporationIDShareholdersOKBodyItemsShareholderTypeCorporation string = "corporation"
)

// prop value enum
func (m *GetCorporationsCorporationIDShareholdersOKBodyItems) validateShareholderTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCorporationsCorporationIdShareholdersOKBodyItemsTypeShareholderTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCorporationsCorporationIDShareholdersOKBodyItems) validateShareholderType(formats strfmt.Registry) error {

	if err := validate.Required("shareholder_type", "body", m.ShareholderType); err != nil {
		return err
	}

	// value enum
	if err := m.validateShareholderTypeEnum("shareholder_type", "body", *m.ShareholderType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCorporationsCorporationIDShareholdersOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCorporationsCorporationIDShareholdersOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDShareholdersOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}