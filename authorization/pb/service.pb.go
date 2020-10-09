// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/plgd-dev/cloud/authorization/pb/service.proto

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
	proto.RegisterFile("github.com/plgd-dev/cloud/authorization/pb/service.proto", fileDescriptor_7d55a688f473dfed)
}

var fileDescriptor_7d55a688f473dfed = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xe7, 0xc1, 0x89, 0x41, 0x04, 0x83, 0xa7, 0x81, 0xa0, 0x55, 0xd4, 0xcb, 0x12, 0xd1,
	0x8b, 0xe0, 0x69, 0x22, 0xc8, 0x4e, 0x83, 0xce, 0x81, 0x08, 0x1e, 0xfa, 0xe3, 0xf5, 0x07, 0x6e,
	0x49, 0x6c, 0x92, 0x1e, 0xfc, 0x67, 0xfd, 0x57, 0x5c, 0xd6, 0xb5, 0xeb, 0x58, 0xda, 0x7a, 0x0b,
	0x79, 0x9f, 0xf7, 0x79, 0x5f, 0x5e, 0x08, 0x7a, 0x8c, 0x53, 0x95, 0x68, 0x9f, 0x04, 0x7c, 0x41,
	0xc5, 0x3c, 0x0e, 0x87, 0x21, 0xe4, 0x34, 0x98, 0x73, 0x1d, 0x52, 0x4f, 0xab, 0x84, 0x67, 0xe9,
	0x8f, 0xa7, 0x52, 0xce, 0xa8, 0xf0, 0xa9, 0x84, 0x2c, 0x4f, 0x03, 0x20, 0x22, 0xe3, 0x8a, 0xe3,
	0x13, 0x1e, 0x44, 0x64, 0x45, 0x12, 0x43, 0x12, 0xe1, 0x0f, 0x50, 0x71, 0x30, 0xe5, 0xfb, 0xdf,
	0x7d, 0x74, 0x3a, 0xaa, 0x1b, 0xa6, 0x45, 0x37, 0x9e, 0xa0, 0xfe, 0x34, 0x8d, 0xd9, 0x4c, 0xe0,
	0x73, 0xb2, 0xa3, 0x20, 0x45, 0xc9, 0x85, 0x6f, 0x0d, 0x52, 0x0d, 0x2e, 0x5a, 0x08, 0x29, 0x38,
	0x93, 0xe0, 0xf4, 0xb0, 0x8b, 0x0e, 0xcc, 0xdd, 0x24, 0x8a, 0x70, 0x13, 0xbf, 0xac, 0x95, 0x4a,
	0xa7, 0x0d, 0xa9, 0x9c, 0xeb, 0x90, 0x63, 0xd6, 0x18, 0x72, 0xcc, 0xba, 0x42, 0x1a, 0x62, 0x27,
	0xa4, 0x56, 0xcd, 0x21, 0xb5, 0xea, 0x0c, 0x69, 0x90, 0xca, 0xe9, 0xa1, 0x23, 0x17, 0xa2, 0x0c,
	0x64, 0xf2, 0xc6, 0xbf, 0x80, 0xe1, 0x6b, 0x4b, 0x57, 0x1d, 0x28, 0xed, 0x37, 0x9d, 0x5c, 0x35,
	0xe2, 0x13, 0x1d, 0xbf, 0x82, 0x9a, 0x2d, 0x5f, 0xfe, 0x05, 0xcc, 0xeb, 0x49, 0x7c, 0x6b, 0x69,
	0xde, 0x46, 0xca, 0x31, 0x67, 0x16, 0x72, 0x83, 0x39, 0xbd, 0xbb, 0x3d, 0xfc, 0x8e, 0x0e, 0x47,
	0x61, 0x58, 0x5c, 0xe0, 0x4b, 0x0b, 0x5f, 0x55, 0x4b, 0xe9, 0x55, 0x3b, 0xb4, 0xbd, 0x9b, 0x05,
	0xcf, 0x61, 0x2d, 0xb7, 0xef, 0x66, 0x03, 0xb4, 0xef, 0xa6, 0xce, 0x95, 0x23, 0x9e, 0xe9, 0xc7,
	0xf0, 0xff, 0x9f, 0xe7, 0x49, 0xf8, 0x7e, 0x7f, 0xf5, 0x33, 0x1e, 0xfe, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x46, 0xba, 0x4a, 0xd2, 0x74, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AuthorizationServiceClient is the client API for AuthorizationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizationServiceClient interface {
	SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error)
	SignOff(ctx context.Context, in *SignOffRequest, opts ...grpc.CallOption) (*SignOffResponse, error)
	SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error)
	SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutResponse, error)
	RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error)
	GetUserDevices(ctx context.Context, in *GetUserDevicesRequest, opts ...grpc.CallOption) (AuthorizationService_GetUserDevicesClient, error)
	AddDevice(ctx context.Context, in *AddDeviceRequest, opts ...grpc.CallOption) (*AddDeviceResponse, error)
	RemoveDevice(ctx context.Context, in *RemoveDeviceRequest, opts ...grpc.CallOption) (*RemoveDeviceResponse, error)
}

type authorizationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizationServiceClient(cc grpc.ClientConnInterface) AuthorizationServiceClient {
	return &authorizationServiceClient{cc}
}

func (c *authorizationServiceClient) SignUp(ctx context.Context, in *SignUpRequest, opts ...grpc.CallOption) (*SignUpResponse, error) {
	out := new(SignUpResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.auth.pb.AuthorizationService/SignUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) SignOff(ctx context.Context, in *SignOffRequest, opts ...grpc.CallOption) (*SignOffResponse, error) {
	out := new(SignOffResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.auth.pb.AuthorizationService/SignOff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) SignIn(ctx context.Context, in *SignInRequest, opts ...grpc.CallOption) (*SignInResponse, error) {
	out := new(SignInResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.auth.pb.AuthorizationService/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) SignOut(ctx context.Context, in *SignOutRequest, opts ...grpc.CallOption) (*SignOutResponse, error) {
	out := new(SignOutResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.auth.pb.AuthorizationService/SignOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...grpc.CallOption) (*RefreshTokenResponse, error) {
	out := new(RefreshTokenResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.auth.pb.AuthorizationService/RefreshToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) GetUserDevices(ctx context.Context, in *GetUserDevicesRequest, opts ...grpc.CallOption) (AuthorizationService_GetUserDevicesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AuthorizationService_serviceDesc.Streams[0], "/ocf.cloud.auth.pb.AuthorizationService/GetUserDevices", opts...)
	if err != nil {
		return nil, err
	}
	x := &authorizationServiceGetUserDevicesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AuthorizationService_GetUserDevicesClient interface {
	Recv() (*UserDevice, error)
	grpc.ClientStream
}

type authorizationServiceGetUserDevicesClient struct {
	grpc.ClientStream
}

func (x *authorizationServiceGetUserDevicesClient) Recv() (*UserDevice, error) {
	m := new(UserDevice)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *authorizationServiceClient) AddDevice(ctx context.Context, in *AddDeviceRequest, opts ...grpc.CallOption) (*AddDeviceResponse, error) {
	out := new(AddDeviceResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.auth.pb.AuthorizationService/AddDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationServiceClient) RemoveDevice(ctx context.Context, in *RemoveDeviceRequest, opts ...grpc.CallOption) (*RemoveDeviceResponse, error) {
	out := new(RemoveDeviceResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.auth.pb.AuthorizationService/RemoveDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServiceServer is the server API for AuthorizationService service.
type AuthorizationServiceServer interface {
	SignUp(context.Context, *SignUpRequest) (*SignUpResponse, error)
	SignOff(context.Context, *SignOffRequest) (*SignOffResponse, error)
	SignIn(context.Context, *SignInRequest) (*SignInResponse, error)
	SignOut(context.Context, *SignOutRequest) (*SignOutResponse, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenResponse, error)
	GetUserDevices(*GetUserDevicesRequest, AuthorizationService_GetUserDevicesServer) error
	AddDevice(context.Context, *AddDeviceRequest) (*AddDeviceResponse, error)
	RemoveDevice(context.Context, *RemoveDeviceRequest) (*RemoveDeviceResponse, error)
}

// UnimplementedAuthorizationServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServiceServer struct {
}

func (*UnimplementedAuthorizationServiceServer) SignUp(ctx context.Context, req *SignUpRequest) (*SignUpResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignUp not implemented")
}
func (*UnimplementedAuthorizationServiceServer) SignOff(ctx context.Context, req *SignOffRequest) (*SignOffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOff not implemented")
}
func (*UnimplementedAuthorizationServiceServer) SignIn(ctx context.Context, req *SignInRequest) (*SignInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (*UnimplementedAuthorizationServiceServer) SignOut(ctx context.Context, req *SignOutRequest) (*SignOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignOut not implemented")
}
func (*UnimplementedAuthorizationServiceServer) RefreshToken(ctx context.Context, req *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshToken not implemented")
}
func (*UnimplementedAuthorizationServiceServer) GetUserDevices(req *GetUserDevicesRequest, srv AuthorizationService_GetUserDevicesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetUserDevices not implemented")
}
func (*UnimplementedAuthorizationServiceServer) AddDevice(ctx context.Context, req *AddDeviceRequest) (*AddDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddDevice not implemented")
}
func (*UnimplementedAuthorizationServiceServer) RemoveDevice(ctx context.Context, req *RemoveDeviceRequest) (*RemoveDeviceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveDevice not implemented")
}

func RegisterAuthorizationServiceServer(s *grpc.Server, srv AuthorizationServiceServer) {
	s.RegisterService(&_AuthorizationService_serviceDesc, srv)
}

func _AuthorizationService_SignUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).SignUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.auth.pb.AuthorizationService/SignUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).SignUp(ctx, req.(*SignUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_SignOff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignOffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).SignOff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.auth.pb.AuthorizationService/SignOff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).SignOff(ctx, req.(*SignOffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.auth.pb.AuthorizationService/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).SignIn(ctx, req.(*SignInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_SignOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).SignOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.auth.pb.AuthorizationService/SignOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).SignOut(ctx, req.(*SignOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_RefreshToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).RefreshToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.auth.pb.AuthorizationService/RefreshToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).RefreshToken(ctx, req.(*RefreshTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_GetUserDevices_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetUserDevicesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AuthorizationServiceServer).GetUserDevices(m, &authorizationServiceGetUserDevicesServer{stream})
}

type AuthorizationService_GetUserDevicesServer interface {
	Send(*UserDevice) error
	grpc.ServerStream
}

type authorizationServiceGetUserDevicesServer struct {
	grpc.ServerStream
}

func (x *authorizationServiceGetUserDevicesServer) Send(m *UserDevice) error {
	return x.ServerStream.SendMsg(m)
}

func _AuthorizationService_AddDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).AddDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.auth.pb.AuthorizationService/AddDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).AddDevice(ctx, req.(*AddDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthorizationService_RemoveDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveDeviceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServiceServer).RemoveDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.auth.pb.AuthorizationService/RemoveDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServiceServer).RemoveDevice(ctx, req.(*RemoveDeviceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AuthorizationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ocf.cloud.auth.pb.AuthorizationService",
	HandlerType: (*AuthorizationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignUp",
			Handler:    _AuthorizationService_SignUp_Handler,
		},
		{
			MethodName: "SignOff",
			Handler:    _AuthorizationService_SignOff_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _AuthorizationService_SignIn_Handler,
		},
		{
			MethodName: "SignOut",
			Handler:    _AuthorizationService_SignOut_Handler,
		},
		{
			MethodName: "RefreshToken",
			Handler:    _AuthorizationService_RefreshToken_Handler,
		},
		{
			MethodName: "AddDevice",
			Handler:    _AuthorizationService_AddDevice_Handler,
		},
		{
			MethodName: "RemoveDevice",
			Handler:    _AuthorizationService_RemoveDevice_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetUserDevices",
			Handler:       _AuthorizationService_GetUserDevices_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/plgd-dev/cloud/authorization/pb/service.proto",
}
