# {{classname}}

All URIs are relative to *http://podman.io/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GenerateKubeLibpod**](PodsApi.md#GenerateKubeLibpod) | **Get** /libpod/generate/kube | Generate a Kubernetes YAML file.
[**GenerateSystemdLibpod**](PodsApi.md#GenerateSystemdLibpod) | **Get** /libpod/generate/{name}/systemd | Generate Systemd Units
[**PlayKubeLibpod**](PodsApi.md#PlayKubeLibpod) | **Post** /libpod/play/kube | Play a Kubernetes YAML file.
[**PodCreateLibpod**](PodsApi.md#PodCreateLibpod) | **Post** /libpod/pods/create | Create a pod
[**PodDeleteLibpod**](PodsApi.md#PodDeleteLibpod) | **Delete** /libpod/pods/{name} | Remove pod
[**PodExistsLibpod**](PodsApi.md#PodExistsLibpod) | **Get** /libpod/pods/{name}/exists | Pod exists
[**PodInspectLibpod**](PodsApi.md#PodInspectLibpod) | **Get** /libpod/pods/{name}/json | Inspect pod
[**PodKillLibpod**](PodsApi.md#PodKillLibpod) | **Post** /libpod/pods/{name}/kill | Kill a pod
[**PodListLibpod**](PodsApi.md#PodListLibpod) | **Get** /libpod/pods/json | List pods
[**PodPauseLibpod**](PodsApi.md#PodPauseLibpod) | **Post** /libpod/pods/{name}/pause | Pause a pod
[**PodPruneLibpod**](PodsApi.md#PodPruneLibpod) | **Post** /libpod/pods/prune | Prune unused pods
[**PodRestartLibpod**](PodsApi.md#PodRestartLibpod) | **Post** /libpod/pods/{name}/restart | Restart a pod
[**PodStartLibpod**](PodsApi.md#PodStartLibpod) | **Post** /libpod/pods/{name}/start | Start a pod
[**PodStatsAllLibpod**](PodsApi.md#PodStatsAllLibpod) | **Get** /libpod/pods/stats | Get stats for one or more pods
[**PodStopLibpod**](PodsApi.md#PodStopLibpod) | **Post** /libpod/pods/{name}/stop | Stop a pod
[**PodTopLibpod**](PodsApi.md#PodTopLibpod) | **Get** /libpod/pods/{name}/top | List processes
[**PodUnpauseLibpod**](PodsApi.md#PodUnpauseLibpod) | **Post** /libpod/pods/{name}/unpause | Unpause a pod

# **GenerateKubeLibpod**
> *os.File GenerateKubeLibpod(ctx, names, optional)
Generate a Kubernetes YAML file.

Generate Kubernetes YAML based on a pod or container.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **names** | [**[]string**](string.md)| Name or ID of the container or pod. | 
 **optional** | ***PodsApiGenerateKubeLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PodsApiGenerateKubeLibpodOpts struct
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
 **optional** | ***PodsApiGenerateSystemdLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PodsApiGenerateSystemdLibpodOpts struct
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

# **PlayKubeLibpod**
> PlayKubeReport PlayKubeLibpod(ctx, optional)
Play a Kubernetes YAML file.

Create and run pods based on a Kubernetes YAML file (pod or service kind).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PodsApiPlayKubeLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PodsApiPlayKubeLibpodOpts struct
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

# **PodCreateLibpod**
> IdResponse PodCreateLibpod(ctx, optional)
Create a pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PodsApiPodCreateLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PodsApiPodCreateLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PodSpecGenerator**](PodSpecGenerator.md)| attributes for creating a pod | 

### Return type

[**IdResponse**](IdResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-tar
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodDeleteLibpod**
> PodRmReport PodDeleteLibpod(ctx, name, optional)
Remove pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the pod | 
 **optional** | ***PodsApiPodDeleteLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PodsApiPodDeleteLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| force removal of a running pod by first stopping all containers, then removing all containers in the pod | 

### Return type

[**PodRmReport**](PodRmReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodExistsLibpod**
> PodExistsLibpod(ctx, name)
Pod exists

Check if a pod exists by name or ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the pod | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodInspectLibpod**
> InlineResponse20013 PodInspectLibpod(ctx, name)
Inspect pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the pod | 

### Return type

[**InlineResponse20013**](inline_response_200_13.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodKillLibpod**
> PodKillReport PodKillLibpod(ctx, name, optional)
Kill a pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the pod | 
 **optional** | ***PodsApiPodKillLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PodsApiPodKillLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **signal** | **optional.String**| signal to be sent to pod | [default to SIGKILL]

### Return type

[**PodKillReport**](PodKillReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodListLibpod**
> []ListPodsReport PodListLibpod(ctx, optional)
List pods

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PodsApiPodListLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PodsApiPodListLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **optional.String**| needs description and plumbing for filters | 

### Return type

[**[]ListPodsReport**](ListPodsReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodPauseLibpod**
> PodPauseReport PodPauseLibpod(ctx, name)
Pause a pod

Pause a pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the pod | 

### Return type

[**PodPauseReport**](PodPauseReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodPruneLibpod**
> PodPruneReport PodPruneLibpod(ctx, )
Prune unused pods

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PodPruneReport**](PodPruneReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodRestartLibpod**
> PodRestartReport PodRestartLibpod(ctx, name)
Restart a pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the pod | 

### Return type

[**PodRestartReport**](PodRestartReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodStartLibpod**
> PodStartReport PodStartLibpod(ctx, name)
Start a pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the pod | 

### Return type

[**PodStartReport**](PodStartReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodStatsAllLibpod**
> InlineResponse2003 PodStatsAllLibpod(ctx, optional)
Get stats for one or more pods

Display a live stream of resource usage statistics for the containers in one or more pods

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PodsApiPodStatsAllLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PodsApiPodStatsAllLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **optional.Bool**| Provide statistics for all running pods. | 
 **namesOrIDs** | [**optional.Interface of []string**](string.md)| Names or IDs of pods. | 

### Return type

[**InlineResponse2003**](inline_response_200_3.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodStopLibpod**
> PodStopReport PodStopLibpod(ctx, name, optional)
Stop a pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the pod | 
 **optional** | ***PodsApiPodStopLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PodsApiPodStopLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **t** | **optional.Int32**| timeout | 

### Return type

[**PodStopReport**](PodStopReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PodTopLibpod**
> InlineResponse2003 PodTopLibpod(ctx, name, optional)
List processes

List processes running inside a pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of pod to query for processes  | 
 **optional** | ***PodsApiPodTopLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PodsApiPodTopLibpodOpts struct
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

# **PodUnpauseLibpod**
> PodUnpauseReport PodUnpauseLibpod(ctx, name)
Unpause a pod

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the pod | 

### Return type

[**PodUnpauseReport**](PodUnpauseReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

