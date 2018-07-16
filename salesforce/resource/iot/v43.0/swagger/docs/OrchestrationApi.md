# \OrchestrationApi

All URIs are relative to *http://salesforce.com/services/data/v43.0/iot*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrchestration**](OrchestrationApi.md#CreateOrchestration) | **Post** /orchestrations | Create an orchestration.
[**CreateOrchestrationActivation**](OrchestrationApi.md#CreateOrchestrationActivation) | **Post** /orchestrations/{orchestrationId}/activations | Activate the specified orchestration, optionally using activation options.
[**DeleteOrchestration**](OrchestrationApi.md#DeleteOrchestration) | **Delete** /orchestrations/{orchestrationId} | Delete the specified orchestration.
[**GetLastOrchestrationActivation**](OrchestrationApi.md#GetLastOrchestrationActivation) | **Get** /orchestrations/{orchestrationId}/activations/last | Get the last activation for the specified orchestration.
[**GetOrchestrationActivationById**](OrchestrationApi.md#GetOrchestrationActivationById) | **Get** /orchestrations/{orchestrationId}/activations/{activationId} | Get activation information for the specified orchestration and activation.
[**GetOrchestrationActivations**](OrchestrationApi.md#GetOrchestrationActivations) | **Get** /orchestrations/{orchestrationId}/activations | Get the activation history for the specified orchestration. The activations are sorted by creation date in descending order.
[**GetOrchestrationById**](OrchestrationApi.md#GetOrchestrationById) | **Get** /orchestrations/{orchestrationId} | Get the specified orchestration.
[**GetOrchestrationStates**](OrchestrationApi.md#GetOrchestrationStates) | **Get** /orchestrations/{orchestrationId}/states | Get the states in this orchestration.
[**GetOrchestrations**](OrchestrationApi.md#GetOrchestrations) | **Get** /orchestrations | List all orchestrations that have been defined in your org.
[**UpdateLastOrchestrationActivation**](OrchestrationApi.md#UpdateLastOrchestrationActivation) | **Patch** /orchestrations/{orchestrationId}/activations/last | Deactivate the specified orchestration. Can only be used to deactivate an active orchestration.
[**UpdateOrchestrationActivation**](OrchestrationApi.md#UpdateOrchestrationActivation) | **Patch** /orchestrations/{orchestrationId}/activations/{activationId} | Deactivate the specified orchestration activation. No other updates are supported.
[**UpdateOrchestrationWithForm**](OrchestrationApi.md#UpdateOrchestrationWithForm) | **Put** /orchestrations/{orchestrationId} | Update the specified orchestration using the full orchestration definition.


# **CreateOrchestration**
> OrchestrationRepresentation CreateOrchestration(ctx, body)
Create an orchestration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **body** | [**OrchestrationInputRepresentation**](OrchestrationInputRepresentation.md)| Orchestration definition. | 

### Return type

[**OrchestrationRepresentation**](OrchestrationRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateOrchestrationActivation**
> OrchestrationActivationRepresentation CreateOrchestrationActivation(ctx, orchestrationId, body)
Activate the specified orchestration, optionally using activation options.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **orchestrationId** | **string**| ID of the orchestration to activate. | 
  **body** | [**OrchestrationActivationInputRepresentation**](OrchestrationActivationInputRepresentation.md)| Optional activation options. Otherwise, an empty body ({}). | 

### Return type

[**OrchestrationActivationRepresentation**](OrchestrationActivationRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteOrchestration**
> DeleteOrchestration(ctx, orchestrationId)
Delete the specified orchestration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **orchestrationId** | **string**| ID of the orchestration to delete. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLastOrchestrationActivation**
> OrchestrationActivationRepresentation GetLastOrchestrationActivation(ctx, orchestrationId)
Get the last activation for the specified orchestration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **orchestrationId** | **string**| ID of the orchestration to get activation information for. | 

### Return type

[**OrchestrationActivationRepresentation**](OrchestrationActivationRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrchestrationActivationById**
> OrchestrationActivationRepresentation GetOrchestrationActivationById(ctx, orchestrationId, activationId)
Get activation information for the specified orchestration and activation.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **orchestrationId** | **string**| ID of the orchestration for which to get activation information. | 
  **activationId** | **string**| ID of the activation to get details of. | 

### Return type

[**OrchestrationActivationRepresentation**](OrchestrationActivationRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrchestrationActivations**
> OrchestrationActivationCollectionRepresentation GetOrchestrationActivations(ctx, orchestrationId, optional)
Get the activation history for the specified orchestration. The activations are sorted by creation date in descending order.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **orchestrationId** | **string**| ID of the orchestration to get the activation history for. | 
 **optional** | **map[string]interface{}** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a map[string]interface{}.

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orchestrationId** | **string**| ID of the orchestration to get the activation history for. | 
 **page** | **string**| For internal use only. | 
 **pageSize** | **int32**| The number of activations returned per page. When a page of the response is returned, it contains the nextPageUrl property which is the URL that you use to get the next page of results. | 

### Return type

[**OrchestrationActivationCollectionRepresentation**](OrchestrationActivationCollectionRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrchestrationById**
> OrchestrationRepresentation GetOrchestrationById(ctx, orchestrationId)
Get the specified orchestration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **orchestrationId** | **string**| ID of the orchestration to retrieve. | 

### Return type

[**OrchestrationRepresentation**](OrchestrationRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrchestrationStates**
> OrchestrationStateCollectionRepresentation GetOrchestrationStates(ctx, orchestrationId)
Get the states in this orchestration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **orchestrationId** | **string**| ID of the orchestration for which to get the states. | 

### Return type

[**OrchestrationStateCollectionRepresentation**](OrchestrationStateCollectionRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrchestrations**
> OrchestrationCollectionRepresentation GetOrchestrations(ctx, optional)
List all orchestrations that have been defined in your org.



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
 **pageSize** | **int32**| The number of orchestrations returned per page. When one page is returned, it contains the URL in the nextPageUrl property that you use to get the next page of results. | 
 **runtimeType** | **string**| The Salesforce IoT runtime used. | 
 **sortDirection** | **string**| The sort direction, either ascending or descending. | 
 **sort** | **string**| The field by which to sort the orchestrations. | 

### Return type

[**OrchestrationCollectionRepresentation**](OrchestrationCollectionRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateLastOrchestrationActivation**
> OrchestrationActivationRepresentation UpdateLastOrchestrationActivation(ctx, orchestrationId, body)
Deactivate the specified orchestration. Can only be used to deactivate an active orchestration.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **orchestrationId** | **string**| ID of the orchestration to deactivate. | 
  **body** | [**OrchestrationActivationInputRepresentation**](OrchestrationActivationInputRepresentation.md)| The body must include the status property with a value of Inactive. Optionally, you can specify whether to delete running orchestration instances by including the resetInstancesOnActivation option. | 

### Return type

[**OrchestrationActivationRepresentation**](OrchestrationActivationRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrchestrationActivation**
> OrchestrationActivationRepresentation UpdateOrchestrationActivation(ctx, orchestrationId, activationId, body)
Deactivate the specified orchestration activation. No other updates are supported.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **orchestrationId** | **string**| ID of the orchestration to deactivate. | 
  **activationId** | **string**| ID of the activation to deactivate. | 
  **body** | [**OrchestrationActivationInputRepresentation**](OrchestrationActivationInputRepresentation.md)| The body must include the status property with a value of Inactive. Optionally, you can specify whether to delete running orchestration instances by including the resetInstancesOnActivation option. | 

### Return type

[**OrchestrationActivationRepresentation**](OrchestrationActivationRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateOrchestrationWithForm**
> OrchestrationRepresentation UpdateOrchestrationWithForm(ctx, orchestrationId, orchestrationInput)
Update the specified orchestration using the full orchestration definition.



### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for logging, tracing, authentication, etc.
  **orchestrationId** | **string**| ID of orchestration to update. | 
  **orchestrationInput** | [**OrchestrationInputRepresentation**](OrchestrationInputRepresentation.md)| The full orchestration definition to replace the current orchestration with. | 

### Return type

[**OrchestrationRepresentation**](OrchestrationRepresentation.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/x-www-form-urlencoded
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

