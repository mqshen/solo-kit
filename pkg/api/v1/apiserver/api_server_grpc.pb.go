// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package apiserver

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

// ApiServerClient is the client API for ApiServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ApiServerClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (ApiServer_WatchClient, error)
}

type apiServerClient struct {
	cc grpc.ClientConnInterface
}

func NewApiServerClient(cc grpc.ClientConnInterface) ApiServerClient {
	return &apiServerClient{cc}
}

func (c *apiServerClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/apiserver.api.v1.ApiServer/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServerClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/apiserver.api.v1.ApiServer/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServerClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteResponse, error) {
	out := new(WriteResponse)
	err := c.cc.Invoke(ctx, "/apiserver.api.v1.ApiServer/Write", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServerClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/apiserver.api.v1.ApiServer/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServerClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, "/apiserver.api.v1.ApiServer/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiServerClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (ApiServer_WatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &ApiServer_ServiceDesc.Streams[0], "/apiserver.api.v1.ApiServer/Watch", opts...)
	if err != nil {
		return nil, err
	}
	x := &apiServerWatchClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ApiServer_WatchClient interface {
	Recv() (*ListResponse, error)
	grpc.ClientStream
}

type apiServerWatchClient struct {
	grpc.ClientStream
}

func (x *apiServerWatchClient) Recv() (*ListResponse, error) {
	m := new(ListResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ApiServerServer is the server API for ApiServer service.
// All implementations must embed UnimplementedApiServerServer
// for forward compatibility
type ApiServerServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Write(context.Context, *WriteRequest) (*WriteResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Watch(*WatchRequest, ApiServer_WatchServer) error
	mustEmbedUnimplementedApiServerServer()
}

// UnimplementedApiServerServer must be embedded to have forward compatible implementations.
type UnimplementedApiServerServer struct {
}

func (UnimplementedApiServerServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedApiServerServer) Read(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedApiServerServer) Write(context.Context, *WriteRequest) (*WriteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Write not implemented")
}
func (UnimplementedApiServerServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedApiServerServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedApiServerServer) Watch(*WatchRequest, ApiServer_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}
func (UnimplementedApiServerServer) mustEmbedUnimplementedApiServerServer() {}

// UnsafeApiServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ApiServerServer will
// result in compilation errors.
type UnsafeApiServerServer interface {
	mustEmbedUnimplementedApiServerServer()
}

func RegisterApiServerServer(s grpc.ServiceRegistrar, srv ApiServerServer) {
	s.RegisterService(&ApiServer_ServiceDesc, srv)
}

func _ApiServer_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServerServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiserver.api.v1.ApiServer/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServerServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiServer_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServerServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiserver.api.v1.ApiServer/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServerServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiServer_Write_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServerServer).Write(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiserver.api.v1.ApiServer/Write",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServerServer).Write(ctx, req.(*WriteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiServer_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiserver.api.v1.ApiServer/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServerServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiServer_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApiServerServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/apiserver.api.v1.ApiServer/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApiServerServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApiServer_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ApiServerServer).Watch(m, &apiServerWatchServer{stream})
}

type ApiServer_WatchServer interface {
	Send(*ListResponse) error
	grpc.ServerStream
}

type apiServerWatchServer struct {
	grpc.ServerStream
}

func (x *apiServerWatchServer) Send(m *ListResponse) error {
	return x.ServerStream.SendMsg(m)
}

// ApiServer_ServiceDesc is the grpc.ServiceDesc for ApiServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ApiServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "apiserver.api.v1.ApiServer",
	HandlerType: (*ApiServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _ApiServer_Register_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ApiServer_Read_Handler,
		},
		{
			MethodName: "Write",
			Handler:    _ApiServer_Write_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ApiServer_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ApiServer_List_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _ApiServer_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api_server.proto",
}
