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

// GetUniverseStargatesStargateIDReader is a Reader for the GetUniverseStargatesStargateID structure.
type GetUniverseStargatesStargateIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseStargatesStargateIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseStargatesStargateIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetUniverseStargatesStargateIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetUniverseStargatesStargateIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetUniverseStargatesStargateIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetUniverseStargatesStargateIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetUniverseStargatesStargateIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetUniverseStargatesStargateIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetUniverseStargatesStargateIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseStargatesStargateIDOK creates a GetUniverseStargatesStargateIDOK with default headers values
func NewGetUniverseStargatesStargateIDOK() *GetUniverseStargatesStargateIDOK {
	return &GetUniverseStargatesStargateIDOK{}
}

/*GetUniverseStargatesStargateIDOK handles this case with default header values.

Information about a stargate
*/
type GetUniverseStargatesStargateIDOK struct {
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

	Payload *models.GetUniverseStargatesStargateIDOKBody
}

func (o *GetUniverseStargatesStargateIDOK) Error() string {
	return fmt.Sprintf("[GET /universe/stargates/{stargate_id}/][%d] getUniverseStargatesStargateIdOK  %+v", 200, o.Payload)
}

func (o *GetUniverseStargatesStargateIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	o.Payload = new(models.GetUniverseStargatesStargateIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStargatesStargateIDNotModified creates a GetUniverseStargatesStargateIDNotModified with default headers values
func NewGetUniverseStargatesStargateIDNotModified() *GetUniverseStargatesStargateIDNotModified {
	return &GetUniverseStargatesStargateIDNotModified{}
}

/*GetUniverseStargatesStargateIDNotModified handles this case with default header values.

Not modified
*/
type GetUniverseStargatesStargateIDNotModified struct {
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

func (o *GetUniverseStargatesStargateIDNotModified) Error() string {
	return fmt.Sprintf("[GET /universe/stargates/{stargate_id}/][%d] getUniverseStargatesStargateIdNotModified ", 304)
}

func (o *GetUniverseStargatesStargateIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseStargatesStargateIDBadRequest creates a GetUniverseStargatesStargateIDBadRequest with default headers values
func NewGetUniverseStargatesStargateIDBadRequest() *GetUniverseStargatesStargateIDBadRequest {
	return &GetUniverseStargatesStargateIDBadRequest{}
}

/*GetUniverseStargatesStargateIDBadRequest handles this case with default header values.

Bad request
*/
type GetUniverseStargatesStargateIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetUniverseStargatesStargateIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /universe/stargates/{stargate_id}/][%d] getUniverseStargatesStargateIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseStargatesStargateIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStargatesStargateIDNotFound creates a GetUniverseStargatesStargateIDNotFound with default headers values
func NewGetUniverseStargatesStargateIDNotFound() *GetUniverseStargatesStargateIDNotFound {
	return &GetUniverseStargatesStargateIDNotFound{}
}

/*GetUniverseStargatesStargateIDNotFound handles this case with default header values.

Stargate not found
*/
type GetUniverseStargatesStargateIDNotFound struct {
	Payload *models.GetUniverseStargatesStargateIDNotFoundBody
}

func (o *GetUniverseStargatesStargateIDNotFound) Error() string {
	return fmt.Sprintf("[GET /universe/stargates/{stargate_id}/][%d] getUniverseStargatesStargateIdNotFound  %+v", 404, o.Payload)
}

func (o *GetUniverseStargatesStargateIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUniverseStargatesStargateIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStargatesStargateIDEnhanceYourCalm creates a GetUniverseStargatesStargateIDEnhanceYourCalm with default headers values
func NewGetUniverseStargatesStargateIDEnhanceYourCalm() *GetUniverseStargatesStargateIDEnhanceYourCalm {
	return &GetUniverseStargatesStargateIDEnhanceYourCalm{}
}

/*GetUniverseStargatesStargateIDEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetUniverseStargatesStargateIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetUniverseStargatesStargateIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /universe/stargates/{stargate_id}/][%d] getUniverseStargatesStargateIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseStargatesStargateIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStargatesStargateIDInternalServerError creates a GetUniverseStargatesStargateIDInternalServerError with default headers values
func NewGetUniverseStargatesStargateIDInternalServerError() *GetUniverseStargatesStargateIDInternalServerError {
	return &GetUniverseStargatesStargateIDInternalServerError{}
}

/*GetUniverseStargatesStargateIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseStargatesStargateIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetUniverseStargatesStargateIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/stargates/{stargate_id}/][%d] getUniverseStargatesStargateIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseStargatesStargateIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStargatesStargateIDServiceUnavailable creates a GetUniverseStargatesStargateIDServiceUnavailable with default headers values
func NewGetUniverseStargatesStargateIDServiceUnavailable() *GetUniverseStargatesStargateIDServiceUnavailable {
	return &GetUniverseStargatesStargateIDServiceUnavailable{}
}

/*GetUniverseStargatesStargateIDServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetUniverseStargatesStargateIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetUniverseStargatesStargateIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /universe/stargates/{stargate_id}/][%d] getUniverseStargatesStargateIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseStargatesStargateIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStargatesStargateIDGatewayTimeout creates a GetUniverseStargatesStargateIDGatewayTimeout with default headers values
func NewGetUniverseStargatesStargateIDGatewayTimeout() *GetUniverseStargatesStargateIDGatewayTimeout {
	return &GetUniverseStargatesStargateIDGatewayTimeout{}
}

/*GetUniverseStargatesStargateIDGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetUniverseStargatesStargateIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetUniverseStargatesStargateIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /universe/stargates/{stargate_id}/][%d] getUniverseStargatesStargateIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseStargatesStargateIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}