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

// Management of Network VPNs Supported Types are:
//
// - site-to-site
//
// NetworkVpnService contains methods and other services that help with interacting
// with the gocloudcix API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkVpnService] method instead.
type NetworkVpnService struct {
	Options []option.RequestOption
}

// NewNetworkVpnService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetworkVpnService(opts ...option.RequestOption) (r NetworkVpnService) {
	r = NetworkVpnService{}
	r.Options = opts
	return
}

// Create a new Network VPN Resource entry using data given by user.
func (r *NetworkVpnService) New(ctx context.Context, body NetworkVpnNewParams, opts ...option.RequestOption) (res *NetworkVpn, err error) {
	var env NetworkVpnResponse
	opts = slices.Concat(r.Options, opts)
	path := "network/vpns/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Attempt to read a Network VPN Resource record by the given `id`, returning a 404
// if it does not exist
func (r *NetworkVpnService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *NetworkVpn, err error) {
	var env NetworkVpnResponse
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("network/vpns/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Attempt to update a Network VPN Resource record by the given `id`, returning a
// 404 if it does not exist
func (r *NetworkVpnService) Update(ctx context.Context, id int64, body NetworkVpnUpdateParams, opts ...option.RequestOption) (res *NetworkVpn, err error) {
	var env NetworkVpnResponse
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("network/vpns/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Retrieve a list of Network VPNs Resources
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
func (r *NetworkVpnService) List(ctx context.Context, query NetworkVpnListParams, opts ...option.RequestOption) (res *NetworkVpnListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "network/vpns/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// This method allows the User to delete a specified Network VPN.
func (r *NetworkVpnService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("network/vpns/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type NetworkVpn struct {
	// The ID of the Network VPN record
	ID int64 `json:"id" api:"required"`
	// Timestamp, in ISO format, of when the Network VPN record was created.
	Created string `json:"created" api:"required"`
	// The metadata for the configuration of the IKE and IPSec phases of the Network
	// VPN.
	Metadata NetworkVpnMetadata `json:"metadata" api:"required"`
	// The user-friendly name given to this Network VPN instance
	Name string `json:"name" api:"required"`
	// The id of the Project that this Network VPN belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	// An array of the specs for the Network VPN
	Specs []Bom `json:"specs" api:"required"`
	// The current state of the Network VPN
	State string `json:"state" api:"required"`
	// The type of the Network VPN
	Type string `json:"type" api:"required"`
	// Timestamp, in ISO format, of when the Network VPN record was last updated.
	Updated string `json:"updated" api:"required"`
	// URL that can be used to run methods in the API associated with the Network VPN
	// instance.
	Uri string `json:"uri" api:"required" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
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
func (r NetworkVpn) RawJSON() string { return r.JSON.raw }
func (r *NetworkVpn) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The metadata for the configuration of the IKE and IPSec phases of the Network
// VPN.
type NetworkVpnMetadata struct {
	// An array of Child SAs (Security Associations) of the routes for the Network VPN.
	ChildSas []NetworkVpnMetadataChildSa `json:"child_sas"`
	// Authentication algorithms for the IKE phase.
	IkeAuthentication string `json:"ike_authentication"`
	// Diffie-Helmen groups for the IKE phase.
	IkeDhGroups string `json:"ike_dh_groups"`
	// Encryption algorithms for the IKE phase.
	IkeEncryption string `json:"ike_encryption"`
	// The gateway type (public_ip or hostname ) for the IKE phase.
	IkeGatewayType string `json:"ike_gateway_type"`
	// The gateway value (IP or hostname) for the IKE phase.
	IkeGatewayValue string `json:"ike_gateway_value"`
	// The lifetime of the IKE phase in seconds.
	IkeLifetime int64 `json:"ike_lifetime"`
	// The local identifier of the IKE phase in the region.
	IkeLocalIdentifier string `json:"ike_local_identifier"`
	// The pre shared key of the IKE phase.
	IkePreSharedKey string `json:"ike_pre_shared_key"`
	// The remote identifier of the IKE phase on the customers side of the VPN.
	IkeRemoteIdentifier string `json:"ike_remote_identifier"`
	// The version of the IKE phase
	IkeVersion int64 `json:"ike_version"`
	// Authentication algorithms for the IPSec phase.
	IpsecAuthentication string `json:"ipsec_authentication"`
	// Encryption algorithms for the IPSec phase.
	IpsecEncryption string `json:"ipsec_encryption"`
	// The establish time for the IPSec phase.
	IpsecEstablishTime string `json:"ipsec_establish_time"`
	// The lifetime of the IPSec phase in seconds.
	IpsecLifetime int64 `json:"ipsec_lifetime"`
	// Perfect Forward Secrecy groups for the IPSec phase.
	IpsecPfsGroups string `json:"ipsec_pfs_groups"`
	// STIF number for the Network VPN.
	StifNumber int64 `json:"stif_number"`
	// Flag for it traffic selectors are enabled.
	TrafficSelector bool `json:"traffic_selector"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ChildSas            respjson.Field
		IkeAuthentication   respjson.Field
		IkeDhGroups         respjson.Field
		IkeEncryption       respjson.Field
		IkeGatewayType      respjson.Field
		IkeGatewayValue     respjson.Field
		IkeLifetime         respjson.Field
		IkeLocalIdentifier  respjson.Field
		IkePreSharedKey     respjson.Field
		IkeRemoteIdentifier respjson.Field
		IkeVersion          respjson.Field
		IpsecAuthentication respjson.Field
		IpsecEncryption     respjson.Field
		IpsecEstablishTime  respjson.Field
		IpsecLifetime       respjson.Field
		IpsecPfsGroups      respjson.Field
		StifNumber          respjson.Field
		TrafficSelector     respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkVpnMetadata) RawJSON() string { return r.JSON.raw }
func (r *NetworkVpnMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkVpnMetadataChildSa struct {
	// The CIDR notation of a subnet configured on the Network Router in the same
	// Project
	LocalTs string `json:"local_ts"`
	// CIDR notation of a subnet on the Customer side of the Network VPN.
	RemoteTs string `json:"remote_ts"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LocalTs     respjson.Field
		RemoteTs    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkVpnMetadataChildSa) RawJSON() string { return r.JSON.raw }
func (r *NetworkVpnMetadataChildSa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkVpnResponse struct {
	Content NetworkVpn `json:"content"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NetworkVpnResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkVpnResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkVpnUpdateParam struct {
	// The user-friendly name for the Network VPN type. If not sent, it will default to
	// current name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Change the state of the Network VPN, triggering the CloudCIX Robot to perform
	// the requested action. Users can only request state changes from certain current
	// states:
	//
	// - running -> update_running or delete
	State param.Opt[string] `json:"state,omitzero"`
	// Optional. The metadata required to configure the Network VPN instance
	Metadata NetworkVpnUpdateMetadataParam `json:"metadata,omitzero"`
	paramObj
}

func (r NetworkVpnUpdateParam) MarshalJSON() (data []byte, err error) {
	type shadow NetworkVpnUpdateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkVpnUpdateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional. The metadata required to configure the Network VPN instance
type NetworkVpnUpdateMetadataParam struct {
	// Optional. A string containing a comma separated array of authentication
	// algorithms for the IKE phase of the Network VPN. If not sent, it will default to
	// the current value.
	//
	// The IKE phase authentication algorithms supported are;
	//
	// - `SHA1`
	// - `SHA256`
	// - `SHA384`
	//
	// Please ensure that each entry in the array matches one of the above strings
	// exactly. Duplicate entries will be ignored.
	IkeAuthentication param.Opt[string] `json:"ike_authentication,omitzero"`
	// Optional. A string containing a comma separated array of Diffie-Helmen groups
	// for the IKE phase of the Network VPN. If not sent, it will default to the
	// current value.
	//
	// The IKE phase Diffie-Helmen groups supported are;
	//
	// - `Group 1`
	// - `Group 2`
	// - `Group 5`
	// - `Group 19`
	// - `Group 20`
	// - `Group 24`
	//
	// Please ensure that each entry in the array matches one of the above strings
	// exactly. Duplicate entries will be ignored.
	IkeDhGroups param.Opt[string] `json:"ike_dh_groups,omitzero"`
	// Optional. A string containing a comma separated array of encryption algorithms
	// for the IKE phase of the Network VPN. If not sent, it will default to the
	// current value.
	//
	// The IKE phase encryption algorithms supported are;
	//
	// - `128 bit AES-CBC`
	// - `192 bit AES-CBC`
	// - `256 bit AES-CBC`
	//
	// Please ensure that each entry in the array matches one of the above strings
	// exactly. Duplicate entries will be ignored.
	IkeEncryption param.Opt[string] `json:"ike_encryption,omitzero"`
	// The type of data that is stored in the `ike_gateway_value` field. It can only
	// either "public_ip" or "hostname". If not sent, it will default to the current
	// value.
	IkeGatewayType param.Opt[string] `json:"ike_gateway_type,omitzero"`
	// The value used as the IKE gateway for the Network VPN. The type for this value
	// depends on what type was sent for the "ike_gateway_type".
	//
	// For "public_ip", this value must be a string containing an IP address. For
	// "hostname", this value must be a valid hostname.
	IkeGatewayValue param.Opt[string] `json:"ike_gateway_value,omitzero"`
	// Optional. The lifetime of the IKE phase in seconds. Must be a value between 180
	// and 86400 inclusive. If not sent, it will default to the current value.
	IkeLifetime param.Opt[int64] `json:"ike_lifetime,omitzero"`
	// The pre shared key to use for setting up the IKE phase of the Network VPN.
	//
	// Note that the pre shared key cannot contain any of the following special
	// characters;
	//
	// - `"`
	// - `'`
	// - `@`
	// - `+`
	// - `-`
	// - `/`
	// - `\`
	// - `|`
	// - `=`
	IkePreSharedKey param.Opt[string] `json:"ike_pre_shared_key,omitzero"`
	// Optional. String value of the chosen version for the IKE phase. If not sent, it
	// will default to the current value.
	//
	// The IKE phase versions supported are;
	//
	// - `v1-only`
	// - `v2-only`
	//
	// Please ensure the sent string matches one of these exactly.
	IkeVersion param.Opt[string] `json:"ike_version,omitzero"`
	// Optional. A string containing a comma separated array of authentication
	// algorithms for the IPSec phase of the Site-to-Site Network VPN. If not sent, it
	// will default to the current value.
	//
	// The IPSec phase authentication algorithms supported are;
	//
	// - `SHA1`
	// - `SHA256`
	//
	// Please ensure that each entry in the array matches one of the above strings
	// exactly. Duplicate entries will be ignored.
	IpsecAuthentication param.Opt[string] `json:"ipsec_authentication,omitzero"`
	// Optional. A string containing a comma separated array of encryption algorithms
	// for the IPSEC phase of the Network VPN. If not sent, it will default to the
	// current value.
	//
	// The IPSEC phase encryption algorithms supported are;
	//
	// - `AES 128`
	// - `AES 192`
	// - `AES 256`
	// - `AES 128 GCM`
	// - `AES 192 GCM`
	// - `AES 256 GCM`
	//
	// Please ensure that each entry in the array matches one of the above strings
	// exactly. Duplicate entries will be ignored.
	IpsecEncryption param.Opt[string] `json:"ipsec_encryption,omitzero"`
	// Optional. String value of the chosen establish_time for the IPSec phase. If not
	// sent, it will default to the current value.
	//
	// The IPSec phase establish time values supported are;
	//
	// - `Immediately`
	// - `On Traffic`
	//
	// Please ensure the sent string matches one of these exactly.
	IpsecEstablishTime param.Opt[string] `json:"ipsec_establish_time,omitzero"`
	// Optional. The lifetime of the IPSec phase in seconds. It be a value between 180
	// and 86400 inclusive. If not sent, it will default to the current value.
	IpsecLifetime param.Opt[int64] `json:"ipsec_lifetime,omitzero"`
	// Optional. A string containing a comma separated array of Perfect Forward Secrecy
	// (PFS) groups for the IPSec phase of the Network VPN. If not sent, it will
	// default to the current value.
	//
	// The IPSec phase PFS groups supported are;
	//
	// - `Group 1`
	// - `Group 2`
	// - `Group 5`
	// - `Group 14`
	// - `Group 19`
	// - `Group 20`
	// - `Group 24`
	//
	// Please ensure that each entry in the array matches one of the above strings
	// exactly. Duplicate entries will be ignored.
	IpsecPfsGroups param.Opt[string] `json:"ipsec_pfs_groups,omitzero"`
	// Optional. Boolean value stating if traffic selectors are to be used in
	// configuring vpn tunnel. If not sent, it will default to the current value.
	//
	// By default, 0.0.0.0/0 will be used for the default local and remote CIDRs. If
	// true, then each of the local and remote CIDRs will be added to the configuration
	// negotiation with peer.
	TrafficSelector param.Opt[bool] `json:"traffic_selector,omitzero"`
	// An array of Child SAs (Security Associations) to create the initial routes for
	// the VPN.
	ChildSas []NetworkVpnUpdateMetadataChildSaParam `json:"child_sas,omitzero"`
	paramObj
}

func (r NetworkVpnUpdateMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow NetworkVpnUpdateMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkVpnUpdateMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkVpnUpdateMetadataChildSaParam struct {
	// The local Subnet in Network VPN's Project that will be configured as part of
	// this route in a Child SA.
	LocalTs param.Opt[string] `json:"local_ts,omitzero"`
	// CIDR notation of the Remote Subnet on the Customer side of the Network VPN that
	// should be given access through the VPN.
	//
	// Note:
	//
	//   - The sent address range can overlap with the subnets configured on the Router
	//     in the Project for the VPNS2S.
	//   - The sent address range cannot overlap with a remote subnet of another VPN in
	//     the same Project.
	RemoteTs param.Opt[string] `json:"remote_ts,omitzero"`
	paramObj
}

func (r NetworkVpnUpdateMetadataChildSaParam) MarshalJSON() (data []byte, err error) {
	type shadow NetworkVpnUpdateMetadataChildSaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkVpnUpdateMetadataChildSaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkVpnListResponse struct {
	Metadata ListMetadata `json:"_metadata"`
	Content  []NetworkVpn `json:"content"`
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
func (r NetworkVpnListResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkVpnListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkVpnNewParams struct {
	// The ID of the User's Project into which this Network VPN should be added.
	ProjectID int64 `json:"project_id" api:"required"`
	// The user-friendly name for the Network VPN. If not sent, it will default to the
	// name 'VPNS2S'
	Name param.Opt[string] `json:"name,omitzero"`
	// The type of Network VPN to create. Valid options are:
	//
	// - "site-to-site"
	Type param.Opt[string] `json:"type,omitzero"`
	// Optional. The metadata required to configure the Network VPN instance.
	Metadata NetworkVpnNewParamsMetadata `json:"metadata,omitzero"`
	paramObj
}

func (r NetworkVpnNewParams) MarshalJSON() (data []byte, err error) {
	type shadow NetworkVpnNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkVpnNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional. The metadata required to configure the Network VPN instance.
type NetworkVpnNewParamsMetadata struct {
	// Optional. A string containing a comma separated array of authentication
	// algorithms for the IKE phase of the Network VPN. If not sent, it will default to
	// `SHA384`.
	//
	// The IKE phase authentication algorithms supported are;
	//
	// - `SHA1`
	// - `SHA256`
	// - `SHA384`
	//
	// Please ensure that each entry in the array matches one of the above strings
	// exactly. Duplicate entries will be ignored.
	IkeAuthentication param.Opt[string] `json:"ike_authentication,omitzero"`
	// Optional. A string containing a comma separated array of Diffie-Helmen groups
	// for the IKE phase of the Network VPN. If not sent, it will default to
	// `Group 24`.
	//
	// The IKE phase Diffie-Helmen groups supported are;
	//
	// - `Group 1`
	// - `Group 2`
	// - `Group 5`
	// - `Group 19`
	// - `Group 20`
	// - `Group 24`
	//
	// Please ensure that each entry in the array matches one of the above strings
	// exactly. Duplicate entries will be ignored.
	IkeDhGroups param.Opt[string] `json:"ike_dh_groups,omitzero"`
	// Optional. A string containing a comma separated array of encryption algorithms
	// for the IKE phase of the Network VPN. If not sent, it will default to "256 bit
	// AES-CBC".
	//
	// The IKE phase encryption algorithms supported are;
	//
	// - `128 bit AES-CBC`
	// - `192 bit AES-CBC`
	// - `256 bit AES-CBC`
	//
	// Please ensure that each entry in the array matches one of the above strings
	// exactly. Duplicate entries will be ignored.
	IkeEncryption param.Opt[string] `json:"ike_encryption,omitzero"`
	// The type of data that is stored in the `ike_gateway_value` field. It can only
	// either "public_ip" or "hostname". If not sent, it will default to "public_ip".
	IkeGatewayType param.Opt[string] `json:"ike_gateway_type,omitzero"`
	// The value used as the IKE gateway for the Network VPN. The type for this value
	// depends on what type was sent for the "ike_gateway_type".
	//
	// For "public_ip", this value must be a string containing an IP address. For
	// "hostname", this value must be a valid hostname.
	IkeGatewayValue param.Opt[string] `json:"ike_gateway_value,omitzero"`
	// Optional. The lifetime of the IKE phase in seconds. Must be a value between 180
	// and 86400 inclusive. If not sent, it will default to 28800.
	IkeLifetime param.Opt[int64] `json:"ike_lifetime,omitzero"`
	// The pre shared key to use for setting up the IKE phase of the Network VPN.
	//
	// Note that the pre shared key cannot contain any of the following special
	// characters;
	//
	// - `"`
	// - `'`
	// - `@`
	// - `+`
	// - `-`
	// - `/`
	// - `\`
	// - `|`
	// - `=`
	IkePreSharedKey param.Opt[string] `json:"ike_pre_shared_key,omitzero"`
	// Optional. String value of the chosen version for the IKE phase. If not sent, it
	// will default to "v2-only".
	//
	// The IKE phase versions supported are;
	//
	// - `v1-only`
	// - `v2-only`
	//
	// Please ensure the sent string matches one of these exactly.
	IkeVersion param.Opt[string] `json:"ike_version,omitzero"`
	// Optional. A string containing a comma separated array of authentication
	// algorithms for the IPSec phase of the Site-to-Site Network VPN. If not sent, it
	// will default to "SHA256".
	//
	// The IPSec phase authentication algorithms supported are;
	//
	// - `SHA1`
	// - `SHA256`
	//
	// Please ensure that each entry in the array matches one of the above strings
	// exactly. Duplicate entries will be ignored.
	IpsecAuthentication param.Opt[string] `json:"ipsec_authentication,omitzero"`
	// Optional. A string containing a comma separated array of encryption algorithms
	// for the IPSEC phase of the Network VPN. If not sent, it will default to "AES
	// 256".
	//
	// The IPSEC phase encryption algorithms supported are;
	//
	// - `AES 128`
	// - `AES 192`
	// - `AES 256`
	// - `AES 128 GCM`
	// - `AES 192 GCM`
	// - `AES 256 GCM`
	//
	// Please ensure that each entry in the array matches one of the above strings
	// exactly. Duplicate entries will be ignored.
	IpsecEncryption param.Opt[string] `json:"ipsec_encryption,omitzero"`
	// Optional. String value of the chosen establish_time for the IPSec phase. If not
	// sent, it will default to "Immediately".
	//
	// The IPSec phase establish time values supported are;
	//
	// - `Immediately`
	// - `On Traffic`
	//
	// Please ensure the sent string matches one of these exactly.
	IpsecEstablishTime param.Opt[string] `json:"ipsec_establish_time,omitzero"`
	// Optional. The lifetime of the IPSec phase in seconds. It be a value between 180
	// and 86400 inclusive. If not sent, it will default to 3600.
	IpsecLifetime param.Opt[int64] `json:"ipsec_lifetime,omitzero"`
	// Optional. A string containing a comma separated array of Perfect Forward Secrecy
	// (PFS) groups for the IPSec phase of the Network VPN. If not sent, it will
	// default to "Group 20".
	//
	// The IPSec phase PFS groups supported are;
	//
	// - `Group 1`
	// - `Group 2`
	// - `Group 5`
	// - `Group 14`
	// - `Group 19`
	// - `Group 20`
	// - `Group 24`
	//
	// Please ensure that each entry in the array matches one of the above strings
	// exactly. Duplicate entries will be ignored.
	IpsecPfsGroups param.Opt[string] `json:"ipsec_pfs_groups,omitzero"`
	// Optional. Boolean value stating if traffic selectors are to be used in
	// configuring vpn tunnel. If not sent, it will default to false.
	//
	// By default, 0.0.0.0/0 will be used for the default local and remote CIDRs. If
	// true, then each of the local and remote CIDRs will be added to the configuration
	// negotiation with peer.
	TrafficSelector param.Opt[bool] `json:"traffic_selector,omitzero"`
	// An array of Child SAs (Security Associations) to create the initial routes for
	// the VPN.
	ChildSas []NetworkVpnNewParamsMetadataChildSa `json:"child_sas,omitzero"`
	paramObj
}

func (r NetworkVpnNewParamsMetadata) MarshalJSON() (data []byte, err error) {
	type shadow NetworkVpnNewParamsMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkVpnNewParamsMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkVpnNewParamsMetadataChildSa struct {
	// The local Subnet in Network VPN's Project that will be configured as part of
	// this route in a Child SA.
	LocalTs param.Opt[string] `json:"local_ts,omitzero"`
	// CIDR notation of the Remote Subnet on the Customer side of the Network VPN that
	// should be given access through the VPN.
	//
	// Note:
	//
	//   - The sent address range can overlap with the subnets configured on the Router
	//     in the Project for the VPNS2S.
	//   - The sent address range cannot overlap with a remote subnet of another VPN in
	//     the same Project.
	RemoteTs param.Opt[string] `json:"remote_ts,omitzero"`
	paramObj
}

func (r NetworkVpnNewParamsMetadataChildSa) MarshalJSON() (data []byte, err error) {
	type shadow NetworkVpnNewParamsMetadataChildSa
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkVpnNewParamsMetadataChildSa) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkVpnUpdateParams struct {
	NetworkVpnUpdate NetworkVpnUpdateParam
	paramObj
}

func (r NetworkVpnUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.NetworkVpnUpdate)
}
func (r *NetworkVpnUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkVpnListParams struct {
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

// URLQuery serializes [NetworkVpnListParams]'s query parameters as `url.Values`.
func (r NetworkVpnListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
