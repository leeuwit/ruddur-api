// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: ruddur/v1/project.proto

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
	// ProjectServiceName is the fully-qualified name of the ProjectService service.
	ProjectServiceName = "ruddur.v1.ProjectService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ProjectServiceListProcedure is the fully-qualified name of the ProjectService's List RPC.
	ProjectServiceListProcedure = "/ruddur.v1.ProjectService/List"
	// ProjectServiceGetProcedure is the fully-qualified name of the ProjectService's Get RPC.
	ProjectServiceGetProcedure = "/ruddur.v1.ProjectService/Get"
	// ProjectServiceRegisterProcedure is the fully-qualified name of the ProjectService's Register RPC.
	ProjectServiceRegisterProcedure = "/ruddur.v1.ProjectService/Register"
	// ProjectServiceCreateProcedure is the fully-qualified name of the ProjectService's Create RPC.
	ProjectServiceCreateProcedure = "/ruddur.v1.ProjectService/Create"
	// ProjectServiceDeleteProcedure is the fully-qualified name of the ProjectService's Delete RPC.
	ProjectServiceDeleteProcedure = "/ruddur.v1.ProjectService/Delete"
	// ProjectServiceAuthorizeProcedure is the fully-qualified name of the ProjectService's Authorize
	// RPC.
	ProjectServiceAuthorizeProcedure = "/ruddur.v1.ProjectService/Authorize"
	// ProjectServiceRevokeProcedure is the fully-qualified name of the ProjectService's Revoke RPC.
	ProjectServiceRevokeProcedure = "/ruddur.v1.ProjectService/Revoke"
)

// ProjectServiceClient is a client for the ruddur.v1.ProjectService service.
type ProjectServiceClient interface {
	List(context.Context, *connect_go.Request[v1.ProjectServiceListRequest]) (*connect_go.Response[v1.ProjectServiceListResponse], error)
	Get(context.Context, *connect_go.Request[v1.ProjectServiceGetRequest]) (*connect_go.Response[v1.ProjectServiceGetResponse], error)
	Register(context.Context, *connect_go.Request[v1.ProjectServiceRegisterRequest]) (*connect_go.Response[v1.ProjectServiceRegisterResponse], error)
	Create(context.Context, *connect_go.Request[v1.ProjectServiceCreateRequest]) (*connect_go.Response[v1.ProjectServiceCreateResponse], error)
	Delete(context.Context, *connect_go.Request[v1.ProjectServiceDeleteRequest]) (*connect_go.Response[v1.ProjectServiceDeleteResponse], error)
	Authorize(context.Context, *connect_go.Request[v1.ProjectServiceAuthorizeRequest]) (*connect_go.Response[v1.ProjectServiceAuthorizeResponse], error)
	Revoke(context.Context, *connect_go.Request[v1.ProjectServiceRevokeRequest]) (*connect_go.Response[v1.ProjectServiceRevokeResponse], error)
}

// NewProjectServiceClient constructs a client for the ruddur.v1.ProjectService service. By default,
// it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and
// sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC()
// or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewProjectServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ProjectServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &projectServiceClient{
		list: connect_go.NewClient[v1.ProjectServiceListRequest, v1.ProjectServiceListResponse](
			httpClient,
			baseURL+ProjectServiceListProcedure,
			opts...,
		),
		get: connect_go.NewClient[v1.ProjectServiceGetRequest, v1.ProjectServiceGetResponse](
			httpClient,
			baseURL+ProjectServiceGetProcedure,
			opts...,
		),
		register: connect_go.NewClient[v1.ProjectServiceRegisterRequest, v1.ProjectServiceRegisterResponse](
			httpClient,
			baseURL+ProjectServiceRegisterProcedure,
			opts...,
		),
		create: connect_go.NewClient[v1.ProjectServiceCreateRequest, v1.ProjectServiceCreateResponse](
			httpClient,
			baseURL+ProjectServiceCreateProcedure,
			opts...,
		),
		delete: connect_go.NewClient[v1.ProjectServiceDeleteRequest, v1.ProjectServiceDeleteResponse](
			httpClient,
			baseURL+ProjectServiceDeleteProcedure,
			opts...,
		),
		authorize: connect_go.NewClient[v1.ProjectServiceAuthorizeRequest, v1.ProjectServiceAuthorizeResponse](
			httpClient,
			baseURL+ProjectServiceAuthorizeProcedure,
			opts...,
		),
		revoke: connect_go.NewClient[v1.ProjectServiceRevokeRequest, v1.ProjectServiceRevokeResponse](
			httpClient,
			baseURL+ProjectServiceRevokeProcedure,
			opts...,
		),
	}
}

// projectServiceClient implements ProjectServiceClient.
type projectServiceClient struct {
	list      *connect_go.Client[v1.ProjectServiceListRequest, v1.ProjectServiceListResponse]
	get       *connect_go.Client[v1.ProjectServiceGetRequest, v1.ProjectServiceGetResponse]
	register  *connect_go.Client[v1.ProjectServiceRegisterRequest, v1.ProjectServiceRegisterResponse]
	create    *connect_go.Client[v1.ProjectServiceCreateRequest, v1.ProjectServiceCreateResponse]
	delete    *connect_go.Client[v1.ProjectServiceDeleteRequest, v1.ProjectServiceDeleteResponse]
	authorize *connect_go.Client[v1.ProjectServiceAuthorizeRequest, v1.ProjectServiceAuthorizeResponse]
	revoke    *connect_go.Client[v1.ProjectServiceRevokeRequest, v1.ProjectServiceRevokeResponse]
}

// List calls ruddur.v1.ProjectService.List.
func (c *projectServiceClient) List(ctx context.Context, req *connect_go.Request[v1.ProjectServiceListRequest]) (*connect_go.Response[v1.ProjectServiceListResponse], error) {
	return c.list.CallUnary(ctx, req)
}

// Get calls ruddur.v1.ProjectService.Get.
func (c *projectServiceClient) Get(ctx context.Context, req *connect_go.Request[v1.ProjectServiceGetRequest]) (*connect_go.Response[v1.ProjectServiceGetResponse], error) {
	return c.get.CallUnary(ctx, req)
}

// Register calls ruddur.v1.ProjectService.Register.
func (c *projectServiceClient) Register(ctx context.Context, req *connect_go.Request[v1.ProjectServiceRegisterRequest]) (*connect_go.Response[v1.ProjectServiceRegisterResponse], error) {
	return c.register.CallUnary(ctx, req)
}

// Create calls ruddur.v1.ProjectService.Create.
func (c *projectServiceClient) Create(ctx context.Context, req *connect_go.Request[v1.ProjectServiceCreateRequest]) (*connect_go.Response[v1.ProjectServiceCreateResponse], error) {
	return c.create.CallUnary(ctx, req)
}

// Delete calls ruddur.v1.ProjectService.Delete.
func (c *projectServiceClient) Delete(ctx context.Context, req *connect_go.Request[v1.ProjectServiceDeleteRequest]) (*connect_go.Response[v1.ProjectServiceDeleteResponse], error) {
	return c.delete.CallUnary(ctx, req)
}

// Authorize calls ruddur.v1.ProjectService.Authorize.
func (c *projectServiceClient) Authorize(ctx context.Context, req *connect_go.Request[v1.ProjectServiceAuthorizeRequest]) (*connect_go.Response[v1.ProjectServiceAuthorizeResponse], error) {
	return c.authorize.CallUnary(ctx, req)
}

// Revoke calls ruddur.v1.ProjectService.Revoke.
func (c *projectServiceClient) Revoke(ctx context.Context, req *connect_go.Request[v1.ProjectServiceRevokeRequest]) (*connect_go.Response[v1.ProjectServiceRevokeResponse], error) {
	return c.revoke.CallUnary(ctx, req)
}

// ProjectServiceHandler is an implementation of the ruddur.v1.ProjectService service.
type ProjectServiceHandler interface {
	List(context.Context, *connect_go.Request[v1.ProjectServiceListRequest]) (*connect_go.Response[v1.ProjectServiceListResponse], error)
	Get(context.Context, *connect_go.Request[v1.ProjectServiceGetRequest]) (*connect_go.Response[v1.ProjectServiceGetResponse], error)
	Register(context.Context, *connect_go.Request[v1.ProjectServiceRegisterRequest]) (*connect_go.Response[v1.ProjectServiceRegisterResponse], error)
	Create(context.Context, *connect_go.Request[v1.ProjectServiceCreateRequest]) (*connect_go.Response[v1.ProjectServiceCreateResponse], error)
	Delete(context.Context, *connect_go.Request[v1.ProjectServiceDeleteRequest]) (*connect_go.Response[v1.ProjectServiceDeleteResponse], error)
	Authorize(context.Context, *connect_go.Request[v1.ProjectServiceAuthorizeRequest]) (*connect_go.Response[v1.ProjectServiceAuthorizeResponse], error)
	Revoke(context.Context, *connect_go.Request[v1.ProjectServiceRevokeRequest]) (*connect_go.Response[v1.ProjectServiceRevokeResponse], error)
}

// NewProjectServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewProjectServiceHandler(svc ProjectServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(ProjectServiceListProcedure, connect_go.NewUnaryHandler(
		ProjectServiceListProcedure,
		svc.List,
		opts...,
	))
	mux.Handle(ProjectServiceGetProcedure, connect_go.NewUnaryHandler(
		ProjectServiceGetProcedure,
		svc.Get,
		opts...,
	))
	mux.Handle(ProjectServiceRegisterProcedure, connect_go.NewUnaryHandler(
		ProjectServiceRegisterProcedure,
		svc.Register,
		opts...,
	))
	mux.Handle(ProjectServiceCreateProcedure, connect_go.NewUnaryHandler(
		ProjectServiceCreateProcedure,
		svc.Create,
		opts...,
	))
	mux.Handle(ProjectServiceDeleteProcedure, connect_go.NewUnaryHandler(
		ProjectServiceDeleteProcedure,
		svc.Delete,
		opts...,
	))
	mux.Handle(ProjectServiceAuthorizeProcedure, connect_go.NewUnaryHandler(
		ProjectServiceAuthorizeProcedure,
		svc.Authorize,
		opts...,
	))
	mux.Handle(ProjectServiceRevokeProcedure, connect_go.NewUnaryHandler(
		ProjectServiceRevokeProcedure,
		svc.Revoke,
		opts...,
	))
	return "/ruddur.v1.ProjectService/", mux
}

// UnimplementedProjectServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedProjectServiceHandler struct{}

func (UnimplementedProjectServiceHandler) List(context.Context, *connect_go.Request[v1.ProjectServiceListRequest]) (*connect_go.Response[v1.ProjectServiceListResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ruddur.v1.ProjectService.List is not implemented"))
}

func (UnimplementedProjectServiceHandler) Get(context.Context, *connect_go.Request[v1.ProjectServiceGetRequest]) (*connect_go.Response[v1.ProjectServiceGetResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ruddur.v1.ProjectService.Get is not implemented"))
}

func (UnimplementedProjectServiceHandler) Register(context.Context, *connect_go.Request[v1.ProjectServiceRegisterRequest]) (*connect_go.Response[v1.ProjectServiceRegisterResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ruddur.v1.ProjectService.Register is not implemented"))
}

func (UnimplementedProjectServiceHandler) Create(context.Context, *connect_go.Request[v1.ProjectServiceCreateRequest]) (*connect_go.Response[v1.ProjectServiceCreateResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ruddur.v1.ProjectService.Create is not implemented"))
}

func (UnimplementedProjectServiceHandler) Delete(context.Context, *connect_go.Request[v1.ProjectServiceDeleteRequest]) (*connect_go.Response[v1.ProjectServiceDeleteResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ruddur.v1.ProjectService.Delete is not implemented"))
}

func (UnimplementedProjectServiceHandler) Authorize(context.Context, *connect_go.Request[v1.ProjectServiceAuthorizeRequest]) (*connect_go.Response[v1.ProjectServiceAuthorizeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ruddur.v1.ProjectService.Authorize is not implemented"))
}

func (UnimplementedProjectServiceHandler) Revoke(context.Context, *connect_go.Request[v1.ProjectServiceRevokeRequest]) (*connect_go.Response[v1.ProjectServiceRevokeResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("ruddur.v1.ProjectService.Revoke is not implemented"))
}