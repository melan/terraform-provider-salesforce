# ContextInputRepresentation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name that&#39;s used to refer to the context through the API. | [optional] [default to null]
**Label** | **string** | The context label that&#39;s displayed in the user interface. | [optional] [default to null]
**Description** | **string** | A description of the context. | [optional] [default to null]
**Runtime** | **string** | The Salesforce IoT runtime used. | [optional] [default to null]
**Events** | [**[]OrchestrationEventTypeInputRepresentation**](OrchestrationEventTypeInputRepresentation.md) | A list of one ore more events associated with this context. | [optional] [default to null]
**ReferenceData** | [**[]ContextDatasetInputRepresentation**](ContextDatasetInputRepresentation.md) | The Salesforce object associated with this context. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


