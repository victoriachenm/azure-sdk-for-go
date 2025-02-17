//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

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

// ConnectionTypeClient contains the methods for the ConnectionType group.
// Don't use this type directly, use NewConnectionTypeClient() instead.
type ConnectionTypeClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewConnectionTypeClient creates a new instance of ConnectionTypeClient with the specified values.
// subscriptionID - Gets subscription credentials which uniquely identify Microsoft Azure subscription. The subscription ID
// forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewConnectionTypeClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConnectionTypeClient, error) {
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
	client := &ConnectionTypeClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create a connection type.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-13-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// connectionTypeName - The parameters supplied to the create or update connection type operation.
// parameters - The parameters supplied to the create or update connection type operation.
// options - ConnectionTypeClientCreateOrUpdateOptions contains the optional parameters for the ConnectionTypeClient.CreateOrUpdate
// method.
func (client *ConnectionTypeClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, automationAccountName string, connectionTypeName string, parameters ConnectionTypeCreateOrUpdateParameters, options *ConnectionTypeClientCreateOrUpdateOptions) (ConnectionTypeClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, automationAccountName, connectionTypeName, parameters, options)
	if err != nil {
		return ConnectionTypeClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectionTypeClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return ConnectionTypeClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConnectionTypeClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, connectionTypeName string, parameters ConnectionTypeCreateOrUpdateParameters, options *ConnectionTypeClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connectionTypes/{connectionTypeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if connectionTypeName == "" {
		return nil, errors.New("parameter connectionTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionTypeName}", url.PathEscape(connectionTypeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ConnectionTypeClient) createOrUpdateHandleResponse(resp *http.Response) (ConnectionTypeClientCreateOrUpdateResponse, error) {
	result := ConnectionTypeClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionType); err != nil {
		return ConnectionTypeClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete the connection type.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-13-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// connectionTypeName - The name of connection type.
// options - ConnectionTypeClientDeleteOptions contains the optional parameters for the ConnectionTypeClient.Delete method.
func (client *ConnectionTypeClient) Delete(ctx context.Context, resourceGroupName string, automationAccountName string, connectionTypeName string, options *ConnectionTypeClientDeleteOptions) (ConnectionTypeClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, automationAccountName, connectionTypeName, options)
	if err != nil {
		return ConnectionTypeClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectionTypeClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return ConnectionTypeClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return ConnectionTypeClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConnectionTypeClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, connectionTypeName string, options *ConnectionTypeClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connectionTypes/{connectionTypeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if connectionTypeName == "" {
		return nil, errors.New("parameter connectionTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionTypeName}", url.PathEscape(connectionTypeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Retrieve the connection type identified by connection type name.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-13-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// connectionTypeName - The name of connection type.
// options - ConnectionTypeClientGetOptions contains the optional parameters for the ConnectionTypeClient.Get method.
func (client *ConnectionTypeClient) Get(ctx context.Context, resourceGroupName string, automationAccountName string, connectionTypeName string, options *ConnectionTypeClientGetOptions) (ConnectionTypeClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, automationAccountName, connectionTypeName, options)
	if err != nil {
		return ConnectionTypeClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ConnectionTypeClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConnectionTypeClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConnectionTypeClient) getCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, connectionTypeName string, options *ConnectionTypeClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connectionTypes/{connectionTypeName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if connectionTypeName == "" {
		return nil, errors.New("parameter connectionTypeName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionTypeName}", url.PathEscape(connectionTypeName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ConnectionTypeClient) getHandleResponse(resp *http.Response) (ConnectionTypeClientGetResponse, error) {
	result := ConnectionTypeClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionType); err != nil {
		return ConnectionTypeClientGetResponse{}, err
	}
	return result, nil
}

// NewListByAutomationAccountPager - Retrieve a list of connection types.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2020-01-13-preview
// resourceGroupName - Name of an Azure Resource group.
// automationAccountName - The name of the automation account.
// options - ConnectionTypeClientListByAutomationAccountOptions contains the optional parameters for the ConnectionTypeClient.ListByAutomationAccount
// method.
func (client *ConnectionTypeClient) NewListByAutomationAccountPager(resourceGroupName string, automationAccountName string, options *ConnectionTypeClientListByAutomationAccountOptions) *runtime.Pager[ConnectionTypeClientListByAutomationAccountResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConnectionTypeClientListByAutomationAccountResponse]{
		More: func(page ConnectionTypeClientListByAutomationAccountResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConnectionTypeClientListByAutomationAccountResponse) (ConnectionTypeClientListByAutomationAccountResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByAutomationAccountCreateRequest(ctx, resourceGroupName, automationAccountName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ConnectionTypeClientListByAutomationAccountResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ConnectionTypeClientListByAutomationAccountResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConnectionTypeClientListByAutomationAccountResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByAutomationAccountHandleResponse(resp)
		},
	})
}

// listByAutomationAccountCreateRequest creates the ListByAutomationAccount request.
func (client *ConnectionTypeClient) listByAutomationAccountCreateRequest(ctx context.Context, resourceGroupName string, automationAccountName string, options *ConnectionTypeClientListByAutomationAccountOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Automation/automationAccounts/{automationAccountName}/connectionTypes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if automationAccountName == "" {
		return nil, errors.New("parameter automationAccountName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{automationAccountName}", url.PathEscape(automationAccountName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-01-13-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByAutomationAccountHandleResponse handles the ListByAutomationAccount response.
func (client *ConnectionTypeClient) listByAutomationAccountHandleResponse(resp *http.Response) (ConnectionTypeClientListByAutomationAccountResponse, error) {
	result := ConnectionTypeClientListByAutomationAccountResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectionTypeListResult); err != nil {
		return ConnectionTypeClientListByAutomationAccountResponse{}, err
	}
	return result, nil
}
