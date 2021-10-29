// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ComputerManagementClient is the client API for ComputerManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComputerManagementClient interface {
	AddComputerResource(ctx context.Context, in *SetComputerResourceRequest, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
	AddComputerProject(ctx context.Context, in *SetComputerProjectRequest, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
	GetComputerProject(ctx context.Context, in *GetComputerProjectRequest, opts ...grpc.CallOption) (*GetComputerProjectResponse, error)
	DeleteComputerProject(ctx context.Context, in *wrapperspb.UInt32Value, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
	DeleteComputerResource(ctx context.Context, in *wrapperspb.UInt32Value, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error)
	GetComputerIdByMac(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*wrapperspb.UInt32Value, error)
}

type computerManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewComputerManagementClient(cc grpc.ClientConnInterface) ComputerManagementClient {
	return &computerManagementClient{cc}
}

func (c *computerManagementClient) AddComputerResource(ctx context.Context, in *SetComputerResourceRequest, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/ComputerManagement/addComputerResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computerManagementClient) AddComputerProject(ctx context.Context, in *SetComputerProjectRequest, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/ComputerManagement/addComputerProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computerManagementClient) GetComputerProject(ctx context.Context, in *GetComputerProjectRequest, opts ...grpc.CallOption) (*GetComputerProjectResponse, error) {
	out := new(GetComputerProjectResponse)
	err := c.cc.Invoke(ctx, "/ComputerManagement/getComputerProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computerManagementClient) DeleteComputerProject(ctx context.Context, in *wrapperspb.UInt32Value, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/ComputerManagement/deleteComputerProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computerManagementClient) DeleteComputerResource(ctx context.Context, in *wrapperspb.UInt32Value, opts ...grpc.CallOption) (*wrapperspb.BoolValue, error) {
	out := new(wrapperspb.BoolValue)
	err := c.cc.Invoke(ctx, "/ComputerManagement/deleteComputerResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *computerManagementClient) GetComputerIdByMac(ctx context.Context, in *wrapperspb.StringValue, opts ...grpc.CallOption) (*wrapperspb.UInt32Value, error) {
	out := new(wrapperspb.UInt32Value)
	err := c.cc.Invoke(ctx, "/ComputerManagement/getComputerIdByMac", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComputerManagementServer is the server API for ComputerManagement service.
// All implementations must embed UnimplementedComputerManagementServer
// for forward compatibility
type ComputerManagementServer interface {
	AddComputerResource(context.Context, *SetComputerResourceRequest) (*wrapperspb.BoolValue, error)
	AddComputerProject(context.Context, *SetComputerProjectRequest) (*wrapperspb.BoolValue, error)
	GetComputerProject(context.Context, *GetComputerProjectRequest) (*GetComputerProjectResponse, error)
	DeleteComputerProject(context.Context, *wrapperspb.UInt32Value) (*wrapperspb.BoolValue, error)
	DeleteComputerResource(context.Context, *wrapperspb.UInt32Value) (*wrapperspb.BoolValue, error)
	GetComputerIdByMac(context.Context, *wrapperspb.StringValue) (*wrapperspb.UInt32Value, error)
	mustEmbedUnimplementedComputerManagementServer()
}

// UnimplementedComputerManagementServer must be embedded to have forward compatible implementations.
type UnimplementedComputerManagementServer struct {
}

func (UnimplementedComputerManagementServer) AddComputerResource(context.Context, *SetComputerResourceRequest) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComputerResource not implemented")
}
func (UnimplementedComputerManagementServer) AddComputerProject(context.Context, *SetComputerProjectRequest) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddComputerProject not implemented")
}
func (UnimplementedComputerManagementServer) GetComputerProject(context.Context, *GetComputerProjectRequest) (*GetComputerProjectResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComputerProject not implemented")
}
func (UnimplementedComputerManagementServer) DeleteComputerProject(context.Context, *wrapperspb.UInt32Value) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComputerProject not implemented")
}
func (UnimplementedComputerManagementServer) DeleteComputerResource(context.Context, *wrapperspb.UInt32Value) (*wrapperspb.BoolValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComputerResource not implemented")
}
func (UnimplementedComputerManagementServer) GetComputerIdByMac(context.Context, *wrapperspb.StringValue) (*wrapperspb.UInt32Value, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComputerIdByMac not implemented")
}
func (UnimplementedComputerManagementServer) mustEmbedUnimplementedComputerManagementServer() {}

// UnsafeComputerManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComputerManagementServer will
// result in compilation errors.
type UnsafeComputerManagementServer interface {
	mustEmbedUnimplementedComputerManagementServer()
}

func RegisterComputerManagementServer(s grpc.ServiceRegistrar, srv ComputerManagementServer) {
	s.RegisterService(&ComputerManagement_ServiceDesc, srv)
}

func _ComputerManagement_AddComputerResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetComputerResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputerManagementServer).AddComputerResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ComputerManagement/addComputerResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputerManagementServer).AddComputerResource(ctx, req.(*SetComputerResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComputerManagement_AddComputerProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetComputerProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputerManagementServer).AddComputerProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ComputerManagement/addComputerProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputerManagementServer).AddComputerProject(ctx, req.(*SetComputerProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComputerManagement_GetComputerProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetComputerProjectRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputerManagementServer).GetComputerProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ComputerManagement/getComputerProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputerManagementServer).GetComputerProject(ctx, req.(*GetComputerProjectRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComputerManagement_DeleteComputerProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.UInt32Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputerManagementServer).DeleteComputerProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ComputerManagement/deleteComputerProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputerManagementServer).DeleteComputerProject(ctx, req.(*wrapperspb.UInt32Value))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComputerManagement_DeleteComputerResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.UInt32Value)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputerManagementServer).DeleteComputerResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ComputerManagement/deleteComputerResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputerManagementServer).DeleteComputerResource(ctx, req.(*wrapperspb.UInt32Value))
	}
	return interceptor(ctx, in, info, handler)
}

func _ComputerManagement_GetComputerIdByMac_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(wrapperspb.StringValue)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputerManagementServer).GetComputerIdByMac(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ComputerManagement/getComputerIdByMac",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputerManagementServer).GetComputerIdByMac(ctx, req.(*wrapperspb.StringValue))
	}
	return interceptor(ctx, in, info, handler)
}

// ComputerManagement_ServiceDesc is the grpc.ServiceDesc for ComputerManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ComputerManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ComputerManagement",
	HandlerType: (*ComputerManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "addComputerResource",
			Handler:    _ComputerManagement_AddComputerResource_Handler,
		},
		{
			MethodName: "addComputerProject",
			Handler:    _ComputerManagement_AddComputerProject_Handler,
		},
		{
			MethodName: "getComputerProject",
			Handler:    _ComputerManagement_GetComputerProject_Handler,
		},
		{
			MethodName: "deleteComputerProject",
			Handler:    _ComputerManagement_DeleteComputerProject_Handler,
		},
		{
			MethodName: "deleteComputerResource",
			Handler:    _ComputerManagement_DeleteComputerResource_Handler,
		},
		{
			MethodName: "getComputerIdByMac",
			Handler:    _ComputerManagement_GetComputerIdByMac_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpc/server/pb/computer.proto",
}
