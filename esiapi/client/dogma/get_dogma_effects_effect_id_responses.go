// Code generated by go-swagger; DO NOT EDIT.

package dogma

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/eve-tmarg/esiapi/models"
)

// GetDogmaEffectsEffectIDReader is a Reader for the GetDogmaEffectsEffectID structure.
type GetDogmaEffectsEffectIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDogmaEffectsEffectIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDogmaEffectsEffectIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetDogmaEffectsEffectIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetDogmaEffectsEffectIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetDogmaEffectsEffectIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetDogmaEffectsEffectIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetDogmaEffectsEffectIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetDogmaEffectsEffectIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetDogmaEffectsEffectIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetDogmaEffectsEffectIDOK creates a GetDogmaEffectsEffectIDOK with default headers values
func NewGetDogmaEffectsEffectIDOK() *GetDogmaEffectsEffectIDOK {
	return &GetDogmaEffectsEffectIDOK{}
}

/*GetDogmaEffectsEffectIDOK handles this case with default header values.

Information about a dogma effect
*/
type GetDogmaEffectsEffectIDOK struct {
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

	Payload *models.GetDogmaEffectsEffectIDOKBody
}

func (o *GetDogmaEffectsEffectIDOK) Error() string {
	return fmt.Sprintf("[GET /dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdOK  %+v", 200, o.Payload)
}

func (o *GetDogmaEffectsEffectIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetDogmaEffectsEffectIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDogmaEffectsEffectIDNotModified creates a GetDogmaEffectsEffectIDNotModified with default headers values
func NewGetDogmaEffectsEffectIDNotModified() *GetDogmaEffectsEffectIDNotModified {
	return &GetDogmaEffectsEffectIDNotModified{}
}

/*GetDogmaEffectsEffectIDNotModified handles this case with default header values.

Not modified
*/
type GetDogmaEffectsEffectIDNotModified struct {
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

func (o *GetDogmaEffectsEffectIDNotModified) Error() string {
	return fmt.Sprintf("[GET /dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdNotModified ", 304)
}

func (o *GetDogmaEffectsEffectIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetDogmaEffectsEffectIDBadRequest creates a GetDogmaEffectsEffectIDBadRequest with default headers values
func NewGetDogmaEffectsEffectIDBadRequest() *GetDogmaEffectsEffectIDBadRequest {
	return &GetDogmaEffectsEffectIDBadRequest{}
}

/*GetDogmaEffectsEffectIDBadRequest handles this case with default header values.

Bad request
*/
type GetDogmaEffectsEffectIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetDogmaEffectsEffectIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetDogmaEffectsEffectIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDogmaEffectsEffectIDNotFound creates a GetDogmaEffectsEffectIDNotFound with default headers values
func NewGetDogmaEffectsEffectIDNotFound() *GetDogmaEffectsEffectIDNotFound {
	return &GetDogmaEffectsEffectIDNotFound{}
}

/*GetDogmaEffectsEffectIDNotFound handles this case with default header values.

Dogma effect not found
*/
type GetDogmaEffectsEffectIDNotFound struct {
	Payload *models.GetDogmaEffectsEffectIDNotFoundBody
}

func (o *GetDogmaEffectsEffectIDNotFound) Error() string {
	return fmt.Sprintf("[GET /dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdNotFound  %+v", 404, o.Payload)
}

func (o *GetDogmaEffectsEffectIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetDogmaEffectsEffectIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDogmaEffectsEffectIDEnhanceYourCalm creates a GetDogmaEffectsEffectIDEnhanceYourCalm with default headers values
func NewGetDogmaEffectsEffectIDEnhanceYourCalm() *GetDogmaEffectsEffectIDEnhanceYourCalm {
	return &GetDogmaEffectsEffectIDEnhanceYourCalm{}
}

/*GetDogmaEffectsEffectIDEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetDogmaEffectsEffectIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetDogmaEffectsEffectIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetDogmaEffectsEffectIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDogmaEffectsEffectIDInternalServerError creates a GetDogmaEffectsEffectIDInternalServerError with default headers values
func NewGetDogmaEffectsEffectIDInternalServerError() *GetDogmaEffectsEffectIDInternalServerError {
	return &GetDogmaEffectsEffectIDInternalServerError{}
}

/*GetDogmaEffectsEffectIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetDogmaEffectsEffectIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetDogmaEffectsEffectIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetDogmaEffectsEffectIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDogmaEffectsEffectIDServiceUnavailable creates a GetDogmaEffectsEffectIDServiceUnavailable with default headers values
func NewGetDogmaEffectsEffectIDServiceUnavailable() *GetDogmaEffectsEffectIDServiceUnavailable {
	return &GetDogmaEffectsEffectIDServiceUnavailable{}
}

/*GetDogmaEffectsEffectIDServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetDogmaEffectsEffectIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetDogmaEffectsEffectIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetDogmaEffectsEffectIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDogmaEffectsEffectIDGatewayTimeout creates a GetDogmaEffectsEffectIDGatewayTimeout with default headers values
func NewGetDogmaEffectsEffectIDGatewayTimeout() *GetDogmaEffectsEffectIDGatewayTimeout {
	return &GetDogmaEffectsEffectIDGatewayTimeout{}
}

/*GetDogmaEffectsEffectIDGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetDogmaEffectsEffectIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetDogmaEffectsEffectIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /dogma/effects/{effect_id}/][%d] getDogmaEffectsEffectIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetDogmaEffectsEffectIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
