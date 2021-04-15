# {{classname}}

All URIs are relative to *http://podman.io/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SystemDataUsageLibpod**](SystemApi.md#SystemDataUsageLibpod) | **Get** /libpod/system/df | Show disk usage
[**SystemEventsLibpod**](SystemApi.md#SystemEventsLibpod) | **Get** /libpod/events | Get events
[**SystemInfoLibpod**](SystemApi.md#SystemInfoLibpod) | **Get** /libpod/info | Get info
[**SystemPing**](SystemApi.md#SystemPing) | **Get** /libpod/_ping | Ping service
[**SystemPruneLibpod**](SystemApi.md#SystemPruneLibpod) | **Post** /libpod/system/prune | Prune unused data
[**SystemVersionLibpod**](SystemApi.md#SystemVersionLibpod) | **Get** /libpod/version | Component Version information

# **SystemDataUsageLibpod**
> InlineResponse20014 SystemDataUsageLibpod(ctx, )
Show disk usage

Return information about disk usage for containers, images, and volumes

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20014**](inline_response_200_14.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemEventsLibpod**
> SystemEventsLibpod(ctx, optional)
Get events

Returns events filtered on query parameters

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SystemApiSystemEventsLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SystemApiSystemEventsLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **optional.String**| start streaming events from this time | 
 **until** | **optional.String**| stop streaming events later than this | 
 **filters** | **optional.String**| JSON encoded map[string][]string of constraints | 
 **stream** | **optional.Bool**| when false, do not follow events | [default to true]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemInfoLibpod**
> Info SystemInfoLibpod(ctx, )
Get info

Returns information on the system and libpod configuration

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**Info**](Info.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemPing**
> string SystemPing(ctx, )
Ping service

Return protocol information in response headers. `HEAD /libpod/_ping` is also supported. `/_ping` is available for compatibility with other engines. The '_ping' endpoints are not versioned. 

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemPruneLibpod**
> InlineResponse20015 SystemPruneLibpod(ctx, )
Prune unused data

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20015**](inline_response_200_15.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SystemVersionLibpod**
> InlineResponse20016 SystemVersionLibpod(ctx, )
Component Version information

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20016**](inline_response_200_16.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

