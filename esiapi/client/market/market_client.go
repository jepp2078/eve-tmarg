// Code generated by go-swagger; DO NOT EDIT.

package market

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new market API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for market API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetCharactersCharacterIDOrders lists open orders from a character

List open market orders placed by a character

---
Alternate route: `/dev/characters/{character_id}/orders/`

Alternate route: `/v2/characters/{character_id}/orders/`

---
This route is cached for up to 1200 seconds
*/
func (a *Client) GetCharactersCharacterIDOrders(params *GetCharactersCharacterIDOrdersParams, authInfo runtime.ClientAuthInfoWriter) (*GetCharactersCharacterIDOrdersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDOrdersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_characters_character_id_orders",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/orders/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDOrdersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCharactersCharacterIDOrdersOK), nil

}

/*
GetCharactersCharacterIDOrdersHistory lists historical orders by a character

List cancelled and expired market orders placed by a character up to 90 days in the past.

---
Alternate route: `/dev/characters/{character_id}/orders/history/`

Alternate route: `/legacy/characters/{character_id}/orders/history/`

Alternate route: `/v1/characters/{character_id}/orders/history/`

---
This route is cached for up to 3600 seconds
*/
func (a *Client) GetCharactersCharacterIDOrdersHistory(params *GetCharactersCharacterIDOrdersHistoryParams, authInfo runtime.ClientAuthInfoWriter) (*GetCharactersCharacterIDOrdersHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCharactersCharacterIDOrdersHistoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_characters_character_id_orders_history",
		Method:             "GET",
		PathPattern:        "/characters/{character_id}/orders/history/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCharactersCharacterIDOrdersHistoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCharactersCharacterIDOrdersHistoryOK), nil

}

/*
GetCorporationsCorporationIDOrders lists open orders from a corporation

List open market orders placed on behalf of a corporation

---
Alternate route: `/v2/corporations/{corporation_id}/orders/`

---
This route is cached for up to 1200 seconds

---
Requires one of the following EVE corporation role(s): Accountant, Trader


---
Warning: This route has an upgrade available.

---
[Diff of the upcoming changes](https://esi.evetech.net/diff/latest/dev/#GET-/corporations/{corporation_id}/orders/)
*/
func (a *Client) GetCorporationsCorporationIDOrders(params *GetCorporationsCorporationIDOrdersParams, authInfo runtime.ClientAuthInfoWriter) (*GetCorporationsCorporationIDOrdersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDOrdersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_orders",
		Method:             "GET",
		PathPattern:        "/corporations/{corporation_id}/orders/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDOrdersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCorporationsCorporationIDOrdersOK), nil

}

/*
GetCorporationsCorporationIDOrdersHistory lists historical orders from a corporation

List cancelled and expired market orders placed on behalf of a corporation up to 90 days in the past.

---
Alternate route: `/legacy/corporations/{corporation_id}/orders/history/`

Alternate route: `/v1/corporations/{corporation_id}/orders/history/`

---
This route is cached for up to 3600 seconds

---
Requires one of the following EVE corporation role(s): Accountant, Trader


---
Warning: This route has an upgrade available.

---
[Diff of the upcoming changes](https://esi.evetech.net/diff/latest/dev/#GET-/corporations/{corporation_id}/orders/history/)
*/
func (a *Client) GetCorporationsCorporationIDOrdersHistory(params *GetCorporationsCorporationIDOrdersHistoryParams, authInfo runtime.ClientAuthInfoWriter) (*GetCorporationsCorporationIDOrdersHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCorporationsCorporationIDOrdersHistoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_corporations_corporation_id_orders_history",
		Method:             "GET",
		PathPattern:        "/corporations/{corporation_id}/orders/history/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCorporationsCorporationIDOrdersHistoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCorporationsCorporationIDOrdersHistoryOK), nil

}

/*
GetMarketsGroups gets item groups

Get a list of item groups

---
Alternate route: `/dev/markets/groups/`

Alternate route: `/legacy/markets/groups/`

Alternate route: `/v1/markets/groups/`

---
This route expires daily at 11:05
*/
func (a *Client) GetMarketsGroups(params *GetMarketsGroupsParams) (*GetMarketsGroupsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMarketsGroupsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_markets_groups",
		Method:             "GET",
		PathPattern:        "/markets/groups/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMarketsGroupsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMarketsGroupsOK), nil

}

/*
GetMarketsGroupsMarketGroupID gets item group information

Get information on an item group

---
Alternate route: `/dev/markets/groups/{market_group_id}/`

Alternate route: `/legacy/markets/groups/{market_group_id}/`

Alternate route: `/v1/markets/groups/{market_group_id}/`

---
This route expires daily at 11:05
*/
func (a *Client) GetMarketsGroupsMarketGroupID(params *GetMarketsGroupsMarketGroupIDParams) (*GetMarketsGroupsMarketGroupIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMarketsGroupsMarketGroupIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_markets_groups_market_group_id",
		Method:             "GET",
		PathPattern:        "/markets/groups/{market_group_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMarketsGroupsMarketGroupIDReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMarketsGroupsMarketGroupIDOK), nil

}

/*
GetMarketsPrices lists market prices

Return a list of prices

---
Alternate route: `/dev/markets/prices/`

Alternate route: `/legacy/markets/prices/`

Alternate route: `/v1/markets/prices/`

---
This route is cached for up to 3600 seconds
*/
func (a *Client) GetMarketsPrices(params *GetMarketsPricesParams) (*GetMarketsPricesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMarketsPricesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_markets_prices",
		Method:             "GET",
		PathPattern:        "/markets/prices/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMarketsPricesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMarketsPricesOK), nil

}

/*
GetMarketsRegionIDHistory lists historical market statistics in a region

Return a list of historical market statistics for the specified type in a region

---
Alternate route: `/dev/markets/{region_id}/history/`

Alternate route: `/legacy/markets/{region_id}/history/`

Alternate route: `/v1/markets/{region_id}/history/`

---
This route expires daily at 11:05
*/
func (a *Client) GetMarketsRegionIDHistory(params *GetMarketsRegionIDHistoryParams) (*GetMarketsRegionIDHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMarketsRegionIDHistoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_markets_region_id_history",
		Method:             "GET",
		PathPattern:        "/markets/{region_id}/history/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMarketsRegionIDHistoryReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMarketsRegionIDHistoryOK), nil

}

/*
GetMarketsRegionIDOrders lists orders in a region

Return a list of orders in a region

---
Alternate route: `/dev/markets/{region_id}/orders/`

Alternate route: `/legacy/markets/{region_id}/orders/`

Alternate route: `/v1/markets/{region_id}/orders/`

---
This route is cached for up to 300 seconds
*/
func (a *Client) GetMarketsRegionIDOrders(params *GetMarketsRegionIDOrdersParams) (*GetMarketsRegionIDOrdersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMarketsRegionIDOrdersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_markets_region_id_orders",
		Method:             "GET",
		PathPattern:        "/markets/{region_id}/orders/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMarketsRegionIDOrdersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMarketsRegionIDOrdersOK), nil

}

/*
GetMarketsRegionIDTypes lists type ids relevant to a market

Return a list of type IDs that have active orders in the region, for efficient market indexing.

---
Alternate route: `/dev/markets/{region_id}/types/`

Alternate route: `/legacy/markets/{region_id}/types/`

Alternate route: `/v1/markets/{region_id}/types/`

---
This route is cached for up to 600 seconds
*/
func (a *Client) GetMarketsRegionIDTypes(params *GetMarketsRegionIDTypesParams) (*GetMarketsRegionIDTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMarketsRegionIDTypesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_markets_region_id_types",
		Method:             "GET",
		PathPattern:        "/markets/{region_id}/types/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMarketsRegionIDTypesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMarketsRegionIDTypesOK), nil

}

/*
GetMarketsStructuresStructureID lists orders in a structure

Return all orders in a structure

---
Alternate route: `/dev/markets/structures/{structure_id}/`

Alternate route: `/legacy/markets/structures/{structure_id}/`

Alternate route: `/v1/markets/structures/{structure_id}/`

---
This route is cached for up to 300 seconds
*/
func (a *Client) GetMarketsStructuresStructureID(params *GetMarketsStructuresStructureIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetMarketsStructuresStructureIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMarketsStructuresStructureIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_markets_structures_structure_id",
		Method:             "GET",
		PathPattern:        "/markets/structures/{structure_id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMarketsStructuresStructureIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetMarketsStructuresStructureIDOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
