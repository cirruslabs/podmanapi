# {{classname}}

All URIs are relative to *http://podman.io/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VolumeCreateLibpod**](VolumesApi.md#VolumeCreateLibpod) | **Post** /libpod/volumes/create | Create a volume
[**VolumeDeleteLibpod**](VolumesApi.md#VolumeDeleteLibpod) | **Delete** /libpod/volumes/{name} | Remove volume
[**VolumeExistsLibpod**](VolumesApi.md#VolumeExistsLibpod) | **Get** /libpod/volumes/{name}/exists | Volume exists
[**VolumeInspectLibpod**](VolumesApi.md#VolumeInspectLibpod) | **Get** /libpod/volumes/{name}/json | Inspect volume
[**VolumeListLibpod**](VolumesApi.md#VolumeListLibpod) | **Get** /libpod/volumes/json | List volumes
[**VolumePruneLibpod**](VolumesApi.md#VolumePruneLibpod) | **Post** /libpod/volumes/prune | Prune volumes

# **VolumeCreateLibpod**
> InlineResponse20017 VolumeCreateLibpod(ctx, optional)
Create a volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VolumesApiVolumeCreateLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumesApiVolumeCreateLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of VolumeCreate**](VolumeCreate.md)| attributes for creating a container | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-tar
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeDeleteLibpod**
> VolumeDeleteLibpod(ctx, name, optional)
Remove volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the volume | 
 **optional** | ***VolumesApiVolumeDeleteLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumesApiVolumeDeleteLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| force removal | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeExistsLibpod**
> VolumeExistsLibpod(ctx, name)
Volume exists

Check if a volume exists

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name of the volume | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeInspectLibpod**
> InlineResponse20017 VolumeInspectLibpod(ctx, name)
Inspect volume

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the volume | 

### Return type

[**InlineResponse20017**](inline_response_200_17.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumeListLibpod**
> []Volume VolumeListLibpod(ctx, optional)
List volumes

Returns a list of volumes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VolumesApiVolumeListLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumesApiVolumeListLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **optional.String**| JSON encoded value of the filters (a map[string][]string) to process on the volumes list. Available filters:   - driver&#x3D;&lt;volume-driver-name&gt; Matches volumes based on their driver.   - label&#x3D;&lt;key&gt; or label&#x3D;&lt;key&gt;:&lt;value&gt; Matches volumes based on the presence of a label alone or a label and a value.   - name&#x3D;&lt;volume-name&gt; Matches all of volume name.   - opt&#x3D;&lt;driver-option&gt; Matches a storage driver options  | 

### Return type

[**[]Volume**](Volume.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **VolumePruneLibpod**
> []PruneReport VolumePruneLibpod(ctx, optional)
Prune volumes

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***VolumesApiVolumePruneLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a VolumesApiVolumePruneLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **optional.String**| JSON encoded value of filters (a map[string][]string) to match volumes against before pruning. Available filters:  - label (label&#x3D;&lt;key&gt;, label&#x3D;&lt;key&gt;&#x3D;&lt;value&gt;, label!&#x3D;&lt;key&gt;, or label!&#x3D;&lt;key&gt;&#x3D;&lt;value&gt;) Prune volumes with (or without, in case label!&#x3D;... is used) the specified labels.  | 

### Return type

[**[]PruneReport**](PruneReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

