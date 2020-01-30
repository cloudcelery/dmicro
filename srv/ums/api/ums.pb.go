// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv/ums/api/ums.proto

package go_micro_srv_ums

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

type G2LRequest struct {
	Proto                int32    `protobuf:"varint,1,opt,name=proto,proto3" json:"proto,omitempty"`
	Appid                int32    `protobuf:"varint,2,opt,name=appid,proto3" json:"appid,omitempty"`
	Uid                  uint64   `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Platform             int32    `protobuf:"varint,4,opt,name=platform,proto3" json:"platform,omitempty"`
	Type                 int32    `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	Cmd                  int32    `protobuf:"varint,6,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Seq                  int32    `protobuf:"varint,7,opt,name=seq,proto3" json:"seq,omitempty"`
	Payload              []byte   `protobuf:"bytes,8,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *G2LRequest) Reset()         { *m = G2LRequest{} }
func (m *G2LRequest) String() string { return proto.CompactTextString(m) }
func (*G2LRequest) ProtoMessage()    {}
func (*G2LRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08f47cff16da6c04, []int{0}
}

func (m *G2LRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_G2LRequest.Unmarshal(m, b)
}
func (m *G2LRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_G2LRequest.Marshal(b, m, deterministic)
}
func (m *G2LRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_G2LRequest.Merge(m, src)
}
func (m *G2LRequest) XXX_Size() int {
	return xxx_messageInfo_G2LRequest.Size(m)
}
func (m *G2LRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_G2LRequest.DiscardUnknown(m)
}

var xxx_messageInfo_G2LRequest proto.InternalMessageInfo

func (m *G2LRequest) GetProto() int32 {
	if m != nil {
		return m.Proto
	}
	return 0
}

func (m *G2LRequest) GetAppid() int32 {
	if m != nil {
		return m.Appid
	}
	return 0
}

func (m *G2LRequest) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *G2LRequest) GetPlatform() int32 {
	if m != nil {
		return m.Platform
	}
	return 0
}

func (m *G2LRequest) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *G2LRequest) GetCmd() int32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *G2LRequest) GetSeq() int32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *G2LRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type G2LResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *G2LResponse) Reset()         { *m = G2LResponse{} }
func (m *G2LResponse) String() string { return proto.CompactTextString(m) }
func (*G2LResponse) ProtoMessage()    {}
func (*G2LResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_08f47cff16da6c04, []int{1}
}

func (m *G2LResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_G2LResponse.Unmarshal(m, b)
}
func (m *G2LResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_G2LResponse.Marshal(b, m, deterministic)
}
func (m *G2LResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_G2LResponse.Merge(m, src)
}
func (m *G2LResponse) XXX_Size() int {
	return xxx_messageInfo_G2LResponse.Size(m)
}
func (m *G2LResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_G2LResponse.DiscardUnknown(m)
}

var xxx_messageInfo_G2LResponse proto.InternalMessageInfo

type A2LRequest struct {
	Proto                int32    `protobuf:"varint,1,opt,name=proto,proto3" json:"proto,omitempty"`
	Appid                int32    `protobuf:"varint,2,opt,name=appid,proto3" json:"appid,omitempty"`
	Uid                  uint64   `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Platform             int32    `protobuf:"varint,4,opt,name=platform,proto3" json:"platform,omitempty"`
	Type                 int32    `protobuf:"varint,5,opt,name=type,proto3" json:"type,omitempty"`
	Cmd                  int32    `protobuf:"varint,6,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Seq                  int32    `protobuf:"varint,7,opt,name=seq,proto3" json:"seq,omitempty"`
	Payload              []byte   `protobuf:"bytes,8,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *A2LRequest) Reset()         { *m = A2LRequest{} }
func (m *A2LRequest) String() string { return proto.CompactTextString(m) }
func (*A2LRequest) ProtoMessage()    {}
func (*A2LRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08f47cff16da6c04, []int{2}
}

func (m *A2LRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_A2LRequest.Unmarshal(m, b)
}
func (m *A2LRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_A2LRequest.Marshal(b, m, deterministic)
}
func (m *A2LRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_A2LRequest.Merge(m, src)
}
func (m *A2LRequest) XXX_Size() int {
	return xxx_messageInfo_A2LRequest.Size(m)
}
func (m *A2LRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_A2LRequest.DiscardUnknown(m)
}

var xxx_messageInfo_A2LRequest proto.InternalMessageInfo

func (m *A2LRequest) GetProto() int32 {
	if m != nil {
		return m.Proto
	}
	return 0
}

func (m *A2LRequest) GetAppid() int32 {
	if m != nil {
		return m.Appid
	}
	return 0
}

func (m *A2LRequest) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *A2LRequest) GetPlatform() int32 {
	if m != nil {
		return m.Platform
	}
	return 0
}

func (m *A2LRequest) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *A2LRequest) GetCmd() int32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *A2LRequest) GetSeq() int32 {
	if m != nil {
		return m.Seq
	}
	return 0
}

func (m *A2LRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type A2LResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *A2LResponse) Reset()         { *m = A2LResponse{} }
func (m *A2LResponse) String() string { return proto.CompactTextString(m) }
func (*A2LResponse) ProtoMessage()    {}
func (*A2LResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_08f47cff16da6c04, []int{3}
}

func (m *A2LResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_A2LResponse.Unmarshal(m, b)
}
func (m *A2LResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_A2LResponse.Marshal(b, m, deterministic)
}
func (m *A2LResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_A2LResponse.Merge(m, src)
}
func (m *A2LResponse) XXX_Size() int {
	return xxx_messageInfo_A2LResponse.Size(m)
}
func (m *A2LResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_A2LResponse.DiscardUnknown(m)
}

var xxx_messageInfo_A2LResponse proto.InternalMessageInfo

// appserver -> ums
type PushRequest struct {
	Typ                  int32    `protobuf:"varint,1,opt,name=typ,proto3" json:"typ,omitempty"`
	Appid                int32    `protobuf:"varint,2,opt,name=appid,proto3" json:"appid,omitempty"`
	Uid                  uint64   `protobuf:"varint,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Cmd                  int32    `protobuf:"varint,4,opt,name=cmd,proto3" json:"cmd,omitempty"`
	Payload              []byte   `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushRequest) Reset()         { *m = PushRequest{} }
func (m *PushRequest) String() string { return proto.CompactTextString(m) }
func (*PushRequest) ProtoMessage()    {}
func (*PushRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08f47cff16da6c04, []int{4}
}

func (m *PushRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushRequest.Unmarshal(m, b)
}
func (m *PushRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushRequest.Marshal(b, m, deterministic)
}
func (m *PushRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushRequest.Merge(m, src)
}
func (m *PushRequest) XXX_Size() int {
	return xxx_messageInfo_PushRequest.Size(m)
}
func (m *PushRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PushRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PushRequest proto.InternalMessageInfo

func (m *PushRequest) GetTyp() int32 {
	if m != nil {
		return m.Typ
	}
	return 0
}

func (m *PushRequest) GetAppid() int32 {
	if m != nil {
		return m.Appid
	}
	return 0
}

func (m *PushRequest) GetUid() uint64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *PushRequest) GetCmd() int32 {
	if m != nil {
		return m.Cmd
	}
	return 0
}

func (m *PushRequest) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type PushResponse struct {
	Errno                int32    `protobuf:"varint,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PushResponse) Reset()         { *m = PushResponse{} }
func (m *PushResponse) String() string { return proto.CompactTextString(m) }
func (*PushResponse) ProtoMessage()    {}
func (*PushResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_08f47cff16da6c04, []int{5}
}

func (m *PushResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushResponse.Unmarshal(m, b)
}
func (m *PushResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushResponse.Marshal(b, m, deterministic)
}
func (m *PushResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushResponse.Merge(m, src)
}
func (m *PushResponse) XXX_Size() int {
	return xxx_messageInfo_PushResponse.Size(m)
}
func (m *PushResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PushResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PushResponse proto.InternalMessageInfo

func (m *PushResponse) GetErrno() int32 {
	if m != nil {
		return m.Errno
	}
	return 0
}

func (m *PushResponse) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func init() {
	proto.RegisterType((*G2LRequest)(nil), "go.micro.srv.ums.G2LRequest")
	proto.RegisterType((*G2LResponse)(nil), "go.micro.srv.ums.G2LResponse")
	proto.RegisterType((*A2LRequest)(nil), "go.micro.srv.ums.A2LRequest")
	proto.RegisterType((*A2LResponse)(nil), "go.micro.srv.ums.A2LResponse")
	proto.RegisterType((*PushRequest)(nil), "go.micro.srv.ums.PushRequest")
	proto.RegisterType((*PushResponse)(nil), "go.micro.srv.ums.PushResponse")
}

func init() { proto.RegisterFile("srv/ums/api/ums.proto", fileDescriptor_08f47cff16da6c04) }

var fileDescriptor_08f47cff16da6c04 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x52, 0x4d, 0x6b, 0xc2, 0x40,
	0x10, 0x35, 0x4d, 0xa2, 0x76, 0xb4, 0x20, 0x43, 0x5b, 0x16, 0xa9, 0x45, 0x3c, 0x79, 0x5a, 0xc1,
	0x5e, 0x7b, 0x09, 0x14, 0xbc, 0x78, 0x28, 0x81, 0xfe, 0x80, 0xd4, 0x6c, 0x6d, 0xc0, 0x75, 0xd7,
	0x9d, 0x44, 0xf0, 0x9f, 0xf5, 0xb7, 0xf4, 0xd7, 0x94, 0xdd, 0x44, 0x4d, 0x3f, 0x2c, 0xbd, 0xf6,
	0x94, 0x79, 0x6f, 0x87, 0x37, 0xef, 0x4d, 0x06, 0xae, 0xc8, 0x6c, 0x27, 0x85, 0xa4, 0x49, 0xa2,
	0x33, 0xfb, 0xe5, 0xda, 0xa8, 0x5c, 0x61, 0x6f, 0xa9, 0xb8, 0xcc, 0x16, 0x46, 0x71, 0x32, 0x5b,
	0x5e, 0x48, 0x1a, 0xbd, 0x79, 0x00, 0xb3, 0xe9, 0x3c, 0x16, 0x9b, 0x42, 0x50, 0x8e, 0x97, 0x10,
	0xba, 0x4e, 0xe6, 0x0d, 0xbd, 0x71, 0x18, 0x97, 0xc0, 0xb2, 0x89, 0xd6, 0x59, 0xca, 0xce, 0x4a,
	0xd6, 0x01, 0xec, 0x81, 0x5f, 0x64, 0x29, 0xf3, 0x87, 0xde, 0x38, 0x88, 0x6d, 0x89, 0x7d, 0x68,
	0xeb, 0x55, 0x92, 0xbf, 0x28, 0x23, 0x59, 0xe0, 0x5a, 0x0f, 0x18, 0x11, 0x82, 0x7c, 0xa7, 0x05,
	0x0b, 0x1d, 0xef, 0x6a, 0xab, 0xb0, 0x90, 0x29, 0x6b, 0x3a, 0xca, 0x96, 0x96, 0x21, 0xb1, 0x61,
	0xad, 0x92, 0x21, 0xb1, 0x41, 0x06, 0x2d, 0x9d, 0xec, 0x56, 0x2a, 0x49, 0x59, 0x7b, 0xe8, 0x8d,
	0xbb, 0xf1, 0x1e, 0x8e, 0x2e, 0xa0, 0xe3, 0x9c, 0x93, 0x56, 0x6b, 0x12, 0x2e, 0x49, 0xf4, 0x6f,
	0x93, 0x44, 0xb5, 0x24, 0x04, 0x9d, 0xc7, 0x82, 0x5e, 0xf7, 0x49, 0x7a, 0xe0, 0xe7, 0x3b, 0x5d,
	0xe5, 0xb0, 0xe5, 0x9f, 0x53, 0x54, 0xae, 0x82, 0xa3, 0xab, 0x9a, 0x87, 0xf0, 0xb3, 0x87, 0x7b,
	0xe8, 0x96, 0x43, 0x4b, 0x13, 0x76, 0x86, 0x30, 0x66, 0x7d, 0xd8, 0x9f, 0x03, 0x78, 0x0d, 0x4d,
	0x61, 0x8c, 0xa4, 0xa5, 0x1b, 0x7d, 0x1e, 0x57, 0x68, 0xfa, 0xee, 0x81, 0xff, 0x24, 0x09, 0x67,
	0x10, 0x58, 0x15, 0x1c, 0xf0, 0xaf, 0x97, 0xc6, 0x6b, 0x91, 0xfa, 0xb7, 0xa7, 0x9e, 0xab, 0x0d,
	0x34, 0xf0, 0x01, 0xfc, 0xd9, 0x74, 0x8e, 0x37, 0xdf, 0x1b, 0x8f, 0xd7, 0xda, 0x1f, 0x9c, 0x78,
	0xad, 0xab, 0x44, 0x3f, 0xab, 0x44, 0xbf, 0xaa, 0xd4, 0xff, 0x46, 0xe3, 0xb9, 0xe9, 0x6e, 0xe7,
	0xee, 0x23, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x57, 0xec, 0x82, 0x55, 0x03, 0x00, 0x00,
}
