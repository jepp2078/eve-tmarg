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

// GetCorporationsCorporationIDContractsContractIDBidsOKBodyItems get_corporations_corporation_id_contracts_contract_id_bids_200_ok
//
// 200 ok object
// swagger:model getCorporationsCorporationIdContractsContractIdBidsOKBodyItems
type GetCorporationsCorporationIDContractsContractIDBidsOKBodyItems struct {

	// get_corporations_corporation_id_contracts_contract_id_bids_amount
	//
	// The amount bid, in ISK
	// Required: true
	Amount *float32 `json:"amount"`

	// get_corporations_corporation_id_contracts_contract_id_bids_bid_id
	//
	// Unique ID for the bid
	// Required: true
	BidID *int32 `json:"bid_id"`

	// get_corporations_corporation_id_contracts_contract_id_bids_bidder_id
	//
	// Character ID of the bidder
	// Required: true
	BidderID *int32 `json:"bidder_id"`

	// get_corporations_corporation_id_contracts_contract_id_bids_date_bid
	//
	// Datetime when the bid was placed
	// Required: true
	// Format: date-time
	DateBid *strfmt.DateTime `json:"date_bid"`
}

// Validate validates this get corporations corporation Id contracts contract Id bids o k body items
func (m *GetCorporationsCorporationIDContractsContractIDBidsOKBodyItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBidID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBidderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDateBid(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetCorporationsCorporationIDContractsContractIDBidsOKBodyItems) validateAmount(formats strfmt.Registry) error {

	if err := validate.Required("amount", "body", m.Amount); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDContractsContractIDBidsOKBodyItems) validateBidID(formats strfmt.Registry) error {

	if err := validate.Required("bid_id", "body", m.BidID); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDContractsContractIDBidsOKBodyItems) validateBidderID(formats strfmt.Registry) error {

	if err := validate.Required("bidder_id", "body", m.BidderID); err != nil {
		return err
	}

	return nil
}

func (m *GetCorporationsCorporationIDContractsContractIDBidsOKBodyItems) validateDateBid(formats strfmt.Registry) error {

	if err := validate.Required("date_bid", "body", m.DateBid); err != nil {
		return err
	}

	if err := validate.FormatOf("date_bid", "body", "date-time", m.DateBid.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetCorporationsCorporationIDContractsContractIDBidsOKBodyItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetCorporationsCorporationIDContractsContractIDBidsOKBodyItems) UnmarshalBinary(b []byte) error {
	var res GetCorporationsCorporationIDContractsContractIDBidsOKBodyItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
