// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: enumprefix.proto

package enumprefix

import (
	fmt "fmt"
	_ "github.com/bittorrent/protobuf/gogoproto"
	proto "github.com/bittorrent/protobuf/proto"
	test "github.com/bittorrent/protobuf/test"
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

type MyMessage struct {
	TheField             test.TheTestEnum `protobuf:"varint,1,opt,name=TheField,enum=test.TheTestEnum" json:"TheField" pg:"TheField"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-" pg:"-"`
	XXX_unrecognized     []byte           `json:"-" pg:"-"`
	XXX_sizecache        int32            `json:"-" pg:"-"`
}

func (m *MyMessage) Reset()         { *m = MyMessage{} }
func (m *MyMessage) String() string { return proto.CompactTextString(m) }
func (*MyMessage) ProtoMessage()    {}
func (*MyMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_d0d23e6cb4323eb3, []int{0}
}
func (m *MyMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyMessage.Unmarshal(m, b)
}
func (m *MyMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MyMessage.Marshal(b, m, deterministic)
}
func (m *MyMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyMessage.Merge(m, src)
}
func (m *MyMessage) XXX_Size() int {
	return xxx_messageInfo_MyMessage.Size(m)
}
func (m *MyMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_MyMessage.DiscardUnknown(m)
}

var xxx_messageInfo_MyMessage proto.InternalMessageInfo

func (m *MyMessage) GetTheField() test.TheTestEnum {
	if m != nil {
		return m.TheField
	}
	return test.A
}

func init() {
	proto.RegisterType((*MyMessage)(nil), "enumprefix.MyMessage")
}

func init() { proto.RegisterFile("enumprefix.proto", fileDescriptor_d0d23e6cb4323eb3) }

var fileDescriptor_d0d23e6cb4323eb3 = []byte{
	// 158 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x48, 0xcd, 0x2b, 0xcd,
	0x2d, 0x28, 0x4a, 0x4d, 0xcb, 0xac, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x42, 0x88,
	0x48, 0x19, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x27, 0x65, 0x96,
	0x94, 0xe4, 0x17, 0x15, 0xa5, 0xe6, 0x95, 0xe8, 0x83, 0x15, 0x26, 0x95, 0xa6, 0xe9, 0x97, 0xa4,
	0x16, 0x97, 0xe8, 0x97, 0x64, 0xa4, 0x82, 0x68, 0x88, 0x76, 0x29, 0x63, 0x02, 0x5a, 0xd2, 0xf3,
	0xd3, 0xf3, 0xc1, 0x1c, 0x30, 0x0b, 0xa2, 0x49, 0xc9, 0x81, 0x8b, 0xd3, 0xb7, 0xd2, 0x37, 0xb5,
	0xb8, 0x38, 0x31, 0x3d, 0x55, 0xc8, 0x98, 0x8b, 0x23, 0x24, 0x23, 0xd5, 0x2d, 0x33, 0x35, 0x27,
	0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0xcf, 0x48, 0x50, 0x0f, 0x6c, 0x41, 0x48, 0x46, 0x6a, 0x48,
	0x6a, 0x71, 0x89, 0x6b, 0x5e, 0x69, 0xae, 0x13, 0xcb, 0x89, 0x7b, 0xf2, 0x0c, 0x41, 0x70, 0x85,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x1a, 0x91, 0x9a, 0xc8, 0x00, 0x00, 0x00,
}
