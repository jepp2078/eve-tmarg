// Code generated by go-swagger; DO NOT EDIT.

package status

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

// NewGetStatusParams creates a new GetStatusParams object
// with the default values initialized.
func NewGetStatusParams() *GetStatusParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetStatusParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStatusParamsWithTimeout creates a new GetStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStatusParamsWithTimeout(timeout time.Duration) *GetStatusParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetStatusParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetStatusParamsWithContext creates a new GetStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStatusParamsWithContext(ctx context.Context) *GetStatusParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetStatusParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetStatusParamsWithHTTPClient creates a new GetStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStatusParamsWithHTTPClient(client *http.Client) *GetStatusParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetStatusParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetStatusParams contains all the parameters to send to the API endpoint
for the get status operation typically these are written to a http.Request
*/
type GetStatusParams struct {

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

// WithTimeout adds the timeout to the get status params
func (o *GetStatusParams) WithTimeout(timeout time.Duration) *GetStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get status params
func (o *GetStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get status params
func (o *GetStatusParams) WithContext(ctx context.Context) *GetStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get status params
func (o *GetStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get status params
func (o *GetStatusParams) WithHTTPClient(client *http.Client) *GetStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get status params
func (o *GetStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get status params
func (o *GetStatusParams) WithIfNoneMatch(ifNoneMatch *string) *GetStatusParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get status params
func (o *GetStatusParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithDatasource adds the datasource to the get status params
func (o *GetStatusParams) WithDatasource(datasource *string) *GetStatusParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get status params
func (o *GetStatusParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
