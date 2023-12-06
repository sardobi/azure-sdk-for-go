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
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/support/armsupport/v2"
	"net/http"
	"net/url"
	"regexp"
)

// FilesNoSubscriptionServer is a fake server for instances of the armsupport.FilesNoSubscriptionClient type.
type FilesNoSubscriptionServer struct {
	// Create is the fake for method FilesNoSubscriptionClient.Create
	// HTTP status codes to indicate success: http.StatusCreated
	Create func(ctx context.Context, fileWorkspaceName string, fileName string, createFileParameters armsupport.FileDetails, options *armsupport.FilesNoSubscriptionClientCreateOptions) (resp azfake.Responder[armsupport.FilesNoSubscriptionClientCreateResponse], errResp azfake.ErrorResponder)

	// Get is the fake for method FilesNoSubscriptionClient.Get
	// HTTP status codes to indicate success: http.StatusOK
	Get func(ctx context.Context, fileWorkspaceName string, fileName string, options *armsupport.FilesNoSubscriptionClientGetOptions) (resp azfake.Responder[armsupport.FilesNoSubscriptionClientGetResponse], errResp azfake.ErrorResponder)

	// NewListPager is the fake for method FilesNoSubscriptionClient.NewListPager
	// HTTP status codes to indicate success: http.StatusOK
	NewListPager func(fileWorkspaceName string, options *armsupport.FilesNoSubscriptionClientListOptions) (resp azfake.PagerResponder[armsupport.FilesNoSubscriptionClientListResponse])

	// Upload is the fake for method FilesNoSubscriptionClient.Upload
	// HTTP status codes to indicate success: http.StatusNoContent
	Upload func(ctx context.Context, fileWorkspaceName string, fileName string, uploadFile armsupport.UploadFile, options *armsupport.FilesNoSubscriptionClientUploadOptions) (resp azfake.Responder[armsupport.FilesNoSubscriptionClientUploadResponse], errResp azfake.ErrorResponder)
}

// NewFilesNoSubscriptionServerTransport creates a new instance of FilesNoSubscriptionServerTransport with the provided implementation.
// The returned FilesNoSubscriptionServerTransport instance is connected to an instance of armsupport.FilesNoSubscriptionClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewFilesNoSubscriptionServerTransport(srv *FilesNoSubscriptionServer) *FilesNoSubscriptionServerTransport {
	return &FilesNoSubscriptionServerTransport{
		srv:          srv,
		newListPager: newTracker[azfake.PagerResponder[armsupport.FilesNoSubscriptionClientListResponse]](),
	}
}

// FilesNoSubscriptionServerTransport connects instances of armsupport.FilesNoSubscriptionClient to instances of FilesNoSubscriptionServer.
// Don't use this type directly, use NewFilesNoSubscriptionServerTransport instead.
type FilesNoSubscriptionServerTransport struct {
	srv          *FilesNoSubscriptionServer
	newListPager *tracker[azfake.PagerResponder[armsupport.FilesNoSubscriptionClientListResponse]]
}

// Do implements the policy.Transporter interface for FilesNoSubscriptionServerTransport.
func (f *FilesNoSubscriptionServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "FilesNoSubscriptionClient.Create":
		resp, err = f.dispatchCreate(req)
	case "FilesNoSubscriptionClient.Get":
		resp, err = f.dispatchGet(req)
	case "FilesNoSubscriptionClient.NewListPager":
		resp, err = f.dispatchNewListPager(req)
	case "FilesNoSubscriptionClient.Upload":
		resp, err = f.dispatchUpload(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (f *FilesNoSubscriptionServerTransport) dispatchCreate(req *http.Request) (*http.Response, error) {
	if f.srv.Create == nil {
		return nil, &nonRetriableError{errors.New("fake for method Create not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Support/fileWorkspaces/(?P<fileWorkspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/files/(?P<fileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsupport.FileDetails](req)
	if err != nil {
		return nil, err
	}
	fileWorkspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fileWorkspaceName")])
	if err != nil {
		return nil, err
	}
	fileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fileName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Create(req.Context(), fileWorkspaceNameParam, fileNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusCreated}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusCreated", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FileDetails, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FilesNoSubscriptionServerTransport) dispatchGet(req *http.Request) (*http.Response, error) {
	if f.srv.Get == nil {
		return nil, &nonRetriableError{errors.New("fake for method Get not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Support/fileWorkspaces/(?P<fileWorkspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/files/(?P<fileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	fileWorkspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fileWorkspaceName")])
	if err != nil {
		return nil, err
	}
	fileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fileName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Get(req.Context(), fileWorkspaceNameParam, fileNameParam, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusOK}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", respContent.HTTPStatus)}
	}
	resp, err := server.MarshalResponseAsJSON(respContent, server.GetResponse(respr).FileDetails, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f *FilesNoSubscriptionServerTransport) dispatchNewListPager(req *http.Request) (*http.Response, error) {
	if f.srv.NewListPager == nil {
		return nil, &nonRetriableError{errors.New("fake for method NewListPager not implemented")}
	}
	newListPager := f.newListPager.get(req)
	if newListPager == nil {
		const regexStr = `/providers/Microsoft\.Support/fileWorkspaces/(?P<fileWorkspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/files`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 1 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		fileWorkspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fileWorkspaceName")])
		if err != nil {
			return nil, err
		}
		resp := f.srv.NewListPager(fileWorkspaceNameParam, nil)
		newListPager = &resp
		f.newListPager.add(req, newListPager)
		server.PagerResponderInjectNextLinks(newListPager, req, func(page *armsupport.FilesNoSubscriptionClientListResponse, createLink func() string) {
			page.NextLink = to.Ptr(createLink())
		})
	}
	resp, err := server.PagerResponderNext(newListPager, req)
	if err != nil {
		return nil, err
	}
	if !contains([]int{http.StatusOK}, resp.StatusCode) {
		f.newListPager.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK", resp.StatusCode)}
	}
	if !server.PagerResponderMore(newListPager) {
		f.newListPager.remove(req)
	}
	return resp, nil
}

func (f *FilesNoSubscriptionServerTransport) dispatchUpload(req *http.Request) (*http.Response, error) {
	if f.srv.Upload == nil {
		return nil, &nonRetriableError{errors.New("fake for method Upload not implemented")}
	}
	const regexStr = `/providers/Microsoft\.Support/fileWorkspaces/(?P<fileWorkspaceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/files/(?P<fileName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/upload`
	regex := regexp.MustCompile(regexStr)
	matches := regex.FindStringSubmatch(req.URL.EscapedPath())
	if matches == nil || len(matches) < 2 {
		return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
	}
	body, err := server.UnmarshalRequestAsJSON[armsupport.UploadFile](req)
	if err != nil {
		return nil, err
	}
	fileWorkspaceNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fileWorkspaceName")])
	if err != nil {
		return nil, err
	}
	fileNameParam, err := url.PathUnescape(matches[regex.SubexpIndex("fileName")])
	if err != nil {
		return nil, err
	}
	respr, errRespr := f.srv.Upload(req.Context(), fileWorkspaceNameParam, fileNameParam, body, nil)
	if respErr := server.GetError(errRespr, req); respErr != nil {
		return nil, respErr
	}
	respContent := server.GetResponseContent(respr)
	if !contains([]int{http.StatusNoContent}, respContent.HTTPStatus) {
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusNoContent", respContent.HTTPStatus)}
	}
	resp, err := server.NewResponse(respContent, req, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}