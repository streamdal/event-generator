// Code generated by protoc-gen-go. DO NOT EDIT.
// source: records/kafka.proto

package records

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

type KafkaHeader struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KafkaHeader) Reset()         { *m = KafkaHeader{} }
func (m *KafkaHeader) String() string { return proto.CompactTextString(m) }
func (*KafkaHeader) ProtoMessage()    {}
func (*KafkaHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_d05b8249ca05dcf4, []int{0}
}

func (m *KafkaHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaHeader.Unmarshal(m, b)
}
func (m *KafkaHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaHeader.Marshal(b, m, deterministic)
}
func (m *KafkaHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaHeader.Merge(m, src)
}
func (m *KafkaHeader) XXX_Size() int {
	return xxx_messageInfo_KafkaHeader.Size(m)
}
func (m *KafkaHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaHeader.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaHeader proto.InternalMessageInfo

func (m *KafkaHeader) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KafkaHeader) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type KafkaSinkRecord struct {
	Topic                string         `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Key                  []byte         `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte         `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp            int64          `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Offset               int64          `protobuf:"varint,5,opt,name=offset,proto3" json:"offset,omitempty"`
	Partition            int32          `protobuf:"varint,6,opt,name=partition,proto3" json:"partition,omitempty"`
	Headers              []*KafkaHeader `protobuf:"bytes,7,rep,name=headers,proto3" json:"headers,omitempty"`
	ForceDeadLetter      bool           `protobuf:"varint,8,opt,name=force_dead_letter,json=forceDeadLetter,proto3" json:"force_dead_letter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *KafkaSinkRecord) Reset()         { *m = KafkaSinkRecord{} }
func (m *KafkaSinkRecord) String() string { return proto.CompactTextString(m) }
func (*KafkaSinkRecord) ProtoMessage()    {}
func (*KafkaSinkRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_d05b8249ca05dcf4, []int{1}
}

func (m *KafkaSinkRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KafkaSinkRecord.Unmarshal(m, b)
}
func (m *KafkaSinkRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KafkaSinkRecord.Marshal(b, m, deterministic)
}
func (m *KafkaSinkRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KafkaSinkRecord.Merge(m, src)
}
func (m *KafkaSinkRecord) XXX_Size() int {
	return xxx_messageInfo_KafkaSinkRecord.Size(m)
}
func (m *KafkaSinkRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_KafkaSinkRecord.DiscardUnknown(m)
}

var xxx_messageInfo_KafkaSinkRecord proto.InternalMessageInfo

func (m *KafkaSinkRecord) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *KafkaSinkRecord) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KafkaSinkRecord) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *KafkaSinkRecord) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *KafkaSinkRecord) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *KafkaSinkRecord) GetPartition() int32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

func (m *KafkaSinkRecord) GetHeaders() []*KafkaHeader {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *KafkaSinkRecord) GetForceDeadLetter() bool {
	if m != nil {
		return m.ForceDeadLetter
	}
	return false
}

func init() {
	proto.RegisterType((*KafkaHeader)(nil), "records.KafkaHeader")
	proto.RegisterType((*KafkaSinkRecord)(nil), "records.KafkaSinkRecord")
}

func init() { proto.RegisterFile("records/kafka.proto", fileDescriptor_d05b8249ca05dcf4) }

var fileDescriptor_d05b8249ca05dcf4 = []byte{
	// 302 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x91, 0xcf, 0x4b, 0xc3, 0x30,
	0x1c, 0xc5, 0xa9, 0x75, 0xbf, 0x32, 0x61, 0x1a, 0x87, 0xe6, 0xe0, 0xa1, 0xec, 0x54, 0x04, 0x13,
	0x50, 0xbc, 0x7a, 0x10, 0x0f, 0x82, 0x9e, 0xea, 0x49, 0x2f, 0x23, 0x4d, 0xbe, 0x5d, 0x43, 0xdb,
	0xa5, 0x24, 0xdf, 0x0d, 0xfc, 0xeb, 0x95, 0x66, 0x2d, 0x13, 0x6f, 0x79, 0x9f, 0xef, 0x7b, 0x0f,
	0x1e, 0x21, 0x97, 0x0e, 0x94, 0x75, 0xda, 0x8b, 0x4a, 0x16, 0x95, 0xe4, 0xad, 0xb3, 0x68, 0xe9,
	0xa4, 0x87, 0xab, 0x47, 0x32, 0x7f, 0xeb, 0xf8, 0x2b, 0x48, 0x0d, 0x8e, 0x9e, 0x93, 0xb8, 0x82,
	0x6f, 0x16, 0x25, 0x51, 0x3a, 0xcb, 0xba, 0x27, 0x5d, 0x92, 0xd1, 0x5e, 0xd6, 0x3b, 0x60, 0x27,
	0x81, 0x1d, 0xc4, 0xea, 0x27, 0x22, 0x8b, 0x90, 0xfb, 0x30, 0xdb, 0x2a, 0x0b, 0x5d, 0x9d, 0x13,
	0x6d, 0x6b, 0x54, 0x9f, 0x3e, 0x88, 0xa1, 0xb1, 0x4b, 0x9f, 0xfd, 0x6b, 0x8c, 0x03, 0x3b, 0x08,
	0x7a, 0x43, 0x66, 0x68, 0x1a, 0xf0, 0x28, 0x9b, 0x96, 0x9d, 0x26, 0x51, 0x1a, 0x67, 0x47, 0x40,
	0xaf, 0xc8, 0xd8, 0x16, 0x85, 0x07, 0x64, 0xa3, 0x70, 0xea, 0x55, 0x97, 0x6a, 0xa5, 0x43, 0x83,
	0xc6, 0x6e, 0xd9, 0x38, 0x89, 0xd2, 0x51, 0x76, 0x04, 0x94, 0x93, 0x49, 0x19, 0x76, 0x79, 0x36,
	0x49, 0xe2, 0x74, 0x7e, 0xbf, 0xe4, 0xfd, 0x6e, 0xfe, 0x67, 0x74, 0x36, 0x98, 0xe8, 0x2d, 0xb9,
	0x28, 0xac, 0x53, 0xb0, 0xd6, 0x20, 0xf5, 0xba, 0x06, 0x44, 0x70, 0x6c, 0x9a, 0x44, 0xe9, 0x34,
	0x5b, 0x84, 0xc3, 0x0b, 0x48, 0xfd, 0x1e, 0xf0, 0xf3, 0x27, 0xb9, 0xf6, 0x25, 0xcf, 0x25, 0xaa,
	0x92, 0xc3, 0x1e, 0xb6, 0xe8, 0x87, 0xee, 0xaf, 0xa7, 0x8d, 0xc1, 0x72, 0x97, 0x73, 0x65, 0x1b,
	0x11, 0x0c, 0xca, 0xba, 0x56, 0x28, 0x5b, 0xd7, 0xa0, 0xd0, 0xba, 0x3b, 0xaf, 0x4a, 0x68, 0xa4,
	0x17, 0xf9, 0xce, 0xd4, 0x5a, 0x6c, 0xac, 0x08, 0x7f, 0xe2, 0x45, 0x9f, 0xcf, 0xc7, 0x41, 0x3f,
	0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0x22, 0x86, 0xe3, 0x57, 0xba, 0x01, 0x00, 0x00,
}
