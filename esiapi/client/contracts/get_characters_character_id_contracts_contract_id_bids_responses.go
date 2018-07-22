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

// GetCharactersCharacterIDContractsContractIDBidsReader is a Reader for the GetCharactersCharacterIDContractsContractIDBids structure.
type GetCharactersCharacterIDContractsContractIDBidsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDContractsContractIDBidsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDContractsContractIDBidsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCharactersCharacterIDContractsContractIDBidsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCharactersCharacterIDContractsContractIDBidsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCharactersCharacterIDContractsContractIDBidsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCharactersCharacterIDContractsContractIDBidsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetCharactersCharacterIDContractsContractIDBidsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCharactersCharacterIDContractsContractIDBidsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDContractsContractIDBidsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCharactersCharacterIDContractsContractIDBidsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCharactersCharacterIDContractsContractIDBidsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDContractsContractIDBidsOK creates a GetCharactersCharacterIDContractsContractIDBidsOK with default headers values
func NewGetCharactersCharacterIDContractsContractIDBidsOK() *GetCharactersCharacterIDContractsContractIDBidsOK {
	return &GetCharactersCharacterIDContractsContractIDBidsOK{}
}

/*GetCharactersCharacterIDContractsContractIDBidsOK handles this case with default header values.

A list of bids
*/
type GetCharactersCharacterIDContractsContractIDBidsOK struct {
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

	Payload []*models.GetCharactersCharacterIDContractsContractIDBidsOKBodyItems
}

func (o *GetCharactersCharacterIDContractsContractIDBidsOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/{contract_id}/bids/][%d] getCharactersCharacterIdContractsContractIdBidsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDBidsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDContractsContractIDBidsNotModified creates a GetCharactersCharacterIDContractsContractIDBidsNotModified with default headers values
func NewGetCharactersCharacterIDContractsContractIDBidsNotModified() *GetCharactersCharacterIDContractsContractIDBidsNotModified {
	return &GetCharactersCharacterIDContractsContractIDBidsNotModified{}
}

/*GetCharactersCharacterIDContractsContractIDBidsNotModified handles this case with default header values.

Not modified
*/
type GetCharactersCharacterIDContractsContractIDBidsNotModified struct {
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

func (o *GetCharactersCharacterIDContractsContractIDBidsNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/{contract_id}/bids/][%d] getCharactersCharacterIdContractsContractIdBidsNotModified ", 304)
}

func (o *GetCharactersCharacterIDContractsContractIDBidsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDContractsContractIDBidsBadRequest creates a GetCharactersCharacterIDContractsContractIDBidsBadRequest with default headers values
func NewGetCharactersCharacterIDContractsContractIDBidsBadRequest() *GetCharactersCharacterIDContractsContractIDBidsBadRequest {
	return &GetCharactersCharacterIDContractsContractIDBidsBadRequest{}
}

/*GetCharactersCharacterIDContractsContractIDBidsBadRequest handles this case with default header values.

Bad request
*/
type GetCharactersCharacterIDContractsContractIDBidsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDContractsContractIDBidsBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/{contract_id}/bids/][%d] getCharactersCharacterIdContractsContractIdBidsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDBidsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsContractIDBidsUnauthorized creates a GetCharactersCharacterIDContractsContractIDBidsUnauthorized with default headers values
func NewGetCharactersCharacterIDContractsContractIDBidsUnauthorized() *GetCharactersCharacterIDContractsContractIDBidsUnauthorized {
	return &GetCharactersCharacterIDContractsContractIDBidsUnauthorized{}
}

/*GetCharactersCharacterIDContractsContractIDBidsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCharactersCharacterIDContractsContractIDBidsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDContractsContractIDBidsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/{contract_id}/bids/][%d] getCharactersCharacterIdContractsContractIdBidsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDBidsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsContractIDBidsForbidden creates a GetCharactersCharacterIDContractsContractIDBidsForbidden with default headers values
func NewGetCharactersCharacterIDContractsContractIDBidsForbidden() *GetCharactersCharacterIDContractsContractIDBidsForbidden {
	return &GetCharactersCharacterIDContractsContractIDBidsForbidden{}
}

/*GetCharactersCharacterIDContractsContractIDBidsForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDContractsContractIDBidsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDContractsContractIDBidsForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/{contract_id}/bids/][%d] getCharactersCharacterIdContractsContractIdBidsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDBidsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsContractIDBidsNotFound creates a GetCharactersCharacterIDContractsContractIDBidsNotFound with default headers values
func NewGetCharactersCharacterIDContractsContractIDBidsNotFound() *GetCharactersCharacterIDContractsContractIDBidsNotFound {
	return &GetCharactersCharacterIDContractsContractIDBidsNotFound{}
}

/*GetCharactersCharacterIDContractsContractIDBidsNotFound handles this case with default header values.

Not found
*/
type GetCharactersCharacterIDContractsContractIDBidsNotFound struct {
	Payload *models.GetCharactersCharacterIDContractsContractIDBidsNotFoundBody
}

func (o *GetCharactersCharacterIDContractsContractIDBidsNotFound) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/{contract_id}/bids/][%d] getCharactersCharacterIdContractsContractIdBidsNotFound  %+v", 404, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDBidsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCharactersCharacterIDContractsContractIDBidsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsContractIDBidsEnhanceYourCalm creates a GetCharactersCharacterIDContractsContractIDBidsEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDContractsContractIDBidsEnhanceYourCalm() *GetCharactersCharacterIDContractsContractIDBidsEnhanceYourCalm {
	return &GetCharactersCharacterIDContractsContractIDBidsEnhanceYourCalm{}
}

/*GetCharactersCharacterIDContractsContractIDBidsEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCharactersCharacterIDContractsContractIDBidsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDContractsContractIDBidsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/{contract_id}/bids/][%d] getCharactersCharacterIdContractsContractIdBidsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDBidsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsContractIDBidsInternalServerError creates a GetCharactersCharacterIDContractsContractIDBidsInternalServerError with default headers values
func NewGetCharactersCharacterIDContractsContractIDBidsInternalServerError() *GetCharactersCharacterIDContractsContractIDBidsInternalServerError {
	return &GetCharactersCharacterIDContractsContractIDBidsInternalServerError{}
}

/*GetCharactersCharacterIDContractsContractIDBidsInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDContractsContractIDBidsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDContractsContractIDBidsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/{contract_id}/bids/][%d] getCharactersCharacterIdContractsContractIdBidsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDBidsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsContractIDBidsServiceUnavailable creates a GetCharactersCharacterIDContractsContractIDBidsServiceUnavailable with default headers values
func NewGetCharactersCharacterIDContractsContractIDBidsServiceUnavailable() *GetCharactersCharacterIDContractsContractIDBidsServiceUnavailable {
	return &GetCharactersCharacterIDContractsContractIDBidsServiceUnavailable{}
}

/*GetCharactersCharacterIDContractsContractIDBidsServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCharactersCharacterIDContractsContractIDBidsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDContractsContractIDBidsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/{contract_id}/bids/][%d] getCharactersCharacterIdContractsContractIdBidsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDBidsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDContractsContractIDBidsGatewayTimeout creates a GetCharactersCharacterIDContractsContractIDBidsGatewayTimeout with default headers values
func NewGetCharactersCharacterIDContractsContractIDBidsGatewayTimeout() *GetCharactersCharacterIDContractsContractIDBidsGatewayTimeout {
	return &GetCharactersCharacterIDContractsContractIDBidsGatewayTimeout{}
}

/*GetCharactersCharacterIDContractsContractIDBidsGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDContractsContractIDBidsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDContractsContractIDBidsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/contracts/{contract_id}/bids/][%d] getCharactersCharacterIdContractsContractIdBidsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDContractsContractIDBidsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
