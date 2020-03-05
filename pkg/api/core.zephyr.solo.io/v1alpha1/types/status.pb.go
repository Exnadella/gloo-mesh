// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/mesh-projects/api/core/v1alpha1/status.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

type ComputedStatus_Status int32

const (
	ComputedStatus_UNKNOWN          ComputedStatus_Status = 0
	ComputedStatus_ACCEPTED         ComputedStatus_Status = 1
	ComputedStatus_INVALID          ComputedStatus_Status = 2
	ComputedStatus_PROCESSING_ERROR ComputedStatus_Status = 3
	ComputedStatus_CONFLICT         ComputedStatus_Status = 4
)

var ComputedStatus_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "ACCEPTED",
	2: "INVALID",
	3: "PROCESSING_ERROR",
	4: "CONFLICT",
}

var ComputedStatus_Status_value = map[string]int32{
	"UNKNOWN":          0,
	"ACCEPTED":         1,
	"INVALID":          2,
	"PROCESSING_ERROR": 3,
	"CONFLICT":         4,
}

func (x ComputedStatus_Status) String() string {
	return proto.EnumName(ComputedStatus_Status_name, int32(x))
}

func (ComputedStatus_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_38a786f3ea3dc5b8, []int{0, 0}
}

// a status set by Service Mesh Hub while processing a resource
type ComputedStatus struct {
	Status ComputedStatus_Status `protobuf:"varint,1,opt,name=status,proto3,enum=core.zephyr.solo.io.ComputedStatus_Status" json:"status,omitempty"`
	// human-readable message to be surfaced to the user
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ComputedStatus) Reset()         { *m = ComputedStatus{} }
func (m *ComputedStatus) String() string { return proto.CompactTextString(m) }
func (*ComputedStatus) ProtoMessage()    {}
func (*ComputedStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_38a786f3ea3dc5b8, []int{0}
}
func (m *ComputedStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ComputedStatus.Unmarshal(m, b)
}
func (m *ComputedStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ComputedStatus.Marshal(b, m, deterministic)
}
func (m *ComputedStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ComputedStatus.Merge(m, src)
}
func (m *ComputedStatus) XXX_Size() int {
	return xxx_messageInfo_ComputedStatus.Size(m)
}
func (m *ComputedStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ComputedStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ComputedStatus proto.InternalMessageInfo

func (m *ComputedStatus) GetStatus() ComputedStatus_Status {
	if m != nil {
		return m.Status
	}
	return ComputedStatus_UNKNOWN
}

func (m *ComputedStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterEnum("core.zephyr.solo.io.ComputedStatus_Status", ComputedStatus_Status_name, ComputedStatus_Status_value)
	proto.RegisterType((*ComputedStatus)(nil), "core.zephyr.solo.io.ComputedStatus")
}

func init() {
	proto.RegisterFile("github.com/solo-io/mesh-projects/api/core/v1alpha1/status.proto", fileDescriptor_38a786f3ea3dc5b8)
}

var fileDescriptor_38a786f3ea3dc5b8 = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xb2, 0x4f, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x2f, 0xce, 0xcf, 0xc9, 0xd7, 0xcd, 0xcc, 0xd7, 0xcf,
	0x4d, 0x2d, 0xce, 0xd0, 0x2d, 0x28, 0xca, 0xcf, 0x4a, 0x4d, 0x2e, 0x29, 0xd6, 0x4f, 0x2c, 0xc8,
	0xd4, 0x4f, 0xce, 0x2f, 0x4a, 0xd5, 0x2f, 0x33, 0x4c, 0xcc, 0x29, 0xc8, 0x48, 0x34, 0xd4, 0x2f,
	0x2e, 0x49, 0x2c, 0x29, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x06, 0xc9, 0xe9,
	0x55, 0xa5, 0x16, 0x64, 0x54, 0x16, 0xe9, 0x81, 0x4c, 0xd0, 0xcb, 0xcc, 0x97, 0x12, 0x49, 0xcf,
	0x4f, 0xcf, 0x07, 0xcb, 0xeb, 0x83, 0x58, 0x10, 0xa5, 0x4a, 0x47, 0x18, 0xb9, 0xf8, 0x9c, 0xf3,
	0x73, 0x0b, 0x4a, 0x4b, 0x52, 0x53, 0x82, 0xc1, 0x66, 0x08, 0x39, 0x71, 0xb1, 0x41, 0x4c, 0x93,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x33, 0xd2, 0xd2, 0xc3, 0x62, 0x9c, 0x1e, 0xaa, 0x26, 0x3d, 0x08,
	0x15, 0x04, 0xd5, 0x29, 0x24, 0xc1, 0xc5, 0x9e, 0x9b, 0x5a, 0x5c, 0x9c, 0x98, 0x9e, 0x2a, 0xc1,
	0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe3, 0x2a, 0x85, 0x70, 0xb1, 0x41, 0xed, 0xe1, 0xe6, 0x62,
	0x0f, 0xf5, 0xf3, 0xf6, 0xf3, 0x0f, 0xf7, 0x13, 0x60, 0x10, 0xe2, 0xe1, 0xe2, 0x70, 0x74, 0x76,
	0x76, 0x0d, 0x08, 0x71, 0x75, 0x11, 0x60, 0x04, 0x49, 0x79, 0xfa, 0x85, 0x39, 0xfa, 0x78, 0xba,
	0x08, 0x30, 0x09, 0x89, 0x70, 0x09, 0x04, 0x04, 0xf9, 0x3b, 0xbb, 0x06, 0x07, 0x7b, 0xfa, 0xb9,
	0xc7, 0xbb, 0x06, 0x05, 0xf9, 0x07, 0x09, 0x30, 0x83, 0x34, 0x38, 0xfb, 0xfb, 0xb9, 0xf9, 0x78,
	0x3a, 0x87, 0x08, 0xb0, 0x38, 0x05, 0xae, 0x78, 0x24, 0xc7, 0x18, 0xe5, 0x4d, 0x30, 0xe0, 0x0a,
	0xb2, 0xd3, 0xe1, 0x81, 0x87, 0xe6, 0x23, 0x44, 0x58, 0x96, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1,
	0x81, 0x03, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x15, 0x9b, 0x45, 0x51, 0x8e, 0x01, 0x00,
	0x00,
}

func (this *ComputedStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ComputedStatus)
	if !ok {
		that2, ok := that.(ComputedStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Status != that1.Status {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
