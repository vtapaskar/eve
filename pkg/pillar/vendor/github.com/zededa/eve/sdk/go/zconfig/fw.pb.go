// Code generated by protoc-gen-go. DO NOT EDIT.
// source: fw.proto

package zconfig

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

type ACEMatch struct {
	// FIXME: We should convert this to enum
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ACEMatch) Reset()         { *m = ACEMatch{} }
func (m *ACEMatch) String() string { return proto.CompactTextString(m) }
func (*ACEMatch) ProtoMessage()    {}
func (*ACEMatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_505e7efac08d3ba9, []int{0}
}

func (m *ACEMatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ACEMatch.Unmarshal(m, b)
}
func (m *ACEMatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ACEMatch.Marshal(b, m, deterministic)
}
func (m *ACEMatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ACEMatch.Merge(m, src)
}
func (m *ACEMatch) XXX_Size() int {
	return xxx_messageInfo_ACEMatch.Size(m)
}
func (m *ACEMatch) XXX_DiscardUnknown() {
	xxx_messageInfo_ACEMatch.DiscardUnknown(m)
}

var xxx_messageInfo_ACEMatch proto.InternalMessageInfo

func (m *ACEMatch) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *ACEMatch) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ACEAction struct {
	Drop bool `protobuf:"varint,1,opt,name=drop,proto3" json:"drop,omitempty"`
	// limit action, and its associated parameter
	Limit      bool   `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Limitrate  uint32 `protobuf:"varint,3,opt,name=limitrate,proto3" json:"limitrate,omitempty"`
	Limitunit  string `protobuf:"bytes,4,opt,name=limitunit,proto3" json:"limitunit,omitempty"`
	Limitburst uint32 `protobuf:"varint,5,opt,name=limitburst,proto3" json:"limitburst,omitempty"`
	// port map action, and its assoicated paramtere
	Portmap              bool     `protobuf:"varint,6,opt,name=portmap,proto3" json:"portmap,omitempty"`
	AppPort              uint32   `protobuf:"varint,7,opt,name=appPort,proto3" json:"appPort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ACEAction) Reset()         { *m = ACEAction{} }
func (m *ACEAction) String() string { return proto.CompactTextString(m) }
func (*ACEAction) ProtoMessage()    {}
func (*ACEAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_505e7efac08d3ba9, []int{1}
}

func (m *ACEAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ACEAction.Unmarshal(m, b)
}
func (m *ACEAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ACEAction.Marshal(b, m, deterministic)
}
func (m *ACEAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ACEAction.Merge(m, src)
}
func (m *ACEAction) XXX_Size() int {
	return xxx_messageInfo_ACEAction.Size(m)
}
func (m *ACEAction) XXX_DiscardUnknown() {
	xxx_messageInfo_ACEAction.DiscardUnknown(m)
}

var xxx_messageInfo_ACEAction proto.InternalMessageInfo

func (m *ACEAction) GetDrop() bool {
	if m != nil {
		return m.Drop
	}
	return false
}

func (m *ACEAction) GetLimit() bool {
	if m != nil {
		return m.Limit
	}
	return false
}

func (m *ACEAction) GetLimitrate() uint32 {
	if m != nil {
		return m.Limitrate
	}
	return 0
}

func (m *ACEAction) GetLimitunit() string {
	if m != nil {
		return m.Limitunit
	}
	return ""
}

func (m *ACEAction) GetLimitburst() uint32 {
	if m != nil {
		return m.Limitburst
	}
	return 0
}

func (m *ACEAction) GetPortmap() bool {
	if m != nil {
		return m.Portmap
	}
	return false
}

func (m *ACEAction) GetAppPort() uint32 {
	if m != nil {
		return m.AppPort
	}
	return 0
}

type ACE struct {
	Matches []*ACEMatch `protobuf:"bytes,1,rep,name=matches,proto3" json:"matches,omitempty"`
	// Expect only single action...repeated here is
	// for future work.
	Actions              []*ACEAction `protobuf:"bytes,2,rep,name=actions,proto3" json:"actions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ACE) Reset()         { *m = ACE{} }
func (m *ACE) String() string { return proto.CompactTextString(m) }
func (*ACE) ProtoMessage()    {}
func (*ACE) Descriptor() ([]byte, []int) {
	return fileDescriptor_505e7efac08d3ba9, []int{2}
}

func (m *ACE) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ACE.Unmarshal(m, b)
}
func (m *ACE) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ACE.Marshal(b, m, deterministic)
}
func (m *ACE) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ACE.Merge(m, src)
}
func (m *ACE) XXX_Size() int {
	return xxx_messageInfo_ACE.Size(m)
}
func (m *ACE) XXX_DiscardUnknown() {
	xxx_messageInfo_ACE.DiscardUnknown(m)
}

var xxx_messageInfo_ACE proto.InternalMessageInfo

func (m *ACE) GetMatches() []*ACEMatch {
	if m != nil {
		return m.Matches
	}
	return nil
}

func (m *ACE) GetActions() []*ACEAction {
	if m != nil {
		return m.Actions
	}
	return nil
}

func init() {
	proto.RegisterType((*ACEMatch)(nil), "ACEMatch")
	proto.RegisterType((*ACEAction)(nil), "ACEAction")
	proto.RegisterType((*ACE)(nil), "ACE")
}

func init() { proto.RegisterFile("fw.proto", fileDescriptor_505e7efac08d3ba9) }

var fileDescriptor_505e7efac08d3ba9 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbb, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x95, 0xde, 0x92, 0x1c, 0xc4, 0x62, 0x31, 0x78, 0x40, 0x50, 0x95, 0x0e, 0x9d, 0x1c,
	0x09, 0x78, 0x81, 0x50, 0x55, 0x4c, 0x48, 0x95, 0x47, 0x36, 0xd7, 0x71, 0x53, 0x8b, 0x24, 0xb6,
	0x1c, 0x3b, 0x88, 0xbe, 0x1f, 0xef, 0x85, 0x72, 0xa2, 0xb4, 0x6c, 0xe7, 0xbf, 0x7c, 0xb2, 0xfc,
	0x43, 0x72, 0xfc, 0x66, 0xd6, 0x19, 0x6f, 0x56, 0xaf, 0x90, 0xe4, 0xdb, 0xdd, 0x87, 0xf0, 0xf2,
	0x44, 0x08, 0xcc, 0xfc, 0x8f, 0x55, 0x34, 0x5a, 0x46, 0x9b, 0x94, 0xe3, 0x4d, 0xee, 0x60, 0xde,
	0x89, 0x2a, 0x28, 0x3a, 0x41, 0x73, 0x10, 0xab, 0xdf, 0x08, 0xd2, 0x7c, 0xbb, 0xcb, 0xa5, 0xd7,
	0xa6, 0xe9, 0xb9, 0xc2, 0x19, 0x8b, 0x5c, 0xc2, 0xf1, 0xee, 0xb9, 0x4a, 0xd7, 0xda, 0x23, 0x97,
	0xf0, 0x41, 0x90, 0x7b, 0x48, 0xf1, 0x70, 0xc2, 0x2b, 0x3a, 0x5d, 0x46, 0x9b, 0x5b, 0x7e, 0x35,
	0x2e, 0x69, 0x68, 0xb4, 0xa7, 0x33, 0x7c, 0xef, 0x6a, 0x90, 0x07, 0x00, 0x14, 0x87, 0xe0, 0x5a,
	0x4f, 0xe7, 0x08, 0xff, 0x73, 0x08, 0x85, 0xd8, 0x1a, 0xe7, 0x6b, 0x61, 0xe9, 0x02, 0xdf, 0x1c,
	0x65, 0x9f, 0x08, 0x6b, 0xf7, 0xc6, 0x79, 0x1a, 0x23, 0x36, 0xca, 0xd5, 0x1e, 0xa6, 0xf9, 0x76,
	0x47, 0x9e, 0x20, 0xae, 0xfb, 0x05, 0x54, 0x4b, 0xa3, 0xe5, 0x74, 0x73, 0xf3, 0x9c, 0xb2, 0x71,
	0x14, 0x3e, 0x26, 0x64, 0x0d, 0xb1, 0xc0, 0xff, 0xb6, 0x74, 0x82, 0x25, 0x60, 0x97, 0x09, 0xf8,
	0x18, 0xbd, 0xbd, 0xc3, 0xa3, 0x34, 0x35, 0x3b, 0xab, 0x42, 0x15, 0x82, 0xc9, 0xca, 0x84, 0x82,
	0x85, 0x56, 0xb9, 0x4e, 0x4b, 0x35, 0x4c, 0xfe, 0xb9, 0x2e, 0xb5, 0x3f, 0x85, 0x03, 0x93, 0xa6,
	0xce, 0x86, 0x5e, 0xa6, 0x3a, 0x95, 0xb5, 0xc5, 0x57, 0x56, 0x9a, 0xec, 0x2c, 0x4d, 0x73, 0xd4,
	0xe5, 0x61, 0x81, 0xe5, 0x97, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x03, 0x30, 0x5a, 0xab,
	0x01, 0x00, 0x00,
}
