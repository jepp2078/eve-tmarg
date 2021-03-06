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

// GetCharactersCharacterIDStandingsReader is a Reader for the GetCharactersCharacterIDStandings structure.
type GetCharactersCharacterIDStandingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDStandingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDStandingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCharactersCharacterIDStandingsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCharactersCharacterIDStandingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCharactersCharacterIDStandingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCharactersCharacterIDStandingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCharactersCharacterIDStandingsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDStandingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCharactersCharacterIDStandingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCharactersCharacterIDStandingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDStandingsOK creates a GetCharactersCharacterIDStandingsOK with default headers values
func NewGetCharactersCharacterIDStandingsOK() *GetCharactersCharacterIDStandingsOK {
	return &GetCharactersCharacterIDStandingsOK{}
}

/*GetCharactersCharacterIDStandingsOK handles this case with default header values.

A list of standings
*/
type GetCharactersCharacterIDStandingsOK struct {
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

	Payload []*models.GetCharactersCharacterIDStandingsOKBodyItems
}

func (o *GetCharactersCharacterIDStandingsOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/standings/][%d] getCharactersCharacterIdStandingsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDStandingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDStandingsNotModified creates a GetCharactersCharacterIDStandingsNotModified with default headers values
func NewGetCharactersCharacterIDStandingsNotModified() *GetCharactersCharacterIDStandingsNotModified {
	return &GetCharactersCharacterIDStandingsNotModified{}
}

/*GetCharactersCharacterIDStandingsNotModified handles this case with default header values.

Not modified
*/
type GetCharactersCharacterIDStandingsNotModified struct {
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

func (o *GetCharactersCharacterIDStandingsNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/standings/][%d] getCharactersCharacterIdStandingsNotModified ", 304)
}

func (o *GetCharactersCharacterIDStandingsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDStandingsBadRequest creates a GetCharactersCharacterIDStandingsBadRequest with default headers values
func NewGetCharactersCharacterIDStandingsBadRequest() *GetCharactersCharacterIDStandingsBadRequest {
	return &GetCharactersCharacterIDStandingsBadRequest{}
}

/*GetCharactersCharacterIDStandingsBadRequest handles this case with default header values.

Bad request
*/
type GetCharactersCharacterIDStandingsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDStandingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/standings/][%d] getCharactersCharacterIdStandingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDStandingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDStandingsUnauthorized creates a GetCharactersCharacterIDStandingsUnauthorized with default headers values
func NewGetCharactersCharacterIDStandingsUnauthorized() *GetCharactersCharacterIDStandingsUnauthorized {
	return &GetCharactersCharacterIDStandingsUnauthorized{}
}

/*GetCharactersCharacterIDStandingsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCharactersCharacterIDStandingsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDStandingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/standings/][%d] getCharactersCharacterIdStandingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDStandingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDStandingsForbidden creates a GetCharactersCharacterIDStandingsForbidden with default headers values
func NewGetCharactersCharacterIDStandingsForbidden() *GetCharactersCharacterIDStandingsForbidden {
	return &GetCharactersCharacterIDStandingsForbidden{}
}

/*GetCharactersCharacterIDStandingsForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDStandingsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDStandingsForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/standings/][%d] getCharactersCharacterIdStandingsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDStandingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDStandingsEnhanceYourCalm creates a GetCharactersCharacterIDStandingsEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDStandingsEnhanceYourCalm() *GetCharactersCharacterIDStandingsEnhanceYourCalm {
	return &GetCharactersCharacterIDStandingsEnhanceYourCalm{}
}

/*GetCharactersCharacterIDStandingsEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCharactersCharacterIDStandingsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDStandingsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/standings/][%d] getCharactersCharacterIdStandingsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDStandingsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDStandingsInternalServerError creates a GetCharactersCharacterIDStandingsInternalServerError with default headers values
func NewGetCharactersCharacterIDStandingsInternalServerError() *GetCharactersCharacterIDStandingsInternalServerError {
	return &GetCharactersCharacterIDStandingsInternalServerError{}
}

/*GetCharactersCharacterIDStandingsInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDStandingsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDStandingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/standings/][%d] getCharactersCharacterIdStandingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDStandingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDStandingsServiceUnavailable creates a GetCharactersCharacterIDStandingsServiceUnavailable with default headers values
func NewGetCharactersCharacterIDStandingsServiceUnavailable() *GetCharactersCharacterIDStandingsServiceUnavailable {
	return &GetCharactersCharacterIDStandingsServiceUnavailable{}
}

/*GetCharactersCharacterIDStandingsServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCharactersCharacterIDStandingsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDStandingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/standings/][%d] getCharactersCharacterIdStandingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDStandingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDStandingsGatewayTimeout creates a GetCharactersCharacterIDStandingsGatewayTimeout with default headers values
func NewGetCharactersCharacterIDStandingsGatewayTimeout() *GetCharactersCharacterIDStandingsGatewayTimeout {
	return &GetCharactersCharacterIDStandingsGatewayTimeout{}
}

/*GetCharactersCharacterIDStandingsGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDStandingsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDStandingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/standings/][%d] getCharactersCharacterIdStandingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDStandingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
