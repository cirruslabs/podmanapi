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

type InlineResponse2007 struct {
	// Automated indicates if the image was created by an automated build.
	Automated string `json:"Automated,omitempty"`
	// Description of the image.
	Description string `json:"Description,omitempty"`
	// Index is the image index (e.g., \"docker.io\" or \"quay.io\")
	Index string `json:"Index,omitempty"`
	// Name is the canonical name of the image (e.g., \"docker.io/library/alpine\").
	Name string `json:"Name,omitempty"`
	// Official indicates if it's an official image.
	Official string `json:"Official,omitempty"`
	// Stars is the number of stars of the image.
	Stars int64 `json:"Stars,omitempty"`
	// Tag is the image tag
	Tag string `json:"Tag,omitempty"`
}
