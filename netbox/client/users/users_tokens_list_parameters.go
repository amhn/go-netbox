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

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewUsersTokensListParams creates a new UsersTokensListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUsersTokensListParams() *UsersTokensListParams {
	return &UsersTokensListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUsersTokensListParamsWithTimeout creates a new UsersTokensListParams object
// with the ability to set a timeout on a request.
func NewUsersTokensListParamsWithTimeout(timeout time.Duration) *UsersTokensListParams {
	return &UsersTokensListParams{
		timeout: timeout,
	}
}

// NewUsersTokensListParamsWithContext creates a new UsersTokensListParams object
// with the ability to set a context for a request.
func NewUsersTokensListParamsWithContext(ctx context.Context) *UsersTokensListParams {
	return &UsersTokensListParams{
		Context: ctx,
	}
}

// NewUsersTokensListParamsWithHTTPClient creates a new UsersTokensListParams object
// with the ability to set a custom HTTPClient for a request.
func NewUsersTokensListParamsWithHTTPClient(client *http.Client) *UsersTokensListParams {
	return &UsersTokensListParams{
		HTTPClient: client,
	}
}

/*
UsersTokensListParams contains all the parameters to send to the API endpoint

	for the users tokens list operation.

	Typically these are written to a http.Request.
*/
type UsersTokensListParams struct {

	// Created.
	Created *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLte.
	CreatedLte *string

	// Description.
	Description *string

	// DescriptionEmpty.
	DescriptionEmpty *string

	// DescriptionIc.
	DescriptionIc *string

	// DescriptionIe.
	DescriptionIe *string

	// DescriptionIew.
	DescriptionIew *string

	// DescriptionIsw.
	DescriptionIsw *string

	// Descriptionn.
	Descriptionn *string

	// DescriptionNic.
	DescriptionNic *string

	// DescriptionNie.
	DescriptionNie *string

	// DescriptionNiew.
	DescriptionNiew *string

	// DescriptionNisw.
	DescriptionNisw *string

	// Expires.
	Expires *string

	// ExpiresGte.
	ExpiresGte *string

	// ExpiresLte.
	ExpiresLte *string

	// ID.
	ID *string

	// IDGt.
	IDGt *string

	// IDGte.
	IDGte *string

	// IDLt.
	IDLt *string

	// IDLte.
	IDLte *string

	// IDn.
	IDn *string

	// Key.
	Key *string

	// KeyEmpty.
	KeyEmpty *string

	// KeyIc.
	KeyIc *string

	// KeyIe.
	KeyIe *string

	// KeyIew.
	KeyIew *string

	// KeyIsw.
	KeyIsw *string

	// Keyn.
	Keyn *string

	// KeyNic.
	KeyNic *string

	// KeyNie.
	KeyNie *string

	// KeyNiew.
	KeyNiew *string

	// KeyNisw.
	KeyNisw *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	/* Ordering.

	   Which field to use when ordering the results.
	*/
	Ordering *string

	// Q.
	Q *string

	// User.
	User *string

	// Usern.
	Usern *string

	// UserID.
	UserID *string

	// UserIDn.
	UserIDn *string

	// WriteEnabled.
	WriteEnabled *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the users tokens list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersTokensListParams) WithDefaults() *UsersTokensListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the users tokens list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UsersTokensListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the users tokens list params
func (o *UsersTokensListParams) WithTimeout(timeout time.Duration) *UsersTokensListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the users tokens list params
func (o *UsersTokensListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the users tokens list params
func (o *UsersTokensListParams) WithContext(ctx context.Context) *UsersTokensListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the users tokens list params
func (o *UsersTokensListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the users tokens list params
func (o *UsersTokensListParams) WithHTTPClient(client *http.Client) *UsersTokensListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the users tokens list params
func (o *UsersTokensListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the users tokens list params
func (o *UsersTokensListParams) WithCreated(created *string) *UsersTokensListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the users tokens list params
func (o *UsersTokensListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGte adds the createdGte to the users tokens list params
func (o *UsersTokensListParams) WithCreatedGte(createdGte *string) *UsersTokensListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the users tokens list params
func (o *UsersTokensListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLte adds the createdLte to the users tokens list params
func (o *UsersTokensListParams) WithCreatedLte(createdLte *string) *UsersTokensListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the users tokens list params
func (o *UsersTokensListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithDescription adds the description to the users tokens list params
func (o *UsersTokensListParams) WithDescription(description *string) *UsersTokensListParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the users tokens list params
func (o *UsersTokensListParams) SetDescription(description *string) {
	o.Description = description
}

// WithDescriptionEmpty adds the descriptionEmpty to the users tokens list params
func (o *UsersTokensListParams) WithDescriptionEmpty(descriptionEmpty *string) *UsersTokensListParams {
	o.SetDescriptionEmpty(descriptionEmpty)
	return o
}

// SetDescriptionEmpty adds the descriptionEmpty to the users tokens list params
func (o *UsersTokensListParams) SetDescriptionEmpty(descriptionEmpty *string) {
	o.DescriptionEmpty = descriptionEmpty
}

// WithDescriptionIc adds the descriptionIc to the users tokens list params
func (o *UsersTokensListParams) WithDescriptionIc(descriptionIc *string) *UsersTokensListParams {
	o.SetDescriptionIc(descriptionIc)
	return o
}

// SetDescriptionIc adds the descriptionIc to the users tokens list params
func (o *UsersTokensListParams) SetDescriptionIc(descriptionIc *string) {
	o.DescriptionIc = descriptionIc
}

// WithDescriptionIe adds the descriptionIe to the users tokens list params
func (o *UsersTokensListParams) WithDescriptionIe(descriptionIe *string) *UsersTokensListParams {
	o.SetDescriptionIe(descriptionIe)
	return o
}

// SetDescriptionIe adds the descriptionIe to the users tokens list params
func (o *UsersTokensListParams) SetDescriptionIe(descriptionIe *string) {
	o.DescriptionIe = descriptionIe
}

// WithDescriptionIew adds the descriptionIew to the users tokens list params
func (o *UsersTokensListParams) WithDescriptionIew(descriptionIew *string) *UsersTokensListParams {
	o.SetDescriptionIew(descriptionIew)
	return o
}

// SetDescriptionIew adds the descriptionIew to the users tokens list params
func (o *UsersTokensListParams) SetDescriptionIew(descriptionIew *string) {
	o.DescriptionIew = descriptionIew
}

// WithDescriptionIsw adds the descriptionIsw to the users tokens list params
func (o *UsersTokensListParams) WithDescriptionIsw(descriptionIsw *string) *UsersTokensListParams {
	o.SetDescriptionIsw(descriptionIsw)
	return o
}

// SetDescriptionIsw adds the descriptionIsw to the users tokens list params
func (o *UsersTokensListParams) SetDescriptionIsw(descriptionIsw *string) {
	o.DescriptionIsw = descriptionIsw
}

// WithDescriptionn adds the descriptionn to the users tokens list params
func (o *UsersTokensListParams) WithDescriptionn(descriptionn *string) *UsersTokensListParams {
	o.SetDescriptionn(descriptionn)
	return o
}

// SetDescriptionn adds the descriptionN to the users tokens list params
func (o *UsersTokensListParams) SetDescriptionn(descriptionn *string) {
	o.Descriptionn = descriptionn
}

// WithDescriptionNic adds the descriptionNic to the users tokens list params
func (o *UsersTokensListParams) WithDescriptionNic(descriptionNic *string) *UsersTokensListParams {
	o.SetDescriptionNic(descriptionNic)
	return o
}

// SetDescriptionNic adds the descriptionNic to the users tokens list params
func (o *UsersTokensListParams) SetDescriptionNic(descriptionNic *string) {
	o.DescriptionNic = descriptionNic
}

// WithDescriptionNie adds the descriptionNie to the users tokens list params
func (o *UsersTokensListParams) WithDescriptionNie(descriptionNie *string) *UsersTokensListParams {
	o.SetDescriptionNie(descriptionNie)
	return o
}

// SetDescriptionNie adds the descriptionNie to the users tokens list params
func (o *UsersTokensListParams) SetDescriptionNie(descriptionNie *string) {
	o.DescriptionNie = descriptionNie
}

// WithDescriptionNiew adds the descriptionNiew to the users tokens list params
func (o *UsersTokensListParams) WithDescriptionNiew(descriptionNiew *string) *UsersTokensListParams {
	o.SetDescriptionNiew(descriptionNiew)
	return o
}

// SetDescriptionNiew adds the descriptionNiew to the users tokens list params
func (o *UsersTokensListParams) SetDescriptionNiew(descriptionNiew *string) {
	o.DescriptionNiew = descriptionNiew
}

// WithDescriptionNisw adds the descriptionNisw to the users tokens list params
func (o *UsersTokensListParams) WithDescriptionNisw(descriptionNisw *string) *UsersTokensListParams {
	o.SetDescriptionNisw(descriptionNisw)
	return o
}

// SetDescriptionNisw adds the descriptionNisw to the users tokens list params
func (o *UsersTokensListParams) SetDescriptionNisw(descriptionNisw *string) {
	o.DescriptionNisw = descriptionNisw
}

// WithExpires adds the expires to the users tokens list params
func (o *UsersTokensListParams) WithExpires(expires *string) *UsersTokensListParams {
	o.SetExpires(expires)
	return o
}

// SetExpires adds the expires to the users tokens list params
func (o *UsersTokensListParams) SetExpires(expires *string) {
	o.Expires = expires
}

// WithExpiresGte adds the expiresGte to the users tokens list params
func (o *UsersTokensListParams) WithExpiresGte(expiresGte *string) *UsersTokensListParams {
	o.SetExpiresGte(expiresGte)
	return o
}

// SetExpiresGte adds the expiresGte to the users tokens list params
func (o *UsersTokensListParams) SetExpiresGte(expiresGte *string) {
	o.ExpiresGte = expiresGte
}

// WithExpiresLte adds the expiresLte to the users tokens list params
func (o *UsersTokensListParams) WithExpiresLte(expiresLte *string) *UsersTokensListParams {
	o.SetExpiresLte(expiresLte)
	return o
}

// SetExpiresLte adds the expiresLte to the users tokens list params
func (o *UsersTokensListParams) SetExpiresLte(expiresLte *string) {
	o.ExpiresLte = expiresLte
}

// WithID adds the id to the users tokens list params
func (o *UsersTokensListParams) WithID(id *string) *UsersTokensListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the users tokens list params
func (o *UsersTokensListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the users tokens list params
func (o *UsersTokensListParams) WithIDGt(iDGt *string) *UsersTokensListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the users tokens list params
func (o *UsersTokensListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the users tokens list params
func (o *UsersTokensListParams) WithIDGte(iDGte *string) *UsersTokensListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the users tokens list params
func (o *UsersTokensListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the users tokens list params
func (o *UsersTokensListParams) WithIDLt(iDLt *string) *UsersTokensListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the users tokens list params
func (o *UsersTokensListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the users tokens list params
func (o *UsersTokensListParams) WithIDLte(iDLte *string) *UsersTokensListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the users tokens list params
func (o *UsersTokensListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the users tokens list params
func (o *UsersTokensListParams) WithIDn(iDn *string) *UsersTokensListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the users tokens list params
func (o *UsersTokensListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithKey adds the key to the users tokens list params
func (o *UsersTokensListParams) WithKey(key *string) *UsersTokensListParams {
	o.SetKey(key)
	return o
}

// SetKey adds the key to the users tokens list params
func (o *UsersTokensListParams) SetKey(key *string) {
	o.Key = key
}

// WithKeyEmpty adds the keyEmpty to the users tokens list params
func (o *UsersTokensListParams) WithKeyEmpty(keyEmpty *string) *UsersTokensListParams {
	o.SetKeyEmpty(keyEmpty)
	return o
}

// SetKeyEmpty adds the keyEmpty to the users tokens list params
func (o *UsersTokensListParams) SetKeyEmpty(keyEmpty *string) {
	o.KeyEmpty = keyEmpty
}

// WithKeyIc adds the keyIc to the users tokens list params
func (o *UsersTokensListParams) WithKeyIc(keyIc *string) *UsersTokensListParams {
	o.SetKeyIc(keyIc)
	return o
}

// SetKeyIc adds the keyIc to the users tokens list params
func (o *UsersTokensListParams) SetKeyIc(keyIc *string) {
	o.KeyIc = keyIc
}

// WithKeyIe adds the keyIe to the users tokens list params
func (o *UsersTokensListParams) WithKeyIe(keyIe *string) *UsersTokensListParams {
	o.SetKeyIe(keyIe)
	return o
}

// SetKeyIe adds the keyIe to the users tokens list params
func (o *UsersTokensListParams) SetKeyIe(keyIe *string) {
	o.KeyIe = keyIe
}

// WithKeyIew adds the keyIew to the users tokens list params
func (o *UsersTokensListParams) WithKeyIew(keyIew *string) *UsersTokensListParams {
	o.SetKeyIew(keyIew)
	return o
}

// SetKeyIew adds the keyIew to the users tokens list params
func (o *UsersTokensListParams) SetKeyIew(keyIew *string) {
	o.KeyIew = keyIew
}

// WithKeyIsw adds the keyIsw to the users tokens list params
func (o *UsersTokensListParams) WithKeyIsw(keyIsw *string) *UsersTokensListParams {
	o.SetKeyIsw(keyIsw)
	return o
}

// SetKeyIsw adds the keyIsw to the users tokens list params
func (o *UsersTokensListParams) SetKeyIsw(keyIsw *string) {
	o.KeyIsw = keyIsw
}

// WithKeyn adds the keyn to the users tokens list params
func (o *UsersTokensListParams) WithKeyn(keyn *string) *UsersTokensListParams {
	o.SetKeyn(keyn)
	return o
}

// SetKeyn adds the keyN to the users tokens list params
func (o *UsersTokensListParams) SetKeyn(keyn *string) {
	o.Keyn = keyn
}

// WithKeyNic adds the keyNic to the users tokens list params
func (o *UsersTokensListParams) WithKeyNic(keyNic *string) *UsersTokensListParams {
	o.SetKeyNic(keyNic)
	return o
}

// SetKeyNic adds the keyNic to the users tokens list params
func (o *UsersTokensListParams) SetKeyNic(keyNic *string) {
	o.KeyNic = keyNic
}

// WithKeyNie adds the keyNie to the users tokens list params
func (o *UsersTokensListParams) WithKeyNie(keyNie *string) *UsersTokensListParams {
	o.SetKeyNie(keyNie)
	return o
}

// SetKeyNie adds the keyNie to the users tokens list params
func (o *UsersTokensListParams) SetKeyNie(keyNie *string) {
	o.KeyNie = keyNie
}

// WithKeyNiew adds the keyNiew to the users tokens list params
func (o *UsersTokensListParams) WithKeyNiew(keyNiew *string) *UsersTokensListParams {
	o.SetKeyNiew(keyNiew)
	return o
}

// SetKeyNiew adds the keyNiew to the users tokens list params
func (o *UsersTokensListParams) SetKeyNiew(keyNiew *string) {
	o.KeyNiew = keyNiew
}

// WithKeyNisw adds the keyNisw to the users tokens list params
func (o *UsersTokensListParams) WithKeyNisw(keyNisw *string) *UsersTokensListParams {
	o.SetKeyNisw(keyNisw)
	return o
}

// SetKeyNisw adds the keyNisw to the users tokens list params
func (o *UsersTokensListParams) SetKeyNisw(keyNisw *string) {
	o.KeyNisw = keyNisw
}

// WithLimit adds the limit to the users tokens list params
func (o *UsersTokensListParams) WithLimit(limit *int64) *UsersTokensListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the users tokens list params
func (o *UsersTokensListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the users tokens list params
func (o *UsersTokensListParams) WithOffset(offset *int64) *UsersTokensListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the users tokens list params
func (o *UsersTokensListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrdering adds the ordering to the users tokens list params
func (o *UsersTokensListParams) WithOrdering(ordering *string) *UsersTokensListParams {
	o.SetOrdering(ordering)
	return o
}

// SetOrdering adds the ordering to the users tokens list params
func (o *UsersTokensListParams) SetOrdering(ordering *string) {
	o.Ordering = ordering
}

// WithQ adds the q to the users tokens list params
func (o *UsersTokensListParams) WithQ(q *string) *UsersTokensListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the users tokens list params
func (o *UsersTokensListParams) SetQ(q *string) {
	o.Q = q
}

// WithUser adds the user to the users tokens list params
func (o *UsersTokensListParams) WithUser(user *string) *UsersTokensListParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the users tokens list params
func (o *UsersTokensListParams) SetUser(user *string) {
	o.User = user
}

// WithUsern adds the usern to the users tokens list params
func (o *UsersTokensListParams) WithUsern(usern *string) *UsersTokensListParams {
	o.SetUsern(usern)
	return o
}

// SetUsern adds the userN to the users tokens list params
func (o *UsersTokensListParams) SetUsern(usern *string) {
	o.Usern = usern
}

// WithUserID adds the userID to the users tokens list params
func (o *UsersTokensListParams) WithUserID(userID *string) *UsersTokensListParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the users tokens list params
func (o *UsersTokensListParams) SetUserID(userID *string) {
	o.UserID = userID
}

// WithUserIDn adds the userIDn to the users tokens list params
func (o *UsersTokensListParams) WithUserIDn(userIDn *string) *UsersTokensListParams {
	o.SetUserIDn(userIDn)
	return o
}

// SetUserIDn adds the userIdN to the users tokens list params
func (o *UsersTokensListParams) SetUserIDn(userIDn *string) {
	o.UserIDn = userIDn
}

// WithWriteEnabled adds the writeEnabled to the users tokens list params
func (o *UsersTokensListParams) WithWriteEnabled(writeEnabled *string) *UsersTokensListParams {
	o.SetWriteEnabled(writeEnabled)
	return o
}

// SetWriteEnabled adds the writeEnabled to the users tokens list params
func (o *UsersTokensListParams) SetWriteEnabled(writeEnabled *string) {
	o.WriteEnabled = writeEnabled
}

// WriteToRequest writes these params to a swagger request
func (o *UsersTokensListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Created != nil {

		// query param created
		var qrCreated string

		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated
		if qCreated != "" {

			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}
	}

	if o.CreatedGte != nil {

		// query param created__gte
		var qrCreatedGte string

		if o.CreatedGte != nil {
			qrCreatedGte = *o.CreatedGte
		}
		qCreatedGte := qrCreatedGte
		if qCreatedGte != "" {

			if err := r.SetQueryParam("created__gte", qCreatedGte); err != nil {
				return err
			}
		}
	}

	if o.CreatedLte != nil {

		// query param created__lte
		var qrCreatedLte string

		if o.CreatedLte != nil {
			qrCreatedLte = *o.CreatedLte
		}
		qCreatedLte := qrCreatedLte
		if qCreatedLte != "" {

			if err := r.SetQueryParam("created__lte", qCreatedLte); err != nil {
				return err
			}
		}
	}

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	if o.DescriptionEmpty != nil {

		// query param description__empty
		var qrDescriptionEmpty string

		if o.DescriptionEmpty != nil {
			qrDescriptionEmpty = *o.DescriptionEmpty
		}
		qDescriptionEmpty := qrDescriptionEmpty
		if qDescriptionEmpty != "" {

			if err := r.SetQueryParam("description__empty", qDescriptionEmpty); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIc != nil {

		// query param description__ic
		var qrDescriptionIc string

		if o.DescriptionIc != nil {
			qrDescriptionIc = *o.DescriptionIc
		}
		qDescriptionIc := qrDescriptionIc
		if qDescriptionIc != "" {

			if err := r.SetQueryParam("description__ic", qDescriptionIc); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIe != nil {

		// query param description__ie
		var qrDescriptionIe string

		if o.DescriptionIe != nil {
			qrDescriptionIe = *o.DescriptionIe
		}
		qDescriptionIe := qrDescriptionIe
		if qDescriptionIe != "" {

			if err := r.SetQueryParam("description__ie", qDescriptionIe); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIew != nil {

		// query param description__iew
		var qrDescriptionIew string

		if o.DescriptionIew != nil {
			qrDescriptionIew = *o.DescriptionIew
		}
		qDescriptionIew := qrDescriptionIew
		if qDescriptionIew != "" {

			if err := r.SetQueryParam("description__iew", qDescriptionIew); err != nil {
				return err
			}
		}
	}

	if o.DescriptionIsw != nil {

		// query param description__isw
		var qrDescriptionIsw string

		if o.DescriptionIsw != nil {
			qrDescriptionIsw = *o.DescriptionIsw
		}
		qDescriptionIsw := qrDescriptionIsw
		if qDescriptionIsw != "" {

			if err := r.SetQueryParam("description__isw", qDescriptionIsw); err != nil {
				return err
			}
		}
	}

	if o.Descriptionn != nil {

		// query param description__n
		var qrDescriptionn string

		if o.Descriptionn != nil {
			qrDescriptionn = *o.Descriptionn
		}
		qDescriptionn := qrDescriptionn
		if qDescriptionn != "" {

			if err := r.SetQueryParam("description__n", qDescriptionn); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNic != nil {

		// query param description__nic
		var qrDescriptionNic string

		if o.DescriptionNic != nil {
			qrDescriptionNic = *o.DescriptionNic
		}
		qDescriptionNic := qrDescriptionNic
		if qDescriptionNic != "" {

			if err := r.SetQueryParam("description__nic", qDescriptionNic); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNie != nil {

		// query param description__nie
		var qrDescriptionNie string

		if o.DescriptionNie != nil {
			qrDescriptionNie = *o.DescriptionNie
		}
		qDescriptionNie := qrDescriptionNie
		if qDescriptionNie != "" {

			if err := r.SetQueryParam("description__nie", qDescriptionNie); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNiew != nil {

		// query param description__niew
		var qrDescriptionNiew string

		if o.DescriptionNiew != nil {
			qrDescriptionNiew = *o.DescriptionNiew
		}
		qDescriptionNiew := qrDescriptionNiew
		if qDescriptionNiew != "" {

			if err := r.SetQueryParam("description__niew", qDescriptionNiew); err != nil {
				return err
			}
		}
	}

	if o.DescriptionNisw != nil {

		// query param description__nisw
		var qrDescriptionNisw string

		if o.DescriptionNisw != nil {
			qrDescriptionNisw = *o.DescriptionNisw
		}
		qDescriptionNisw := qrDescriptionNisw
		if qDescriptionNisw != "" {

			if err := r.SetQueryParam("description__nisw", qDescriptionNisw); err != nil {
				return err
			}
		}
	}

	if o.Expires != nil {

		// query param expires
		var qrExpires string

		if o.Expires != nil {
			qrExpires = *o.Expires
		}
		qExpires := qrExpires
		if qExpires != "" {

			if err := r.SetQueryParam("expires", qExpires); err != nil {
				return err
			}
		}
	}

	if o.ExpiresGte != nil {

		// query param expires__gte
		var qrExpiresGte string

		if o.ExpiresGte != nil {
			qrExpiresGte = *o.ExpiresGte
		}
		qExpiresGte := qrExpiresGte
		if qExpiresGte != "" {

			if err := r.SetQueryParam("expires__gte", qExpiresGte); err != nil {
				return err
			}
		}
	}

	if o.ExpiresLte != nil {

		// query param expires__lte
		var qrExpiresLte string

		if o.ExpiresLte != nil {
			qrExpiresLte = *o.ExpiresLte
		}
		qExpiresLte := qrExpiresLte
		if qExpiresLte != "" {

			if err := r.SetQueryParam("expires__lte", qExpiresLte); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.IDGt != nil {

		// query param id__gt
		var qrIDGt string

		if o.IDGt != nil {
			qrIDGt = *o.IDGt
		}
		qIDGt := qrIDGt
		if qIDGt != "" {

			if err := r.SetQueryParam("id__gt", qIDGt); err != nil {
				return err
			}
		}
	}

	if o.IDGte != nil {

		// query param id__gte
		var qrIDGte string

		if o.IDGte != nil {
			qrIDGte = *o.IDGte
		}
		qIDGte := qrIDGte
		if qIDGte != "" {

			if err := r.SetQueryParam("id__gte", qIDGte); err != nil {
				return err
			}
		}
	}

	if o.IDLt != nil {

		// query param id__lt
		var qrIDLt string

		if o.IDLt != nil {
			qrIDLt = *o.IDLt
		}
		qIDLt := qrIDLt
		if qIDLt != "" {

			if err := r.SetQueryParam("id__lt", qIDLt); err != nil {
				return err
			}
		}
	}

	if o.IDLte != nil {

		// query param id__lte
		var qrIDLte string

		if o.IDLte != nil {
			qrIDLte = *o.IDLte
		}
		qIDLte := qrIDLte
		if qIDLte != "" {

			if err := r.SetQueryParam("id__lte", qIDLte); err != nil {
				return err
			}
		}
	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string

		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {

			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}
	}

	if o.Key != nil {

		// query param key
		var qrKey string

		if o.Key != nil {
			qrKey = *o.Key
		}
		qKey := qrKey
		if qKey != "" {

			if err := r.SetQueryParam("key", qKey); err != nil {
				return err
			}
		}
	}

	if o.KeyEmpty != nil {

		// query param key__empty
		var qrKeyEmpty string

		if o.KeyEmpty != nil {
			qrKeyEmpty = *o.KeyEmpty
		}
		qKeyEmpty := qrKeyEmpty
		if qKeyEmpty != "" {

			if err := r.SetQueryParam("key__empty", qKeyEmpty); err != nil {
				return err
			}
		}
	}

	if o.KeyIc != nil {

		// query param key__ic
		var qrKeyIc string

		if o.KeyIc != nil {
			qrKeyIc = *o.KeyIc
		}
		qKeyIc := qrKeyIc
		if qKeyIc != "" {

			if err := r.SetQueryParam("key__ic", qKeyIc); err != nil {
				return err
			}
		}
	}

	if o.KeyIe != nil {

		// query param key__ie
		var qrKeyIe string

		if o.KeyIe != nil {
			qrKeyIe = *o.KeyIe
		}
		qKeyIe := qrKeyIe
		if qKeyIe != "" {

			if err := r.SetQueryParam("key__ie", qKeyIe); err != nil {
				return err
			}
		}
	}

	if o.KeyIew != nil {

		// query param key__iew
		var qrKeyIew string

		if o.KeyIew != nil {
			qrKeyIew = *o.KeyIew
		}
		qKeyIew := qrKeyIew
		if qKeyIew != "" {

			if err := r.SetQueryParam("key__iew", qKeyIew); err != nil {
				return err
			}
		}
	}

	if o.KeyIsw != nil {

		// query param key__isw
		var qrKeyIsw string

		if o.KeyIsw != nil {
			qrKeyIsw = *o.KeyIsw
		}
		qKeyIsw := qrKeyIsw
		if qKeyIsw != "" {

			if err := r.SetQueryParam("key__isw", qKeyIsw); err != nil {
				return err
			}
		}
	}

	if o.Keyn != nil {

		// query param key__n
		var qrKeyn string

		if o.Keyn != nil {
			qrKeyn = *o.Keyn
		}
		qKeyn := qrKeyn
		if qKeyn != "" {

			if err := r.SetQueryParam("key__n", qKeyn); err != nil {
				return err
			}
		}
	}

	if o.KeyNic != nil {

		// query param key__nic
		var qrKeyNic string

		if o.KeyNic != nil {
			qrKeyNic = *o.KeyNic
		}
		qKeyNic := qrKeyNic
		if qKeyNic != "" {

			if err := r.SetQueryParam("key__nic", qKeyNic); err != nil {
				return err
			}
		}
	}

	if o.KeyNie != nil {

		// query param key__nie
		var qrKeyNie string

		if o.KeyNie != nil {
			qrKeyNie = *o.KeyNie
		}
		qKeyNie := qrKeyNie
		if qKeyNie != "" {

			if err := r.SetQueryParam("key__nie", qKeyNie); err != nil {
				return err
			}
		}
	}

	if o.KeyNiew != nil {

		// query param key__niew
		var qrKeyNiew string

		if o.KeyNiew != nil {
			qrKeyNiew = *o.KeyNiew
		}
		qKeyNiew := qrKeyNiew
		if qKeyNiew != "" {

			if err := r.SetQueryParam("key__niew", qKeyNiew); err != nil {
				return err
			}
		}
	}

	if o.KeyNisw != nil {

		// query param key__nisw
		var qrKeyNisw string

		if o.KeyNisw != nil {
			qrKeyNisw = *o.KeyNisw
		}
		qKeyNisw := qrKeyNisw
		if qKeyNisw != "" {

			if err := r.SetQueryParam("key__nisw", qKeyNisw); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Ordering != nil {

		// query param ordering
		var qrOrdering string

		if o.Ordering != nil {
			qrOrdering = *o.Ordering
		}
		qOrdering := qrOrdering
		if qOrdering != "" {

			if err := r.SetQueryParam("ordering", qOrdering); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.User != nil {

		// query param user
		var qrUser string

		if o.User != nil {
			qrUser = *o.User
		}
		qUser := qrUser
		if qUser != "" {

			if err := r.SetQueryParam("user", qUser); err != nil {
				return err
			}
		}
	}

	if o.Usern != nil {

		// query param user__n
		var qrUsern string

		if o.Usern != nil {
			qrUsern = *o.Usern
		}
		qUsern := qrUsern
		if qUsern != "" {

			if err := r.SetQueryParam("user__n", qUsern); err != nil {
				return err
			}
		}
	}

	if o.UserID != nil {

		// query param user_id
		var qrUserID string

		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := qrUserID
		if qUserID != "" {

			if err := r.SetQueryParam("user_id", qUserID); err != nil {
				return err
			}
		}
	}

	if o.UserIDn != nil {

		// query param user_id__n
		var qrUserIDn string

		if o.UserIDn != nil {
			qrUserIDn = *o.UserIDn
		}
		qUserIDn := qrUserIDn
		if qUserIDn != "" {

			if err := r.SetQueryParam("user_id__n", qUserIDn); err != nil {
				return err
			}
		}
	}

	if o.WriteEnabled != nil {

		// query param write_enabled
		var qrWriteEnabled string

		if o.WriteEnabled != nil {
			qrWriteEnabled = *o.WriteEnabled
		}
		qWriteEnabled := qrWriteEnabled
		if qWriteEnabled != "" {

			if err := r.SetQueryParam("write_enabled", qWriteEnabled); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
