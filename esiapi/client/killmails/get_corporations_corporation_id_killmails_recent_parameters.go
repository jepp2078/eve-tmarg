// Code generated by go-swagger; DO NOT EDIT.

package killmails

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

// NewGetCorporationsCorporationIDKillmailsRecentParams creates a new GetCorporationsCorporationIDKillmailsRecentParams object
// with the default values initialized.
func NewGetCorporationsCorporationIDKillmailsRecentParams() *GetCorporationsCorporationIDKillmailsRecentParams {
	var (
		datasourceDefault = string("tranquility")
		pageDefault       = int32(1)
	)
	return &GetCorporationsCorporationIDKillmailsRecentParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCorporationsCorporationIDKillmailsRecentParamsWithTimeout creates a new GetCorporationsCorporationIDKillmailsRecentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCorporationsCorporationIDKillmailsRecentParamsWithTimeout(timeout time.Duration) *GetCorporationsCorporationIDKillmailsRecentParams {
	var (
		datasourceDefault = string("tranquility")
		pageDefault       = int32(1)
	)
	return &GetCorporationsCorporationIDKillmailsRecentParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,

		timeout: timeout,
	}
}

// NewGetCorporationsCorporationIDKillmailsRecentParamsWithContext creates a new GetCorporationsCorporationIDKillmailsRecentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCorporationsCorporationIDKillmailsRecentParamsWithContext(ctx context.Context) *GetCorporationsCorporationIDKillmailsRecentParams {
	var (
		datasourceDefault = string("tranquility")
		pageDefault       = int32(1)
	)
	return &GetCorporationsCorporationIDKillmailsRecentParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,

		Context: ctx,
	}
}

// NewGetCorporationsCorporationIDKillmailsRecentParamsWithHTTPClient creates a new GetCorporationsCorporationIDKillmailsRecentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCorporationsCorporationIDKillmailsRecentParamsWithHTTPClient(client *http.Client) *GetCorporationsCorporationIDKillmailsRecentParams {
	var (
		datasourceDefault = string("tranquility")
		pageDefault       = int32(1)
	)
	return &GetCorporationsCorporationIDKillmailsRecentParams{
		Datasource: &datasourceDefault,
		Page:       &pageDefault,
		HTTPClient: client,
	}
}

/*GetCorporationsCorporationIDKillmailsRecentParams contains all the parameters to send to the API endpoint
for the get corporations corporation id killmails recent operation typically these are written to a http.Request
*/
type GetCorporationsCorporationIDKillmailsRecentParams struct {

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

// WithTimeout adds the timeout to the get corporations corporation id killmails recent params
func (o *GetCorporationsCorporationIDKillmailsRecentParams) WithTimeout(timeout time.Duration) *GetCorporationsCorporationIDKillmailsRecentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get corporations corporation id killmails recent params
func (o *GetCorporationsCorporationIDKillmailsRecentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get corporations corporation id killmails recent params
func (o *GetCorporationsCorporationIDKillmailsRecentParams) WithContext(ctx context.Context) *GetCorporationsCorporationIDKillmailsRecentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get corporations corporation id killmails recent params
func (o *GetCorporationsCorporationIDKillmailsRecentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get corporations corporation id killmails recent params
func (o *GetCorporationsCorporationIDKillmailsRecentParams) WithHTTPClient(client *http.Client) *GetCorporationsCorporationIDKillmailsRecentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get corporations corporation id killmails recent params
func (o *GetCorporationsCorporationIDKillmailsRecentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get corporations corporation id killmails recent params
func (o *GetCorporationsCorporationIDKillmailsRecentParams) WithIfNoneMatch(ifNoneMatch *string) *GetCorporationsCorporationIDKillmailsRecentParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get corporations corporation id killmails recent params
func (o *GetCorporationsCorporationIDKillmailsRecentParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithCorporationID adds the corporationID to the get corporations corporation id killmails recent params
func (o *GetCorporationsCorporationIDKillmailsRecentParams) WithCorporationID(corporationID int32) *GetCorporationsCorporationIDKillmailsRecentParams {
	o.SetCorporationID(corporationID)
	return o
}

// SetCorporationID adds the corporationId to the get corporations corporation id killmails recent params
func (o *GetCorporationsCorporationIDKillmailsRecentParams) SetCorporationID(corporationID int32) {
	o.CorporationID = corporationID
}

// WithDatasource adds the datasource to the get corporations corporation id killmails recent params
func (o *GetCorporationsCorporationIDKillmailsRecentParams) WithDatasource(datasource *string) *GetCorporationsCorporationIDKillmailsRecentParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get corporations corporation id killmails recent params
func (o *GetCorporationsCorporationIDKillmailsRecentParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithPage adds the page to the get corporations corporation id killmails recent params
func (o *GetCorporationsCorporationIDKillmailsRecentParams) WithPage(page *int32) *GetCorporationsCorporationIDKillmailsRecentParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get corporations corporation id killmails recent params
func (o *GetCorporationsCorporationIDKillmailsRecentParams) SetPage(page *int32) {
	o.Page = page
}

// WithToken adds the token to the get corporations corporation id killmails recent params
func (o *GetCorporationsCorporationIDKillmailsRecentParams) WithToken(token *string) *GetCorporationsCorporationIDKillmailsRecentParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the get corporations corporation id killmails recent params
func (o *GetCorporationsCorporationIDKillmailsRecentParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *GetCorporationsCorporationIDKillmailsRecentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
