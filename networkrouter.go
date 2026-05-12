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

// Management of Virtual Network Routers
//
// This module provides API endpoints for managing virtual network routers in the
// CloudCIX Compute platform. Each project can have one virtual router that
// provides network connectivity and routing between your cloud resources and
// external networks and one or more static routes. The router manages one or more
// private networks (subnets) and handles traffic routing, NAT, and network
// isolation for your project's virtual machines and containers.
//
// Network Router Type:
//
//  1. Project Router (type: "router") - Manage IP forwarding, and participate in
//     routing decisions within isolated network topologies.
//  2. Static Routes (type: "static_route") - Define a fixed route entry within the
//     Project Router's routing table. It maps a destination network to a nexthop
//     IP, enabling deterministic packet forwarding without dynamic updates.
//
// Available operations:
//
//   - List and filter virtual routers from all your projects
//   - Create a project's router with one or more private network definitions (RFC
//     1918 address ranges)
//   - Retrieve detailed router configuration including networks, IP addresses, and
//     state
//   - Update router by adding networks, changing network names, or changing router
//     state
//
// Network Management: When creating or adding networks, you only specify the IPv4
// CIDR range and name. The system automatically generates VLAN IDs and IPv6 ranges
// based on regional availability. When updating a router to add new networks, you
// must include all existing networks (with their auto-generated VLAN and IPv6
// properties) to preserve them, plus any new networks (with only IPv4 CIDR and
// name specified). Existing network IPv4/IPv6 ranges and VLANs cannot be modified,
// but network names can be updated by including the name field with existing
// networks.
//
// Each router includes its associated project, public IP addresses (IPv4/IPv6),
// private networks with VLANs, and current state. You can add additional private
// networks to an existing router through the update operation.
//
// NetworkRouterService contains methods and other services that help with
// interacting with the gocloudcix API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkRouterService] method instead.
type NetworkRouterService struct {
	Options []option.RequestOption
}

// NewNetworkRouterService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetworkRouterService(opts ...option.RequestOption) (r NetworkRouterService) {
	r = NetworkRouterService{}
	r.Options = opts
	return
}

// Create a new Network Routers Resource entry using data given by user for a
// specified type, "router" or "static_route".
//
// Create a new netwotk router for a project. Each project can have only one of the
// type "router". For this type, optionally define one or more private networks
// using RFC 1918 addresses (10.x.x.x, 172.16-31.x.x, 192.168.x.x). If no networks
// are specified, a default 10.0.0.1/24 network is created. Additional networks can
// be added later through the update operation.
//
// One or more of the Network Router of the type "static_route" can be added to the
// Project to manage mapping a destination network to a nexthop IP.
func (r *NetworkRouterService) New(ctx context.Context, body NetworkRouterNewParams, opts ...option.RequestOption) (res *NetworkRouter, err error) {
	var env NetworkRouterResponse
	opts = slices.Concat(r.Options, opts)
	path := "network/routers/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Retrieve detailed information about a specific network router configuration,
// including its type (router or static_route), associated networking details, and
// current state.
func (r *NetworkRouterService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *NetworkRouter, err error) {
	var env NetworkRouterResponse
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("network/routers/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Update a network router to modify its name, add new private networks, rename
// existing networks, or change its state. When adding networks, specify new IPv4
// CIDR ranges. When updating existing networks, reference them by VLAN ID.
//
// For static routes, they can be renamed or have there state changed,
func (r *NetworkRouterService) Update(ctx context.Context, id int64, body NetworkRouterUpdateParams, opts ...option.RequestOption) (res *NetworkRouter, err error) {
	var env NetworkRouterResponse
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("network/routers/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Retrieve a paginated list of network routers from all your projects. Each
// project has at most one router. Supports filtering by type (router or
// static_route), project, state, name, and other attributes. Results can be
// ordered and paginated.
//
// ## Filtering
//
// The following fields and modifiers can be used to filter records from the list;
//
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
//
// To reverse the ordering, simply prepend a `-` character to the request. So
// `?order=field` orders by `field` in ascending order, while `?order=-field`
// orders in descending order instead.
func (r *NetworkRouterService) List(ctx context.Context, query NetworkRouterListParams, opts ...option.RequestOption) (res *NetworkRouterListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "network/routers/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// This method allows the User to delete a Network Router immediately. The grace
// period will be set to 0 which will facilitate the virtual infrastructure for the
// Network Router to be scrubbed from the region without the option to restore.
func (r *NetworkRouterService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("network/routers/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type BaseIPAddress struct {
	// The ID of the IPAddress record.
	ID int64 `json:"id" api:"required"`
	// The IP address of the IPAddress record.
	Address string `json:"address" api:"required"`
	// Timestamp, in ISO format, of when the IPAddress record was created.
	Created string `json:"created" api:"required"`
	// A verbose name given to the IPAddress record.
	Name string `json:"name" api:"required"`
	// The note attached to IPAddress that made it.
	Notes    string                `json:"notes" api:"required"`
	PublicIP BaseIPAddressPublicIP `json:"public_ip" api:"required"`
	// The ID of the Public IPAddress record.
	PublicIPID int64 `json:"public_ip_id" api:"required"`
	// The ID of the Subnet record.
	SubnetID int64 `json:"subnet_id" api:"required"`
	// Timestamp, in ISO format, of when the IPAddress record was last updated.
	Updated string `json:"updated" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Address     respjson.Field
		Created     respjson.Field
		Name        respjson.Field
		Notes       respjson.Field
		PublicIP    respjson.Field
		PublicIPID  respjson.Field
		SubnetID    respjson.Field
		Updated     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseIPAddress) RawJSON() string { return r.JSON.raw }
func (r *BaseIPAddress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BaseIPAddressPublicIP struct {
	// The ID of the Public IPAddress record.
	ID int64 `json:"id" api:"required"`
	// The actual address of the Public IPAddress record.
	Address string `json:"address" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Address     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BaseIPAddressPublicIP) RawJSON() string { return r.JSON.raw }
func (r *BaseIPAddressPublicIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkRouter struct {
	// The ID of the Router Resource record
	ID int64 `json:"id" api:"required"`
	// Timestamp, in ISO format, of when the Router Resource record was created.
	Created string `json:"created" api:"required"`
	// Number of days after a user sets the state of the Router to Scrub (8) before it
	// is executed by robot. The default value is 7 days for a Router.
	GracePeriod int64                 `json:"grace_period" api:"required"`
	Metadata    NetworkRouterMetadata `json:"metadata" api:"required"`
	// The user-friendly name given to this Router Resource instance
	Name string `json:"name" api:"required"`
	// An array of the list of networks defined on the Router
	Networks []NetworkRouterNetwork `json:"networks" api:"required"`
	// The id of the Project that this Router Resource belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// An array of the specs for the Router Resource
	Specs []Bom `json:"specs" api:"required"`
	// The current state of the Router Resource
	State string `json:"state" api:"required"`
	// The type of the Network Router
	Type string `json:"type" api:"required"`
	// Timestamp, in ISO format, of when the Router Resource record was last updated.
	Updated string `json:"updated" api:"required"`
	// URL that can be used to run methods in the API associated with the Network
	// Routers instance.
	Uri string `json:"uri" api:"required" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		GracePeriod respjson.Field
		Metadata    respjson.Field
		Name        respjson.Field
		Networks    respjson.Field
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
func (r NetworkRouter) RawJSON() string { return r.JSON.raw }
func (r *NetworkRouter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkRouterMetadata struct {
	Ipv4Address BaseIPAddress `json:"ipv4_address" api:"required"`
	// The ID of the assigned public IPv4 address for the Router.
	Ipv4AddressID int64         `json:"ipv4_address_id" api:"required"`
	Ipv6Address   BaseIPAddress `json:"ipv6_address" api:"required"`
	// The ID of the assigned public IPv6 address for the Router.
	Ipv6AddressID int64 `json:"ipv6_address_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ipv4Address   respjson.Field
		Ipv4AddressID respjson.Field
		Ipv6Address   respjson.Field
		Ipv6AddressID respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkRouterMetadata) RawJSON() string { return r.JSON.raw }
func (r *NetworkRouterMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkRouterNetwork struct {
	// The destination address range of the target network of the static route.
	// Returned if the type is "static_route".
	Destination string `json:"destination"`
	// The IPv4 address range of the network. Returned if the type is "router".
	Ipv4 string `json:"ipv4"`
	// The IPv6 address range of the network. Returned if the type is "router".
	Ipv6 string `json:"ipv6"`
	// The name of the network. Returned if the type is "router".
	Name string `json:"name"`
	// Flag indicating if traffic from the destination can route to the Public
	// internet. Returned if the type is "static_route".
	Nat bool `json:"nat"`
	// An IP address from one of the networks configured on the Router in the Project
	// to forward the packet to. Returned if the type is "static_route".
	Nexthop string `json:"nexthop"`
	// The VLAN ID of the network. Returned if the type is "router".
	Vlan int64 `json:"vlan"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Destination respjson.Field
		Ipv4        respjson.Field
		Ipv6        respjson.Field
		Name        respjson.Field
		Nat         respjson.Field
		Nexthop     respjson.Field
		Vlan        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkRouterNetwork) RawJSON() string { return r.JSON.raw }
func (r *NetworkRouterNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkRouterResponse struct {
	Content NetworkRouter `json:"content"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkRouterResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkRouterResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property State is required.
type NetworkRouterUpdateParam struct {
	// Change the state of the Network Router, triggering the CloudCIX Robot to perform
	// the requested action.
	//
	// Available state transitions:
	//
	// From running state, you can transition to:
	//
	//   - update_running - Apply pending configuration changes while keeping the router
	//     operational
	//   - delete - Mark the router for deletion (requires all other project resources to
	//     be deleted first)
	//
	// From delete_queue state, you can transition to:
	//
	// - restart - Restore a router that was previously marked for deletion
	//
	// Note: To delete a router, all other resources in the project must first be in
	// one of these states: delete, delete_queue, or deleting.
	State string `json:"state" api:"required"`
	// Metadata for the Netork Routers of the type "static_route".
	Metadata NetworkRouterUpdateMetadataParam `json:"metadata,omitzero"`
	// Networks for the Netork Routers of the type "router".
	//
	// An array of the list of networks defined on the Network Router. Existing
	// networks (vlan property is not None) can have their names updated but IPv4/IPv6
	// ranges and VLAN cannot be modified. To create a new network on the Network
	// Router, append an object to the list with an `ipv4` key for an available RFC
	// 1918 address range. The `ipv6` and `vlan` values will be generated based on what
	// is available in the region.
	Networks []NetworkRouterUpdateNetworkParam `json:"networks,omitzero"`
	paramObj
}

func (r NetworkRouterUpdateParam) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouterUpdateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkRouterUpdateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata for the Netork Routers of the type "static_route".
type NetworkRouterUpdateMetadataParam struct {
	// CIDR notation of the destination address range of the target network of the
	// static route. The destination cannot be updated.
	Destination param.Opt[string] `json:"destination,omitzero"`
	// Flag indicating if traffic from the destination can be routed to the Public
	// internet via the Project's Router.
	Nat param.Opt[bool] `json:"nat,omitzero"`
	// An IP address from one of the networks configured on the Router in the Project
	// to forward the packet to. The nexthop cannot be updated.
	Nexthop param.Opt[string] `json:"nexthop,omitzero"`
	paramObj
}

func (r NetworkRouterUpdateMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouterUpdateMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkRouterUpdateMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkRouterUpdateNetworkParam struct {
	// The IPv4 address range of the network
	Ipv4 param.Opt[string] `json:"ipv4,omitzero"`
	// The IPv6 address range of the network
	Ipv6 param.Opt[string] `json:"ipv6,omitzero"`
	// The name of the network
	Name param.Opt[string] `json:"name,omitzero"`
	// The VLAN ID of the network.
	Vlan param.Opt[int64] `json:"vlan,omitzero"`
	paramObj
}

func (r NetworkRouterUpdateNetworkParam) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouterUpdateNetworkParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkRouterUpdateNetworkParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkRouterListResponse struct {
	Metadata ListMetadata    `json:"_metadata"`
	Content  []NetworkRouter `json:"content"`
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
func (r NetworkRouterListResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkRouterListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkRouterNewParams struct {
	// The ID of the User's Project into which this Network Router should be added.
	ProjectID int64 `json:"project_id" api:"required"`
	// The type of Network Router to create. Valid options are:
	//
	//   - "router" A virtual route that manages IP forwarding, and participate in
	//     routing decisions for the Project.
	//   - "static_route" Maps a destination network to a nexthop IP, enabling
	//     deterministic packet forwarding.
	Type param.Opt[string] `json:"type,omitzero"`
	// Required if type is "static_route".
	//
	// Metadata for the Static Route resource.
	Metadata NetworkRouterNewParamsMetadata `json:"metadata,omitzero"`
	// Option if type is "router". If not sent, defaults will be applied.
	//
	// An array of the list of networks defined on the Router. To create a new network
	// on the Network Router, append an object to the list with an `ipv4` key for an
	// available RFC 1918 address range. The `ipv6` and `vlan` values will be generated
	// based on what is available in the region. If networks is not sent, the default
	// address range 10.0.0.1/24 will be assigned to `ipv4`.
	Networks []NetworkRouterNewParamsNetwork `json:"networks,omitzero"`
	paramObj
}

func (r NetworkRouterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkRouterNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Required if type is "static_route".
//
// Metadata for the Static Route resource.
type NetworkRouterNewParamsMetadata struct {
	// CIDR notation of the destination address range of the target network of the
	// static route.
	//
	// Note:
	//
	//   - The sent address range cannot overlap with the destination of other Static
	//     Routes in the same Project.
	//   - The sent address range can overlap with the networks configured on the Router
	//     in the Project.
	//   - The sent address range cannot overlap with the "remote_ts" of Network VPNs in
	//     the same Project.
	Destination param.Opt[string] `json:"destination,omitzero"`
	// Flag indicating if traffic from the destination can be routed to the Public
	// internet via the Project's Router. It will default to False if not sent.
	Nat param.Opt[bool] `json:"nat,omitzero"`
	// An IP address from one of the networks configured on the Router in the Project
	// to forward the packet to.
	Nexthop param.Opt[string] `json:"nexthop,omitzero"`
	paramObj
}

func (r NetworkRouterNewParamsMetadata) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouterNewParamsMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkRouterNewParamsMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkRouterNewParamsNetwork struct {
	// The IPv4 address range of the network
	Ipv4 param.Opt[string] `json:"ipv4,omitzero"`
	// The IPv6 address range of the network
	Ipv6 param.Opt[string] `json:"ipv6,omitzero"`
	// The name of the network
	Name param.Opt[string] `json:"name,omitzero"`
	// The VLAN of the network
	Vlan param.Opt[int64] `json:"vlan,omitzero"`
	paramObj
}

func (r NetworkRouterNewParamsNetwork) MarshalJSON() (data []byte, err error) {
	type shadow NetworkRouterNewParamsNetwork
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkRouterNewParamsNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkRouterUpdateParams struct {
	NetworkRouterUpdate NetworkRouterUpdateParam
	paramObj
}

func (r NetworkRouterUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.NetworkRouterUpdate)
}
func (r *NetworkRouterUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkRouterListParams struct {
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

// URLQuery serializes [NetworkRouterListParams]'s query parameters as
// `url.Values`.
func (r NetworkRouterListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
