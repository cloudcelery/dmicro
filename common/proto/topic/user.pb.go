// Code generated by protoc-gen-go. DO NOT EDIT.
// source: common/proto/topic/user.proto

package dmicro_topic

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

type UserInfo struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Mobile               string   `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserInfo) Reset()         { *m = UserInfo{} }
func (m *UserInfo) String() string { return proto.CompactTextString(m) }
func (*UserInfo) ProtoMessage()    {}
func (*UserInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_85bec14ef9159631, []int{0}
}

func (m *UserInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserInfo.Unmarshal(m, b)
}
func (m *UserInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserInfo.Marshal(b, m, deterministic)
}
func (m *UserInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserInfo.Merge(m, src)
}
func (m *UserInfo) XXX_Size() int {
	return xxx_messageInfo_UserInfo.Size(m)
}
func (m *UserInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_UserInfo.DiscardUnknown(m)
}

var xxx_messageInfo_UserInfo proto.InternalMessageInfo

func (m *UserInfo) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *UserInfo) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

type UserCreated struct {
	Id                   int64     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Topic                string    `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
	Info                 *UserInfo `protobuf:"bytes,3,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *UserCreated) Reset()         { *m = UserCreated{} }
func (m *UserCreated) String() string { return proto.CompactTextString(m) }
func (*UserCreated) ProtoMessage()    {}
func (*UserCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_85bec14ef9159631, []int{1}
}

func (m *UserCreated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCreated.Unmarshal(m, b)
}
func (m *UserCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCreated.Marshal(b, m, deterministic)
}
func (m *UserCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCreated.Merge(m, src)
}
func (m *UserCreated) XXX_Size() int {
	return xxx_messageInfo_UserCreated.Size(m)
}
func (m *UserCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCreated.DiscardUnknown(m)
}

var xxx_messageInfo_UserCreated proto.InternalMessageInfo

func (m *UserCreated) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserCreated) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *UserCreated) GetInfo() *UserInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*UserInfo)(nil), "dmicro.topic.UserInfo")
	proto.RegisterType((*UserCreated)(nil), "dmicro.topic.UserCreated")
}

func init() { proto.RegisterFile("common/proto/topic/user.proto", fileDescriptor_85bec14ef9159631) }

var fileDescriptor_85bec14ef9159631 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0xb1, 0x0b, 0x82, 0x40,
	0x14, 0x87, 0x39, 0x2d, 0xab, 0x67, 0x34, 0x3c, 0xc2, 0x5c, 0x02, 0x71, 0x92, 0x86, 0x13, 0x6a,
	0x6c, 0x6c, 0x72, 0x15, 0x9a, 0x25, 0xbd, 0x13, 0x1e, 0x74, 0x3e, 0x39, 0xed, 0xff, 0x0f, 0x2f,
	0x8b, 0xc6, 0xef, 0xee, 0xf7, 0xf1, 0x3e, 0x38, 0x36, 0x6c, 0x0c, 0x77, 0x79, 0x6f, 0x79, 0xe4,
	0x7c, 0xe4, 0x9e, 0x9a, 0xfc, 0x35, 0x68, 0x2b, 0xdd, 0x03, 0x6e, 0x95, 0xa1, 0xc6, 0xb2, 0x74,
	0x1f, 0xe9, 0x15, 0xd6, 0xf7, 0x41, 0xdb, 0xa2, 0x6b, 0x19, 0x0f, 0xb0, 0x9a, 0x76, 0x15, 0xa9,
	0x58, 0x24, 0x22, 0xf3, 0xcb, 0x60, 0xc2, 0x42, 0x61, 0x04, 0x81, 0xe1, 0x9a, 0x9e, 0x3a, 0xf6,
	0x12, 0x91, 0x6d, 0xca, 0x99, 0xd2, 0x0a, 0xc2, 0x49, 0xbe, 0x59, 0xfd, 0x18, 0xb5, 0xc2, 0x1d,
	0x78, 0x3f, 0xd5, 0x23, 0x85, 0x7b, 0x58, 0xba, 0x23, 0xb3, 0xf5, 0x01, 0x3c, 0xc1, 0x82, 0xba,
	0x96, 0x63, 0x3f, 0x11, 0x59, 0x78, 0x8e, 0xe4, 0x7f, 0x8e, 0xfc, 0xb6, 0x94, 0x6e, 0x53, 0x07,
	0x2e, 0xf9, 0xf2, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x42, 0xe9, 0x3d, 0xa1, 0xd3, 0x00, 0x00, 0x00,
}