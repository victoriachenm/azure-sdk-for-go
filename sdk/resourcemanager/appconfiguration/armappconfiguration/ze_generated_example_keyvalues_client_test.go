//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armappconfiguration_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/appconfiguration/armappconfiguration"
)

// x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-03-01-preview/examples/ConfigurationStoresListKeyValues.json
func ExampleKeyValuesClient_ListByConfigurationStore() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappconfiguration.NewKeyValuesClient("<subscription-id>", cred, nil)
	pager := client.ListByConfigurationStore("<resource-group-name>",
		"<config-store-name>",
		&armappconfiguration.KeyValuesListByConfigurationStoreOptions{SkipToken: nil})
	for pager.NextPage(ctx) {
		if err := pager.Err(); err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range pager.PageResponse().Value {
			log.Printf("KeyValue.ID: %s\n", *v.ID)
		}
	}
}

// x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-03-01-preview/examples/ConfigurationStoresGetKeyValue.json
func ExampleKeyValuesClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappconfiguration.NewKeyValuesClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<config-store-name>",
		"<key-value-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("KeyValue.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-03-01-preview/examples/ConfigurationStoresCreateKeyValue.json
func ExampleKeyValuesClient_CreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappconfiguration.NewKeyValuesClient("<subscription-id>", cred, nil)
	res, err := client.CreateOrUpdate(ctx,
		"<resource-group-name>",
		"<config-store-name>",
		"<key-value-name>",
		&armappconfiguration.KeyValuesCreateOrUpdateOptions{KeyValueParameters: &armappconfiguration.KeyValue{
			Properties: &armappconfiguration.KeyValueProperties{
				Tags: map[string]*string{
					"tag1": to.StringPtr("tagValue1"),
					"tag2": to.StringPtr("tagValue2"),
				},
				Value: to.StringPtr("<value>"),
			},
		},
		})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("KeyValue.ID: %s\n", *res.ID)
}

// x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/preview/2021-03-01-preview/examples/ConfigurationStoresDeleteKeyValue.json
func ExampleKeyValuesClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armappconfiguration.NewKeyValuesClient("<subscription-id>", cred, nil)
	poller, err := client.BeginDelete(ctx,
		"<resource-group-name>",
		"<config-store-name>",
		"<key-value-name>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	_, err = poller.PollUntilDone(ctx, 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}
}