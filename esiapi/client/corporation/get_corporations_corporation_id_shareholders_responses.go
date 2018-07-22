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

// GetCorporationsCorporationIDShareholdersReader is a Reader for the GetCorporationsCorporationIDShareholders structure.
type GetCorporationsCorporationIDShareholdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDShareholdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCorporationsCorporationIDShareholdersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCorporationsCorporationIDShareholdersNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCorporationsCorporationIDShareholdersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCorporationsCorporationIDShareholdersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCorporationsCorporationIDShareholdersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCorporationsCorporationIDShareholdersEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCorporationsCorporationIDShareholdersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCorporationsCorporationIDShareholdersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCorporationsCorporationIDShareholdersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDShareholdersOK creates a GetCorporationsCorporationIDShareholdersOK with default headers values
func NewGetCorporationsCorporationIDShareholdersOK() *GetCorporationsCorporationIDShareholdersOK {
	return &GetCorporationsCorporationIDShareholdersOK{
		XPages: 1,
	}
}

/*GetCorporationsCorporationIDShareholdersOK handles this case with default header values.

List of shareholders
*/
type GetCorporationsCorporationIDShareholdersOK struct {
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

	Payload []*models.GetCorporationsCorporationIDShareholdersOKBodyItems
}

func (o *GetCorporationsCorporationIDShareholdersOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/shareholders/][%d] getCorporationsCorporationIdShareholdersOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDShareholdersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDShareholdersNotModified creates a GetCorporationsCorporationIDShareholdersNotModified with default headers values
func NewGetCorporationsCorporationIDShareholdersNotModified() *GetCorporationsCorporationIDShareholdersNotModified {
	return &GetCorporationsCorporationIDShareholdersNotModified{}
}

/*GetCorporationsCorporationIDShareholdersNotModified handles this case with default header values.

Not modified
*/
type GetCorporationsCorporationIDShareholdersNotModified struct {
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

func (o *GetCorporationsCorporationIDShareholdersNotModified) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/shareholders/][%d] getCorporationsCorporationIdShareholdersNotModified ", 304)
}

func (o *GetCorporationsCorporationIDShareholdersNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDShareholdersBadRequest creates a GetCorporationsCorporationIDShareholdersBadRequest with default headers values
func NewGetCorporationsCorporationIDShareholdersBadRequest() *GetCorporationsCorporationIDShareholdersBadRequest {
	return &GetCorporationsCorporationIDShareholdersBadRequest{}
}

/*GetCorporationsCorporationIDShareholdersBadRequest handles this case with default header values.

Bad request
*/
type GetCorporationsCorporationIDShareholdersBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDShareholdersBadRequest) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/shareholders/][%d] getCorporationsCorporationIdShareholdersBadRequest  %+v", 400, o.Payload)
}

func (o *GetCorporationsCorporationIDShareholdersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDShareholdersUnauthorized creates a GetCorporationsCorporationIDShareholdersUnauthorized with default headers values
func NewGetCorporationsCorporationIDShareholdersUnauthorized() *GetCorporationsCorporationIDShareholdersUnauthorized {
	return &GetCorporationsCorporationIDShareholdersUnauthorized{}
}

/*GetCorporationsCorporationIDShareholdersUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDShareholdersUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDShareholdersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/shareholders/][%d] getCorporationsCorporationIdShareholdersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCorporationsCorporationIDShareholdersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDShareholdersForbidden creates a GetCorporationsCorporationIDShareholdersForbidden with default headers values
func NewGetCorporationsCorporationIDShareholdersForbidden() *GetCorporationsCorporationIDShareholdersForbidden {
	return &GetCorporationsCorporationIDShareholdersForbidden{}
}

/*GetCorporationsCorporationIDShareholdersForbidden handles this case with default header values.

Forbidden
*/
type GetCorporationsCorporationIDShareholdersForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDShareholdersForbidden) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/shareholders/][%d] getCorporationsCorporationIdShareholdersForbidden  %+v", 403, o.Payload)
}

func (o *GetCorporationsCorporationIDShareholdersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDShareholdersEnhanceYourCalm creates a GetCorporationsCorporationIDShareholdersEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDShareholdersEnhanceYourCalm() *GetCorporationsCorporationIDShareholdersEnhanceYourCalm {
	return &GetCorporationsCorporationIDShareholdersEnhanceYourCalm{}
}

/*GetCorporationsCorporationIDShareholdersEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCorporationsCorporationIDShareholdersEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDShareholdersEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/shareholders/][%d] getCorporationsCorporationIdShareholdersEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCorporationsCorporationIDShareholdersEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDShareholdersInternalServerError creates a GetCorporationsCorporationIDShareholdersInternalServerError with default headers values
func NewGetCorporationsCorporationIDShareholdersInternalServerError() *GetCorporationsCorporationIDShareholdersInternalServerError {
	return &GetCorporationsCorporationIDShareholdersInternalServerError{}
}

/*GetCorporationsCorporationIDShareholdersInternalServerError handles this case with default header values.

Internal server error
*/
type GetCorporationsCorporationIDShareholdersInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDShareholdersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/shareholders/][%d] getCorporationsCorporationIdShareholdersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDShareholdersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDShareholdersServiceUnavailable creates a GetCorporationsCorporationIDShareholdersServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDShareholdersServiceUnavailable() *GetCorporationsCorporationIDShareholdersServiceUnavailable {
	return &GetCorporationsCorporationIDShareholdersServiceUnavailable{}
}

/*GetCorporationsCorporationIDShareholdersServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDShareholdersServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDShareholdersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/shareholders/][%d] getCorporationsCorporationIdShareholdersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCorporationsCorporationIDShareholdersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDShareholdersGatewayTimeout creates a GetCorporationsCorporationIDShareholdersGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDShareholdersGatewayTimeout() *GetCorporationsCorporationIDShareholdersGatewayTimeout {
	return &GetCorporationsCorporationIDShareholdersGatewayTimeout{}
}

/*GetCorporationsCorporationIDShareholdersGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDShareholdersGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDShareholdersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/shareholders/][%d] getCorporationsCorporationIdShareholdersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCorporationsCorporationIDShareholdersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
