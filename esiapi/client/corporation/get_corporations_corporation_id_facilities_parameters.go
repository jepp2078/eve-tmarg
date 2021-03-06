// Code generated by go-swagger; DO NOT EDIT.

package corporation

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

// NewGetCorporationsCorporationIDFacilitiesParams creates a new GetCorporationsCorporationIDFacilitiesParams object
// with the default values initialized.
func NewGetCorporationsCorporationIDFacilitiesParams() *GetCorporationsCorporationIDFacilitiesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDFacilitiesParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCorporationsCorporationIDFacilitiesParamsWithTimeout creates a new GetCorporationsCorporationIDFacilitiesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCorporationsCorporationIDFacilitiesParamsWithTimeout(timeout time.Duration) *GetCorporationsCorporationIDFacilitiesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDFacilitiesParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetCorporationsCorporationIDFacilitiesParamsWithContext creates a new GetCorporationsCorporationIDFacilitiesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCorporationsCorporationIDFacilitiesParamsWithContext(ctx context.Context) *GetCorporationsCorporationIDFacilitiesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDFacilitiesParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetCorporationsCorporationIDFacilitiesParamsWithHTTPClient creates a new GetCorporationsCorporationIDFacilitiesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCorporationsCorporationIDFacilitiesParamsWithHTTPClient(client *http.Client) *GetCorporationsCorporationIDFacilitiesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetCorporationsCorporationIDFacilitiesParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetCorporationsCorporationIDFacilitiesParams contains all the parameters to send to the API endpoint
for the get corporations corporation id facilities operation typically these are written to a http.Request
*/
type GetCorporationsCorporationIDFacilitiesParams struct {

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

// WithTimeout adds the timeout to the get corporations corporation id facilities params
func (o *GetCorporationsCorporationIDFacilitiesParams) WithTimeout(timeout time.Duration) *GetCorporationsCorporationIDFacilitiesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get corporations corporation id facilities params
func (o *GetCorporationsCorporationIDFacilitiesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get corporations corporation id facilities params
func (o *GetCorporationsCorporationIDFacilitiesParams) WithContext(ctx context.Context) *GetCorporationsCorporationIDFacilitiesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get corporations corporation id facilities params
func (o *GetCorporationsCorporationIDFacilitiesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get corporations corporation id facilities params
func (o *GetCorporationsCorporationIDFacilitiesParams) WithHTTPClient(client *http.Client) *GetCorporationsCorporationIDFacilitiesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get corporations corporation id facilities params
func (o *GetCorporationsCorporationIDFacilitiesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get corporations corporation id facilities params
func (o *GetCorporationsCorporationIDFacilitiesParams) WithIfNoneMatch(ifNoneMatch *string) *GetCorporationsCorporationIDFacilitiesParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get corporations corporation id facilities params
func (o *GetCorporationsCorporationIDFacilitiesParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCorporationID adds the corporationID to the get corporations corporation id facilities params
func (o *GetCorporationsCorporationIDFacilitiesParams) WithCorporationID(corporationID int32) *GetCorporationsCorporationIDFacilitiesParams {
	o.SetCorporationID(corporationID)
	return o
}

// SetCorporationID adds the corporationId to the get corporations corporation id facilities params
func (o *GetCorporationsCorporationIDFacilitiesParams) SetCorporationID(corporationID int32) {
	o.CorporationID = corporationID
}

// WithDatasource adds the datasource to the get corporations corporation id facilities params
func (o *GetCorporationsCorporationIDFacilitiesParams) WithDatasource(datasource *string) *GetCorporationsCorporationIDFacilitiesParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get corporations corporation id facilities params
func (o *GetCorporationsCorporationIDFacilitiesParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithToken adds the token to the get corporations corporation id facilities params
func (o *GetCorporationsCorporationIDFacilitiesParams) WithToken(token *string) *GetCorporationsCorporationIDFacilitiesParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get corporations corporation id facilities params
func (o *GetCorporationsCorporationIDFacilitiesParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetCorporationsCorporationIDFacilitiesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
