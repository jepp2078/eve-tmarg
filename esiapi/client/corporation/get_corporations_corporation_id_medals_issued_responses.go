// Code generated by go-swagger; DO NOT EDIT.

package corporation

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

// GetCorporationsCorporationIDMedalsIssuedReader is a Reader for the GetCorporationsCorporationIDMedalsIssued structure.
type GetCorporationsCorporationIDMedalsIssuedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDMedalsIssuedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCorporationsCorporationIDMedalsIssuedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCorporationsCorporationIDMedalsIssuedNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCorporationsCorporationIDMedalsIssuedBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCorporationsCorporationIDMedalsIssuedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCorporationsCorporationIDMedalsIssuedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCorporationsCorporationIDMedalsIssuedEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCorporationsCorporationIDMedalsIssuedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCorporationsCorporationIDMedalsIssuedServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCorporationsCorporationIDMedalsIssuedGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDMedalsIssuedOK creates a GetCorporationsCorporationIDMedalsIssuedOK with default headers values
func NewGetCorporationsCorporationIDMedalsIssuedOK() *GetCorporationsCorporationIDMedalsIssuedOK {
	return &GetCorporationsCorporationIDMedalsIssuedOK{
		XPages: 1,
	}
}

/*GetCorporationsCorporationIDMedalsIssuedOK handles this case with default header values.

A list of issued medals
*/
type GetCorporationsCorporationIDMedalsIssuedOK struct {
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

	Payload []*models.GetCorporationsCorporationIDMedalsIssuedOKBodyItems
}

func (o *GetCorporationsCorporationIDMedalsIssuedOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/issued/][%d] getCorporationsCorporationIdMedalsIssuedOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDMedalsIssuedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDMedalsIssuedNotModified creates a GetCorporationsCorporationIDMedalsIssuedNotModified with default headers values
func NewGetCorporationsCorporationIDMedalsIssuedNotModified() *GetCorporationsCorporationIDMedalsIssuedNotModified {
	return &GetCorporationsCorporationIDMedalsIssuedNotModified{}
}

/*GetCorporationsCorporationIDMedalsIssuedNotModified handles this case with default header values.

Not modified
*/
type GetCorporationsCorporationIDMedalsIssuedNotModified struct {
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

func (o *GetCorporationsCorporationIDMedalsIssuedNotModified) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/issued/][%d] getCorporationsCorporationIdMedalsIssuedNotModified ", 304)
}

func (o *GetCorporationsCorporationIDMedalsIssuedNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDMedalsIssuedBadRequest creates a GetCorporationsCorporationIDMedalsIssuedBadRequest with default headers values
func NewGetCorporationsCorporationIDMedalsIssuedBadRequest() *GetCorporationsCorporationIDMedalsIssuedBadRequest {
	return &GetCorporationsCorporationIDMedalsIssuedBadRequest{}
}

/*GetCorporationsCorporationIDMedalsIssuedBadRequest handles this case with default header values.

Bad request
*/
type GetCorporationsCorporationIDMedalsIssuedBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDMedalsIssuedBadRequest) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/issued/][%d] getCorporationsCorporationIdMedalsIssuedBadRequest  %+v", 400, o.Payload)
}

func (o *GetCorporationsCorporationIDMedalsIssuedBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsIssuedUnauthorized creates a GetCorporationsCorporationIDMedalsIssuedUnauthorized with default headers values
func NewGetCorporationsCorporationIDMedalsIssuedUnauthorized() *GetCorporationsCorporationIDMedalsIssuedUnauthorized {
	return &GetCorporationsCorporationIDMedalsIssuedUnauthorized{}
}

/*GetCorporationsCorporationIDMedalsIssuedUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDMedalsIssuedUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDMedalsIssuedUnauthorized) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/issued/][%d] getCorporationsCorporationIdMedalsIssuedUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCorporationsCorporationIDMedalsIssuedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsIssuedForbidden creates a GetCorporationsCorporationIDMedalsIssuedForbidden with default headers values
func NewGetCorporationsCorporationIDMedalsIssuedForbidden() *GetCorporationsCorporationIDMedalsIssuedForbidden {
	return &GetCorporationsCorporationIDMedalsIssuedForbidden{}
}

/*GetCorporationsCorporationIDMedalsIssuedForbidden handles this case with default header values.

Forbidden
*/
type GetCorporationsCorporationIDMedalsIssuedForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDMedalsIssuedForbidden) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/issued/][%d] getCorporationsCorporationIdMedalsIssuedForbidden  %+v", 403, o.Payload)
}

func (o *GetCorporationsCorporationIDMedalsIssuedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsIssuedEnhanceYourCalm creates a GetCorporationsCorporationIDMedalsIssuedEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDMedalsIssuedEnhanceYourCalm() *GetCorporationsCorporationIDMedalsIssuedEnhanceYourCalm {
	return &GetCorporationsCorporationIDMedalsIssuedEnhanceYourCalm{}
}

/*GetCorporationsCorporationIDMedalsIssuedEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCorporationsCorporationIDMedalsIssuedEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDMedalsIssuedEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/issued/][%d] getCorporationsCorporationIdMedalsIssuedEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCorporationsCorporationIDMedalsIssuedEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsIssuedInternalServerError creates a GetCorporationsCorporationIDMedalsIssuedInternalServerError with default headers values
func NewGetCorporationsCorporationIDMedalsIssuedInternalServerError() *GetCorporationsCorporationIDMedalsIssuedInternalServerError {
	return &GetCorporationsCorporationIDMedalsIssuedInternalServerError{}
}

/*GetCorporationsCorporationIDMedalsIssuedInternalServerError handles this case with default header values.

Internal server error
*/
type GetCorporationsCorporationIDMedalsIssuedInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDMedalsIssuedInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/issued/][%d] getCorporationsCorporationIdMedalsIssuedInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDMedalsIssuedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsIssuedServiceUnavailable creates a GetCorporationsCorporationIDMedalsIssuedServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDMedalsIssuedServiceUnavailable() *GetCorporationsCorporationIDMedalsIssuedServiceUnavailable {
	return &GetCorporationsCorporationIDMedalsIssuedServiceUnavailable{}
}

/*GetCorporationsCorporationIDMedalsIssuedServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDMedalsIssuedServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDMedalsIssuedServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/issued/][%d] getCorporationsCorporationIdMedalsIssuedServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCorporationsCorporationIDMedalsIssuedServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMedalsIssuedGatewayTimeout creates a GetCorporationsCorporationIDMedalsIssuedGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDMedalsIssuedGatewayTimeout() *GetCorporationsCorporationIDMedalsIssuedGatewayTimeout {
	return &GetCorporationsCorporationIDMedalsIssuedGatewayTimeout{}
}

/*GetCorporationsCorporationIDMedalsIssuedGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDMedalsIssuedGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDMedalsIssuedGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/medals/issued/][%d] getCorporationsCorporationIdMedalsIssuedGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCorporationsCorporationIDMedalsIssuedGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
