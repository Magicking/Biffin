package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

// MapFile map file
// swagger:model MapFile
type MapFile struct {

	// Hex-formatted color (#RRGGBB) (Optional)
	Backgroundcolor string `json:"backgroundcolor,omitempty"`

	// Number of tile rows
	Height int64 `json:"height,omitempty"`

	// Array of Layers
	Layers []*Layer `json:"layers"`

	// Auto-increments for each placed object
	Nextobjectid int64 `json:"nextobjectid,omitempty"`

	// Orthogonal, isometric, or staggered
	Orientation string `json:"orientation,omitempty"`

	// String key-value pairs
	Properties interface{} `json:"properties,omitempty"`

	// Rendering direction (orthogonal maps only)
	Renderorder string `json:"renderorder,omitempty"`

	// Map grid height.
	Tileheight int64 `json:"tileheight,omitempty"`

	// Array of Tilesets
	Tilesets []*Tileset `json:"tilesets"`

	// Map grid width.
	Tilewidth int64 `json:"tilewidth,omitempty"`

	// Number of tile columns
	Width int64 `json:"width,omitempty"`
}

// Validate validates this map file
func (m *MapFile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLayers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTilesets(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MapFile) validateLayers(formats strfmt.Registry) error {

	if swag.IsZero(m.Layers) { // not required
		return nil
	}

	for i := 0; i < len(m.Layers); i++ {

		if swag.IsZero(m.Layers[i]) { // not required
			continue
		}

		if m.Layers[i] != nil {

			if err := m.Layers[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *MapFile) validateTilesets(formats strfmt.Registry) error {

	if swag.IsZero(m.Tilesets) { // not required
		return nil
	}

	for i := 0; i < len(m.Tilesets); i++ {

		if swag.IsZero(m.Tilesets[i]) { // not required
			continue
		}

		if m.Tilesets[i] != nil {

			if err := m.Tilesets[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}