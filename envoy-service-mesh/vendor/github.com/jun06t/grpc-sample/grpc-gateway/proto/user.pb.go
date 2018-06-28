// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package gateway

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Group int32

const (
	Group_USER      Group = 0
	Group_ADMIN     Group = 1
	Group_DEVELOPER Group = 2
)

var Group_name = map[int32]string{
	0: "USER",
	1: "ADMIN",
	2: "DEVELOPER",
}
var Group_value = map[string]int32{
	"USER":      0,
	"ADMIN":     1,
	"DEVELOPER": 2,
}

func (x Group) String() string {
	return proto.EnumName(Group_name, int32(x))
}
func (Group) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type GetUserRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetUserRequest) Reset()                    { *m = GetUserRequest{} }
func (m *GetUserRequest) String() string            { return proto.CompactTextString(m) }
func (*GetUserRequest) ProtoMessage()               {}
func (*GetUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *GetUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type UpdateUserRequest struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,3,opt,name=age" json:"age,omitempty"`
}

func (m *UpdateUserRequest) Reset()                    { *m = UpdateUserRequest{} }
func (m *UpdateUserRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateUserRequest) ProtoMessage()               {}
func (*UpdateUserRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *UpdateUserRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UpdateUserRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateUserRequest) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type UserResponse struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,3,opt,name=age" json:"age,omitempty"`
}

func (m *UserResponse) Reset()                    { *m = UserResponse{} }
func (m *UserResponse) String() string            { return proto.CompactTextString(m) }
func (*UserResponse) ProtoMessage()               {}
func (*UserResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *UserResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserResponse) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

type UserGroupRequest struct {
	Group Group `protobuf:"varint,1,opt,name=group,enum=gateway.Group" json:"group,omitempty"`
}

func (m *UserGroupRequest) Reset()                    { *m = UserGroupRequest{} }
func (m *UserGroupRequest) String() string            { return proto.CompactTextString(m) }
func (*UserGroupRequest) ProtoMessage()               {}
func (*UserGroupRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *UserGroupRequest) GetGroup() Group {
	if m != nil {
		return m.Group
	}
	return Group_USER
}

type UsersResponse struct {
	Group Group           `protobuf:"varint,1,opt,name=group,enum=gateway.Group" json:"group,omitempty"`
	Users []*UserResponse `protobuf:"bytes,2,rep,name=users" json:"users,omitempty"`
}

func (m *UsersResponse) Reset()                    { *m = UsersResponse{} }
func (m *UsersResponse) String() string            { return proto.CompactTextString(m) }
func (*UsersResponse) ProtoMessage()               {}
func (*UsersResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *UsersResponse) GetGroup() Group {
	if m != nil {
		return m.Group
	}
	return Group_USER
}

func (m *UsersResponse) GetUsers() []*UserResponse {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*GetUserRequest)(nil), "gateway.GetUserRequest")
	proto.RegisterType((*UpdateUserRequest)(nil), "gateway.UpdateUserRequest")
	proto.RegisterType((*UserResponse)(nil), "gateway.UserResponse")
	proto.RegisterType((*UserGroupRequest)(nil), "gateway.UserGroupRequest")
	proto.RegisterType((*UsersResponse)(nil), "gateway.UsersResponse")
	proto.RegisterEnum("gateway.Group", Group_name, Group_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for UserService service

type UserServiceClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetUsersByGroup(ctx context.Context, in *UserGroupRequest, opts ...grpc.CallOption) (*UsersResponse, error)
	UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*Empty, error)
}

type userServiceClient struct {
	cc *grpc.ClientConn
}

func NewUserServiceClient(cc *grpc.ClientConn) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := grpc.Invoke(ctx, "/gateway.UserService/GetUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUsersByGroup(ctx context.Context, in *UserGroupRequest, opts ...grpc.CallOption) (*UsersResponse, error) {
	out := new(UsersResponse)
	err := grpc.Invoke(ctx, "/gateway.UserService/GetUsersByGroup", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUser(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/gateway.UserService/UpdateUser", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceServer interface {
	GetUser(context.Context, *GetUserRequest) (*UserResponse, error)
	GetUsersByGroup(context.Context, *UserGroupRequest) (*UsersResponse, error)
	UpdateUser(context.Context, *UpdateUserRequest) (*Empty, error)
}

func RegisterUserServiceServer(s *grpc.Server, srv UserServiceServer) {
	s.RegisterService(&_UserService_serviceDesc, srv)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.UserService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUsersByGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserGroupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUsersByGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.UserService/GetUsersByGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUsersByGroup(ctx, req.(*UserGroupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.UserService/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUser(ctx, req.(*UpdateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "GetUsersByGroup",
			Handler:    _UserService_GetUsersByGroup_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserService_UpdateUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}

func init() { proto.RegisterFile("user.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xd1, 0x6a, 0xe2, 0x40,
	0x18, 0x85, 0x37, 0xd1, 0xac, 0xeb, 0xaf, 0xc6, 0xec, 0x0f, 0xee, 0x66, 0xc3, 0x5e, 0x84, 0xb0,
	0x17, 0xa2, 0x60, 0xc0, 0xbd, 0x29, 0xbd, 0x6b, 0x31, 0x88, 0x50, 0x6b, 0x89, 0xd8, 0xfb, 0xb1,
	0x19, 0x42, 0x40, 0x33, 0x69, 0x66, 0xb4, 0x48, 0xe9, 0x8d, 0xaf, 0xd0, 0x47, 0xeb, 0x2b, 0xf4,
	0x41, 0x4a, 0x26, 0xd6, 0x68, 0x29, 0xa5, 0xf4, 0x6e, 0x26, 0x73, 0xce, 0x97, 0x33, 0x67, 0x7e,
	0x80, 0x15, 0xa7, 0x69, 0x2f, 0x49, 0x99, 0x60, 0x58, 0x09, 0x89, 0xa0, 0x77, 0x64, 0x63, 0xfd,
	0x0d, 0x19, 0x0b, 0x17, 0xd4, 0x25, 0x49, 0xe4, 0x92, 0x38, 0x66, 0x82, 0x88, 0x88, 0xc5, 0x3c,
	0x97, 0x59, 0x35, 0xb2, 0x88, 0xd6, 0x34, 0xdf, 0x38, 0x36, 0xe8, 0x43, 0x2a, 0x66, 0x9c, 0xa6,
	0x3e, 0xbd, 0x5d, 0x51, 0x2e, 0x50, 0x07, 0x35, 0x0a, 0x4c, 0xc5, 0x56, 0xda, 0x55, 0x5f, 0x8d,
	0x02, 0x67, 0x04, 0x3f, 0x67, 0x49, 0x40, 0x04, 0xfd, 0x40, 0x84, 0x08, 0xe5, 0x98, 0x2c, 0xa9,
	0xa9, 0xca, 0x2f, 0x72, 0x8d, 0x06, 0x94, 0x48, 0x48, 0xcd, 0x92, 0xad, 0xb4, 0x35, 0x3f, 0x5b,
	0x3a, 0x03, 0xa8, 0xe7, 0x10, 0x9e, 0xb0, 0x98, 0xd3, 0x2f, 0x52, 0x4e, 0xc0, 0xc8, 0x28, 0xc3,
	0x94, 0xad, 0x92, 0xd7, 0x3c, 0xff, 0x40, 0x0b, 0xb3, 0xbd, 0x84, 0xe9, 0x7d, 0xbd, 0xb7, 0xab,
	0xa2, 0x97, 0xab, 0xf2, 0x43, 0x67, 0x0e, 0x8d, 0xcc, 0xc9, 0xf7, 0x01, 0x3e, 0x65, 0xc3, 0x2e,
	0x68, 0x59, 0xcb, 0xdc, 0x54, 0xed, 0x52, 0xbb, 0xd6, 0x6f, 0xed, 0x55, 0x87, 0x97, 0xf1, 0x73,
	0x4d, 0xa7, 0x0b, 0x9a, 0x34, 0xe3, 0x0f, 0x28, 0xcf, 0xa6, 0x9e, 0x6f, 0x7c, 0xc3, 0x2a, 0x68,
	0x67, 0x83, 0xf1, 0xe8, 0xd2, 0x50, 0xb0, 0x01, 0xd5, 0x81, 0x77, 0xed, 0x5d, 0x4c, 0xae, 0x3c,
	0xdf, 0x50, 0xfb, 0x5b, 0x15, 0x6a, 0x19, 0x64, 0x4a, 0xd3, 0x75, 0x74, 0x43, 0x71, 0x0c, 0x95,
	0xdd, 0x6b, 0xe0, 0xef, 0x22, 0xcb, 0xd1, 0xfb, 0x58, 0xef, 0xff, 0xde, 0xc1, 0xed, 0xd3, 0xf3,
	0xa3, 0x5a, 0x47, 0x70, 0xb3, 0x1c, 0xee, 0x7d, 0x14, 0x3c, 0xe0, 0x14, 0x9a, 0x3b, 0x33, 0x3f,
	0xdf, 0xe4, 0xa9, 0xfe, 0x1c, 0xb9, 0x0f, 0x3b, 0xb4, 0x7e, 0x1d, 0x1d, 0xed, 0x4b, 0x72, 0x1a,
	0x92, 0x5c, 0x41, 0x4d, 0x92, 0x71, 0x02, 0x50, 0xcc, 0x03, 0x5a, 0x85, 0xe9, 0xed, 0x90, 0x58,
	0x45, 0x9d, 0xde, 0x32, 0x11, 0x1b, 0xa7, 0x25, 0x41, 0x4d, 0xeb, 0x20, 0xe2, 0xa9, 0xd2, 0x99,
	0x7f, 0x97, 0x93, 0xf8, 0xff, 0x25, 0x00, 0x00, 0xff, 0xff, 0x09, 0xa3, 0x0b, 0xd3, 0xcb, 0x02,
	0x00, 0x00,
}
