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

// GetCorporationsCorporationIDOrdersOKBodyItems get_corporations_corporation_id_orders_200_ok
//
// 200 ok object
// swagger:model getCorporationsCorporationIdOrdersOKBodyItems
type GetCorporationsCorporationIDOrdersOKBodyItems struct {

	// get_corporations_corporation_id_orders_duration
	//
	// Number of days for which order is valid (starting from the issued date). An order expires at time issued + duration
	// Required: true
	Duration *int32 `json:"duration"`

	// get_corporations_corporation_id_orders_escrow
	//
	// For buy orders, the amount of ISK in escrow
	Escrow float64 `json:"escrow,omitempty"`

	// get_corporations_corporation_id_orders_is_buy_order
	//
	// True if the order is a bid (buy) order
	IsBuyOrder bool `json:"is_buy_order,omitempty"`

	// get_corporations_corporation_id_orders_issued
	//
	// Date and time when this order was issued
	// Required: true
	// Format: date-time
	Issued *strfmt.DateTime `json:"issued"`

	// get_corporations_corporation_id_orders_location_id
	//
	// ID of the location where order was placed
	// Required: true
	LocationID *int64 `json:"location_id"`

	// get_corporations_corporation_id_orders_min_volume
	//
	// For buy orders, the minimum quantity that will be accepted in a matching sell order
	MinVolume int32 `json:"min_volume,omitempty"`

	// get_corporations_corporation_id_orders_order_id
	//
	// Unique order ID
	// Required: true
	OrderID *int64 `json:"order_id"`

	// get_corporations_corporation_id_orders_price
	//
	// Cost per unit for this order
	// Required: true
	Price *float64 `json:"price"`

	// get_corporations_corporation_id_orders_range
	//
	// Valid order range, numbers are ranges in jumps
	// Required: true
	// Enum: [1 10 2 20 3 30 4 40 5 region solarsystem station]
	Range *string `json:"range"`

	// get_corporations_corporation_id_orders_region_id
	//
	// ID of the region where order was placed
	// Required: true
	RegionID *int32 `json:"region_id"`

	// get_corporations_corporation_id_orders_type_id
	//
	// The type ID of the item transacted in this order
	// Required: true
	TypeID *int32 `json:"type_id"`

	// get_corporations_corporation_id_orders_volume_remain
	//
	// Quantity of items still required or offered
	// Required: true
	VolumeRemain *int32 `json:"volume_remain"`

	// get_corporations_corporation_id_orders_volume_total
	//
	// Quantity of items required or offered at time order was placed
	// Required: true
	VolumeTotal *int32 `json:"volume_total"`

	// get_corporations_corporation_id_orders_wallet_division
	//
	// The corporation wallet division used for this order.
	// Required: true
	// Maximum: 7
	// Minimum: 1
	WalletDivision *int32 `json:"wallet_division"`
}

// Validate validates this get corporations corporation Id orders o k body items
func (m *GetCorporationsCorporationIDOrdersOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssued(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTypeID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeRemain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVolumeTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWalletDivision(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCorporationsCorporationIDOrdersOKBodyItems) validateDuration(formats strfmt.Registry) error {

	if err := validate.Required("duration", "body", m.Duration); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOrdersOKBodyItems) validateIssued(formats strfmt.Registry) error {

	if err := validate.Required("issued", "body", m.Issued); err != nil {
		return err
	}

	if err := validate.FormatOf("issued", "body", "date-time", m.Issued.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOrdersOKBodyItems) validateLocationID(formats strfmt.Registry) error {

	if err := validate.Required("location_id", "body", m.LocationID); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOrdersOKBodyItems) validateOrderID(formats strfmt.Registry) error {

	if err := validate.Required("order_id", "body", m.OrderID); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOrdersOKBodyItems) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", m.Price); err != nil {
		return err
	}

	return nil
}

var getCorporationsCorporationIdOrdersOKBodyItemsTypeRangePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["1","10","2","20","3","30","4","40","5","region","solarsystem","station"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getCorporationsCorporationIdOrdersOKBodyItemsTypeRangePropEnum = append(getCorporationsCorporationIdOrdersOKBodyItemsTypeRangePropEnum, v)
	}
}

const (

	// GetCorporationsCorporationIDOrdersOKBodyItemsRangeNr1 captures enum value "1"
	GetCorporationsCorporationIDOrdersOKBodyItemsRangeNr1 string = "1"

	// GetCorporationsCorporationIDOrdersOKBodyItemsRangeNr10 captures enum value "10"
	GetCorporationsCorporationIDOrdersOKBodyItemsRangeNr10 string = "10"

	// GetCorporationsCorporationIDOrdersOKBodyItemsRangeNr2 captures enum value "2"
	GetCorporationsCorporationIDOrdersOKBodyItemsRangeNr2 string = "2"

	// GetCorporationsCorporationIDOrdersOKBodyItemsRangeNr20 captures enum value "20"
	GetCorporationsCorporationIDOrdersOKBodyItemsRangeNr20 string = "20"

	// GetCorporationsCorporationIDOrdersOKBodyItemsRangeNr3 captures enum value "3"
	GetCorporationsCorporationIDOrdersOKBodyItemsRangeNr3 string = "3"

	// GetCorporationsCorporationIDOrdersOKBodyItemsRangeNr30 captures enum value "30"
	GetCorporationsCorporationIDOrdersOKBodyItemsRangeNr30 string = "30"

	// GetCorporationsCorporationIDOrdersOKBodyItemsRangeNr4 captures enum value "4"
	GetCorporationsCorporationIDOrdersOKBodyItemsRangeNr4 string = "4"

	// GetCorporationsCorporationIDOrdersOKBodyItemsRangeNr40 captures enum value "40"
	GetCorporationsCorporationIDOrdersOKBodyItemsRangeNr40 string = "40"

	// GetCorporationsCorporationIDOrdersOKBodyItemsRangeNr5 captures enum value "5"
	GetCorporationsCorporationIDOrdersOKBodyItemsRangeNr5 string = "5"

	// GetCorporationsCorporationIDOrdersOKBodyItemsRangeRegion captures enum value "region"
	GetCorporationsCorporationIDOrdersOKBodyItemsRangeRegion string = "region"

	// GetCorporationsCorporationIDOrdersOKBodyItemsRangeSolarsystem captures enum value "solarsystem"
	GetCorporationsCorporationIDOrdersOKBodyItemsRangeSolarsystem string = "solarsystem"

	// GetCorporationsCorporationIDOrdersOKBodyItemsRangeStation captures enum value "station"
	GetCorporationsCorporationIDOrdersOKBodyItemsRangeStation string = "station"
)

// prop value enum
func (m *GetCorporationsCorporationIDOrdersOKBodyItems) validateRangeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, getCorporationsCorporationIdOrdersOKBodyItemsTypeRangePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *GetCorporationsCorporationIDOrdersOKBodyItems) validateRange(formats strfmt.Registry) error {

	if err := validate.Required("range", "body", m.Range); err != nil {
		return err
	}

	// value enum
	if err := m.validateRangeEnum("range", "body", *m.Range); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOrdersOKBodyItems) validateRegionID(formats strfmt.Registry) error {

	if err := validate.Required("region_id", "body", m.RegionID); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOrdersOKBodyItems) validateTypeID(formats strfmt.Registry) error {

	if err := validate.Required("type_id", "body", m.TypeID); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOrdersOKBodyItems) validateVolumeRemain(formats strfmt.Registry) error {

	if err := validate.Required("volume_remain", "body", m.VolumeRemain); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOrdersOKBodyItems) validateVolumeTotal(formats strfmt.Registry) error {

	if err := validate.Required("volume_total", "body", m.VolumeTotal); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDOrdersOKBodyItems) validateWalletDivision(formats strfmt.Registry) error {

	if err := validate.Required("wallet_division", "body", m.WalletDivision); err != nil {
		return err
	}

	if err := validate.MinimumInt("wallet_division", "body", int64(*m.WalletDivision), 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("wallet_division", "body", int64(*m.WalletDivision), 7, false); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCorporationsCorporationIDOrdersOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCorporationsCorporationIDOrdersOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDOrdersOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}