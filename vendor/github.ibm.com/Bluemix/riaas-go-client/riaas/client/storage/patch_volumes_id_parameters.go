// Code generated by go-swagger; DO NOT EDIT.

package storage

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

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// NewPatchVolumesIDParams creates a new PatchVolumesIDParams object
// with the default values initialized.
func NewPatchVolumesIDParams() *PatchVolumesIDParams {
	var ()
	return &PatchVolumesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchVolumesIDParamsWithTimeout creates a new PatchVolumesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchVolumesIDParamsWithTimeout(timeout time.Duration) *PatchVolumesIDParams {
	var ()
	return &PatchVolumesIDParams{

		timeout: timeout,
	}
}

// NewPatchVolumesIDParamsWithContext creates a new PatchVolumesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchVolumesIDParamsWithContext(ctx context.Context) *PatchVolumesIDParams {
	var ()
	return &PatchVolumesIDParams{

		Context: ctx,
	}
}

// NewPatchVolumesIDParamsWithHTTPClient creates a new PatchVolumesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchVolumesIDParamsWithHTTPClient(client *http.Client) *PatchVolumesIDParams {
	var ()
	return &PatchVolumesIDParams{
		HTTPClient: client,
	}
}

/*PatchVolumesIDParams contains all the parameters to send to the API endpoint
for the patch volumes ID operation typically these are written to a http.Request
*/
type PatchVolumesIDParams struct {

	/*Body*/
	Body *models.PatchVolumesIDParamsBody
	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The volume identifier

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

// WithTimeout adds the timeout to the patch volumes ID params
func (o *PatchVolumesIDParams) WithTimeout(timeout time.Duration) *PatchVolumesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch volumes ID params
func (o *PatchVolumesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch volumes ID params
func (o *PatchVolumesIDParams) WithContext(ctx context.Context) *PatchVolumesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch volumes ID params
func (o *PatchVolumesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch volumes ID params
func (o *PatchVolumesIDParams) WithHTTPClient(client *http.Client) *PatchVolumesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch volumes ID params
func (o *PatchVolumesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch volumes ID params
func (o *PatchVolumesIDParams) WithBody(body *models.PatchVolumesIDParamsBody) *PatchVolumesIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch volumes ID params
func (o *PatchVolumesIDParams) SetBody(body *models.PatchVolumesIDParamsBody) {
	o.Body = body
}

// WithGeneration adds the generation to the patch volumes ID params
func (o *PatchVolumesIDParams) WithGeneration(generation int64) *PatchVolumesIDParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the patch volumes ID params
func (o *PatchVolumesIDParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the patch volumes ID params
func (o *PatchVolumesIDParams) WithID(id string) *PatchVolumesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch volumes ID params
func (o *PatchVolumesIDParams) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the patch volumes ID params
func (o *PatchVolumesIDParams) WithVersion(version string) *PatchVolumesIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the patch volumes ID params
func (o *PatchVolumesIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PatchVolumesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

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
