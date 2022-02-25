// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc.proto

package command

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("grpc.proto", fileDescriptor_bedfbfc9b54e5600) }

var fileDescriptor_bedfbfc9b54e5600 = []byte{
	// 183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x2f, 0x2a, 0x48,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xce, 0xcf, 0xcd, 0x4d, 0xcc, 0x4b, 0x91,
	0xe2, 0x4d, 0x4f, 0x2c, 0x49, 0x2d, 0x4f, 0xac, 0x84, 0x88, 0x4b, 0xf1, 0xe4, 0xe7, 0xe5, 0x64,
	0xe6, 0xa5, 0x42, 0x79, 0x9c, 0x05, 0x65, 0x05, 0x10, 0xa6, 0x51, 0x18, 0x17, 0xa7, 0x3f, 0x58,
	0x2a, 0xa0, 0xac, 0x40, 0xc8, 0x93, 0x8b, 0x1f, 0xce, 0x09, 0x2e, 0x29, 0x4a, 0x4d, 0xcc, 0x15,
	0x12, 0xd7, 0x83, 0x9a, 0xa8, 0x07, 0x91, 0x31, 0x0a, 0x28, 0x2b, 0xf0, 0xcc, 0x4b, 0xcb, 0x97,
	0x42, 0x48, 0x04, 0x94, 0x15, 0x18, 0x41, 0x24, 0x41, 0x12, 0x4a, 0x0c, 0x1a, 0x8c, 0x06, 0x8c,
	0x46, 0x19, 0x5c, 0xbc, 0xee, 0x10, 0x17, 0x40, 0xa4, 0x84, 0xc2, 0xb9, 0x84, 0x51, 0x04, 0xa0,
	0xe6, 0xcb, 0xc0, 0x8d, 0x71, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x09, 0xc8, 0x49, 0xac, 0x4c, 0x2d,
	0xf2, 0x2d, 0x4e, 0x77, 0x49, 0x2c, 0x49, 0x94, 0x42, 0xc8, 0x06, 0xa7, 0x16, 0x95, 0xa5, 0x16,
	0xa1, 0xc8, 0x42, 0x6c, 0x4a, 0x62, 0x03, 0x7b, 0xc4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x6b,
	0x69, 0xa1, 0xb0, 0x07, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OnlinePvpClient is the client API for OnlinePvp service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OnlinePvpClient interface {
	OnlinePvpStream(ctx context.Context, opts ...grpc.CallOption) (OnlinePvp_OnlinePvpStreamClient, error)
}

type onlinePvpClient struct {
	cc *grpc.ClientConn
}

func NewOnlinePvpClient(cc *grpc.ClientConn) OnlinePvpClient {
	return &onlinePvpClient{cc}
}

func (c *onlinePvpClient) OnlinePvpStream(ctx context.Context, opts ...grpc.CallOption) (OnlinePvp_OnlinePvpStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_OnlinePvp_serviceDesc.Streams[0], "/command.OnlinePvp/OnlinePvpStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &onlinePvpOnlinePvpStreamClient{stream}
	return x, nil
}

type OnlinePvp_OnlinePvpStreamClient interface {
	Send(*Online2PvpInfo) error
	Recv() (*Pvp2OnlineInfo, error)
	grpc.ClientStream
}

type onlinePvpOnlinePvpStreamClient struct {
	grpc.ClientStream
}

func (x *onlinePvpOnlinePvpStreamClient) Send(m *Online2PvpInfo) error {
	return x.ClientStream.SendMsg(m)
}

func (x *onlinePvpOnlinePvpStreamClient) Recv() (*Pvp2OnlineInfo, error) {
	m := new(Pvp2OnlineInfo)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OnlinePvpServer is the server API for OnlinePvp service.
type OnlinePvpServer interface {
	OnlinePvpStream(OnlinePvp_OnlinePvpStreamServer) error
}

func RegisterOnlinePvpServer(s *grpc.Server, srv OnlinePvpServer) {
	s.RegisterService(&_OnlinePvp_serviceDesc, srv)
}

func _OnlinePvp_OnlinePvpStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OnlinePvpServer).OnlinePvpStream(&onlinePvpOnlinePvpStreamServer{stream})
}

type OnlinePvp_OnlinePvpStreamServer interface {
	Send(*Pvp2OnlineInfo) error
	Recv() (*Online2PvpInfo, error)
	grpc.ServerStream
}

type onlinePvpOnlinePvpStreamServer struct {
	grpc.ServerStream
}

func (x *onlinePvpOnlinePvpStreamServer) Send(m *Pvp2OnlineInfo) error {
	return x.ServerStream.SendMsg(m)
}

func (x *onlinePvpOnlinePvpStreamServer) Recv() (*Online2PvpInfo, error) {
	m := new(Online2PvpInfo)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _OnlinePvp_serviceDesc = grpc.ServiceDesc{
	ServiceName: "command.OnlinePvp",
	HandlerType: (*OnlinePvpServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "OnlinePvpStream",
			Handler:       _OnlinePvp_OnlinePvpStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc.proto",
}

// GatewayOnlineClient is the client API for GatewayOnline service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GatewayOnlineClient interface {
	GatewayOnlineStream(ctx context.Context, opts ...grpc.CallOption) (GatewayOnline_GatewayOnlineStreamClient, error)
}

type gatewayOnlineClient struct {
	cc *grpc.ClientConn
}

func NewGatewayOnlineClient(cc *grpc.ClientConn) GatewayOnlineClient {
	return &gatewayOnlineClient{cc}
}

func (c *gatewayOnlineClient) GatewayOnlineStream(ctx context.Context, opts ...grpc.CallOption) (GatewayOnline_GatewayOnlineStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_GatewayOnline_serviceDesc.Streams[0], "/command.GatewayOnline/GatewayOnlineStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &gatewayOnlineGatewayOnlineStreamClient{stream}
	return x, nil
}

type GatewayOnline_GatewayOnlineStreamClient interface {
	Send(*ClientPlayerMsgData) error
	Recv() (*ServerPlayerMsgData, error)
	grpc.ClientStream
}

type gatewayOnlineGatewayOnlineStreamClient struct {
	grpc.ClientStream
}

func (x *gatewayOnlineGatewayOnlineStreamClient) Send(m *ClientPlayerMsgData) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gatewayOnlineGatewayOnlineStreamClient) Recv() (*ServerPlayerMsgData, error) {
	m := new(ServerPlayerMsgData)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// GatewayOnlineServer is the server API for GatewayOnline service.
type GatewayOnlineServer interface {
	GatewayOnlineStream(GatewayOnline_GatewayOnlineStreamServer) error
}

func RegisterGatewayOnlineServer(s *grpc.Server, srv GatewayOnlineServer) {
	s.RegisterService(&_GatewayOnline_serviceDesc, srv)
}

func _GatewayOnline_GatewayOnlineStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GatewayOnlineServer).GatewayOnlineStream(&gatewayOnlineGatewayOnlineStreamServer{stream})
}

type GatewayOnline_GatewayOnlineStreamServer interface {
	Send(*ServerPlayerMsgData) error
	Recv() (*ClientPlayerMsgData, error)
	grpc.ServerStream
}

type gatewayOnlineGatewayOnlineStreamServer struct {
	grpc.ServerStream
}

func (x *gatewayOnlineGatewayOnlineStreamServer) Send(m *ServerPlayerMsgData) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gatewayOnlineGatewayOnlineStreamServer) Recv() (*ClientPlayerMsgData, error) {
	m := new(ClientPlayerMsgData)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _GatewayOnline_serviceDesc = grpc.ServiceDesc{
	ServiceName: "command.GatewayOnline",
	HandlerType: (*GatewayOnlineServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GatewayOnlineStream",
			Handler:       _GatewayOnline_GatewayOnlineStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc.proto",
}
