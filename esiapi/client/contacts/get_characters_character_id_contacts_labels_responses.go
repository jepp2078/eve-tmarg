// Code generated by go-swagger; DO NOT EDIT.

package contacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/eve-tmarg/esiapi/models"
)

// GetCharactersCharacterIDContactsLabelsReader is a Reader for the GetCharactersCharacterIDContactsLabels structure.
type GetCharactersCharacterIDContactsLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDContactsLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDContactsLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCharactersCharacterIDContactsLabelsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCharactersCharacterIDContactsLabelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCharactersCharacterIDContactsLabelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCharactersCharacterIDContactsLabelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCharactersCharacterIDContactsLabelsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDContactsLabelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCharactersCharacterIDContactsLabelsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCharactersCharacterIDContactsLabelsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDContactsLabelsOK creates a GetCharactersCharacterIDContactsLabelsOK with default headers values
func NewGetCharactersCharacterIDContactsLabelsOK() *GetCharactersCharacterIDContactsLabelsOK {
	return &GetCharactersCharacterIDContactsLabelsOK{}
}

/*GetCharactersCharacterIDContactsLabelsOK handles this case with default header values.

A list of contact labels
*/
type GetCharactersCharacterIDContactsLabelsOK struct {
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

	Payload []*models.GetCharactersCharacterIDContactsLabelsOKBodyItems
}

func (o *GetCharactersCharacterIDContactsLabelsOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contacts/labels/][%d] getCharactersCharacterIdContactsLabelsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDContactsLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDContactsLabelsNotModified creates a GetCharactersCharacterIDContactsLabelsNotModified with default headers values
func NewGetCharactersCharacterIDContactsLabelsNotModified() *GetCharactersCharacterIDContactsLabelsNotModified {
	return &GetCharactersCharacterIDContactsLabelsNotModified{}
}

/*GetCharactersCharacterIDContactsLabelsNotModified handles this case with default header values.

Not modified
*/
type GetCharactersCharacterIDContactsLabelsNotModified struct {
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

func (o *GetCharactersCharacterIDContactsLabelsNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contacts/labels/][%d] getCharactersCharacterIdContactsLabelsNotModified ", 304)
}

func (o *GetCharactersCharacterIDContactsLabelsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDContactsLabelsBadRequest creates a GetCharactersCharacterIDContactsLabelsBadRequest with default headers values
func NewGetCharactersCharacterIDContactsLabelsBadRequest() *GetCharactersCharacterIDContactsLabelsBadRequest {
	return &GetCharactersCharacterIDContactsLabelsBadRequest{}
}

/*GetCharactersCharacterIDContactsLabelsBadRequest handles this case with default header values.

Bad request
*/
type GetCharactersCharacterIDContactsLabelsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDContactsLabelsBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contacts/labels/][%d] getCharactersCharacterIdContactsLabelsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDContactsLabelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContactsLabelsUnauthorized creates a GetCharactersCharacterIDContactsLabelsUnauthorized with default headers values
func NewGetCharactersCharacterIDContactsLabelsUnauthorized() *GetCharactersCharacterIDContactsLabelsUnauthorized {
	return &GetCharactersCharacterIDContactsLabelsUnauthorized{}
}

/*GetCharactersCharacterIDContactsLabelsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCharactersCharacterIDContactsLabelsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDContactsLabelsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contacts/labels/][%d] getCharactersCharacterIdContactsLabelsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDContactsLabelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContactsLabelsForbidden creates a GetCharactersCharacterIDContactsLabelsForbidden with default headers values
func NewGetCharactersCharacterIDContactsLabelsForbidden() *GetCharactersCharacterIDContactsLabelsForbidden {
	return &GetCharactersCharacterIDContactsLabelsForbidden{}
}

/*GetCharactersCharacterIDContactsLabelsForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDContactsLabelsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDContactsLabelsForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contacts/labels/][%d] getCharactersCharacterIdContactsLabelsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDContactsLabelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContactsLabelsEnhanceYourCalm creates a GetCharactersCharacterIDContactsLabelsEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDContactsLabelsEnhanceYourCalm() *GetCharactersCharacterIDContactsLabelsEnhanceYourCalm {
	return &GetCharactersCharacterIDContactsLabelsEnhanceYourCalm{}
}

/*GetCharactersCharacterIDContactsLabelsEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCharactersCharacterIDContactsLabelsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDContactsLabelsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contacts/labels/][%d] getCharactersCharacterIdContactsLabelsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDContactsLabelsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContactsLabelsInternalServerError creates a GetCharactersCharacterIDContactsLabelsInternalServerError with default headers values
func NewGetCharactersCharacterIDContactsLabelsInternalServerError() *GetCharactersCharacterIDContactsLabelsInternalServerError {
	return &GetCharactersCharacterIDContactsLabelsInternalServerError{}
}

/*GetCharactersCharacterIDContactsLabelsInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDContactsLabelsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDContactsLabelsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contacts/labels/][%d] getCharactersCharacterIdContactsLabelsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDContactsLabelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContactsLabelsServiceUnavailable creates a GetCharactersCharacterIDContactsLabelsServiceUnavailable with default headers values
func NewGetCharactersCharacterIDContactsLabelsServiceUnavailable() *GetCharactersCharacterIDContactsLabelsServiceUnavailable {
	return &GetCharactersCharacterIDContactsLabelsServiceUnavailable{}
}

/*GetCharactersCharacterIDContactsLabelsServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCharactersCharacterIDContactsLabelsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDContactsLabelsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contacts/labels/][%d] getCharactersCharacterIdContactsLabelsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDContactsLabelsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContactsLabelsGatewayTimeout creates a GetCharactersCharacterIDContactsLabelsGatewayTimeout with default headers values
func NewGetCharactersCharacterIDContactsLabelsGatewayTimeout() *GetCharactersCharacterIDContactsLabelsGatewayTimeout {
	return &GetCharactersCharacterIDContactsLabelsGatewayTimeout{}
}

/*GetCharactersCharacterIDContactsLabelsGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDContactsLabelsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDContactsLabelsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contacts/labels/][%d] getCharactersCharacterIdContactsLabelsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDContactsLabelsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
