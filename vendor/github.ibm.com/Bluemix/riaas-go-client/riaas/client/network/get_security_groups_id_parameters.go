// Code generated by go-swagger; DO NOT EDIT.

package network

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

// NewGetSecurityGroupsIDParams creates a new GetSecurityGroupsIDParams object
// with the default values initialized.
func NewGetSecurityGroupsIDParams() *GetSecurityGroupsIDParams {
	var ()
	return &GetSecurityGroupsIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSecurityGroupsIDParamsWithTimeout creates a new GetSecurityGroupsIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSecurityGroupsIDParamsWithTimeout(timeout time.Duration) *GetSecurityGroupsIDParams {
	var ()
	return &GetSecurityGroupsIDParams{

		timeout: timeout,
	}
}

// NewGetSecurityGroupsIDParamsWithContext creates a new GetSecurityGroupsIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSecurityGroupsIDParamsWithContext(ctx context.Context) *GetSecurityGroupsIDParams {
	var ()
	return &GetSecurityGroupsIDParams{

		Context: ctx,
	}
}

// NewGetSecurityGroupsIDParamsWithHTTPClient creates a new GetSecurityGroupsIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSecurityGroupsIDParamsWithHTTPClient(client *http.Client) *GetSecurityGroupsIDParams {
	var ()
	return &GetSecurityGroupsIDParams{
		HTTPClient: client,
	}
}

/*GetSecurityGroupsIDParams contains all the parameters to send to the API endpoint
for the get security groups ID operation typically these are written to a http.Request
*/
type GetSecurityGroupsIDParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The security group identifier

	*/
	ID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get security groups ID params
func (o *GetSecurityGroupsIDParams) WithTimeout(timeout time.Duration) *GetSecurityGroupsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get security groups ID params
func (o *GetSecurityGroupsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get security groups ID params
func (o *GetSecurityGroupsIDParams) WithContext(ctx context.Context) *GetSecurityGroupsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get security groups ID params
func (o *GetSecurityGroupsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get security groups ID params
func (o *GetSecurityGroupsIDParams) WithHTTPClient(client *http.Client) *GetSecurityGroupsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get security groups ID params
func (o *GetSecurityGroupsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get security groups ID params
func (o *GetSecurityGroupsIDParams) WithGeneration(generation int64) *GetSecurityGroupsIDParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get security groups ID params
func (o *GetSecurityGroupsIDParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the get security groups ID params
func (o *GetSecurityGroupsIDParams) WithID(id string) *GetSecurityGroupsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get security groups ID params
func (o *GetSecurityGroupsIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the get security groups ID params
func (o *GetSecurityGroupsIDParams) WithVersion(version string) *GetSecurityGroupsIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get security groups ID params
func (o *GetSecurityGroupsIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetSecurityGroupsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
