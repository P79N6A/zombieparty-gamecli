// Code generated by protoc-gen-go. DO NOT EDIT.
// source: SyncRequestProto.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type SyncRequest struct {
	RoundId string       `protobuf:"bytes,1,opt,name=roundId" json:"roundId,omitempty"`
	Data    []*InputData `protobuf:"bytes,2,rep,name=data" json:"data,omitempty"`
}

func (m *SyncRequest) Reset()                    { *m = SyncRequest{} }
func (m *SyncRequest) String() string            { return proto.CompactTextString(m) }
func (*SyncRequest) ProtoMessage()               {}
func (*SyncRequest) Descriptor() ([]byte, []int) { return fileDescriptor7, []int{0} }

func (m *SyncRequest) GetRoundId() string {
	if m != nil {
		return m.RoundId
	}
	return ""
}

func (m *SyncRequest) GetData() []*InputData {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*SyncRequest)(nil), "SyncRequest")
}

func init() { proto.RegisterFile("SyncRequestProto.proto", fileDescriptor7) }

var fileDescriptor7 = []byte{
	// 156 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x0b, 0xae, 0xcc, 0x4b,
	0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x09, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2b, 0x00, 0x91,
	0x52, 0x22, 0x9e, 0x79, 0x05, 0xa5, 0x25, 0x2e, 0x89, 0x25, 0x89, 0x48, 0xa2, 0x4a, 0xee, 0x5c,
	0xdc, 0x48, 0xea, 0x85, 0x24, 0xb8, 0xd8, 0x8b, 0xf2, 0x4b, 0xf3, 0x52, 0x3c, 0x53, 0x24, 0x18,
	0x15, 0x18, 0x35, 0x38, 0x83, 0x60, 0x5c, 0x21, 0x39, 0x2e, 0x96, 0x94, 0xc4, 0x92, 0x44, 0x09,
	0x26, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x2e, 0x3d, 0xb8, 0x69, 0x41, 0x60, 0x71, 0x27, 0x13, 0x2e,
	0x8d, 0xe4, 0xfc, 0x5c, 0xbd, 0xcc, 0xdc, 0xdc, 0xfc, 0xdc, 0x7c, 0xbd, 0xaa, 0xfc, 0xdc, 0xa4,
	0xcc, 0xd4, 0x82, 0xc4, 0xa2, 0x92, 0x4a, 0xbd, 0xf4, 0xc4, 0xdc, 0xd4, 0xe2, 0xd4, 0xa2, 0xb2,
	0xd4, 0x22, 0x88, 0x9d, 0xc5, 0x01, 0x8c, 0x51, 0x6c, 0x10, 0x56, 0x12, 0x84, 0x36, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0x9b, 0x10, 0x15, 0xea, 0xb5, 0x00, 0x00, 0x00,
}