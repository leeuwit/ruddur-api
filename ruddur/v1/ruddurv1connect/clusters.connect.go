// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: ruddur/v1/clusters.proto

package ruddurv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/leeuwit/ruddur-api/ruddur/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ClusterServiceName is the fully-qualified name of the ClusterService service.
	ClusterServiceName = "ruddur.v1.ClusterService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ClusterServiceListProcedure is the fully-qualified name of the ClusterService's List RPC.
	ClusterServiceListProcedure = "/ruddur.v1.ClusterService/List"
	// ClusterServiceCreateProcedure is the fully-qualified name of the ClusterService's Create RPC.
	ClusterServiceCreateProcedure = "/ruddur.v1.ClusterService/Create"
	// ClusterServiceDeleteProcedure is the fully-qualified name of the ClusterService's Delete RPC.
	ClusterServiceDeleteProcedure = "/ruddur.v1.ClusterService/Delete"
	// ClusterServiceKubeconfigProcedure is the fully-qualified name of the ClusterService's Kubeconfig
	// RPC.
	ClusterServiceKubeconfigProcedure = "/ruddur.v1.ClusterService/Kubeconfig"
)

// ClusterServiceClient is a client for the ruddur.v1.ClusterService service.
type ClusterServiceClient interface {
	List(context.Context, *connect_go.Request[v1.ClusterServiceListRequest]) (*connect_go.Response[v1.ClusterServiceListResponse], error)
	Create(context.Context, *connect_go.Request[v1.ClusterServiceCreateRequest]) (*connect_go.Response[v1.ClusterServiceCreateResponse], error)
	// TODO: Edit (extend / shrink)
	Delete(context.Context, *connect_go.Request[v1.ClusterServiceDeleteRequest]) (*connect_go.Response[v1.ClusterServiceDeleteResponse], error)
	Kubeconfig(context.Context, *connect_go.Request[v1.ClusterServiceKubeconfigRequest]) (*connect_go.Response[v1.ClusterServiceKubeconfigResponse], error)
}

// NewClusterServiceClient constructs a client for the ruddur.v1.ClusterService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewClusterServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ClusterServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &clusterServiceClient{
		list: connect_go.NewClient[v1.ClusterServiceListRequest, v1.ClusterServiceListResponse](
			httpClient,
			baseURL+ClusterServiceListProcedure,
			opts...,
		),
		create: connect_go.NewClient[v1.ClusterServiceCreateRequest, v1.ClusterServiceCreateResponse](
			httpClient,
			baseURL+ClusterServiceCreateProcedure,
			opts...,
		),
		delete: connect_go.NewClient[v1.ClusterServiceDeleteRequest, v1.ClusterServiceDeleteResponse](
			httpClient,
			baseURL+ClusterServiceDeleteProcedure,
			opts...,
		),
		kubeconfig: connect_go.NewClient[v1.ClusterServiceKubeconfigRequest, v1.ClusterServiceKubeconfigResponse](
			httpClient,
			baseURL+ClusterServiceKubeconfigProcedure,
			opts...,
		),
	}
}

// clusterServiceClient implements ClusterServiceClient.
type clusterServiceClient struct {
	list       *connect_go.Client[v1.ClusterServiceListRequest, v1.ClusterServiceListResponse]
	create     *connect_go.Client[v1.ClusterServiceCreateRequest, v1.ClusterServiceCreateResponse]
	delete     *connect_go.Client[v1.ClusterServiceDeleteRequest, v1.ClusterServiceDeleteResponse]
	kubeconfig *connect_go.Client[v1.ClusterServiceKubeconfigRequest, v1.ClusterServiceKubeconfigResponse]
}

// List calls ruddur.v1.ClusterService.List.
func (c *clusterServiceClient) List(ctx context.Context, req *connect_go.Request[v1.ClusterServiceListRequest]) (*connect_go.Response[v1.ClusterServiceListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// Create calls ruddur.v1.ClusterService.Create.
func (c *clusterServiceClient) Create(ctx context.Context, req *connect_go.Request[v1.ClusterServiceCreateRequest]) (*connect_go.Response[v1.ClusterServiceCreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// Delete calls ruddur.v1.ClusterService.Delete.
func (c *clusterServiceClient) Delete(ctx context.Context, req *connect_go.Request[v1.ClusterServiceDeleteRequest]) (*connect_go.Response[v1.ClusterServiceDeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// Kubeconfig calls ruddur.v1.ClusterService.Kubeconfig.
func (c *clusterServiceClient) Kubeconfig(ctx context.Context, req *connect_go.Request[v1.ClusterServiceKubeconfigRequest]) (*connect_go.Response[v1.ClusterServiceKubeconfigResponse], error) {
	return c.kubeconfig.CallUnary(ctx, req)
}

// ClusterServiceHandler is an implementation of the ruddur.v1.ClusterService service.
type ClusterServiceHandler interface {
	List(context.Context, *connect_go.Request[v1.ClusterServiceListRequest]) (*connect_go.Response[v1.ClusterServiceListResponse], error)
	Create(context.Context, *connect_go.Request[v1.ClusterServiceCreateRequest]) (*connect_go.Response[v1.ClusterServiceCreateResponse], error)
	// TODO: Edit (extend / shrink)
	Delete(context.Context, *connect_go.Request[v1.ClusterServiceDeleteRequest]) (*connect_go.Response[v1.ClusterServiceDeleteResponse], error)
	Kubeconfig(context.Context, *connect_go.Request[v1.ClusterServiceKubeconfigRequest]) (*connect_go.Response[v1.ClusterServiceKubeconfigResponse], error)
}

// NewClusterServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewClusterServiceHandler(svc ClusterServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(ClusterServiceListProcedure, connect_go.NewUnaryHandler(
		ClusterServiceListProcedure,
		svc.List,
		opts...,
	))
	mux.Handle(ClusterServiceCreateProcedure, connect_go.NewUnaryHandler(
		ClusterServiceCreateProcedure,
		svc.Create,
		opts...,
	))
	mux.Handle(ClusterServiceDeleteProcedure, connect_go.NewUnaryHandler(
		ClusterServiceDeleteProcedure,
		svc.Delete,
		opts...,
	))
	mux.Handle(ClusterServiceKubeconfigProcedure, connect_go.NewUnaryHandler(
		ClusterServiceKubeconfigProcedure,
		svc.Kubeconfig,
		opts...,
	))
	return "/ruddur.v1.ClusterService/", mux
}

// UnimplementedClusterServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedClusterServiceHandler struct{}

func (UnimplementedClusterServiceHandler) List(context.Context, *connect_go.Request[v1.ClusterServiceListRequest]) (*connect_go.Response[v1.ClusterServiceListResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ruddur.v1.ClusterService.List is not implemented"))
}

func (UnimplementedClusterServiceHandler) Create(context.Context, *connect_go.Request[v1.ClusterServiceCreateRequest]) (*connect_go.Response[v1.ClusterServiceCreateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ruddur.v1.ClusterService.Create is not implemented"))
}

func (UnimplementedClusterServiceHandler) Delete(context.Context, *connect_go.Request[v1.ClusterServiceDeleteRequest]) (*connect_go.Response[v1.ClusterServiceDeleteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ruddur.v1.ClusterService.Delete is not implemented"))
}

func (UnimplementedClusterServiceHandler) Kubeconfig(context.Context, *connect_go.Request[v1.ClusterServiceKubeconfigRequest]) (*connect_go.Response[v1.ClusterServiceKubeconfigResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ruddur.v1.ClusterService.Kubeconfig is not implemented"))
}
