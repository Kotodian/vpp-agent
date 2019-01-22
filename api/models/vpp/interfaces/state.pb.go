// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: models/vpp/interfaces/state.proto

package vpp_interfaces // import "github.com/ligato/vpp-agent/api/models/vpp/interfaces"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type InterfaceState_Status int32

const (
	InterfaceState_UNKNOWN_STATUS InterfaceState_Status = 0
	InterfaceState_UP             InterfaceState_Status = 1
	InterfaceState_DOWN           InterfaceState_Status = 2
	InterfaceState_DELETED        InterfaceState_Status = 3
)

var InterfaceState_Status_name = map[int32]string{
	0: "UNKNOWN_STATUS",
	1: "UP",
	2: "DOWN",
	3: "DELETED",
}
var InterfaceState_Status_value = map[string]int32{
	"UNKNOWN_STATUS": 0,
	"UP":             1,
	"DOWN":           2,
	"DELETED":        3,
}

func (x InterfaceState_Status) String() string {
	return proto.EnumName(InterfaceState_Status_name, int32(x))
}
func (InterfaceState_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_state_bb2097df664231a9, []int{0, 0}
}

type InterfaceState_Duplex int32

const (
	InterfaceState_UNKNOWN_DUPLEX InterfaceState_Duplex = 0
	InterfaceState_HALF           InterfaceState_Duplex = 1
	InterfaceState_FULL           InterfaceState_Duplex = 2
)

var InterfaceState_Duplex_name = map[int32]string{
	0: "UNKNOWN_DUPLEX",
	1: "HALF",
	2: "FULL",
}
var InterfaceState_Duplex_value = map[string]int32{
	"UNKNOWN_DUPLEX": 0,
	"HALF":           1,
	"FULL":           2,
}

func (x InterfaceState_Duplex) String() string {
	return proto.EnumName(InterfaceState_Duplex_name, int32(x))
}
func (InterfaceState_Duplex) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_state_bb2097df664231a9, []int{0, 1}
}

type InterfaceNotification_NotifType int32

const (
	InterfaceNotification_UNKNOWN  InterfaceNotification_NotifType = 0
	InterfaceNotification_UPDOWN   InterfaceNotification_NotifType = 1
	InterfaceNotification_COUNTERS InterfaceNotification_NotifType = 2
)

var InterfaceNotification_NotifType_name = map[int32]string{
	0: "UNKNOWN",
	1: "UPDOWN",
	2: "COUNTERS",
}
var InterfaceNotification_NotifType_value = map[string]int32{
	"UNKNOWN":  0,
	"UPDOWN":   1,
	"COUNTERS": 2,
}

func (x InterfaceNotification_NotifType) String() string {
	return proto.EnumName(InterfaceNotification_NotifType_name, int32(x))
}
func (InterfaceNotification_NotifType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_state_bb2097df664231a9, []int{1, 0}
}

type InterfaceState struct {
	Name                 string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	InternalName         string                     `protobuf:"bytes,2,opt,name=internal_name,json=internalName,proto3" json:"internal_name,omitempty"`
	Type                 Interface_Type             `protobuf:"varint,3,opt,name=type,proto3,enum=vpp.interfaces.Interface_Type" json:"type,omitempty"`
	IfIndex              uint32                     `protobuf:"varint,4,opt,name=if_index,json=ifIndex,proto3" json:"if_index,omitempty"`
	AdminStatus          InterfaceState_Status      `protobuf:"varint,5,opt,name=admin_status,json=adminStatus,proto3,enum=vpp.interfaces.InterfaceState_Status" json:"admin_status,omitempty"`
	OperStatus           InterfaceState_Status      `protobuf:"varint,6,opt,name=oper_status,json=operStatus,proto3,enum=vpp.interfaces.InterfaceState_Status" json:"oper_status,omitempty"`
	LastChange           int64                      `protobuf:"varint,7,opt,name=last_change,json=lastChange,proto3" json:"last_change,omitempty"`
	PhysAddress          string                     `protobuf:"bytes,8,opt,name=phys_address,json=physAddress,proto3" json:"phys_address,omitempty"`
	Speed                uint64                     `protobuf:"varint,9,opt,name=speed,proto3" json:"speed,omitempty"`
	Mtu                  uint32                     `protobuf:"varint,10,opt,name=mtu,proto3" json:"mtu,omitempty"`
	Duplex               InterfaceState_Duplex      `protobuf:"varint,11,opt,name=duplex,proto3,enum=vpp.interfaces.InterfaceState_Duplex" json:"duplex,omitempty"`
	Statistics           *InterfaceState_Statistics `protobuf:"bytes,100,opt,name=statistics,proto3" json:"statistics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *InterfaceState) Reset()         { *m = InterfaceState{} }
func (m *InterfaceState) String() string { return proto.CompactTextString(m) }
func (*InterfaceState) ProtoMessage()    {}
func (*InterfaceState) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_bb2097df664231a9, []int{0}
}
func (m *InterfaceState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfaceState.Unmarshal(m, b)
}
func (m *InterfaceState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfaceState.Marshal(b, m, deterministic)
}
func (dst *InterfaceState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfaceState.Merge(dst, src)
}
func (m *InterfaceState) XXX_Size() int {
	return xxx_messageInfo_InterfaceState.Size(m)
}
func (m *InterfaceState) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfaceState.DiscardUnknown(m)
}

var xxx_messageInfo_InterfaceState proto.InternalMessageInfo

func (m *InterfaceState) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *InterfaceState) GetInternalName() string {
	if m != nil {
		return m.InternalName
	}
	return ""
}

func (m *InterfaceState) GetType() Interface_Type {
	if m != nil {
		return m.Type
	}
	return Interface_UNDEFINED_TYPE
}

func (m *InterfaceState) GetIfIndex() uint32 {
	if m != nil {
		return m.IfIndex
	}
	return 0
}

func (m *InterfaceState) GetAdminStatus() InterfaceState_Status {
	if m != nil {
		return m.AdminStatus
	}
	return InterfaceState_UNKNOWN_STATUS
}

func (m *InterfaceState) GetOperStatus() InterfaceState_Status {
	if m != nil {
		return m.OperStatus
	}
	return InterfaceState_UNKNOWN_STATUS
}

func (m *InterfaceState) GetLastChange() int64 {
	if m != nil {
		return m.LastChange
	}
	return 0
}

func (m *InterfaceState) GetPhysAddress() string {
	if m != nil {
		return m.PhysAddress
	}
	return ""
}

func (m *InterfaceState) GetSpeed() uint64 {
	if m != nil {
		return m.Speed
	}
	return 0
}

func (m *InterfaceState) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *InterfaceState) GetDuplex() InterfaceState_Duplex {
	if m != nil {
		return m.Duplex
	}
	return InterfaceState_UNKNOWN_DUPLEX
}

func (m *InterfaceState) GetStatistics() *InterfaceState_Statistics {
	if m != nil {
		return m.Statistics
	}
	return nil
}

func (*InterfaceState) XXX_MessageName() string {
	return "vpp.interfaces.InterfaceState"
}

type InterfaceState_Statistics struct {
	InPackets            uint64   `protobuf:"varint,1,opt,name=in_packets,json=inPackets,proto3" json:"in_packets,omitempty"`
	InBytes              uint64   `protobuf:"varint,2,opt,name=in_bytes,json=inBytes,proto3" json:"in_bytes,omitempty"`
	OutPackets           uint64   `protobuf:"varint,3,opt,name=out_packets,json=outPackets,proto3" json:"out_packets,omitempty"`
	OutBytes             uint64   `protobuf:"varint,4,opt,name=out_bytes,json=outBytes,proto3" json:"out_bytes,omitempty"`
	DropPackets          uint64   `protobuf:"varint,5,opt,name=drop_packets,json=dropPackets,proto3" json:"drop_packets,omitempty"`
	PuntPackets          uint64   `protobuf:"varint,6,opt,name=punt_packets,json=puntPackets,proto3" json:"punt_packets,omitempty"`
	Ipv4Packets          uint64   `protobuf:"varint,7,opt,name=ipv4_packets,json=ipv4Packets,proto3" json:"ipv4_packets,omitempty"`
	Ipv6Packets          uint64   `protobuf:"varint,8,opt,name=ipv6_packets,json=ipv6Packets,proto3" json:"ipv6_packets,omitempty"`
	InNobufPackets       uint64   `protobuf:"varint,9,opt,name=in_nobuf_packets,json=inNobufPackets,proto3" json:"in_nobuf_packets,omitempty"`
	InMissPackets        uint64   `protobuf:"varint,10,opt,name=in_miss_packets,json=inMissPackets,proto3" json:"in_miss_packets,omitempty"`
	InErrorPackets       uint64   `protobuf:"varint,11,opt,name=in_error_packets,json=inErrorPackets,proto3" json:"in_error_packets,omitempty"`
	OutErrorPackets      uint64   `protobuf:"varint,12,opt,name=out_error_packets,json=outErrorPackets,proto3" json:"out_error_packets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InterfaceState_Statistics) Reset()         { *m = InterfaceState_Statistics{} }
func (m *InterfaceState_Statistics) String() string { return proto.CompactTextString(m) }
func (*InterfaceState_Statistics) ProtoMessage()    {}
func (*InterfaceState_Statistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_bb2097df664231a9, []int{0, 0}
}
func (m *InterfaceState_Statistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfaceState_Statistics.Unmarshal(m, b)
}
func (m *InterfaceState_Statistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfaceState_Statistics.Marshal(b, m, deterministic)
}
func (dst *InterfaceState_Statistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfaceState_Statistics.Merge(dst, src)
}
func (m *InterfaceState_Statistics) XXX_Size() int {
	return xxx_messageInfo_InterfaceState_Statistics.Size(m)
}
func (m *InterfaceState_Statistics) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfaceState_Statistics.DiscardUnknown(m)
}

var xxx_messageInfo_InterfaceState_Statistics proto.InternalMessageInfo

func (m *InterfaceState_Statistics) GetInPackets() uint64 {
	if m != nil {
		return m.InPackets
	}
	return 0
}

func (m *InterfaceState_Statistics) GetInBytes() uint64 {
	if m != nil {
		return m.InBytes
	}
	return 0
}

func (m *InterfaceState_Statistics) GetOutPackets() uint64 {
	if m != nil {
		return m.OutPackets
	}
	return 0
}

func (m *InterfaceState_Statistics) GetOutBytes() uint64 {
	if m != nil {
		return m.OutBytes
	}
	return 0
}

func (m *InterfaceState_Statistics) GetDropPackets() uint64 {
	if m != nil {
		return m.DropPackets
	}
	return 0
}

func (m *InterfaceState_Statistics) GetPuntPackets() uint64 {
	if m != nil {
		return m.PuntPackets
	}
	return 0
}

func (m *InterfaceState_Statistics) GetIpv4Packets() uint64 {
	if m != nil {
		return m.Ipv4Packets
	}
	return 0
}

func (m *InterfaceState_Statistics) GetIpv6Packets() uint64 {
	if m != nil {
		return m.Ipv6Packets
	}
	return 0
}

func (m *InterfaceState_Statistics) GetInNobufPackets() uint64 {
	if m != nil {
		return m.InNobufPackets
	}
	return 0
}

func (m *InterfaceState_Statistics) GetInMissPackets() uint64 {
	if m != nil {
		return m.InMissPackets
	}
	return 0
}

func (m *InterfaceState_Statistics) GetInErrorPackets() uint64 {
	if m != nil {
		return m.InErrorPackets
	}
	return 0
}

func (m *InterfaceState_Statistics) GetOutErrorPackets() uint64 {
	if m != nil {
		return m.OutErrorPackets
	}
	return 0
}

func (*InterfaceState_Statistics) XXX_MessageName() string {
	return "vpp.interfaces.InterfaceState.Statistics"
}

type InterfaceNotification struct {
	Type                 InterfaceNotification_NotifType `protobuf:"varint,1,opt,name=type,proto3,enum=vpp.interfaces.InterfaceNotification_NotifType" json:"type,omitempty"`
	State                *InterfaceState                 `protobuf:"bytes,2,opt,name=state,proto3" json:"state,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *InterfaceNotification) Reset()         { *m = InterfaceNotification{} }
func (m *InterfaceNotification) String() string { return proto.CompactTextString(m) }
func (*InterfaceNotification) ProtoMessage()    {}
func (*InterfaceNotification) Descriptor() ([]byte, []int) {
	return fileDescriptor_state_bb2097df664231a9, []int{1}
}
func (m *InterfaceNotification) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InterfaceNotification.Unmarshal(m, b)
}
func (m *InterfaceNotification) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InterfaceNotification.Marshal(b, m, deterministic)
}
func (dst *InterfaceNotification) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InterfaceNotification.Merge(dst, src)
}
func (m *InterfaceNotification) XXX_Size() int {
	return xxx_messageInfo_InterfaceNotification.Size(m)
}
func (m *InterfaceNotification) XXX_DiscardUnknown() {
	xxx_messageInfo_InterfaceNotification.DiscardUnknown(m)
}

var xxx_messageInfo_InterfaceNotification proto.InternalMessageInfo

func (m *InterfaceNotification) GetType() InterfaceNotification_NotifType {
	if m != nil {
		return m.Type
	}
	return InterfaceNotification_UNKNOWN
}

func (m *InterfaceNotification) GetState() *InterfaceState {
	if m != nil {
		return m.State
	}
	return nil
}

func (*InterfaceNotification) XXX_MessageName() string {
	return "vpp.interfaces.InterfaceNotification"
}
func init() {
	proto.RegisterType((*InterfaceState)(nil), "vpp.interfaces.InterfaceState")
	proto.RegisterType((*InterfaceState_Statistics)(nil), "vpp.interfaces.InterfaceState.Statistics")
	proto.RegisterType((*InterfaceNotification)(nil), "vpp.interfaces.InterfaceNotification")
	proto.RegisterEnum("vpp.interfaces.InterfaceState_Status", InterfaceState_Status_name, InterfaceState_Status_value)
	proto.RegisterEnum("vpp.interfaces.InterfaceState_Duplex", InterfaceState_Duplex_name, InterfaceState_Duplex_value)
	proto.RegisterEnum("vpp.interfaces.InterfaceNotification_NotifType", InterfaceNotification_NotifType_name, InterfaceNotification_NotifType_value)
}

func init() {
	proto.RegisterFile("models/vpp/interfaces/state.proto", fileDescriptor_state_bb2097df664231a9)
}

var fileDescriptor_state_bb2097df664231a9 = []byte{
	// 749 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xe1, 0x6e, 0xe3, 0x44,
	0x10, 0xc7, 0xeb, 0xc4, 0x75, 0x92, 0x71, 0x9a, 0x33, 0x2b, 0x90, 0x4c, 0x11, 0x47, 0x1a, 0x74,
	0x28, 0x20, 0x9d, 0x8d, 0xca, 0xe9, 0xbe, 0x9c, 0xf8, 0xd0, 0x6b, 0x52, 0x5d, 0x21, 0xb8, 0x95,
	0x93, 0xe8, 0x10, 0x5f, 0x2c, 0x37, 0xde, 0xa4, 0x2b, 0x92, 0xdd, 0x95, 0x77, 0x5d, 0x5d, 0xdf,
	0x8c, 0x27, 0x40, 0x88, 0xd7, 0xe0, 0x45, 0xd0, 0x8e, 0x63, 0x27, 0x91, 0xee, 0xd0, 0xf1, 0x6d,
	0xe7, 0x3f, 0xbf, 0xff, 0x6c, 0x76, 0x3c, 0x19, 0x38, 0xdb, 0x88, 0x8c, 0xae, 0x55, 0xf8, 0x20,
	0x65, 0xc8, 0xb8, 0xa6, 0xf9, 0x32, 0x5d, 0x50, 0x15, 0x2a, 0x9d, 0x6a, 0x1a, 0xc8, 0x5c, 0x68,
	0x41, 0x7a, 0x0f, 0x52, 0x06, 0xbb, 0xdc, 0xe9, 0xf3, 0x15, 0xd3, 0xf7, 0xc5, 0x5d, 0xb0, 0x10,
	0x9b, 0x70, 0x25, 0x56, 0x22, 0x44, 0xec, 0xae, 0x58, 0x62, 0x84, 0x01, 0x9e, 0x4a, 0xfb, 0xe9,
	0xb3, 0xf7, 0xdf, 0x50, 0x1f, 0x4b, 0x6c, 0xf0, 0x77, 0x1b, 0x7a, 0xd7, 0x95, 0x36, 0x35, 0xd7,
	0x13, 0x02, 0x36, 0x4f, 0x37, 0xd4, 0xb7, 0xfa, 0xd6, 0xb0, 0x13, 0xe3, 0x99, 0x7c, 0x0d, 0x27,
	0xe8, 0xe4, 0xe9, 0x3a, 0xc1, 0x64, 0x03, 0x93, 0xdd, 0x4a, 0x8c, 0x0c, 0x74, 0x0e, 0xb6, 0x7e,
	0x94, 0xd4, 0x6f, 0xf6, 0xad, 0x61, 0xef, 0xfc, 0x69, 0x70, 0xf8, 0x80, 0xa0, 0xbe, 0x26, 0x98,
	0x3d, 0x4a, 0x1a, 0x23, 0x4b, 0x3e, 0x87, 0x36, 0x5b, 0x26, 0x8c, 0x67, 0xf4, 0x9d, 0x6f, 0xf7,
	0xad, 0xe1, 0x49, 0xdc, 0x62, 0xcb, 0x6b, 0x13, 0x92, 0x37, 0xd0, 0x4d, 0xb3, 0x0d, 0xe3, 0x89,
	0xe9, 0x4a, 0xa1, 0xfc, 0x63, 0x2c, 0xfb, 0xec, 0x83, 0x65, 0xf1, 0xd7, 0x07, 0x53, 0x84, 0x63,
	0x17, 0xad, 0x65, 0x40, 0xae, 0xc0, 0x15, 0x92, 0xe6, 0x55, 0x21, 0xe7, 0xff, 0x14, 0x02, 0xe3,
	0xdc, 0xd6, 0xf9, 0x0a, 0xdc, 0x75, 0xaa, 0x74, 0xb2, 0xb8, 0x4f, 0xf9, 0x8a, 0xfa, 0xad, 0xbe,
	0x35, 0x6c, 0xc6, 0x60, 0xa4, 0x4b, 0x54, 0xc8, 0x19, 0x74, 0xe5, 0xfd, 0xa3, 0x4a, 0xd2, 0x2c,
	0xcb, 0xa9, 0x52, 0x7e, 0x1b, 0xbb, 0xe4, 0x1a, 0xed, 0xa2, 0x94, 0xc8, 0xa7, 0x70, 0xac, 0x24,
	0xa5, 0x99, 0xdf, 0xe9, 0x5b, 0x43, 0x3b, 0x2e, 0x03, 0xe2, 0x41, 0x73, 0xa3, 0x0b, 0x1f, 0xb0,
	0x03, 0xe6, 0x48, 0x7e, 0x04, 0x27, 0x2b, 0xe4, 0x9a, 0xbe, 0xf3, 0xdd, 0x8f, 0xfa, 0xb9, 0x23,
	0x84, 0xe3, 0xad, 0x89, 0x5c, 0x03, 0x98, 0xd7, 0x32, 0xa5, 0xd9, 0x42, 0xf9, 0x59, 0xdf, 0x1a,
	0xba, 0xe7, 0xdf, 0x7e, 0xc4, 0x8b, 0x4b, 0x43, 0xbc, 0x67, 0x3e, 0xfd, 0xa3, 0x09, 0xb0, 0x4b,
	0x91, 0x2f, 0x01, 0x18, 0x4f, 0x64, 0xba, 0xf8, 0x9d, 0x6a, 0x85, 0x43, 0x62, 0xc7, 0x1d, 0xc6,
	0x6f, 0x4b, 0x01, 0x3f, 0x28, 0x4f, 0xee, 0x1e, 0x35, 0x55, 0x38, 0x24, 0x76, 0xdc, 0x62, 0xfc,
	0xb5, 0x09, 0x4d, 0xfb, 0x44, 0xa1, 0x6b, 0x6b, 0x13, 0xb3, 0x20, 0x0a, 0x5d, 0x79, 0xbf, 0x80,
	0x8e, 0x01, 0x4a, 0xb3, 0x8d, 0xe9, 0xb6, 0x28, 0x74, 0xe9, 0x3e, 0x83, 0x6e, 0x96, 0x0b, 0x59,
	0xdb, 0x8f, 0x31, 0xef, 0x1a, 0xad, 0xf2, 0x9b, 0xf6, 0x17, 0x7c, 0x77, 0x83, 0x53, 0x22, 0x46,
	0xdb, 0x43, 0x98, 0x7c, 0x78, 0x51, 0x23, 0xad, 0x12, 0x31, 0xda, 0x21, 0xf2, 0xb2, 0x46, 0xda,
	0x35, 0xf2, 0xb2, 0x42, 0x86, 0xe0, 0x31, 0x9e, 0x70, 0xf3, 0xef, 0xab, 0xb1, 0xf2, 0x7b, 0xf6,
	0x18, 0x8f, 0x8c, 0x5c, 0x91, 0xdf, 0xc0, 0x13, 0xc6, 0x93, 0x0d, 0x53, 0xaa, 0x06, 0x01, 0xc1,
	0x13, 0xc6, 0x7f, 0x61, 0x4a, 0x1d, 0x56, 0xa4, 0x79, 0x2e, 0xf2, 0x1a, 0x74, 0xab, 0x8a, 0x63,
	0x23, 0x57, 0xe4, 0x77, 0xf0, 0x89, 0x69, 0xd2, 0x21, 0xda, 0x45, 0xf4, 0x89, 0x28, 0xf4, 0x3e,
	0x3b, 0x78, 0x05, 0xce, 0x76, 0x74, 0x09, 0xf4, 0xe6, 0xd1, 0xcf, 0xd1, 0xcd, 0xdb, 0x28, 0x99,
	0xce, 0x2e, 0x66, 0xf3, 0xa9, 0x77, 0x44, 0x1c, 0x68, 0xcc, 0x6f, 0x3d, 0x8b, 0xb4, 0xc1, 0x1e,
	0xdd, 0xbc, 0x8d, 0xbc, 0x06, 0x71, 0xa1, 0x35, 0x1a, 0x4f, 0xc6, 0xb3, 0xf1, 0xc8, 0x6b, 0x0e,
	0xbe, 0x07, 0xa7, 0x1c, 0xaa, 0x7d, 0xf3, 0x68, 0x7e, 0x3b, 0x19, 0xff, 0xea, 0x1d, 0x19, 0xd3,
	0x9b, 0x8b, 0xc9, 0x55, 0x69, 0xbf, 0x9a, 0x4f, 0x26, 0x5e, 0x63, 0xf0, 0xa7, 0x05, 0x9f, 0xd5,
	0x33, 0x15, 0x09, 0xcd, 0x96, 0x6c, 0x91, 0x6a, 0x26, 0x38, 0xb9, 0xdc, 0xae, 0x06, 0x0b, 0x67,
	0x39, 0xfc, 0xe0, 0x20, 0xee, 0x9b, 0x02, 0x0c, 0xf6, 0x76, 0xc5, 0x0b, 0x38, 0xc6, 0x05, 0x89,
	0x73, 0xe5, 0xfe, 0xc7, 0x82, 0xc1, 0x71, 0x8e, 0x4b, 0x78, 0x70, 0x0e, 0x9d, 0xba, 0x90, 0x79,
	0xe0, 0xf6, 0x25, 0xde, 0x11, 0x01, 0x70, 0xe6, 0xb7, 0xf8, 0x72, 0x8b, 0x74, 0xa1, 0x7d, 0x79,
	0x33, 0x8f, 0x66, 0xe3, 0x78, 0xea, 0x35, 0x5e, 0xff, 0xf4, 0xd7, 0x3f, 0x4f, 0xad, 0xdf, 0x46,
	0x7b, 0x1b, 0x77, 0xcd, 0x56, 0xa9, 0x16, 0x66, 0x9b, 0x3e, 0x4f, 0x57, 0x94, 0xeb, 0x30, 0x95,
	0x2c, 0x7c, 0xef, 0x8a, 0x7d, 0xf5, 0x20, 0x65, 0xb2, 0x0b, 0xef, 0x1c, 0x5c, 0xb4, 0x3f, 0xfc,
	0x1b, 0x00, 0x00, 0xff, 0xff, 0x18, 0xae, 0x97, 0x14, 0xf3, 0x05, 0x00, 0x00,
}
