// Code generated by go-swagger; DO NOT EDIT.

// Copyright (c) 2020 Samuel Mutel <12967891+smutel@users.noreply.github.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// VRF v r f
//
// swagger:model VRF
type VRF struct {

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// Display name
	// Read Only: true
	DisplayName string `json:"display_name,omitempty"`

	// Enforce unique space
	//
	// Prevent duplicate prefixes/IP addresses within this VRF
	EnforceUnique bool `json:"enforce_unique,omitempty"`

	// export targets
	// Unique: true
	ExportTargets []*NestedRouteTarget `json:"export_targets"`

	// Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// import targets
	// Unique: true
	ImportTargets []*NestedRouteTarget `json:"import_targets"`

	// Ipaddress count
	// Read Only: true
	IpaddressCount int64 `json:"ipaddress_count,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Name
	// Required: true
	// Max Length: 100
	// Min Length: 1
	Name *string `json:"name"`

	// Prefix count
	// Read Only: true
	PrefixCount int64 `json:"prefix_count,omitempty"`

	// Route distinguisher
	//
	// Unique route distinguisher (as defined in RFC 4364)
	// Max Length: 21
	Rd *string `json:"rd,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// tenant
	Tenant *NestedTenant `json:"tenant,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this v r f
func (m *VRF) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExportTargets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImportTargets(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTenant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VRF) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VRF) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *VRF) validateExportTargets(formats strfmt.Registry) error {
	if swag.IsZero(m.ExportTargets) { // not required
		return nil
	}

	if err := validate.UniqueItems("export_targets", "body", m.ExportTargets); err != nil {
		return err
	}

	for i := 0; i < len(m.ExportTargets); i++ {
		if swag.IsZero(m.ExportTargets[i]) { // not required
			continue
		}

		if m.ExportTargets[i] != nil {
			if err := m.ExportTargets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("export_targets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("export_targets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VRF) validateImportTargets(formats strfmt.Registry) error {
	if swag.IsZero(m.ImportTargets) { // not required
		return nil
	}

	if err := validate.UniqueItems("import_targets", "body", m.ImportTargets); err != nil {
		return err
	}

	for i := 0; i < len(m.ImportTargets); i++ {
		if swag.IsZero(m.ImportTargets[i]) { // not required
			continue
		}

		if m.ImportTargets[i] != nil {
			if err := m.ImportTargets[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("import_targets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("import_targets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VRF) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *VRF) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 100); err != nil {
		return err
	}

	return nil
}

func (m *VRF) validateRd(formats strfmt.Registry) error {
	if swag.IsZero(m.Rd) { // not required
		return nil
	}

	if err := validate.MaxLength("rd", "body", *m.Rd, 21); err != nil {
		return err
	}

	return nil
}

func (m *VRF) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VRF) validateTenant(formats strfmt.Registry) error {
	if swag.IsZero(m.Tenant) { // not required
		return nil
	}

	if m.Tenant != nil {
		if err := m.Tenant.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *VRF) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this v r f based on the context it is used
func (m *VRF) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplayName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExportTargets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateImportTargets(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIpaddressCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrefixCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTenant(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VRF) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.Date(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *VRF) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *VRF) contextValidateDisplayName(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display_name", "body", string(m.DisplayName)); err != nil {
		return err
	}

	return nil
}

func (m *VRF) contextValidateExportTargets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ExportTargets); i++ {

		if m.ExportTargets[i] != nil {
			if err := m.ExportTargets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("export_targets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("export_targets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VRF) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *VRF) contextValidateImportTargets(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.ImportTargets); i++ {

		if m.ImportTargets[i] != nil {
			if err := m.ImportTargets[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("import_targets" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("import_targets" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VRF) contextValidateIpaddressCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "ipaddress_count", "body", int64(m.IpaddressCount)); err != nil {
		return err
	}

	return nil
}

func (m *VRF) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *VRF) contextValidatePrefixCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "prefix_count", "body", int64(m.PrefixCount)); err != nil {
		return err
	}

	return nil
}

func (m *VRF) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *VRF) contextValidateTenant(ctx context.Context, formats strfmt.Registry) error {

	if m.Tenant != nil {
		if err := m.Tenant.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("tenant")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("tenant")
			}
			return err
		}
	}

	return nil
}

func (m *VRF) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VRF) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VRF) UnmarshalBinary(b []byte) error {
	var res VRF
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
