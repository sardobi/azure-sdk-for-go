//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmobilenetwork

import (
	"context"
	"errors"
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

// SlicesClient contains the methods for the Slices group.
// Don't use this type directly, use NewSlicesClient() instead.
type SlicesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSlicesClient creates a new instance of SlicesClient with the specified values.
// subscriptionID - The ID of the target subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSlicesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*SlicesClient, error) {
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
	client := &SlicesClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a network slice.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// mobileNetworkName - The name of the mobile network.
// sliceName - The name of the network slice.
// parameters - Parameters supplied to the create or update network slice operation.
// options - SlicesClientBeginCreateOrUpdateOptions contains the optional parameters for the SlicesClient.BeginCreateOrUpdate
// method.
func (client *SlicesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, mobileNetworkName string, sliceName string, parameters Slice, options *SlicesClientBeginCreateOrUpdateOptions) (*runtime.Poller[SlicesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, mobileNetworkName, sliceName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[SlicesClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SlicesClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a network slice.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01-preview
func (client *SlicesClient) createOrUpdate(ctx context.Context, resourceGroupName string, mobileNetworkName string, sliceName string, parameters Slice, options *SlicesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, mobileNetworkName, sliceName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SlicesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, sliceName string, parameters Slice, options *SlicesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/slices/{sliceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if sliceName == "" {
		return nil, errors.New("parameter sliceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sliceName}", url.PathEscape(sliceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified network slice.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// mobileNetworkName - The name of the mobile network.
// sliceName - The name of the network slice.
// options - SlicesClientBeginDeleteOptions contains the optional parameters for the SlicesClient.BeginDelete method.
func (client *SlicesClient) BeginDelete(ctx context.Context, resourceGroupName string, mobileNetworkName string, sliceName string, options *SlicesClientBeginDeleteOptions) (*runtime.Poller[SlicesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, mobileNetworkName, sliceName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.pl, &runtime.NewPollerOptions[SlicesClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[SlicesClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified network slice.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01-preview
func (client *SlicesClient) deleteOperation(ctx context.Context, resourceGroupName string, mobileNetworkName string, sliceName string, options *SlicesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, mobileNetworkName, sliceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SlicesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, sliceName string, options *SlicesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/slices/{sliceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if sliceName == "" {
		return nil, errors.New("parameter sliceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sliceName}", url.PathEscape(sliceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets information about the specified network slice.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// mobileNetworkName - The name of the mobile network.
// sliceName - The name of the network slice.
// options - SlicesClientGetOptions contains the optional parameters for the SlicesClient.Get method.
func (client *SlicesClient) Get(ctx context.Context, resourceGroupName string, mobileNetworkName string, sliceName string, options *SlicesClientGetOptions) (SlicesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, mobileNetworkName, sliceName, options)
	if err != nil {
		return SlicesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SlicesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SlicesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SlicesClient) getCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, sliceName string, options *SlicesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/slices/{sliceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if sliceName == "" {
		return nil, errors.New("parameter sliceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sliceName}", url.PathEscape(sliceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SlicesClient) getHandleResponse(resp *http.Response) (SlicesClientGetResponse, error) {
	result := SlicesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Slice); err != nil {
		return SlicesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByMobileNetworkPager - Lists all slices in the mobile network.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// mobileNetworkName - The name of the mobile network.
// options - SlicesClientListByMobileNetworkOptions contains the optional parameters for the SlicesClient.ListByMobileNetwork
// method.
func (client *SlicesClient) NewListByMobileNetworkPager(resourceGroupName string, mobileNetworkName string, options *SlicesClientListByMobileNetworkOptions) *runtime.Pager[SlicesClientListByMobileNetworkResponse] {
	return runtime.NewPager(runtime.PagingHandler[SlicesClientListByMobileNetworkResponse]{
		More: func(page SlicesClientListByMobileNetworkResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *SlicesClientListByMobileNetworkResponse) (SlicesClientListByMobileNetworkResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByMobileNetworkCreateRequest(ctx, resourceGroupName, mobileNetworkName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return SlicesClientListByMobileNetworkResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return SlicesClientListByMobileNetworkResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return SlicesClientListByMobileNetworkResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByMobileNetworkHandleResponse(resp)
		},
	})
}

// listByMobileNetworkCreateRequest creates the ListByMobileNetwork request.
func (client *SlicesClient) listByMobileNetworkCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, options *SlicesClientListByMobileNetworkOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/slices"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByMobileNetworkHandleResponse handles the ListByMobileNetwork response.
func (client *SlicesClient) listByMobileNetworkHandleResponse(resp *http.Response) (SlicesClientListByMobileNetworkResponse, error) {
	result := SlicesClientListByMobileNetworkResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.SliceListResult); err != nil {
		return SlicesClientListByMobileNetworkResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates slice tags.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-04-01-preview
// resourceGroupName - The name of the resource group. The name is case insensitive.
// mobileNetworkName - The name of the mobile network.
// sliceName - The name of the network slice.
// parameters - Parameters supplied to update network slice tags.
// options - SlicesClientUpdateTagsOptions contains the optional parameters for the SlicesClient.UpdateTags method.
func (client *SlicesClient) UpdateTags(ctx context.Context, resourceGroupName string, mobileNetworkName string, sliceName string, parameters TagsObject, options *SlicesClientUpdateTagsOptions) (SlicesClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, mobileNetworkName, sliceName, parameters, options)
	if err != nil {
		return SlicesClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SlicesClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SlicesClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *SlicesClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, mobileNetworkName string, sliceName string, parameters TagsObject, options *SlicesClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MobileNetwork/mobileNetworks/{mobileNetworkName}/slices/{sliceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if mobileNetworkName == "" {
		return nil, errors.New("parameter mobileNetworkName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{mobileNetworkName}", url.PathEscape(mobileNetworkName))
	if sliceName == "" {
		return nil, errors.New("parameter sliceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{sliceName}", url.PathEscape(sliceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-04-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *SlicesClient) updateTagsHandleResponse(resp *http.Response) (SlicesClientUpdateTagsResponse, error) {
	result := SlicesClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Slice); err != nil {
		return SlicesClientUpdateTagsResponse{}, err
	}
	return result, nil
}