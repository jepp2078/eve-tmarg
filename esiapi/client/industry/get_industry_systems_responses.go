// Code generated by go-swagger; DO NOT EDIT.

package industry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/eve-tmarg/esiapi/models"
)

// GetIndustrySystemsReader is a Reader for the GetIndustrySystems structure.
type GetIndustrySystemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIndustrySystemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetIndustrySystemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetIndustrySystemsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetIndustrySystemsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetIndustrySystemsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetIndustrySystemsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetIndustrySystemsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetIndustrySystemsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetIndustrySystemsOK creates a GetIndustrySystemsOK with default headers values
func NewGetIndustrySystemsOK() *GetIndustrySystemsOK {
	return &GetIndustrySystemsOK{}
}

/*GetIndustrySystemsOK handles this case with default header values.

A list of cost indicies
*/
type GetIndustrySystemsOK struct {
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

	Payload []*models.GetIndustrySystemsOKBodyItems
}

func (o *GetIndustrySystemsOK) Error() string {
	return fmt.Sprintf("[GET /industry/systems/][%d] getIndustrySystemsOK  %+v", 200, o.Payload)
}

func (o *GetIndustrySystemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIndustrySystemsNotModified creates a GetIndustrySystemsNotModified with default headers values
func NewGetIndustrySystemsNotModified() *GetIndustrySystemsNotModified {
	return &GetIndustrySystemsNotModified{}
}

/*GetIndustrySystemsNotModified handles this case with default header values.

Not modified
*/
type GetIndustrySystemsNotModified struct {
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

func (o *GetIndustrySystemsNotModified) Error() string {
	return fmt.Sprintf("[GET /industry/systems/][%d] getIndustrySystemsNotModified ", 304)
}

func (o *GetIndustrySystemsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetIndustrySystemsBadRequest creates a GetIndustrySystemsBadRequest with default headers values
func NewGetIndustrySystemsBadRequest() *GetIndustrySystemsBadRequest {
	return &GetIndustrySystemsBadRequest{}
}

/*GetIndustrySystemsBadRequest handles this case with default header values.

Bad request
*/
type GetIndustrySystemsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetIndustrySystemsBadRequest) Error() string {
	return fmt.Sprintf("[GET /industry/systems/][%d] getIndustrySystemsBadRequest  %+v", 400, o.Payload)
}

func (o *GetIndustrySystemsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIndustrySystemsEnhanceYourCalm creates a GetIndustrySystemsEnhanceYourCalm with default headers values
func NewGetIndustrySystemsEnhanceYourCalm() *GetIndustrySystemsEnhanceYourCalm {
	return &GetIndustrySystemsEnhanceYourCalm{}
}

/*GetIndustrySystemsEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetIndustrySystemsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetIndustrySystemsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /industry/systems/][%d] getIndustrySystemsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetIndustrySystemsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIndustrySystemsInternalServerError creates a GetIndustrySystemsInternalServerError with default headers values
func NewGetIndustrySystemsInternalServerError() *GetIndustrySystemsInternalServerError {
	return &GetIndustrySystemsInternalServerError{}
}

/*GetIndustrySystemsInternalServerError handles this case with default header values.

Internal server error
*/
type GetIndustrySystemsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetIndustrySystemsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /industry/systems/][%d] getIndustrySystemsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIndustrySystemsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIndustrySystemsServiceUnavailable creates a GetIndustrySystemsServiceUnavailable with default headers values
func NewGetIndustrySystemsServiceUnavailable() *GetIndustrySystemsServiceUnavailable {
	return &GetIndustrySystemsServiceUnavailable{}
}

/*GetIndustrySystemsServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetIndustrySystemsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetIndustrySystemsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /industry/systems/][%d] getIndustrySystemsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIndustrySystemsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIndustrySystemsGatewayTimeout creates a GetIndustrySystemsGatewayTimeout with default headers values
func NewGetIndustrySystemsGatewayTimeout() *GetIndustrySystemsGatewayTimeout {
	return &GetIndustrySystemsGatewayTimeout{}
}

/*GetIndustrySystemsGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetIndustrySystemsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetIndustrySystemsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /industry/systems/][%d] getIndustrySystemsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIndustrySystemsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
