// Code generated by go-swagger; DO NOT EDIT.

package opportunities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/eve-tmarg/esiapi/models"
)

// GetCharactersCharacterIDOpportunitiesReader is a Reader for the GetCharactersCharacterIDOpportunities structure.
type GetCharactersCharacterIDOpportunitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDOpportunitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDOpportunitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCharactersCharacterIDOpportunitiesNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCharactersCharacterIDOpportunitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCharactersCharacterIDOpportunitiesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCharactersCharacterIDOpportunitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCharactersCharacterIDOpportunitiesEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDOpportunitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCharactersCharacterIDOpportunitiesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCharactersCharacterIDOpportunitiesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDOpportunitiesOK creates a GetCharactersCharacterIDOpportunitiesOK with default headers values
func NewGetCharactersCharacterIDOpportunitiesOK() *GetCharactersCharacterIDOpportunitiesOK {
	return &GetCharactersCharacterIDOpportunitiesOK{}
}

/*GetCharactersCharacterIDOpportunitiesOK handles this case with default header values.

A list of opportunities task ids
*/
type GetCharactersCharacterIDOpportunitiesOK struct {
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

	Payload []*models.GetCharactersCharacterIDOpportunitiesOKBodyItems
}

func (o *GetCharactersCharacterIDOpportunitiesOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/opportunities/][%d] getCharactersCharacterIdOpportunitiesOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDOpportunitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDOpportunitiesNotModified creates a GetCharactersCharacterIDOpportunitiesNotModified with default headers values
func NewGetCharactersCharacterIDOpportunitiesNotModified() *GetCharactersCharacterIDOpportunitiesNotModified {
	return &GetCharactersCharacterIDOpportunitiesNotModified{}
}

/*GetCharactersCharacterIDOpportunitiesNotModified handles this case with default header values.

Not modified
*/
type GetCharactersCharacterIDOpportunitiesNotModified struct {
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

func (o *GetCharactersCharacterIDOpportunitiesNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/opportunities/][%d] getCharactersCharacterIdOpportunitiesNotModified ", 304)
}

func (o *GetCharactersCharacterIDOpportunitiesNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDOpportunitiesBadRequest creates a GetCharactersCharacterIDOpportunitiesBadRequest with default headers values
func NewGetCharactersCharacterIDOpportunitiesBadRequest() *GetCharactersCharacterIDOpportunitiesBadRequest {
	return &GetCharactersCharacterIDOpportunitiesBadRequest{}
}

/*GetCharactersCharacterIDOpportunitiesBadRequest handles this case with default header values.

Bad request
*/
type GetCharactersCharacterIDOpportunitiesBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDOpportunitiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/opportunities/][%d] getCharactersCharacterIdOpportunitiesBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDOpportunitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOpportunitiesUnauthorized creates a GetCharactersCharacterIDOpportunitiesUnauthorized with default headers values
func NewGetCharactersCharacterIDOpportunitiesUnauthorized() *GetCharactersCharacterIDOpportunitiesUnauthorized {
	return &GetCharactersCharacterIDOpportunitiesUnauthorized{}
}

/*GetCharactersCharacterIDOpportunitiesUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCharactersCharacterIDOpportunitiesUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDOpportunitiesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/opportunities/][%d] getCharactersCharacterIdOpportunitiesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDOpportunitiesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOpportunitiesForbidden creates a GetCharactersCharacterIDOpportunitiesForbidden with default headers values
func NewGetCharactersCharacterIDOpportunitiesForbidden() *GetCharactersCharacterIDOpportunitiesForbidden {
	return &GetCharactersCharacterIDOpportunitiesForbidden{}
}

/*GetCharactersCharacterIDOpportunitiesForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDOpportunitiesForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDOpportunitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/opportunities/][%d] getCharactersCharacterIdOpportunitiesForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDOpportunitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOpportunitiesEnhanceYourCalm creates a GetCharactersCharacterIDOpportunitiesEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDOpportunitiesEnhanceYourCalm() *GetCharactersCharacterIDOpportunitiesEnhanceYourCalm {
	return &GetCharactersCharacterIDOpportunitiesEnhanceYourCalm{}
}

/*GetCharactersCharacterIDOpportunitiesEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCharactersCharacterIDOpportunitiesEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDOpportunitiesEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/opportunities/][%d] getCharactersCharacterIdOpportunitiesEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDOpportunitiesEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOpportunitiesInternalServerError creates a GetCharactersCharacterIDOpportunitiesInternalServerError with default headers values
func NewGetCharactersCharacterIDOpportunitiesInternalServerError() *GetCharactersCharacterIDOpportunitiesInternalServerError {
	return &GetCharactersCharacterIDOpportunitiesInternalServerError{}
}

/*GetCharactersCharacterIDOpportunitiesInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDOpportunitiesInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDOpportunitiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/opportunities/][%d] getCharactersCharacterIdOpportunitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDOpportunitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOpportunitiesServiceUnavailable creates a GetCharactersCharacterIDOpportunitiesServiceUnavailable with default headers values
func NewGetCharactersCharacterIDOpportunitiesServiceUnavailable() *GetCharactersCharacterIDOpportunitiesServiceUnavailable {
	return &GetCharactersCharacterIDOpportunitiesServiceUnavailable{}
}

/*GetCharactersCharacterIDOpportunitiesServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCharactersCharacterIDOpportunitiesServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDOpportunitiesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/opportunities/][%d] getCharactersCharacterIdOpportunitiesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDOpportunitiesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDOpportunitiesGatewayTimeout creates a GetCharactersCharacterIDOpportunitiesGatewayTimeout with default headers values
func NewGetCharactersCharacterIDOpportunitiesGatewayTimeout() *GetCharactersCharacterIDOpportunitiesGatewayTimeout {
	return &GetCharactersCharacterIDOpportunitiesGatewayTimeout{}
}

/*GetCharactersCharacterIDOpportunitiesGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDOpportunitiesGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDOpportunitiesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/opportunities/][%d] getCharactersCharacterIdOpportunitiesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDOpportunitiesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
