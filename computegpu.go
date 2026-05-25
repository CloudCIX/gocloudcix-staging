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

// Management of GPU Resources
//
// This module provides API endpoints for managing GPU (Graphics Processing Unit)
// resources within the CloudCIX Compute platform. GPUs are physical hardware
// accelerators that can be attached to LXD instances to provide enhanced
// computational capabilities for workloads such as machine learning, AI training,
// scientific simulations, and graphics rendering.
//
// Available operations:
//
// - List and filter GPU resources across your projects
// - Attach GPUs to running LXD instances by creating new GPU resources
// - Retrieve individual GPU configuration and status details
// - Detach GPUs from instances by updating the state to delete
//
// Each GPU resource includes its associated LXD instance, capacity specifications
// (SKUs), current state, and project information.
//
// ComputeGPUService contains methods and other services that help with interacting
// with the gocloudcix API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeGPUService] method instead.
type ComputeGPUService struct {
	Options []option.RequestOption
}

// NewComputeGPUService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewComputeGPUService(opts ...option.RequestOption) (r ComputeGPUService) {
	r = ComputeGPUService{}
	r.Options = opts
	return
}

// Retrieve detailed information about a specific GPU resource, including its
// attached LXD instance, capacity specifications, current state, and project
// information.
func (r *ComputeGPUService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *ComputeGPU, err error) {
	var env ComputeGPUResponse
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("compute/gpus/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Update a GPU resource to change its state. Set state to delete to initiate
// detachment from the LXD instance.
func (r *ComputeGPUService) Update(ctx context.Context, id int64, body ComputeGPUUpdateParams, opts ...option.RequestOption) (res *ComputeGPU, err error) {
	var env ComputeGPUResponse
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("compute/gpus/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Retrieve a paginated list of GPU resources across your projects. Supports
// filtering by project, region, state, and other attributes. Results can be
// ordered and paginated.
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
//
// To reverse the ordering, simply prepend a `-` character to the request. So
// `?order=field` orders by `field` in ascending order, while `?order=-field`
// orders in descending order instead.
func (r *ComputeGPUService) List(ctx context.Context, query ComputeGPUListParams, opts ...option.RequestOption) (res *ComputeGPUListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/gpus/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// This method allows the User to delete a specified Compute GPU. Requesting to
// delete a GPU will result in the GPU being detached from the LXD Compute Instance
// it is attached to.
func (r *ComputeGPUService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("compute/gpus/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Attach a GPU hardware accelerator to a running LXD instance. Specify the target
// LXD instance ID, project, and GPU specifications (SKUs) to provision and attach
// the GPU resource.
func (r *ComputeGPUService) Attach(ctx context.Context, body ComputeGPUAttachParams, opts ...option.RequestOption) (res *ComputeGPU, err error) {
	var env ComputeGPUResponse
	opts = slices.Concat(r.Options, opts)
	path := "compute/gpus/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

type ComputeGPU struct {
	// The ID of the Compute GPU record
	ID int64 `json:"id" api:"required"`
	// Timestamp, in ISO format, of when the Compute GPU record was created.
	Created string `json:"created" api:"required"`
	// The "lxd" Compute Instance the Compute GPU is attached to.
	Instance ComputeGPUInstance `json:"instance" api:"required"`
	// The user-friendly name given to this Compute GPU
	Name string `json:"name" api:"required"`
	// The id of the Project that this Compute GPU belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// An array of the specs for the Compute GPU
	Specs []Bom `json:"specs" api:"required"`
	// The current state of the Compute GPU
	State string `json:"state" api:"required"`
	// Timestamp, in ISO format, of when the Compute GPU record was last updated.
	Updated string `json:"updated" api:"required"`
	// URL that can be used to run methods in the API associated with the Compute GPU
	// instance.
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
		Updated     respjson.Field
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeGPU) RawJSON() string { return r.JSON.raw }
func (r *ComputeGPU) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The "lxd" Compute Instance the Compute GPU is attached to.
type ComputeGPUInstance struct {
	// The ID of the "lxd" Compute Instance the Compute GPU is attached to.
	ID int64 `json:"id"`
	// The user-friendly name of the "lxd" Compute Instance the Compute GPU is attached
	// to.
	Name string `json:"name"`
	// The current state of the "lxd" Compute Instance the Compute GPU is attached to.
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
func (r ComputeGPUInstance) RawJSON() string { return r.JSON.raw }
func (r *ComputeGPUInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeGPUResponse struct {
	Content ComputeGPU `json:"content"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeGPUResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeGPUResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeGPUUpdateParam struct {
	// The user-friendly name for the GPU Resource.
	Name param.Opt[string] `json:"name,omitzero"`
	// Change the state of the Compute GPU, triggering the CloudCIX Robot to perform
	// the requested action.
	//
	// To detach a Compute GPU from an lxd Compute Instance, send request to change the
	// state to delete
	State param.Opt[string] `json:"state,omitzero"`
	paramObj
}

func (r ComputeGPUUpdateParam) MarshalJSON() (data []byte, err error) {
	type shadow ComputeGPUUpdateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeGPUUpdateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeGPUListResponse struct {
	Metadata ListMetadata `json:"_metadata"`
	Content  []ComputeGPU `json:"content"`
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
func (r ComputeGPUListResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeGPUListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeGPUUpdateParams struct {
	ComputeGPUUpdate ComputeGPUUpdateParam
	paramObj
}

func (r ComputeGPUUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ComputeGPUUpdate)
}
func (r *ComputeGPUUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeGPUListParams struct {
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

// URLQuery serializes [ComputeGPUListParams]'s query parameters as `url.Values`.
func (r ComputeGPUListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ComputeGPUAttachParams struct {
	// The ID of the LXD Resource the GPU is requested to be mounted to.
	InstanceID int64 `json:"instance_id" api:"required"`
	// The ID of the Project which this GPU Resource should be in.
	ProjectID int64 `json:"project_id" api:"required"`
	// List of specs (SKUs) for the GPU resource.
	Specs []ComputeGPUAttachParamsSpec `json:"specs,omitzero" api:"required"`
	// The user-friendly name for the Compute GPU. If not sent, it will default to the
	// name 'GPU'
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ComputeGPUAttachParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputeGPUAttachParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeGPUAttachParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeGPUAttachParamsSpec struct {
	// The name of the SKU for the GPU
	SKUName param.Opt[string] `json:"sku_name,omitzero"`
	paramObj
}

func (r ComputeGPUAttachParamsSpec) MarshalJSON() (data []byte, err error) {
	type shadow ComputeGPUAttachParamsSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeGPUAttachParamsSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
