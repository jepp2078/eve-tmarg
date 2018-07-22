// Code generated by go-swagger; DO NOT EDIT.

package character

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/eve-tmarg/esiapi/models"
)

// GetCharactersCharacterIDMedalsReader is a Reader for the GetCharactersCharacterIDMedals structure.
type GetCharactersCharacterIDMedalsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDMedalsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDMedalsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCharactersCharacterIDMedalsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCharactersCharacterIDMedalsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCharactersCharacterIDMedalsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCharactersCharacterIDMedalsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCharactersCharacterIDMedalsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDMedalsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCharactersCharacterIDMedalsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCharactersCharacterIDMedalsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDMedalsOK creates a GetCharactersCharacterIDMedalsOK with default headers values
func NewGetCharactersCharacterIDMedalsOK() *GetCharactersCharacterIDMedalsOK {
	return &GetCharactersCharacterIDMedalsOK{}
}

/*GetCharactersCharacterIDMedalsOK handles this case with default header values.

A list of medals
*/
type GetCharactersCharacterIDMedalsOK struct {
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

	Payload []*models.GetCharactersCharacterIDMedalsOKBodyItems
}

func (o *GetCharactersCharacterIDMedalsOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/medals/][%d] getCharactersCharacterIdMedalsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDMedalsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMedalsNotModified creates a GetCharactersCharacterIDMedalsNotModified with default headers values
func NewGetCharactersCharacterIDMedalsNotModified() *GetCharactersCharacterIDMedalsNotModified {
	return &GetCharactersCharacterIDMedalsNotModified{}
}

/*GetCharactersCharacterIDMedalsNotModified handles this case with default header values.

Not modified
*/
type GetCharactersCharacterIDMedalsNotModified struct {
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

func (o *GetCharactersCharacterIDMedalsNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/medals/][%d] getCharactersCharacterIdMedalsNotModified ", 304)
}

func (o *GetCharactersCharacterIDMedalsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDMedalsBadRequest creates a GetCharactersCharacterIDMedalsBadRequest with default headers values
func NewGetCharactersCharacterIDMedalsBadRequest() *GetCharactersCharacterIDMedalsBadRequest {
	return &GetCharactersCharacterIDMedalsBadRequest{}
}

/*GetCharactersCharacterIDMedalsBadRequest handles this case with default header values.

Bad request
*/
type GetCharactersCharacterIDMedalsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDMedalsBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/medals/][%d] getCharactersCharacterIdMedalsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDMedalsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMedalsUnauthorized creates a GetCharactersCharacterIDMedalsUnauthorized with default headers values
func NewGetCharactersCharacterIDMedalsUnauthorized() *GetCharactersCharacterIDMedalsUnauthorized {
	return &GetCharactersCharacterIDMedalsUnauthorized{}
}

/*GetCharactersCharacterIDMedalsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCharactersCharacterIDMedalsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDMedalsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/medals/][%d] getCharactersCharacterIdMedalsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDMedalsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMedalsForbidden creates a GetCharactersCharacterIDMedalsForbidden with default headers values
func NewGetCharactersCharacterIDMedalsForbidden() *GetCharactersCharacterIDMedalsForbidden {
	return &GetCharactersCharacterIDMedalsForbidden{}
}

/*GetCharactersCharacterIDMedalsForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDMedalsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDMedalsForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/medals/][%d] getCharactersCharacterIdMedalsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDMedalsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMedalsEnhanceYourCalm creates a GetCharactersCharacterIDMedalsEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDMedalsEnhanceYourCalm() *GetCharactersCharacterIDMedalsEnhanceYourCalm {
	return &GetCharactersCharacterIDMedalsEnhanceYourCalm{}
}

/*GetCharactersCharacterIDMedalsEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCharactersCharacterIDMedalsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDMedalsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/medals/][%d] getCharactersCharacterIdMedalsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDMedalsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMedalsInternalServerError creates a GetCharactersCharacterIDMedalsInternalServerError with default headers values
func NewGetCharactersCharacterIDMedalsInternalServerError() *GetCharactersCharacterIDMedalsInternalServerError {
	return &GetCharactersCharacterIDMedalsInternalServerError{}
}

/*GetCharactersCharacterIDMedalsInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDMedalsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDMedalsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/medals/][%d] getCharactersCharacterIdMedalsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDMedalsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMedalsServiceUnavailable creates a GetCharactersCharacterIDMedalsServiceUnavailable with default headers values
func NewGetCharactersCharacterIDMedalsServiceUnavailable() *GetCharactersCharacterIDMedalsServiceUnavailable {
	return &GetCharactersCharacterIDMedalsServiceUnavailable{}
}

/*GetCharactersCharacterIDMedalsServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCharactersCharacterIDMedalsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDMedalsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/medals/][%d] getCharactersCharacterIdMedalsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDMedalsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMedalsGatewayTimeout creates a GetCharactersCharacterIDMedalsGatewayTimeout with default headers values
func NewGetCharactersCharacterIDMedalsGatewayTimeout() *GetCharactersCharacterIDMedalsGatewayTimeout {
	return &GetCharactersCharacterIDMedalsGatewayTimeout{}
}

/*GetCharactersCharacterIDMedalsGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDMedalsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDMedalsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/medals/][%d] getCharactersCharacterIdMedalsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDMedalsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
