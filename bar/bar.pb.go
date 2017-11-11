// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: bar/bar.proto

/*
	Package bar is a generated protocol buffer package.

	It is generated from these files:
		bar/bar.proto

	It has these top-level messages:
		Two
*/
package bar

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import google_protobuf1 "go.pedge.io/pb/gogo/google/protobuf"
import foo "."
import _ "github.com/gogo/protobuf/gogoproto"

import context "golang.org/x/net/context"
import grpc "google.golang.org/grpc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Two struct {
	One *foo.One `protobuf:"bytes,1,opt,name=one" json:"one,omitempty"`
	J   int64    `protobuf:"varint,2,opt,name=j,proto3" json:"j,omitempty"`
}

func (m *Two) Reset()                    { *m = Two{} }
func (m *Two) String() string            { return proto.CompactTextString(m) }
func (*Two) ProtoMessage()               {}
func (*Two) Descriptor() ([]byte, []int) { return fileDescriptorBar, []int{0} }

func (m *Two) GetOne() *foo.One {
	if m != nil {
		return m.One
	}
	return nil
}

func (m *Two) GetJ() int64 {
	if m != nil {
		return m.J
	}
	return 0
}

func init() {
	proto.RegisterType((*Two)(nil), "bar.Two")
}
func (this *Two) Equal(that interface{}) bool {
	if that == nil {
		if this == nil {
			return true
		}
		return false
	}

	that1, ok := that.(*Two)
	if !ok {
		that2, ok := that.(Two)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		if this == nil {
			return true
		}
		return false
	} else if this == nil {
		return false
	}
	if !this.One.Equal(that1.One) {
		return false
	}
	if this.J != that1.J {
		return false
	}
	return true
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for API service

type APIClient interface {
	Do(ctx context.Context, in *Two, opts ...grpc.CallOption) (*google_protobuf1.Empty, error)
}

type aPIClient struct {
	cc *grpc.ClientConn
}

func NewAPIClient(cc *grpc.ClientConn) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) Do(ctx context.Context, in *Two, opts ...grpc.CallOption) (*google_protobuf1.Empty, error) {
	out := new(google_protobuf1.Empty)
	err := grpc.Invoke(ctx, "/bar.API/Do", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for API service

type APIServer interface {
	Do(context.Context, *Two) (*google_protobuf1.Empty, error)
}

func RegisterAPIServer(s *grpc.Server, srv APIServer) {
	s.RegisterService(&_API_serviceDesc, srv)
}

func _API_Do_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Two)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Do(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/bar.API/Do",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Do(ctx, req.(*Two))
	}
	return interceptor(ctx, in, info, handler)
}

var _API_serviceDesc = grpc.ServiceDesc{
	ServiceName: "bar.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Do",
			Handler:    _API_Do_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "bar/bar.proto",
}

func (m *Two) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Two) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.One != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintBar(dAtA, i, uint64(m.One.Size()))
		n1, err := m.One.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.J != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintBar(dAtA, i, uint64(m.J))
	}
	return i, nil
}

func encodeVarintBar(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func NewPopulatedTwo(r randyBar, easy bool) *Two {
	this := &Two{}
	if r.Intn(10) != 0 {
		this.One = foo.NewPopulatedOne(r, easy)
	}
	this.J = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.J *= -1
	}
	if !easy && r.Intn(10) != 0 {
	}
	return this
}

type randyBar interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneBar(r randyBar) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringBar(r randyBar) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneBar(r)
	}
	return string(tmps)
}
func randUnrecognizedBar(r randyBar, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldBar(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldBar(dAtA []byte, r randyBar, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateBar(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateBar(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateBar(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateBar(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateBar(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateBar(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateBar(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *Two) Size() (n int) {
	var l int
	_ = l
	if m.One != nil {
		l = m.One.Size()
		n += 1 + l + sovBar(uint64(l))
	}
	if m.J != 0 {
		n += 1 + sovBar(uint64(m.J))
	}
	return n
}

func sovBar(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBar(x uint64) (n int) {
	return sovBar(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Two) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBar
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Two: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Two: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field One", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBar
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBar
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.One == nil {
				m.One = &foo.One{}
			}
			if err := m.One.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field J", wireType)
			}
			m.J = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBar
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.J |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBar(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBar
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBar(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBar
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBar
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowBar
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthBar
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBar
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipBar(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthBar = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBar   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("bar/bar.proto", fileDescriptorBar) }

var fileDescriptorBar = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4a, 0x2c, 0xd2,
	0x4f, 0x4a, 0x2c, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x4a, 0x2c, 0x92, 0x92,
	0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f,
	0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0x28, 0x91, 0x92, 0x86, 0xca, 0x82, 0x79, 0x49, 0xa5,
	0x69, 0xfa, 0xa9, 0xb9, 0x05, 0x25, 0x95, 0x50, 0x49, 0xce, 0xb4, 0xfc, 0x7c, 0x28, 0x53, 0x24,
	0x3d, 0x3f, 0x3d, 0x1f, 0xcc, 0xd4, 0x07, 0xb1, 0x20, 0xa2, 0x4a, 0xfa, 0x5c, 0xcc, 0x21, 0xe5,
	0xf9, 0x42, 0x52, 0x5c, 0xcc, 0xf9, 0x79, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x1c,
	0x7a, 0x20, 0x5d, 0xfe, 0x79, 0xa9, 0x41, 0x20, 0x41, 0x21, 0x1e, 0x2e, 0xc6, 0x2c, 0x09, 0x26,
	0x05, 0x46, 0x0d, 0xe6, 0x20, 0xc6, 0x2c, 0x23, 0x5b, 0x2e, 0x66, 0xc7, 0x00, 0x4f, 0x21, 0x33,
	0x2e, 0x26, 0x97, 0x7c, 0x21, 0x0e, 0x3d, 0x90, 0x53, 0x43, 0xca, 0xf3, 0xa5, 0xc4, 0xf4, 0x20,
	0xce, 0xd0, 0x83, 0x39, 0x43, 0xcf, 0x15, 0xe4, 0x0c, 0x25, 0xbe, 0xa6, 0xcb, 0x4f, 0x26, 0x33,
	0x71, 0x28, 0x31, 0xeb, 0xa7, 0xe4, 0x5b, 0x31, 0x6a, 0x39, 0x89, 0xfc, 0x78, 0x28, 0xc7, 0xb8,
	0xe2, 0x91, 0x1c, 0xe3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7,
	0x98, 0xc4, 0x06, 0xd6, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x70, 0x4d, 0xab, 0xf7, 0xfe,
	0x00, 0x00, 0x00,
}
