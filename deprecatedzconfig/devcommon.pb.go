// Code generated by protoc-gen-go. DO NOT EDIT.
// source: devcommon.proto

package deprecatedzconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type UUIDandVersion struct {
	Uuid    string `protobuf:"bytes,1,opt,name=uuid" json:"uuid,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
}

func (m *UUIDandVersion) Reset()                    { *m = UUIDandVersion{} }
func (m *UUIDandVersion) String() string            { return proto.CompactTextString(m) }
func (*UUIDandVersion) ProtoMessage()               {}
func (*UUIDandVersion) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

func (m *UUIDandVersion) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

func (m *UUIDandVersion) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*UUIDandVersion)(nil), "UUIDandVersion")
}

func init() { proto.RegisterFile("devcommon.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8d, 0x3f, 0x0b, 0xc2, 0x30,
	0x10, 0x47, 0xa9, 0x88, 0x62, 0x06, 0x85, 0x4e, 0xdd, 0x14, 0x17, 0x9d, 0xd2, 0xc1, 0xdd, 0x41,
	0x5c, 0x74, 0x14, 0xea, 0xe0, 0x96, 0xde, 0x9d, 0xf5, 0xc0, 0xe4, 0x4a, 0xda, 0x64, 0xe8, 0xa7,
	0x17, 0x12, 0xdc, 0x7e, 0x7f, 0x1e, 0x3c, 0xb5, 0x41, 0x8a, 0x20, 0xd6, 0x8a, 0xd3, 0xbd, 0x97,
	0x51, 0xf6, 0x67, 0xb5, 0x6e, 0x9a, 0xdb, 0xd5, 0x38, 0x7c, 0x92, 0x1f, 0x58, 0x5c, 0x59, 0xaa,
	0x79, 0x08, 0x8c, 0x55, 0xb1, 0x2b, 0x8e, 0xab, 0x47, 0xca, 0x65, 0xa5, 0x96, 0x31, 0xdf, 0xd5,
	0x2c, 0xcd, 0xff, 0x7a, 0xb9, 0xab, 0x2d, 0x88, 0xd5, 0x13, 0x21, 0xa1, 0xd1, 0xf0, 0x95, 0x80,
	0x3a, 0x0c, 0xe4, 0x23, 0x03, 0x65, 0xc5, 0xeb, 0xd0, 0xf1, 0xf8, 0x09, 0xad, 0x06, 0xb1, 0x75,
	0xe6, 0x6a, 0xd3, 0x73, 0x8d, 0xd4, 0x7b, 0x02, 0x33, 0x12, 0x4e, 0x20, 0xee, 0xcd, 0x5d, 0xbb,
	0x48, 0xfc, 0xe9, 0x17, 0x00, 0x00, 0xff, 0xff, 0x0b, 0xe1, 0xdc, 0x64, 0xa5, 0x00, 0x00, 0x00,
}
