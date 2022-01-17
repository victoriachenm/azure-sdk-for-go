//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtrafficmanager

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// GeographicHierarchiesClient contains the methods for the GeographicHierarchies group.
// Don't use this type directly, use NewGeographicHierarchiesClient() instead.
type GeographicHierarchiesClient struct {
	host string
	pl   runtime.Pipeline
}

// NewGeographicHierarchiesClient creates a new instance of GeographicHierarchiesClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewGeographicHierarchiesClient(credential azcore.TokenCredential, options *arm.ClientOptions) *GeographicHierarchiesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &GeographicHierarchiesClient{
		host: string(cp.Endpoint),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// GetDefault - Gets the default Geographic Hierarchy used by the Geographic traffic routing method.
// If the operation fails it returns an *azcore.ResponseError type.
// options - GeographicHierarchiesClientGetDefaultOptions contains the optional parameters for the GeographicHierarchiesClient.GetDefault
// method.
func (client *GeographicHierarchiesClient) GetDefault(ctx context.Context, options *GeographicHierarchiesClientGetDefaultOptions) (GeographicHierarchiesClientGetDefaultResponse, error) {
	req, err := client.getDefaultCreateRequest(ctx, options)
	if err != nil {
		return GeographicHierarchiesClientGetDefaultResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return GeographicHierarchiesClientGetDefaultResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return GeographicHierarchiesClientGetDefaultResponse{}, runtime.NewResponseError(resp)
	}
	return client.getDefaultHandleResponse(resp)
}

// getDefaultCreateRequest creates the GetDefault request.
func (client *GeographicHierarchiesClient) getDefaultCreateRequest(ctx context.Context, options *GeographicHierarchiesClientGetDefaultOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Network/trafficManagerGeographicHierarchies/default"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-08-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getDefaultHandleResponse handles the GetDefault response.
func (client *GeographicHierarchiesClient) getDefaultHandleResponse(resp *http.Response) (GeographicHierarchiesClientGetDefaultResponse, error) {
	result := GeographicHierarchiesClientGetDefaultResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.GeographicHierarchy); err != nil {
		return GeographicHierarchiesClientGetDefaultResponse{}, err
	}
	return result, nil
}