// Code generated by go-swagger; DO NOT EDIT.

package fleets

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

// NewDeleteFleetsFleetIDMembersMemberIDParams creates a new DeleteFleetsFleetIDMembersMemberIDParams object
// with the default values initialized.
func NewDeleteFleetsFleetIDMembersMemberIDParams() *DeleteFleetsFleetIDMembersMemberIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &DeleteFleetsFleetIDMembersMemberIDParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteFleetsFleetIDMembersMemberIDParamsWithTimeout creates a new DeleteFleetsFleetIDMembersMemberIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteFleetsFleetIDMembersMemberIDParamsWithTimeout(timeout time.Duration) *DeleteFleetsFleetIDMembersMemberIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &DeleteFleetsFleetIDMembersMemberIDParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewDeleteFleetsFleetIDMembersMemberIDParamsWithContext creates a new DeleteFleetsFleetIDMembersMemberIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteFleetsFleetIDMembersMemberIDParamsWithContext(ctx context.Context) *DeleteFleetsFleetIDMembersMemberIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &DeleteFleetsFleetIDMembersMemberIDParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewDeleteFleetsFleetIDMembersMemberIDParamsWithHTTPClient creates a new DeleteFleetsFleetIDMembersMemberIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteFleetsFleetIDMembersMemberIDParamsWithHTTPClient(client *http.Client) *DeleteFleetsFleetIDMembersMemberIDParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &DeleteFleetsFleetIDMembersMemberIDParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*DeleteFleetsFleetIDMembersMemberIDParams contains all the parameters to send to the API endpoint
for the delete fleets fleet id members member id operation typically these are written to a http.Request
*/
type DeleteFleetsFleetIDMembersMemberIDParams struct {

	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*FleetID
	  ID for a fleet

	*/
	FleetID int64
	/*MemberID
	  The character ID of a member in this fleet

	*/
	MemberID int32
	/*Token
	  Access token to use if unable to set a header

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete fleets fleet id members member id params
func (o *DeleteFleetsFleetIDMembersMemberIDParams) WithTimeout(timeout time.Duration) *DeleteFleetsFleetIDMembersMemberIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete fleets fleet id members member id params
func (o *DeleteFleetsFleetIDMembersMemberIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete fleets fleet id members member id params
func (o *DeleteFleetsFleetIDMembersMemberIDParams) WithContext(ctx context.Context) *DeleteFleetsFleetIDMembersMemberIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete fleets fleet id members member id params
func (o *DeleteFleetsFleetIDMembersMemberIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete fleets fleet id members member id params
func (o *DeleteFleetsFleetIDMembersMemberIDParams) WithHTTPClient(client *http.Client) *DeleteFleetsFleetIDMembersMemberIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete fleets fleet id members member id params
func (o *DeleteFleetsFleetIDMembersMemberIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatasource adds the datasource to the delete fleets fleet id members member id params
func (o *DeleteFleetsFleetIDMembersMemberIDParams) WithDatasource(datasource *string) *DeleteFleetsFleetIDMembersMemberIDParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the delete fleets fleet id members member id params
func (o *DeleteFleetsFleetIDMembersMemberIDParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithFleetID adds the fleetID to the delete fleets fleet id members member id params
func (o *DeleteFleetsFleetIDMembersMemberIDParams) WithFleetID(fleetID int64) *DeleteFleetsFleetIDMembersMemberIDParams {
	o.SetFleetID(fleetID)
	return o
}

// SetFleetID adds the fleetId to the delete fleets fleet id members member id params
func (o *DeleteFleetsFleetIDMembersMemberIDParams) SetFleetID(fleetID int64) {
	o.FleetID = fleetID
}

// WithMemberID adds the memberID to the delete fleets fleet id members member id params
func (o *DeleteFleetsFleetIDMembersMemberIDParams) WithMemberID(memberID int32) *DeleteFleetsFleetIDMembersMemberIDParams {
	o.SetMemberID(memberID)
	return o
}

// SetMemberID adds the memberId to the delete fleets fleet id members member id params
func (o *DeleteFleetsFleetIDMembersMemberIDParams) SetMemberID(memberID int32) {
	o.MemberID = memberID
}

// WithToken adds the token to the delete fleets fleet id members member id params
func (o *DeleteFleetsFleetIDMembersMemberIDParams) WithToken(token *string) *DeleteFleetsFleetIDMembersMemberIDParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the delete fleets fleet id members member id params
func (o *DeleteFleetsFleetIDMembersMemberIDParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteFleetsFleetIDMembersMemberIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param fleet_id
	if err := r.SetPathParam("fleet_id", swag.FormatInt64(o.FleetID)); err != nil {
		return err
	}

	// path param member_id
	if err := r.SetPathParam("member_id", swag.FormatInt32(o.MemberID)); err != nil {
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
