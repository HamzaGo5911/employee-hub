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

// NewGetEmployeeParams creates a new GetEmployeeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetEmployeeParams() *GetEmployeeParams {
	return &GetEmployeeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetEmployeeParamsWithTimeout creates a new GetEmployeeParams object
// with the ability to set a timeout on a request.
func NewGetEmployeeParamsWithTimeout(timeout time.Duration) *GetEmployeeParams {
	return &GetEmployeeParams{
		timeout: timeout,
	}
}

// NewGetEmployeeParamsWithContext creates a new GetEmployeeParams object
// with the ability to set a context for a request.
func NewGetEmployeeParamsWithContext(ctx context.Context) *GetEmployeeParams {
	return &GetEmployeeParams{
		Context: ctx,
	}
}

// NewGetEmployeeParamsWithHTTPClient creates a new GetEmployeeParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetEmployeeParamsWithHTTPClient(client *http.Client) *GetEmployeeParams {
	return &GetEmployeeParams{
		HTTPClient: client,
	}
}

/*
GetEmployeeParams contains all the parameters to send to the API endpoint

	for the get employee operation.

	Typically these are written to a http.Request.
*/
type GetEmployeeParams struct {

	/* ID.

	   ID of the employee
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get employee params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEmployeeParams) WithDefaults() *GetEmployeeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get employee params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetEmployeeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get employee params
func (o *GetEmployeeParams) WithTimeout(timeout time.Duration) *GetEmployeeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get employee params
func (o *GetEmployeeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get employee params
func (o *GetEmployeeParams) WithContext(ctx context.Context) *GetEmployeeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get employee params
func (o *GetEmployeeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get employee params
func (o *GetEmployeeParams) WithHTTPClient(client *http.Client) *GetEmployeeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get employee params
func (o *GetEmployeeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get employee params
func (o *GetEmployeeParams) WithID(id string) *GetEmployeeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get employee params
func (o *GetEmployeeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetEmployeeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param ID
	if err := r.SetPathParam("ID", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
