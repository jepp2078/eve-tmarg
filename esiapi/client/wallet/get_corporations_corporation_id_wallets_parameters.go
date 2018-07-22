// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCorporationsCorporationIDWalletsParams creates a new GetCorporationsCorporationIDWalletsParams object
// with the default values initialized.
func NewGetCorporationsCorporationIDWalletsParams() *GetCorporationsCorporationIDWalletsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDWalletsParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCorporationsCorporationIDWalletsParamsWithTimeout creates a new GetCorporationsCorporationIDWalletsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCorporationsCorporationIDWalletsParamsWithTimeout(timeout time.Duration) *GetCorporationsCorporationIDWalletsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDWalletsParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCorporationsCorporationIDWalletsParamsWithContext creates a new GetCorporationsCorporationIDWalletsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCorporationsCorporationIDWalletsParamsWithContext(ctx context.Context) *GetCorporationsCorporationIDWalletsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDWalletsParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetCorporationsCorporationIDWalletsParamsWithHTTPClient creates a new GetCorporationsCorporationIDWalletsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCorporationsCorporationIDWalletsParamsWithHTTPClient(client *http.Client) *GetCorporationsCorporationIDWalletsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDWalletsParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetCorporationsCorporationIDWalletsParams contains all the parameters to send to the API endpoint
for the get corporations corporation id wallets operation typically these are written to a http.Request
*/
type GetCorporationsCorporationIDWalletsParams struct {

	/*IfNoneMatch
	  ETag from a previous request. A 304 will be returned if this matches the current ETag

	*/
	IfNoneMatch *string
	/*CorporationID
	  An EVE corporation ID

	*/
	CorporationID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*Token
	  Access token to use if unable to set a header

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get corporations corporation id wallets params
func (o *GetCorporationsCorporationIDWalletsParams) WithTimeout(timeout time.Duration) *GetCorporationsCorporationIDWalletsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get corporations corporation id wallets params
func (o *GetCorporationsCorporationIDWalletsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get corporations corporation id wallets params
func (o *GetCorporationsCorporationIDWalletsParams) WithContext(ctx context.Context) *GetCorporationsCorporationIDWalletsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get corporations corporation id wallets params
func (o *GetCorporationsCorporationIDWalletsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get corporations corporation id wallets params
func (o *GetCorporationsCorporationIDWalletsParams) WithHTTPClient(client *http.Client) *GetCorporationsCorporationIDWalletsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get corporations corporation id wallets params
func (o *GetCorporationsCorporationIDWalletsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get corporations corporation id wallets params
func (o *GetCorporationsCorporationIDWalletsParams) WithIfNoneMatch(ifNoneMatch *string) *GetCorporationsCorporationIDWalletsParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get corporations corporation id wallets params
func (o *GetCorporationsCorporationIDWalletsParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCorporationID adds the corporationID to the get corporations corporation id wallets params
func (o *GetCorporationsCorporationIDWalletsParams) WithCorporationID(corporationID int32) *GetCorporationsCorporationIDWalletsParams {
	o.SetCorporationID(corporationID)
	return o
}

// SetCorporationID adds the corporationId to the get corporations corporation id wallets params
func (o *GetCorporationsCorporationIDWalletsParams) SetCorporationID(corporationID int32) {
	o.CorporationID = corporationID
}

// WithDatasource adds the datasource to the get corporations corporation id wallets params
func (o *GetCorporationsCorporationIDWalletsParams) WithDatasource(datasource *string) *GetCorporationsCorporationIDWalletsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get corporations corporation id wallets params
func (o *GetCorporationsCorporationIDWalletsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithToken adds the token to the get corporations corporation id wallets params
func (o *GetCorporationsCorporationIDWalletsParams) WithToken(token *string) *GetCorporationsCorporationIDWalletsParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get corporations corporation id wallets params
func (o *GetCorporationsCorporationIDWalletsParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetCorporationsCorporationIDWalletsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.IfNoneMatch != nil {

		// header param If-None-Match
		if err := r.SetHeaderParam("If-None-Match", *o.IfNoneMatch); err != nil {
			return err
		}

	}

	// path param corporation_id
	if err := r.SetPathParam("corporation_id", swag.FormatInt32(o.CorporationID)); err != nil {
		return err
	}

	if o.Datasource != nil {

		// query param datasource
		var qrDatasource string
		if o.Datasource != nil {
			qrDatasource = *o.Datasource
		}
		qDatasource := qrDatasource
		if qDatasource != "" {
			if err := r.SetQueryParam("datasource", qDatasource); err != nil {
				return err
			}
		}

	}

	if o.Token != nil {

		// query param token
		var qrToken string
		if o.Token != nil {
			qrToken = *o.Token
		}
		qToken := qrToken
		if qToken != "" {
			if err := r.SetQueryParam("token", qToken); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}