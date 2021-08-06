// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: myblockchain2/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("myblockchain2/query.proto", fileDescriptor_11e5baba9049d7c9) }

var fileDescriptor_11e5baba9049d7c9 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xad, 0x4c, 0xca,
	0xc9, 0x4f, 0xce, 0x4e, 0xce, 0x48, 0xcc, 0xcc, 0x33, 0xd2, 0x2f, 0x2c, 0x4d, 0x2d, 0xaa, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xd2, 0x28, 0x48, 0x2c, 0xc8, 0xcc, 0xc9, 0xc9, 0xcf, 0x33,
	0xb3, 0x30, 0x34, 0xd1, 0x43, 0x51, 0x87, 0xca, 0x93, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49,
	0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b,
	0x86, 0x98, 0x23, 0xa5, 0x95, 0x9c, 0x5f, 0x9c, 0x9b, 0x5f, 0xac, 0x9f, 0x94, 0x58, 0x9c, 0x0a,
	0xb1, 0x40, 0xbf, 0xcc, 0x30, 0x29, 0xb5, 0x24, 0xd1, 0x50, 0xbf, 0x20, 0x31, 0x3d, 0x33, 0x0f,
	0xac, 0x18, 0xa2, 0xd6, 0x88, 0x9d, 0x8b, 0x35, 0x10, 0xa4, 0xc2, 0x29, 0xf4, 0xc4, 0x23, 0x39,
	0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63,
	0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xac, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92,
	0xf3, 0x73, 0xf5, 0x91, 0x5d, 0xa8, 0x8f, 0xea, 0x93, 0x0a, 0x34, 0x7e, 0x49, 0x65, 0x41, 0x6a,
	0x71, 0x12, 0x1b, 0xd8, 0x1a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x16, 0x5e, 0x9f, 0x61,
	0xf7, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "papillon6814.myblockchain2.myblockchain2.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "myblockchain2/query.proto",
}
