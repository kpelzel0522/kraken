// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testmutate.proto

package proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TestMutateConfig struct {
	ScalingUrl           string   `protobuf:"bytes,1,opt,name=scaling_url,json=scalingUrl,proto3" json:"scaling_url,omitempty"`
	TempUrl              string   `protobuf:"bytes,2,opt,name=temp_url,json=tempUrl,proto3" json:"temp_url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TestMutateConfig) Reset()         { *m = TestMutateConfig{} }
func (m *TestMutateConfig) String() string { return proto.CompactTextString(m) }
func (*TestMutateConfig) ProtoMessage()    {}
func (*TestMutateConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_testmutate_3ee645b6e9047898, []int{0}
}
func (m *TestMutateConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestMutateConfig.Unmarshal(m, b)
}
func (m *TestMutateConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestMutateConfig.Marshal(b, m, deterministic)
}
func (dst *TestMutateConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestMutateConfig.Merge(dst, src)
}
func (m *TestMutateConfig) XXX_Size() int {
	return xxx_messageInfo_TestMutateConfig.Size(m)
}
func (m *TestMutateConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_TestMutateConfig.DiscardUnknown(m)
}

var xxx_messageInfo_TestMutateConfig proto.InternalMessageInfo

func (m *TestMutateConfig) GetScalingUrl() string {
	if m != nil {
		return m.ScalingUrl
	}
	return ""
}

func (m *TestMutateConfig) GetTempUrl() string {
	if m != nil {
		return m.TempUrl
	}
	return ""
}

func init() {
	proto.RegisterType((*TestMutateConfig)(nil), "proto.TestMutateConfig")
}

func init() { proto.RegisterFile("testmutate.proto", fileDescriptor_testmutate_3ee645b6e9047898) }

var fileDescriptor_testmutate_3ee645b6e9047898 = []byte{
	// 114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x49, 0x2d, 0x2e,
	0xc9, 0x2d, 0x2d, 0x49, 0x2c, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53,
	0x4a, 0x7e, 0x5c, 0x02, 0x21, 0xa9, 0xc5, 0x25, 0xbe, 0x60, 0x29, 0xe7, 0xfc, 0xbc, 0xb4, 0xcc,
	0x74, 0x21, 0x79, 0x2e, 0xee, 0xe2, 0xe4, 0xc4, 0x9c, 0xcc, 0xbc, 0xf4, 0xf8, 0xd2, 0xa2, 0x1c,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x2e, 0xa8, 0x50, 0x68, 0x51, 0x8e, 0x90, 0x24, 0x17,
	0x47, 0x49, 0x6a, 0x6e, 0x01, 0x58, 0x96, 0x09, 0x2c, 0xcb, 0x0e, 0xe2, 0x87, 0x16, 0xe5, 0x24,
	0xb1, 0x81, 0x8d, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x49, 0x7e, 0x9b, 0x71, 0x00,
	0x00, 0x00,
}