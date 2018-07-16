# OrchestrationInstanceDetailRepresentation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceKey** | **string** | A key that uniquely identifies an orchestration instance. The instanceKey is the value of the event key field. | [optional] [default to null]
**Id** | **string** | A unique identifier for the orchestration instance. | [optional] [default to null]
**Url** | **string** | Resource URL to retrieve this orchestration instance. | [optional] [default to null]
**CreatedDate** | **string** |  | [optional] [default to null]
**CreatedBy** | [***UserReference**](UserReference.md) |  | [optional] [default to null]
**LastModifiedDate** | **string** |  | [optional] [default to null]
**LastModifiedBy** | [***UserReference**](UserReference.md) |  | [optional] [default to null]
**Description** | **string** | A description for the orchestration instance. | [optional] [default to null]
**Variables** | [**[]VariableRepresentation**](VariableRepresentation.md) | A list of variables that are part of the orchestration instance. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


