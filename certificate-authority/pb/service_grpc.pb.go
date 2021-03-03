// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// CertificateAuthorityClient is the client API for CertificateAuthority service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
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
// All implementations must embed UnimplementedCertificateAuthorityServer
// for forward compatibility
type CertificateAuthorityServer interface {
	// SignIdentityCertificate sends a Identity Certificate Signing Request to the certificate authority
	// and obtains a signed certificate. Both in the PEM format. It adds EKU: '1.3.6.1.4.1.44924.1.6' .
	SignIdentityCertificate(context.Context, *SignCertificateRequest) (*SignCertificateResponse, error)
	// SignCertificate sends a Certificate Signing Request to the certificate authority
	// and obtains a signed certificate. Both in the PEM format.
	SignCertificate(context.Context, *SignCertificateRequest) (*SignCertificateResponse, error)
	mustEmbedUnimplementedCertificateAuthorityServer()
}

// UnimplementedCertificateAuthorityServer must be embedded to have forward compatible implementations.
type UnimplementedCertificateAuthorityServer struct {
}

func (UnimplementedCertificateAuthorityServer) SignIdentityCertificate(context.Context, *SignCertificateRequest) (*SignCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIdentityCertificate not implemented")
}
func (UnimplementedCertificateAuthorityServer) SignCertificate(context.Context, *SignCertificateRequest) (*SignCertificateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignCertificate not implemented")
}
func (UnimplementedCertificateAuthorityServer) mustEmbedUnimplementedCertificateAuthorityServer() {}

// UnsafeCertificateAuthorityServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CertificateAuthorityServer will
// result in compilation errors.
type UnsafeCertificateAuthorityServer interface {
	mustEmbedUnimplementedCertificateAuthorityServer()
}

func RegisterCertificateAuthorityServer(s grpc.ServiceRegistrar, srv CertificateAuthorityServer) {
	s.RegisterService(&CertificateAuthority_ServiceDesc, srv)
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

// CertificateAuthority_ServiceDesc is the grpc.ServiceDesc for CertificateAuthority service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CertificateAuthority_ServiceDesc = grpc.ServiceDesc{
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
	Metadata: "github.com/plgd-dev/cloud/certificate-authority/pb/service.proto",
}
