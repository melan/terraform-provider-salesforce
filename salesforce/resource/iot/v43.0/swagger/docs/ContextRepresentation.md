# ContextRepresentation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the context. | [optional] [default to null]
**Name** | **string** | A unique name that&#39;s used to refer to the context through the API. | [optional] [default to null]
**Namespace** | **string** | A namespace prefix used for packages. | [optional] [default to null]
**Label** | **string** | The context label that&#39;s displayed in the user interface. | [optional] [default to null]
**ManageableState** | **string** | An enum that denotes if the orchestration is part of a managed package, and the package state: released, deleted, deprecated, installed, beta, or unmanaged. | [optional] [default to null]
**Url** | **string** | Resource URL to retrieve this context. | [optional] [default to null]
**CreatedDate** | **string** |  | [optional] [default to null]
**CreatedBy** | [***UserReference**](UserReference.md) |  | [optional] [default to null]
**LastModifiedDate** | **string** |  | [optional] [default to null]
**LastModifiedBy** | [***UserReference**](UserReference.md) |  | [optional] [default to null]
**Description** | **string** | A context description. | [optional] [default to null]
**Status** | **string** |  | [optional] [default to null]
**Runtime** | **string** | The Salesforce IoT runtime used. | [optional] [default to null]
**Events** | [**[]OrchestrationEventType**](OrchestrationEventType.md) | A list of events and key fields associated with the context. | [optional] [default to null]
**ReferenceData** | [**[]ContextDataset**](ContextDataset.md) | The Salesforce object and key field associated with the context. | [optional] [default to null]
**OrchestrationsUrl** | **string** | Resource URL to retrieve all orchestrations using this context. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


