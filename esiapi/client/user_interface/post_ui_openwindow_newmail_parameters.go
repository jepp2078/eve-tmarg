// Code generated by go-swagger; DO NOT EDIT.

package user_interface

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

	models "github.com/eve-tmarg/esiapi/models"
)

// NewPostUIOpenwindowNewmailParams creates a new PostUIOpenwindowNewmailParams object
// with the default values initialized.
func NewPostUIOpenwindowNewmailParams() *PostUIOpenwindowNewmailParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostUIOpenwindowNewmailParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostUIOpenwindowNewmailParamsWithTimeout creates a new PostUIOpenwindowNewmailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostUIOpenwindowNewmailParamsWithTimeout(timeout time.Duration) *PostUIOpenwindowNewmailParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostUIOpenwindowNewmailParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewPostUIOpenwindowNewmailParamsWithContext creates a new PostUIOpenwindowNewmailParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostUIOpenwindowNewmailParamsWithContext(ctx context.Context) *PostUIOpenwindowNewmailParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostUIOpenwindowNewmailParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewPostUIOpenwindowNewmailParamsWithHTTPClient creates a new PostUIOpenwindowNewmailParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostUIOpenwindowNewmailParamsWithHTTPClient(client *http.Client) *PostUIOpenwindowNewmailParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostUIOpenwindowNewmailParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*PostUIOpenwindowNewmailParams contains all the parameters to send to the API endpoint
for the post ui openwindow newmail operation typically these are written to a http.Request
*/
type PostUIOpenwindowNewmailParams struct {

	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string
	/*NewMail
	  The details of mail to create

	*/
	NewMail *models.PostUIOpenwindowNewmailParamsBody
	/*Token
	  Access token to use if unable to set a header

	*/
	Token *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post ui openwindow newmail params
func (o *PostUIOpenwindowNewmailParams) WithTimeout(timeout time.Duration) *PostUIOpenwindowNewmailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post ui openwindow newmail params
func (o *PostUIOpenwindowNewmailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post ui openwindow newmail params
func (o *PostUIOpenwindowNewmailParams) WithContext(ctx context.Context) *PostUIOpenwindowNewmailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post ui openwindow newmail params
func (o *PostUIOpenwindowNewmailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post ui openwindow newmail params
func (o *PostUIOpenwindowNewmailParams) WithHTTPClient(client *http.Client) *PostUIOpenwindowNewmailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post ui openwindow newmail params
func (o *PostUIOpenwindowNewmailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatasource adds the datasource to the post ui openwindow newmail params
func (o *PostUIOpenwindowNewmailParams) WithDatasource(datasource *string) *PostUIOpenwindowNewmailParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the post ui openwindow newmail params
func (o *PostUIOpenwindowNewmailParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WithNewMail adds the newMail to the post ui openwindow newmail params
func (o *PostUIOpenwindowNewmailParams) WithNewMail(newMail *models.PostUIOpenwindowNewmailParamsBody) *PostUIOpenwindowNewmailParams {
	o.SetNewMail(newMail)
	return o
}

// SetNewMail adds the newMail to the post ui openwindow newmail params
func (o *PostUIOpenwindowNewmailParams) SetNewMail(newMail *models.PostUIOpenwindowNewmailParamsBody) {
	o.NewMail = newMail
}

// WithToken adds the token to the post ui openwindow newmail params
func (o *PostUIOpenwindowNewmailParams) WithToken(token *string) *PostUIOpenwindowNewmailParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the post ui openwindow newmail params
func (o *PostUIOpenwindowNewmailParams) SetToken(token *string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *PostUIOpenwindowNewmailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.NewMail != nil {
		if err := r.SetBodyParam(o.NewMail); err != nil {
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