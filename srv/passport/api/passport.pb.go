// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv/passport/api/passport.proto

package go_micro_srv_passport

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

type TokenInfo struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	RefreshToken         string   `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	ExpiresAt            int64    `protobuf:"varint,3,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TokenInfo) Reset()         { *m = TokenInfo{} }
func (m *TokenInfo) String() string { return proto.CompactTextString(m) }
func (*TokenInfo) ProtoMessage()    {}
func (*TokenInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e4d00cde15d6b1e, []int{0}
}

func (m *TokenInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenInfo.Unmarshal(m, b)
}
func (m *TokenInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenInfo.Marshal(b, m, deterministic)
}
func (m *TokenInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenInfo.Merge(m, src)
}
func (m *TokenInfo) XXX_Size() int {
	return xxx_messageInfo_TokenInfo.Size(m)
}
func (m *TokenInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TokenInfo proto.InternalMessageInfo

func (m *TokenInfo) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *TokenInfo) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *TokenInfo) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

type Request struct {
	Mobile               string   `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e4d00cde15d6b1e, []int{1}
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

func (m *Request) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

type Response struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e4d00cde15d6b1e, []int{2}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type SmsLoginRequest struct {
	Appid                int32    `protobuf:"varint,1,opt,name=appid,proto3" json:"appid,omitempty"`
	Plat                 int32    `protobuf:"varint,2,opt,name=plat,proto3" json:"plat,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Code                 string   `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SmsLoginRequest) Reset()         { *m = SmsLoginRequest{} }
func (m *SmsLoginRequest) String() string { return proto.CompactTextString(m) }
func (*SmsLoginRequest) ProtoMessage()    {}
func (*SmsLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e4d00cde15d6b1e, []int{3}
}

func (m *SmsLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmsLoginRequest.Unmarshal(m, b)
}
func (m *SmsLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmsLoginRequest.Marshal(b, m, deterministic)
}
func (m *SmsLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmsLoginRequest.Merge(m, src)
}
func (m *SmsLoginRequest) XXX_Size() int {
	return xxx_messageInfo_SmsLoginRequest.Size(m)
}
func (m *SmsLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SmsLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SmsLoginRequest proto.InternalMessageInfo

func (m *SmsLoginRequest) GetAppid() int32 {
	if m != nil {
		return m.Appid
	}
	return 0
}

func (m *SmsLoginRequest) GetPlat() int32 {
	if m != nil {
		return m.Plat
	}
	return 0
}

func (m *SmsLoginRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *SmsLoginRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type SmsLoginResponse struct {
	TokenInfo            *TokenInfo `protobuf:"bytes,1,opt,name=token_info,json=tokenInfo,proto3" json:"token_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SmsLoginResponse) Reset()         { *m = SmsLoginResponse{} }
func (m *SmsLoginResponse) String() string { return proto.CompactTextString(m) }
func (*SmsLoginResponse) ProtoMessage()    {}
func (*SmsLoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e4d00cde15d6b1e, []int{4}
}

func (m *SmsLoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SmsLoginResponse.Unmarshal(m, b)
}
func (m *SmsLoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SmsLoginResponse.Marshal(b, m, deterministic)
}
func (m *SmsLoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SmsLoginResponse.Merge(m, src)
}
func (m *SmsLoginResponse) XXX_Size() int {
	return xxx_messageInfo_SmsLoginResponse.Size(m)
}
func (m *SmsLoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SmsLoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SmsLoginResponse proto.InternalMessageInfo

func (m *SmsLoginResponse) GetTokenInfo() *TokenInfo {
	if m != nil {
		return m.TokenInfo
	}
	return nil
}

type LoginRequest struct {
	Appid                int32    `protobuf:"varint,1,opt,name=appid,proto3" json:"appid,omitempty"`
	Plat                 int32    `protobuf:"varint,2,opt,name=plat,proto3" json:"plat,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Passwd               string   `protobuf:"bytes,4,opt,name=passwd,proto3" json:"passwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e4d00cde15d6b1e, []int{5}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetAppid() int32 {
	if m != nil {
		return m.Appid
	}
	return 0
}

func (m *LoginRequest) GetPlat() int32 {
	if m != nil {
		return m.Plat
	}
	return 0
}

func (m *LoginRequest) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *LoginRequest) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

type LoginResponse struct {
	TokenInfo            *TokenInfo `protobuf:"bytes,1,opt,name=token_info,json=tokenInfo,proto3" json:"token_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e4d00cde15d6b1e, []int{6}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetTokenInfo() *TokenInfo {
	if m != nil {
		return m.TokenInfo
	}
	return nil
}

type OAuthLoginRequest struct {
	Platform             string   `protobuf:"bytes,1,opt,name=platform,proto3" json:"platform,omitempty"`
	Code                 string   `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OAuthLoginRequest) Reset()         { *m = OAuthLoginRequest{} }
func (m *OAuthLoginRequest) String() string { return proto.CompactTextString(m) }
func (*OAuthLoginRequest) ProtoMessage()    {}
func (*OAuthLoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e4d00cde15d6b1e, []int{7}
}

func (m *OAuthLoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthLoginRequest.Unmarshal(m, b)
}
func (m *OAuthLoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthLoginRequest.Marshal(b, m, deterministic)
}
func (m *OAuthLoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthLoginRequest.Merge(m, src)
}
func (m *OAuthLoginRequest) XXX_Size() int {
	return xxx_messageInfo_OAuthLoginRequest.Size(m)
}
func (m *OAuthLoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthLoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthLoginRequest proto.InternalMessageInfo

func (m *OAuthLoginRequest) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *OAuthLoginRequest) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

type OAuthLoginResponse struct {
	TokenInfo            *TokenInfo `protobuf:"bytes,1,opt,name=token_info,json=tokenInfo,proto3" json:"token_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *OAuthLoginResponse) Reset()         { *m = OAuthLoginResponse{} }
func (m *OAuthLoginResponse) String() string { return proto.CompactTextString(m) }
func (*OAuthLoginResponse) ProtoMessage()    {}
func (*OAuthLoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e4d00cde15d6b1e, []int{8}
}

func (m *OAuthLoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OAuthLoginResponse.Unmarshal(m, b)
}
func (m *OAuthLoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OAuthLoginResponse.Marshal(b, m, deterministic)
}
func (m *OAuthLoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OAuthLoginResponse.Merge(m, src)
}
func (m *OAuthLoginResponse) XXX_Size() int {
	return xxx_messageInfo_OAuthLoginResponse.Size(m)
}
func (m *OAuthLoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OAuthLoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OAuthLoginResponse proto.InternalMessageInfo

func (m *OAuthLoginResponse) GetTokenInfo() *TokenInfo {
	if m != nil {
		return m.TokenInfo
	}
	return nil
}

type SetPwdRequest struct {
	Passwd               string   `protobuf:"bytes,1,opt,name=passwd,proto3" json:"passwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetPwdRequest) Reset()         { *m = SetPwdRequest{} }
func (m *SetPwdRequest) String() string { return proto.CompactTextString(m) }
func (*SetPwdRequest) ProtoMessage()    {}
func (*SetPwdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e4d00cde15d6b1e, []int{9}
}

func (m *SetPwdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetPwdRequest.Unmarshal(m, b)
}
func (m *SetPwdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetPwdRequest.Marshal(b, m, deterministic)
}
func (m *SetPwdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetPwdRequest.Merge(m, src)
}
func (m *SetPwdRequest) XXX_Size() int {
	return xxx_messageInfo_SetPwdRequest.Size(m)
}
func (m *SetPwdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetPwdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetPwdRequest proto.InternalMessageInfo

func (m *SetPwdRequest) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

type SetPwdResponse struct {
	TokenInfo            *TokenInfo `protobuf:"bytes,1,opt,name=token_info,json=tokenInfo,proto3" json:"token_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SetPwdResponse) Reset()         { *m = SetPwdResponse{} }
func (m *SetPwdResponse) String() string { return proto.CompactTextString(m) }
func (*SetPwdResponse) ProtoMessage()    {}
func (*SetPwdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e4d00cde15d6b1e, []int{10}
}

func (m *SetPwdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetPwdResponse.Unmarshal(m, b)
}
func (m *SetPwdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetPwdResponse.Marshal(b, m, deterministic)
}
func (m *SetPwdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetPwdResponse.Merge(m, src)
}
func (m *SetPwdResponse) XXX_Size() int {
	return xxx_messageInfo_SetPwdResponse.Size(m)
}
func (m *SetPwdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetPwdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetPwdResponse proto.InternalMessageInfo

func (m *SetPwdResponse) GetTokenInfo() *TokenInfo {
	if m != nil {
		return m.TokenInfo
	}
	return nil
}

type AuthTokenRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthTokenRequest) Reset()         { *m = AuthTokenRequest{} }
func (m *AuthTokenRequest) String() string { return proto.CompactTextString(m) }
func (*AuthTokenRequest) ProtoMessage()    {}
func (*AuthTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e4d00cde15d6b1e, []int{11}
}

func (m *AuthTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthTokenRequest.Unmarshal(m, b)
}
func (m *AuthTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthTokenRequest.Marshal(b, m, deterministic)
}
func (m *AuthTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthTokenRequest.Merge(m, src)
}
func (m *AuthTokenRequest) XXX_Size() int {
	return xxx_messageInfo_AuthTokenRequest.Size(m)
}
func (m *AuthTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthTokenRequest proto.InternalMessageInfo

type AuthTokenResponse struct {
	Appid                int32    `protobuf:"varint,1,opt,name=appid,proto3" json:"appid,omitempty"`
	Uid                  int64    `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Plat                 int32    `protobuf:"varint,3,opt,name=plat,proto3" json:"plat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthTokenResponse) Reset()         { *m = AuthTokenResponse{} }
func (m *AuthTokenResponse) String() string { return proto.CompactTextString(m) }
func (*AuthTokenResponse) ProtoMessage()    {}
func (*AuthTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e4d00cde15d6b1e, []int{12}
}

func (m *AuthTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthTokenResponse.Unmarshal(m, b)
}
func (m *AuthTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthTokenResponse.Marshal(b, m, deterministic)
}
func (m *AuthTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthTokenResponse.Merge(m, src)
}
func (m *AuthTokenResponse) XXX_Size() int {
	return xxx_messageInfo_AuthTokenResponse.Size(m)
}
func (m *AuthTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthTokenResponse proto.InternalMessageInfo

func (m *AuthTokenResponse) GetAppid() int32 {
	if m != nil {
		return m.Appid
	}
	return 0
}

func (m *AuthTokenResponse) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *AuthTokenResponse) GetPlat() int32 {
	if m != nil {
		return m.Plat
	}
	return 0
}

type AuthCookieRequest struct {
	Cookie               string   `protobuf:"bytes,1,opt,name=Cookie,proto3" json:"Cookie,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthCookieRequest) Reset()         { *m = AuthCookieRequest{} }
func (m *AuthCookieRequest) String() string { return proto.CompactTextString(m) }
func (*AuthCookieRequest) ProtoMessage()    {}
func (*AuthCookieRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e4d00cde15d6b1e, []int{13}
}

func (m *AuthCookieRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthCookieRequest.Unmarshal(m, b)
}
func (m *AuthCookieRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthCookieRequest.Marshal(b, m, deterministic)
}
func (m *AuthCookieRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthCookieRequest.Merge(m, src)
}
func (m *AuthCookieRequest) XXX_Size() int {
	return xxx_messageInfo_AuthCookieRequest.Size(m)
}
func (m *AuthCookieRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthCookieRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthCookieRequest proto.InternalMessageInfo

func (m *AuthCookieRequest) GetCookie() string {
	if m != nil {
		return m.Cookie
	}
	return ""
}

type AuthCookieResponse struct {
	Appid                int32    `protobuf:"varint,1,opt,name=appid,proto3" json:"appid,omitempty"`
	Uid                  int64    `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Plat                 int32    `protobuf:"varint,3,opt,name=plat,proto3" json:"plat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthCookieResponse) Reset()         { *m = AuthCookieResponse{} }
func (m *AuthCookieResponse) String() string { return proto.CompactTextString(m) }
func (*AuthCookieResponse) ProtoMessage()    {}
func (*AuthCookieResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e4d00cde15d6b1e, []int{14}
}

func (m *AuthCookieResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthCookieResponse.Unmarshal(m, b)
}
func (m *AuthCookieResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthCookieResponse.Marshal(b, m, deterministic)
}
func (m *AuthCookieResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthCookieResponse.Merge(m, src)
}
func (m *AuthCookieResponse) XXX_Size() int {
	return xxx_messageInfo_AuthCookieResponse.Size(m)
}
func (m *AuthCookieResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthCookieResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthCookieResponse proto.InternalMessageInfo

func (m *AuthCookieResponse) GetAppid() int32 {
	if m != nil {
		return m.Appid
	}
	return 0
}

func (m *AuthCookieResponse) GetUid() int64 {
	if m != nil {
		return m.Uid
	}
	return 0
}

func (m *AuthCookieResponse) GetPlat() int32 {
	if m != nil {
		return m.Plat
	}
	return 0
}

func init() {
	proto.RegisterType((*TokenInfo)(nil), "go.micro.srv.passport.TokenInfo")
	proto.RegisterType((*Request)(nil), "go.micro.srv.passport.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.passport.Response")
	proto.RegisterType((*SmsLoginRequest)(nil), "go.micro.srv.passport.SmsLoginRequest")
	proto.RegisterType((*SmsLoginResponse)(nil), "go.micro.srv.passport.SmsLoginResponse")
	proto.RegisterType((*LoginRequest)(nil), "go.micro.srv.passport.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "go.micro.srv.passport.LoginResponse")
	proto.RegisterType((*OAuthLoginRequest)(nil), "go.micro.srv.passport.OAuthLoginRequest")
	proto.RegisterType((*OAuthLoginResponse)(nil), "go.micro.srv.passport.OAuthLoginResponse")
	proto.RegisterType((*SetPwdRequest)(nil), "go.micro.srv.passport.SetPwdRequest")
	proto.RegisterType((*SetPwdResponse)(nil), "go.micro.srv.passport.SetPwdResponse")
	proto.RegisterType((*AuthTokenRequest)(nil), "go.micro.srv.passport.AuthTokenRequest")
	proto.RegisterType((*AuthTokenResponse)(nil), "go.micro.srv.passport.AuthTokenResponse")
	proto.RegisterType((*AuthCookieRequest)(nil), "go.micro.srv.passport.AuthCookieRequest")
	proto.RegisterType((*AuthCookieResponse)(nil), "go.micro.srv.passport.AuthCookieResponse")
}

func init() { proto.RegisterFile("srv/passport/api/passport.proto", fileDescriptor_1e4d00cde15d6b1e) }

var fileDescriptor_1e4d00cde15d6b1e = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x61, 0x6f, 0xd3, 0x30,
	0x10, 0x5d, 0x97, 0xa5, 0x34, 0xc7, 0x06, 0xdd, 0x69, 0x4c, 0x53, 0x24, 0xb6, 0xe2, 0x01, 0x2b,
	0x42, 0xca, 0xa4, 0xf1, 0x03, 0xd0, 0xb4, 0x2f, 0x20, 0x21, 0xad, 0xa4, 0x43, 0x7c, 0x42, 0x25,
	0x6b, 0xdc, 0x36, 0xda, 0x12, 0x87, 0xd8, 0xdd, 0xf8, 0x15, 0xfc, 0x66, 0x14, 0xdb, 0x71, 0x9d,
	0x6d, 0x59, 0x91, 0xe8, 0x37, 0x9f, 0xef, 0xdd, 0xbd, 0x77, 0xcf, 0x17, 0x05, 0x0e, 0x78, 0x71,
	0x73, 0x9c, 0x47, 0x9c, 0xe7, 0xac, 0x10, 0xc7, 0x51, 0x9e, 0x98, 0x20, 0xc8, 0x0b, 0x26, 0x18,
	0xbe, 0x98, 0xb2, 0x20, 0x4d, 0xc6, 0x05, 0x0b, 0x78, 0x71, 0x13, 0x54, 0x49, 0x42, 0xc1, 0xbb,
	0x60, 0x57, 0x34, 0xfb, 0x9c, 0x4d, 0x18, 0xee, 0x80, 0x2b, 0xca, 0x60, 0xaf, 0xd5, 0x6b, 0xf5,
	0xbd, 0x50, 0x05, 0x78, 0x08, 0x5b, 0x05, 0x9d, 0x14, 0x94, 0xcf, 0x46, 0x2a, 0xbb, 0x2e, 0xb3,
	0x9b, 0xfa, 0x52, 0x96, 0xe3, 0x4b, 0x00, 0xfa, 0x3b, 0x4f, 0x0a, 0xca, 0x47, 0x91, 0xd8, 0x73,
	0x7a, 0xad, 0xbe, 0x13, 0x7a, 0xfa, 0xe6, 0x54, 0x90, 0x57, 0xf0, 0x24, 0xa4, 0xbf, 0xe6, 0x94,
	0x0b, 0xdc, 0x85, 0x76, 0xca, 0x2e, 0x93, 0x6b, 0xaa, 0x59, 0x74, 0x44, 0xf6, 0xa1, 0x13, 0x52,
	0x9e, 0xb3, 0x8c, 0x53, 0x44, 0xd8, 0x18, 0xb3, 0xb8, 0x42, 0xc8, 0x33, 0x99, 0xc2, 0xf3, 0x61,
	0xca, 0xbf, 0xb0, 0x69, 0x92, 0x55, 0xad, 0x76, 0xc0, 0x8d, 0xf2, 0x3c, 0x89, 0x25, 0xce, 0x0d,
	0x55, 0x50, 0x16, 0xe7, 0xd7, 0x91, 0x90, 0x32, 0xdd, 0x50, 0x9e, 0x2d, 0x52, 0xc7, 0x26, 0x35,
	0x44, 0x1b, 0x16, 0xd1, 0x10, 0xba, 0x0b, 0x22, 0x2d, 0xe8, 0x23, 0x80, 0x9c, 0x7d, 0x94, 0x64,
	0x13, 0x26, 0xe9, 0x9e, 0x9e, 0xf4, 0x82, 0x07, 0x2d, 0x0d, 0x8c, 0x9f, 0xa1, 0x27, 0xaa, 0x23,
	0x99, 0xc1, 0xe6, 0x8a, 0xa5, 0xef, 0x42, 0xbb, 0xa4, 0xbc, 0x8d, 0xb5, 0x78, 0x1d, 0x91, 0x01,
	0x6c, 0xad, 0x58, 0xfb, 0x19, 0x6c, 0x9f, 0x9f, 0xce, 0xc5, 0xac, 0x36, 0x80, 0x0f, 0x9d, 0x52,
	0xde, 0x84, 0x15, 0xa9, 0x7e, 0x26, 0x13, 0x1b, 0x57, 0xd7, 0x2d, 0x57, 0xbf, 0x01, 0xda, 0x4d,
	0x56, 0xa5, 0xed, 0x08, 0xb6, 0x86, 0x54, 0x0c, 0x6e, 0x63, 0x6b, 0xbd, 0xb4, 0x2d, 0xad, 0x9a,
	0x2d, 0x5f, 0xe1, 0x59, 0x05, 0x5c, 0x15, 0x37, 0x42, 0xb7, 0x9c, 0x48, 0xe6, 0x34, 0x3d, 0x39,
	0x87, 0x6d, 0xeb, 0x4e, 0x33, 0x3d, 0xfc, 0xd8, 0x5d, 0x70, 0xe6, 0x49, 0x2c, 0x4d, 0x72, 0xc2,
	0xf2, 0x68, 0x9e, 0xdf, 0x59, 0x3c, 0x3f, 0x79, 0xaf, 0x1a, 0x9e, 0x31, 0x76, 0x95, 0x50, 0x6b,
	0x48, 0x75, 0x51, 0x0d, 0xa9, 0x22, 0x32, 0x00, 0xb4, 0xc1, 0xff, 0x4f, 0x7f, 0xf2, 0xc7, 0x85,
	0xce, 0x40, 0xbb, 0x80, 0x9f, 0xc0, 0x19, 0xa6, 0x1c, 0xf7, 0x1b, 0x4c, 0xd2, 0xea, 0xfc, 0x83,
	0xc6, 0xbc, 0x12, 0x44, 0xd6, 0xf0, 0x07, 0x74, 0xaa, 0x6f, 0x0c, 0xdf, 0x36, 0xc0, 0xef, 0x7c,
	0xed, 0xfe, 0xd1, 0x52, 0x9c, 0x69, 0x7f, 0x01, 0xae, 0xea, 0x7d, 0xd8, 0x50, 0x53, 0x6b, 0xfc,
	0xfa, 0x71, 0x90, 0xe9, 0x3a, 0x06, 0x58, 0xac, 0x30, 0xf6, 0x1b, 0xaa, 0xee, 0x7d, 0x2a, 0xfe,
	0xbb, 0x7f, 0x40, 0x1a, 0x92, 0xef, 0xd0, 0x56, 0x7b, 0x8a, 0x4d, 0xb2, 0x6a, 0xfb, 0xee, 0xbf,
	0x59, 0x82, 0x32, 0x8d, 0x7f, 0x82, 0x67, 0x36, 0x13, 0x9b, 0xbc, 0xbc, 0xbb, 0xcf, 0x7e, 0x7f,
	0x39, 0xd0, 0xf6, 0x67, 0xb1, 0x7d, 0xf8, 0x58, 0x65, 0x6d, 0x9b, 0x1b, 0xfd, 0xb9, 0xbf, 0xca,
	0x64, 0xed, 0xb2, 0x2d, 0x7f, 0x67, 0x1f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x70, 0x24, 0x09,
	0xba, 0xf1, 0x06, 0x00, 0x00,
}
