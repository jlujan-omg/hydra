// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewRevokeAllUserConsentSessionsParams creates a new RevokeAllUserConsentSessionsParams object
// with the default values initialized.
func NewRevokeAllUserConsentSessionsParams() *RevokeAllUserConsentSessionsParams {
	var ()
	return &RevokeAllUserConsentSessionsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRevokeAllUserConsentSessionsParamsWithTimeout creates a new RevokeAllUserConsentSessionsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRevokeAllUserConsentSessionsParamsWithTimeout(timeout time.Duration) *RevokeAllUserConsentSessionsParams {
	var ()
	return &RevokeAllUserConsentSessionsParams{

		timeout: timeout,
	}
}

// NewRevokeAllUserConsentSessionsParamsWithContext creates a new RevokeAllUserConsentSessionsParams object
// with the default values initialized, and the ability to set a context for a request
func NewRevokeAllUserConsentSessionsParamsWithContext(ctx context.Context) *RevokeAllUserConsentSessionsParams {
	var ()
	return &RevokeAllUserConsentSessionsParams{

		Context: ctx,
	}
}

// NewRevokeAllUserConsentSessionsParamsWithHTTPClient creates a new RevokeAllUserConsentSessionsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRevokeAllUserConsentSessionsParamsWithHTTPClient(client *http.Client) *RevokeAllUserConsentSessionsParams {
	var ()
	return &RevokeAllUserConsentSessionsParams{
		HTTPClient: client,
	}
}

/*RevokeAllUserConsentSessionsParams contains all the parameters to send to the API endpoint
for the revoke all user consent sessions operation typically these are written to a http.Request
*/
type RevokeAllUserConsentSessionsParams struct {

	/*User*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the revoke all user consent sessions params
func (o *RevokeAllUserConsentSessionsParams) WithTimeout(timeout time.Duration) *RevokeAllUserConsentSessionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the revoke all user consent sessions params
func (o *RevokeAllUserConsentSessionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the revoke all user consent sessions params
func (o *RevokeAllUserConsentSessionsParams) WithContext(ctx context.Context) *RevokeAllUserConsentSessionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the revoke all user consent sessions params
func (o *RevokeAllUserConsentSessionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the revoke all user consent sessions params
func (o *RevokeAllUserConsentSessionsParams) WithHTTPClient(client *http.Client) *RevokeAllUserConsentSessionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the revoke all user consent sessions params
func (o *RevokeAllUserConsentSessionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUser adds the user to the revoke all user consent sessions params
func (o *RevokeAllUserConsentSessionsParams) WithUser(user string) *RevokeAllUserConsentSessionsParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the revoke all user consent sessions params
func (o *RevokeAllUserConsentSessionsParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *RevokeAllUserConsentSessionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user
	if err := r.SetPathParam("user", o.User); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
