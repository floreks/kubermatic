// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DigitaloceanNodeSpec DigitaloceanNodeSpec digitalocean node settings
// swagger:model DigitaloceanNodeSpec
type DigitaloceanNodeSpec struct {

	// enable backups for the droplet
	Backups bool `json:"backups,omitempty"`

	// enable ipv6 for the droplet
	IPV6 bool `json:"ipv6,omitempty"`

	// enable monitoring for the droplet
	Monitoring bool `json:"monitoring,omitempty"`

	// droplet size slug
	// Required: true
	Size *string `json:"size"`

	// additional droplet tags
	Tags []string `json:"tags"`
}

// Validate validates this digitalocean node spec
func (m *DigitaloceanNodeSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DigitaloceanNodeSpec) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("size", "body", m.Size); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DigitaloceanNodeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DigitaloceanNodeSpec) UnmarshalBinary(b []byte) error {
	var res DigitaloceanNodeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
