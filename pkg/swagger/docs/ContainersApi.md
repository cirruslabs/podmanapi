# {{classname}}

All URIs are relative to *http://podman.io/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContainerAttachLibpod**](ContainersApi.md#ContainerAttachLibpod) | **Post** /libpod/containers/{name}/attach | Attach to a container
[**ContainerChangesLibpod**](ContainersApi.md#ContainerChangesLibpod) | **Get** /libpod/containers/{name}/changes | Report on changes to container&#x27;s filesystem; adds, deletes or modifications.
[**ContainerCheckpointLibpod**](ContainersApi.md#ContainerCheckpointLibpod) | **Post** /libpod/containers/{name}/checkpoint | Checkpoint a container
[**ContainerCreateLibpod**](ContainersApi.md#ContainerCreateLibpod) | **Post** /libpod/containers/create | Create a container
[**ContainerDeleteLibpod**](ContainersApi.md#ContainerDeleteLibpod) | **Delete** /libpod/containers/{name} | Delete container
[**ContainerExistsLibpod**](ContainersApi.md#ContainerExistsLibpod) | **Get** /libpod/containers/{name}/exists | Check if container exists
[**ContainerExportLibpod**](ContainersApi.md#ContainerExportLibpod) | **Get** /libpod/containers/{name}/export | Export a container
[**ContainerHealthcheckLibpod**](ContainersApi.md#ContainerHealthcheckLibpod) | **Get** /libpod/containers/{name}/healthcheck | Run a container&#x27;s healthcheck
[**ContainerInitLibpod**](ContainersApi.md#ContainerInitLibpod) | **Post** /libpod/containers/{name}/init | Initialize a container
[**ContainerInspectLibpod**](ContainersApi.md#ContainerInspectLibpod) | **Get** /libpod/containers/{name}/json | Inspect container
[**ContainerKillLibpod**](ContainersApi.md#ContainerKillLibpod) | **Post** /libpod/containers/{name}/kill | Kill container
[**ContainerListLibpod**](ContainersApi.md#ContainerListLibpod) | **Get** /libpod/containers/json | List containers
[**ContainerLogsLibpod**](ContainersApi.md#ContainerLogsLibpod) | **Get** /libpod/containers/{name}/logs | Get container logs
[**ContainerMountLibpod**](ContainersApi.md#ContainerMountLibpod) | **Post** /libpod/containers/{name}/mount | Mount a container
[**ContainerPauseLibpod**](ContainersApi.md#ContainerPauseLibpod) | **Post** /libpod/containers/{name}/pause | Pause a container
[**ContainerPruneLibpod**](ContainersApi.md#ContainerPruneLibpod) | **Post** /libpod/containers/prune | Delete stopped containers
[**ContainerRenameLibpod**](ContainersApi.md#ContainerRenameLibpod) | **Post** /libpod/containers/{name}/rename | Rename an existing container
[**ContainerResizeLibpod**](ContainersApi.md#ContainerResizeLibpod) | **Post** /libpod/containers/{name}/resize | Resize a container&#x27;s TTY
[**ContainerRestartLibpod**](ContainersApi.md#ContainerRestartLibpod) | **Post** /libpod/containers/{name}/restart | Restart a container
[**ContainerRestoreLibpod**](ContainersApi.md#ContainerRestoreLibpod) | **Post** /libpod/containers/{name}/restore | Restore a container
[**ContainerShowMountedLibpod**](ContainersApi.md#ContainerShowMountedLibpod) | **Get** /libpod/containers/showmounted | Show mounted containers
[**ContainerStartLibpod**](ContainersApi.md#ContainerStartLibpod) | **Post** /libpod/containers/{name}/start | Start a container
[**ContainerStatsLibpod**](ContainersApi.md#ContainerStatsLibpod) | **Get** /libpod/containers/{name}/stats | Get stats for a container
[**ContainerStopLibpod**](ContainersApi.md#ContainerStopLibpod) | **Post** /libpod/containers/{name}/stop | Stop a container
[**ContainerTopLibpod**](ContainersApi.md#ContainerTopLibpod) | **Get** /libpod/containers/{name}/top | List processes
[**ContainerUnmountLibpod**](ContainersApi.md#ContainerUnmountLibpod) | **Post** /libpod/containers/{name}/unmount | Unmount a container
[**ContainerUnpauseLibpod**](ContainersApi.md#ContainerUnpauseLibpod) | **Post** /libpod/containers/{name}/unpause | Unpause Container
[**ContainerWaitLibpod**](ContainersApi.md#ContainerWaitLibpod) | **Post** /libpod/containers/{name}/wait | Wait on a container
[**ContainersStatsAllLibpod**](ContainersApi.md#ContainersStatsAllLibpod) | **Get** /libpod/containers/stats | Get stats for one or more containers
[**GenerateKubeLibpod**](ContainersApi.md#GenerateKubeLibpod) | **Get** /libpod/generate/kube | Generate a Kubernetes YAML file.
[**GenerateSystemdLibpod**](ContainersApi.md#GenerateSystemdLibpod) | **Get** /libpod/generate/{name}/systemd | Generate Systemd Units
[**ImageCommitLibpod**](ContainersApi.md#ImageCommitLibpod) | **Post** /libpod/commit | Commit
[**PlayKubeLibpod**](ContainersApi.md#PlayKubeLibpod) | **Post** /libpod/play/kube | Play a Kubernetes YAML file.
[**PutContainerArchiveLibpod**](ContainersApi.md#PutContainerArchiveLibpod) | **Put** /libpod/containers/{name}/archive | Copy files into a container

# **ContainerAttachLibpod**
> ContainerAttachLibpod(ctx, name, optional)
Attach to a container

Hijacks the connection to forward the container's standard streams to the client.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ContainersApiContainerAttachLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersApiContainerAttachLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **detachKeys** | **optional.String**| keys to use for detaching from the container | 
 **logs** | **optional.Bool**| Stream all logs from the container across the connection. Happens before streaming attach (if requested). At least one of logs or stream must be set | 
 **stream** | **optional.Bool**| Attach to the container. If unset, and logs is set, only the container&#x27;s logs will be sent. At least one of stream or logs must be set | [default to true]
 **stdout** | **optional.Bool**| Attach to container STDOUT | 
 **stderr** | **optional.Bool**| Attach to container STDERR | 
 **stdin** | **optional.Bool**| Attach to container STDIN | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerChangesLibpod**
> ContainerChangesLibpod(ctx, name)
Report on changes to container's filesystem; adds, deletes or modifications.

Returns which files in a container's filesystem have been added, deleted, or modified. The Kind of modification can be one of:  0: Modified 1: Added 2: Deleted 

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or id of the container | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, application/octet-stream, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerCheckpointLibpod**
> ContainerCheckpointLibpod(ctx, name, optional)
Checkpoint a container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ContainersApiContainerCheckpointLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersApiContainerCheckpointLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **keep** | **optional.Bool**| keep all temporary checkpoint files | 
 **leaveRunning** | **optional.Bool**| leave the container running after writing checkpoint to disk | 
 **tcpEstablished** | **optional.Bool**| checkpoint a container with established TCP connections | 
 **export** | **optional.Bool**| export the checkpoint image to a tar.gz | 
 **ignoreRootFS** | **optional.Bool**| do not include root file-system changes when exporting | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerCreateLibpod**
> InlineResponse201 ContainerCreateLibpod(ctx, optional)
Create a container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ContainersApiContainerCreateLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersApiContainerCreateLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SpecGenerator**](SpecGenerator.md)| attributes for creating a container | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-tar
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerDeleteLibpod**
> ContainerDeleteLibpod(ctx, name, optional)
Delete container

Delete container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ContainersApiContainerDeleteLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersApiContainerDeleteLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| need something | 
 **v** | **optional.Bool**| delete volumes | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerExistsLibpod**
> ContainerExistsLibpod(ctx, name)
Check if container exists

Quick way to determine if a container exists by name or ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerExportLibpod**
> ContainerExportLibpod(ctx, name)
Export a container

Export the contents of a container as a tarball.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerHealthcheckLibpod**
> InlineResponse2009 ContainerHealthcheckLibpod(ctx, name)
Run a container's healthcheck

Execute the defined healthcheck and return information about the results

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 

### Return type

[**InlineResponse2009**](inline_response_200_9.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerInitLibpod**
> ContainerInitLibpod(ctx, name)
Initialize a container

Performs all tasks necessary for initializing the container but does not start the container.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerInspectLibpod**
> InlineResponse20010 ContainerInspectLibpod(ctx, name, optional)
Inspect container

Return low-level information about a container.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ContainersApiContainerInspectLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersApiContainerInspectLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **optional.Bool**| display filesystem usage | 

### Return type

[**InlineResponse20010**](inline_response_200_10.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerKillLibpod**
> ContainerKillLibpod(ctx, name, optional)
Kill container

send a signal to a container, defaults to killing the container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ContainersApiContainerKillLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersApiContainerKillLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signal** | **optional.String**| signal to be sent to container, either by integer or SIG_ name | [default to TERM]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerListLibpod**
> []ListContainer ContainerListLibpod(ctx, optional)
List containers

Returns a list of containers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ContainersApiContainerListLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersApiContainerListLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **optional.Bool**| Return all containers. By default, only running containers are shown | [default to false]
 **limit** | **optional.Int32**| Return this number of most recently created containers, including non-running ones. | 
 **pod** | **optional.Bool**| Ignored. Previously included details on pod name and ID that are currently included by default. | [default to false]
 **size** | **optional.Bool**| Return the size of container as fields SizeRw and SizeRootFs. | [default to false]
 **sync** | **optional.Bool**| Sync container state with OCI runtime | [default to false]
 **filters** | **optional.String**| A JSON encoded value of the filters (a &#x60;map[string][]string&#x60;) to process on the containers list. Available filters: - &#x60;ancestor&#x60;&#x3D;(&#x60;&lt;image-name&gt;[:&lt;tag&gt;]&#x60;, &#x60;&lt;image id&gt;&#x60;, or &#x60;&lt;image@digest&gt;&#x60;) - &#x60;before&#x60;&#x3D;(&#x60;&lt;container id&gt;&#x60; or &#x60;&lt;container name&gt;&#x60;) - &#x60;expose&#x60;&#x3D;(&#x60;&lt;port&gt;[/&lt;proto&gt;]&#x60; or &#x60;&lt;startport-endport&gt;/[&lt;proto&gt;]&#x60;) - &#x60;exited&#x3D;&lt;int&gt;&#x60; containers with exit code of &#x60;&lt;int&gt;&#x60; - &#x60;health&#x60;&#x3D;(&#x60;starting&#x60;, &#x60;healthy&#x60;, &#x60;unhealthy&#x60; or &#x60;none&#x60;) - &#x60;id&#x3D;&lt;ID&gt;&#x60; a container&#x27;s ID - &#x60;is-task&#x60;&#x3D;(&#x60;true&#x60; or &#x60;false&#x60;) - &#x60;label&#x60;&#x3D;(&#x60;key&#x60; or &#x60;\&quot;key&#x3D;value\&quot;&#x60;) of an container label - &#x60;name&#x3D;&lt;name&gt;&#x60; a container&#x27;s name - &#x60;network&#x60;&#x3D;(&#x60;&lt;network id&gt;&#x60; or &#x60;&lt;network name&gt;&#x60;) - &#x60;pod&#x60;&#x3D;(&#x60;&lt;pod id&gt;&#x60; or &#x60;&lt;pod name&gt;&#x60;) - &#x60;publish&#x60;&#x3D;(&#x60;&lt;port&gt;[/&lt;proto&gt;]&#x60; or &#x60;&lt;startport-endport&gt;/[&lt;proto&gt;]&#x60;) - &#x60;since&#x60;&#x3D;(&#x60;&lt;container id&gt;&#x60; or &#x60;&lt;container name&gt;&#x60;) - &#x60;status&#x60;&#x3D;(&#x60;created&#x60;, &#x60;restarting&#x60;, &#x60;running&#x60;, &#x60;removing&#x60;, &#x60;paused&#x60;, &#x60;exited&#x60; or &#x60;dead&#x60;) - &#x60;volume&#x60;&#x3D;(&#x60;&lt;volume name&gt;&#x60; or &#x60;&lt;mount point destination&gt;&#x60;)  | 

### Return type

[**[]ListContainer**](ListContainer.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerLogsLibpod**
> ContainerLogsLibpod(ctx, name, optional)
Get container logs

Get stdout and stderr logs from a container.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ContainersApiContainerLogsLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersApiContainerLogsLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **follow** | **optional.Bool**| Keep connection after returning logs. | 
 **stdout** | **optional.Bool**| Return logs from stdout | 
 **stderr** | **optional.Bool**| Return logs from stderr | 
 **since** | **optional.String**| Only return logs since this time, as a UNIX timestamp | 
 **until** | **optional.String**| Only return logs before this time, as a UNIX timestamp | 
 **timestamps** | **optional.Bool**| Add timestamps to every log line | [default to false]
 **tail** | **optional.String**| Only return this number of log lines from the end of the logs | [default to all]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerMountLibpod**
> string ContainerMountLibpod(ctx, name)
Mount a container

Mount a container to the filesystem

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerPauseLibpod**
> ContainerPauseLibpod(ctx, name)
Pause a container

Use the cgroups freezer to suspend all processes in a container.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerPruneLibpod**
> []LibpodContainersPruneReport ContainerPruneLibpod(ctx, optional)
Delete stopped containers

Remove containers not in use

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ContainersApiContainerPruneLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersApiContainerPruneLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **optional.String**| Filters to process on the prune list, encoded as JSON (a &#x60;map[string][]string&#x60;).  Available filters:  - &#x60;until&#x3D;&lt;timestamp&gt;&#x60; Prune containers created before this timestamp. The &#x60;&lt;timestamp&gt;&#x60; can be Unix timestamps, date formatted timestamps, or Go duration strings (e.g. &#x60;10m&#x60;, &#x60;1h30m&#x60;) computed relative to the daemon machine’s time.  - &#x60;label&#x60; (&#x60;label&#x3D;&lt;key&gt;&#x60;, &#x60;label&#x3D;&lt;key&gt;&#x3D;&lt;value&gt;&#x60;, &#x60;label!&#x3D;&lt;key&gt;&#x60;, or &#x60;label!&#x3D;&lt;key&gt;&#x3D;&lt;value&gt;&#x60;) Prune containers with (or without, in case &#x60;label!&#x3D;...&#x60; is used) the specified labels.  | 

### Return type

[**[]LibpodContainersPruneReport**](LibpodContainersPruneReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerRenameLibpod**
> ContainerRenameLibpod(ctx, name, name)
Rename an existing container

Change the name of an existing container.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Full or partial ID or full name of the container to rename | 
  **name** | **string**| New name for the container | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerResizeLibpod**
> interface{} ContainerResizeLibpod(ctx, name, optional)
Resize a container's TTY

Resize the terminal attached to a container (for use with Attach).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ContainersApiContainerResizeLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersApiContainerResizeLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **h** | **optional.Int32**| Height to set for the terminal, in characters | 
 **w** | **optional.Int32**| Width to set for the terminal, in characters | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerRestartLibpod**
> ContainerRestartLibpod(ctx, name, optional)
Restart a container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ContainersApiContainerRestartLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersApiContainerRestartLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t** | **optional.Int32**| timeout before sending kill signal to container | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerRestoreLibpod**
> ContainerRestoreLibpod(ctx, name, optional)
Restore a container

Restore a container from a checkpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or id of the container | 
 **optional** | ***ContainersApiContainerRestoreLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersApiContainerRestoreLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| the name of the container when restored from a tar. can only be used with import | 
 **keep** | **optional.Bool**| keep all temporary checkpoint files | 
 **leaveRunning** | **optional.Bool**| leave the container running after writing checkpoint to disk | 
 **tcpEstablished** | **optional.Bool**| checkpoint a container with established TCP connections | 
 **import_** | **optional.Bool**| import the restore from a checkpoint tar.gz | 
 **ignoreRootFS** | **optional.Bool**| do not include root file-system changes when exporting | 
 **ignoreStaticIP** | **optional.Bool**| ignore IP address if set statically | 
 **ignoreStaticMAC** | **optional.Bool**| ignore MAC address if set statically | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerShowMountedLibpod**
> map[string]string ContainerShowMountedLibpod(ctx, )
Show mounted containers

Lists all mounted containers mount points

### Required Parameters
This endpoint does not need any parameter.

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerStartLibpod**
> ContainerStartLibpod(ctx, name, optional)
Start a container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ContainersApiContainerStartLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersApiContainerStartLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **detachKeys** | **optional.String**| Override the key sequence for detaching a container. Format is a single character [a-Z] or ctrl-&lt;value&gt; where &lt;value&gt; is one of: a-z, @, ^, [, , or _. | [default to ctrl-p,ctrl-q]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerStatsLibpod**
> ContainerStatsLibpod(ctx, name, optional)
Get stats for a container

DEPRECATED. This endpoint will be removed with the next major release. Please use /libpod/containers/stats instead.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ContainersApiContainerStatsLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersApiContainerStatsLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stream** | **optional.Bool**| Stream the output | [default to true]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerStopLibpod**
> ContainerStopLibpod(ctx, name, optional)
Stop a container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ContainersApiContainerStopLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersApiContainerStopLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **all** | **optional.Bool**| Stop all containers | [default to false]
 **timeout** | **optional.Int32**| number of seconds to wait before killing container | [default to 10]
 **ignore** | **optional.Bool**| do not return error if container is already stopped | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerTopLibpod**
> InlineResponse2003 ContainerTopLibpod(ctx, name, optional)
List processes

List processes running inside a container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of container to query for processes (As of version 1.xx)  | 
 **optional** | ***ContainersApiContainerTopLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersApiContainerTopLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stream** | **optional.Bool**| Stream the output | [default to true]
 **psArgs** | **optional.String**| arguments to pass to ps such as aux. Requires ps(1) to be installed in the container if no ps(1) compatible AIX descriptors are used. | [default to -ef]

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerUnmountLibpod**
> ContainerUnmountLibpod(ctx, name)
Unmount a container

Unmount a container from the filesystem

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerUnpauseLibpod**
> ContainerUnpauseLibpod(ctx, name)
Unpause Container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerWaitLibpod**
> int32 ContainerWaitLibpod(ctx, name, optional)
Wait on a container

Wait on a container to met a given condition

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ContainersApiContainerWaitLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersApiContainerWaitLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **condition** | [**optional.Interface of []string**](string.md)| Conditions to wait for. If no condition provided the &#x27;exited&#x27; condition is assumed. | 
 **interval** | **optional.String**| Time Interval to wait before polling for completion. | [default to 250ms]

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainersStatsAllLibpod**
> ContainersStatsAllLibpod(ctx, optional)
Get stats for one or more containers

Return a live stream of resource usage statistics of one or more container. If no container is specified, the statistics of all containers are returned.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ContainersApiContainersStatsAllLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersApiContainersStatsAllLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containers** | [**optional.Interface of []string**](string.md)| names or IDs of containers | 
 **stream** | **optional.Bool**| Stream the output | [default to true]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateKubeLibpod**
> *os.File GenerateKubeLibpod(ctx, names, optional)
Generate a Kubernetes YAML file.

Generate Kubernetes YAML based on a pod or container.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **names** | [**[]string**](string.md)| Name or ID of the container or pod. | 
 **optional** | ***ContainersApiGenerateKubeLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersApiGenerateKubeLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **service** | **optional.Bool**| Generate YAML for a Kubernetes service object. | [default to false]

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GenerateSystemdLibpod**
> map[string]string GenerateSystemdLibpod(ctx, name, optional)
Generate Systemd Units

Generate Systemd Units based on a pod or container.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name or ID of the container or pod. | 
 **optional** | ***ContainersApiGenerateSystemdLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersApiGenerateSystemdLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **useName** | **optional.Bool**| Use container/pod names instead of IDs. | [default to false]
 **new** | **optional.Bool**| Create a new container instead of starting an existing one. | [default to false]
 **noHeader** | **optional.Bool**| Do not generate the header including the Podman version and the timestamp. | [default to false]
 **time** | **optional.Int32**| Stop timeout override. | [default to 10]
 **restartPolicy** | **optional.String**| Systemd restart-policy. | [default to on-failure]
 **containerPrefix** | **optional.String**| Systemd unit name prefix for containers. | [default to container]
 **podPrefix** | **optional.String**| Systemd unit name prefix for pods. | [default to pod]
 **separator** | **optional.String**| Systemd unit name separator between name/id and prefix. | [default to -]

### Return type

**map[string]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageCommitLibpod**
> ImageCommitLibpod(ctx, container, optional)
Commit

Create a new image from a container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **container** | **string**| the name or ID of a container | 
 **optional** | ***ContainersApiImageCommitLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersApiImageCommitLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repo** | **optional.String**| the repository name for the created image | 
 **tag** | **optional.String**| tag name for the created image | 
 **comment** | **optional.String**| commit message | 
 **author** | **optional.String**| author of the image | 
 **pause** | **optional.Bool**| pause the container before committing it | 
 **changes** | [**optional.Interface of []string**](string.md)| instructions to apply while committing in Dockerfile format (i.e. \&quot;CMD&#x3D;/bin/foo\&quot;) | 
 **format** | **optional.String**| format of the image manifest and metadata (default \&quot;oci\&quot;) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlayKubeLibpod**
> PlayKubeReport PlayKubeLibpod(ctx, optional)
Play a Kubernetes YAML file.

Create and run pods based on a Kubernetes YAML file (pod or service kind).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ContainersApiPlayKubeLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersApiPlayKubeLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of string**](string.md)| Kubernetes YAML file. | 
 **network** | **optional.**| Connect the pod to this network. | 
 **tlsVerify** | **optional.**| Require HTTPS and verify signatures when contacting registries. | [default to true]
 **logDriver** | **optional.**| Logging driver for the containers in the pod. | 
 **start** | **optional.**| Start the pod after creating it. | [default to true]

### Return type

[**PlayKubeReport**](PlayKubeReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-tar
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutContainerArchiveLibpod**
> PutContainerArchiveLibpod(ctx, name, path, optional)
Copy files into a container

Copy a tar archive of files into a container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| container name or id | 
  **path** | **string**| Path to a directory in the container to extract | 
 **optional** | ***ContainersApiPutContainerArchiveLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersApiPutContainerArchiveLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of string**](string.md)| tarfile of files to copy into the container | 
 **pause** | **optional.**| pause the container while copying (defaults to true) | [default to true]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-tar
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

