// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package extras

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

// NewExtrasWebhooksBulkDeleteParams creates a new ExtrasWebhooksBulkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExtrasWebhooksBulkDeleteParams() *ExtrasWebhooksBulkDeleteParams {
	return &ExtrasWebhooksBulkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasWebhooksBulkDeleteParamsWithTimeout creates a new ExtrasWebhooksBulkDeleteParams object
// with the ability to set a timeout on a request.
func NewExtrasWebhooksBulkDeleteParamsWithTimeout(timeout time.Duration) *ExtrasWebhooksBulkDeleteParams {
	return &ExtrasWebhooksBulkDeleteParams{
		timeout: timeout,
	}
}

// NewExtrasWebhooksBulkDeleteParamsWithContext creates a new ExtrasWebhooksBulkDeleteParams object
// with the ability to set a context for a request.
func NewExtrasWebhooksBulkDeleteParamsWithContext(ctx context.Context) *ExtrasWebhooksBulkDeleteParams {
	return &ExtrasWebhooksBulkDeleteParams{
		Context: ctx,
	}
}

// NewExtrasWebhooksBulkDeleteParamsWithHTTPClient creates a new ExtrasWebhooksBulkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewExtrasWebhooksBulkDeleteParamsWithHTTPClient(client *http.Client) *ExtrasWebhooksBulkDeleteParams {
	return &ExtrasWebhooksBulkDeleteParams{
		HTTPClient: client,
	}
}

/* ExtrasWebhooksBulkDeleteParams contains all the parameters to send to the API endpoint
   for the extras webhooks bulk delete operation.

   Typically these are written to a http.Request.
*/
type ExtrasWebhooksBulkDeleteParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the extras webhooks bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasWebhooksBulkDeleteParams) WithDefaults() *ExtrasWebhooksBulkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the extras webhooks bulk delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExtrasWebhooksBulkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the extras webhooks bulk delete params
func (o *ExtrasWebhooksBulkDeleteParams) WithTimeout(timeout time.Duration) *ExtrasWebhooksBulkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras webhooks bulk delete params
func (o *ExtrasWebhooksBulkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras webhooks bulk delete params
func (o *ExtrasWebhooksBulkDeleteParams) WithContext(ctx context.Context) *ExtrasWebhooksBulkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras webhooks bulk delete params
func (o *ExtrasWebhooksBulkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras webhooks bulk delete params
func (o *ExtrasWebhooksBulkDeleteParams) WithHTTPClient(client *http.Client) *ExtrasWebhooksBulkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras webhooks bulk delete params
func (o *ExtrasWebhooksBulkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasWebhooksBulkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}