// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: raystack/optimus/core/v1beta1/runtime.proto

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

// RuntimeServiceClient is the client API for RuntimeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RuntimeServiceClient interface {
	// server ping with version
	Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error)
}

type runtimeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRuntimeServiceClient(cc grpc.ClientConnInterface) RuntimeServiceClient {
	return &runtimeServiceClient{cc}
}

func (c *runtimeServiceClient) Version(ctx context.Context, in *VersionRequest, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/raystack.optimus.core.v1beta1.RuntimeService/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RuntimeServiceServer is the server API for RuntimeService service.
// All implementations must embed UnimplementedRuntimeServiceServer
// for forward compatibility
type RuntimeServiceServer interface {
	// server ping with version
	Version(context.Context, *VersionRequest) (*VersionResponse, error)
	mustEmbedUnimplementedRuntimeServiceServer()
}

// UnimplementedRuntimeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRuntimeServiceServer struct {
}

func (UnimplementedRuntimeServiceServer) Version(context.Context, *VersionRequest) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedRuntimeServiceServer) mustEmbedUnimplementedRuntimeServiceServer() {}

// UnsafeRuntimeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RuntimeServiceServer will
// result in compilation errors.
type UnsafeRuntimeServiceServer interface {
	mustEmbedUnimplementedRuntimeServiceServer()
}

func RegisterRuntimeServiceServer(s grpc.ServiceRegistrar, srv RuntimeServiceServer) {
	s.RegisterService(&RuntimeService_ServiceDesc, srv)
}

func _RuntimeService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuntimeServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/raystack.optimus.core.v1beta1.RuntimeService/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuntimeServiceServer).Version(ctx, req.(*VersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RuntimeService_ServiceDesc is the grpc.ServiceDesc for RuntimeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RuntimeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "raystack.optimus.core.v1beta1.RuntimeService",
	HandlerType: (*RuntimeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _RuntimeService_Version_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "raystack/optimus/core/v1beta1/runtime.proto",
}
