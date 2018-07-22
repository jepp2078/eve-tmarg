// Code generated by go-swagger; DO NOT EDIT.

package character

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

// NewGetCharactersNamesParams creates a new GetCharactersNamesParams object
// with the default values initialized.
func NewGetCharactersNamesParams() *GetCharactersNamesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersNamesParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCharactersNamesParamsWithTimeout creates a new GetCharactersNamesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCharactersNamesParamsWithTimeout(timeout time.Duration) *GetCharactersNamesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersNamesParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCharactersNamesParamsWithContext creates a new GetCharactersNamesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCharactersNamesParamsWithContext(ctx context.Context) *GetCharactersNamesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersNamesParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetCharactersNamesParamsWithHTTPClient creates a new GetCharactersNamesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCharactersNamesParamsWithHTTPClient(client *http.Client) *GetCharactersNamesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCharactersNamesParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetCharactersNamesParams contains all the parameters to send to the API endpoint
for the get characters names operation typically these are written to a http.Request
*/
type GetCharactersNamesParams struct {

	/*IfNoneMatch
	  ETag from a previous request. A 304 will be returned if this matches the current ETag

	*/
	IfNoneMatch *string
	/*CharacterIds
	  A comma separated list of character IDs

	*/
	CharacterIds []int64
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get characters names params
func (o *GetCharactersNamesParams) WithTimeout(timeout time.Duration) *GetCharactersNamesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get characters names params
func (o *GetCharactersNamesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get characters names params
func (o *GetCharactersNamesParams) WithContext(ctx context.Context) *GetCharactersNamesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get characters names params
func (o *GetCharactersNamesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get characters names params
func (o *GetCharactersNamesParams) WithHTTPClient(client *http.Client) *GetCharactersNamesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get characters names params
func (o *GetCharactersNamesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get characters names params
func (o *GetCharactersNamesParams) WithIfNoneMatch(ifNoneMatch *string) *GetCharactersNamesParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get characters names params
func (o *GetCharactersNamesParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCharacterIds adds the characterIds to the get characters names params
func (o *GetCharactersNamesParams) WithCharacterIds(characterIds []int64) *GetCharactersNamesParams {
	o.SetCharacterIds(characterIds)
	return o
}

// SetCharacterIds adds the characterIds to the get characters names params
func (o *GetCharactersNamesParams) SetCharacterIds(characterIds []int64) {
	o.CharacterIds = characterIds
}

// WithDatasource adds the datasource to the get characters names params
func (o *GetCharactersNamesParams) WithDatasource(datasource *string) *GetCharactersNamesParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get characters names params
func (o *GetCharactersNamesParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetCharactersNamesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	var valuesCharacterIds []string
	for _, v := range o.CharacterIds {
		valuesCharacterIds = append(valuesCharacterIds, swag.FormatInt64(v))
	}

	joinedCharacterIds := swag.JoinByFormat(valuesCharacterIds, "")
	// query array param character_ids
	if err := r.SetQueryParam("character_ids", joinedCharacterIds...); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}