// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv/recorder/proto/recorder/recorder.proto

package recordersrv

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
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

type TransactionReadRequest struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionReadRequest) Reset()         { *m = TransactionReadRequest{} }
func (m *TransactionReadRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionReadRequest) ProtoMessage()    {}
func (*TransactionReadRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65991f949dbc3f7a, []int{0}
}

func (m *TransactionReadRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionReadRequest.Unmarshal(m, b)
}
func (m *TransactionReadRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionReadRequest.Marshal(b, m, deterministic)
}
func (m *TransactionReadRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionReadRequest.Merge(m, src)
}
func (m *TransactionReadRequest) XXX_Size() int {
	return xxx_messageInfo_TransactionReadRequest.Size(m)
}
func (m *TransactionReadRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionReadRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionReadRequest proto.InternalMessageInfo

func (m *TransactionReadRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type TransactionWriteRequest struct {
	Key                  string            `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Event                *TransactionEvent `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TransactionWriteRequest) Reset()         { *m = TransactionWriteRequest{} }
func (m *TransactionWriteRequest) String() string { return proto.CompactTextString(m) }
func (*TransactionWriteRequest) ProtoMessage()    {}
func (*TransactionWriteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65991f949dbc3f7a, []int{1}
}

func (m *TransactionWriteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionWriteRequest.Unmarshal(m, b)
}
func (m *TransactionWriteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionWriteRequest.Marshal(b, m, deterministic)
}
func (m *TransactionWriteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionWriteRequest.Merge(m, src)
}
func (m *TransactionWriteRequest) XXX_Size() int {
	return xxx_messageInfo_TransactionWriteRequest.Size(m)
}
func (m *TransactionWriteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionWriteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionWriteRequest proto.InternalMessageInfo

func (m *TransactionWriteRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *TransactionWriteRequest) GetEvent() *TransactionEvent {
	if m != nil {
		return m.Event
	}
	return nil
}

// Transaction Event
type TransactionEvent struct {
	// request
	Req []byte `protobuf:"bytes,1,opt,name=req,proto3" json:"req,omitempty"`
	// responce
	Rsp                  []byte   `protobuf:"bytes,2,opt,name=rsp,proto3" json:"rsp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransactionEvent) Reset()         { *m = TransactionEvent{} }
func (m *TransactionEvent) String() string { return proto.CompactTextString(m) }
func (*TransactionEvent) ProtoMessage()    {}
func (*TransactionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_65991f949dbc3f7a, []int{2}
}

func (m *TransactionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransactionEvent.Unmarshal(m, b)
}
func (m *TransactionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransactionEvent.Marshal(b, m, deterministic)
}
func (m *TransactionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransactionEvent.Merge(m, src)
}
func (m *TransactionEvent) XXX_Size() int {
	return xxx_messageInfo_TransactionEvent.Size(m)
}
func (m *TransactionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TransactionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TransactionEvent proto.InternalMessageInfo

func (m *TransactionEvent) GetReq() []byte {
	if m != nil {
		return m.Req
	}
	return nil
}

func (m *TransactionEvent) GetRsp() []byte {
	if m != nil {
		return m.Rsp
	}
	return nil
}

func init() {
	proto.RegisterType((*TransactionReadRequest)(nil), "recordersrv.TransactionReadRequest")
	proto.RegisterType((*TransactionWriteRequest)(nil), "recordersrv.TransactionWriteRequest")
	proto.RegisterType((*TransactionEvent)(nil), "recordersrv.TransactionEvent")
}

func init() {
	proto.RegisterFile("srv/recorder/proto/recorder/recorder.proto", fileDescriptor_65991f949dbc3f7a)
}

var fileDescriptor_65991f949dbc3f7a = []byte{
	// 267 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x6d, 0xac, 0x15, 0x9c, 0xf5, 0x50, 0x72, 0x68, 0xcb, 0x8a, 0x50, 0x56, 0x0f, 0xc5, 0x43,
	0x02, 0x2d, 0x78, 0xf0, 0x28, 0x14, 0x2f, 0x9e, 0x82, 0xe0, 0x39, 0xed, 0x8e, 0x65, 0xb1, 0x6e,
	0xb6, 0x93, 0x74, 0xa1, 0xff, 0xe3, 0x47, 0x7a, 0x94, 0x24, 0x2c, 0x06, 0x61, 0xbd, 0xbd, 0xbc,
	0x79, 0xef, 0xcd, 0x4c, 0x06, 0xee, 0x2d, 0xb5, 0x92, 0x70, 0x6b, 0xa8, 0x44, 0x92, 0x0d, 0x19,
	0x67, 0x7e, 0x9f, 0x1d, 0x10, 0x81, 0xe7, 0x59, 0xf7, 0xb6, 0xd4, 0xe6, 0xd7, 0x3b, 0x63, 0x76,
	0x7b, 0x8c, 0x96, 0xcd, 0xf1, 0x5d, 0xe2, 0x67, 0xe3, 0x4e, 0x51, 0x99, 0x4f, 0x5b, 0xbd, 0xaf,
	0x4a, 0xed, 0x50, 0x76, 0x20, 0x16, 0x0a, 0x09, 0x93, 0x57, 0xd2, 0xb5, 0xd5, 0x5b, 0x57, 0x99,
	0x5a, 0xa1, 0x2e, 0x15, 0x1e, 0x8e, 0x68, 0x1d, 0x1f, 0xc3, 0xf0, 0x03, 0x4f, 0x33, 0x36, 0x67,
	0x8b, 0x4b, 0xe5, 0xe1, 0xe3, 0xf0, 0xfb, 0x89, 0x15, 0x08, 0xd3, 0xc4, 0xf0, 0x46, 0x95, 0xc3,
	0x5e, 0x07, 0x5f, 0xc1, 0x08, 0x5b, 0xac, 0xdd, 0xec, 0x6c, 0xce, 0x16, 0xd9, 0xf2, 0x46, 0x24,
	0x03, 0x8b, 0x24, 0x66, 0xed, 0x45, 0x2a, 0x6a, 0x63, 0x9b, 0x07, 0x18, 0xff, 0xad, 0xfb, 0x7c,
	0xc2, 0x43, 0xc8, 0xbf, 0x52, 0x1e, 0x06, 0xc6, 0x36, 0x21, 0xdd, 0x33, 0xb6, 0x59, 0x7e, 0x31,
	0xc8, 0x12, 0x23, 0x7f, 0x81, 0x73, 0xbf, 0x14, 0xbf, 0xed, 0x6b, 0x9d, 0xac, 0x9c, 0xff, 0x3f,
	0x5f, 0x31, 0xe0, 0xcf, 0x30, 0x0a, 0x1b, 0xf3, 0xbb, 0x3e, 0x65, 0xfa, 0x21, 0xf9, 0x44, 0xc4,
	0x9b, 0x88, 0xee, 0x26, 0x62, 0xed, 0x6f, 0x52, 0x0c, 0x36, 0x17, 0x81, 0x59, 0xfd, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x5d, 0x0f, 0x1d, 0xf5, 0xee, 0x01, 0x00, 0x00,
}
