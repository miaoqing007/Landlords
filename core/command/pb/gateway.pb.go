// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gateway.proto

package command

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type ClientPlayerMsgData struct {
	PlayerId             uint64   `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	ClientAddr           string   `protobuf:"bytes,3,opt,name=clientAddr,proto3" json:"clientAddr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientPlayerMsgData) Reset()         { *m = ClientPlayerMsgData{} }
func (m *ClientPlayerMsgData) String() string { return proto.CompactTextString(m) }
func (*ClientPlayerMsgData) ProtoMessage()    {}
func (*ClientPlayerMsgData) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{0}
}

func (m *ClientPlayerMsgData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientPlayerMsgData.Unmarshal(m, b)
}
func (m *ClientPlayerMsgData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientPlayerMsgData.Marshal(b, m, deterministic)
}
func (m *ClientPlayerMsgData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientPlayerMsgData.Merge(m, src)
}
func (m *ClientPlayerMsgData) XXX_Size() int {
	return xxx_messageInfo_ClientPlayerMsgData.Size(m)
}
func (m *ClientPlayerMsgData) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientPlayerMsgData.DiscardUnknown(m)
}

var xxx_messageInfo_ClientPlayerMsgData proto.InternalMessageInfo

func (m *ClientPlayerMsgData) GetPlayerId() uint64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *ClientPlayerMsgData) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ClientPlayerMsgData) GetClientAddr() string {
	if m != nil {
		return m.ClientAddr
	}
	return ""
}

type GatewayMsgToOnline struct {
	RemoteAddr           string   `protobuf:"bytes,1,opt,name=remoteAddr,proto3" json:"remoteAddr,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GatewayMsgToOnline) Reset()         { *m = GatewayMsgToOnline{} }
func (m *GatewayMsgToOnline) String() string { return proto.CompactTextString(m) }
func (*GatewayMsgToOnline) ProtoMessage()    {}
func (*GatewayMsgToOnline) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{1}
}

func (m *GatewayMsgToOnline) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GatewayMsgToOnline.Unmarshal(m, b)
}
func (m *GatewayMsgToOnline) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GatewayMsgToOnline.Marshal(b, m, deterministic)
}
func (m *GatewayMsgToOnline) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GatewayMsgToOnline.Merge(m, src)
}
func (m *GatewayMsgToOnline) XXX_Size() int {
	return xxx_messageInfo_GatewayMsgToOnline.Size(m)
}
func (m *GatewayMsgToOnline) XXX_DiscardUnknown() {
	xxx_messageInfo_GatewayMsgToOnline.DiscardUnknown(m)
}

var xxx_messageInfo_GatewayMsgToOnline proto.InternalMessageInfo

func (m *GatewayMsgToOnline) GetRemoteAddr() string {
	if m != nil {
		return m.RemoteAddr
	}
	return ""
}

func (m *GatewayMsgToOnline) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type OnlineMsgToGateway struct {
	RemoteAddr           string   `protobuf:"bytes,1,opt,name=remoteAddr,proto3" json:"remoteAddr,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OnlineMsgToGateway) Reset()         { *m = OnlineMsgToGateway{} }
func (m *OnlineMsgToGateway) String() string { return proto.CompactTextString(m) }
func (*OnlineMsgToGateway) ProtoMessage()    {}
func (*OnlineMsgToGateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{2}
}

func (m *OnlineMsgToGateway) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OnlineMsgToGateway.Unmarshal(m, b)
}
func (m *OnlineMsgToGateway) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OnlineMsgToGateway.Marshal(b, m, deterministic)
}
func (m *OnlineMsgToGateway) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OnlineMsgToGateway.Merge(m, src)
}
func (m *OnlineMsgToGateway) XXX_Size() int {
	return xxx_messageInfo_OnlineMsgToGateway.Size(m)
}
func (m *OnlineMsgToGateway) XXX_DiscardUnknown() {
	xxx_messageInfo_OnlineMsgToGateway.DiscardUnknown(m)
}

var xxx_messageInfo_OnlineMsgToGateway proto.InternalMessageInfo

func (m *OnlineMsgToGateway) GetRemoteAddr() string {
	if m != nil {
		return m.RemoteAddr
	}
	return ""
}

func (m *OnlineMsgToGateway) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type CSPlayerLogin_Gateway struct {
	PlayerId             uint64   `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CSPlayerLogin_Gateway) Reset()         { *m = CSPlayerLogin_Gateway{} }
func (m *CSPlayerLogin_Gateway) String() string { return proto.CompactTextString(m) }
func (*CSPlayerLogin_Gateway) ProtoMessage()    {}
func (*CSPlayerLogin_Gateway) Descriptor() ([]byte, []int) {
	return fileDescriptor_f1a937782ebbded5, []int{3}
}

func (m *CSPlayerLogin_Gateway) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CSPlayerLogin_Gateway.Unmarshal(m, b)
}
func (m *CSPlayerLogin_Gateway) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CSPlayerLogin_Gateway.Marshal(b, m, deterministic)
}
func (m *CSPlayerLogin_Gateway) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CSPlayerLogin_Gateway.Merge(m, src)
}
func (m *CSPlayerLogin_Gateway) XXX_Size() int {
	return xxx_messageInfo_CSPlayerLogin_Gateway.Size(m)
}
func (m *CSPlayerLogin_Gateway) XXX_DiscardUnknown() {
	xxx_messageInfo_CSPlayerLogin_Gateway.DiscardUnknown(m)
}

var xxx_messageInfo_CSPlayerLogin_Gateway proto.InternalMessageInfo

func (m *CSPlayerLogin_Gateway) GetPlayerId() uint64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func init() {
	proto.RegisterType((*ClientPlayerMsgData)(nil), "command.ClientPlayerMsgData")
	proto.RegisterType((*GatewayMsgToOnline)(nil), "command.GatewayMsgToOnline")
	proto.RegisterType((*OnlineMsgToGateway)(nil), "command.OnlineMsgToGateway")
	proto.RegisterType((*CSPlayerLogin_Gateway)(nil), "command.CSPlayerLogin_Gateway")
}

func init() { proto.RegisterFile("gateway.proto", fileDescriptor_f1a937782ebbded5) }

var fileDescriptor_f1a937782ebbded5 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x4f, 0x2c, 0x49,
	0x2d, 0x4f, 0xac, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xce, 0xcf, 0xcd, 0x4d,
	0xcc, 0x4b, 0x51, 0x4a, 0xe5, 0x12, 0x76, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x09, 0xc8, 0x49, 0xac,
	0x4c, 0x2d, 0xf2, 0x2d, 0x4e, 0x77, 0x49, 0x2c, 0x49, 0x14, 0x92, 0xe2, 0xe2, 0x28, 0x00, 0x0b,
	0x78, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0xb0, 0x04, 0xc1, 0xf9, 0x42, 0x42, 0x5c, 0x2c, 0x29,
	0x89, 0x25, 0x89, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x3c, 0x41, 0x60, 0xb6, 0x90, 0x1c, 0x17, 0x57,
	0x32, 0xd8, 0x18, 0xc7, 0x94, 0x94, 0x22, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x24, 0x11,
	0x25, 0x0f, 0x2e, 0x21, 0x77, 0x88, 0x03, 0x7c, 0x8b, 0xd3, 0x43, 0xf2, 0xfd, 0xf3, 0x72, 0x32,
	0xf3, 0x52, 0x41, 0xba, 0x8a, 0x52, 0x73, 0xf3, 0x4b, 0x52, 0xc1, 0xba, 0x18, 0x21, 0xba, 0x10,
	0x22, 0xd8, 0x6c, 0x02, 0x99, 0x04, 0xd1, 0x0d, 0x36, 0x08, 0x6a, 0x28, 0x59, 0x26, 0x19, 0x73,
	0x89, 0x3a, 0x07, 0x43, 0xbc, 0xed, 0x93, 0x9f, 0x9e, 0x99, 0x17, 0x0f, 0x33, 0x0c, 0x8f, 0xe7,
	0x93, 0xd8, 0xc0, 0xe1, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x72, 0x3c, 0x98, 0x4d, 0x50,
	0x01, 0x00, 0x00,
}