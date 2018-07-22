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

// GetIndustryFacilitiesOKBodyItems get_industry_facilities_200_ok
//
// 200 ok object
// swagger:model getIndustryFacilitiesOKBodyItems
type GetIndustryFacilitiesOKBodyItems struct {

	// get_industry_facilities_facility_id
	//
	// ID of the facility
	// Required: true
	FacilityID *int64 `json:"facility_id"`

	// get_industry_facilities_owner_id
	//
	// Owner of the facility
	// Required: true
	OwnerID *int32 `json:"owner_id"`

	// get_industry_facilities_region_id
	//
	// Region ID where the facility is
	// Required: true
	RegionID *int32 `json:"region_id"`

	// get_industry_facilities_solar_system_id
	//
	// Solar system ID where the facility is
	// Required: true
	SolarSystemID *int32 `json:"solar_system_id"`

	// get_industry_facilities_tax
	//
	// Tax imposed by the facility
	Tax float32 `json:"tax,omitempty"`

	// get_industry_facilities_type_id
	//
	// Type ID of the facility
	// Required: true
	TypeID *int32 `json:"type_id"`
}

// Validate validates this get industry facilities o k body items
func (m *GetIndustryFacilitiesOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFacilityID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwnerID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSolarSystemID(formats); err != nil {
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

func (m *GetIndustryFacilitiesOKBodyItems) validateFacilityID(formats strfmt.Registry) error {

	if err := validate.Required("facility_id", "body", m.FacilityID); err != nil {
		return err
	}

	return nil
}

func (m *GetIndustryFacilitiesOKBodyItems) validateOwnerID(formats strfmt.Registry) error {

	if err := validate.Required("owner_id", "body", m.OwnerID); err != nil {
		return err
	}

	return nil
}

func (m *GetIndustryFacilitiesOKBodyItems) validateRegionID(formats strfmt.Registry) error {

	if err := validate.Required("region_id", "body", m.RegionID); err != nil {
		return err
	}

	return nil
}

func (m *GetIndustryFacilitiesOKBodyItems) validateSolarSystemID(formats strfmt.Registry) error {

	if err := validate.Required("solar_system_id", "body", m.SolarSystemID); err != nil {
		return err
	}

	return nil
}

func (m *GetIndustryFacilitiesOKBodyItems) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", m.TypeID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetIndustryFacilitiesOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetIndustryFacilitiesOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetIndustryFacilitiesOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}