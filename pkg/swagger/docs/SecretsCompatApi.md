# {{classname}}

All URIs are relative to *http://podman.io/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecretCreate**](SecretsCompatApi.md#SecretCreate) | **Post** /secrets/create | Create a secret
[**SecretDelete**](SecretsCompatApi.md#SecretDelete) | **Delete** /secrets/{name} | Remove secret
[**SecretInspect**](SecretsCompatApi.md#SecretInspect) | **Get** /secrets/{name} | Inspect secret
[**SecretList**](SecretsCompatApi.md#SecretList) | **Get** /secrets | List secrets

# **SecretCreate**
> InlineResponse2011 SecretCreate(ctx, optional)
Create a secret

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SecretsCompatApiSecretCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretsCompatApiSecretCreateOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of SecretCreate**](SecretCreate.md)| attributes for creating a secret
 | 

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-tar
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SecretDelete**
> SecretDelete(ctx, name)
Remove secret

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the secret | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SecretInspect**
> SecretInfoReportCompat SecretInspect(ctx, name)
Inspect secret

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the secret | 

### Return type

[**SecretInfoReportCompat**](SecretInfoReportCompat.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SecretList**
> []SecretInfoReportCompat SecretList(ctx, )
List secrets

Returns a list of secrets

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]SecretInfoReportCompat**](SecretInfoReportCompat.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

