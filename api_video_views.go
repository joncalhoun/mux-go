// Mux Go - Copyright 2019 Mux Inc.
// NOTE: This file is auto generated. Do not edit this file manually.

package muxgo

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type VideoViewsApiService service

/*
VideoViewsApiService Get a Video View
Returns the details of a video view
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param vIDEOVIEWID ID of the Video View
@return VideoViewResponse
*/
func (a *VideoViewsApiService) GetVideoView(ctx context.Context, vIDEOVIEWID string) (VideoViewResponse, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  VideoViewResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.basePath + "/data/v1/video-views/{VIDEO_VIEW_ID}"
	localVarPath = strings.Replace(localVarPath, "{"+"VIDEO_VIEW_ID"+"}", fmt.Sprintf("%v", vIDEOVIEWID), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v VideoViewResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.model = v
			return localVarReturnValue, newErr
		}
		return localVarReturnValue, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}

/*
VideoViewsApiService List Video Views
Returns a list of video views
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param optional nil or *ListVideoViewsOpts - Optional Parameters:
 * @param "Limit" (optional.Int32) -  Number of items to include in the response
 * @param "Page" (optional.Int32) -  Offset by this many pages, of the size of `limit`
 * @param "ViewerId" (optional.String) -  Viewer ID to filter results by. This value may be provided by the integration, or may be created by Mux.
 * @param "ErrorId" (optional.Int32) -  Filter video views by the provided error ID (as returned in the error_type_id field in the list video views endpoint). If you provide any as the error ID, this will filter the results to those with any error.
 * @param "OrderDirection" (optional.String) -  Sort order.
 * @param "Filters" (optional.Interface of []string) -  Filter key:value pairs. Must be provided as an array query string parameter (e.g. filters[]=operating_system:windows&filters[]=country:US).  Possible filter names are the same as returned by the List Filters endpoint.
 * @param "Timeframe" (optional.Interface of []string) -  Timeframe window to limit results by. Must be provided as an array query string parameter (e.g. timeframe[]=). Accepted formats are...   * array of epoch timestamps e.g. timeframe[]=1498867200&timeframe[]=1498953600    * duration string e.g. timeframe[]=24:hours or timeframe[]=7:days.
@return ListVideoViewsResponse
*/

type ListVideoViewsOpts struct {
	Limit          optional.Int32
	Page           optional.Int32
	ViewerId       optional.String
	ErrorId        optional.Int32
	OrderDirection optional.String
	Filters        optional.Interface
	Timeframe      optional.Interface
}

func (a *VideoViewsApiService) ListVideoViews(ctx context.Context, localVarOptionals *ListVideoViewsOpts) (ListVideoViewsResponse, error) {
	var (
		localVarHttpMethod   = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  ListVideoViewsResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.basePath + "/data/v1/video-views"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Limit.IsSet() {
		localVarQueryParams.Add("limit", parameterToString(localVarOptionals.Limit.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Page.IsSet() {
		localVarQueryParams.Add("page", parameterToString(localVarOptionals.Page.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ViewerId.IsSet() {
		localVarQueryParams.Add("viewer_id", parameterToString(localVarOptionals.ViewerId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ErrorId.IsSet() {
		localVarQueryParams.Add("error_id", parameterToString(localVarOptionals.ErrorId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OrderDirection.IsSet() {
		localVarQueryParams.Add("order_direction", parameterToString(localVarOptionals.OrderDirection.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Filters.IsSet() {
		localVarQueryParams.Add("filters[]", parameterToString(localVarOptionals.Filters.Value(), "multi"))
	}
	if localVarOptionals != nil && localVarOptionals.Timeframe.IsSet() {
		localVarQueryParams.Add("timeframe[]", parameterToString(localVarOptionals.Timeframe.Value(), "multi"))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v ListVideoViewsResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, newErr
			}
			newErr.model = v
			return localVarReturnValue, newErr
		}
		return localVarReturnValue, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, newErr
	}

	return localVarReturnValue, nil
}
