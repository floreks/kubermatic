// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// OpenstackSize OpenstackSize is the object representing openstack's sizes.
// swagger:model OpenstackSize
type OpenstackSize struct {

	// Disk is the amount of root disk, measured in GB
	Disk int64 `json:"disk,omitempty"`

	// IsPublic indicates whether the size is public (available to all projects) or scoped to a set of projects
	IsPublic bool `json:"isPublic,omitempty"`

	// Memory is the amount of memory, measured in MB
	Memory int64 `json:"memory,omitempty"`

	// Region specifies the geographic region in which the size resides
	Region string `json:"region,omitempty"`

	// Slug holds  the name of the size
	Slug string `json:"slug,omitempty"`

	// Swap is the amount of swap space, measured in MB
	Swap int64 `json:"swap,omitempty"`

	// VCPUs indicates how many (virtual) CPUs are available for this flavor
	VCpus int64 `json:"vcpus,omitempty"`
}

// Validate validates this openstack size
func (m *OpenstackSize) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OpenstackSize) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenstackSize) UnmarshalBinary(b []byte) error {
	var res OpenstackSize
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
