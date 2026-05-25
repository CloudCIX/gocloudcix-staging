// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gocloudcix

import (
	"github.com/CloudCIX/gocloudcix/option"
)

// ComputeService contains methods and other services that help with interacting
// with the gocloudcix API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewComputeService] method instead.
type ComputeService struct {
	Options []option.RequestOption
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
	Backups ComputeBackupService
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
	GPUs ComputeGPUService
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
	Images ComputeImageService
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
	Instances ComputeInstanceService
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
	Snapshots ComputeSnapshotService
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
	SSHKeys ComputeSSHKeyService
}

// NewComputeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewComputeService(opts ...option.RequestOption) (r ComputeService) {
	r = ComputeService{}
	r.Options = opts
	r.Backups = NewComputeBackupService(opts...)
	r.GPUs = NewComputeGPUService(opts...)
	r.Images = NewComputeImageService(opts...)
	r.Instances = NewComputeInstanceService(opts...)
	r.Snapshots = NewComputeSnapshotService(opts...)
	r.SSHKeys = NewComputeSSHKeyService(opts...)
	return
}
