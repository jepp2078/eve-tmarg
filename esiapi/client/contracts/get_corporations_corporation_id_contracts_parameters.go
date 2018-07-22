// Code generated by go-swagger; DO NOT EDIT.

package contracts

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

// NewGetCorporationsCorporationIDContractsParams creates a new GetCorporationsCorporationIDContractsParams object
// with the default values initialized.
func NewGetCorporationsCorporationIDContractsParams() *GetCorporationsCorporationIDContractsParams {
	var (
		datasourceDefault = string("tranquility")
		pageDefault       = int32(1)
	)
	return &GetCorporationsCorporationIDContractsParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCorporationsCorporationIDContractsParamsWithTimeout creates a new GetCorporationsCorporationIDContractsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCorporationsCorporationIDContractsParamsWithTimeout(timeout time.Duration) *GetCorporationsCorporationIDContractsParams {
	var (
		datasourceDefault = string("tranquility")
		pageDefault       = int32(1)
	)
	return &GetCorporationsCorporationIDContractsParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,

		timeout: timeout,
	}
}

// NewGetCorporationsCorporationIDContractsParamsWithContext creates a new GetCorporationsCorporationIDContractsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCorporationsCorporationIDContractsParamsWithContext(ctx context.Context) *GetCorporationsCorporationIDContractsParams {
	var (
		datasourceDefault = string("tranquility")
		pageDefault       = int32(1)
	)
	return &GetCorporationsCorporationIDContractsParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,

		Context: ctx,
	}
}

// NewGetCorporationsCorporationIDContractsParamsWithHTTPClient creates a new GetCorporationsCorporationIDContractsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCorporationsCorporationIDContractsParamsWithHTTPClient(client *http.Client) *GetCorporationsCorporationIDContractsParams {
	var (
		datasourceDefault = string("tranquility")
		pageDefault       = int32(1)
	)
	return &GetCorporationsCorporationIDContractsParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,
		HTTPClient: client,
	}
}

/*GetCorporationsCorporationIDContractsParams contains all the parameters to send to the API endpoint
for the get corporations corporation id contracts operation typically these are written to a http.Request
*/
type GetCorporationsCorporationIDContractsParams struct {

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
	/*Page
	  Which page of results to return

	*/
	Page *int32
	/*Token
	  Access token to use if unable to set a header

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get corporations corporation id contracts params
func (o *GetCorporationsCorporationIDContractsParams) WithTimeout(timeout time.Duration) *GetCorporationsCorporationIDContractsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get corporations corporation id contracts params
func (o *GetCorporationsCorporationIDContractsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get corporations corporation id contracts params
func (o *GetCorporationsCorporationIDContractsParams) WithContext(ctx context.Context) *GetCorporationsCorporationIDContractsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get corporations corporation id contracts params
func (o *GetCorporationsCorporationIDContractsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get corporations corporation id contracts params
func (o *GetCorporationsCorporationIDContractsParams) WithHTTPClient(client *http.Client) *GetCorporationsCorporationIDContractsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get corporations corporation id contracts params
func (o *GetCorporationsCorporationIDContractsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get corporations corporation id contracts params
func (o *GetCorporationsCorporationIDContractsParams) WithIfNoneMatch(ifNoneMatch *string) *GetCorporationsCorporationIDContractsParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get corporations corporation id contracts params
func (o *GetCorporationsCorporationIDContractsParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCorporationID adds the corporationID to the get corporations corporation id contracts params
func (o *GetCorporationsCorporationIDContractsParams) WithCorporationID(corporationID int32) *GetCorporationsCorporationIDContractsParams {
	o.SetCorporationID(corporationID)
	return o
}

// SetCorporationID adds the corporationId to the get corporations corporation id contracts params
func (o *GetCorporationsCorporationIDContractsParams) SetCorporationID(corporationID int32) {
	o.CorporationID = corporationID
}

// WithDatasource adds the datasource to the get corporations corporation id contracts params
func (o *GetCorporationsCorporationIDContractsParams) WithDatasource(datasource *string) *GetCorporationsCorporationIDContractsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get corporations corporation id contracts params
func (o *GetCorporationsCorporationIDContractsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithPage adds the page to the get corporations corporation id contracts params
func (o *GetCorporationsCorporationIDContractsParams) WithPage(page *int32) *GetCorporationsCorporationIDContractsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get corporations corporation id contracts params
func (o *GetCorporationsCorporationIDContractsParams) SetPage(page *int32) {
	o.Page = page
}

// WithToken adds the token to the get corporations corporation id contracts params
func (o *GetCorporationsCorporationIDContractsParams) WithToken(token *string) *GetCorporationsCorporationIDContractsParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get corporations corporation id contracts params
func (o *GetCorporationsCorporationIDContractsParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetCorporationsCorporationIDContractsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Page != nil {

		// query param page
		var qrPage int32
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatInt32(qrPage)
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
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