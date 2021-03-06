// Code generated by go-swagger; DO NOT EDIT.

package corporation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/eve-tmarg/esiapi/models"
)

// GetCorporationsCorporationIDMembertrackingReader is a Reader for the GetCorporationsCorporationIDMembertracking structure.
type GetCorporationsCorporationIDMembertrackingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCorporationsCorporationIDMembertrackingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCorporationsCorporationIDMembertrackingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetCorporationsCorporationIDMembertrackingNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetCorporationsCorporationIDMembertrackingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetCorporationsCorporationIDMembertrackingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetCorporationsCorporationIDMembertrackingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetCorporationsCorporationIDMembertrackingEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetCorporationsCorporationIDMembertrackingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetCorporationsCorporationIDMembertrackingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetCorporationsCorporationIDMembertrackingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCorporationsCorporationIDMembertrackingOK creates a GetCorporationsCorporationIDMembertrackingOK with default headers values
func NewGetCorporationsCorporationIDMembertrackingOK() *GetCorporationsCorporationIDMembertrackingOK {
	return &GetCorporationsCorporationIDMembertrackingOK{}
}

/*GetCorporationsCorporationIDMembertrackingOK handles this case with default header values.

List of member character IDs
*/
type GetCorporationsCorporationIDMembertrackingOK struct {
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

	Payload []*models.GetCorporationsCorporationIDMembertrackingOKBodyItems
}

func (o *GetCorporationsCorporationIDMembertrackingOK) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/membertracking/][%d] getCorporationsCorporationIdMembertrackingOK  %+v", 200, o.Payload)
}

func (o *GetCorporationsCorporationIDMembertrackingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDMembertrackingNotModified creates a GetCorporationsCorporationIDMembertrackingNotModified with default headers values
func NewGetCorporationsCorporationIDMembertrackingNotModified() *GetCorporationsCorporationIDMembertrackingNotModified {
	return &GetCorporationsCorporationIDMembertrackingNotModified{}
}

/*GetCorporationsCorporationIDMembertrackingNotModified handles this case with default header values.

Not modified
*/
type GetCorporationsCorporationIDMembertrackingNotModified struct {
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

func (o *GetCorporationsCorporationIDMembertrackingNotModified) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/membertracking/][%d] getCorporationsCorporationIdMembertrackingNotModified ", 304)
}

func (o *GetCorporationsCorporationIDMembertrackingNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetCorporationsCorporationIDMembertrackingBadRequest creates a GetCorporationsCorporationIDMembertrackingBadRequest with default headers values
func NewGetCorporationsCorporationIDMembertrackingBadRequest() *GetCorporationsCorporationIDMembertrackingBadRequest {
	return &GetCorporationsCorporationIDMembertrackingBadRequest{}
}

/*GetCorporationsCorporationIDMembertrackingBadRequest handles this case with default header values.

Bad request
*/
type GetCorporationsCorporationIDMembertrackingBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetCorporationsCorporationIDMembertrackingBadRequest) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/membertracking/][%d] getCorporationsCorporationIdMembertrackingBadRequest  %+v", 400, o.Payload)
}

func (o *GetCorporationsCorporationIDMembertrackingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembertrackingUnauthorized creates a GetCorporationsCorporationIDMembertrackingUnauthorized with default headers values
func NewGetCorporationsCorporationIDMembertrackingUnauthorized() *GetCorporationsCorporationIDMembertrackingUnauthorized {
	return &GetCorporationsCorporationIDMembertrackingUnauthorized{}
}

/*GetCorporationsCorporationIDMembertrackingUnauthorized handles this case with default header values.

Unauthorized
*/
type GetCorporationsCorporationIDMembertrackingUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *GetCorporationsCorporationIDMembertrackingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/membertracking/][%d] getCorporationsCorporationIdMembertrackingUnauthorized  %+v", 401, o.Payload)
}

func (o *GetCorporationsCorporationIDMembertrackingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembertrackingForbidden creates a GetCorporationsCorporationIDMembertrackingForbidden with default headers values
func NewGetCorporationsCorporationIDMembertrackingForbidden() *GetCorporationsCorporationIDMembertrackingForbidden {
	return &GetCorporationsCorporationIDMembertrackingForbidden{}
}

/*GetCorporationsCorporationIDMembertrackingForbidden handles this case with default header values.

Forbidden
*/
type GetCorporationsCorporationIDMembertrackingForbidden struct {
	Payload *models.Forbidden
}

func (o *GetCorporationsCorporationIDMembertrackingForbidden) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/membertracking/][%d] getCorporationsCorporationIdMembertrackingForbidden  %+v", 403, o.Payload)
}

func (o *GetCorporationsCorporationIDMembertrackingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembertrackingEnhanceYourCalm creates a GetCorporationsCorporationIDMembertrackingEnhanceYourCalm with default headers values
func NewGetCorporationsCorporationIDMembertrackingEnhanceYourCalm() *GetCorporationsCorporationIDMembertrackingEnhanceYourCalm {
	return &GetCorporationsCorporationIDMembertrackingEnhanceYourCalm{}
}

/*GetCorporationsCorporationIDMembertrackingEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetCorporationsCorporationIDMembertrackingEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetCorporationsCorporationIDMembertrackingEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/membertracking/][%d] getCorporationsCorporationIdMembertrackingEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetCorporationsCorporationIDMembertrackingEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembertrackingInternalServerError creates a GetCorporationsCorporationIDMembertrackingInternalServerError with default headers values
func NewGetCorporationsCorporationIDMembertrackingInternalServerError() *GetCorporationsCorporationIDMembertrackingInternalServerError {
	return &GetCorporationsCorporationIDMembertrackingInternalServerError{}
}

/*GetCorporationsCorporationIDMembertrackingInternalServerError handles this case with default header values.

Internal server error
*/
type GetCorporationsCorporationIDMembertrackingInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetCorporationsCorporationIDMembertrackingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/membertracking/][%d] getCorporationsCorporationIdMembertrackingInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCorporationsCorporationIDMembertrackingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembertrackingServiceUnavailable creates a GetCorporationsCorporationIDMembertrackingServiceUnavailable with default headers values
func NewGetCorporationsCorporationIDMembertrackingServiceUnavailable() *GetCorporationsCorporationIDMembertrackingServiceUnavailable {
	return &GetCorporationsCorporationIDMembertrackingServiceUnavailable{}
}

/*GetCorporationsCorporationIDMembertrackingServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetCorporationsCorporationIDMembertrackingServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetCorporationsCorporationIDMembertrackingServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/membertracking/][%d] getCorporationsCorporationIdMembertrackingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetCorporationsCorporationIDMembertrackingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCorporationsCorporationIDMembertrackingGatewayTimeout creates a GetCorporationsCorporationIDMembertrackingGatewayTimeout with default headers values
func NewGetCorporationsCorporationIDMembertrackingGatewayTimeout() *GetCorporationsCorporationIDMembertrackingGatewayTimeout {
	return &GetCorporationsCorporationIDMembertrackingGatewayTimeout{}
}

/*GetCorporationsCorporationIDMembertrackingGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetCorporationsCorporationIDMembertrackingGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetCorporationsCorporationIDMembertrackingGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /corporations/{corporation_id}/membertracking/][%d] getCorporationsCorporationIdMembertrackingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetCorporationsCorporationIDMembertrackingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
