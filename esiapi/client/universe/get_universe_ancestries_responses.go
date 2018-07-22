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

// GetUniverseAncestriesReader is a Reader for the GetUniverseAncestries structure.
type GetUniverseAncestriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseAncestriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseAncestriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetUniverseAncestriesNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetUniverseAncestriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetUniverseAncestriesEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetUniverseAncestriesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetUniverseAncestriesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetUniverseAncestriesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseAncestriesOK creates a GetUniverseAncestriesOK with default headers values
func NewGetUniverseAncestriesOK() *GetUniverseAncestriesOK {
	return &GetUniverseAncestriesOK{}
}

/*GetUniverseAncestriesOK handles this case with default header values.

A list of ancestries
*/
type GetUniverseAncestriesOK struct {
	/*The caching mechanism used
	 */
	CacheControl string
	/*The language used in the response
	 */
	ContentLanguage string
	/*RFC7232 compliant entity tag
	 */
	ETag string
	/*RFC7231 formatted datetime string
	 */
	Expires string
	/*RFC7231 formatted datetime string
	 */
	LastModified string

	Payload []*models.GetUniverseAncestriesOKBodyItems
}

func (o *GetUniverseAncestriesOK) Error() string {
	return fmt.Sprintf("[GET /universe/ancestries/][%d] getUniverseAncestriesOK  %+v", 200, o.Payload)
}

func (o *GetUniverseAncestriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Cache-Control
	o.CacheControl = response.GetHeader("Cache-Control")

	// response header Content-Language
	o.ContentLanguage = response.GetHeader("Content-Language")

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

// NewGetUniverseAncestriesNotModified creates a GetUniverseAncestriesNotModified with default headers values
func NewGetUniverseAncestriesNotModified() *GetUniverseAncestriesNotModified {
	return &GetUniverseAncestriesNotModified{}
}

/*GetUniverseAncestriesNotModified handles this case with default header values.

Not modified
*/
type GetUniverseAncestriesNotModified struct {
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

func (o *GetUniverseAncestriesNotModified) Error() string {
	return fmt.Sprintf("[GET /universe/ancestries/][%d] getUniverseAncestriesNotModified ", 304)
}

func (o *GetUniverseAncestriesNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseAncestriesBadRequest creates a GetUniverseAncestriesBadRequest with default headers values
func NewGetUniverseAncestriesBadRequest() *GetUniverseAncestriesBadRequest {
	return &GetUniverseAncestriesBadRequest{}
}

/*GetUniverseAncestriesBadRequest handles this case with default header values.

Bad request
*/
type GetUniverseAncestriesBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetUniverseAncestriesBadRequest) Error() string {
	return fmt.Sprintf("[GET /universe/ancestries/][%d] getUniverseAncestriesBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseAncestriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseAncestriesEnhanceYourCalm creates a GetUniverseAncestriesEnhanceYourCalm with default headers values
func NewGetUniverseAncestriesEnhanceYourCalm() *GetUniverseAncestriesEnhanceYourCalm {
	return &GetUniverseAncestriesEnhanceYourCalm{}
}

/*GetUniverseAncestriesEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetUniverseAncestriesEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetUniverseAncestriesEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /universe/ancestries/][%d] getUniverseAncestriesEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseAncestriesEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseAncestriesInternalServerError creates a GetUniverseAncestriesInternalServerError with default headers values
func NewGetUniverseAncestriesInternalServerError() *GetUniverseAncestriesInternalServerError {
	return &GetUniverseAncestriesInternalServerError{}
}

/*GetUniverseAncestriesInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseAncestriesInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetUniverseAncestriesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/ancestries/][%d] getUniverseAncestriesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseAncestriesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseAncestriesServiceUnavailable creates a GetUniverseAncestriesServiceUnavailable with default headers values
func NewGetUniverseAncestriesServiceUnavailable() *GetUniverseAncestriesServiceUnavailable {
	return &GetUniverseAncestriesServiceUnavailable{}
}

/*GetUniverseAncestriesServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetUniverseAncestriesServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetUniverseAncestriesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /universe/ancestries/][%d] getUniverseAncestriesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseAncestriesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseAncestriesGatewayTimeout creates a GetUniverseAncestriesGatewayTimeout with default headers values
func NewGetUniverseAncestriesGatewayTimeout() *GetUniverseAncestriesGatewayTimeout {
	return &GetUniverseAncestriesGatewayTimeout{}
}

/*GetUniverseAncestriesGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetUniverseAncestriesGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetUniverseAncestriesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /universe/ancestries/][%d] getUniverseAncestriesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseAncestriesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
