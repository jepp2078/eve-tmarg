// Code generated by go-swagger; DO NOT EDIT.

package universe

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

// NewGetUniverseStructuresParams creates a new GetUniverseStructuresParams object
// with the default values initialized.
func NewGetUniverseStructuresParams() *GetUniverseStructuresParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseStructuresParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUniverseStructuresParamsWithTimeout creates a new GetUniverseStructuresParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUniverseStructuresParamsWithTimeout(timeout time.Duration) *GetUniverseStructuresParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseStructuresParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetUniverseStructuresParamsWithContext creates a new GetUniverseStructuresParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUniverseStructuresParamsWithContext(ctx context.Context) *GetUniverseStructuresParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseStructuresParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetUniverseStructuresParamsWithHTTPClient creates a new GetUniverseStructuresParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUniverseStructuresParamsWithHTTPClient(client *http.Client) *GetUniverseStructuresParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseStructuresParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetUniverseStructuresParams contains all the parameters to send to the API endpoint
for the get universe structures operation typically these are written to a http.Request
*/
type GetUniverseStructuresParams struct {

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

// WithTimeout adds the timeout to the get universe structures params
func (o *GetUniverseStructuresParams) WithTimeout(timeout time.Duration) *GetUniverseStructuresParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get universe structures params
func (o *GetUniverseStructuresParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get universe structures params
func (o *GetUniverseStructuresParams) WithContext(ctx context.Context) *GetUniverseStructuresParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get universe structures params
func (o *GetUniverseStructuresParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get universe structures params
func (o *GetUniverseStructuresParams) WithHTTPClient(client *http.Client) *GetUniverseStructuresParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get universe structures params
func (o *GetUniverseStructuresParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get universe structures params
func (o *GetUniverseStructuresParams) WithIfNoneMatch(ifNoneMatch *string) *GetUniverseStructuresParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get universe structures params
func (o *GetUniverseStructuresParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithDatasource adds the datasource to the get universe structures params
func (o *GetUniverseStructuresParams) WithDatasource(datasource *string) *GetUniverseStructuresParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get universe structures params
func (o *GetUniverseStructuresParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetUniverseStructuresParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
