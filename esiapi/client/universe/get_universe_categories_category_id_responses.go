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

// GetUniverseCategoriesCategoryIDReader is a Reader for the GetUniverseCategoriesCategoryID structure.
type GetUniverseCategoriesCategoryIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseCategoriesCategoryIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseCategoriesCategoryIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetUniverseCategoriesCategoryIDNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetUniverseCategoriesCategoryIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetUniverseCategoriesCategoryIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetUniverseCategoriesCategoryIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetUniverseCategoriesCategoryIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetUniverseCategoriesCategoryIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetUniverseCategoriesCategoryIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseCategoriesCategoryIDOK creates a GetUniverseCategoriesCategoryIDOK with default headers values
func NewGetUniverseCategoriesCategoryIDOK() *GetUniverseCategoriesCategoryIDOK {
	return &GetUniverseCategoriesCategoryIDOK{}
}

/*GetUniverseCategoriesCategoryIDOK handles this case with default header values.

Information about an item category
*/
type GetUniverseCategoriesCategoryIDOK struct {
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

	Payload *models.GetUniverseCategoriesCategoryIDOKBody
}

func (o *GetUniverseCategoriesCategoryIDOK) Error() string {
	return fmt.Sprintf("[GET /universe/categories/{category_id}/][%d] getUniverseCategoriesCategoryIdOK  %+v", 200, o.Payload)
}

func (o *GetUniverseCategoriesCategoryIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.GetUniverseCategoriesCategoryIDOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseCategoriesCategoryIDNotModified creates a GetUniverseCategoriesCategoryIDNotModified with default headers values
func NewGetUniverseCategoriesCategoryIDNotModified() *GetUniverseCategoriesCategoryIDNotModified {
	return &GetUniverseCategoriesCategoryIDNotModified{}
}

/*GetUniverseCategoriesCategoryIDNotModified handles this case with default header values.

Not modified
*/
type GetUniverseCategoriesCategoryIDNotModified struct {
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

func (o *GetUniverseCategoriesCategoryIDNotModified) Error() string {
	return fmt.Sprintf("[GET /universe/categories/{category_id}/][%d] getUniverseCategoriesCategoryIdNotModified ", 304)
}

func (o *GetUniverseCategoriesCategoryIDNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseCategoriesCategoryIDBadRequest creates a GetUniverseCategoriesCategoryIDBadRequest with default headers values
func NewGetUniverseCategoriesCategoryIDBadRequest() *GetUniverseCategoriesCategoryIDBadRequest {
	return &GetUniverseCategoriesCategoryIDBadRequest{}
}

/*GetUniverseCategoriesCategoryIDBadRequest handles this case with default header values.

Bad request
*/
type GetUniverseCategoriesCategoryIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetUniverseCategoriesCategoryIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /universe/categories/{category_id}/][%d] getUniverseCategoriesCategoryIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseCategoriesCategoryIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseCategoriesCategoryIDNotFound creates a GetUniverseCategoriesCategoryIDNotFound with default headers values
func NewGetUniverseCategoriesCategoryIDNotFound() *GetUniverseCategoriesCategoryIDNotFound {
	return &GetUniverseCategoriesCategoryIDNotFound{}
}

/*GetUniverseCategoriesCategoryIDNotFound handles this case with default header values.

Category not found
*/
type GetUniverseCategoriesCategoryIDNotFound struct {
	Payload *models.GetUniverseCategoriesCategoryIDNotFoundBody
}

func (o *GetUniverseCategoriesCategoryIDNotFound) Error() string {
	return fmt.Sprintf("[GET /universe/categories/{category_id}/][%d] getUniverseCategoriesCategoryIdNotFound  %+v", 404, o.Payload)
}

func (o *GetUniverseCategoriesCategoryIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetUniverseCategoriesCategoryIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseCategoriesCategoryIDEnhanceYourCalm creates a GetUniverseCategoriesCategoryIDEnhanceYourCalm with default headers values
func NewGetUniverseCategoriesCategoryIDEnhanceYourCalm() *GetUniverseCategoriesCategoryIDEnhanceYourCalm {
	return &GetUniverseCategoriesCategoryIDEnhanceYourCalm{}
}

/*GetUniverseCategoriesCategoryIDEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetUniverseCategoriesCategoryIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetUniverseCategoriesCategoryIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /universe/categories/{category_id}/][%d] getUniverseCategoriesCategoryIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseCategoriesCategoryIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseCategoriesCategoryIDInternalServerError creates a GetUniverseCategoriesCategoryIDInternalServerError with default headers values
func NewGetUniverseCategoriesCategoryIDInternalServerError() *GetUniverseCategoriesCategoryIDInternalServerError {
	return &GetUniverseCategoriesCategoryIDInternalServerError{}
}

/*GetUniverseCategoriesCategoryIDInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseCategoriesCategoryIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetUniverseCategoriesCategoryIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/categories/{category_id}/][%d] getUniverseCategoriesCategoryIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseCategoriesCategoryIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseCategoriesCategoryIDServiceUnavailable creates a GetUniverseCategoriesCategoryIDServiceUnavailable with default headers values
func NewGetUniverseCategoriesCategoryIDServiceUnavailable() *GetUniverseCategoriesCategoryIDServiceUnavailable {
	return &GetUniverseCategoriesCategoryIDServiceUnavailable{}
}

/*GetUniverseCategoriesCategoryIDServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetUniverseCategoriesCategoryIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetUniverseCategoriesCategoryIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /universe/categories/{category_id}/][%d] getUniverseCategoriesCategoryIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseCategoriesCategoryIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseCategoriesCategoryIDGatewayTimeout creates a GetUniverseCategoriesCategoryIDGatewayTimeout with default headers values
func NewGetUniverseCategoriesCategoryIDGatewayTimeout() *GetUniverseCategoriesCategoryIDGatewayTimeout {
	return &GetUniverseCategoriesCategoryIDGatewayTimeout{}
}

/*GetUniverseCategoriesCategoryIDGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetUniverseCategoriesCategoryIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetUniverseCategoriesCategoryIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /universe/categories/{category_id}/][%d] getUniverseCategoriesCategoryIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseCategoriesCategoryIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
