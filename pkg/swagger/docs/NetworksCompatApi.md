# {{classname}}

All URIs are relative to *http://podman.io/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetworkConnect**](NetworksCompatApi.md#NetworkConnect) | **Post** /networks/{name}/connect | Connect container to network
[**NetworkCreate**](NetworksCompatApi.md#NetworkCreate) | **Post** /networks/create | Create network
[**NetworkDelete**](NetworksCompatApi.md#NetworkDelete) | **Delete** /networks/{name} | Remove a network
[**NetworkDisconnect**](NetworksCompatApi.md#NetworkDisconnect) | **Post** /networks/{name}/disconnect | Disconnect container from network
[**NetworkInspect**](NetworksCompatApi.md#NetworkInspect) | **Get** /networks/{name} | Inspect a network
[**NetworkList**](NetworksCompatApi.md#NetworkList) | **Get** /networks | List networks
[**NetworkPrune**](NetworksCompatApi.md#NetworkPrune) | **Post** /networks/prune | Delete unused networks

# **NetworkConnect**
> NetworkConnect(ctx, name, optional)
Connect container to network

Connect a container to a network.  This endpoint is current a no-op

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name of the network | 
 **optional** | ***NetworksCompatApiNetworkConnectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworksCompatApiNetworkConnectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NetworkConnectRequest**](NetworkConnectRequest.md)| attributes for connecting a container to a network | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-tar
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NetworkCreate**
> InlineResponse20018 NetworkCreate(ctx, optional)
Create network

Create a network configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NetworksCompatApiNetworkCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworksCompatApiNetworkCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NetworkCreateRequest**](NetworkCreateRequest.md)| attributes for creating a container | 

### Return type

[**InlineResponse20018**](inline_response_200_18.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-tar
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NetworkDelete**
> NetworkDelete(ctx, name)
Remove a network

Remove a network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name of the network | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NetworkDisconnect**
> NetworkDisconnect(ctx, name, optional)
Disconnect container from network

Disconnect a container from a network.  This endpoint is current a no-op

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name of the network | 
 **optional** | ***NetworksCompatApiNetworkDisconnectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworksCompatApiNetworkDisconnectOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of NetworkDisconnectRequest**](NetworkDisconnectRequest.md)| attributes for disconnecting a container from a network | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-tar
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NetworkInspect**
> NetworkResource NetworkInspect(ctx, name)
Inspect a network

Display low level configuration network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name of the network | 

### Return type

[**NetworkResource**](NetworkResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NetworkList**
> []NetworkResource NetworkList(ctx, optional)
List networks

Display summary of network configurations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NetworksCompatApiNetworkListOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworksCompatApiNetworkListOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **optional.String**| JSON encoded value of the filters (a map[string][]string) to process on the network list. Currently available filters:   - name&#x3D;[name] Matches network name (accepts regex).   - id&#x3D;[id] Matches for full or partial ID.   - driver&#x3D;[driver] Only bridge is supported.   - label&#x3D;[key] or label&#x3D;[key&#x3D;value] Matches networks based on the presence of a label alone or a label and a value.  | 

### Return type

[**[]NetworkResource**](NetworkResource.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NetworkPrune**
> []string NetworkPrune(ctx, optional)
Delete unused networks

Remove CNI networks that do not have containers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NetworksCompatApiNetworkPruneOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworksCompatApiNetworkPruneOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **optional.String**| Filters to process on the prune list, encoded as JSON (a map[string][]string). Available filters:   - until&#x3D;&lt;timestamp&gt; Prune networks created before this timestamp. The &lt;timestamp&gt; can be Unix timestamps, date formatted timestamps, or Go duration strings (e.g. 10m, 1h30m) computed relative to the daemon machine’s time.  - label (label&#x3D;&lt;key&gt;, label&#x3D;&lt;key&gt;&#x3D;&lt;value&gt;, label!&#x3D;&lt;key&gt;, or label!&#x3D;&lt;key&gt;&#x3D;&lt;value&gt;) Prune networks with (or without, in case label!&#x3D;... is used) the specified labels.  | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

