// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/go-ocf/cloud/certificate-authority/pb/service.proto

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
	proto.RegisterFile("github.com/go-ocf/cloud/certificate-authority/pb/service.proto", fileDescriptor_c420cbb2c19797ed)
}

var fileDescriptor_c420cbb2c19797ed = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xb2, 0x4b, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0xd7, 0xcd, 0x4f, 0x4e, 0xd3, 0x4f, 0xce,
	0xc9, 0x2f, 0x4d, 0xd1, 0x4f, 0x4e, 0x2d, 0x2a, 0xc9, 0x4c, 0xcb, 0x4c, 0x4e, 0x2c, 0x49, 0xd5,
	0x4d, 0x2c, 0x2d, 0xc9, 0xc8, 0x2f, 0xca, 0x2c, 0xa9, 0xd4, 0x2f, 0x48, 0xd2, 0x2f, 0x4e, 0x2d,
	0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x52, 0x04, 0xea, 0xd0, 0x03,
	0xeb, 0xd0, 0x43, 0xd2, 0x01, 0xd7, 0xa0, 0x57, 0x90, 0x24, 0x65, 0x4d, 0xb2, 0x15, 0x20, 0x09,
	0x88, 0xf9, 0x46, 0x3b, 0x98, 0xb8, 0x44, 0x9c, 0x11, 0xea, 0x1c, 0x61, 0xca, 0x84, 0x26, 0x31,
	0x72, 0x89, 0x07, 0x67, 0xa6, 0xe7, 0x79, 0xa6, 0xa4, 0xe6, 0x95, 0x00, 0x05, 0x90, 0x14, 0x09,
	0x59, 0xea, 0x11, 0x74, 0x95, 0x1e, 0x48, 0x2f, 0x92, 0x9e, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2,
	0x12, 0x29, 0x2b, 0x72, 0xb4, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0xa7, 0x2a, 0x31, 0x08, 0x75, 0x31,
	0x72, 0xf1, 0xa3, 0xc9, 0x0e, 0x98, 0x63, 0x9c, 0x4c, 0xa3, 0x8c, 0x49, 0x0d, 0x79, 0xeb, 0x82,
	0xa4, 0x24, 0x36, 0x70, 0xc0, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x4f, 0x78, 0xa9,
	0x1a, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CertificateAuthorityClient is the client API for CertificateAuthority service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CertificateAuthorityClient interface {
	// SignIdentityCertificate sends a Identity Certificate Signing Request to the certificate authority
	// and obtains a signed certificate. Both in the PEM format. It adds EKU: '1.3.6.1.4.1.44924.1.6' .
	SignIdentityCertificate(ctx context.Context, in *SignCertificateRequest, opts ...grpc.CallOption) (*SignCertificateResponse, error)
	// SignCertificate sends a Certificate Signing Request to the certificate authority
	// and obtains a signed certificate. Both in the PEM format.
	SignCertificate(ctx context.Context, in *SignCertificateRequest, opts ...grpc.CallOption) (*SignCertificateResponse, error)
}

type certificateAuthorityClient struct {
	cc grpc.ClientConnInterface
}

func NewCertificateAuthorityClient(cc grpc.ClientConnInterface) CertificateAuthorityClient {
	return &certificateAuthorityClient{cc}
}

func (c *certificateAuthorityClient) SignIdentityCertificate(ctx context.Context, in *SignCertificateRequest, opts ...grpc.CallOption) (*SignCertificateResponse, error) {
	out := new(SignCertificateResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.certificateauthority.pb.CertificateAuthority/SignIdentityCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certificateAuthorityClient) SignCertificate(ctx context.Context, in *SignCertificateRequest, opts ...grpc.CallOption) (*SignCertificateResponse, error) {
	out := new(SignCertificateResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.certificateauthority.pb.CertificateAuthority/SignCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertificateAuthorityServer is the server API for CertificateAuthority service.
type CertificateAuthorityServer interface {
	// SignIdentityCertificate sends a Identity Certificate Signing Request to the certificate authority
	// and obtains a signed certificate. Both in the PEM format. It adds EKU: '1.3.6.1.4.1.44924.1.6' .
	SignIdentityCertificate(context.Context, *SignCertificateRequest) (*SignCertificateResponse, error)
	// SignCertificate sends a Certificate Signing Request to the certificate authority
	// and obtains a signed certificate. Both in the PEM format.
	SignCertificate(context.Context, *SignCertificateRequest) (*SignCertificateResponse, error)
}

// UnimplementedCertificateAuthorityServer can be embedded to have forward compatible implementations.
type UnimplementedCertificateAuthorityServer struct {
}

func (*UnimplementedCertificateAuthorityServer) SignIdentityCertificate(ctx context.Context, req *SignCertificateRequest) (*SignCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIdentityCertificate not implemented")
}
func (*UnimplementedCertificateAuthorityServer) SignCertificate(ctx context.Context, req *SignCertificateRequest) (*SignCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignCertificate not implemented")
}

func RegisterCertificateAuthorityServer(s *grpc.Server, srv CertificateAuthorityServer) {
	s.RegisterService(&_CertificateAuthority_serviceDesc, srv)
}

func _CertificateAuthority_SignIdentityCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServer).SignIdentityCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.certificateauthority.pb.CertificateAuthority/SignIdentityCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServer).SignIdentityCertificate(ctx, req.(*SignCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CertificateAuthority_SignCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertificateAuthorityServer).SignCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.certificateauthority.pb.CertificateAuthority/SignCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertificateAuthorityServer).SignCertificate(ctx, req.(*SignCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CertificateAuthority_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ocf.cloud.certificateauthority.pb.CertificateAuthority",
	HandlerType: (*CertificateAuthorityServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SignIdentityCertificate",
			Handler:    _CertificateAuthority_SignIdentityCertificate_Handler,
		},
		{
			MethodName: "SignCertificate",
			Handler:    _CertificateAuthority_SignCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/go-ocf/cloud/certificate-authority/pb/service.proto",
}
