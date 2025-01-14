//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtrafficmanager_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/trafficmanager/armtrafficmanager"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-08-01/examples/Endpoint-PATCH-External-Target.json
func ExampleEndpointsClient_Update() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armtrafficmanager.NewEndpointsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Update(ctx,
		"azuresdkfornetautoresttrafficmanager1421",
		"azsmnet6386",
		armtrafficmanager.EndpointTypeExternalEndpoints,
		"azsmnet7187",
		armtrafficmanager.Endpoint{
			Name: to.Ptr("azsmnet7187"),
			Type: to.Ptr("Microsoft.Network/trafficManagerProfiles/externalEndpoints"),
			ID:   to.Ptr("/subscriptions/{subscription-id}/resourceGroups/azuresdkfornetautoresttrafficmanager1421/providers/Microsoft.Network/trafficManagerProfiles/azsmnet6386/externalEndpoints/azsmnet7187"),
			Properties: &armtrafficmanager.EndpointProperties{
				Target: to.Ptr("another.foobar.contoso.com"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-08-01/examples/Endpoint-GET-External-WithGeoMapping.json
func ExampleEndpointsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armtrafficmanager.NewEndpointsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"azuresdkfornetautoresttrafficmanager2191",
		"azuresdkfornetautoresttrafficmanager8224",
		armtrafficmanager.EndpointTypeExternalEndpoints,
		"My%20external%20endpoint",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-08-01/examples/Endpoint-PUT-External-WithCustomHeaders.json
func ExampleEndpointsClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armtrafficmanager.NewEndpointsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.CreateOrUpdate(ctx,
		"azuresdkfornetautoresttrafficmanager1421",
		"azsmnet6386",
		armtrafficmanager.EndpointTypeExternalEndpoints,
		"azsmnet7187",
		armtrafficmanager.Endpoint{
			Name: to.Ptr("azsmnet7187"),
			Type: to.Ptr("Microsoft.network/TrafficManagerProfiles/ExternalEndpoints"),
			Properties: &armtrafficmanager.EndpointProperties{
				CustomHeaders: []*armtrafficmanager.EndpointPropertiesCustomHeadersItem{
					{
						Name:  to.Ptr("header-1"),
						Value: to.Ptr("value-1"),
					},
					{
						Name:  to.Ptr("header-2"),
						Value: to.Ptr("value-2"),
					}},
				EndpointLocation: to.Ptr("North Europe"),
				EndpointStatus:   to.Ptr(armtrafficmanager.EndpointStatusEnabled),
				Target:           to.Ptr("foobar.contoso.com"),
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-08-01/examples/Endpoint-DELETE-External.json
func ExampleEndpointsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armtrafficmanager.NewEndpointsClient("{subscription-id}", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Delete(ctx,
		"azuresdkfornetautoresttrafficmanager1421",
		"azsmnet6386",
		armtrafficmanager.EndpointTypeExternalEndpoints,
		"azsmnet7187",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}
