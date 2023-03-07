// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1alpha1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProjectsClient is the client API for Projects service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProjectsClient interface {
	ListProjects(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListProjectsResponse, error)
	CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	AuthorizeUser(ctx context.Context, in *AuthorizeUserForProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	RevokeUser(ctx context.Context, in *RevokeUserForProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type projectsClient struct {
	cc grpc.ClientConnInterface
}

func NewProjectsClient(cc grpc.ClientConnInterface) ProjectsClient {
	return &projectsClient{cc}
}

func (c *projectsClient) ListProjects(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ListProjectsResponse, error) {
	out := new(ListProjectsResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.Projects/ListProjects", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) CreateProject(ctx context.Context, in *CreateProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/v1alpha1.Projects/CreateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) AuthorizeUser(ctx context.Context, in *AuthorizeUserForProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/v1alpha1.Projects/AuthorizeUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectsClient) RevokeUser(ctx context.Context, in *RevokeUserForProjectRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/v1alpha1.Projects/RevokeUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProjectsServer is the server API for Projects service.
// All implementations must embed UnimplementedProjectsServer
// for forward compatibility
type ProjectsServer interface {
	ListProjects(context.Context, *emptypb.Empty) (*ListProjectsResponse, error)
	CreateProject(context.Context, *CreateProjectRequest) (*emptypb.Empty, error)
	AuthorizeUser(context.Context, *AuthorizeUserForProjectRequest) (*emptypb.Empty, error)
	RevokeUser(context.Context, *RevokeUserForProjectRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedProjectsServer()
}

// UnimplementedProjectsServer must be embedded to have forward compatible implementations.
type UnimplementedProjectsServer struct {
}

func (UnimplementedProjectsServer) ListProjects(context.Context, *emptypb.Empty) (*ListProjectsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProjects not implemented")
}
func (UnimplementedProjectsServer) CreateProject(context.Context, *CreateProjectRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedProjectsServer) AuthorizeUser(context.Context, *AuthorizeUserForProjectRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthorizeUser not implemented")
}
func (UnimplementedProjectsServer) RevokeUser(context.Context, *RevokeUserForProjectRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeUser not implemented")
}
func (UnimplementedProjectsServer) mustEmbedUnimplementedProjectsServer() {}

// UnsafeProjectsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProjectsServer will
// result in compilation errors.
type UnsafeProjectsServer interface {
	mustEmbedUnimplementedProjectsServer()
}

func RegisterProjectsServer(s grpc.ServiceRegistrar, srv ProjectsServer) {
	s.RegisterService(&Projects_ServiceDesc, srv)
}

func _Projects_ListProjects_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).ListProjects(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Projects/ListProjects",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).ListProjects(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Projects/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).CreateProject(ctx, req.(*CreateProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_AuthorizeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeUserForProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).AuthorizeUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Projects/AuthorizeUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).AuthorizeUser(ctx, req.(*AuthorizeUserForProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Projects_RevokeUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeUserForProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProjectsServer).RevokeUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Projects/RevokeUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProjectsServer).RevokeUser(ctx, req.(*RevokeUserForProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Projects_ServiceDesc is the grpc.ServiceDesc for Projects service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Projects_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha1.Projects",
	HandlerType: (*ProjectsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProjects",
			Handler:    _Projects_ListProjects_Handler,
		},
		{
			MethodName: "CreateProject",
			Handler:    _Projects_CreateProject_Handler,
		},
		{
			MethodName: "AuthorizeUser",
			Handler:    _Projects_AuthorizeUser_Handler,
		},
		{
			MethodName: "RevokeUser",
			Handler:    _Projects_RevokeUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ruddur.proto",
}

// ClustersClient is the client API for Clusters service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClustersClient interface {
	ListMachineTypes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*MachineTypesResponse, error)
	ListAvailableMachines(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AvailableMachinesResponse, error)
	ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error)
	CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*CreateClusterResponse, error)
	DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*DeleteClusterResponse, error)
	Kubeconfig(ctx context.Context, in *KubeconfigRequest, opts ...grpc.CallOption) (*KubeconfigResponse, error)
}

type clustersClient struct {
	cc grpc.ClientConnInterface
}

func NewClustersClient(cc grpc.ClientConnInterface) ClustersClient {
	return &clustersClient{cc}
}

func (c *clustersClient) ListMachineTypes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*MachineTypesResponse, error) {
	out := new(MachineTypesResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.Clusters/ListMachineTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) ListAvailableMachines(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*AvailableMachinesResponse, error) {
	out := new(AvailableMachinesResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.Clusters/ListAvailableMachines", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error) {
	out := new(ListClustersResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.Clusters/ListClusters", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) CreateCluster(ctx context.Context, in *CreateClusterRequest, opts ...grpc.CallOption) (*CreateClusterResponse, error) {
	out := new(CreateClusterResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.Clusters/CreateCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) DeleteCluster(ctx context.Context, in *DeleteClusterRequest, opts ...grpc.CallOption) (*DeleteClusterResponse, error) {
	out := new(DeleteClusterResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.Clusters/DeleteCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clustersClient) Kubeconfig(ctx context.Context, in *KubeconfigRequest, opts ...grpc.CallOption) (*KubeconfigResponse, error) {
	out := new(KubeconfigResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.Clusters/Kubeconfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClustersServer is the server API for Clusters service.
// All implementations must embed UnimplementedClustersServer
// for forward compatibility
type ClustersServer interface {
	ListMachineTypes(context.Context, *emptypb.Empty) (*MachineTypesResponse, error)
	ListAvailableMachines(context.Context, *emptypb.Empty) (*AvailableMachinesResponse, error)
	ListClusters(context.Context, *ListClustersRequest) (*ListClustersResponse, error)
	CreateCluster(context.Context, *CreateClusterRequest) (*CreateClusterResponse, error)
	DeleteCluster(context.Context, *DeleteClusterRequest) (*DeleteClusterResponse, error)
	Kubeconfig(context.Context, *KubeconfigRequest) (*KubeconfigResponse, error)
	mustEmbedUnimplementedClustersServer()
}

// UnimplementedClustersServer must be embedded to have forward compatible implementations.
type UnimplementedClustersServer struct {
}

func (UnimplementedClustersServer) ListMachineTypes(context.Context, *emptypb.Empty) (*MachineTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMachineTypes not implemented")
}
func (UnimplementedClustersServer) ListAvailableMachines(context.Context, *emptypb.Empty) (*AvailableMachinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListAvailableMachines not implemented")
}
func (UnimplementedClustersServer) ListClusters(context.Context, *ListClustersRequest) (*ListClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClusters not implemented")
}
func (UnimplementedClustersServer) CreateCluster(context.Context, *CreateClusterRequest) (*CreateClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCluster not implemented")
}
func (UnimplementedClustersServer) DeleteCluster(context.Context, *DeleteClusterRequest) (*DeleteClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCluster not implemented")
}
func (UnimplementedClustersServer) Kubeconfig(context.Context, *KubeconfigRequest) (*KubeconfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Kubeconfig not implemented")
}
func (UnimplementedClustersServer) mustEmbedUnimplementedClustersServer() {}

// UnsafeClustersServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClustersServer will
// result in compilation errors.
type UnsafeClustersServer interface {
	mustEmbedUnimplementedClustersServer()
}

func RegisterClustersServer(s grpc.ServiceRegistrar, srv ClustersServer) {
	s.RegisterService(&Clusters_ServiceDesc, srv)
}

func _Clusters_ListMachineTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).ListMachineTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Clusters/ListMachineTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).ListMachineTypes(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_ListAvailableMachines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).ListAvailableMachines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Clusters/ListAvailableMachines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).ListAvailableMachines(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_ListClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).ListClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Clusters/ListClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).ListClusters(ctx, req.(*ListClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_CreateCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).CreateCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Clusters/CreateCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).CreateCluster(ctx, req.(*CreateClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_DeleteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).DeleteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Clusters/DeleteCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).DeleteCluster(ctx, req.(*DeleteClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Clusters_Kubeconfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KubeconfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClustersServer).Kubeconfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.Clusters/Kubeconfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClustersServer).Kubeconfig(ctx, req.(*KubeconfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Clusters_ServiceDesc is the grpc.ServiceDesc for Clusters service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Clusters_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha1.Clusters",
	HandlerType: (*ClustersServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMachineTypes",
			Handler:    _Clusters_ListMachineTypes_Handler,
		},
		{
			MethodName: "ListAvailableMachines",
			Handler:    _Clusters_ListAvailableMachines_Handler,
		},
		{
			MethodName: "ListClusters",
			Handler:    _Clusters_ListClusters_Handler,
		},
		{
			MethodName: "CreateCluster",
			Handler:    _Clusters_CreateCluster_Handler,
		},
		{
			MethodName: "DeleteCluster",
			Handler:    _Clusters_DeleteCluster_Handler,
		},
		{
			MethodName: "Kubeconfig",
			Handler:    _Clusters_Kubeconfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ruddur.proto",
}

// VirtualMachinesClient is the client API for VirtualMachines service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VirtualMachinesClient interface {
	ListVirtualMachineTypes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VirtualMachineTypesResponse, error)
	ListVirtualMachineOperatingSystems(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VirtualMachineOperatingSystemsResponse, error)
	ListVirtualMachines(ctx context.Context, in *ListVirtualMachinesRequest, opts ...grpc.CallOption) (*ListVirtualMachinesResponse, error)
	CreateVirtualMachine(ctx context.Context, in *CreateVirtualMachineRequest, opts ...grpc.CallOption) (*CreateVirtualMachineResponse, error)
	DeleteVirtualMachine(ctx context.Context, in *DeleteVirtualMachineRequest, opts ...grpc.CallOption) (*DeleteVirtualMachineResponse, error)
	OpenStream(ctx context.Context, opts ...grpc.CallOption) (VirtualMachines_OpenStreamClient, error)
}

type virtualMachinesClient struct {
	cc grpc.ClientConnInterface
}

func NewVirtualMachinesClient(cc grpc.ClientConnInterface) VirtualMachinesClient {
	return &virtualMachinesClient{cc}
}

func (c *virtualMachinesClient) ListVirtualMachineTypes(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VirtualMachineTypesResponse, error) {
	out := new(VirtualMachineTypesResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.VirtualMachines/ListVirtualMachineTypes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachinesClient) ListVirtualMachineOperatingSystems(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VirtualMachineOperatingSystemsResponse, error) {
	out := new(VirtualMachineOperatingSystemsResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.VirtualMachines/ListVirtualMachineOperatingSystems", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachinesClient) ListVirtualMachines(ctx context.Context, in *ListVirtualMachinesRequest, opts ...grpc.CallOption) (*ListVirtualMachinesResponse, error) {
	out := new(ListVirtualMachinesResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.VirtualMachines/ListVirtualMachines", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachinesClient) CreateVirtualMachine(ctx context.Context, in *CreateVirtualMachineRequest, opts ...grpc.CallOption) (*CreateVirtualMachineResponse, error) {
	out := new(CreateVirtualMachineResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.VirtualMachines/CreateVirtualMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachinesClient) DeleteVirtualMachine(ctx context.Context, in *DeleteVirtualMachineRequest, opts ...grpc.CallOption) (*DeleteVirtualMachineResponse, error) {
	out := new(DeleteVirtualMachineResponse)
	err := c.cc.Invoke(ctx, "/v1alpha1.VirtualMachines/DeleteVirtualMachine", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *virtualMachinesClient) OpenStream(ctx context.Context, opts ...grpc.CallOption) (VirtualMachines_OpenStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &VirtualMachines_ServiceDesc.Streams[0], "/v1alpha1.VirtualMachines/OpenStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &virtualMachinesOpenStreamClient{stream}
	return x, nil
}

type VirtualMachines_OpenStreamClient interface {
	Send(*Chunk) error
	Recv() (*Chunk, error)
	grpc.ClientStream
}

type virtualMachinesOpenStreamClient struct {
	grpc.ClientStream
}

func (x *virtualMachinesOpenStreamClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *virtualMachinesOpenStreamClient) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VirtualMachinesServer is the server API for VirtualMachines service.
// All implementations must embed UnimplementedVirtualMachinesServer
// for forward compatibility
type VirtualMachinesServer interface {
	ListVirtualMachineTypes(context.Context, *emptypb.Empty) (*VirtualMachineTypesResponse, error)
	ListVirtualMachineOperatingSystems(context.Context, *emptypb.Empty) (*VirtualMachineOperatingSystemsResponse, error)
	ListVirtualMachines(context.Context, *ListVirtualMachinesRequest) (*ListVirtualMachinesResponse, error)
	CreateVirtualMachine(context.Context, *CreateVirtualMachineRequest) (*CreateVirtualMachineResponse, error)
	DeleteVirtualMachine(context.Context, *DeleteVirtualMachineRequest) (*DeleteVirtualMachineResponse, error)
	OpenStream(VirtualMachines_OpenStreamServer) error
	mustEmbedUnimplementedVirtualMachinesServer()
}

// UnimplementedVirtualMachinesServer must be embedded to have forward compatible implementations.
type UnimplementedVirtualMachinesServer struct {
}

func (UnimplementedVirtualMachinesServer) ListVirtualMachineTypes(context.Context, *emptypb.Empty) (*VirtualMachineTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVirtualMachineTypes not implemented")
}
func (UnimplementedVirtualMachinesServer) ListVirtualMachineOperatingSystems(context.Context, *emptypb.Empty) (*VirtualMachineOperatingSystemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVirtualMachineOperatingSystems not implemented")
}
func (UnimplementedVirtualMachinesServer) ListVirtualMachines(context.Context, *ListVirtualMachinesRequest) (*ListVirtualMachinesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVirtualMachines not implemented")
}
func (UnimplementedVirtualMachinesServer) CreateVirtualMachine(context.Context, *CreateVirtualMachineRequest) (*CreateVirtualMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVirtualMachine not implemented")
}
func (UnimplementedVirtualMachinesServer) DeleteVirtualMachine(context.Context, *DeleteVirtualMachineRequest) (*DeleteVirtualMachineResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVirtualMachine not implemented")
}
func (UnimplementedVirtualMachinesServer) OpenStream(VirtualMachines_OpenStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method OpenStream not implemented")
}
func (UnimplementedVirtualMachinesServer) mustEmbedUnimplementedVirtualMachinesServer() {}

// UnsafeVirtualMachinesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VirtualMachinesServer will
// result in compilation errors.
type UnsafeVirtualMachinesServer interface {
	mustEmbedUnimplementedVirtualMachinesServer()
}

func RegisterVirtualMachinesServer(s grpc.ServiceRegistrar, srv VirtualMachinesServer) {
	s.RegisterService(&VirtualMachines_ServiceDesc, srv)
}

func _VirtualMachines_ListVirtualMachineTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachinesServer).ListVirtualMachineTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.VirtualMachines/ListVirtualMachineTypes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachinesServer).ListVirtualMachineTypes(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachines_ListVirtualMachineOperatingSystems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachinesServer).ListVirtualMachineOperatingSystems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.VirtualMachines/ListVirtualMachineOperatingSystems",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachinesServer).ListVirtualMachineOperatingSystems(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachines_ListVirtualMachines_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVirtualMachinesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachinesServer).ListVirtualMachines(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.VirtualMachines/ListVirtualMachines",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachinesServer).ListVirtualMachines(ctx, req.(*ListVirtualMachinesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachines_CreateVirtualMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVirtualMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachinesServer).CreateVirtualMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.VirtualMachines/CreateVirtualMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachinesServer).CreateVirtualMachine(ctx, req.(*CreateVirtualMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachines_DeleteVirtualMachine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVirtualMachineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VirtualMachinesServer).DeleteVirtualMachine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v1alpha1.VirtualMachines/DeleteVirtualMachine",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VirtualMachinesServer).DeleteVirtualMachine(ctx, req.(*DeleteVirtualMachineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VirtualMachines_OpenStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VirtualMachinesServer).OpenStream(&virtualMachinesOpenStreamServer{stream})
}

type VirtualMachines_OpenStreamServer interface {
	Send(*Chunk) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type virtualMachinesOpenStreamServer struct {
	grpc.ServerStream
}

func (x *virtualMachinesOpenStreamServer) Send(m *Chunk) error {
	return x.ServerStream.SendMsg(m)
}

func (x *virtualMachinesOpenStreamServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VirtualMachines_ServiceDesc is the grpc.ServiceDesc for VirtualMachines service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VirtualMachines_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1alpha1.VirtualMachines",
	HandlerType: (*VirtualMachinesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListVirtualMachineTypes",
			Handler:    _VirtualMachines_ListVirtualMachineTypes_Handler,
		},
		{
			MethodName: "ListVirtualMachineOperatingSystems",
			Handler:    _VirtualMachines_ListVirtualMachineOperatingSystems_Handler,
		},
		{
			MethodName: "ListVirtualMachines",
			Handler:    _VirtualMachines_ListVirtualMachines_Handler,
		},
		{
			MethodName: "CreateVirtualMachine",
			Handler:    _VirtualMachines_CreateVirtualMachine_Handler,
		},
		{
			MethodName: "DeleteVirtualMachine",
			Handler:    _VirtualMachines_DeleteVirtualMachine_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OpenStream",
			Handler:       _VirtualMachines_OpenStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ruddur.proto",
}
