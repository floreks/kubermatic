// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AuthInfo AuthInfo contains information that describes identity information.  This is use to tell the kubernetes cluster who you are.
// swagger:model AuthInfo
type AuthInfo struct {

	// ClientCertificate is the path to a client cert file for TLS.
	// +optional
	ClientCertificate string `json:"client-certificate,omitempty"`

	// ClientCertificateData contains PEM-encoded data from a client cert file for TLS. Overrides ClientCertificate
	// +optional
	ClientCertificateData []uint8 `json:"client-certificate-data"`

	// ClientKey is the path to a client key file for TLS.
	// +optional
	ClientKey string `json:"client-key,omitempty"`

	// ClientKeyData contains PEM-encoded data from a client key file for TLS. Overrides ClientKey
	// +optional
	ClientKeyData []uint8 `json:"client-key-data"`

	// Extensions holds additional information. This is useful for extenders so that reads and writes don't clobber unknown fields
	// +optional
	Extensions []*NamedExtension `json:"extensions"`

	// Impersonate is the username to imperonate.  The name matches the flag.
	// +optional
	Impersonate string `json:"as,omitempty"`

	// ImpersonateGroups is the groups to imperonate.
	// +optional
	ImpersonateGroups []string `json:"as-groups"`

	// ImpersonateUserExtra contains additional information for impersonated user.
	// +optional
	ImpersonateUserExtra map[string][]string `json:"as-user-extra,omitempty"`

	// Password is the password for basic authentication to the kubernetes cluster.
	// +optional
	Password string `json:"password,omitempty"`

	// Token is the bearer token for authentication to the kubernetes cluster.
	// +optional
	Token string `json:"token,omitempty"`

	// TokenFile is a pointer to a file that contains a bearer token (as described above).  If both Token and TokenFile are present, Token takes precedence.
	// +optional
	TokenFile string `json:"tokenFile,omitempty"`

	// Username is the username for basic authentication to the kubernetes cluster.
	// +optional
	Username string `json:"username,omitempty"`

	// auth provider
	AuthProvider *AuthProviderConfig `json:"auth-provider,omitempty"`

	// exec
	Exec *ExecConfig `json:"exec,omitempty"`
}

// Validate validates this auth info
func (m *AuthInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExec(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AuthInfo) validateExtensions(formats strfmt.Registry) error {

	if swag.IsZero(m.Extensions) { // not required
		return nil
	}

	for i := 0; i < len(m.Extensions); i++ {
		if swag.IsZero(m.Extensions[i]) { // not required
			continue
		}

		if m.Extensions[i] != nil {
			if err := m.Extensions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("extensions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AuthInfo) validateAuthProvider(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthProvider) { // not required
		return nil
	}

	if m.AuthProvider != nil {
		if err := m.AuthProvider.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("auth-provider")
			}
			return err
		}
	}

	return nil
}

func (m *AuthInfo) validateExec(formats strfmt.Registry) error {

	if swag.IsZero(m.Exec) { // not required
		return nil
	}

	if m.Exec != nil {
		if err := m.Exec.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("exec")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuthInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuthInfo) UnmarshalBinary(b []byte) error {
	var res AuthInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
