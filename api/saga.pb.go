// Code generated by protoc-gen-go. DO NOT EDIT.
// source: saga.proto

package hellogrpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
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

type Event struct {
	OperationId          string   `protobuf:"bytes,1,opt,name=operationId,proto3" json:"operationId,omitempty"`
	EventType            string   `protobuf:"bytes,2,opt,name=eventType,proto3" json:"eventType,omitempty"`
	Payload              *any.Any `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_9818be635ac82bc9, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetOperationId() string {
	if m != nil {
		return m.OperationId
	}
	return ""
}

func (m *Event) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *Event) GetPayload() *any.Any {
	if m != nil {
		return m.Payload
	}
	return nil
}

type CompensationRequest struct {
	OperationId          string   `protobuf:"bytes,1,opt,name=operationId,proto3" json:"operationId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompensationRequest) Reset()         { *m = CompensationRequest{} }
func (m *CompensationRequest) String() string { return proto.CompactTextString(m) }
func (*CompensationRequest) ProtoMessage()    {}
func (*CompensationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9818be635ac82bc9, []int{1}
}

func (m *CompensationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompensationRequest.Unmarshal(m, b)
}
func (m *CompensationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompensationRequest.Marshal(b, m, deterministic)
}
func (m *CompensationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompensationRequest.Merge(m, src)
}
func (m *CompensationRequest) XXX_Size() int {
	return xxx_messageInfo_CompensationRequest.Size(m)
}
func (m *CompensationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CompensationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CompensationRequest proto.InternalMessageInfo

func (m *CompensationRequest) GetOperationId() string {
	if m != nil {
		return m.OperationId
	}
	return ""
}

func init() {
	proto.RegisterType((*Event)(nil), "hellogrpc.Event")
	proto.RegisterType((*CompensationRequest)(nil), "hellogrpc.CompensationRequest")
}

func init() { proto.RegisterFile("saga.proto", fileDescriptor_9818be635ac82bc9) }

var fileDescriptor_9818be635ac82bc9 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x4e, 0x4c, 0x4f,
	0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xcc, 0x48, 0xcd, 0xc9, 0xc9, 0x4f, 0x2f, 0x2a,
	0x48, 0x96, 0x92, 0x4c, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x4b, 0x24, 0x95, 0xa6, 0xe9,
	0x27, 0xe6, 0x55, 0x42, 0x54, 0x29, 0x95, 0x73, 0xb1, 0xba, 0x96, 0xa5, 0xe6, 0x95, 0x08, 0x29,
	0x70, 0x71, 0xe7, 0x17, 0xa4, 0x16, 0x25, 0x96, 0x64, 0xe6, 0xe7, 0x79, 0xa6, 0x48, 0x30, 0x2a,
	0x30, 0x6a, 0x70, 0x06, 0x21, 0x0b, 0x09, 0xc9, 0x70, 0x71, 0xa6, 0x82, 0x94, 0x86, 0x54, 0x16,
	0xa4, 0x4a, 0x30, 0x81, 0xe5, 0x11, 0x02, 0x42, 0x7a, 0x5c, 0xec, 0x05, 0x89, 0x95, 0x39, 0xf9,
	0x89, 0x29, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x22, 0x7a, 0x10, 0x5b, 0xf5, 0x60, 0xb6,
	0xea, 0x39, 0xe6, 0x55, 0x06, 0xc1, 0x14, 0x29, 0x99, 0x73, 0x09, 0x3b, 0xe7, 0xe7, 0x16, 0xa4,
	0xe6, 0x15, 0x83, 0xcd, 0x0f, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x26, 0xc2, 0x19, 0x49, 0x6c, 0x60,
	0xf3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x78, 0x39, 0x7d, 0xd7, 0xec, 0x00, 0x00, 0x00,
}