// Code generated by go-swagger; DO NOT EDIT.

package contracts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new contracts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for contracts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetCharactersCharacterIDContracts gets contracts

Returns contracts available to a character, only if the character is issuer, acceptor or assignee. Only returns contracts no older than 30 days, or if the status is "in_progress".

---
Alternate route: `/dev/characters/{character_id}/contracts/`

Alternate route: `/legacy/characters/{character_id}/contracts/`

Alternate route: `/v1/characters/{character_id}/contracts/`

---
This route is cached for up to 300 seconds
*/
func (a *Client) GetCharactersCharacterIDContracts(params *GetCharactersCharacterIDContractsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCharactersCharacterIDContractsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDContractsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_characters_character_id_contracts",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/contracts/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDContractsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCharactersCharacterIDContractsOK), nil

}

/*
GetCharactersCharacterIDContractsContractIDBids gets contract bids

Lists bids on a particular auction contract

---
Alternate route: `/dev/characters/{character_id}/contracts/{contract_id}/bids/`

Alternate route: `/legacy/characters/{character_id}/contracts/{contract_id}/bids/`

Alternate route: `/v1/characters/{character_id}/contracts/{contract_id}/bids/`

---
This route is cached for up to 300 seconds
*/
func (a *Client) GetCharactersCharacterIDContractsContractIDBids(params *GetCharactersCharacterIDContractsContractIDBidsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCharactersCharacterIDContractsContractIDBidsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDContractsContractIDBidsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_characters_character_id_contracts_contract_id_bids",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/contracts/{contract_id}/bids/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDContractsContractIDBidsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCharactersCharacterIDContractsContractIDBidsOK), nil

}

/*
GetCharactersCharacterIDContractsContractIDItems gets contract items

Lists items of a particular contract

---
Alternate route: `/dev/characters/{character_id}/contracts/{contract_id}/items/`

Alternate route: `/legacy/characters/{character_id}/contracts/{contract_id}/items/`

Alternate route: `/v1/characters/{character_id}/contracts/{contract_id}/items/`

---
This route is cached for up to 3600 seconds
*/
func (a *Client) GetCharactersCharacterIDContractsContractIDItems(params *GetCharactersCharacterIDContractsContractIDItemsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCharactersCharacterIDContractsContractIDItemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDContractsContractIDItemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_characters_character_id_contracts_contract_id_items",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/contracts/{contract_id}/items/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDContractsContractIDItemsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCharactersCharacterIDContractsContractIDItemsOK), nil

}

/*
GetCorporationsCorporationIDContracts gets corporation contracts

Returns contracts available to a corporation, only if the corporation is issuer, acceptor or assignee. Only returns contracts no older than 30 days, or if the status is "in_progress".

---
Alternate route: `/dev/corporations/{corporation_id}/contracts/`

Alternate route: `/legacy/corporations/{corporation_id}/contracts/`

Alternate route: `/v1/corporations/{corporation_id}/contracts/`

---
This route is cached for up to 300 seconds
*/
func (a *Client) GetCorporationsCorporationIDContracts(params *GetCorporationsCorporationIDContractsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCorporationsCorporationIDContractsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDContractsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_contracts",
		Method:             "GET",
		PathPattern:        "/corporations/{corporation_id}/contracts/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDContractsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCorporationsCorporationIDContractsOK), nil

}

/*
GetCorporationsCorporationIDContractsContractIDBids gets corporation contract bids

Lists bids on a particular auction contract

---
Alternate route: `/dev/corporations/{corporation_id}/contracts/{contract_id}/bids/`

Alternate route: `/legacy/corporations/{corporation_id}/contracts/{contract_id}/bids/`

Alternate route: `/v1/corporations/{corporation_id}/contracts/{contract_id}/bids/`

---
This route is cached for up to 3600 seconds
*/
func (a *Client) GetCorporationsCorporationIDContractsContractIDBids(params *GetCorporationsCorporationIDContractsContractIDBidsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCorporationsCorporationIDContractsContractIDBidsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDContractsContractIDBidsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_contracts_contract_id_bids",
		Method:             "GET",
		PathPattern:        "/corporations/{corporation_id}/contracts/{contract_id}/bids/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDContractsContractIDBidsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCorporationsCorporationIDContractsContractIDBidsOK), nil

}

/*
GetCorporationsCorporationIDContractsContractIDItems gets corporation contract items

Lists items of a particular contract

---
Alternate route: `/dev/corporations/{corporation_id}/contracts/{contract_id}/items/`

Alternate route: `/legacy/corporations/{corporation_id}/contracts/{contract_id}/items/`

Alternate route: `/v1/corporations/{corporation_id}/contracts/{contract_id}/items/`

---
This route is cached for up to 3600 seconds
*/
func (a *Client) GetCorporationsCorporationIDContractsContractIDItems(params *GetCorporationsCorporationIDContractsContractIDItemsParams, authInfo runtime.ClientAuthInfoWriter) (*GetCorporationsCorporationIDContractsContractIDItemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDContractsContractIDItemsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_contracts_contract_id_items",
		Method:             "GET",
		PathPattern:        "/corporations/{corporation_id}/contracts/{contract_id}/items/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDContractsContractIDItemsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCorporationsCorporationIDContractsContractIDItemsOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
