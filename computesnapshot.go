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
	shimjson "github.com/CloudCIX/gocloudcix/internal/encoding/json"
	"github.com/CloudCIX/gocloudcix/internal/requestconfig"
	"github.com/CloudCIX/gocloudcix/option"
	"github.com/CloudCIX/gocloudcix/packages/param"
	"github.com/CloudCIX/gocloudcix/packages/respjson"
)

// Management of Instance Snapshots
//
// This module provides API endpoints for managing snapshots of virtual machine and
// container instances in the CloudCIX Compute platform. Snapshots are
// point-in-time backups of running instances that can be used for data recovery
// from a known state. Two snapshot types are supported: LXD snapshots (for Linux
// containers and VMs) and Hyper-V snapshots (for Windows VMs).
//
// Available operations:
//
// - List and filter snapshots across your projects by type, instance, or state
// - Create new snapshots from running LXD or Hyper-V instances
// - Retrieve detailed information about individual snapshots
// - Delete snapshots by updating their state
//
// Each snapshot includes its associated instance, project, creation timestamp, and
// current state.
//
// ComputeSnapshotService contains methods and other services that help with
// interacting with the gocloudcix API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeSnapshotService] method instead.
type ComputeSnapshotService struct {
	Options []option.RequestOption
}

// NewComputeSnapshotService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewComputeSnapshotService(opts ...option.RequestOption) (r ComputeSnapshotService) {
	r = ComputeSnapshotService{}
	r.Options = opts
	return
}

// Create a new snapshot from a running compute instance (LXD or Hyper-V). Specify
// the snapshot type, instance to snapshot, and project. The instance must be in a
// running state to create a snapshot.
func (r *ComputeSnapshotService) New(ctx context.Context, body ComputeSnapshotNewParams, opts ...option.RequestOption) (res *ComputeSnapshot, err error) {
	var env ComputeSnapshotResponse
	opts = slices.Concat(r.Options, opts)
	path := "compute/snapshots/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Retrieve detailed information about a specific snapshot, including its type (LXD
// or Hyper-V), associated instance, project, creation timestamp, and current
// state.
func (r *ComputeSnapshotService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *ComputeSnapshot, err error) {
	var env ComputeSnapshotResponse
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("compute/snapshots/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Update a snapshot to change its name or state. Set state to delete to initiate
// snapshot deletion and free up storage space.
func (r *ComputeSnapshotService) Update(ctx context.Context, id int64, body ComputeSnapshotUpdateParams, opts ...option.RequestOption) (res *ComputeSnapshot, err error) {
	var env ComputeSnapshotResponse
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("compute/snapshots/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Retrieve a paginated list of instance snapshots (LXD and/or Hyper-V) across your
// projects. Supports filtering by type, project, instance, state, and other
// attributes. Results can be ordered and paginated.
//
// ## Filtering
//
// The following fields and modifiers can be used to filter records from the list;
//
// - resource_links\_\_contra_resource_id (gt, gte, in, isnull, lt, lte, range)
// - created (gt, gte, in, isnull, lt, lte, range)
// - id (gt, gte, in, isnull, lt, lte, range)
// - name (in, icontains, iendswith, iexact, istartswith)
// - project_id (gt, gte, in, isnull, lt, lte, range)
// - project\_\_address_id (gt, gte, in, isnull, lt, lte, range)
// - project\_\_name (in, icontains, iendswith, iexact, istartswith)
// - project\_\_region_id (gt, gte, in, isnull, lt, lte, range)
// - project\_\_reseller_id (gt, gte, in, isnull, lt, lte, range)
// - state (gt, gte, in, isnull, lt, lte, range)
// - type
// - updated (gt, gte, in, isnull, lt, lte, range)
//
// To search, simply add `?search[field]=value` to include records that match the
// request, or `?exclude[field]=value` to exclude them. To use modifiers, simply
// add `?search[field__modifier]` and `?exclude[field__modifier]`
//
// ## Ordering
//
// The following fields can be used to order the results of the list;
//
// - name (default)
// - id
// - project_id
// - state
//
// To reverse the ordering, simply prepend a `-` character to the request. So
// `?order=field` orders by `field` in ascending order, while `?order=-field`
// orders in descending order instead.
func (r *ComputeSnapshotService) List(ctx context.Context, query ComputeSnapshotListParams, opts ...option.RequestOption) (res *ComputeSnapshotListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/snapshots/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// This method allows the User to delete a specified Compute Snapshot.
func (r *ComputeSnapshotService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("compute/snapshots/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type ComputeSnapshot struct {
	// The ID of the Compute Snapshots record
	ID int64 `json:"id" api:"required"`
	// Timestamp, in ISO format, of when the Compute Snapshots record was created.
	Created string `json:"created" api:"required"`
	// The Compute Instance the Compute Snapshot record is of.
	Instance ComputeSnapshotInstance `json:"instance" api:"required"`
	// The metadata details of the The metadata details of the "hyperv" Compute
	// Snapshot. Returned if the type is "hyperv".
	Metadata ComputeSnapshotMetadata `json:"metadata" api:"required"`
	// The user-friendly name given to this Compute Snapshots instance
	Name string `json:"name" api:"required"`
	// The id of the Project that this Compute Snapshots belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// An array of the specs for the Compute Snapshots
	Specs []Bom `json:"specs" api:"required"`
	// The current state of the Compute Snapshots
	State string `json:"state" api:"required"`
	// The type of the Compute Snapshots
	Type string `json:"type" api:"required"`
	// Timestamp, in ISO format, of when the Compute Snapshots record was last updated.
	Updated string `json:"updated" api:"required"`
	// URL that can be used to run methods in the API associated with the Compute
	// Snapshots instance.
	Uri string `json:"uri" api:"required" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Instance    respjson.Field
		Metadata    respjson.Field
		Name        respjson.Field
		ProjectID   respjson.Field
		Specs       respjson.Field
		State       respjson.Field
		Type        respjson.Field
		Updated     respjson.Field
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSnapshot) RawJSON() string { return r.JSON.raw }
func (r *ComputeSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Compute Instance the Compute Snapshot record is of.
type ComputeSnapshotInstance struct {
	// The ID of the Compute Instance the Compute Snapshot is of.
	ID int64 `json:"id"`
	// The user-friendly name of the Compute Instance the Compute Snapshot is of.
	Name string `json:"name"`
	// The current state of the Compute Instance the Compute Snapshot is of.
	State string `json:"state"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		State       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSnapshotInstance) RawJSON() string { return r.JSON.raw }
func (r *ComputeSnapshotInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The metadata details of the The metadata details of the "hyperv" Compute
// Snapshot. Returned if the type is "hyperv".
type ComputeSnapshotMetadata struct {
	// Indicates if the "hyperv" Compute Snapshot is currently active.
	Active bool `json:"active"`
	// Indicates if the "hyperv" Compute Snapshot should remove the subtree when
	// deleted.
	RemoveSubtree bool `json:"remove_subtree"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active        respjson.Field
		RemoveSubtree respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSnapshotMetadata) RawJSON() string { return r.JSON.raw }
func (r *ComputeSnapshotMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSnapshotResponse struct {
	Content ComputeSnapshot `json:"content"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeSnapshotResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeSnapshotResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property State is required.
type ComputeSnapshotUpdateParam struct {
	// Change the state of the Compute Snapshot, triggering the CloudCIX Robot to
	// perform the requested action. Users can only request state changes from certain
	// current states:
	//
	// - running -> update_running or delete
	State string `json:"state" api:"required"`
	// The user-friendly name for the Compute Snapshots Resource. If not sent, it will
	// default to the name current name.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ComputeSnapshotUpdateParam) MarshalJSON() (data []byte, err error) {
	type shadow ComputeSnapshotUpdateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeSnapshotUpdateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSnapshotListResponse struct {
	Metadata ListMetadata      `json:"_metadata"`
	Content  []ComputeSnapshot `json:"content"`
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
func (r ComputeSnapshotListResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeSnapshotListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSnapshotNewParams struct {
	// The id of the Compute Instance the Compute Snapshot is to be taken of.
	InstanceID int64 `json:"instance_id" api:"required"`
	// The ID of the User's Project into which this new Compute Snapshots should be
	// added.
	ProjectID int64 `json:"project_id" api:"required"`
	// The user-friendly name for the Compute Snapshot Resource. If not sent, it will
	// default to the name "Compute Snapshot"
	Name param.Opt[string] `json:"name,omitzero"`
	// The type of Compute Snapshot to create. Valid options are:
	//
	// - "hyperv"
	// - "lxd"
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r ComputeSnapshotNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputeSnapshotNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeSnapshotNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSnapshotUpdateParams struct {
	ComputeSnapshotUpdate ComputeSnapshotUpdateParam
	paramObj
}

func (r ComputeSnapshotUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ComputeSnapshotUpdate)
}
func (r *ComputeSnapshotUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeSnapshotListParams struct {
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

// URLQuery serializes [ComputeSnapshotListParams]'s query parameters as
// `url.Values`.
func (r ComputeSnapshotListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
