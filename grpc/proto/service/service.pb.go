// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/service/service.proto

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	proto/service/service.proto

It has these top-level messages:
	CreateUserRequest
	GetUserRequest
	GreetUserRequest
	GreetUserResponse
	User
*/
package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type CreateUserRequest struct {
	User *User `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
}

func (m *CreateUserRequest) Reset()                    { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()               {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type GetUserRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
}

func (m *GetUserRequest) Reset()                    { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()               {}
func (*GetUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

type GreetUserRequest struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Greeting string `protobuf:"bytes,2,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *GreetUserRequest) Reset()                    { *m = GreetUserRequest{} }
func (m *GreetUserRequest) String() string            { return proto.CompactTextString(m) }
func (*GreetUserRequest) ProtoMessage()               {}
func (*GreetUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GreetUserRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *GreetUserRequest) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

type GreetUserResponse struct {
	Greeting string `protobuf:"bytes,1,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *GreetUserResponse) Reset()                    { *m = GreetUserResponse{} }
func (m *GreetUserResponse) String() string            { return proto.CompactTextString(m) }
func (*GreetUserResponse) ProtoMessage()               {}
func (*GreetUserResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *GreetUserResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

type User struct {
	Username string `protobuf:"bytes,1,opt,name=username" json:"username,omitempty"`
	Role     string `protobuf:"bytes,2,opt,name=role" json:"role,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateUserRequest)(nil), "myself659.golab.grpc.service.CreateUserRequest")
	proto.RegisterType((*GetUserRequest)(nil), "myself659.golab.grpc.service.GetUserRequest")
	proto.RegisterType((*GreetUserRequest)(nil), "myself659.golab.grpc.service.GreetUserRequest")
	proto.RegisterType((*GreetUserResponse)(nil), "myself659.golab.grpc.service.GreetUserResponse")
	proto.RegisterType((*User)(nil), "myself659.golab.grpc.service.User")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SimpleServer service

type SimpleServerClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	GreetUser(ctx context.Context, in *GreetUserRequest, opts ...grpc.CallOption) (*GreetUserResponse, error)
}

type simpleServerClient struct {
	cc *grpc.ClientConn
}

func NewSimpleServerClient(cc *grpc.ClientConn) SimpleServerClient {
	return &simpleServerClient{cc}
}

func (c *simpleServerClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/myself659.golab.grpc.service.SimpleServer/CreateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleServerClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := grpc.Invoke(ctx, "/myself659.golab.grpc.service.SimpleServer/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *simpleServerClient) GreetUser(ctx context.Context, in *GreetUserRequest, opts ...grpc.CallOption) (*GreetUserResponse, error) {
	out := new(GreetUserResponse)
	err := grpc.Invoke(ctx, "/myself659.golab.grpc.service.SimpleServer/GreetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SimpleServer service

type SimpleServerServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*google_protobuf.Empty, error)
	GetUser(context.Context, *GetUserRequest) (*User, error)
	GreetUser(context.Context, *GreetUserRequest) (*GreetUserResponse, error)
}

func RegisterSimpleServerServer(s *grpc.Server, srv SimpleServerServer) {
	s.RegisterService(&_SimpleServer_serviceDesc, srv)
}

func _SimpleServer_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleServerServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myself659.golab.grpc.service.SimpleServer/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleServerServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleServer_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleServerServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myself659.golab.grpc.service.SimpleServer/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleServerServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SimpleServer_GreetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SimpleServerServer).GreetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/myself659.golab.grpc.service.SimpleServer/GreetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SimpleServerServer).GreetUser(ctx, req.(*GreetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SimpleServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "myself659.golab.grpc.service.SimpleServer",
	HandlerType: (*SimpleServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _SimpleServer_CreateUser_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _SimpleServer_GetUser_Handler,
		},
		{
			MethodName: "GreetUser",
			Handler:    _SimpleServer_GreetUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service/service.proto",
}

func init() { proto.RegisterFile("proto/service/service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xdf, 0x4a, 0xe3, 0x40,
	0x14, 0xc6, 0x49, 0xb7, 0x74, 0xb7, 0xd3, 0x65, 0xd9, 0xcc, 0xc2, 0x6e, 0x37, 0xed, 0x85, 0xcc,
	0x95, 0x94, 0x3a, 0x07, 0x2a, 0x16, 0xec, 0xa5, 0x22, 0x05, 0xbd, 0x6b, 0xf1, 0xc6, 0xbb, 0xa4,
	0x9c, 0xa6, 0x81, 0x24, 0x13, 0x67, 0x26, 0x85, 0x22, 0xde, 0xe8, 0x0b, 0x08, 0x3e, 0x9a, 0xaf,
	0xe0, 0x83, 0x48, 0x26, 0x49, 0xed, 0x1f, 0x88, 0xbd, 0x4a, 0x4e, 0xce, 0x77, 0xce, 0xef, 0xf0,
	0x7d, 0x84, 0x74, 0x12, 0x29, 0xb4, 0x00, 0x85, 0x72, 0x19, 0xcc, 0xb0, 0x7c, 0x72, 0xf3, 0x95,
	0x76, 0xa3, 0x95, 0xc2, 0x70, 0x3e, 0x3c, 0x3b, 0xe7, 0xbe, 0x08, 0x5d, 0x8f, 0xfb, 0x32, 0x99,
	0xf1, 0x42, 0xe3, 0x74, 0x7c, 0x21, 0xfc, 0x10, 0xc1, 0x68, 0xbd, 0x74, 0x0e, 0x18, 0x25, 0x7a,
	0x95, 0x8f, 0x3a, 0xdd, 0xa2, 0xe9, 0x26, 0x01, 0xb8, 0x71, 0x2c, 0xb4, 0xab, 0x03, 0x11, 0xab,
	0xbc, 0xcb, 0x6e, 0x88, 0x7d, 0x29, 0xd1, 0xd5, 0x78, 0xab, 0x50, 0x4e, 0xf0, 0x3e, 0x45, 0xa5,
	0xe9, 0x90, 0xd4, 0x53, 0x85, 0xb2, 0x6d, 0x1d, 0x59, 0xc7, 0xad, 0x01, 0xe3, 0x55, 0x70, 0x6e,
	0x06, 0x8d, 0x9e, 0xf5, 0xc9, 0xaf, 0x31, 0xea, 0xcd, 0x4d, 0x0e, 0xf9, 0x91, 0x75, 0x62, 0x37,
	0x42, 0xb3, 0xad, 0x39, 0x59, 0xd7, 0xec, 0x9a, 0xfc, 0x1e, 0x4b, 0x3c, 0x58, 0x9f, 0xf5, 0xfc,
	0x4c, 0x1f, 0xc4, 0x7e, 0xbb, 0x96, 0xf7, 0xca, 0x9a, 0x01, 0xb1, 0x37, 0x76, 0xa9, 0x44, 0xc4,
	0x6a, 0x7b, 0xc0, 0xda, 0x19, 0x18, 0x92, 0x7a, 0xa6, 0xad, 0x04, 0x52, 0x52, 0x97, 0x22, 0xc4,
	0x02, 0x66, 0xde, 0x07, 0xcf, 0xdf, 0xc8, 0xcf, 0x69, 0x10, 0x25, 0x21, 0x4e, 0x51, 0x2e, 0x51,
	0xd2, 0x05, 0x21, 0x9f, 0x06, 0x52, 0xa8, 0xf6, 0x6a, 0xcf, 0x6a, 0xe7, 0x2f, 0xcf, 0xe3, 0xe1,
	0x65, 0x76, 0xfc, 0x2a, 0xcb, 0x8e, 0xd9, 0x4f, 0x6f, 0xef, 0xaf, 0xb5, 0x16, 0x6b, 0x40, 0x76,
	0x8e, 0x1a, 0x59, 0x3d, 0xba, 0x24, 0xdf, 0x0b, 0x77, 0x69, 0xbf, 0x1a, 0xb3, 0x1d, 0x82, 0x73,
	0x40, 0x80, 0xec, 0xbf, 0xe1, 0xfd, 0xa1, 0x76, 0xce, 0x83, 0x87, 0xd2, 0x85, 0x47, 0xfa, 0x62,
	0x91, 0xe6, 0xda, 0x5c, 0xca, 0xbf, 0x40, 0xef, 0x24, 0xea, 0xc0, 0xc1, 0xfa, 0x3c, 0x35, 0xc6,
	0xcc, 0x25, 0x5d, 0xf6, 0x6f, 0xef, 0x12, 0x30, 0xe9, 0x8d, 0xac, 0xde, 0x05, 0xdc, 0x9d, 0xf8,
	0x81, 0x5e, 0xa4, 0x1e, 0x9f, 0x89, 0x08, 0xd6, 0x00, 0x30, 0x00, 0xc8, 0x00, 0xb0, 0xf5, 0x37,
	0x79, 0x0d, 0x53, 0x9e, 0x7e, 0x04, 0x00, 0x00, 0xff, 0xff, 0x62, 0x27, 0x04, 0x42, 0x65, 0x03,
	0x00, 0x00,
}
