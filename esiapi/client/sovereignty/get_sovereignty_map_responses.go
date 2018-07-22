// Code generated by go-swagger; DO NOT EDIT.

package sovereignty

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/eve-tmarg/esiapi/models"
)

// GetSovereigntyMapReader is a Reader for the GetSovereigntyMap structure.
type GetSovereigntyMapReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSovereigntyMapReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSovereigntyMapOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetSovereigntyMapNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetSovereigntyMapBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetSovereigntyMapEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetSovereigntyMapInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetSovereigntyMapServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetSovereigntyMapGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSovereigntyMapOK creates a GetSovereigntyMapOK with default headers values
func NewGetSovereigntyMapOK() *GetSovereigntyMapOK {
	return &GetSovereigntyMapOK{}
}

/*GetSovereigntyMapOK handles this case with default header values.

A list of sovereignty information for solar systems in New Eden
*/
type GetSovereigntyMapOK struct {
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

	Payload []*models.GetSovereigntyMapOKBodyItems
}

func (o *GetSovereigntyMapOK) Error() string {
	return fmt.Sprintf("[GET /sovereignty/map/][%d] getSovereigntyMapOK  %+v", 200, o.Payload)
}

func (o *GetSovereigntyMapOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSovereigntyMapNotModified creates a GetSovereigntyMapNotModified with default headers values
func NewGetSovereigntyMapNotModified() *GetSovereigntyMapNotModified {
	return &GetSovereigntyMapNotModified{}
}

/*GetSovereigntyMapNotModified handles this case with default header values.

Not modified
*/
type GetSovereigntyMapNotModified struct {
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

func (o *GetSovereigntyMapNotModified) Error() string {
	return fmt.Sprintf("[GET /sovereignty/map/][%d] getSovereigntyMapNotModified ", 304)
}

func (o *GetSovereigntyMapNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSovereigntyMapBadRequest creates a GetSovereigntyMapBadRequest with default headers values
func NewGetSovereigntyMapBadRequest() *GetSovereigntyMapBadRequest {
	return &GetSovereigntyMapBadRequest{}
}

/*GetSovereigntyMapBadRequest handles this case with default header values.

Bad request
*/
type GetSovereigntyMapBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetSovereigntyMapBadRequest) Error() string {
	return fmt.Sprintf("[GET /sovereignty/map/][%d] getSovereigntyMapBadRequest  %+v", 400, o.Payload)
}

func (o *GetSovereigntyMapBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSovereigntyMapEnhanceYourCalm creates a GetSovereigntyMapEnhanceYourCalm with default headers values
func NewGetSovereigntyMapEnhanceYourCalm() *GetSovereigntyMapEnhanceYourCalm {
	return &GetSovereigntyMapEnhanceYourCalm{}
}

/*GetSovereigntyMapEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetSovereigntyMapEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetSovereigntyMapEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /sovereignty/map/][%d] getSovereigntyMapEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetSovereigntyMapEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSovereigntyMapInternalServerError creates a GetSovereigntyMapInternalServerError with default headers values
func NewGetSovereigntyMapInternalServerError() *GetSovereigntyMapInternalServerError {
	return &GetSovereigntyMapInternalServerError{}
}

/*GetSovereigntyMapInternalServerError handles this case with default header values.

Internal server error
*/
type GetSovereigntyMapInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetSovereigntyMapInternalServerError) Error() string {
	return fmt.Sprintf("[GET /sovereignty/map/][%d] getSovereigntyMapInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSovereigntyMapInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSovereigntyMapServiceUnavailable creates a GetSovereigntyMapServiceUnavailable with default headers values
func NewGetSovereigntyMapServiceUnavailable() *GetSovereigntyMapServiceUnavailable {
	return &GetSovereigntyMapServiceUnavailable{}
}

/*GetSovereigntyMapServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetSovereigntyMapServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetSovereigntyMapServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /sovereignty/map/][%d] getSovereigntyMapServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSovereigntyMapServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSovereigntyMapGatewayTimeout creates a GetSovereigntyMapGatewayTimeout with default headers values
func NewGetSovereigntyMapGatewayTimeout() *GetSovereigntyMapGatewayTimeout {
	return &GetSovereigntyMapGatewayTimeout{}
}

/*GetSovereigntyMapGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetSovereigntyMapGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetSovereigntyMapGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /sovereignty/map/][%d] getSovereigntyMapGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSovereigntyMapGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
