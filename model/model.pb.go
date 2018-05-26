// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model.proto

package model

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CallRecord struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	StartTime            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=StartTime" json:"StartTime,omitempty"`
	EndTime              *timestamp.Timestamp `protobuf:"bytes,3,opt,name=EndTime" json:"EndTime,omitempty"`
	Duration             *timestamp.Timestamp `protobuf:"bytes,4,opt,name=Duration" json:"Duration,omitempty"`
	CallerLineIdentity   string               `protobuf:"bytes,5,opt,name=callerLineIdentity" json:"callerLineIdentity,omitempty"`
	NonChargedParty      string               `protobuf:"bytes,6,opt,name=NonChargedParty" json:"NonChargedParty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CallRecord) Reset()         { *m = CallRecord{} }
func (m *CallRecord) String() string { return proto.CompactTextString(m) }
func (*CallRecord) ProtoMessage()    {}
func (*CallRecord) Descriptor() ([]byte, []int) {
	return fileDescriptor_model_c14e3d1453919291, []int{0}
}
func (m *CallRecord) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CallRecord.Unmarshal(m, b)
}
func (m *CallRecord) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CallRecord.Marshal(b, m, deterministic)
}
func (dst *CallRecord) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CallRecord.Merge(dst, src)
}
func (m *CallRecord) XXX_Size() int {
	return xxx_messageInfo_CallRecord.Size(m)
}
func (m *CallRecord) XXX_DiscardUnknown() {
	xxx_messageInfo_CallRecord.DiscardUnknown(m)
}

var xxx_messageInfo_CallRecord proto.InternalMessageInfo

func (m *CallRecord) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *CallRecord) GetStartTime() *timestamp.Timestamp {
	if m != nil {
		return m.StartTime
	}
	return nil
}

func (m *CallRecord) GetEndTime() *timestamp.Timestamp {
	if m != nil {
		return m.EndTime
	}
	return nil
}

func (m *CallRecord) GetDuration() *timestamp.Timestamp {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *CallRecord) GetCallerLineIdentity() string {
	if m != nil {
		return m.CallerLineIdentity
	}
	return ""
}

func (m *CallRecord) GetNonChargedParty() string {
	if m != nil {
		return m.NonChargedParty
	}
	return ""
}

func init() {
	proto.RegisterType((*CallRecord)(nil), "model.CallRecord")
}

func init() { proto.RegisterFile("model.proto", fileDescriptor_model_c14e3d1453919291) }

var fileDescriptor_model_c14e3d1453919291 = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8d, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0xd9, 0xd5, 0xad, 0x76, 0x0a, 0x0a, 0x73, 0x0a, 0xbd, 0xb8, 0x78, 0xda, 0x53, 0x0a,
	0x2a, 0xe2, 0xbd, 0x7a, 0x10, 0x44, 0x24, 0xf6, 0x05, 0xd2, 0x66, 0x5c, 0x03, 0x49, 0xa6, 0xc4,
	0xe9, 0xa1, 0x2f, 0xe3, 0xb3, 0x0a, 0x29, 0xab, 0x20, 0x85, 0x3d, 0xe6, 0xcf, 0xf7, 0x7d, 0x03,
	0xb3, 0xc8, 0x8e, 0x82, 0xde, 0x66, 0x16, 0xc6, 0xa6, 0x3c, 0xe6, 0x57, 0x3d, 0x73, 0x1f, 0x68,
	0x51, 0xc6, 0xf5, 0xee, 0x63, 0x21, 0x3e, 0xd2, 0x97, 0xd8, 0xb8, 0x3d, 0x70, 0xd7, 0xdf, 0x35,
	0xc0, 0xd2, 0x86, 0x60, 0x68, 0xc3, 0xd9, 0xe1, 0x05, 0xd4, 0xde, 0xa9, 0xaa, 0xad, 0xba, 0xc6,
	0xd4, 0xde, 0xe1, 0x03, 0x4c, 0xdf, 0xc5, 0x66, 0x59, 0xf9, 0x48, 0xaa, 0x6e, 0xab, 0x6e, 0x76,
	0x33, 0xd7, 0x87, 0xa6, 0x1e, 0x9a, 0x7a, 0x35, 0x34, 0xcd, 0x1f, 0x8c, 0x77, 0x70, 0xf6, 0x94,
	0x5c, 0xf1, 0x4e, 0x46, 0xbd, 0x01, 0xc5, 0x7b, 0x38, 0x7f, 0xdc, 0x65, 0x2b, 0x9e, 0x93, 0x3a,
	0x1d, 0xd5, 0x7e, 0x59, 0xd4, 0x80, 0x1b, 0x1b, 0x02, 0xe5, 0x17, 0x9f, 0xe8, 0xd9, 0x51, 0x12,
	0x2f, 0x7b, 0xd5, 0xb4, 0x55, 0x37, 0x35, 0x47, 0x7e, 0xb0, 0x83, 0xcb, 0x57, 0x4e, 0xcb, 0x4f,
	0x9b, 0x7b, 0x72, 0x6f, 0x36, 0xcb, 0x5e, 0x4d, 0x0a, 0xfc, 0x7f, 0x5e, 0x4f, 0xca, 0xdd, 0xdb,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc3, 0x1b, 0x1c, 0xca, 0x5e, 0x01, 0x00, 0x00,
}
