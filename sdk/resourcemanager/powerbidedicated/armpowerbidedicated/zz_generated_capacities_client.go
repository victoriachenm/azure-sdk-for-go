//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armpowerbidedicated

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// CapacitiesClient contains the methods for the Capacities group.
// Don't use this type directly, use NewCapacitiesClient() instead.
type CapacitiesClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewCapacitiesClient creates a new instance of CapacitiesClient with the specified values.
func NewCapacitiesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *CapacitiesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Host) == 0 {
		cp.Host = arm.AzurePublicCloud
	}
	return &CapacitiesClient{subscriptionID: subscriptionID, ep: string(cp.Host), pl: armruntime.NewPipeline(module, version, credential, &cp)}
}

// CheckNameAvailability - Check the name availability in the target location.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CapacitiesClient) CheckNameAvailability(ctx context.Context, location string, capacityParameters CheckCapacityNameAvailabilityParameters, options *CapacitiesCheckNameAvailabilityOptions) (CapacitiesCheckNameAvailabilityResponse, error) {
	req, err := client.checkNameAvailabilityCreateRequest(ctx, location, capacityParameters, options)
	if err != nil {
		return CapacitiesCheckNameAvailabilityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CapacitiesCheckNameAvailabilityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CapacitiesCheckNameAvailabilityResponse{}, client.checkNameAvailabilityHandleError(resp)
	}
	return client.checkNameAvailabilityHandleResponse(resp)
}

// checkNameAvailabilityCreateRequest creates the CheckNameAvailability request.
func (client *CapacitiesClient) checkNameAvailabilityCreateRequest(ctx context.Context, location string, capacityParameters CheckCapacityNameAvailabilityParameters, options *CapacitiesCheckNameAvailabilityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.PowerBIDedicated/locations/{location}/checkNameAvailability"
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, capacityParameters)
}

// checkNameAvailabilityHandleResponse handles the CheckNameAvailability response.
func (client *CapacitiesClient) checkNameAvailabilityHandleResponse(resp *http.Response) (CapacitiesCheckNameAvailabilityResponse, error) {
	result := CapacitiesCheckNameAvailabilityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CheckCapacityNameAvailabilityResult); err != nil {
		return CapacitiesCheckNameAvailabilityResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// checkNameAvailabilityHandleError handles the CheckNameAvailability error response.
func (client *CapacitiesClient) checkNameAvailabilityHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginCreate - Provisions the specified Dedicated capacity based on the configuration specified in the request.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CapacitiesClient) BeginCreate(ctx context.Context, resourceGroupName string, dedicatedCapacityName string, capacityParameters DedicatedCapacity, options *CapacitiesBeginCreateOptions) (CapacitiesCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, dedicatedCapacityName, capacityParameters, options)
	if err != nil {
		return CapacitiesCreatePollerResponse{}, err
	}
	result := CapacitiesCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CapacitiesClient.Create", "", resp, client.pl, client.createHandleError)
	if err != nil {
		return CapacitiesCreatePollerResponse{}, err
	}
	result.Poller = &CapacitiesCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Provisions the specified Dedicated capacity based on the configuration specified in the request.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CapacitiesClient) create(ctx context.Context, resourceGroupName string, dedicatedCapacityName string, capacityParameters DedicatedCapacity, options *CapacitiesBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, dedicatedCapacityName, capacityParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, client.createHandleError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *CapacitiesClient) createCreateRequest(ctx context.Context, resourceGroupName string, dedicatedCapacityName string, capacityParameters DedicatedCapacity, options *CapacitiesBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities/{dedicatedCapacityName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dedicatedCapacityName == "" {
		return nil, errors.New("parameter dedicatedCapacityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dedicatedCapacityName}", url.PathEscape(dedicatedCapacityName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, capacityParameters)
}

// createHandleError handles the Create error response.
func (client *CapacitiesClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginDelete - Deletes the specified Dedicated capacity.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CapacitiesClient) BeginDelete(ctx context.Context, resourceGroupName string, dedicatedCapacityName string, options *CapacitiesBeginDeleteOptions) (CapacitiesDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, dedicatedCapacityName, options)
	if err != nil {
		return CapacitiesDeletePollerResponse{}, err
	}
	result := CapacitiesDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CapacitiesClient.Delete", "", resp, client.pl, client.deleteHandleError)
	if err != nil {
		return CapacitiesDeletePollerResponse{}, err
	}
	result.Poller = &CapacitiesDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified Dedicated capacity.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CapacitiesClient) deleteOperation(ctx context.Context, resourceGroupName string, dedicatedCapacityName string, options *CapacitiesBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, dedicatedCapacityName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.deleteHandleError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CapacitiesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, dedicatedCapacityName string, options *CapacitiesBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities/{dedicatedCapacityName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dedicatedCapacityName == "" {
		return nil, errors.New("parameter dedicatedCapacityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dedicatedCapacityName}", url.PathEscape(dedicatedCapacityName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *CapacitiesClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// GetDetails - Gets details about the specified dedicated capacity.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CapacitiesClient) GetDetails(ctx context.Context, resourceGroupName string, dedicatedCapacityName string, options *CapacitiesGetDetailsOptions) (CapacitiesGetDetailsResponse, error) {
	req, err := client.getDetailsCreateRequest(ctx, resourceGroupName, dedicatedCapacityName, options)
	if err != nil {
		return CapacitiesGetDetailsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CapacitiesGetDetailsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CapacitiesGetDetailsResponse{}, client.getDetailsHandleError(resp)
	}
	return client.getDetailsHandleResponse(resp)
}

// getDetailsCreateRequest creates the GetDetails request.
func (client *CapacitiesClient) getDetailsCreateRequest(ctx context.Context, resourceGroupName string, dedicatedCapacityName string, options *CapacitiesGetDetailsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities/{dedicatedCapacityName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dedicatedCapacityName == "" {
		return nil, errors.New("parameter dedicatedCapacityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dedicatedCapacityName}", url.PathEscape(dedicatedCapacityName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getDetailsHandleResponse handles the GetDetails response.
func (client *CapacitiesClient) getDetailsHandleResponse(resp *http.Response) (CapacitiesGetDetailsResponse, error) {
	result := CapacitiesGetDetailsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedCapacity); err != nil {
		return CapacitiesGetDetailsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// getDetailsHandleError handles the GetDetails error response.
func (client *CapacitiesClient) getDetailsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// List - Lists all the Dedicated capacities for the given subscription.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CapacitiesClient) List(ctx context.Context, options *CapacitiesListOptions) (CapacitiesListResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return CapacitiesListResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CapacitiesListResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CapacitiesListResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *CapacitiesClient) listCreateRequest(ctx context.Context, options *CapacitiesListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.PowerBIDedicated/capacities"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *CapacitiesClient) listHandleResponse(resp *http.Response) (CapacitiesListResponse, error) {
	result := CapacitiesListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedCapacities); err != nil {
		return CapacitiesListResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listHandleError handles the List error response.
func (client *CapacitiesClient) listHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - Gets all the Dedicated capacities for the given resource group.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CapacitiesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string, options *CapacitiesListByResourceGroupOptions) (CapacitiesListByResourceGroupResponse, error) {
	req, err := client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
	if err != nil {
		return CapacitiesListByResourceGroupResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CapacitiesListByResourceGroupResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CapacitiesListByResourceGroupResponse{}, client.listByResourceGroupHandleError(resp)
	}
	return client.listByResourceGroupHandleResponse(resp)
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *CapacitiesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *CapacitiesListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *CapacitiesClient) listByResourceGroupHandleResponse(resp *http.Response) (CapacitiesListByResourceGroupResponse, error) {
	result := CapacitiesListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.DedicatedCapacities); err != nil {
		return CapacitiesListByResourceGroupResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *CapacitiesClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListSKUs - Lists eligible SKUs for PowerBI Dedicated resource provider.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CapacitiesClient) ListSKUs(ctx context.Context, options *CapacitiesListSKUsOptions) (CapacitiesListSKUsResponse, error) {
	req, err := client.listSKUsCreateRequest(ctx, options)
	if err != nil {
		return CapacitiesListSKUsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CapacitiesListSKUsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CapacitiesListSKUsResponse{}, client.listSKUsHandleError(resp)
	}
	return client.listSKUsHandleResponse(resp)
}

// listSKUsCreateRequest creates the ListSKUs request.
func (client *CapacitiesClient) listSKUsCreateRequest(ctx context.Context, options *CapacitiesListSKUsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.PowerBIDedicated/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listSKUsHandleResponse handles the ListSKUs response.
func (client *CapacitiesClient) listSKUsHandleResponse(resp *http.Response) (CapacitiesListSKUsResponse, error) {
	result := CapacitiesListSKUsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SKUEnumerationForNewResourceResult); err != nil {
		return CapacitiesListSKUsResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listSKUsHandleError handles the ListSKUs error response.
func (client *CapacitiesClient) listSKUsHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListSKUsForCapacity - Lists eligible SKUs for a PowerBI Dedicated resource.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CapacitiesClient) ListSKUsForCapacity(ctx context.Context, resourceGroupName string, dedicatedCapacityName string, options *CapacitiesListSKUsForCapacityOptions) (CapacitiesListSKUsForCapacityResponse, error) {
	req, err := client.listSKUsForCapacityCreateRequest(ctx, resourceGroupName, dedicatedCapacityName, options)
	if err != nil {
		return CapacitiesListSKUsForCapacityResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CapacitiesListSKUsForCapacityResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CapacitiesListSKUsForCapacityResponse{}, client.listSKUsForCapacityHandleError(resp)
	}
	return client.listSKUsForCapacityHandleResponse(resp)
}

// listSKUsForCapacityCreateRequest creates the ListSKUsForCapacity request.
func (client *CapacitiesClient) listSKUsForCapacityCreateRequest(ctx context.Context, resourceGroupName string, dedicatedCapacityName string, options *CapacitiesListSKUsForCapacityOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities/{dedicatedCapacityName}/skus"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dedicatedCapacityName == "" {
		return nil, errors.New("parameter dedicatedCapacityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dedicatedCapacityName}", url.PathEscape(dedicatedCapacityName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listSKUsForCapacityHandleResponse handles the ListSKUsForCapacity response.
func (client *CapacitiesClient) listSKUsForCapacityHandleResponse(resp *http.Response) (CapacitiesListSKUsForCapacityResponse, error) {
	result := CapacitiesListSKUsForCapacityResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.SKUEnumerationForExistingResourceResult); err != nil {
		return CapacitiesListSKUsForCapacityResponse{}, runtime.NewResponseError(err, resp)
	}
	return result, nil
}

// listSKUsForCapacityHandleError handles the ListSKUsForCapacity error response.
func (client *CapacitiesClient) listSKUsForCapacityHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginResume - Resumes operation of the specified Dedicated capacity instance.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CapacitiesClient) BeginResume(ctx context.Context, resourceGroupName string, dedicatedCapacityName string, options *CapacitiesBeginResumeOptions) (CapacitiesResumePollerResponse, error) {
	resp, err := client.resume(ctx, resourceGroupName, dedicatedCapacityName, options)
	if err != nil {
		return CapacitiesResumePollerResponse{}, err
	}
	result := CapacitiesResumePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CapacitiesClient.Resume", "", resp, client.pl, client.resumeHandleError)
	if err != nil {
		return CapacitiesResumePollerResponse{}, err
	}
	result.Poller = &CapacitiesResumePoller{
		pt: pt,
	}
	return result, nil
}

// Resume - Resumes operation of the specified Dedicated capacity instance.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CapacitiesClient) resume(ctx context.Context, resourceGroupName string, dedicatedCapacityName string, options *CapacitiesBeginResumeOptions) (*http.Response, error) {
	req, err := client.resumeCreateRequest(ctx, resourceGroupName, dedicatedCapacityName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.resumeHandleError(resp)
	}
	return resp, nil
}

// resumeCreateRequest creates the Resume request.
func (client *CapacitiesClient) resumeCreateRequest(ctx context.Context, resourceGroupName string, dedicatedCapacityName string, options *CapacitiesBeginResumeOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities/{dedicatedCapacityName}/resume"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dedicatedCapacityName == "" {
		return nil, errors.New("parameter dedicatedCapacityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dedicatedCapacityName}", url.PathEscape(dedicatedCapacityName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// resumeHandleError handles the Resume error response.
func (client *CapacitiesClient) resumeHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginSuspend - Suspends operation of the specified dedicated capacity instance.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CapacitiesClient) BeginSuspend(ctx context.Context, resourceGroupName string, dedicatedCapacityName string, options *CapacitiesBeginSuspendOptions) (CapacitiesSuspendPollerResponse, error) {
	resp, err := client.suspend(ctx, resourceGroupName, dedicatedCapacityName, options)
	if err != nil {
		return CapacitiesSuspendPollerResponse{}, err
	}
	result := CapacitiesSuspendPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CapacitiesClient.Suspend", "", resp, client.pl, client.suspendHandleError)
	if err != nil {
		return CapacitiesSuspendPollerResponse{}, err
	}
	result.Poller = &CapacitiesSuspendPoller{
		pt: pt,
	}
	return result, nil
}

// Suspend - Suspends operation of the specified dedicated capacity instance.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CapacitiesClient) suspend(ctx context.Context, resourceGroupName string, dedicatedCapacityName string, options *CapacitiesBeginSuspendOptions) (*http.Response, error) {
	req, err := client.suspendCreateRequest(ctx, resourceGroupName, dedicatedCapacityName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.suspendHandleError(resp)
	}
	return resp, nil
}

// suspendCreateRequest creates the Suspend request.
func (client *CapacitiesClient) suspendCreateRequest(ctx context.Context, resourceGroupName string, dedicatedCapacityName string, options *CapacitiesBeginSuspendOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities/{dedicatedCapacityName}/suspend"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dedicatedCapacityName == "" {
		return nil, errors.New("parameter dedicatedCapacityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dedicatedCapacityName}", url.PathEscape(dedicatedCapacityName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// suspendHandleError handles the Suspend error response.
func (client *CapacitiesClient) suspendHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// BeginUpdate - Updates the current state of the specified Dedicated capacity.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CapacitiesClient) BeginUpdate(ctx context.Context, resourceGroupName string, dedicatedCapacityName string, capacityUpdateParameters DedicatedCapacityUpdateParameters, options *CapacitiesBeginUpdateOptions) (CapacitiesUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, dedicatedCapacityName, capacityUpdateParameters, options)
	if err != nil {
		return CapacitiesUpdatePollerResponse{}, err
	}
	result := CapacitiesUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("CapacitiesClient.Update", "", resp, client.pl, client.updateHandleError)
	if err != nil {
		return CapacitiesUpdatePollerResponse{}, err
	}
	result.Poller = &CapacitiesUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Updates the current state of the specified Dedicated capacity.
// If the operation fails it returns the *ErrorResponse error type.
func (client *CapacitiesClient) update(ctx context.Context, resourceGroupName string, dedicatedCapacityName string, capacityUpdateParameters DedicatedCapacityUpdateParameters, options *CapacitiesBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, dedicatedCapacityName, capacityUpdateParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, client.updateHandleError(resp)
	}
	return resp, nil
}

// updateCreateRequest creates the Update request.
func (client *CapacitiesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, dedicatedCapacityName string, capacityUpdateParameters DedicatedCapacityUpdateParameters, options *CapacitiesBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.PowerBIDedicated/capacities/{dedicatedCapacityName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if dedicatedCapacityName == "" {
		return nil, errors.New("parameter dedicatedCapacityName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{dedicatedCapacityName}", url.PathEscape(dedicatedCapacityName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-01-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, capacityUpdateParameters)
}

// updateHandleError handles the Update error response.
func (client *CapacitiesClient) updateHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := ErrorResponse{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}