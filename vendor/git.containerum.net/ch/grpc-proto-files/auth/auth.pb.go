// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth/auth.proto

/*
Package auth is a generated protocol buffer package.

It is generated from these files:
	auth/auth.proto
	auth/auth_types.proto

It has these top-level messages:
	CreateTokenRequest
	CreateTokenResponse
	CheckTokenRequest
	CheckTokenResponse
	ExtendTokenRequest
	ExtendTokenResponse
	UpdateAccessRequestElement
	UpdateAccessRequest
	GetUserTokensRequest
	GetUserTokensResponse
	DeleteTokenRequest
	DeleteUserTokensRequest
	StoredToken
	AccessObject
	ResourcesAccess
	StoredTokenForUser
*/
package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import uuid "git.containerum.net/ch/grpc-proto-files/common"
import google_protobuf1 "github.com/golang/protobuf/ptypes/empty"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateTokenRequest struct {
	UserAgent   string           `protobuf:"bytes,1,opt,name=user_agent,json=userAgent" json:"user_agent,omitempty"`
	Fingerprint string           `protobuf:"bytes,2,opt,name=fingerprint" json:"fingerprint,omitempty"`
	UserId      *uuid.UUID       `protobuf:"bytes,3,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserIp      string           `protobuf:"bytes,4,opt,name=user_ip,json=userIp" json:"user_ip,omitempty"`
	UserRole    string           `protobuf:"bytes,5,opt,name=user_role,json=userRole" json:"user_role,omitempty"`
	RwAccess    bool             `protobuf:"varint,6,opt,name=rw_access,json=rwAccess" json:"rw_access,omitempty"`
	Access      *ResourcesAccess `protobuf:"bytes,7,opt,name=access" json:"access,omitempty"`
	PartTokenId *uuid.UUID       `protobuf:"bytes,8,opt,name=part_token_id,json=partTokenId" json:"part_token_id,omitempty"`
}

func (m *CreateTokenRequest) Reset()                    { *m = CreateTokenRequest{} }
func (m *CreateTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateTokenRequest) ProtoMessage()               {}
func (*CreateTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateTokenRequest) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *CreateTokenRequest) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

func (m *CreateTokenRequest) GetUserId() *uuid.UUID {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *CreateTokenRequest) GetUserIp() string {
	if m != nil {
		return m.UserIp
	}
	return ""
}

func (m *CreateTokenRequest) GetUserRole() string {
	if m != nil {
		return m.UserRole
	}
	return ""
}

func (m *CreateTokenRequest) GetRwAccess() bool {
	if m != nil {
		return m.RwAccess
	}
	return false
}

func (m *CreateTokenRequest) GetAccess() *ResourcesAccess {
	if m != nil {
		return m.Access
	}
	return nil
}

func (m *CreateTokenRequest) GetPartTokenId() *uuid.UUID {
	if m != nil {
		return m.PartTokenId
	}
	return nil
}

type CreateTokenResponse struct {
	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
}

func (m *CreateTokenResponse) Reset()                    { *m = CreateTokenResponse{} }
func (m *CreateTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateTokenResponse) ProtoMessage()               {}
func (*CreateTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CreateTokenResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *CreateTokenResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type CheckTokenRequest struct {
	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	UserAgent   string `protobuf:"bytes,2,opt,name=user_agent,json=userAgent" json:"user_agent,omitempty"`
	FingerPrint string `protobuf:"bytes,3,opt,name=finger_print,json=fingerPrint" json:"finger_print,omitempty"`
	UserIp      string `protobuf:"bytes,4,opt,name=user_ip,json=userIp" json:"user_ip,omitempty"`
}

func (m *CheckTokenRequest) Reset()                    { *m = CheckTokenRequest{} }
func (m *CheckTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckTokenRequest) ProtoMessage()               {}
func (*CheckTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CheckTokenRequest) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *CheckTokenRequest) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

func (m *CheckTokenRequest) GetFingerPrint() string {
	if m != nil {
		return m.FingerPrint
	}
	return ""
}

func (m *CheckTokenRequest) GetUserIp() string {
	if m != nil {
		return m.UserIp
	}
	return ""
}

type CheckTokenResponse struct {
	Access      *ResourcesAccess `protobuf:"bytes,1,opt,name=access" json:"access,omitempty"`
	UserId      *uuid.UUID       `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	UserRole    string           `protobuf:"bytes,3,opt,name=user_role,json=userRole" json:"user_role,omitempty"`
	TokenId     *uuid.UUID       `protobuf:"bytes,4,opt,name=token_id,json=tokenId" json:"token_id,omitempty"`
	PartTokenId *uuid.UUID       `protobuf:"bytes,5,opt,name=part_token_id,json=partTokenId" json:"part_token_id,omitempty"`
}

func (m *CheckTokenResponse) Reset()                    { *m = CheckTokenResponse{} }
func (m *CheckTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*CheckTokenResponse) ProtoMessage()               {}
func (*CheckTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CheckTokenResponse) GetAccess() *ResourcesAccess {
	if m != nil {
		return m.Access
	}
	return nil
}

func (m *CheckTokenResponse) GetUserId() *uuid.UUID {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *CheckTokenResponse) GetUserRole() string {
	if m != nil {
		return m.UserRole
	}
	return ""
}

func (m *CheckTokenResponse) GetTokenId() *uuid.UUID {
	if m != nil {
		return m.TokenId
	}
	return nil
}

func (m *CheckTokenResponse) GetPartTokenId() *uuid.UUID {
	if m != nil {
		return m.PartTokenId
	}
	return nil
}

type ExtendTokenRequest struct {
	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
	Fingerprint  string `protobuf:"bytes,2,opt,name=fingerprint" json:"fingerprint,omitempty"`
}

func (m *ExtendTokenRequest) Reset()                    { *m = ExtendTokenRequest{} }
func (m *ExtendTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*ExtendTokenRequest) ProtoMessage()               {}
func (*ExtendTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ExtendTokenRequest) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *ExtendTokenRequest) GetFingerprint() string {
	if m != nil {
		return m.Fingerprint
	}
	return ""
}

type ExtendTokenResponse struct {
	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken" json:"refresh_token,omitempty"`
}

func (m *ExtendTokenResponse) Reset()                    { *m = ExtendTokenResponse{} }
func (m *ExtendTokenResponse) String() string            { return proto.CompactTextString(m) }
func (*ExtendTokenResponse) ProtoMessage()               {}
func (*ExtendTokenResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ExtendTokenResponse) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *ExtendTokenResponse) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

type UpdateAccessRequestElement struct {
	UserId *uuid.UUID       `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Access *ResourcesAccess `protobuf:"bytes,2,opt,name=access" json:"access,omitempty"`
}

func (m *UpdateAccessRequestElement) Reset()                    { *m = UpdateAccessRequestElement{} }
func (m *UpdateAccessRequestElement) String() string            { return proto.CompactTextString(m) }
func (*UpdateAccessRequestElement) ProtoMessage()               {}
func (*UpdateAccessRequestElement) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *UpdateAccessRequestElement) GetUserId() *uuid.UUID {
	if m != nil {
		return m.UserId
	}
	return nil
}

func (m *UpdateAccessRequestElement) GetAccess() *ResourcesAccess {
	if m != nil {
		return m.Access
	}
	return nil
}

type UpdateAccessRequest struct {
	Users []*UpdateAccessRequestElement `protobuf:"bytes,1,rep,name=users" json:"users,omitempty"`
}

func (m *UpdateAccessRequest) Reset()                    { *m = UpdateAccessRequest{} }
func (m *UpdateAccessRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateAccessRequest) ProtoMessage()               {}
func (*UpdateAccessRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UpdateAccessRequest) GetUsers() []*UpdateAccessRequestElement {
	if m != nil {
		return m.Users
	}
	return nil
}

type GetUserTokensRequest struct {
	UserId *uuid.UUID `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *GetUserTokensRequest) Reset()                    { *m = GetUserTokensRequest{} }
func (m *GetUserTokensRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserTokensRequest) ProtoMessage()               {}
func (*GetUserTokensRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *GetUserTokensRequest) GetUserId() *uuid.UUID {
	if m != nil {
		return m.UserId
	}
	return nil
}

type GetUserTokensResponse struct {
	Tokens []*StoredTokenForUser `protobuf:"bytes,1,rep,name=tokens" json:"tokens,omitempty"`
}

func (m *GetUserTokensResponse) Reset()                    { *m = GetUserTokensResponse{} }
func (m *GetUserTokensResponse) String() string            { return proto.CompactTextString(m) }
func (*GetUserTokensResponse) ProtoMessage()               {}
func (*GetUserTokensResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GetUserTokensResponse) GetTokens() []*StoredTokenForUser {
	if m != nil {
		return m.Tokens
	}
	return nil
}

type DeleteTokenRequest struct {
	TokenId *uuid.UUID `protobuf:"bytes,1,opt,name=token_id,json=tokenId" json:"token_id,omitempty"`
	UserId  *uuid.UUID `protobuf:"bytes,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *DeleteTokenRequest) Reset()                    { *m = DeleteTokenRequest{} }
func (m *DeleteTokenRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteTokenRequest) ProtoMessage()               {}
func (*DeleteTokenRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DeleteTokenRequest) GetTokenId() *uuid.UUID {
	if m != nil {
		return m.TokenId
	}
	return nil
}

func (m *DeleteTokenRequest) GetUserId() *uuid.UUID {
	if m != nil {
		return m.UserId
	}
	return nil
}

type DeleteUserTokensRequest struct {
	UserId *uuid.UUID `protobuf:"bytes,1,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *DeleteUserTokensRequest) Reset()                    { *m = DeleteUserTokensRequest{} }
func (m *DeleteUserTokensRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteUserTokensRequest) ProtoMessage()               {}
func (*DeleteUserTokensRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DeleteUserTokensRequest) GetUserId() *uuid.UUID {
	if m != nil {
		return m.UserId
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateTokenRequest)(nil), "CreateTokenRequest")
	proto.RegisterType((*CreateTokenResponse)(nil), "CreateTokenResponse")
	proto.RegisterType((*CheckTokenRequest)(nil), "CheckTokenRequest")
	proto.RegisterType((*CheckTokenResponse)(nil), "CheckTokenResponse")
	proto.RegisterType((*ExtendTokenRequest)(nil), "ExtendTokenRequest")
	proto.RegisterType((*ExtendTokenResponse)(nil), "ExtendTokenResponse")
	proto.RegisterType((*UpdateAccessRequestElement)(nil), "UpdateAccessRequestElement")
	proto.RegisterType((*UpdateAccessRequest)(nil), "UpdateAccessRequest")
	proto.RegisterType((*GetUserTokensRequest)(nil), "GetUserTokensRequest")
	proto.RegisterType((*GetUserTokensResponse)(nil), "GetUserTokensResponse")
	proto.RegisterType((*DeleteTokenRequest)(nil), "DeleteTokenRequest")
	proto.RegisterType((*DeleteUserTokensRequest)(nil), "DeleteUserTokensRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Auth service

type AuthClient interface {
	CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error)
	CheckToken(ctx context.Context, in *CheckTokenRequest, opts ...grpc.CallOption) (*CheckTokenResponse, error)
	ExtendToken(ctx context.Context, in *ExtendTokenRequest, opts ...grpc.CallOption) (*ExtendTokenResponse, error)
	UpdateAccess(ctx context.Context, in *UpdateAccessRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	GetUserTokens(ctx context.Context, in *GetUserTokensRequest, opts ...grpc.CallOption) (*GetUserTokensResponse, error)
	DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
	DeleteUserTokens(ctx context.Context, in *DeleteUserTokensRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) CreateToken(ctx context.Context, in *CreateTokenRequest, opts ...grpc.CallOption) (*CreateTokenResponse, error) {
	out := new(CreateTokenResponse)
	err := grpc.Invoke(ctx, "/Auth/CreateToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) CheckToken(ctx context.Context, in *CheckTokenRequest, opts ...grpc.CallOption) (*CheckTokenResponse, error) {
	out := new(CheckTokenResponse)
	err := grpc.Invoke(ctx, "/Auth/CheckToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ExtendToken(ctx context.Context, in *ExtendTokenRequest, opts ...grpc.CallOption) (*ExtendTokenResponse, error) {
	out := new(ExtendTokenResponse)
	err := grpc.Invoke(ctx, "/Auth/ExtendToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) UpdateAccess(ctx context.Context, in *UpdateAccessRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/Auth/UpdateAccess", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) GetUserTokens(ctx context.Context, in *GetUserTokensRequest, opts ...grpc.CallOption) (*GetUserTokensResponse, error) {
	out := new(GetUserTokensResponse)
	err := grpc.Invoke(ctx, "/Auth/GetUserTokens", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteToken(ctx context.Context, in *DeleteTokenRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/Auth/DeleteToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) DeleteUserTokens(ctx context.Context, in *DeleteUserTokensRequest, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/Auth/DeleteUserTokens", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	CreateToken(context.Context, *CreateTokenRequest) (*CreateTokenResponse, error)
	CheckToken(context.Context, *CheckTokenRequest) (*CheckTokenResponse, error)
	ExtendToken(context.Context, *ExtendTokenRequest) (*ExtendTokenResponse, error)
	UpdateAccess(context.Context, *UpdateAccessRequest) (*google_protobuf1.Empty, error)
	GetUserTokens(context.Context, *GetUserTokensRequest) (*GetUserTokensResponse, error)
	DeleteToken(context.Context, *DeleteTokenRequest) (*google_protobuf1.Empty, error)
	DeleteUserTokens(context.Context, *DeleteUserTokensRequest) (*google_protobuf1.Empty, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_CreateToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CreateToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/CreateToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CreateToken(ctx, req.(*CreateTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_CheckToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).CheckToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/CheckToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).CheckToken(ctx, req.(*CheckTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_ExtendToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExtendTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).ExtendToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/ExtendToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).ExtendToken(ctx, req.(*ExtendTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_UpdateAccess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateAccessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).UpdateAccess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/UpdateAccess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).UpdateAccess(ctx, req.(*UpdateAccessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_GetUserTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).GetUserTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/GetUserTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).GetUserTokens(ctx, req.(*GetUserTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DeleteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DeleteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/DeleteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DeleteToken(ctx, req.(*DeleteTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Auth_DeleteUserTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).DeleteUserTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Auth/DeleteUserTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).DeleteUserTokens(ctx, req.(*DeleteUserTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateToken",
			Handler:    _Auth_CreateToken_Handler,
		},
		{
			MethodName: "CheckToken",
			Handler:    _Auth_CheckToken_Handler,
		},
		{
			MethodName: "ExtendToken",
			Handler:    _Auth_ExtendToken_Handler,
		},
		{
			MethodName: "UpdateAccess",
			Handler:    _Auth_UpdateAccess_Handler,
		},
		{
			MethodName: "GetUserTokens",
			Handler:    _Auth_GetUserTokens_Handler,
		},
		{
			MethodName: "DeleteToken",
			Handler:    _Auth_DeleteToken_Handler,
		},
		{
			MethodName: "DeleteUserTokens",
			Handler:    _Auth_DeleteUserTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth/auth.proto",
}

func init() { proto.RegisterFile("auth/auth.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 713 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x95, 0xdd, 0x6e, 0xd3, 0x4a,
	0x10, 0xc7, 0xe3, 0xa4, 0x49, 0xd3, 0x49, 0xaa, 0xd3, 0x6e, 0xd2, 0x36, 0x4a, 0x75, 0x8e, 0xd2,
	0x3d, 0x37, 0x41, 0xd0, 0x8d, 0x28, 0x12, 0x08, 0x84, 0x10, 0xa5, 0x1f, 0xd0, 0x3b, 0x64, 0x08,
	0x17, 0x20, 0x14, 0xb9, 0xce, 0x24, 0xb1, 0x9a, 0x78, 0xcd, 0x7a, 0xad, 0xd2, 0xc7, 0xe0, 0x01,
	0x78, 0x16, 0x6e, 0x79, 0x2c, 0xb4, 0xbb, 0x6e, 0x6b, 0xc7, 0x4e, 0x03, 0x12, 0x37, 0x96, 0x3c,
	0x33, 0x3b, 0x3b, 0xfb, 0xff, 0xcd, 0xce, 0xc2, 0x3f, 0x4e, 0x24, 0x27, 0x3d, 0xf5, 0x61, 0x81,
	0xe0, 0x92, 0xb7, 0x37, 0x5d, 0x3e, 0x9b, 0x71, 0xbf, 0x17, 0x45, 0xde, 0x30, 0x36, 0x6d, 0xdd,
	0xc4, 0x0c, 0xe4, 0x55, 0x80, 0x61, 0x6c, 0xde, 0x1d, 0x73, 0x3e, 0x9e, 0x62, 0x4f, 0xff, 0x9d,
	0x47, 0xa3, 0x1e, 0xce, 0x02, 0x79, 0x65, 0x9c, 0xf4, 0x7b, 0x11, 0xc8, 0x91, 0x40, 0x47, 0xe2,
	0x7b, 0x7e, 0x81, 0xbe, 0x8d, 0x5f, 0x22, 0x0c, 0x25, 0xf9, 0x17, 0x20, 0x0a, 0x51, 0x0c, 0x9c,
	0x31, 0xfa, 0xb2, 0x65, 0x75, 0xac, 0xee, 0x9a, 0xbd, 0xa6, 0x2c, 0x87, 0xca, 0x40, 0x3a, 0x50,
	0x1b, 0x79, 0xfe, 0x18, 0x45, 0x20, 0x3c, 0x5f, 0xb6, 0x8a, 0xda, 0x9f, 0x34, 0x91, 0xff, 0x60,
	0x55, 0x27, 0xf0, 0x86, 0xad, 0x52, 0xc7, 0xea, 0xd6, 0x0e, 0xca, 0xac, 0xdf, 0x3f, 0x3b, 0xb6,
	0x2b, 0xca, 0x7a, 0x36, 0x24, 0x3b, 0xd7, 0xfe, 0xa0, 0xb5, 0xa2, 0x57, 0x1b, 0x47, 0x40, 0x76,
	0x41, 0xef, 0x33, 0x10, 0x7c, 0x8a, 0xad, 0xb2, 0x76, 0x55, 0x95, 0xc1, 0xe6, 0x53, 0x54, 0x4e,
	0x71, 0x39, 0x70, 0x5c, 0x17, 0xc3, 0xb0, 0x55, 0xe9, 0x58, 0xdd, 0xaa, 0x5d, 0x15, 0x97, 0x87,
	0xfa, 0x9f, 0x74, 0xa1, 0x12, 0x7b, 0x56, 0xf5, 0x8e, 0x1b, 0xcc, 0xc6, 0x90, 0x47, 0xc2, 0xc5,
	0xd0, 0x44, 0xd8, 0xb1, 0x9f, 0xdc, 0x83, 0xf5, 0xc0, 0x11, 0x72, 0x20, 0xd5, 0x91, 0x55, 0x89,
	0xd5, 0x64, 0x89, 0x35, 0xe5, 0xd3, 0x6a, 0x9c, 0x0d, 0xe9, 0x67, 0x68, 0xa4, 0xe4, 0x09, 0x03,
	0xee, 0x87, 0x48, 0xf6, 0xa0, 0x6e, 0x72, 0x99, 0x1c, 0xb1, 0x42, 0x35, 0x63, 0xd3, 0xa1, 0xe4,
	0x7f, 0x58, 0x17, 0x38, 0x12, 0x18, 0x4e, 0xe2, 0x18, 0xa3, 0x52, 0x3d, 0x36, 0xea, 0x20, 0xfa,
	0xcd, 0x82, 0xcd, 0xa3, 0x09, 0xba, 0x17, 0x29, 0xf5, 0x7f, 0x23, 0x7b, 0x1a, 0x50, 0x71, 0x1e,
	0xd0, 0x1e, 0xd4, 0x0d, 0x8d, 0x81, 0x21, 0x54, 0x4a, 0x12, 0x7a, 0xab, 0x09, 0x2d, 0x22, 0x40,
	0x7f, 0x5a, 0x40, 0x92, 0x35, 0xc5, 0x47, 0xbe, 0x95, 0xd7, 0x5a, 0x22, 0x6f, 0x82, 0x7d, 0x31,
	0x8f, 0x7d, 0x0a, 0x71, 0x69, 0x0e, 0x71, 0x07, 0xaa, 0x37, 0x58, 0x56, 0x92, 0xab, 0x57, 0xa5,
	0x41, 0x92, 0xa5, 0x57, 0x5e, 0x48, 0xef, 0x13, 0x90, 0x93, 0xaf, 0x12, 0xfd, 0x61, 0x4a, 0xde,
	0x0c, 0x19, 0x2b, 0x4b, 0x66, 0x79, 0x8b, 0xab, 0xd6, 0x48, 0x25, 0xff, 0xcb, 0xad, 0x31, 0x82,
	0x76, 0x3f, 0x18, 0x3a, 0x12, 0x63, 0x75, 0x4d, 0xf1, 0x27, 0x53, 0x9c, 0x61, 0xfa, 0x7e, 0x59,
	0x79, 0x1a, 0xdf, 0xd2, 0x2a, 0xde, 0x4d, 0x8b, 0xbe, 0x81, 0x46, 0xce, 0x3e, 0xe4, 0x21, 0x94,
	0x55, 0x2a, 0x45, 0xbb, 0xd4, 0xad, 0x1d, 0xec, 0xb2, 0xc5, 0xc5, 0xd8, 0x26, 0x92, 0x3e, 0x86,
	0xe6, 0x6b, 0x94, 0xfd, 0x10, 0x85, 0x3e, 0xc1, 0x4d, 0xaa, 0x25, 0xb5, 0xd2, 0x63, 0xd8, 0x9a,
	0x5b, 0x17, 0x4b, 0x79, 0x1f, 0x2a, 0x5a, 0x9f, 0xeb, 0x22, 0x1a, 0xec, 0x9d, 0xe4, 0x02, 0x8d,
	0xe0, 0xa7, 0x5c, 0xa8, 0x25, 0x76, 0x1c, 0x42, 0x3f, 0x00, 0x39, 0xc6, 0x29, 0xce, 0x0d, 0xb2,
	0x64, 0x3b, 0x59, 0xb9, 0xed, 0xb4, 0xa4, 0x5b, 0xe9, 0x53, 0xd8, 0x31, 0x79, 0xff, 0xf8, 0x60,
	0x07, 0x3f, 0x4a, 0xb0, 0x72, 0x18, 0xc9, 0x09, 0x79, 0x06, 0xb5, 0xc4, 0x14, 0x21, 0x0d, 0x96,
	0x1d, 0xb9, 0xed, 0x26, 0xcb, 0x19, 0x34, 0xb4, 0x40, 0x9e, 0x00, 0xdc, 0xde, 0x46, 0x42, 0x58,
	0x66, 0x5c, 0xb4, 0x1b, 0x2c, 0x7b, 0x5d, 0x69, 0x41, 0x6d, 0x9a, 0xe8, 0x4f, 0xd2, 0x60, 0xd9,
	0xab, 0xd0, 0x6e, 0xb2, 0x9c, 0x16, 0xa6, 0x05, 0xf2, 0x02, 0xea, 0x49, 0xde, 0xa4, 0x99, 0x87,
	0xbf, 0xbd, 0xcd, 0xcc, 0xd3, 0xc2, 0xae, 0x9f, 0x16, 0x76, 0xa2, 0x9e, 0x16, 0x5a, 0x20, 0x2f,
	0x61, 0x3d, 0x85, 0x94, 0x6c, 0xb1, 0xbc, 0xd6, 0x68, 0x6f, 0xb3, 0x5c, 0xf2, 0xb4, 0x40, 0x9e,
	0x43, 0x2d, 0x81, 0x93, 0x34, 0x58, 0x16, 0xee, 0x1d, 0xfb, 0x9f, 0xc2, 0xc6, 0x3c, 0x34, 0xd2,
	0x62, 0x0b, 0x38, 0x2e, 0xce, 0xf3, 0x8a, 0x7d, 0x7c, 0x30, 0xf6, 0x24, 0x73, 0xb9, 0x2f, 0x1d,
	0xcf, 0x47, 0x11, 0xcd, 0x98, 0x8f, 0xb2, 0xe7, 0x4e, 0x7a, 0x63, 0x11, 0xb8, 0xfb, 0x3a, 0x7e,
	0x7f, 0xe4, 0x4d, 0x31, 0xd4, 0xef, 0xee, 0x79, 0x45, 0x5b, 0x1e, 0xfd, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0x6b, 0xb6, 0x19, 0xe3, 0xaf, 0x07, 0x00, 0x00,
}
