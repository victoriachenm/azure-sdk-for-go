//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armapplicationinsights_test

import (
	"context"
	"log"

	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/applicationinsights/armapplicationinsights"
)

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnnotationsList.json
func ExampleAnnotationsClient_List() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewAnnotationsClient("<subscription-id>", cred, nil)
	res, err := client.List(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<start>",
		"<end>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AnnotationsClientListResult)
}

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnnotationsCreate.json
func ExampleAnnotationsClient_Create() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewAnnotationsClient("<subscription-id>", cred, nil)
	res, err := client.Create(ctx,
		"<resource-group-name>",
		"<resource-name>",
		armapplicationinsights.Annotation{
			AnnotationName: to.StringPtr("<annotation-name>"),
			Category:       to.StringPtr("<category>"),
			EventTime:      to.TimePtr(func() time.Time { t, _ := time.Parse(time.RFC3339Nano, "2018-01-31T13:41:38.657Z"); return t }()),
			ID:             to.StringPtr("<id>"),
			Properties:     to.StringPtr("<properties>"),
		},
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AnnotationsClientCreateResult)
}

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnnotationsDelete.json
func ExampleAnnotationsClient_Delete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewAnnotationsClient("<subscription-id>", cred, nil)
	_, err = client.Delete(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<annotation-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
}

// x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/AnnotationsGet.json
func ExampleAnnotationsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client := armapplicationinsights.NewAnnotationsClient("<subscription-id>", cred, nil)
	res, err := client.Get(ctx,
		"<resource-group-name>",
		"<resource-name>",
		"<annotation-id>",
		nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Response result: %#v\n", res.AnnotationsClientGetResult)
}