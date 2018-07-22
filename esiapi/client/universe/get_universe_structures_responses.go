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

// GetUniverseStructuresReader is a Reader for the GetUniverseStructures structure.
type GetUniverseStructuresReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUniverseStructuresReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUniverseStructuresOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetUniverseStructuresNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetUniverseStructuresBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetUniverseStructuresEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetUniverseStructuresInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetUniverseStructuresServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetUniverseStructuresGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUniverseStructuresOK creates a GetUniverseStructuresOK with default headers values
func NewGetUniverseStructuresOK() *GetUniverseStructuresOK {
	return &GetUniverseStructuresOK{}
}

/*GetUniverseStructuresOK handles this case with default header values.

List of public structure IDs
*/
type GetUniverseStructuresOK struct {
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

	Payload []*int64
}

func (o *GetUniverseStructuresOK) Error() string {
	return fmt.Sprintf("[GET /universe/structures/][%d] getUniverseStructuresOK  %+v", 200, o.Payload)
}

func (o *GetUniverseStructuresOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseStructuresNotModified creates a GetUniverseStructuresNotModified with default headers values
func NewGetUniverseStructuresNotModified() *GetUniverseStructuresNotModified {
	return &GetUniverseStructuresNotModified{}
}

/*GetUniverseStructuresNotModified handles this case with default header values.

Not modified
*/
type GetUniverseStructuresNotModified struct {
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

func (o *GetUniverseStructuresNotModified) Error() string {
	return fmt.Sprintf("[GET /universe/structures/][%d] getUniverseStructuresNotModified ", 304)
}

func (o *GetUniverseStructuresNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetUniverseStructuresBadRequest creates a GetUniverseStructuresBadRequest with default headers values
func NewGetUniverseStructuresBadRequest() *GetUniverseStructuresBadRequest {
	return &GetUniverseStructuresBadRequest{}
}

/*GetUniverseStructuresBadRequest handles this case with default header values.

Bad request
*/
type GetUniverseStructuresBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetUniverseStructuresBadRequest) Error() string {
	return fmt.Sprintf("[GET /universe/structures/][%d] getUniverseStructuresBadRequest  %+v", 400, o.Payload)
}

func (o *GetUniverseStructuresBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStructuresEnhanceYourCalm creates a GetUniverseStructuresEnhanceYourCalm with default headers values
func NewGetUniverseStructuresEnhanceYourCalm() *GetUniverseStructuresEnhanceYourCalm {
	return &GetUniverseStructuresEnhanceYourCalm{}
}

/*GetUniverseStructuresEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetUniverseStructuresEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetUniverseStructuresEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /universe/structures/][%d] getUniverseStructuresEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetUniverseStructuresEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStructuresInternalServerError creates a GetUniverseStructuresInternalServerError with default headers values
func NewGetUniverseStructuresInternalServerError() *GetUniverseStructuresInternalServerError {
	return &GetUniverseStructuresInternalServerError{}
}

/*GetUniverseStructuresInternalServerError handles this case with default header values.

Internal server error
*/
type GetUniverseStructuresInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetUniverseStructuresInternalServerError) Error() string {
	return fmt.Sprintf("[GET /universe/structures/][%d] getUniverseStructuresInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUniverseStructuresInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStructuresServiceUnavailable creates a GetUniverseStructuresServiceUnavailable with default headers values
func NewGetUniverseStructuresServiceUnavailable() *GetUniverseStructuresServiceUnavailable {
	return &GetUniverseStructuresServiceUnavailable{}
}

/*GetUniverseStructuresServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetUniverseStructuresServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetUniverseStructuresServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /universe/structures/][%d] getUniverseStructuresServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUniverseStructuresServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUniverseStructuresGatewayTimeout creates a GetUniverseStructuresGatewayTimeout with default headers values
func NewGetUniverseStructuresGatewayTimeout() *GetUniverseStructuresGatewayTimeout {
	return &GetUniverseStructuresGatewayTimeout{}
}

/*GetUniverseStructuresGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetUniverseStructuresGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetUniverseStructuresGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /universe/structures/][%d] getUniverseStructuresGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUniverseStructuresGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}