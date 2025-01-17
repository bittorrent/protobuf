// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: enumcustomname.proto

// Package enumcustomname tests the behavior of enum_customname and
// enumvalue_customname extensions.

package enumcustomname

import (
	fmt "fmt"
	_ "github.com/bittorrent/protobuf/gogoproto"
	proto "github.com/bittorrent/protobuf/proto"
	test "github.com/bittorrent/protobuf/test"
	math "math"
	strconv "strconv"
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

type MyCustomEnum int32

const (
	// The following field will take on the custom name and the prefix, joined
	// by an underscore.
	MyCustomEnum_MyBetterNameA MyCustomEnum = 0
	MyCustomEnum_B             MyCustomEnum = 1
)

var MyCustomEnum_name = map[int32]string{
	0: "A",
	1: "B",
}

var MyCustomEnum_value = map[string]int32{
	"A": 0,
	"B": 1,
}

func (x MyCustomEnum) Enum() *MyCustomEnum {
	p := new(MyCustomEnum)
	*p = x
	return p
}

func (x MyCustomEnum) String() string {
	return proto.EnumName(MyCustomEnum_name, int32(x))
}

func (x *MyCustomEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MyCustomEnum_value, data, "MyCustomEnum")
	if err != nil {
		return err
	}
	*x = MyCustomEnum(value)
	return nil
}

func (MyCustomEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_49eed3c955d68b51, []int{0}
}

type MyCustomUnprefixedEnum int32

const (
	MyBetterNameUnprefixedA MyCustomUnprefixedEnum = 0
	UNPREFIXED_B            MyCustomUnprefixedEnum = 1
)

var MyCustomUnprefixedEnum_name = map[int32]string{
	0: "UNPREFIXED_A",
	1: "UNPREFIXED_B",
}

var MyCustomUnprefixedEnum_value = map[string]int32{
	"UNPREFIXED_A": 0,
	"UNPREFIXED_B": 1,
}

func (x MyCustomUnprefixedEnum) Enum() *MyCustomUnprefixedEnum {
	p := new(MyCustomUnprefixedEnum)
	*p = x
	return p
}

func (x MyCustomUnprefixedEnum) MarshalJSON() ([]byte, error) {
	return proto.MarshalJSONEnum(MyCustomUnprefixedEnum_name, int32(x))
}

func (x *MyCustomUnprefixedEnum) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MyCustomUnprefixedEnum_value, data, "MyCustomUnprefixedEnum")
	if err != nil {
		return err
	}
	*x = MyCustomUnprefixedEnum(value)
	return nil
}

func (MyCustomUnprefixedEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_49eed3c955d68b51, []int{1}
}

type MyEnumWithEnumStringer int32

const (
	MyEnumWithEnumStringer_EnumValueStringerA MyEnumWithEnumStringer = 0
	MyEnumWithEnumStringer_STRINGER_B         MyEnumWithEnumStringer = 1
)

var MyEnumWithEnumStringer_name = map[int32]string{
	0: "STRINGER_A",
	1: "STRINGER_B",
}

var MyEnumWithEnumStringer_value = map[string]int32{
	"STRINGER_A": 0,
	"STRINGER_B": 1,
}

func (x MyEnumWithEnumStringer) Enum() *MyEnumWithEnumStringer {
	p := new(MyEnumWithEnumStringer)
	*p = x
	return p
}

func (x MyEnumWithEnumStringer) MarshalJSON() ([]byte, error) {
	return proto.MarshalJSONEnum(MyEnumWithEnumStringer_name, int32(x))
}

func (x *MyEnumWithEnumStringer) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(MyEnumWithEnumStringer_value, data, "MyEnumWithEnumStringer")
	if err != nil {
		return err
	}
	*x = MyEnumWithEnumStringer(value)
	return nil
}

func (MyEnumWithEnumStringer) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_49eed3c955d68b51, []int{2}
}

type OnlyEnums struct {
	MyEnum                         *MyCustomEnum               `protobuf:"varint,1,opt,name=my_enum,json=myEnum,enum=enumcustomname.MyCustomEnum" json:"my_enum,omitempty" pg:"my_enum"`
	MyEnumDefaultA                 *MyCustomEnum               `protobuf:"varint,2,opt,name=my_enum_default_a,json=myEnumDefaultA,enum=enumcustomname.MyCustomEnum,def=0" json:"my_enum_default_a,omitempty" pg:"my_enum_default_a"`
	MyEnumDefaultB                 *MyCustomEnum               `protobuf:"varint,3,opt,name=my_enum_default_b,json=myEnumDefaultB,enum=enumcustomname.MyCustomEnum,def=1" json:"my_enum_default_b,omitempty" pg:"my_enum_default_b"`
	MyUnprefixedEnum               *MyCustomUnprefixedEnum     `protobuf:"varint,4,opt,name=my_unprefixed_enum,json=myUnprefixedEnum,enum=enumcustomname.MyCustomUnprefixedEnum" json:"my_unprefixed_enum,omitempty" pg:"my_unprefixed_enum"`
	MyUnprefixedEnumDefaultA       *MyCustomUnprefixedEnum     `protobuf:"varint,5,opt,name=my_unprefixed_enum_default_a,json=myUnprefixedEnumDefaultA,enum=enumcustomname.MyCustomUnprefixedEnum,def=0" json:"my_unprefixed_enum_default_a,omitempty" pg:"my_unprefixed_enum_default_a"`
	MyUnprefixedEnumDefaultB       *MyCustomUnprefixedEnum     `protobuf:"varint,6,opt,name=my_unprefixed_enum_default_b,json=myUnprefixedEnumDefaultB,enum=enumcustomname.MyCustomUnprefixedEnum,def=1" json:"my_unprefixed_enum_default_b,omitempty" pg:"my_unprefixed_enum_default_b"`
	YetAnotherTestEnum             *test.YetAnotherTestEnum    `protobuf:"varint,7,opt,name=yet_another_test_enum,json=yetAnotherTestEnum,enum=test.YetAnotherTestEnum" json:"yet_another_test_enum,omitempty" pg:"yet_another_test_enum"`
	YetAnotherTestEnumDefaultAa    *test.YetAnotherTestEnum    `protobuf:"varint,8,opt,name=yet_another_test_enum_default_aa,json=yetAnotherTestEnumDefaultAa,enum=test.YetAnotherTestEnum,def=0" json:"yet_another_test_enum_default_aa,omitempty" pg:"yet_another_test_enum_default_aa"`
	YetAnotherTestEnumDefaultBb    *test.YetAnotherTestEnum    `protobuf:"varint,9,opt,name=yet_another_test_enum_default_bb,json=yetAnotherTestEnumDefaultBb,enum=test.YetAnotherTestEnum,def=1" json:"yet_another_test_enum_default_bb,omitempty" pg:"yet_another_test_enum_default_bb"`
	YetYetAnotherTestEnum          *test.YetYetAnotherTestEnum `protobuf:"varint,10,opt,name=yet_yet_another_test_enum,json=yetYetAnotherTestEnum,enum=test.YetYetAnotherTestEnum" json:"yet_yet_another_test_enum,omitempty" pg:"yet_yet_another_test_enum"`
	YetYetAnotherTestEnumDefaultCc *test.YetYetAnotherTestEnum `protobuf:"varint,11,opt,name=yet_yet_another_test_enum_default_cc,json=yetYetAnotherTestEnumDefaultCc,enum=test.YetYetAnotherTestEnum,def=0" json:"yet_yet_another_test_enum_default_cc,omitempty" pg:"yet_yet_another_test_enum_default_cc"`
	YetYetAnotherTestEnumDefaultDd *test.YetYetAnotherTestEnum `protobuf:"varint,12,opt,name=yet_yet_another_test_enum_default_dd,json=yetYetAnotherTestEnumDefaultDd,enum=test.YetYetAnotherTestEnum,def=1" json:"yet_yet_another_test_enum_default_dd,omitempty" pg:"yet_yet_another_test_enum_default_dd"`
	XXX_NoUnkeyedLiteral           struct{}                    `json:"-" pg:"-"`
	XXX_unrecognized               []byte                      `json:"-" pg:"-"`
	XXX_sizecache                  int32                       `json:"-" pg:"-"`
}

func (m *OnlyEnums) Reset()         { *m = OnlyEnums{} }
func (m *OnlyEnums) String() string { return proto.CompactTextString(m) }
func (*OnlyEnums) ProtoMessage()    {}
func (*OnlyEnums) Descriptor() ([]byte, []int) {
	return fileDescriptor_49eed3c955d68b51, []int{0}
}
func (m *OnlyEnums) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OnlyEnums.Unmarshal(m, b)
}
func (m *OnlyEnums) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OnlyEnums.Marshal(b, m, deterministic)
}
func (m *OnlyEnums) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnlyEnums.Merge(m, src)
}
func (m *OnlyEnums) XXX_Size() int {
	return xxx_messageInfo_OnlyEnums.Size(m)
}
func (m *OnlyEnums) XXX_DiscardUnknown() {
	xxx_messageInfo_OnlyEnums.DiscardUnknown(m)
}

var xxx_messageInfo_OnlyEnums proto.InternalMessageInfo

const Default_OnlyEnums_MyEnumDefaultA MyCustomEnum = MyCustomEnum_MyBetterNameA
const Default_OnlyEnums_MyEnumDefaultB MyCustomEnum = MyCustomEnum_B
const Default_OnlyEnums_MyUnprefixedEnumDefaultA MyCustomUnprefixedEnum = MyBetterNameUnprefixedA
const Default_OnlyEnums_MyUnprefixedEnumDefaultB MyCustomUnprefixedEnum = UNPREFIXED_B
const Default_OnlyEnums_YetAnotherTestEnumDefaultAa test.YetAnotherTestEnum = test.AA
const Default_OnlyEnums_YetAnotherTestEnumDefaultBb test.YetAnotherTestEnum = test.BetterYetBB
const Default_OnlyEnums_YetYetAnotherTestEnumDefaultCc test.YetYetAnotherTestEnum = test.YetYetAnotherTestEnum_CC
const Default_OnlyEnums_YetYetAnotherTestEnumDefaultDd test.YetYetAnotherTestEnum = test.YetYetAnotherTestEnum_BetterYetDD

func (m *OnlyEnums) GetMyEnum() MyCustomEnum {
	if m != nil && m.MyEnum != nil {
		return *m.MyEnum
	}
	return MyCustomEnum_MyBetterNameA
}

func (m *OnlyEnums) GetMyEnumDefaultA() MyCustomEnum {
	if m != nil && m.MyEnumDefaultA != nil {
		return *m.MyEnumDefaultA
	}
	return Default_OnlyEnums_MyEnumDefaultA
}

func (m *OnlyEnums) GetMyEnumDefaultB() MyCustomEnum {
	if m != nil && m.MyEnumDefaultB != nil {
		return *m.MyEnumDefaultB
	}
	return Default_OnlyEnums_MyEnumDefaultB
}

func (m *OnlyEnums) GetMyUnprefixedEnum() MyCustomUnprefixedEnum {
	if m != nil && m.MyUnprefixedEnum != nil {
		return *m.MyUnprefixedEnum
	}
	return MyBetterNameUnprefixedA
}

func (m *OnlyEnums) GetMyUnprefixedEnumDefaultA() MyCustomUnprefixedEnum {
	if m != nil && m.MyUnprefixedEnumDefaultA != nil {
		return *m.MyUnprefixedEnumDefaultA
	}
	return Default_OnlyEnums_MyUnprefixedEnumDefaultA
}

func (m *OnlyEnums) GetMyUnprefixedEnumDefaultB() MyCustomUnprefixedEnum {
	if m != nil && m.MyUnprefixedEnumDefaultB != nil {
		return *m.MyUnprefixedEnumDefaultB
	}
	return Default_OnlyEnums_MyUnprefixedEnumDefaultB
}

func (m *OnlyEnums) GetYetAnotherTestEnum() test.YetAnotherTestEnum {
	if m != nil && m.YetAnotherTestEnum != nil {
		return *m.YetAnotherTestEnum
	}
	return test.AA
}

func (m *OnlyEnums) GetYetAnotherTestEnumDefaultAa() test.YetAnotherTestEnum {
	if m != nil && m.YetAnotherTestEnumDefaultAa != nil {
		return *m.YetAnotherTestEnumDefaultAa
	}
	return Default_OnlyEnums_YetAnotherTestEnumDefaultAa
}

func (m *OnlyEnums) GetYetAnotherTestEnumDefaultBb() test.YetAnotherTestEnum {
	if m != nil && m.YetAnotherTestEnumDefaultBb != nil {
		return *m.YetAnotherTestEnumDefaultBb
	}
	return Default_OnlyEnums_YetAnotherTestEnumDefaultBb
}

func (m *OnlyEnums) GetYetYetAnotherTestEnum() test.YetYetAnotherTestEnum {
	if m != nil && m.YetYetAnotherTestEnum != nil {
		return *m.YetYetAnotherTestEnum
	}
	return test.YetYetAnotherTestEnum_CC
}

func (m *OnlyEnums) GetYetYetAnotherTestEnumDefaultCc() test.YetYetAnotherTestEnum {
	if m != nil && m.YetYetAnotherTestEnumDefaultCc != nil {
		return *m.YetYetAnotherTestEnumDefaultCc
	}
	return Default_OnlyEnums_YetYetAnotherTestEnumDefaultCc
}

func (m *OnlyEnums) GetYetYetAnotherTestEnumDefaultDd() test.YetYetAnotherTestEnum {
	if m != nil && m.YetYetAnotherTestEnumDefaultDd != nil {
		return *m.YetYetAnotherTestEnumDefaultDd
	}
	return Default_OnlyEnums_YetYetAnotherTestEnumDefaultDd
}

func init() {
	proto.RegisterEnum("enumcustomname.MyCustomEnum", MyCustomEnum_name, MyCustomEnum_value)
	proto.RegisterEnum("enumcustomname.MyCustomUnprefixedEnum", MyCustomUnprefixedEnum_name, MyCustomUnprefixedEnum_value)
	proto.RegisterEnum("enumcustomname.MyEnumWithEnumStringer", MyEnumWithEnumStringer_name, MyEnumWithEnumStringer_value)
	proto.RegisterType((*OnlyEnums)(nil), "enumcustomname.OnlyEnums")
}

func init() { proto.RegisterFile("enumcustomname.proto", fileDescriptor_49eed3c955d68b51) }

var fileDescriptor_49eed3c955d68b51 = []byte{
	// 558 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x93, 0x4f, 0x8f, 0xd2, 0x40,
	0x18, 0xc6, 0xe9, 0xea, 0xb2, 0xec, 0x88, 0xa4, 0x3b, 0x51, 0x1c, 0xc1, 0x34, 0xcd, 0xc6, 0x18,
	0x43, 0x22, 0x44, 0xbd, 0xe1, 0x69, 0x4a, 0xd1, 0x6c, 0x0c, 0x68, 0xba, 0x8b, 0xff, 0x2e, 0x4d,
	0x5b, 0x06, 0x68, 0xc2, 0xb4, 0x9b, 0x32, 0x8d, 0xf6, 0x1b, 0x18, 0xbe, 0x03, 0x27, 0x39, 0x78,
	0xf4, 0xbc, 0x67, 0x3f, 0x98, 0x99, 0xe9, 0xc2, 0x42, 0xdb, 0x2d, 0xc4, 0xd3, 0xc0, 0xe4, 0x79,
	0x9e, 0xdf, 0xbc, 0x4f, 0xde, 0x82, 0x07, 0xc4, 0x0b, 0xa9, 0x13, 0xce, 0x98, 0x4f, 0x3d, 0x8b,
	0x92, 0xe6, 0x65, 0xe0, 0x33, 0x1f, 0x56, 0xb6, 0x6f, 0x6b, 0xaf, 0xc7, 0x2e, 0x9b, 0x84, 0x76,
	0xd3, 0xf1, 0x69, 0xcb, 0x76, 0x19, 0xf3, 0x83, 0x80, 0x78, 0xac, 0x25, 0xc4, 0x76, 0x38, 0x6a,
	0x8d, 0xfd, 0xb1, 0x2f, 0xfe, 0x88, 0x5f, 0x71, 0x48, 0xed, 0xe5, 0x0e, 0x13, 0x23, 0x33, 0xd6,
	0x62, 0x13, 0xc2, 0xcf, 0xd8, 0x72, 0xfa, 0xb7, 0x04, 0x8e, 0x3f, 0x78, 0xd3, 0xa8, 0xeb, 0x85,
	0x74, 0x06, 0x5b, 0xe0, 0x88, 0x46, 0x26, 0x7f, 0x0a, 0x92, 0x54, 0xe9, 0x79, 0xe5, 0x55, 0xb5,
	0x99, 0x78, 0x6d, 0x4f, 0x28, 0x8d, 0x22, 0x15, 0x27, 0xd4, 0xc1, 0xc9, 0xb5, 0xc1, 0x1c, 0x92,
	0x91, 0x15, 0x4e, 0x99, 0x69, 0xa1, 0x83, 0x3c, 0x6b, 0x5b, 0xc2, 0x46, 0x25, 0x76, 0xeb, 0xb1,
	0x03, 0x67, 0xa5, 0xd8, 0xe8, 0x4e, 0x7e, 0x8a, 0x96, 0x48, 0xd1, 0x60, 0x1f, 0x40, 0x1a, 0x99,
	0xa1, 0x77, 0x19, 0x90, 0x91, 0xfb, 0x83, 0x0c, 0xe3, 0x39, 0xee, 0x8a, 0x18, 0x35, 0x1d, 0x33,
	0x58, 0x0b, 0xc5, 0x44, 0x32, 0x4d, 0xdc, 0x40, 0x0f, 0x3c, 0x49, 0xe7, 0x6d, 0x8c, 0x79, 0xb8,
	0x5f, 0x72, 0xbb, 0x3c, 0xe8, 0x7f, 0x34, 0xba, 0x6f, 0xcf, 0xbe, 0x74, 0x75, 0x13, 0x1b, 0x28,
	0xc9, 0x59, 0xb7, 0x90, 0xcf, 0xb3, 0x51, 0xf1, 0x3f, 0x78, 0xda, 0xad, 0x3c, 0x0d, 0xbe, 0x07,
	0x0f, 0x23, 0xc2, 0x4c, 0xcb, 0xf3, 0xd9, 0x84, 0x04, 0x26, 0x5f, 0x8a, 0xb8, 0xb2, 0x23, 0x01,
	0x42, 0x4d, 0xb1, 0x26, 0x5f, 0x09, 0xc3, 0xb1, 0xe2, 0x82, 0xcc, 0x98, 0xa8, 0x0a, 0x46, 0xa9,
	0x3b, 0xe8, 0x00, 0x35, 0x33, 0xec, 0xa6, 0x2f, 0x0b, 0x95, 0xf2, 0x73, 0xdb, 0x07, 0x18, 0x1b,
	0xf5, 0x74, 0xf6, 0xaa, 0x20, 0x6b, 0x37, 0xc4, 0xb6, 0xd1, 0xf1, 0x2e, 0x88, 0xa6, 0xe5, 0x40,
	0x34, 0x1b, 0x0e, 0xc0, 0x63, 0x0e, 0xc9, 0xae, 0x06, 0x88, 0xf4, 0xfa, 0x3a, 0x3d, 0xa3, 0x1d,
	0x5e, 0x6a, 0xfa, 0x1a, 0x52, 0xf0, 0xf4, 0xd6, 0xd8, 0xf5, 0xfb, 0x1d, 0x07, 0xdd, 0xdb, 0x49,
	0x68, 0x1f, 0x74, 0x3a, 0x86, 0x92, 0x49, 0xb9, 0x9e, 0xa2, 0xe3, 0xec, 0x87, 0x1b, 0x0e, 0x51,
	0x79, 0x0f, 0x9c, 0xae, 0xe7, 0xe3, 0xf4, 0x61, 0xe3, 0x0d, 0x28, 0xc6, 0x1f, 0x26, 0x44, 0x40,
	0xc2, 0x72, 0xa1, 0x76, 0x32, 0x5f, 0xa8, 0xf7, 0x7b, 0x91, 0x46, 0x18, 0x23, 0x41, 0xdf, 0xa2,
	0x04, 0xc3, 0x43, 0x20, 0x69, 0xb2, 0x54, 0x93, 0xaf, 0x96, 0x4a, 0xb9, 0x17, 0x75, 0xc4, 0x06,
	0x73, 0x4b, 0xe3, 0x3b, 0x90, 0x93, 0x4b, 0x0c, 0x5f, 0x80, 0xad, 0xcf, 0x46, 0x2e, 0xd4, 0xea,
	0xf3, 0x85, 0xfa, 0x68, 0x33, 0xf1, 0xc6, 0x81, 0xa1, 0xbc, 0x25, 0xe7, 0x98, 0xd3, 0x9f, 0xbf,
	0x94, 0xc2, 0xef, 0xa5, 0x52, 0xb8, 0x5a, 0x2a, 0xd5, 0x15, 0x6e, 0x1b, 0xd2, 0xf8, 0x06, 0xaa,
	0xf1, 0xab, 0x3f, 0xbb, 0x6c, 0xc2, 0xcf, 0x73, 0x16, 0xb8, 0xde, 0x98, 0x04, 0xf0, 0x19, 0x00,
	0xe7, 0x17, 0xc6, 0x59, 0xff, 0x5d, 0xd7, 0x10, 0xf0, 0xea, 0x7c, 0xa1, 0x42, 0xae, 0xf8, 0x64,
	0x4d, 0x43, 0xb2, 0x92, 0x61, 0x58, 0xd9, 0xd0, 0x71, 0x6a, 0x89, 0x13, 0xff, 0x2c, 0x15, 0xe9,
	0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x85, 0xfc, 0x12, 0xfb, 0xe7, 0x05, 0x00, 0x00,
}

func (x MyEnumWithEnumStringer) String() string {
	s, ok := MyEnumWithEnumStringer_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
