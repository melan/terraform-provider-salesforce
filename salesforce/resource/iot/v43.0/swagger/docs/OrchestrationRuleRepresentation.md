# OrchestrationRuleRepresentation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** | A unique identifier for the rule. | [optional] [default to null]
**Description** | **string** | A rule description. | [optional] [default to null]
**Condition** | **string** | Specifies the condition to be evaluated. The Salesforce Expression Language is supported. | [optional] [default to null]
**Actions** | [**[]OrchestrationActionRepresentation**](OrchestrationActionRepresentation.md) | A list of actions to execute, such as assigning a variable value or creating a Salesforce record, when a condition is met. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


