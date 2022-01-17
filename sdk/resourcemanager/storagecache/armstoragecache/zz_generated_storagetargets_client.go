//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstoragecache

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// StorageTargetsClient contains the methods for the StorageTargets group.
// Don't use this type directly, use NewStorageTargetsClient() instead.
type StorageTargetsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewStorageTargetsClient creates a new instance of StorageTargetsClient with the specified values.
// subscriptionID - Subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID forms
// part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewStorageTargetsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *StorageTargetsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &StorageTargetsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Create or update a Storage Target. This operation is allowed at any time, but if the Cache is down
// or unhealthy, the actual creation/modification of the Storage Target may be delayed until the Cache
// is healthy again.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Target resource group.
// cacheName - Name of Cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
// storageTargetName - Name of Storage Target.
// options - StorageTargetsClientBeginCreateOrUpdateOptions contains the optional parameters for the StorageTargetsClient.BeginCreateOrUpdate
// method.
func (client *StorageTargetsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, options *StorageTargetsClientBeginCreateOrUpdateOptions) (StorageTargetsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, cacheName, storageTargetName, options)
	if err != nil {
		return StorageTargetsClientCreateOrUpdatePollerResponse{}, err
	}
	result := StorageTargetsClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("StorageTargetsClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return StorageTargetsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &StorageTargetsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Create or update a Storage Target. This operation is allowed at any time, but if the Cache is down or
// unhealthy, the actual creation/modification of the Storage Target may be delayed until the Cache
// is healthy again.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *StorageTargetsClient) createOrUpdate(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, options *StorageTargetsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, cacheName, storageTargetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *StorageTargetsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, options *StorageTargetsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	if storageTargetName == "" {
		return nil, errors.New("parameter storageTargetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageTargetName}", url.PathEscape(storageTargetName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	if options != nil && options.Storagetarget != nil {
		return req, runtime.MarshalAsJSON(req, *options.Storagetarget)
	}
	return req, nil
}

// BeginDNSRefresh - Tells a storage target to refresh its DNS information.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Target resource group.
// cacheName - Name of Cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
// storageTargetName - Name of Storage Target.
// options - StorageTargetsClientBeginDNSRefreshOptions contains the optional parameters for the StorageTargetsClient.BeginDNSRefresh
// method.
func (client *StorageTargetsClient) BeginDNSRefresh(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, options *StorageTargetsClientBeginDNSRefreshOptions) (StorageTargetsClientDNSRefreshPollerResponse, error) {
	resp, err := client.dNSRefresh(ctx, resourceGroupName, cacheName, storageTargetName, options)
	if err != nil {
		return StorageTargetsClientDNSRefreshPollerResponse{}, err
	}
	result := StorageTargetsClientDNSRefreshPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("StorageTargetsClient.DNSRefresh", "azure-async-operation", resp, client.pl)
	if err != nil {
		return StorageTargetsClientDNSRefreshPollerResponse{}, err
	}
	result.Poller = &StorageTargetsClientDNSRefreshPoller{
		pt: pt,
	}
	return result, nil
}

// DNSRefresh - Tells a storage target to refresh its DNS information.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *StorageTargetsClient) dNSRefresh(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, options *StorageTargetsClientBeginDNSRefreshOptions) (*http.Response, error) {
	req, err := client.dnsRefreshCreateRequest(ctx, resourceGroupName, cacheName, storageTargetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// dnsRefreshCreateRequest creates the DNSRefresh request.
func (client *StorageTargetsClient) dnsRefreshCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, options *StorageTargetsClientBeginDNSRefreshOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}/dnsRefresh"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	if storageTargetName == "" {
		return nil, errors.New("parameter storageTargetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageTargetName}", url.PathEscape(storageTargetName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginDelete - Removes a Storage Target from a Cache. This operation is allowed at any time, but if the Cache is down or
// unhealthy, the actual removal of the Storage Target may be delayed until the Cache is healthy
// again. Note that if the Cache has data to flush to the Storage Target, the data will be flushed before the Storage Target
// will be deleted.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Target resource group.
// cacheName - Name of Cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
// storageTargetName - Name of Storage Target.
// options - StorageTargetsClientBeginDeleteOptions contains the optional parameters for the StorageTargetsClient.BeginDelete
// method.
func (client *StorageTargetsClient) BeginDelete(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, options *StorageTargetsClientBeginDeleteOptions) (StorageTargetsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, cacheName, storageTargetName, options)
	if err != nil {
		return StorageTargetsClientDeletePollerResponse{}, err
	}
	result := StorageTargetsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("StorageTargetsClient.Delete", "", resp, client.pl)
	if err != nil {
		return StorageTargetsClientDeletePollerResponse{}, err
	}
	result.Poller = &StorageTargetsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Removes a Storage Target from a Cache. This operation is allowed at any time, but if the Cache is down or unhealthy,
// the actual removal of the Storage Target may be delayed until the Cache is healthy
// again. Note that if the Cache has data to flush to the Storage Target, the data will be flushed before the Storage Target
// will be deleted.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *StorageTargetsClient) deleteOperation(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, options *StorageTargetsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, cacheName, storageTargetName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *StorageTargetsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, options *StorageTargetsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	if storageTargetName == "" {
		return nil, errors.New("parameter storageTargetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageTargetName}", url.PathEscape(storageTargetName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	if options != nil && options.Force != nil {
		reqQP.Set("force", *options.Force)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Returns a Storage Target from a Cache.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Target resource group.
// cacheName - Name of Cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
// storageTargetName - Name of Storage Target.
// options - StorageTargetsClientGetOptions contains the optional parameters for the StorageTargetsClient.Get method.
func (client *StorageTargetsClient) Get(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, options *StorageTargetsClientGetOptions) (StorageTargetsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, cacheName, storageTargetName, options)
	if err != nil {
		return StorageTargetsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return StorageTargetsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return StorageTargetsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *StorageTargetsClient) getCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, storageTargetName string, options *StorageTargetsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets/{storageTargetName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	if storageTargetName == "" {
		return nil, errors.New("parameter storageTargetName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{storageTargetName}", url.PathEscape(storageTargetName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *StorageTargetsClient) getHandleResponse(resp *http.Response) (StorageTargetsClientGetResponse, error) {
	result := StorageTargetsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageTarget); err != nil {
		return StorageTargetsClientGetResponse{}, err
	}
	return result, nil
}

// ListByCache - Returns a list of Storage Targets for the specified Cache.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - Target resource group.
// cacheName - Name of Cache. Length of name must not be greater than 80 and chars must be from the [-0-9a-zA-Z_] char class.
// options - StorageTargetsClientListByCacheOptions contains the optional parameters for the StorageTargetsClient.ListByCache
// method.
func (client *StorageTargetsClient) ListByCache(resourceGroupName string, cacheName string, options *StorageTargetsClientListByCacheOptions) *StorageTargetsClientListByCachePager {
	return &StorageTargetsClientListByCachePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByCacheCreateRequest(ctx, resourceGroupName, cacheName, options)
		},
		advancer: func(ctx context.Context, resp StorageTargetsClientListByCacheResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.StorageTargetsResult.NextLink)
		},
	}
}

// listByCacheCreateRequest creates the ListByCache request.
func (client *StorageTargetsClient) listByCacheCreateRequest(ctx context.Context, resourceGroupName string, cacheName string, options *StorageTargetsClientListByCacheOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.StorageCache/caches/{cacheName}/storageTargets"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if cacheName == "" {
		return nil, errors.New("parameter cacheName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{cacheName}", url.PathEscape(cacheName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByCacheHandleResponse handles the ListByCache response.
func (client *StorageTargetsClient) listByCacheHandleResponse(resp *http.Response) (StorageTargetsClientListByCacheResponse, error) {
	result := StorageTargetsClientListByCacheResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.StorageTargetsResult); err != nil {
		return StorageTargetsClientListByCacheResponse{}, err
	}
	return result, nil
}