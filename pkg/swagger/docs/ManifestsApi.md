# {{classname}}

All URIs are relative to *http://podman.io/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ManifestAddLibpod**](ManifestsApi.md#ManifestAddLibpod) | **Post** /libpod/manifests/{name}/add | Add image
[**ManifestCreateLibpod**](ManifestsApi.md#ManifestCreateLibpod) | **Post** /libpod/manifests/create | Create
[**ManifestDeleteLibpod**](ManifestsApi.md#ManifestDeleteLibpod) | **Delete** /libpod/manifests/{name} | Remove
[**ManifestExistsLibpod**](ManifestsApi.md#ManifestExistsLibpod) | **Get** /libpod/manifests/{name}/exists | Exists
[**ManifestInspectLibpod**](ManifestsApi.md#ManifestInspectLibpod) | **Get** /libpod/manifests/{name}/json | Inspect
[**ManifestPushLibpod**](ManifestsApi.md#ManifestPushLibpod) | **Post** /libpod/manifests/{name}/push | Push

# **ManifestAddLibpod**
> IdResponse ManifestAddLibpod(ctx, name, optional)
Add image

Add an image to a manifest list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the manifest | 
 **optional** | ***ManifestsApiManifestAddLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManifestsApiManifestAddLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of ManifestAddOpts**](ManifestAddOpts.md)| options for creating a manifest | 

### Return type

[**IdResponse**](IDResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-tar
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ManifestCreateLibpod**
> IdResponse ManifestCreateLibpod(ctx, name, optional)
Create

Create a manifest list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| manifest list name | 
 **optional** | ***ManifestsApiManifestCreateLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManifestsApiManifestCreateLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **image** | **optional.String**| name of the image | 
 **all** | **optional.Bool**| add all contents if given list | 

### Return type

[**IdResponse**](IDResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ManifestDeleteLibpod**
> IdResponse ManifestDeleteLibpod(ctx, name, optional)
Remove

Remove an image from a manifest list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the image associated with the manifest | 
 **optional** | ***ManifestsApiManifestDeleteLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManifestsApiManifestDeleteLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **digest** | **optional.String**| image digest to be removed | 

### Return type

[**IdResponse**](IDResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ManifestExistsLibpod**
> ManifestExistsLibpod(ctx, name)
Exists

Check if manifest list exists

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name of the manifest list | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ManifestInspectLibpod**
> Schema2List ManifestInspectLibpod(ctx, name)
Inspect

Display a manifest list

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the manifest | 

### Return type

[**Schema2List**](Schema2List.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ManifestPushLibpod**
> IdResponse ManifestPushLibpod(ctx, name, destination, optional)
Push

Push a manifest list or image index to a registry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the manifest | 
  **destination** | **string**| the destination for the manifest | 
 **optional** | ***ManifestsApiManifestPushLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ManifestsApiManifestPushLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **all** | **optional.Bool**| push all images | 

### Return type

[**IdResponse**](IDResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

