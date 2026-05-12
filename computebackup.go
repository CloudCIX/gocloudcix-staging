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

// Management of Instance Backups
//
// Supported backup types:
//
// - "lxd" - LXD backups for Linux containers and VMs
// - "hyperv" - Hyper-V backups for Windows VMs
//
// This module provides API endpoints for managing backups of virtual machine and
// container instances in the CloudCIX Compute platform. Backups are on-demand
// copies of running instances stored in backup repositories for disaster recovery
// and data protection.
//
// Available operations:
//
// - List and filter backups across your projects by type, instance, or repository
// - Create new backups from running LXD or Hyper-V instances
// - Retrieve detailed information about individual backups including validity time
// - Delete backups by updating their state to delete
//
// Each backup includes its associated instance, project, repository location, and
// time valid timestamp.
//
// ComputeBackupService contains methods and other services that help with
// interacting with the gocloudcix API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeBackupService] method instead.
type ComputeBackupService struct {
	Options []option.RequestOption
}

// NewComputeBackupService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewComputeBackupService(opts ...option.RequestOption) (r ComputeBackupService) {
	r = ComputeBackupService{}
	r.Options = opts
	return
}

// Create a new backup from a running compute instance (LXD or Hyper-V). Specify
// the backup type, instance to backup and project. The instance must be in a
// running state to create a backup.
func (r *ComputeBackupService) New(ctx context.Context, body ComputeBackupNewParams, opts ...option.RequestOption) (res *ComputeBackup, err error) {
	var env ComputeBackupResponse
	opts = slices.Concat(r.Options, opts)
	path := "compute/backups/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Retrieve detailed information about a specific backup, including its type (LXD
// or Hyper-V), associated instance, project, validity timestamp, and current
// state.
func (r *ComputeBackupService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *ComputeBackup, err error) {
	var env ComputeBackupResponse
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("compute/backups/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Update a backup to change its name or state. Set state to delete to initiate
// backup deletion and free up repository storage space.
func (r *ComputeBackupService) Update(ctx context.Context, id int64, body ComputeBackupUpdateParams, opts ...option.RequestOption) (res *ComputeBackup, err error) {
	var env ComputeBackupResponse
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("compute/backups/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Retrieve a paginated list of instance backups (LXD and/or Hyper-V) across your
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
func (r *ComputeBackupService) List(ctx context.Context, query ComputeBackupListParams, opts ...option.RequestOption) (res *ComputeBackupListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/backups/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// This method allows the User to delete a specified Compute Backup.
func (r *ComputeBackupService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("compute/backups/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type ComputeBackup struct {
	// The ID of the Compute Backups record
	ID int64 `json:"id" api:"required"`
	// Timestamp, in ISO format, of when the Compute Backups record was created.
	Created string `json:"created" api:"required"`
	// The Compute Instance the Compute Backup record is of.
	Instance ComputeBackupInstance `json:"instance" api:"required"`
	// The user-friendly name given to this Compute Backups instance
	Name string `json:"name" api:"required"`
	// The id of the Project that this Compute Backups belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// An array of the specs for the Compute Backups
	Specs []Bom `json:"specs" api:"required"`
	// The current state of the Compute Backups
	State string `json:"state" api:"required"`
	// The type of the Compute Backups
	Type string `json:"type" api:"required"`
	// Timestamp, in ISO format, of when the Compute Backups record was last updated.
	Updated string `json:"updated" api:"required"`
	// URL that can be used to run methods in the API associated with the Compute
	// Backups instance.
	Uri string `json:"uri" api:"required" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		Instance    respjson.Field
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
func (r ComputeBackup) RawJSON() string { return r.JSON.raw }
func (r *ComputeBackup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The Compute Instance the Compute Backup record is of.
type ComputeBackupInstance struct {
	// The ID of the Compute Instance the Compute Backup is of.
	ID int64 `json:"id"`
	// The user-friendly name of the Compute Instance the Compute Backup is of.
	Name string `json:"name"`
	// The current state of the Compute Instance the Compute Backup is of.
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
func (r ComputeBackupInstance) RawJSON() string { return r.JSON.raw }
func (r *ComputeBackupInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeBackupResponse struct {
	Content ComputeBackup `json:"content"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeBackupResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeBackupResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeBackupUpdateParam struct {
	// The user-friendly name for the Compute Backups Resource. If not sent, it will
	// default to the current name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Change the state of the Compute Backup, triggering the CloudCIX Robot to perform
	// the requested action. Users can only request state changes from certain current
	// states:
	//
	// - running -> delete
	State param.Opt[string] `json:"state,omitzero"`
	paramObj
}

func (r ComputeBackupUpdateParam) MarshalJSON() (data []byte, err error) {
	type shadow ComputeBackupUpdateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeBackupUpdateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListMetadata struct {
	// The value of limit that was used for the request
	Limit int64 `json:"limit" api:"required"`
	// The value of order that was used for the request
	Order string `json:"order" api:"required"`
	// The value of page that was used for the request
	Page int64 `json:"page" api:"required"`
	// The total number of records found for the given search
	TotalRecords int64 `json:"total_records" api:"required"`
	// A list of warnings generated during execution. Any invalid search filters used
	// will cause a warning to be generated, for example.
	Warnings []string `json:"warnings" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limit        respjson.Field
		Order        respjson.Field
		Page         respjson.Field
		TotalRecords respjson.Field
		Warnings     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListMetadata) RawJSON() string { return r.JSON.raw }
func (r *ListMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeBackupListResponse struct {
	Metadata ListMetadata    `json:"_metadata"`
	Content  []ComputeBackup `json:"content"`
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
func (r ComputeBackupListResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeBackupListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeBackupNewParams struct {
	// The id of the Compute Instance the Compute Backup is to be taken of.
	InstanceID int64 `json:"instance_id" api:"required"`
	// The ID of the User's Project into which this new Compute Backups should be
	// added.
	ProjectID int64 `json:"project_id" api:"required"`
	// The user-friendly name for the Compute Backup. If not sent, it will default to
	// the name "Backup HyperV" or "Backup LXD" depending on the type chosen.
	Name param.Opt[string] `json:"name,omitzero"`
	// The type of Compute Backup to create. Valid options are:
	//
	// - "hyperv"
	// - "lxd"
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r ComputeBackupNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputeBackupNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeBackupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeBackupUpdateParams struct {
	ComputeBackupUpdate ComputeBackupUpdateParam
	paramObj
}

func (r ComputeBackupUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ComputeBackupUpdate)
}
func (r *ComputeBackupUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeBackupListParams struct {
	// The limit of the number of objects returned per page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// The field to use for ordering. Possible fields and the default are outlined in
	// the individual method descriptions.
	Order param.Opt[string] `query:"order,omitzero" json:"-"`
	// The page of records to return, assuming `limit` number of records per page.
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// Filter the result to objects that do not match the specified filters. Possible
	// filters are outlined in the individual list method descriptions.
	Exclude map[string]any `query:"exclude,omitzero" json:"-"`
	// Filter the result to objects that match the specified filters. Possible filters
	// are outlined in the individual list method descriptions.
	Search map[string]any `query:"search,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ComputeBackupListParams]'s query parameters as
// `url.Values`.
func (r ComputeBackupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
