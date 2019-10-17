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

// NewDeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams creates a new DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams object
// with the default values initialized.
func NewDeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams() *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams {
	var ()
	return &DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParamsWithTimeout creates a new DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParamsWithTimeout(timeout time.Duration) *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams {
	var ()
	return &DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams{

		timeout: timeout,
	}
}

// NewDeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParamsWithContext creates a new DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParamsWithContext(ctx context.Context) *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams {
	var ()
	return &DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams{

		Context: ctx,
	}
}

// NewDeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParamsWithHTTPClient creates a new DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParamsWithHTTPClient(client *http.Client) *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams {
	var ()
	return &DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams{
		HTTPClient: client,
	}
}

/*DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams contains all the parameters to send to the API endpoint
for the delete security groups security group ID network interfaces ID operation typically these are written to a http.Request
*/
type DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams struct {

	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*ID
	  The network interface identifier

	*/
	ID string
	/*SecurityGroupID
	  The security group identifier

	*/
	SecurityGroupID string
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete security groups security group ID network interfaces ID params
func (o *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) WithTimeout(timeout time.Duration) *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete security groups security group ID network interfaces ID params
func (o *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete security groups security group ID network interfaces ID params
func (o *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) WithContext(ctx context.Context) *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete security groups security group ID network interfaces ID params
func (o *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete security groups security group ID network interfaces ID params
func (o *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) WithHTTPClient(client *http.Client) *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete security groups security group ID network interfaces ID params
func (o *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithGeneration adds the generation to the delete security groups security group ID network interfaces ID params
func (o *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) WithGeneration(generation int64) *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the delete security groups security group ID network interfaces ID params
func (o *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithID adds the id to the delete security groups security group ID network interfaces ID params
func (o *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) WithID(id string) *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete security groups security group ID network interfaces ID params
func (o *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) SetID(id string) {
	o.ID = id
}

// WithSecurityGroupID adds the securityGroupID to the delete security groups security group ID network interfaces ID params
func (o *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) WithSecurityGroupID(securityGroupID string) *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams {
	o.SetSecurityGroupID(securityGroupID)
	return o
}

// SetSecurityGroupID adds the securityGroupId to the delete security groups security group ID network interfaces ID params
func (o *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) SetSecurityGroupID(securityGroupID string) {
	o.SecurityGroupID = securityGroupID
}

// WithVersion adds the version to the delete security groups security group ID network interfaces ID params
func (o *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) WithVersion(version string) *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the delete security groups security group ID network interfaces ID params
func (o *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteSecurityGroupsSecurityGroupIDNetworkInterfacesIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param security_group_id
	if err := r.SetPathParam("security_group_id", o.SecurityGroupID); err != nil {
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
