//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armstorsimple1200series_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/storsimple1200series/armstorsimple1200series"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/ChapSettingsListByDevice.json
func ExampleChapSettingsClient_NewListByDevicePager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple1200series.NewChapSettingsClient("9eb689cd-7243-43b4-b6f6-5c65cb296641", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListByDevicePager("HSDK-0NZI14MDTF",
		"ResourceGroupForSDKTest",
		"hAzureSDKOperations",
		nil)
	for pager.More() {
		nextResult, err := pager.NextPage(ctx)
		if err != nil {
			log.Fatalf("failed to advance page: %v", err)
		}
		for _, v := range nextResult.Value {
			// TODO: use page item
			_ = v
		}
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/ChapSettingsGet.json
func ExampleChapSettingsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple1200series.NewChapSettingsClient("9eb689cd-7243-43b4-b6f6-5c65cb296641", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"HSDK-WSJQERQW3F",
		"ChapSettingForSDK",
		"ResourceGroupForSDKTest",
		"hAzureSDKOperations",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/ChapSettingsCreateOrUpdate.json
func ExampleChapSettingsClient_BeginCreateOrUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple1200series.NewChapSettingsClient("9eb689cd-7243-43b4-b6f6-5c65cb296641", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreateOrUpdate(ctx,
		"HSDK-WSJQERQW3F",
		"ChapSettingForSDK",
		"ResourceGroupForSDKTest",
		"hAzureSDKOperations",
		armstorsimple1200series.ChapSettings{
			Name: to.Ptr("ChapSettingForSDK"),
			Properties: &armstorsimple1200series.ChapProperties{
				Password: &armstorsimple1200series.AsymmetricEncryptedSecret{
					EncryptionAlgorithm:             to.Ptr(armstorsimple1200series.EncryptionAlgorithmRSAESPKCS1V15),
					EncryptionCertificateThumbprint: to.Ptr("D73DB57C4CDD6761E159F8D1E8A7D759424983FD"),
					Value:                           to.Ptr("<value>"),
				},
			},
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	res, err := poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/storSimple1200Series/resource-manager/Microsoft.StorSimple/stable/2016-10-01/examples/ChapSettingsDelete.json
func ExampleChapSettingsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armstorsimple1200series.NewChapSettingsClient("9eb689cd-7243-43b4-b6f6-5c65cb296641", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"HSDK-WSJQERQW3F",
		"ChapSettingForSDKForDelete",
		"ResourceGroupForSDKTest",
		"hAzureSDKOperations",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}
