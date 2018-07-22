// Code generated by go-swagger; DO NOT EDIT.

package fittings

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

// NewDeleteCharactersCharacterIDFittingsFittingIDParams creates a new DeleteCharactersCharacterIDFittingsFittingIDParams object
// with the default values initialized.
func NewDeleteCharactersCharacterIDFittingsFittingIDParams() *DeleteCharactersCharacterIDFittingsFittingIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &DeleteCharactersCharacterIDFittingsFittingIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteCharactersCharacterIDFittingsFittingIDParamsWithTimeout creates a new DeleteCharactersCharacterIDFittingsFittingIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteCharactersCharacterIDFittingsFittingIDParamsWithTimeout(timeout time.Duration) *DeleteCharactersCharacterIDFittingsFittingIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &DeleteCharactersCharacterIDFittingsFittingIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewDeleteCharactersCharacterIDFittingsFittingIDParamsWithContext creates a new DeleteCharactersCharacterIDFittingsFittingIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteCharactersCharacterIDFittingsFittingIDParamsWithContext(ctx context.Context) *DeleteCharactersCharacterIDFittingsFittingIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &DeleteCharactersCharacterIDFittingsFittingIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewDeleteCharactersCharacterIDFittingsFittingIDParamsWithHTTPClient creates a new DeleteCharactersCharacterIDFittingsFittingIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteCharactersCharacterIDFittingsFittingIDParamsWithHTTPClient(client *http.Client) *DeleteCharactersCharacterIDFittingsFittingIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &DeleteCharactersCharacterIDFittingsFittingIDParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*DeleteCharactersCharacterIDFittingsFittingIDParams contains all the parameters to send to the API endpoint
for the delete characters character id fittings fitting id operation typically these are written to a http.Request
*/
type DeleteCharactersCharacterIDFittingsFittingIDParams struct {

	/*CharacterID
	  An EVE character ID

	*/
	CharacterID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*FittingID
	  ID for a fitting of this character

	*/
	FittingID int32
	/*Token
	  Access token to use if unable to set a header

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete characters character id fittings fitting id params
func (o *DeleteCharactersCharacterIDFittingsFittingIDParams) WithTimeout(timeout time.Duration) *DeleteCharactersCharacterIDFittingsFittingIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete characters character id fittings fitting id params
func (o *DeleteCharactersCharacterIDFittingsFittingIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete characters character id fittings fitting id params
func (o *DeleteCharactersCharacterIDFittingsFittingIDParams) WithContext(ctx context.Context) *DeleteCharactersCharacterIDFittingsFittingIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete characters character id fittings fitting id params
func (o *DeleteCharactersCharacterIDFittingsFittingIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete characters character id fittings fitting id params
func (o *DeleteCharactersCharacterIDFittingsFittingIDParams) WithHTTPClient(client *http.Client) *DeleteCharactersCharacterIDFittingsFittingIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete characters character id fittings fitting id params
func (o *DeleteCharactersCharacterIDFittingsFittingIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCharacterID adds the characterID to the delete characters character id fittings fitting id params
func (o *DeleteCharactersCharacterIDFittingsFittingIDParams) WithCharacterID(characterID int32) *DeleteCharactersCharacterIDFittingsFittingIDParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the delete characters character id fittings fitting id params
func (o *DeleteCharactersCharacterIDFittingsFittingIDParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the delete characters character id fittings fitting id params
func (o *DeleteCharactersCharacterIDFittingsFittingIDParams) WithDatasource(datasource *string) *DeleteCharactersCharacterIDFittingsFittingIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the delete characters character id fittings fitting id params
func (o *DeleteCharactersCharacterIDFittingsFittingIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithFittingID adds the fittingID to the delete characters character id fittings fitting id params
func (o *DeleteCharactersCharacterIDFittingsFittingIDParams) WithFittingID(fittingID int32) *DeleteCharactersCharacterIDFittingsFittingIDParams {
	o.SetFittingID(fittingID)
	return o
}

// SetFittingID adds the fittingId to the delete characters character id fittings fitting id params
func (o *DeleteCharactersCharacterIDFittingsFittingIDParams) SetFittingID(fittingID int32) {
	o.FittingID = fittingID
}

// WithToken adds the token to the delete characters character id fittings fitting id params
func (o *DeleteCharactersCharacterIDFittingsFittingIDParams) WithToken(token *string) *DeleteCharactersCharacterIDFittingsFittingIDParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the delete characters character id fittings fitting id params
func (o *DeleteCharactersCharacterIDFittingsFittingIDParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteCharactersCharacterIDFittingsFittingIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param fitting_id
	if err := r.SetPathParam("fitting_id", swag.FormatInt32(o.FittingID)); err != nil {
		return err
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
