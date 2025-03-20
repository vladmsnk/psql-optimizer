// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: optimizer/optimizer.proto

package pb

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

const (
	OptimizerAPI_ListIndexesStat_FullMethodName = "/api.optimizer.v1.OptimizerAPI/ListIndexesStat"
)

// OptimizerAPIClient is the client API for OptimizerAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OptimizerAPIClient interface {
	ListIndexesStat(ctx context.Context, in *ListIndexesStatRequest, opts ...grpc.CallOption) (*ListIndexesStatResponse, error)
}

type optimizerAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewOptimizerAPIClient(cc grpc.ClientConnInterface) OptimizerAPIClient {
	return &optimizerAPIClient{cc}
}

func (c *optimizerAPIClient) ListIndexesStat(ctx context.Context, in *ListIndexesStatRequest, opts ...grpc.CallOption) (*ListIndexesStatResponse, error) {
	out := new(ListIndexesStatResponse)
	err := c.cc.Invoke(ctx, OptimizerAPI_ListIndexesStat_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OptimizerAPIServer is the server API for OptimizerAPI service.
// All implementations must embed UnimplementedOptimizerAPIServer
// for forward compatibility
type OptimizerAPIServer interface {
	ListIndexesStat(context.Context, *ListIndexesStatRequest) (*ListIndexesStatResponse, error)
	mustEmbedUnimplementedOptimizerAPIServer()
}

// UnimplementedOptimizerAPIServer must be embedded to have forward compatible implementations.
type UnimplementedOptimizerAPIServer struct {
}

func (UnimplementedOptimizerAPIServer) ListIndexesStat(context.Context, *ListIndexesStatRequest) (*ListIndexesStatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIndexesStat not implemented")
}
func (UnimplementedOptimizerAPIServer) mustEmbedUnimplementedOptimizerAPIServer() {}

// UnsafeOptimizerAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OptimizerAPIServer will
// result in compilation errors.
type UnsafeOptimizerAPIServer interface {
	mustEmbedUnimplementedOptimizerAPIServer()
}

func RegisterOptimizerAPIServer(s grpc.ServiceRegistrar, srv OptimizerAPIServer) {
	s.RegisterService(&OptimizerAPI_ServiceDesc, srv)
}

func _OptimizerAPI_ListIndexesStat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIndexesStatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OptimizerAPIServer).ListIndexesStat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OptimizerAPI_ListIndexesStat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OptimizerAPIServer).ListIndexesStat(ctx, req.(*ListIndexesStatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OptimizerAPI_ServiceDesc is the grpc.ServiceDesc for OptimizerAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OptimizerAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.optimizer.v1.OptimizerAPI",
	HandlerType: (*OptimizerAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListIndexesStat",
			Handler:    _OptimizerAPI_ListIndexesStat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "optimizer/optimizer.proto",
}
