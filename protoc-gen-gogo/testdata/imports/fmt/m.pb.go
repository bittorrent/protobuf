// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: imports/fmt/m.proto

package fmt

import (
	fmt "fmt"
	proto "github.com/bittorrent/protobuf/proto"
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

type M struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" pg:"-"`
	XXX_unrecognized     []byte   `json:"-" pg:"-"`
	XXX_sizecache        int32    `json:"-" pg:"-"`
}

func (m *M) Reset()         { *m = M{} }
func (m *M) String() string { return proto.CompactTextString(m) }
func (*M) ProtoMessage()    {}
func (*M) Descriptor() ([]byte, []int) {
	return fileDescriptor_72c126fcd452e392, []int{0}
}
func (m *M) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M.Unmarshal(m, b)
}
func (m *M) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M.Marshal(b, m, deterministic)
}
func (m *M) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M.Merge(m, src)
}
func (m *M) XXX_Size() int {
	return xxx_messageInfo_M.Size(m)
}
func (m *M) XXX_DiscardUnknown() {
	xxx_messageInfo_M.DiscardUnknown(m)
}

var xxx_messageInfo_M proto.InternalMessageInfo

func init() {
	proto.RegisterType((*M)(nil), "fmt.M")
}

func init() { proto.RegisterFile("imports/fmt/m.proto", fileDescriptor_72c126fcd452e392) }

var fileDescriptor_72c126fcd452e392 = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0x29, 0xd6, 0x4f, 0xcb, 0x2d, 0xd1, 0xcf, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x4e, 0xcb, 0x2d, 0x51, 0x62, 0xe6, 0x62, 0xf4, 0x75, 0x72, 0x8d, 0x72, 0x4e, 0xcf, 0x2c,
	0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xca, 0x2c, 0x29, 0xc9, 0x2f, 0x2a, 0x4a,
	0xcd, 0x2b, 0xd1, 0x07, 0x2b, 0x4c, 0x2a, 0x4d, 0x83, 0x30, 0x92, 0x75, 0xd3, 0x53, 0xf3, 0x74,
	0xd3, 0xf3, 0xd3, 0xf3, 0xf5, 0x4b, 0x52, 0x8b, 0x4b, 0x52, 0x12, 0x4b, 0x12, 0xf5, 0x91, 0x0c,
	0x4e, 0x62, 0x03, 0xab, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x53, 0x2e, 0x46, 0x1c, 0x6e,
	0x00, 0x00, 0x00,
}
