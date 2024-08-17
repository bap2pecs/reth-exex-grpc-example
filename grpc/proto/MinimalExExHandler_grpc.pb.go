// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.0
// source: MinimalExExHandler.proto

package proto

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

// ExExHandlerClient is the client API for ExExHandler service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExExHandlerClient interface {
	HandleChainCommitted(ctx context.Context, in *ChainCommittedRequest, opts ...grpc.CallOption) (*HandlerResponse, error)
	HandleChainReorged(ctx context.Context, in *ChainReorgedRequest, opts ...grpc.CallOption) (*HandlerResponse, error)
	HandleChainReverted(ctx context.Context, in *ChainRevertedRequest, opts ...grpc.CallOption) (*HandlerResponse, error)
}

type exExHandlerClient struct {
	cc grpc.ClientConnInterface
}

func NewExExHandlerClient(cc grpc.ClientConnInterface) ExExHandlerClient {
	return &exExHandlerClient{cc}
}

func (c *exExHandlerClient) HandleChainCommitted(ctx context.Context, in *ChainCommittedRequest, opts ...grpc.CallOption) (*HandlerResponse, error) {
	out := new(HandlerResponse)
	err := c.cc.Invoke(ctx, "/exex.ExExHandler/HandleChainCommitted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exExHandlerClient) HandleChainReorged(ctx context.Context, in *ChainReorgedRequest, opts ...grpc.CallOption) (*HandlerResponse, error) {
	out := new(HandlerResponse)
	err := c.cc.Invoke(ctx, "/exex.ExExHandler/HandleChainReorged", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exExHandlerClient) HandleChainReverted(ctx context.Context, in *ChainRevertedRequest, opts ...grpc.CallOption) (*HandlerResponse, error) {
	out := new(HandlerResponse)
	err := c.cc.Invoke(ctx, "/exex.ExExHandler/HandleChainReverted", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExExHandlerServer is the server API for ExExHandler service.
// All implementations must embed UnimplementedExExHandlerServer
// for forward compatibility
type ExExHandlerServer interface {
	HandleChainCommitted(context.Context, *ChainCommittedRequest) (*HandlerResponse, error)
	HandleChainReorged(context.Context, *ChainReorgedRequest) (*HandlerResponse, error)
	HandleChainReverted(context.Context, *ChainRevertedRequest) (*HandlerResponse, error)
	mustEmbedUnimplementedExExHandlerServer()
}

// UnimplementedExExHandlerServer must be embedded to have forward compatible implementations.
type UnimplementedExExHandlerServer struct {
}

func (UnimplementedExExHandlerServer) HandleChainCommitted(context.Context, *ChainCommittedRequest) (*HandlerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleChainCommitted not implemented")
}
func (UnimplementedExExHandlerServer) HandleChainReorged(context.Context, *ChainReorgedRequest) (*HandlerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleChainReorged not implemented")
}
func (UnimplementedExExHandlerServer) HandleChainReverted(context.Context, *ChainRevertedRequest) (*HandlerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HandleChainReverted not implemented")
}
func (UnimplementedExExHandlerServer) mustEmbedUnimplementedExExHandlerServer() {}

// UnsafeExExHandlerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExExHandlerServer will
// result in compilation errors.
type UnsafeExExHandlerServer interface {
	mustEmbedUnimplementedExExHandlerServer()
}

func RegisterExExHandlerServer(s grpc.ServiceRegistrar, srv ExExHandlerServer) {
	s.RegisterService(&ExExHandler_ServiceDesc, srv)
}

func _ExExHandler_HandleChainCommitted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChainCommittedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExExHandlerServer).HandleChainCommitted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exex.ExExHandler/HandleChainCommitted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExExHandlerServer).HandleChainCommitted(ctx, req.(*ChainCommittedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExExHandler_HandleChainReorged_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChainReorgedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExExHandlerServer).HandleChainReorged(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exex.ExExHandler/HandleChainReorged",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExExHandlerServer).HandleChainReorged(ctx, req.(*ChainReorgedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExExHandler_HandleChainReverted_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChainRevertedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExExHandlerServer).HandleChainReverted(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/exex.ExExHandler/HandleChainReverted",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExExHandlerServer).HandleChainReverted(ctx, req.(*ChainRevertedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExExHandler_ServiceDesc is the grpc.ServiceDesc for ExExHandler service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExExHandler_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "exex.ExExHandler",
	HandlerType: (*ExExHandlerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleChainCommitted",
			Handler:    _ExExHandler_HandleChainCommitted_Handler,
		},
		{
			MethodName: "HandleChainReorged",
			Handler:    _ExExHandler_HandleChainReorged_Handler,
		},
		{
			MethodName: "HandleChainReverted",
			Handler:    _ExExHandler_HandleChainReverted_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "MinimalExExHandler.proto",
}
