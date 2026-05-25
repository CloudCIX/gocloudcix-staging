// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gocloudcix

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/CloudCIX/gocloudcix/internal/apijson"
	"github.com/CloudCIX/gocloudcix/internal/apiquery"
	"github.com/CloudCIX/gocloudcix/internal/requestconfig"
	"github.com/CloudCIX/gocloudcix/option"
	"github.com/CloudCIX/gocloudcix/packages/param"
	"github.com/CloudCIX/gocloudcix/packages/respjson"
)

// Management of SSH Key records.
//
// This module provides API endpoints for managing SSH Key pairs used when
// provisioning compute instances. SSH Keys are stored in the Membership service
// and proxied through here.
//
// Available operations:
//
// - List SSH Keys belonging to the requesting User's Address
// - Create a new SSH Key (optionally auto-generate an Ed25519 key pair)
// - Read a single SSH Key record
// - Delete an SSH Key record
//
// ComputeSSHKeyService contains methods and other services that help with
// interacting with the gocloudcix API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeSSHKeyService] method instead.
type ComputeSSHKeyService struct {
	Options []option.RequestOption
}

// NewComputeSSHKeyService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewComputeSSHKeyService(opts ...option.RequestOption) (r ComputeSSHKeyService) {
	r = ComputeSSHKeyService{}
	r.Options = opts
	return
}

// Create a new SSH Key. If no `public_key` is provided, an Ed25519 key pair is
// generated automatically. The `private_key` is returned **once** in this response
// and is never stored — save it immediately.
func (r *ComputeSSHKeyService) New(ctx context.Context, body ComputeSSHKeyNewParams, opts ...option.RequestOption) (res *SSHKey, err error) {
	var env SSHKeyResponse
	opts = slices.Concat(r.Options, opts)
	path := "compute/ssh_keys/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Retrieve detailed information about a specific SSH Key record.
func (r *ComputeSSHKeyService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *SSHKey, err error) {
	var env SSHKeyResponse
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("compute/ssh_keys/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Returns a paginated list of all SSH Keys belonging to the requesting User's
// Address.
func (r *ComputeSSHKeyService) List(ctx context.Context, query ComputeSSHKeyListParams, opts ...option.RequestOption) (res *ComputeSSHKeyListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/ssh_keys/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Permanently delete an SSH Key record. This action cannot be undone.
func (r *ComputeSSHKeyService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("compute/ssh_keys/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type SSHKey struct {
	// The ID of the SSH Key record.
	ID int64 `json:"id" api:"required"`
	// Timestamp, in ISO format, of when the SSH Key record was created.
	Created string `json:"created" api:"required"`
	// The user-friendly name for the SSH Key.
	Name string `json:"name" api:"required"`
	// The PEM-encoded Ed25519 private key. Only present in the create (POST) response
	// when no public_key was supplied and the key pair was auto-generated. Not stored
	// by the API — save it immediately.
	PrivateKey string `json:"private_key" api:"required"`
	// The SSH public key string.
	PublicKey string `json:"public_key" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Name        respjson.Field
		PrivateKey  respjson.Field
		PublicKey   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SSHKey) RawJSON() string { return r.JSON.raw }
func (r *SSHKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SSHKeyResponse struct {
	Content SSHKey `json:"content"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SSHKeyResponse) RawJSON() string { return r.JSON.raw }
func (r *SSHKeyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSSHKeyListResponse struct {
	Metadata ListMetadata `json:"_metadata"`
	Content  []SSHKey     `json:"content"`
	// Maximum number of records returned per page.
	Count int64 `json:"count"`
	// Page number of the current results.
	Page int64 `json:"page"`
	// Total number of matching records.
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Content     respjson.Field
		Count       respjson.Field
		Page        respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSSHKeyListResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeSSHKeyListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSSHKeyNewParams struct {
	paramObj
}

func (r ComputeSSHKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputeSSHKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeSSHKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSSHKeyListParams struct {
	// The limit of the number of objects returned per page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The field to use for ordering. Possible fields and the default are outlined in
	// the individual method descriptions.
	Order param.Opt[string] `query:"order,omitzero" json:"-"`
	// The page of records to return, assuming `limit` number of records per page.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Filter the result to objects that do not match the specified filters. Possible
	// filters are outlined in the individual list method descriptions.
	Exclude any `query:"exclude,omitzero" json:"-"`
	// Filter the result to objects that match the specified filters. Possible filters
	// are outlined in the individual list method descriptions.
	Search any `query:"search,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ComputeSSHKeyListParams]'s query parameters as
// `url.Values`.
func (r ComputeSSHKeyListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
