// Code generated by protoc-gen-go. DO NOT EDIT.
// source: FrameProto.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Frame struct {
	No   int64        `protobuf:"varint,1,opt,name=no" json:"no,omitempty"`
	Data []*InputData `protobuf:"bytes,2,rep,name=data" json:"data,omitempty"`
}

func (m *Frame) Reset()                    { *m = Frame{} }
func (m *Frame) String() string            { return proto.CompactTextString(m) }
func (*Frame) ProtoMessage()               {}
func (*Frame) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Frame) GetNo() int64 {
	if m != nil {
		return m.No
	}
	return 0
}

func (m *Frame) GetData() []*InputData {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*Frame)(nil), "Frame")
}

func init() { proto.RegisterFile("FrameProto.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x70, 0x2b, 0x4a, 0xcc,
	0x4d, 0x0d, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2b, 0x00, 0x91, 0x52, 0x22, 0x9e, 0x79, 0x05, 0xa5,
	0x25, 0x2e, 0x89, 0x25, 0x89, 0x48, 0xa2, 0x4a, 0xe6, 0x5c, 0xac, 0x60, 0x95, 0x42, 0x7c, 0x5c,
	0x4c, 0x79, 0xf9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xcc, 0x41, 0x4c, 0x79, 0xf9, 0x42, 0x72, 0x5c,
	0x2c, 0x29, 0x89, 0x25, 0x89, 0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x5c, 0x7a, 0x70, 0xdd,
	0x41, 0x60, 0x71, 0x27, 0x13, 0x2e, 0x8d, 0xe4, 0xfc, 0x5c, 0xbd, 0xcc, 0xdc, 0xdc, 0xfc, 0xdc,
	0x7c, 0xbd, 0xaa, 0xfc, 0xdc, 0xa4, 0xcc, 0xd4, 0x82, 0xc4, 0xa2, 0x92, 0x4a, 0xbd, 0xf4, 0xc4,
	0xdc, 0xd4, 0xe2, 0xd4, 0xa2, 0xb2, 0xd4, 0x22, 0x88, 0x1d, 0xc5, 0x01, 0x8c, 0x51, 0x6c, 0x10,
	0x56, 0x12, 0x84, 0x36, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa6, 0xe7, 0x59, 0x1e, 0x9f, 0x00,
	0x00, 0x00,
}