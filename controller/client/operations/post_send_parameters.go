// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewPostSendParams creates a new PostSendParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostSendParams() *PostSendParams {
	return &PostSendParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostSendParamsWithTimeout creates a new PostSendParams object
// with the ability to set a timeout on a request.
func NewPostSendParamsWithTimeout(timeout time.Duration) *PostSendParams {
	return &PostSendParams{
		timeout: timeout,
	}
}

// NewPostSendParamsWithContext creates a new PostSendParams object
// with the ability to set a context for a request.
func NewPostSendParamsWithContext(ctx context.Context) *PostSendParams {
	return &PostSendParams{
		Context: ctx,
	}
}

// NewPostSendParamsWithHTTPClient creates a new PostSendParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostSendParamsWithHTTPClient(client *http.Client) *PostSendParams {
	return &PostSendParams{
		HTTPClient: client,
	}
}

/*
PostSendParams contains all the parameters to send to the API endpoint

	for the post send operation.

	Typically these are written to a http.Request.
*/
type PostSendParams struct {

	// Upfile1.
	Upfile1 runtime.NamedReadCloser

	// Upfile2.
	Upfile2 runtime.NamedReadCloser

	// Upfile3.
	Upfile3 runtime.NamedReadCloser

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post send params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSendParams) WithDefaults() *PostSendParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post send params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostSendParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post send params
func (o *PostSendParams) WithTimeout(timeout time.Duration) *PostSendParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post send params
func (o *PostSendParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post send params
func (o *PostSendParams) WithContext(ctx context.Context) *PostSendParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post send params
func (o *PostSendParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post send params
func (o *PostSendParams) WithHTTPClient(client *http.Client) *PostSendParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post send params
func (o *PostSendParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUpfile1 adds the upfile1 to the post send params
func (o *PostSendParams) WithUpfile1(upfile1 runtime.NamedReadCloser) *PostSendParams {
	o.SetUpfile1(upfile1)
	return o
}

// SetUpfile1 adds the upfile1 to the post send params
func (o *PostSendParams) SetUpfile1(upfile1 runtime.NamedReadCloser) {
	o.Upfile1 = upfile1
}

// WithUpfile2 adds the upfile2 to the post send params
func (o *PostSendParams) WithUpfile2(upfile2 runtime.NamedReadCloser) *PostSendParams {
	o.SetUpfile2(upfile2)
	return o
}

// SetUpfile2 adds the upfile2 to the post send params
func (o *PostSendParams) SetUpfile2(upfile2 runtime.NamedReadCloser) {
	o.Upfile2 = upfile2
}

// WithUpfile3 adds the upfile3 to the post send params
func (o *PostSendParams) WithUpfile3(upfile3 runtime.NamedReadCloser) *PostSendParams {
	o.SetUpfile3(upfile3)
	return o
}

// SetUpfile3 adds the upfile3 to the post send params
func (o *PostSendParams) SetUpfile3(upfile3 runtime.NamedReadCloser) {
	o.Upfile3 = upfile3
}

// WriteToRequest writes these params to a swagger request
func (o *PostSendParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	// form file param upfile1
	if err := r.SetFileParam("upfile1", o.Upfile1); err != nil {
		return err
	}

	if o.Upfile2 != nil {

		if o.Upfile2 != nil {
			// form file param upfile2
			if err := r.SetFileParam("upfile2", o.Upfile2); err != nil {
				return err
			}
		}
	}

	if o.Upfile3 != nil {

		if o.Upfile3 != nil {
			// form file param upfile3
			if err := r.SetFileParam("upfile3", o.Upfile3); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
