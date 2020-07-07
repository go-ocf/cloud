// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/go-ocf/cloud/resource-aggregate/pb/service.proto

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
	proto.RegisterFile("github.com/go-ocf/cloud/resource-aggregate/pb/service.proto", fileDescriptor_51b717c14aa1bbfc)
}

var fileDescriptor_51b717c14aa1bbfc = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x94, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xed, 0x45, 0x24, 0x07, 0x75, 0x05, 0x11, 0x76, 0xf0, 0xe0, 0x7d, 0xa9, 0x4c, 0x9c,
	0x93, 0x4d, 0x45, 0x77, 0x17, 0x19, 0xec, 0xe2, 0xad, 0x4d, 0x5f, 0xb3, 0xc0, 0xda, 0x17, 0x93,
	0x74, 0xe0, 0x51, 0x10, 0x05, 0x3f, 0x80, 0x27, 0x3f, 0xac, 0xb0, 0x99, 0x8c, 0xd5, 0x69, 0x4d,
	0xcf, 0xef, 0xff, 0x7b, 0xfd, 0xe5, 0x4f, 0x79, 0x64, 0xc0, 0x85, 0x99, 0x96, 0x09, 0x65, 0x98,
	0x47, 0x1c, 0x3b, 0xc8, 0xb2, 0x88, 0xcd, 0xb0, 0x4c, 0x23, 0x05, 0x1a, 0x4b, 0xc5, 0xa0, 0x13,
	0x73, 0xae, 0x80, 0xc7, 0x06, 0x22, 0x99, 0x44, 0x1a, 0xd4, 0x5c, 0x30, 0xa0, 0x52, 0xa1, 0xc1,
	0xf0, 0x08, 0x59, 0x46, 0x17, 0x71, 0x6a, 0xe3, 0x2e, 0x4d, 0x65, 0xd2, 0x1e, 0xfa, 0x2d, 0x67,
	0x98, 0xe7, 0x71, 0x91, 0xea, 0xe5, 0xf6, 0xee, 0xeb, 0x0e, 0x69, 0x8d, 0xbf, 0x83, 0x37, 0x36,
	0x17, 0xbe, 0x04, 0x64, 0xef, 0xbe, 0x4c, 0x66, 0x42, 0x4f, 0xed, 0x30, 0xec, 0xd1, 0xbf, 0x45,
	0x68, 0x05, 0x18, 0xc3, 0x63, 0x09, 0xda, 0xb4, 0xcf, 0xbd, 0x39, 0x2d, 0xb1, 0xd0, 0x70, 0xbc,
	0x15, 0xbe, 0x07, 0xa4, 0x35, 0x29, 0x64, 0x45, 0xa4, 0x5f, 0xb7, 0xf0, 0x07, 0x62, 0x55, 0x2e,
	0x1a, 0x90, 0x4e, 0xe6, 0x23, 0x20, 0x07, 0x77, 0x68, 0x44, 0xf6, 0x64, 0x87, 0xa3, 0x69, 0x5c,
	0x70, 0x48, 0xc3, 0x61, 0xdd, 0xda, 0x8d, 0x98, 0x95, 0xba, 0x6c, 0x48, 0x3b, 0xb1, 0xe7, 0x80,
	0xec, 0x4e, 0x64, 0x1a, 0x1b, 0x70, 0x15, 0x9d, 0xd5, 0x3e, 0x74, 0x2d, 0x6f, 0x55, 0x7a, 0xbe,
	0xd8, 0x5a, 0x39, 0x23, 0x2c, 0x32, 0xa1, 0x72, 0x3b, 0x5d, 0x66, 0xeb, 0xcb, 0xd9, 0x88, 0xfd,
	0xbb, 0x9c, 0x5f, 0x68, 0x27, 0xf6, 0x16, 0x90, 0xfd, 0x31, 0x18, 0x25, 0x60, 0xbe, 0xaa, 0xa7,
	0xf6, 0x97, 0xac, 0x12, 0x56, 0xa7, 0xef, 0x0f, 0x3a, 0x93, 0xcf, 0x80, 0x1c, 0x56, 0x6c, 0x6d,
	0x3a, 0xbc, 0xf2, 0x7c, 0xe6, 0xea, 0x33, 0x4b, 0xaf, 0xeb, 0xc6, 0xbc, 0xd5, 0xbb, 0xed, 0x3e,
	0x9c, 0x78, 0x1d, 0x92, 0x81, 0x4c, 0x92, 0xed, 0xc5, 0x0d, 0x39, 0xfd, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0x4e, 0x57, 0x57, 0xe9, 0xe0, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ResourceAggregateClient is the client API for ResourceAggregate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResourceAggregateClient interface {
	PublishResource(ctx context.Context, in *PublishResourceRequest, opts ...grpc.CallOption) (*PublishResourceResponse, error)
	UnpublishResource(ctx context.Context, in *UnpublishResourceRequest, opts ...grpc.CallOption) (*UnpublishResourceResponse, error)
	NotifyResourceChanged(ctx context.Context, in *NotifyResourceChangedRequest, opts ...grpc.CallOption) (*NotifyResourceChangedResponse, error)
	UpdateResource(ctx context.Context, in *UpdateResourceRequest, opts ...grpc.CallOption) (*UpdateResourceResponse, error)
	ConfirmResourceUpdate(ctx context.Context, in *ConfirmResourceUpdateRequest, opts ...grpc.CallOption) (*ConfirmResourceUpdateResponse, error)
	RetrieveResource(ctx context.Context, in *RetrieveResourceRequest, opts ...grpc.CallOption) (*RetrieveResourceResponse, error)
	ConfirmResourceRetrieve(ctx context.Context, in *ConfirmResourceRetrieveRequest, opts ...grpc.CallOption) (*ConfirmResourceRetrieveResponse, error)
}

type resourceAggregateClient struct {
	cc *grpc.ClientConn
}

func NewResourceAggregateClient(cc *grpc.ClientConn) ResourceAggregateClient {
	return &resourceAggregateClient{cc}
}

func (c *resourceAggregateClient) PublishResource(ctx context.Context, in *PublishResourceRequest, opts ...grpc.CallOption) (*PublishResourceResponse, error) {
	out := new(PublishResourceResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/PublishResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) UnpublishResource(ctx context.Context, in *UnpublishResourceRequest, opts ...grpc.CallOption) (*UnpublishResourceResponse, error) {
	out := new(UnpublishResourceResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/UnpublishResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) NotifyResourceChanged(ctx context.Context, in *NotifyResourceChangedRequest, opts ...grpc.CallOption) (*NotifyResourceChangedResponse, error) {
	out := new(NotifyResourceChangedResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/NotifyResourceChanged", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) UpdateResource(ctx context.Context, in *UpdateResourceRequest, opts ...grpc.CallOption) (*UpdateResourceResponse, error) {
	out := new(UpdateResourceResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/UpdateResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) ConfirmResourceUpdate(ctx context.Context, in *ConfirmResourceUpdateRequest, opts ...grpc.CallOption) (*ConfirmResourceUpdateResponse, error) {
	out := new(ConfirmResourceUpdateResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/ConfirmResourceUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) RetrieveResource(ctx context.Context, in *RetrieveResourceRequest, opts ...grpc.CallOption) (*RetrieveResourceResponse, error) {
	out := new(RetrieveResourceResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/RetrieveResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceAggregateClient) ConfirmResourceRetrieve(ctx context.Context, in *ConfirmResourceRetrieveRequest, opts ...grpc.CallOption) (*ConfirmResourceRetrieveResponse, error) {
	out := new(ConfirmResourceRetrieveResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/ConfirmResourceRetrieve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceAggregateServer is the server API for ResourceAggregate service.
type ResourceAggregateServer interface {
	PublishResource(context.Context, *PublishResourceRequest) (*PublishResourceResponse, error)
	UnpublishResource(context.Context, *UnpublishResourceRequest) (*UnpublishResourceResponse, error)
	NotifyResourceChanged(context.Context, *NotifyResourceChangedRequest) (*NotifyResourceChangedResponse, error)
	UpdateResource(context.Context, *UpdateResourceRequest) (*UpdateResourceResponse, error)
	ConfirmResourceUpdate(context.Context, *ConfirmResourceUpdateRequest) (*ConfirmResourceUpdateResponse, error)
	RetrieveResource(context.Context, *RetrieveResourceRequest) (*RetrieveResourceResponse, error)
	ConfirmResourceRetrieve(context.Context, *ConfirmResourceRetrieveRequest) (*ConfirmResourceRetrieveResponse, error)
}

// UnimplementedResourceAggregateServer can be embedded to have forward compatible implementations.
type UnimplementedResourceAggregateServer struct {
}

func (*UnimplementedResourceAggregateServer) PublishResource(ctx context.Context, req *PublishResourceRequest) (*PublishResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishResource not implemented")
}
func (*UnimplementedResourceAggregateServer) UnpublishResource(ctx context.Context, req *UnpublishResourceRequest) (*UnpublishResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnpublishResource not implemented")
}
func (*UnimplementedResourceAggregateServer) NotifyResourceChanged(ctx context.Context, req *NotifyResourceChangedRequest) (*NotifyResourceChangedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyResourceChanged not implemented")
}
func (*UnimplementedResourceAggregateServer) UpdateResource(ctx context.Context, req *UpdateResourceRequest) (*UpdateResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateResource not implemented")
}
func (*UnimplementedResourceAggregateServer) ConfirmResourceUpdate(ctx context.Context, req *ConfirmResourceUpdateRequest) (*ConfirmResourceUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmResourceUpdate not implemented")
}
func (*UnimplementedResourceAggregateServer) RetrieveResource(ctx context.Context, req *RetrieveResourceRequest) (*RetrieveResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveResource not implemented")
}
func (*UnimplementedResourceAggregateServer) ConfirmResourceRetrieve(ctx context.Context, req *ConfirmResourceRetrieveRequest) (*ConfirmResourceRetrieveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmResourceRetrieve not implemented")
}

func RegisterResourceAggregateServer(s *grpc.Server, srv ResourceAggregateServer) {
	s.RegisterService(&_ResourceAggregate_serviceDesc, srv)
}

func _ResourceAggregate_PublishResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).PublishResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/PublishResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).PublishResource(ctx, req.(*PublishResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_UnpublishResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnpublishResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).UnpublishResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/UnpublishResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).UnpublishResource(ctx, req.(*UnpublishResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_NotifyResourceChanged_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NotifyResourceChangedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).NotifyResourceChanged(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/NotifyResourceChanged",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).NotifyResourceChanged(ctx, req.(*NotifyResourceChangedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_UpdateResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).UpdateResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/UpdateResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).UpdateResource(ctx, req.(*UpdateResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_ConfirmResourceUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmResourceUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).ConfirmResourceUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/ConfirmResourceUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).ConfirmResourceUpdate(ctx, req.(*ConfirmResourceUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_RetrieveResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).RetrieveResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/RetrieveResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).RetrieveResource(ctx, req.(*RetrieveResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceAggregate_ConfirmResourceRetrieve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmResourceRetrieveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceAggregateServer).ConfirmResourceRetrieve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.resourceaggregate.pb.ResourceAggregate/ConfirmResourceRetrieve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceAggregateServer).ConfirmResourceRetrieve(ctx, req.(*ConfirmResourceRetrieveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResourceAggregate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ocf.cloud.resourceaggregate.pb.ResourceAggregate",
	HandlerType: (*ResourceAggregateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PublishResource",
			Handler:    _ResourceAggregate_PublishResource_Handler,
		},
		{
			MethodName: "UnpublishResource",
			Handler:    _ResourceAggregate_UnpublishResource_Handler,
		},
		{
			MethodName: "NotifyResourceChanged",
			Handler:    _ResourceAggregate_NotifyResourceChanged_Handler,
		},
		{
			MethodName: "UpdateResource",
			Handler:    _ResourceAggregate_UpdateResource_Handler,
		},
		{
			MethodName: "ConfirmResourceUpdate",
			Handler:    _ResourceAggregate_ConfirmResourceUpdate_Handler,
		},
		{
			MethodName: "RetrieveResource",
			Handler:    _ResourceAggregate_RetrieveResource_Handler,
		},
		{
			MethodName: "ConfirmResourceRetrieve",
			Handler:    _ResourceAggregate_ConfirmResourceRetrieve_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/go-ocf/cloud/resource-aggregate/pb/service.proto",
}
