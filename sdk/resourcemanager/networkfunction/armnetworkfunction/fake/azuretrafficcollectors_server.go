//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/networkfunction/armnetworkfunction/v2"
	"net/http"
	"net/url"
	"regexp"
)

// AzureTrafficCollectorsServer is a fake server for instances of the armnetworkfunction.AzureTrafficCollectorsClient type.
type AzureTrafficCollectorsServer struct {
	// BeginCreateOrUpdate is the fake for method AzureTrafficCollectorsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, azureTrafficCollectorName string, parameters armnetworkfunction.AzureTrafficCollector, options *armnetworkfunction.AzureTrafficCollectorsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armnetworkfunction.AzureTrafficCollectorsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method AzureTrafficCollectorsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, azureTrafficCollectorName string, options *armnetworkfunction.AzureTrafficCollectorsClientBeginDeleteOptions) (resp azfake.PollerResponder[armnetworkfunction.AzureTrafficCollectorsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method AzureTrafficCollectorsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, azureTrafficCollectorName string, options *armnetworkfunction.AzureTrafficCollectorsClientGetOptions) (resp azfake.Responder[armnetworkfunction.AzureTrafficCollectorsClientGetResponse], errResp azfake.ErrorResponder)

	// UpdateTags is the fake for method AzureTrafficCollectorsClient.UpdateTags
	// HTTP status codes to indicate success: http.StatusOK
	UpdateTags func(ctx context.Context, resourceGroupName string, azureTrafficCollectorName string, parameters armnetworkfunction.TagsObject, options *armnetworkfunction.AzureTrafficCollectorsClientUpdateTagsOptions) (resp azfake.Responder[armnetworkfunction.AzureTrafficCollectorsClientUpdateTagsResponse], errResp azfake.ErrorResponder)
}

// NewAzureTrafficCollectorsServerTransport creates a new instance of AzureTrafficCollectorsServerTransport with the provided implementation.
// The returned AzureTrafficCollectorsServerTransport instance is connected to an instance of armnetworkfunction.AzureTrafficCollectorsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewAzureTrafficCollectorsServerTransport(srv *AzureTrafficCollectorsServer) *AzureTrafficCollectorsServerTransport {
	return &AzureTrafficCollectorsServerTransport{
		srv:                 srv,
		beginCreateOrUpdate: newTracker[azfake.PollerResponder[armnetworkfunction.AzureTrafficCollectorsClientCreateOrUpdateResponse]](),
		beginDelete:         newTracker[azfake.PollerResponder[armnetworkfunction.AzureTrafficCollectorsClientDeleteResponse]](),
	}
}

// AzureTrafficCollectorsServerTransport connects instances of armnetworkfunction.AzureTrafficCollectorsClient to instances of AzureTrafficCollectorsServer.
// Don't use this type directly, use NewAzureTrafficCollectorsServerTransport instead.
type AzureTrafficCollectorsServerTransport struct {
	srv                 *AzureTrafficCollectorsServer
	beginCreateOrUpdate *tracker[azfake.PollerResponder[armnetworkfunction.AzureTrafficCollectorsClientCreateOrUpdateResponse]]
	beginDelete         *tracker[azfake.PollerResponder[armnetworkfunction.AzureTrafficCollectorsClientDeleteResponse]]
}

// Do implements the policy.Transporter interface for AzureTrafficCollectorsServerTransport.
func (a *AzureTrafficCollectorsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "AzureTrafficCollectorsClient.BeginCreateOrUpdate":
		resp, err = a.dispatchBeginCreateOrUpdate(req)
	case "AzureTrafficCollectorsClient.BeginDelete":
		resp, err = a.dispatchBeginDelete(req)
	case "AzureTrafficCollectorsClient.Get":
		resp, err = a.dispatchGet(req)
	case "AzureTrafficCollectorsClient.UpdateTags":
		resp, err = a.dispatchUpdateTags(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (a *AzureTrafficCollectorsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if a.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := a.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkFunction/azureTrafficCollectors/(?P<azureTrafficCollectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armnetworkfunction.AzureTrafficCollector](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		azureTrafficCollectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureTrafficCollectorName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, azureTrafficCollectorNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		a.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated}, resp.StatusCode) {
		a.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		a.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (a *AzureTrafficCollectorsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if a.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := a.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkFunction/azureTrafficCollectors/(?P<azureTrafficCollectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		azureTrafficCollectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureTrafficCollectorName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := a.srv.BeginDelete(req.Context(), resourceGroupNameParam, azureTrafficCollectorNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		a.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		a.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		a.beginDelete.remove(req)
	}

	return resp, nil
}

func (a *AzureTrafficCollectorsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if a.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkFunction/azureTrafficCollectors/(?P<azureTrafficCollectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	azureTrafficCollectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureTrafficCollectorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.Get(req.Context(), resourceGroupNameParam, azureTrafficCollectorNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AzureTrafficCollector, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AzureTrafficCollectorsServerTransport) dispatchUpdateTags(req *http.Request) (*http.Response, error) {
	if a.srv.UpdateTags == nil {
		return nil, &nonRetriableError{errors.New("fake for method UpdateTags not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.NetworkFunction/azureTrafficCollectors/(?P<azureTrafficCollectorName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armnetworkfunction.TagsObject](req)
	if err != nil {
		return nil, err
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	azureTrafficCollectorNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("azureTrafficCollectorName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := a.srv.UpdateTags(req.Context(), resourceGroupNameParam, azureTrafficCollectorNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).AzureTrafficCollector, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}