# \OrchestrationTrackerApi

All URIs are relative to *http://salesforce.com/services/data/v43.0/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTracker**](OrchestrationTrackerApi.md#CreateTracker) | **Post** /orchestrations/{orchestrationId}/trackers | Create a new tracker for an orchestration.
[**DeleteTracker**](OrchestrationTrackerApi.md#DeleteTracker) | **Delete** /orchestrations/{orchestrationId}/trackers/{trackerId} | Delete a tracker.
[**GetOrchestrationInstanceLog**](OrchestrationTrackerApi.md#GetOrchestrationInstanceLog) | **Get** /orchestrations/{orchestrationId}/instances/{instanceKey}/log | Get a log of orchestration events for a specific orchestration and key value.
[**GetOrchestrationTrackers**](OrchestrationTrackerApi.md#GetOrchestrationTrackers) | **Get** /orchestrations/{orchestrationId}/trackers | Get a list of trackers setup for this orchestration.
[**GetTrackerById**](OrchestrationTrackerApi.md#GetTrackerById) | **Get** /orchestrations/{orchestrationId}/trackers/{trackerId} | Get the specified tracker corresponding to the specified orchestration.
[**UpdateTracker**](OrchestrationTrackerApi.md#UpdateTracker) | **Patch** /orchestrations/{orchestrationId}/trackers/{trackerId} | Update a tracker.


# **CreateTracker**
> OrchestrationTrackerRepresentation CreateTracker(ctx, orchestrationId, body)
Create a new tracker for an orchestration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **orchestrationId** | **string**| ID of the orchestration for which to create a tracker. | 
  **body** | [**OrchestrationTrackerPostInputRepresentation**](OrchestrationTrackerPostInputRepresentation.md)| Tracker definition. | 

### Return type

[**OrchestrationTrackerRepresentation**](OrchestrationTrackerRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteTracker**
> DeleteTracker(ctx, orchestrationId, trackerId)
Delete a tracker.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **orchestrationId** | **string**| ID of the tracker&#39;s orchestration for which to delete the tracker. | 
  **trackerId** | **string**| ID of the tracker to delete. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrchestrationInstanceLog**
> OrchestrationLogEventCollectionRepresentation GetOrchestrationInstanceLog(ctx, orchestrationId, instanceKey, fromDate, toDate, optional)
Get a log of orchestration events for a specific orchestration and key value.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **orchestrationId** | **string**| Orchestration ID for which to get the log of events for a specific orchestration instance key value. | 
  **instanceKey** | **string**| Orchestration instance key value for which to get a log of orchestration events. | 
  **fromDate** | **string**| Get log events for which the timestamp is later than fromDate. | 
  **toDate** | **string**| Get log events for which the timestamp is prior to toDate | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orchestrationId** | **string**| Orchestration ID for which to get the log of events for a specific orchestration instance key value. | 
 **instanceKey** | **string**| Orchestration instance key value for which to get a log of orchestration events. | 
 **fromDate** | **string**| Get log events for which the timestamp is later than fromDate. | 
 **toDate** | **string**| Get log events for which the timestamp is prior to toDate | 
 **page** | **string**| For internal use only. | 
 **pageSize** | **int32**| The number of log events returned per page. When one page is returned, it contains the URL in the nextPageUrl property that you use to get the next page of results. | 

### Return type

[**OrchestrationLogEventCollectionRepresentation**](OrchestrationLogEventCollectionRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrchestrationTrackers**
> OrchestrationTrackerCollectionRepresentation GetOrchestrationTrackers(ctx, orchestrationId, optional)
Get a list of trackers setup for this orchestration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **orchestrationId** | **string**| ID of the orchestration for which to get trackers. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orchestrationId** | **string**| ID of the orchestration for which to get trackers. | 
 **page** | **string**| For internal use only. | 
 **pageSize** | **int32**| The number of trackers returned per page. When a page of the response is returned, it contains the nextPageUrl property which is the URL that you use to get the next page of results. | 

### Return type

[**OrchestrationTrackerCollectionRepresentation**](OrchestrationTrackerCollectionRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetTrackerById**
> OrchestrationTrackerRepresentation GetTrackerById(ctx, orchestrationId, trackerId)
Get the specified tracker corresponding to the specified orchestration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **orchestrationId** | **string**| ID of the tracker&#39;s orchestration. | 
  **trackerId** | **string**| ID of the tracker. | 

### Return type

[**OrchestrationTrackerRepresentation**](OrchestrationTrackerRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateTracker**
> OrchestrationTrackerRepresentation UpdateTracker(ctx, orchestrationId, trackerId, body)
Update a tracker.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **orchestrationId** | **string**| ID of a tracker&#39;s orchestration. | 
  **trackerId** | **string**| ID of a tracker. | 
  **body** | [**OrchestrationTrackerPatchInputRepresentation**](OrchestrationTrackerPatchInputRepresentation.md)| Tracker to update. | 

### Return type

[**OrchestrationTrackerRepresentation**](OrchestrationTrackerRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

