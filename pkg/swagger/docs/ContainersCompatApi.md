# {{classname}}

All URIs are relative to *http://podman.io/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ContainerArchive**](ContainersCompatApi.md#ContainerArchive) | **Get** /containers/{name}/archive | Get files from a container
[**ContainerArchiveLibpod**](ContainersCompatApi.md#ContainerArchiveLibpod) | **Get** /libpod/containers/{name}/archive | Copy files from a container
[**ContainerAttach**](ContainersCompatApi.md#ContainerAttach) | **Post** /containers/{name}/attach | Attach to a container
[**ContainerChangesLibpod**](ContainersCompatApi.md#ContainerChangesLibpod) | **Get** /libpod/containers/{name}/changes | Report on changes to container&#x27;s filesystem; adds, deletes or modifications.
[**ContainerCreate**](ContainersCompatApi.md#ContainerCreate) | **Post** /containers/create | Create a container
[**ContainerDelete**](ContainersCompatApi.md#ContainerDelete) | **Delete** /containers/{name} | Remove a container
[**ContainerExport**](ContainersCompatApi.md#ContainerExport) | **Get** /containers/{name}/export | Export a container
[**ContainerInspect**](ContainersCompatApi.md#ContainerInspect) | **Get** /containers/{name}/json | Inspect container
[**ContainerKill**](ContainersCompatApi.md#ContainerKill) | **Post** /containers/{name}/kill | Kill container
[**ContainerList**](ContainersCompatApi.md#ContainerList) | **Get** /containers/json | List containers
[**ContainerLogs**](ContainersCompatApi.md#ContainerLogs) | **Get** /containers/{name}/logs | Get container logs
[**ContainerPause**](ContainersCompatApi.md#ContainerPause) | **Post** /containers/{name}/pause | Pause container
[**ContainerPrune**](ContainersCompatApi.md#ContainerPrune) | **Post** /containers/prune | Delete stopped containers
[**ContainerRename**](ContainersCompatApi.md#ContainerRename) | **Post** /containers/{name}/rename | Rename an existing container
[**ContainerResize**](ContainersCompatApi.md#ContainerResize) | **Post** /containers/{name}/resize | Resize a container&#x27;s TTY
[**ContainerRestart**](ContainersCompatApi.md#ContainerRestart) | **Post** /containers/{name}/restart | Restart container
[**ContainerStart**](ContainersCompatApi.md#ContainerStart) | **Post** /containers/{name}/start | Start a container
[**ContainerStats**](ContainersCompatApi.md#ContainerStats) | **Get** /containers/{name}/stats | Get stats for a container
[**ContainerStop**](ContainersCompatApi.md#ContainerStop) | **Post** /containers/{name}/stop | Stop a container
[**ContainerTop**](ContainersCompatApi.md#ContainerTop) | **Get** /containers/{name}/top | List processes running inside a container
[**ContainerUnpause**](ContainersCompatApi.md#ContainerUnpause) | **Post** /containers/{name}/unpause | Unpause container
[**ContainerWait**](ContainersCompatApi.md#ContainerWait) | **Post** /containers/{name}/wait | Wait on a container
[**ImageCommit**](ContainersCompatApi.md#ImageCommit) | **Post** /commit | New Image
[**PutContainerArchive**](ContainersCompatApi.md#PutContainerArchive) | **Put** /containers/{name}/archive | Put files into a container

# **ContainerArchive**
> *os.File ContainerArchive(ctx, name, path)
Get files from a container

Get a tar archive of files from a container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| container name or id | 
  **path** | **string**| Path to a directory in the container to extract | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerArchiveLibpod**
> *os.File ContainerArchiveLibpod(ctx, name, path)
Copy files from a container

Copy a tar archive of files from a container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| container name or id | 
  **path** | **string**| Path to a directory in the container to extract | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerAttach**
> ContainerAttach(ctx, name, optional)
Attach to a container

Hijacks the connection to forward the container's standard streams to the client.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ContainersCompatApiContainerAttachOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersCompatApiContainerAttachOpts struct
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

# **ContainerCreate**
> InlineResponse201 ContainerCreate(ctx, optional)
Create a container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ContainersCompatApiContainerCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersCompatApiContainerCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **optional.String**| container name | 

### Return type

[**InlineResponse201**](inline_response_201.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerDelete**
> ContainerDelete(ctx, name, optional)
Remove a container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ContainersCompatApiContainerDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersCompatApiContainerDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| If the container is running, kill it before removing it. | [default to false]
 **v** | **optional.Bool**| Remove the volumes associated with the container. | [default to false]
 **link** | **optional.Bool**| not supported | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerExport**
> ContainerExport(ctx, name)
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

# **ContainerInspect**
> InlineResponse2002 ContainerInspect(ctx, name, optional)
Inspect container

Return low-level information about a container.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or id of the container | 
 **optional** | ***ContainersCompatApiContainerInspectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersCompatApiContainerInspectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **optional.Bool**| include the size of the container | [default to false]

### Return type

[**InlineResponse2002**](inline_response_200_2.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerKill**
> ContainerKill(ctx, name, optional)
Kill container

Signal to send to the container as an integer or string (e.g. SIGINT)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ContainersCompatApiContainerKillOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersCompatApiContainerKillOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **all** | **optional.Bool**| Send kill signal to all containers | [default to false]
 **signal** | **optional.String**| signal to be sent to container | [default to SIGKILL]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerList**
> interface{} ContainerList(ctx, optional)
List containers

Returns a list of containers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ContainersCompatApiContainerListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersCompatApiContainerListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **optional.Bool**| Return all containers. By default, only running containers are shown | [default to false]
 **external** | **optional.Bool**| Return containers in storage not controlled by Podman | [default to false]
 **limit** | **optional.Int32**| Return this number of most recently created containers, including non-running ones. | 
 **size** | **optional.Bool**| Return the size of container as fields SizeRw and SizeRootFs. | [default to false]
 **filters** | **optional.String**| Returns a list of containers.  - ancestor&#x3D;(&lt;image-name&gt;[:&lt;tag&gt;], &lt;image id&gt;, or &lt;image@digest&gt;)  - before&#x3D;(&lt;container id&gt; or &lt;container name&gt;)  - expose&#x3D;(&lt;port&gt;[/&lt;proto&gt;]|&lt;startport-endport&gt;/[&lt;proto&gt;])  - exited&#x3D;&lt;int&gt; containers with exit code of &lt;int&gt;  - health&#x3D;(starting|healthy|unhealthy|none)  - id&#x3D;&lt;ID&gt; a container&#x27;s ID  - is-task&#x3D;(true|false)  - label&#x3D;key or label&#x3D;\&quot;key&#x3D;value\&quot; of a container label  - name&#x3D;&lt;name&gt; a container&#x27;s name  - network&#x3D;(&lt;network id&gt; or &lt;network name&gt;)  - publish&#x3D;(&lt;port&gt;[/&lt;proto&gt;]|&lt;startport-endport&gt;/[&lt;proto&gt;])  - since&#x3D;(&lt;container id&gt; or &lt;container name&gt;)  - status&#x3D;(created|restarting|running|removing|paused|exited|dead)  - volume&#x3D;(&lt;volume name&gt; or &lt;mount point destination&gt;)  | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerLogs**
> ContainerLogs(ctx, name, optional)
Get container logs

Get stdout and stderr logs from a container.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ContainersCompatApiContainerLogsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersCompatApiContainerLogsOpts struct
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

# **ContainerPause**
> ContainerPause(ctx, name)
Pause container

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

# **ContainerPrune**
> []ContainersPruneReport ContainerPrune(ctx, optional)
Delete stopped containers

Remove containers not in use

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ContainersCompatApiContainerPruneOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersCompatApiContainerPruneOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **optional.String**| Filters to process on the prune list, encoded as JSON (a &#x60;map[string][]string&#x60;).  Available filters:  - &#x60;until&#x3D;&lt;timestamp&gt;&#x60; Prune containers created before this timestamp. The &#x60;&lt;timestamp&gt;&#x60; can be Unix timestamps, date formatted timestamps, or Go duration strings (e.g. &#x60;10m&#x60;, &#x60;1h30m&#x60;) computed relative to the daemon machine’s time.  - &#x60;label&#x60; (&#x60;label&#x3D;&lt;key&gt;&#x60;, &#x60;label&#x3D;&lt;key&gt;&#x3D;&lt;value&gt;&#x60;, &#x60;label!&#x3D;&lt;key&gt;&#x60;, or &#x60;label!&#x3D;&lt;key&gt;&#x3D;&lt;value&gt;&#x60;) Prune containers with (or without, in case &#x60;label!&#x3D;...&#x60; is used) the specified labels.  | 

### Return type

[**[]ContainersPruneReport**](ContainersPruneReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerRename**
> ContainerRename(ctx, name, name)
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

# **ContainerResize**
> interface{} ContainerResize(ctx, name, optional)
Resize a container's TTY

Resize the terminal attached to a container (for use with Attach).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ContainersCompatApiContainerResizeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersCompatApiContainerResizeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **h** | **optional.Int32**| Height to set for the terminal, in characters | 
 **w** | **optional.Int32**| Width to set for the terminal, in characters | 
 **running** | **optional.Bool**| Ignore containers not running errors | 

### Return type

[**interface{}**](interface{}.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerRestart**
> ContainerRestart(ctx, name, optional)
Restart container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ContainersCompatApiContainerRestartOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersCompatApiContainerRestartOpts struct
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

# **ContainerStart**
> ContainerStart(ctx, name, optional)
Start a container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ContainersCompatApiContainerStartOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersCompatApiContainerStartOpts struct
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

# **ContainerStats**
> ContainerStats(ctx, name, optional)
Get stats for a container

This returns a live stream of a container’s resource usage statistics.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ContainersCompatApiContainerStatsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersCompatApiContainerStatsOpts struct
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

# **ContainerStop**
> ContainerStop(ctx, name, optional)
Stop a container

Stop a container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ContainersCompatApiContainerStopOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersCompatApiContainerStopOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t** | **optional.Int32**| number of seconds to wait before killing container | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerTop**
> InlineResponse2003 ContainerTop(ctx, name, optional)
List processes running inside a container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ContainersCompatApiContainerTopOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersCompatApiContainerTopOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **psArgs** | **optional.String**| arguments to pass to ps such as aux. Requires ps(1) to be installed in the container if no ps(1) compatible AIX descriptors are used. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContainerUnpause**
> ContainerUnpause(ctx, name)
Unpause container

Resume a paused container

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

# **ContainerWait**
> InlineResponse2004 ContainerWait(ctx, name, optional)
Wait on a container

Block until a container stops or given condition is met.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ContainersCompatApiContainerWaitOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersCompatApiContainerWaitOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **condition** | **optional.String**| wait until container is to a given condition. default is stopped. valid conditions are:   - configured   - created   - exited   - paused   - running   - stopped  | 
 **interval** | **optional.String**| Time Interval to wait before polling for completion. | [default to 250ms]

### Return type

[**InlineResponse2004**](inline_response_200_4.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageCommit**
> ImageCommit(ctx, optional)
New Image

Create a new image from a container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ContainersCompatApiImageCommitOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersCompatApiImageCommitOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **container** | **optional.String**| the name or ID of a container | 
 **repo** | **optional.String**| the repository name for the created image | 
 **tag** | **optional.String**| tag name for the created image | 
 **comment** | **optional.String**| commit message | 
 **author** | **optional.String**| author of the image | 
 **pause** | **optional.Bool**| pause the container before committing it | 
 **changes** | **optional.String**| instructions to apply while committing in Dockerfile format | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PutContainerArchive**
> PutContainerArchive(ctx, name, path, optional)
Put files into a container

Put a tar archive of files into a container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| container name or id | 
  **path** | **string**| Path to a directory in the container to extract | 
 **optional** | ***ContainersCompatApiPutContainerArchiveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ContainersCompatApiPutContainerArchiveOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of string**](string.md)| tarfile of files to copy into the container | 
 **noOverwriteDirNonDir** | **optional.**| if unpacking the given content would cause an existing directory to be replaced with a non-directory and vice versa (1 or true) | 
 **copyUIDGID** | **optional.**| copy UID/GID maps to the dest file or di (1 or true) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-tar
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

