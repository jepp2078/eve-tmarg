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

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostCharactersAffiliationParams creates a new PostCharactersAffiliationParams object
// with the default values initialized.
func NewPostCharactersAffiliationParams() *PostCharactersAffiliationParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostCharactersAffiliationParams{
		Datasource: &datasourceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPostCharactersAffiliationParamsWithTimeout creates a new PostCharactersAffiliationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostCharactersAffiliationParamsWithTimeout(timeout time.Duration) *PostCharactersAffiliationParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostCharactersAffiliationParams{
		Datasource: &datasourceDefault,

		timeout: timeout,
	}
}

// NewPostCharactersAffiliationParamsWithContext creates a new PostCharactersAffiliationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostCharactersAffiliationParamsWithContext(ctx context.Context) *PostCharactersAffiliationParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostCharactersAffiliationParams{
		Datasource: &datasourceDefault,

		Context: ctx,
	}
}

// NewPostCharactersAffiliationParamsWithHTTPClient creates a new PostCharactersAffiliationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostCharactersAffiliationParamsWithHTTPClient(client *http.Client) *PostCharactersAffiliationParams {
	var (
		datasourceDefault = string("tranquility")
	)
	return &PostCharactersAffiliationParams{
		Datasource: &datasourceDefault,
		HTTPClient: client,
	}
}

/*PostCharactersAffiliationParams contains all the parameters to send to the API endpoint
for the post characters affiliation operation typically these are written to a http.Request
*/
type PostCharactersAffiliationParams struct {

	/*Characters
	  The character IDs to fetch affiliations for. All characters must exist, or none will be returned.

	*/
	Characters []int32
	/*Datasource
	  The server name you would like data from

	*/
	Datasource *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post characters affiliation params
func (o *PostCharactersAffiliationParams) WithTimeout(timeout time.Duration) *PostCharactersAffiliationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post characters affiliation params
func (o *PostCharactersAffiliationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post characters affiliation params
func (o *PostCharactersAffiliationParams) WithContext(ctx context.Context) *PostCharactersAffiliationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post characters affiliation params
func (o *PostCharactersAffiliationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post characters affiliation params
func (o *PostCharactersAffiliationParams) WithHTTPClient(client *http.Client) *PostCharactersAffiliationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post characters affiliation params
func (o *PostCharactersAffiliationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCharacters adds the characters to the post characters affiliation params
func (o *PostCharactersAffiliationParams) WithCharacters(characters []int32) *PostCharactersAffiliationParams {
	o.SetCharacters(characters)
	return o
}

// SetCharacters adds the characters to the post characters affiliation params
func (o *PostCharactersAffiliationParams) SetCharacters(characters []int32) {
	o.Characters = characters
}

// WithDatasource adds the datasource to the post characters affiliation params
func (o *PostCharactersAffiliationParams) WithDatasource(datasource *string) *PostCharactersAffiliationParams {
	o.SetDatasource(datasource)
	return o
}

// SetDatasource adds the datasource to the post characters affiliation params
func (o *PostCharactersAffiliationParams) SetDatasource(datasource *string) {
	o.Datasource = datasource
}

// WriteToRequest writes these params to a swagger request
func (o *PostCharactersAffiliationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Characters != nil {
		if err := r.SetBodyParam(o.Characters); err != nil {
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