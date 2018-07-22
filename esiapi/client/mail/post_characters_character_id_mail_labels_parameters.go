// Code generated by go-swagger; DO NOT EDIT.

package mail

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

	models "github.com/eve-tmarg/esiapi/models"
)

// NewPostCharactersCharacterIDMailLabelsParams creates a new PostCharactersCharacterIDMailLabelsParams object
// with the default values initialized.
func NewPostCharactersCharacterIDMailLabelsParams() *PostCharactersCharacterIDMailLabelsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostCharactersCharacterIDMailLabelsParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCharactersCharacterIDMailLabelsParamsWithTimeout creates a new PostCharactersCharacterIDMailLabelsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCharactersCharacterIDMailLabelsParamsWithTimeout(timeout time.Duration) *PostCharactersCharacterIDMailLabelsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostCharactersCharacterIDMailLabelsParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewPostCharactersCharacterIDMailLabelsParamsWithContext creates a new PostCharactersCharacterIDMailLabelsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCharactersCharacterIDMailLabelsParamsWithContext(ctx context.Context) *PostCharactersCharacterIDMailLabelsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostCharactersCharacterIDMailLabelsParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewPostCharactersCharacterIDMailLabelsParamsWithHTTPClient creates a new PostCharactersCharacterIDMailLabelsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCharactersCharacterIDMailLabelsParamsWithHTTPClient(client *http.Client) *PostCharactersCharacterIDMailLabelsParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostCharactersCharacterIDMailLabelsParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*PostCharactersCharacterIDMailLabelsParams contains all the parameters to send to the API endpoint
for the post characters character id mail labels operation typically these are written to a http.Request
*/
type PostCharactersCharacterIDMailLabelsParams struct {

	/*CharacterID
	  An EVE character ID

	*/
	CharacterID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*Label
	  Label to create

	*/
	Label *models.PostCharactersCharacterIDMailLabelsParamsBody
	/*Token
	  Access token to use if unable to set a header

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post characters character id mail labels params
func (o *PostCharactersCharacterIDMailLabelsParams) WithTimeout(timeout time.Duration) *PostCharactersCharacterIDMailLabelsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post characters character id mail labels params
func (o *PostCharactersCharacterIDMailLabelsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post characters character id mail labels params
func (o *PostCharactersCharacterIDMailLabelsParams) WithContext(ctx context.Context) *PostCharactersCharacterIDMailLabelsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post characters character id mail labels params
func (o *PostCharactersCharacterIDMailLabelsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post characters character id mail labels params
func (o *PostCharactersCharacterIDMailLabelsParams) WithHTTPClient(client *http.Client) *PostCharactersCharacterIDMailLabelsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post characters character id mail labels params
func (o *PostCharactersCharacterIDMailLabelsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCharacterID adds the characterID to the post characters character id mail labels params
func (o *PostCharactersCharacterIDMailLabelsParams) WithCharacterID(characterID int32) *PostCharactersCharacterIDMailLabelsParams {
	o.SetCharacterID(characterID)
	return o
}

// SetCharacterID adds the characterId to the post characters character id mail labels params
func (o *PostCharactersCharacterIDMailLabelsParams) SetCharacterID(characterID int32) {
	o.CharacterID = characterID
}

// WithDatasource adds the datasource to the post characters character id mail labels params
func (o *PostCharactersCharacterIDMailLabelsParams) WithDatasource(datasource *string) *PostCharactersCharacterIDMailLabelsParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the post characters character id mail labels params
func (o *PostCharactersCharacterIDMailLabelsParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithLabel adds the label to the post characters character id mail labels params
func (o *PostCharactersCharacterIDMailLabelsParams) WithLabel(label *models.PostCharactersCharacterIDMailLabelsParamsBody) *PostCharactersCharacterIDMailLabelsParams {
	o.SetLabel(label)
	return o
}

// SetLabel adds the label to the post characters character id mail labels params
func (o *PostCharactersCharacterIDMailLabelsParams) SetLabel(label *models.PostCharactersCharacterIDMailLabelsParamsBody) {
	o.Label = label
}

// WithToken adds the token to the post characters character id mail labels params
func (o *PostCharactersCharacterIDMailLabelsParams) WithToken(token *string) *PostCharactersCharacterIDMailLabelsParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the post characters character id mail labels params
func (o *PostCharactersCharacterIDMailLabelsParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *PostCharactersCharacterIDMailLabelsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Label != nil {
		if err := r.SetBodyParam(o.Label); err != nil {
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