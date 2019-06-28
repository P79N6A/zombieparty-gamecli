// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ResponseProto.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// use to response for AuRequestProto and SyncRequestProto
type Response struct {
	// 1-> success, 0->failure
	Success int32  `protobuf:"varint,1,opt,name=success" json:"success,omitempty"`
	Errcode int32  `protobuf:"varint,2,opt,name=errcode" json:"errcode,omitempty"`
	Errmsg  string `protobuf:"bytes,3,opt,name=errmsg" json:"errmsg,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor6, []int{0} }

func (m *Response) GetSuccess() int32 {
	if m != nil {
		return m.Success
	}
	return 0
}

func (m *Response) GetErrcode() int32 {
	if m != nil {
		return m.Errcode
	}
	return 0
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func init() {
	proto.RegisterType((*Response)(nil), "Response")
}

func init() { proto.RegisterFile("ResponseProto.proto", fileDescriptor6) }

var fileDescriptor6 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x0e, 0x4a, 0x2d, 0x2e,
	0xc8, 0xcf, 0x2b, 0x4e, 0x0d, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2b, 0x00, 0x91, 0x4a, 0x61, 0x5c,
	0x1c, 0x30, 0x61, 0x21, 0x09, 0x2e, 0xf6, 0xe2, 0xd2, 0xe4, 0xe4, 0xd4, 0xe2, 0x62, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xd6, 0x20, 0x18, 0x17, 0x24, 0x93, 0x5a, 0x54, 0x94, 0x9c, 0x9f, 0x92, 0x2a,
	0xc1, 0x04, 0x91, 0x81, 0x72, 0x85, 0xc4, 0xb8, 0xd8, 0x52, 0x8b, 0x8a, 0x72, 0x8b, 0xd3, 0x25,
	0x98, 0x15, 0x18, 0x35, 0x38, 0x83, 0xa0, 0x3c, 0x27, 0x13, 0x2e, 0x8d, 0xe4, 0xfc, 0x5c, 0xbd,
	0xcc, 0xdc, 0xdc, 0xfc, 0xdc, 0x7c, 0xbd, 0xaa, 0xfc, 0xdc, 0xa4, 0xcc, 0xd4, 0x82, 0xc4, 0xa2,
	0x92, 0x4a, 0xbd, 0xf4, 0xc4, 0xdc, 0xd4, 0xe2, 0xd4, 0xa2, 0xb2, 0xd4, 0x22, 0x88, 0x13, 0x8a,
	0x03, 0x18, 0xa3, 0xd8, 0x20, 0xac, 0x24, 0x08, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xa3,
	0x03, 0x16, 0xed, 0xab, 0x00, 0x00, 0x00,
}