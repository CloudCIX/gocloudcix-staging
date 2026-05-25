# Compute

## Backups

Params Types:

- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeBackupUpdateParam">ComputeBackupUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeBackup">ComputeBackup</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeBackupResponse">ComputeBackupResponse</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ListMetadata">ListMetadata</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeBackupListResponse">ComputeBackupListResponse</a>

Methods:

- <code title="post /compute/backups/">client.Compute.Backups.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeBackupService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeBackupNewParams">ComputeBackupNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeBackup">ComputeBackup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/backups/{id}/">client.Compute.Backups.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeBackupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeBackup">ComputeBackup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /compute/backups/{id}/">client.Compute.Backups.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeBackupService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeBackupUpdateParams">ComputeBackupUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeBackup">ComputeBackup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/backups/">client.Compute.Backups.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeBackupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeBackupListParams">ComputeBackupListParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeBackupListResponse">ComputeBackupListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /compute/backups/{id}/">client.Compute.Backups.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeBackupService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## GPUs

Params Types:

- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeGPUUpdateParam">ComputeGPUUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeGPU">ComputeGPU</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeGPUResponse">ComputeGPUResponse</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeGPUListResponse">ComputeGPUListResponse</a>

Methods:

- <code title="get /compute/gpus/{id}/">client.Compute.GPUs.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeGPUService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeGPU">ComputeGPU</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /compute/gpus/{id}/">client.Compute.GPUs.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeGPUService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeGPUUpdateParams">ComputeGPUUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeGPU">ComputeGPU</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/gpus/">client.Compute.GPUs.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeGPUService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeGPUListParams">ComputeGPUListParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeGPUListResponse">ComputeGPUListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /compute/gpus/{id}/">client.Compute.GPUs.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeGPUService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="post /compute/gpus/">client.Compute.GPUs.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeGPUService.Attach">Attach</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeGPUAttachParams">ComputeGPUAttachParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeGPU">ComputeGPU</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Images

Response Types:

- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeImage">ComputeImage</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeImageListResponse">ComputeImageListResponse</a>

Methods:

- <code title="get /compute/images/{id}/">client.Compute.Images.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeImageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeImage">ComputeImage</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/images/">client.Compute.Images.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeImageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeImageListParams">ComputeImageListParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeImageListResponse">ComputeImageListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Instances

Params Types:

- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeInstanceUpdateParam">ComputeInstanceUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#Bom">Bom</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeInstance">ComputeInstance</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeInstanceResponse">ComputeInstanceResponse</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeInstanceListResponse">ComputeInstanceListResponse</a>

Methods:

- <code title="post /compute/instances/">client.Compute.Instances.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeInstanceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeInstanceNewParams">ComputeInstanceNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeInstance">ComputeInstance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/instances/{id}/">client.Compute.Instances.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeInstanceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeInstance">ComputeInstance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /compute/instances/{id}/">client.Compute.Instances.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeInstanceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeInstanceUpdateParams">ComputeInstanceUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeInstance">ComputeInstance</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/instances/">client.Compute.Instances.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeInstanceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeInstanceListParams">ComputeInstanceListParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeInstanceListResponse">ComputeInstanceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /compute/instances/{id}/">client.Compute.Instances.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeInstanceService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Snapshots

Params Types:

- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeSnapshotUpdateParam">ComputeSnapshotUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeSnapshot">ComputeSnapshot</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeSnapshotResponse">ComputeSnapshotResponse</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeSnapshotListResponse">ComputeSnapshotListResponse</a>

Methods:

- <code title="post /compute/snapshots/">client.Compute.Snapshots.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeSnapshotService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeSnapshotNewParams">ComputeSnapshotNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeSnapshot">ComputeSnapshot</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/snapshots/{id}/">client.Compute.Snapshots.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeSnapshotService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeSnapshot">ComputeSnapshot</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /compute/snapshots/{id}/">client.Compute.Snapshots.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeSnapshotService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeSnapshotUpdateParams">ComputeSnapshotUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeSnapshot">ComputeSnapshot</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/snapshots/">client.Compute.Snapshots.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeSnapshotService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeSnapshotListParams">ComputeSnapshotListParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeSnapshotListResponse">ComputeSnapshotListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /compute/snapshots/{id}/">client.Compute.Snapshots.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ComputeSnapshotService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Network

## Firewalls

Params Types:

- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkFirewallUpdateParam">NetworkFirewallUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkFirewall">NetworkFirewall</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkFirewallResponse">NetworkFirewallResponse</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkFirewallListResponse">NetworkFirewallListResponse</a>

Methods:

- <code title="post /network/firewalls/">client.Network.Firewalls.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkFirewallService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkFirewallNewParams">NetworkFirewallNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkFirewall">NetworkFirewall</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/firewalls/{id}/">client.Network.Firewalls.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkFirewallService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkFirewall">NetworkFirewall</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /network/firewalls/{id}/">client.Network.Firewalls.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkFirewallService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkFirewallUpdateParams">NetworkFirewallUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkFirewall">NetworkFirewall</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/firewalls/">client.Network.Firewalls.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkFirewallService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkFirewallListParams">NetworkFirewallListParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkFirewallListResponse">NetworkFirewallListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /network/firewalls/{id}/">client.Network.Firewalls.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkFirewallService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## IPGroups

Params Types:

- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkIPGroupUpdateParam">NetworkIPGroupUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkIPGroup">NetworkIPGroup</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkIPGroupResponse">NetworkIPGroupResponse</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkIPGroupListResponse">NetworkIPGroupListResponse</a>

Methods:

- <code title="post /network/ip_groups/">client.Network.IPGroups.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkIPGroupService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkIPGroupNewParams">NetworkIPGroupNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkIPGroup">NetworkIPGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/ip_groups/{id}/">client.Network.IPGroups.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkIPGroupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkIPGroup">NetworkIPGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /network/ip_groups/{id}/">client.Network.IPGroups.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkIPGroupService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkIPGroupUpdateParams">NetworkIPGroupUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkIPGroup">NetworkIPGroup</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/ip_groups/">client.Network.IPGroups.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkIPGroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkIPGroupListParams">NetworkIPGroupListParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkIPGroupListResponse">NetworkIPGroupListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /network/ip_groups/{id}/">client.Network.IPGroups.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkIPGroupService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Routers

Params Types:

- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkRouterUpdateParam">NetworkRouterUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#BaseIPAddress">BaseIPAddress</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkRouter">NetworkRouter</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkRouterResponse">NetworkRouterResponse</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#RouterMetadata">RouterMetadata</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkRouterListResponse">NetworkRouterListResponse</a>

Methods:

- <code title="post /network/routers/">client.Network.Routers.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkRouterService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkRouterNewParams">NetworkRouterNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkRouter">NetworkRouter</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/routers/{id}/">client.Network.Routers.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkRouterService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkRouter">NetworkRouter</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /network/routers/{id}/">client.Network.Routers.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkRouterService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkRouterUpdateParams">NetworkRouterUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkRouter">NetworkRouter</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/routers/">client.Network.Routers.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkRouterService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkRouterListParams">NetworkRouterListParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkRouterListResponse">NetworkRouterListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /network/routers/{id}/">client.Network.Routers.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkRouterService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Vpns

Params Types:

- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkVpnUpdateParam">NetworkVpnUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkVpn">NetworkVpn</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkVpnResponse">NetworkVpnResponse</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkVpnListResponse">NetworkVpnListResponse</a>

Methods:

- <code title="post /network/vpns/">client.Network.Vpns.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkVpnService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkVpnNewParams">NetworkVpnNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkVpn">NetworkVpn</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/vpns/{id}/">client.Network.Vpns.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkVpnService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkVpn">NetworkVpn</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /network/vpns/{id}/">client.Network.Vpns.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkVpnService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkVpnUpdateParams">NetworkVpnUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkVpn">NetworkVpn</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/vpns/">client.Network.Vpns.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkVpnService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkVpnListParams">NetworkVpnListParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkVpnListResponse">NetworkVpnListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /network/vpns/{id}/">client.Network.Vpns.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#NetworkVpnService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Project

Params Types:

- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ProjectUpdateParam">ProjectUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#Project">Project</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ProjectResponse">ProjectResponse</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ProjectListResponse">ProjectListResponse</a>

Methods:

- <code title="post /project/">client.Project.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ProjectService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ProjectNewParams">ProjectNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /project/{id}/">client.Project.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ProjectService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /project/{id}/">client.Project.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ProjectService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ProjectUpdateParams">ProjectUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#Project">Project</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /project/">client.Project.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ProjectService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ProjectListParams">ProjectListParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ProjectListResponse">ProjectListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /project/{id}/">client.Project.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#ProjectService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Storage

## Volumes

Params Types:

- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#StorageVolumesUpdateParam">StorageVolumesUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#StorageVolumes">StorageVolumes</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#StorageVolumesResponse">StorageVolumesResponse</a>
- <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#StorageVolumeListResponse">StorageVolumeListResponse</a>

Methods:

- <code title="post /storage/volumes/">client.Storage.Volumes.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#StorageVolumeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#StorageVolumeNewParams">StorageVolumeNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#StorageVolumes">StorageVolumes</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /storage/volumes/{id}/">client.Storage.Volumes.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#StorageVolumeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#StorageVolumes">StorageVolumes</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /storage/volumes/{id}/">client.Storage.Volumes.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#StorageVolumeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#StorageVolumeUpdateParams">StorageVolumeUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#StorageVolumes">StorageVolumes</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /storage/volumes/">client.Storage.Volumes.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#StorageVolumeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#StorageVolumeListParams">StorageVolumeListParams</a>) (\*<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#StorageVolumeListResponse">StorageVolumeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /storage/volumes/{id}/">client.Storage.Volumes.<a href="https://pkg.go.dev/github.com/CloudCIX/gocloudcix#StorageVolumeService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
