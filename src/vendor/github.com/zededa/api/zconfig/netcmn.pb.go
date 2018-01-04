// Code generated by protoc-gen-go. DO NOT EDIT.
// source: netcmn.proto

package zconfig

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type NetworkType int32

const (
	NetworkType_NETWORKTYPENOOP NetworkType = 0
	NetworkType_V4              NetworkType = 4
	NetworkType_V6              NetworkType = 6
	NetworkType_L2              NetworkType = 2
	NetworkType_LISP            NetworkType = 10
)

var NetworkType_name = map[int32]string{
	0:  "NETWORKTYPENOOP",
	4:  "V4",
	6:  "V6",
	2:  "L2",
	10: "LISP",
}
var NetworkType_value = map[string]int32{
	"NETWORKTYPENOOP": 0,
	"V4":              4,
	"V6":              6,
	"L2":              2,
	"LISP":            10,
}

func (x NetworkType) String() string {
	return proto.EnumName(NetworkType_name, int32(x))
}
func (NetworkType) EnumDescriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

type Ipv4Spec struct {
	Dhcp    bool     `protobuf:"varint,2,opt,name=dhcp" json:"dhcp,omitempty"`
	Subnet  string   `protobuf:"bytes,3,opt,name=subnet" json:"subnet,omitempty"`
	Mask    string   `protobuf:"bytes,4,opt,name=mask" json:"mask,omitempty"`
	Gateway string   `protobuf:"bytes,5,opt,name=gateway" json:"gateway,omitempty"`
	Domain  string   `protobuf:"bytes,6,opt,name=domain" json:"domain,omitempty"`
	Ntp     string   `protobuf:"bytes,7,opt,name=ntp" json:"ntp,omitempty"`
	Dns     []string `protobuf:"bytes,8,rep,name=dns" json:"dns,omitempty"`
}

func (m *Ipv4Spec) Reset()                    { *m = Ipv4Spec{} }
func (m *Ipv4Spec) String() string            { return proto.CompactTextString(m) }
func (*Ipv4Spec) ProtoMessage()               {}
func (*Ipv4Spec) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Ipv4Spec) GetDhcp() bool {
	if m != nil {
		return m.Dhcp
	}
	return false
}

func (m *Ipv4Spec) GetSubnet() string {
	if m != nil {
		return m.Subnet
	}
	return ""
}

func (m *Ipv4Spec) GetMask() string {
	if m != nil {
		return m.Mask
	}
	return ""
}

func (m *Ipv4Spec) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *Ipv4Spec) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Ipv4Spec) GetNtp() string {
	if m != nil {
		return m.Ntp
	}
	return ""
}

func (m *Ipv4Spec) GetDns() []string {
	if m != nil {
		return m.Dns
	}
	return nil
}

type Ipv6Spec struct {
	Dhcp    bool     `protobuf:"varint,2,opt,name=dhcp" json:"dhcp,omitempty"`
	Subnet  string   `protobuf:"bytes,3,opt,name=subnet" json:"subnet,omitempty"`
	Mask    string   `protobuf:"bytes,4,opt,name=mask" json:"mask,omitempty"`
	Gateway string   `protobuf:"bytes,5,opt,name=gateway" json:"gateway,omitempty"`
	Domain  string   `protobuf:"bytes,6,opt,name=domain" json:"domain,omitempty"`
	Ntp     string   `protobuf:"bytes,7,opt,name=ntp" json:"ntp,omitempty"`
	Dns     []string `protobuf:"bytes,8,rep,name=dns" json:"dns,omitempty"`
}

func (m *Ipv6Spec) Reset()                    { *m = Ipv6Spec{} }
func (m *Ipv6Spec) String() string            { return proto.CompactTextString(m) }
func (*Ipv6Spec) ProtoMessage()               {}
func (*Ipv6Spec) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{1} }

func (m *Ipv6Spec) GetDhcp() bool {
	if m != nil {
		return m.Dhcp
	}
	return false
}

func (m *Ipv6Spec) GetSubnet() string {
	if m != nil {
		return m.Subnet
	}
	return ""
}

func (m *Ipv6Spec) GetMask() string {
	if m != nil {
		return m.Mask
	}
	return ""
}

func (m *Ipv6Spec) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *Ipv6Spec) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Ipv6Spec) GetNtp() string {
	if m != nil {
		return m.Ntp
	}
	return ""
}

func (m *Ipv6Spec) GetDns() []string {
	if m != nil {
		return m.Dns
	}
	return nil
}

type NameToEid struct {
	Hostname string   `protobuf:"bytes,1,opt,name=hostname" json:"hostname,omitempty"`
	Eids     []string `protobuf:"bytes,2,rep,name=eids" json:"eids,omitempty"`
}

func (m *NameToEid) Reset()                    { *m = NameToEid{} }
func (m *NameToEid) String() string            { return proto.CompactTextString(m) }
func (*NameToEid) ProtoMessage()               {}
func (*NameToEid) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{2} }

func (m *NameToEid) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *NameToEid) GetEids() []string {
	if m != nil {
		return m.Eids
	}
	return nil
}

type EIDAllocation struct {
	Allocate            bool   `protobuf:"varint,1,opt,name=allocate" json:"allocate,omitempty"`
	Exportprivate       bool   `protobuf:"varint,2,opt,name=exportprivate" json:"exportprivate,omitempty"`
	Allocationprefix    []byte `protobuf:"bytes,3,opt,name=allocationprefix,proto3" json:"allocationprefix,omitempty"`
	Allocationprefixlen uint32 `protobuf:"varint,4,opt,name=allocationprefixlen" json:"allocationprefixlen,omitempty"`
}

func (m *EIDAllocation) Reset()                    { *m = EIDAllocation{} }
func (m *EIDAllocation) String() string            { return proto.CompactTextString(m) }
func (*EIDAllocation) ProtoMessage()               {}
func (*EIDAllocation) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{3} }

func (m *EIDAllocation) GetAllocate() bool {
	if m != nil {
		return m.Allocate
	}
	return false
}

func (m *EIDAllocation) GetExportprivate() bool {
	if m != nil {
		return m.Exportprivate
	}
	return false
}

func (m *EIDAllocation) GetAllocationprefix() []byte {
	if m != nil {
		return m.Allocationprefix
	}
	return nil
}

func (m *EIDAllocation) GetAllocationprefixlen() uint32 {
	if m != nil {
		return m.Allocationprefixlen
	}
	return 0
}

type Lispspec struct {
	Iid      uint32         `protobuf:"varint,1,opt,name=iid" json:"iid,omitempty"`
	Eidalloc *EIDAllocation `protobuf:"bytes,2,opt,name=eidalloc" json:"eidalloc,omitempty"`
	Nmtoeid  []*NameToEid   `protobuf:"bytes,3,rep,name=nmtoeid" json:"nmtoeid,omitempty"`
}

func (m *Lispspec) Reset()                    { *m = Lispspec{} }
func (m *Lispspec) String() string            { return proto.CompactTextString(m) }
func (*Lispspec) ProtoMessage()               {}
func (*Lispspec) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{4} }

func (m *Lispspec) GetIid() uint32 {
	if m != nil {
		return m.Iid
	}
	return 0
}

func (m *Lispspec) GetEidalloc() *EIDAllocation {
	if m != nil {
		return m.Eidalloc
	}
	return nil
}

func (m *Lispspec) GetNmtoeid() []*NameToEid {
	if m != nil {
		return m.Nmtoeid
	}
	return nil
}

type L2Spec struct {
}

func (m *L2Spec) Reset()                    { *m = L2Spec{} }
func (m *L2Spec) String() string            { return proto.CompactTextString(m) }
func (*L2Spec) ProtoMessage()               {}
func (*L2Spec) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{5} }

func init() {
	proto.RegisterType((*Ipv4Spec)(nil), "ipv4spec")
	proto.RegisterType((*Ipv6Spec)(nil), "ipv6spec")
	proto.RegisterType((*NameToEid)(nil), "NameToEid")
	proto.RegisterType((*EIDAllocation)(nil), "EIDAllocation")
	proto.RegisterType((*Lispspec)(nil), "lispspec")
	proto.RegisterType((*L2Spec)(nil), "l2spec")
	proto.RegisterEnum("NetworkType", NetworkType_name, NetworkType_value)
}

func init() { proto.RegisterFile("netcmn.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x25, 0x71, 0x70, 0xdc, 0x49, 0x03, 0xd6, 0x56, 0x42, 0x2b, 0x24, 0x44, 0x14, 0xf5, 0x10,
	0xe5, 0xe0, 0xa0, 0x50, 0xf5, 0xc2, 0x05, 0x50, 0x73, 0xa8, 0xa8, 0x92, 0xc8, 0x44, 0x20, 0xb8,
	0x6d, 0x76, 0xa7, 0xc9, 0xaa, 0xf6, 0xee, 0xca, 0xbb, 0x49, 0x3f, 0x7e, 0x0d, 0xff, 0x80, 0xbf,
	0x88, 0x76, 0x4d, 0x2c, 0x15, 0xf8, 0x01, 0x9c, 0xe6, 0xbd, 0x37, 0x33, 0xcf, 0xcf, 0x63, 0x19,
	0x8e, 0x15, 0x3a, 0x5e, 0xaa, 0xcc, 0x54, 0xda, 0xe9, 0xe1, 0x8f, 0x16, 0x24, 0xd2, 0xec, 0xcf,
	0xac, 0x41, 0x4e, 0x08, 0x74, 0xc4, 0x96, 0x1b, 0xda, 0x1e, 0xb4, 0x46, 0x49, 0x1e, 0x30, 0x79,
	0x01, 0xb1, 0xdd, 0xad, 0x15, 0x3a, 0x1a, 0x0d, 0x5a, 0xa3, 0xa3, 0xfc, 0x37, 0xf3, 0xb3, 0x25,
	0xb3, 0x37, 0xb4, 0x13, 0xd4, 0x80, 0x09, 0x85, 0xee, 0x86, 0x39, 0xbc, 0x65, 0xf7, 0xf4, 0x69,
	0x90, 0x0f, 0xd4, 0xbb, 0x08, 0x5d, 0x32, 0xa9, 0x68, 0x5c, 0xbb, 0xd4, 0x8c, 0xa4, 0x10, 0x29,
	0x67, 0x68, 0x37, 0x88, 0x1e, 0x7a, 0x45, 0x28, 0x4b, 0x93, 0x41, 0xe4, 0x15, 0xa1, 0xec, 0x21,
	0xe2, 0xf9, 0x7f, 0x1c, 0xf1, 0x1d, 0x1c, 0xcd, 0x59, 0x89, 0x2b, 0x3d, 0x93, 0x82, 0xbc, 0x84,
	0x64, 0xab, 0xad, 0x53, 0xac, 0x44, 0xda, 0x0a, 0x5b, 0x0d, 0xf7, 0x91, 0x50, 0x0a, 0x4b, 0xdb,
	0x61, 0x37, 0xe0, 0xe1, 0xcf, 0x16, 0xf4, 0x67, 0x97, 0x17, 0x1f, 0x8a, 0x42, 0x73, 0xe6, 0xa4,
	0x56, 0xde, 0x81, 0xd5, 0xac, 0x76, 0x48, 0xf2, 0x86, 0x93, 0x53, 0xe8, 0xe3, 0x9d, 0xd1, 0x95,
	0x33, 0x95, 0xdc, 0xfb, 0x81, 0xfa, 0x12, 0x8f, 0x45, 0x32, 0x86, 0x94, 0x35, 0x7e, 0xa6, 0xc2,
	0x6b, 0x79, 0x17, 0x8e, 0x73, 0x9c, 0xff, 0xa5, 0x93, 0x37, 0x70, 0xf2, 0xa7, 0x56, 0xa0, 0x0a,
	0x57, 0xeb, 0xe7, 0xff, 0x6a, 0x0d, 0x15, 0x24, 0x85, 0xb4, 0x26, 0x7c, 0x90, 0x14, 0x22, 0x29,
	0x45, 0x88, 0xd9, 0xcf, 0x3d, 0x24, 0x63, 0x48, 0x50, 0x8a, 0xb0, 0x17, 0xc2, 0xf5, 0xa6, 0xcf,
	0xb2, 0x47, 0xef, 0x97, 0x37, 0x7d, 0x72, 0x0a, 0x5d, 0x55, 0x3a, 0x8d, 0x52, 0xd0, 0x68, 0x10,
	0x8d, 0x7a, 0x53, 0xc8, 0x9a, 0x43, 0xe6, 0x87, 0xd6, 0x30, 0x81, 0xb8, 0x98, 0xfa, 0xa7, 0x8d,
	0x2f, 0xa0, 0x37, 0x47, 0x77, 0xab, 0xab, 0x9b, 0xd5, 0xbd, 0x41, 0x72, 0x02, 0xcf, 0xe7, 0xb3,
	0xd5, 0xd7, 0x45, 0xfe, 0x69, 0xf5, 0x6d, 0x39, 0x9b, 0x2f, 0x16, 0xcb, 0xf4, 0x09, 0x89, 0xa1,
	0xfd, 0xe5, 0x2c, 0xed, 0x84, 0x7a, 0x9e, 0xc6, 0xbe, 0x5e, 0x4d, 0xd3, 0x36, 0x49, 0xa0, 0x73,
	0x75, 0xf9, 0x79, 0x99, 0xc2, 0xc7, 0xf7, 0xf0, 0x9a, 0xeb, 0x32, 0x7b, 0x40, 0x81, 0x82, 0x65,
	0xbc, 0xd0, 0x3b, 0x91, 0xed, 0x2c, 0x56, 0x7b, 0xc9, 0xb1, 0xfe, 0x2f, 0xbe, 0xbf, 0xda, 0x48,
	0xb7, 0xdd, 0xad, 0x33, 0xae, 0xcb, 0x49, 0x3d, 0x37, 0x61, 0x46, 0x4e, 0x1e, 0xb8, 0x56, 0xd7,
	0x72, 0xb3, 0x8e, 0xc3, 0xd4, 0xdb, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x0a, 0x31, 0x47, 0x32,
	0x4d, 0x03, 0x00, 0x00,
}