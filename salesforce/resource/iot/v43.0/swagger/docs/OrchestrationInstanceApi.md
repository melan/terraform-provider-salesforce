# \OrchestrationInstanceApi

All URIs are relative to *http://salesforce.com/services/data/v43.0/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteOrchestrationInstanceByKey**](OrchestrationInstanceApi.md#DeleteOrchestrationInstanceByKey) | **Delete** /orchestrations/{orchestrationId}/instances/{instanceKey} | Delete the orchestration instance corresponding to the specified orchestration and instance key.
[**GetOrchestrationInstanceByKey**](OrchestrationInstanceApi.md#GetOrchestrationInstanceByKey) | **Get** /orchestrations/{orchestrationId}/instances/{instanceKey} | Get the orchestration instance corresponding to the specified orchestration and instance key.
[**GetOrchestrationInstances**](OrchestrationInstanceApi.md#GetOrchestrationInstances) | **Get** /orchestrations/{orchestrationId}/instances | Get the orchestration instances in the specified orchestration.


# **DeleteOrchestrationInstanceByKey**
> DeleteOrchestrationInstanceByKey(ctx, orchestrationId, instanceKey)
Delete the orchestration instance corresponding to the specified orchestration and instance key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **orchestrationId** | **string**| A unique identifier for the orchestration instance. | 
  **instanceKey** | **string**| A key that uniquely identifies an orchestration instance. The instanceKey is the value of the event key field. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrchestrationInstanceByKey**
> OrchestrationInstanceDetailRepresentation GetOrchestrationInstanceByKey(ctx, orchestrationId, instanceKey)
Get the orchestration instance corresponding to the specified orchestration and instance key.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **orchestrationId** | **string**| A unique identifier for the orchestration instance. | 
  **instanceKey** | **string**| A key that uniquely identifies an orchestration instance. The instanceKey is the value of the event key field. | 

### Return type

[**OrchestrationInstanceDetailRepresentation**](OrchestrationInstanceDetailRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrchestrationInstances**
> OrchestrationInstanceCollectionRepresentation GetOrchestrationInstances(ctx, orchestrationId, optional)
Get the orchestration instances in the specified orchestration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **orchestrationId** | **string**| ID of orchestration for which to get the instances. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orchestrationId** | **string**| ID of orchestration for which to get the instances. | 
 **stateId** | **string**| ID of state for which to get orchestration instances. | 
 **page** | **string**| For internal use only. | 
 **pageSize** | **int32**| The number of orchestration instances returned per page. When one page is returned, it contains the URL in the nextPageUrl property that you use to get the next page of results. | 

### Return type

[**OrchestrationInstanceCollectionRepresentation**](OrchestrationInstanceCollectionRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

