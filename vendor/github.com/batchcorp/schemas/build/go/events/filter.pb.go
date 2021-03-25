// Code generated by protoc-gen-go. DO NOT EDIT.
// source: filter.proto

package events

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Filter_Type int32

const (
	Filter_LUCENE Filter_Type = 0
	Filter_ES_SQL Filter_Type = 1
)

var Filter_Type_name = map[int32]string{
	0: "LUCENE",
	1: "ES_SQL",
}

var Filter_Type_value = map[string]int32{
	"LUCENE": 0,
	"ES_SQL": 1,
}

func (x Filter_Type) String() string {
	return proto.EnumName(Filter_Type_name, int32(x))
}

func (Filter_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1f5303cab7a20d6f, []int{0, 0}
}

type Filter struct {
	Type                 Filter_Type `protobuf:"varint,1,opt,name=type,proto3,enum=events.Filter_Type" json:"type,omitempty"`
	Query                string      `protobuf:"bytes,2,opt,name=query,proto3" json:"query,omitempty"`
	Pagesize             int32       `protobuf:"varint,3,opt,name=pagesize,proto3" json:"pagesize,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f5303cab7a20d6f, []int{0}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetType() Filter_Type {
	if m != nil {
		return m.Type
	}
	return Filter_LUCENE
}

func (m *Filter) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *Filter) GetPagesize() int32 {
	if m != nil {
		return m.Pagesize
	}
	return 0
}

func init() {
	proto.RegisterEnum("events.Filter_Type", Filter_Type_name, Filter_Type_value)
	proto.RegisterType((*Filter)(nil), "events.Filter")
}

func init() { proto.RegisterFile("filter.proto", fileDescriptor_1f5303cab7a20d6f) }

var fileDescriptor_1f5303cab7a20d6f = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xcb, 0xcc, 0x29,
	0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4b, 0x2d, 0x4b, 0xcd, 0x2b, 0x29,
	0x56, 0x6a, 0x66, 0xe4, 0x62, 0x73, 0x03, 0x4b, 0x08, 0xa9, 0x73, 0xb1, 0x94, 0x54, 0x16, 0xa4,
	0x4a, 0x30, 0x2a, 0x30, 0x6a, 0xf0, 0x19, 0x09, 0xeb, 0x41, 0x54, 0xe8, 0x41, 0x64, 0xf5, 0x42,
	0x2a, 0x0b, 0x52, 0x83, 0xc0, 0x0a, 0x84, 0x44, 0xb8, 0x58, 0x0b, 0x4b, 0x53, 0x8b, 0x2a, 0x25,
	0x98, 0x14, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x21, 0x29, 0x2e, 0x8e, 0x82, 0xc4, 0xf4, 0xd4,
	0xe2, 0xcc, 0xaa, 0x54, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x38, 0x5f, 0x49, 0x8e, 0x8b,
	0x05, 0xa4, 0x5f, 0x88, 0x8b, 0x8b, 0xcd, 0x27, 0xd4, 0xd9, 0xd5, 0xcf, 0x55, 0x80, 0x01, 0xc4,
	0x76, 0x0d, 0x8e, 0x0f, 0x0e, 0xf4, 0x11, 0x60, 0x74, 0xd2, 0x8b, 0xd2, 0x49, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0x4a, 0x2c, 0x49, 0xce, 0x48, 0xce, 0x2f, 0x2a,
	0xd0, 0x2f, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x2c, 0xd6, 0x4f, 0x2a, 0xcd, 0xcc, 0x49, 0xd1, 0x4f,
	0xcf, 0xd7, 0x87, 0xb8, 0x29, 0x89, 0x0d, 0xec, 0x09, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xab, 0x1c, 0xec, 0xc3, 0xd4, 0x00, 0x00, 0x00,
}
