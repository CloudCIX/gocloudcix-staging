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

// Management of Network IP Groups
//
// IP address groups organise sets of CIDR networks for use in firewall rules and
// access control. Two types are available:
//
// - Geo Groups (type="geo"): Maintained by CloudCIX and accessible to all members
//
//   - Used for geo-filtering based on country IP ranges (e.g., 'Ireland', 'USA',
//     'China')
//   - Essential for creating geo firewalls that block/allow traffic from specific
//     countries
//   - To list country groups: GET /ip_address_groups?search[member_id]=0
//   - Referenced in geo firewall rules by numeric ID: "ip_address_group_id": 123
//
//   - Project Groups (type="project"): Created and managed by individual members for
//     their own use
//   - Used for project firewalls with fine-grained access control
//   - Examples: office networks, VPN endpoints, admin workstations
//   - Referenced in project firewall rules using @groupname syntax: "source":
//     "@office_networks"
//
// Usage in Firewall Rules:
//
// - Project Firewall: "source": "@office_networks" (uses project type groups only)
// - Geo Firewall: "group_name": "@ie_v4" (uses geo type groups only)
//
// Examples:
//
//   - Block traffic from Ireland: Create geo firewall rule with group_name of
//     Ireland group
//   - Allow access from office: Create project firewall rule with source
//     "@office_networks"
//   - Compliance geo-blocking: Use global country groups referenced by ID in geo
//     firewall rules
//
// NetworkIPGroupService contains methods and other services that help with
// interacting with the gocloudcix API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkIPGroupService] method instead.
type NetworkIPGroupService struct {
	Options []option.RequestOption
}

// NewNetworkIPGroupService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetworkIPGroupService(opts ...option.RequestOption) (r NetworkIPGroupService) {
	r = NetworkIPGroupService{}
	r.Options = opts
	return
}

// Create a new Network IP Group. Specify a name, IP version (IPv4 or IPv6), and a
// list of CIDR networks to include in the group. Groups can be used in firewall
// rules for access control. Only project type Network IP Groups can be created.
func (r *NetworkIPGroupService) New(ctx context.Context, body NetworkIPGroupNewParams, opts ...option.RequestOption) (res *NetworkIPGroup, err error) {
	var env NetworkIPGroupResponse
	opts = slices.Concat(r.Options, opts)
	path := "network/ip_groups/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Retrieve detailed information about a specific Network IP Group including its
// name, IP version, and list of CIDR networks.
func (r *NetworkIPGroupService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *NetworkIPGroup, err error) {
	var env NetworkIPGroupResponse
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("network/ip_groups/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Update a Network IP Group's configuration. Modify the group name, IP version, or
// the list of CIDR networks included in the group. Only project type Network IP
// Groups can be updated.
func (r *NetworkIPGroupService) Update(ctx context.Context, id int64, body NetworkIPGroupUpdateParams, opts ...option.RequestOption) (res *NetworkIPGroup, err error) {
	var env NetworkIPGroupResponse
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("network/ip_groups/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Retrieve a paginated list of Network IP Groups.
//
// ## Filtering
//
// The following fields and modifiers can be used to filter records from the list;
//
// - created (gt, gte, in, isnull, lt, lte, range)
// - id (gt, gte, in, isnull, lt, lte, range)
// - name (in, icontains, iendswith, iexact, istartswith)
// - type
// - updated (gt, gte, in, isnull, lt, lte, range)
// - version (gt, gte, in, isnull, lt, lte, range)
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
// - created
// - updated
// - version
//
// To reverse the ordering, simply prepend a `-` character to the request. So
// `?order=field` orders by `field` in ascending order, while `?order=-field`
// orders in descending order instead.
func (r *NetworkIPGroupService) List(ctx context.Context, query NetworkIPGroupListParams, opts ...option.RequestOption) (res *NetworkIPGroupListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "network/ip_groups/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete an Network IP Group from the system. Only project type Network IP Groups
// can be updated. Groups currently in use by firewall rules cannot be deleted.
func (r *NetworkIPGroupService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("network/ip_groups/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type NetworkIPGroup struct {
	// The ID of the Network IP Goup record
	ID int64 `json:"id" api:"required"`
	// An array of CIDR addresses in the Network IP Group.
	Cidrs []string `json:"cidrs" api:"required"`
	// Timestamp, in ISO format, of when the Network IP Group was created.
	Created string `json:"created" api:"required"`
	// The name of the Network IP Group.
	Name string `json:"name" api:"required"`
	// The type of the Network IP Group
	Type string `json:"type" api:"required"`
	// Timestamp, in ISO format, of when the Network IP Group was last updated.
	Updated string `json:"updated" api:"required"`
	// The absolute URL of the Network IP Group record that can be used to perform
	// `Read`, `Update` and `Delete`
	Uri string `json:"uri" api:"required"`
	// The IP Version of the CIDRs in the group.
	Version int64 `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cidrs       respjson.Field
		Created     respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		Updated     respjson.Field
		Uri         respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkIPGroup) RawJSON() string { return r.JSON.raw }
func (r *NetworkIPGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkIPGroupResponse struct {
	Content NetworkIPGroup `json:"content"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkIPGroupResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkIPGroupResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Cidrs, Name are required.
type NetworkIPGroupUpdateParam struct {
	// An array of CIDR addresses in the IP Address Group. Can include individual IP
	// addresses (e.g., "91.103.3.36") or network ranges (e.g., "90.103.2.0/24"). All
	// addresses must match the specified IP version. Use these groups in firewall
	// rules to allow/block traffic from multiple locations with a single rule.
	Cidrs []string `json:"cidrs,omitzero" api:"required"`
	// The name to be given to the new IP Address Group. Used to identify the group
	// when creating firewall rules or geo-filters. Must start with a letter and
	// contain only letters, numbers, underscores, and hyphens.
	Name string `json:"name" api:"required"`
	// The IP version of the IP Address Group Objects in the IP Address Group. Accepted
	// versions are 4 and 6. If not sent, it will default to 4.
	Version param.Opt[int64] `json:"version,omitzero"`
	paramObj
}

func (r NetworkIPGroupUpdateParam) MarshalJSON() (data []byte, err error) {
	type shadow NetworkIPGroupUpdateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkIPGroupUpdateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkIPGroupListResponse struct {
	Metadata ListMetadata     `json:"_metadata"`
	Content  []NetworkIPGroup `json:"content"`
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
func (r NetworkIPGroupListResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkIPGroupListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkIPGroupNewParams struct {
	// An array of CIDR addresses in the IP Address Group. Can include individual IP
	// addresses (e.g., "91.103.3.36") or network ranges (e.g., "90.103.2.0/24"). All
	// addresses must match the specified IP version. Use these groups in firewall
	// rules to allow/block traffic from multiple locations with a single rule.
	Cidrs []string `json:"cidrs,omitzero" api:"required"`
	// The name to be given to the new IP Address Group. Used to identify the group
	// when creating firewall rules or geo-filters. Must start with a letter and
	// contain only letters, numbers, underscores, and hyphens.
	Name string `json:"name" api:"required"`
	// The IP version of the IP Address Group Objects in the IP Address Group. Accepted
	// versions are 4 and 6. If not sent, it will default to 4.
	Version param.Opt[int64] `json:"version,omitzero"`
	paramObj
}

func (r NetworkIPGroupNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NetworkIPGroupNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkIPGroupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkIPGroupUpdateParams struct {
	NetworkIPGroupUpdate NetworkIPGroupUpdateParam
	paramObj
}

func (r NetworkIPGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.NetworkIPGroupUpdate)
}
func (r *NetworkIPGroupUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkIPGroupListParams struct {
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

// URLQuery serializes [NetworkIPGroupListParams]'s query parameters as
// `url.Values`.
func (r NetworkIPGroupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
