// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto3/proto3.proto

package proto3

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

type Request_Flavour int32

const (
	Request_SWEET         Request_Flavour = 0
	Request_SOUR          Request_Flavour = 1
	Request_UMAMI         Request_Flavour = 2
	Request_GOPHERLICIOUS Request_Flavour = 3
)

var Request_Flavour_name = map[int32]string{
	0: "SWEET",
	1: "SOUR",
	2: "UMAMI",
	3: "GOPHERLICIOUS",
}

var Request_Flavour_value = map[string]int32{
	"SWEET":         0,
	"SOUR":          1,
	"UMAMI":         2,
	"GOPHERLICIOUS": 3,
}

func (x Request_Flavour) String() string {
	return proto.EnumName(Request_Flavour_name, int32(x))
}

func (Request_Flavour) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ab04eb4084a521db, []int{0, 0}
}

type Request struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty" pg:"name"`
	Key                  []int64         `protobuf:"varint,2,rep,packed,name=key,proto3" json:"key,omitempty" pg:"key"`
	Taste                Request_Flavour `protobuf:"varint,3,opt,name=taste,proto3,enum=proto3.Request_Flavour" json:"taste,omitempty" pg:"taste"`
	Book                 *Book           `protobuf:"bytes,4,opt,name=book,proto3" json:"book,omitempty" pg:"book"`
	Unpacked             []int64         `protobuf:"varint,5,rep,name=unpacked,proto3" json:"unpacked,omitempty" pg:"unpacked"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-" pg:"-"`
	XXX_unrecognized     []byte          `json:"-" pg:"-"`
	XXX_sizecache        int32           `json:"-" pg:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab04eb4084a521db, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Request) GetKey() []int64 {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Request) GetTaste() Request_Flavour {
	if m != nil {
		return m.Taste
	}
	return Request_SWEET
}

func (m *Request) GetBook() *Book {
	if m != nil {
		return m.Book
	}
	return nil
}

func (m *Request) GetUnpacked() []int64 {
	if m != nil {
		return m.Unpacked
	}
	return nil
}

type Book struct {
	Title                string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" pg:"title"`
	RawData              []byte   `protobuf:"bytes,2,opt,name=raw_data,json=rawData,proto3" json:"raw_data,omitempty" pg:"raw_data"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" pg:"-"`
	XXX_unrecognized     []byte   `json:"-" pg:"-"`
	XXX_sizecache        int32    `json:"-" pg:"-"`
}

func (m *Book) Reset()         { *m = Book{} }
func (m *Book) String() string { return proto.CompactTextString(m) }
func (*Book) ProtoMessage()    {}
func (*Book) Descriptor() ([]byte, []int) {
	return fileDescriptor_ab04eb4084a521db, []int{1}
}
func (m *Book) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Book.Unmarshal(m, b)
}
func (m *Book) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Book.Marshal(b, m, deterministic)
}
func (m *Book) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Book.Merge(m, src)
}
func (m *Book) XXX_Size() int {
	return xxx_messageInfo_Book.Size(m)
}
func (m *Book) XXX_DiscardUnknown() {
	xxx_messageInfo_Book.DiscardUnknown(m)
}

var xxx_messageInfo_Book proto.InternalMessageInfo

func (m *Book) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Book) GetRawData() []byte {
	if m != nil {
		return m.RawData
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto3.Request_Flavour", Request_Flavour_name, Request_Flavour_value)
	proto.RegisterType((*Request)(nil), "proto3.Request")
	proto.RegisterType((*Book)(nil), "proto3.Book")
}

func init() { proto.RegisterFile("proto3/proto3.proto", fileDescriptor_ab04eb4084a521db) }

var fileDescriptor_ab04eb4084a521db = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x90, 0xcf, 0x4f, 0xfa, 0x40,
	0x10, 0xc5, 0xd9, 0xfe, 0xf8, 0x02, 0xf3, 0x45, 0x53, 0x47, 0x13, 0xeb, 0xc5, 0x34, 0x9c, 0x7a,
	0xa1, 0x24, 0x78, 0xf0, 0x62, 0x8c, 0xa2, 0xa8, 0x24, 0x12, 0xcc, 0x22, 0x31, 0xf1, 0x62, 0xb6,
	0xb0, 0x56, 0x52, 0xe8, 0xe2, 0x32, 0x95, 0xf8, 0xcf, 0xfa, 0xb7, 0x98, 0x76, 0x8b, 0xa7, 0x79,
	0x33, 0xfb, 0xf2, 0x79, 0xd9, 0x07, 0x87, 0x6b, 0xad, 0x48, 0x9d, 0x75, 0xcd, 0x88, 0xca, 0x81,
	0xff, 0xcc, 0xd6, 0xfe, 0x61, 0x50, 0xe7, 0xf2, 0x33, 0x97, 0x1b, 0x42, 0x04, 0x27, 0x13, 0x2b,
	0xe9, 0xb3, 0x80, 0x85, 0x4d, 0x5e, 0x6a, 0xf4, 0xc0, 0x4e, 0xe5, 0xb7, 0x6f, 0x05, 0x76, 0x68,
	0xf3, 0x42, 0x62, 0x07, 0x5c, 0x12, 0x1b, 0x92, 0xbe, 0x1d, 0xb0, 0x70, 0xbf, 0x77, 0x1c, 0x55,
	0xdc, 0x8a, 0x12, 0xdd, 0x2d, 0xc5, 0x97, 0xca, 0x35, 0x37, 0x2e, 0x0c, 0xc0, 0x89, 0x95, 0x4a,
	0x7d, 0x27, 0x60, 0xe1, 0xff, 0x5e, 0x6b, 0xe7, 0xee, 0x2b, 0x95, 0xf2, 0xf2, 0x05, 0x4f, 0xa1,
	0x91, 0x67, 0x6b, 0x31, 0x4b, 0xe5, 0xdc, 0x77, 0x8b, 0x9c, 0xbe, 0xe5, 0xd5, 0xf8, 0xdf, 0xad,
	0x7d, 0x01, 0xf5, 0x8a, 0x89, 0x4d, 0x70, 0x27, 0x2f, 0x83, 0xc1, 0xb3, 0x57, 0xc3, 0x06, 0x38,
	0x93, 0xf1, 0x94, 0x7b, 0xac, 0x38, 0x4e, 0x47, 0xd7, 0xa3, 0xa1, 0x67, 0xe1, 0x01, 0xec, 0xdd,
	0x8f, 0x9f, 0x1e, 0x06, 0xfc, 0x71, 0x78, 0x33, 0x1c, 0x4f, 0x27, 0x9e, 0xdd, 0x3e, 0x07, 0xa7,
	0xc8, 0xc2, 0x23, 0x70, 0x69, 0x41, 0xcb, 0xdd, 0xef, 0xcc, 0x82, 0x27, 0xd0, 0xd0, 0x62, 0xfb,
	0x36, 0x17, 0x24, 0x7c, 0x2b, 0x60, 0x61, 0x8b, 0xd7, 0xb5, 0xd8, 0xde, 0x0a, 0x12, 0xfd, 0xab,
	0xd7, 0xcb, 0x64, 0x41, 0x1f, 0x79, 0x1c, 0xcd, 0xd4, 0xaa, 0x1b, 0x2f, 0x88, 0x94, 0xd6, 0x32,
	0x23, 0xd3, 0x63, 0x9c, 0xbf, 0x1b, 0x31, 0xeb, 0x24, 0x32, 0xeb, 0x24, 0x2a, 0x51, 0x5d, 0x92,
	0x1b, 0x2a, 0x48, 0x55, 0xd3, 0x71, 0xd5, 0xf1, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x41, 0xa3,
	0xaa, 0x9c, 0x81, 0x01, 0x00, 0x00,
}
