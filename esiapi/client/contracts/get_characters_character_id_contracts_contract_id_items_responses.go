// Code generated by go-swagger; DO NOT EDIT.

package contracts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/eve-tmarg/esiapi/models"
)

// GetCharactersCharacterIDContractsContractIDItemsReader is a Reader for the GetCharactersCharacterIDContractsContractIDItems structure.
type GetCharactersCharacterIDContractsContractIDItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDContractsContractIDItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDContractsContractIDItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCharactersCharacterIDContractsContractIDItemsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCharactersCharacterIDContractsContractIDItemsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCharactersCharacterIDContractsContractIDItemsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCharactersCharacterIDContractsContractIDItemsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetCharactersCharacterIDContractsContractIDItemsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDContractsContractIDItemsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCharactersCharacterIDContractsContractIDItemsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCharactersCharacterIDContractsContractIDItemsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDContractsContractIDItemsOK creates a GetCharactersCharacterIDContractsContractIDItemsOK with default headers values
func NewGetCharactersCharacterIDContractsContractIDItemsOK() *GetCharactersCharacterIDContractsContractIDItemsOK {
	return &GetCharactersCharacterIDContractsContractIDItemsOK{}
}

/*GetCharactersCharacterIDContractsContractIDItemsOK handles this case with default header values.

A list of items in this contract
*/
type GetCharactersCharacterIDContractsContractIDItemsOK struct {
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

	Payload []*models.GetCharactersCharacterIDContractsContractIDItemsOKBodyItems
}

func (o *GetCharactersCharacterIDContractsContractIDItemsOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDContractsContractIDItemsNotModified creates a GetCharactersCharacterIDContractsContractIDItemsNotModified with default headers values
func NewGetCharactersCharacterIDContractsContractIDItemsNotModified() *GetCharactersCharacterIDContractsContractIDItemsNotModified {
	return &GetCharactersCharacterIDContractsContractIDItemsNotModified{}
}

/*GetCharactersCharacterIDContractsContractIDItemsNotModified handles this case with default header values.

Not modified
*/
type GetCharactersCharacterIDContractsContractIDItemsNotModified struct {
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

func (o *GetCharactersCharacterIDContractsContractIDItemsNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsNotModified ", 304)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDContractsContractIDItemsBadRequest creates a GetCharactersCharacterIDContractsContractIDItemsBadRequest with default headers values
func NewGetCharactersCharacterIDContractsContractIDItemsBadRequest() *GetCharactersCharacterIDContractsContractIDItemsBadRequest {
	return &GetCharactersCharacterIDContractsContractIDItemsBadRequest{}
}

/*GetCharactersCharacterIDContractsContractIDItemsBadRequest handles this case with default header values.

Bad request
*/
type GetCharactersCharacterIDContractsContractIDItemsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDContractsContractIDItemsBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsContractIDItemsUnauthorized creates a GetCharactersCharacterIDContractsContractIDItemsUnauthorized with default headers values
func NewGetCharactersCharacterIDContractsContractIDItemsUnauthorized() *GetCharactersCharacterIDContractsContractIDItemsUnauthorized {
	return &GetCharactersCharacterIDContractsContractIDItemsUnauthorized{}
}

/*GetCharactersCharacterIDContractsContractIDItemsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCharactersCharacterIDContractsContractIDItemsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDContractsContractIDItemsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsContractIDItemsForbidden creates a GetCharactersCharacterIDContractsContractIDItemsForbidden with default headers values
func NewGetCharactersCharacterIDContractsContractIDItemsForbidden() *GetCharactersCharacterIDContractsContractIDItemsForbidden {
	return &GetCharactersCharacterIDContractsContractIDItemsForbidden{}
}

/*GetCharactersCharacterIDContractsContractIDItemsForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDContractsContractIDItemsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDContractsContractIDItemsForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsContractIDItemsNotFound creates a GetCharactersCharacterIDContractsContractIDItemsNotFound with default headers values
func NewGetCharactersCharacterIDContractsContractIDItemsNotFound() *GetCharactersCharacterIDContractsContractIDItemsNotFound {
	return &GetCharactersCharacterIDContractsContractIDItemsNotFound{}
}

/*GetCharactersCharacterIDContractsContractIDItemsNotFound handles this case with default header values.

Not found
*/
type GetCharactersCharacterIDContractsContractIDItemsNotFound struct {
	Payload *models.GetCharactersCharacterIDContractsContractIDItemsNotFoundBody
}

func (o *GetCharactersCharacterIDContractsContractIDItemsNotFound) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsNotFound  %+v", 404, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDContractsContractIDItemsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm creates a GetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm() *GetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm {
	return &GetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm{}
}

/*GetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsContractIDItemsInternalServerError creates a GetCharactersCharacterIDContractsContractIDItemsInternalServerError with default headers values
func NewGetCharactersCharacterIDContractsContractIDItemsInternalServerError() *GetCharactersCharacterIDContractsContractIDItemsInternalServerError {
	return &GetCharactersCharacterIDContractsContractIDItemsInternalServerError{}
}

/*GetCharactersCharacterIDContractsContractIDItemsInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDContractsContractIDItemsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDContractsContractIDItemsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsContractIDItemsServiceUnavailable creates a GetCharactersCharacterIDContractsContractIDItemsServiceUnavailable with default headers values
func NewGetCharactersCharacterIDContractsContractIDItemsServiceUnavailable() *GetCharactersCharacterIDContractsContractIDItemsServiceUnavailable {
	return &GetCharactersCharacterIDContractsContractIDItemsServiceUnavailable{}
}

/*GetCharactersCharacterIDContractsContractIDItemsServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCharactersCharacterIDContractsContractIDItemsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDContractsContractIDItemsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsContractIDItemsGatewayTimeout creates a GetCharactersCharacterIDContractsContractIDItemsGatewayTimeout with default headers values
func NewGetCharactersCharacterIDContractsContractIDItemsGatewayTimeout() *GetCharactersCharacterIDContractsContractIDItemsGatewayTimeout {
	return &GetCharactersCharacterIDContractsContractIDItemsGatewayTimeout{}
}

/*GetCharactersCharacterIDContractsContractIDItemsGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDContractsContractIDItemsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDContractsContractIDItemsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/{contract_id}/items/][%d] getCharactersCharacterIdContractsContractIdItemsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDItemsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
