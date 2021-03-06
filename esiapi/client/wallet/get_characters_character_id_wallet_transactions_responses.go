// Code generated by go-swagger; DO NOT EDIT.

package wallet

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/eve-tmarg/esiapi/models"
)

// GetCharactersCharacterIDWalletTransactionsReader is a Reader for the GetCharactersCharacterIDWalletTransactions structure.
type GetCharactersCharacterIDWalletTransactionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCharactersCharacterIDWalletTransactionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCharactersCharacterIDWalletTransactionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCharactersCharacterIDWalletTransactionsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCharactersCharacterIDWalletTransactionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCharactersCharacterIDWalletTransactionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCharactersCharacterIDWalletTransactionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCharactersCharacterIDWalletTransactionsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCharactersCharacterIDWalletTransactionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCharactersCharacterIDWalletTransactionsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCharactersCharacterIDWalletTransactionsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCharactersCharacterIDWalletTransactionsOK creates a GetCharactersCharacterIDWalletTransactionsOK with default headers values
func NewGetCharactersCharacterIDWalletTransactionsOK() *GetCharactersCharacterIDWalletTransactionsOK {
	return &GetCharactersCharacterIDWalletTransactionsOK{}
}

/*GetCharactersCharacterIDWalletTransactionsOK handles this case with default header values.

Wallet transactions
*/
type GetCharactersCharacterIDWalletTransactionsOK struct {
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

	Payload []*models.GetCharactersCharacterIDWalletTransactionsOKBodyItems
}

func (o *GetCharactersCharacterIDWalletTransactionsOK) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsOK  %+v", 200, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDWalletTransactionsNotModified creates a GetCharactersCharacterIDWalletTransactionsNotModified with default headers values
func NewGetCharactersCharacterIDWalletTransactionsNotModified() *GetCharactersCharacterIDWalletTransactionsNotModified {
	return &GetCharactersCharacterIDWalletTransactionsNotModified{}
}

/*GetCharactersCharacterIDWalletTransactionsNotModified handles this case with default header values.

Not modified
*/
type GetCharactersCharacterIDWalletTransactionsNotModified struct {
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

func (o *GetCharactersCharacterIDWalletTransactionsNotModified) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsNotModified ", 304)
}

func (o *GetCharactersCharacterIDWalletTransactionsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCharactersCharacterIDWalletTransactionsBadRequest creates a GetCharactersCharacterIDWalletTransactionsBadRequest with default headers values
func NewGetCharactersCharacterIDWalletTransactionsBadRequest() *GetCharactersCharacterIDWalletTransactionsBadRequest {
	return &GetCharactersCharacterIDWalletTransactionsBadRequest{}
}

/*GetCharactersCharacterIDWalletTransactionsBadRequest handles this case with default header values.

Bad request
*/
type GetCharactersCharacterIDWalletTransactionsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCharactersCharacterIDWalletTransactionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDWalletTransactionsUnauthorized creates a GetCharactersCharacterIDWalletTransactionsUnauthorized with default headers values
func NewGetCharactersCharacterIDWalletTransactionsUnauthorized() *GetCharactersCharacterIDWalletTransactionsUnauthorized {
	return &GetCharactersCharacterIDWalletTransactionsUnauthorized{}
}

/*GetCharactersCharacterIDWalletTransactionsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCharactersCharacterIDWalletTransactionsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCharactersCharacterIDWalletTransactionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDWalletTransactionsForbidden creates a GetCharactersCharacterIDWalletTransactionsForbidden with default headers values
func NewGetCharactersCharacterIDWalletTransactionsForbidden() *GetCharactersCharacterIDWalletTransactionsForbidden {
	return &GetCharactersCharacterIDWalletTransactionsForbidden{}
}

/*GetCharactersCharacterIDWalletTransactionsForbidden handles this case with default header values.

Forbidden
*/
type GetCharactersCharacterIDWalletTransactionsForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCharactersCharacterIDWalletTransactionsForbidden) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsForbidden  %+v", 403, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDWalletTransactionsEnhanceYourCalm creates a GetCharactersCharacterIDWalletTransactionsEnhanceYourCalm with default headers values
func NewGetCharactersCharacterIDWalletTransactionsEnhanceYourCalm() *GetCharactersCharacterIDWalletTransactionsEnhanceYourCalm {
	return &GetCharactersCharacterIDWalletTransactionsEnhanceYourCalm{}
}

/*GetCharactersCharacterIDWalletTransactionsEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCharactersCharacterIDWalletTransactionsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCharactersCharacterIDWalletTransactionsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDWalletTransactionsInternalServerError creates a GetCharactersCharacterIDWalletTransactionsInternalServerError with default headers values
func NewGetCharactersCharacterIDWalletTransactionsInternalServerError() *GetCharactersCharacterIDWalletTransactionsInternalServerError {
	return &GetCharactersCharacterIDWalletTransactionsInternalServerError{}
}

/*GetCharactersCharacterIDWalletTransactionsInternalServerError handles this case with default header values.

Internal server error
*/
type GetCharactersCharacterIDWalletTransactionsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCharactersCharacterIDWalletTransactionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDWalletTransactionsServiceUnavailable creates a GetCharactersCharacterIDWalletTransactionsServiceUnavailable with default headers values
func NewGetCharactersCharacterIDWalletTransactionsServiceUnavailable() *GetCharactersCharacterIDWalletTransactionsServiceUnavailable {
	return &GetCharactersCharacterIDWalletTransactionsServiceUnavailable{}
}

/*GetCharactersCharacterIDWalletTransactionsServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCharactersCharacterIDWalletTransactionsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCharactersCharacterIDWalletTransactionsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCharactersCharacterIDWalletTransactionsGatewayTimeout creates a GetCharactersCharacterIDWalletTransactionsGatewayTimeout with default headers values
func NewGetCharactersCharacterIDWalletTransactionsGatewayTimeout() *GetCharactersCharacterIDWalletTransactionsGatewayTimeout {
	return &GetCharactersCharacterIDWalletTransactionsGatewayTimeout{}
}

/*GetCharactersCharacterIDWalletTransactionsGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCharactersCharacterIDWalletTransactionsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCharactersCharacterIDWalletTransactionsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /characters/{character_id}/wallet/transactions/][%d] getCharactersCharacterIdWalletTransactionsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCharactersCharacterIDWalletTransactionsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
