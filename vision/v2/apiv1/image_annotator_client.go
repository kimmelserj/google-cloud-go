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

package vision

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"time"

	"cloud.google.com/go/longrunning"
	lroauto "cloud.google.com/go/longrunning/autogen"
	visionpb "cloud.google.com/go/vision/v2/apiv1/visionpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	httptransport "google.golang.org/api/transport/http"
	longrunningpb "google.golang.org/genproto/googleapis/longrunning"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
)

var newImageAnnotatorClientHook clientHook

// ImageAnnotatorCallOptions contains the retry settings for each method of ImageAnnotatorClient.
type ImageAnnotatorCallOptions struct {
	BatchAnnotateImages      []gax.CallOption
	BatchAnnotateFiles       []gax.CallOption
	AsyncBatchAnnotateImages []gax.CallOption
	AsyncBatchAnnotateFiles  []gax.CallOption
}

func defaultImageAnnotatorGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("vision.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("vision.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://vision.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultImageAnnotatorCallOptions() *ImageAnnotatorCallOptions {
	return &ImageAnnotatorCallOptions{
		BatchAnnotateImages: []gax.CallOption{
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
		BatchAnnotateFiles: []gax.CallOption{
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
		AsyncBatchAnnotateImages: []gax.CallOption{
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
		AsyncBatchAnnotateFiles: []gax.CallOption{
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

func defaultImageAnnotatorRESTCallOptions() *ImageAnnotatorCallOptions {
	return &ImageAnnotatorCallOptions{
		BatchAnnotateImages: []gax.CallOption{
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
		BatchAnnotateFiles: []gax.CallOption{
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
		AsyncBatchAnnotateImages: []gax.CallOption{
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
		AsyncBatchAnnotateFiles: []gax.CallOption{
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

// internalImageAnnotatorClient is an interface that defines the methods available from Cloud Vision API.
type internalImageAnnotatorClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	BatchAnnotateImages(context.Context, *visionpb.BatchAnnotateImagesRequest, ...gax.CallOption) (*visionpb.BatchAnnotateImagesResponse, error)
	BatchAnnotateFiles(context.Context, *visionpb.BatchAnnotateFilesRequest, ...gax.CallOption) (*visionpb.BatchAnnotateFilesResponse, error)
	AsyncBatchAnnotateImages(context.Context, *visionpb.AsyncBatchAnnotateImagesRequest, ...gax.CallOption) (*AsyncBatchAnnotateImagesOperation, error)
	AsyncBatchAnnotateImagesOperation(name string) *AsyncBatchAnnotateImagesOperation
	AsyncBatchAnnotateFiles(context.Context, *visionpb.AsyncBatchAnnotateFilesRequest, ...gax.CallOption) (*AsyncBatchAnnotateFilesOperation, error)
	AsyncBatchAnnotateFilesOperation(name string) *AsyncBatchAnnotateFilesOperation
}

// ImageAnnotatorClient is a client for interacting with Cloud Vision API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// Service that performs Google Cloud Vision API detection tasks over client
// images, such as face, landmark, logo, label, and text detection. The
// ImageAnnotator service returns detected entities from the images.
type ImageAnnotatorClient struct {
	// The internal transport-dependent client.
	internalClient internalImageAnnotatorClient

	// The call options for this service.
	CallOptions *ImageAnnotatorCallOptions

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient *lroauto.OperationsClient
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *ImageAnnotatorClient) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *ImageAnnotatorClient) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *ImageAnnotatorClient) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// BatchAnnotateImages run image detection and annotation for a batch of images.
func (c *ImageAnnotatorClient) BatchAnnotateImages(ctx context.Context, req *visionpb.BatchAnnotateImagesRequest, opts ...gax.CallOption) (*visionpb.BatchAnnotateImagesResponse, error) {
	return c.internalClient.BatchAnnotateImages(ctx, req, opts...)
}

// BatchAnnotateFiles service that performs image detection and annotation for a batch of files.
// Now only “application/pdf”, “image/tiff” and “image/gif” are supported.
//
// This service will extract at most 5 (customers can specify which 5 in
// AnnotateFileRequest.pages) frames (gif) or pages (pdf or tiff) from each
// file provided and perform detection and annotation for each image
// extracted.
func (c *ImageAnnotatorClient) BatchAnnotateFiles(ctx context.Context, req *visionpb.BatchAnnotateFilesRequest, opts ...gax.CallOption) (*visionpb.BatchAnnotateFilesResponse, error) {
	return c.internalClient.BatchAnnotateFiles(ctx, req, opts...)
}

// AsyncBatchAnnotateImages run asynchronous image detection and annotation for a list of images.
//
// Progress and results can be retrieved through the
// google.longrunning.Operations interface.
// Operation.metadata contains OperationMetadata (metadata).
// Operation.response contains AsyncBatchAnnotateImagesResponse (results).
//
// This service will write image annotation outputs to json files in customer
// GCS bucket, each json file containing BatchAnnotateImagesResponse proto.
func (c *ImageAnnotatorClient) AsyncBatchAnnotateImages(ctx context.Context, req *visionpb.AsyncBatchAnnotateImagesRequest, opts ...gax.CallOption) (*AsyncBatchAnnotateImagesOperation, error) {
	return c.internalClient.AsyncBatchAnnotateImages(ctx, req, opts...)
}

// AsyncBatchAnnotateImagesOperation returns a new AsyncBatchAnnotateImagesOperation from a given name.
// The name must be that of a previously created AsyncBatchAnnotateImagesOperation, possibly from a different process.
func (c *ImageAnnotatorClient) AsyncBatchAnnotateImagesOperation(name string) *AsyncBatchAnnotateImagesOperation {
	return c.internalClient.AsyncBatchAnnotateImagesOperation(name)
}

// AsyncBatchAnnotateFiles run asynchronous image detection and annotation for a list of generic
// files, such as PDF files, which may contain multiple pages and multiple
// images per page. Progress and results can be retrieved through the
// google.longrunning.Operations interface.
// Operation.metadata contains OperationMetadata (metadata).
// Operation.response contains AsyncBatchAnnotateFilesResponse (results).
func (c *ImageAnnotatorClient) AsyncBatchAnnotateFiles(ctx context.Context, req *visionpb.AsyncBatchAnnotateFilesRequest, opts ...gax.CallOption) (*AsyncBatchAnnotateFilesOperation, error) {
	return c.internalClient.AsyncBatchAnnotateFiles(ctx, req, opts...)
}

// AsyncBatchAnnotateFilesOperation returns a new AsyncBatchAnnotateFilesOperation from a given name.
// The name must be that of a previously created AsyncBatchAnnotateFilesOperation, possibly from a different process.
func (c *ImageAnnotatorClient) AsyncBatchAnnotateFilesOperation(name string) *AsyncBatchAnnotateFilesOperation {
	return c.internalClient.AsyncBatchAnnotateFilesOperation(name)
}

// imageAnnotatorGRPCClient is a client for interacting with Cloud Vision API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type imageAnnotatorGRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing ImageAnnotatorClient
	CallOptions **ImageAnnotatorCallOptions

	// The gRPC API client.
	imageAnnotatorClient visionpb.ImageAnnotatorClient

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewImageAnnotatorClient creates a new image annotator client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// Service that performs Google Cloud Vision API detection tasks over client
// images, such as face, landmark, logo, label, and text detection. The
// ImageAnnotator service returns detected entities from the images.
func NewImageAnnotatorClient(ctx context.Context, opts ...option.ClientOption) (*ImageAnnotatorClient, error) {
	clientOpts := defaultImageAnnotatorGRPCClientOptions()
	if newImageAnnotatorClientHook != nil {
		hookOpts, err := newImageAnnotatorClientHook(ctx, clientHookParams{})
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
	client := ImageAnnotatorClient{CallOptions: defaultImageAnnotatorCallOptions()}

	c := &imageAnnotatorGRPCClient{
		connPool:             connPool,
		disableDeadlines:     disableDeadlines,
		imageAnnotatorClient: visionpb.NewImageAnnotatorClient(connPool),
		CallOptions:          &client.CallOptions,
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
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *imageAnnotatorGRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *imageAnnotatorGRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *imageAnnotatorGRPCClient) Close() error {
	return c.connPool.Close()
}

// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type imageAnnotatorRESTClient struct {
	// The http endpoint to connect to.
	endpoint string

	// The http client.
	httpClient *http.Client

	// LROClient is used internally to handle long-running operations.
	// It is exposed so that its CallOptions can be modified if required.
	// Users should not Close this client.
	LROClient **lroauto.OperationsClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD

	// Points back to the CallOptions field of the containing ImageAnnotatorClient
	CallOptions **ImageAnnotatorCallOptions
}

// NewImageAnnotatorRESTClient creates a new image annotator rest client.
//
// Service that performs Google Cloud Vision API detection tasks over client
// images, such as face, landmark, logo, label, and text detection. The
// ImageAnnotator service returns detected entities from the images.
func NewImageAnnotatorRESTClient(ctx context.Context, opts ...option.ClientOption) (*ImageAnnotatorClient, error) {
	clientOpts := append(defaultImageAnnotatorRESTClientOptions(), opts...)
	httpClient, endpoint, err := httptransport.NewClient(ctx, clientOpts...)
	if err != nil {
		return nil, err
	}

	callOpts := defaultImageAnnotatorRESTCallOptions()
	c := &imageAnnotatorRESTClient{
		endpoint:    endpoint,
		httpClient:  httpClient,
		CallOptions: &callOpts,
	}
	c.setGoogleClientInfo()

	lroOpts := []option.ClientOption{
		option.WithHTTPClient(httpClient),
		option.WithEndpoint(endpoint),
	}
	opClient, err := lroauto.NewOperationsRESTClient(ctx, lroOpts...)
	if err != nil {
		return nil, err
	}
	c.LROClient = &opClient

	return &ImageAnnotatorClient{internalClient: c, CallOptions: callOpts}, nil
}

func defaultImageAnnotatorRESTClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("https://vision.googleapis.com"),
		internaloption.WithDefaultMTLSEndpoint("https://vision.mtls.googleapis.com"),
		internaloption.WithDefaultAudience("https://vision.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
	}
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *imageAnnotatorRESTClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "rest", "UNKNOWN")
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *imageAnnotatorRESTClient) Close() error {
	// Replace httpClient with nil to force cleanup.
	c.httpClient = nil
	return nil
}

// Connection returns a connection to the API service.
//
// Deprecated: This method always returns nil.
func (c *imageAnnotatorRESTClient) Connection() *grpc.ClientConn {
	return nil
}
func (c *imageAnnotatorGRPCClient) BatchAnnotateImages(ctx context.Context, req *visionpb.BatchAnnotateImagesRequest, opts ...gax.CallOption) (*visionpb.BatchAnnotateImagesResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).BatchAnnotateImages[0:len((*c.CallOptions).BatchAnnotateImages):len((*c.CallOptions).BatchAnnotateImages)], opts...)
	var resp *visionpb.BatchAnnotateImagesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.imageAnnotatorClient.BatchAnnotateImages(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *imageAnnotatorGRPCClient) BatchAnnotateFiles(ctx context.Context, req *visionpb.BatchAnnotateFilesRequest, opts ...gax.CallOption) (*visionpb.BatchAnnotateFilesResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).BatchAnnotateFiles[0:len((*c.CallOptions).BatchAnnotateFiles):len((*c.CallOptions).BatchAnnotateFiles)], opts...)
	var resp *visionpb.BatchAnnotateFilesResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.imageAnnotatorClient.BatchAnnotateFiles(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *imageAnnotatorGRPCClient) AsyncBatchAnnotateImages(ctx context.Context, req *visionpb.AsyncBatchAnnotateImagesRequest, opts ...gax.CallOption) (*AsyncBatchAnnotateImagesOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).AsyncBatchAnnotateImages[0:len((*c.CallOptions).AsyncBatchAnnotateImages):len((*c.CallOptions).AsyncBatchAnnotateImages)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.imageAnnotatorClient.AsyncBatchAnnotateImages(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &AsyncBatchAnnotateImagesOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

func (c *imageAnnotatorGRPCClient) AsyncBatchAnnotateFiles(ctx context.Context, req *visionpb.AsyncBatchAnnotateFilesRequest, opts ...gax.CallOption) (*AsyncBatchAnnotateFilesOperation, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 600000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	ctx = insertMetadata(ctx, c.xGoogMetadata, md)
	opts = append((*c.CallOptions).AsyncBatchAnnotateFiles[0:len((*c.CallOptions).AsyncBatchAnnotateFiles):len((*c.CallOptions).AsyncBatchAnnotateFiles)], opts...)
	var resp *longrunningpb.Operation
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.imageAnnotatorClient.AsyncBatchAnnotateFiles(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return &AsyncBatchAnnotateFilesOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, resp),
	}, nil
}

// BatchAnnotateImages run image detection and annotation for a batch of images.
func (c *imageAnnotatorRESTClient) BatchAnnotateImages(ctx context.Context, req *visionpb.BatchAnnotateImagesRequest, opts ...gax.CallOption) (*visionpb.BatchAnnotateImagesResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/images:annotate")

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).BatchAnnotateImages[0:len((*c.CallOptions).BatchAnnotateImages):len((*c.CallOptions).BatchAnnotateImages)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &visionpb.BatchAnnotateImagesResponse{}
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

// BatchAnnotateFiles service that performs image detection and annotation for a batch of files.
// Now only “application/pdf”, “image/tiff” and “image/gif” are supported.
//
// This service will extract at most 5 (customers can specify which 5 in
// AnnotateFileRequest.pages) frames (gif) or pages (pdf or tiff) from each
// file provided and perform detection and annotation for each image
// extracted.
func (c *imageAnnotatorRESTClient) BatchAnnotateFiles(ctx context.Context, req *visionpb.BatchAnnotateFilesRequest, opts ...gax.CallOption) (*visionpb.BatchAnnotateFilesResponse, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/files:annotate")

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	opts = append((*c.CallOptions).BatchAnnotateFiles[0:len((*c.CallOptions).BatchAnnotateFiles):len((*c.CallOptions).BatchAnnotateFiles)], opts...)
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &visionpb.BatchAnnotateFilesResponse{}
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

// AsyncBatchAnnotateImages run asynchronous image detection and annotation for a list of images.
//
// Progress and results can be retrieved through the
// google.longrunning.Operations interface.
// Operation.metadata contains OperationMetadata (metadata).
// Operation.response contains AsyncBatchAnnotateImagesResponse (results).
//
// This service will write image annotation outputs to json files in customer
// GCS bucket, each json file containing BatchAnnotateImagesResponse proto.
func (c *imageAnnotatorRESTClient) AsyncBatchAnnotateImages(ctx context.Context, req *visionpb.AsyncBatchAnnotateImagesRequest, opts ...gax.CallOption) (*AsyncBatchAnnotateImagesOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/images:asyncBatchAnnotate")

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
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

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &AsyncBatchAnnotateImagesOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// AsyncBatchAnnotateFiles run asynchronous image detection and annotation for a list of generic
// files, such as PDF files, which may contain multiple pages and multiple
// images per page. Progress and results can be retrieved through the
// google.longrunning.Operations interface.
// Operation.metadata contains OperationMetadata (metadata).
// Operation.response contains AsyncBatchAnnotateFilesResponse (results).
func (c *imageAnnotatorRESTClient) AsyncBatchAnnotateFiles(ctx context.Context, req *visionpb.AsyncBatchAnnotateFilesRequest, opts ...gax.CallOption) (*AsyncBatchAnnotateFilesOperation, error) {
	m := protojson.MarshalOptions{AllowPartial: true, UseEnumNumbers: true}
	jsonReq, err := m.Marshal(req)
	if err != nil {
		return nil, err
	}

	baseUrl, err := url.Parse(c.endpoint)
	if err != nil {
		return nil, err
	}
	baseUrl.Path += fmt.Sprintf("/v1/files:asyncBatchAnnotate")

	params := url.Values{}
	params.Add("$alt", "json;enum-encoding=int")

	baseUrl.RawQuery = params.Encode()

	// Build HTTP headers from client and context metadata.
	md := metadata.Pairs("x-goog-request-params", fmt.Sprintf("%s=%v", "parent", url.QueryEscape(req.GetParent())))

	headers := buildHeaders(ctx, c.xGoogMetadata, md, metadata.Pairs("Content-Type", "application/json"))
	unm := protojson.UnmarshalOptions{AllowPartial: true, DiscardUnknown: true}
	resp := &longrunningpb.Operation{}
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

	override := fmt.Sprintf("/v1/%s", resp.GetName())
	return &AsyncBatchAnnotateFilesOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, resp),
		pollPath: override,
	}, nil
}

// AsyncBatchAnnotateFilesOperation manages a long-running operation from AsyncBatchAnnotateFiles.
type AsyncBatchAnnotateFilesOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// AsyncBatchAnnotateFilesOperation returns a new AsyncBatchAnnotateFilesOperation from a given name.
// The name must be that of a previously created AsyncBatchAnnotateFilesOperation, possibly from a different process.
func (c *imageAnnotatorGRPCClient) AsyncBatchAnnotateFilesOperation(name string) *AsyncBatchAnnotateFilesOperation {
	return &AsyncBatchAnnotateFilesOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// AsyncBatchAnnotateFilesOperation returns a new AsyncBatchAnnotateFilesOperation from a given name.
// The name must be that of a previously created AsyncBatchAnnotateFilesOperation, possibly from a different process.
func (c *imageAnnotatorRESTClient) AsyncBatchAnnotateFilesOperation(name string) *AsyncBatchAnnotateFilesOperation {
	override := fmt.Sprintf("/v1/%s", name)
	return &AsyncBatchAnnotateFilesOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *AsyncBatchAnnotateFilesOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*visionpb.AsyncBatchAnnotateFilesResponse, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp visionpb.AsyncBatchAnnotateFilesResponse
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
func (op *AsyncBatchAnnotateFilesOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*visionpb.AsyncBatchAnnotateFilesResponse, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp visionpb.AsyncBatchAnnotateFilesResponse
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
func (op *AsyncBatchAnnotateFilesOperation) Metadata() (*visionpb.OperationMetadata, error) {
	var meta visionpb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *AsyncBatchAnnotateFilesOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *AsyncBatchAnnotateFilesOperation) Name() string {
	return op.lro.Name()
}

// AsyncBatchAnnotateImagesOperation manages a long-running operation from AsyncBatchAnnotateImages.
type AsyncBatchAnnotateImagesOperation struct {
	lro      *longrunning.Operation
	pollPath string
}

// AsyncBatchAnnotateImagesOperation returns a new AsyncBatchAnnotateImagesOperation from a given name.
// The name must be that of a previously created AsyncBatchAnnotateImagesOperation, possibly from a different process.
func (c *imageAnnotatorGRPCClient) AsyncBatchAnnotateImagesOperation(name string) *AsyncBatchAnnotateImagesOperation {
	return &AsyncBatchAnnotateImagesOperation{
		lro: longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
	}
}

// AsyncBatchAnnotateImagesOperation returns a new AsyncBatchAnnotateImagesOperation from a given name.
// The name must be that of a previously created AsyncBatchAnnotateImagesOperation, possibly from a different process.
func (c *imageAnnotatorRESTClient) AsyncBatchAnnotateImagesOperation(name string) *AsyncBatchAnnotateImagesOperation {
	override := fmt.Sprintf("/v1/%s", name)
	return &AsyncBatchAnnotateImagesOperation{
		lro:      longrunning.InternalNewOperation(*c.LROClient, &longrunningpb.Operation{Name: name}),
		pollPath: override,
	}
}

// Wait blocks until the long-running operation is completed, returning the response and any errors encountered.
//
// See documentation of Poll for error-handling information.
func (op *AsyncBatchAnnotateImagesOperation) Wait(ctx context.Context, opts ...gax.CallOption) (*visionpb.AsyncBatchAnnotateImagesResponse, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp visionpb.AsyncBatchAnnotateImagesResponse
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
func (op *AsyncBatchAnnotateImagesOperation) Poll(ctx context.Context, opts ...gax.CallOption) (*visionpb.AsyncBatchAnnotateImagesResponse, error) {
	opts = append([]gax.CallOption{gax.WithPath(op.pollPath)}, opts...)
	var resp visionpb.AsyncBatchAnnotateImagesResponse
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
func (op *AsyncBatchAnnotateImagesOperation) Metadata() (*visionpb.OperationMetadata, error) {
	var meta visionpb.OperationMetadata
	if err := op.lro.Metadata(&meta); err == longrunning.ErrNoMetadata {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &meta, nil
}

// Done reports whether the long-running operation has completed.
func (op *AsyncBatchAnnotateImagesOperation) Done() bool {
	return op.lro.Done()
}

// Name returns the name of the long-running operation.
// The name is assigned by the server and is unique within the service from which the operation is created.
func (op *AsyncBatchAnnotateImagesOperation) Name() string {
	return op.lro.Name()
}
