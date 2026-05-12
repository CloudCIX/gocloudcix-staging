// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package gocloudcix

import (
	"github.com/CloudCIX/gocloudcix/option"
)

// StorageService contains methods and other services that help with interacting
// with the gocloudcix API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewStorageService] method instead.
type StorageService struct {
	Options []option.RequestOption
	// Management of Storage Volumes
	//
	// Storage volumes provide additional storage capacity for compute instances.
	//
	// Three types are supported:
	//
	//   - CephFS: Network-attached file system volumes that can be mounted to multiple
	//     LXD instances
	//   - CephRBD: Block Storage volume that can be mounted to a virtual-machine LXD
	//     instance.
	//   - HyperV: Secondary drives attached to Hyper-V instances
	//
	// SKU Configuration: Storage capacity is specified using SKUs (Stock Keeping
	// Units) with quantity in GB.
	//
	//   - Ceph volumes use Ceph storage SKUs (CEPH_001 for HDD, CEPH_002 for SSD)
	//   - HyperV volumes use storage SKUs (e.g., SSD_001, HDD_001) Available SKUs depend
	//     on your region's configured devices.
	Volumes StorageVolumeService
}

// NewStorageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewStorageService(opts ...option.RequestOption) (r StorageService) {
	r = StorageService{}
	r.Options = opts
	r.Volumes = NewStorageVolumeService(opts...)
	return
}
