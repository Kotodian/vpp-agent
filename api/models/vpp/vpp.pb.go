// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: models/vpp/vpp.proto

package vpp // import "github.com/ligato/vpp-agent/api/models/vpp"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import acl "github.com/ligato/vpp-agent/api/models/vpp/acl"
import interfaces "github.com/ligato/vpp-agent/api/models/vpp/interfaces"
import ipsec "github.com/ligato/vpp-agent/api/models/vpp/ipsec"
import l2 "github.com/ligato/vpp-agent/api/models/vpp/l2"
import l3 "github.com/ligato/vpp-agent/api/models/vpp/l3"
import nat "github.com/ligato/vpp-agent/api/models/vpp/nat"
import punt "github.com/ligato/vpp-agent/api/models/vpp/punt"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type ConfigData struct {
	Interfaces           []*interfaces.Interface         `protobuf:"bytes,10,rep,name=interfaces,proto3" json:"interfaces,omitempty"`
	Acls                 []*acl.Acl                      `protobuf:"bytes,20,rep,name=acls,proto3" json:"acls,omitempty"`
	BridgeDomains        []*l2.BridgeDomain              `protobuf:"bytes,30,rep,name=bridge_domains,json=bridgeDomains,proto3" json:"bridge_domains,omitempty"`
	Fibs                 []*l2.FIBEntry                  `protobuf:"bytes,31,rep,name=fibs,proto3" json:"fibs,omitempty"`
	XconnectPairs        []*l2.XConnectPair              `protobuf:"bytes,32,rep,name=xconnect_pairs,json=xconnectPairs,proto3" json:"xconnect_pairs,omitempty"`
	Routes               []*l3.Route                     `protobuf:"bytes,40,rep,name=routes,proto3" json:"routes,omitempty"`
	Arps                 []*l3.ARPEntry                  `protobuf:"bytes,41,rep,name=arps,proto3" json:"arps,omitempty"`
	ProxyArp             *l3.ProxyARP                    `protobuf:"bytes,42,opt,name=proxy_arp,json=proxyArp,proto3" json:"proxy_arp,omitempty"`
	IpscanNeighbor       *l3.IPScanNeighbor              `protobuf:"bytes,43,opt,name=ipscan_neighbor,json=ipscanNeighbor,proto3" json:"ipscan_neighbor,omitempty"`
	Nat44Global          *nat.Nat44Global                `protobuf:"bytes,50,opt,name=nat44_global,json=nat44Global,proto3" json:"nat44_global,omitempty"`
	Dnat44S              []*nat.DNat44                   `protobuf:"bytes,51,rep,name=dnat44s,proto3" json:"dnat44s,omitempty"`
	IpsecSpds            []*ipsec.SecurityPolicyDatabase `protobuf:"bytes,60,rep,name=ipsec_spds,json=ipsecSpds,proto3" json:"ipsec_spds,omitempty"`
	IpsecSas             []*ipsec.SecurityAssociation    `protobuf:"bytes,61,rep,name=ipsec_sas,json=ipsecSas,proto3" json:"ipsec_sas,omitempty"`
	PuntIpredirects      []*punt.IpRedirect              `protobuf:"bytes,70,rep,name=punt_ipredirects,json=puntIpredirects,proto3" json:"punt_ipredirects,omitempty"`
	PuntTohosts          []*punt.ToHost                  `protobuf:"bytes,71,rep,name=punt_tohosts,json=puntTohosts,proto3" json:"punt_tohosts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ConfigData) Reset()         { *m = ConfigData{} }
func (m *ConfigData) String() string { return proto.CompactTextString(m) }
func (*ConfigData) ProtoMessage()    {}
func (*ConfigData) Descriptor() ([]byte, []int) {
	return fileDescriptor_vpp_3389ccf9551be495, []int{0}
}
func (m *ConfigData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigData.Unmarshal(m, b)
}
func (m *ConfigData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigData.Marshal(b, m, deterministic)
}
func (dst *ConfigData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigData.Merge(dst, src)
}
func (m *ConfigData) XXX_Size() int {
	return xxx_messageInfo_ConfigData.Size(m)
}
func (m *ConfigData) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigData.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigData proto.InternalMessageInfo

func (m *ConfigData) GetInterfaces() []*interfaces.Interface {
	if m != nil {
		return m.Interfaces
	}
	return nil
}

func (m *ConfigData) GetAcls() []*acl.Acl {
	if m != nil {
		return m.Acls
	}
	return nil
}

func (m *ConfigData) GetBridgeDomains() []*l2.BridgeDomain {
	if m != nil {
		return m.BridgeDomains
	}
	return nil
}

func (m *ConfigData) GetFibs() []*l2.FIBEntry {
	if m != nil {
		return m.Fibs
	}
	return nil
}

func (m *ConfigData) GetXconnectPairs() []*l2.XConnectPair {
	if m != nil {
		return m.XconnectPairs
	}
	return nil
}

func (m *ConfigData) GetRoutes() []*l3.Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

func (m *ConfigData) GetArps() []*l3.ARPEntry {
	if m != nil {
		return m.Arps
	}
	return nil
}

func (m *ConfigData) GetProxyArp() *l3.ProxyARP {
	if m != nil {
		return m.ProxyArp
	}
	return nil
}

func (m *ConfigData) GetIpscanNeighbor() *l3.IPScanNeighbor {
	if m != nil {
		return m.IpscanNeighbor
	}
	return nil
}

func (m *ConfigData) GetNat44Global() *nat.Nat44Global {
	if m != nil {
		return m.Nat44Global
	}
	return nil
}

func (m *ConfigData) GetDnat44S() []*nat.DNat44 {
	if m != nil {
		return m.Dnat44S
	}
	return nil
}

func (m *ConfigData) GetIpsecSpds() []*ipsec.SecurityPolicyDatabase {
	if m != nil {
		return m.IpsecSpds
	}
	return nil
}

func (m *ConfigData) GetIpsecSas() []*ipsec.SecurityAssociation {
	if m != nil {
		return m.IpsecSas
	}
	return nil
}

func (m *ConfigData) GetPuntIpredirects() []*punt.IpRedirect {
	if m != nil {
		return m.PuntIpredirects
	}
	return nil
}

func (m *ConfigData) GetPuntTohosts() []*punt.ToHost {
	if m != nil {
		return m.PuntTohosts
	}
	return nil
}

func (*ConfigData) XXX_MessageName() string {
	return "vpp.ConfigData"
}

type Notification struct {
	Interface            *interfaces.InterfaceNotification `protobuf:"bytes,1,opt,name=interface,proto3" json:"interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *Notification) Reset()         { *m = Notification{} }
func (m *Notification) String() string { return proto.CompactTextString(m) }
func (*Notification) ProtoMessage()    {}
func (*Notification) Descriptor() ([]byte, []int) {
	return fileDescriptor_vpp_3389ccf9551be495, []int{1}
}
func (m *Notification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Notification.Unmarshal(m, b)
}
func (m *Notification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Notification.Marshal(b, m, deterministic)
}
func (dst *Notification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Notification.Merge(dst, src)
}
func (m *Notification) XXX_Size() int {
	return xxx_messageInfo_Notification.Size(m)
}
func (m *Notification) XXX_DiscardUnknown() {
	xxx_messageInfo_Notification.DiscardUnknown(m)
}

var xxx_messageInfo_Notification proto.InternalMessageInfo

func (m *Notification) GetInterface() *interfaces.InterfaceNotification {
	if m != nil {
		return m.Interface
	}
	return nil
}

func (*Notification) XXX_MessageName() string {
	return "vpp.Notification"
}
func init() {
	proto.RegisterType((*ConfigData)(nil), "vpp.ConfigData")
	proto.RegisterType((*Notification)(nil), "vpp.Notification")
}

func init() { proto.RegisterFile("models/vpp/vpp.proto", fileDescriptor_vpp_3389ccf9551be495) }

var fileDescriptor_vpp_3389ccf9551be495 = []byte{
	// 671 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x94, 0xdf, 0x4e, 0xdb, 0x3e,
	0x14, 0xc7, 0x85, 0x40, 0xfc, 0xa8, 0x29, 0x7f, 0x64, 0x21, 0x7e, 0xa6, 0x9a, 0x58, 0x41, 0x43,
	0x82, 0x4d, 0x4d, 0xa4, 0x16, 0x69, 0x9a, 0xd8, 0xc4, 0x0a, 0x0c, 0xd6, 0x1b, 0x54, 0xb9, 0x5c,
	0x4c, 0xbb, 0x89, 0x1c, 0x27, 0x0d, 0x96, 0x82, 0x6d, 0xd9, 0x2e, 0xa2, 0x4f, 0xb3, 0xd7, 0xd9,
	0x7b, 0xec, 0x45, 0x26, 0x9f, 0x24, 0x24, 0x95, 0xba, 0x8b, 0x54, 0x39, 0xfe, 0x7c, 0xbf, 0x27,
	0xf6, 0xf1, 0x39, 0x45, 0x7b, 0x4f, 0x2a, 0x49, 0x73, 0x1b, 0x3e, 0x6b, 0xed, 0x9f, 0x40, 0x1b,
	0xe5, 0x14, 0x5e, 0x7d, 0xd6, 0xba, 0xd3, 0xcb, 0x84, 0x7b, 0x9c, 0xc5, 0x01, 0x57, 0x4f, 0x61,
	0xa6, 0x32, 0x15, 0x02, 0x8b, 0x67, 0x53, 0x88, 0x20, 0x80, 0xb7, 0xc2, 0xd3, 0x21, 0x8d, 0x4c,
	0x8c, 0xe7, 0xfe, 0x29, 0xc9, 0x49, 0x83, 0x08, 0xe9, 0x52, 0x33, 0x65, 0x3c, 0xb5, 0xf5, 0x6b,
	0x29, 0x3b, 0x5a, 0x2e, 0xb3, 0x8e, 0xb9, 0x4a, 0xf2, 0xa6, 0x29, 0xd1, 0x36, 0xe5, 0xc5, 0xef,
	0x92, 0x04, 0x79, 0x3f, 0x8c, 0x8d, 0x48, 0xb2, 0xb4, 0x97, 0xa8, 0x27, 0x26, 0x64, 0x29, 0xf9,
	0x7f, 0x51, 0x32, 0x15, 0xf1, 0x92, 0xcc, 0x79, 0x3f, 0x7c, 0xe1, 0x4a, 0xca, 0x94, 0xbb, 0x65,
	0xb6, 0x41, 0xc8, 0x4c, 0x59, 0xa8, 0xce, 0xfe, 0x22, 0xc8, 0x07, 0xe5, 0xfa, 0xc1, 0xe2, 0xba,
	0x51, 0xb3, 0xd7, 0x33, 0x34, 0xeb, 0x24, 0x99, 0xf3, 0x4f, 0x49, 0x3a, 0x0d, 0xa2, 0x67, 0xd2,
	0xc1, 0x4f, 0xc1, 0x8e, 0x7f, 0xad, 0x23, 0x74, 0xad, 0xe4, 0x54, 0x64, 0x37, 0xcc, 0x31, 0xfc,
	0x09, 0xa1, 0xba, 0x44, 0x04, 0x75, 0x57, 0x4f, 0x37, 0xfb, 0x07, 0x81, 0xbf, 0xc0, 0x7a, 0x39,
	0x18, 0x55, 0xaf, 0xb4, 0x21, 0xc6, 0x5d, 0xb4, 0xc6, 0x78, 0x6e, 0xc9, 0x1e, 0x98, 0xda, 0x60,
	0xf2, 0x77, 0x35, 0xe4, 0x39, 0x05, 0x82, 0x2f, 0xd0, 0x76, 0x51, 0xbb, 0xa8, 0xa8, 0x9d, 0x25,
	0x87, 0xa0, 0xdd, 0x03, 0x6d, 0xde, 0x0f, 0xae, 0x80, 0xde, 0x00, 0xa4, 0x5b, 0x71, 0x23, 0xb2,
	0xf8, 0x1d, 0x5a, 0x9b, 0x8a, 0xd8, 0x92, 0xb7, 0x60, 0xd9, 0xad, 0x2c, 0xb7, 0xa3, 0xab, 0x6f,
	0xd2, 0x99, 0x39, 0x05, 0xea, 0x3f, 0x51, 0x95, 0x38, 0xd2, 0x4c, 0x18, 0x4b, 0xba, 0x8b, 0x9f,
	0xf8, 0x71, 0x5d, 0xd0, 0x31, 0x13, 0x86, 0x6e, 0x55, 0x5a, 0x1f, 0x59, 0x7c, 0x82, 0xd6, 0xa1,
	0xa0, 0x96, 0x9c, 0x82, 0x69, 0xab, 0x30, 0x0d, 0x02, 0xea, 0x57, 0x69, 0x09, 0xfd, 0x4e, 0x98,
	0xd1, 0x96, 0x9c, 0x35, 0x77, 0x32, 0x08, 0x86, 0x74, 0x5c, 0xee, 0xc4, 0x53, 0xdc, 0x43, 0x2d,
	0x6d, 0xd4, 0xcb, 0x3c, 0x62, 0x46, 0x93, 0xf7, 0xdd, 0x95, 0xa6, 0x74, 0xec, 0xc1, 0x90, 0x8e,
	0xe9, 0x06, 0x48, 0x86, 0x46, 0xe3, 0x4b, 0xb4, 0x23, 0xb4, 0xe5, 0x4c, 0x46, 0x32, 0x15, 0xd9,
	0x63, 0xac, 0x0c, 0xf9, 0x00, 0xa6, 0xfd, 0xca, 0x34, 0x1a, 0x4f, 0x38, 0x93, 0xf7, 0x25, 0xa5,
	0xdb, 0x85, 0xbc, 0x8a, 0xf1, 0x47, 0xd4, 0x96, 0xcc, 0x9d, 0x9f, 0x47, 0x59, 0xae, 0x62, 0x96,
	0x93, 0x3e, 0xb8, 0x8b, 0x73, 0xfb, 0x56, 0xb8, 0xf7, 0xf0, 0x0e, 0x18, 0xdd, 0x94, 0x75, 0x80,
	0xcf, 0xd0, 0x7f, 0x09, 0xc4, 0x96, 0x0c, 0xe0, 0x44, 0x3b, 0xaf, 0x9e, 0x1b, 0x30, 0xd1, 0x8a,
	0xe3, 0xaf, 0x08, 0xc1, 0x5c, 0x44, 0x56, 0x27, 0x96, 0x7c, 0x06, 0xf5, 0x51, 0xd1, 0x1d, 0x30,
	0x2e, 0x93, 0x94, 0xcf, 0x8c, 0x70, 0xf3, 0xb1, 0xca, 0x05, 0x9f, 0xfb, 0x86, 0x8a, 0x99, 0x4d,
	0x69, 0x0b, 0xe8, 0x44, 0x27, 0xfe, 0x7e, 0x5a, 0x65, 0x06, 0x66, 0xc9, 0x17, 0x48, 0x70, 0xb8,
	0x24, 0xc1, 0xd0, 0x5a, 0xc5, 0x05, 0x73, 0x42, 0x49, 0xba, 0x51, 0xb8, 0x99, 0xc5, 0x97, 0x68,
	0xd7, 0x77, 0x6e, 0x24, 0xb4, 0x49, 0x13, 0x61, 0x52, 0xee, 0x2c, 0xb9, 0x6d, 0x5c, 0x2f, 0xb4,
	0xf5, 0x48, 0xd3, 0x12, 0xd2, 0x1d, 0xbf, 0x30, 0xaa, 0xc5, 0x78, 0x80, 0xda, 0x90, 0xc0, 0xa9,
	0x47, 0x65, 0x9d, 0x25, 0x77, 0x8d, 0x1b, 0x04, 0xf3, 0x83, 0xfa, 0xae, 0xac, 0xa3, 0x9b, 0x3e,
	0x78, 0x28, 0x44, 0xc7, 0x13, 0xd4, 0xbe, 0x57, 0x4e, 0x4c, 0x05, 0x87, 0xfd, 0xe0, 0x6b, 0xd4,
	0x7a, 0xed, 0x7a, 0xb2, 0x02, 0x55, 0x3e, 0xf9, 0xe7, 0x84, 0x34, 0x9d, 0xb4, 0xf6, 0x5d, 0x9d,
	0xff, 0xfe, 0x73, 0xb8, 0xf2, 0x33, 0x68, 0xfc, 0x13, 0xe6, 0x22, 0x63, 0x4e, 0xf9, 0x19, 0xed,
	0xb1, 0x2c, 0x95, 0x2e, 0x64, 0x5a, 0x84, 0xf5, 0xe0, 0x5e, 0x3c, 0x6b, 0x1d, 0xaf, 0xc3, 0xcc,
	0x0e, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xfd, 0x45, 0x0c, 0x5d, 0x05, 0x00, 0x00,
}
