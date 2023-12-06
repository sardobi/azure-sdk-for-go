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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/labservices/armlabservices"
	"net/http"
	"net/url"
	"regexp"
)

// LabsServer is a fake server for instances of the armlabservices.LabsClient type.
type LabsServer struct {
	// BeginCreateOrUpdate is the fake for method LabsClient.BeginCreateOrUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusCreated, http.StatusAccepted
	BeginCreateOrUpdate func(ctx context.Context, resourceGroupName string, labName string, body armlabservices.Lab, options *armlabservices.LabsClientBeginCreateOrUpdateOptions) (resp azfake.PollerResponder[armlabservices.LabsClientCreateOrUpdateResponse], errResp azfake.ErrorResponder)

	// BeginDelete is the fake for method LabsClient.BeginDelete
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted, http.StatusNoContent
	BeginDelete func(ctx context.Context, resourceGroupName string, labName string, options *armlabservices.LabsClientBeginDeleteOptions) (resp azfake.PollerResponder[armlabservices.LabsClientDeleteResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method LabsClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, resourceGroupName string, labName string, options *armlabservices.LabsClientGetOptions) (resp azfake.Responder[armlabservices.LabsClientGetResponse], errResp azfake.ErrorResponder)

	// NewListByResourceGroupPager is the fake for method LabsClient.NewListByResourceGroupPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListByResourceGroupPager func(resourceGroupName string, options *armlabservices.LabsClientListByResourceGroupOptions) (resp azfake.PagerResponder[armlabservices.LabsClientListByResourceGroupResponse])

	// NewListBySubscriptionPager is the fake for method LabsClient.NewListBySubscriptionPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListBySubscriptionPager func(options *armlabservices.LabsClientListBySubscriptionOptions) (resp azfake.PagerResponder[armlabservices.LabsClientListBySubscriptionResponse])

	// BeginPublish is the fake for method LabsClient.BeginPublish
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginPublish func(ctx context.Context, resourceGroupName string, labName string, options *armlabservices.LabsClientBeginPublishOptions) (resp azfake.PollerResponder[armlabservices.LabsClientPublishResponse], errResp azfake.ErrorResponder)

	// BeginSyncGroup is the fake for method LabsClient.BeginSyncGroup
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginSyncGroup func(ctx context.Context, resourceGroupName string, labName string, options *armlabservices.LabsClientBeginSyncGroupOptions) (resp azfake.PollerResponder[armlabservices.LabsClientSyncGroupResponse], errResp azfake.ErrorResponder)

	// BeginUpdate is the fake for method LabsClient.BeginUpdate
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginUpdate func(ctx context.Context, resourceGroupName string, labName string, body armlabservices.LabUpdate, options *armlabservices.LabsClientBeginUpdateOptions) (resp azfake.PollerResponder[armlabservices.LabsClientUpdateResponse], errResp azfake.ErrorResponder)
}

// NewLabsServerTransport creates a new instance of LabsServerTransport with the provided implementation.
// The returned LabsServerTransport instance is connected to an instance of armlabservices.LabsClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewLabsServerTransport(srv *LabsServer) *LabsServerTransport {
	return &LabsServerTransport{
		srv:                         srv,
		beginCreateOrUpdate:         newTracker[azfake.PollerResponder[armlabservices.LabsClientCreateOrUpdateResponse]](),
		beginDelete:                 newTracker[azfake.PollerResponder[armlabservices.LabsClientDeleteResponse]](),
		newListByResourceGroupPager: newTracker[azfake.PagerResponder[armlabservices.LabsClientListByResourceGroupResponse]](),
		newListBySubscriptionPager:  newTracker[azfake.PagerResponder[armlabservices.LabsClientListBySubscriptionResponse]](),
		beginPublish:                newTracker[azfake.PollerResponder[armlabservices.LabsClientPublishResponse]](),
		beginSyncGroup:              newTracker[azfake.PollerResponder[armlabservices.LabsClientSyncGroupResponse]](),
		beginUpdate:                 newTracker[azfake.PollerResponder[armlabservices.LabsClientUpdateResponse]](),
	}
}

// LabsServerTransport connects instances of armlabservices.LabsClient to instances of LabsServer.
// Don't use this type directly, use NewLabsServerTransport instead.
type LabsServerTransport struct {
	srv                         *LabsServer
	beginCreateOrUpdate         *tracker[azfake.PollerResponder[armlabservices.LabsClientCreateOrUpdateResponse]]
	beginDelete                 *tracker[azfake.PollerResponder[armlabservices.LabsClientDeleteResponse]]
	newListByResourceGroupPager *tracker[azfake.PagerResponder[armlabservices.LabsClientListByResourceGroupResponse]]
	newListBySubscriptionPager  *tracker[azfake.PagerResponder[armlabservices.LabsClientListBySubscriptionResponse]]
	beginPublish                *tracker[azfake.PollerResponder[armlabservices.LabsClientPublishResponse]]
	beginSyncGroup              *tracker[azfake.PollerResponder[armlabservices.LabsClientSyncGroupResponse]]
	beginUpdate                 *tracker[azfake.PollerResponder[armlabservices.LabsClientUpdateResponse]]
}

// Do implements the policy.Transporter interface for LabsServerTransport.
func (l *LabsServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "LabsClient.BeginCreateOrUpdate":
		resp, err = l.dispatchBeginCreateOrUpdate(req)
	case "LabsClient.BeginDelete":
		resp, err = l.dispatchBeginDelete(req)
	case "LabsClient.Get":
		resp, err = l.dispatchGet(req)
	case "LabsClient.NewListByResourceGroupPager":
		resp, err = l.dispatchNewListByResourceGroupPager(req)
	case "LabsClient.NewListBySubscriptionPager":
		resp, err = l.dispatchNewListBySubscriptionPager(req)
	case "LabsClient.BeginPublish":
		resp, err = l.dispatchBeginPublish(req)
	case "LabsClient.BeginSyncGroup":
		resp, err = l.dispatchBeginSyncGroup(req)
	case "LabsClient.BeginUpdate":
		resp, err = l.dispatchBeginUpdate(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (l *LabsServerTransport) dispatchBeginCreateOrUpdate(req *http.Request) (*http.Response, error) {
	if l.srv.BeginCreateOrUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginCreateOrUpdate not implemented")}
	}
	beginCreateOrUpdate := l.beginCreateOrUpdate.get(req)
	if beginCreateOrUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LabServices/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armlabservices.Lab](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginCreateOrUpdate(req.Context(), resourceGroupNameParam, labNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginCreateOrUpdate = &respr
		l.beginCreateOrUpdate.add(req, beginCreateOrUpdate)
	}

	resp, err := server.PollerResponderNext(beginCreateOrUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusCreated, http.StatusAccepted}, resp.StatusCode) {
		l.beginCreateOrUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusCreated, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginCreateOrUpdate) {
		l.beginCreateOrUpdate.remove(req)
	}

	return resp, nil
}

func (l *LabsServerTransport) dispatchBeginDelete(req *http.Request) (*http.Response, error) {
	if l.srv.BeginDelete == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginDelete not implemented")}
	}
	beginDelete := l.beginDelete.get(req)
	if beginDelete == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LabServices/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginDelete(req.Context(), resourceGroupNameParam, labNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginDelete = &respr
		l.beginDelete.add(req, beginDelete)
	}

	resp, err := server.PollerResponderNext(beginDelete, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted, http.StatusNoContent}, resp.StatusCode) {
		l.beginDelete.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted, http.StatusNoContent", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginDelete) {
		l.beginDelete.remove(req)
	}

	return resp, nil
}

func (l *LabsServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if l.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LabServices/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 3 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
	if err != nil {
		return nil, err
	}
	labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := l.srv.Get(req.Context(), resourceGroupNameParam, labNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).Lab, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (l *LabsServerTransport) dispatchNewListByResourceGroupPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListByResourceGroupPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListByResourceGroupPager not implemented")}
	}
	newListByResourceGroupPager := l.newListByResourceGroupPager.get(req)
	if newListByResourceGroupPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LabServices/labs`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 2 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		resp := l.srv.NewListByResourceGroupPager(resourceGroupNameParam, nil)
		newListByResourceGroupPager = &resp
		l.newListByResourceGroupPager.add(req, newListByResourceGroupPager)
		server.PagerResponderInjectNextLinks(newListByResourceGroupPager, req, func(page *armlabservices.LabsClientListByResourceGroupResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListByResourceGroupPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListByResourceGroupPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListByResourceGroupPager) {
		l.newListByResourceGroupPager.remove(req)
	}
	return resp, nil
}

func (l *LabsServerTransport) dispatchNewListBySubscriptionPager(req *http.Request) (*http.Response, error) {
	if l.srv.NewListBySubscriptionPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListBySubscriptionPager not implemented")}
	}
	newListBySubscriptionPager := l.newListBySubscriptionPager.get(req)
	if newListBySubscriptionPager == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LabServices/labs`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		qp := req.URL.Query()
		filterUnescaped, err := url.QueryUnescape(qp.Get("$filter"))
		if err != nil {
			return nil, err
		}
		filterParam := getOptional(filterUnescaped)
		var options *armlabservices.LabsClientListBySubscriptionOptions
		if filterParam != nil {
			options = &armlabservices.LabsClientListBySubscriptionOptions{
				Filter: filterParam,
			}
		}
		resp := l.srv.NewListBySubscriptionPager(options)
		newListBySubscriptionPager = &resp
		l.newListBySubscriptionPager.add(req, newListBySubscriptionPager)
		server.PagerResponderInjectNextLinks(newListBySubscriptionPager, req, func(page *armlabservices.LabsClientListBySubscriptionResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListBySubscriptionPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		l.newListBySubscriptionPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListBySubscriptionPager) {
		l.newListBySubscriptionPager.remove(req)
	}
	return resp, nil
}

func (l *LabsServerTransport) dispatchBeginPublish(req *http.Request) (*http.Response, error) {
	if l.srv.BeginPublish == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginPublish not implemented")}
	}
	beginPublish := l.beginPublish.get(req)
	if beginPublish == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LabServices/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/publish`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginPublish(req.Context(), resourceGroupNameParam, labNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginPublish = &respr
		l.beginPublish.add(req, beginPublish)
	}

	resp, err := server.PollerResponderNext(beginPublish, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		l.beginPublish.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginPublish) {
		l.beginPublish.remove(req)
	}

	return resp, nil
}

func (l *LabsServerTransport) dispatchBeginSyncGroup(req *http.Request) (*http.Response, error) {
	if l.srv.BeginSyncGroup == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginSyncGroup not implemented")}
	}
	beginSyncGroup := l.beginSyncGroup.get(req)
	if beginSyncGroup == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LabServices/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/syncGroup`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginSyncGroup(req.Context(), resourceGroupNameParam, labNameParam, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginSyncGroup = &respr
		l.beginSyncGroup.add(req, beginSyncGroup)
	}

	resp, err := server.PollerResponderNext(beginSyncGroup, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		l.beginSyncGroup.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginSyncGroup) {
		l.beginSyncGroup.remove(req)
	}

	return resp, nil
}

func (l *LabsServerTransport) dispatchBeginUpdate(req *http.Request) (*http.Response, error) {
	if l.srv.BeginUpdate == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginUpdate not implemented")}
	}
	beginUpdate := l.beginUpdate.get(req)
	if beginUpdate == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.LabServices/labs/(?P<labName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armlabservices.LabUpdate](req)
		if err != nil {
			return nil, err
		}
		resourceGroupNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		labNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("labName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := l.srv.BeginUpdate(req.Context(), resourceGroupNameParam, labNameParam, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginUpdate = &respr
		l.beginUpdate.add(req, beginUpdate)
	}

	resp, err := server.PollerResponderNext(beginUpdate, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		l.beginUpdate.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginUpdate) {
		l.beginUpdate.remove(req)
	}

	return resp, nil
}