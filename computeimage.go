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

// Management of Operating System Images
//
// This module provides API endpoints for browsing available operating system
// images that can be used when creating virtual instances in the CloudCIX Compute
// platform. Images represent pre-configured OS templates including various Linux
// distributions and Windows versions.
//
// Available operations:
//
//   - List and filter available OS images by region, name, or variant
//   - Retrieve detailed information about a specific image including its SKU and OS
//     variant
//
// ComputeImageService contains methods and other services that help with
// interacting with the gocloudcix API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeImageService] method instead.
type ComputeImageService struct {
	Options []option.RequestOption
}

// NewComputeImageService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewComputeImageService(opts ...option.RequestOption) (r ComputeImageService) {
	r = ComputeImageService{}
	r.Options = opts
	return
}

// Retrieve detailed information about a specific operating system image, including
// its SKU name, filename, and OS variant.
func (r *ComputeImageService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *ComputeImage, err error) {
	var env ComputeImageGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("compute/images/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Retrieve a paginated list of available operating system images. Supports
// filtering by region, SKU name, and OS variant. Results can be ordered and
// paginated.
//
// ## Filtering
//
// The following fields and modifiers can be used to filter records from the list;
//
// - id (gt, gte, in, isnull, lt, lte, range)
// - member_id (gt, gte, in, isnull, lt, lte, range)
// - os_variant (in, icontains, iendswith, iexact, istartswith)
// - region_id (gt, gte, in, isnull, lt, lte, range)
// - sku_name (in, icontains, iendswith, iexact, istartswith)
// - type
//
// To search, simply add `?search[field]=value` to include records that match the
// request, or `?exclude[field]=value` to exclude them. To use modifiers, simply
// add `?search[field__modifier]` and `?exclude[field__modifier]`
//
// ## Ordering
//
// The following fields can be used to order the results of the list;
//
// - sku_name (default)
// - id
//
// To reverse the ordering, simply prepend a `-` character to the request. So
// `?order=field` orders by `field` in ascending order, while `?order=-field`
// orders in descending order instead.
func (r *ComputeImageService) List(ctx context.Context, query ComputeImageListParams, opts ...option.RequestOption) (res *ComputeImageListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/images/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type ComputeImage struct {
	// The ID of the Image.
	ID int64 `json:"id" api:"required"`
	// The name of the file containing the Image.
	Filename string `json:"filename" api:"required"`
	// Is a unique word to define each Image.
	OsVariant string `json:"os_variant" api:"required"`
	// The protocl for images supported by LXD Hypervisor device types.
	Protocol string `json:"protocol" api:"required"`
	// The SKU for the Image.
	SKUName string `json:"sku_name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Filename    respjson.Field
		OsVariant   respjson.Field
		Protocol    respjson.Field
		SKUName     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeImage) RawJSON() string { return r.JSON.raw }
func (r *ComputeImage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeImageListResponse struct {
	Metadata ListMetadata   `json:"_metadata"`
	Content  []ComputeImage `json:"content"`
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
func (r ComputeImageListResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeImageListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeImageGetResponseEnvelope struct {
	Content ComputeImage `json:"content"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeImageGetResponseEnvelope) RawJSON() string { return r.JSON.raw }
func (r *ComputeImageGetResponseEnvelope) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeImageListParams struct {
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

// URLQuery serializes [ComputeImageListParams]'s query parameters as `url.Values`.
func (r ComputeImageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
