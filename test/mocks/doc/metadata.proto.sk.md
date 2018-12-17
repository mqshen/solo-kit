<!-- Code generated by protoc-gen-solo-kit. DO NOT EDIT. -->

## Package:
core.solo.io

## Source File:
github.com/solo-io/solo-kit/api/v1/metadata.proto 

## Description:  

## Contents:
- Messages:  
	- [Metadata](#Metadata)  
	- [LabelsEntry](#LabelsEntry)  
	- [AnnotationsEntry](#AnnotationsEntry)

---
  
### <a name="Metadata">Metadata</a>

Description: *
Metadata contains general properties of resources for purposes of versioning, annotating, and namespacing.

```yaml
"name": string
"namespace": string
"resource_version": string
"labels": map<string, string>
"annotations": map<string, string>

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| name | string | Name of the resource.

Names must be unique and follow the following syntax rules:

One or more lowercase rfc1035/rfc1123 labels separated by '.' with a maximum length of 253 characters. |  |
| namespace | string | Namespace is used for the namespacing of resources. |  |
| resource_version | string | An opaque value that represents the internal version of this object that can be used by clients to determine when objects have changed. |  |
| labels | [map<string, string>](metadata.proto.sk.md#LabelsEntry) | Map of string keys and values that can be used to organize and categorize (scope and select) objects. Some resources contain `selectors` which can be linked with other resources by their labels |  |
| annotations | [map<string, string>](metadata.proto.sk.md#AnnotationsEntry) | Annotations is an unstructured key value map stored with a resource that may be set by external tools to store and retrieve arbitrary metadata. |  |
  
### <a name="LabelsEntry">LabelsEntry</a>

Description: 

```yaml
"key": string
"value": string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| key | string |  |  |
| value | string |  |  |
  
### <a name="AnnotationsEntry">AnnotationsEntry</a>

Description: 

```yaml
"key": string
"value": string

```

| Field | Type | Description | Default |
| ----- | ---- | ----------- |----------- | 
| key | string |  |  |
| value | string |  |  |


<!-- Start of HubSpot Embed Code -->
<script type="text/javascript" id="hs-script-loader" async defer src="//js.hs-scripts.com/5130874.js"></script>
<!-- End of HubSpot Embed Code -->