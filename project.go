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

// Management of Cloud Projects
//
// This module provides API endpoints for managing cloud projects in the CloudCIX
// Compute platform. Projects are logical containers that organise and group your
// cloud infrastructure resources such as virtual machines, routers, firewalls, and
// storage. Each project belongs to a specific region and has its own isolated
// network environment.
//
// Available operations:
//
// - List and filter projects across your organization
// - Create new projects in available cloud regions
// - Retrieve detailed project information including region and manager
// - Update project details such as name and notes
//
// Each project includes its associated address, region, manager, and creation
// metadata.
//
// ProjectService contains methods and other services that help with interacting
// with the gocloudcix API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectService] method instead.
type ProjectService struct {
	Options []option.RequestOption
}

// NewProjectService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProjectService(opts ...option.RequestOption) (r ProjectService) {
	r = ProjectService{}
	r.Options = opts
	return
}

// Create a new cloud project in a specified region. Projects provide isolated
// network environments for deploying and managing your cloud resources. Specify
// the region and optionally provide a name and notes.
func (r *ProjectService) New(ctx context.Context, body ProjectNewParams, opts ...option.RequestOption) (res *Project, err error) {
	var env ProjectResponse
	opts = slices.Concat(r.Options, opts)
	path := "project/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Retrieve detailed information about a specific project, including its region,
// address, manager, creation date, and associated notes.
func (r *ProjectService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *Project, err error) {
	var env ProjectResponse
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("project/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Update a project to modify its name or notes. Projects cannot be moved between
// regions after creation.
func (r *ProjectService) Update(ctx context.Context, id int64, body ProjectUpdateParams, opts ...option.RequestOption) (res *Project, err error) {
	var env ProjectResponse
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("project/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Retrieve a paginated list of your cloud projects. Supports filtering by region,
// name, manager, and other attributes. Results can be ordered and paginated.
//
// ## Filtering
//
// The following fields and modifiers can be used to filter records from the list;
//
// - address_id (gt, gte, in, isnull, lt, lte, range)
// - created (gt, gte, in, isnull, lt, lte, range)
// - id (gt, gte, in, isnull, lt, lte, range)
// - manager_id (gt, gte, in, isnull, lt, lte, range)
// - name (in, icontains, iendswith, iexact, istartswith)
// - region_id (gt, gte, in, isnull, lt, lte, range)
// - reseller_id (gt, gte, in, isnull, lt, lte, range)
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
// - region_id
// - address_id
// - created
// - manager_id
//
// To reverse the ordering, simply prepend a `-` character to the request. So
// `?order=field` orders by `field` in ascending order, while `?order=-field`
// orders in descending order instead.
func (r *ProjectService) List(ctx context.Context, query ProjectListParams, opts ...option.RequestOption) (res *ProjectListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "project/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a project. Any resources in the project that are not already in a scrub
// state will be queued for deletion. If all resources are already closed, the
// project will be soft-deleted immediately.
func (r *ProjectService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("project/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type Project struct {
	// The ID of the Project.
	ID int64 `json:"id" api:"required"`
	// The ID of the Project address.
	AddressID int64 `json:"address_id" api:"required"`
	// A flag stating whether or not the Project is classified as closed. A Project is
	// classified as closed when all the infrastructure in it is in a Closed (99)
	// state.
	Closed bool `json:"closed" api:"required"`
	// The date that the Project entry was created
	Created string `json:"created" api:"required"`
	// The ID of the User that manages the Project
	ManagerID int64 `json:"manager_id" api:"required"`
	// The name of the Project.
	Name string `json:"name" api:"required"`
	// The note attached to the Project.
	Note string `json:"note" api:"required"`
	// The region ID that the Project is in.
	RegionID int64 `json:"region_id" api:"required"`
	// The Address ID that will send the bill for the Project to the customer.
	ResellerID int64 `json:"reseller_id" api:"required"`
	// The date that the Project entry was last updated
	Updated string `json:"updated" api:"required"`
	// The absolute URL of the Project that can be used to perform `Read`, `Update` and
	// `Delete`
	Uri string `json:"uri" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AddressID   respjson.Field
		Closed      respjson.Field
		Created     respjson.Field
		ManagerID   respjson.Field
		Name        respjson.Field
		Note        respjson.Field
		RegionID    respjson.Field
		ResellerID  respjson.Field
		Updated     respjson.Field
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Project) RawJSON() string { return r.JSON.raw }
func (r *Project) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectResponse struct {
	Content Project `json:"content"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProjectResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectUpdateParam struct {
	// The name of the Project. Must be unique within an Address' Project collection.
	Name param.Opt[string] `json:"name,omitzero"`
	// A note about the project to store information. No length limit.
	Note param.Opt[string] `json:"note,omitzero"`
	paramObj
}

func (r ProjectUpdateParam) MarshalJSON() (data []byte, err error) {
	type shadow ProjectUpdateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectUpdateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectListResponse struct {
	Metadata ListMetadata `json:"_metadata"`
	Content  []Project    `json:"content"`
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
func (r ProjectListResponse) RawJSON() string { return r.JSON.raw }
func (r *ProjectListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectNewParams struct {
	// The name of the Project. Must be unique within an Address' Project collection.
	Name string `json:"name" api:"required"`
	// The Address ID of the CloudCIX region that the Project will be deployed in.
	RegionID int64 `json:"region_id" api:"required"`
	// An optional note providing a description of what the Project is used for.
	Note param.Opt[string] `json:"note,omitzero"`
	paramObj
}

func (r ProjectNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ProjectNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProjectNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectUpdateParams struct {
	ProjectUpdate ProjectUpdateParam
	paramObj
}

func (r ProjectUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ProjectUpdate)
}
func (r *ProjectUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProjectListParams struct {
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

// URLQuery serializes [ProjectListParams]'s query parameters as `url.Values`.
func (r ProjectListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
