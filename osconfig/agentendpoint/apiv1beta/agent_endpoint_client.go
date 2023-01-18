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

package agentendpoint

import (
	"context"
	"math"
	"time"

	agentendpointpb "cloud.google.com/go/osconfig/agentendpoint/apiv1beta/agentendpointpb"
	gax "github.com/googleapis/gax-go/v2"
	"google.golang.org/api/option"
	"google.golang.org/api/option/internaloption"
	gtransport "google.golang.org/api/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

var newClientHook clientHook

// CallOptions contains the retry settings for each method of Client.
type CallOptions struct {
	ReceiveTaskNotification    []gax.CallOption
	StartNextTask              []gax.CallOption
	ReportTaskProgress         []gax.CallOption
	ReportTaskComplete         []gax.CallOption
	LookupEffectiveGuestPolicy []gax.CallOption
	RegisterAgent              []gax.CallOption
}

func defaultGRPCClientOptions() []option.ClientOption {
	return []option.ClientOption{
		internaloption.WithDefaultEndpoint("osconfig.googleapis.com:443"),
		internaloption.WithDefaultMTLSEndpoint("osconfig.mtls.googleapis.com:443"),
		internaloption.WithDefaultAudience("https://osconfig.googleapis.com/"),
		internaloption.WithDefaultScopes(DefaultAuthScopes()...),
		internaloption.EnableJwtWithScope(),
		option.WithGRPCDialOption(grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(math.MaxInt32))),
	}
}

func defaultCallOptions() *CallOptions {
	return &CallOptions{
		ReceiveTaskNotification: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.DeadlineExceeded,
					codes.Canceled,
					codes.Aborted,
					codes.Internal,
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		StartNextTask: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ReportTaskProgress: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		ReportTaskComplete: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		LookupEffectiveGuestPolicy: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
		RegisterAgent: []gax.CallOption{
			gax.WithRetry(func() gax.Retryer {
				return gax.OnCodes([]codes.Code{
					codes.Unavailable,
				}, gax.Backoff{
					Initial:    1000 * time.Millisecond,
					Max:        60000 * time.Millisecond,
					Multiplier: 1.30,
				})
			}),
		},
	}
}

// internalClient is an interface that defines the methods available from OS Config API.
type internalClient interface {
	Close() error
	setGoogleClientInfo(...string)
	Connection() *grpc.ClientConn
	ReceiveTaskNotification(context.Context, *agentendpointpb.ReceiveTaskNotificationRequest, ...gax.CallOption) (agentendpointpb.AgentEndpointService_ReceiveTaskNotificationClient, error)
	StartNextTask(context.Context, *agentendpointpb.StartNextTaskRequest, ...gax.CallOption) (*agentendpointpb.StartNextTaskResponse, error)
	ReportTaskProgress(context.Context, *agentendpointpb.ReportTaskProgressRequest, ...gax.CallOption) (*agentendpointpb.ReportTaskProgressResponse, error)
	ReportTaskComplete(context.Context, *agentendpointpb.ReportTaskCompleteRequest, ...gax.CallOption) (*agentendpointpb.ReportTaskCompleteResponse, error)
	LookupEffectiveGuestPolicy(context.Context, *agentendpointpb.LookupEffectiveGuestPolicyRequest, ...gax.CallOption) (*agentendpointpb.EffectiveGuestPolicy, error)
	RegisterAgent(context.Context, *agentendpointpb.RegisterAgentRequest, ...gax.CallOption) (*agentendpointpb.RegisterAgentResponse, error)
}

// Client is a client for interacting with OS Config API.
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
//
// OS Config agent endpoint API.
type Client struct {
	// The internal transport-dependent client.
	internalClient internalClient

	// The call options for this service.
	CallOptions *CallOptions
}

// Wrapper methods routed to the internal client.

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *Client) Close() error {
	return c.internalClient.Close()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *Client) setGoogleClientInfo(keyval ...string) {
	c.internalClient.setGoogleClientInfo(keyval...)
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *Client) Connection() *grpc.ClientConn {
	return c.internalClient.Connection()
}

// ReceiveTaskNotification stream established by client to receive Task notifications.
func (c *Client) ReceiveTaskNotification(ctx context.Context, req *agentendpointpb.ReceiveTaskNotificationRequest, opts ...gax.CallOption) (agentendpointpb.AgentEndpointService_ReceiveTaskNotificationClient, error) {
	return c.internalClient.ReceiveTaskNotification(ctx, req, opts...)
}

// StartNextTask signals the start of a task execution and returns the task info.
func (c *Client) StartNextTask(ctx context.Context, req *agentendpointpb.StartNextTaskRequest, opts ...gax.CallOption) (*agentendpointpb.StartNextTaskResponse, error) {
	return c.internalClient.StartNextTask(ctx, req, opts...)
}

// ReportTaskProgress signals an intermediary progress checkpoint in task execution.
func (c *Client) ReportTaskProgress(ctx context.Context, req *agentendpointpb.ReportTaskProgressRequest, opts ...gax.CallOption) (*agentendpointpb.ReportTaskProgressResponse, error) {
	return c.internalClient.ReportTaskProgress(ctx, req, opts...)
}

// ReportTaskComplete signals that the task execution is complete and optionally returns the next
// task.
func (c *Client) ReportTaskComplete(ctx context.Context, req *agentendpointpb.ReportTaskCompleteRequest, opts ...gax.CallOption) (*agentendpointpb.ReportTaskCompleteResponse, error) {
	return c.internalClient.ReportTaskComplete(ctx, req, opts...)
}

// LookupEffectiveGuestPolicy lookup the effective guest policy that applies to a VM instance. This
// lookup merges all policies that are assigned to the instance ancestry.
func (c *Client) LookupEffectiveGuestPolicy(ctx context.Context, req *agentendpointpb.LookupEffectiveGuestPolicyRequest, opts ...gax.CallOption) (*agentendpointpb.EffectiveGuestPolicy, error) {
	return c.internalClient.LookupEffectiveGuestPolicy(ctx, req, opts...)
}

// RegisterAgent registers the agent running on the VM.
func (c *Client) RegisterAgent(ctx context.Context, req *agentendpointpb.RegisterAgentRequest, opts ...gax.CallOption) (*agentendpointpb.RegisterAgentResponse, error) {
	return c.internalClient.RegisterAgent(ctx, req, opts...)
}

// gRPCClient is a client for interacting with OS Config API over gRPC transport.
//
// Methods, except Close, may be called concurrently. However, fields must not be modified concurrently with method calls.
type gRPCClient struct {
	// Connection pool of gRPC connections to the service.
	connPool gtransport.ConnPool

	// flag to opt out of default deadlines via GOOGLE_API_GO_EXPERIMENTAL_DISABLE_DEFAULT_DEADLINE
	disableDeadlines bool

	// Points back to the CallOptions field of the containing Client
	CallOptions **CallOptions

	// The gRPC API client.
	client agentendpointpb.AgentEndpointServiceClient

	// The x-goog-* metadata to be sent with each request.
	xGoogMetadata metadata.MD
}

// NewClient creates a new agent endpoint service client based on gRPC.
// The returned client must be Closed when it is done being used to clean up its underlying connections.
//
// OS Config agent endpoint API.
func NewClient(ctx context.Context, opts ...option.ClientOption) (*Client, error) {
	clientOpts := defaultGRPCClientOptions()
	if newClientHook != nil {
		hookOpts, err := newClientHook(ctx, clientHookParams{})
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
	client := Client{CallOptions: defaultCallOptions()}

	c := &gRPCClient{
		connPool:         connPool,
		disableDeadlines: disableDeadlines,
		client:           agentendpointpb.NewAgentEndpointServiceClient(connPool),
		CallOptions:      &client.CallOptions,
	}
	c.setGoogleClientInfo()

	client.internalClient = c

	return &client, nil
}

// Connection returns a connection to the API service.
//
// Deprecated: Connections are now pooled so this method does not always
// return the same resource.
func (c *gRPCClient) Connection() *grpc.ClientConn {
	return c.connPool.Conn()
}

// setGoogleClientInfo sets the name and version of the application in
// the `x-goog-api-client` header passed on each request. Intended for
// use by Google-written clients.
func (c *gRPCClient) setGoogleClientInfo(keyval ...string) {
	kv := append([]string{"gl-go", versionGo()}, keyval...)
	kv = append(kv, "gapic", getVersionClient(), "gax", gax.Version, "grpc", grpc.Version)
	c.xGoogMetadata = metadata.Pairs("x-goog-api-client", gax.XGoogHeader(kv...))
}

// Close closes the connection to the API service. The user should invoke this when
// the client is no longer required.
func (c *gRPCClient) Close() error {
	return c.connPool.Close()
}

func (c *gRPCClient) ReceiveTaskNotification(ctx context.Context, req *agentendpointpb.ReceiveTaskNotificationRequest, opts ...gax.CallOption) (agentendpointpb.AgentEndpointService_ReceiveTaskNotificationClient, error) {
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ReceiveTaskNotification[0:len((*c.CallOptions).ReceiveTaskNotification):len((*c.CallOptions).ReceiveTaskNotification)], opts...)
	var resp agentendpointpb.AgentEndpointService_ReceiveTaskNotificationClient
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ReceiveTaskNotification(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) StartNextTask(ctx context.Context, req *agentendpointpb.StartNextTaskRequest, opts ...gax.CallOption) (*agentendpointpb.StartNextTaskResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).StartNextTask[0:len((*c.CallOptions).StartNextTask):len((*c.CallOptions).StartNextTask)], opts...)
	var resp *agentendpointpb.StartNextTaskResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.StartNextTask(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ReportTaskProgress(ctx context.Context, req *agentendpointpb.ReportTaskProgressRequest, opts ...gax.CallOption) (*agentendpointpb.ReportTaskProgressResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ReportTaskProgress[0:len((*c.CallOptions).ReportTaskProgress):len((*c.CallOptions).ReportTaskProgress)], opts...)
	var resp *agentendpointpb.ReportTaskProgressResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ReportTaskProgress(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) ReportTaskComplete(ctx context.Context, req *agentendpointpb.ReportTaskCompleteRequest, opts ...gax.CallOption) (*agentendpointpb.ReportTaskCompleteResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).ReportTaskComplete[0:len((*c.CallOptions).ReportTaskComplete):len((*c.CallOptions).ReportTaskComplete)], opts...)
	var resp *agentendpointpb.ReportTaskCompleteResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.ReportTaskComplete(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) LookupEffectiveGuestPolicy(ctx context.Context, req *agentendpointpb.LookupEffectiveGuestPolicyRequest, opts ...gax.CallOption) (*agentendpointpb.EffectiveGuestPolicy, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).LookupEffectiveGuestPolicy[0:len((*c.CallOptions).LookupEffectiveGuestPolicy):len((*c.CallOptions).LookupEffectiveGuestPolicy)], opts...)
	var resp *agentendpointpb.EffectiveGuestPolicy
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.LookupEffectiveGuestPolicy(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *gRPCClient) RegisterAgent(ctx context.Context, req *agentendpointpb.RegisterAgentRequest, opts ...gax.CallOption) (*agentendpointpb.RegisterAgentResponse, error) {
	if _, ok := ctx.Deadline(); !ok && !c.disableDeadlines {
		cctx, cancel := context.WithTimeout(ctx, 60000*time.Millisecond)
		defer cancel()
		ctx = cctx
	}
	ctx = insertMetadata(ctx, c.xGoogMetadata)
	opts = append((*c.CallOptions).RegisterAgent[0:len((*c.CallOptions).RegisterAgent):len((*c.CallOptions).RegisterAgent)], opts...)
	var resp *agentendpointpb.RegisterAgentResponse
	err := gax.Invoke(ctx, func(ctx context.Context, settings gax.CallSettings) error {
		var err error
		resp, err = c.client.RegisterAgent(ctx, req, settings.GRPC...)
		return err
	}, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
