//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armazurestackhci

import "encoding/json"

// ArcSettingsClientConsentAndInstallDefaultExtensionsResponse contains the response from method ArcSettingsClient.ConsentAndInstallDefaultExtensions.
type ArcSettingsClientConsentAndInstallDefaultExtensionsResponse struct {
	// ArcSetting details.
	ArcSetting
}

// ArcSettingsClientCreateIdentityResponse contains the response from method ArcSettingsClient.BeginCreateIdentity.
type ArcSettingsClientCreateIdentityResponse struct {
	// ArcIdentity details.
	ArcIdentityResponse
}

// ArcSettingsClientCreateResponse contains the response from method ArcSettingsClient.Create.
type ArcSettingsClientCreateResponse struct {
	// ArcSetting details.
	ArcSetting
}

// ArcSettingsClientDeleteResponse contains the response from method ArcSettingsClient.BeginDelete.
type ArcSettingsClientDeleteResponse struct {
	// placeholder for future response values
}

// ArcSettingsClientGeneratePasswordResponse contains the response from method ArcSettingsClient.GeneratePassword.
type ArcSettingsClientGeneratePasswordResponse struct {
	PasswordCredential
}

// ArcSettingsClientGetResponse contains the response from method ArcSettingsClient.Get.
type ArcSettingsClientGetResponse struct {
	// ArcSetting details.
	ArcSetting
}

// ArcSettingsClientInitializeDisableProcessResponse contains the response from method ArcSettingsClient.BeginInitializeDisableProcess.
type ArcSettingsClientInitializeDisableProcessResponse struct {
	// placeholder for future response values
}

// ArcSettingsClientListByClusterResponse contains the response from method ArcSettingsClient.NewListByClusterPager.
type ArcSettingsClientListByClusterResponse struct {
	// List of ArcSetting proxy resources for the HCI cluster.
	ArcSettingList
}

// ArcSettingsClientUpdateResponse contains the response from method ArcSettingsClient.Update.
type ArcSettingsClientUpdateResponse struct {
	// ArcSetting details.
	ArcSetting
}

// ClustersClientConfigureRemoteSupportResponse contains the response from method ClustersClient.BeginConfigureRemoteSupport.
type ClustersClientConfigureRemoteSupportResponse struct {
	// Cluster details.
	Cluster
}

// ClustersClientCreateIdentityResponse contains the response from method ClustersClient.BeginCreateIdentity.
type ClustersClientCreateIdentityResponse struct {
	// Cluster Identity details.
	ClusterIdentityResponse
}

// ClustersClientCreateResponse contains the response from method ClustersClient.Create.
type ClustersClientCreateResponse struct {
	// Cluster details.
	Cluster
}

// ClustersClientDeleteResponse contains the response from method ClustersClient.BeginDelete.
type ClustersClientDeleteResponse struct {
	// placeholder for future response values
}

// ClustersClientExtendSoftwareAssuranceBenefitResponse contains the response from method ClustersClient.BeginExtendSoftwareAssuranceBenefit.
type ClustersClientExtendSoftwareAssuranceBenefitResponse struct {
	// Cluster details.
	Cluster
}

// ClustersClientGetResponse contains the response from method ClustersClient.Get.
type ClustersClientGetResponse struct {
	// Cluster details.
	Cluster
}

// ClustersClientListByResourceGroupResponse contains the response from method ClustersClient.NewListByResourceGroupPager.
type ClustersClientListByResourceGroupResponse struct {
	// List of clusters.
	ClusterList
}

// ClustersClientListBySubscriptionResponse contains the response from method ClustersClient.NewListBySubscriptionPager.
type ClustersClientListBySubscriptionResponse struct {
	// List of clusters.
	ClusterList
}

// ClustersClientTriggerLogCollectionResponse contains the response from method ClustersClient.BeginTriggerLogCollection.
type ClustersClientTriggerLogCollectionResponse struct {
	// Cluster details.
	Cluster
}

// ClustersClientUpdateResponse contains the response from method ClustersClient.Update.
type ClustersClientUpdateResponse struct {
	// Cluster details.
	Cluster
}

// ClustersClientUploadCertificateResponse contains the response from method ClustersClient.BeginUploadCertificate.
type ClustersClientUploadCertificateResponse struct {
	// placeholder for future response values
}

// DeploymentSettingsClientCreateOrUpdateResponse contains the response from method DeploymentSettingsClient.BeginCreateOrUpdate.
type DeploymentSettingsClientCreateOrUpdateResponse struct {
	// Edge device resource
	DeploymentSetting
}

// DeploymentSettingsClientDeleteResponse contains the response from method DeploymentSettingsClient.BeginDelete.
type DeploymentSettingsClientDeleteResponse struct {
	// placeholder for future response values
}

// DeploymentSettingsClientGetResponse contains the response from method DeploymentSettingsClient.Get.
type DeploymentSettingsClientGetResponse struct {
	// Edge device resource
	DeploymentSetting
}

// DeploymentSettingsClientListByClustersResponse contains the response from method DeploymentSettingsClient.NewListByClustersPager.
type DeploymentSettingsClientListByClustersResponse struct {
	// The response of a DeploymentSetting list operation.
	DeploymentSettingListResult
}

// EdgeDevicesClientCreateOrUpdateResponse contains the response from method EdgeDevicesClient.BeginCreateOrUpdate.
type EdgeDevicesClientCreateOrUpdateResponse struct {
	// Edge device resource.
	EdgeDeviceClassification
}

// MarshalJSON implements the json.Marshaller interface for type EdgeDevicesClientCreateOrUpdateResponse.
func (e EdgeDevicesClientCreateOrUpdateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.EdgeDeviceClassification)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EdgeDevicesClientCreateOrUpdateResponse.
func (e *EdgeDevicesClientCreateOrUpdateResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalEdgeDeviceClassification(data)
	if err != nil {
		return err
	}
	e.EdgeDeviceClassification = res
	return nil
}

// EdgeDevicesClientDeleteResponse contains the response from method EdgeDevicesClient.BeginDelete.
type EdgeDevicesClientDeleteResponse struct {
	// placeholder for future response values
}

// EdgeDevicesClientGetResponse contains the response from method EdgeDevicesClient.Get.
type EdgeDevicesClientGetResponse struct {
	// Edge device resource.
	EdgeDeviceClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type EdgeDevicesClientGetResponse.
func (e *EdgeDevicesClientGetResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalEdgeDeviceClassification(data)
	if err != nil {
		return err
	}
	e.EdgeDeviceClassification = res
	return nil
}

// EdgeDevicesClientListResponse contains the response from method EdgeDevicesClient.NewListPager.
type EdgeDevicesClientListResponse struct {
	// The response of a EdgeDevice list operation.
	EdgeDeviceListResult
}

// EdgeDevicesClientValidateResponse contains the response from method EdgeDevicesClient.BeginValidate.
type EdgeDevicesClientValidateResponse struct {
	// An Accepted response with an Operation-Location header.
	ValidateResponse
}

// ExtensionsClientCreateResponse contains the response from method ExtensionsClient.BeginCreate.
type ExtensionsClientCreateResponse struct {
	// Details of a particular extension in HCI Cluster.
	Extension
}

// ExtensionsClientDeleteResponse contains the response from method ExtensionsClient.BeginDelete.
type ExtensionsClientDeleteResponse struct {
	// placeholder for future response values
}

// ExtensionsClientGetResponse contains the response from method ExtensionsClient.Get.
type ExtensionsClientGetResponse struct {
	// Details of a particular extension in HCI Cluster.
	Extension
}

// ExtensionsClientListByArcSettingResponse contains the response from method ExtensionsClient.NewListByArcSettingPager.
type ExtensionsClientListByArcSettingResponse struct {
	// List of Extensions in HCI cluster.
	ExtensionList
}

// ExtensionsClientUpdateResponse contains the response from method ExtensionsClient.BeginUpdate.
type ExtensionsClientUpdateResponse struct {
	// Details of a particular extension in HCI Cluster.
	Extension
}

// ExtensionsClientUpgradeResponse contains the response from method ExtensionsClient.BeginUpgrade.
type ExtensionsClientUpgradeResponse struct {
	// placeholder for future response values
}

// OffersClientGetResponse contains the response from method OffersClient.Get.
type OffersClientGetResponse struct {
	// Offer details.
	Offer
}

// OffersClientListByClusterResponse contains the response from method OffersClient.NewListByClusterPager.
type OffersClientListByClusterResponse struct {
	// List of Offer proxy resources for the HCI cluster.
	OfferList
}

// OffersClientListByPublisherResponse contains the response from method OffersClient.NewListByPublisherPager.
type OffersClientListByPublisherResponse struct {
	// List of Offer proxy resources for the HCI cluster.
	OfferList
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	// A list of REST API operations supported by an Azure Resource Provider. It contains an URL link to get the next set of results.
	OperationListResult
}

// PublishersClientGetResponse contains the response from method PublishersClient.Get.
type PublishersClientGetResponse struct {
	// Publisher details.
	Publisher
}

// PublishersClientListByClusterResponse contains the response from method PublishersClient.NewListByClusterPager.
type PublishersClientListByClusterResponse struct {
	// List of Publisher proxy resources for the HCI cluster.
	PublisherList
}

// SKUsClientGetResponse contains the response from method SKUsClient.Get.
type SKUsClientGetResponse struct {
	// Sku details.
	SKU
}

// SKUsClientListByOfferResponse contains the response from method SKUsClient.NewListByOfferPager.
type SKUsClientListByOfferResponse struct {
	// List of SKU proxy resources for the HCI cluster.
	SKUList
}

// SecuritySettingsClientCreateOrUpdateResponse contains the response from method SecuritySettingsClient.BeginCreateOrUpdate.
type SecuritySettingsClientCreateOrUpdateResponse struct {
	// Security settings proxy resource
	SecuritySetting
}

// SecuritySettingsClientDeleteResponse contains the response from method SecuritySettingsClient.BeginDelete.
type SecuritySettingsClientDeleteResponse struct {
	// placeholder for future response values
}

// SecuritySettingsClientGetResponse contains the response from method SecuritySettingsClient.Get.
type SecuritySettingsClientGetResponse struct {
	// Security settings proxy resource
	SecuritySetting
}

// SecuritySettingsClientListByClustersResponse contains the response from method SecuritySettingsClient.NewListByClustersPager.
type SecuritySettingsClientListByClustersResponse struct {
	// The response of a SecuritySetting list operation.
	SecuritySettingListResult
}

// UpdateRunsClientDeleteResponse contains the response from method UpdateRunsClient.BeginDelete.
type UpdateRunsClientDeleteResponse struct {
	// placeholder for future response values
}

// UpdateRunsClientGetResponse contains the response from method UpdateRunsClient.Get.
type UpdateRunsClientGetResponse struct {
	// Details of an Update run
	UpdateRun
}

// UpdateRunsClientListResponse contains the response from method UpdateRunsClient.NewListPager.
type UpdateRunsClientListResponse struct {
	// List of Update runs
	UpdateRunList
}

// UpdateRunsClientPutResponse contains the response from method UpdateRunsClient.Put.
type UpdateRunsClientPutResponse struct {
	// Details of an Update run
	UpdateRun
}

// UpdateSummariesClientDeleteResponse contains the response from method UpdateSummariesClient.BeginDelete.
type UpdateSummariesClientDeleteResponse struct {
	// placeholder for future response values
}

// UpdateSummariesClientGetResponse contains the response from method UpdateSummariesClient.Get.
type UpdateSummariesClientGetResponse struct {
	// Get the update summaries for the cluster
	UpdateSummaries
}

// UpdateSummariesClientListResponse contains the response from method UpdateSummariesClient.NewListPager.
type UpdateSummariesClientListResponse struct {
	// List of Update Summaries
	UpdateSummariesList
}

// UpdateSummariesClientPutResponse contains the response from method UpdateSummariesClient.Put.
type UpdateSummariesClientPutResponse struct {
	// Get the update summaries for the cluster
	UpdateSummaries
}

// UpdatesClientDeleteResponse contains the response from method UpdatesClient.BeginDelete.
type UpdatesClientDeleteResponse struct {
	// placeholder for future response values
}

// UpdatesClientGetResponse contains the response from method UpdatesClient.Get.
type UpdatesClientGetResponse struct {
	// Update details
	Update
}

// UpdatesClientListResponse contains the response from method UpdatesClient.NewListPager.
type UpdatesClientListResponse struct {
	// List of Updates
	UpdateList
}

// UpdatesClientPostResponse contains the response from method UpdatesClient.BeginPost.
type UpdatesClientPostResponse struct {
	// placeholder for future response values
}

// UpdatesClientPutResponse contains the response from method UpdatesClient.Put.
type UpdatesClientPutResponse struct {
	// Update details
	Update
}