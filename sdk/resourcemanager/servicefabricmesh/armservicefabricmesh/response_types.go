//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armservicefabricmesh

// ApplicationClientCreateResponse contains the response from method ApplicationClient.Create.
type ApplicationClientCreateResponse struct {
	ApplicationResourceDescription
}

// ApplicationClientDeleteResponse contains the response from method ApplicationClient.Delete.
type ApplicationClientDeleteResponse struct {
	// placeholder for future response values
}

// ApplicationClientGetResponse contains the response from method ApplicationClient.Get.
type ApplicationClientGetResponse struct {
	ApplicationResourceDescription
}

// ApplicationClientListByResourceGroupResponse contains the response from method ApplicationClient.NewListByResourceGroupPager.
type ApplicationClientListByResourceGroupResponse struct {
	ApplicationResourceDescriptionList
}

// ApplicationClientListBySubscriptionResponse contains the response from method ApplicationClient.NewListBySubscriptionPager.
type ApplicationClientListBySubscriptionResponse struct {
	ApplicationResourceDescriptionList
}

// CodePackageClientGetContainerLogsResponse contains the response from method CodePackageClient.GetContainerLogs.
type CodePackageClientGetContainerLogsResponse struct {
	ContainerLogs
}

// GatewayClientCreateResponse contains the response from method GatewayClient.Create.
type GatewayClientCreateResponse struct {
	GatewayResourceDescription
}

// GatewayClientDeleteResponse contains the response from method GatewayClient.Delete.
type GatewayClientDeleteResponse struct {
	// placeholder for future response values
}

// GatewayClientGetResponse contains the response from method GatewayClient.Get.
type GatewayClientGetResponse struct {
	GatewayResourceDescription
}

// GatewayClientListByResourceGroupResponse contains the response from method GatewayClient.NewListByResourceGroupPager.
type GatewayClientListByResourceGroupResponse struct {
	GatewayResourceDescriptionList
}

// GatewayClientListBySubscriptionResponse contains the response from method GatewayClient.NewListBySubscriptionPager.
type GatewayClientListBySubscriptionResponse struct {
	GatewayResourceDescriptionList
}

// NetworkClientCreateResponse contains the response from method NetworkClient.Create.
type NetworkClientCreateResponse struct {
	NetworkResourceDescription
}

// NetworkClientDeleteResponse contains the response from method NetworkClient.Delete.
type NetworkClientDeleteResponse struct {
	// placeholder for future response values
}

// NetworkClientGetResponse contains the response from method NetworkClient.Get.
type NetworkClientGetResponse struct {
	NetworkResourceDescription
}

// NetworkClientListByResourceGroupResponse contains the response from method NetworkClient.NewListByResourceGroupPager.
type NetworkClientListByResourceGroupResponse struct {
	NetworkResourceDescriptionList
}

// NetworkClientListBySubscriptionResponse contains the response from method NetworkClient.NewListBySubscriptionPager.
type NetworkClientListBySubscriptionResponse struct {
	NetworkResourceDescriptionList
}

// OperationsClientListResponse contains the response from method OperationsClient.NewListPager.
type OperationsClientListResponse struct {
	OperationListResult
}

// SecretClientCreateResponse contains the response from method SecretClient.Create.
type SecretClientCreateResponse struct {
	SecretResourceDescription
}

// SecretClientDeleteResponse contains the response from method SecretClient.Delete.
type SecretClientDeleteResponse struct {
	// placeholder for future response values
}

// SecretClientGetResponse contains the response from method SecretClient.Get.
type SecretClientGetResponse struct {
	SecretResourceDescription
}

// SecretClientListByResourceGroupResponse contains the response from method SecretClient.NewListByResourceGroupPager.
type SecretClientListByResourceGroupResponse struct {
	SecretResourceDescriptionList
}

// SecretClientListBySubscriptionResponse contains the response from method SecretClient.NewListBySubscriptionPager.
type SecretClientListBySubscriptionResponse struct {
	SecretResourceDescriptionList
}

// SecretValueClientCreateResponse contains the response from method SecretValueClient.Create.
type SecretValueClientCreateResponse struct {
	SecretValueResourceDescription
}

// SecretValueClientDeleteResponse contains the response from method SecretValueClient.Delete.
type SecretValueClientDeleteResponse struct {
	// placeholder for future response values
}

// SecretValueClientGetResponse contains the response from method SecretValueClient.Get.
type SecretValueClientGetResponse struct {
	SecretValueResourceDescription
}

// SecretValueClientListResponse contains the response from method SecretValueClient.NewListPager.
type SecretValueClientListResponse struct {
	SecretValueResourceDescriptionList
}

// SecretValueClientListValueResponse contains the response from method SecretValueClient.ListValue.
type SecretValueClientListValueResponse struct {
	SecretValue
}

// ServiceClientGetResponse contains the response from method ServiceClient.Get.
type ServiceClientGetResponse struct {
	ServiceResourceDescription
}

// ServiceClientListResponse contains the response from method ServiceClient.NewListPager.
type ServiceClientListResponse struct {
	ServiceResourceDescriptionList
}

// ServiceReplicaClientGetResponse contains the response from method ServiceReplicaClient.Get.
type ServiceReplicaClientGetResponse struct {
	ServiceReplicaDescription
}

// ServiceReplicaClientListResponse contains the response from method ServiceReplicaClient.NewListPager.
type ServiceReplicaClientListResponse struct {
	ServiceReplicaDescriptionList
}

// VolumeClientCreateResponse contains the response from method VolumeClient.Create.
type VolumeClientCreateResponse struct {
	VolumeResourceDescription
}

// VolumeClientDeleteResponse contains the response from method VolumeClient.Delete.
type VolumeClientDeleteResponse struct {
	// placeholder for future response values
}

// VolumeClientGetResponse contains the response from method VolumeClient.Get.
type VolumeClientGetResponse struct {
	VolumeResourceDescription
}

// VolumeClientListByResourceGroupResponse contains the response from method VolumeClient.NewListByResourceGroupPager.
type VolumeClientListByResourceGroupResponse struct {
	VolumeResourceDescriptionList
}

// VolumeClientListBySubscriptionResponse contains the response from method VolumeClient.NewListBySubscriptionPager.
type VolumeClientListBySubscriptionResponse struct {
	VolumeResourceDescriptionList
}