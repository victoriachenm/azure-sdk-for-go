// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azcosmos

import (
	"context"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
)

// cosmosClientConnection maintains a Pipeline for the client.
// The Pipeline is build based on the CosmosClientOptions.
type cosmosClientConnection struct {
	endpoint string
	Pipeline azcore.Pipeline
}

// newConnection creates an instance of the connection type with the specified endpoint.
// Pass nil to accept the default options; this is the same as passing a zero-value options.
func newCosmosClientConnection(endpoint string, cred azcore.Credential, options *CosmosClientOptions) *cosmosClientConnection {
	if options == nil {
		options = &CosmosClientOptions{}
	}
	policies := []azcore.Policy{
		azcore.NewTelemetryPolicy(options.enrichTelemetryOptions()),
	}
	policies = append(policies, options.PerCallPolicies...)
	policies = append(policies, azcore.NewRetryPolicy(&options.Retry))
	policies = append(policies, options.PerRetryPolicies...)
	policies = append(policies, options.getSDKInternalPolicies()...)
	policies = append(policies, cred.AuthenticationPolicy(azcore.AuthenticationPolicyOptions{Options: azcore.TokenRequestOptions{Scopes: []string{"none"}}}))
	policies = append(policies, azcore.NewLogPolicy(&options.Logging))
	return &cosmosClientConnection{endpoint: endpoint, Pipeline: azcore.NewPipeline(options.HTTPClient, policies...)}
}

func (c *cosmosClientConnection) sendPostRequest(
	path string,
	ctx context.Context,
	content interface{},
	operationContext cosmosOperationContext,
	requestOptions cosmosRequestOptions) (*azcore.Response, error) {
	req, err := c.createRequest(path, ctx, http.MethodPost, operationContext, requestOptions)
	if err != nil {
		return nil, err
	}

	err = req.MarshalAsJSON(content)
	if err != nil {
		return nil, err
	}

	return c.Pipeline.Do(req)
}

func (c *cosmosClientConnection) sendPutRequest(
	path string,
	ctx context.Context,
	content interface{},
	operationContext cosmosOperationContext,
	requestOptions cosmosRequestOptions) (*azcore.Response, error) {
	req, err := c.createRequest(path, ctx, http.MethodPut, operationContext, requestOptions)
	if err != nil {
		return nil, err
	}

	err = req.MarshalAsJSON(content)
	if err != nil {
		return nil, err
	}

	return c.Pipeline.Do(req)
}

func (c *cosmosClientConnection) sendGetRequest(
	path string,
	ctx context.Context,
	operationContext cosmosOperationContext,
	requestOptions cosmosRequestOptions) (*azcore.Response, error) {
	req, err := c.createRequest(path, ctx, http.MethodGet, operationContext, requestOptions)
	if err != nil {
		return nil, err
	}

	return c.Pipeline.Do(req)
}

func (c *cosmosClientConnection) sendDeleteRequest(
	path string,
	ctx context.Context,
	operationContext cosmosOperationContext,
	requestOptions cosmosRequestOptions) (*azcore.Response, error) {
	req, err := c.createRequest(path, ctx, http.MethodDelete, operationContext, requestOptions)
	if err != nil {
		return nil, err
	}

	return c.Pipeline.Do(req)
}

func (c *cosmosClientConnection) createRequest(
	path string,
	ctx context.Context,
	method string,
	operationContext cosmosOperationContext,
	requestOptions cosmosRequestOptions) (*azcore.Request, error) {

	// todo: endpoint will be set originally by globalendpointmanager
	finalURL := c.endpoint

	if path != "" {
		finalURL = azcore.JoinPaths(c.endpoint, path)
	}

	req, err := azcore.NewRequest(ctx, method, finalURL)
	if err != nil {
		return nil, err
	}

	headers := requestOptions.toHeaders()
	if headers != nil {
		for k, v := range *headers {
			req.Header.Set(k, v)
		}
	}

	req.Request.Header.Set(azcore.HeaderXmsDate, time.Now().UTC().Format(http.TimeFormat))
	req.Request.Header.Set(azcore.HeaderXmsVersion, "2020-11-05")

	req.SetOperationValue(operationContext)
	return req, nil
}