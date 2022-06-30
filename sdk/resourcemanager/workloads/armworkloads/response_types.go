//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armworkloads

// ClientSAPAvailabilityZoneDetailsResponse contains the response from method Client.SAPAvailabilityZoneDetails.
type ClientSAPAvailabilityZoneDetailsResponse struct {
	SAPAvailabilityZoneDetailsResult
}

// ClientSAPDiskConfigurationsResponse contains the response from method Client.SAPDiskConfigurations.
type ClientSAPDiskConfigurationsResponse struct {
	SAPDiskConfigurationsResult
}

// ClientSAPSizingRecommendationsResponse contains the response from method Client.SAPSizingRecommendations.
type ClientSAPSizingRecommendationsResponse struct {
	SAPSizingRecommendationResultClassification
}

// UnmarshalJSON implements the json.Unmarshaller interface for type ClientSAPSizingRecommendationsResponse.
func (c *ClientSAPSizingRecommendationsResponse) UnmarshalJSON(data []byte) error {
	res, err := unmarshalSAPSizingRecommendationResultClassification(data)
	if err != nil {
		return err
	}
	c.SAPSizingRecommendationResultClassification = res
	return nil
}

// ClientSAPSupportedSKUResponse contains the response from method Client.SAPSupportedSKU.
type ClientSAPSupportedSKUResponse struct {
	SAPSupportedResourceSKUsResult
}

// MonitorsClientCreateResponse contains the response from method MonitorsClient.Create.
type MonitorsClientCreateResponse struct {
	Monitor
}

// MonitorsClientDeleteResponse contains the response from method MonitorsClient.Delete.
type MonitorsClientDeleteResponse struct {
	OperationStatusResult
}

// MonitorsClientGetResponse contains the response from method MonitorsClient.Get.
type MonitorsClientGetResponse struct {
	Monitor
}

// MonitorsClientListByResourceGroupResponse contains the response from method MonitorsClient.ListByResourceGroup.
type MonitorsClientListByResourceGroupResponse struct {
	MonitorListResult
}

// MonitorsClientListResponse contains the response from method MonitorsClient.List.
type MonitorsClientListResponse struct {
	MonitorListResult
}

// MonitorsClientUpdateResponse contains the response from method MonitorsClient.Update.
type MonitorsClientUpdateResponse struct {
	Monitor
}

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// PhpWorkloadsClientCreateOrUpdateResponse contains the response from method PhpWorkloadsClient.CreateOrUpdate.
type PhpWorkloadsClientCreateOrUpdateResponse struct {
	PhpWorkloadResource
}

// PhpWorkloadsClientDeleteResponse contains the response from method PhpWorkloadsClient.Delete.
type PhpWorkloadsClientDeleteResponse struct {
	// placeholder for future response values
}

// PhpWorkloadsClientGetResponse contains the response from method PhpWorkloadsClient.Get.
type PhpWorkloadsClientGetResponse struct {
	PhpWorkloadResource
}

// PhpWorkloadsClientListByResourceGroupResponse contains the response from method PhpWorkloadsClient.ListByResourceGroup.
type PhpWorkloadsClientListByResourceGroupResponse struct {
	PhpWorkloadResourceList
}

// PhpWorkloadsClientListBySubscriptionResponse contains the response from method PhpWorkloadsClient.ListBySubscription.
type PhpWorkloadsClientListBySubscriptionResponse struct {
	PhpWorkloadResourceList
}

// PhpWorkloadsClientUpdateResponse contains the response from method PhpWorkloadsClient.Update.
type PhpWorkloadsClientUpdateResponse struct {
	PhpWorkloadResource
}

// ProviderInstancesClientCreateResponse contains the response from method ProviderInstancesClient.Create.
type ProviderInstancesClientCreateResponse struct {
	ProviderInstance
}

// ProviderInstancesClientDeleteResponse contains the response from method ProviderInstancesClient.Delete.
type ProviderInstancesClientDeleteResponse struct {
	OperationStatusResult
}

// ProviderInstancesClientGetResponse contains the response from method ProviderInstancesClient.Get.
type ProviderInstancesClientGetResponse struct {
	ProviderInstance
}

// ProviderInstancesClientListResponse contains the response from method ProviderInstancesClient.List.
type ProviderInstancesClientListResponse struct {
	ProviderInstanceListResult
}

// SAPApplicationServerInstancesClientCreateResponse contains the response from method SAPApplicationServerInstancesClient.Create.
type SAPApplicationServerInstancesClientCreateResponse struct {
	SAPApplicationServerInstance
}

// SAPApplicationServerInstancesClientDeleteResponse contains the response from method SAPApplicationServerInstancesClient.Delete.
type SAPApplicationServerInstancesClientDeleteResponse struct {
	OperationStatusResult
}

// SAPApplicationServerInstancesClientGetResponse contains the response from method SAPApplicationServerInstancesClient.Get.
type SAPApplicationServerInstancesClientGetResponse struct {
	SAPApplicationServerInstance
}

// SAPApplicationServerInstancesClientListResponse contains the response from method SAPApplicationServerInstancesClient.List.
type SAPApplicationServerInstancesClientListResponse struct {
	SAPApplicationServerInstanceList
}

// SAPApplicationServerInstancesClientUpdateResponse contains the response from method SAPApplicationServerInstancesClient.Update.
type SAPApplicationServerInstancesClientUpdateResponse struct {
	SAPApplicationServerInstance
}

// SAPCentralInstancesClientCreateResponse contains the response from method SAPCentralInstancesClient.Create.
type SAPCentralInstancesClientCreateResponse struct {
	SAPCentralServerInstance
}

// SAPCentralInstancesClientDeleteResponse contains the response from method SAPCentralInstancesClient.Delete.
type SAPCentralInstancesClientDeleteResponse struct {
	OperationStatusResult
}

// SAPCentralInstancesClientGetResponse contains the response from method SAPCentralInstancesClient.Get.
type SAPCentralInstancesClientGetResponse struct {
	SAPCentralServerInstance
}

// SAPCentralInstancesClientListResponse contains the response from method SAPCentralInstancesClient.List.
type SAPCentralInstancesClientListResponse struct {
	SAPCentralInstanceList
}

// SAPCentralInstancesClientUpdateResponse contains the response from method SAPCentralInstancesClient.Update.
type SAPCentralInstancesClientUpdateResponse struct {
	SAPCentralServerInstance
}

// SAPDatabaseInstancesClientCreateResponse contains the response from method SAPDatabaseInstancesClient.Create.
type SAPDatabaseInstancesClientCreateResponse struct {
	SAPDatabaseInstance
}

// SAPDatabaseInstancesClientDeleteResponse contains the response from method SAPDatabaseInstancesClient.Delete.
type SAPDatabaseInstancesClientDeleteResponse struct {
	OperationStatusResult
}

// SAPDatabaseInstancesClientGetResponse contains the response from method SAPDatabaseInstancesClient.Get.
type SAPDatabaseInstancesClientGetResponse struct {
	SAPDatabaseInstance
}

// SAPDatabaseInstancesClientListResponse contains the response from method SAPDatabaseInstancesClient.List.
type SAPDatabaseInstancesClientListResponse struct {
	SAPDatabaseInstanceList
}

// SAPDatabaseInstancesClientUpdateResponse contains the response from method SAPDatabaseInstancesClient.Update.
type SAPDatabaseInstancesClientUpdateResponse struct {
	SAPDatabaseInstance
}

// SAPVirtualInstancesClientCreateResponse contains the response from method SAPVirtualInstancesClient.Create.
type SAPVirtualInstancesClientCreateResponse struct {
	SAPVirtualInstance
}

// SAPVirtualInstancesClientDeleteResponse contains the response from method SAPVirtualInstancesClient.Delete.
type SAPVirtualInstancesClientDeleteResponse struct {
	OperationStatusResult
}

// SAPVirtualInstancesClientGetResponse contains the response from method SAPVirtualInstancesClient.Get.
type SAPVirtualInstancesClientGetResponse struct {
	SAPVirtualInstance
}

// SAPVirtualInstancesClientListByResourceGroupResponse contains the response from method SAPVirtualInstancesClient.ListByResourceGroup.
type SAPVirtualInstancesClientListByResourceGroupResponse struct {
	SAPVirtualInstanceList
}

// SAPVirtualInstancesClientListBySubscriptionResponse contains the response from method SAPVirtualInstancesClient.ListBySubscription.
type SAPVirtualInstancesClientListBySubscriptionResponse struct {
	SAPVirtualInstanceList
}

// SAPVirtualInstancesClientStartResponse contains the response from method SAPVirtualInstancesClient.Start.
type SAPVirtualInstancesClientStartResponse struct {
	OperationStatusResult
}

// SAPVirtualInstancesClientStopResponse contains the response from method SAPVirtualInstancesClient.Stop.
type SAPVirtualInstancesClientStopResponse struct {
	OperationStatusResult
}

// SAPVirtualInstancesClientUpdateResponse contains the response from method SAPVirtualInstancesClient.Update.
type SAPVirtualInstancesClientUpdateResponse struct {
	SAPVirtualInstance
}

// SKUsClientListResponse contains the response from method SKUsClient.List.
type SKUsClientListResponse struct {
	SKUsListResult
}

// WordpressInstancesClientCreateOrUpdateResponse contains the response from method WordpressInstancesClient.CreateOrUpdate.
type WordpressInstancesClientCreateOrUpdateResponse struct {
	WordpressInstanceResource
}

// WordpressInstancesClientDeleteResponse contains the response from method WordpressInstancesClient.Delete.
type WordpressInstancesClientDeleteResponse struct {
	// placeholder for future response values
}

// WordpressInstancesClientGetResponse contains the response from method WordpressInstancesClient.Get.
type WordpressInstancesClientGetResponse struct {
	WordpressInstanceResource
}

// WordpressInstancesClientListResponse contains the response from method WordpressInstancesClient.List.
type WordpressInstancesClientListResponse struct {
	WordpressInstanceResourceList
}