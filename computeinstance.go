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

// Management of Virtual Machine and Container Instances
//
// This module provides API endpoints for managing compute instances in the
// CloudCIX Compute platform. Compute instances are virtual machines or containers
// that run workloads in the cloud. Two instance types are supported: LXD (Linux
// containers and VMs) and Hyper-V (Windows virtual machines).
//
// Available operations:
//
//   - List and filter compute instances across your projects by type, state, or
//     other attributes
//   - Create new LXD or Hyper-V instances with specified resources, network
//     interfaces, and OS images
//   - Retrieve detailed configuration and status information for individual
//     instances
//   - Update instance specifications, network configuration, or change instance
//     state (stop, restart, delete)
//
// Each instance includes its associated project, resource specifications (CPU,
// RAM, storage), network interfaces, current state, and OS image information.
//
// Additional Resources:
//
// - [Information on available SKUs](https://www.cix.ie/#/services/cloud/public_cloud)
// - [Cloud-init User Data Tutorial](https://docs.cloudcix.com/tutorials/cloudinit_userdata_tutorial.html)
//
// ComputeInstanceService contains methods and other services that help with
// interacting with the gocloudcix API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeInstanceService] method instead.
type ComputeInstanceService struct {
	Options []option.RequestOption
}

// NewComputeInstanceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewComputeInstanceService(opts ...option.RequestOption) (r ComputeInstanceService) {
	r = ComputeInstanceService{}
	r.Options = opts
	return
}

// Create a new compute instance (LXD or Hyper-V). Specify the instance type,
// project, OS image, resource specifications (SKUs for CPU, RAM, storage), and
// network interfaces. The instance will be provisioned and started automatically.
func (r *ComputeInstanceService) New(ctx context.Context, body ComputeInstanceNewParams, opts ...option.RequestOption) (res *ComputeInstance, err error) {
	var env ComputeInstanceResponse
	opts = slices.Concat(r.Options, opts)
	path := "compute/instances/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Retrieve detailed information about a specific compute instance, including its
// type (LXD or Hyper-V), resource specifications, network interfaces, OS image,
// current state, and project information.
func (r *ComputeInstanceService) Get(ctx context.Context, id int64, opts ...option.RequestOption) (res *ComputeInstance, err error) {
	var env ComputeInstanceResponse
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("compute/instances/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Update a compute instance to modify its configuration or change its state. You
// can update resource specifications (CPU, RAM, storage), network interfaces, or
// change the instance state (stop, restart, delete, update_running,
// update_stopped).
func (r *ComputeInstanceService) Update(ctx context.Context, id int64, body ComputeInstanceUpdateParams, opts ...option.RequestOption) (res *ComputeInstance, err error) {
	var env ComputeInstanceResponse
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("compute/instances/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Content
	return res, nil
}

// Retrieve a paginated list of compute instances (LXD and/or Hyper-V) across your
// projects. Supports filtering by type, project, state, name, and other
// attributes. Results can be ordered and paginated.
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
// - state
//
// To reverse the ordering, simply prepend a `-` character to the request. So
// `?order=field` orders by `field` in ascending order, while `?order=-field`
// orders in descending order instead.
func (r *ComputeInstanceService) List(ctx context.Context, query ComputeInstanceListParams, opts ...option.RequestOption) (res *ComputeInstanceListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "compute/instances/"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// This method allows the User to delete a Compute Instance immediately. The grace
// period will be set to 0 which will facilitate the virtual infrastructure for the
// Compute Instance to be scrubbed from the region without the option to restore.
func (r *ComputeInstanceService) Delete(ctx context.Context, id int64, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := fmt.Sprintf("compute/instances/%v/", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type Bom struct {
	// How many units of a billable entity that a Resource utilises
	Quantity int64 `json:"quantity" api:"required"`
	// An identifier for a billable entity that a Resource utilises
	SKUName string `json:"sku_name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Quantity    respjson.Field
		SKUName     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Bom) RawJSON() string { return r.JSON.raw }
func (r *Bom) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeInstance struct {
	// The ID of the Compute Instance record
	ID int64 `json:"id" api:"required"`
	// Timestamp, in ISO format, of when the Compute Instance record was created.
	Created string `json:"created" api:"required"`
	// Number of days after a user sets the state of the Compute Instance Resource to
	// Scrub (8) before it is executed by robot.
	GracePeriod int64 `json:"grace_period" api:"required"`
	// Array of the interfaces for the Compute Instance
	Interfaces []ComputeInstanceInterface `json:"interfaces" api:"required"`
	// The metadata details of the Compute Instance
	Metadata ComputeInstanceMetadata `json:"metadata" api:"required"`
	// The human-friendly name given to this Compute Instance
	Name string `json:"name" api:"required"`
	// The id of the Project that this Compute Instance belongs to
	ProjectID int64 `json:"project_id" api:"required"`
	Specs     Bom   `json:"specs" api:"required"`
	// The current state of the Compute Instance
	State string `json:"state" api:"required"`
	// The type of the Compute Instance
	Type string `json:"type" api:"required"`
	// Timestamp, in ISO format, of when the Compute Instance record was last updated.
	Updated string `json:"updated" api:"required"`
	// URL that can be used to run methods in the API associated with the Compute
	// Instance.
	Uri string `json:"uri" api:"required" format:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Created     respjson.Field
		GracePeriod respjson.Field
		Interfaces  respjson.Field
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
func (r ComputeInstance) RawJSON() string { return r.JSON.raw }
func (r *ComputeInstance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeInstanceInterface struct {
	// Indicates if this interface functions as the network gateway for the Compute
	// instance.
	Gateway bool `json:"gateway"`
	// An array of the IPv4 addresses on the Interface.
	Ipv4Addresses []ComputeInstanceInterfaceIpv4Address `json:"ipv4_addresses"`
	// An array of the IPv6 addresses on the Interface.
	Ipv6Addresses []ComputeInstanceInterfaceIpv6Address `json:"ipv6_addresses"`
	// The VLAN assigned to this Interface
	Vlan int64 `json:"vlan"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Gateway       respjson.Field
		Ipv4Addresses respjson.Field
		Ipv6Addresses respjson.Field
		Vlan          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeInstanceInterface) RawJSON() string { return r.JSON.raw }
func (r *ComputeInstanceInterface) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeInstanceInterfaceIpv4Address struct {
	// The ID of the IP Address record on the Interface.
	ID int64 `json:"id"`
	// The IP address on the Interface.
	Address string `json:"address"`
	// Flag indicating if the IP Address is NATted to a Public IP Address
	Nat bool `json:"nat"`
	// The Public IP address that the address is NATted to.
	PublicIP string `json:"public_ip"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Address     respjson.Field
		Nat         respjson.Field
		PublicIP    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeInstanceInterfaceIpv4Address) RawJSON() string { return r.JSON.raw }
func (r *ComputeInstanceInterfaceIpv4Address) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeInstanceInterfaceIpv6Address struct {
	// The ID of the IP Address record on the Interface.
	ID int64 `json:"id"`
	// The IP address on the Interface.
	Address string `json:"address"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Address     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeInstanceInterfaceIpv6Address) RawJSON() string { return r.JSON.raw }
func (r *ComputeInstanceInterfaceIpv6Address) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The metadata details of the Compute Instance
type ComputeInstanceMetadata struct {
	// DNS server IP addresses for the Compute Instance. Returned if the type is "lxd".
	DNS string `json:"dns"`
	// Primary DNS server IPv4 address for the VM. Returned if the type is "hyperv".
	Dns4 string `json:"dns4"`
	// Primary DNS server IPv6 address for the VM. Returned if the type is "hyperv".
	Dns6 string `json:"dns6"`
	// The name of the instance type. It will be either "container" or "vm". Returned
	// if the type is "lxd".
	InstanceType string `json:"instance_type"`
	// Configuration file to be used by Cloud Init. Returned if the type is "lxd".
	Userdata string `json:"userdata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DNS          respjson.Field
		Dns4         respjson.Field
		Dns6         respjson.Field
		InstanceType respjson.Field
		Userdata     respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeInstanceMetadata) RawJSON() string { return r.JSON.raw }
func (r *ComputeInstanceMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeInstanceResponse struct {
	Content ComputeInstance `json:"content"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ComputeInstanceResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeInstanceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Metadata, Specs are required.
type ComputeInstanceUpdateParam struct {
	// Optional. The metadata required to configure in an Compute Instance.
	Metadata ComputeInstanceUpdateMetadataParam `json:"metadata,omitzero" api:"required"`
	// List of specs (SKUs) for the Compute Instance resource.
	Specs []ComputeInstanceUpdateSpecParam `json:"specs,omitzero" api:"required"`
	// The number of days after a Compute Intsance is closed before it is permanently
	// deleted.
	GracePeriod param.Opt[int64] `json:"grace_period,omitzero"`
	// The user-friendly name for the Compute Intsance type. If not sent and the type
	// is "lxd", it will default to the name 'LXD'. If not sent and the type is
	// "hyperv", it will default to the name 'VM HyperV'.
	Name param.Opt[string] `json:"name,omitzero"`
	// Change the state of the Compute Instance, triggering the CloudCIX Robot to
	// perform the requested action. Users can only request state changes from certain
	// current states, with specific allowed target states:
	//
	// - running -> stop, delete, or update_running
	// - stopped -> restart, delete, or update_stopped
	// - delete_queue -> restart or stop
	State param.Opt[string] `json:"state,omitzero"`
	// Optional. A list of network interfaces that represent the interfaces that will
	// be configured on the LXD instance.
	Interfaces []ComputeInstanceUpdateInterfaceParam `json:"interfaces,omitzero"`
	paramObj
}

func (r ComputeInstanceUpdateParam) MarshalJSON() (data []byte, err error) {
	type shadow ComputeInstanceUpdateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeInstanceUpdateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional. The metadata required to configure in an Compute Instance.
type ComputeInstanceUpdateMetadataParam struct {
	// Optional. A string containing IP Addresses, separated by commas, that represent
	// the DNS servers that the Compute Instance will use.
	DNS param.Opt[string] `json:"dns,omitzero"`
	// Optional, The Compute Instance instance type of the VM. Valid options are
	// "container" or "virtual-machine". If not sent it will default to "container".
	InstanceType param.Opt[string] `json:"instance_type,omitzero"`
	// Cloud Init allows Mime Multi-part messages, or files that start with a given set
	// of strings. It is a requirement to configure at minimum one user with a password
	// or ssh key that is in the sudo group.
	//
	// Reference: https://cloudinit.readthedocs.io/en/latest/explanation/format.html
	//
	// A hashed password can be generated using `openssl passwd -6 'yourpassword'`
	//
	// Example:
	//
	// ```yaml
	// #cloud-config
	// users:
	//   - name: administrator
	//     groups: sudo
	//     shell: /bin/bash
	//     lock_passwd: false
	//     passwd: < HASHED PASWWORD >
	//     ssh_authorized_keys:
	//   - < HASHED PASWWORD >
	//
	// chpasswd:
	//
	//	expire: false
	//
	// ssh_pwauth: true
	// ```
	Userdata param.Opt[string] `json:"userdata,omitzero"`
	paramObj
}

func (r ComputeInstanceUpdateMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow ComputeInstanceUpdateMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeInstanceUpdateMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeInstanceUpdateSpecParam struct {
	// The quantity of the SKU to configure the Compute Instance instance with.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// The name of the SKU to configure on the Compute Instance instacne
	SKUName param.Opt[string] `json:"sku_name,omitzero"`
	paramObj
}

func (r ComputeInstanceUpdateSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow ComputeInstanceUpdateSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeInstanceUpdateSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeInstanceUpdateInterfaceParam struct {
	// Flag representing if this interface will be the Gateway Interface to the Public
	// Internet.
	Gateway param.Opt[bool] `json:"gateway,omitzero"`
	// A list of IPv4 address objects to be assigned to this interface. All addresses
	// in this list must be from the same network.
	Ipv4Addresses []ComputeInstanceUpdateInterfaceIpv4AddressParam `json:"ipv4_addresses,omitzero"`
	// A list of IPv6 address objects to be assigned to this interface. All addresses
	// in this list must be from the same network as the `ipv4_addresses`.
	Ipv6Addresses []ComputeInstanceUpdateInterfaceIpv6AddressParam `json:"ipv6_addresses,omitzero"`
	paramObj
}

func (r ComputeInstanceUpdateInterfaceParam) MarshalJSON() (data []byte, err error) {
	type shadow ComputeInstanceUpdateInterfaceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeInstanceUpdateInterfaceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeInstanceUpdateInterfaceIpv4AddressParam struct {
	// An RFC1918 IPv4 address to be configured on this interface on the Compute
	// Instance instance.
	Address param.Opt[string] `json:"address,omitzero"`
	// Optional, Flag indicating if this address should be NATted to a Public IP
	// Address. If not sent, it will default to False.
	Nat param.Opt[bool] `json:"nat,omitzero"`
	paramObj
}

func (r ComputeInstanceUpdateInterfaceIpv4AddressParam) MarshalJSON() (data []byte, err error) {
	type shadow ComputeInstanceUpdateInterfaceIpv4AddressParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeInstanceUpdateInterfaceIpv4AddressParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeInstanceUpdateInterfaceIpv6AddressParam struct {
	// An IPv6 address to be configured on this interface on the Compute Instance
	// instance.
	Address param.Opt[string] `json:"address,omitzero"`
	paramObj
}

func (r ComputeInstanceUpdateInterfaceIpv6AddressParam) MarshalJSON() (data []byte, err error) {
	type shadow ComputeInstanceUpdateInterfaceIpv6AddressParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeInstanceUpdateInterfaceIpv6AddressParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeInstanceListResponse struct {
	Metadata ListMetadata      `json:"_metadata"`
	Content  []ComputeInstance `json:"content"`
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
func (r ComputeInstanceListResponse) RawJSON() string { return r.JSON.raw }
func (r *ComputeInstanceListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeInstanceNewParams struct {
	// Optional. The metadata required to configure in an Compute Instance.
	Metadata ComputeInstanceNewParamsMetadata `json:"metadata,omitzero" api:"required"`
	// The ID of the Project which this Compute Intsance Resource should be in.
	ProjectID int64 `json:"project_id" api:"required"`
	// List of specs (SKUs) for the Compute Instance resource.
	Specs []ComputeInstanceNewParamsSpec `json:"specs,omitzero" api:"required"`
	// The number of days after a Compute Intsance is closed before it is permanently
	// deleted.
	GracePeriod param.Opt[int64] `json:"grace_period,omitzero"`
	// The user-friendly name for the Compute Intsance type. If not sent and the type
	// is "lxd", it will default to the name 'LXD'. If not sent and the type is
	// "hyperv", it will default to the name 'HyperV'.
	Name param.Opt[string] `json:"name,omitzero"`
	// The type of Compute Instance to create. Valid options are:
	//
	// - "hyperv"
	// - "lxd"
	Type param.Opt[string] `json:"type,omitzero"`
	// Optional. A list of network interfaces that represent the interfaces that will
	// be configured on the LXD instance.
	Interfaces []ComputeInstanceNewParamsInterface `json:"interfaces,omitzero"`
	paramObj
}

func (r ComputeInstanceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ComputeInstanceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeInstanceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Optional. The metadata required to configure in an Compute Instance.
type ComputeInstanceNewParamsMetadata struct {
	// Optional. A string containing IP Addresses, separated by commas, that represent
	// the DNS servers that the Compute Instance will use.
	DNS param.Opt[string] `json:"dns,omitzero"`
	// Optional, The Compute Instance instance type of the VM. Valid options are
	// "container" or "virtual-machine". If not sent it will default to "container".
	InstanceType param.Opt[string] `json:"instance_type,omitzero"`
	// Cloud Init allows Mime Multi-part messages, or files that start with a given set
	// of strings. It is a requirement to configure at minimum one user with a password
	// or ssh key that is in the sudo group.
	//
	// Reference: https://cloudinit.readthedocs.io/en/latest/explanation/format.html
	//
	// A hashed password can be generated using `openssl passwd -6 'yourpassword'`
	//
	// Example:
	//
	// ```yaml
	// #cloud-config
	// users:
	//   - name: administrator
	//     groups: sudo
	//     shell: /bin/bash
	//     lock_passwd: false
	//     passwd: < HASHED PASWWORD >
	//     ssh_authorized_keys:
	//   - < HASHED PASWWORD >
	//
	// chpasswd:
	//
	//	expire: false
	//
	// ssh_pwauth: true
	// ```
	Userdata param.Opt[string] `json:"userdata,omitzero"`
	// Optional. A list of SSH key names (previously registered via POST
	// /compute/ssh_keys/) to inject into the instance during provisioning. The public
	// keys are appended to `ssh_authorized_keys` for the first sudo user defined in
	// the cloud-init userdata. If no sudo user exists, an `administrator` user is
	// created automatically. Only supported for LXD instances with #cloud-config
	// userdata.
	SSHKeyNames []string `json:"ssh_key_names,omitzero"`
	paramObj
}

func (r ComputeInstanceNewParamsMetadata) MarshalJSON() (data []byte, err error) {
	type shadow ComputeInstanceNewParamsMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeInstanceNewParamsMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeInstanceNewParamsSpec struct {
	// The quantity of the SKU to configure the Compute Instance instance with.
	Quantity param.Opt[int64] `json:"quantity,omitzero"`
	// The name of the SKU to configure on the Compute Instance instacne
	SKUName param.Opt[string] `json:"sku_name,omitzero"`
	paramObj
}

func (r ComputeInstanceNewParamsSpec) MarshalJSON() (data []byte, err error) {
	type shadow ComputeInstanceNewParamsSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeInstanceNewParamsSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeInstanceNewParamsInterface struct {
	// Flag representing if this interface will be the Gateway Interface to the Public
	// Internet.
	Gateway param.Opt[bool] `json:"gateway,omitzero"`
	// A list of IPv4 address objects to be assigned to this interface. All addresses
	// in this list must be from the same network.
	Ipv4Addresses []ComputeInstanceNewParamsInterfaceIpv4Address `json:"ipv4_addresses,omitzero"`
	// A list of IPv6 address objects to be assigned to this interface. All addresses
	// in this list must be from the same network as the `ipv4_addresses`.
	Ipv6Addresses []ComputeInstanceNewParamsInterfaceIpv6Address `json:"ipv6_addresses,omitzero"`
	paramObj
}

func (r ComputeInstanceNewParamsInterface) MarshalJSON() (data []byte, err error) {
	type shadow ComputeInstanceNewParamsInterface
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeInstanceNewParamsInterface) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeInstanceNewParamsInterfaceIpv4Address struct {
	// An RFC1918 IPv4 address to be configured on this interface on the Compute
	// Instance instance.
	Address param.Opt[string] `json:"address,omitzero"`
	// Optional, Flag indicating if this address should be NATted to a Public IP
	// Address. If not sent, it will default to False.
	Nat param.Opt[bool] `json:"nat,omitzero"`
	paramObj
}

func (r ComputeInstanceNewParamsInterfaceIpv4Address) MarshalJSON() (data []byte, err error) {
	type shadow ComputeInstanceNewParamsInterfaceIpv4Address
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeInstanceNewParamsInterfaceIpv4Address) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeInstanceNewParamsInterfaceIpv6Address struct {
	// An IPv6 address to be configured on this interface on the Compute Instance
	// instance.
	Address param.Opt[string] `json:"address,omitzero"`
	paramObj
}

func (r ComputeInstanceNewParamsInterfaceIpv6Address) MarshalJSON() (data []byte, err error) {
	type shadow ComputeInstanceNewParamsInterfaceIpv6Address
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ComputeInstanceNewParamsInterfaceIpv6Address) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeInstanceUpdateParams struct {
	ComputeInstanceUpdate ComputeInstanceUpdateParam
	paramObj
}

func (r ComputeInstanceUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ComputeInstanceUpdate)
}
func (r *ComputeInstanceUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ComputeInstanceListParams struct {
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

// URLQuery serializes [ComputeInstanceListParams]'s query parameters as
// `url.Values`.
func (r ComputeInstanceListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
