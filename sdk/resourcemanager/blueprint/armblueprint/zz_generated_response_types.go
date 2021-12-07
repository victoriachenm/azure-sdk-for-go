//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armblueprint

import "net/http"

// ArtifactsCreateOrUpdateResponse contains the response from method Artifacts.CreateOrUpdate.
type ArtifactsCreateOrUpdateResponse struct {
	ArtifactsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArtifactsCreateOrUpdateResult contains the result from method Artifacts.CreateOrUpdate.
type ArtifactsCreateOrUpdateResult struct {
	ArtifactClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ArtifactsCreateOrUpdateResult.
func (a *ArtifactsCreateOrUpdateResult) UnmarshalJSON(data []byte) error {
	res, err := unmarshalArtifactClassification(data)
	if err != nil {
		return err
	}
	a.ArtifactClassification = res
	return nil
}

// ArtifactsDeleteResponse contains the response from method Artifacts.Delete.
type ArtifactsDeleteResponse struct {
	ArtifactsDeleteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArtifactsDeleteResult contains the result from method Artifacts.Delete.
type ArtifactsDeleteResult struct {
	ArtifactClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ArtifactsDeleteResult.
func (a *ArtifactsDeleteResult) UnmarshalJSON(data []byte) error {
	res, err := unmarshalArtifactClassification(data)
	if err != nil {
		return err
	}
	a.ArtifactClassification = res
	return nil
}

// ArtifactsGetResponse contains the response from method Artifacts.Get.
type ArtifactsGetResponse struct {
	ArtifactsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArtifactsGetResult contains the result from method Artifacts.Get.
type ArtifactsGetResult struct {
	ArtifactClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ArtifactsGetResult.
func (a *ArtifactsGetResult) UnmarshalJSON(data []byte) error {
	res, err := unmarshalArtifactClassification(data)
	if err != nil {
		return err
	}
	a.ArtifactClassification = res
	return nil
}

// ArtifactsListResponse contains the response from method Artifacts.List.
type ArtifactsListResponse struct {
	ArtifactsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// ArtifactsListResult contains the result from method Artifacts.List.
type ArtifactsListResult struct {
	ArtifactList
}

// AssignmentOperationsGetResponse contains the response from method AssignmentOperations.Get.
type AssignmentOperationsGetResponse struct {
	AssignmentOperationsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AssignmentOperationsGetResult contains the result from method AssignmentOperations.Get.
type AssignmentOperationsGetResult struct {
	AssignmentOperation
}

// AssignmentOperationsListResponse contains the response from method AssignmentOperations.List.
type AssignmentOperationsListResponse struct {
	AssignmentOperationsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AssignmentOperationsListResult contains the result from method AssignmentOperations.List.
type AssignmentOperationsListResult struct {
	AssignmentOperationList
}

// AssignmentsCreateOrUpdateResponse contains the response from method Assignments.CreateOrUpdate.
type AssignmentsCreateOrUpdateResponse struct {
	AssignmentsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AssignmentsCreateOrUpdateResult contains the result from method Assignments.CreateOrUpdate.
type AssignmentsCreateOrUpdateResult struct {
	Assignment
}

// AssignmentsDeleteResponse contains the response from method Assignments.Delete.
type AssignmentsDeleteResponse struct {
	AssignmentsDeleteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AssignmentsDeleteResult contains the result from method Assignments.Delete.
type AssignmentsDeleteResult struct {
	Assignment
}

// AssignmentsGetResponse contains the response from method Assignments.Get.
type AssignmentsGetResponse struct {
	AssignmentsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AssignmentsGetResult contains the result from method Assignments.Get.
type AssignmentsGetResult struct {
	Assignment
}

// AssignmentsListResponse contains the response from method Assignments.List.
type AssignmentsListResponse struct {
	AssignmentsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AssignmentsListResult contains the result from method Assignments.List.
type AssignmentsListResult struct {
	AssignmentList
}

// AssignmentsWhoIsBlueprintResponse contains the response from method Assignments.WhoIsBlueprint.
type AssignmentsWhoIsBlueprintResponse struct {
	AssignmentsWhoIsBlueprintResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// AssignmentsWhoIsBlueprintResult contains the result from method Assignments.WhoIsBlueprint.
type AssignmentsWhoIsBlueprintResult struct {
	WhoIsBlueprintContract
}

// BlueprintsCreateOrUpdateResponse contains the response from method Blueprints.CreateOrUpdate.
type BlueprintsCreateOrUpdateResponse struct {
	BlueprintsCreateOrUpdateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BlueprintsCreateOrUpdateResult contains the result from method Blueprints.CreateOrUpdate.
type BlueprintsCreateOrUpdateResult struct {
	Blueprint
}

// BlueprintsDeleteResponse contains the response from method Blueprints.Delete.
type BlueprintsDeleteResponse struct {
	BlueprintsDeleteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BlueprintsDeleteResult contains the result from method Blueprints.Delete.
type BlueprintsDeleteResult struct {
	Blueprint
}

// BlueprintsGetResponse contains the response from method Blueprints.Get.
type BlueprintsGetResponse struct {
	BlueprintsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BlueprintsGetResult contains the result from method Blueprints.Get.
type BlueprintsGetResult struct {
	Blueprint
}

// BlueprintsListResponse contains the response from method Blueprints.List.
type BlueprintsListResponse struct {
	BlueprintsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// BlueprintsListResult contains the result from method Blueprints.List.
type BlueprintsListResult struct {
	BlueprintList
}

// PublishedArtifactsGetResponse contains the response from method PublishedArtifacts.Get.
type PublishedArtifactsGetResponse struct {
	PublishedArtifactsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PublishedArtifactsGetResult contains the result from method PublishedArtifacts.Get.
type PublishedArtifactsGetResult struct {
	ArtifactClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type PublishedArtifactsGetResult.
func (p *PublishedArtifactsGetResult) UnmarshalJSON(data []byte) error {
	res, err := unmarshalArtifactClassification(data)
	if err != nil {
		return err
	}
	p.ArtifactClassification = res
	return nil
}

// PublishedArtifactsListResponse contains the response from method PublishedArtifacts.List.
type PublishedArtifactsListResponse struct {
	PublishedArtifactsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PublishedArtifactsListResult contains the result from method PublishedArtifacts.List.
type PublishedArtifactsListResult struct {
	ArtifactList
}

// PublishedBlueprintsCreateResponse contains the response from method PublishedBlueprints.Create.
type PublishedBlueprintsCreateResponse struct {
	PublishedBlueprintsCreateResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PublishedBlueprintsCreateResult contains the result from method PublishedBlueprints.Create.
type PublishedBlueprintsCreateResult struct {
	PublishedBlueprint
}

// PublishedBlueprintsDeleteResponse contains the response from method PublishedBlueprints.Delete.
type PublishedBlueprintsDeleteResponse struct {
	PublishedBlueprintsDeleteResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PublishedBlueprintsDeleteResult contains the result from method PublishedBlueprints.Delete.
type PublishedBlueprintsDeleteResult struct {
	PublishedBlueprint
}

// PublishedBlueprintsGetResponse contains the response from method PublishedBlueprints.Get.
type PublishedBlueprintsGetResponse struct {
	PublishedBlueprintsGetResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PublishedBlueprintsGetResult contains the result from method PublishedBlueprints.Get.
type PublishedBlueprintsGetResult struct {
	PublishedBlueprint
}

// PublishedBlueprintsListResponse contains the response from method PublishedBlueprints.List.
type PublishedBlueprintsListResponse struct {
	PublishedBlueprintsListResult
	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// PublishedBlueprintsListResult contains the result from method PublishedBlueprints.List.
type PublishedBlueprintsListResult struct {
	PublishedBlueprintList
}