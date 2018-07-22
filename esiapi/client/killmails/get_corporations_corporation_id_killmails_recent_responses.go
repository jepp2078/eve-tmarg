// Code generated by go-swagger; DO NOT EDIT.

package killmails

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/eve-tmarg/esiapi/models"
)

// GetCorporationsCorporationIDKillmailsRecentReader is a Reader for the GetCorporationsCorporationIDKillmailsRecent structure.
type GetCorporationsCorporationIDKillmailsRecentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDKillmailsRecentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCorporationsCorporationIDKillmailsRecentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCorporationsCorporationIDKillmailsRecentNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCorporationsCorporationIDKillmailsRecentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCorporationsCorporationIDKillmailsRecentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCorporationsCorporationIDKillmailsRecentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCorporationsCorporationIDKillmailsRecentEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCorporationsCorporationIDKillmailsRecentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCorporationsCorporationIDKillmailsRecentServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCorporationsCorporationIDKillmailsRecentGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDKillmailsRecentOK creates a GetCorporationsCorporationIDKillmailsRecentOK with default headers values
func NewGetCorporationsCorporationIDKillmailsRecentOK() *GetCorporationsCorporationIDKillmailsRecentOK {
	return &GetCorporationsCorporationIDKillmailsRecentOK{
		XPages: 1,
	}
}

/*GetCorporationsCorporationIDKillmailsRecentOK handles this case with default header values.

A list of killmail IDs and hashes
*/
type GetCorporationsCorporationIDKillmailsRecentOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7232 compliant entity tag
	 */
	ETag string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string
	/*Maximum page number
	 */
	XPages int32

	Payload []*models.GetCorporationsCorporationIDKillmailsRecentOKBodyItems
}

func (o *GetCorporationsCorporationIDKillmailsRecentOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/killmails/recent/][%d] getCorporationsCorporationIdKillmailsRecentOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDKillmailsRecentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	// response header X-Pages
	xPages, err := swag.ConvertInt32(response.GetHeader("X-Pages"))
	if err != nil {
		return errors.InvalidType("X-Pages", "header", "int32", response.GetHeader("X-Pages"))
	}
	o.XPages = xPages

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDKillmailsRecentNotModified creates a GetCorporationsCorporationIDKillmailsRecentNotModified with default headers values
func NewGetCorporationsCorporationIDKillmailsRecentNotModified() *GetCorporationsCorporationIDKillmailsRecentNotModified {
	return &GetCorporationsCorporationIDKillmailsRecentNotModified{}
}

/*GetCorporationsCorporationIDKillmailsRecentNotModified handles this case with default header values.

Not modified
*/
type GetCorporationsCorporationIDKillmailsRecentNotModified struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*RFC7232 compliant entity tag
	 */
	ETag string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string
}

func (o *GetCorporationsCorporationIDKillmailsRecentNotModified) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/killmails/recent/][%d] getCorporationsCorporationIdKillmailsRecentNotModified ", 304)
}

func (o *GetCorporationsCorporationIDKillmailsRecentNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	return nil
}

// NewGetCorporationsCorporationIDKillmailsRecentBadRequest creates a GetCorporationsCorporationIDKillmailsRecentBadRequest with default headers values
func NewGetCorporationsCorporationIDKillmailsRecentBadRequest() *GetCorporationsCorporationIDKillmailsRecentBadRequest {
	return &GetCorporationsCorporationIDKillmailsRecentBadRequest{}
}

/*GetCorporationsCorporationIDKillmailsRecentBadRequest handles this case with default header values.

Bad request
*/
type GetCorporationsCorporationIDKillmailsRecentBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDKillmailsRecentBadRequest) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/killmails/recent/][%d] getCorporationsCorporationIdKillmailsRecentBadRequest  %+v", 400, o.Payload)
}

func (o *GetCorporationsCorporationIDKillmailsRecentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDKillmailsRecentUnauthorized creates a GetCorporationsCorporationIDKillmailsRecentUnauthorized with default headers values
func NewGetCorporationsCorporationIDKillmailsRecentUnauthorized() *GetCorporationsCorporationIDKillmailsRecentUnauthorized {
	return &GetCorporationsCorporationIDKillmailsRecentUnauthorized{}
}

/*GetCorporationsCorporationIDKillmailsRecentUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDKillmailsRecentUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDKillmailsRecentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/killmails/recent/][%d] getCorporationsCorporationIdKillmailsRecentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCorporationsCorporationIDKillmailsRecentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDKillmailsRecentForbidden creates a GetCorporationsCorporationIDKillmailsRecentForbidden with default headers values
func NewGetCorporationsCorporationIDKillmailsRecentForbidden() *GetCorporationsCorporationIDKillmailsRecentForbidden {
	return &GetCorporationsCorporationIDKillmailsRecentForbidden{}
}

/*GetCorporationsCorporationIDKillmailsRecentForbidden handles this case with default header values.

Forbidden
*/
type GetCorporationsCorporationIDKillmailsRecentForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDKillmailsRecentForbidden) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/killmails/recent/][%d] getCorporationsCorporationIdKillmailsRecentForbidden  %+v", 403, o.Payload)
}

func (o *GetCorporationsCorporationIDKillmailsRecentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDKillmailsRecentEnhanceYourCalm creates a GetCorporationsCorporationIDKillmailsRecentEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDKillmailsRecentEnhanceYourCalm() *GetCorporationsCorporationIDKillmailsRecentEnhanceYourCalm {
	return &GetCorporationsCorporationIDKillmailsRecentEnhanceYourCalm{}
}

/*GetCorporationsCorporationIDKillmailsRecentEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCorporationsCorporationIDKillmailsRecentEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDKillmailsRecentEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/killmails/recent/][%d] getCorporationsCorporationIdKillmailsRecentEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCorporationsCorporationIDKillmailsRecentEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDKillmailsRecentInternalServerError creates a GetCorporationsCorporationIDKillmailsRecentInternalServerError with default headers values
func NewGetCorporationsCorporationIDKillmailsRecentInternalServerError() *GetCorporationsCorporationIDKillmailsRecentInternalServerError {
	return &GetCorporationsCorporationIDKillmailsRecentInternalServerError{}
}

/*GetCorporationsCorporationIDKillmailsRecentInternalServerError handles this case with default header values.

Internal server error
*/
type GetCorporationsCorporationIDKillmailsRecentInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDKillmailsRecentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/killmails/recent/][%d] getCorporationsCorporationIdKillmailsRecentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDKillmailsRecentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDKillmailsRecentServiceUnavailable creates a GetCorporationsCorporationIDKillmailsRecentServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDKillmailsRecentServiceUnavailable() *GetCorporationsCorporationIDKillmailsRecentServiceUnavailable {
	return &GetCorporationsCorporationIDKillmailsRecentServiceUnavailable{}
}

/*GetCorporationsCorporationIDKillmailsRecentServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDKillmailsRecentServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDKillmailsRecentServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/killmails/recent/][%d] getCorporationsCorporationIdKillmailsRecentServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCorporationsCorporationIDKillmailsRecentServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDKillmailsRecentGatewayTimeout creates a GetCorporationsCorporationIDKillmailsRecentGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDKillmailsRecentGatewayTimeout() *GetCorporationsCorporationIDKillmailsRecentGatewayTimeout {
	return &GetCorporationsCorporationIDKillmailsRecentGatewayTimeout{}
}

/*GetCorporationsCorporationIDKillmailsRecentGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDKillmailsRecentGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDKillmailsRecentGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/killmails/recent/][%d] getCorporationsCorporationIdKillmailsRecentGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCorporationsCorporationIDKillmailsRecentGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}