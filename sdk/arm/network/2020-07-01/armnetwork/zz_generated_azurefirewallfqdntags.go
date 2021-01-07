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

// AzureFirewallFqdnTagsClient contains the methods for the AzureFirewallFqdnTags group.
// Don't use this type directly, use NewAzureFirewallFqdnTagsClient() instead.
type AzureFirewallFqdnTagsClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewAzureFirewallFqdnTagsClient creates a new instance of AzureFirewallFqdnTagsClient with the specified values.
func NewAzureFirewallFqdnTagsClient(con *armcore.Connection, subscriptionID string) AzureFirewallFqdnTagsClient {
	return AzureFirewallFqdnTagsClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client AzureFirewallFqdnTagsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// ListAll - Gets all the Azure Firewall FQDN Tags in a subscription.
func (client AzureFirewallFqdnTagsClient) ListAll(options *AzureFirewallFqdnTagsListAllOptions) AzureFirewallFqdnTagListResultPager {
	return &azureFirewallFqdnTagListResultPager{
		pipeline: client.con.Pipeline(),
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.listAllCreateRequest(ctx, options)
		},
		responder: client.listAllHandleResponse,
		errorer:   client.listAllHandleError,
		advancer: func(ctx context.Context, resp AzureFirewallFqdnTagListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.AzureFirewallFqdnTagListResult.NextLink)
		},
		statusCodes: []int{http.StatusOK},
	}
}

// listAllCreateRequest creates the ListAll request.
func (client AzureFirewallFqdnTagsClient) listAllCreateRequest(ctx context.Context, options *AzureFirewallFqdnTagsListAllOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/azureFirewallFqdnTags"
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

// listAllHandleResponse handles the ListAll response.
func (client AzureFirewallFqdnTagsClient) listAllHandleResponse(resp *azcore.Response) (AzureFirewallFqdnTagListResultResponse, error) {
	result := AzureFirewallFqdnTagListResultResponse{RawResponse: resp.Response}
	err := resp.UnmarshalAsJSON(&result.AzureFirewallFqdnTagListResult)
	return result, err
}

// listAllHandleError handles the ListAll error response.
func (client AzureFirewallFqdnTagsClient) listAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}