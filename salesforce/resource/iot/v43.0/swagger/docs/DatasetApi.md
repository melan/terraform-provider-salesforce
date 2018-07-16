# \DatasetApi

All URIs are relative to *http://salesforce.com/services/data/v43.0/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDatasetById**](DatasetApi.md#GetDatasetById) | **Get** /datasets/{name} | Get the specified dataset.
[**GetDatasets**](DatasetApi.md#GetDatasets) | **Get** /datasets | List all datasets that have been defined in your org.


# **GetDatasetById**
> OrchestrationDatasetDetailRepresentation GetDatasetById(ctx, name)
Get the specified dataset.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **name** | **string**| Unique name of the dataset to retrieve. | 

### Return type

[**OrchestrationDatasetDetailRepresentation**](OrchestrationDatasetDetailRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDatasets**
> OrchestrationDatasetCollectionRepresentation GetDatasets(ctx, optional)
List all datasets that have been defined in your org.



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
 **pageSize** | **int32**| The number of datasets returned per page. When one page is returned, it contains the URL in the nextPageUrl property that you use to get the next page of results. | 

### Return type

[**OrchestrationDatasetCollectionRepresentation**](OrchestrationDatasetCollectionRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

