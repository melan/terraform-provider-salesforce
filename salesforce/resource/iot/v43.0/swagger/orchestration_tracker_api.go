/*
 * Salesforce IoT API
 *
 * Use the Salesforce IoT REST API to create and manage orchestrations and contexts, and retrieve usage data.
 *
 * API version: 43.0.0
 * Contact: support@salesforce.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"io/ioutil"
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"encoding/json"
	"fmt"
)

// Linger please
var (
	_ context.Context
)

type OrchestrationTrackerApiService service


/* OrchestrationTrackerApiService Create a new tracker for an orchestration.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param orchestrationId ID of the orchestration for which to create a tracker.
 @param body Tracker definition.
 @return OrchestrationTrackerRepresentation*/
func (a *OrchestrationTrackerApiService) CreateTracker(ctx context.Context, orchestrationId string, body OrchestrationTrackerPostInputRepresentation) (OrchestrationTrackerRepresentation,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  OrchestrationTrackerRepresentation
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/orchestrations/{orchestrationId}/trackers"
	localVarPath = strings.Replace(localVarPath, "{"+"orchestrationId"+"}", fmt.Sprintf("%v", orchestrationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* OrchestrationTrackerApiService Delete a tracker.
 
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param orchestrationId ID of the tracker&#39;s orchestration for which to delete the tracker.
 @param trackerId ID of the tracker to delete.
 @return */
func (a *OrchestrationTrackerApiService) DeleteTracker(ctx context.Context, orchestrationId string, trackerId string) ( *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/orchestrations/{orchestrationId}/trackers/{trackerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orchestrationId"+"}", fmt.Sprintf("%v", orchestrationId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"trackerId"+"}", fmt.Sprintf("%v", trackerId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	return localVarHttpResponse, err
}

/* OrchestrationTrackerApiService Get a log of orchestration events for a specific orchestration and key value.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param orchestrationId Orchestration ID for which to get the log of events for a specific orchestration instance key value.
 @param instanceKey Orchestration instance key value for which to get a log of orchestration events.
 @param fromDate Get log events for which the timestamp is later than fromDate.
 @param toDate Get log events for which the timestamp is prior to toDate
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "page" (string) For internal use only.
     @param "pageSize" (int32) The number of log events returned per page. When one page is returned, it contains the URL in the nextPageUrl property that you use to get the next page of results.
 @return OrchestrationLogEventCollectionRepresentation*/
func (a *OrchestrationTrackerApiService) GetOrchestrationInstanceLog(ctx context.Context, orchestrationId string, instanceKey string, fromDate string, toDate string, localVarOptionals map[string]interface{}) (OrchestrationLogEventCollectionRepresentation,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  OrchestrationLogEventCollectionRepresentation
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/orchestrations/{orchestrationId}/instances/{instanceKey}/log"
	localVarPath = strings.Replace(localVarPath, "{"+"orchestrationId"+"}", fmt.Sprintf("%v", orchestrationId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instanceKey"+"}", fmt.Sprintf("%v", instanceKey), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["page"], "string", "page"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["pageSize"], "int32", "pageSize"); err != nil {
		return successPayload, nil, err
	}

	localVarQueryParams.Add("fromDate", parameterToString(fromDate, ""))
	localVarQueryParams.Add("toDate", parameterToString(toDate, ""))
	if localVarTempParam, localVarOk := localVarOptionals["page"].(string); localVarOk {
		localVarQueryParams.Add("page", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["pageSize"].(int32); localVarOk {
		localVarQueryParams.Add("pageSize", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* OrchestrationTrackerApiService Get a list of trackers setup for this orchestration.
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param orchestrationId ID of the orchestration for which to get trackers.
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "page" (string) For internal use only.
     @param "pageSize" (int32) The number of trackers returned per page. When a page of the response is returned, it contains the nextPageUrl property which is the URL that you use to get the next page of results.
 @return OrchestrationTrackerCollectionRepresentation*/
func (a *OrchestrationTrackerApiService) GetOrchestrationTrackers(ctx context.Context, orchestrationId string, localVarOptionals map[string]interface{}) (OrchestrationTrackerCollectionRepresentation,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  OrchestrationTrackerCollectionRepresentation
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/orchestrations/{orchestrationId}/trackers"
	localVarPath = strings.Replace(localVarPath, "{"+"orchestrationId"+"}", fmt.Sprintf("%v", orchestrationId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["page"], "string", "page"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["pageSize"], "int32", "pageSize"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["page"].(string); localVarOk {
		localVarQueryParams.Add("page", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["pageSize"].(int32); localVarOk {
		localVarQueryParams.Add("pageSize", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* OrchestrationTrackerApiService Get the specified tracker corresponding to the specified orchestration.
 
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param orchestrationId ID of the tracker&#39;s orchestration.
 @param trackerId ID of the tracker.
 @return OrchestrationTrackerRepresentation*/
func (a *OrchestrationTrackerApiService) GetTrackerById(ctx context.Context, orchestrationId string, trackerId string) (OrchestrationTrackerRepresentation,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  OrchestrationTrackerRepresentation
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/orchestrations/{orchestrationId}/trackers/{trackerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orchestrationId"+"}", fmt.Sprintf("%v", orchestrationId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"trackerId"+"}", fmt.Sprintf("%v", trackerId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

/* OrchestrationTrackerApiService Update a tracker.
 
 * @param ctx context.Context for authentication, logging, tracing, etc.
 @param orchestrationId ID of a tracker&#39;s orchestration.
 @param trackerId ID of a tracker.
 @param body Tracker to update.
 @return OrchestrationTrackerRepresentation*/
func (a *OrchestrationTrackerApiService) UpdateTracker(ctx context.Context, orchestrationId string, trackerId string, body OrchestrationTrackerPatchInputRepresentation) (OrchestrationTrackerRepresentation,  *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Patch")
		localVarPostBody interface{}
		localVarFileName string
		localVarFileBytes []byte
	 	successPayload  OrchestrationTrackerRepresentation
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/orchestrations/{orchestrationId}/trackers/{trackerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"orchestrationId"+"}", fmt.Sprintf("%v", orchestrationId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"trackerId"+"}", fmt.Sprintf("%v", trackerId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}


	// to determine the Content-Type header
	localVarHttpContentTypes := []string{ "application/json",  }

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	localVarPostBody = &body
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		bodyBytes, _ := ioutil.ReadAll(localVarHttpResponse.Body)
		return successPayload, localVarHttpResponse, reportError("Status: %v, Body: %s", localVarHttpResponse.Status, bodyBytes)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}


	return successPayload, localVarHttpResponse, err
}

