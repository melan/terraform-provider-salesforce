# \OrchestrationContextApi

All URIs are relative to *http://salesforce.com/services/data/v43.0/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateContext**](OrchestrationContextApi.md#CreateContext) | **Post** /contexts | Create a context.
[**CreateContextActivation**](OrchestrationContextApi.md#CreateContextActivation) | **Post** /contexts/{contextId}/activations | Activate a context.
[**DeleteContext**](OrchestrationContextApi.md#DeleteContext) | **Delete** /contexts/{contextId} | Delete the specified context.
[**GetContextActivationById**](OrchestrationContextApi.md#GetContextActivationById) | **Get** /contexts/{contextId}/activations/{activationId} | Get the specified context activation.
[**GetContextActivations**](OrchestrationContextApi.md#GetContextActivations) | **Get** /contexts/{contextId}/activations | List of all activations for this context.
[**GetContextById**](OrchestrationContextApi.md#GetContextById) | **Get** /contexts/{contextId} | Get the specified context.
[**GetContexts**](OrchestrationContextApi.md#GetContexts) | **Get** /contexts | List all contexts that have been defined in your org.
[**GetLastContextActivation**](OrchestrationContextApi.md#GetLastContextActivation) | **Get** /contexts/{contextId}/activations/last | Get the last context activation.
[**UpdateContext**](OrchestrationContextApi.md#UpdateContext) | **Put** /contexts/{contextId} | Update the specified context using the specified full context definition.


# **CreateContext**
> ContextRepresentation CreateContext(ctx, body)
Create a context.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**ContextInputRepresentation**](ContextInputRepresentation.md)| Context definition. | 

### Return type

[**ContextRepresentation**](ContextRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateContextActivation**
> ContextActivationRepresentation CreateContextActivation(ctx, contextId, body)
Activate a context.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **contextId** | **string**| ID of the context to activate. | 
  **body** | [**ContextActivationInputRepresentation**](ContextActivationInputRepresentation.md)| Context activation definition. | 

### Return type

[**ContextActivationRepresentation**](ContextActivationRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteContext**
> DeleteContext(ctx, contextId)
Delete the specified context.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **contextId** | **string**| ID of the context to delete. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContextActivationById**
> ContextActivationRepresentation GetContextActivationById(ctx, contextId, activationId)
Get the specified context activation.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **contextId** | **string**| ID of the context for the context activation to retrieve. | 
  **activationId** | **string**| ID of context activation to retrieve. | 

### Return type

[**ContextActivationRepresentation**](ContextActivationRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContextActivations**
> ContextActivationCollectionRepresentation GetContextActivations(ctx, contextId, optional)
List of all activations for this context.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **contextId** | **string**| ID of the context activation record to retrieve. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **contextId** | **string**| ID of the context activation record to retrieve. | 
 **page** | **string**| For internal use only. | 
 **pageSize** | **int32**| The number of context activations returned per page. When one page is returned, it contains the URL in the nextPageUrl property that you use to get the next page of results. | 

### Return type

[**ContextActivationCollectionRepresentation**](ContextActivationCollectionRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContextById**
> ContextRepresentation GetContextById(ctx, contextId)
Get the specified context.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **contextId** | **string**| ID of the context to retrieve. | 

### Return type

[**ContextRepresentation**](ContextRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetContexts**
> ContextCollectionRepresentation GetContexts(ctx, optional)
List all contexts that have been defined in your org.



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
 **pageSize** | **int32**| The number of contexts returned per page. When one page is returned, it contains the URL in the nextPageUrl property that you use to get the next page of results. | 
 **runtime** | **string**| The Salesforce IoT runtime used. | 
 **sortDirection** | **string**| The sort direction, either ascending or descending. | 
 **sort** | **string**| The field by which the contexts are sorted. | 

### Return type

[**ContextCollectionRepresentation**](ContextCollectionRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLastContextActivation**
> ContextActivationRepresentation GetLastContextActivation(ctx, contextId)
Get the last context activation.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **contextId** | **string**| ID of the context activation record to retrieve. | 

### Return type

[**ContextActivationRepresentation**](ContextActivationRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateContext**
> ContextRepresentation UpdateContext(ctx, contextId, body)
Update the specified context using the specified full context definition.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **contextId** | **string**| ID of context to update. | 
  **body** | [**ContextInputRepresentation**](ContextInputRepresentation.md)| Full context definition to replace previous context. | 

### Return type

[**ContextRepresentation**](ContextRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

