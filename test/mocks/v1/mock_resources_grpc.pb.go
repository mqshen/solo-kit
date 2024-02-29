// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	v2 "github.com/solo-io/solo-kit/pkg/api/external/envoy/api/v2"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MockXdsResourceDiscoveryServiceClient is the client API for MockXdsResourceDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MockXdsResourceDiscoveryServiceClient interface {
	StreamMockXdsResourceConfig(ctx context.Context, opts ...grpc.CallOption) (MockXdsResourceDiscoveryService_StreamMockXdsResourceConfigClient, error)
	DeltaMockXdsResourceConfig(ctx context.Context, opts ...grpc.CallOption) (MockXdsResourceDiscoveryService_DeltaMockXdsResourceConfigClient, error)
	FetchMockXdsResourceConfig(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error)
}

type mockXdsResourceDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMockXdsResourceDiscoveryServiceClient(cc grpc.ClientConnInterface) MockXdsResourceDiscoveryServiceClient {
	return &mockXdsResourceDiscoveryServiceClient{cc}
}

func (c *mockXdsResourceDiscoveryServiceClient) StreamMockXdsResourceConfig(ctx context.Context, opts ...grpc.CallOption) (MockXdsResourceDiscoveryService_StreamMockXdsResourceConfigClient, error) {
	stream, err := c.cc.NewStream(ctx, &MockXdsResourceDiscoveryService_ServiceDesc.Streams[0], "/testing.solo.io.MockXdsResourceDiscoveryService/StreamMockXdsResourceConfig", opts...)
	if err != nil {
		return nil, err
	}
	x := &mockXdsResourceDiscoveryServiceStreamMockXdsResourceConfigClient{stream}
	return x, nil
}

type MockXdsResourceDiscoveryService_StreamMockXdsResourceConfigClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type mockXdsResourceDiscoveryServiceStreamMockXdsResourceConfigClient struct {
	grpc.ClientStream
}

func (x *mockXdsResourceDiscoveryServiceStreamMockXdsResourceConfigClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mockXdsResourceDiscoveryServiceStreamMockXdsResourceConfigClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mockXdsResourceDiscoveryServiceClient) DeltaMockXdsResourceConfig(ctx context.Context, opts ...grpc.CallOption) (MockXdsResourceDiscoveryService_DeltaMockXdsResourceConfigClient, error) {
	stream, err := c.cc.NewStream(ctx, &MockXdsResourceDiscoveryService_ServiceDesc.Streams[1], "/testing.solo.io.MockXdsResourceDiscoveryService/DeltaMockXdsResourceConfig", opts...)
	if err != nil {
		return nil, err
	}
	x := &mockXdsResourceDiscoveryServiceDeltaMockXdsResourceConfigClient{stream}
	return x, nil
}

type MockXdsResourceDiscoveryService_DeltaMockXdsResourceConfigClient interface {
	Send(*v2.DeltaDiscoveryRequest) error
	Recv() (*v2.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type mockXdsResourceDiscoveryServiceDeltaMockXdsResourceConfigClient struct {
	grpc.ClientStream
}

func (x *mockXdsResourceDiscoveryServiceDeltaMockXdsResourceConfigClient) Send(m *v2.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *mockXdsResourceDiscoveryServiceDeltaMockXdsResourceConfigClient) Recv() (*v2.DeltaDiscoveryResponse, error) {
	m := new(v2.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *mockXdsResourceDiscoveryServiceClient) FetchMockXdsResourceConfig(ctx context.Context, in *v2.DiscoveryRequest, opts ...grpc.CallOption) (*v2.DiscoveryResponse, error) {
	out := new(v2.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/testing.solo.io.MockXdsResourceDiscoveryService/FetchMockXdsResourceConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MockXdsResourceDiscoveryServiceServer is the server API for MockXdsResourceDiscoveryService service.
// All implementations must embed UnimplementedMockXdsResourceDiscoveryServiceServer
// for forward compatibility
type MockXdsResourceDiscoveryServiceServer interface {
	StreamMockXdsResourceConfig(MockXdsResourceDiscoveryService_StreamMockXdsResourceConfigServer) error
	DeltaMockXdsResourceConfig(MockXdsResourceDiscoveryService_DeltaMockXdsResourceConfigServer) error
	FetchMockXdsResourceConfig(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error)
	mustEmbedUnimplementedMockXdsResourceDiscoveryServiceServer()
}

// UnimplementedMockXdsResourceDiscoveryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMockXdsResourceDiscoveryServiceServer struct {
}

func (UnimplementedMockXdsResourceDiscoveryServiceServer) StreamMockXdsResourceConfig(MockXdsResourceDiscoveryService_StreamMockXdsResourceConfigServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamMockXdsResourceConfig not implemented")
}
func (UnimplementedMockXdsResourceDiscoveryServiceServer) DeltaMockXdsResourceConfig(MockXdsResourceDiscoveryService_DeltaMockXdsResourceConfigServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaMockXdsResourceConfig not implemented")
}
func (UnimplementedMockXdsResourceDiscoveryServiceServer) FetchMockXdsResourceConfig(context.Context, *v2.DiscoveryRequest) (*v2.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchMockXdsResourceConfig not implemented")
}
func (UnimplementedMockXdsResourceDiscoveryServiceServer) mustEmbedUnimplementedMockXdsResourceDiscoveryServiceServer() {
}

// UnsafeMockXdsResourceDiscoveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MockXdsResourceDiscoveryServiceServer will
// result in compilation errors.
type UnsafeMockXdsResourceDiscoveryServiceServer interface {
	mustEmbedUnimplementedMockXdsResourceDiscoveryServiceServer()
}

func RegisterMockXdsResourceDiscoveryServiceServer(s grpc.ServiceRegistrar, srv MockXdsResourceDiscoveryServiceServer) {
	s.RegisterService(&MockXdsResourceDiscoveryService_ServiceDesc, srv)
}

func _MockXdsResourceDiscoveryService_StreamMockXdsResourceConfig_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MockXdsResourceDiscoveryServiceServer).StreamMockXdsResourceConfig(&mockXdsResourceDiscoveryServiceStreamMockXdsResourceConfigServer{stream})
}

type MockXdsResourceDiscoveryService_StreamMockXdsResourceConfigServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type mockXdsResourceDiscoveryServiceStreamMockXdsResourceConfigServer struct {
	grpc.ServerStream
}

func (x *mockXdsResourceDiscoveryServiceStreamMockXdsResourceConfigServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mockXdsResourceDiscoveryServiceStreamMockXdsResourceConfigServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MockXdsResourceDiscoveryService_DeltaMockXdsResourceConfig_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MockXdsResourceDiscoveryServiceServer).DeltaMockXdsResourceConfig(&mockXdsResourceDiscoveryServiceDeltaMockXdsResourceConfigServer{stream})
}

type MockXdsResourceDiscoveryService_DeltaMockXdsResourceConfigServer interface {
	Send(*v2.DeltaDiscoveryResponse) error
	Recv() (*v2.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type mockXdsResourceDiscoveryServiceDeltaMockXdsResourceConfigServer struct {
	grpc.ServerStream
}

func (x *mockXdsResourceDiscoveryServiceDeltaMockXdsResourceConfigServer) Send(m *v2.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *mockXdsResourceDiscoveryServiceDeltaMockXdsResourceConfigServer) Recv() (*v2.DeltaDiscoveryRequest, error) {
	m := new(v2.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MockXdsResourceDiscoveryService_FetchMockXdsResourceConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v2.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MockXdsResourceDiscoveryServiceServer).FetchMockXdsResourceConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testing.solo.io.MockXdsResourceDiscoveryService/FetchMockXdsResourceConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MockXdsResourceDiscoveryServiceServer).FetchMockXdsResourceConfig(ctx, req.(*v2.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MockXdsResourceDiscoveryService_ServiceDesc is the grpc.ServiceDesc for MockXdsResourceDiscoveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MockXdsResourceDiscoveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "testing.solo.io.MockXdsResourceDiscoveryService",
	HandlerType: (*MockXdsResourceDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchMockXdsResourceConfig",
			Handler:    _MockXdsResourceDiscoveryService_FetchMockXdsResourceConfig_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamMockXdsResourceConfig",
			Handler:       _MockXdsResourceDiscoveryService_StreamMockXdsResourceConfig_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaMockXdsResourceConfig",
			Handler:       _MockXdsResourceDiscoveryService_DeltaMockXdsResourceConfig_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/solo-io/solo-kit/test/mocks/api/v1/mock_resources.proto",
}
