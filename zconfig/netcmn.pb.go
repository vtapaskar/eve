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

type ProxyProto int32

const (
	ProxyProto_PROXY_HTTP  ProxyProto = 0
	ProxyProto_PROXY_HTTPS ProxyProto = 1
	ProxyProto_PROXY_SOCKS ProxyProto = 2
	ProxyProto_PROXY_FTP   ProxyProto = 3
	ProxyProto_PROXY_OTHER ProxyProto = 255
)

var ProxyProto_name = map[int32]string{
	0:   "PROXY_HTTP",
	1:   "PROXY_HTTPS",
	2:   "PROXY_SOCKS",
	3:   "PROXY_FTP",
	255: "PROXY_OTHER",
}
var ProxyProto_value = map[string]int32{
	"PROXY_HTTP":  0,
	"PROXY_HTTPS": 1,
	"PROXY_SOCKS": 2,
	"PROXY_FTP":   3,
	"PROXY_OTHER": 255,
}

func (x ProxyProto) String() string {
	return proto.EnumName(ProxyProto_name, int32(x))
}
func (ProxyProto) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

type DHCPType int32

const (
	DHCPType_DHCPNoop DHCPType = 0
	// used for device adapter
	DHCPType_Static DHCPType = 1
	// used for application simulation
	DHCPType_PassThrough DHCPType = 2
	// used for application simulation
	DHCPType_Server DHCPType = 3
	// used for device adapter
	DHCPType_Client DHCPType = 4
)

var DHCPType_name = map[int32]string{
	0: "DHCPNoop",
	1: "Static",
	2: "PassThrough",
	3: "Server",
	4: "Client",
}
var DHCPType_value = map[string]int32{
	"DHCPNoop":    0,
	"Static":      1,
	"PassThrough": 2,
	"Server":      3,
	"Client":      4,
}

func (x DHCPType) String() string {
	return proto.EnumName(DHCPType_name, int32(x))
}
func (DHCPType) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

type NetworkType int32

const (
	NetworkType_NETWORKTYPENOOP NetworkType = 0
	NetworkType_V4              NetworkType = 4
	NetworkType_V6              NetworkType = 6
	NetworkType_CryptoEID       NetworkType = 14
)

var NetworkType_name = map[int32]string{
	0:  "NETWORKTYPENOOP",
	4:  "V4",
	6:  "V6",
	14: "CryptoEID",
}
var NetworkType_value = map[string]int32{
	"NETWORKTYPENOOP": 0,
	"V4":              4,
	"V6":              6,
	"CryptoEID":       14,
}

func (x NetworkType) String() string {
	return proto.EnumName(NetworkType_name, int32(x))
}
func (NetworkType) EnumDescriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

type IpRange struct {
	Start string `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	End   string `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
}

func (m *IpRange) Reset()                    { *m = IpRange{} }
func (m *IpRange) String() string            { return proto.CompactTextString(m) }
func (*IpRange) ProtoMessage()               {}
func (*IpRange) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *IpRange) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *IpRange) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

type ProxyServer struct {
	Proto  ProxyProto `protobuf:"varint,1,opt,name=proto,enum=ProxyProto" json:"proto,omitempty"`
	Server string     `protobuf:"bytes,2,opt,name=server" json:"server,omitempty"`
	Port   uint32     `protobuf:"varint,3,opt,name=port" json:"port,omitempty"`
}

func (m *ProxyServer) Reset()                    { *m = ProxyServer{} }
func (m *ProxyServer) String() string            { return proto.CompactTextString(m) }
func (*ProxyServer) ProtoMessage()               {}
func (*ProxyServer) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

func (m *ProxyServer) GetProto() ProxyProto {
	if m != nil {
		return m.Proto
	}
	return ProxyProto_PROXY_HTTP
}

func (m *ProxyServer) GetServer() string {
	if m != nil {
		return m.Server
	}
	return ""
}

func (m *ProxyServer) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

type ProxyConfig struct {
	// enable network level proxy
	NetworkProxyEnable bool `protobuf:"varint,1,opt,name=networkProxyEnable" json:"networkProxyEnable,omitempty"`
	// dedicated per protocol information
	Proxies []*ProxyServer `protobuf:"bytes,2,rep,name=proxies" json:"proxies,omitempty"`
	// exceptions seperated by commas
	Exceptions string `protobuf:"bytes,3,opt,name=exceptions" json:"exceptions,omitempty"`
	// or pacfile can be in place of others
	// base64 encoded
	Pacfile string `protobuf:"bytes,4,opt,name=pacfile" json:"pacfile,omitempty"`
}

func (m *ProxyConfig) Reset()                    { *m = ProxyConfig{} }
func (m *ProxyConfig) String() string            { return proto.CompactTextString(m) }
func (*ProxyConfig) ProtoMessage()               {}
func (*ProxyConfig) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *ProxyConfig) GetNetworkProxyEnable() bool {
	if m != nil {
		return m.NetworkProxyEnable
	}
	return false
}

func (m *ProxyConfig) GetProxies() []*ProxyServer {
	if m != nil {
		return m.Proxies
	}
	return nil
}

func (m *ProxyConfig) GetExceptions() string {
	if m != nil {
		return m.Exceptions
	}
	return ""
}

func (m *ProxyConfig) GetPacfile() string {
	if m != nil {
		return m.Pacfile
	}
	return ""
}

// These are list of static mapping that can be added to network
type ZnetStaticDNSEntry struct {
	HostName string   `protobuf:"bytes,1,opt,name=HostName" json:"HostName,omitempty"`
	Address  []string `protobuf:"bytes,2,rep,name=Address" json:"Address,omitempty"`
}

func (m *ZnetStaticDNSEntry) Reset()                    { *m = ZnetStaticDNSEntry{} }
func (m *ZnetStaticDNSEntry) String() string            { return proto.CompactTextString(m) }
func (*ZnetStaticDNSEntry) ProtoMessage()               {}
func (*ZnetStaticDNSEntry) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

func (m *ZnetStaticDNSEntry) GetHostName() string {
	if m != nil {
		return m.HostName
	}
	return ""
}

func (m *ZnetStaticDNSEntry) GetAddress() []string {
	if m != nil {
		return m.Address
	}
	return nil
}

// Common for IPv4 and IPv6
type Ipspec struct {
	Dhcp DHCPType `protobuf:"varint,2,opt,name=dhcp,enum=DHCPType" json:"dhcp,omitempty"`
	// subnet is CIDR format...x.y.z.l/nn
	Subnet  string   `protobuf:"bytes,3,opt,name=subnet" json:"subnet,omitempty"`
	Gateway string   `protobuf:"bytes,5,opt,name=gateway" json:"gateway,omitempty"`
	Domain  string   `protobuf:"bytes,6,opt,name=domain" json:"domain,omitempty"`
	Ntp     string   `protobuf:"bytes,7,opt,name=ntp" json:"ntp,omitempty"`
	Dns     []string `protobuf:"bytes,8,rep,name=dns" json:"dns,omitempty"`
	// for IPAM management when dhcp is turned on.
	// If none provided, system will default pool.
	DhcpRange *IpRange `protobuf:"bytes,9,opt,name=dhcpRange" json:"dhcpRange,omitempty"`
}

func (m *Ipspec) Reset()                    { *m = Ipspec{} }
func (m *Ipspec) String() string            { return proto.CompactTextString(m) }
func (*Ipspec) ProtoMessage()               {}
func (*Ipspec) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{4} }

func (m *Ipspec) GetDhcp() DHCPType {
	if m != nil {
		return m.Dhcp
	}
	return DHCPType_DHCPNoop
}

func (m *Ipspec) GetSubnet() string {
	if m != nil {
		return m.Subnet
	}
	return ""
}

func (m *Ipspec) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *Ipspec) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *Ipspec) GetNtp() string {
	if m != nil {
		return m.Ntp
	}
	return ""
}

func (m *Ipspec) GetDns() []string {
	if m != nil {
		return m.Dns
	}
	return nil
}

func (m *Ipspec) GetDhcpRange() *IpRange {
	if m != nil {
		return m.DhcpRange
	}
	return nil
}

func init() {
	proto.RegisterType((*IpRange)(nil), "ipRange")
	proto.RegisterType((*ProxyServer)(nil), "ProxyServer")
	proto.RegisterType((*ProxyConfig)(nil), "ProxyConfig")
	proto.RegisterType((*ZnetStaticDNSEntry)(nil), "ZnetStaticDNSEntry")
	proto.RegisterType((*Ipspec)(nil), "ipspec")
	proto.RegisterEnum("ProxyProto", ProxyProto_name, ProxyProto_value)
	proto.RegisterEnum("DHCPType", DHCPType_name, DHCPType_value)
	proto.RegisterEnum("NetworkType", NetworkType_name, NetworkType_value)
}

func init() { proto.RegisterFile("netcmn.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 576 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0x51, 0x6f, 0xd3, 0x3c,
	0x14, 0x5d, 0xda, 0x2e, 0x6d, 0x6e, 0xb7, 0xce, 0xf2, 0xf7, 0x09, 0x45, 0x48, 0x83, 0xd2, 0x87,
	0xa9, 0xda, 0x43, 0x26, 0x0a, 0xe2, 0x99, 0xd1, 0x15, 0x15, 0x26, 0xda, 0xc8, 0x8d, 0x80, 0x4d,
	0x48, 0x93, 0x9b, 0x78, 0xad, 0x45, 0x6b, 0x5b, 0x89, 0xcb, 0x96, 0xfd, 0x19, 0x7e, 0x09, 0xbf,
	0x0d, 0x64, 0x3b, 0xa5, 0x7b, 0xe0, 0x29, 0xf7, 0x9c, 0x63, 0x1f, 0x1f, 0xdd, 0x7b, 0x03, 0x07,
	0x82, 0xe9, 0x74, 0x2d, 0x22, 0x95, 0x4b, 0x2d, 0x7b, 0x2f, 0xa1, 0xc9, 0x15, 0xa1, 0x62, 0xc1,
	0xf0, 0xff, 0xb0, 0x5f, 0x68, 0x9a, 0xeb, 0xd0, 0xeb, 0x7a, 0xfd, 0x80, 0x38, 0x80, 0x11, 0xd4,
	0x99, 0xc8, 0xc2, 0x9a, 0xe5, 0x4c, 0xd9, 0xfb, 0x06, 0xed, 0x38, 0x97, 0xf7, 0xe5, 0x8c, 0xe5,
	0x3f, 0x58, 0x8e, 0x5f, 0xc0, 0xbe, 0xb5, 0xb2, 0xd7, 0x3a, 0x83, 0xb6, 0x31, 0xbe, 0x2f, 0x63,
	0x43, 0x11, 0xa7, 0xe0, 0x27, 0xe0, 0x17, 0xf6, 0x70, 0x65, 0x53, 0x21, 0x8c, 0xa1, 0xa1, 0x64,
	0xae, 0xc3, 0x7a, 0xd7, 0xeb, 0x1f, 0x12, 0x5b, 0xf7, 0x7e, 0x7a, 0x95, 0xfd, 0x50, 0x8a, 0x5b,
	0xbe, 0xc0, 0x11, 0x60, 0xc1, 0xf4, 0x9d, 0xcc, 0xbf, 0x5b, 0x76, 0x24, 0xe8, 0x7c, 0xc5, 0xec,
	0x5b, 0x2d, 0xf2, 0x0f, 0x05, 0x9f, 0x40, 0xd3, 0x04, 0xe0, 0xac, 0x08, 0x6b, 0xdd, 0x7a, 0xbf,
	0x3d, 0x38, 0x88, 0x1e, 0xa5, 0x25, 0x5b, 0x11, 0x3f, 0x03, 0x60, 0xf7, 0x29, 0x53, 0x9a, 0x4b,
	0x51, 0xd8, 0x04, 0x01, 0x79, 0xc4, 0xe0, 0x10, 0x9a, 0x8a, 0xa6, 0xb7, 0x7c, 0xc5, 0xc2, 0x86,
	0x15, 0xb7, 0xb0, 0xf7, 0x11, 0xf0, 0xb5, 0x60, 0x7a, 0xa6, 0xa9, 0xe6, 0xe9, 0xc5, 0x64, 0x36,
	0x12, 0x3a, 0x2f, 0xf1, 0x53, 0x68, 0x8d, 0x65, 0xa1, 0x27, 0x74, 0xcd, 0xaa, 0x06, 0xfe, 0xc5,
	0xc6, 0xeb, 0x3c, 0xcb, 0x72, 0x56, 0xb8, 0x4c, 0x01, 0xd9, 0xc2, 0xde, 0x2f, 0x0f, 0x7c, 0xae,
	0x0a, 0xc5, 0x52, 0x7c, 0x0c, 0x8d, 0x6c, 0x99, 0x2a, 0xdb, 0xa2, 0xce, 0x20, 0x88, 0x2e, 0xc6,
	0xc3, 0x38, 0x29, 0x15, 0x23, 0x96, 0xb6, 0x3d, 0xdc, 0xcc, 0x05, 0xd3, 0x55, 0xd6, 0x0a, 0x19,
	0xef, 0x05, 0xd5, 0xec, 0x8e, 0x96, 0xe1, 0xbe, 0xcb, 0x59, 0x41, 0x73, 0x23, 0x93, 0x6b, 0xca,
	0x45, 0xe8, 0xbb, 0x1b, 0x0e, 0x99, 0x89, 0x0a, 0xad, 0xc2, 0xa6, 0x9b, 0xa8, 0xd0, 0xca, 0x30,
	0x99, 0x28, 0xc2, 0x96, 0xcd, 0x66, 0x4a, 0x7c, 0x02, 0x81, 0x79, 0xd5, 0x2e, 0x46, 0x18, 0x74,
	0xbd, 0x7e, 0x7b, 0xd0, 0x8a, 0xaa, 0x45, 0x21, 0x3b, 0xe9, 0xf4, 0x06, 0x60, 0x37, 0x6e, 0xdc,
	0x01, 0x88, 0xc9, 0xf4, 0xeb, 0xd5, 0xcd, 0x38, 0x49, 0x62, 0xb4, 0x87, 0x8f, 0xa0, 0xbd, 0xc3,
	0x33, 0xe4, 0xed, 0x88, 0xd9, 0x74, 0x78, 0x39, 0x43, 0x35, 0x7c, 0x08, 0x81, 0x23, 0xde, 0x27,
	0x31, 0xaa, 0x63, 0xb4, 0xd5, 0xa7, 0xc9, 0x78, 0x44, 0xd0, 0x6f, 0xef, 0xf4, 0x13, 0xb4, 0xb6,
	0x8d, 0xc0, 0x07, 0xae, 0x9e, 0x48, 0xa9, 0xd0, 0x1e, 0x06, 0xf0, 0xdd, 0x08, 0x2a, 0x5f, 0x5a,
	0x14, 0xc9, 0x32, 0x97, 0x9b, 0xc5, 0x12, 0xd5, 0xac, 0x68, 0x07, 0x8e, 0xea, 0xa6, 0x1e, 0xae,
	0x38, 0x13, 0x1a, 0x35, 0x4e, 0xcf, 0xa1, 0x3d, 0x71, 0x3b, 0x63, 0x1d, 0xff, 0x83, 0xa3, 0xc9,
	0x28, 0xf9, 0x32, 0x25, 0x97, 0xc9, 0x55, 0x3c, 0x9a, 0x4c, 0xa7, 0x26, 0xb5, 0x0f, 0xb5, 0xcf,
	0xaf, 0x51, 0xc3, 0x7e, 0xdf, 0x20, 0xdf, 0x64, 0x1c, 0xe6, 0xa5, 0xd2, 0x72, 0xf4, 0xe1, 0x02,
	0x75, 0xde, 0xbd, 0x85, 0xe7, 0xa9, 0x5c, 0x47, 0x0f, 0x2c, 0x63, 0x19, 0x8d, 0xd2, 0x95, 0xdc,
	0x64, 0xd1, 0xc6, 0x2c, 0x34, 0x4f, 0x99, 0xfb, 0xa9, 0xae, 0x8f, 0x17, 0x5c, 0x2f, 0x37, 0xf3,
	0x28, 0x95, 0xeb, 0x33, 0x77, 0xee, 0x8c, 0x2a, 0x7e, 0xf6, 0x90, 0xda, 0x95, 0x9e, 0xfb, 0xf6,
	0xd4, 0xab, 0x3f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x81, 0x54, 0xa1, 0xf9, 0x8a, 0x03, 0x00, 0x00,
}
