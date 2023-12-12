// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/HamzaGo5911/wanclouds-employee-hub/gen/models"
)

// AddHomeReader is a Reader for the AddHome structure.
type AddHomeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddHomeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewAddHomeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddHomeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddHomeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddHomeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /home] addHome", response, response.Code())
	}
}

// NewAddHomeCreated creates a AddHomeCreated with default headers values
func NewAddHomeCreated() *AddHomeCreated {
	return &AddHomeCreated{}
}

/*
AddHomeCreated describes a response with status code 201, with default header values.

Home added
*/
type AddHomeCreated struct {
	Payload *models.Home
}

// IsSuccess returns true when this add home created response has a 2xx status code
func (o *AddHomeCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add home created response has a 3xx status code
func (o *AddHomeCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add home created response has a 4xx status code
func (o *AddHomeCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this add home created response has a 5xx status code
func (o *AddHomeCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this add home created response a status code equal to that given
func (o *AddHomeCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the add home created response
func (o *AddHomeCreated) Code() int {
	return 201
}

func (o *AddHomeCreated) Error() string {
	return fmt.Sprintf("[POST /home][%d] addHomeCreated  %+v", 201, o.Payload)
}

func (o *AddHomeCreated) String() string {
	return fmt.Sprintf("[POST /home][%d] addHomeCreated  %+v", 201, o.Payload)
}

func (o *AddHomeCreated) GetPayload() *models.Home {
	return o.Payload
}

func (o *AddHomeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Home)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddHomeBadRequest creates a AddHomeBadRequest with default headers values
func NewAddHomeBadRequest() *AddHomeBadRequest {
	return &AddHomeBadRequest{}
}

/*
AddHomeBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type AddHomeBadRequest struct {
}

// IsSuccess returns true when this add home bad request response has a 2xx status code
func (o *AddHomeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add home bad request response has a 3xx status code
func (o *AddHomeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add home bad request response has a 4xx status code
func (o *AddHomeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add home bad request response has a 5xx status code
func (o *AddHomeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add home bad request response a status code equal to that given
func (o *AddHomeBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the add home bad request response
func (o *AddHomeBadRequest) Code() int {
	return 400
}

func (o *AddHomeBadRequest) Error() string {
	return fmt.Sprintf("[POST /home][%d] addHomeBadRequest ", 400)
}

func (o *AddHomeBadRequest) String() string {
	return fmt.Sprintf("[POST /home][%d] addHomeBadRequest ", 400)
}

func (o *AddHomeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddHomeConflict creates a AddHomeConflict with default headers values
func NewAddHomeConflict() *AddHomeConflict {
	return &AddHomeConflict{}
}

/*
AddHomeConflict describes a response with status code 409, with default header values.

Home already exists
*/
type AddHomeConflict struct {
}

// IsSuccess returns true when this add home conflict response has a 2xx status code
func (o *AddHomeConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add home conflict response has a 3xx status code
func (o *AddHomeConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add home conflict response has a 4xx status code
func (o *AddHomeConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this add home conflict response has a 5xx status code
func (o *AddHomeConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this add home conflict response a status code equal to that given
func (o *AddHomeConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the add home conflict response
func (o *AddHomeConflict) Code() int {
	return 409
}

func (o *AddHomeConflict) Error() string {
	return fmt.Sprintf("[POST /home][%d] addHomeConflict ", 409)
}

func (o *AddHomeConflict) String() string {
	return fmt.Sprintf("[POST /home][%d] addHomeConflict ", 409)
}

func (o *AddHomeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAddHomeInternalServerError creates a AddHomeInternalServerError with default headers values
func NewAddHomeInternalServerError() *AddHomeInternalServerError {
	return &AddHomeInternalServerError{}
}

/*
AddHomeInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type AddHomeInternalServerError struct {
}

// IsSuccess returns true when this add home internal server error response has a 2xx status code
func (o *AddHomeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add home internal server error response has a 3xx status code
func (o *AddHomeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add home internal server error response has a 4xx status code
func (o *AddHomeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add home internal server error response has a 5xx status code
func (o *AddHomeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add home internal server error response a status code equal to that given
func (o *AddHomeInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the add home internal server error response
func (o *AddHomeInternalServerError) Code() int {
	return 500
}

func (o *AddHomeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /home][%d] addHomeInternalServerError ", 500)
}

func (o *AddHomeInternalServerError) String() string {
	return fmt.Sprintf("[POST /home][%d] addHomeInternalServerError ", 500)
}

func (o *AddHomeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
