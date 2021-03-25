// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stub.proto

package grpc

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TestRequest struct {
	Test string `protobuf:"bytes,1,opt,name=test,proto3" json:"test,omitempty"`
}

func (m *TestRequest) Reset()         { *m = TestRequest{} }
func (m *TestRequest) String() string { return proto.CompactTextString(m) }
func (*TestRequest) ProtoMessage()    {}
func (*TestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f6c10bbcad5131e, []int{0}
}
func (m *TestRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestRequest.Merge(m, src)
}
func (m *TestRequest) XXX_Size() int {
	return m.Size()
}
func (m *TestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TestRequest proto.InternalMessageInfo

func (m *TestRequest) GetTest() string {
	if m != nil {
		return m.Test
	}
	return ""
}

type TestResponse struct {
	Test string `protobuf:"bytes,1,opt,name=test,proto3" json:"test,omitempty"`
}

func (m *TestResponse) Reset()         { *m = TestResponse{} }
func (m *TestResponse) String() string { return proto.CompactTextString(m) }
func (*TestResponse) ProtoMessage()    {}
func (*TestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f6c10bbcad5131e, []int{1}
}
func (m *TestResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TestResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestResponse.Merge(m, src)
}
func (m *TestResponse) XXX_Size() int {
	return m.Size()
}
func (m *TestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TestResponse proto.InternalMessageInfo

func (m *TestResponse) GetTest() string {
	if m != nil {
		return m.Test
	}
	return ""
}

func init() {
	proto.RegisterType((*TestRequest)(nil), "ocf.cloud.test.pb.TestRequest")
	proto.RegisterType((*TestResponse)(nil), "ocf.cloud.test.pb.TestResponse")
}

func init() { proto.RegisterFile("stub.proto", fileDescriptor_9f6c10bbcad5131e) }

var fileDescriptor_9f6c10bbcad5131e = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x50, 0x31, 0x4e, 0xc4, 0x30,
	0x10, 0x8c, 0x25, 0x84, 0x60, 0x8f, 0x06, 0x57, 0xe8, 0x0a, 0x03, 0xa9, 0xae, 0x39, 0x1b, 0x41,
	0x43, 0x0d, 0x35, 0xcd, 0x85, 0x8a, 0xee, 0x6c, 0x36, 0x21, 0x22, 0x77, 0x1b, 0xec, 0x35, 0xef,
	0xe0, 0x2b, 0xfc, 0x82, 0x32, 0x25, 0x25, 0x4a, 0x3e, 0x82, 0xe2, 0x34, 0x20, 0xa0, 0xa3, 0x9b,
	0xd9, 0x9d, 0x9d, 0x59, 0x0d, 0x40, 0xe0, 0x68, 0x75, 0xeb, 0x89, 0x49, 0x1e, 0x92, 0x2b, 0xb5,
	0x6b, 0x28, 0xde, 0x6b, 0xc6, 0xc0, 0xba, 0xb5, 0xf3, 0x65, 0x55, 0xf3, 0x43, 0xb4, 0xda, 0xd1,
	0xc6, 0x54, 0x54, 0x91, 0x49, 0x4a, 0x1b, 0xcb, 0xc4, 0x12, 0x49, 0x68, 0x72, 0xc8, 0x4f, 0x61,
	0x76, 0x8b, 0x81, 0x57, 0xf8, 0x14, 0x31, 0xb0, 0x94, 0xb0, 0x33, 0x1a, 0x1d, 0x89, 0x13, 0xb1,
	0xd8, 0x5f, 0x25, 0x9c, 0xe7, 0x70, 0x30, 0x49, 0x42, 0x4b, 0xdb, 0x80, 0xbf, 0x69, 0xce, 0x5f,
	0x05, 0xcc, 0x0a, 0x8e, 0xb6, 0x40, 0xff, 0x5c, 0x3b, 0x94, 0x37, 0xb0, 0x37, 0xde, 0x5c, 0xaf,
	0x9b, 0x46, 0x2a, 0xfd, 0xe3, 0x4b, 0xfd, 0x25, 0x73, 0x7e, 0xfc, 0xe7, 0x7e, 0x0a, 0xcc, 0x33,
	0x59, 0x00, 0x8c, 0x93, 0x82, 0x3d, 0xae, 0x37, 0xff, 0x60, 0xb8, 0x10, 0x67, 0xe2, 0xea, 0xf2,
	0xad, 0x57, 0xa2, 0xeb, 0x95, 0xf8, 0xe8, 0x95, 0x78, 0x19, 0x54, 0xd6, 0x0d, 0x2a, 0x7b, 0x1f,
	0x54, 0x76, 0xa7, 0xbe, 0x75, 0xb8, 0x24, 0x57, 0x9a, 0xc7, 0x9a, 0xcd, 0x16, 0xd9, 0x54, 0xbe,
	0x75, 0x76, 0x37, 0x75, 0x77, 0xf1, 0x19, 0x00, 0x00, 0xff, 0xff, 0x28, 0x67, 0x6b, 0x58, 0x8b,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StubServiceClient is the client API for StubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StubServiceClient interface {
	TestCall(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error)
	TestStream(ctx context.Context, opts ...grpc.CallOption) (StubService_TestStreamClient, error)
}

type stubServiceClient struct {
	cc *grpc.ClientConn
}

func NewStubServiceClient(cc *grpc.ClientConn) StubServiceClient {
	return &stubServiceClient{cc}
}

func (c *stubServiceClient) TestCall(ctx context.Context, in *TestRequest, opts ...grpc.CallOption) (*TestResponse, error) {
	out := new(TestResponse)
	err := c.cc.Invoke(ctx, "/ocf.cloud.test.pb.StubService/TestCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stubServiceClient) TestStream(ctx context.Context, opts ...grpc.CallOption) (StubService_TestStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_StubService_serviceDesc.Streams[0], "/ocf.cloud.test.pb.StubService/TestStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &stubServiceTestStreamClient{stream}
	return x, nil
}

type StubService_TestStreamClient interface {
	Send(*TestRequest) error
	Recv() (*TestResponse, error)
	grpc.ClientStream
}

type stubServiceTestStreamClient struct {
	grpc.ClientStream
}

func (x *stubServiceTestStreamClient) Send(m *TestRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *stubServiceTestStreamClient) Recv() (*TestResponse, error) {
	m := new(TestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StubServiceServer is the server API for StubService service.
type StubServiceServer interface {
	TestCall(context.Context, *TestRequest) (*TestResponse, error)
	TestStream(StubService_TestStreamServer) error
}

// UnimplementedStubServiceServer can be embedded to have forward compatible implementations.
type UnimplementedStubServiceServer struct {
}

func (*UnimplementedStubServiceServer) TestCall(ctx context.Context, req *TestRequest) (*TestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TestCall not implemented")
}
func (*UnimplementedStubServiceServer) TestStream(srv StubService_TestStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method TestStream not implemented")
}

func RegisterStubServiceServer(s *grpc.Server, srv StubServiceServer) {
	s.RegisterService(&_StubService_serviceDesc, srv)
}

func _StubService_TestCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StubServiceServer).TestCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ocf.cloud.test.pb.StubService/TestCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StubServiceServer).TestCall(ctx, req.(*TestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StubService_TestStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StubServiceServer).TestStream(&stubServiceTestStreamServer{stream})
}

type StubService_TestStreamServer interface {
	Send(*TestResponse) error
	Recv() (*TestRequest, error)
	grpc.ServerStream
}

type stubServiceTestStreamServer struct {
	grpc.ServerStream
}

func (x *stubServiceTestStreamServer) Send(m *TestResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *stubServiceTestStreamServer) Recv() (*TestRequest, error) {
	m := new(TestRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _StubService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ocf.cloud.test.pb.StubService",
	HandlerType: (*StubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TestCall",
			Handler:    _StubService_TestCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TestStream",
			Handler:       _StubService_TestStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "stub.proto",
}

func (m *TestRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Test) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStub(dAtA, i, uint64(len(m.Test)))
		i += copy(dAtA[i:], m.Test)
	}
	return i, nil
}

func (m *TestResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TestResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Test) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintStub(dAtA, i, uint64(len(m.Test)))
		i += copy(dAtA[i:], m.Test)
	}
	return i, nil
}

func encodeVarintStub(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TestRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Test)
	if l > 0 {
		n += 1 + l + sovStub(uint64(l))
	}
	return n
}

func (m *TestResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Test)
	if l > 0 {
		n += 1 + l + sovStub(uint64(l))
	}
	return n
}

func sovStub(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStub(x uint64) (n int) {
	return sovStub(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TestRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStub
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TestRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Test", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStub
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStub
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStub
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Test = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStub(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStub
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStub
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TestResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStub
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TestResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TestResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Test", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStub
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthStub
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStub
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Test = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStub(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStub
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthStub
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipStub(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStub
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStub
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowStub
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthStub
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthStub
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStub
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipStub(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthStub
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthStub = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStub   = fmt.Errorf("proto: integer overflow")
)
