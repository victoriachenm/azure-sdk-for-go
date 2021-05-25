package migrate

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// SitesClient is the discover your workloads for Azure.
type SitesClient struct {
	BaseClient
}

// NewSitesClient creates an instance of the SitesClient client.
func NewSitesClient() SitesClient {
	return NewSitesClientWithBaseURI(DefaultBaseURI)
}

// NewSitesClientWithBaseURI creates an instance of the SitesClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewSitesClientWithBaseURI(baseURI string) SitesClient {
	return SitesClient{NewWithBaseURI(baseURI)}
}

// DeleteSite sends the delete site request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// APIVersion - the API version to use for this operation.
func (client SitesClient) DeleteSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SitesClient.DeleteSite")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeleteSitePreparer(ctx, subscriptionID, resourceGroupName, siteName, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.SitesClient", "DeleteSite", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSiteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "migrate.SitesClient", "DeleteSite", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteSiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.SitesClient", "DeleteSite", resp, "Failure responding to request")
		return
	}

	return
}

// DeleteSitePreparer prepares the DeleteSite request.
func (client SitesClient) DeleteSitePreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSiteSender sends the DeleteSite request. The method will close the
// http.Response Body if it receives an error.
func (client SitesClient) DeleteSiteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteSiteResponder handles the response to the DeleteSite request. The method always
// closes the http.Response Body.
func (client SitesClient) DeleteSiteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// GetSite sends the get site request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// APIVersion - the API version to use for this operation.
func (client SitesClient) GetSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result VMwareSite, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SitesClient.GetSite")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetSitePreparer(ctx, subscriptionID, resourceGroupName, siteName, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.SitesClient", "GetSite", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSiteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.SitesClient", "GetSite", resp, "Failure sending request")
		return
	}

	result, err = client.GetSiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.SitesClient", "GetSite", resp, "Failure responding to request")
		return
	}

	return
}

// GetSitePreparer prepares the GetSite request.
func (client SitesClient) GetSitePreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSiteSender sends the GetSite request. The method will close the
// http.Response Body if it receives an error.
func (client SitesClient) GetSiteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetSiteResponder handles the response to the GetSite request. The method always
// closes the http.Response Body.
func (client SitesClient) GetSiteResponder(resp *http.Response) (result VMwareSite, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetSiteHealthSummary sends the get site health summary request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// APIVersion - the API version to use for this operation.
func (client SitesClient) GetSiteHealthSummary(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result SiteHealthSummaryCollectionPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SitesClient.GetSiteHealthSummary")
		defer func() {
			sc := -1
			if result.shsc.Response.Response != nil {
				sc = result.shsc.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.getSiteHealthSummaryNextResults
	req, err := client.GetSiteHealthSummaryPreparer(ctx, subscriptionID, resourceGroupName, siteName, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.SitesClient", "GetSiteHealthSummary", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSiteHealthSummarySender(req)
	if err != nil {
		result.shsc.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.SitesClient", "GetSiteHealthSummary", resp, "Failure sending request")
		return
	}

	result.shsc, err = client.GetSiteHealthSummaryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.SitesClient", "GetSiteHealthSummary", resp, "Failure responding to request")
		return
	}
	if result.shsc.hasNextLink() && result.shsc.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// GetSiteHealthSummaryPreparer prepares the GetSiteHealthSummary request.
func (client SitesClient) GetSiteHealthSummaryPreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}/healthSummary", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSiteHealthSummarySender sends the GetSiteHealthSummary request. The method will close the
// http.Response Body if it receives an error.
func (client SitesClient) GetSiteHealthSummarySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetSiteHealthSummaryResponder handles the response to the GetSiteHealthSummary request. The method always
// closes the http.Response Body.
func (client SitesClient) GetSiteHealthSummaryResponder(resp *http.Response) (result SiteHealthSummaryCollection, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// getSiteHealthSummaryNextResults retrieves the next set of results, if any.
func (client SitesClient) getSiteHealthSummaryNextResults(ctx context.Context, lastResults SiteHealthSummaryCollection) (result SiteHealthSummaryCollection, err error) {
	req, err := lastResults.siteHealthSummaryCollectionPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "migrate.SitesClient", "getSiteHealthSummaryNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.GetSiteHealthSummarySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "migrate.SitesClient", "getSiteHealthSummaryNextResults", resp, "Failure sending next results request")
	}
	result, err = client.GetSiteHealthSummaryResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.SitesClient", "getSiteHealthSummaryNextResults", resp, "Failure responding to next results request")
	}
	return
}

// GetSiteHealthSummaryComplete enumerates all values, automatically crossing page boundaries as required.
func (client SitesClient) GetSiteHealthSummaryComplete(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result SiteHealthSummaryCollectionIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SitesClient.GetSiteHealthSummary")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.GetSiteHealthSummary(ctx, subscriptionID, resourceGroupName, siteName, APIVersion)
	return
}

// GetSiteUsage sends the get site usage request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
func (client SitesClient) GetSiteUsage(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result VMwareSiteUsage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SitesClient.GetSiteUsage")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetSiteUsagePreparer(ctx, subscriptionID, resourceGroupName, siteName, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.SitesClient", "GetSiteUsage", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSiteUsageSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.SitesClient", "GetSiteUsage", resp, "Failure sending request")
		return
	}

	result, err = client.GetSiteUsageResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.SitesClient", "GetSiteUsage", resp, "Failure responding to request")
		return
	}

	return
}

// GetSiteUsagePreparer prepares the GetSiteUsage request.
func (client SitesClient) GetSiteUsagePreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}/summary", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSiteUsageSender sends the GetSiteUsage request. The method will close the
// http.Response Body if it receives an error.
func (client SitesClient) GetSiteUsageSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetSiteUsageResponder handles the response to the GetSiteUsage request. The method always
// closes the http.Response Body.
func (client SitesClient) GetSiteUsageResponder(resp *http.Response) (result VMwareSiteUsage, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PatchSite sends the patch site request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// body - body with site details.
// APIVersion - the API version to use for this operation.
func (client SitesClient) PatchSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, body VMwareSite, APIVersion string) (result VMwareSite, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SitesClient.PatchSite")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PatchSitePreparer(ctx, subscriptionID, resourceGroupName, siteName, body, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.SitesClient", "PatchSite", nil, "Failure preparing request")
		return
	}

	resp, err := client.PatchSiteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.SitesClient", "PatchSite", resp, "Failure sending request")
		return
	}

	result, err = client.PatchSiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.SitesClient", "PatchSite", resp, "Failure responding to request")
		return
	}

	return
}

// PatchSitePreparer prepares the PatchSite request.
func (client SitesClient) PatchSitePreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, body VMwareSite, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	body.ID = nil
	body.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PatchSiteSender sends the PatchSite request. The method will close the
// http.Response Body if it receives an error.
func (client SitesClient) PatchSiteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// PatchSiteResponder handles the response to the PatchSite request. The method always
// closes the http.Response Body.
func (client SitesClient) PatchSiteResponder(resp *http.Response) (result VMwareSite, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PutSite sends the put site request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// body - body with site details.
// APIVersion - the API version to use for this operation.
func (client SitesClient) PutSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, body VMwareSite, APIVersion string) (result VMwareSite, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SitesClient.PutSite")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.PutSitePreparer(ctx, subscriptionID, resourceGroupName, siteName, body, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.SitesClient", "PutSite", nil, "Failure preparing request")
		return
	}

	resp, err := client.PutSiteSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "migrate.SitesClient", "PutSite", resp, "Failure sending request")
		return
	}

	result, err = client.PutSiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.SitesClient", "PutSite", resp, "Failure responding to request")
		return
	}

	return
}

// PutSitePreparer prepares the PutSite request.
func (client SitesClient) PutSitePreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, body VMwareSite, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	body.ID = nil
	body.Type = nil
	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}", pathParameters),
		autorest.WithJSON(body),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PutSiteSender sends the PutSite request. The method will close the
// http.Response Body if it receives an error.
func (client SitesClient) PutSiteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// PutSiteResponder handles the response to the PutSite request. The method always
// closes the http.Response Body.
func (client SitesClient) PutSiteResponder(resp *http.Response) (result VMwareSite, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// RefreshSite sends the refresh site request.
// Parameters:
// subscriptionID - the ID of the target subscription.
// resourceGroupName - the name of the resource group. The name is case insensitive.
// siteName - site name.
// APIVersion - the API version to use for this operation.
func (client SitesClient) RefreshSite(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SitesClient.RefreshSite")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.RefreshSitePreparer(ctx, subscriptionID, resourceGroupName, siteName, APIVersion)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.SitesClient", "RefreshSite", nil, "Failure preparing request")
		return
	}

	resp, err := client.RefreshSiteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "migrate.SitesClient", "RefreshSite", resp, "Failure sending request")
		return
	}

	result, err = client.RefreshSiteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "migrate.SitesClient", "RefreshSite", resp, "Failure responding to request")
		return
	}

	return
}

// RefreshSitePreparer prepares the RefreshSite request.
func (client SitesClient) RefreshSitePreparer(ctx context.Context, subscriptionID string, resourceGroupName string, siteName string, APIVersion string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"siteName":          autorest.Encode("path", siteName),
		"subscriptionId":    autorest.Encode("path", subscriptionID),
	}

	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OffAzure/VMwareSites/{siteName}/refresh", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RefreshSiteSender sends the RefreshSite request. The method will close the
// http.Response Body if it receives an error.
func (client SitesClient) RefreshSiteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// RefreshSiteResponder handles the response to the RefreshSite request. The method always
// closes the http.Response Body.
func (client SitesClient) RefreshSiteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByClosing())
	result.Response = resp
	return
}