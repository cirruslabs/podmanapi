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

// NetworkCreate is the expected body of the \"create network\" http request message
type NetworkCreate struct {
	Attachable bool `json:"Attachable,omitempty"`
	// Check for networks with duplicate names. Network is primarily keyed based on a random ID and not on the name. Network name is strictly a user-friendly alias to the network which is uniquely identified using ID. And there is no guaranteed way to check for duplicates. Option CheckDuplicate is there to provide a best effort checking of any networks which has the same name but it is not guaranteed to catch all name collisions.
	CheckDuplicate bool `json:"CheckDuplicate,omitempty"`
	ConfigFrom *ConfigReference `json:"ConfigFrom,omitempty"`
	ConfigOnly bool `json:"ConfigOnly,omitempty"`
	Driver string `json:"Driver,omitempty"`
	EnableIPv6 bool `json:"EnableIPv6,omitempty"`
	IPAM *Ipam `json:"IPAM,omitempty"`
	Ingress bool `json:"Ingress,omitempty"`
	Internal bool `json:"Internal,omitempty"`
	Labels map[string]string `json:"Labels,omitempty"`
	Options map[string]string `json:"Options,omitempty"`
	Scope string `json:"Scope,omitempty"`
}
