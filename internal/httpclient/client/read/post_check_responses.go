// Code generated by go-swagger; DO NOT EDIT.

package read

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/keto/internal/httpclient/models"
)

// PostCheckReader is a Reader for the PostCheck structure.
type PostCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostCheckBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostCheckInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostCheckOK creates a PostCheckOK with default headers values
func NewPostCheckOK() *PostCheckOK {
	return &PostCheckOK{}
}

/* PostCheckOK describes a response with status code 200, with default header values.

getCheckResponse
*/
type PostCheckOK struct {
	Payload *models.GetCheckResponse
}

func (o *PostCheckOK) Error() string {
	return fmt.Sprintf("[POST /relation-tuples/check/openapi][%d] postCheckOK  %+v", 200, o.Payload)
}
func (o *PostCheckOK) GetPayload() *models.GetCheckResponse {
	return o.Payload
}

func (o *PostCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetCheckResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCheckBadRequest creates a PostCheckBadRequest with default headers values
func NewPostCheckBadRequest() *PostCheckBadRequest {
	return &PostCheckBadRequest{}
}

/* PostCheckBadRequest describes a response with status code 400, with default header values.

genericError
*/
type PostCheckBadRequest struct {
	Payload *models.GenericError
}

func (o *PostCheckBadRequest) Error() string {
	return fmt.Sprintf("[POST /relation-tuples/check/openapi][%d] postCheckBadRequest  %+v", 400, o.Payload)
}
func (o *PostCheckBadRequest) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *PostCheckBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostCheckInternalServerError creates a PostCheckInternalServerError with default headers values
func NewPostCheckInternalServerError() *PostCheckInternalServerError {
	return &PostCheckInternalServerError{}
}

/* PostCheckInternalServerError describes a response with status code 500, with default header values.

genericError
*/
type PostCheckInternalServerError struct {
	Payload *models.GenericError
}

func (o *PostCheckInternalServerError) Error() string {
	return fmt.Sprintf("[POST /relation-tuples/check/openapi][%d] postCheckInternalServerError  %+v", 500, o.Payload)
}
func (o *PostCheckInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *PostCheckInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
