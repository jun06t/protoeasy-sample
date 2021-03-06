// Code generated by protoc-gen-go. DO NOT EDIT.
// source: uploader.proto

package upload

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

type Chunk struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Chunk) Reset()         { *m = Chunk{} }
func (m *Chunk) String() string { return proto.CompactTextString(m) }
func (*Chunk) ProtoMessage()    {}
func (*Chunk) Descriptor() ([]byte, []int) {
	return fileDescriptor_b055a52f625709c9, []int{0}
}

func (m *Chunk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Chunk.Unmarshal(m, b)
}
func (m *Chunk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Chunk.Marshal(b, m, deterministic)
}
func (m *Chunk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Chunk.Merge(m, src)
}
func (m *Chunk) XXX_Size() int {
	return xxx_messageInfo_Chunk.Size(m)
}
func (m *Chunk) XXX_DiscardUnknown() {
	xxx_messageInfo_Chunk.DiscardUnknown(m)
}

var xxx_messageInfo_Chunk proto.InternalMessageInfo

func (m *Chunk) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type UploadResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadResponse) Reset()         { *m = UploadResponse{} }
func (m *UploadResponse) String() string { return proto.CompactTextString(m) }
func (*UploadResponse) ProtoMessage()    {}
func (*UploadResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b055a52f625709c9, []int{1}
}

func (m *UploadResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadResponse.Unmarshal(m, b)
}
func (m *UploadResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadResponse.Marshal(b, m, deterministic)
}
func (m *UploadResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadResponse.Merge(m, src)
}
func (m *UploadResponse) XXX_Size() int {
	return xxx_messageInfo_UploadResponse.Size(m)
}
func (m *UploadResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadResponse proto.InternalMessageInfo

func (m *UploadResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*Chunk)(nil), "upload.Chunk")
	proto.RegisterType((*UploadResponse)(nil), "upload.UploadResponse")
}

func init() { proto.RegisterFile("uploader.proto", fileDescriptor_b055a52f625709c9) }

var fileDescriptor_b055a52f625709c9 = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x2d, 0xc8, 0xc9,
	0x4f, 0x4c, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0xf0, 0x95, 0xa4,
	0xb9, 0x58, 0x9d, 0x33, 0x4a, 0xf3, 0xb2, 0x85, 0x84, 0xb8, 0x58, 0x52, 0x12, 0x4b, 0x12, 0x25,
	0x18, 0x15, 0x18, 0x35, 0x78, 0x82, 0xc0, 0x6c, 0x25, 0x0d, 0x2e, 0xbe, 0x50, 0xb0, 0xb2, 0xa0,
	0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x31, 0x2e, 0xb6, 0xe2, 0x92, 0xc4, 0x92, 0xd2,
	0x62, 0xb0, 0x3a, 0xce, 0x20, 0x28, 0xcf, 0xc8, 0x9e, 0x8b, 0x23, 0x14, 0x6a, 0x81, 0x90, 0x31,
	0x17, 0x1b, 0x84, 0x2d, 0xc4, 0xab, 0x07, 0xb1, 0x45, 0x0f, 0x6c, 0x85, 0x94, 0x18, 0x8c, 0x8b,
	0x6a, 0xa8, 0x12, 0x83, 0x06, 0x63, 0x12, 0x1b, 0xd8, 0x59, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xda, 0xb2, 0x75, 0x29, 0xa8, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// UploaderClient is the client API for Uploader service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UploaderClient interface {
	Upload(ctx context.Context, opts ...grpc.CallOption) (Uploader_UploadClient, error)
}

type uploaderClient struct {
	cc *grpc.ClientConn
}

func NewUploaderClient(cc *grpc.ClientConn) UploaderClient {
	return &uploaderClient{cc}
}

func (c *uploaderClient) Upload(ctx context.Context, opts ...grpc.CallOption) (Uploader_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Uploader_serviceDesc.Streams[0], "/upload.Uploader/Upload", opts...)
	if err != nil {
		return nil, err
	}
	x := &uploaderUploadClient{stream}
	return x, nil
}

type Uploader_UploadClient interface {
	Send(*Chunk) error
	CloseAndRecv() (*UploadResponse, error)
	grpc.ClientStream
}

type uploaderUploadClient struct {
	grpc.ClientStream
}

func (x *uploaderUploadClient) Send(m *Chunk) error {
	return x.ClientStream.SendMsg(m)
}

func (x *uploaderUploadClient) CloseAndRecv() (*UploadResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// UploaderServer is the server API for Uploader service.
type UploaderServer interface {
	Upload(Uploader_UploadServer) error
}

// UnimplementedUploaderServer can be embedded to have forward compatible implementations.
type UnimplementedUploaderServer struct {
}

func (*UnimplementedUploaderServer) Upload(srv Uploader_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}

func RegisterUploaderServer(s *grpc.Server, srv UploaderServer) {
	s.RegisterService(&_Uploader_serviceDesc, srv)
}

func _Uploader_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UploaderServer).Upload(&uploaderUploadServer{stream})
}

type Uploader_UploadServer interface {
	SendAndClose(*UploadResponse) error
	Recv() (*Chunk, error)
	grpc.ServerStream
}

type uploaderUploadServer struct {
	grpc.ServerStream
}

func (x *uploaderUploadServer) SendAndClose(m *UploadResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *uploaderUploadServer) Recv() (*Chunk, error) {
	m := new(Chunk)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Uploader_serviceDesc = grpc.ServiceDesc{
	ServiceName: "upload.Uploader",
	HandlerType: (*UploaderServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Upload",
			Handler:       _Uploader_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "uploader.proto",
}
