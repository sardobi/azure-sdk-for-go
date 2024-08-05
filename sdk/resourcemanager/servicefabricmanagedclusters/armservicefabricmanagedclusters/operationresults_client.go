//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armservicefabricmanagedclusters

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// OperationResultsClient contains the methods for the OperationResults group.
// Don't use this type directly, use NewOperationResultsClient() instead.
type OperationResultsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewOperationResultsClient creates a new instance of OperationResultsClient with the specified values.
//   - subscriptionID - The customer subscription identifier.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewOperationResultsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*OperationResultsClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &OperationResultsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// Get - Get long running operation result.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2024-04-01
//   - location - The location for the cluster code versions. This is different from cluster location.
//   - operationID - operation identifier.
//   - options - OperationResultsClientGetOptions contains the optional parameters for the OperationResultsClient.Get method.
func (client *OperationResultsClient) Get(ctx context.Context, location string, operationID string, options *OperationResultsClientGetOptions) (OperationResultsClientGetResponse, error) {
	var err error
	const operationName = "OperationResultsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, location, operationID, options)
	if err != nil {
		return OperationResultsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return OperationResultsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return OperationResultsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *OperationResultsClient) getCreateRequest(ctx context.Context, location string, operationID string, options *OperationResultsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.ServiceFabric/locations/{location}/managedClusterOperationResults/{operationId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if location == "" {
		return nil, errors.New("parameter location cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	if operationID == "" {
		return nil, errors.New("parameter operationID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{operationId}", url.PathEscape(operationID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2024-04-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *OperationResultsClient) getHandleResponse(resp *http.Response) (OperationResultsClientGetResponse, error) {
	result := OperationResultsClientGetResponse{}
	if val := resp.Header.Get("Location"); val != "" {
		result.Location = &val
	}
	return result, nil
}