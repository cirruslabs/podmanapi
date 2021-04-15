# {{classname}}

All URIs are relative to *http://podman.io/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**NetworkConnectLibpod**](NetworksApi.md#NetworkConnectLibpod) | **Post** /libpod/networks/{name}/connect | Connect container to network
[**NetworkCreateLibpod**](NetworksApi.md#NetworkCreateLibpod) | **Post** /libpod/networks/create | Create network
[**NetworkDeleteLibpod**](NetworksApi.md#NetworkDeleteLibpod) | **Delete** /libpod/networks/{name} | Remove a network
[**NetworkDisconnectLibpod**](NetworksApi.md#NetworkDisconnectLibpod) | **Post** /libpod/networks/{name}/disconnect | Disconnect container from network
[**NetworkExistsLibpod**](NetworksApi.md#NetworkExistsLibpod) | **Get** /libpod/networks/{name}/exists | Network exists
[**NetworkInspectLibpod**](NetworksApi.md#NetworkInspectLibpod) | **Get** /libpod/networks/{name}/json | Inspect a network
[**NetworkListLibpod**](NetworksApi.md#NetworkListLibpod) | **Get** /libpod/networks/json | List networks
[**NetworkPruneLibpod**](NetworksApi.md#NetworkPruneLibpod) | **Post** /libpod/networks/prune | Delete unused networks

# **NetworkConnectLibpod**
> NetworkConnectLibpod(ctx, name, optional)
Connect container to network

Connect a container to a network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name of the network | 
 **optional** | ***NetworksApiNetworkConnectLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworksApiNetworkConnectLibpodOpts struct
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

# **NetworkCreateLibpod**
> NetworkCreateReport NetworkCreateLibpod(ctx, optional)
Create network

Create a new CNI network configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NetworksApiNetworkCreateLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworksApiNetworkCreateLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of NetworkCreateOptions**](NetworkCreateOptions.md)| attributes for creating a container | 
 **name** | **optional.**| optional name for new network | 

### Return type

[**NetworkCreateReport**](NetworkCreateReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-tar
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NetworkDeleteLibpod**
> NetworkRmReport NetworkDeleteLibpod(ctx, name, optional)
Remove a network

Remove a CNI configured network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name of the network | 
 **optional** | ***NetworksApiNetworkDeleteLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworksApiNetworkDeleteLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| remove containers associated with network | 

### Return type

[**NetworkRmReport**](NetworkRmReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NetworkDisconnectLibpod**
> NetworkDisconnectLibpod(ctx, name, optional)
Disconnect container from network

Disconnect a container from a network.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name of the network | 
 **optional** | ***NetworksApiNetworkDisconnectLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworksApiNetworkDisconnectLibpodOpts struct
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

# **NetworkExistsLibpod**
> NetworkExistsLibpod(ctx, name)
Network exists

Check if network exists

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the network | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NetworkInspectLibpod**
> map[string]interface{} NetworkInspectLibpod(ctx, name)
Inspect a network

Display low level configuration for a CNI network

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name of the network | 

### Return type

[**map[string]interface{}**](map.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NetworkListLibpod**
> []NetworkListReport NetworkListLibpod(ctx, optional)
List networks

Display summary of network configurations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NetworksApiNetworkListLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworksApiNetworkListLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **optional.String**| JSON encoded value of the filters (a map[string][]string) to process on the network list. Available filters:   - name&#x3D;[name] Matches network name (accepts regex).   - id&#x3D;[id] Matches for full or partial ID.   - driver&#x3D;[driver] Only bridge is supported.   - label&#x3D;[key] or label&#x3D;[key&#x3D;value] Matches networks based on the presence of a label alone or a label and a value.   - plugin&#x3D;[plugin] Matches CNI plugins included in a network (e.g &#x60;bridge&#x60;,&#x60;portmap&#x60;,&#x60;firewall&#x60;,&#x60;tuning&#x60;,&#x60;dnsname&#x60;,&#x60;macvlan&#x60;)  | 

### Return type

[**[]NetworkListReport**](NetworkListReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **NetworkPruneLibpod**
> []string NetworkPruneLibpod(ctx, optional)
Delete unused networks

Remove CNI networks that do not have containers

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***NetworksApiNetworkPruneLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a NetworksApiNetworkPruneLibpodOpts struct
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

