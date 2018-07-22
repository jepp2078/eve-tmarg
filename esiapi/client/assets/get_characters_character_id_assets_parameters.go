// Code generated by go-swagger; DO NOT EDIT.

package assets

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

// NewGetCharactersCharacterIDAssetsParams creates a new GetCharactersCharacterIDAssetsParams object
// with the default values initialized.
func NewGetCharactersCharacterIDAssetsParams() *GetCharactersCharacterIDAssetsParams {
	var (
		datasourceDefault = string("tranquility")
		pageDefault       = int32(1)
	)
	return &GetCharactersCharacterIDAssetsParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersCharacterIDAssetsParamsWithTimeout creates a new GetCharactersCharacterIDAssetsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersCharacterIDAssetsParamsWithTimeout(timeout time.Duration) *GetCharactersCharacterIDAssetsParams {
	var (
		datasourceDefault = string("tranquility")
		pageDefault       = int32(1)
	)
	return &GetCharactersCharacterIDAssetsParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,

		timeout: timeout,
	}
}

// NewGetCharactersCharacterIDAssetsParamsWithContext creates a new GetCharactersCharacterIDAssetsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersCharacterIDAssetsParamsWithContext(ctx context.Context) *GetCharactersCharacterIDAssetsParams {
	var (
		datasourceDefault = string("tranquility")
		pageDefault       = int32(1)
	)
	return &GetCharactersCharacterIDAssetsParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,

		Context: ctx,
	}
}

// NewGetCharactersCharacterIDAssetsParamsWithHTTPClient creates a new GetCharactersCharacterIDAssetsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCharactersCharacterIDAssetsParamsWithHTTPClient(client *http.Client) *GetCharactersCharacterIDAssetsParams {
	var (
		datasourceDefault = string("tranquility")
		pageDefault       = int32(1)
	)
	return &GetCharactersCharacterIDAssetsParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,
		HTTPClient: client,
	}
}

/*GetCharactersCharacterIDAssetsParams contains all the parameters to send to the API endpoint
for the get characters character id assets operation typically these are written to a http.Request
*/
type GetCharactersCharacterIDAssetsParams struct {

	/*IfNoneMatch
	  ETag from a previous request. A 304 will be returned if this matches the current ETag

	*/
	IfNoneMatch *string
	/*CharacterID
	  An EVE character ID

	*/
	CharacterID int32
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

// WithTimeout adds the timeout to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) WithTimeout(timeout time.Duration) *GetCharactersCharacterIDAssetsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) WithContext(ctx context.Context) *GetCharactersCharacterIDAssetsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) WithHTTPClient(client *http.Client) *GetCharactersCharacterIDAssetsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) WithIfNoneMatch(ifNoneMatch *string) *GetCharactersCharacterIDAssetsParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCharacterID adds the characterID to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) WithCharacterID(characterID int32) *GetCharactersCharacterIDAssetsParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) WithDatasource(datasource *string) *GetCharactersCharacterIDAssetsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithPage adds the page to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) WithPage(page *int32) *GetCharactersCharacterIDAssetsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) SetPage(page *int32) {
	o.Page = page
}

// WithToken adds the token to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) WithToken(token *string) *GetCharactersCharacterIDAssetsParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get characters character id assets params
func (o *GetCharactersCharacterIDAssetsParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersCharacterIDAssetsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param character_id
	if err := r.SetPathParam("character_id", swag.FormatInt32(o.CharacterID)); err != nil {
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