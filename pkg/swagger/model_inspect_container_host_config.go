/*
 * Provides an API for the  Libpod library
 *
 * This documentation describes the Podman v2.0 RESTful API. It replaces the Podman v1.0 API and was initially delivered along with Podman v2.0.  It consists of a Docker-compatible API and a Libpod API providing support for Podman’s unique features such as pods.  To start the service and keep it running for 5,000 seconds (-t 0 runs forever):  podman system service -t 5000 &  You can then use cURL on the socket using requests documented below.  NOTE: if you install the package podman-docker, it will create a symbolic link for /run/docker.sock to /run/podman/podman.sock  See podman-service(1) for more information.  Quick Examples:  'podman info'  curl --unix-socket /run/podman/podman.sock http://d/v3.0.0/libpod/info  'podman pull quay.io/containers/podman'  curl -XPOST --unix-socket /run/podman/podman.sock -v 'http://d/v3.0.0/images/create?fromImage=quay.io%2Fcontainers%2Fpodman'  'podman list images'  curl --unix-socket /run/podman/podman.sock -v 'http://d/v3.0.0/libpod/images/json' | jq
 *
 * API version: 3.2.0
 * Contact: podman@lists.podman.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// InspectContainerHostConfig holds information used when the container was created. It's very much a Docker-specific struct, retained (mostly) as-is for compatibility. We fill individual fields as best as we can, inferring as much as possible from the spec and container config. Some things cannot be inferred. These will be populated by spec annotations (if available). Field names are fixed for compatibility and cannot be changed. As such, silence lint warnings about them. nolint
type InspectContainerHostConfig struct {
	// AutoRemove is whether the container will be automatically removed on exiting. It is not handled directly within libpod and is stored in an annotation.
	AutoRemove bool `json:"AutoRemove,omitempty"`
	// Binds contains an array of user-added mounts. Both volume mounts and named volumes are included. Tmpfs mounts are NOT included. In 'docker inspect' this is separated into 'Binds' and 'Mounts' based on how a mount was added. We do not make this distinction and do not include a Mounts field in inspect. Format: <src>:<destination>[:<comma-separated options>]
	Binds []string `json:"Binds,omitempty"`
	// BlkioDeviceReadBps is an array of I/O throttle parameters for individual device nodes. This specifically sets read rate cap in bytes per second for device nodes. As with BlkioWeightDevice, we pull the path from /sys/dev, and we don't guarantee the path will be identical to the original (though the node will be).
	BlkioDeviceReadBps []InspectBlkioThrottleDevice `json:"BlkioDeviceReadBps,omitempty"`
	// BlkioDeviceReadIOps is an array of I/O throttle parameters for individual device nodes. This specifically sets the read rate cap in iops per second for device nodes. As with BlkioWeightDevice, we pull the path from /sys/dev, and we don't guarantee the path will be identical to the original (though the node will be).
	BlkioDeviceReadIOps []InspectBlkioThrottleDevice `json:"BlkioDeviceReadIOps,omitempty"`
	// BlkioDeviceWriteBps is an array of I/O throttle parameters for individual device nodes. this specifically sets write rate cap in bytes per second for device nodes. as with BlkioWeightDevice, we pull the path from /sys/dev, and we don't guarantee the path will be identical to the original (though the node will be).
	BlkioDeviceWriteBps []InspectBlkioThrottleDevice `json:"BlkioDeviceWriteBps,omitempty"`
	// BlkioDeviceWriteIOps is an array of I/O throttle parameters for individual device nodes. This specifically sets the write rate cap in iops per second for device nodes. As with BlkioWeightDevice, we pull the path from /sys/dev, and we don't guarantee the path will be identical to the original (though the node will be).
	BlkioDeviceWriteIOps []InspectBlkioThrottleDevice `json:"BlkioDeviceWriteIOps,omitempty"`
	// BlkioWeight indicates the I/O resources allocated to the container. It is a relative weight in the scheduler for assigning I/O time versus other CGroups.
	BlkioWeight int32 `json:"BlkioWeight,omitempty"`
	// BlkioWeightDevice is an array of I/O resource priorities for individual device nodes. Unfortunately, the spec only stores the device's Major/Minor numbers and not the path, which is used here. Fortunately, the kernel provides an interface for retrieving the path of a given node by major:minor at /sys/dev/. However, the exact path in use may not be what was used in the original CLI invocation - though it is guaranteed that the device node will be the same, and using the given path will be functionally identical.
	BlkioWeightDevice []InspectBlkioWeightDevice `json:"BlkioWeightDevice,omitempty"`
	// CapAdd is a list of capabilities added to the container. It is not directly stored by Libpod, and instead computed from the capabilities listed in the container's spec, compared against a set of default capabilities.
	CapAdd []string `json:"CapAdd,omitempty"`
	// CapDrop is a list of capabilities removed from the container. It is not directly stored by libpod, and instead computed from the capabilities listed in the container's spec, compared against a set of default capabilities.
	CapDrop []string `json:"CapDrop,omitempty"`
	// Cgroup contains the container's cgroup. It is presently not populated. TODO.
	Cgroup string `json:"Cgroup,omitempty"`
	// CgroupConf is the configuration for cgroup v2.
	CgroupConf map[string]string `json:"CgroupConf,omitempty"`
	// CgroupManager is the cgroup manager used by the container. At present, allowed values are either \"cgroupfs\" or \"systemd\".
	CgroupManager string `json:"CgroupManager,omitempty"`
	// CgroupMode is the configuration of the container's cgroup namespace. Populated as follows: private - a cgroup namespace has been created host - No cgroup namespace created container:<id> - Using another container's cgroup namespace ns:<path> - A path to a cgroup namespace has been specified
	CgroupMode string `json:"CgroupMode,omitempty"`
	// CgroupParent is the CGroup parent of the container. Only set if not default.
	CgroupParent string `json:"CgroupParent,omitempty"`
	// Cgroups contains the container's CGroup mode. Allowed values are \"default\" (container is creating CGroups) and \"disabled\" (container is not creating CGroups). This is Libpod-specific and not included in `docker inspect`.
	Cgroups string `json:"Cgroups,omitempty"`
	// ConsoleSize is an array of 2 integers showing the size of the container's console. It is only set if the container is creating a terminal. TODO.
	ConsoleSize []int32 `json:"ConsoleSize,omitempty"`
	// ContainerIDFile is a file created during container creation to hold the ID of the created container. This is not handled within libpod and is stored in an annotation.
	ContainerIDFile string `json:"ContainerIDFile,omitempty"`
	// CpuCount is Windows-only and not presently implemented.
	CpuCount int32 `json:"CpuCount,omitempty"`
	// CpuPercent is Windows-only and not presently implemented.
	CpuPercent int32 `json:"CpuPercent,omitempty"`
	// CpuPeriod is the length of a CPU period in microseconds. It relates directly to CpuQuota.
	CpuPeriod int32 `json:"CpuPeriod,omitempty"`
	// CpuPeriod is the amount of time (in microseconds) that a container can use the CPU in every CpuPeriod.
	CpuQuota int64 `json:"CpuQuota,omitempty"`
	// CpuRealtimePeriod is the length of time (in microseconds) of the CPU realtime period. If set to 0, no time will be allocated to realtime tasks.
	CpuRealtimePeriod int32 `json:"CpuRealtimePeriod,omitempty"`
	// CpuRealtimeRuntime is the length of time (in microseconds) allocated for realtime tasks within every CpuRealtimePeriod.
	CpuRealtimeRuntime int64 `json:"CpuRealtimeRuntime,omitempty"`
	// CpuShares indicates the CPU resources allocated to the container. It is a relative weight in the scheduler for assigning CPU time versus other CGroups.
	CpuShares int32 `json:"CpuShares,omitempty"`
	// CpusetCpus is the is the set of CPUs that the container will execute on. Formatted as `0-3` or `0,2`. Default (if unset) is all CPUs.
	CpusetCpus string `json:"CpusetCpus,omitempty"`
	// CpusetMems is the set of memory nodes the container will use. Formatted as `0-3` or `0,2`. Default (if unset) is all memory nodes.
	CpusetMems string `json:"CpusetMems,omitempty"`
	// Devices is a list of device nodes that will be added to the container. These are stored in the OCI spec only as type, major, minor while we display the host path. We convert this with /sys/dev, but we cannot guarantee that the host path will be identical - only that the actual device will be.
	Devices []InspectDevice `json:"Devices,omitempty"`
	// DiskQuota is the maximum amount of disk space the container may use (in bytes). Presently not populated. TODO.
	DiskQuota int32 `json:"DiskQuota,omitempty"`
	// Dns is a list of DNS nameservers that will be added to the container's resolv.conf
	Dns []string `json:"Dns,omitempty"`
	// DnsOptions is a list of DNS options that will be set in the container's resolv.conf
	DnsOptions []string `json:"DnsOptions,omitempty"`
	// DnsSearch is a list of DNS search domains that will be set in the container's resolv.conf
	DnsSearch []string `json:"DnsSearch,omitempty"`
	// ExtraHosts contains hosts that will be aded to the container's etc/hosts.
	ExtraHosts []string `json:"ExtraHosts,omitempty"`
	// GroupAdd contains groups that the user inside the container will be added to.
	GroupAdd []string `json:"GroupAdd,omitempty"`
	// IOMaximumBandwidth is Windows-only and not presently implemented.
	IOMaximumBandwidth int32 `json:"IOMaximumBandwidth,omitempty"`
	// IOMaximumIOps is Windows-only and not presently implemented.
	IOMaximumIOps int32 `json:"IOMaximumIOps,omitempty"`
	// Init indicates whether the container has an init mounted into it.
	Init bool `json:"Init,omitempty"`
	// IpcMode represents the configuration of the container's IPC namespace. Populated as follows: \"\" (empty string) - Default, an IPC namespace will be created host - No IPC namespace created container:<id> - Using another container's IPC namespace ns:<path> - A path to an IPC namespace has been specified
	IpcMode string `json:"IpcMode,omitempty"`
	// Isolation is presently unused and provided solely for Docker compatibility.
	Isolation string `json:"Isolation,omitempty"`
	// KernelMemory is the maximum amount of memory the kernel will devote to the container.
	KernelMemory int64 `json:"KernelMemory,omitempty"`
	// Links is unused, and provided purely for Docker compatibility.
	Links []string `json:"Links,omitempty"`
	LogConfig *InspectLogConfig `json:"LogConfig,omitempty"`
	// Memory indicates the memory resources allocated to the container. This is the limit (in bytes) of RAM the container may use.
	Memory int64 `json:"Memory,omitempty"`
	// MemoryReservation is the reservation (soft limit) of memory available to the container. Soft limits are warnings only and can be exceeded.
	MemoryReservation int64 `json:"MemoryReservation,omitempty"`
	// MemorySwap is the total limit for all memory available to the container, including swap. 0 indicates that there is no limit to the amount of memory available.
	MemorySwap int64 `json:"MemorySwap,omitempty"`
	// MemorySwappiness is the willingness of the kernel to page container memory to swap. It is an integer from 0 to 100, with low numbers being more likely to be put into swap. 1, the default, will not set swappiness and use the system defaults.
	MemorySwappiness int64 `json:"MemorySwappiness,omitempty"`
	// NanoCpus indicates number of CPUs allocated to the container. It is an integer where one full CPU is indicated by 1000000000 (one billion). Thus, 2.5 CPUs (fractional portions of CPUs are allowed) would be 2500000000 (2.5 billion). In 'docker inspect' this is set exclusively of two further options in the output (CpuPeriod and CpuQuota) which are both used to implement this functionality. We can't distinguish here, so if CpuQuota is set to the default of 100000, we will set both CpuQuota, CpuPeriod, and NanoCpus. If CpuQuota is not the default, we will not set NanoCpus.
	NanoCpus int64 `json:"NanoCpus,omitempty"`
	// NetworkMode is the configuration of the container's network namespace. Populated as follows: default - A network namespace is being created and configured via CNI none - A network namespace is being created, not configured via CNI host - No network namespace created container:<id> - Using another container's network namespace ns:<path> - A path to a network namespace has been specified
	NetworkMode string `json:"NetworkMode,omitempty"`
	// OomKillDisable indicates whether the kernel OOM killer is disabled for the container.
	OomKillDisable bool `json:"OomKillDisable,omitempty"`
	// OOMScoreAdj is an adjustment that will be made to the container's OOM score.
	OomScoreAdj int64 `json:"OomScoreAdj,omitempty"`
	// PidMode represents the configuration of the container's PID namespace. Populated as follows: \"\" (empty string) - Default, a PID namespace will be created host - No PID namespace created container:<id> - Using another container's PID namespace ns:<path> - A path to a PID namespace has been specified
	PidMode string `json:"PidMode,omitempty"`
	// PidsLimit is the maximum number of PIDs what may be created within the container. 0, the default, indicates no limit.
	PidsLimit int64 `json:"PidsLimit,omitempty"`
	// PortBindings contains the container's port bindings. It is formatted as map[string][]InspectHostPort. The string key here is formatted as <integer port number>/<protocol> and represents the container port. A single container port may be bound to multiple host ports (on different IPs).
	PortBindings map[string][]InspectHostPort `json:"PortBindings,omitempty"`
	// Privileged indicates whether the container is running with elevated privileges. This has a very specific meaning in the Docker sense, so it's very difficult to decode from the spec and config, and so is stored as an annotation.
	Privileged bool `json:"Privileged,omitempty"`
	// PublishAllPorts indicates whether image ports are being published. This is not directly stored in libpod and is saved as an annotation.
	PublishAllPorts bool `json:"PublishAllPorts,omitempty"`
	// ReadonlyRootfs is whether the container will be mounted read-only.
	ReadonlyRootfs bool `json:"ReadonlyRootfs,omitempty"`
	RestartPolicy *InspectRestartPolicy `json:"RestartPolicy,omitempty"`
	// Runtime is provided purely for Docker compatibility. It is set unconditionally to \"oci\" as Podman does not presently support non-OCI runtimes.
	Runtime string `json:"Runtime,omitempty"`
	// SecurityOpt is a list of security-related options that are set in the container.
	SecurityOpt []string `json:"SecurityOpt,omitempty"`
	// ShmSize is the size of the container's SHM device.
	ShmSize int64 `json:"ShmSize,omitempty"`
	// Tmpfs is a list of tmpfs filesystems that will be mounted into the container. It is a map of destination path to options for the mount.
	Tmpfs map[string]string `json:"Tmpfs,omitempty"`
	// UTSMode represents the configuration of the container's UID namespace. Populated as follows: \"\" (empty string) - Default, a UTS namespace will be created host - no UTS namespace created container:<id> - Using another container's UTS namespace ns:<path> - A path to a UTS namespace has been specified
	UTSMode string `json:"UTSMode,omitempty"`
	// Ulimits is a set of ulimits that will be set within the container.
	Ulimits []InspectUlimit `json:"Ulimits,omitempty"`
	// UsernsMode represents the configuration of the container's user namespace. When running rootless, a user namespace is created outside of libpod to allow some privileged operations. This will not be reflected here. Populated as follows: \"\" (empty string) - No user namespace will be created private - The container will be run in a user namespace container:<id> - Using another container's user namespace ns:<path> - A path to a user namespace has been specified TODO Rootless has an additional 'keep-id' option, presently not reflected here.
	UsernsMode string `json:"UsernsMode,omitempty"`
	// VolumeDriver is presently unused and is retained for Docker compatibility.
	VolumeDriver string `json:"VolumeDriver,omitempty"`
	// VolumesFrom is a list of containers which this container uses volumes from. This is not handled directly within libpod and is stored in an annotation. It is formatted as an array of container names and IDs.
	VolumesFrom []string `json:"VolumesFrom,omitempty"`
}
