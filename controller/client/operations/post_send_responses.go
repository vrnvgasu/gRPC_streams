// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostSendReader is a Reader for the PostSend structure.
type PostSendReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *PostSendReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSendOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostSendBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostSendInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /send] PostSend", response, response.Code())
	}
}

// NewPostSendOK creates a PostSendOK with default headers values
func NewPostSendOK(writer io.Writer) *PostSendOK {
	return &PostSendOK{

		Payload: writer,
	}
}

/*
PostSendOK describes a response with status code 200, with default header values.

A PDF file.
*/
type PostSendOK struct {
	Payload io.Writer
}

// IsSuccess returns true when this post send o k response has a 2xx status code
func (o *PostSendOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post send o k response has a 3xx status code
func (o *PostSendOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post send o k response has a 4xx status code
func (o *PostSendOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post send o k response has a 5xx status code
func (o *PostSendOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post send o k response a status code equal to that given
func (o *PostSendOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post send o k response
func (o *PostSendOK) Code() int {
	return 200
}

func (o *PostSendOK) Error() string {
	return fmt.Sprintf("[POST /send][%d] postSendOK  %+v", 200, o.Payload)
}

func (o *PostSendOK) String() string {
	return fmt.Sprintf("[POST /send][%d] postSendOK  %+v", 200, o.Payload)
}

func (o *PostSendOK) GetPayload() io.Writer {
	return o.Payload
}

func (o *PostSendOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSendBadRequest creates a PostSendBadRequest with default headers values
func NewPostSendBadRequest() *PostSendBadRequest {
	return &PostSendBadRequest{}
}

/*
PostSendBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PostSendBadRequest struct {
}

// IsSuccess returns true when this post send bad request response has a 2xx status code
func (o *PostSendBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post send bad request response has a 3xx status code
func (o *PostSendBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post send bad request response has a 4xx status code
func (o *PostSendBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post send bad request response has a 5xx status code
func (o *PostSendBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post send bad request response a status code equal to that given
func (o *PostSendBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post send bad request response
func (o *PostSendBadRequest) Code() int {
	return 400
}

func (o *PostSendBadRequest) Error() string {
	return fmt.Sprintf("[POST /send][%d] postSendBadRequest ", 400)
}

func (o *PostSendBadRequest) String() string {
	return fmt.Sprintf("[POST /send][%d] postSendBadRequest ", 400)
}

func (o *PostSendBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostSendInternalServerError creates a PostSendInternalServerError with default headers values
func NewPostSendInternalServerError() *PostSendInternalServerError {
	return &PostSendInternalServerError{}
}

/*
PostSendInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PostSendInternalServerError struct {
}

// IsSuccess returns true when this post send internal server error response has a 2xx status code
func (o *PostSendInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post send internal server error response has a 3xx status code
func (o *PostSendInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post send internal server error response has a 4xx status code
func (o *PostSendInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post send internal server error response has a 5xx status code
func (o *PostSendInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post send internal server error response a status code equal to that given
func (o *PostSendInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the post send internal server error response
func (o *PostSendInternalServerError) Code() int {
	return 500
}

func (o *PostSendInternalServerError) Error() string {
	return fmt.Sprintf("[POST /send][%d] postSendInternalServerError ", 500)
}

func (o *PostSendInternalServerError) String() string {
	return fmt.Sprintf("[POST /send][%d] postSendInternalServerError ", 500)
}

func (o *PostSendInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
