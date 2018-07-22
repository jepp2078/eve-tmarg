// Code generated by go-swagger; DO NOT EDIT.

package sovereignty

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/eve-tmarg/esiapi/models"
)

// GetSovereigntyCampaignsReader is a Reader for the GetSovereigntyCampaigns structure.
type GetSovereigntyCampaignsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSovereigntyCampaignsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSovereigntyCampaignsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 304:
		result := NewGetSovereigntyCampaignsNotModified()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 400:
		result := NewGetSovereigntyCampaignsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewGetSovereigntyCampaignsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetSovereigntyCampaignsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewGetSovereigntyCampaignsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewGetSovereigntyCampaignsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSovereigntyCampaignsOK creates a GetSovereigntyCampaignsOK with default headers values
func NewGetSovereigntyCampaignsOK() *GetSovereigntyCampaignsOK {
	return &GetSovereigntyCampaignsOK{}
}

/*GetSovereigntyCampaignsOK handles this case with default header values.

A list of sovereignty campaigns
*/
type GetSovereigntyCampaignsOK struct {
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

	Payload []*models.GetSovereigntyCampaignsOKBodyItems
}

func (o *GetSovereigntyCampaignsOK) Error() string {
	return fmt.Sprintf("[GET /sovereignty/campaigns/][%d] getSovereigntyCampaignsOK  %+v", 200, o.Payload)
}

func (o *GetSovereigntyCampaignsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSovereigntyCampaignsNotModified creates a GetSovereigntyCampaignsNotModified with default headers values
func NewGetSovereigntyCampaignsNotModified() *GetSovereigntyCampaignsNotModified {
	return &GetSovereigntyCampaignsNotModified{}
}

/*GetSovereigntyCampaignsNotModified handles this case with default header values.

Not modified
*/
type GetSovereigntyCampaignsNotModified struct {
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

func (o *GetSovereigntyCampaignsNotModified) Error() string {
	return fmt.Sprintf("[GET /sovereignty/campaigns/][%d] getSovereigntyCampaignsNotModified ", 304)
}

func (o *GetSovereigntyCampaignsNotModified) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSovereigntyCampaignsBadRequest creates a GetSovereigntyCampaignsBadRequest with default headers values
func NewGetSovereigntyCampaignsBadRequest() *GetSovereigntyCampaignsBadRequest {
	return &GetSovereigntyCampaignsBadRequest{}
}

/*GetSovereigntyCampaignsBadRequest handles this case with default header values.

Bad request
*/
type GetSovereigntyCampaignsBadRequest struct {
	Payload *models.BadRequest
}

func (o *GetSovereigntyCampaignsBadRequest) Error() string {
	return fmt.Sprintf("[GET /sovereignty/campaigns/][%d] getSovereigntyCampaignsBadRequest  %+v", 400, o.Payload)
}

func (o *GetSovereigntyCampaignsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSovereigntyCampaignsEnhanceYourCalm creates a GetSovereigntyCampaignsEnhanceYourCalm with default headers values
func NewGetSovereigntyCampaignsEnhanceYourCalm() *GetSovereigntyCampaignsEnhanceYourCalm {
	return &GetSovereigntyCampaignsEnhanceYourCalm{}
}

/*GetSovereigntyCampaignsEnhanceYourCalm handles this case with default header values.

Error limited
*/
type GetSovereigntyCampaignsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *GetSovereigntyCampaignsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[GET /sovereignty/campaigns/][%d] getSovereigntyCampaignsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *GetSovereigntyCampaignsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSovereigntyCampaignsInternalServerError creates a GetSovereigntyCampaignsInternalServerError with default headers values
func NewGetSovereigntyCampaignsInternalServerError() *GetSovereigntyCampaignsInternalServerError {
	return &GetSovereigntyCampaignsInternalServerError{}
}

/*GetSovereigntyCampaignsInternalServerError handles this case with default header values.

Internal server error
*/
type GetSovereigntyCampaignsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *GetSovereigntyCampaignsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /sovereignty/campaigns/][%d] getSovereigntyCampaignsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSovereigntyCampaignsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSovereigntyCampaignsServiceUnavailable creates a GetSovereigntyCampaignsServiceUnavailable with default headers values
func NewGetSovereigntyCampaignsServiceUnavailable() *GetSovereigntyCampaignsServiceUnavailable {
	return &GetSovereigntyCampaignsServiceUnavailable{}
}

/*GetSovereigntyCampaignsServiceUnavailable handles this case with default header values.

Service unavailable
*/
type GetSovereigntyCampaignsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *GetSovereigntyCampaignsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /sovereignty/campaigns/][%d] getSovereigntyCampaignsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSovereigntyCampaignsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSovereigntyCampaignsGatewayTimeout creates a GetSovereigntyCampaignsGatewayTimeout with default headers values
func NewGetSovereigntyCampaignsGatewayTimeout() *GetSovereigntyCampaignsGatewayTimeout {
	return &GetSovereigntyCampaignsGatewayTimeout{}
}

/*GetSovereigntyCampaignsGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type GetSovereigntyCampaignsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *GetSovereigntyCampaignsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /sovereignty/campaigns/][%d] getSovereigntyCampaignsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSovereigntyCampaignsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}