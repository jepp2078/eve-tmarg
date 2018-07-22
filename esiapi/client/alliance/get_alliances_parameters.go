// Code generated by go-swagger; DO NOT EDIT.

package alliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAlliancesParams creates a new GetAlliancesParams object
// with the default values initialized.
func NewGetAlliancesParams() *GetAlliancesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetAlliancesParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAlliancesParamsWithTimeout creates a new GetAlliancesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAlliancesParamsWithTimeout(timeout time.Duration) *GetAlliancesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetAlliancesParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetAlliancesParamsWithContext creates a new GetAlliancesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAlliancesParamsWithContext(ctx context.Context) *GetAlliancesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetAlliancesParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetAlliancesParamsWithHTTPClient creates a new GetAlliancesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAlliancesParamsWithHTTPClient(client *http.Client) *GetAlliancesParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetAlliancesParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetAlliancesParams contains all the parameters to send to the API endpoint
for the get alliances operation typically these are written to a http.Request
*/
type GetAlliancesParams struct {

	/*IfNoneMatch
	  ETag from a previous request. A 304 will be returned if this matches the current ETag

	*/
	IfNoneMatch *string
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get alliances params
func (o *GetAlliancesParams) WithTimeout(timeout time.Duration) *GetAlliancesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get alliances params
func (o *GetAlliancesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get alliances params
func (o *GetAlliancesParams) WithContext(ctx context.Context) *GetAlliancesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get alliances params
func (o *GetAlliancesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get alliances params
func (o *GetAlliancesParams) WithHTTPClient(client *http.Client) *GetAlliancesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get alliances params
func (o *GetAlliancesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get alliances params
func (o *GetAlliancesParams) WithIfNoneMatch(ifNoneMatch *string) *GetAlliancesParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get alliances params
func (o *GetAlliancesParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithDatasource adds the datasource to the get alliances params
func (o *GetAlliancesParams) WithDatasource(datasource *string) *GetAlliancesParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get alliances params
func (o *GetAlliancesParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetAlliancesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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