# InlineResponse20017

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Anonymous** | **bool** | Anonymous indicates that the volume was created as an anonymous volume for a specific container, and will be be removed when any container using it is removed. | [optional] [default to null]
**CreatedAt** | [**time.Time**](time.Time.md) | CreatedAt is the date and time the volume was created at. This is not stored for older Libpod volumes; if so, it will be omitted. | [optional] [default to null]
**Driver** | **string** | Driver is the driver used to create the volume. If set to \&quot;local\&quot; or \&quot;\&quot;, the Local driver (Podman built-in code) is used to service the volume; otherwise, a volume plugin with the given name is used to mount and manage the volume. | [optional] [default to null]
**GID** | **int64** | GID is the GID that the volume was created with. | [optional] [default to null]
**Labels** | **map[string]string** | Labels includes the volume&#x27;s configured labels, key:value pairs that can be passed during volume creation to provide information for third party tools. | [optional] [default to null]
**Mountpoint** | **string** | Mountpoint is the path on the host where the volume is mounted. | [optional] [default to null]
**Name** | **string** | Name is the name of the volume. | [optional] [default to null]
**Options** | **map[string]string** | Options is a set of options that were used when creating the volume. For the Local driver, these are mount options that will be used to determine how a local filesystem is mounted; they are handled as parameters to Mount in a manner described in the volume create manpage. For non-local drivers, these are passed as-is to the volume plugin. | [optional] [default to null]
**Scope** | **string** | Scope is unused and provided solely for Docker compatibility. It is unconditionally set to \&quot;local\&quot;. | [optional] [default to null]
**Status** | [**map[string]interface{}**](interface{}.md) | Status is used to return information on the volume&#x27;s current state, if the volume was created using a volume plugin (uses a Driver that is not the local driver). Status is provided to us by an external program, so no guarantees are made about its format or contents. Further, it is an optional field, so it may not be set even in cases where a volume plugin is in use. | [optional] [default to null]
**UID** | **int64** | UID is the UID that the volume was created with. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

