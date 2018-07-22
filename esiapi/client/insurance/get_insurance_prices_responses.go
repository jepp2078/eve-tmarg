// Code generated by go-swagger; DO NOT EDIT.

package insurance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/eve-tmarg/esiapi/models"
)

// GetInsurancePricesReader is a Reader for the GetInsurancePrices structure.
type GetInsurancePricesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInsurancePricesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInsurancePricesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetInsurancePricesNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetInsurancePricesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetInsurancePricesEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetInsurancePricesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetInsurancePricesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetInsurancePricesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInsurancePricesOK creates a GetInsurancePricesOK with default headers values
func NewGetInsurancePricesOK() *GetInsurancePricesOK {
	return &GetInsurancePricesOK{}
}

/*GetInsurancePricesOK handles this case with default header values.

A list of insurance levels for all ship types
*/
type GetInsurancePricesOK struct {
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

	Payload []*models.GetInsurancePricesOKBodyItems
}

func (o *GetInsurancePricesOK) Error() string {
	return fmt.Sprintf("[GET /insurance/prices/][%d] getInsurancePricesOK  %+v", 200, o.Payload)
}

func (o *GetInsurancePricesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetInsurancePricesNotModified creates a GetInsurancePricesNotModified with default headers values
func NewGetInsurancePricesNotModified() *GetInsurancePricesNotModified {
	return &GetInsurancePricesNotModified{}
}

/*GetInsurancePricesNotModified handles this case with default header values.

Not modified
*/
type GetInsurancePricesNotModified struct {
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

func (o *GetInsurancePricesNotModified) Error() string {
	return fmt.Sprintf("[GET /insurance/prices/][%d] getInsurancePricesNotModified ", 304)
}

func (o *GetInsurancePricesNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetInsurancePricesBadRequest creates a GetInsurancePricesBadRequest with default headers values
func NewGetInsurancePricesBadRequest() *GetInsurancePricesBadRequest {
	return &GetInsurancePricesBadRequest{}
}

/*GetInsurancePricesBadRequest handles this case with default header values.

Bad request
*/
type GetInsurancePricesBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetInsurancePricesBadRequest) Error() string {
	return fmt.Sprintf("[GET /insurance/prices/][%d] getInsurancePricesBadRequest  %+v", 400, o.Payload)
}

func (o *GetInsurancePricesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInsurancePricesEnhanceYourCalm creates a GetInsurancePricesEnhanceYourCalm with default headers values
func NewGetInsurancePricesEnhanceYourCalm() *GetInsurancePricesEnhanceYourCalm {
	return &GetInsurancePricesEnhanceYourCalm{}
}

/*GetInsurancePricesEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetInsurancePricesEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetInsurancePricesEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /insurance/prices/][%d] getInsurancePricesEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetInsurancePricesEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInsurancePricesInternalServerError creates a GetInsurancePricesInternalServerError with default headers values
func NewGetInsurancePricesInternalServerError() *GetInsurancePricesInternalServerError {
	return &GetInsurancePricesInternalServerError{}
}

/*GetInsurancePricesInternalServerError handles this case with default header values.

Internal server error
*/
type GetInsurancePricesInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetInsurancePricesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /insurance/prices/][%d] getInsurancePricesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInsurancePricesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInsurancePricesServiceUnavailable creates a GetInsurancePricesServiceUnavailable with default headers values
func NewGetInsurancePricesServiceUnavailable() *GetInsurancePricesServiceUnavailable {
	return &GetInsurancePricesServiceUnavailable{}
}

/*GetInsurancePricesServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetInsurancePricesServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetInsurancePricesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /insurance/prices/][%d] getInsurancePricesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetInsurancePricesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInsurancePricesGatewayTimeout creates a GetInsurancePricesGatewayTimeout with default headers values
func NewGetInsurancePricesGatewayTimeout() *GetInsurancePricesGatewayTimeout {
	return &GetInsurancePricesGatewayTimeout{}
}

/*GetInsurancePricesGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetInsurancePricesGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetInsurancePricesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /insurance/prices/][%d] getInsurancePricesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetInsurancePricesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
