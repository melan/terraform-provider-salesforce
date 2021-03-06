# Go API client for swagger

Use the Salesforce IoT REST API to create and manage orchestrations and contexts, and retrieve usage data.

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 43.0.0
- Package version: 1.0.0
- Build package: io.swagger.codegen.languages.GoClientCodegen

## Installation
Put the package under your project folder and add the following in import:
```
    "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *http://salesforce.com/services/data/v43.0/iot*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*ActivityApi* | [**ActivitiesGet**](docs/ActivityApi.md#activitiesget) | **Get** /activities | Get a list of all activities correlating to matching filter parameters.
*ActivityApi* | [**ActivitiesIdGet**](docs/ActivityApi.md#activitiesidget) | **Get** /activities/{id} | Get a single activity record matching the ID.
*DatasetApi* | [**GetDatasetById**](docs/DatasetApi.md#getdatasetbyid) | **Get** /datasets/{name} | Get the specified dataset.
*DatasetApi* | [**GetDatasets**](docs/DatasetApi.md#getdatasets) | **Get** /datasets | List all datasets that have been defined in your org.
*OrchestrationApi* | [**CreateOrchestration**](docs/OrchestrationApi.md#createorchestration) | **Post** /orchestrations | Create an orchestration.
*OrchestrationApi* | [**CreateOrchestrationActivation**](docs/OrchestrationApi.md#createorchestrationactivation) | **Post** /orchestrations/{orchestrationId}/activations | Activate the specified orchestration, optionally using activation options.
*OrchestrationApi* | [**DeleteOrchestration**](docs/OrchestrationApi.md#deleteorchestration) | **Delete** /orchestrations/{orchestrationId} | Delete the specified orchestration.
*OrchestrationApi* | [**GetLastOrchestrationActivation**](docs/OrchestrationApi.md#getlastorchestrationactivation) | **Get** /orchestrations/{orchestrationId}/activations/last | Get the last activation for the specified orchestration.
*OrchestrationApi* | [**GetOrchestrationActivationById**](docs/OrchestrationApi.md#getorchestrationactivationbyid) | **Get** /orchestrations/{orchestrationId}/activations/{activationId} | Get activation information for the specified orchestration and activation.
*OrchestrationApi* | [**GetOrchestrationActivations**](docs/OrchestrationApi.md#getorchestrationactivations) | **Get** /orchestrations/{orchestrationId}/activations | Get the activation history for the specified orchestration. The activations are sorted by creation date in descending order.
*OrchestrationApi* | [**GetOrchestrationById**](docs/OrchestrationApi.md#getorchestrationbyid) | **Get** /orchestrations/{orchestrationId} | Get the specified orchestration.
*OrchestrationApi* | [**GetOrchestrationStates**](docs/OrchestrationApi.md#getorchestrationstates) | **Get** /orchestrations/{orchestrationId}/states | Get the states in this orchestration.
*OrchestrationApi* | [**GetOrchestrations**](docs/OrchestrationApi.md#getorchestrations) | **Get** /orchestrations | List all orchestrations that have been defined in your org.
*OrchestrationApi* | [**UpdateLastOrchestrationActivation**](docs/OrchestrationApi.md#updatelastorchestrationactivation) | **Patch** /orchestrations/{orchestrationId}/activations/last | Deactivate the specified orchestration. Can only be used to deactivate an active orchestration.
*OrchestrationApi* | [**UpdateOrchestrationActivation**](docs/OrchestrationApi.md#updateorchestrationactivation) | **Patch** /orchestrations/{orchestrationId}/activations/{activationId} | Deactivate the specified orchestration activation. No other updates are supported.
*OrchestrationApi* | [**UpdateOrchestrationWithForm**](docs/OrchestrationApi.md#updateorchestrationwithform) | **Put** /orchestrations/{orchestrationId} | Update the specified orchestration using the full orchestration definition.
*OrchestrationContextApi* | [**CreateContext**](docs/OrchestrationContextApi.md#createcontext) | **Post** /contexts | Create a context.
*OrchestrationContextApi* | [**CreateContextActivation**](docs/OrchestrationContextApi.md#createcontextactivation) | **Post** /contexts/{contextId}/activations | Activate a context.
*OrchestrationContextApi* | [**DeleteContext**](docs/OrchestrationContextApi.md#deletecontext) | **Delete** /contexts/{contextId} | Delete the specified context.
*OrchestrationContextApi* | [**GetContextActivationById**](docs/OrchestrationContextApi.md#getcontextactivationbyid) | **Get** /contexts/{contextId}/activations/{activationId} | Get the specified context activation.
*OrchestrationContextApi* | [**GetContextActivations**](docs/OrchestrationContextApi.md#getcontextactivations) | **Get** /contexts/{contextId}/activations | List of all activations for this context.
*OrchestrationContextApi* | [**GetContextById**](docs/OrchestrationContextApi.md#getcontextbyid) | **Get** /contexts/{contextId} | Get the specified context.
*OrchestrationContextApi* | [**GetContexts**](docs/OrchestrationContextApi.md#getcontexts) | **Get** /contexts | List all contexts that have been defined in your org.
*OrchestrationContextApi* | [**GetLastContextActivation**](docs/OrchestrationContextApi.md#getlastcontextactivation) | **Get** /contexts/{contextId}/activations/last | Get the last context activation.
*OrchestrationContextApi* | [**UpdateContext**](docs/OrchestrationContextApi.md#updatecontext) | **Put** /contexts/{contextId} | Update the specified context using the specified full context definition.
*OrchestrationInstanceApi* | [**DeleteOrchestrationInstanceByKey**](docs/OrchestrationInstanceApi.md#deleteorchestrationinstancebykey) | **Delete** /orchestrations/{orchestrationId}/instances/{instanceKey} | Delete the orchestration instance corresponding to the specified orchestration and instance key.
*OrchestrationInstanceApi* | [**GetOrchestrationInstanceByKey**](docs/OrchestrationInstanceApi.md#getorchestrationinstancebykey) | **Get** /orchestrations/{orchestrationId}/instances/{instanceKey} | Get the orchestration instance corresponding to the specified orchestration and instance key.
*OrchestrationInstanceApi* | [**GetOrchestrationInstances**](docs/OrchestrationInstanceApi.md#getorchestrationinstances) | **Get** /orchestrations/{orchestrationId}/instances | Get the orchestration instances in the specified orchestration.
*OrchestrationMetricsApi* | [**GetOrchestrationTransitionCounts**](docs/OrchestrationMetricsApi.md#getorchestrationtransitioncounts) | **Get** /orchestrations/{orchestrationId}/metrics/transitions | Get state transition counts for the specified orchestration.
*OrchestrationTrackerApi* | [**CreateTracker**](docs/OrchestrationTrackerApi.md#createtracker) | **Post** /orchestrations/{orchestrationId}/trackers | Create a new tracker for an orchestration.
*OrchestrationTrackerApi* | [**DeleteTracker**](docs/OrchestrationTrackerApi.md#deletetracker) | **Delete** /orchestrations/{orchestrationId}/trackers/{trackerId} | Delete a tracker.
*OrchestrationTrackerApi* | [**GetOrchestrationInstanceLog**](docs/OrchestrationTrackerApi.md#getorchestrationinstancelog) | **Get** /orchestrations/{orchestrationId}/instances/{instanceKey}/log | Get a log of orchestration events for a specific orchestration and key value.
*OrchestrationTrackerApi* | [**GetOrchestrationTrackers**](docs/OrchestrationTrackerApi.md#getorchestrationtrackers) | **Get** /orchestrations/{orchestrationId}/trackers | Get a list of trackers setup for this orchestration.
*OrchestrationTrackerApi* | [**GetTrackerById**](docs/OrchestrationTrackerApi.md#gettrackerbyid) | **Get** /orchestrations/{orchestrationId}/trackers/{trackerId} | Get the specified tracker corresponding to the specified orchestration.
*OrchestrationTrackerApi* | [**UpdateTracker**](docs/OrchestrationTrackerApi.md#updatetracker) | **Patch** /orchestrations/{orchestrationId}/trackers/{trackerId} | Update a tracker.
*UsageAndAllocationsApi* | [**UsageGet**](docs/UsageAndAllocationsApi.md#usageget) | **Get** /usage | Get IoT usage information for the entire org. Usage information includes allocations for events and orchestrations, the number of events processed and skipped across all orchestrations, and a link to usage information for orchestration instances.
*UsageAndAllocationsApi* | [**UsageOrchestrationUsageGet**](docs/UsageAndAllocationsApi.md#usageorchestrationusageget) | **Get** /usage/orchestration-usage | Get usage data for each orchestration in the org, such as the number of instances, and the number of processed and skipped events per orchestration.


## Documentation For Models

 - [Activity](docs/Activity.md)
 - [ActivityInfo](docs/ActivityInfo.md)
 - [AllocationCountRepresentation](docs/AllocationCountRepresentation.md)
 - [AllocationsRepresentation](docs/AllocationsRepresentation.md)
 - [ContextActivationCollectionRepresentation](docs/ContextActivationCollectionRepresentation.md)
 - [ContextActivationInputRepresentation](docs/ContextActivationInputRepresentation.md)
 - [ContextActivationOrchestrationResultRepresentation](docs/ContextActivationOrchestrationResultRepresentation.md)
 - [ContextActivationRepresentation](docs/ContextActivationRepresentation.md)
 - [ContextActivationResultRepresentation](docs/ContextActivationResultRepresentation.md)
 - [ContextActivationSummaryRepresentation](docs/ContextActivationSummaryRepresentation.md)
 - [ContextCollectionRepresentation](docs/ContextCollectionRepresentation.md)
 - [ContextDataset](docs/ContextDataset.md)
 - [ContextDatasetInputRepresentation](docs/ContextDatasetInputRepresentation.md)
 - [ContextDefinitionRepresentation](docs/ContextDefinitionRepresentation.md)
 - [ContextInputRepresentation](docs/ContextInputRepresentation.md)
 - [ContextReferenceInputRepresentation](docs/ContextReferenceInputRepresentation.md)
 - [ContextReferenceRepresentation](docs/ContextReferenceRepresentation.md)
 - [ContextRepresentation](docs/ContextRepresentation.md)
 - [ContextSummaryRepresentation](docs/ContextSummaryRepresentation.md)
 - [DatasetReference](docs/DatasetReference.md)
 - [DatasetReferenceInputRepresentation](docs/DatasetReferenceInputRepresentation.md)
 - [Dependencies](docs/Dependencies.md)
 - [ErrorRepresentation](docs/ErrorRepresentation.md)
 - [EventTypeReference](docs/EventTypeReference.md)
 - [EventTypeReferenceInputRepresentation](docs/EventTypeReferenceInputRepresentation.md)
 - [EventUsageRepresentation](docs/EventUsageRepresentation.md)
 - [InstanceActivityInfo](docs/InstanceActivityInfo.md)
 - [MetadataActivityInfo](docs/MetadataActivityInfo.md)
 - [NotificationActivityInfo](docs/NotificationActivityInfo.md)
 - [OrchestrationActionRepresentation](docs/OrchestrationActionRepresentation.md)
 - [OrchestrationActivationCollectionRepresentation](docs/OrchestrationActivationCollectionRepresentation.md)
 - [OrchestrationActivationInputRepresentation](docs/OrchestrationActivationInputRepresentation.md)
 - [OrchestrationActivationOptions](docs/OrchestrationActivationOptions.md)
 - [OrchestrationActivationRepresentation](docs/OrchestrationActivationRepresentation.md)
 - [OrchestrationActivationSummaryRepresentation](docs/OrchestrationActivationSummaryRepresentation.md)
 - [OrchestrationActivationValidationFailureRepresentation](docs/OrchestrationActivationValidationFailureRepresentation.md)
 - [OrchestrationActivationValidationFailureRepresentationOutput](docs/OrchestrationActivationValidationFailureRepresentationOutput.md)
 - [OrchestrationAssignVariableActionRepresentation](docs/OrchestrationAssignVariableActionRepresentation.md)
 - [OrchestrationChangeStateActionRepresentation](docs/OrchestrationChangeStateActionRepresentation.md)
 - [OrchestrationCollectionRepresentation](docs/OrchestrationCollectionRepresentation.md)
 - [OrchestrationCountAggregationVariableValueRepresentation](docs/OrchestrationCountAggregationVariableValueRepresentation.md)
 - [OrchestrationCustomVariableValueRepresentation](docs/OrchestrationCustomVariableValueRepresentation.md)
 - [OrchestrationDatasetCollectionRepresentation](docs/OrchestrationDatasetCollectionRepresentation.md)
 - [OrchestrationDatasetDefinitionFieldRepresentation](docs/OrchestrationDatasetDefinitionFieldRepresentation.md)
 - [OrchestrationDatasetDefinitionRepresentation](docs/OrchestrationDatasetDefinitionRepresentation.md)
 - [OrchestrationDatasetDetailRepresentation](docs/OrchestrationDatasetDetailRepresentation.md)
 - [OrchestrationDatasetSummaryRepresentation](docs/OrchestrationDatasetSummaryRepresentation.md)
 - [OrchestrationDefinitionLocation](docs/OrchestrationDefinitionLocation.md)
 - [OrchestrationDefinitionRepresentation](docs/OrchestrationDefinitionRepresentation.md)
 - [OrchestrationEventFieldVariableValueRepresentation](docs/OrchestrationEventFieldVariableValueRepresentation.md)
 - [OrchestrationEventType](docs/OrchestrationEventType.md)
 - [OrchestrationEventTypeInputRepresentation](docs/OrchestrationEventTypeInputRepresentation.md)
 - [OrchestrationInputRepresentation](docs/OrchestrationInputRepresentation.md)
 - [OrchestrationInstanceCollectionRepresentation](docs/OrchestrationInstanceCollectionRepresentation.md)
 - [OrchestrationInstanceDetailRepresentation](docs/OrchestrationInstanceDetailRepresentation.md)
 - [OrchestrationInstanceKey](docs/OrchestrationInstanceKey.md)
 - [OrchestrationInstanceLifecycleStatus](docs/OrchestrationInstanceLifecycleStatus.md)
 - [OrchestrationInstanceSummaryRepresentation](docs/OrchestrationInstanceSummaryRepresentation.md)
 - [OrchestrationLogEventCollectionRepresentation](docs/OrchestrationLogEventCollectionRepresentation.md)
 - [OrchestrationLogEventInputEventKind](docs/OrchestrationLogEventInputEventKind.md)
 - [OrchestrationLogEventInputEventProcessingResult](docs/OrchestrationLogEventInputEventProcessingResult.md)
 - [OrchestrationLogEventRepresentation](docs/OrchestrationLogEventRepresentation.md)
 - [OrchestrationLogEventTerminationKind](docs/OrchestrationLogEventTerminationKind.md)
 - [OrchestrationOutputActionRepresentation](docs/OrchestrationOutputActionRepresentation.md)
 - [OrchestrationPropertiesRepresentation](docs/OrchestrationPropertiesRepresentation.md)
 - [OrchestrationRepresentation](docs/OrchestrationRepresentation.md)
 - [OrchestrationResetVariableActionRepresentation](docs/OrchestrationResetVariableActionRepresentation.md)
 - [OrchestrationRuleRepresentation](docs/OrchestrationRuleRepresentation.md)
 - [OrchestrationRuleThrottleRepresentation](docs/OrchestrationRuleThrottleRepresentation.md)
 - [OrchestrationRuleThrottleStateRepresentation](docs/OrchestrationRuleThrottleStateRepresentation.md)
 - [OrchestrationRuleWhenCustomRepresentation](docs/OrchestrationRuleWhenCustomRepresentation.md)
 - [OrchestrationRuleWhenEventRepresentation](docs/OrchestrationRuleWhenEventRepresentation.md)
 - [OrchestrationRuleWhenRepresentation](docs/OrchestrationRuleWhenRepresentation.md)
 - [OrchestrationRuleWhenSimpleRepresentation](docs/OrchestrationRuleWhenSimpleRepresentation.md)
 - [OrchestrationRuleWhenTemporalRepresentation](docs/OrchestrationRuleWhenTemporalRepresentation.md)
 - [OrchestrationSalesforceOutputActionField](docs/OrchestrationSalesforceOutputActionField.md)
 - [OrchestrationSalesforceOutputActionRepresentation](docs/OrchestrationSalesforceOutputActionRepresentation.md)
 - [OrchestrationStateCollectionRepresentation](docs/OrchestrationStateCollectionRepresentation.md)
 - [OrchestrationStateReferenceRepresentation](docs/OrchestrationStateReferenceRepresentation.md)
 - [OrchestrationStateRepresentation](docs/OrchestrationStateRepresentation.md)
 - [OrchestrationStateSummaryRepresentation](docs/OrchestrationStateSummaryRepresentation.md)
 - [OrchestrationSummaryRepresentation](docs/OrchestrationSummaryRepresentation.md)
 - [OrchestrationTemporalAggregationVariableValueRepresentation](docs/OrchestrationTemporalAggregationVariableValueRepresentation.md)
 - [OrchestrationTerminateActionRepresentation](docs/OrchestrationTerminateActionRepresentation.md)
 - [OrchestrationTrackerCollectionRepresentation](docs/OrchestrationTrackerCollectionRepresentation.md)
 - [OrchestrationTrackerPatchInputRepresentation](docs/OrchestrationTrackerPatchInputRepresentation.md)
 - [OrchestrationTrackerPostInputRepresentation](docs/OrchestrationTrackerPostInputRepresentation.md)
 - [OrchestrationTrackerRepresentation](docs/OrchestrationTrackerRepresentation.md)
 - [OrchestrationTransitionCount](docs/OrchestrationTransitionCount.md)
 - [OrchestrationUsageCollectionRepresentation](docs/OrchestrationUsageCollectionRepresentation.md)
 - [OrchestrationUsageRepresentation](docs/OrchestrationUsageRepresentation.md)
 - [OrchestrationValidationMessageRepresentation](docs/OrchestrationValidationMessageRepresentation.md)
 - [OrchestrationValidationMessagesRepresentation](docs/OrchestrationValidationMessagesRepresentation.md)
 - [OrchestrationVariableRepresentation](docs/OrchestrationVariableRepresentation.md)
 - [OrchestrationVariableValueRepresentation](docs/OrchestrationVariableValueRepresentation.md)
 - [SalesforceExpression](docs/SalesforceExpression.md)
 - [SalesforceTemporalExpression](docs/SalesforceTemporalExpression.md)
 - [Source](docs/Source.md)
 - [StateReferenceRepresentation](docs/StateReferenceRepresentation.md)
 - [UsageDetailRepresentation](docs/UsageDetailRepresentation.md)
 - [UserReference](docs/UserReference.md)
 - [VariableRepresentation](docs/VariableRepresentation.md)
 - [OrchestrationLogBeginEventRepresentation](docs/OrchestrationLogBeginEventRepresentation.md)
 - [OrchestrationLogEndEventRepresentation](docs/OrchestrationLogEndEventRepresentation.md)
 - [OrchestrationLogErrorEventRepresentation](docs/OrchestrationLogErrorEventRepresentation.md)
 - [OrchestrationLogFullStateEventRepresentation](docs/OrchestrationLogFullStateEventRepresentation.md)
 - [OrchestrationLogInfoEventRepresentation](docs/OrchestrationLogInfoEventRepresentation.md)
 - [OrchestrationLogInputEventRepresentation](docs/OrchestrationLogInputEventRepresentation.md)
 - [OrchestrationLogOutputEventRepresentation](docs/OrchestrationLogOutputEventRepresentation.md)
 - [OrchestrationLogRuleConditionEventRepresentation](docs/OrchestrationLogRuleConditionEventRepresentation.md)
 - [OrchestrationLogTerminateEventRepresentation](docs/OrchestrationLogTerminateEventRepresentation.md)
 - [OrchestrationLogThrottleEventRepresentation](docs/OrchestrationLogThrottleEventRepresentation.md)
 - [OrchestrationLogTimerArmEventRepresentation](docs/OrchestrationLogTimerArmEventRepresentation.md)
 - [OrchestrationLogTimerUnarmEventRepresentation](docs/OrchestrationLogTimerUnarmEventRepresentation.md)
 - [OrchestrationLogTransitionEventRepresentation](docs/OrchestrationLogTransitionEventRepresentation.md)
 - [OrchestrationLogVariableChangeEventRepresentation](docs/OrchestrationLogVariableChangeEventRepresentation.md)


## Documentation For Authorization
 Endpoints do not require authorization.


## Author

support@salesforce.com

