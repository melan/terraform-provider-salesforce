# \OrchestrationMetricsApi

All URIs are relative to *http://salesforce.com/services/data/v43.0/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrchestrationTransitionCounts**](OrchestrationMetricsApi.md#GetOrchestrationTransitionCounts) | **Get** /orchestrations/{orchestrationId}/metrics/transitions | Get state transition counts for the specified orchestration.


# **GetOrchestrationTransitionCounts**
> []OrchestrationTransitionCount GetOrchestrationTransitionCounts(ctx, orchestrationId)
Get state transition counts for the specified orchestration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **orchestrationId** | **string**| ID of the orchestration to get transition counts for. | 

### Return type

[**[]OrchestrationTransitionCount**](OrchestrationTransitionCount.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

