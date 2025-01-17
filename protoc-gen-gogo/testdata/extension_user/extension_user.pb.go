// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: extension_user/extension_user.proto

package extension_user

import (
	fmt "fmt"
	proto "github.com/bittorrent/protobuf/proto"
	extension_base "github.com/bittorrent/protobuf/protoc-gen-gogo/testdata/extension_base"
	extension_extra "github.com/bittorrent/protobuf/protoc-gen-gogo/testdata/extension_extra"
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

type UserMessage struct {
	Name                 *string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty" pg:"name"`
	Rank                 *string  `protobuf:"bytes,2,opt,name=rank" json:"rank,omitempty" pg:"rank"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" pg:"-"`
	XXX_unrecognized     []byte   `json:"-" pg:"-"`
	XXX_sizecache        int32    `json:"-" pg:"-"`
}

func (m *UserMessage) Reset()         { *m = UserMessage{} }
func (m *UserMessage) String() string { return proto.CompactTextString(m) }
func (*UserMessage) ProtoMessage()    {}
func (*UserMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_359ba8abf543ca10, []int{0}
}
func (m *UserMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserMessage.Unmarshal(m, b)
}
func (m *UserMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserMessage.Marshal(b, m, deterministic)
}
func (m *UserMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserMessage.Merge(m, src)
}
func (m *UserMessage) XXX_Size() int {
	return xxx_messageInfo_UserMessage.Size(m)
}
func (m *UserMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_UserMessage.DiscardUnknown(m)
}

var xxx_messageInfo_UserMessage proto.InternalMessageInfo

func (m *UserMessage) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *UserMessage) GetRank() string {
	if m != nil && m.Rank != nil {
		return *m.Rank
	}
	return ""
}

// Extend inside the scope of another type
type LoudMessage struct {
	XXX_NoUnkeyedLiteral         struct{} `json:"-" pg:"-"`
	proto.XXX_InternalExtensions `json:"-" pg:"-"`
	XXX_unrecognized             []byte `json:"-" pg:"-"`
	XXX_sizecache                int32  `json:"-" pg:"-"`
}

func (m *LoudMessage) Reset()         { *m = LoudMessage{} }
func (m *LoudMessage) String() string { return proto.CompactTextString(m) }
func (*LoudMessage) ProtoMessage()    {}
func (*LoudMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_359ba8abf543ca10, []int{1}
}

var extRange_LoudMessage = []proto.ExtensionRange{
	{Start: 100, End: 536870911},
}

func (*LoudMessage) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_LoudMessage
}

func (m *LoudMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoudMessage.Unmarshal(m, b)
}
func (m *LoudMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoudMessage.Marshal(b, m, deterministic)
}
func (m *LoudMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoudMessage.Merge(m, src)
}
func (m *LoudMessage) XXX_Size() int {
	return xxx_messageInfo_LoudMessage.Size(m)
}
func (m *LoudMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_LoudMessage.DiscardUnknown(m)
}

var xxx_messageInfo_LoudMessage proto.InternalMessageInfo

var E_LoudMessage_Volume = &proto.ExtensionDesc{
	ExtendedType:  (*extension_base.BaseMessage)(nil),
	ExtensionType: (*uint32)(nil),
	Field:         8,
	Name:          "extension_user.LoudMessage.volume",
	Tag:           "varint,8,opt,name=volume",
	Filename:      "extension_user/extension_user.proto",
}

// Extend inside the scope of another type, using a message.
type LoginMessage struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" pg:"-"`
	XXX_unrecognized     []byte   `json:"-" pg:"-"`
	XXX_sizecache        int32    `json:"-" pg:"-"`
}

func (m *LoginMessage) Reset()         { *m = LoginMessage{} }
func (m *LoginMessage) String() string { return proto.CompactTextString(m) }
func (*LoginMessage) ProtoMessage()    {}
func (*LoginMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_359ba8abf543ca10, []int{2}
}
func (m *LoginMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginMessage.Unmarshal(m, b)
}
func (m *LoginMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginMessage.Marshal(b, m, deterministic)
}
func (m *LoginMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginMessage.Merge(m, src)
}
func (m *LoginMessage) XXX_Size() int {
	return xxx_messageInfo_LoginMessage.Size(m)
}
func (m *LoginMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginMessage.DiscardUnknown(m)
}

var xxx_messageInfo_LoginMessage proto.InternalMessageInfo

var E_LoginMessage_UserMessage = &proto.ExtensionDesc{
	ExtendedType:  (*extension_base.BaseMessage)(nil),
	ExtensionType: (*UserMessage)(nil),
	Field:         16,
	Name:          "extension_user.LoginMessage.user_message",
	Tag:           "bytes,16,opt,name=user_message",
	Filename:      "extension_user/extension_user.proto",
}

type Detail struct {
	Color                *string  `protobuf:"bytes,1,opt,name=color" json:"color,omitempty" pg:"color"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" pg:"-"`
	XXX_unrecognized     []byte   `json:"-" pg:"-"`
	XXX_sizecache        int32    `json:"-" pg:"-"`
}

func (m *Detail) Reset()         { *m = Detail{} }
func (m *Detail) String() string { return proto.CompactTextString(m) }
func (*Detail) ProtoMessage()    {}
func (*Detail) Descriptor() ([]byte, []int) {
	return fileDescriptor_359ba8abf543ca10, []int{3}
}
func (m *Detail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Detail.Unmarshal(m, b)
}
func (m *Detail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Detail.Marshal(b, m, deterministic)
}
func (m *Detail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Detail.Merge(m, src)
}
func (m *Detail) XXX_Size() int {
	return xxx_messageInfo_Detail.Size(m)
}
func (m *Detail) XXX_DiscardUnknown() {
	xxx_messageInfo_Detail.DiscardUnknown(m)
}

var xxx_messageInfo_Detail proto.InternalMessageInfo

func (m *Detail) GetColor() string {
	if m != nil && m.Color != nil {
		return *m.Color
	}
	return ""
}

// An extension of an extension
type Announcement struct {
	Words                *string  `protobuf:"bytes,1,opt,name=words" json:"words,omitempty" pg:"words"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" pg:"-"`
	XXX_unrecognized     []byte   `json:"-" pg:"-"`
	XXX_sizecache        int32    `json:"-" pg:"-"`
}

func (m *Announcement) Reset()         { *m = Announcement{} }
func (m *Announcement) String() string { return proto.CompactTextString(m) }
func (*Announcement) ProtoMessage()    {}
func (*Announcement) Descriptor() ([]byte, []int) {
	return fileDescriptor_359ba8abf543ca10, []int{4}
}
func (m *Announcement) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Announcement.Unmarshal(m, b)
}
func (m *Announcement) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Announcement.Marshal(b, m, deterministic)
}
func (m *Announcement) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Announcement.Merge(m, src)
}
func (m *Announcement) XXX_Size() int {
	return xxx_messageInfo_Announcement.Size(m)
}
func (m *Announcement) XXX_DiscardUnknown() {
	xxx_messageInfo_Announcement.DiscardUnknown(m)
}

var xxx_messageInfo_Announcement proto.InternalMessageInfo

func (m *Announcement) GetWords() string {
	if m != nil && m.Words != nil {
		return *m.Words
	}
	return ""
}

var E_Announcement_LoudExt = &proto.ExtensionDesc{
	ExtendedType:  (*LoudMessage)(nil),
	ExtensionType: (*Announcement)(nil),
	Field:         100,
	Name:          "extension_user.Announcement.loud_ext",
	Tag:           "bytes,100,opt,name=loud_ext",
	Filename:      "extension_user/extension_user.proto",
}

// Something that can be put in a message set.
type OldStyleParcel struct {
	Name                 *string  `protobuf:"bytes,1,req,name=name" json:"name,omitempty" pg:"name"`
	Height               *int32   `protobuf:"varint,2,opt,name=height" json:"height,omitempty" pg:"height"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" pg:"-"`
	XXX_unrecognized     []byte   `json:"-" pg:"-"`
	XXX_sizecache        int32    `json:"-" pg:"-"`
}

func (m *OldStyleParcel) Reset()         { *m = OldStyleParcel{} }
func (m *OldStyleParcel) String() string { return proto.CompactTextString(m) }
func (*OldStyleParcel) ProtoMessage()    {}
func (*OldStyleParcel) Descriptor() ([]byte, []int) {
	return fileDescriptor_359ba8abf543ca10, []int{5}
}
func (m *OldStyleParcel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OldStyleParcel.Unmarshal(m, b)
}
func (m *OldStyleParcel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OldStyleParcel.Marshal(b, m, deterministic)
}
func (m *OldStyleParcel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OldStyleParcel.Merge(m, src)
}
func (m *OldStyleParcel) XXX_Size() int {
	return xxx_messageInfo_OldStyleParcel.Size(m)
}
func (m *OldStyleParcel) XXX_DiscardUnknown() {
	xxx_messageInfo_OldStyleParcel.DiscardUnknown(m)
}

var xxx_messageInfo_OldStyleParcel proto.InternalMessageInfo

func (m *OldStyleParcel) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *OldStyleParcel) GetHeight() int32 {
	if m != nil && m.Height != nil {
		return *m.Height
	}
	return 0
}

var E_OldStyleParcel_MessageSetExtension = &proto.ExtensionDesc{
	ExtendedType:  (*extension_base.OldStyleMessage)(nil),
	ExtensionType: (*OldStyleParcel)(nil),
	Field:         2001,
	Name:          "extension_user.OldStyleParcel",
	Tag:           "bytes,2001,opt,name=message_set_extension",
	Filename:      "extension_user/extension_user.proto",
}

var E_UserMessage = &proto.ExtensionDesc{
	ExtendedType:  (*extension_base.BaseMessage)(nil),
	ExtensionType: (*UserMessage)(nil),
	Field:         5,
	Name:          "extension_user.user_message",
	Tag:           "bytes,5,opt,name=user_message",
	Filename:      "extension_user/extension_user.proto",
}

var E_ExtraMessage = &proto.ExtensionDesc{
	ExtendedType:  (*extension_base.BaseMessage)(nil),
	ExtensionType: (*extension_extra.ExtraMessage)(nil),
	Field:         9,
	Name:          "extension_user.extra_message",
	Tag:           "bytes,9,opt,name=extra_message",
	Filename:      "extension_user/extension_user.proto",
}

var E_Width = &proto.ExtensionDesc{
	ExtendedType:  (*extension_base.BaseMessage)(nil),
	ExtensionType: (*int32)(nil),
	Field:         6,
	Name:          "extension_user.width",
	Tag:           "varint,6,opt,name=width",
	Filename:      "extension_user/extension_user.proto",
}

var E_Area = &proto.ExtensionDesc{
	ExtendedType:  (*extension_base.BaseMessage)(nil),
	ExtensionType: (*int64)(nil),
	Field:         7,
	Name:          "extension_user.area",
	Tag:           "varint,7,opt,name=area",
	Filename:      "extension_user/extension_user.proto",
}

var E_Detail = &proto.ExtensionDesc{
	ExtendedType:  (*extension_base.BaseMessage)(nil),
	ExtensionType: ([]*Detail)(nil),
	Field:         17,
	Name:          "extension_user.detail",
	Tag:           "bytes,17,rep,name=detail",
	Filename:      "extension_user/extension_user.proto",
}

func init() {
	proto.RegisterType((*UserMessage)(nil), "extension_user.UserMessage")
	proto.RegisterExtension(E_LoudMessage_Volume)
	proto.RegisterType((*LoudMessage)(nil), "extension_user.LoudMessage")
	proto.RegisterExtension(E_LoginMessage_UserMessage)
	proto.RegisterType((*LoginMessage)(nil), "extension_user.LoginMessage")
	proto.RegisterType((*Detail)(nil), "extension_user.Detail")
	proto.RegisterExtension(E_Announcement_LoudExt)
	proto.RegisterType((*Announcement)(nil), "extension_user.Announcement")
	proto.RegisterExtension(E_OldStyleParcel_MessageSetExtension)
	proto.RegisterType((*OldStyleParcel)(nil), "extension_user.OldStyleParcel")
	proto.RegisterExtension(E_UserMessage)
	proto.RegisterExtension(E_ExtraMessage)
	proto.RegisterExtension(E_Width)
	proto.RegisterExtension(E_Area)
	proto.RegisterExtension(E_Detail)
}

func init() {
	proto.RegisterFile("extension_user/extension_user.proto", fileDescriptor_359ba8abf543ca10)
}

var fileDescriptor_359ba8abf543ca10 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x51, 0x6f, 0x94, 0x4c,
	0x14, 0x0d, 0xdb, 0x2e, 0xdd, 0x0e, 0xdb, 0x7e, 0xfd, 0x50, 0x9b, 0x4d, 0xd5, 0x4a, 0x30, 0x26,
	0xc4, 0xa4, 0x4b, 0xc4, 0xf8, 0xc2, 0x9b, 0x8d, 0x6b, 0x4c, 0x5c, 0xa3, 0xa1, 0xfa, 0xa2, 0x0f,
	0x64, 0x16, 0xae, 0x2c, 0x29, 0xcc, 0x98, 0x99, 0x8b, 0x5d, 0x7d, 0xda, 0xdf, 0xe4, 0x3f, 0xf1,
	0x1f, 0x19, 0x86, 0xa1, 0x65, 0x31, 0xd9, 0xd8, 0x17, 0x32, 0xe7, 0x72, 0xee, 0xb9, 0x77, 0xce,
	0xbd, 0x40, 0x1e, 0xc3, 0x0a, 0x81, 0xc9, 0x9c, 0xb3, 0xb8, 0x92, 0x20, 0xfc, 0x4d, 0x38, 0xfd,
	0x26, 0x38, 0x72, 0xfb, 0x70, 0x33, 0x7a, 0xd2, 0x49, 0x5a, 0x50, 0x09, 0xfe, 0x26, 0x6c, 0x92,
	0x4e, 0x9e, 0xdc, 0x44, 0x61, 0x85, 0x82, 0xfa, 0x3d, 0xdc, 0xd0, 0xdc, 0x17, 0xc4, 0xfa, 0x24,
	0x41, 0xbc, 0x03, 0x29, 0x69, 0x06, 0xb6, 0x4d, 0x76, 0x19, 0x2d, 0x61, 0x62, 0x38, 0x86, 0xb7,
	0x1f, 0xa9, 0x73, 0x1d, 0x13, 0x94, 0x5d, 0x4e, 0x06, 0x4d, 0xac, 0x3e, 0xbb, 0x73, 0x62, 0xcd,
	0x79, 0x95, 0xea, 0xb4, 0xa7, 0xa3, 0x51, 0x7a, 0xb4, 0x5e, 0xaf, 0xd7, 0x83, 0xe0, 0x39, 0x31,
	0xbf, 0xf3, 0xa2, 0x2a, 0xc1, 0xbe, 0x3f, 0xed, 0xf5, 0x75, 0x4e, 0x25, 0xe8, 0x84, 0xc9, 0xc8,
	0x31, 0xbc, 0x83, 0x48, 0x53, 0xdd, 0x4b, 0x32, 0x9e, 0xf3, 0x2c, 0x67, 0xfa, 0x6d, 0xf0, 0x85,
	0x8c, 0xeb, 0x8b, 0xc6, 0xa5, 0xee, 0x6a, 0xab, 0xd4, 0x91, 0x63, 0x78, 0x56, 0xd0, 0xa5, 0x28,
	0xeb, 0x3a, 0xb7, 0x8a, 0xac, 0xea, 0x06, 0xb8, 0xa7, 0xc4, 0x7c, 0x05, 0x48, 0xf3, 0xc2, 0xbe,
	0x4b, 0x86, 0x09, 0x2f, 0xb8, 0xd0, 0xb7, 0x6d, 0x80, 0xfb, 0x93, 0x8c, 0x5f, 0x32, 0xc6, 0x2b,
	0x96, 0x40, 0x09, 0x0c, 0x6b, 0xd6, 0x15, 0x17, 0xa9, 0x6c, 0x59, 0x0a, 0x04, 0x1f, 0xc9, 0xa8,
	0xe0, 0x55, 0x5a, 0x7b, 0x69, 0xff, 0x55, 0xbb, 0x63, 0xcd, 0x24, 0x55, 0xed, 0x3d, 0xe8, 0x53,
	0xba, 0x25, 0xa2, 0xbd, 0x5a, 0x6a, 0xb6, 0x42, 0xf7, 0x97, 0x41, 0x0e, 0xdf, 0x17, 0xe9, 0x05,
	0xfe, 0x28, 0xe0, 0x03, 0x15, 0x09, 0x14, 0x9d, 0x89, 0x0c, 0xae, 0x27, 0x72, 0x4c, 0xcc, 0x25,
	0xe4, 0xd9, 0x12, 0xd5, 0x4c, 0x86, 0x91, 0x46, 0x01, 0x92, 0x7b, 0xda, 0xb2, 0x58, 0x02, 0xc6,
	0xd7, 0x25, 0xed, 0x47, 0x7d, 0x03, 0xdb, 0x22, 0x6d, 0x97, 0xbf, 0xff, 0x53, 0x6d, 0x9e, 0xf6,
	0xdb, 0xdc, 0x6c, 0x26, 0xba, 0xa3, 0xe5, 0x2f, 0x00, 0x67, 0x2d, 0x31, 0xbc, 0xd5, 0xb4, 0x86,
	0xb7, 0x9b, 0x56, 0x18, 0x93, 0x03, 0xb5, 0xae, 0xff, 0xa6, 0xbe, 0xaf, 0xd4, 0x1f, 0x4e, 0xfb,
	0xbb, 0x3e, 0xab, 0x9f, 0xad, 0xfe, 0x18, 0x3a, 0x28, 0x7c, 0x46, 0x86, 0x57, 0x79, 0x8a, 0xcb,
	0xed, 0xc2, 0xa6, 0xf2, 0xb9, 0x61, 0x86, 0x3e, 0xd9, 0xa5, 0x02, 0xe8, 0xf6, 0x8c, 0x3d, 0xc7,
	0xf0, 0x76, 0x22, 0x45, 0x0c, 0xdf, 0x12, 0x33, 0x6d, 0x56, 0x6e, 0x6b, 0xca, 0xff, 0xce, 0x8e,
	0x67, 0x05, 0xc7, 0x7d, 0x6f, 0x9a, 0x6d, 0x8d, 0xb4, 0xc4, 0xf9, 0x9b, 0xcf, 0xaf, 0xb3, 0x1c,
	0x97, 0xd5, 0x62, 0x9a, 0xf0, 0xd2, 0x5f, 0xe4, 0x88, 0x5c, 0x08, 0x60, 0xe8, 0xab, 0x0f, 0x7a,
	0x51, 0x7d, 0x6d, 0x0e, 0xc9, 0x59, 0x06, 0xec, 0x2c, 0xe3, 0x19, 0xf7, 0x11, 0x24, 0xa6, 0x14,
	0x69, 0xef, 0xef, 0xf2, 0x27, 0x00, 0x00, 0xff, 0xff, 0x30, 0xb2, 0xe9, 0x9d, 0x7d, 0x04, 0x00,
	0x00,
}
