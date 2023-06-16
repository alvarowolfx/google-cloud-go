// Copyright 2023 Google LLC
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

package storage

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"time"

	storagepb "cloud.google.com/go/bigquery/storage/apiv1beta1/storagepb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newBigQueryStorageClientHook clientHook

// BigQueryStorageCallOptions contains the retry settings for each method of BigQueryStorageClient.
type BigQueryStorageCallOptions struct {
	CreateReadSession             []gax.CallOption
	ReadRows                      []gax.CallOption
	BatchCreateReadSessionStreams []gax.CallOption
	FinalizeStream                []gax.CallOption
	SplitReadStream               []gax.CallOption
}

func defaultBigQueryStorageGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("bigquerystorage.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("bigquerystorage.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://bigquerystorage.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultBigQueryStorageCallOptions() *BigQueryStorageCallOptions {
	return &BigQueryStorageCallOptions{
		CreateReadSession: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ReadRows: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		BatchCreateReadSessionStreams: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		FinalizeStream: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		SplitReadStream: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

func defaultBigQueryStorageRESTCallOptions() *BigQueryStorageCallOptions {
	return &BigQueryStorageCallOptions{
		CreateReadSession: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
		ReadRows: []gax.CallOption{
			gax.WithTimeout(86400000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusServiceUnavailable)
			}),
		},
		BatchCreateReadSessionStreams: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
		FinalizeStream: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
		SplitReadStream: []gax.CallOption{
			gax.WithTimeout(600000 * time.Millisecond),
			gax.WithRetry(func() gax.Retryer {
				return gax.OnHTTPCodes(gax.Backoff{
					Initial:    100 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				},
					http.StatusGatewayTimeout,
					http.StatusServiceUnavailable)
			}),
		},
	}
}

// internalBigQueryStorageClient is an interface that defines the methods available from BigQuery Storage API.
type internalBigQueryStorageClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	CreateReadSession(context.Context, *storagepb.CreateReadSessionRequest, ...gax.CallOption) (*storagepb.ReadSession, error)
	ReadRows(context.Context, *storagepb.ReadRowsRequest, ...gax.CallOption) (storagepb.BigQueryStorage_ReadRowsClient, error)
	BatchCreateReadSessionStreams(context.Context, *storagepb.BatchCreateReadSessionStreamsRequest, ...gax.CallOption) (*storagepb.BatchCreateReadSessionStreamsResponse, error)
	FinalizeStream(context.Context, *storagepb.FinalizeStreamRequest, ...gax.CallOption) error
	SplitReadStream(context.Context, *storagepb.SplitReadStreamRequest, ...gax.CallOption) (*storagepb.SplitReadStreamResponse, error)
}

// BigQueryStorageClient is a client for interacting with BigQuery Storage API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// BigQuery storage API.
//
// The BigQuery storage API can be used to read data stored in BigQuery.
//
// The v1beta1 API is not yet officially deprecated, and will go through a full
// deprecation cycle (https://cloud.google.com/products#product-launch-stages (at https://cloud.google.com/products#product-launch-stages))
// before the service is turned down. However, new code should use the v1 API
// going forward.
type BigQueryStorageClient struct {
	// The internal transport-dependent client.
	internalClient internalBigQueryStorageClient

	// The call options for this service.
	CallOptions *BigQueryStorageCallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *BigQueryStorageClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *BigQueryStorageClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *BigQueryStorageClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// CreateReadSession creates a new read session. A read session divides the contents of a
// BigQuery table into one or more streams, which can then be used to read
// data from the table. The read session also specifies properties of the
// data to be read, such as a list of columns or a push-down filter describing
// the rows to be returned.
//
// A particular row can be read by at most one stream. When the caller has
// reached the end of each stream in the session, then all the data in the
// table has been read.
//
// Read sessions automatically expire 6 hours after they are created and do
// not require manual clean-up by the caller.
func (c *BigQueryStorageClient) CreateReadSession(ctx context.Context, req *storagepb.CreateReadSessionRequest, opts ...gax.CallOption) (*storagepb.ReadSession, error) {
	return c.internalClient.CreateReadSession(ctx, req, opts...)
}

// ReadRows reads rows from the table in the format prescribed by the read session.
// Each response contains one or more table rows, up to a maximum of 10 MiB
// per response; read requests which attempt to read individual rows larger
// than this will fail.
//
// Each request also returns a set of stream statistics reflecting the
// estimated total number of rows in the read stream. This number is computed
// based on the total table size and the number of active streams in the read
// session, and may change as other streams continue to read data.
func (c *BigQueryStorageClient) ReadRows(ctx context.Context, req *storagepb.ReadRowsRequest, opts ...gax.CallOption) (storagepb.BigQueryStorage_ReadRowsClient, error) {
	return c.internalClient.ReadRows(ctx, req, opts...)
}

// BatchCreateReadSessionStreams creates additional streams for a ReadSession. This API can be used to
// dynamically adjust the parallelism of a batch processing task upwards by
// adding additional workers.
func (c *BigQueryStorageClient) BatchCreateReadSessionStreams(ctx context.Context, req *storagepb.BatchCreateReadSessionStreamsRequest, opts ...gax.CallOption) (*storagepb.BatchCreateReadSessionStreamsResponse, error) {
	return c.internalClient.BatchCreateReadSessionStreams(ctx, req, opts...)
}

// FinalizeStream causes a single stream in a ReadSession to gracefully stop. This
// API can be used to dynamically adjust the parallelism of a batch processing
// task downwards without losing data.
//
// This API does not delete the stream – it remains visible in the
// ReadSession, and any data processed by the stream is not released to other
// streams. However, no additional data will be assigned to the stream once
// this call completes. Callers must continue reading data on the stream until
// the end of the stream is reached so that data which has already been
// assigned to the stream will be processed.
//
// This method will return an error if there are no other live streams
// in the Session, or if SplitReadStream() has been called on the given
// Stream.
func (c *BigQueryStorageClient) FinalizeStream(ctx context.Context, req *storagepb.FinalizeStreamRequest, opts ...gax.CallOption) error {
	return c.internalClient.FinalizeStream(ctx, req, opts...)
}

// SplitReadStream splits a given read stream into two Streams. These streams are referred to
// as the primary and the residual of the split. The original stream can still
// be read from in the same manner as before. Both of the returned streams can
// also be read from, and the total rows return by both child streams will be
// the same as the rows read from the original stream.
//
// Moreover, the two child streams will be allocated back to back in the
// original Stream. Concretely, it is guaranteed that for streams Original,
// Primary, and Residual, that Original[0-j] = Primary[0-j] and
// Original[j-n] = Residual[0-m] once the streams have been read to
// completion.
//
// This method is guaranteed to be idempotent.
func (c *BigQueryStorageClient) SplitReadStream(ctx context.Context, req *storagepb.SplitReadStreamRequest, opts ...gax.CallOption) (*storagepb.SplitReadStreamResponse, error) {
	return c.internalClient.SplitReadStream(ctx, req, opts...)
}

// bigQueryStorageGRPCClient is a client for interacting with BigQuery Storage API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type bigQueryStorageGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// Points back to the CallOptions field of the containing BigQueryStorageClient
	CallOptions **BigQueryStorageCallOptions

	// The gRPC API client.
	bigQueryStorageClient storagepb.BigQueryStorageClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewBigQueryStorageClient creates a new big query storage client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// BigQuery storage API.
//
// The BigQuery storage API can be used to read data stored in BigQuery.
//
// The v1beta1 API is not yet officially deprecated, and will go through a full
// deprecation cycle (https://cloud.google.com/products#product-launch-stages (at https://cloud.google.com/products#product-launch-stages))
// before the service is turned down. However, new code should use the v1 API
// going forward.
func NewBigQueryStorageClient(ctx context.Context, opts ...option.ClientOption) (*BigQueryStorageClient, error) {
	clientOpts := defaultBigQueryStorageGRPCClientOptions()
	if newBigQueryStorageClientHook != nil {
		hookOpts, err := newBigQueryStorageClientHook(ctx, clientHookParams{})
		if err != nil {
			return nil, err
		}
		clientOpts = append(clientOpts, hookOpts...)
	}

	connPool, err := gtransport.DialPool(ctx, append(clientOpts, opts...)...)
	if err != nil {
		return nil, err
	}
	client := BigQueryStorageClient{CallOptions: defaultBigQueryStorageCallOptions()}

	c := &bigQueryStorageGRPCClient{
		connPool:              connPool,
		bigQueryStorageClient: storagepb.NewBigQueryStorageClient(connPool),
		CallOptions:           &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *bigQueryStorageGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *bigQueryStorageGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *bigQueryStorageGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type bigQueryStorageRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing BigQueryStorageClient
	CallOptions **BigQueryStorageCallOptions
}

// NewBigQueryStorageRESTClient creates a new big query storage rest client.
//
// BigQuery storage API.
//
// The BigQuery storage API can be used to read data stored in BigQuery.
//
// The v1beta1 API is not yet officially deprecated, and will go through a full
// deprecation cycle (https://cloud.google.com/products#product-launch-stages (at https://cloud.google.com/products#product-launch-stages))
// before the service is turned down. However, new code should use the v1 API
// going forward.
func NewBigQueryStorageRESTClient(ctx context.Context, opts ...option.ClientOption) (*BigQueryStorageClient, error) {
	clientOpts := append(defaultBigQueryStorageRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultBigQueryStorageRESTCallOptions()
	c := &bigQueryStorageRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	return &BigQueryStorageClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultBigQueryStorageRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://bigquerystorage.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://bigquerystorage.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://bigquerystorage.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *bigQueryStorageRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *bigQueryStorageRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *bigQueryStorageRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *bigQueryStorageGRPCClient) CreateReadSession(ctx context.Context, req *storagepb.CreateReadSessionRequest, opts ...gax.CallOption) (*storagepb.ReadSession, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "table_reference.project_id", url.QueryEscape(req.GetTableReference().GetProjectId()), "table_reference.dataset_id", url.QueryEscape(req.GetTableReference().GetDatasetId())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).CreateReadSession[0:len((*c.CallOptions).CreateReadSession):len((*c.CallOptions).CreateReadSession)], opts...)
	var resp *storagepb.ReadSession
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.bigQueryStorageClient.CreateReadSession(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *bigQueryStorageGRPCClient) ReadRows(ctx context.Context, req *storagepb.ReadRowsRequest, opts ...gax.CallOption) (storagepb.BigQueryStorage_ReadRowsClient, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "read_position.stream.name", url.QueryEscape(req.GetReadPosition().GetStream().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).ReadRows[0:len((*c.CallOptions).ReadRows):len((*c.CallOptions).ReadRows)], opts...)
	var resp storagepb.BigQueryStorage_ReadRowsClient
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.bigQueryStorageClient.ReadRows(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *bigQueryStorageGRPCClient) BatchCreateReadSessionStreams(ctx context.Context, req *storagepb.BatchCreateReadSessionStreamsRequest, opts ...gax.CallOption) (*storagepb.BatchCreateReadSessionStreamsResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "session.name", url.QueryEscape(req.GetSession().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).BatchCreateReadSessionStreams[0:len((*c.CallOptions).BatchCreateReadSessionStreams):len((*c.CallOptions).BatchCreateReadSessionStreams)], opts...)
	var resp *storagepb.BatchCreateReadSessionStreamsResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.bigQueryStorageClient.BatchCreateReadSessionStreams(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *bigQueryStorageGRPCClient) FinalizeStream(ctx context.Context, req *storagepb.FinalizeStreamRequest, opts ...gax.CallOption) error {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "stream.name", url.QueryEscape(req.GetStream().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).FinalizeStream[0:len((*c.CallOptions).FinalizeStream):len((*c.CallOptions).FinalizeStream)], opts...)
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		_, err = c.bigQueryStorageClient.FinalizeStream(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	return err
}

func (c *bigQueryStorageGRPCClient) SplitReadStream(ctx context.Context, req *storagepb.SplitReadStreamRequest, opts ...gax.CallOption) (*storagepb.SplitReadStreamResponse, error) {
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "original_stream.name", url.QueryEscape(req.GetOriginalStream().GetName())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).SplitReadStream[0:len((*c.CallOptions).SplitReadStream):len((*c.CallOptions).SplitReadStream)], opts...)
	var resp *storagepb.SplitReadStreamResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.bigQueryStorageClient.SplitReadStream(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// CreateReadSession creates a new read session. A read session divides the contents of a
// BigQuery table into one or more streams, which can then be used to read
// data from the table. The read session also specifies properties of the
// data to be read, such as a list of columns or a push-down filter describing
// the rows to be returned.
//
// A particular row can be read by at most one stream. When the caller has
// reached the end of each stream in the session, then all the data in the
// table has been read.
//
// Read sessions automatically expire 6 hours after they are created and do
// not require manual clean-up by the caller.
func (c *bigQueryStorageRESTClient) CreateReadSession(ctx context.Context, req *storagepb.CreateReadSessionRequest, opts ...gax.CallOption) (*storagepb.ReadSession, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v", req.GetTableReference().GetProjectId())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v&%s=%v", "table_reference.project_id", url.QueryEscape(req.GetTableReference().GetProjectId()), "table_reference.dataset_id", url.QueryEscape(req.GetTableReference().GetDatasetId())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).CreateReadSession[0:len((*c.CallOptions).CreateReadSession):len((*c.CallOptions).CreateReadSession)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &storagepb.ReadSession{}
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

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// ReadRows reads rows from the table in the format prescribed by the read session.
// Each response contains one or more table rows, up to a maximum of 10 MiB
// per response; read requests which attempt to read individual rows larger
// than this will fail.
//
// Each request also returns a set of stream statistics reflecting the
// estimated total number of rows in the read stream. This number is computed
// based on the total table size and the number of active streams in the read
// session, and may change as other streams continue to read data.
func (c *bigQueryStorageRESTClient) ReadRows(ctx context.Context, req *storagepb.ReadRowsRequest, opts ...gax.CallOption) (storagepb.BigQueryStorage_ReadRowsClient, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v", req.GetReadPosition().GetStream().GetName())

	params := url.Values{}
	if req.GetReadPosition().GetOffset() != 0 {
		params.Add("readPosition.offset", fmt.Sprintf("%v", req.GetReadPosition().GetOffset()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "read_position.stream.name", url.QueryEscape(req.GetReadPosition().GetStream().GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	var streamClient *readRowsRESTClient
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

		if err = googleapi.CheckResponse(httpRsp); err != nil {
			return err
		}

		streamClient = &readRowsRESTClient{
			ctx:    ctx,
			md:     metadata.MD(httpRsp.Header),
			stream: gax.NewProtoJSONStreamReader(httpRsp.Body, (&storagepb.ReadRowsResponse{}).ProtoReflect().Type()),
		}
		return nil
	}, opts...)

	return streamClient, e
}

// readRowsRESTClient is the stream client used to consume the server stream created by
// the REST implementation of ReadRows.
type readRowsRESTClient struct {
	ctx    context.Context
	md     metadata.MD
	stream *gax.ProtoJSONStream
}

func (c *readRowsRESTClient) Recv() (*storagepb.ReadRowsResponse, error) {
	if err := c.ctx.Err(); err != nil {
		defer c.stream.Close()
		return nil, err
	}
	msg, err := c.stream.Recv()
	if err != nil {
		defer c.stream.Close()
		return nil, err
	}
	res := msg.(*storagepb.ReadRowsResponse)
	return res, nil
}

func (c *readRowsRESTClient) Header() (metadata.MD, error) {
	return c.md, nil
}

func (c *readRowsRESTClient) Trailer() metadata.MD {
	return c.md
}

func (c *readRowsRESTClient) CloseSend() error {
	// This is a no-op to fulfill the interface.
	return fmt.Errorf("this method is not implemented for a server-stream")
}

func (c *readRowsRESTClient) Context() context.Context {
	return c.ctx
}

func (c *readRowsRESTClient) SendMsg(m interface{}) error {
	// This is a no-op to fulfill the interface.
	return fmt.Errorf("this method is not implemented for a server-stream")
}

func (c *readRowsRESTClient) RecvMsg(m interface{}) error {
	// This is a no-op to fulfill the interface.
	return fmt.Errorf("this method is not implemented, use Recv")
}

// BatchCreateReadSessionStreams creates additional streams for a ReadSession. This API can be used to
// dynamically adjust the parallelism of a batch processing task upwards by
// adding additional workers.
func (c *bigQueryStorageRESTClient) BatchCreateReadSessionStreams(ctx context.Context, req *storagepb.BatchCreateReadSessionStreamsRequest, opts ...gax.CallOption) (*storagepb.BatchCreateReadSessionStreamsResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v", req.GetSession().GetName())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "session.name", url.QueryEscape(req.GetSession().GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).BatchCreateReadSessionStreams[0:len((*c.CallOptions).BatchCreateReadSessionStreams):len((*c.CallOptions).BatchCreateReadSessionStreams)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &storagepb.BatchCreateReadSessionStreamsResponse{}
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

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}

// FinalizeStream causes a single stream in a ReadSession to gracefully stop. This
// API can be used to dynamically adjust the parallelism of a batch processing
// task downwards without losing data.
//
// This API does not delete the stream – it remains visible in the
// ReadSession, and any data processed by the stream is not released to other
// streams. However, no additional data will be assigned to the stream once
// this call completes. Callers must continue reading data on the stream until
// the end of the stream is reached so that data which has already been
// assigned to the stream will be processed.
//
// This method will return an error if there are no other live streams
// in the Session, or if SplitReadStream() has been called on the given
// Stream.
func (c *bigQueryStorageRESTClient) FinalizeStream(ctx context.Context, req *storagepb.FinalizeStreamRequest, opts ...gax.CallOption) error {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v", req.GetStream().GetName())

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "stream.name", url.QueryEscape(req.GetStream().GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	return gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
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

		// Returns nil if there is no error, otherwise wraps
		// the response code and body into a non-nil error
		return googleapi.CheckResponse(httpRsp)
	}, opts...)
}

// SplitReadStream splits a given read stream into two Streams. These streams are referred to
// as the primary and the residual of the split. The original stream can still
// be read from in the same manner as before. Both of the returned streams can
// also be read from, and the total rows return by both child streams will be
// the same as the rows read from the original stream.
//
// Moreover, the two child streams will be allocated back to back in the
// original Stream. Concretely, it is guaranteed that for streams Original,
// Primary, and Residual, that Original[0-j] = Primary[0-j] and
// Original[j-n] = Residual[0-m] once the streams have been read to
// completion.
//
// This method is guaranteed to be idempotent.
func (c *bigQueryStorageRESTClient) SplitReadStream(ctx context.Context, req *storagepb.SplitReadStreamRequest, opts ...gax.CallOption) (*storagepb.SplitReadStreamResponse, error) {
	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1beta1/%v", req.GetOriginalStream().GetName())

	params := url.Values{}
	if req.GetFraction() != 0 {
		params.Add("fraction", fmt.Sprintf("%v", req.GetFraction()))
	}

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "original_stream.name", url.QueryEscape(req.GetOriginalStream().GetName())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).SplitReadStream[0:len((*c.CallOptions).SplitReadStream):len((*c.CallOptions).SplitReadStream)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &storagepb.SplitReadStreamResponse{}
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

		buf, err := ioutil.ReadAll(httpRsp.Body)
		if err != nil {
			return err
		}

		if err := unm.Unmarshal(buf, resp); err != nil {
			return maybeUnknownEnum(err)
		}

		return nil
	}, opts...)
	if e != nil {
		return nil, e
	}
	return resp, nil
}
