// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetSecurityGroupsSecurityGroupIDNetworkInterfacesOKBody NetworkInterfaceCollection
// swagger:model getSecurityGroupsSecurityGroupIdNetworkInterfacesOKBody
type GetSecurityGroupsSecurityGroupIDNetworkInterfacesOKBody struct {

	// Collection of network interfaces
	// Required: true
	NetworkInterfaces []*ServerNetworkInterface `json:"network_interfaces"`
}

// Validate validates this get security groups security group Id network interfaces o k body
func (m *GetSecurityGroupsSecurityGroupIDNetworkInterfacesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworkInterfaces(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetSecurityGroupsSecurityGroupIDNetworkInterfacesOKBody) validateNetworkInterfaces(formats strfmt.Registry) error {

	if err := validate.Required("network_interfaces", "body", m.NetworkInterfaces); err != nil {
		return err
	}

	for i := 0; i < len(m.NetworkInterfaces); i++ {
		if swag.IsZero(m.NetworkInterfaces[i]) { // not required
			continue
		}

		if m.NetworkInterfaces[i] != nil {
			if err := m.NetworkInterfaces[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("network_interfaces" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetSecurityGroupsSecurityGroupIDNetworkInterfacesOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetSecurityGroupsSecurityGroupIDNetworkInterfacesOKBody) UnmarshalBinary(b []byte) error {
	var res GetSecurityGroupsSecurityGroupIDNetworkInterfacesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
