// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package dataflow

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"math"
	"net/http"
	"net/url"
	"time"

	dataflowpb "cloud.google.com/go/dataflow/apiv1beta3/dataflowpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

var newTemplatesClientHook clientHook

// TemplatesCallOptions contains the retry settings for each method of TemplatesClient.
type TemplatesCallOptions struct {
	CreateJobFromTemplate []gax.CallOption
	LaunchTemplate        []gax.CallOption
	GetTemplate           []gax.CallOption
}

func defaultTemplatesGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("dataflow.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("dataflow.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://dataflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultTemplatesCallOptions() *TemplatesCallOptions {
	return &TemplatesCallOptions{
		CreateJobFromTemplate: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		LaunchTemplate: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetTemplate: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

func defaultTemplatesRESTCallOptions() *TemplatesCallOptions {
	return &TemplatesCallOptions{
		CreateJobFromTemplate: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		LaunchTemplate: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
		GetTemplate: []gax.CallOption{
			gax.WithTimeout(60000 * time.Millisecond),
		},
	}
}

// internalTemplatesClient is an interface that defines the methods available from Dataflow API.
type internalTemplatesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateJobFromTemplate(context.Context, *dataflowpb.CreateJobFromTemplateRequest, ...gax.CallOption) (*dataflowpb.Job, error)
	LaunchTemplate(context.Context, *dataflowpb.LaunchTemplateRequest, ...gax.CallOption) (*dataflowpb.LaunchTemplateResponse, error)
	GetTemplate(context.Context, *dataflowpb.GetTemplateRequest, ...gax.CallOption) (*dataflowpb.GetTemplateResponse, error)
}

// TemplatesClient is a client for interacting with Dataflow API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Provides a method to create Cloud Dataflow jobs from templates.
type TemplatesClient struct {
	// The internal transport-dependent client.
	internalClient internalTemplatesClient

	// The call options for this service.
	CallOptions *TemplatesCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *TemplatesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *TemplatesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *TemplatesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateJobFromTemplate creates a Cloud Dataflow job from a template.
func (c *TemplatesClient) CreateJobFromTemplate(ctx context.Context, req *dataflowpb.CreateJobFromTemplateRequest, opts ...gax.CallOption) (*dataflowpb.Job, error) {
	return c.internalClient.CreateJobFromTemplate(ctx, req, opts...)
}

// LaunchTemplate launch a template.
func (c *TemplatesClient) LaunchTemplate(ctx context.Context, req *dataflowpb.LaunchTemplateRequest, opts ...gax.CallOption) (*dataflowpb.LaunchTemplateResponse, error) {
	return c.internalClient.LaunchTemplate(ctx, req, opts...)
}

// GetTemplate get the template associated with a template.
func (c *TemplatesClient) GetTemplate(ctx context.Context, req *dataflowpb.GetTemplateRequest, opts ...gax.CallOption) (*dataflowpb.GetTemplateResponse, error) {
	return c.internalClient.GetTemplate(ctx, req, opts...)
}

// templatesGRPCClient is a client for interacting with Dataflow API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type templatesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing TemplatesClient
	CallOptions **TemplatesCallOptions

	// The gRPC API client.
	templatesClient dataflowpb.TemplatesServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogHeaders []string
}

// NewTemplatesClient creates a new templates service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Provides a method to create Cloud Dataflow jobs from templates.
func NewTemplatesClient(ctx context.Context, opts ...option.ClientOption) (*TemplatesClient, error) {
	clientOpts := defaultTemplatesGRPCClientOptions()
	if newTemplatesClientHook != nil {
		hookOpts, err := newTemplatesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := TemplatesClient{CallOptions: defaultTemplatesCallOptions()}

	c := &templatesGRPCClient{
		connPool:        connPool,
		templatesClient: dataflowpb.NewTemplatesServiceClient(connPool),
		CallOptions:     &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *templatesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *templatesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *templatesGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type templatesRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* headers to be sent with each request.
	xGoogHeaders []string

	// Points back to the CallOptions field of the containing TemplatesClient
	CallOptions **TemplatesCallOptions
}

// NewTemplatesRESTClient creates a new templates service rest client.
//
// Provides a method to create Cloud Dataflow jobs from templates.
func NewTemplatesRESTClient(ctx context.Context, opts ...option.ClientOption) (*TemplatesClient, error) {
	clientOpts := append(defaultTemplatesRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultTemplatesRESTCallOptions()
	c := &templatesRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &TemplatesClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultTemplatesRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://dataflow.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://dataflow.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://dataflow.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *templatesRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", gax.GoVersion}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogHeaders = []string{"x-goog-api-client", gax.XGoogHeader(kv...)}
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *templatesRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *templatesRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *templatesGRPCClient) CreateJobFromTemplate(ctx context.Context, req *dataflowpb.CreateJobFromTemplateRequest, opts ...gax.CallOption) (*dataflowpb.Job, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).CreateJobFromTemplate[0:len((*c.CallOptions).CreateJobFromTemplate):len((*c.CallOptions).CreateJobFromTemplate)], opts...)
	var resp *dataflowpb.Job
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.templatesClient.CreateJobFromTemplate(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *templatesGRPCClient) LaunchTemplate(ctx context.Context, req *dataflowpb.LaunchTemplateRequest, opts ...gax.CallOption) (*dataflowpb.LaunchTemplateResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).LaunchTemplate[0:len((*c.CallOptions).LaunchTemplate):len((*c.CallOptions).LaunchTemplate)], opts...)
	var resp *dataflowpb.LaunchTemplateResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.templatesClient.LaunchTemplate(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *templatesGRPCClient) GetTemplate(ctx context.Context, req *dataflowpb.GetTemplateRequest, opts ...gax.CallOption) (*dataflowpb.GetTemplateResponse, error) {
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()))}

	hds = append(c.xGoogHeaders, hds...)
	ctx = gax.InsertMetadataIntoOutgoingContext(ctx, hds...)
	opts = append((*c.CallOptions).GetTemplate[0:len((*c.CallOptions).GetTemplate):len((*c.CallOptions).GetTemplate)], opts...)
	var resp *dataflowpb.GetTemplateResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.templatesClient.GetTemplate(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateJobFromTemplate creates a Cloud Dataflow job from a template.
func (c *templatesRESTClient) CreateJobFromTemplate(ctx context.Context, req *dataflowpb.CreateJobFromTemplateRequest, opts ...gax.CallOption) (*dataflowpb.Job, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1b3/projects/%v/locations/%v/templates", req.GetProjectId(), req.GetLocation())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).CreateJobFromTemplate[0:len((*c.CallOptions).CreateJobFromTemplate):len((*c.CallOptions).CreateJobFromTemplate)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &dataflowpb.Job{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// LaunchTemplate launch a template.
func (c *templatesRESTClient) LaunchTemplate(ctx context.Context, req *dataflowpb.LaunchTemplateRequest, opts ...gax.CallOption) (*dataflowpb.LaunchTemplateResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	body := req.GetLaunchParameters()
	jsonReq, err := m.Marshal(body)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1b3/projects/%v/locations/%v/templates:launch", req.GetProjectId(), req.GetLocation())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetDynamicTemplate().GetGcsPath() != "" {
		params.Add("dynamicTemplate.gcsPath", fmt.Sprintf("%v", req.GetDynamicTemplate().GetGcsPath()))
	}
	if req.GetDynamicTemplate().GetStagingLocation() != "" {
		params.Add("dynamicTemplate.stagingLocation", fmt.Sprintf("%v", req.GetDynamicTemplate().GetStagingLocation()))
	}
	if req.GetGcsPath() != "" {
		params.Add("gcsPath", fmt.Sprintf("%v", req.GetGcsPath()))
	}
	if req.GetValidateOnly() {
		params.Add("validateOnly", fmt.Sprintf("%v", req.GetValidateOnly()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).LaunchTemplate[0:len((*c.CallOptions).LaunchTemplate):len((*c.CallOptions).LaunchTemplate)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &dataflowpb.LaunchTemplateResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("POST", baseUrl.String(), bytes.NewReader(jsonReq))
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// GetTemplate get the template associated with a template.
func (c *templatesRESTClient) GetTemplate(ctx context.Context, req *dataflowpb.GetTemplateRequest, opts ...gax.CallOption) (*dataflowpb.GetTemplateResponse, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1b3/projects/%v/locations/%v/templates:get", req.GetProjectId(), req.GetLocation())

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")
	if req.GetGcsPath() != "" {
		params.Add("gcsPath", fmt.Sprintf("%v", req.GetGcsPath()))
	}
	if req.GetView() != 0 {
		params.Add("view", fmt.Sprintf("%v", req.GetView()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	hds := []string{"x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "project_id", url.QueryEscape(req.GetProjectId()), "location", url.QueryEscape(req.GetLocation()))}

	hds = append(c.xGoogHeaders, hds...)
	hds = append(hds, "Content-Type", "application/json")
	headers := gax.BuildHeaders(ctx, hds...)
	opts = append((*c.CallOptions).GetTemplate[0:len((*c.CallOptions).GetTemplate):len((*c.CallOptions).GetTemplate)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &dataflowpb.GetTemplateResponse{}
	e := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		if settings.Path != "" {
			baseUrl.Path = settings.Path
		}
		httpReq, err := http.NewRequest("GET", baseUrl.String(), nil)
		if err != nil {
			return err
		}
		httpReq = httpReq.WithContext(ctx)
		httpReq.Header = headers

		httpRsp, err := c.httpClient.Do(httpReq)
		if err != nil {
			return err
		}
		defer httpRsp.Body.Close()

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		buf, err := io.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return err
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
