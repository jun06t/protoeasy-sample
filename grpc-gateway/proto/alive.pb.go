// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alive.proto

package gateway

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_90fb53809164d43e, []int{0}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type AliveResponse struct {
	Status               bool     `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AliveResponse) Reset()         { *m = AliveResponse{} }
func (m *AliveResponse) String() string { return proto.CompactTextString(m) }
func (*AliveResponse) ProtoMessage()    {}
func (*AliveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_90fb53809164d43e, []int{1}
}

func (m *AliveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AliveResponse.Unmarshal(m, b)
}
func (m *AliveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AliveResponse.Marshal(b, m, deterministic)
}
func (m *AliveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AliveResponse.Merge(m, src)
}
func (m *AliveResponse) XXX_Size() int {
	return xxx_messageInfo_AliveResponse.Size(m)
}
func (m *AliveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AliveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AliveResponse proto.InternalMessageInfo

func (m *AliveResponse) GetStatus() bool {
	if m != nil {
		return m.Status
	}
	return false
}

func init() {
	proto.RegisterType((*Empty)(nil), "gateway.Empty")
	proto.RegisterType((*AliveResponse)(nil), "gateway.AliveResponse")
}

func init() { proto.RegisterFile("alive.proto", fileDescriptor_90fb53809164d43e) }

var fileDescriptor_90fb53809164d43e = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4e, 0xcc, 0xc9, 0x2c,
	0x4b, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0x4f, 0x2c, 0x49, 0x2d, 0x4f, 0xac,
	0x94, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb,
	0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0x28, 0x53, 0x62, 0xe7, 0x62, 0x75, 0xcd,
	0x2d, 0x28, 0xa9, 0x54, 0x52, 0xe7, 0xe2, 0x75, 0x04, 0x69, 0x0f, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf,
	0x2b, 0x4e, 0x15, 0x12, 0xe3, 0x62, 0x2b, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x96, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x08, 0x82, 0xf2, 0x8c, 0x82, 0xb9, 0x78, 0xc0, 0x0a, 0x83, 0x53, 0x8b, 0xca, 0x32,
	0x93, 0x53, 0x85, 0x9c, 0xb9, 0x38, 0xdd, 0x53, 0x4b, 0x82, 0xc1, 0x92, 0x42, 0x7c, 0x7a, 0x50,
	0x6b, 0xf5, 0xc0, 0xa6, 0x4a, 0x89, 0xc1, 0xf9, 0x28, 0x86, 0x2b, 0xf1, 0x35, 0x5d, 0x7e, 0x32,
	0x99, 0x89, 0x43, 0x88, 0x4d, 0x1f, 0xec, 0xe6, 0x24, 0x36, 0xb0, 0x6b, 0x8c, 0x01, 0x01, 0x00,
	0x00, 0xff, 0xff, 0xf1, 0x17, 0xe4, 0xdc, 0xc3, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AliveServiceClient is the client API for AliveService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AliveServiceClient interface {
	GetStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AliveResponse, error)
}

type aliveServiceClient struct {
	cc *grpc.ClientConn
}

func NewAliveServiceClient(cc *grpc.ClientConn) AliveServiceClient {
	return &aliveServiceClient{cc}
}

func (c *aliveServiceClient) GetStatus(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AliveResponse, error) {
	out := new(AliveResponse)
	err := c.cc.Invoke(ctx, "/gateway.AliveService/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AliveServiceServer is the server API for AliveService service.
type AliveServiceServer interface {
	GetStatus(context.Context, *Empty) (*AliveResponse, error)
}

// UnimplementedAliveServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAliveServiceServer struct {
}

func (*UnimplementedAliveServiceServer) GetStatus(ctx context.Context, req *Empty) (*AliveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}

func RegisterAliveServiceServer(s *grpc.Server, srv AliveServiceServer) {
	s.RegisterService(&_AliveService_serviceDesc, srv)
}

func _AliveService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AliveServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gateway.AliveService/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AliveServiceServer).GetStatus(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _AliveService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "gateway.AliveService",
	HandlerType: (*AliveServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetStatus",
			Handler:    _AliveService_GetStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alive.proto",
}
