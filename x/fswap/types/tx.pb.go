// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lbm/fswap/v1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/Finschia/finschia-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgSwapRequest struct {
	// holder's address
	FromAddress string `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
	// amount of old currency
	Amount types.Coin `protobuf:"bytes,2,opt,name=amount,proto3,castrepeated=github.com/Finschia/finschia-sdk/types.Coin" json:"amount"`
}

func (m *MsgSwapRequest) Reset()         { *m = MsgSwapRequest{} }
func (m *MsgSwapRequest) String() string { return proto.CompactTextString(m) }
func (*MsgSwapRequest) ProtoMessage()    {}
func (*MsgSwapRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65c77cf1d9b67323, []int{0}
}
func (m *MsgSwapRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapRequest.Merge(m, src)
}
func (m *MsgSwapRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapRequest proto.InternalMessageInfo

func (m *MsgSwapRequest) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

func (m *MsgSwapRequest) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

type MsgSwapResponse struct {
}

func (m *MsgSwapResponse) Reset()         { *m = MsgSwapResponse{} }
func (m *MsgSwapResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSwapResponse) ProtoMessage()    {}
func (*MsgSwapResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_65c77cf1d9b67323, []int{1}
}
func (m *MsgSwapResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapResponse.Merge(m, src)
}
func (m *MsgSwapResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapResponse proto.InternalMessageInfo

type MsgSwapAllRequest struct {
	// holder's address
	FromAddress string `protobuf:"bytes,1,opt,name=from_address,json=fromAddress,proto3" json:"from_address,omitempty"`
}

func (m *MsgSwapAllRequest) Reset()         { *m = MsgSwapAllRequest{} }
func (m *MsgSwapAllRequest) String() string { return proto.CompactTextString(m) }
func (*MsgSwapAllRequest) ProtoMessage()    {}
func (*MsgSwapAllRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_65c77cf1d9b67323, []int{2}
}
func (m *MsgSwapAllRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapAllRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapAllRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapAllRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapAllRequest.Merge(m, src)
}
func (m *MsgSwapAllRequest) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapAllRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapAllRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapAllRequest proto.InternalMessageInfo

func (m *MsgSwapAllRequest) GetFromAddress() string {
	if m != nil {
		return m.FromAddress
	}
	return ""
}

type MsgSwapAllResponse struct {
	// amount of swapped in new currency
	Amount types.Coin `protobuf:"bytes,1,opt,name=amount,proto3,castrepeated=github.com/Finschia/finschia-sdk/types.Coin" json:"amount"`
}

func (m *MsgSwapAllResponse) Reset()         { *m = MsgSwapAllResponse{} }
func (m *MsgSwapAllResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSwapAllResponse) ProtoMessage()    {}
func (*MsgSwapAllResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_65c77cf1d9b67323, []int{3}
}
func (m *MsgSwapAllResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSwapAllResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSwapAllResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSwapAllResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSwapAllResponse.Merge(m, src)
}
func (m *MsgSwapAllResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSwapAllResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSwapAllResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSwapAllResponse proto.InternalMessageInfo

func (m *MsgSwapAllResponse) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

func init() {
	proto.RegisterType((*MsgSwapRequest)(nil), "lbm.fswap.v1.MsgSwapRequest")
	proto.RegisterType((*MsgSwapResponse)(nil), "lbm.fswap.v1.MsgSwapResponse")
	proto.RegisterType((*MsgSwapAllRequest)(nil), "lbm.fswap.v1.MsgSwapAllRequest")
	proto.RegisterType((*MsgSwapAllResponse)(nil), "lbm.fswap.v1.MsgSwapAllResponse")
}

func init() { proto.RegisterFile("lbm/fswap/v1/tx.proto", fileDescriptor_65c77cf1d9b67323) }

var fileDescriptor_65c77cf1d9b67323 = []byte{
	// 356 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xcd, 0x49, 0xca, 0xd5,
	0x4f, 0x2b, 0x2e, 0x4f, 0x2c, 0xd0, 0x2f, 0x33, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0xc9, 0x49, 0xca, 0xd5, 0x03, 0x0b, 0xeb, 0x95, 0x19, 0x4a, 0x89, 0xa4, 0xe7,
	0xa7, 0xe7, 0x83, 0x25, 0xf4, 0x41, 0x2c, 0x88, 0x1a, 0x29, 0xb9, 0xe4, 0xfc, 0xe2, 0xdc, 0xfc,
	0x62, 0xfd, 0xa4, 0xc4, 0xe2, 0x54, 0xfd, 0x32, 0xc3, 0xa4, 0xd4, 0x92, 0x44, 0x43, 0xfd, 0xe4,
	0xfc, 0xcc, 0x3c, 0x88, 0xbc, 0xd2, 0x6c, 0x46, 0x2e, 0x3e, 0xdf, 0xe2, 0xf4, 0xe0, 0xf2, 0xc4,
	0x82, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x45, 0x2e, 0x9e, 0xb4, 0xa2, 0xfc, 0xdc,
	0xf8, 0xc4, 0x94, 0x94, 0xa2, 0xd4, 0xe2, 0x62, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x6e,
	0x90, 0x98, 0x23, 0x44, 0x48, 0x28, 0x8d, 0x8b, 0x2d, 0x31, 0x37, 0xbf, 0x34, 0xaf, 0x44, 0x82,
	0x49, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x52, 0x0f, 0x62, 0x8d, 0x1e, 0xc8, 0x1a, 0x3d, 0xa8, 0x35,
	0x7a, 0xce, 0xf9, 0x99, 0x79, 0x4e, 0xc6, 0x27, 0xee, 0xc9, 0x33, 0xac, 0xba, 0x2f, 0xaf, 0x9d,
	0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0xef, 0x96, 0x99, 0x57, 0x9c, 0x9c,
	0x91, 0x99, 0xa8, 0x9f, 0x06, 0x65, 0xe8, 0x16, 0xa7, 0x64, 0xeb, 0x97, 0x54, 0x16, 0xa4, 0x16,
	0x83, 0x35, 0x05, 0x41, 0x4d, 0x57, 0x12, 0xe4, 0xe2, 0x87, 0x3b, 0xae, 0xb8, 0x20, 0x3f, 0xaf,
	0x38, 0x55, 0xc9, 0x8c, 0x4b, 0x10, 0x2a, 0xe4, 0x98, 0x93, 0x43, 0xbc, 0x93, 0x95, 0x6a, 0xb8,
	0x84, 0x90, 0xf5, 0x41, 0x4c, 0x43, 0xf2, 0x08, 0x23, 0x2d, 0x3d, 0x62, 0x34, 0x83, 0x91, 0x8b,
	0xd9, 0xb7, 0x38, 0x5d, 0xc8, 0x99, 0x8b, 0x05, 0xe4, 0x04, 0x21, 0x19, 0x3d, 0xe4, 0xb8, 0xd3,
	0x43, 0x8d, 0x01, 0x29, 0x59, 0x1c, 0xb2, 0x50, 0x47, 0xfb, 0x70, 0xb1, 0x43, 0xfd, 0x21, 0x24,
	0x8f, 0x55, 0x25, 0x22, 0x64, 0xa4, 0x14, 0x70, 0x2b, 0x80, 0x98, 0xe6, 0xe4, 0x71, 0xe2, 0x91,
	0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1,
	0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0x7a, 0x04, 0x7d, 0x5a, 0x01, 0x4d, 0x95, 0x60,
	0x1f, 0x27, 0xb1, 0x81, 0x93, 0x94, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x66, 0x25, 0x56, 0x7f,
	0xaf, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	Swap(ctx context.Context, in *MsgSwapRequest, opts ...grpc.CallOption) (*MsgSwapResponse, error)
	SwapAll(ctx context.Context, in *MsgSwapAllRequest, opts ...grpc.CallOption) (*MsgSwapAllResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Swap(ctx context.Context, in *MsgSwapRequest, opts ...grpc.CallOption) (*MsgSwapResponse, error) {
	out := new(MsgSwapResponse)
	err := c.cc.Invoke(ctx, "/lbm.fswap.v1.Msg/Swap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SwapAll(ctx context.Context, in *MsgSwapAllRequest, opts ...grpc.CallOption) (*MsgSwapAllResponse, error) {
	out := new(MsgSwapAllResponse)
	err := c.cc.Invoke(ctx, "/lbm.fswap.v1.Msg/SwapAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Swap(context.Context, *MsgSwapRequest) (*MsgSwapResponse, error)
	SwapAll(context.Context, *MsgSwapAllRequest) (*MsgSwapAllResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Swap(ctx context.Context, req *MsgSwapRequest) (*MsgSwapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Swap not implemented")
}
func (*UnimplementedMsgServer) SwapAll(ctx context.Context, req *MsgSwapAllRequest) (*MsgSwapAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SwapAll not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Swap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSwapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Swap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lbm.fswap.v1.Msg/Swap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Swap(ctx, req.(*MsgSwapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SwapAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSwapAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SwapAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/lbm.fswap.v1.Msg/SwapAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SwapAll(ctx, req.(*MsgSwapAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "lbm.fswap.v1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Swap",
			Handler:    _Msg_Swap_Handler,
		},
		{
			MethodName: "SwapAll",
			Handler:    _Msg_SwapAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lbm/fswap/v1/tx.proto",
}

func (m *MsgSwapRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSwapResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSwapAllRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapAllRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapAllRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FromAddress) > 0 {
		i -= len(m.FromAddress)
		copy(dAtA[i:], m.FromAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FromAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSwapAllResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSwapAllResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSwapAllResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSwapRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgSwapResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSwapAllRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FromAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSwapAllResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSwapRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSwapRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSwapResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSwapResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSwapAllRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSwapAllRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapAllRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgSwapAllResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSwapAllResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSwapAllResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
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
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)