# \ActivityApi

All URIs are relative to *http://salesforce.com/services/data/v43.0/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivitiesGet**](ActivityApi.md#ActivitiesGet) | **Get** /activities | Get a list of all activities correlating to matching filter parameters.
[**ActivitiesIdGet**](ActivityApi.md#ActivitiesIdGet) | **Get** /activities/{id} | Get a single activity record matching the ID.


# **ActivitiesGet**
> Activity ActivitiesGet(ctx, optional)
Get a list of all activities correlating to matching filter parameters.

This endpoint will return a list of all the activities matching the filter parameters. Optional source record ids can be an array.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sourceIds** | [**[]string**](string.md)| IDs of the source objects to be retrieved. | 
 **search** | **string**| Single-word term search. | 
 **level** | **string**| Activity log level. | 
 **activityType** | **string**| Type of activity. | 
 **instanceKey** | **string**| Orchestration instance key. | 
 **startTime** | **time.Time**| Activity log start-time filter. | 
 **endTime** | **time.Time**| Activity log end-time filter. | 
 **pageSize** | **int32**| The number of items returned per page. When one page is returned, it contains the URL in the nextPageUrl property that you use to get the next page of results. | 

### Return type

[**Activity**](Activity.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ActivitiesIdGet**
> ActivityInfo ActivitiesIdGet(ctx, id)
Get a single activity record matching the ID.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **id** | **string**| ID for the log entry. | 

### Return type

[**ActivityInfo**](ActivityInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

