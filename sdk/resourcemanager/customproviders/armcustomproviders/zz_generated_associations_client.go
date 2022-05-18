//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcustomproviders

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

// AssociationsClient contains the methods for the Associations group.
// Don't use this type directly, use NewAssociationsClient() instead.
type AssociationsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewAssociationsClient creates a new instance of AssociationsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewAssociationsClient(credential azcore.TokenCredential, options *arm.ClientOptions) (*AssociationsClient, error) {
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
	client := &AssociationsClient{
		host: ep,
		pl:   pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Create or update an association.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-01-preview
// scope - The scope of the association. The scope can be any valid REST resource instance. For example, use
// '/subscriptions/{subscription-id}/resourceGroups/{resource-group-name}/providers/Microsoft.Compute/virtualMachines/{vm-name}'
// for a virtual machine resource.
// associationName - The name of the association.
// association - The parameters required to create or update an association.
// options - AssociationsClientBeginCreateOrUpdateOptions contains the optional parameters for the AssociationsClient.BeginCreateOrUpdate
// method.
func (client *AssociationsClient) BeginCreateOrUpdate(ctx context.Context, scope string, associationName string, association Association, options *AssociationsClientBeginCreateOrUpdateOptions) (*runtime.Poller[AssociationsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, scope, associationName, association, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[AssociationsClientCreateOrUpdateResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[AssociationsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Create or update an association.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-01-preview
func (client *AssociationsClient) createOrUpdate(ctx context.Context, scope string, associationName string, association Association, options *AssociationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, scope, associationName, association, options)
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
func (client *AssociationsClient) createOrUpdateCreateRequest(ctx context.Context, scope string, associationName string, association Association, options *AssociationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CustomProviders/associations/{associationName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, association)
}

// BeginDelete - Delete an association.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-01-preview
// scope - The scope of the association.
// associationName - The name of the association.
// options - AssociationsClientBeginDeleteOptions contains the optional parameters for the AssociationsClient.BeginDelete
// method.
func (client *AssociationsClient) BeginDelete(ctx context.Context, scope string, associationName string, options *AssociationsClientBeginDeleteOptions) (*runtime.Poller[AssociationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, scope, associationName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[AssociationsClientDeleteResponse](resp, client.pl, nil)
	} else {
		return runtime.NewPollerFromResumeToken[AssociationsClientDeleteResponse](options.ResumeToken, client.pl, nil)
	}
}

// Delete - Delete an association.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-01-preview
func (client *AssociationsClient) deleteOperation(ctx context.Context, scope string, associationName string, options *AssociationsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, scope, associationName, options)
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
func (client *AssociationsClient) deleteCreateRequest(ctx context.Context, scope string, associationName string, options *AssociationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CustomProviders/associations/{associationName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Get an association.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-01-preview
// scope - The scope of the association.
// associationName - The name of the association.
// options - AssociationsClientGetOptions contains the optional parameters for the AssociationsClient.Get method.
func (client *AssociationsClient) Get(ctx context.Context, scope string, associationName string, options *AssociationsClientGetOptions) (AssociationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, associationName, options)
	if err != nil {
		return AssociationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return AssociationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return AssociationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *AssociationsClient) getCreateRequest(ctx context.Context, scope string, associationName string, options *AssociationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CustomProviders/associations/{associationName}"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	if associationName == "" {
		return nil, errors.New("parameter associationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{associationName}", url.PathEscape(associationName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *AssociationsClient) getHandleResponse(resp *http.Response) (AssociationsClientGetResponse, error) {
	result := AssociationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Association); err != nil {
		return AssociationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListAllPager - Gets all association for the given scope.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2018-09-01-preview
// scope - The scope of the association.
// options - AssociationsClientListAllOptions contains the optional parameters for the AssociationsClient.ListAll method.
func (client *AssociationsClient) NewListAllPager(scope string, options *AssociationsClientListAllOptions) *runtime.Pager[AssociationsClientListAllResponse] {
	return runtime.NewPager(runtime.PagingHandler[AssociationsClientListAllResponse]{
		More: func(page AssociationsClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *AssociationsClientListAllResponse) (AssociationsClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, scope, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return AssociationsClientListAllResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return AssociationsClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return AssociationsClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *AssociationsClient) listAllCreateRequest(ctx context.Context, scope string, options *AssociationsClientListAllOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.CustomProviders/associations"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *AssociationsClient) listAllHandleResponse(resp *http.Response) (AssociationsClientListAllResponse, error) {
	result := AssociationsClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.AssociationsList); err != nil {
		return AssociationsClientListAllResponse{}, err
	}
	return result, nil
}
