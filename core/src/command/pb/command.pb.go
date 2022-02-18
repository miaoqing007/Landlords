// Code generated by protoc-gen-go. DO NOT EDIT.
// source: command.proto

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

type Command int32

const (
	Command_dummy Command = 0
	//
	Command_CSStartGame Command = 1
)

var Command_name = map[int32]string{
	0: "dummy",
	1: "CSStartGame",
}

var Command_value = map[string]int32{
	"dummy":       0,
	"CSStartGame": 1,
}

func (x Command) String() string {
	return proto.EnumName(Command_name, int32(x))
}

func (Command) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_213c0bb044472049, []int{0}
}

func init() {
	proto.RegisterEnum("command.Command", Command_name, Command_value)
}

func init() { proto.RegisterFile("command.proto", fileDescriptor_213c0bb044472049) }

var fileDescriptor_213c0bb044472049 = []byte{
	// 83 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0xcf, 0xcd,
	0x4d, 0xcc, 0x4b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0xb5, 0x54, 0xb9,
	0xd8, 0x9d, 0x21, 0x4c, 0x21, 0x4e, 0x2e, 0xd6, 0x94, 0xd2, 0xdc, 0xdc, 0x4a, 0x01, 0x06, 0x21,
	0x7e, 0x2e, 0x6e, 0xe7, 0xe0, 0xe0, 0x92, 0xc4, 0xa2, 0x12, 0xf7, 0xc4, 0xdc, 0x54, 0x01, 0xc6,
	0x24, 0x36, 0xb0, 0x36, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x73, 0x0e, 0xc1, 0x9c, 0x47,
	0x00, 0x00, 0x00,
}
