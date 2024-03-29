// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: service_sharecycle.proto

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
	Sharecycle_CreateUser_FullMethodName = "/pb.Sharecycle/CreateUser"
	Sharecycle_GetUser_FullMethodName    = "/pb.Sharecycle/GetUser"
)

// SharecycleClient is the client API for Sharecycle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SharecycleClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
}

type sharecycleClient struct {
	cc grpc.ClientConnInterface
}

func NewSharecycleClient(cc grpc.ClientConnInterface) SharecycleClient {
	return &sharecycleClient{cc}
}

func (c *sharecycleClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, Sharecycle_CreateUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sharecycleClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, Sharecycle_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SharecycleServer is the server API for Sharecycle service.
// All implementations must embed UnimplementedSharecycleServer
// for forward compatibility
type SharecycleServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	mustEmbedUnimplementedSharecycleServer()
}

// UnimplementedSharecycleServer must be embedded to have forward compatible implementations.
type UnimplementedSharecycleServer struct {
}

func (UnimplementedSharecycleServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedSharecycleServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedSharecycleServer) mustEmbedUnimplementedSharecycleServer() {}

// UnsafeSharecycleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SharecycleServer will
// result in compilation errors.
type UnsafeSharecycleServer interface {
	mustEmbedUnimplementedSharecycleServer()
}

func RegisterSharecycleServer(s grpc.ServiceRegistrar, srv SharecycleServer) {
	s.RegisterService(&Sharecycle_ServiceDesc, srv)
}

func _Sharecycle_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharecycleServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sharecycle_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharecycleServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Sharecycle_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SharecycleServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Sharecycle_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SharecycleServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sharecycle_ServiceDesc is the grpc.ServiceDesc for Sharecycle service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sharecycle_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Sharecycle",
	HandlerType: (*SharecycleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _Sharecycle_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _Sharecycle_GetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_sharecycle.proto",
}
