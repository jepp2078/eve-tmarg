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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetUniverseAsteroidBeltsAsteroidBeltIDParams creates a new GetUniverseAsteroidBeltsAsteroidBeltIDParams object
// with the default values initialized.
func NewGetUniverseAsteroidBeltsAsteroidBeltIDParams() *GetUniverseAsteroidBeltsAsteroidBeltIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseAsteroidBeltsAsteroidBeltIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUniverseAsteroidBeltsAsteroidBeltIDParamsWithTimeout creates a new GetUniverseAsteroidBeltsAsteroidBeltIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUniverseAsteroidBeltsAsteroidBeltIDParamsWithTimeout(timeout time.Duration) *GetUniverseAsteroidBeltsAsteroidBeltIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseAsteroidBeltsAsteroidBeltIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewGetUniverseAsteroidBeltsAsteroidBeltIDParamsWithContext creates a new GetUniverseAsteroidBeltsAsteroidBeltIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUniverseAsteroidBeltsAsteroidBeltIDParamsWithContext(ctx context.Context) *GetUniverseAsteroidBeltsAsteroidBeltIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseAsteroidBeltsAsteroidBeltIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewGetUniverseAsteroidBeltsAsteroidBeltIDParamsWithHTTPClient creates a new GetUniverseAsteroidBeltsAsteroidBeltIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUniverseAsteroidBeltsAsteroidBeltIDParamsWithHTTPClient(client *http.Client) *GetUniverseAsteroidBeltsAsteroidBeltIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &GetUniverseAsteroidBeltsAsteroidBeltIDParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*GetUniverseAsteroidBeltsAsteroidBeltIDParams contains all the parameters to send to the API endpoint
for the get universe asteroid belts asteroid belt id operation typically these are written to a http.Request
*/
type GetUniverseAsteroidBeltsAsteroidBeltIDParams struct {

	/*IfNoneMatch
	  ETag from a previous request. A 304 will be returned if this matches the current ETag

	*/
	IfNoneMatch *string
	/*AsteroidBeltID
	  asteroid_belt_id integer

	*/
	AsteroidBeltID int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get universe asteroid belts asteroid belt id params
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDParams) WithTimeout(timeout time.Duration) *GetUniverseAsteroidBeltsAsteroidBeltIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get universe asteroid belts asteroid belt id params
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get universe asteroid belts asteroid belt id params
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDParams) WithContext(ctx context.Context) *GetUniverseAsteroidBeltsAsteroidBeltIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get universe asteroid belts asteroid belt id params
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get universe asteroid belts asteroid belt id params
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDParams) WithHTTPClient(client *http.Client) *GetUniverseAsteroidBeltsAsteroidBeltIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get universe asteroid belts asteroid belt id params
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIfNoneMatch adds the ifNoneMatch to the get universe asteroid belts asteroid belt id params
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDParams) WithIfNoneMatch(ifNoneMatch *string) *GetUniverseAsteroidBeltsAsteroidBeltIDParams {
	o.SetIfNoneMatch(ifNoneMatch)
	return o
}

// SetIfNoneMatch adds the ifNoneMatch to the get universe asteroid belts asteroid belt id params
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDParams) SetIfNoneMatch(ifNoneMatch *string) {
	o.IfNoneMatch = ifNoneMatch
}

// WithAsteroidBeltID adds the asteroidBeltID to the get universe asteroid belts asteroid belt id params
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDParams) WithAsteroidBeltID(asteroidBeltID int32) *GetUniverseAsteroidBeltsAsteroidBeltIDParams {
	o.SetAsteroidBeltID(asteroidBeltID)
	return o
}

// SetAsteroidBeltID adds the asteroidBeltId to the get universe asteroid belts asteroid belt id params
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDParams) SetAsteroidBeltID(asteroidBeltID int32) {
	o.AsteroidBeltID = asteroidBeltID
}

// WithDatasource adds the datasource to the get universe asteroid belts asteroid belt id params
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDParams) WithDatasource(datasource *string) *GetUniverseAsteroidBeltsAsteroidBeltIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the get universe asteroid belts asteroid belt id params
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param asteroid_belt_id
	if err := r.SetPathParam("asteroid_belt_id", swag.FormatInt32(o.AsteroidBeltID)); err != nil {
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