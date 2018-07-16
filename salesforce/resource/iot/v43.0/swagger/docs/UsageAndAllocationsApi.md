# \UsageAndAllocationsApi

All URIs are relative to *http://salesforce.com/services/data/v43.0/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UsageGet**](UsageAndAllocationsApi.md#UsageGet) | **Get** /usage | Get IoT usage information for the entire org. Usage information includes allocations for events and orchestrations, the number of events processed and skipped across all orchestrations, and a link to usage information for orchestration instances.
[**UsageOrchestrationUsageGet**](UsageAndAllocationsApi.md#UsageOrchestrationUsageGet) | **Get** /usage/orchestration-usage | Get usage data for each orchestration in the org, such as the number of instances, and the number of processed and skipped events per orchestration.


# **UsageGet**
> UsageDetailRepresentation UsageGet(ctx, )
Get IoT usage information for the entire org. Usage information includes allocations for events and orchestrations, the number of events processed and skipped across all orchestrations, and a link to usage information for orchestration instances.



### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UsageDetailRepresentation**](UsageDetailRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UsageOrchestrationUsageGet**
> OrchestrationUsageCollectionRepresentation UsageOrchestrationUsageGet(ctx, optional)
Get usage data for each orchestration in the org, such as the number of instances, and the number of processed and skipped events per orchestration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **string**| For internal use only. | 
 **pageSize** | **int32**| The number of items returned per page. When one page is returned, it contains the URL in the nextPageUrl property that you use to get the next page of results. | 

### Return type

[**OrchestrationUsageCollectionRepresentation**](OrchestrationUsageCollectionRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

