// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.1
// source: pkg/fornaxcore/fornaxcore.proto

package grpc

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// FornaxCoreServiceClient is the client API for FornaxCoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FornaxCoreServiceClient interface {
	GetMessage(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (FornaxCoreService_GetMessageClient, error)
	PutMessage(ctx context.Context, in *FornaxCoreMessage, opts ...grpc.CallOption) (*empty.Empty, error)
}

type fornaxCoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFornaxCoreServiceClient(cc grpc.ClientConnInterface) FornaxCoreServiceClient {
	return &fornaxCoreServiceClient{cc}
}

func (c *fornaxCoreServiceClient) GetMessage(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (FornaxCoreService_GetMessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &FornaxCoreService_ServiceDesc.Streams[0], "/centaurusinfra.io.fornaxcore.service.FornaxCoreService/getMessage", opts...)
	if err != nil {
		return nil, err
	}
	x := &fornaxCoreServiceGetMessageClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FornaxCoreService_GetMessageClient interface {
	Recv() (*FornaxCoreMessage, error)
	grpc.ClientStream
}

type fornaxCoreServiceGetMessageClient struct {
	grpc.ClientStream
}

func (x *fornaxCoreServiceGetMessageClient) Recv() (*FornaxCoreMessage, error) {
	m := new(FornaxCoreMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *fornaxCoreServiceClient) PutMessage(ctx context.Context, in *FornaxCoreMessage, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/centaurusinfra.io.fornaxcore.service.FornaxCoreService/putMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FornaxCoreServiceServer is the server API for FornaxCoreService service.
// All implementations must embed UnimplementedFornaxCoreServiceServer
// for forward compatibility
type FornaxCoreServiceServer interface {
	GetMessage(*empty.Empty, FornaxCoreService_GetMessageServer) error
	PutMessage(context.Context, *FornaxCoreMessage) (*empty.Empty, error)
	mustEmbedUnimplementedFornaxCoreServiceServer()
}

// UnimplementedFornaxCoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFornaxCoreServiceServer struct {
}

func (UnimplementedFornaxCoreServiceServer) GetMessage(*empty.Empty, FornaxCoreService_GetMessageServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedFornaxCoreServiceServer) PutMessage(context.Context, *FornaxCoreMessage) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutMessage not implemented")
}
func (UnimplementedFornaxCoreServiceServer) mustEmbedUnimplementedFornaxCoreServiceServer() {}

// UnsafeFornaxCoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FornaxCoreServiceServer will
// result in compilation errors.
type UnsafeFornaxCoreServiceServer interface {
	mustEmbedUnimplementedFornaxCoreServiceServer()
}

func RegisterFornaxCoreServiceServer(s grpc.ServiceRegistrar, srv FornaxCoreServiceServer) {
	s.RegisterService(&FornaxCoreService_ServiceDesc, srv)
}

func _FornaxCoreService_GetMessage_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FornaxCoreServiceServer).GetMessage(m, &fornaxCoreServiceGetMessageServer{stream})
}

type FornaxCoreService_GetMessageServer interface {
	Send(*FornaxCoreMessage) error
	grpc.ServerStream
}

type fornaxCoreServiceGetMessageServer struct {
	grpc.ServerStream
}

func (x *fornaxCoreServiceGetMessageServer) Send(m *FornaxCoreMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _FornaxCoreService_PutMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FornaxCoreMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FornaxCoreServiceServer).PutMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/centaurusinfra.io.fornaxcore.service.FornaxCoreService/putMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FornaxCoreServiceServer).PutMessage(ctx, req.(*FornaxCoreMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// FornaxCoreService_ServiceDesc is the grpc.ServiceDesc for FornaxCoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FornaxCoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "centaurusinfra.io.fornaxcore.service.FornaxCoreService",
	HandlerType: (*FornaxCoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "putMessage",
			Handler:    _FornaxCoreService_PutMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "getMessage",
			Handler:       _FornaxCoreService_GetMessage_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "pkg/fornaxcore/fornaxcore.proto",
}
