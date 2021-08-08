// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: voter/tx.proto

package types

import (
	context "context"
	fmt "fmt"
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

// this line is used by starport scaffolding # proto/tx/message
type MsgCreatePoll struct {
	Creator string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Title   string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Options []string `protobuf:"bytes,3,rep,name=options,proto3" json:"options,omitempty"`
}

func (m *MsgCreatePoll) Reset()         { *m = MsgCreatePoll{} }
func (m *MsgCreatePoll) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePoll) ProtoMessage()    {}
func (*MsgCreatePoll) Descriptor() ([]byte, []int) {
	return fileDescriptor_779ae7fcafad750a, []int{0}
}
func (m *MsgCreatePoll) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePoll) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePoll.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePoll) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePoll.Merge(m, src)
}
func (m *MsgCreatePoll) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePoll) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePoll.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePoll proto.InternalMessageInfo

func (m *MsgCreatePoll) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreatePoll) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MsgCreatePoll) GetOptions() []string {
	if m != nil {
		return m.Options
	}
	return nil
}

type MsgCreatePollResponse struct {
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgCreatePollResponse) Reset()         { *m = MsgCreatePollResponse{} }
func (m *MsgCreatePollResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreatePollResponse) ProtoMessage()    {}
func (*MsgCreatePollResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_779ae7fcafad750a, []int{1}
}
func (m *MsgCreatePollResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreatePollResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreatePollResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreatePollResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreatePollResponse.Merge(m, src)
}
func (m *MsgCreatePollResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreatePollResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreatePollResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreatePollResponse proto.InternalMessageInfo

func (m *MsgCreatePollResponse) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type MsgUpdatePoll struct {
	Creator string   `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      uint64   `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Title   string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Options []string `protobuf:"bytes,4,rep,name=options,proto3" json:"options,omitempty"`
}

func (m *MsgUpdatePoll) Reset()         { *m = MsgUpdatePoll{} }
func (m *MsgUpdatePoll) String() string { return proto.CompactTextString(m) }
func (*MsgUpdatePoll) ProtoMessage()    {}
func (*MsgUpdatePoll) Descriptor() ([]byte, []int) {
	return fileDescriptor_779ae7fcafad750a, []int{2}
}
func (m *MsgUpdatePoll) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdatePoll) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdatePoll.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdatePoll) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdatePoll.Merge(m, src)
}
func (m *MsgUpdatePoll) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdatePoll) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdatePoll.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdatePoll proto.InternalMessageInfo

func (m *MsgUpdatePoll) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdatePoll) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MsgUpdatePoll) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MsgUpdatePoll) GetOptions() []string {
	if m != nil {
		return m.Options
	}
	return nil
}

type MsgUpdatePollResponse struct {
}

func (m *MsgUpdatePollResponse) Reset()         { *m = MsgUpdatePollResponse{} }
func (m *MsgUpdatePollResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdatePollResponse) ProtoMessage()    {}
func (*MsgUpdatePollResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_779ae7fcafad750a, []int{3}
}
func (m *MsgUpdatePollResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdatePollResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdatePollResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdatePollResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdatePollResponse.Merge(m, src)
}
func (m *MsgUpdatePollResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdatePollResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdatePollResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdatePollResponse proto.InternalMessageInfo

type MsgDeletePoll struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgDeletePoll) Reset()         { *m = MsgDeletePoll{} }
func (m *MsgDeletePoll) String() string { return proto.CompactTextString(m) }
func (*MsgDeletePoll) ProtoMessage()    {}
func (*MsgDeletePoll) Descriptor() ([]byte, []int) {
	return fileDescriptor_779ae7fcafad750a, []int{4}
}
func (m *MsgDeletePoll) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeletePoll) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeletePoll.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeletePoll) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeletePoll.Merge(m, src)
}
func (m *MsgDeletePoll) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeletePoll) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeletePoll.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeletePoll proto.InternalMessageInfo

func (m *MsgDeletePoll) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDeletePoll) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type MsgDeletePollResponse struct {
}

func (m *MsgDeletePollResponse) Reset()         { *m = MsgDeletePollResponse{} }
func (m *MsgDeletePollResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeletePollResponse) ProtoMessage()    {}
func (*MsgDeletePollResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_779ae7fcafad750a, []int{5}
}
func (m *MsgDeletePollResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeletePollResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeletePollResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeletePollResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeletePollResponse.Merge(m, src)
}
func (m *MsgDeletePollResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeletePollResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeletePollResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeletePollResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreatePoll)(nil), "papillon6814.voter.voter.MsgCreatePoll")
	proto.RegisterType((*MsgCreatePollResponse)(nil), "papillon6814.voter.voter.MsgCreatePollResponse")
	proto.RegisterType((*MsgUpdatePoll)(nil), "papillon6814.voter.voter.MsgUpdatePoll")
	proto.RegisterType((*MsgUpdatePollResponse)(nil), "papillon6814.voter.voter.MsgUpdatePollResponse")
	proto.RegisterType((*MsgDeletePoll)(nil), "papillon6814.voter.voter.MsgDeletePoll")
	proto.RegisterType((*MsgDeletePollResponse)(nil), "papillon6814.voter.voter.MsgDeletePollResponse")
}

func init() { proto.RegisterFile("voter/tx.proto", fileDescriptor_779ae7fcafad750a) }

var fileDescriptor_779ae7fcafad750a = []byte{
	// 328 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0xcb, 0x2f, 0x49,
	0x2d, 0xd2, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x28, 0x48, 0x2c, 0xc8,
	0xcc, 0xc9, 0xc9, 0xcf, 0x33, 0xb3, 0x30, 0x34, 0xd1, 0x03, 0x4b, 0x42, 0x48, 0x29, 0x01, 0x88,
	0xca, 0x82, 0xfc, 0x9c, 0x1c, 0x88, 0x5a, 0xa5, 0x48, 0x2e, 0x5e, 0xdf, 0xe2, 0x74, 0xe7, 0xa2,
	0xd4, 0xc4, 0x92, 0xd4, 0x80, 0xfc, 0x9c, 0x1c, 0x21, 0x09, 0x2e, 0xf6, 0x64, 0x10, 0x2f, 0xbf,
	0x48, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xc6, 0x15, 0x12, 0xe1, 0x62, 0x2d, 0xc9, 0x2c,
	0xc9, 0x49, 0x95, 0x60, 0x02, 0x8b, 0x43, 0x38, 0x20, 0xf5, 0xf9, 0x05, 0x25, 0x99, 0xf9, 0x79,
	0xc5, 0x12, 0xcc, 0x0a, 0xcc, 0x20, 0xf5, 0x50, 0xae, 0x92, 0x3a, 0x97, 0x28, 0x8a, 0xd1, 0x41,
	0xa9, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x60, 0xd3, 0x59,
	0x82, 0x98, 0x32, 0x53, 0x94, 0x32, 0xc1, 0x6e, 0x08, 0x2d, 0x48, 0x21, 0xec, 0x06, 0x88, 0x56,
	0x26, 0x98, 0x56, 0x84, 0x9b, 0x98, 0x71, 0xb8, 0x89, 0x05, 0xd5, 0x4d, 0xe2, 0x60, 0x37, 0x21,
	0xac, 0x82, 0xb9, 0x49, 0xc9, 0x12, 0xec, 0x06, 0x97, 0xd4, 0x9c, 0x54, 0x52, 0xdd, 0x00, 0x35,
	0x13, 0xa1, 0x15, 0x66, 0xa6, 0xd1, 0x5e, 0x26, 0x2e, 0x66, 0xdf, 0xe2, 0x74, 0xa1, 0x34, 0x2e,
	0x2e, 0xa4, 0x00, 0x56, 0xd7, 0xc3, 0x15, 0x3d, 0x7a, 0x28, 0xc1, 0x25, 0xa5, 0x4f, 0xa4, 0x42,
	0x78, 0xb8, 0xa6, 0x71, 0x71, 0x21, 0x05, 0x22, 0x7e, 0x7b, 0x10, 0x0a, 0x09, 0xd8, 0x83, 0x19,
	0x56, 0x20, 0x7b, 0x90, 0x02, 0x0a, 0xbf, 0x3d, 0x08, 0x85, 0x04, 0xec, 0xc1, 0x0c, 0x3f, 0x27,
	0xd7, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63,
	0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x4e, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x47, 0x36, 0x54, 0x1f, 0x92, 0xbe, 0x2b, 0xa0, 0x74,
	0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0xa5, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x19, 0x97, 0xba, 0x05, 0x27, 0x03, 0x00, 0x00,
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
	// this line is used by starport scaffolding # proto/tx/rpc
	CreatePoll(ctx context.Context, in *MsgCreatePoll, opts ...grpc.CallOption) (*MsgCreatePollResponse, error)
	UpdatePoll(ctx context.Context, in *MsgUpdatePoll, opts ...grpc.CallOption) (*MsgUpdatePollResponse, error)
	DeletePoll(ctx context.Context, in *MsgDeletePoll, opts ...grpc.CallOption) (*MsgDeletePollResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreatePoll(ctx context.Context, in *MsgCreatePoll, opts ...grpc.CallOption) (*MsgCreatePollResponse, error) {
	out := new(MsgCreatePollResponse)
	err := c.cc.Invoke(ctx, "/papillon6814.voter.voter.Msg/CreatePoll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdatePoll(ctx context.Context, in *MsgUpdatePoll, opts ...grpc.CallOption) (*MsgUpdatePollResponse, error) {
	out := new(MsgUpdatePollResponse)
	err := c.cc.Invoke(ctx, "/papillon6814.voter.voter.Msg/UpdatePoll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeletePoll(ctx context.Context, in *MsgDeletePoll, opts ...grpc.CallOption) (*MsgDeletePollResponse, error) {
	out := new(MsgDeletePollResponse)
	err := c.cc.Invoke(ctx, "/papillon6814.voter.voter.Msg/DeletePoll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	CreatePoll(context.Context, *MsgCreatePoll) (*MsgCreatePollResponse, error)
	UpdatePoll(context.Context, *MsgUpdatePoll) (*MsgUpdatePollResponse, error)
	DeletePoll(context.Context, *MsgDeletePoll) (*MsgDeletePollResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreatePoll(ctx context.Context, req *MsgCreatePoll) (*MsgCreatePollResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePoll not implemented")
}
func (*UnimplementedMsgServer) UpdatePoll(ctx context.Context, req *MsgUpdatePoll) (*MsgUpdatePollResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePoll not implemented")
}
func (*UnimplementedMsgServer) DeletePoll(ctx context.Context, req *MsgDeletePoll) (*MsgDeletePollResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePoll not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreatePoll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreatePoll)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreatePoll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/papillon6814.voter.voter.Msg/CreatePoll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreatePoll(ctx, req.(*MsgCreatePoll))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdatePoll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdatePoll)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdatePoll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/papillon6814.voter.voter.Msg/UpdatePoll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdatePoll(ctx, req.(*MsgUpdatePoll))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeletePoll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeletePoll)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeletePoll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/papillon6814.voter.voter.Msg/DeletePoll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeletePoll(ctx, req.(*MsgDeletePoll))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "papillon6814.voter.voter.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePoll",
			Handler:    _Msg_CreatePoll_Handler,
		},
		{
			MethodName: "UpdatePoll",
			Handler:    _Msg_UpdatePoll_Handler,
		},
		{
			MethodName: "DeletePoll",
			Handler:    _Msg_DeletePoll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "voter/tx.proto",
}

func (m *MsgCreatePoll) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePoll) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePoll) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Options) > 0 {
		for iNdEx := len(m.Options) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Options[iNdEx])
			copy(dAtA[i:], m.Options[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Options[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreatePollResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreatePollResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreatePollResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdatePoll) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdatePoll) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdatePoll) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Options) > 0 {
		for iNdEx := len(m.Options) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Options[iNdEx])
			copy(dAtA[i:], m.Options[iNdEx])
			i = encodeVarintTx(dAtA, i, uint64(len(m.Options[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdatePollResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdatePollResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdatePollResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeletePoll) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeletePoll) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeletePoll) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeletePollResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeletePollResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeletePollResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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
func (m *MsgCreatePoll) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Options) > 0 {
		for _, s := range m.Options {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgCreatePollResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	return n
}

func (m *MsgUpdatePoll) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Options) > 0 {
		for _, s := range m.Options {
			l = len(s)
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgUpdatePollResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeletePoll) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	return n
}

func (m *MsgDeletePollResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreatePoll) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreatePoll: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePoll: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
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
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
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
			m.Options = append(m.Options, string(dAtA[iNdEx:postIndex]))
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
func (m *MsgCreatePollResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreatePollResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreatePollResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *MsgUpdatePoll) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdatePoll: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdatePoll: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
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
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Options", wireType)
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
			m.Options = append(m.Options, string(dAtA[iNdEx:postIndex]))
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
func (m *MsgUpdatePollResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUpdatePollResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdatePollResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgDeletePoll) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDeletePoll: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeletePoll: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *MsgDeletePollResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDeletePollResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeletePollResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
