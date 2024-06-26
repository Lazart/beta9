// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: container.proto

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

const (
	ContainerService_ExecuteCommand_FullMethodName = "/container.ContainerService/ExecuteCommand"
)

// ContainerServiceClient is the client API for ContainerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContainerServiceClient interface {
	ExecuteCommand(ctx context.Context, in *CommandExecutionRequest, opts ...grpc.CallOption) (ContainerService_ExecuteCommandClient, error)
}

type containerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContainerServiceClient(cc grpc.ClientConnInterface) ContainerServiceClient {
	return &containerServiceClient{cc}
}

func (c *containerServiceClient) ExecuteCommand(ctx context.Context, in *CommandExecutionRequest, opts ...grpc.CallOption) (ContainerService_ExecuteCommandClient, error) {
	stream, err := c.cc.NewStream(ctx, &ContainerService_ServiceDesc.Streams[0], ContainerService_ExecuteCommand_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &containerServiceExecuteCommandClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ContainerService_ExecuteCommandClient interface {
	Recv() (*CommandExecutionResponse, error)
	grpc.ClientStream
}

type containerServiceExecuteCommandClient struct {
	grpc.ClientStream
}

func (x *containerServiceExecuteCommandClient) Recv() (*CommandExecutionResponse, error) {
	m := new(CommandExecutionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ContainerServiceServer is the server API for ContainerService service.
// All implementations must embed UnimplementedContainerServiceServer
// for forward compatibility
type ContainerServiceServer interface {
	ExecuteCommand(*CommandExecutionRequest, ContainerService_ExecuteCommandServer) error
	mustEmbedUnimplementedContainerServiceServer()
}

// UnimplementedContainerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContainerServiceServer struct {
}

func (UnimplementedContainerServiceServer) ExecuteCommand(*CommandExecutionRequest, ContainerService_ExecuteCommandServer) error {
	return status.Errorf(codes.Unimplemented, "method ExecuteCommand not implemented")
}
func (UnimplementedContainerServiceServer) mustEmbedUnimplementedContainerServiceServer() {}

// UnsafeContainerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContainerServiceServer will
// result in compilation errors.
type UnsafeContainerServiceServer interface {
	mustEmbedUnimplementedContainerServiceServer()
}

func RegisterContainerServiceServer(s grpc.ServiceRegistrar, srv ContainerServiceServer) {
	s.RegisterService(&ContainerService_ServiceDesc, srv)
}

func _ContainerService_ExecuteCommand_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CommandExecutionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ContainerServiceServer).ExecuteCommand(m, &containerServiceExecuteCommandServer{stream})
}

type ContainerService_ExecuteCommandServer interface {
	Send(*CommandExecutionResponse) error
	grpc.ServerStream
}

type containerServiceExecuteCommandServer struct {
	grpc.ServerStream
}

func (x *containerServiceExecuteCommandServer) Send(m *CommandExecutionResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ContainerService_ServiceDesc is the grpc.ServiceDesc for ContainerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ContainerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "container.ContainerService",
	HandlerType: (*ContainerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ExecuteCommand",
			Handler:       _ContainerService_ExecuteCommand_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "container.proto",
}
