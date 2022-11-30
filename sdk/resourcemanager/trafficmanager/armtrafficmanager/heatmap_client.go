//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armtrafficmanager

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// HeatMapClient contains the methods for the HeatMap group.
// Don't use this type directly, use NewHeatMapClient() instead.
type HeatMapClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewHeatMapClient creates a new instance of HeatMapClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewHeatMapClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*HeatMapClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &HeatMapClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// Get - Gets latest heatmap for Traffic Manager profile.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// profileName - The name of the Traffic Manager profile.
// options - HeatMapClientGetOptions contains the optional parameters for the HeatMapClient.Get method.
func (client *HeatMapClient) Get(ctx context.Context, resourceGroupName string, profileName string, options *HeatMapClientGetOptions) (HeatMapClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, profileName, options)
	if err != nil {
		return HeatMapClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HeatMapClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return HeatMapClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *HeatMapClient) getCreateRequest(ctx context.Context, resourceGroupName string, profileName string, options *HeatMapClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/trafficmanagerprofiles/{profileName}/heatMaps/{heatMapType}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if profileName == "" {
		return nil, errors.New("parameter profileName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{profileName}", url.PathEscape(profileName))
	urlPath = strings.ReplaceAll(urlPath, "{heatMapType}", url.PathEscape("default"))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.TopLeft != nil {
		reqQP.Set("topLeft", strings.Join(strings.Fields(strings.Trim(fmt.Sprint(options.TopLeft), "[]")), ","))
	}
	if options != nil && options.BotRight != nil {
		reqQP.Set("botRight", strings.Join(strings.Fields(strings.Trim(fmt.Sprint(options.BotRight), "[]")), ","))
	}
	reqQP.Set("api-version", "2022-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *HeatMapClient) getHandleResponse(resp *http.Response) (HeatMapClientGetResponse, error) {
	result := HeatMapClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.HeatMapModel); err != nil {
		return HeatMapClientGetResponse{}, err
	}
	return result, nil
}