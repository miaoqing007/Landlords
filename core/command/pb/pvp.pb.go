// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pvp.proto

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

//
type Online2PvpInfo struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	PlayerId             uint64   `protobuf:"varint,2,opt,name=playerId,proto3" json:"playerId,omitempty"`
	Addr                 string   `protobuf:"bytes,3,opt,name=addr,proto3" json:"addr,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Online2PvpInfo) Reset()         { *m = Online2PvpInfo{} }
func (m *Online2PvpInfo) String() string { return proto.CompactTextString(m) }
func (*Online2PvpInfo) ProtoMessage()    {}
func (*Online2PvpInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e815925a12c5bb2d, []int{0}
}

func (m *Online2PvpInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Online2PvpInfo.Unmarshal(m, b)
}
func (m *Online2PvpInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Online2PvpInfo.Marshal(b, m, deterministic)
}
func (m *Online2PvpInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Online2PvpInfo.Merge(m, src)
}
func (m *Online2PvpInfo) XXX_Size() int {
	return xxx_messageInfo_Online2PvpInfo.Size(m)
}
func (m *Online2PvpInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Online2PvpInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Online2PvpInfo proto.InternalMessageInfo

func (m *Online2PvpInfo) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Online2PvpInfo) GetPlayerId() uint64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *Online2PvpInfo) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

//
type Pvp2OnlineInfo struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	PlayerId             uint64   `protobuf:"varint,2,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pvp2OnlineInfo) Reset()         { *m = Pvp2OnlineInfo{} }
func (m *Pvp2OnlineInfo) String() string { return proto.CompactTextString(m) }
func (*Pvp2OnlineInfo) ProtoMessage()    {}
func (*Pvp2OnlineInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_e815925a12c5bb2d, []int{1}
}

func (m *Pvp2OnlineInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pvp2OnlineInfo.Unmarshal(m, b)
}
func (m *Pvp2OnlineInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pvp2OnlineInfo.Marshal(b, m, deterministic)
}
func (m *Pvp2OnlineInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pvp2OnlineInfo.Merge(m, src)
}
func (m *Pvp2OnlineInfo) XXX_Size() int {
	return xxx_messageInfo_Pvp2OnlineInfo.Size(m)
}
func (m *Pvp2OnlineInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_Pvp2OnlineInfo.DiscardUnknown(m)
}

var xxx_messageInfo_Pvp2OnlineInfo proto.InternalMessageInfo

func (m *Pvp2OnlineInfo) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Pvp2OnlineInfo) GetPlayerId() uint64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

type CSJoinPvpPool_Pvp struct {
	PlayerId             uint64   `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CSJoinPvpPool_Pvp) Reset()         { *m = CSJoinPvpPool_Pvp{} }
func (m *CSJoinPvpPool_Pvp) String() string { return proto.CompactTextString(m) }
func (*CSJoinPvpPool_Pvp) ProtoMessage()    {}
func (*CSJoinPvpPool_Pvp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e815925a12c5bb2d, []int{2}
}

func (m *CSJoinPvpPool_Pvp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CSJoinPvpPool_Pvp.Unmarshal(m, b)
}
func (m *CSJoinPvpPool_Pvp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CSJoinPvpPool_Pvp.Marshal(b, m, deterministic)
}
func (m *CSJoinPvpPool_Pvp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CSJoinPvpPool_Pvp.Merge(m, src)
}
func (m *CSJoinPvpPool_Pvp) XXX_Size() int {
	return xxx_messageInfo_CSJoinPvpPool_Pvp.Size(m)
}
func (m *CSJoinPvpPool_Pvp) XXX_DiscardUnknown() {
	xxx_messageInfo_CSJoinPvpPool_Pvp.DiscardUnknown(m)
}

var xxx_messageInfo_CSJoinPvpPool_Pvp proto.InternalMessageInfo

func (m *CSJoinPvpPool_Pvp) GetPlayerId() uint64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

type CSExitPvpPool_Pvp struct {
	PlayerId             uint64   `protobuf:"varint,1,opt,name=playerId,proto3" json:"playerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CSExitPvpPool_Pvp) Reset()         { *m = CSExitPvpPool_Pvp{} }
func (m *CSExitPvpPool_Pvp) String() string { return proto.CompactTextString(m) }
func (*CSExitPvpPool_Pvp) ProtoMessage()    {}
func (*CSExitPvpPool_Pvp) Descriptor() ([]byte, []int) {
	return fileDescriptor_e815925a12c5bb2d, []int{3}
}

func (m *CSExitPvpPool_Pvp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CSExitPvpPool_Pvp.Unmarshal(m, b)
}
func (m *CSExitPvpPool_Pvp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CSExitPvpPool_Pvp.Marshal(b, m, deterministic)
}
func (m *CSExitPvpPool_Pvp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CSExitPvpPool_Pvp.Merge(m, src)
}
func (m *CSExitPvpPool_Pvp) XXX_Size() int {
	return xxx_messageInfo_CSExitPvpPool_Pvp.Size(m)
}
func (m *CSExitPvpPool_Pvp) XXX_DiscardUnknown() {
	xxx_messageInfo_CSExitPvpPool_Pvp.DiscardUnknown(m)
}

var xxx_messageInfo_CSExitPvpPool_Pvp proto.InternalMessageInfo

func (m *CSExitPvpPool_Pvp) GetPlayerId() uint64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func init() {
	proto.RegisterType((*Online2PvpInfo)(nil), "command.Online2PvpInfo")
	proto.RegisterType((*Pvp2OnlineInfo)(nil), "command.Pvp2OnlineInfo")
	proto.RegisterType((*CSJoinPvpPool_Pvp)(nil), "command.CSJoinPvpPool_Pvp")
	proto.RegisterType((*CSExitPvpPool_Pvp)(nil), "command.CSExitPvpPool_Pvp")
}

func init() { proto.RegisterFile("pvp.proto", fileDescriptor_e815925a12c5bb2d) }

var fileDescriptor_e815925a12c5bb2d = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x28, 0x2b, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xce, 0xcf, 0xcd, 0x4d, 0xcc, 0x4b, 0x51, 0x0a,
	0xe1, 0xe2, 0xf3, 0xcf, 0xcb, 0xc9, 0xcc, 0x4b, 0x35, 0x0a, 0x28, 0x2b, 0xf0, 0xcc, 0x4b, 0xcb,
	0x17, 0x12, 0xe2, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x09, 0x02,
	0xb3, 0x85, 0xa4, 0xb8, 0x38, 0x0a, 0x72, 0x12, 0x2b, 0x53, 0x8b, 0x3c, 0x53, 0x24, 0x98, 0x14,
	0x18, 0x35, 0x58, 0x82, 0xe0, 0x7c, 0x90, 0xfa, 0xc4, 0x94, 0x94, 0x22, 0x09, 0x66, 0x05, 0x46,
	0x0d, 0xce, 0x20, 0x30, 0x5b, 0xc9, 0x81, 0x8b, 0x2f, 0xa0, 0xac, 0xc0, 0x08, 0x62, 0x32, 0x39,
	0xa6, 0x2a, 0xe9, 0x73, 0x09, 0x3a, 0x07, 0x7b, 0xe5, 0x67, 0xe6, 0x05, 0x94, 0x15, 0x04, 0xe4,
	0xe7, 0xe7, 0xc4, 0x07, 0x94, 0x15, 0xa0, 0x68, 0x60, 0xc4, 0xa6, 0xc1, 0xb5, 0x22, 0xb3, 0x84,
	0x48, 0x0d, 0x49, 0x6c, 0xe0, 0x90, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x12, 0x0a, 0x7c,
	0xea, 0x16, 0x01, 0x00, 0x00,
}