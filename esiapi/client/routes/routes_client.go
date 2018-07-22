// Code generated by go-swagger; DO NOT EDIT.

package routes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new routes API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for routes API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetRouteOriginDestination gets route

Get the systems between origin and destination

---
Alternate route: `/dev/route/{origin}/{destination}/`

Alternate route: `/legacy/route/{origin}/{destination}/`

Alternate route: `/v1/route/{origin}/{destination}/`

---
This route is cached for up to 86400 seconds
*/
func (a *Client) GetRouteOriginDestination(params *GetRouteOriginDestinationParams) (*GetRouteOriginDestinationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRouteOriginDestinationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_route_origin_destination",
		Method:             "GET",
		PathPattern:        "/route/{origin}/{destination}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetRouteOriginDestinationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRouteOriginDestinationOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}