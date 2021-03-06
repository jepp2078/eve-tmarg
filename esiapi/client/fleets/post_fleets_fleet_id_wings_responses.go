// Code generated by go-swagger; DO NOT EDIT.

package fleets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/eve-tmarg/esiapi/models"
)

// PostFleetsFleetIDWingsReader is a Reader for the PostFleetsFleetIDWings structure.
type PostFleetsFleetIDWingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostFleetsFleetIDWingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostFleetsFleetIDWingsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostFleetsFleetIDWingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostFleetsFleetIDWingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostFleetsFleetIDWingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostFleetsFleetIDWingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewPostFleetsFleetIDWingsEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostFleetsFleetIDWingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewPostFleetsFleetIDWingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewPostFleetsFleetIDWingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostFleetsFleetIDWingsCreated creates a PostFleetsFleetIDWingsCreated with default headers values
func NewPostFleetsFleetIDWingsCreated() *PostFleetsFleetIDWingsCreated {
	return &PostFleetsFleetIDWingsCreated{}
}

/*PostFleetsFleetIDWingsCreated handles this case with default header values.

Wing created
*/
type PostFleetsFleetIDWingsCreated struct {
	Payload *models.PostFleetsFleetIDWingsCreatedBody
}

func (o *PostFleetsFleetIDWingsCreated) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsCreated  %+v", 201, o.Payload)
}

func (o *PostFleetsFleetIDWingsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostFleetsFleetIDWingsCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDWingsBadRequest creates a PostFleetsFleetIDWingsBadRequest with default headers values
func NewPostFleetsFleetIDWingsBadRequest() *PostFleetsFleetIDWingsBadRequest {
	return &PostFleetsFleetIDWingsBadRequest{}
}

/*PostFleetsFleetIDWingsBadRequest handles this case with default header values.

Bad request
*/
type PostFleetsFleetIDWingsBadRequest struct {
	Payload *models.BadRequest
}

func (o *PostFleetsFleetIDWingsBadRequest) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsBadRequest  %+v", 400, o.Payload)
}

func (o *PostFleetsFleetIDWingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDWingsUnauthorized creates a PostFleetsFleetIDWingsUnauthorized with default headers values
func NewPostFleetsFleetIDWingsUnauthorized() *PostFleetsFleetIDWingsUnauthorized {
	return &PostFleetsFleetIDWingsUnauthorized{}
}

/*PostFleetsFleetIDWingsUnauthorized handles this case with default header values.

Unauthorized
*/
type PostFleetsFleetIDWingsUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *PostFleetsFleetIDWingsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostFleetsFleetIDWingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDWingsForbidden creates a PostFleetsFleetIDWingsForbidden with default headers values
func NewPostFleetsFleetIDWingsForbidden() *PostFleetsFleetIDWingsForbidden {
	return &PostFleetsFleetIDWingsForbidden{}
}

/*PostFleetsFleetIDWingsForbidden handles this case with default header values.

Forbidden
*/
type PostFleetsFleetIDWingsForbidden struct {
	Payload *models.Forbidden
}

func (o *PostFleetsFleetIDWingsForbidden) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsForbidden  %+v", 403, o.Payload)
}

func (o *PostFleetsFleetIDWingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDWingsNotFound creates a PostFleetsFleetIDWingsNotFound with default headers values
func NewPostFleetsFleetIDWingsNotFound() *PostFleetsFleetIDWingsNotFound {
	return &PostFleetsFleetIDWingsNotFound{}
}

/*PostFleetsFleetIDWingsNotFound handles this case with default header values.

The fleet does not exist or you don't have access to it
*/
type PostFleetsFleetIDWingsNotFound struct {
	Payload *models.PostFleetsFleetIDWingsNotFoundBody
}

func (o *PostFleetsFleetIDWingsNotFound) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsNotFound  %+v", 404, o.Payload)
}

func (o *PostFleetsFleetIDWingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostFleetsFleetIDWingsNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDWingsEnhanceYourCalm creates a PostFleetsFleetIDWingsEnhanceYourCalm with default headers values
func NewPostFleetsFleetIDWingsEnhanceYourCalm() *PostFleetsFleetIDWingsEnhanceYourCalm {
	return &PostFleetsFleetIDWingsEnhanceYourCalm{}
}

/*PostFleetsFleetIDWingsEnhanceYourCalm handles this case with default header values.

Error limited
*/
type PostFleetsFleetIDWingsEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *PostFleetsFleetIDWingsEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PostFleetsFleetIDWingsEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDWingsInternalServerError creates a PostFleetsFleetIDWingsInternalServerError with default headers values
func NewPostFleetsFleetIDWingsInternalServerError() *PostFleetsFleetIDWingsInternalServerError {
	return &PostFleetsFleetIDWingsInternalServerError{}
}

/*PostFleetsFleetIDWingsInternalServerError handles this case with default header values.

Internal server error
*/
type PostFleetsFleetIDWingsInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *PostFleetsFleetIDWingsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostFleetsFleetIDWingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDWingsServiceUnavailable creates a PostFleetsFleetIDWingsServiceUnavailable with default headers values
func NewPostFleetsFleetIDWingsServiceUnavailable() *PostFleetsFleetIDWingsServiceUnavailable {
	return &PostFleetsFleetIDWingsServiceUnavailable{}
}

/*PostFleetsFleetIDWingsServiceUnavailable handles this case with default header values.

Service unavailable
*/
type PostFleetsFleetIDWingsServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *PostFleetsFleetIDWingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostFleetsFleetIDWingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostFleetsFleetIDWingsGatewayTimeout creates a PostFleetsFleetIDWingsGatewayTimeout with default headers values
func NewPostFleetsFleetIDWingsGatewayTimeout() *PostFleetsFleetIDWingsGatewayTimeout {
	return &PostFleetsFleetIDWingsGatewayTimeout{}
}

/*PostFleetsFleetIDWingsGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type PostFleetsFleetIDWingsGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *PostFleetsFleetIDWingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /fleets/{fleet_id}/wings/][%d] postFleetsFleetIdWingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostFleetsFleetIDWingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
