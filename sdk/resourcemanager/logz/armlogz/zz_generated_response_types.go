//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogz

import (
	"context"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"net/http"
	"time"
)

// MonitorClientListVMHostUpdateResponse contains the response from method MonitorClient.ListVMHostUpdate.
type MonitorClientListVMHostUpdateResponse struct {
	MonitorClientListVMHostUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorClientListVMHostUpdateResult contains the result from method MonitorClient.ListVMHostUpdate.
type MonitorClientListVMHostUpdateResult struct {
	VMResourcesListResponse
}

// MonitorClientListVMHostsResponse contains the response from method MonitorClient.ListVMHosts.
type MonitorClientListVMHostsResponse struct {
	MonitorClientListVMHostsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorClientListVMHostsResult contains the result from method MonitorClient.ListVMHosts.
type MonitorClientListVMHostsResult struct {
	VMResourcesListResponse
}

// MonitorClientVMHostPayloadResponse contains the response from method MonitorClient.VMHostPayload.
type MonitorClientVMHostPayloadResponse struct {
	MonitorClientVMHostPayloadResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorClientVMHostPayloadResult contains the result from method MonitorClient.VMHostPayload.
type MonitorClientVMHostPayloadResult struct {
	VMExtensionPayload
}

// MonitorsClientCreatePollerResponse contains the response from method MonitorsClient.Create.
type MonitorsClientCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *MonitorsClientCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l MonitorsClientCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (MonitorsClientCreateResponse, error) {
	respType := MonitorsClientCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.MonitorResource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a MonitorsClientCreatePollerResponse from the provided client and resume token.
func (l *MonitorsClientCreatePollerResponse) Resume(ctx context.Context, client *MonitorsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("MonitorsClient.Create", token, client.pl)
	if err != nil {
		return err
	}
	poller := &MonitorsClientCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// MonitorsClientCreateResponse contains the response from method MonitorsClient.Create.
type MonitorsClientCreateResponse struct {
	MonitorsClientCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsClientCreateResult contains the result from method MonitorsClient.Create.
type MonitorsClientCreateResult struct {
	MonitorResource
}

// MonitorsClientDeletePollerResponse contains the response from method MonitorsClient.Delete.
type MonitorsClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *MonitorsClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l MonitorsClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (MonitorsClientDeleteResponse, error) {
	respType := MonitorsClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a MonitorsClientDeletePollerResponse from the provided client and resume token.
func (l *MonitorsClientDeletePollerResponse) Resume(ctx context.Context, client *MonitorsClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("MonitorsClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &MonitorsClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// MonitorsClientDeleteResponse contains the response from method MonitorsClient.Delete.
type MonitorsClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsClientGetResponse contains the response from method MonitorsClient.Get.
type MonitorsClientGetResponse struct {
	MonitorsClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsClientGetResult contains the result from method MonitorsClient.Get.
type MonitorsClientGetResult struct {
	MonitorResource
}

// MonitorsClientListByResourceGroupResponse contains the response from method MonitorsClient.ListByResourceGroup.
type MonitorsClientListByResourceGroupResponse struct {
	MonitorsClientListByResourceGroupResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsClientListByResourceGroupResult contains the result from method MonitorsClient.ListByResourceGroup.
type MonitorsClientListByResourceGroupResult struct {
	MonitorResourceListResponse
}

// MonitorsClientListBySubscriptionResponse contains the response from method MonitorsClient.ListBySubscription.
type MonitorsClientListBySubscriptionResponse struct {
	MonitorsClientListBySubscriptionResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsClientListBySubscriptionResult contains the result from method MonitorsClient.ListBySubscription.
type MonitorsClientListBySubscriptionResult struct {
	MonitorResourceListResponse
}

// MonitorsClientListMonitoredResourcesResponse contains the response from method MonitorsClient.ListMonitoredResources.
type MonitorsClientListMonitoredResourcesResponse struct {
	MonitorsClientListMonitoredResourcesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsClientListMonitoredResourcesResult contains the result from method MonitorsClient.ListMonitoredResources.
type MonitorsClientListMonitoredResourcesResult struct {
	MonitoredResourceListResponse
}

// MonitorsClientListUserRolesResponse contains the response from method MonitorsClient.ListUserRoles.
type MonitorsClientListUserRolesResponse struct {
	MonitorsClientListUserRolesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsClientListUserRolesResult contains the result from method MonitorsClient.ListUserRoles.
type MonitorsClientListUserRolesResult struct {
	UserRoleListResponse
}

// MonitorsClientUpdateResponse contains the response from method MonitorsClient.Update.
type MonitorsClientUpdateResponse struct {
	MonitorsClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// MonitorsClientUpdateResult contains the result from method MonitorsClient.Update.
type MonitorsClientUpdateResult struct {
	MonitorResource
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationsClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// OperationsClientListResult contains the result from method OperationsClient.List.
type OperationsClientListResult struct {
	OperationListResult
}

// SingleSignOnClientCreateOrUpdatePollerResponse contains the response from method SingleSignOnClient.CreateOrUpdate.
type SingleSignOnClientCreateOrUpdatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SingleSignOnClientCreateOrUpdatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l SingleSignOnClientCreateOrUpdatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SingleSignOnClientCreateOrUpdateResponse, error) {
	respType := SingleSignOnClientCreateOrUpdateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.SingleSignOnResource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SingleSignOnClientCreateOrUpdatePollerResponse from the provided client and resume token.
func (l *SingleSignOnClientCreateOrUpdatePollerResponse) Resume(ctx context.Context, client *SingleSignOnClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SingleSignOnClient.CreateOrUpdate", token, client.pl)
	if err != nil {
		return err
	}
	poller := &SingleSignOnClientCreateOrUpdatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SingleSignOnClientCreateOrUpdateResponse contains the response from method SingleSignOnClient.CreateOrUpdate.
type SingleSignOnClientCreateOrUpdateResponse struct {
	SingleSignOnClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SingleSignOnClientCreateOrUpdateResult contains the result from method SingleSignOnClient.CreateOrUpdate.
type SingleSignOnClientCreateOrUpdateResult struct {
	SingleSignOnResource
}

// SingleSignOnClientGetResponse contains the response from method SingleSignOnClient.Get.
type SingleSignOnClientGetResponse struct {
	SingleSignOnClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SingleSignOnClientGetResult contains the result from method SingleSignOnClient.Get.
type SingleSignOnClientGetResult struct {
	SingleSignOnResource
}

// SingleSignOnClientListResponse contains the response from method SingleSignOnClient.List.
type SingleSignOnClientListResponse struct {
	SingleSignOnClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SingleSignOnClientListResult contains the result from method SingleSignOnClient.List.
type SingleSignOnClientListResult struct {
	SingleSignOnResourceListResponse
}

// SubAccountClientCreatePollerResponse contains the response from method SubAccountClient.Create.
type SubAccountClientCreatePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SubAccountClientCreatePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l SubAccountClientCreatePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SubAccountClientCreateResponse, error) {
	respType := SubAccountClientCreateResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, &respType.MonitorResource)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SubAccountClientCreatePollerResponse from the provided client and resume token.
func (l *SubAccountClientCreatePollerResponse) Resume(ctx context.Context, client *SubAccountClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SubAccountClient.Create", token, client.pl)
	if err != nil {
		return err
	}
	poller := &SubAccountClientCreatePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SubAccountClientCreateResponse contains the response from method SubAccountClient.Create.
type SubAccountClientCreateResponse struct {
	SubAccountClientCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubAccountClientCreateResult contains the result from method SubAccountClient.Create.
type SubAccountClientCreateResult struct {
	MonitorResource
}

// SubAccountClientDeletePollerResponse contains the response from method SubAccountClient.Delete.
type SubAccountClientDeletePollerResponse struct {
	// Poller contains an initialized poller.
	Poller *SubAccountClientDeletePoller

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PollUntilDone will poll the service endpoint until a terminal state is reached or an error is received.
// freq: the time to wait between intervals in absence of a Retry-After header. Allowed minimum is one second.
// A good starting value is 30 seconds. Note that some resources might benefit from a different value.
func (l SubAccountClientDeletePollerResponse) PollUntilDone(ctx context.Context, freq time.Duration) (SubAccountClientDeleteResponse, error) {
	respType := SubAccountClientDeleteResponse{}
	resp, err := l.Poller.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return respType, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// Resume rehydrates a SubAccountClientDeletePollerResponse from the provided client and resume token.
func (l *SubAccountClientDeletePollerResponse) Resume(ctx context.Context, client *SubAccountClient, token string) error {
	pt, err := armruntime.NewPollerFromResumeToken("SubAccountClient.Delete", token, client.pl)
	if err != nil {
		return err
	}
	poller := &SubAccountClientDeletePoller{
		pt: pt,
	}
	resp, err := poller.Poll(ctx)
	if err != nil {
		return err
	}
	l.Poller = poller
	l.RawResponse = resp
	return nil
}

// SubAccountClientDeleteResponse contains the response from method SubAccountClient.Delete.
type SubAccountClientDeleteResponse struct {
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubAccountClientGetResponse contains the response from method SubAccountClient.Get.
type SubAccountClientGetResponse struct {
	SubAccountClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubAccountClientGetResult contains the result from method SubAccountClient.Get.
type SubAccountClientGetResult struct {
	MonitorResource
}

// SubAccountClientListMonitoredResourcesResponse contains the response from method SubAccountClient.ListMonitoredResources.
type SubAccountClientListMonitoredResourcesResponse struct {
	SubAccountClientListMonitoredResourcesResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubAccountClientListMonitoredResourcesResult contains the result from method SubAccountClient.ListMonitoredResources.
type SubAccountClientListMonitoredResourcesResult struct {
	MonitoredResourceListResponse
}

// SubAccountClientListResponse contains the response from method SubAccountClient.List.
type SubAccountClientListResponse struct {
	SubAccountClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubAccountClientListResult contains the result from method SubAccountClient.List.
type SubAccountClientListResult struct {
	MonitorResourceListResponse
}

// SubAccountClientListVMHostUpdateResponse contains the response from method SubAccountClient.ListVMHostUpdate.
type SubAccountClientListVMHostUpdateResponse struct {
	SubAccountClientListVMHostUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubAccountClientListVMHostUpdateResult contains the result from method SubAccountClient.ListVMHostUpdate.
type SubAccountClientListVMHostUpdateResult struct {
	VMResourcesListResponse
}

// SubAccountClientListVMHostsResponse contains the response from method SubAccountClient.ListVMHosts.
type SubAccountClientListVMHostsResponse struct {
	SubAccountClientListVMHostsResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubAccountClientListVMHostsResult contains the result from method SubAccountClient.ListVMHosts.
type SubAccountClientListVMHostsResult struct {
	VMResourcesListResponse
}

// SubAccountClientUpdateResponse contains the response from method SubAccountClient.Update.
type SubAccountClientUpdateResponse struct {
	SubAccountClientUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubAccountClientUpdateResult contains the result from method SubAccountClient.Update.
type SubAccountClientUpdateResult struct {
	MonitorResource
}

// SubAccountClientVMHostPayloadResponse contains the response from method SubAccountClient.VMHostPayload.
type SubAccountClientVMHostPayloadResponse struct {
	SubAccountClientVMHostPayloadResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubAccountClientVMHostPayloadResult contains the result from method SubAccountClient.VMHostPayload.
type SubAccountClientVMHostPayloadResult struct {
	VMExtensionPayload
}

// SubAccountTagRulesClientCreateOrUpdateResponse contains the response from method SubAccountTagRulesClient.CreateOrUpdate.
type SubAccountTagRulesClientCreateOrUpdateResponse struct {
	SubAccountTagRulesClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubAccountTagRulesClientCreateOrUpdateResult contains the result from method SubAccountTagRulesClient.CreateOrUpdate.
type SubAccountTagRulesClientCreateOrUpdateResult struct {
	MonitoringTagRules
}

// SubAccountTagRulesClientDeleteResponse contains the response from method SubAccountTagRulesClient.Delete.
type SubAccountTagRulesClientDeleteResponse struct {
	SubAccountTagRulesClientDeleteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubAccountTagRulesClientDeleteResult contains the result from method SubAccountTagRulesClient.Delete.
type SubAccountTagRulesClientDeleteResult struct {
	// Location contains the information returned from the location header response.
	Location *string
}

// SubAccountTagRulesClientGetResponse contains the response from method SubAccountTagRulesClient.Get.
type SubAccountTagRulesClientGetResponse struct {
	SubAccountTagRulesClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubAccountTagRulesClientGetResult contains the result from method SubAccountTagRulesClient.Get.
type SubAccountTagRulesClientGetResult struct {
	MonitoringTagRules
}

// SubAccountTagRulesClientListResponse contains the response from method SubAccountTagRulesClient.List.
type SubAccountTagRulesClientListResponse struct {
	SubAccountTagRulesClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// SubAccountTagRulesClientListResult contains the result from method SubAccountTagRulesClient.List.
type SubAccountTagRulesClientListResult struct {
	MonitoringTagRulesListResponse
}

// TagRulesClientCreateOrUpdateResponse contains the response from method TagRulesClient.CreateOrUpdate.
type TagRulesClientCreateOrUpdateResponse struct {
	TagRulesClientCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TagRulesClientCreateOrUpdateResult contains the result from method TagRulesClient.CreateOrUpdate.
type TagRulesClientCreateOrUpdateResult struct {
	MonitoringTagRules
}

// TagRulesClientDeleteResponse contains the response from method TagRulesClient.Delete.
type TagRulesClientDeleteResponse struct {
	TagRulesClientDeleteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TagRulesClientDeleteResult contains the result from method TagRulesClient.Delete.
type TagRulesClientDeleteResult struct {
	// Location contains the information returned from the location header response.
	Location *string
}

// TagRulesClientGetResponse contains the response from method TagRulesClient.Get.
type TagRulesClientGetResponse struct {
	TagRulesClientGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TagRulesClientGetResult contains the result from method TagRulesClient.Get.
type TagRulesClientGetResult struct {
	MonitoringTagRules
}

// TagRulesClientListResponse contains the response from method TagRulesClient.List.
type TagRulesClientListResponse struct {
	TagRulesClientListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// TagRulesClientListResult contains the result from method TagRulesClient.List.
type TagRulesClientListResult struct {
	MonitoringTagRulesListResponse
}