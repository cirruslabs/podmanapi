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
import (
	"time"
)

// InspectContainerState provides a detailed record of a container's current state. It is returned as part of InspectContainerData. As with InspectContainerData, many portions of this struct are matched to Docker, but here we see more fields that are unused (nonsensical in the context of Libpod).
type InspectContainerState struct {
	ConmonPid int64 `json:"ConmonPid,omitempty"`
	Dead bool `json:"Dead,omitempty"`
	Error_ string `json:"Error,omitempty"`
	ExitCode int32 `json:"ExitCode,omitempty"`
	FinishedAt time.Time `json:"FinishedAt,omitempty"`
	Healthcheck *HealthCheckResults `json:"Healthcheck,omitempty"`
	OOMKilled bool `json:"OOMKilled,omitempty"`
	OciVersion string `json:"OciVersion,omitempty"`
	Paused bool `json:"Paused,omitempty"`
	Pid int64 `json:"Pid,omitempty"`
	Restarting bool `json:"Restarting,omitempty"`
	Running bool `json:"Running,omitempty"`
	StartedAt time.Time `json:"StartedAt,omitempty"`
	Status string `json:"Status,omitempty"`
}
