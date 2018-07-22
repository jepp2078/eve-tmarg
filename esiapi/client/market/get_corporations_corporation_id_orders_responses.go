// Code generated by go-swagger; DO NOT EDIT.

package market

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

// GetCorporationsCorporationIDOrdersReader is a Reader for the GetCorporationsCorporationIDOrders structure.
type GetCorporationsCorporationIDOrdersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDOrdersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCorporationsCorporationIDOrdersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCorporationsCorporationIDOrdersNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCorporationsCorporationIDOrdersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCorporationsCorporationIDOrdersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCorporationsCorporationIDOrdersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCorporationsCorporationIDOrdersEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCorporationsCorporationIDOrdersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCorporationsCorporationIDOrdersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCorporationsCorporationIDOrdersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDOrdersOK creates a GetCorporationsCorporationIDOrdersOK with default headers values
func NewGetCorporationsCorporationIDOrdersOK() *GetCorporationsCorporationIDOrdersOK {
	return &GetCorporationsCorporationIDOrdersOK{
		XPages: 1,
	}
}

/*GetCorporationsCorporationIDOrdersOK handles this case with default header values.

A list of open market orders
*/
type GetCorporationsCorporationIDOrdersOK struct {
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

	Payload []*models.GetCorporationsCorporationIDOrdersOKBodyItems
}

func (o *GetCorporationsCorporationIDOrdersOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/][%d] getCorporationsCorporationIdOrdersOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDOrdersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDOrdersNotModified creates a GetCorporationsCorporationIDOrdersNotModified with default headers values
func NewGetCorporationsCorporationIDOrdersNotModified() *GetCorporationsCorporationIDOrdersNotModified {
	return &GetCorporationsCorporationIDOrdersNotModified{}
}

/*GetCorporationsCorporationIDOrdersNotModified handles this case with default header values.

Not modified
*/
type GetCorporationsCorporationIDOrdersNotModified struct {
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

func (o *GetCorporationsCorporationIDOrdersNotModified) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/][%d] getCorporationsCorporationIdOrdersNotModified ", 304)
}

func (o *GetCorporationsCorporationIDOrdersNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDOrdersBadRequest creates a GetCorporationsCorporationIDOrdersBadRequest with default headers values
func NewGetCorporationsCorporationIDOrdersBadRequest() *GetCorporationsCorporationIDOrdersBadRequest {
	return &GetCorporationsCorporationIDOrdersBadRequest{}
}

/*GetCorporationsCorporationIDOrdersBadRequest handles this case with default header values.

Bad request
*/
type GetCorporationsCorporationIDOrdersBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDOrdersBadRequest) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/][%d] getCorporationsCorporationIdOrdersBadRequest  %+v", 400, o.Payload)
}

func (o *GetCorporationsCorporationIDOrdersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDOrdersUnauthorized creates a GetCorporationsCorporationIDOrdersUnauthorized with default headers values
func NewGetCorporationsCorporationIDOrdersUnauthorized() *GetCorporationsCorporationIDOrdersUnauthorized {
	return &GetCorporationsCorporationIDOrdersUnauthorized{}
}

/*GetCorporationsCorporationIDOrdersUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDOrdersUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDOrdersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/][%d] getCorporationsCorporationIdOrdersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCorporationsCorporationIDOrdersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDOrdersForbidden creates a GetCorporationsCorporationIDOrdersForbidden with default headers values
func NewGetCorporationsCorporationIDOrdersForbidden() *GetCorporationsCorporationIDOrdersForbidden {
	return &GetCorporationsCorporationIDOrdersForbidden{}
}

/*GetCorporationsCorporationIDOrdersForbidden handles this case with default header values.

Forbidden
*/
type GetCorporationsCorporationIDOrdersForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDOrdersForbidden) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/][%d] getCorporationsCorporationIdOrdersForbidden  %+v", 403, o.Payload)
}

func (o *GetCorporationsCorporationIDOrdersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDOrdersEnhanceYourCalm creates a GetCorporationsCorporationIDOrdersEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDOrdersEnhanceYourCalm() *GetCorporationsCorporationIDOrdersEnhanceYourCalm {
	return &GetCorporationsCorporationIDOrdersEnhanceYourCalm{}
}

/*GetCorporationsCorporationIDOrdersEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCorporationsCorporationIDOrdersEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDOrdersEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/][%d] getCorporationsCorporationIdOrdersEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCorporationsCorporationIDOrdersEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDOrdersInternalServerError creates a GetCorporationsCorporationIDOrdersInternalServerError with default headers values
func NewGetCorporationsCorporationIDOrdersInternalServerError() *GetCorporationsCorporationIDOrdersInternalServerError {
	return &GetCorporationsCorporationIDOrdersInternalServerError{}
}

/*GetCorporationsCorporationIDOrdersInternalServerError handles this case with default header values.

Internal server error
*/
type GetCorporationsCorporationIDOrdersInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDOrdersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/][%d] getCorporationsCorporationIdOrdersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDOrdersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDOrdersServiceUnavailable creates a GetCorporationsCorporationIDOrdersServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDOrdersServiceUnavailable() *GetCorporationsCorporationIDOrdersServiceUnavailable {
	return &GetCorporationsCorporationIDOrdersServiceUnavailable{}
}

/*GetCorporationsCorporationIDOrdersServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDOrdersServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDOrdersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/][%d] getCorporationsCorporationIdOrdersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCorporationsCorporationIDOrdersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDOrdersGatewayTimeout creates a GetCorporationsCorporationIDOrdersGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDOrdersGatewayTimeout() *GetCorporationsCorporationIDOrdersGatewayTimeout {
	return &GetCorporationsCorporationIDOrdersGatewayTimeout{}
}

/*GetCorporationsCorporationIDOrdersGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDOrdersGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDOrdersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/orders/][%d] getCorporationsCorporationIdOrdersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCorporationsCorporationIDOrdersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
