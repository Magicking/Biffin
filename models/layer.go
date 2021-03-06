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

// Layer layer
// swagger:model Layer
type Layer struct {

	// Array of GIDs. tilelayer only.
	Data []int64 `json:"data"`

	// topdown (default) or index. objectgroup only.
	Draworder string `json:"draworder,omitempty"`

	// Row count. Same as map height in Tiled Qt.
	Height int64 `json:"height,omitempty"`

	// Name assigned to this layer
	Name string `json:"name,omitempty"`

	// Array of Object units. objectgroup only.
	Objects []*ObjectUnit `json:"objects"`

	// Value between 0 and 1
	Opacity float64 `json:"opacity,omitempty"`

	// string key-value pairs.
	Properties interface{} `json:"properties,omitempty"`

	// tilelayer, objectgroup, or imagelayer
	Type string `json:"type,omitempty"`

	// Whether layer is shown or hidden in editor
	Visible bool `json:"visible,omitempty"`

	// Column count. Same as map width in Tiled Qt.
	Width int64 `json:"width,omitempty"`

	// Horizontal layer offset. Always 0 in Tiled Qt.
	X int64 `json:"x,omitempty"`

	// Vertical layer offset. Always 0 in Tiled Qt.
	Y int64 `json:"y,omitempty"`
}

// Validate validates this layer
func (m *Layer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Layer) validateObjects(formats strfmt.Registry) error {

	if swag.IsZero(m.Objects) { // not required
		return nil
	}

	for i := 0; i < len(m.Objects); i++ {
		if swag.IsZero(m.Objects[i]) { // not required
			continue
		}

		if m.Objects[i] != nil {
			if err := m.Objects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Layer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Layer) UnmarshalBinary(b []byte) error {
	var res Layer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
