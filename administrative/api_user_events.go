/*
Administrative API

Manage users, accounts, and account groups in the ThousandEyes platform using the Administrative API. This API provides the following operations to manage your organization:     * `/account-groups`: Account groups are used to divide an organization into different sections. These operations can be used to create, retrieve, update and delete account groups.   * `/users`: Create, retrieve, update and delete users within an organization.    * `/roles`: Create, retrieve and update roles for the current user.    * `/permissions`: Retrieve all assignable permissions. Used in the context of modifying roles.    * `/audit-user-events`: Retrieve all activity log events.    For more information about the administrative models, see [Account Management](https://docs.thousandeyes.com/product-documentation/user-management).

API version: 7.0.36
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package administrative

import (
	"bytes"
    "github.com/thousandeyes/thousandeyes-sdk-go/v3/core"
	"io"
	"net/http"
	"net/url"
	"time"
)


// UserEventsAPIService UserEventsAPI service
type UserEventsAPIService core.Service

type ApiGetUserEventsRequest struct {

	ApiService *UserEventsAPIService
	aid *string
	useAllPermittedAids *bool
	window *string
	startDate *time.Time
	endDate *time.Time
	cursor *string
}

// A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. Note that you must be assigned to the target account group. Specifying this parameter without being assigned to the target account group will result in an error response.
func (r ApiGetUserEventsRequest) Aid(aid string) ApiGetUserEventsRequest {
	r.aid = &aid
	return r
}

// Set to &#x60;true&#x60; to load data from all accounts the user has access to.
func (r ApiGetUserEventsRequest) UseAllPermittedAids(useAllPermittedAids bool) ApiGetUserEventsRequest {
	r.useAllPermittedAids = &useAllPermittedAids
	return r
}

// A dynamic time interval up to the current time of the request. Specify the interval as a number followed by an optional type: &#x60;s&#x60; for seconds (default if no type is specified), &#x60;m&#x60; for minutes, &#x60;h&#x60; for hours, &#x60;d&#x60; for days, and &#x60;w&#x60; for weeks. For a precise date range, use &#x60;startDate&#x60; and &#x60;endDate&#x60;.
func (r ApiGetUserEventsRequest) Window(window string) ApiGetUserEventsRequest {
	r.window = &window
	return r
}

// Use with the &#x60;endDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;.
func (r ApiGetUserEventsRequest) StartDate(startDate time.Time) ApiGetUserEventsRequest {
	r.startDate = &startDate
	return r
}

// Defaults to current time the request is made. Use with the &#x60;startDate&#x60; parameter. Include the complete time (hours, minutes, and seconds) in UTC time zone, following the ISO 8601 date-time format. See the example for reference. Please note that this parameter can&#39;t be used with &#x60;window&#x60;.
func (r ApiGetUserEventsRequest) EndDate(endDate time.Time) ApiGetUserEventsRequest {
	r.endDate = &endDate
	return r
}

// (Optional) Opaque cursor used for pagination. Clients should use &#x60;next&#x60; value from &#x60;_links&#x60; instead of this parameter.
func (r ApiGetUserEventsRequest) Cursor(cursor string) ApiGetUserEventsRequest {
	r.cursor = &cursor
	return r
}

func (r ApiGetUserEventsRequest) Execute() (*AuditUserEvents, *http.Response, error) {
	return r.ApiService.GetUserEventsExecute(r)
}

/*
GetUserEvents List activity log events

Returns a list of activity log events in the current account group. 

If `useAllPermittedAids=true` query parameter is passed and the user has permission `View activity log for all users in account group` the logs returned include events across all the account groups they belong to.

For more information about changing the account group context, see [Account Context](https://developer.thousandeyes.com/v7/#/accountcontext).


 @return ApiGetUserEventsRequest
*/
func (a *UserEventsAPIService) GetUserEvents() ApiGetUserEventsRequest {
	return ApiGetUserEventsRequest{
		ApiService: a,
	}
}

// Execute executes the request
//  @return AuditUserEvents
func (a *UserEventsAPIService) GetUserEventsExecute(r ApiGetUserEventsRequest) (*AuditUserEvents, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarReturnValue  *AuditUserEvents
	)

	localBasePath := a.Client.GetConfig().ServerURL

	localVarPath := localBasePath + "/audit-user-events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.aid != nil {
		core.ParameterAddToHeaderOrQuery(localVarQueryParams, "aid", r.aid, "")
	}
	if r.useAllPermittedAids != nil {
		core.ParameterAddToHeaderOrQuery(localVarQueryParams, "useAllPermittedAids", r.useAllPermittedAids, "")
	} else {
		var defaultValue bool = false
		r.useAllPermittedAids = &defaultValue
	}
	if r.window != nil {
		core.ParameterAddToHeaderOrQuery(localVarQueryParams, "window", r.window, "")
	}
	if r.startDate != nil {
		core.ParameterAddToHeaderOrQuery(localVarQueryParams, "startDate", r.startDate, "")
	}
	if r.endDate != nil {
		core.ParameterAddToHeaderOrQuery(localVarQueryParams, "endDate", r.endDate, "")
	}
	if r.cursor != nil {
		core.ParameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := core.SelectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/hal+json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := core.SelectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.Client.PrepareRequest(localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.Client.CallAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &core.GenericAPIError{
			Body:  localVarBody,
			ErrorMessage: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v ValidationError
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.ErrorMessage = core.FormatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.Model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UnauthorizedError
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.ErrorMessage = core.FormatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.Model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.ErrorMessage = core.FormatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.Model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.ErrorMessage = core.FormatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.Model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.ErrorMessage = core.FormatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.Model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.Client.Decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.ErrorMessage = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.ErrorMessage = core.FormatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.Model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.Client.Decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &core.GenericAPIError{
			Body:  localVarBody,
			ErrorMessage: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
