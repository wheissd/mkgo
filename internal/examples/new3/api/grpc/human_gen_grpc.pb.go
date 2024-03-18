// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: human_gen.proto

package grpc

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
	HumanService_ReadHuman_FullMethodName = "/grpc.HumanService/ReadHuman"
)

// HumanServiceClient is the client API for HumanService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HumanServiceClient interface {
	ReadHuman(ctx context.Context, in *ReadHumanRequest, opts ...grpc.CallOption) (*Human, error)
}

type humanServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHumanServiceClient(cc grpc.ClientConnInterface) HumanServiceClient {
	return &humanServiceClient{cc}
}

func (c *humanServiceClient) ReadHuman(ctx context.Context, in *ReadHumanRequest, opts ...grpc.CallOption) (*Human, error) {
	out := new(Human)
	err := c.cc.Invoke(ctx, HumanService_ReadHuman_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HumanServiceServer is the server API for HumanService service.
// All implementations must embed UnimplementedHumanServiceServer
// for forward compatibility
type HumanServiceServer interface {
	ReadHuman(context.Context, *ReadHumanRequest) (*Human, error)
	mustEmbedUnimplementedHumanServiceServer()
}

// UnimplementedHumanServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHumanServiceServer struct {
}

func (UnimplementedHumanServiceServer) ReadHuman(context.Context, *ReadHumanRequest) (*Human, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadHuman not implemented")
}
func (UnimplementedHumanServiceServer) mustEmbedUnimplementedHumanServiceServer() {}

// UnsafeHumanServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HumanServiceServer will
// result in compilation errors.
type UnsafeHumanServiceServer interface {
	mustEmbedUnimplementedHumanServiceServer()
}

func RegisterHumanServiceServer(s grpc.ServiceRegistrar, srv HumanServiceServer) {
	s.RegisterService(&HumanService_ServiceDesc, srv)
}

func _HumanService_ReadHuman_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadHumanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HumanServiceServer).ReadHuman(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HumanService_ReadHuman_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HumanServiceServer).ReadHuman(ctx, req.(*ReadHumanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HumanService_ServiceDesc is the grpc.ServiceDesc for HumanService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HumanService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.HumanService",
	HandlerType: (*HumanServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ReadHuman",
			Handler:    _HumanService_ReadHuman_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "human_gen.proto",
}
