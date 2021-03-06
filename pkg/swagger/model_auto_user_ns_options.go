/*
 * Provides a container compatible interface.
 *
 * This documentation describes the Podman v2.0 RESTful API. It replaces the Podman v1.0 API and was initially delivered along with Podman v2.0.  It consists of a Docker-compatible API and a Libpod API providing support for Podman’s unique features such as pods.  To start the service and keep it running for 5,000 seconds (-t 0 runs forever):  podman system service -t 5000 &  You can then use cURL on the socket using requests documented below.  NOTE: if you install the package podman-docker, it will create a symbolic link for /var/run/docker.sock to /run/podman/podman.sock  See podman-service(1) for more information.  Quick Examples:  'podman info'  curl --unix-socket /run/podman/podman.sock http://d/v1.0.0/libpod/info  'podman pull quay.io/containers/podman'  curl -XPOST --unix-socket /run/podman/podman.sock -v 'http://d/v1.0.0/images/create?fromImage=quay.io%2Fcontainers%2Fpodman'  'podman list images'  curl --unix-socket /run/podman/podman.sock -v 'http://d/v1.0.0/libpod/images/json' | jq
 *
 * API version: 0.0.1
 * Contact: podman@lists.podman.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type AutoUserNsOptions struct {
	// AdditionalGIDMappings specified additional GID mappings to include in the generated user namespace.
	AdditionalGIDMappings []IdMap `json:"AdditionalGIDMappings,omitempty"`
	// AdditionalUIDMappings specified additional UID mappings to include in the generated user namespace.
	AdditionalUIDMappings []IdMap `json:"AdditionalUIDMappings,omitempty"`
	// GroupFile to use if the container uses a volume.
	GroupFile string `json:"GroupFile,omitempty"`
	// InitialSize defines the minimum size for the user namespace. The created user namespace will have at least this size.
	InitialSize int32 `json:"InitialSize,omitempty"`
	// PasswdFile to use if the container uses a volume.
	PasswdFile string `json:"PasswdFile,omitempty"`
	// Size defines the size for the user namespace.  If it is set to a value bigger than 0, the user namespace will have exactly this size. If it is not set, some heuristics will be used to find its size.
	Size int32 `json:"Size,omitempty"`
}
