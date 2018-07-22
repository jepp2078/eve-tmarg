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

// GetCharactersCharacterIDMailLabelsReader is a Reader for the GetCharactersCharacterIDMailLabels structure.
type GetCharactersCharacterIDMailLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDMailLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDMailLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCharactersCharacterIDMailLabelsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCharactersCharacterIDMailLabelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCharactersCharacterIDMailLabelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCharactersCharacterIDMailLabelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCharactersCharacterIDMailLabelsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDMailLabelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCharactersCharacterIDMailLabelsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCharactersCharacterIDMailLabelsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDMailLabelsOK creates a GetCharactersCharacterIDMailLabelsOK with default headers values
func NewGetCharactersCharacterIDMailLabelsOK() *GetCharactersCharacterIDMailLabelsOK {
	return &GetCharactersCharacterIDMailLabelsOK{}
}

/*GetCharactersCharacterIDMailLabelsOK handles this case with default header values.

A list of mail labels and unread counts
*/
type GetCharactersCharacterIDMailLabelsOK struct {
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

	Payload *models.GetCharactersCharacterIDMailLabelsOKBody
}

func (o *GetCharactersCharacterIDMailLabelsOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/labels/][%d] getCharactersCharacterIdMailLabelsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDMailLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetCharactersCharacterIDMailLabelsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailLabelsNotModified creates a GetCharactersCharacterIDMailLabelsNotModified with default headers values
func NewGetCharactersCharacterIDMailLabelsNotModified() *GetCharactersCharacterIDMailLabelsNotModified {
	return &GetCharactersCharacterIDMailLabelsNotModified{}
}

/*GetCharactersCharacterIDMailLabelsNotModified handles this case with default header values.

Not modified
*/
type GetCharactersCharacterIDMailLabelsNotModified struct {
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

func (o *GetCharactersCharacterIDMailLabelsNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/labels/][%d] getCharactersCharacterIdMailLabelsNotModified ", 304)
}

func (o *GetCharactersCharacterIDMailLabelsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDMailLabelsBadRequest creates a GetCharactersCharacterIDMailLabelsBadRequest with default headers values
func NewGetCharactersCharacterIDMailLabelsBadRequest() *GetCharactersCharacterIDMailLabelsBadRequest {
	return &GetCharactersCharacterIDMailLabelsBadRequest{}
}

/*GetCharactersCharacterIDMailLabelsBadRequest handles this case with default header values.

Bad request
*/
type GetCharactersCharacterIDMailLabelsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDMailLabelsBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/labels/][%d] getCharactersCharacterIdMailLabelsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDMailLabelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailLabelsUnauthorized creates a GetCharactersCharacterIDMailLabelsUnauthorized with default headers values
func NewGetCharactersCharacterIDMailLabelsUnauthorized() *GetCharactersCharacterIDMailLabelsUnauthorized {
	return &GetCharactersCharacterIDMailLabelsUnauthorized{}
}

/*GetCharactersCharacterIDMailLabelsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCharactersCharacterIDMailLabelsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDMailLabelsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/labels/][%d] getCharactersCharacterIdMailLabelsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDMailLabelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailLabelsForbidden creates a GetCharactersCharacterIDMailLabelsForbidden with default headers values
func NewGetCharactersCharacterIDMailLabelsForbidden() *GetCharactersCharacterIDMailLabelsForbidden {
	return &GetCharactersCharacterIDMailLabelsForbidden{}
}

/*GetCharactersCharacterIDMailLabelsForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDMailLabelsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDMailLabelsForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/labels/][%d] getCharactersCharacterIdMailLabelsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDMailLabelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailLabelsEnhanceYourCalm creates a GetCharactersCharacterIDMailLabelsEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDMailLabelsEnhanceYourCalm() *GetCharactersCharacterIDMailLabelsEnhanceYourCalm {
	return &GetCharactersCharacterIDMailLabelsEnhanceYourCalm{}
}

/*GetCharactersCharacterIDMailLabelsEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCharactersCharacterIDMailLabelsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDMailLabelsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/labels/][%d] getCharactersCharacterIdMailLabelsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDMailLabelsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailLabelsInternalServerError creates a GetCharactersCharacterIDMailLabelsInternalServerError with default headers values
func NewGetCharactersCharacterIDMailLabelsInternalServerError() *GetCharactersCharacterIDMailLabelsInternalServerError {
	return &GetCharactersCharacterIDMailLabelsInternalServerError{}
}

/*GetCharactersCharacterIDMailLabelsInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDMailLabelsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDMailLabelsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/labels/][%d] getCharactersCharacterIdMailLabelsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDMailLabelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailLabelsServiceUnavailable creates a GetCharactersCharacterIDMailLabelsServiceUnavailable with default headers values
func NewGetCharactersCharacterIDMailLabelsServiceUnavailable() *GetCharactersCharacterIDMailLabelsServiceUnavailable {
	return &GetCharactersCharacterIDMailLabelsServiceUnavailable{}
}

/*GetCharactersCharacterIDMailLabelsServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCharactersCharacterIDMailLabelsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDMailLabelsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/labels/][%d] getCharactersCharacterIdMailLabelsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDMailLabelsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDMailLabelsGatewayTimeout creates a GetCharactersCharacterIDMailLabelsGatewayTimeout with default headers values
func NewGetCharactersCharacterIDMailLabelsGatewayTimeout() *GetCharactersCharacterIDMailLabelsGatewayTimeout {
	return &GetCharactersCharacterIDMailLabelsGatewayTimeout{}
}

/*GetCharactersCharacterIDMailLabelsGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDMailLabelsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDMailLabelsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/mail/labels/][%d] getCharactersCharacterIdMailLabelsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDMailLabelsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}