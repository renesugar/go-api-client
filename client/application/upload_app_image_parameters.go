// Code generated by go-swagger; DO NOT EDIT.

package application

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewUploadAppImageParams creates a new UploadAppImageParams object
// with the default values initialized.
func NewUploadAppImageParams() *UploadAppImageParams {
	var ()
	return &UploadAppImageParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUploadAppImageParamsWithTimeout creates a new UploadAppImageParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUploadAppImageParamsWithTimeout(timeout time.Duration) *UploadAppImageParams {
	var ()
	return &UploadAppImageParams{

		timeout: timeout,
	}
}

// NewUploadAppImageParamsWithContext creates a new UploadAppImageParams object
// with the default values initialized, and the ability to set a context for a request
func NewUploadAppImageParamsWithContext(ctx context.Context) *UploadAppImageParams {
	var ()
	return &UploadAppImageParams{

		Context: ctx,
	}
}

// NewUploadAppImageParamsWithHTTPClient creates a new UploadAppImageParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUploadAppImageParamsWithHTTPClient(client *http.Client) *UploadAppImageParams {
	var ()
	return &UploadAppImageParams{
		HTTPClient: client,
	}
}

/*UploadAppImageParams contains all the parameters to send to the API endpoint
for the upload app image operation typically these are written to a http.Request
*/
type UploadAppImageParams struct {

	/*File
	  the application image

	*/
	File runtime.NamedReadCloser
	/*ID
	  the application id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the upload app image params
func (o *UploadAppImageParams) WithTimeout(timeout time.Duration) *UploadAppImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upload app image params
func (o *UploadAppImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upload app image params
func (o *UploadAppImageParams) WithContext(ctx context.Context) *UploadAppImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upload app image params
func (o *UploadAppImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upload app image params
func (o *UploadAppImageParams) WithHTTPClient(client *http.Client) *UploadAppImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upload app image params
func (o *UploadAppImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFile adds the file to the upload app image params
func (o *UploadAppImageParams) WithFile(file runtime.NamedReadCloser) *UploadAppImageParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the upload app image params
func (o *UploadAppImageParams) SetFile(file runtime.NamedReadCloser) {
	o.File = file
}

// WithID adds the id to the upload app image params
func (o *UploadAppImageParams) WithID(id int64) *UploadAppImageParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the upload app image params
func (o *UploadAppImageParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UploadAppImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// form file param file
	if err := r.SetFileParam("file", o.File); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
