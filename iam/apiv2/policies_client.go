// Copyright 2022 Google LLC
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

package iam

import (
	"context"
	"fmt"
	"math"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	iampb "google.golang.org/genproto/googleapis/iam/v2"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
)

var newPoliciesClientHook clientHook

// PoliciesCallOptions contains the retry settings for each method of PoliciesClient.
type PoliciesCallOptions struct {
	ListPolicies           []gax.CallOption
	GetPolicy              []gax.CallOption
	CreatePolicy           []gax.CallOption
	UpdatePolicy           []gax.CallOption
	DeletePolicy           []gax.CallOption
	ListApplicablePolicies []gax.CallOption
	GetOperation           []gax.CallOption
}

func defaultPoliciesGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("iam.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("iam.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://iam.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultPoliciesCallOptions() *PoliciesCallOptions {
	return &PoliciesCallOptions{
		ListPolicies: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		GetPolicy: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		CreatePolicy: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		UpdatePolicy: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		DeletePolicy: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        10000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ListApplicablePolicies: []gax.CallOption{},
		GetOperation:           []gax.CallOption{},
	}
}

// internalPoliciesClient is an interface that defines the methods available from Identity and Access Management (IAM) API.
type internalPoliciesClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ListPolicies(context.Context, *iampb.ListPoliciesRequest, ...gax.CallOption) *PolicyIterator
	GetPolicy(context.Context, *iampb.GetPolicyRequest, ...gax.CallOption) (*iampb.Policy, error)
	CreatePolicy(context.Context, *iampb.CreatePolicyRequest, ...gax.CallOption) (*CreatePolicyOperation, error)
	CreatePolicyOperation(name string) *CreatePolicyOperation
	UpdatePolicy(context.Context, *iampb.UpdatePolicyRequest, ...gax.CallOption) (*UpdatePolicyOperation, error)
	UpdatePolicyOperation(name string) *UpdatePolicyOperation
	DeletePolicy(context.Context, *iampb.DeletePolicyRequest, ...gax.CallOption) (*DeletePolicyOperation, error)
	DeletePolicyOperation(name string) *DeletePolicyOperation
	ListApplicablePolicies(context.Context, *iampb.ListApplicablePoliciesRequest, ...gax.CallOption) *PolicyIterator
	GetOperation(context.Context, *longrunningpb.GetOperationRequest, ...gax.CallOption) (*longrunningpb.Operation, error)
}

// PoliciesClient is a client for interacting with Identity and Access Management (IAM) API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// An interface for managing Identity and Access Management (IAM) policies.
type PoliciesClient struct {
	// The internal transport-dependent client.
	internalClient internalPoliciesClient

	// The call options for this service.
	CallOptions *PoliciesCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *PoliciesClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *PoliciesClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *PoliciesClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ListPolicies retrieves the policies of the specified kind that are attached to a
// resource.
//
// The response lists only policy metadata. In particular, policy rules are
// omitted.
func (c *PoliciesClient) ListPolicies(ctx context.Context, req *iampb.ListPoliciesRequest, opts ...gax.CallOption) *PolicyIterator {
	return c.internalClient.ListPolicies(ctx, req, opts...)
}

// GetPolicy gets a policy.
func (c *PoliciesClient) GetPolicy(ctx context.Context, req *iampb.GetPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	return c.internalClient.GetPolicy(ctx, req, opts...)
}

// CreatePolicy creates a policy.
func (c *PoliciesClient) CreatePolicy(ctx context.Context, req *iampb.CreatePolicyRequest, opts ...gax.CallOption) (*CreatePolicyOperation, error) {
	return c.internalClient.CreatePolicy(ctx, req, opts...)
}

// CreatePolicyOperation returns a new CreatePolicyOperation from a given name.
// The name must be that of a previously created CreatePolicyOperation, possibly from a different process.
func (c *PoliciesClient) CreatePolicyOperation(name string) *CreatePolicyOperation {
	return c.internalClient.CreatePolicyOperation(name)
}

// UpdatePolicy updates the specified policy.
//
// You can update only the rules and the display name for the policy.
//
// To update a policy, you should use a read-modify-write loop:
//
// Use GetPolicy to read the current
// version of the policy.
//
// Modify the policy as needed.
//
// Use UpdatePolicy to write the updated policy.
//
// This pattern helps prevent conflicts between concurrent updates.
func (c *PoliciesClient) UpdatePolicy(ctx context.Context, req *iampb.UpdatePolicyRequest, opts ...gax.CallOption) (*UpdatePolicyOperation, error) {
	return c.internalClient.UpdatePolicy(ctx, req, opts...)
}

// UpdatePolicyOperation returns a new UpdatePolicyOperation from a given name.
// The name must be that of a previously created UpdatePolicyOperation, possibly from a different process.
func (c *PoliciesClient) UpdatePolicyOperation(name string) *UpdatePolicyOperation {
	return c.internalClient.UpdatePolicyOperation(name)
}

// DeletePolicy deletes a policy. This action is permanent.
func (c *PoliciesClient) DeletePolicy(ctx context.Context, req *iampb.DeletePolicyRequest, opts ...gax.CallOption) (*DeletePolicyOperation, error) {
	return c.internalClient.DeletePolicy(ctx, req, opts...)
}

// DeletePolicyOperation returns a new DeletePolicyOperation from a given name.
// The name must be that of a previously created DeletePolicyOperation, possibly from a different process.
func (c *PoliciesClient) DeletePolicyOperation(name string) *DeletePolicyOperation {
	return c.internalClient.DeletePolicyOperation(name)
}

// ListApplicablePolicies retrieves all the policies that are attached to the specified resource,
// or anywhere in the ancestry of the resource. For example, for a project
// this endpoint would return all the denyPolicy kind policies attached to
// the project, its parent folder (if any), and its parent organization (if
// any).
// The endpoint requires the same permissions that it would take to call
// ListPolicies or GetPolicy.
//
// The main reason to use this endpoint is as a policy admin to debug access
// issues for a resource.
func (c *PoliciesClient) ListApplicablePolicies(ctx context.Context, req *iampb.ListApplicablePoliciesRequest, opts ...gax.CallOption) *PolicyIterator {
	return c.internalClient.ListApplicablePolicies(ctx, req, opts...)
}

// GetOperation is a utility method from google.longrunning.Operations.
func (c *PoliciesClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	return c.internalClient.GetOperation(ctx, req, opts...)
}

// policiesGRPCClient is a client for interacting with Identity and Access Management (IAM) API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type policiesGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing PoliciesClient
	CallOptions **PoliciesCallOptions

	// The gRPC API client.
	policiesClient iampb.PoliciesClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	operationsClient longrunningpb.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewPoliciesClient creates a new policies client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// An interface for managing Identity and Access Management (IAM) policies.
func NewPoliciesClient(ctx context.Context, opts ...option.ClientOption) (*PoliciesClient, error) {
	clientOpts := defaultPoliciesGRPCClientOptions()
	if newPoliciesClientHook != nil {
		hookOpts, err := newPoliciesClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	disableDeadlines, err := checkDisableDeadlines()
	if err != nil {
		return nil, err
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := PoliciesClient{CallOptions: defaultPoliciesCallOptions()}

	c := &policiesGRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		policiesClient:   iampb.NewPoliciesClient(connPool),
		CallOptions:      &client.CallOptions,
		operationsClient: longrunningpb.NewOperationsClient(connPool),
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	client.LROClient, err = lroauto.NewOperationsClient(ctx, gtransport.WithConnPool(connPool))
	if err != nil {
		// This error "should not happen", since we are just reusing old connection pool
		// and never actually need to dial.
		// If this does happen, we could leak connp. However, we cannot close conn:
		// If the user invoked the constructor with option.WithGRPCConn,
		// we would close a connection that's still in use.
		// TODO: investigate error conditions.
		return nil, err
	}
	c.LROClient = &client.LROClient
	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated.
func (c *policiesGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *policiesGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *policiesGRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *policiesGRPCClient) ListPolicies(ctx context.Context, req *iampb.ListPoliciesRequest, opts ...gax.CallOption) *PolicyIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListPolicies[0:len((*c.CallOptions).ListPolicies):len((*c.CallOptions).ListPolicies)], opts...)
	it := &PolicyIterator{}
	req = proto.Clone(req).(*iampb.ListPoliciesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*iampb.Policy, string, error) {
		resp := &iampb.ListPoliciesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.policiesClient.ListPolicies(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetPolicies(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *policiesGRPCClient) GetPolicy(ctx context.Context, req *iampb.GetPolicyRequest, opts ...gax.CallOption) (*iampb.Policy, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetPolicy[0:len((*c.CallOptions).GetPolicy):len((*c.CallOptions).GetPolicy)], opts...)
	var resp *iampb.Policy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.policiesClient.GetPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *policiesGRPCClient) CreatePolicy(ctx context.Context, req *iampb.CreatePolicyRequest, opts ...gax.CallOption) (*CreatePolicyOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreatePolicy[0:len((*c.CallOptions).CreatePolicy):len((*c.CallOptions).CreatePolicy)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.policiesClient.CreatePolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &CreatePolicyOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *policiesGRPCClient) UpdatePolicy(ctx context.Context, req *iampb.UpdatePolicyRequest, opts ...gax.CallOption) (*UpdatePolicyOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "policy.name", url.QueryEscape(req.GetPolicy().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).UpdatePolicy[0:len((*c.CallOptions).UpdatePolicy):len((*c.CallOptions).UpdatePolicy)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.policiesClient.UpdatePolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &UpdatePolicyOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *policiesGRPCClient) DeletePolicy(ctx context.Context, req *iampb.DeletePolicyRequest, opts ...gax.CallOption) (*DeletePolicyOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).DeletePolicy[0:len((*c.CallOptions).DeletePolicy):len((*c.CallOptions).DeletePolicy)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.policiesClient.DeletePolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &DeletePolicyOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *policiesGRPCClient) ListApplicablePolicies(ctx context.Context, req *iampb.ListApplicablePoliciesRequest, opts ...gax.CallOption) *PolicyIterator {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "attachment_point", url.QueryEscape(req.GetAttachmentPoint())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ListApplicablePolicies[0:len((*c.CallOptions).ListApplicablePolicies):len((*c.CallOptions).ListApplicablePolicies)], opts...)
	it := &PolicyIterator{}
	req = proto.Clone(req).(*iampb.ListApplicablePoliciesRequest)
	it.InternalFetch = func(pageSize int, pageToken string) ([]*iampb.Policy, string, error) {
		resp := &iampb.ListApplicablePoliciesResponse{}
		if pageToken != "" {
			req.PageToken = pageToken
		}
		if pageSize > math.MaxInt32 {
			req.PageSize = math.MaxInt32
		} else if pageSize != 0 {
			req.PageSize = int32(pageSize)
		}
		err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
			var err error
			resp, err = c.policiesClient.ListApplicablePolicies(ctx, req, settings.GRPC...)
			return err
		}, opts...)
		if err != nil {
			return nil, "", err
		}

		it.Response = resp
		return resp.GetPolicies(), resp.GetNextPageToken(), nil
	}
	fetch := func(pageSize int, pageToken string) (string, error) {
		items, nextPageToken, err := it.InternalFetch(pageSize, pageToken)
		if err != nil {
			return "", err
		}
		it.items = append(it.items, items...)
		return nextPageToken, nil
	}

	it.pageInfo, it.nextFunc = iterator.NewPageInfo(fetch, it.bufLen, it.takeBuf)
	it.pageInfo.MaxSize = int(req.GetPageSize())
	it.pageInfo.Token = req.GetPageToken()

	return it
}

func (c *policiesGRPCClient) GetOperation(ctx context.Context, req *longrunningpb.GetOperationRequest, opts ...gax.CallOption) (*longrunningpb.Operation, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "name", url.QueryEscape(req.GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).GetOperation[0:len((*c.CallOptions).GetOperation):len((*c.CallOptions).GetOperation)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.operationsClient.GetOperation(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreatePolicyOperation manages a long-running operation from CreatePolicy.
type CreatePolicyOperation struct {
	lro *longrunning.Operation
}

// CreatePolicyOperation returns a new CreatePolicyOperation from a given name.
// The name must be that of a previously created CreatePolicyOperation, possibly from a different process.
func (c *policiesGRPCClient) CreatePolicyOperation(name string) *CreatePolicyOperation {
	return &CreatePolicyOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *CreatePolicyOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*iampb.Policy, error) {
	var resp iampb.Policy
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *CreatePolicyOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*iampb.Policy, error) {
	var resp iampb.Policy
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *CreatePolicyOperation) Metadata() (*iampb.PolicyOperationMetadata, error) {
	var meta iampb.PolicyOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *CreatePolicyOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *CreatePolicyOperation) Name() string {
	return op.lro.Name()
}

// DeletePolicyOperation manages a long-running operation from DeletePolicy.
type DeletePolicyOperation struct {
	lro *longrunning.Operation
}

// DeletePolicyOperation returns a new DeletePolicyOperation from a given name.
// The name must be that of a previously created DeletePolicyOperation, possibly from a different process.
func (c *policiesGRPCClient) DeletePolicyOperation(name string) *DeletePolicyOperation {
	return &DeletePolicyOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *DeletePolicyOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*iampb.Policy, error) {
	var resp iampb.Policy
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *DeletePolicyOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*iampb.Policy, error) {
	var resp iampb.Policy
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *DeletePolicyOperation) Metadata() (*iampb.PolicyOperationMetadata, error) {
	var meta iampb.PolicyOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *DeletePolicyOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *DeletePolicyOperation) Name() string {
	return op.lro.Name()
}

// UpdatePolicyOperation manages a long-running operation from UpdatePolicy.
type UpdatePolicyOperation struct {
	lro *longrunning.Operation
}

// UpdatePolicyOperation returns a new UpdatePolicyOperation from a given name.
// The name must be that of a previously created UpdatePolicyOperation, possibly from a different process.
func (c *policiesGRPCClient) UpdatePolicyOperation(name string) *UpdatePolicyOperation {
	return &UpdatePolicyOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *UpdatePolicyOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*iampb.Policy, error) {
	var resp iampb.Policy
	if err := op.lro.WaitWithInterval(ctx, &resp, time.Minute, opts...); err != nil {
		return nil, err
	}
	return &resp, nil
}

// Poll fetches the latest state of the long-running operation.
//
// Poll also fetches the latest metadata, which can be retrieved by Metadata.
//
// If Poll fails, the error is returned and op is unmodified. If Poll succeeds and
// the operation has completed with failure, the error is returned and op.Done will return true.
// If Poll succeeds and the operation has completed successfully,
// op.Done will return true, and the response of the operation is returned.
// If Poll succeeds and the operation has not completed, the returned response and error are both nil.
func (op *UpdatePolicyOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*iampb.Policy, error) {
	var resp iampb.Policy
	if err := op.lro.Poll(ctx, &resp, opts...); err != nil {
		return nil, err
	}
	if !op.Done() {
		return nil, nil
	}
	return &resp, nil
}

// Metadata returns metadata associated with the long-running operation.
// Metadata itself does not contact the server, but Poll does.
// To get the latest metadata, call this method after a successful call to Poll.
// If the metadata is not available, the returned metadata and error are both nil.
func (op *UpdatePolicyOperation) Metadata() (*iampb.PolicyOperationMetadata, error) {
	var meta iampb.PolicyOperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *UpdatePolicyOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *UpdatePolicyOperation) Name() string {
	return op.lro.Name()
}

// PolicyIterator manages a stream of *iampb.Policy.
type PolicyIterator struct {
	items    []*iampb.Policy
	pageInfo *iterator.PageInfo
	nextFunc func() error

	// Response is the raw response for the current page.
	// It must be cast to the RPC response type.
	// Calling Next() or InternalFetch() updates this value.
	Response interface{}

	// InternalFetch is for use by the Google Cloud Libraries only.
	// It is not part of the stable interface of this package.
	//
	// InternalFetch returns results from a single call to the underlying RPC.
	// The number of results is no greater than pageSize.
	// If there are no more results, nextPageToken is empty and err is nil.
	InternalFetch func(pageSize int, pageToken string) (results []*iampb.Policy, nextPageToken string, err error)
}

// PageInfo supports pagination. See the google.golang.org/api/iterator package for details.
func (it *PolicyIterator) PageInfo() *iterator.PageInfo {
	return it.pageInfo
}

// Next returns the next result. Its second return value is iterator.Done if there are no more
// results. Once Next returns Done, all subsequent calls will return Done.
func (it *PolicyIterator) Next() (*iampb.Policy, error) {
	var item *iampb.Policy
	if err := it.nextFunc(); err != nil {
		return item, err
	}
	item = it.items[0]
	it.items = it.items[1:]
	return item, nil
}

func (it *PolicyIterator) bufLen() int {
	return len(it.items)
}

func (it *PolicyIterator) takeBuf() interface{} {
	b := it.items
	it.items = nil
	return b
}
