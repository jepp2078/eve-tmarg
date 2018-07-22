// Code generated by go-swagger; DO NOT EDIT.

package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/eve-tmarg/esiapi/models"
)

// GetUniverseAsteroidBeltsAsteroidBeltIDReader is a Reader for the GetUniverseAsteroidBeltsAsteroidBeltID structure.
type GetUniverseAsteroidBeltsAsteroidBeltIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseAsteroidBeltsAsteroidBeltIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseAsteroidBeltsAsteroidBeltIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetUniverseAsteroidBeltsAsteroidBeltIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetUniverseAsteroidBeltsAsteroidBeltIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetUniverseAsteroidBeltsAsteroidBeltIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetUniverseAsteroidBeltsAsteroidBeltIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetUniverseAsteroidBeltsAsteroidBeltIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetUniverseAsteroidBeltsAsteroidBeltIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetUniverseAsteroidBeltsAsteroidBeltIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseAsteroidBeltsAsteroidBeltIDOK creates a GetUniverseAsteroidBeltsAsteroidBeltIDOK with default headers values
func NewGetUniverseAsteroidBeltsAsteroidBeltIDOK() *GetUniverseAsteroidBeltsAsteroidBeltIDOK {
	return &GetUniverseAsteroidBeltsAsteroidBeltIDOK{}
}

/*GetUniverseAsteroidBeltsAsteroidBeltIDOK handles this case with default header values.

Information about an asteroid belt
*/
type GetUniverseAsteroidBeltsAsteroidBeltIDOK struct {
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

	Payload *models.GetUniverseAsteroidBeltsAsteroidBeltIDOKBody
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDOK) Error() string {
	return fmt.Sprintf("[GET /universe/asteroid_belts/{asteroid_belt_id}/][%d] getUniverseAsteroidBeltsAsteroidBeltIdOK  %+v", 200, o.Payload)
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetUniverseAsteroidBeltsAsteroidBeltIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseAsteroidBeltsAsteroidBeltIDNotModified creates a GetUniverseAsteroidBeltsAsteroidBeltIDNotModified with default headers values
func NewGetUniverseAsteroidBeltsAsteroidBeltIDNotModified() *GetUniverseAsteroidBeltsAsteroidBeltIDNotModified {
	return &GetUniverseAsteroidBeltsAsteroidBeltIDNotModified{}
}

/*GetUniverseAsteroidBeltsAsteroidBeltIDNotModified handles this case with default header values.

Not modified
*/
type GetUniverseAsteroidBeltsAsteroidBeltIDNotModified struct {
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

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDNotModified) Error() string {
	return fmt.Sprintf("[GET /universe/asteroid_belts/{asteroid_belt_id}/][%d] getUniverseAsteroidBeltsAsteroidBeltIdNotModified ", 304)
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseAsteroidBeltsAsteroidBeltIDBadRequest creates a GetUniverseAsteroidBeltsAsteroidBeltIDBadRequest with default headers values
func NewGetUniverseAsteroidBeltsAsteroidBeltIDBadRequest() *GetUniverseAsteroidBeltsAsteroidBeltIDBadRequest {
	return &GetUniverseAsteroidBeltsAsteroidBeltIDBadRequest{}
}

/*GetUniverseAsteroidBeltsAsteroidBeltIDBadRequest handles this case with default header values.

Bad request
*/
type GetUniverseAsteroidBeltsAsteroidBeltIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /universe/asteroid_belts/{asteroid_belt_id}/][%d] getUniverseAsteroidBeltsAsteroidBeltIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseAsteroidBeltsAsteroidBeltIDNotFound creates a GetUniverseAsteroidBeltsAsteroidBeltIDNotFound with default headers values
func NewGetUniverseAsteroidBeltsAsteroidBeltIDNotFound() *GetUniverseAsteroidBeltsAsteroidBeltIDNotFound {
	return &GetUniverseAsteroidBeltsAsteroidBeltIDNotFound{}
}

/*GetUniverseAsteroidBeltsAsteroidBeltIDNotFound handles this case with default header values.

Asteroid belt not found
*/
type GetUniverseAsteroidBeltsAsteroidBeltIDNotFound struct {
	Payload *models.GetUniverseAsteroidBeltsAsteroidBeltIDNotFoundBody
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDNotFound) Error() string {
	return fmt.Sprintf("[GET /universe/asteroid_belts/{asteroid_belt_id}/][%d] getUniverseAsteroidBeltsAsteroidBeltIdNotFound  %+v", 404, o.Payload)
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUniverseAsteroidBeltsAsteroidBeltIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseAsteroidBeltsAsteroidBeltIDEnhanceYourCalm creates a GetUniverseAsteroidBeltsAsteroidBeltIDEnhanceYourCalm with default headers values
func NewGetUniverseAsteroidBeltsAsteroidBeltIDEnhanceYourCalm() *GetUniverseAsteroidBeltsAsteroidBeltIDEnhanceYourCalm {
	return &GetUniverseAsteroidBeltsAsteroidBeltIDEnhanceYourCalm{}
}

/*GetUniverseAsteroidBeltsAsteroidBeltIDEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetUniverseAsteroidBeltsAsteroidBeltIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /universe/asteroid_belts/{asteroid_belt_id}/][%d] getUniverseAsteroidBeltsAsteroidBeltIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseAsteroidBeltsAsteroidBeltIDInternalServerError creates a GetUniverseAsteroidBeltsAsteroidBeltIDInternalServerError with default headers values
func NewGetUniverseAsteroidBeltsAsteroidBeltIDInternalServerError() *GetUniverseAsteroidBeltsAsteroidBeltIDInternalServerError {
	return &GetUniverseAsteroidBeltsAsteroidBeltIDInternalServerError{}
}

/*GetUniverseAsteroidBeltsAsteroidBeltIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseAsteroidBeltsAsteroidBeltIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/asteroid_belts/{asteroid_belt_id}/][%d] getUniverseAsteroidBeltsAsteroidBeltIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseAsteroidBeltsAsteroidBeltIDServiceUnavailable creates a GetUniverseAsteroidBeltsAsteroidBeltIDServiceUnavailable with default headers values
func NewGetUniverseAsteroidBeltsAsteroidBeltIDServiceUnavailable() *GetUniverseAsteroidBeltsAsteroidBeltIDServiceUnavailable {
	return &GetUniverseAsteroidBeltsAsteroidBeltIDServiceUnavailable{}
}

/*GetUniverseAsteroidBeltsAsteroidBeltIDServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetUniverseAsteroidBeltsAsteroidBeltIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /universe/asteroid_belts/{asteroid_belt_id}/][%d] getUniverseAsteroidBeltsAsteroidBeltIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseAsteroidBeltsAsteroidBeltIDGatewayTimeout creates a GetUniverseAsteroidBeltsAsteroidBeltIDGatewayTimeout with default headers values
func NewGetUniverseAsteroidBeltsAsteroidBeltIDGatewayTimeout() *GetUniverseAsteroidBeltsAsteroidBeltIDGatewayTimeout {
	return &GetUniverseAsteroidBeltsAsteroidBeltIDGatewayTimeout{}
}

/*GetUniverseAsteroidBeltsAsteroidBeltIDGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetUniverseAsteroidBeltsAsteroidBeltIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /universe/asteroid_belts/{asteroid_belt_id}/][%d] getUniverseAsteroidBeltsAsteroidBeltIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseAsteroidBeltsAsteroidBeltIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}