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

// GetCorporationsCorporationIDStructuresReader is a Reader for the GetCorporationsCorporationIDStructures structure.
type GetCorporationsCorporationIDStructuresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDStructuresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCorporationsCorporationIDStructuresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCorporationsCorporationIDStructuresNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCorporationsCorporationIDStructuresBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCorporationsCorporationIDStructuresUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCorporationsCorporationIDStructuresForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCorporationsCorporationIDStructuresEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCorporationsCorporationIDStructuresInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCorporationsCorporationIDStructuresServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCorporationsCorporationIDStructuresGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDStructuresOK creates a GetCorporationsCorporationIDStructuresOK with default headers values
func NewGetCorporationsCorporationIDStructuresOK() *GetCorporationsCorporationIDStructuresOK {
	return &GetCorporationsCorporationIDStructuresOK{
		XPages: 1,
	}
}

/*GetCorporationsCorporationIDStructuresOK handles this case with default header values.

List of corporation structures' information
*/
type GetCorporationsCorporationIDStructuresOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*The language used in the response
	 */
	ContentLanguage string
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

	Payload []*models.GetCorporationsCorporationIDStructuresOKBodyItems
}

func (o *GetCorporationsCorporationIDStructuresOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/structures/][%d] getCorporationsCorporationIdStructuresOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDStructuresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Content-Language
	o.ContentLanguage = response.GetHeader("Content-Language")

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

// NewGetCorporationsCorporationIDStructuresNotModified creates a GetCorporationsCorporationIDStructuresNotModified with default headers values
func NewGetCorporationsCorporationIDStructuresNotModified() *GetCorporationsCorporationIDStructuresNotModified {
	return &GetCorporationsCorporationIDStructuresNotModified{}
}

/*GetCorporationsCorporationIDStructuresNotModified handles this case with default header values.

Not modified
*/
type GetCorporationsCorporationIDStructuresNotModified struct {
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

func (o *GetCorporationsCorporationIDStructuresNotModified) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/structures/][%d] getCorporationsCorporationIdStructuresNotModified ", 304)
}

func (o *GetCorporationsCorporationIDStructuresNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDStructuresBadRequest creates a GetCorporationsCorporationIDStructuresBadRequest with default headers values
func NewGetCorporationsCorporationIDStructuresBadRequest() *GetCorporationsCorporationIDStructuresBadRequest {
	return &GetCorporationsCorporationIDStructuresBadRequest{}
}

/*GetCorporationsCorporationIDStructuresBadRequest handles this case with default header values.

Bad request
*/
type GetCorporationsCorporationIDStructuresBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDStructuresBadRequest) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/structures/][%d] getCorporationsCorporationIdStructuresBadRequest  %+v", 400, o.Payload)
}

func (o *GetCorporationsCorporationIDStructuresBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDStructuresUnauthorized creates a GetCorporationsCorporationIDStructuresUnauthorized with default headers values
func NewGetCorporationsCorporationIDStructuresUnauthorized() *GetCorporationsCorporationIDStructuresUnauthorized {
	return &GetCorporationsCorporationIDStructuresUnauthorized{}
}

/*GetCorporationsCorporationIDStructuresUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDStructuresUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDStructuresUnauthorized) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/structures/][%d] getCorporationsCorporationIdStructuresUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCorporationsCorporationIDStructuresUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDStructuresForbidden creates a GetCorporationsCorporationIDStructuresForbidden with default headers values
func NewGetCorporationsCorporationIDStructuresForbidden() *GetCorporationsCorporationIDStructuresForbidden {
	return &GetCorporationsCorporationIDStructuresForbidden{}
}

/*GetCorporationsCorporationIDStructuresForbidden handles this case with default header values.

Forbidden
*/
type GetCorporationsCorporationIDStructuresForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDStructuresForbidden) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/structures/][%d] getCorporationsCorporationIdStructuresForbidden  %+v", 403, o.Payload)
}

func (o *GetCorporationsCorporationIDStructuresForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDStructuresEnhanceYourCalm creates a GetCorporationsCorporationIDStructuresEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDStructuresEnhanceYourCalm() *GetCorporationsCorporationIDStructuresEnhanceYourCalm {
	return &GetCorporationsCorporationIDStructuresEnhanceYourCalm{}
}

/*GetCorporationsCorporationIDStructuresEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCorporationsCorporationIDStructuresEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDStructuresEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/structures/][%d] getCorporationsCorporationIdStructuresEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCorporationsCorporationIDStructuresEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDStructuresInternalServerError creates a GetCorporationsCorporationIDStructuresInternalServerError with default headers values
func NewGetCorporationsCorporationIDStructuresInternalServerError() *GetCorporationsCorporationIDStructuresInternalServerError {
	return &GetCorporationsCorporationIDStructuresInternalServerError{}
}

/*GetCorporationsCorporationIDStructuresInternalServerError handles this case with default header values.

Internal server error
*/
type GetCorporationsCorporationIDStructuresInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDStructuresInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/structures/][%d] getCorporationsCorporationIdStructuresInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDStructuresInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDStructuresServiceUnavailable creates a GetCorporationsCorporationIDStructuresServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDStructuresServiceUnavailable() *GetCorporationsCorporationIDStructuresServiceUnavailable {
	return &GetCorporationsCorporationIDStructuresServiceUnavailable{}
}

/*GetCorporationsCorporationIDStructuresServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDStructuresServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDStructuresServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/structures/][%d] getCorporationsCorporationIdStructuresServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCorporationsCorporationIDStructuresServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDStructuresGatewayTimeout creates a GetCorporationsCorporationIDStructuresGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDStructuresGatewayTimeout() *GetCorporationsCorporationIDStructuresGatewayTimeout {
	return &GetCorporationsCorporationIDStructuresGatewayTimeout{}
}

/*GetCorporationsCorporationIDStructuresGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDStructuresGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDStructuresGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/structures/][%d] getCorporationsCorporationIdStructuresGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCorporationsCorporationIDStructuresGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
