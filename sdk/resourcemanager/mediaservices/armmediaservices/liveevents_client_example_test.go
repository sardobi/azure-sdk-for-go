//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armmediaservices_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/mediaservices/armmediaservices/v3"
)

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/liveevent-list-all.json
func ExampleLiveEventsClient_NewListPager() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewLiveEventsClient("0a6ec948-5a62-437d-b9df-934dc7c1b722", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	pager := client.NewListPager("mediaresources",
		"slitestmedia10",
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/liveevent-list-by-name.json
func ExampleLiveEventsClient_Get() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewLiveEventsClient("0a6ec948-5a62-437d-b9df-934dc7c1b722", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	res, err := client.Get(ctx,
		"mediaresources",
		"slitestmedia10",
		"myLiveEvent1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	// TODO: use response item
	_ = res
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/liveevent-create.json
func ExampleLiveEventsClient_BeginCreate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewLiveEventsClient("0a6ec948-5a62-437d-b9df-934dc7c1b722", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginCreate(ctx,
		"mediaresources",
		"slitestmedia10",
		"myLiveEvent1",
		armmediaservices.LiveEvent{
			Location: to.Ptr("West US"),
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
			},
			Properties: &armmediaservices.LiveEventProperties{
				Description: to.Ptr("test event 1"),
				Input: &armmediaservices.LiveEventInput{
					AccessControl: &armmediaservices.LiveEventInputAccessControl{
						IP: &armmediaservices.IPAccessControl{
							Allow: []*armmediaservices.IPRange{
								{
									Name:               to.Ptr("AllowAll"),
									Address:            to.Ptr("0.0.0.0"),
									SubnetPrefixLength: to.Ptr[int32](0),
								}},
						},
					},
					KeyFrameIntervalDuration: to.Ptr("PT6S"),
					StreamingProtocol:        to.Ptr(armmediaservices.LiveEventInputProtocolRTMP),
				},
				Preview: &armmediaservices.LiveEventPreview{
					AccessControl: &armmediaservices.LiveEventPreviewAccessControl{
						IP: &armmediaservices.IPAccessControl{
							Allow: []*armmediaservices.IPRange{
								{
									Name:               to.Ptr("AllowAll"),
									Address:            to.Ptr("0.0.0.0"),
									SubnetPrefixLength: to.Ptr[int32](0),
								}},
						},
					},
				},
			},
		},
		&armmediaservices.LiveEventsClientBeginCreateOptions{AutoStart: nil})
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/liveevent-update.json
func ExampleLiveEventsClient_BeginUpdate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewLiveEventsClient("0a6ec948-5a62-437d-b9df-934dc7c1b722", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginUpdate(ctx,
		"mediaresources",
		"slitestmedia10",
		"myLiveEvent1",
		armmediaservices.LiveEvent{
			Location: to.Ptr("West US"),
			Tags: map[string]*string{
				"tag1": to.Ptr("value1"),
				"tag2": to.Ptr("value2"),
				"tag3": to.Ptr("value3"),
			},
			Properties: &armmediaservices.LiveEventProperties{
				Description: to.Ptr("test event updated"),
				Input: &armmediaservices.LiveEventInput{
					AccessControl: &armmediaservices.LiveEventInputAccessControl{
						IP: &armmediaservices.IPAccessControl{
							Allow: []*armmediaservices.IPRange{
								{
									Name:    to.Ptr("AllowOne"),
									Address: to.Ptr("192.1.1.0"),
								}},
						},
					},
					KeyFrameIntervalDuration: to.Ptr("PT6S"),
					StreamingProtocol:        to.Ptr(armmediaservices.LiveEventInputProtocolFragmentedMP4),
				},
				Preview: &armmediaservices.LiveEventPreview{
					AccessControl: &armmediaservices.LiveEventPreviewAccessControl{
						IP: &armmediaservices.IPAccessControl{
							Allow: []*armmediaservices.IPRange{
								{
									Name:    to.Ptr("AllowOne"),
									Address: to.Ptr("192.1.1.0"),
								}},
						},
					},
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

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/liveevent-delete.json
func ExampleLiveEventsClient_BeginDelete() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewLiveEventsClient("0a6ec948-5a62-437d-b9df-934dc7c1b722", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginDelete(ctx,
		"mediaresources",
		"slitestmedia10",
		"myLiveEvent1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/liveevent-allocate.json
func ExampleLiveEventsClient_BeginAllocate() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewLiveEventsClient("0a6ec948-5a62-437d-b9df-934dc7c1b722", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginAllocate(ctx,
		"mediaresources",
		"slitestmedia10",
		"myLiveEvent1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/liveevent-start.json
func ExampleLiveEventsClient_BeginStart() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewLiveEventsClient("0a6ec948-5a62-437d-b9df-934dc7c1b722", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginStart(ctx,
		"mediaresources",
		"slitestmedia10",
		"myLiveEvent1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/liveevent-stop.json
func ExampleLiveEventsClient_BeginStop() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewLiveEventsClient("0a6ec948-5a62-437d-b9df-934dc7c1b722", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginStop(ctx,
		"mediaresources",
		"slitestmedia10",
		"myLiveEvent1",
		armmediaservices.LiveEventActionInput{
			RemoveOutputsOnStop: to.Ptr(false),
		},
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}

// Generated from example definition: https://github.com/Azure/azure-rest-api-specs/tree/main/specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/liveevent-reset.json
func ExampleLiveEventsClient_BeginReset() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	ctx := context.Background()
	client, err := armmediaservices.NewLiveEventsClient("0a6ec948-5a62-437d-b9df-934dc7c1b722", cred, nil)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
	poller, err := client.BeginReset(ctx,
		"mediaresources",
		"slitestmedia10",
		"myLiveEvent1",
		nil)
	if err != nil {
		log.Fatalf("failed to finish the request: %v", err)
	}
	_, err = poller.PollUntilDone(ctx, nil)
	if err != nil {
		log.Fatalf("failed to pull the result: %v", err)
	}
}