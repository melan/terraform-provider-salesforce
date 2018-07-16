# OrchestrationPropertiesRepresentation

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestartOnError** | **bool** | Set to true to create a new orchestration instance for an incoming event with a matching key after an error occurs. By default (false), incoming events are ignored after an error occurs in the matching orchestration instance. | [optional] [default to null]
**RunOnce** | **bool** | Set to true to ignore incoming events with a matching key for an exited orchestration instance. By default (false), incoming events with a matching key delete an exited orchestration instance and create a new instance. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


