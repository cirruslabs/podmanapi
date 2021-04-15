# {{classname}}

All URIs are relative to *http://podman.io/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SecretCreateLibpod**](SecretsApi.md#SecretCreateLibpod) | **Post** /libpod/secrets/create | Create a secret
[**SecretDeleteLibpod**](SecretsApi.md#SecretDeleteLibpod) | **Delete** /libpod/secrets/{name} | Remove secret
[**SecretInspectLibpod**](SecretsApi.md#SecretInspectLibpod) | **Get** /libpod/secrets/{name}/json | Inspect secret
[**SecretListLibpod**](SecretsApi.md#SecretListLibpod) | **Get** /libpod/secrets/json | List secrets

# **SecretCreateLibpod**
> InlineResponse2011 SecretCreateLibpod(ctx, name, optional)
Create a secret

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| User-defined name of the secret. | 
 **optional** | ***SecretsApiSecretCreateLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretsApiSecretCreateLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of string**](string.md)| Secret | 
 **driver** | **optional.**| Secret driver | [default to file]

### Return type

[**InlineResponse2011**](inline_response_201_1.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json, application/x-tar
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SecretDeleteLibpod**
> SecretDeleteLibpod(ctx, name, optional)
Remove secret

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the secret | 
 **optional** | ***SecretsApiSecretDeleteLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SecretsApiSecretDeleteLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **all** | **optional.Bool**| Remove all secrets | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SecretInspectLibpod**
> SecretInfoReport SecretInspectLibpod(ctx, name)
Inspect secret

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the secret | 

### Return type

[**SecretInfoReport**](SecretInfoReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SecretListLibpod**
> []SecretInfoReport SecretListLibpod(ctx, )
List secrets

Returns a list of secrets

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]SecretInfoReport**](SecretInfoReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

