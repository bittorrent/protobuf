// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: grpc/grpc.proto

package testing

import (
	context "context"
	fmt "fmt"
	proto "github.com/bittorrent/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type SimpleRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" pg:"-"`
	XXX_unrecognized     []byte   `json:"-" pg:"-"`
	XXX_sizecache        int32    `json:"-" pg:"-"`
}

func (m *SimpleRequest) Reset()         { *m = SimpleRequest{} }
func (m *SimpleRequest) String() string { return proto.CompactTextString(m) }
func (*SimpleRequest) ProtoMessage()    {}
func (*SimpleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_81ea47a3f88c2082, []int{0}
}
func (m *SimpleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleRequest.Unmarshal(m, b)
}
func (m *SimpleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleRequest.Marshal(b, m, deterministic)
}
func (m *SimpleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleRequest.Merge(m, src)
}
func (m *SimpleRequest) XXX_Size() int {
	return xxx_messageInfo_SimpleRequest.Size(m)
}
func (m *SimpleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleRequest proto.InternalMessageInfo

type SimpleResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" pg:"-"`
	XXX_unrecognized     []byte   `json:"-" pg:"-"`
	XXX_sizecache        int32    `json:"-" pg:"-"`
}

func (m *SimpleResponse) Reset()         { *m = SimpleResponse{} }
func (m *SimpleResponse) String() string { return proto.CompactTextString(m) }
func (*SimpleResponse) ProtoMessage()    {}
func (*SimpleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_81ea47a3f88c2082, []int{1}
}
func (m *SimpleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleResponse.Unmarshal(m, b)
}
func (m *SimpleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleResponse.Marshal(b, m, deterministic)
}
func (m *SimpleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleResponse.Merge(m, src)
}
func (m *SimpleResponse) XXX_Size() int {
	return xxx_messageInfo_SimpleResponse.Size(m)
}
func (m *SimpleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleResponse proto.InternalMessageInfo

type StreamMsg struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" pg:"-"`
	XXX_unrecognized     []byte   `json:"-" pg:"-"`
	XXX_sizecache        int32    `json:"-" pg:"-"`
}

func (m *StreamMsg) Reset()         { *m = StreamMsg{} }
func (m *StreamMsg) String() string { return proto.CompactTextString(m) }
func (*StreamMsg) ProtoMessage()    {}
func (*StreamMsg) Descriptor() ([]byte, []int) {
	return fileDescriptor_81ea47a3f88c2082, []int{2}
}
func (m *StreamMsg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMsg.Unmarshal(m, b)
}
func (m *StreamMsg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMsg.Marshal(b, m, deterministic)
}
func (m *StreamMsg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMsg.Merge(m, src)
}
func (m *StreamMsg) XXX_Size() int {
	return xxx_messageInfo_StreamMsg.Size(m)
}
func (m *StreamMsg) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMsg.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMsg proto.InternalMessageInfo

type StreamMsg2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" pg:"-"`
	XXX_unrecognized     []byte   `json:"-" pg:"-"`
	XXX_sizecache        int32    `json:"-" pg:"-"`
}

func (m *StreamMsg2) Reset()         { *m = StreamMsg2{} }
func (m *StreamMsg2) String() string { return proto.CompactTextString(m) }
func (*StreamMsg2) ProtoMessage()    {}
func (*StreamMsg2) Descriptor() ([]byte, []int) {
	return fileDescriptor_81ea47a3f88c2082, []int{3}
}
func (m *StreamMsg2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamMsg2.Unmarshal(m, b)
}
func (m *StreamMsg2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamMsg2.Marshal(b, m, deterministic)
}
func (m *StreamMsg2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamMsg2.Merge(m, src)
}
func (m *StreamMsg2) XXX_Size() int {
	return xxx_messageInfo_StreamMsg2.Size(m)
}
func (m *StreamMsg2) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamMsg2.DiscardUnknown(m)
}

var xxx_messageInfo_StreamMsg2 proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SimpleRequest)(nil), "grpc.testing.SimpleRequest")
	proto.RegisterType((*SimpleResponse)(nil), "grpc.testing.SimpleResponse")
	proto.RegisterType((*StreamMsg)(nil), "grpc.testing.StreamMsg")
	proto.RegisterType((*StreamMsg2)(nil), "grpc.testing.StreamMsg2")
}

func init() { proto.RegisterFile("grpc/grpc.proto", fileDescriptor_81ea47a3f88c2082) }

var fileDescriptor_81ea47a3f88c2082 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0x15, 0x54, 0x21, 0x7a, 0x14, 0x8a, 0xbc, 0x80, 0x80, 0x29, 0x53, 0x97, 0x26, 0x55,
	0x18, 0x11, 0x4b, 0x1b, 0x75, 0x63, 0xa1, 0x74, 0x61, 0xb3, 0xd3, 0xc3, 0x58, 0x4a, 0x7c, 0xc6,
	0xbe, 0x08, 0xf1, 0x4f, 0xf8, 0xb9, 0xa8, 0x6e, 0x8b, 0x00, 0x35, 0x5d, 0xac, 0xf3, 0xbd, 0xd3,
	0xf7, 0x9e, 0xf4, 0x60, 0xa8, 0xbd, 0xab, 0xf2, 0xf5, 0x93, 0x39, 0x4f, 0x4c, 0x62, 0x10, 0x67,
	0xc6, 0xc0, 0xc6, 0xea, 0x74, 0x08, 0x67, 0x0b, 0xd3, 0xb8, 0x1a, 0x9f, 0xf0, 0xbd, 0xc5, 0xc0,
	0xe9, 0x05, 0x9c, 0xef, 0x16, 0xc1, 0x91, 0x0d, 0x98, 0x9e, 0x42, 0x7f, 0xc1, 0x1e, 0x65, 0xf3,
	0x18, 0x74, 0x3a, 0x00, 0xf8, 0xf9, 0x14, 0xc5, 0xd7, 0x11, 0xf4, 0x9e, 0x31, 0xb0, 0x98, 0x43,
	0x7f, 0x69, 0xa5, 0xff, 0x9c, 0xc9, 0xba, 0x16, 0x37, 0xd9, 0x6f, 0x8b, 0xec, 0x0f, 0xff, 0xfa,
	0x76, 0xbf, 0xb8, 0xf1, 0x12, 0x25, 0x40, 0x49, 0x1f, 0x36, 0x44, 0x8b, 0xc3, 0xa0, 0xcb, 0x7f,
	0xe2, 0x2e, 0xd5, 0x24, 0x11, 0x33, 0x38, 0x59, 0xba, 0x2d, 0xa3, 0xeb, 0xec, 0x70, 0x90, 0x51,
	0x22, 0x1e, 0xa0, 0x37, 0x35, 0x2b, 0xd3, 0x0d, 0xb8, 0xea, 0x10, 0x8a, 0x51, 0x32, 0x49, 0xa6,
	0xf3, 0x97, 0x52, 0x1b, 0x7e, 0x6b, 0x55, 0x56, 0x51, 0x93, 0x2b, 0xc3, 0x4c, 0xde, 0xa3, 0xe5,
	0x3c, 0xb6, 0xa0, 0xda, 0xd7, 0xcd, 0x50, 0x8d, 0x35, 0xda, 0xb1, 0x26, 0x4d, 0xf9, 0x1a, 0xb4,
	0x92, 0x2c, 0x63, 0x59, 0xf7, 0x5b, 0xac, 0x3a, 0x8e, 0x67, 0x77, 0xdf, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x21, 0x4f, 0xf7, 0x41, 0xc8, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TestClient is the client API for Test service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TestClient interface {
	UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error)
	// This RPC streams from the server only.
	Downstream(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (Test_DownstreamClient, error)
	// This RPC streams from the client.
	Upstream(ctx context.Context, opts ...grpc.CallOption) (Test_UpstreamClient, error)
	// This one streams in both directions.
	Bidi(ctx context.Context, opts ...grpc.CallOption) (Test_BidiClient, error)
}

type testClient struct {
	cc *grpc.ClientConn
}

func NewTestClient(cc *grpc.ClientConn) TestClient {
	return &testClient{cc}
}

func (c *testClient) UnaryCall(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (*SimpleResponse, error) {
	out := new(SimpleResponse)
	err := c.cc.Invoke(ctx, "/grpc.testing.Test/UnaryCall", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *testClient) Downstream(ctx context.Context, in *SimpleRequest, opts ...grpc.CallOption) (Test_DownstreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Test_serviceDesc.Streams[0], "/grpc.testing.Test/Downstream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testDownstreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Test_DownstreamClient interface {
	Recv() (*StreamMsg, error)
	grpc.ClientStream
}

type testDownstreamClient struct {
	grpc.ClientStream
}

func (x *testDownstreamClient) Recv() (*StreamMsg, error) {
	m := new(StreamMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testClient) Upstream(ctx context.Context, opts ...grpc.CallOption) (Test_UpstreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Test_serviceDesc.Streams[1], "/grpc.testing.Test/Upstream", opts...)
	if err != nil {
		return nil, err
	}
	x := &testUpstreamClient{stream}
	return x, nil
}

type Test_UpstreamClient interface {
	Send(*StreamMsg) error
	CloseAndRecv() (*SimpleResponse, error)
	grpc.ClientStream
}

type testUpstreamClient struct {
	grpc.ClientStream
}

func (x *testUpstreamClient) Send(m *StreamMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testUpstreamClient) CloseAndRecv() (*SimpleResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(SimpleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *testClient) Bidi(ctx context.Context, opts ...grpc.CallOption) (Test_BidiClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Test_serviceDesc.Streams[2], "/grpc.testing.Test/Bidi", opts...)
	if err != nil {
		return nil, err
	}
	x := &testBidiClient{stream}
	return x, nil
}

type Test_BidiClient interface {
	Send(*StreamMsg) error
	Recv() (*StreamMsg2, error)
	grpc.ClientStream
}

type testBidiClient struct {
	grpc.ClientStream
}

func (x *testBidiClient) Send(m *StreamMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *testBidiClient) Recv() (*StreamMsg2, error) {
	m := new(StreamMsg2)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TestServer is the server API for Test service.
type TestServer interface {
	UnaryCall(context.Context, *SimpleRequest) (*SimpleResponse, error)
	// This RPC streams from the server only.
	Downstream(*SimpleRequest, Test_DownstreamServer) error
	// This RPC streams from the client.
	Upstream(Test_UpstreamServer) error
	// This one streams in both directions.
	Bidi(Test_BidiServer) error
}

// UnimplementedTestServer can be embedded to have forward compatible implementations.
type UnimplementedTestServer struct {
}

func (*UnimplementedTestServer) UnaryCall(ctx context.Context, req *SimpleRequest) (*SimpleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryCall not implemented")
}
func (*UnimplementedTestServer) Downstream(req *SimpleRequest, srv Test_DownstreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Downstream not implemented")
}
func (*UnimplementedTestServer) Upstream(srv Test_UpstreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Upstream not implemented")
}
func (*UnimplementedTestServer) Bidi(srv Test_BidiServer) error {
	return status.Errorf(codes.Unimplemented, "method Bidi not implemented")
}

func RegisterTestServer(s *grpc.Server, srv TestServer) {
	s.RegisterService(&_Test_serviceDesc, srv)
}

func _Test_UnaryCall_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SimpleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TestServer).UnaryCall(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.testing.Test/UnaryCall",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TestServer).UnaryCall(ctx, req.(*SimpleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Test_Downstream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SimpleRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TestServer).Downstream(m, &testDownstreamServer{stream})
}

type Test_DownstreamServer interface {
	Send(*StreamMsg) error
	grpc.ServerStream
}

type testDownstreamServer struct {
	grpc.ServerStream
}

func (x *testDownstreamServer) Send(m *StreamMsg) error {
	return x.ServerStream.SendMsg(m)
}

func _Test_Upstream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServer).Upstream(&testUpstreamServer{stream})
}

type Test_UpstreamServer interface {
	SendAndClose(*SimpleResponse) error
	Recv() (*StreamMsg, error)
	grpc.ServerStream
}

type testUpstreamServer struct {
	grpc.ServerStream
}

func (x *testUpstreamServer) SendAndClose(m *SimpleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testUpstreamServer) Recv() (*StreamMsg, error) {
	m := new(StreamMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Test_Bidi_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TestServer).Bidi(&testBidiServer{stream})
}

type Test_BidiServer interface {
	Send(*StreamMsg2) error
	Recv() (*StreamMsg, error)
	grpc.ServerStream
}

type testBidiServer struct {
	grpc.ServerStream
}

func (x *testBidiServer) Send(m *StreamMsg2) error {
	return x.ServerStream.SendMsg(m)
}

func (x *testBidiServer) Recv() (*StreamMsg, error) {
	m := new(StreamMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Test_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.Test",
	HandlerType: (*TestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryCall",
			Handler:    _Test_UnaryCall_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Downstream",
			Handler:       _Test_Downstream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Upstream",
			Handler:       _Test_Upstream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "Bidi",
			Handler:       _Test_Bidi_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/grpc.proto",
}
