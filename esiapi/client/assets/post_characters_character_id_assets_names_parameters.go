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

// NewPostCharactersCharacterIDAssetsNamesParams creates a new PostCharactersCharacterIDAssetsNamesParams object
// with the default values initialized.
func NewPostCharactersCharacterIDAssetsNamesParams() *PostCharactersCharacterIDAssetsNamesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostCharactersCharacterIDAssetsNamesParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCharactersCharacterIDAssetsNamesParamsWithTimeout creates a new PostCharactersCharacterIDAssetsNamesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCharactersCharacterIDAssetsNamesParamsWithTimeout(timeout time.Duration) *PostCharactersCharacterIDAssetsNamesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostCharactersCharacterIDAssetsNamesParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewPostCharactersCharacterIDAssetsNamesParamsWithContext creates a new PostCharactersCharacterIDAssetsNamesParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCharactersCharacterIDAssetsNamesParamsWithContext(ctx context.Context) *PostCharactersCharacterIDAssetsNamesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostCharactersCharacterIDAssetsNamesParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewPostCharactersCharacterIDAssetsNamesParamsWithHTTPClient creates a new PostCharactersCharacterIDAssetsNamesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCharactersCharacterIDAssetsNamesParamsWithHTTPClient(client *http.Client) *PostCharactersCharacterIDAssetsNamesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostCharactersCharacterIDAssetsNamesParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*PostCharactersCharacterIDAssetsNamesParams contains all the parameters to send to the API endpoint
for the post characters character id assets names operation typically these are written to a http.Request
*/
type PostCharactersCharacterIDAssetsNamesParams struct {

	/*CharacterID
	  An EVE character ID

	*/
	CharacterID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*ItemIds
	  A list of item ids

	*/
	ItemIds []int64
	/*Token
	  Access token to use if unable to set a header

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post characters character id assets names params
func (o *PostCharactersCharacterIDAssetsNamesParams) WithTimeout(timeout time.Duration) *PostCharactersCharacterIDAssetsNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post characters character id assets names params
func (o *PostCharactersCharacterIDAssetsNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post characters character id assets names params
func (o *PostCharactersCharacterIDAssetsNamesParams) WithContext(ctx context.Context) *PostCharactersCharacterIDAssetsNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post characters character id assets names params
func (o *PostCharactersCharacterIDAssetsNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post characters character id assets names params
func (o *PostCharactersCharacterIDAssetsNamesParams) WithHTTPClient(client *http.Client) *PostCharactersCharacterIDAssetsNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post characters character id assets names params
func (o *PostCharactersCharacterIDAssetsNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCharacterID adds the characterID to the post characters character id assets names params
func (o *PostCharactersCharacterIDAssetsNamesParams) WithCharacterID(characterID int32) *PostCharactersCharacterIDAssetsNamesParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the post characters character id assets names params
func (o *PostCharactersCharacterIDAssetsNamesParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the post characters character id assets names params
func (o *PostCharactersCharacterIDAssetsNamesParams) WithDatasource(datasource *string) *PostCharactersCharacterIDAssetsNamesParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the post characters character id assets names params
func (o *PostCharactersCharacterIDAssetsNamesParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithItemIds adds the itemIds to the post characters character id assets names params
func (o *PostCharactersCharacterIDAssetsNamesParams) WithItemIds(itemIds []int64) *PostCharactersCharacterIDAssetsNamesParams {
	o.SetItemIds(itemIds)
	return o
}

// SetItemIds adds the itemIds to the post characters character id assets names params
func (o *PostCharactersCharacterIDAssetsNamesParams) SetItemIds(itemIds []int64) {
	o.ItemIds = itemIds
}

// WithToken adds the token to the post characters character id assets names params
func (o *PostCharactersCharacterIDAssetsNamesParams) WithToken(token *string) *PostCharactersCharacterIDAssetsNamesParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the post characters character id assets names params
func (o *PostCharactersCharacterIDAssetsNamesParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *PostCharactersCharacterIDAssetsNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	if o.ItemIds != nil {
		if err := r.SetBodyParam(o.ItemIds); err != nil {
			return err
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
