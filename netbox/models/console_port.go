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
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ConsolePort console port
//
// swagger:model ConsolePort
type ConsolePort struct {

	// occupied
	// Read Only: true
	Occupied *bool `json:"_occupied,omitempty"`

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Cable peer
	//
	//
	// Return the appropriate serializer for the cable termination model.
	//
	// Read Only: true
	CablePeer map[string]*string `json:"cable_peer,omitempty"`

	// Cable peer type
	// Read Only: true
	CablePeerType string `json:"cable_peer_type,omitempty"`

	// Connected endpoint
	//
	//
	// Return the appropriate serializer for the type of connected object.
	//
	// Read Only: true
	ConnectedEndpoint map[string]*string `json:"connected_endpoint,omitempty"`

	// Connected endpoint reachable
	// Read Only: true
	ConnectedEndpointReachable *bool `json:"connected_endpoint_reachable,omitempty"`

	// Connected endpoint type
	// Read Only: true
	ConnectedEndpointType string `json:"connected_endpoint_type,omitempty"`

	// Created
	// Read Only: true
	// Format: date
	Created strfmt.Date `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// device
	// Required: true
	Device *NestedDevice `json:"device"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Label
	//
	// Physical label
	// Max Length: 64
	Label string `json:"label,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated strfmt.DateTime `json:"last_updated,omitempty"`

	// Mark connected
	//
	// Treat as if a cable is connected
	MarkConnected bool `json:"mark_connected,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// speed
	Speed *ConsolePortSpeed `json:"speed,omitempty"`

	// tags
	Tags []*NestedTag `json:"tags"`

	// type
	Type *ConsolePortType `json:"type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this console port
func (m *ConsolePort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpeed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
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

func (m *ConsolePort) validateCable(formats strfmt.Registry) error {
	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *ConsolePort) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConsolePort) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *ConsolePort) validateDevice(formats strfmt.Registry) error {

	if err := validate.Required("device", "body", m.Device); err != nil {
		return err
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *ConsolePort) validateLabel(formats strfmt.Registry) error {
	if swag.IsZero(m.Label) { // not required
		return nil
	}

	if err := validate.MaxLength("label", "body", m.Label, 64); err != nil {
		return err
	}

	return nil
}

func (m *ConsolePort) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ConsolePort) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 64); err != nil {
		return err
	}

	return nil
}

func (m *ConsolePort) validateSpeed(formats strfmt.Registry) error {
	if swag.IsZero(m.Speed) { // not required
		return nil
	}

	if m.Speed != nil {
		if err := m.Speed.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("speed")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("speed")
			}
			return err
		}
	}

	return nil
}

func (m *ConsolePort) validateTags(formats strfmt.Registry) error {
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

func (m *ConsolePort) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if m.Type != nil {
		if err := m.Type.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *ConsolePort) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this console port based on the context it is used
func (m *ConsolePort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOccupied(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCablePeer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCablePeerType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpoint(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpointReachable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectedEndpointType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpeed(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
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

func (m *ConsolePort) contextValidateOccupied(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "_occupied", "body", m.Occupied); err != nil {
		return err
	}

	return nil
}

func (m *ConsolePort) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

	if m.Cable != nil {
		if err := m.Cable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *ConsolePort) contextValidateCablePeer(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ConsolePort) contextValidateCablePeerType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cable_peer_type", "body", string(m.CablePeerType)); err != nil {
		return err
	}

	return nil
}

func (m *ConsolePort) contextValidateConnectedEndpoint(ctx context.Context, formats strfmt.Registry) error {

	return nil
}

func (m *ConsolePort) contextValidateConnectedEndpointReachable(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoint_reachable", "body", m.ConnectedEndpointReachable); err != nil {
		return err
	}

	return nil
}

func (m *ConsolePort) contextValidateConnectedEndpointType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connected_endpoint_type", "body", string(m.ConnectedEndpointType)); err != nil {
		return err
	}

	return nil
}

func (m *ConsolePort) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.Date(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *ConsolePort) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.Device != nil {
		if err := m.Device.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *ConsolePort) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *ConsolePort) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *ConsolePort) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", strfmt.DateTime(m.LastUpdated)); err != nil {
		return err
	}

	return nil
}

func (m *ConsolePort) contextValidateSpeed(ctx context.Context, formats strfmt.Registry) error {

	if m.Speed != nil {
		if err := m.Speed.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("speed")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("speed")
			}
			return err
		}
	}

	return nil
}

func (m *ConsolePort) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ConsolePort) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if m.Type != nil {
		if err := m.Type.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("type")
			}
			return err
		}
	}

	return nil
}

func (m *ConsolePort) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsolePort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsolePort) UnmarshalBinary(b []byte) error {
	var res ConsolePort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConsolePortSpeed Speed
//
// swagger:model ConsolePortSpeed
type ConsolePortSpeed struct {

	// label
	// Required: true
	// Enum: [1200 bps 2400 bps 4800 bps 9600 bps 19.2 kbps 38.4 kbps 57.6 kbps 115.2 kbps]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [1200 2400 4800 9600 19200 38400 57600 115200]
	Value *int64 `json:"value"`
}

// Validate validates this console port speed
func (m *ConsolePortSpeed) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var consolePortSpeedTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["1200 bps","2400 bps","4800 bps","9600 bps","19.2 kbps","38.4 kbps","57.6 kbps","115.2 kbps"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consolePortSpeedTypeLabelPropEnum = append(consolePortSpeedTypeLabelPropEnum, v)
	}
}

const (

	// ConsolePortSpeedLabelNr1200Bps captures enum value "1200 bps"
	ConsolePortSpeedLabelNr1200Bps string = "1200 bps"

	// ConsolePortSpeedLabelNr2400Bps captures enum value "2400 bps"
	ConsolePortSpeedLabelNr2400Bps string = "2400 bps"

	// ConsolePortSpeedLabelNr4800Bps captures enum value "4800 bps"
	ConsolePortSpeedLabelNr4800Bps string = "4800 bps"

	// ConsolePortSpeedLabelNr9600Bps captures enum value "9600 bps"
	ConsolePortSpeedLabelNr9600Bps string = "9600 bps"

	// ConsolePortSpeedLabelNr19Dot2Kbps captures enum value "19.2 kbps"
	ConsolePortSpeedLabelNr19Dot2Kbps string = "19.2 kbps"

	// ConsolePortSpeedLabelNr38Dot4Kbps captures enum value "38.4 kbps"
	ConsolePortSpeedLabelNr38Dot4Kbps string = "38.4 kbps"

	// ConsolePortSpeedLabelNr57Dot6Kbps captures enum value "57.6 kbps"
	ConsolePortSpeedLabelNr57Dot6Kbps string = "57.6 kbps"

	// ConsolePortSpeedLabelNr115Dot2Kbps captures enum value "115.2 kbps"
	ConsolePortSpeedLabelNr115Dot2Kbps string = "115.2 kbps"
)

// prop value enum
func (m *ConsolePortSpeed) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consolePortSpeedTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsolePortSpeed) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("speed"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("speed"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var consolePortSpeedTypeValuePropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1200,2400,4800,9600,19200,38400,57600,115200]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consolePortSpeedTypeValuePropEnum = append(consolePortSpeedTypeValuePropEnum, v)
	}
}

// prop value enum
func (m *ConsolePortSpeed) validateValueEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, consolePortSpeedTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsolePortSpeed) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("speed"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("speed"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this console port speed based on context it is used
func (m *ConsolePortSpeed) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConsolePortSpeed) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsolePortSpeed) UnmarshalBinary(b []byte) error {
	var res ConsolePortSpeed
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConsolePortType Type
//
// swagger:model ConsolePortType
type ConsolePortType struct {

	// label
	// Required: true
	// Enum: [DE-9 DB-25 RJ-11 RJ-12 RJ-45 USB Type A USB Type B USB Type C USB Mini A USB Mini B USB Micro A USB Micro B Other]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [de-9 db-25 rj-11 rj-12 rj-45 usb-a usb-b usb-c usb-mini-a usb-mini-b usb-micro-a usb-micro-b other]
	Value *string `json:"value"`
}

// Validate validates this console port type
func (m *ConsolePortType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var consolePortTypeTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DE-9","DB-25","RJ-11","RJ-12","RJ-45","USB Type A","USB Type B","USB Type C","USB Mini A","USB Mini B","USB Micro A","USB Micro B","Other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consolePortTypeTypeLabelPropEnum = append(consolePortTypeTypeLabelPropEnum, v)
	}
}

const (

	// ConsolePortTypeLabelDEDash9 captures enum value "DE-9"
	ConsolePortTypeLabelDEDash9 string = "DE-9"

	// ConsolePortTypeLabelDBDash25 captures enum value "DB-25"
	ConsolePortTypeLabelDBDash25 string = "DB-25"

	// ConsolePortTypeLabelRJDash11 captures enum value "RJ-11"
	ConsolePortTypeLabelRJDash11 string = "RJ-11"

	// ConsolePortTypeLabelRJDash12 captures enum value "RJ-12"
	ConsolePortTypeLabelRJDash12 string = "RJ-12"

	// ConsolePortTypeLabelRJDash45 captures enum value "RJ-45"
	ConsolePortTypeLabelRJDash45 string = "RJ-45"

	// ConsolePortTypeLabelUSBTypeA captures enum value "USB Type A"
	ConsolePortTypeLabelUSBTypeA string = "USB Type A"

	// ConsolePortTypeLabelUSBTypeB captures enum value "USB Type B"
	ConsolePortTypeLabelUSBTypeB string = "USB Type B"

	// ConsolePortTypeLabelUSBTypeC captures enum value "USB Type C"
	ConsolePortTypeLabelUSBTypeC string = "USB Type C"

	// ConsolePortTypeLabelUSBMiniA captures enum value "USB Mini A"
	ConsolePortTypeLabelUSBMiniA string = "USB Mini A"

	// ConsolePortTypeLabelUSBMiniB captures enum value "USB Mini B"
	ConsolePortTypeLabelUSBMiniB string = "USB Mini B"

	// ConsolePortTypeLabelUSBMicroA captures enum value "USB Micro A"
	ConsolePortTypeLabelUSBMicroA string = "USB Micro A"

	// ConsolePortTypeLabelUSBMicroB captures enum value "USB Micro B"
	ConsolePortTypeLabelUSBMicroB string = "USB Micro B"

	// ConsolePortTypeLabelOther captures enum value "Other"
	ConsolePortTypeLabelOther string = "Other"
)

// prop value enum
func (m *ConsolePortType) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consolePortTypeTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsolePortType) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("type"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var consolePortTypeTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["de-9","db-25","rj-11","rj-12","rj-45","usb-a","usb-b","usb-c","usb-mini-a","usb-mini-b","usb-micro-a","usb-micro-b","other"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		consolePortTypeTypeValuePropEnum = append(consolePortTypeTypeValuePropEnum, v)
	}
}

const (

	// ConsolePortTypeValueDeDash9 captures enum value "de-9"
	ConsolePortTypeValueDeDash9 string = "de-9"

	// ConsolePortTypeValueDbDash25 captures enum value "db-25"
	ConsolePortTypeValueDbDash25 string = "db-25"

	// ConsolePortTypeValueRjDash11 captures enum value "rj-11"
	ConsolePortTypeValueRjDash11 string = "rj-11"

	// ConsolePortTypeValueRjDash12 captures enum value "rj-12"
	ConsolePortTypeValueRjDash12 string = "rj-12"

	// ConsolePortTypeValueRjDash45 captures enum value "rj-45"
	ConsolePortTypeValueRjDash45 string = "rj-45"

	// ConsolePortTypeValueUsbDasha captures enum value "usb-a"
	ConsolePortTypeValueUsbDasha string = "usb-a"

	// ConsolePortTypeValueUsbDashb captures enum value "usb-b"
	ConsolePortTypeValueUsbDashb string = "usb-b"

	// ConsolePortTypeValueUsbDashc captures enum value "usb-c"
	ConsolePortTypeValueUsbDashc string = "usb-c"

	// ConsolePortTypeValueUsbDashMiniDasha captures enum value "usb-mini-a"
	ConsolePortTypeValueUsbDashMiniDasha string = "usb-mini-a"

	// ConsolePortTypeValueUsbDashMiniDashb captures enum value "usb-mini-b"
	ConsolePortTypeValueUsbDashMiniDashb string = "usb-mini-b"

	// ConsolePortTypeValueUsbDashMicroDasha captures enum value "usb-micro-a"
	ConsolePortTypeValueUsbDashMicroDasha string = "usb-micro-a"

	// ConsolePortTypeValueUsbDashMicroDashb captures enum value "usb-micro-b"
	ConsolePortTypeValueUsbDashMicroDashb string = "usb-micro-b"

	// ConsolePortTypeValueOther captures enum value "other"
	ConsolePortTypeValueOther string = "other"
)

// prop value enum
func (m *ConsolePortType) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, consolePortTypeTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ConsolePortType) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("type"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("type"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this console port type based on context it is used
func (m *ConsolePortType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConsolePortType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsolePortType) UnmarshalBinary(b []byte) error {
	var res ConsolePortType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
