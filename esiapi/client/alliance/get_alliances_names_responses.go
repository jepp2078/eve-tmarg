// Code generated by go-swagger; DO NOT EDIT.

package alliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/eve-tmarg/esiapi/models"
)

// GetAlliancesNamesReader is a Reader for the GetAlliancesNames structure.
type GetAlliancesNamesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlliancesNamesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAlliancesNamesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetAlliancesNamesNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetAlliancesNamesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetAlliancesNamesEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAlliancesNamesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetAlliancesNamesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetAlliancesNamesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAlliancesNamesOK creates a GetAlliancesNamesOK with default headers values
func NewGetAlliancesNamesOK() *GetAlliancesNamesOK {
	return &GetAlliancesNamesOK{}
}

/*GetAlliancesNamesOK handles this case with default header values.

List of id/name associations
*/
type GetAlliancesNamesOK struct {
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

	Payload []*models.GetAlliancesNamesOKBodyItems
}

func (o *GetAlliancesNamesOK) Error() string {
	return fmt.Sprintf("[GET /alliances/names/][%d] getAlliancesNamesOK  %+v", 200, o.Payload)
}

func (o *GetAlliancesNamesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAlliancesNamesNotModified creates a GetAlliancesNamesNotModified with default headers values
func NewGetAlliancesNamesNotModified() *GetAlliancesNamesNotModified {
	return &GetAlliancesNamesNotModified{}
}

/*GetAlliancesNamesNotModified handles this case with default header values.

Not modified
*/
type GetAlliancesNamesNotModified struct {
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

func (o *GetAlliancesNamesNotModified) Error() string {
	return fmt.Sprintf("[GET /alliances/names/][%d] getAlliancesNamesNotModified ", 304)
}

func (o *GetAlliancesNamesNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetAlliancesNamesBadRequest creates a GetAlliancesNamesBadRequest with default headers values
func NewGetAlliancesNamesBadRequest() *GetAlliancesNamesBadRequest {
	return &GetAlliancesNamesBadRequest{}
}

/*GetAlliancesNamesBadRequest handles this case with default header values.

Bad request
*/
type GetAlliancesNamesBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetAlliancesNamesBadRequest) Error() string {
	return fmt.Sprintf("[GET /alliances/names/][%d] getAlliancesNamesBadRequest  %+v", 400, o.Payload)
}

func (o *GetAlliancesNamesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesNamesEnhanceYourCalm creates a GetAlliancesNamesEnhanceYourCalm with default headers values
func NewGetAlliancesNamesEnhanceYourCalm() *GetAlliancesNamesEnhanceYourCalm {
	return &GetAlliancesNamesEnhanceYourCalm{}
}

/*GetAlliancesNamesEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetAlliancesNamesEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetAlliancesNamesEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /alliances/names/][%d] getAlliancesNamesEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetAlliancesNamesEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesNamesInternalServerError creates a GetAlliancesNamesInternalServerError with default headers values
func NewGetAlliancesNamesInternalServerError() *GetAlliancesNamesInternalServerError {
	return &GetAlliancesNamesInternalServerError{}
}

/*GetAlliancesNamesInternalServerError handles this case with default header values.

Internal server error
*/
type GetAlliancesNamesInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetAlliancesNamesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /alliances/names/][%d] getAlliancesNamesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAlliancesNamesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesNamesServiceUnavailable creates a GetAlliancesNamesServiceUnavailable with default headers values
func NewGetAlliancesNamesServiceUnavailable() *GetAlliancesNamesServiceUnavailable {
	return &GetAlliancesNamesServiceUnavailable{}
}

/*GetAlliancesNamesServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetAlliancesNamesServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetAlliancesNamesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /alliances/names/][%d] getAlliancesNamesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAlliancesNamesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlliancesNamesGatewayTimeout creates a GetAlliancesNamesGatewayTimeout with default headers values
func NewGetAlliancesNamesGatewayTimeout() *GetAlliancesNamesGatewayTimeout {
	return &GetAlliancesNamesGatewayTimeout{}
}

/*GetAlliancesNamesGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetAlliancesNamesGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetAlliancesNamesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /alliances/names/][%d] getAlliancesNamesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAlliancesNamesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
