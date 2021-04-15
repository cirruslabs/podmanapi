# {{classname}}

All URIs are relative to *http://podman.io/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ImageBuildLibpod**](ImagesApi.md#ImageBuildLibpod) | **Post** /libpod/build | Create image
[**ImageChangesLibpod**](ImagesApi.md#ImageChangesLibpod) | **Get** /libpod/images/{name}/changes | Report on changes to images&#x27;s filesystem; adds, deletes or modifications.
[**ImageDeleteAllLibpod**](ImagesApi.md#ImageDeleteAllLibpod) | **Delete** /libpod/images/remove | Remove one or more images from the storage.
[**ImageDeleteLibpod**](ImagesApi.md#ImageDeleteLibpod) | **Delete** /libpod/images/{name} | Remove an image from the local storage.
[**ImageExistsLibpod**](ImagesApi.md#ImageExistsLibpod) | **Get** /libpod/images/{name}/exists | Image exists
[**ImageExportLibpod**](ImagesApi.md#ImageExportLibpod) | **Get** /libpod/images/export | Export multiple images
[**ImageGetLibpod**](ImagesApi.md#ImageGetLibpod) | **Get** /libpod/images/{name}/get | Export an image
[**ImageHistoryLibpod**](ImagesApi.md#ImageHistoryLibpod) | **Get** /libpod/images/{name}/history | History of an image
[**ImageImportLibpod**](ImagesApi.md#ImageImportLibpod) | **Post** /libpod/images/import | Import image
[**ImageInspectLibpod**](ImagesApi.md#ImageInspectLibpod) | **Get** /libpod/images/{name}/json | Inspect an image
[**ImageListLibpod**](ImagesApi.md#ImageListLibpod) | **Get** /libpod/images/json | List Images
[**ImageLoadLibpod**](ImagesApi.md#ImageLoadLibpod) | **Post** /libpod/images/load | Load image
[**ImagePruneLibpod**](ImagesApi.md#ImagePruneLibpod) | **Post** /libpod/images/prune | Prune unused images
[**ImagePullLibpod**](ImagesApi.md#ImagePullLibpod) | **Post** /libpod/images/pull | Pull images
[**ImagePushLibpod**](ImagesApi.md#ImagePushLibpod) | **Post** /libpod/images/{name}/push | Push Image
[**ImageSearchLibpod**](ImagesApi.md#ImageSearchLibpod) | **Get** /libpod/images/search | Search images
[**ImageTagLibpod**](ImagesApi.md#ImageTagLibpod) | **Post** /libpod/images/{name}/tag | Tag an image
[**ImageTreeLibpod**](ImagesApi.md#ImageTreeLibpod) | **Get** /libpod/images/{name}/tree | Image tree
[**ImageUntagLibpod**](ImagesApi.md#ImageUntagLibpod) | **Post** /libpod/images/{name}/untag | Untag an image

# **ImageBuildLibpod**
> InlineResponse2008 ImageBuildLibpod(ctx, optional)
Create image

Build an image from the given Dockerfile(s)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImagesApiImageBuildLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImagesApiImageBuildLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dockerfile** | **optional.String**| Path within the build context to the &#x60;Dockerfile&#x60;. This is ignored if remote is specified and points to an external &#x60;Dockerfile&#x60;.  | [default to Dockerfile]
 **t** | **optional.String**| A name and optional tag to apply to the image in the &#x60;name:tag&#x60; format.  If you omit the tag the default latest value is assumed. You can provide several t parameters. | [default to latest]
 **extrahosts** | **optional.String**| TBD Extra hosts to add to /etc/hosts (As of version 1.xx)  | 
 **remote** | **optional.String**| A Git repository URI or HTTP/HTTPS context URI. If the URI points to a single text file, the file’s contents are placed into a file called Dockerfile and the image is built from that file. If the URI points to a tarball, the file is downloaded by the daemon and the contents therein used as the context for the build. If the URI points to a tarball and the dockerfile parameter is also specified, there must be a file with the corresponding path inside the tarball. (As of version 1.xx)  | 
 **q** | **optional.Bool**| Suppress verbose build output  | [default to false]
 **nocache** | **optional.Bool**| Do not use the cache when building the image (As of version 1.xx)  | [default to false]
 **cachefrom** | **optional.String**| JSON array of images used to build cache resolution (As of version 1.xx)  | 
 **pull** | **optional.Bool**| Attempt to pull the image even if an older image exists locally (As of version 1.xx)  | [default to false]
 **rm** | **optional.Bool**| Remove intermediate containers after a successful build (As of version 1.xx)  | [default to true]
 **forcerm** | **optional.Bool**| Always remove intermediate containers, even upon failure (As of version 1.xx)  | [default to false]
 **memory** | **optional.Int32**| Memory is the upper limit (in bytes) on how much memory running containers can use (As of version 1.xx)  | 
 **memswap** | **optional.Int32**| MemorySwap limits the amount of memory and swap together (As of version 1.xx)  | 
 **cpushares** | **optional.Int32**| CPUShares (relative weight (As of version 1.xx)  | 
 **cpusetcpus** | **optional.String**| CPUSetCPUs in which to allow execution (0-3, 0,1) (As of version 1.xx)  | 
 **cpuperiod** | **optional.Int32**| CPUPeriod limits the CPU CFS (Completely Fair Scheduler) period (As of version 1.xx)  | 
 **cpuquota** | **optional.Int32**| CPUQuota limits the CPU CFS (Completely Fair Scheduler) quota (As of version 1.xx)  | 
 **buildargs** | **optional.String**| JSON map of string pairs denoting build-time variables. For example, the build argument &#x60;Foo&#x60; with the value of &#x60;bar&#x60; would be encoded in JSON as &#x60;[\&quot;Foo\&quot;:\&quot;bar\&quot;]&#x60;.  For example, buildargs&#x3D;{\&quot;Foo\&quot;:\&quot;bar\&quot;}.  Note(s): * This should not be used to pass secrets. * The value of buildargs should be URI component encoded before being passed to the API.  (As of version 1.xx)  | 
 **shmsize** | **optional.Int32**| ShmSize is the \&quot;size\&quot; value to use when mounting an shmfs on the container&#x27;s /dev/shm directory. Default is 64MB (As of version 1.xx)  | [default to 67108864]
 **squash** | **optional.Bool**| Silently ignored. Squash the resulting images layers into a single layer (As of version 1.xx)  | [default to false]
 **labels** | **optional.String**| JSON map of key, value pairs to set as labels on the new image (As of version 1.xx)  | 
 **networkmode** | **optional.String**| Sets the networking mode for the run commands during build. Supported standard values are:   * &#x60;bridge&#x60; limited to containers within a single host, port mapping required for external access   * &#x60;host&#x60; no isolation between host and containers on this network   * &#x60;none&#x60; disable all networking for this container   * container:&lt;nameOrID&gt; share networking with given container   ---All other values are assumed to be a custom network&#x27;s name (As of version 1.xx)  | [default to bridge]
 **platform** | **optional.String**| Platform format os[/arch[/variant]] (As of version 1.xx)  | 
 **target** | **optional.String**| Target build stage (As of version 1.xx)  | 
 **outputs** | **optional.String**| output configuration TBD (As of version 1.xx)  | 
 **httpproxy** | **optional.Bool**| Inject http proxy environment variables into container (As of version 2.0.0)  | 

### Return type

[**InlineResponse2008**](inline_response_200_8.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageChangesLibpod**
> ImageChangesLibpod(ctx, name)
Report on changes to images's filesystem; adds, deletes or modifications.

Returns which files in a images's filesystem have been added, deleted, or modified. The Kind of modification can be one of:  0: Modified 1: Added 2: Deleted 

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

# **ImageDeleteAllLibpod**
> LibpodImagesRemoveReport ImageDeleteAllLibpod(ctx, optional)
Remove one or more images from the storage.

Remove one or more images from the storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImagesApiImageDeleteAllLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImagesApiImageDeleteAllLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **images** | [**optional.Interface of []string**](string.md)| Images IDs or names to remove. | 
 **all** | **optional.Bool**| Remove all images. | [default to true]
 **force** | **optional.Bool**| Force image removal (including containers using the images). | 

### Return type

[**LibpodImagesRemoveReport**](LibpodImagesRemoveReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageDeleteLibpod**
> []ImageDeleteResponse ImageDeleteLibpod(ctx, name, optional)
Remove an image from the local storage.

Remove an image from the local storage.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| name or ID of image to remove | 
 **optional** | ***ImagesApiImageDeleteLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImagesApiImageDeleteLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **force** | **optional.Bool**| remove the image even if used by containers or has other tags | 

### Return type

[**[]ImageDeleteResponse**](ImageDeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageExistsLibpod**
> ImageExistsLibpod(ctx, name)
Image exists

Check if image exists in local store

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

# **ImageExportLibpod**
> *os.File ImageExportLibpod(ctx, optional)
Export multiple images

Export multiple images into a single object. Only `docker-archive` is currently supported.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImagesApiImageExportLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImagesApiImageExportLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **optional.String**| format for exported image (only docker-archive is supported) | 
 **references** | [**optional.Interface of []string**](string.md)| references to images to export | 
 **compress** | **optional.Bool**| use compression on image | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageGetLibpod**
> *os.File ImageGetLibpod(ctx, name, optional)
Export an image

Export an image

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ImagesApiImageGetLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImagesApiImageGetLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **format** | **optional.String**| format for exported image | 
 **compress** | **optional.Bool**| use compression on image | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/x-tar

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageHistoryLibpod**
> InlineResponse2005 ImageHistoryLibpod(ctx, name)
History of an image

Return parent layers of an image.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 

### Return type

[**InlineResponse2005**](inline_response_200_5.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageImportLibpod**
> ImageImportReport ImageImportLibpod(ctx, body, optional)
Import image

Import a previously exported tarball as an image.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| tarball for imported image | 
 **optional** | ***ImagesApiImageImportLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImagesApiImageImportLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **changes** | [**optional.Interface of []string**](string.md)| Apply the following possible instructions to the created image: CMD | ENTRYPOINT | ENV | EXPOSE | LABEL | STOPSIGNAL | USER | VOLUME | WORKDIR.  JSON encoded string | 
 **message** | **optional.**| Set commit message for imported image | 
 **reference** | **optional.**| Optional Name[:TAG] for the image | 
 **url** | **optional.**| Load image from the specified URL | 

### Return type

[**ImageImportReport**](ImageImportReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-tar
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageInspectLibpod**
> InlineResponse20011 ImageInspectLibpod(ctx, name)
Inspect an image

Obtain low-level information about an image

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 

### Return type

[**InlineResponse20011**](inline_response_200_11.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageListLibpod**
> []ImageSummary ImageListLibpod(ctx, optional)
List Images

Returns a list of images on the server

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImagesApiImageListLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImagesApiImageListLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **all** | **optional.Bool**| Show all images. Only images from a final layer (no children) are shown by default. | [default to false]
 **filters** | **optional.String**| A JSON encoded value of the filters (a &#x60;map[string][]string&#x60;) to process on the images list. Available filters: - &#x60;before&#x60;&#x3D;(&#x60;&lt;image-name&gt;[:&lt;tag&gt;]&#x60;,  &#x60;&lt;image id&gt;&#x60; or &#x60;&lt;image@digest&gt;&#x60;) - &#x60;dangling&#x3D;true&#x60; - &#x60;label&#x3D;key&#x60; or &#x60;label&#x3D;\&quot;key&#x3D;value\&quot;&#x60; of an image label - &#x60;reference&#x60;&#x3D;(&#x60;&lt;image-name&gt;[:&lt;tag&gt;]&#x60;) - &#x60;id&#x60;&#x3D;(&#x60;&lt;image-id&gt;&#x60;) - &#x60;since&#x60;&#x3D;(&#x60;&lt;image-name&gt;[:&lt;tag&gt;]&#x60;,  &#x60;&lt;image id&gt;&#x60; or &#x60;&lt;image@digest&gt;&#x60;)  | 

### Return type

[**[]ImageSummary**](ImageSummary.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageLoadLibpod**
> ImageLoadReport ImageLoadLibpod(ctx, body)
Load image

Load an image (oci-archive or docker-archive) stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**string**](string.md)| tarball of container image | 

### Return type

[**ImageLoadReport**](ImageLoadReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-tar
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImagePruneLibpod**
> []ImageDeleteResponse ImagePruneLibpod(ctx, optional)
Prune unused images

Remove images that are not being used by a container

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImagesApiImagePruneLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImagesApiImagePruneLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filters** | **optional.String**| filters to apply to image pruning, encoded as JSON (map[string][]string). Available filters:   - &#x60;dangling&#x3D;&lt;boolean&gt;&#x60; When set to &#x60;true&#x60; (or &#x60;1&#x60;), prune only      unused *and* untagged images. When set to &#x60;false&#x60;      (or &#x60;0&#x60;), all unused images are pruned.   - &#x60;until&#x3D;&lt;string&gt;&#x60; Prune images created before this timestamp. The &#x60;&lt;timestamp&gt;&#x60; can be Unix timestamps, date formatted timestamps, or Go duration strings (e.g. &#x60;10m&#x60;, &#x60;1h30m&#x60;) computed relative to the daemon machine’s time.   - &#x60;label&#x60; (&#x60;label&#x3D;&lt;key&gt;&#x60;, &#x60;label&#x3D;&lt;key&gt;&#x3D;&lt;value&gt;&#x60;, &#x60;label!&#x3D;&lt;key&gt;&#x60;, or &#x60;label!&#x3D;&lt;key&gt;&#x3D;&lt;value&gt;&#x60;) Prune images with (or without, in case &#x60;label!&#x3D;...&#x60; is used) the specified labels.  | 

### Return type

[**[]ImageDeleteResponse**](ImageDeleteResponse.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImagePullLibpod**
> LibpodImagesPullReport ImagePullLibpod(ctx, optional)
Pull images

Pull one or more images from a container registry.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImagesApiImagePullLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImagesApiImagePullLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reference** | **optional.String**| Mandatory reference to the image (e.g., quay.io/image/name:tag) | 
 **credentials** | **optional.String**| username:password for the registry | 
 **arch** | **optional.String**| Pull image for the specified architecture. | 
 **oS** | **optional.String**| Pull image for the specified operating system. | 
 **variant** | **optional.String**| Pull image for the specified variant. | 
 **tlsVerify** | **optional.Bool**| Require TLS verification. | [default to true]
 **allTags** | **optional.Bool**| Pull all tagged images in the repository. | 

### Return type

[**LibpodImagesPullReport**](LibpodImagesPullReport.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImagePushLibpod**
> *os.File ImagePushLibpod(ctx, name, optional)
Push Image

Push an image to a container registry

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| Name of image to push. | 
 **optional** | ***ImagesApiImagePushLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImagesApiImagePushLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **destination** | **optional.String**| Allows for pushing the image to a different destination than the image refers to. | 
 **tlsVerify** | **optional.Bool**| Require TLS verification. | [default to true]
 **xRegistryAuth** | **optional.String**| A base64-encoded auth configuration. | 

### Return type

[***os.File**](*os.File.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageSearchLibpod**
> InlineResponse2007 ImageSearchLibpod(ctx, optional)
Search images

Search registries for images

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ImagesApiImageSearchLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImagesApiImageSearchLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **term** | **optional.String**| term to search | 
 **limit** | **optional.Int32**| maximum number of results | 
 **noTrunc** | **optional.Bool**| do not truncate any of the result strings | 
 **filters** | **optional.String**| A JSON encoded value of the filters (a &#x60;map[string][]string&#x60;) to process on the images list. Available filters: - &#x60;is-automated&#x3D;(true|false)&#x60; - &#x60;is-official&#x3D;(true|false)&#x60; - &#x60;stars&#x3D;&lt;number&gt;&#x60; Matches images that has at least &#x27;number&#x27; stars.  | 

### Return type

[**InlineResponse2007**](inline_response_200_7.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageTagLibpod**
> ImageTagLibpod(ctx, name, optional)
Tag an image

Tag an image so that it becomes part of a repository.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ImagesApiImageTagLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImagesApiImageTagLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repo** | **optional.String**| the repository to tag in | 
 **tag** | **optional.String**| the name of the new tag | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageTreeLibpod**
> InlineResponse20012 ImageTreeLibpod(ctx, name, optional)
Image tree

Retrieve the image tree for the provided image name or ID

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ImagesApiImageTreeLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImagesApiImageTreeLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **whatrequires** | **optional.Bool**| show all child images and layers of the specified image | 

### Return type

[**InlineResponse20012**](inline_response_200_12.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImageUntagLibpod**
> ImageUntagLibpod(ctx, name, optional)
Untag an image

Untag an image. If not repo and tag are specified, all tags are removed from the image.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| the name or ID of the container | 
 **optional** | ***ImagesApiImageUntagLibpodOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ImagesApiImageUntagLibpodOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **repo** | **optional.String**| the repository to untag | 
 **tag** | **optional.String**| the name of the tag to untag | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

