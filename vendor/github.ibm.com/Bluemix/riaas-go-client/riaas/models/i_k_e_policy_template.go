// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IKEPolicyTemplate i k e policy template
// swagger:model IKEPolicyTemplate
type IKEPolicyTemplate struct {

	// The authentication algorithm
	// Enum: [md5 sha1 sha256]
	AuthenticationAlgorithm string `json:"authentication_algorithm,omitempty"`

	// The Diffie-Hellman group
	// Enum: [2 5 14]
	DhGroup int64 `json:"dh_group,omitempty"`

	// The encryption algorithm
	// Enum: [3des aes128 aes256]
	EncryptionAlgorithm string `json:"encryption_algorithm,omitempty"`

	// The IKE protocol version
	// Enum: [1 2]
	IkeVersion int64 `json:"ike_version,omitempty"`

	// The key lifetime in seconds
	// Maximum: 86400
	// Minimum: 300
	KeyLifetime int64 `json:"key_lifetime,omitempty"`

	// The name given to this IKE policy
	// Pattern: ^[A-Za-z][-A-Za-z0-9_]*$
	Name string `json:"name,omitempty"`

	// resource group
	ResourceGroup *IKEPolicyTemplateResourceGroup `json:"resource_group,omitempty"`
}

// Validate validates this i k e policy template
func (m *IKEPolicyTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthenticationAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDhGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEncryptionAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIkeVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKeyLifetime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateResourceGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var iKEPolicyTemplateTypeAuthenticationAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["md5","sha1","sha256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iKEPolicyTemplateTypeAuthenticationAlgorithmPropEnum = append(iKEPolicyTemplateTypeAuthenticationAlgorithmPropEnum, v)
	}
}

const (

	// IKEPolicyTemplateAuthenticationAlgorithmMd5 captures enum value "md5"
	IKEPolicyTemplateAuthenticationAlgorithmMd5 string = "md5"

	// IKEPolicyTemplateAuthenticationAlgorithmSha1 captures enum value "sha1"
	IKEPolicyTemplateAuthenticationAlgorithmSha1 string = "sha1"

	// IKEPolicyTemplateAuthenticationAlgorithmSha256 captures enum value "sha256"
	IKEPolicyTemplateAuthenticationAlgorithmSha256 string = "sha256"
)

// prop value enum
func (m *IKEPolicyTemplate) validateAuthenticationAlgorithmEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, iKEPolicyTemplateTypeAuthenticationAlgorithmPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IKEPolicyTemplate) validateAuthenticationAlgorithm(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthenticationAlgorithm) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthenticationAlgorithmEnum("authentication_algorithm", "body", m.AuthenticationAlgorithm); err != nil {
		return err
	}

	return nil
}

var iKEPolicyTemplateTypeDhGroupPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[2,5,14]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iKEPolicyTemplateTypeDhGroupPropEnum = append(iKEPolicyTemplateTypeDhGroupPropEnum, v)
	}
}

// prop value enum
func (m *IKEPolicyTemplate) validateDhGroupEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, iKEPolicyTemplateTypeDhGroupPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IKEPolicyTemplate) validateDhGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.DhGroup) { // not required
		return nil
	}

	// value enum
	if err := m.validateDhGroupEnum("dh_group", "body", m.DhGroup); err != nil {
		return err
	}

	return nil
}

var iKEPolicyTemplateTypeEncryptionAlgorithmPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["3des","aes128","aes256"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iKEPolicyTemplateTypeEncryptionAlgorithmPropEnum = append(iKEPolicyTemplateTypeEncryptionAlgorithmPropEnum, v)
	}
}

const (

	// IKEPolicyTemplateEncryptionAlgorithmNr3des captures enum value "3des"
	IKEPolicyTemplateEncryptionAlgorithmNr3des string = "3des"

	// IKEPolicyTemplateEncryptionAlgorithmAes128 captures enum value "aes128"
	IKEPolicyTemplateEncryptionAlgorithmAes128 string = "aes128"

	// IKEPolicyTemplateEncryptionAlgorithmAes256 captures enum value "aes256"
	IKEPolicyTemplateEncryptionAlgorithmAes256 string = "aes256"
)

// prop value enum
func (m *IKEPolicyTemplate) validateEncryptionAlgorithmEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, iKEPolicyTemplateTypeEncryptionAlgorithmPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IKEPolicyTemplate) validateEncryptionAlgorithm(formats strfmt.Registry) error {

	if swag.IsZero(m.EncryptionAlgorithm) { // not required
		return nil
	}

	// value enum
	if err := m.validateEncryptionAlgorithmEnum("encryption_algorithm", "body", m.EncryptionAlgorithm); err != nil {
		return err
	}

	return nil
}

var iKEPolicyTemplateTypeIkeVersionPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		iKEPolicyTemplateTypeIkeVersionPropEnum = append(iKEPolicyTemplateTypeIkeVersionPropEnum, v)
	}
}

// prop value enum
func (m *IKEPolicyTemplate) validateIkeVersionEnum(path, location string, value int64) error {
	if err := validate.Enum(path, location, value, iKEPolicyTemplateTypeIkeVersionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IKEPolicyTemplate) validateIkeVersion(formats strfmt.Registry) error {

	if swag.IsZero(m.IkeVersion) { // not required
		return nil
	}

	// value enum
	if err := m.validateIkeVersionEnum("ike_version", "body", m.IkeVersion); err != nil {
		return err
	}

	return nil
}

func (m *IKEPolicyTemplate) validateKeyLifetime(formats strfmt.Registry) error {

	if swag.IsZero(m.KeyLifetime) { // not required
		return nil
	}

	if err := validate.MinimumInt("key_lifetime", "body", int64(m.KeyLifetime), 300, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("key_lifetime", "body", int64(m.KeyLifetime), 86400, false); err != nil {
		return err
	}

	return nil
}

func (m *IKEPolicyTemplate) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.Pattern("name", "body", string(m.Name), `^[A-Za-z][-A-Za-z0-9_]*$`); err != nil {
		return err
	}

	return nil
}

func (m *IKEPolicyTemplate) validateResourceGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.ResourceGroup) { // not required
		return nil
	}

	if m.ResourceGroup != nil {
		if err := m.ResourceGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("resource_group")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IKEPolicyTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IKEPolicyTemplate) UnmarshalBinary(b []byte) error {
	var res IKEPolicyTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
