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

// PutFleetsFleetIDSquadsSquadIDReader is a Reader for the PutFleetsFleetIDSquadsSquadID structure.
type PutFleetsFleetIDSquadsSquadIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutFleetsFleetIDSquadsSquadIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewPutFleetsFleetIDSquadsSquadIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutFleetsFleetIDSquadsSquadIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPutFleetsFleetIDSquadsSquadIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPutFleetsFleetIDSquadsSquadIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutFleetsFleetIDSquadsSquadIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 420:
		result := NewPutFleetsFleetIDSquadsSquadIDEnhanceYourCalm()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPutFleetsFleetIDSquadsSquadIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 503:
		result := NewPutFleetsFleetIDSquadsSquadIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 504:
		result := NewPutFleetsFleetIDSquadsSquadIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutFleetsFleetIDSquadsSquadIDNoContent creates a PutFleetsFleetIDSquadsSquadIDNoContent with default headers values
func NewPutFleetsFleetIDSquadsSquadIDNoContent() *PutFleetsFleetIDSquadsSquadIDNoContent {
	return &PutFleetsFleetIDSquadsSquadIDNoContent{}
}

/*PutFleetsFleetIDSquadsSquadIDNoContent handles this case with default header values.

Squad renamed
*/
type PutFleetsFleetIDSquadsSquadIDNoContent struct {
}

func (o *PutFleetsFleetIDSquadsSquadIDNoContent) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/squads/{squad_id}/][%d] putFleetsFleetIdSquadsSquadIdNoContent ", 204)
}

func (o *PutFleetsFleetIDSquadsSquadIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutFleetsFleetIDSquadsSquadIDBadRequest creates a PutFleetsFleetIDSquadsSquadIDBadRequest with default headers values
func NewPutFleetsFleetIDSquadsSquadIDBadRequest() *PutFleetsFleetIDSquadsSquadIDBadRequest {
	return &PutFleetsFleetIDSquadsSquadIDBadRequest{}
}

/*PutFleetsFleetIDSquadsSquadIDBadRequest handles this case with default header values.

Bad request
*/
type PutFleetsFleetIDSquadsSquadIDBadRequest struct {
	Payload *models.BadRequest
}

func (o *PutFleetsFleetIDSquadsSquadIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/squads/{squad_id}/][%d] putFleetsFleetIdSquadsSquadIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutFleetsFleetIDSquadsSquadIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.BadRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDSquadsSquadIDUnauthorized creates a PutFleetsFleetIDSquadsSquadIDUnauthorized with default headers values
func NewPutFleetsFleetIDSquadsSquadIDUnauthorized() *PutFleetsFleetIDSquadsSquadIDUnauthorized {
	return &PutFleetsFleetIDSquadsSquadIDUnauthorized{}
}

/*PutFleetsFleetIDSquadsSquadIDUnauthorized handles this case with default header values.

Unauthorized
*/
type PutFleetsFleetIDSquadsSquadIDUnauthorized struct {
	Payload *models.Unauthorized
}

func (o *PutFleetsFleetIDSquadsSquadIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/squads/{squad_id}/][%d] putFleetsFleetIdSquadsSquadIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PutFleetsFleetIDSquadsSquadIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Unauthorized)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDSquadsSquadIDForbidden creates a PutFleetsFleetIDSquadsSquadIDForbidden with default headers values
func NewPutFleetsFleetIDSquadsSquadIDForbidden() *PutFleetsFleetIDSquadsSquadIDForbidden {
	return &PutFleetsFleetIDSquadsSquadIDForbidden{}
}

/*PutFleetsFleetIDSquadsSquadIDForbidden handles this case with default header values.

Forbidden
*/
type PutFleetsFleetIDSquadsSquadIDForbidden struct {
	Payload *models.Forbidden
}

func (o *PutFleetsFleetIDSquadsSquadIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/squads/{squad_id}/][%d] putFleetsFleetIdSquadsSquadIdForbidden  %+v", 403, o.Payload)
}

func (o *PutFleetsFleetIDSquadsSquadIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Forbidden)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDSquadsSquadIDNotFound creates a PutFleetsFleetIDSquadsSquadIDNotFound with default headers values
func NewPutFleetsFleetIDSquadsSquadIDNotFound() *PutFleetsFleetIDSquadsSquadIDNotFound {
	return &PutFleetsFleetIDSquadsSquadIDNotFound{}
}

/*PutFleetsFleetIDSquadsSquadIDNotFound handles this case with default header values.

The fleet does not exist or you don't have access to it
*/
type PutFleetsFleetIDSquadsSquadIDNotFound struct {
	Payload *models.PutFleetsFleetIDSquadsSquadIDNotFoundBody
}

func (o *PutFleetsFleetIDSquadsSquadIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/squads/{squad_id}/][%d] putFleetsFleetIdSquadsSquadIdNotFound  %+v", 404, o.Payload)
}

func (o *PutFleetsFleetIDSquadsSquadIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PutFleetsFleetIDSquadsSquadIDNotFoundBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDSquadsSquadIDEnhanceYourCalm creates a PutFleetsFleetIDSquadsSquadIDEnhanceYourCalm with default headers values
func NewPutFleetsFleetIDSquadsSquadIDEnhanceYourCalm() *PutFleetsFleetIDSquadsSquadIDEnhanceYourCalm {
	return &PutFleetsFleetIDSquadsSquadIDEnhanceYourCalm{}
}

/*PutFleetsFleetIDSquadsSquadIDEnhanceYourCalm handles this case with default header values.

Error limited
*/
type PutFleetsFleetIDSquadsSquadIDEnhanceYourCalm struct {
	Payload *models.ErrorLimited
}

func (o *PutFleetsFleetIDSquadsSquadIDEnhanceYourCalm) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/squads/{squad_id}/][%d] putFleetsFleetIdSquadsSquadIdEnhanceYourCalm  %+v", 420, o.Payload)
}

func (o *PutFleetsFleetIDSquadsSquadIDEnhanceYourCalm) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorLimited)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDSquadsSquadIDInternalServerError creates a PutFleetsFleetIDSquadsSquadIDInternalServerError with default headers values
func NewPutFleetsFleetIDSquadsSquadIDInternalServerError() *PutFleetsFleetIDSquadsSquadIDInternalServerError {
	return &PutFleetsFleetIDSquadsSquadIDInternalServerError{}
}

/*PutFleetsFleetIDSquadsSquadIDInternalServerError handles this case with default header values.

Internal server error
*/
type PutFleetsFleetIDSquadsSquadIDInternalServerError struct {
	Payload *models.InternalServerError
}

func (o *PutFleetsFleetIDSquadsSquadIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/squads/{squad_id}/][%d] putFleetsFleetIdSquadsSquadIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PutFleetsFleetIDSquadsSquadIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalServerError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDSquadsSquadIDServiceUnavailable creates a PutFleetsFleetIDSquadsSquadIDServiceUnavailable with default headers values
func NewPutFleetsFleetIDSquadsSquadIDServiceUnavailable() *PutFleetsFleetIDSquadsSquadIDServiceUnavailable {
	return &PutFleetsFleetIDSquadsSquadIDServiceUnavailable{}
}

/*PutFleetsFleetIDSquadsSquadIDServiceUnavailable handles this case with default header values.

Service unavailable
*/
type PutFleetsFleetIDSquadsSquadIDServiceUnavailable struct {
	Payload *models.ServiceUnavailable
}

func (o *PutFleetsFleetIDSquadsSquadIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/squads/{squad_id}/][%d] putFleetsFleetIdSquadsSquadIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutFleetsFleetIDSquadsSquadIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ServiceUnavailable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutFleetsFleetIDSquadsSquadIDGatewayTimeout creates a PutFleetsFleetIDSquadsSquadIDGatewayTimeout with default headers values
func NewPutFleetsFleetIDSquadsSquadIDGatewayTimeout() *PutFleetsFleetIDSquadsSquadIDGatewayTimeout {
	return &PutFleetsFleetIDSquadsSquadIDGatewayTimeout{}
}

/*PutFleetsFleetIDSquadsSquadIDGatewayTimeout handles this case with default header values.

Gateway timeout
*/
type PutFleetsFleetIDSquadsSquadIDGatewayTimeout struct {
	Payload *models.GatewayTimeout
}

func (o *PutFleetsFleetIDSquadsSquadIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /fleets/{fleet_id}/squads/{squad_id}/][%d] putFleetsFleetIdSquadsSquadIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutFleetsFleetIDSquadsSquadIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GatewayTimeout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}