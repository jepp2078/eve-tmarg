// Code generated by go-swagger; DO NOT EDIT.

package universe

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/eve-tmarg/esiapi/models"
)

// GetUniverseTypesReader is a Reader for the GetUniverseTypes structure.
type GetUniverseTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetUniverseTypesNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetUniverseTypesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetUniverseTypesEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetUniverseTypesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetUniverseTypesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetUniverseTypesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseTypesOK creates a GetUniverseTypesOK with default headers values
func NewGetUniverseTypesOK() *GetUniverseTypesOK {
	return &GetUniverseTypesOK{
		XPages: 1,
	}
}

/*GetUniverseTypesOK handles this case with default header values.

A list of type ids
*/
type GetUniverseTypesOK struct {
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
	/*Maximum page number
	 */
	XPages int32

	Payload []int32
}

func (o *GetUniverseTypesOK) Error() string {
	return fmt.Sprintf("[GET /universe/types/][%d] getUniverseTypesOK  %+v", 200, o.Payload)
}

func (o *GetUniverseTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header ETag
	o.ETag = response.GetHeader("ETag")

	// response header Expires
	o.Expires = response.GetHeader("Expires")

	// response header Last-Modified
	o.LastModified = response.GetHeader("Last-Modified")

	// response header X-Pages
	xPages, err := swag.ConvertInt32(response.GetHeader("X-Pages"))
	if err != nil {
		return errors.InvalidType("X-Pages", "header", "int32", response.GetHeader("X-Pages"))
	}
	o.XPages = xPages

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseTypesNotModified creates a GetUniverseTypesNotModified with default headers values
func NewGetUniverseTypesNotModified() *GetUniverseTypesNotModified {
	return &GetUniverseTypesNotModified{}
}

/*GetUniverseTypesNotModified handles this case with default header values.

Not modified
*/
type GetUniverseTypesNotModified struct {
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

func (o *GetUniverseTypesNotModified) Error() string {
	return fmt.Sprintf("[GET /universe/types/][%d] getUniverseTypesNotModified ", 304)
}

func (o *GetUniverseTypesNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseTypesBadRequest creates a GetUniverseTypesBadRequest with default headers values
func NewGetUniverseTypesBadRequest() *GetUniverseTypesBadRequest {
	return &GetUniverseTypesBadRequest{}
}

/*GetUniverseTypesBadRequest handles this case with default header values.

Bad request
*/
type GetUniverseTypesBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetUniverseTypesBadRequest) Error() string {
	return fmt.Sprintf("[GET /universe/types/][%d] getUniverseTypesBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseTypesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseTypesEnhanceYourCalm creates a GetUniverseTypesEnhanceYourCalm with default headers values
func NewGetUniverseTypesEnhanceYourCalm() *GetUniverseTypesEnhanceYourCalm {
	return &GetUniverseTypesEnhanceYourCalm{}
}

/*GetUniverseTypesEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetUniverseTypesEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetUniverseTypesEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /universe/types/][%d] getUniverseTypesEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseTypesEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseTypesInternalServerError creates a GetUniverseTypesInternalServerError with default headers values
func NewGetUniverseTypesInternalServerError() *GetUniverseTypesInternalServerError {
	return &GetUniverseTypesInternalServerError{}
}

/*GetUniverseTypesInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseTypesInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetUniverseTypesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/types/][%d] getUniverseTypesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseTypesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseTypesServiceUnavailable creates a GetUniverseTypesServiceUnavailable with default headers values
func NewGetUniverseTypesServiceUnavailable() *GetUniverseTypesServiceUnavailable {
	return &GetUniverseTypesServiceUnavailable{}
}

/*GetUniverseTypesServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetUniverseTypesServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetUniverseTypesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /universe/types/][%d] getUniverseTypesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseTypesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseTypesGatewayTimeout creates a GetUniverseTypesGatewayTimeout with default headers values
func NewGetUniverseTypesGatewayTimeout() *GetUniverseTypesGatewayTimeout {
	return &GetUniverseTypesGatewayTimeout{}
}

/*GetUniverseTypesGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetUniverseTypesGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetUniverseTypesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /universe/types/][%d] getUniverseTypesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseTypesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}