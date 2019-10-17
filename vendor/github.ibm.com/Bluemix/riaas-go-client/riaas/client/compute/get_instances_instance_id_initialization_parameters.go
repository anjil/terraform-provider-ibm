// Code generated by go-swagger; DO NOT EDIT.

package compute

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

// NewGetInstancesInstanceIDInitializationParams creates a new GetInstancesInstanceIDInitializationParams object
// with the default values initialized.
func NewGetInstancesInstanceIDInitializationParams() *GetInstancesInstanceIDInitializationParams {
	var ()
	return &GetInstancesInstanceIDInitializationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetInstancesInstanceIDInitializationParamsWithTimeout creates a new GetInstancesInstanceIDInitializationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetInstancesInstanceIDInitializationParamsWithTimeout(timeout time.Duration) *GetInstancesInstanceIDInitializationParams {
	var ()
	return &GetInstancesInstanceIDInitializationParams{

		timeout: timeout,
	}
}

// NewGetInstancesInstanceIDInitializationParamsWithContext creates a new GetInstancesInstanceIDInitializationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetInstancesInstanceIDInitializationParamsWithContext(ctx context.Context) *GetInstancesInstanceIDInitializationParams {
	var ()
	return &GetInstancesInstanceIDInitializationParams{

		Context: ctx,
	}
}

// NewGetInstancesInstanceIDInitializationParamsWithHTTPClient creates a new GetInstancesInstanceIDInitializationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetInstancesInstanceIDInitializationParamsWithHTTPClient(client *http.Client) *GetInstancesInstanceIDInitializationParams {
	var ()
	return &GetInstancesInstanceIDInitializationParams{
		HTTPClient: client,
	}
}

/*GetInstancesInstanceIDInitializationParams contains all the parameters to send to the API endpoint
for the get instances instance ID initialization operation typically these are written to a http.Request
*/
type GetInstancesInstanceIDInitializationParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*InstanceID
	  The instance identifier

	*/
	InstanceID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get instances instance ID initialization params
func (o *GetInstancesInstanceIDInitializationParams) WithTimeout(timeout time.Duration) *GetInstancesInstanceIDInitializationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get instances instance ID initialization params
func (o *GetInstancesInstanceIDInitializationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get instances instance ID initialization params
func (o *GetInstancesInstanceIDInitializationParams) WithContext(ctx context.Context) *GetInstancesInstanceIDInitializationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get instances instance ID initialization params
func (o *GetInstancesInstanceIDInitializationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get instances instance ID initialization params
func (o *GetInstancesInstanceIDInitializationParams) WithHTTPClient(client *http.Client) *GetInstancesInstanceIDInitializationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get instances instance ID initialization params
func (o *GetInstancesInstanceIDInitializationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the get instances instance ID initialization params
func (o *GetInstancesInstanceIDInitializationParams) WithGeneration(generation int64) *GetInstancesInstanceIDInitializationParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the get instances instance ID initialization params
func (o *GetInstancesInstanceIDInitializationParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithInstanceID adds the instanceID to the get instances instance ID initialization params
func (o *GetInstancesInstanceIDInitializationParams) WithInstanceID(instanceID string) *GetInstancesInstanceIDInitializationParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the get instances instance ID initialization params
func (o *GetInstancesInstanceIDInitializationParams) SetInstanceID(instanceID string) {
	o.InstanceID = instanceID
}

// WithVersion adds the version to the get instances instance ID initialization params
func (o *GetInstancesInstanceIDInitializationParams) WithVersion(version string) *GetInstancesInstanceIDInitializationParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get instances instance ID initialization params
func (o *GetInstancesInstanceIDInitializationParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetInstancesInstanceIDInitializationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param instance_id
	if err := r.SetPathParam("instance_id", o.InstanceID); err != nil {
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
