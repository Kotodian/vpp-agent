// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

// Package lcp contains generated bindings for API file lcp.api.
//
// Contents:
// -  1 enum
// - 15 messages
package lcp

import (
	"strconv"

	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
	interface_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2306/interface_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "lcp"
	APIVersion = "1.0.0"
	VersionCrc = 0x64780a3
)

// LcpItfHostType defines enum 'lcp_itf_host_type'.
type LcpItfHostType uint8

const (
	LCP_API_ITF_HOST_TAP LcpItfHostType = 0
	LCP_API_ITF_HOST_TUN LcpItfHostType = 1
)

var (
	LcpItfHostType_name = map[uint8]string{
		0: "LCP_API_ITF_HOST_TAP",
		1: "LCP_API_ITF_HOST_TUN",
	}
	LcpItfHostType_value = map[string]uint8{
		"LCP_API_ITF_HOST_TAP": 0,
		"LCP_API_ITF_HOST_TUN": 1,
	}
)

func (x LcpItfHostType) String() string {
	s, ok := LcpItfHostType_name[uint8(x)]
	if ok {
		return s
	}
	return "LcpItfHostType(" + strconv.Itoa(int(x)) + ")"
}

// get the default Linux Control Plane netns
// LcpDefaultNsGet defines message 'lcp_default_ns_get'.
type LcpDefaultNsGet struct{}

func (m *LcpDefaultNsGet) Reset()               { *m = LcpDefaultNsGet{} }
func (*LcpDefaultNsGet) GetMessageName() string { return "lcp_default_ns_get" }
func (*LcpDefaultNsGet) GetCrcString() string   { return "51077d14" }
func (*LcpDefaultNsGet) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *LcpDefaultNsGet) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *LcpDefaultNsGet) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *LcpDefaultNsGet) Unmarshal(b []byte) error {
	return nil
}

// get the default Linux Control Plane netns
//   - netns - the default netns; netns[0] == 0 if none
//
// LcpDefaultNsGetReply defines message 'lcp_default_ns_get_reply'.
// InProgress: the message form may change in the future versions
type LcpDefaultNsGetReply struct {
	Netns string `binapi:"string[32],name=netns" json:"netns,omitempty"`
}

func (m *LcpDefaultNsGetReply) Reset()               { *m = LcpDefaultNsGetReply{} }
func (*LcpDefaultNsGetReply) GetMessageName() string { return "lcp_default_ns_get_reply" }
func (*LcpDefaultNsGetReply) GetCrcString() string   { return "5102feee" }
func (*LcpDefaultNsGetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *LcpDefaultNsGetReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 32 // m.Netns
	return size
}
func (m *LcpDefaultNsGetReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.Netns, 32)
	return buf.Bytes(), nil
}
func (m *LcpDefaultNsGetReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Netns = buf.DecodeString(32)
	return nil
}

// Set the default Linux Control Plane netns
//   - netns - the new default netns; netns[0] == 0 if none
//
// LcpDefaultNsSet defines message 'lcp_default_ns_set'.
// InProgress: the message form may change in the future versions
type LcpDefaultNsSet struct {
	Netns string `binapi:"string[32],name=netns" json:"netns,omitempty"`
}

func (m *LcpDefaultNsSet) Reset()               { *m = LcpDefaultNsSet{} }
func (*LcpDefaultNsSet) GetMessageName() string { return "lcp_default_ns_set" }
func (*LcpDefaultNsSet) GetCrcString() string   { return "69749409" }
func (*LcpDefaultNsSet) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *LcpDefaultNsSet) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 32 // m.Netns
	return size
}
func (m *LcpDefaultNsSet) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeString(m.Netns, 32)
	return buf.Bytes(), nil
}
func (m *LcpDefaultNsSet) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Netns = buf.DecodeString(32)
	return nil
}

// LcpDefaultNsSetReply defines message 'lcp_default_ns_set_reply'.
// InProgress: the message form may change in the future versions
type LcpDefaultNsSetReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *LcpDefaultNsSetReply) Reset()               { *m = LcpDefaultNsSetReply{} }
func (*LcpDefaultNsSetReply) GetMessageName() string { return "lcp_default_ns_set_reply" }
func (*LcpDefaultNsSetReply) GetCrcString() string   { return "e8d4e804" }
func (*LcpDefaultNsSetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *LcpDefaultNsSetReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *LcpDefaultNsSetReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *LcpDefaultNsSetReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// Add or delete a Linux Conrol Plane interface pair
//   - is_add - 0 if deleting, != 0 if adding
//   - sw_if_index - index of VPP PHY SW interface
//   - host_if_name - host tap interface name
//   - host_if_type - the type of host interface to create (tun, tap)
//   - netns - optional tap netns; netns[0] == 0 if none
//
// LcpItfPairAddDel defines message 'lcp_itf_pair_add_del'.
// InProgress: the message form may change in the future versions
type LcpItfPairAddDel struct {
	IsAdd      bool                           `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	SwIfIndex  interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	HostIfName string                         `binapi:"string[16],name=host_if_name" json:"host_if_name,omitempty"`
	HostIfType LcpItfHostType                 `binapi:"lcp_itf_host_type,name=host_if_type" json:"host_if_type,omitempty"`
	Netns      string                         `binapi:"string[32],name=netns" json:"netns,omitempty"`
}

func (m *LcpItfPairAddDel) Reset()               { *m = LcpItfPairAddDel{} }
func (*LcpItfPairAddDel) GetMessageName() string { return "lcp_itf_pair_add_del" }
func (*LcpItfPairAddDel) GetCrcString() string   { return "40482b80" }
func (*LcpItfPairAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *LcpItfPairAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1  // m.IsAdd
	size += 4  // m.SwIfIndex
	size += 16 // m.HostIfName
	size += 1  // m.HostIfType
	size += 32 // m.Netns
	return size
}
func (m *LcpItfPairAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeString(m.HostIfName, 16)
	buf.EncodeUint8(uint8(m.HostIfType))
	buf.EncodeString(m.Netns, 32)
	return buf.Bytes(), nil
}
func (m *LcpItfPairAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.HostIfName = buf.DecodeString(16)
	m.HostIfType = LcpItfHostType(buf.DecodeUint8())
	m.Netns = buf.DecodeString(32)
	return nil
}

// LcpItfPairAddDelReply defines message 'lcp_itf_pair_add_del_reply'.
// InProgress: the message form may change in the future versions
type LcpItfPairAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *LcpItfPairAddDelReply) Reset()               { *m = LcpItfPairAddDelReply{} }
func (*LcpItfPairAddDelReply) GetMessageName() string { return "lcp_itf_pair_add_del_reply" }
func (*LcpItfPairAddDelReply) GetCrcString() string   { return "e8d4e804" }
func (*LcpItfPairAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *LcpItfPairAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *LcpItfPairAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *LcpItfPairAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// LcpItfPairAddDelV2 defines message 'lcp_itf_pair_add_del_v2'.
// InProgress: the message form may change in the future versions
type LcpItfPairAddDelV2 struct {
	IsAdd      bool                           `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	SwIfIndex  interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	HostIfName string                         `binapi:"string[16],name=host_if_name" json:"host_if_name,omitempty"`
	HostIfType LcpItfHostType                 `binapi:"lcp_itf_host_type,name=host_if_type" json:"host_if_type,omitempty"`
	Netns      string                         `binapi:"string[32],name=netns" json:"netns,omitempty"`
}

func (m *LcpItfPairAddDelV2) Reset()               { *m = LcpItfPairAddDelV2{} }
func (*LcpItfPairAddDelV2) GetMessageName() string { return "lcp_itf_pair_add_del_v2" }
func (*LcpItfPairAddDelV2) GetCrcString() string   { return "40482b80" }
func (*LcpItfPairAddDelV2) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *LcpItfPairAddDelV2) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1  // m.IsAdd
	size += 4  // m.SwIfIndex
	size += 16 // m.HostIfName
	size += 1  // m.HostIfType
	size += 32 // m.Netns
	return size
}
func (m *LcpItfPairAddDelV2) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeString(m.HostIfName, 16)
	buf.EncodeUint8(uint8(m.HostIfType))
	buf.EncodeString(m.Netns, 32)
	return buf.Bytes(), nil
}
func (m *LcpItfPairAddDelV2) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.HostIfName = buf.DecodeString(16)
	m.HostIfType = LcpItfHostType(buf.DecodeUint8())
	m.Netns = buf.DecodeString(32)
	return nil
}

// LcpItfPairAddDelV2Reply defines message 'lcp_itf_pair_add_del_v2_reply'.
type LcpItfPairAddDelV2Reply struct {
	Retval        int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	HostSwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=host_sw_if_index" json:"host_sw_if_index,omitempty"`
}

func (m *LcpItfPairAddDelV2Reply) Reset()               { *m = LcpItfPairAddDelV2Reply{} }
func (*LcpItfPairAddDelV2Reply) GetMessageName() string { return "lcp_itf_pair_add_del_v2_reply" }
func (*LcpItfPairAddDelV2Reply) GetCrcString() string   { return "39452f52" }
func (*LcpItfPairAddDelV2Reply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *LcpItfPairAddDelV2Reply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.HostSwIfIndex
	return size
}
func (m *LcpItfPairAddDelV2Reply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.HostSwIfIndex))
	return buf.Bytes(), nil
}
func (m *LcpItfPairAddDelV2Reply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.HostSwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// Linux Control Plane interface pair dump response
//   - phy_sw_if_index - VPP's sw_if_index for the PHY
//   - host_sw_if_index - VPP's sw_if_index for the host tap
//   - vif_index - tap linux index
//   - host_if_name - host interface name
//   - host_if_type - host interface type (tun, tap)
//   - netns - host interface netns
//
// LcpItfPairDetails defines message 'lcp_itf_pair_details'.
// InProgress: the message form may change in the future versions
type LcpItfPairDetails struct {
	PhySwIfIndex  interface_types.InterfaceIndex `binapi:"interface_index,name=phy_sw_if_index" json:"phy_sw_if_index,omitempty"`
	HostSwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=host_sw_if_index" json:"host_sw_if_index,omitempty"`
	VifIndex      uint32                         `binapi:"u32,name=vif_index" json:"vif_index,omitempty"`
	HostIfName    string                         `binapi:"string[16],name=host_if_name" json:"host_if_name,omitempty"`
	HostIfType    LcpItfHostType                 `binapi:"lcp_itf_host_type,name=host_if_type" json:"host_if_type,omitempty"`
	Netns         string                         `binapi:"string[32],name=netns" json:"netns,omitempty"`
}

func (m *LcpItfPairDetails) Reset()               { *m = LcpItfPairDetails{} }
func (*LcpItfPairDetails) GetMessageName() string { return "lcp_itf_pair_details" }
func (*LcpItfPairDetails) GetCrcString() string   { return "8b5481af" }
func (*LcpItfPairDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *LcpItfPairDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4  // m.PhySwIfIndex
	size += 4  // m.HostSwIfIndex
	size += 4  // m.VifIndex
	size += 16 // m.HostIfName
	size += 1  // m.HostIfType
	size += 32 // m.Netns
	return size
}
func (m *LcpItfPairDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.PhySwIfIndex))
	buf.EncodeUint32(uint32(m.HostSwIfIndex))
	buf.EncodeUint32(m.VifIndex)
	buf.EncodeString(m.HostIfName, 16)
	buf.EncodeUint8(uint8(m.HostIfType))
	buf.EncodeString(m.Netns, 32)
	return buf.Bytes(), nil
}
func (m *LcpItfPairDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.PhySwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.HostSwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.VifIndex = buf.DecodeUint32()
	m.HostIfName = buf.DecodeString(16)
	m.HostIfType = LcpItfHostType(buf.DecodeUint8())
	m.Netns = buf.DecodeString(32)
	return nil
}

// Dump Linux Control Plane interface pair data
//   - sw_if_index - interface to use as filter (~0 == "all")
//
// LcpItfPairGet defines message 'lcp_itf_pair_get'.
type LcpItfPairGet struct {
	Cursor uint32 `binapi:"u32,name=cursor" json:"cursor,omitempty"`
}

func (m *LcpItfPairGet) Reset()               { *m = LcpItfPairGet{} }
func (*LcpItfPairGet) GetMessageName() string { return "lcp_itf_pair_get" }
func (*LcpItfPairGet) GetCrcString() string   { return "f75ba505" }
func (*LcpItfPairGet) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *LcpItfPairGet) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Cursor
	return size
}
func (m *LcpItfPairGet) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Cursor)
	return buf.Bytes(), nil
}
func (m *LcpItfPairGet) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Cursor = buf.DecodeUint32()
	return nil
}

// LcpItfPairGetReply defines message 'lcp_itf_pair_get_reply'.
type LcpItfPairGetReply struct {
	Retval int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	Cursor uint32 `binapi:"u32,name=cursor" json:"cursor,omitempty"`
}

func (m *LcpItfPairGetReply) Reset()               { *m = LcpItfPairGetReply{} }
func (*LcpItfPairGetReply) GetMessageName() string { return "lcp_itf_pair_get_reply" }
func (*LcpItfPairGetReply) GetCrcString() string   { return "53b48f5d" }
func (*LcpItfPairGetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *LcpItfPairGetReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.Cursor
	return size
}
func (m *LcpItfPairGetReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.Cursor)
	return buf.Bytes(), nil
}
func (m *LcpItfPairGetReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.Cursor = buf.DecodeUint32()
	return nil
}

// Replace end/begin
// LcpItfPairReplaceBegin defines message 'lcp_itf_pair_replace_begin'.
type LcpItfPairReplaceBegin struct{}

func (m *LcpItfPairReplaceBegin) Reset()               { *m = LcpItfPairReplaceBegin{} }
func (*LcpItfPairReplaceBegin) GetMessageName() string { return "lcp_itf_pair_replace_begin" }
func (*LcpItfPairReplaceBegin) GetCrcString() string   { return "51077d14" }
func (*LcpItfPairReplaceBegin) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *LcpItfPairReplaceBegin) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *LcpItfPairReplaceBegin) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *LcpItfPairReplaceBegin) Unmarshal(b []byte) error {
	return nil
}

// LcpItfPairReplaceBeginReply defines message 'lcp_itf_pair_replace_begin_reply'.
type LcpItfPairReplaceBeginReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *LcpItfPairReplaceBeginReply) Reset() { *m = LcpItfPairReplaceBeginReply{} }
func (*LcpItfPairReplaceBeginReply) GetMessageName() string {
	return "lcp_itf_pair_replace_begin_reply"
}
func (*LcpItfPairReplaceBeginReply) GetCrcString() string { return "e8d4e804" }
func (*LcpItfPairReplaceBeginReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *LcpItfPairReplaceBeginReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *LcpItfPairReplaceBeginReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *LcpItfPairReplaceBeginReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// LcpItfPairReplaceEnd defines message 'lcp_itf_pair_replace_end'.
type LcpItfPairReplaceEnd struct{}

func (m *LcpItfPairReplaceEnd) Reset()               { *m = LcpItfPairReplaceEnd{} }
func (*LcpItfPairReplaceEnd) GetMessageName() string { return "lcp_itf_pair_replace_end" }
func (*LcpItfPairReplaceEnd) GetCrcString() string   { return "51077d14" }
func (*LcpItfPairReplaceEnd) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *LcpItfPairReplaceEnd) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *LcpItfPairReplaceEnd) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *LcpItfPairReplaceEnd) Unmarshal(b []byte) error {
	return nil
}

// LcpItfPairReplaceEndReply defines message 'lcp_itf_pair_replace_end_reply'.
type LcpItfPairReplaceEndReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *LcpItfPairReplaceEndReply) Reset()               { *m = LcpItfPairReplaceEndReply{} }
func (*LcpItfPairReplaceEndReply) GetMessageName() string { return "lcp_itf_pair_replace_end_reply" }
func (*LcpItfPairReplaceEndReply) GetCrcString() string   { return "e8d4e804" }
func (*LcpItfPairReplaceEndReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *LcpItfPairReplaceEndReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *LcpItfPairReplaceEndReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *LcpItfPairReplaceEndReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_lcp_binapi_init() }
func file_lcp_binapi_init() {
	api.RegisterMessage((*LcpDefaultNsGet)(nil), "lcp_default_ns_get_51077d14")
	api.RegisterMessage((*LcpDefaultNsGetReply)(nil), "lcp_default_ns_get_reply_5102feee")
	api.RegisterMessage((*LcpDefaultNsSet)(nil), "lcp_default_ns_set_69749409")
	api.RegisterMessage((*LcpDefaultNsSetReply)(nil), "lcp_default_ns_set_reply_e8d4e804")
	api.RegisterMessage((*LcpItfPairAddDel)(nil), "lcp_itf_pair_add_del_40482b80")
	api.RegisterMessage((*LcpItfPairAddDelReply)(nil), "lcp_itf_pair_add_del_reply_e8d4e804")
	api.RegisterMessage((*LcpItfPairAddDelV2)(nil), "lcp_itf_pair_add_del_v2_40482b80")
	api.RegisterMessage((*LcpItfPairAddDelV2Reply)(nil), "lcp_itf_pair_add_del_v2_reply_39452f52")
	api.RegisterMessage((*LcpItfPairDetails)(nil), "lcp_itf_pair_details_8b5481af")
	api.RegisterMessage((*LcpItfPairGet)(nil), "lcp_itf_pair_get_f75ba505")
	api.RegisterMessage((*LcpItfPairGetReply)(nil), "lcp_itf_pair_get_reply_53b48f5d")
	api.RegisterMessage((*LcpItfPairReplaceBegin)(nil), "lcp_itf_pair_replace_begin_51077d14")
	api.RegisterMessage((*LcpItfPairReplaceBeginReply)(nil), "lcp_itf_pair_replace_begin_reply_e8d4e804")
	api.RegisterMessage((*LcpItfPairReplaceEnd)(nil), "lcp_itf_pair_replace_end_51077d14")
	api.RegisterMessage((*LcpItfPairReplaceEndReply)(nil), "lcp_itf_pair_replace_end_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*LcpDefaultNsGet)(nil),
		(*LcpDefaultNsGetReply)(nil),
		(*LcpDefaultNsSet)(nil),
		(*LcpDefaultNsSetReply)(nil),
		(*LcpItfPairAddDel)(nil),
		(*LcpItfPairAddDelReply)(nil),
		(*LcpItfPairAddDelV2)(nil),
		(*LcpItfPairAddDelV2Reply)(nil),
		(*LcpItfPairDetails)(nil),
		(*LcpItfPairGet)(nil),
		(*LcpItfPairGetReply)(nil),
		(*LcpItfPairReplaceBegin)(nil),
		(*LcpItfPairReplaceBeginReply)(nil),
		(*LcpItfPairReplaceEnd)(nil),
		(*LcpItfPairReplaceEndReply)(nil),
	}
}
