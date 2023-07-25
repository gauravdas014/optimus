// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: raystack/optimus/core/v1beta1/namespace.proto

package optimus

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// NamespaceServiceClient is the client API for NamespaceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NamespaceServiceClient interface {
	// RegisterProjectNamespace creates a new namespace for a project
	RegisterProjectNamespace(ctx context.Context, in *RegisterProjectNamespaceRequest, opts ...grpc.CallOption) (*RegisterProjectNamespaceResponse, error)
	// ListProjectNamespaces returns list of namespaces of a project
	ListProjectNamespaces(ctx context.Context, in *ListProjectNamespacesRequest, opts ...grpc.CallOption) (*ListProjectNamespacesResponse, error)
	// GetNamespace returns namespace details based on project_name and namespace_name
	GetNamespace(ctx context.Context, in *GetNamespaceRequest, opts ...grpc.CallOption) (*GetNamespaceResponse, error)
}

type namespaceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNamespaceServiceClient(cc grpc.ClientConnInterface) NamespaceServiceClient {
	return &namespaceServiceClient{cc}
}

func (c *namespaceServiceClient) RegisterProjectNamespace(ctx context.Context, in *RegisterProjectNamespaceRequest, opts ...grpc.CallOption) (*RegisterProjectNamespaceResponse, error) {
	out := new(RegisterProjectNamespaceResponse)
	err := c.cc.Invoke(ctx, "/raystack.optimus.core.v1beta1.NamespaceService/RegisterProjectNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceServiceClient) ListProjectNamespaces(ctx context.Context, in *ListProjectNamespacesRequest, opts ...grpc.CallOption) (*ListProjectNamespacesResponse, error) {
	out := new(ListProjectNamespacesResponse)
	err := c.cc.Invoke(ctx, "/raystack.optimus.core.v1beta1.NamespaceService/ListProjectNamespaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *namespaceServiceClient) GetNamespace(ctx context.Context, in *GetNamespaceRequest, opts ...grpc.CallOption) (*GetNamespaceResponse, error) {
	out := new(GetNamespaceResponse)
	err := c.cc.Invoke(ctx, "/raystack.optimus.core.v1beta1.NamespaceService/GetNamespace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NamespaceServiceServer is the server API for NamespaceService service.
// All implementations must embed UnimplementedNamespaceServiceServer
// for forward compatibility
type NamespaceServiceServer interface {
	// RegisterProjectNamespace creates a new namespace for a project
	RegisterProjectNamespace(context.Context, *RegisterProjectNamespaceRequest) (*RegisterProjectNamespaceResponse, error)
	// ListProjectNamespaces returns list of namespaces of a project
	ListProjectNamespaces(context.Context, *ListProjectNamespacesRequest) (*ListProjectNamespacesResponse, error)
	// GetNamespace returns namespace details based on project_name and namespace_name
	GetNamespace(context.Context, *GetNamespaceRequest) (*GetNamespaceResponse, error)
	mustEmbedUnimplementedNamespaceServiceServer()
}

// UnimplementedNamespaceServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNamespaceServiceServer struct {
}

func (UnimplementedNamespaceServiceServer) RegisterProjectNamespace(context.Context, *RegisterProjectNamespaceRequest) (*RegisterProjectNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterProjectNamespace not implemented")
}
func (UnimplementedNamespaceServiceServer) ListProjectNamespaces(context.Context, *ListProjectNamespacesRequest) (*ListProjectNamespacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProjectNamespaces not implemented")
}
func (UnimplementedNamespaceServiceServer) GetNamespace(context.Context, *GetNamespaceRequest) (*GetNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNamespace not implemented")
}
func (UnimplementedNamespaceServiceServer) mustEmbedUnimplementedNamespaceServiceServer() {}

// UnsafeNamespaceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NamespaceServiceServer will
// result in compilation errors.
type UnsafeNamespaceServiceServer interface {
	mustEmbedUnimplementedNamespaceServiceServer()
}

func RegisterNamespaceServiceServer(s grpc.ServiceRegistrar, srv NamespaceServiceServer) {
	s.RegisterService(&NamespaceService_ServiceDesc, srv)
}

func _NamespaceService_RegisterProjectNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterProjectNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).RegisterProjectNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raystack.optimus.core.v1beta1.NamespaceService/RegisterProjectNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).RegisterProjectNamespace(ctx, req.(*RegisterProjectNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceService_ListProjectNamespaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProjectNamespacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).ListProjectNamespaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raystack.optimus.core.v1beta1.NamespaceService/ListProjectNamespaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).ListProjectNamespaces(ctx, req.(*ListProjectNamespacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NamespaceService_GetNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NamespaceServiceServer).GetNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raystack.optimus.core.v1beta1.NamespaceService/GetNamespace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NamespaceServiceServer).GetNamespace(ctx, req.(*GetNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NamespaceService_ServiceDesc is the grpc.ServiceDesc for NamespaceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NamespaceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "raystack.optimus.core.v1beta1.NamespaceService",
	HandlerType: (*NamespaceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterProjectNamespace",
			Handler:    _NamespaceService_RegisterProjectNamespace_Handler,
		},
		{
			MethodName: "ListProjectNamespaces",
			Handler:    _NamespaceService_ListProjectNamespaces_Handler,
		},
		{
			MethodName: "GetNamespace",
			Handler:    _NamespaceService_GetNamespace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "raystack/optimus/core/v1beta1/namespace.proto",
}
