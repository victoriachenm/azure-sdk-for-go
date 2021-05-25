// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// WebCategoriesClient contains the methods for the WebCategories group.
// Don't use this type directly, use NewWebCategoriesClient() instead.
type WebCategoriesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewWebCategoriesClient creates a new instance of WebCategoriesClient with the specified values.
func NewWebCategoriesClient(con *armcore.Connection, subscriptionID string) *WebCategoriesClient {
	return &WebCategoriesClient{con: con, subscriptionID: subscriptionID}
}

// Get - Gets the specified Azure Web Category.
func (client *WebCategoriesClient) Get(ctx context.Context, name string, options *WebCategoriesGetOptions) (AzureWebCategoryResponse, error) {
	req, err := client.getCreateRequest(ctx, name, options)
	if err != nil {
		return AzureWebCategoryResponse{}, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return AzureWebCategoryResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return AzureWebCategoryResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WebCategoriesClient) getCreateRequest(ctx context.Context, name string, options *WebCategoriesGetOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureWebCategories/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	if options != nil && options.Expand != nil {
		query.Set("$expand", *options.Expand)
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WebCategoriesClient) getHandleResponse(resp *azcore.Response) (AzureWebCategoryResponse, error) {
	var val *AzureWebCategory
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return AzureWebCategoryResponse{}, err
	}
	return AzureWebCategoryResponse{RawResponse: resp.Response, AzureWebCategory: val}, nil
}

// getHandleError handles the Get error response.
func (client *WebCategoriesClient) getHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// ListBySubscription - Gets all the Azure Web Categories in a subscription.
func (client *WebCategoriesClient) ListBySubscription(options *WebCategoriesListBySubscriptionOptions) AzureWebCategoryListResultPager {
	return &azureWebCategoryListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		responder: client.listBySubscriptionHandleResponse,
		errorer:   client.listBySubscriptionHandleError,
		advancer: func(ctx context.Context, resp AzureWebCategoryListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AzureWebCategoryListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *WebCategoriesClient) listBySubscriptionCreateRequest(ctx context.Context, options *WebCategoriesListBySubscriptionOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureWebCategories"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2020-07-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *WebCategoriesClient) listBySubscriptionHandleResponse(resp *azcore.Response) (AzureWebCategoryListResultResponse, error) {
	var val *AzureWebCategoryListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return AzureWebCategoryListResultResponse{}, err
	}
	return AzureWebCategoryListResultResponse{RawResponse: resp.Response, AzureWebCategoryListResult: val}, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *WebCategoriesClient) listBySubscriptionHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}