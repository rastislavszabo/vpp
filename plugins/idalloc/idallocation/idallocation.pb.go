// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: idallocation.proto

package idallocation

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// IDAllocation represents a VXLAN VNI allocation made for the specified unique VXLAN name
type AllocationPool struct {
	Name  string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Range *AllocationPool_Range `protobuf:"bytes,2,opt,name=range,proto3" json:"range,omitempty"`
	// map of all allocations, key is the allocation "label"
	// describing its purpose (e.g. nework name for vrf pool, etc.)
	IdAllocations        map[string]uint32 `protobuf:"bytes,3,rep,name=id_allocations,json=idAllocations,proto3" json:"id_allocations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AllocationPool) Reset()         { *m = AllocationPool{} }
func (m *AllocationPool) String() string { return proto.CompactTextString(m) }
func (*AllocationPool) ProtoMessage()    {}
func (*AllocationPool) Descriptor() ([]byte, []int) {
	return fileDescriptor_4826bf6deca43300, []int{0}
}
func (m *AllocationPool) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllocationPool.Unmarshal(m, b)
}
func (m *AllocationPool) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllocationPool.Marshal(b, m, deterministic)
}
func (m *AllocationPool) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocationPool.Merge(m, src)
}
func (m *AllocationPool) XXX_Size() int {
	return xxx_messageInfo_AllocationPool.Size(m)
}
func (m *AllocationPool) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocationPool.DiscardUnknown(m)
}

var xxx_messageInfo_AllocationPool proto.InternalMessageInfo

func (m *AllocationPool) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *AllocationPool) GetRange() *AllocationPool_Range {
	if m != nil {
		return m.Range
	}
	return nil
}

func (m *AllocationPool) GetIdAllocations() map[string]uint32 {
	if m != nil {
		return m.IdAllocations
	}
	return nil
}

type AllocationPool_Range struct {
	MinId                uint32   `protobuf:"varint,1,opt,name=min_id,json=minId,proto3" json:"min_id,omitempty"`
	MaxId                uint32   `protobuf:"varint,2,opt,name=max_id,json=maxId,proto3" json:"max_id,omitempty"`
	Reserved             []uint32 `protobuf:"varint,3,rep,packed,name=reserved,proto3" json:"reserved,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AllocationPool_Range) Reset()         { *m = AllocationPool_Range{} }
func (m *AllocationPool_Range) String() string { return proto.CompactTextString(m) }
func (*AllocationPool_Range) ProtoMessage()    {}
func (*AllocationPool_Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_4826bf6deca43300, []int{0, 0}
}
func (m *AllocationPool_Range) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AllocationPool_Range.Unmarshal(m, b)
}
func (m *AllocationPool_Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AllocationPool_Range.Marshal(b, m, deterministic)
}
func (m *AllocationPool_Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllocationPool_Range.Merge(m, src)
}
func (m *AllocationPool_Range) XXX_Size() int {
	return xxx_messageInfo_AllocationPool_Range.Size(m)
}
func (m *AllocationPool_Range) XXX_DiscardUnknown() {
	xxx_messageInfo_AllocationPool_Range.DiscardUnknown(m)
}

var xxx_messageInfo_AllocationPool_Range proto.InternalMessageInfo

func (m *AllocationPool_Range) GetMinId() uint32 {
	if m != nil {
		return m.MinId
	}
	return 0
}

func (m *AllocationPool_Range) GetMaxId() uint32 {
	if m != nil {
		return m.MaxId
	}
	return 0
}

func (m *AllocationPool_Range) GetReserved() []uint32 {
	if m != nil {
		return m.Reserved
	}
	return nil
}

func init() {
	proto.RegisterType((*AllocationPool)(nil), "idallocation.AllocationPool")
	proto.RegisterMapType((map[string]uint32)(nil), "idallocation.AllocationPool.IdAllocationsEntry")
	proto.RegisterType((*AllocationPool_Range)(nil), "idallocation.AllocationPool.Range")
}

func init() { proto.RegisterFile("idallocation.proto", fileDescriptor_4826bf6deca43300) }

var fileDescriptor_4826bf6deca43300 = []byte{
	// 238 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xcf, 0x4a, 0xc4, 0x30,
	0x10, 0x87, 0x49, 0x6a, 0x16, 0x9d, 0xb5, 0x8b, 0x0c, 0x0a, 0xa5, 0xa7, 0xb2, 0xa7, 0x9e, 0x2a,
	0xac, 0x97, 0xc5, 0x93, 0x1e, 0x3c, 0xf4, 0xa6, 0x39, 0x78, 0x5d, 0xa2, 0x09, 0x12, 0x4c, 0x13,
	0x49, 0xeb, 0xd2, 0x3e, 0xa3, 0x2f, 0x25, 0x8d, 0xb5, 0x7f, 0x10, 0xbc, 0xcd, 0xef, 0x63, 0xe6,
	0x9b, 0x61, 0x00, 0xb5, 0x14, 0xc6, 0xb8, 0x57, 0xd1, 0x68, 0x67, 0x8b, 0x0f, 0xef, 0x1a, 0x87,
	0xe7, 0x73, 0xb6, 0xfd, 0xa2, 0xb0, 0xb9, 0x1f, 0xe3, 0xa3, 0x73, 0x06, 0x11, 0x4e, 0xac, 0xa8,
	0x54, 0x42, 0x32, 0x92, 0x9f, 0xf1, 0x50, 0xe3, 0x1e, 0x98, 0x17, 0xf6, 0x4d, 0x25, 0x34, 0x23,
	0xf9, 0x7a, 0xb7, 0x2d, 0x16, 0xe2, 0xa5, 0xa0, 0xe0, 0x7d, 0x27, 0xff, 0x19, 0xc0, 0x67, 0xd8,
	0x68, 0x79, 0x98, 0x9a, 0xeb, 0x24, 0xca, 0xa2, 0x7c, 0xbd, 0xbb, 0xfe, 0x57, 0x51, 0xca, 0x09,
	0xd4, 0x0f, 0xb6, 0xf1, 0x1d, 0x8f, 0xf5, 0x9c, 0xa5, 0x4f, 0xc0, 0xc2, 0x1e, 0xbc, 0x82, 0x55,
	0xa5, 0xed, 0x41, 0xcb, 0x70, 0x70, 0xcc, 0x59, 0xa5, 0x6d, 0x29, 0x03, 0x16, 0x6d, 0x8f, 0xe9,
	0x80, 0x45, 0x5b, 0x4a, 0x4c, 0xe1, 0xd4, 0xab, 0x5a, 0xf9, 0xa3, 0x92, 0xe1, 0x90, 0x98, 0x8f,
	0x39, 0xbd, 0x03, 0xfc, 0xbb, 0x17, 0x2f, 0x20, 0x7a, 0x57, 0xdd, 0xf0, 0x8d, 0xbe, 0xc4, 0x4b,
	0x60, 0x47, 0x61, 0x3e, 0xd5, 0xaf, 0x39, 0x84, 0x5b, 0xba, 0x27, 0x2f, 0xab, 0xf0, 0xe2, 0x9b,
	0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa2, 0x40, 0x1d, 0xe2, 0x78, 0x01, 0x00, 0x00,
}