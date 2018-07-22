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

// GetUniverseCategoriesCategoryIDOKBody get_universe_categories_category_id_ok
//
// 200 ok object
// swagger:model getUniverseCategoriesCategoryIdOKBody
type GetUniverseCategoriesCategoryIDOKBody struct {

	// get_universe_categories_category_id_category_id
	//
	// category_id integer
	// Required: true
	CategoryID *int32 `json:"category_id"`

	// get_universe_categories_category_id_groups
	//
	// groups array
	// Required: true
	// Max Items: 10000
	Groups []int32 `json:"groups"`

	// get_universe_categories_category_id_name
	//
	// name string
	// Required: true
	Name *string `json:"name"`

	// get_universe_categories_category_id_published
	//
	// published boolean
	// Required: true
	Published *bool `json:"published"`
}

// Validate validates this get universe categories category Id o k body
func (m *GetUniverseCategoriesCategoryIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCategoryID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePublished(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetUniverseCategoriesCategoryIDOKBody) validateCategoryID(formats strfmt.Registry) error {

	if err := validate.Required("category_id", "body", m.CategoryID); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseCategoriesCategoryIDOKBody) validateGroups(formats strfmt.Registry) error {

	if err := validate.Required("groups", "body", m.Groups); err != nil {
		return err
	}

	iGroupsSize := int64(len(m.Groups))

	if err := validate.MaxItems("groups", "body", iGroupsSize, 10000); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseCategoriesCategoryIDOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *GetUniverseCategoriesCategoryIDOKBody) validatePublished(formats strfmt.Registry) error {

	if err := validate.Required("published", "body", m.Published); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetUniverseCategoriesCategoryIDOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUniverseCategoriesCategoryIDOKBody) UnmarshalBinary(b []byte) error {
	var res GetUniverseCategoriesCategoryIDOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}