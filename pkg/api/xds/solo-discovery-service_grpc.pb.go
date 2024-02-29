// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package xds

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

// SoloDiscoveryServiceClient is the client API for SoloDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SoloDiscoveryServiceClient interface {
	// This is a gRPC-only API.
	StreamAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (SoloDiscoveryService_StreamAggregatedResourcesClient, error)
	DeltaAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (SoloDiscoveryService_DeltaAggregatedResourcesClient, error)
}

type soloDiscoveryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSoloDiscoveryServiceClient(cc grpc.ClientConnInterface) SoloDiscoveryServiceClient {
	return &soloDiscoveryServiceClient{cc}
}

func (c *soloDiscoveryServiceClient) StreamAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (SoloDiscoveryService_StreamAggregatedResourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &SoloDiscoveryService_ServiceDesc.Streams[0], "/solo.io.xds.SoloDiscoveryService/StreamAggregatedResources", opts...)
	if err != nil {
		return nil, err
	}
	x := &soloDiscoveryServiceStreamAggregatedResourcesClient{stream}
	return x, nil
}

type SoloDiscoveryService_StreamAggregatedResourcesClient interface {
	Send(*v2.DiscoveryRequest) error
	Recv() (*v2.DiscoveryResponse, error)
	grpc.ClientStream
}

type soloDiscoveryServiceStreamAggregatedResourcesClient struct {
	grpc.ClientStream
}

func (x *soloDiscoveryServiceStreamAggregatedResourcesClient) Send(m *v2.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *soloDiscoveryServiceStreamAggregatedResourcesClient) Recv() (*v2.DiscoveryResponse, error) {
	m := new(v2.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *soloDiscoveryServiceClient) DeltaAggregatedResources(ctx context.Context, opts ...grpc.CallOption) (SoloDiscoveryService_DeltaAggregatedResourcesClient, error) {
	stream, err := c.cc.NewStream(ctx, &SoloDiscoveryService_ServiceDesc.Streams[1], "/solo.io.xds.SoloDiscoveryService/DeltaAggregatedResources", opts...)
	if err != nil {
		return nil, err
	}
	x := &soloDiscoveryServiceDeltaAggregatedResourcesClient{stream}
	return x, nil
}

type SoloDiscoveryService_DeltaAggregatedResourcesClient interface {
	Send(*v2.DeltaDiscoveryRequest) error
	Recv() (*v2.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type soloDiscoveryServiceDeltaAggregatedResourcesClient struct {
	grpc.ClientStream
}

func (x *soloDiscoveryServiceDeltaAggregatedResourcesClient) Send(m *v2.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *soloDiscoveryServiceDeltaAggregatedResourcesClient) Recv() (*v2.DeltaDiscoveryResponse, error) {
	m := new(v2.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SoloDiscoveryServiceServer is the server API for SoloDiscoveryService service.
// All implementations must embed UnimplementedSoloDiscoveryServiceServer
// for forward compatibility
type SoloDiscoveryServiceServer interface {
	// This is a gRPC-only API.
	StreamAggregatedResources(SoloDiscoveryService_StreamAggregatedResourcesServer) error
	DeltaAggregatedResources(SoloDiscoveryService_DeltaAggregatedResourcesServer) error
	mustEmbedUnimplementedSoloDiscoveryServiceServer()
}

// UnimplementedSoloDiscoveryServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSoloDiscoveryServiceServer struct {
}

func (UnimplementedSoloDiscoveryServiceServer) StreamAggregatedResources(SoloDiscoveryService_StreamAggregatedResourcesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamAggregatedResources not implemented")
}
func (UnimplementedSoloDiscoveryServiceServer) DeltaAggregatedResources(SoloDiscoveryService_DeltaAggregatedResourcesServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaAggregatedResources not implemented")
}
func (UnimplementedSoloDiscoveryServiceServer) mustEmbedUnimplementedSoloDiscoveryServiceServer() {}

// UnsafeSoloDiscoveryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SoloDiscoveryServiceServer will
// result in compilation errors.
type UnsafeSoloDiscoveryServiceServer interface {
	mustEmbedUnimplementedSoloDiscoveryServiceServer()
}

func RegisterSoloDiscoveryServiceServer(s grpc.ServiceRegistrar, srv SoloDiscoveryServiceServer) {
	s.RegisterService(&SoloDiscoveryService_ServiceDesc, srv)
}

func _SoloDiscoveryService_StreamAggregatedResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SoloDiscoveryServiceServer).StreamAggregatedResources(&soloDiscoveryServiceStreamAggregatedResourcesServer{stream})
}

type SoloDiscoveryService_StreamAggregatedResourcesServer interface {
	Send(*v2.DiscoveryResponse) error
	Recv() (*v2.DiscoveryRequest, error)
	grpc.ServerStream
}

type soloDiscoveryServiceStreamAggregatedResourcesServer struct {
	grpc.ServerStream
}

func (x *soloDiscoveryServiceStreamAggregatedResourcesServer) Send(m *v2.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *soloDiscoveryServiceStreamAggregatedResourcesServer) Recv() (*v2.DiscoveryRequest, error) {
	m := new(v2.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _SoloDiscoveryService_DeltaAggregatedResources_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SoloDiscoveryServiceServer).DeltaAggregatedResources(&soloDiscoveryServiceDeltaAggregatedResourcesServer{stream})
}

type SoloDiscoveryService_DeltaAggregatedResourcesServer interface {
	Send(*v2.DeltaDiscoveryResponse) error
	Recv() (*v2.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type soloDiscoveryServiceDeltaAggregatedResourcesServer struct {
	grpc.ServerStream
}

func (x *soloDiscoveryServiceDeltaAggregatedResourcesServer) Send(m *v2.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *soloDiscoveryServiceDeltaAggregatedResourcesServer) Recv() (*v2.DeltaDiscoveryRequest, error) {
	m := new(v2.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// SoloDiscoveryService_ServiceDesc is the grpc.ServiceDesc for SoloDiscoveryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SoloDiscoveryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "solo.io.xds.SoloDiscoveryService",
	HandlerType: (*SoloDiscoveryServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamAggregatedResources",
			Handler:       _SoloDiscoveryService_StreamAggregatedResources_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaAggregatedResources",
			Handler:       _SoloDiscoveryService_DeltaAggregatedResources_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/solo-io/solo-kit/api/xds/solo-discovery-service.proto",
}
