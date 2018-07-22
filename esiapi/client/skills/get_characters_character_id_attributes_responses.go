// Code generated by go-swagger; DO NOT EDIT.

package skills

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/eve-tmarg/esiapi/models"
)

// GetCharactersCharacterIDAttributesReader is a Reader for the GetCharactersCharacterIDAttributes structure.
type GetCharactersCharacterIDAttributesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDAttributesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDAttributesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCharactersCharacterIDAttributesNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCharactersCharacterIDAttributesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCharactersCharacterIDAttributesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCharactersCharacterIDAttributesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCharactersCharacterIDAttributesEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDAttributesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCharactersCharacterIDAttributesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCharactersCharacterIDAttributesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDAttributesOK creates a GetCharactersCharacterIDAttributesOK with default headers values
func NewGetCharactersCharacterIDAttributesOK() *GetCharactersCharacterIDAttributesOK {
	return &GetCharactersCharacterIDAttributesOK{}
}

/*GetCharactersCharacterIDAttributesOK handles this case with default header values.

Attributes of a character
*/
type GetCharactersCharacterIDAttributesOK struct {
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

	Payload *models.GetCharactersCharacterIDAttributesOKBody
}

func (o *GetCharactersCharacterIDAttributesOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/attributes/][%d] getCharactersCharacterIdAttributesOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDAttributesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetCharactersCharacterIDAttributesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDAttributesNotModified creates a GetCharactersCharacterIDAttributesNotModified with default headers values
func NewGetCharactersCharacterIDAttributesNotModified() *GetCharactersCharacterIDAttributesNotModified {
	return &GetCharactersCharacterIDAttributesNotModified{}
}

/*GetCharactersCharacterIDAttributesNotModified handles this case with default header values.

Not modified
*/
type GetCharactersCharacterIDAttributesNotModified struct {
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

func (o *GetCharactersCharacterIDAttributesNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/attributes/][%d] getCharactersCharacterIdAttributesNotModified ", 304)
}

func (o *GetCharactersCharacterIDAttributesNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDAttributesBadRequest creates a GetCharactersCharacterIDAttributesBadRequest with default headers values
func NewGetCharactersCharacterIDAttributesBadRequest() *GetCharactersCharacterIDAttributesBadRequest {
	return &GetCharactersCharacterIDAttributesBadRequest{}
}

/*GetCharactersCharacterIDAttributesBadRequest handles this case with default header values.

Bad request
*/
type GetCharactersCharacterIDAttributesBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDAttributesBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/attributes/][%d] getCharactersCharacterIdAttributesBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDAttributesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDAttributesUnauthorized creates a GetCharactersCharacterIDAttributesUnauthorized with default headers values
func NewGetCharactersCharacterIDAttributesUnauthorized() *GetCharactersCharacterIDAttributesUnauthorized {
	return &GetCharactersCharacterIDAttributesUnauthorized{}
}

/*GetCharactersCharacterIDAttributesUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCharactersCharacterIDAttributesUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDAttributesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/attributes/][%d] getCharactersCharacterIdAttributesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDAttributesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDAttributesForbidden creates a GetCharactersCharacterIDAttributesForbidden with default headers values
func NewGetCharactersCharacterIDAttributesForbidden() *GetCharactersCharacterIDAttributesForbidden {
	return &GetCharactersCharacterIDAttributesForbidden{}
}

/*GetCharactersCharacterIDAttributesForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDAttributesForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDAttributesForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/attributes/][%d] getCharactersCharacterIdAttributesForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDAttributesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDAttributesEnhanceYourCalm creates a GetCharactersCharacterIDAttributesEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDAttributesEnhanceYourCalm() *GetCharactersCharacterIDAttributesEnhanceYourCalm {
	return &GetCharactersCharacterIDAttributesEnhanceYourCalm{}
}

/*GetCharactersCharacterIDAttributesEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCharactersCharacterIDAttributesEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDAttributesEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/attributes/][%d] getCharactersCharacterIdAttributesEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDAttributesEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDAttributesInternalServerError creates a GetCharactersCharacterIDAttributesInternalServerError with default headers values
func NewGetCharactersCharacterIDAttributesInternalServerError() *GetCharactersCharacterIDAttributesInternalServerError {
	return &GetCharactersCharacterIDAttributesInternalServerError{}
}

/*GetCharactersCharacterIDAttributesInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDAttributesInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDAttributesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/attributes/][%d] getCharactersCharacterIdAttributesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDAttributesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDAttributesServiceUnavailable creates a GetCharactersCharacterIDAttributesServiceUnavailable with default headers values
func NewGetCharactersCharacterIDAttributesServiceUnavailable() *GetCharactersCharacterIDAttributesServiceUnavailable {
	return &GetCharactersCharacterIDAttributesServiceUnavailable{}
}

/*GetCharactersCharacterIDAttributesServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCharactersCharacterIDAttributesServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDAttributesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/attributes/][%d] getCharactersCharacterIdAttributesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDAttributesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDAttributesGatewayTimeout creates a GetCharactersCharacterIDAttributesGatewayTimeout with default headers values
func NewGetCharactersCharacterIDAttributesGatewayTimeout() *GetCharactersCharacterIDAttributesGatewayTimeout {
	return &GetCharactersCharacterIDAttributesGatewayTimeout{}
}

/*GetCharactersCharacterIDAttributesGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDAttributesGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDAttributesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/attributes/][%d] getCharactersCharacterIdAttributesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDAttributesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
