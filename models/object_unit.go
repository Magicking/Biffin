// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ObjectUnit object unit
// swagger:model ObjectUnit
type ObjectUnit struct {

	// GID, only if object comes from a Tilemap
	Gid int64 `json:"gid,omitempty"`

	// Height in pixels. Ignored if using a gid.
	Height int64 `json:"height,omitempty"`

	// Incremental id - unique across all objects
	ID int64 `json:"id,omitempty"`

	// String assigned to name field in editor
	Name string `json:"name,omitempty"`

	// String key-value pairs
	Properties interface{} `json:"properties,omitempty"`

	// Angle in degrees clockwise
	Rotation float64 `json:"rotation,omitempty"`

	// String assigned to type field in editor
	Type string `json:"type,omitempty"`

	// Whether object is shown in editor.
	Visible bool `json:"visible,omitempty"`

	// Width in pixels. Ignored if using a gid.
	Width int64 `json:"width,omitempty"`

	// x coordinate in pixels
	X int64 `json:"x,omitempty"`

	// y coordinate in pixels
	Y int64 `json:"y,omitempty"`
}

// Validate validates this object unit
func (m *ObjectUnit) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ObjectUnit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObjectUnit) UnmarshalBinary(b []byte) error {
	var res ObjectUnit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
