// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package helloworld

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RoleType int32

const (
	RoleType_NONE     RoleType = 0
	RoleType_STANDARD RoleType = 1
	RoleType_PREMIUM  RoleType = 2
	RoleType_ULTIMATE RoleType = 3
)

var RoleType_name = map[int32]string{
	0: "NONE",
	1: "STANDARD",
	2: "PREMIUM",
	3: "ULTIMATE",
}

var RoleType_value = map[string]int32{
	"NONE":     0,
	"STANDARD": 1,
	"PREMIUM":  2,
	"ULTIMATE": 3,
}

func (x RoleType) String() string {
	return proto.EnumName(RoleType_name, int32(x))
}

func (RoleType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{0}
}

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloReply struct {
	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	//
	//enum Role {
	//NONE     = 0;
	//STANDARD = 1;
	//PREMIUM  = 2;
	//ULTIMATE = 3;
	//}
	//Role role = 2;
	Role                 RoleType `protobuf:"varint,2,opt,name=role,proto3,enum=helloworld.RoleType" json:"role,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *HelloReply) GetRole() RoleType {
	if m != nil {
		return m.Role
	}
	return RoleType_NONE
}

func init() {
	proto.RegisterEnum("helloworld.RoleType", RoleType_name, RoleType_value)
	proto.RegisterType((*HelloRequest)(nil), "helloworld.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "helloworld.HelloReply")
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor_17b8c58d586b62f2) }

var fileDescriptor_17b8c58d586b62f2 = []byte{
	// 225 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x37, 0x6b, 0xb1, 0x71, 0x5c, 0x24, 0x0c, 0x22, 0xc1, 0xd3, 0x92, 0x53, 0xf1, 0xb0,
	0x87, 0xf5, 0xac, 0x50, 0xd8, 0xa2, 0x05, 0x5b, 0x97, 0x6c, 0xf6, 0x07, 0x54, 0x1c, 0xf4, 0x90,
	0x35, 0x31, 0xad, 0x48, 0xfe, 0xbd, 0x6c, 0xb0, 0xd8, 0x83, 0xb7, 0x79, 0xef, 0x7d, 0xf0, 0x66,
	0x06, 0xc4, 0x3b, 0x59, 0xeb, 0xbe, 0x5d, 0xb0, 0xaf, 0x2b, 0x1f, 0xdc, 0xe0, 0x10, 0xfe, 0x1c,
	0xa5, 0x60, 0xf1, 0x78, 0x54, 0x9a, 0x3e, 0xbf, 0xa8, 0x1f, 0x10, 0x21, 0xfb, 0xe8, 0x0e, 0x24,
	0xd9, 0x92, 0x15, 0x67, 0x3a, 0xcd, 0x6a, 0x0b, 0xf0, 0xcb, 0x78, 0x1b, 0x51, 0x42, 0x7e, 0xa0,
	0xbe, 0xef, 0xde, 0x46, 0x68, 0x94, 0x58, 0x40, 0x16, 0x9c, 0x25, 0x39, 0x5f, 0xb2, 0xe2, 0x62,
	0x7d, 0xb9, 0x9a, 0x14, 0x6b, 0x67, 0xc9, 0x44, 0x4f, 0x3a, 0x11, 0x37, 0x77, 0xc0, 0x47, 0x07,
	0x39, 0x64, 0xed, 0x73, 0x5b, 0x89, 0x19, 0x2e, 0x80, 0xef, 0x4c, 0xd9, 0x6e, 0x4a, 0xbd, 0x11,
	0x0c, 0xcf, 0x21, 0xdf, 0xea, 0xaa, 0xa9, 0xf7, 0x8d, 0x98, 0x1f, 0xa3, 0xfd, 0x93, 0xa9, 0x9b,
	0xd2, 0x54, 0xe2, 0x64, 0x5d, 0x43, 0xfe, 0x10, 0x88, 0x06, 0x0a, 0x78, 0x0f, 0x7c, 0xd7, 0xc5,
	0xb4, 0x1e, 0xca, 0x69, 0xe3, 0xf4, 0xaa, 0xeb, 0xab, 0x7f, 0x12, 0x6f, 0xa3, 0x9a, 0xbd, 0x9c,
	0xa6, 0x97, 0xdc, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x70, 0x7d, 0x6f, 0x26, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type greeterClient struct {
	cc *grpc.ClientConn
}

func NewGreeterClient(cc *grpc.ClientConn) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/helloworld.Greeter/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/helloworld.Greeter/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "helloworld.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Greeter_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "helloworld.proto",
}
