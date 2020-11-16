// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/plgd-dev/cloud/grpc-gateway/pb/service.proto

package pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() {
	proto.RegisterFile("github.com/plgd-dev/cloud/grpc-gateway/pb/service.proto", fileDescriptor_0ca0cadb01fbb01f)
}

var fileDescriptor_0ca0cadb01fbb01f = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x94, 0x4d, 0x4f, 0xdb, 0x30,
	0x18, 0xc7, 0x9b, 0xcb, 0x26, 0x79, 0xd2, 0x34, 0x59, 0xda, 0x8b, 0x72, 0xd9, 0xd4, 0xc3, 0x36,
	0x69, 0xad, 0xd3, 0x6d, 0x14, 0x90, 0xca, 0x89, 0x96, 0xf6, 0xc2, 0xa9, 0x08, 0x0e, 0x5c, 0x50,
	0xe2, 0x3c, 0x0d, 0x16, 0x69, 0xec, 0xda, 0x4e, 0x10, 0x37, 0x6e, 0xf0, 0x15, 0xf8, 0xa4, 0x5c,
	0x31, 0x49, 0xd3, 0x57, 0x5c, 0x52, 0xce, 0xfe, 0xfd, 0x5f, 0x9e, 0x47, 0xb6, 0xd1, 0x5e, 0xc4,
	0xf4, 0x65, 0x1a, 0x10, 0xca, 0xc7, 0x9e, 0x88, 0xa3, 0xb0, 0x19, 0x42, 0xe6, 0xd1, 0x98, 0xa7,
	0xa1, 0x17, 0x49, 0x41, 0x9b, 0x91, 0xaf, 0xe1, 0xda, 0xbf, 0xf1, 0x44, 0xe0, 0x29, 0x90, 0x19,
	0xa3, 0x40, 0x84, 0xe4, 0x9a, 0xe3, 0x6f, 0x9c, 0x8e, 0x48, 0x0e, 0x92, 0x67, 0x70, 0xca, 0x11,
	0x11, 0xb8, 0xdd, 0xea, 0x96, 0x34, 0x66, 0x90, 0xe8, 0x2e, 0x4f, 0x46, 0x2c, 0x4a, 0xa5, 0xaf,
	0x19, 0x4f, 0x0a, 0x7b, 0x77, 0x8b, 0x5e, 0xe6, 0xc4, 0xd4, 0x52, 0x85, 0xf0, 0xdf, 0xe3, 0x7b,
	0xf4, 0x61, 0x60, 0x88, 0x41, 0x01, 0xe0, 0x0b, 0x84, 0x06, 0xa0, 0x7b, 0x05, 0x83, 0xff, 0x10,
	0x5b, 0x6d, 0x32, 0xa7, 0x86, 0x30, 0x49, 0x41, 0x69, 0xf7, 0x87, 0x1d, 0x2e, 0xc8, 0x7a, 0xad,
	0xe5, 0xe0, 0x09, 0xfa, 0x64, 0xa4, 0x43, 0x50, 0x3c, 0x95, 0x14, 0x8e, 0x59, 0x72, 0xa5, 0xf0,
	0xdf, 0x8d, 0x31, 0x4b, 0x6c, 0x19, 0xf6, 0xd3, 0x2e, 0x59, 0xe4, 0xf3, 0xc8, 0x07, 0x07, 0xb9,
	0x43, 0xd0, 0x92, 0x41, 0x06, 0xe5, 0x61, 0x5f, 0xf2, 0x71, 0xd1, 0x0a, 0x77, 0x36, 0x59, 0xd9,
	0x54, 0x65, 0x8f, 0x83, 0xb7, 0x89, 0x95, 0xe0, 0x89, 0x32, 0x0b, 0xc1, 0xb7, 0x0e, 0xfa, 0xba,
	0x0a, 0xaa, 0x33, 0x3f, 0x36, 0xe6, 0x78, 0xbf, 0xba, 0xf7, 0x54, 0x52, 0xb6, 0xfa, 0xf5, 0xfa,
	0x76, 0x72, 0x41, 0xbe, 0x9e, 0x7b, 0x07, 0x7d, 0x3e, 0x15, 0xa1, 0x41, 0x56, 0x0b, 0xb4, 0xed,
	0x36, 0xcb, 0x82, 0xe5, 0xf4, 0xdd, 0x6d, 0x65, 0xb3, 0x6d, 0x30, 0x84, 0x4f, 0xd2, 0x40, 0x51,
	0xc9, 0x02, 0xe8, 0x73, 0x79, 0x94, 0x99, 0xeb, 0xae, 0x70, 0xc3, 0xee, 0xb7, 0x4e, 0xbb, 0xdf,
	0xed, 0x74, 0x4e, 0xd4, 0x6b, 0xbf, 0x1d, 0x33, 0xf5, 0x9d, 0x83, 0xbe, 0x98, 0xcb, 0xd5, 0x5d,
	0x7f, 0x52, 0x78, 0xc7, 0xee, 0xf0, 0x02, 0x5e, 0x4e, 0xdd, 0xde, 0x52, 0x35, 0x1b, 0x3a, 0x45,
	0x1f, 0x7b, 0x10, 0xc3, 0x7c, 0x2d, 0xd8, 0xdb, 0xf4, 0x92, 0x16, 0xc9, 0x32, 0xbb, 0x55, 0x5d,
	0x50, 0xc6, 0x1e, 0x92, 0xf3, 0x46, 0xe5, 0x4f, 0xa3, 0x23, 0x82, 0xe0, 0x5d, 0xfe, 0x61, 0xfc,
	0x7f, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x50, 0x64, 0x44, 0xf7, 0x03, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GrpcGatewayClient is the client API for GrpcGateway service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GrpcGatewayClient interface {
	// Get all devices
	GetDevices(ctx context.Context, in *GetDevicesRequest, opts ...grpc.CallOption) (GrpcGateway_GetDevicesClient, error)
	// Get resource links of devices.
	GetResourceLinks(ctx context.Context, in *GetResourceLinksRequest, opts ...grpc.CallOption) (GrpcGateway_GetResourceLinksClient, error)
	RetrieveResourceFromDevice(ctx context.Context, in *RetrieveResourceFromDeviceRequest, opts ...grpc.CallOption) (*RetrieveResourceFromDeviceResponse, error)
	// Retrieve resources values from resource shadow
	RetrieveResourcesValues(ctx context.Context, in *RetrieveResourcesValuesRequest, opts ...grpc.CallOption) (GrpcGateway_RetrieveResourcesValuesClient, error)
	// Update resource values
	UpdateResourcesValues(ctx context.Context, in *UpdateResourceValuesRequest, opts ...grpc.CallOption) (*UpdateResourceValuesResponse, error)
	// Subscribe to events
	SubscribeForEvents(ctx context.Context, opts ...grpc.CallOption) (GrpcGateway_SubscribeForEventsClient, error)
	// Get client configuration
	GetClientConfiguration(ctx context.Context, in *ClientConfigurationRequest, opts ...grpc.CallOption) (*ClientConfigurationResponse, error)
	// Delete resource at the device.
	DeleteResource(ctx context.Context, in *DeleteResourceRequest, opts ...grpc.CallOption) (*DeleteResourceResponse, error)
}

type grpcGatewayClient struct {
	cc grpc.ClientConnInterface
}

func NewGrpcGatewayClient(cc grpc.ClientConnInterface) GrpcGatewayClient {
	return &grpcGatewayClient{cc}
}

func (c *grpcGatewayClient) GetDevices(ctx context.Context, in *GetDevicesRequest, opts ...grpc.CallOption) (GrpcGateway_GetDevicesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GrpcGateway_serviceDesc.Streams[0], "/ocf.cloud.grpcgateway.pb.GrpcGateway/GetDevices", opts...)
	if err != nil {
		return nil, err
	}
	x := &grpcGatewayGetDevicesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GrpcGateway_GetDevicesClient interface {
	Recv() (*Device, error)
	grpc.ClientStream
}

type grpcGatewayGetDevicesClient struct {
	grpc.ClientStream
}

func (x *grpcGatewayGetDevicesClient) Recv() (*Device, error) {
	m := new(Device)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *grpcGatewayClient) GetResourceLinks(ctx context.Context, in *GetResourceLinksRequest, opts ...grpc.CallOption) (GrpcGateway_GetResourceLinksClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GrpcGateway_serviceDesc.Streams[1], "/ocf.cloud.grpcgateway.pb.GrpcGateway/GetResourceLinks", opts...)
	if err != nil {
		return nil, err
	}
	x := &grpcGatewayGetResourceLinksClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GrpcGateway_GetResourceLinksClient interface {
	Recv() (*ResourceLink, error)
	grpc.ClientStream
}

type grpcGatewayGetResourceLinksClient struct {
	grpc.ClientStream
}

func (x *grpcGatewayGetResourceLinksClient) Recv() (*ResourceLink, error) {
	m := new(ResourceLink)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *grpcGatewayClient) RetrieveResourceFromDevice(ctx context.Context, in *RetrieveResourceFromDeviceRequest, opts ...grpc.CallOption) (*RetrieveResourceFromDeviceResponse, error) {
	out := new(RetrieveResourceFromDeviceResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.grpcgateway.pb.GrpcGateway/RetrieveResourceFromDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcGatewayClient) RetrieveResourcesValues(ctx context.Context, in *RetrieveResourcesValuesRequest, opts ...grpc.CallOption) (GrpcGateway_RetrieveResourcesValuesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GrpcGateway_serviceDesc.Streams[2], "/ocf.cloud.grpcgateway.pb.GrpcGateway/RetrieveResourcesValues", opts...)
	if err != nil {
		return nil, err
	}
	x := &grpcGatewayRetrieveResourcesValuesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GrpcGateway_RetrieveResourcesValuesClient interface {
	Recv() (*ResourceValue, error)
	grpc.ClientStream
}

type grpcGatewayRetrieveResourcesValuesClient struct {
	grpc.ClientStream
}

func (x *grpcGatewayRetrieveResourcesValuesClient) Recv() (*ResourceValue, error) {
	m := new(ResourceValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *grpcGatewayClient) UpdateResourcesValues(ctx context.Context, in *UpdateResourceValuesRequest, opts ...grpc.CallOption) (*UpdateResourceValuesResponse, error) {
	out := new(UpdateResourceValuesResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.grpcgateway.pb.GrpcGateway/UpdateResourcesValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcGatewayClient) SubscribeForEvents(ctx context.Context, opts ...grpc.CallOption) (GrpcGateway_SubscribeForEventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GrpcGateway_serviceDesc.Streams[3], "/ocf.cloud.grpcgateway.pb.GrpcGateway/SubscribeForEvents", opts...)
	if err != nil {
		return nil, err
	}
	x := &grpcGatewaySubscribeForEventsClient{stream}
	return x, nil
}

type GrpcGateway_SubscribeForEventsClient interface {
	Send(*SubscribeForEvents) error
	Recv() (*Event, error)
	grpc.ClientStream
}

type grpcGatewaySubscribeForEventsClient struct {
	grpc.ClientStream
}

func (x *grpcGatewaySubscribeForEventsClient) Send(m *SubscribeForEvents) error {
	return x.ClientStream.SendMsg(m)
}

func (x *grpcGatewaySubscribeForEventsClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *grpcGatewayClient) GetClientConfiguration(ctx context.Context, in *ClientConfigurationRequest, opts ...grpc.CallOption) (*ClientConfigurationResponse, error) {
	out := new(ClientConfigurationResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.grpcgateway.pb.GrpcGateway/GetClientConfiguration", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcGatewayClient) DeleteResource(ctx context.Context, in *DeleteResourceRequest, opts ...grpc.CallOption) (*DeleteResourceResponse, error) {
	out := new(DeleteResourceResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.grpcgateway.pb.GrpcGateway/DeleteResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrpcGatewayServer is the server API for GrpcGateway service.
type GrpcGatewayServer interface {
	// Get all devices
	GetDevices(*GetDevicesRequest, GrpcGateway_GetDevicesServer) error
	// Get resource links of devices.
	GetResourceLinks(*GetResourceLinksRequest, GrpcGateway_GetResourceLinksServer) error
	RetrieveResourceFromDevice(context.Context, *RetrieveResourceFromDeviceRequest) (*RetrieveResourceFromDeviceResponse, error)
	// Retrieve resources values from resource shadow
	RetrieveResourcesValues(*RetrieveResourcesValuesRequest, GrpcGateway_RetrieveResourcesValuesServer) error
	// Update resource values
	UpdateResourcesValues(context.Context, *UpdateResourceValuesRequest) (*UpdateResourceValuesResponse, error)
	// Subscribe to events
	SubscribeForEvents(GrpcGateway_SubscribeForEventsServer) error
	// Get client configuration
	GetClientConfiguration(context.Context, *ClientConfigurationRequest) (*ClientConfigurationResponse, error)
	// Delete resource at the device.
	DeleteResource(context.Context, *DeleteResourceRequest) (*DeleteResourceResponse, error)
}

// UnimplementedGrpcGatewayServer can be embedded to have forward compatible implementations.
type UnimplementedGrpcGatewayServer struct {
}

func (*UnimplementedGrpcGatewayServer) GetDevices(req *GetDevicesRequest, srv GrpcGateway_GetDevicesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetDevices not implemented")
}
func (*UnimplementedGrpcGatewayServer) GetResourceLinks(req *GetResourceLinksRequest, srv GrpcGateway_GetResourceLinksServer) error {
	return status.Errorf(codes.Unimplemented, "method GetResourceLinks not implemented")
}
func (*UnimplementedGrpcGatewayServer) RetrieveResourceFromDevice(ctx context.Context, req *RetrieveResourceFromDeviceRequest) (*RetrieveResourceFromDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveResourceFromDevice not implemented")
}
func (*UnimplementedGrpcGatewayServer) RetrieveResourcesValues(req *RetrieveResourcesValuesRequest, srv GrpcGateway_RetrieveResourcesValuesServer) error {
	return status.Errorf(codes.Unimplemented, "method RetrieveResourcesValues not implemented")
}
func (*UnimplementedGrpcGatewayServer) UpdateResourcesValues(ctx context.Context, req *UpdateResourceValuesRequest) (*UpdateResourceValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateResourcesValues not implemented")
}
func (*UnimplementedGrpcGatewayServer) SubscribeForEvents(srv GrpcGateway_SubscribeForEventsServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeForEvents not implemented")
}
func (*UnimplementedGrpcGatewayServer) GetClientConfiguration(ctx context.Context, req *ClientConfigurationRequest) (*ClientConfigurationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClientConfiguration not implemented")
}
func (*UnimplementedGrpcGatewayServer) DeleteResource(ctx context.Context, req *DeleteResourceRequest) (*DeleteResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteResource not implemented")
}

func RegisterGrpcGatewayServer(s *grpc.Server, srv GrpcGatewayServer) {
	s.RegisterService(&_GrpcGateway_serviceDesc, srv)
}

func _GrpcGateway_GetDevices_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetDevicesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GrpcGatewayServer).GetDevices(m, &grpcGatewayGetDevicesServer{stream})
}

type GrpcGateway_GetDevicesServer interface {
	Send(*Device) error
	grpc.ServerStream
}

type grpcGatewayGetDevicesServer struct {
	grpc.ServerStream
}

func (x *grpcGatewayGetDevicesServer) Send(m *Device) error {
	return x.ServerStream.SendMsg(m)
}

func _GrpcGateway_GetResourceLinks_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetResourceLinksRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GrpcGatewayServer).GetResourceLinks(m, &grpcGatewayGetResourceLinksServer{stream})
}

type GrpcGateway_GetResourceLinksServer interface {
	Send(*ResourceLink) error
	grpc.ServerStream
}

type grpcGatewayGetResourceLinksServer struct {
	grpc.ServerStream
}

func (x *grpcGatewayGetResourceLinksServer) Send(m *ResourceLink) error {
	return x.ServerStream.SendMsg(m)
}

func _GrpcGateway_RetrieveResourceFromDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveResourceFromDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcGatewayServer).RetrieveResourceFromDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.grpcgateway.pb.GrpcGateway/RetrieveResourceFromDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcGatewayServer).RetrieveResourceFromDevice(ctx, req.(*RetrieveResourceFromDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrpcGateway_RetrieveResourcesValues_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RetrieveResourcesValuesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GrpcGatewayServer).RetrieveResourcesValues(m, &grpcGatewayRetrieveResourcesValuesServer{stream})
}

type GrpcGateway_RetrieveResourcesValuesServer interface {
	Send(*ResourceValue) error
	grpc.ServerStream
}

type grpcGatewayRetrieveResourcesValuesServer struct {
	grpc.ServerStream
}

func (x *grpcGatewayRetrieveResourcesValuesServer) Send(m *ResourceValue) error {
	return x.ServerStream.SendMsg(m)
}

func _GrpcGateway_UpdateResourcesValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateResourceValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcGatewayServer).UpdateResourcesValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.grpcgateway.pb.GrpcGateway/UpdateResourcesValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcGatewayServer).UpdateResourcesValues(ctx, req.(*UpdateResourceValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrpcGateway_SubscribeForEvents_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GrpcGatewayServer).SubscribeForEvents(&grpcGatewaySubscribeForEventsServer{stream})
}

type GrpcGateway_SubscribeForEventsServer interface {
	Send(*Event) error
	Recv() (*SubscribeForEvents, error)
	grpc.ServerStream
}

type grpcGatewaySubscribeForEventsServer struct {
	grpc.ServerStream
}

func (x *grpcGatewaySubscribeForEventsServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func (x *grpcGatewaySubscribeForEventsServer) Recv() (*SubscribeForEvents, error) {
	m := new(SubscribeForEvents)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GrpcGateway_GetClientConfiguration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientConfigurationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcGatewayServer).GetClientConfiguration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.grpcgateway.pb.GrpcGateway/GetClientConfiguration",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcGatewayServer).GetClientConfiguration(ctx, req.(*ClientConfigurationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrpcGateway_DeleteResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcGatewayServer).DeleteResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.grpcgateway.pb.GrpcGateway/DeleteResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcGatewayServer).DeleteResource(ctx, req.(*DeleteResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GrpcGateway_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ocf.cloud.grpcgateway.pb.GrpcGateway",
	HandlerType: (*GrpcGatewayServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RetrieveResourceFromDevice",
			Handler:    _GrpcGateway_RetrieveResourceFromDevice_Handler,
		},
		{
			MethodName: "UpdateResourcesValues",
			Handler:    _GrpcGateway_UpdateResourcesValues_Handler,
		},
		{
			MethodName: "GetClientConfiguration",
			Handler:    _GrpcGateway_GetClientConfiguration_Handler,
		},
		{
			MethodName: "DeleteResource",
			Handler:    _GrpcGateway_DeleteResource_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetDevices",
			Handler:       _GrpcGateway_GetDevices_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetResourceLinks",
			Handler:       _GrpcGateway_GetResourceLinks_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RetrieveResourcesValues",
			Handler:       _GrpcGateway_RetrieveResourcesValues_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeForEvents",
			Handler:       _GrpcGateway_SubscribeForEvents_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "github.com/plgd-dev/cloud/grpc-gateway/pb/service.proto",
}
