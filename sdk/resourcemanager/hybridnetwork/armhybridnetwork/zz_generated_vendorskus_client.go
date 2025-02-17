//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhybridnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VendorSKUsClient contains the methods for the VendorSKUs group.
// Don't use this type directly, use NewVendorSKUsClient() instead.
type VendorSKUsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVendorSKUsClient creates a new instance of VendorSKUsClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVendorSKUsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VendorSKUsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &VendorSKUsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a sku. This operation can take up to 2 hours to complete. This is expected service
// behavior.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// vendorName - The name of the vendor.
// skuName - The name of the sku.
// parameters - Parameters supplied to the create or update sku operation.
// options - VendorSKUsClientBeginCreateOrUpdateOptions contains the optional parameters for the VendorSKUsClient.BeginCreateOrUpdate
// method.
func (client *VendorSKUsClient) BeginCreateOrUpdate(ctx context.Context, vendorName string, skuName string, parameters VendorSKU, options *VendorSKUsClientBeginCreateOrUpdateOptions) (*runtime.Poller[VendorSKUsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, vendorName, skuName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VendorSKUsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VendorSKUsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a sku. This operation can take up to 2 hours to complete. This is expected service
// behavior.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
func (client *VendorSKUsClient) createOrUpdate(ctx context.Context, vendorName string, skuName string, parameters VendorSKU, options *VendorSKUsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, vendorName, skuName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VendorSKUsClient) createOrUpdateCreateRequest(ctx context.Context, vendorName string, skuName string, parameters VendorSKU, options *VendorSKUsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/vendors/{vendorName}/vendorSkus/{skuName}"
	if vendorName == "" {
		return nil, errors.New("parameter vendorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vendorName}", url.PathEscape(vendorName))
	if skuName == "" {
		return nil, errors.New("parameter skuName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skuName}", url.PathEscape(skuName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified sku. This operation can take up to 2 hours to complete. This is expected service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// vendorName - The name of the vendor.
// skuName - The name of the sku.
// options - VendorSKUsClientBeginDeleteOptions contains the optional parameters for the VendorSKUsClient.BeginDelete method.
func (client *VendorSKUsClient) BeginDelete(ctx context.Context, vendorName string, skuName string, options *VendorSKUsClientBeginDeleteOptions) (*runtime.Poller[VendorSKUsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, vendorName, skuName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[VendorSKUsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[VendorSKUsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified sku. This operation can take up to 2 hours to complete. This is expected service behavior.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
func (client *VendorSKUsClient) deleteOperation(ctx context.Context, vendorName string, skuName string, options *VendorSKUsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, vendorName, skuName, options)
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
func (client *VendorSKUsClient) deleteCreateRequest(ctx context.Context, vendorName string, skuName string, options *VendorSKUsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/vendors/{vendorName}/vendorSkus/{skuName}"
	if vendorName == "" {
		return nil, errors.New("parameter vendorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vendorName}", url.PathEscape(vendorName))
	if skuName == "" {
		return nil, errors.New("parameter skuName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skuName}", url.PathEscape(skuName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about the specified sku.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// vendorName - The name of the vendor.
// skuName - The name of the sku.
// options - VendorSKUsClientGetOptions contains the optional parameters for the VendorSKUsClient.Get method.
func (client *VendorSKUsClient) Get(ctx context.Context, vendorName string, skuName string, options *VendorSKUsClientGetOptions) (VendorSKUsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, vendorName, skuName, options)
	if err != nil {
		return VendorSKUsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VendorSKUsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VendorSKUsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VendorSKUsClient) getCreateRequest(ctx context.Context, vendorName string, skuName string, options *VendorSKUsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/vendors/{vendorName}/vendorSkus/{skuName}"
	if vendorName == "" {
		return nil, errors.New("parameter vendorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vendorName}", url.PathEscape(vendorName))
	if skuName == "" {
		return nil, errors.New("parameter skuName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skuName}", url.PathEscape(skuName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *VendorSKUsClient) getHandleResponse(resp *http.Response) (VendorSKUsClientGetResponse, error) {
	result := VendorSKUsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VendorSKU); err != nil {
		return VendorSKUsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the skus of a vendor.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// vendorName - The name of the vendor.
// options - VendorSKUsClientListOptions contains the optional parameters for the VendorSKUsClient.List method.
func (client *VendorSKUsClient) NewListPager(vendorName string, options *VendorSKUsClientListOptions) *runtime.Pager[VendorSKUsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[VendorSKUsClientListResponse]{
		More: func(page VendorSKUsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VendorSKUsClientListResponse) (VendorSKUsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, vendorName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VendorSKUsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VendorSKUsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VendorSKUsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VendorSKUsClient) listCreateRequest(ctx context.Context, vendorName string, options *VendorSKUsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/vendors/{vendorName}/vendorSkus"
	if vendorName == "" {
		return nil, errors.New("parameter vendorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vendorName}", url.PathEscape(vendorName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VendorSKUsClient) listHandleResponse(resp *http.Response) (VendorSKUsClientListResponse, error) {
	result := VendorSKUsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VendorSKUListResult); err != nil {
		return VendorSKUsClientListResponse{}, err
	}
	return result, nil
}

// ListCredential - Generate credentials for publishing SKU images.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-01-preview
// vendorName - The name of the vendor.
// skuName - The name of the sku.
// options - VendorSKUsClientListCredentialOptions contains the optional parameters for the VendorSKUsClient.ListCredential
// method.
func (client *VendorSKUsClient) ListCredential(ctx context.Context, vendorName string, skuName string, options *VendorSKUsClientListCredentialOptions) (VendorSKUsClientListCredentialResponse, error) {
	req, err := client.listCredentialCreateRequest(ctx, vendorName, skuName, options)
	if err != nil {
		return VendorSKUsClientListCredentialResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VendorSKUsClientListCredentialResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VendorSKUsClientListCredentialResponse{}, runtime.NewResponseError(resp)
	}
	return client.listCredentialHandleResponse(resp)
}

// listCredentialCreateRequest creates the ListCredential request.
func (client *VendorSKUsClient) listCredentialCreateRequest(ctx context.Context, vendorName string, skuName string, options *VendorSKUsClientListCredentialOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/vendors/{vendorName}/vendorSkus/{skuName}/listCredential"
	if vendorName == "" {
		return nil, errors.New("parameter vendorName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vendorName}", url.PathEscape(vendorName))
	if skuName == "" {
		return nil, errors.New("parameter skuName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{skuName}", url.PathEscape(skuName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listCredentialHandleResponse handles the ListCredential response.
func (client *VendorSKUsClient) listCredentialHandleResponse(resp *http.Response) (VendorSKUsClientListCredentialResponse, error) {
	result := VendorSKUsClientListCredentialResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SKUCredential); err != nil {
		return VendorSKUsClientListCredentialResponse{}, err
	}
	return result, nil
}
