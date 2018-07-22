// Code generated by go-swagger; DO NOT EDIT.

package mail

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/eve-tmarg/esiapi/models"
)

// GetCharactersCharacterIDMailReader is a Reader for the GetCharactersCharacterIDMail structure.
type GetCharactersCharacterIDMailReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDMailReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDMailOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCharactersCharacterIDMailNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCharactersCharacterIDMailBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCharactersCharacterIDMailUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCharactersCharacterIDMailForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCharactersCharacterIDMailEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDMailInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCharactersCharacterIDMailServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCharactersCharacterIDMailGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDMailOK creates a GetCharactersCharacterIDMailOK with default headers values
func NewGetCharactersCharacterIDMailOK() *GetCharactersCharacterIDMailOK {
	return &GetCharactersCharacterIDMailOK{}
}

/*GetCharactersCharacterIDMailOK handles this case with default header values.

The requested mail
*/
type GetCharactersCharacterIDMailOK struct {
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

	Payload []*models.GetCharactersCharacterIDMailOKBodyItems
}

func (o *GetCharactersCharacterIDMailOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/][%d] getCharactersCharacterIdMailOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDMailOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDMailNotModified creates a GetCharactersCharacterIDMailNotModified with default headers values
func NewGetCharactersCharacterIDMailNotModified() *GetCharactersCharacterIDMailNotModified {
	return &GetCharactersCharacterIDMailNotModified{}
}

/*GetCharactersCharacterIDMailNotModified handles this case with default header values.

Not modified
*/
type GetCharactersCharacterIDMailNotModified struct {
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

func (o *GetCharactersCharacterIDMailNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/][%d] getCharactersCharacterIdMailNotModified ", 304)
}

func (o *GetCharactersCharacterIDMailNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDMailBadRequest creates a GetCharactersCharacterIDMailBadRequest with default headers values
func NewGetCharactersCharacterIDMailBadRequest() *GetCharactersCharacterIDMailBadRequest {
	return &GetCharactersCharacterIDMailBadRequest{}
}

/*GetCharactersCharacterIDMailBadRequest handles this case with default header values.

Bad request
*/
type GetCharactersCharacterIDMailBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDMailBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/][%d] getCharactersCharacterIdMailBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDMailBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailUnauthorized creates a GetCharactersCharacterIDMailUnauthorized with default headers values
func NewGetCharactersCharacterIDMailUnauthorized() *GetCharactersCharacterIDMailUnauthorized {
	return &GetCharactersCharacterIDMailUnauthorized{}
}

/*GetCharactersCharacterIDMailUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCharactersCharacterIDMailUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDMailUnauthorized) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/][%d] getCharactersCharacterIdMailUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDMailUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailForbidden creates a GetCharactersCharacterIDMailForbidden with default headers values
func NewGetCharactersCharacterIDMailForbidden() *GetCharactersCharacterIDMailForbidden {
	return &GetCharactersCharacterIDMailForbidden{}
}

/*GetCharactersCharacterIDMailForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDMailForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDMailForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/][%d] getCharactersCharacterIdMailForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDMailForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailEnhanceYourCalm creates a GetCharactersCharacterIDMailEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDMailEnhanceYourCalm() *GetCharactersCharacterIDMailEnhanceYourCalm {
	return &GetCharactersCharacterIDMailEnhanceYourCalm{}
}

/*GetCharactersCharacterIDMailEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCharactersCharacterIDMailEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDMailEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/][%d] getCharactersCharacterIdMailEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDMailEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailInternalServerError creates a GetCharactersCharacterIDMailInternalServerError with default headers values
func NewGetCharactersCharacterIDMailInternalServerError() *GetCharactersCharacterIDMailInternalServerError {
	return &GetCharactersCharacterIDMailInternalServerError{}
}

/*GetCharactersCharacterIDMailInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDMailInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDMailInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/][%d] getCharactersCharacterIdMailInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDMailInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailServiceUnavailable creates a GetCharactersCharacterIDMailServiceUnavailable with default headers values
func NewGetCharactersCharacterIDMailServiceUnavailable() *GetCharactersCharacterIDMailServiceUnavailable {
	return &GetCharactersCharacterIDMailServiceUnavailable{}
}

/*GetCharactersCharacterIDMailServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCharactersCharacterIDMailServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDMailServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/][%d] getCharactersCharacterIdMailServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDMailServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailGatewayTimeout creates a GetCharactersCharacterIDMailGatewayTimeout with default headers values
func NewGetCharactersCharacterIDMailGatewayTimeout() *GetCharactersCharacterIDMailGatewayTimeout {
	return &GetCharactersCharacterIDMailGatewayTimeout{}
}

/*GetCharactersCharacterIDMailGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDMailGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDMailGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/][%d] getCharactersCharacterIdMailGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDMailGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}