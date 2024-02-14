// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

// Package lisp_gpe contains generated bindings for API file lisp_gpe.api.
//
// Contents:
// -  3 structs
// - 20 messages
package lisp_gpe

import (
	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
	_ "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/ethernet_types"
	interface_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/interface_types"
	ip_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/ip_types"
	lisp_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/lisp_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "lisp_gpe"
	APIVersion = "2.0.0"
	VersionCrc = 0x92611b0
)

// GpeFwdEntry defines type 'gpe_fwd_entry'.
type GpeFwdEntry struct {
	FwdEntryIndex uint32         `binapi:"u32,name=fwd_entry_index" json:"fwd_entry_index,omitempty"`
	DpTable       uint32         `binapi:"u32,name=dp_table" json:"dp_table,omitempty"`
	Leid          lisp_types.Eid `binapi:"eid,name=leid" json:"leid,omitempty"`
	Reid          lisp_types.Eid `binapi:"eid,name=reid" json:"reid,omitempty"`
	Vni           uint32         `binapi:"u32,name=vni" json:"vni,omitempty"`
	Action        uint8          `binapi:"u8,name=action" json:"action,omitempty"`
}

// GpeLocator defines type 'gpe_locator'.
type GpeLocator struct {
	Weight uint8            `binapi:"u8,name=weight" json:"weight,omitempty"`
	Addr   ip_types.Address `binapi:"address,name=addr" json:"addr,omitempty"`
}

// GpeNativeFwdRpath defines type 'gpe_native_fwd_rpath'.
type GpeNativeFwdRpath struct {
	FibIndex    uint32                         `binapi:"u32,name=fib_index" json:"fib_index,omitempty"`
	NhSwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=nh_sw_if_index" json:"nh_sw_if_index,omitempty"`
	NhAddr      ip_types.Address               `binapi:"address,name=nh_addr" json:"nh_addr,omitempty"`
}

// add or delete GPE tunnel
//   - is_add - add address if non-zero, else delete
//   - rmt_eid - remote eid
//   - lcl_eid - local eid
//   - vni - virtual network identifier
//   - dp_table - vrf/bridge domain id
//   - action - negative action when 0 locators configured
//   - loc_num - number of locators
//   - locs - array of remote locators
//
// GpeAddDelFwdEntry defines message 'gpe_add_del_fwd_entry'.
type GpeAddDelFwdEntry struct {
	IsAdd   bool           `binapi:"bool,name=is_add,default=true" json:"is_add,omitempty"`
	RmtEid  lisp_types.Eid `binapi:"eid,name=rmt_eid" json:"rmt_eid,omitempty"`
	LclEid  lisp_types.Eid `binapi:"eid,name=lcl_eid" json:"lcl_eid,omitempty"`
	Vni     uint32         `binapi:"u32,name=vni" json:"vni,omitempty"`
	DpTable uint32         `binapi:"u32,name=dp_table" json:"dp_table,omitempty"`
	Action  uint8          `binapi:"u8,name=action" json:"action,omitempty"`
	LocNum  uint32         `binapi:"u32,name=loc_num" json:"-"`
	Locs    []GpeLocator   `binapi:"gpe_locator[loc_num],name=locs" json:"locs,omitempty"`
}

func (m *GpeAddDelFwdEntry) Reset()               { *m = GpeAddDelFwdEntry{} }
func (*GpeAddDelFwdEntry) GetMessageName() string { return "gpe_add_del_fwd_entry" }
func (*GpeAddDelFwdEntry) GetCrcString() string   { return "f0847644" }
func (*GpeAddDelFwdEntry) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GpeAddDelFwdEntry) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.IsAdd
	size += 1      // m.RmtEid.Type
	size += 1 * 18 // m.RmtEid.Address
	size += 1      // m.LclEid.Type
	size += 1 * 18 // m.LclEid.Address
	size += 4      // m.Vni
	size += 4      // m.DpTable
	size += 1      // m.Action
	size += 4      // m.LocNum
	for j1 := 0; j1 < len(m.Locs); j1++ {
		var s1 GpeLocator
		_ = s1
		if j1 < len(m.Locs) {
			s1 = m.Locs[j1]
		}
		size += 1      // s1.Weight
		size += 1      // s1.Addr.Af
		size += 1 * 16 // s1.Addr.Un
	}
	return size
}
func (m *GpeAddDelFwdEntry) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint8(uint8(m.RmtEid.Type))
	buf.EncodeBytes(m.RmtEid.Address.XXX_UnionData[:], 18)
	buf.EncodeUint8(uint8(m.LclEid.Type))
	buf.EncodeBytes(m.LclEid.Address.XXX_UnionData[:], 18)
	buf.EncodeUint32(m.Vni)
	buf.EncodeUint32(m.DpTable)
	buf.EncodeUint8(m.Action)
	buf.EncodeUint32(uint32(len(m.Locs)))
	for j0 := 0; j0 < len(m.Locs); j0++ {
		var v0 GpeLocator // Locs
		if j0 < len(m.Locs) {
			v0 = m.Locs[j0]
		}
		buf.EncodeUint8(v0.Weight)
		buf.EncodeUint8(uint8(v0.Addr.Af))
		buf.EncodeBytes(v0.Addr.Un.XXX_UnionData[:], 16)
	}
	return buf.Bytes(), nil
}
func (m *GpeAddDelFwdEntry) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.RmtEid.Type = lisp_types.EidType(buf.DecodeUint8())
	copy(m.RmtEid.Address.XXX_UnionData[:], buf.DecodeBytes(18))
	m.LclEid.Type = lisp_types.EidType(buf.DecodeUint8())
	copy(m.LclEid.Address.XXX_UnionData[:], buf.DecodeBytes(18))
	m.Vni = buf.DecodeUint32()
	m.DpTable = buf.DecodeUint32()
	m.Action = buf.DecodeUint8()
	m.LocNum = buf.DecodeUint32()
	m.Locs = make([]GpeLocator, m.LocNum)
	for j0 := 0; j0 < len(m.Locs); j0++ {
		m.Locs[j0].Weight = buf.DecodeUint8()
		m.Locs[j0].Addr.Af = ip_types.AddressFamily(buf.DecodeUint8())
		copy(m.Locs[j0].Addr.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	}
	return nil
}

// GpeAddDelFwdEntryReply defines message 'gpe_add_del_fwd_entry_reply'.
type GpeAddDelFwdEntryReply struct {
	Retval        int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	FwdEntryIndex uint32 `binapi:"u32,name=fwd_entry_index" json:"fwd_entry_index,omitempty"`
}

func (m *GpeAddDelFwdEntryReply) Reset()               { *m = GpeAddDelFwdEntryReply{} }
func (*GpeAddDelFwdEntryReply) GetMessageName() string { return "gpe_add_del_fwd_entry_reply" }
func (*GpeAddDelFwdEntryReply) GetCrcString() string   { return "efe5f176" }
func (*GpeAddDelFwdEntryReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GpeAddDelFwdEntryReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.FwdEntryIndex
	return size
}
func (m *GpeAddDelFwdEntryReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.FwdEntryIndex)
	return buf.Bytes(), nil
}
func (m *GpeAddDelFwdEntryReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.FwdEntryIndex = buf.DecodeUint32()
	return nil
}

// add or delete gpe_iface
//   - is_add - add address if non-zero, else delete
//
// GpeAddDelIface defines message 'gpe_add_del_iface'.
type GpeAddDelIface struct {
	IsAdd   bool   `binapi:"bool,name=is_add,default=true" json:"is_add,omitempty"`
	IsL2    bool   `binapi:"bool,name=is_l2" json:"is_l2,omitempty"`
	DpTable uint32 `binapi:"u32,name=dp_table" json:"dp_table,omitempty"`
	Vni     uint32 `binapi:"u32,name=vni" json:"vni,omitempty"`
}

func (m *GpeAddDelIface) Reset()               { *m = GpeAddDelIface{} }
func (*GpeAddDelIface) GetMessageName() string { return "gpe_add_del_iface" }
func (*GpeAddDelIface) GetCrcString() string   { return "3ccff273" }
func (*GpeAddDelIface) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GpeAddDelIface) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsAdd
	size += 1 // m.IsL2
	size += 4 // m.DpTable
	size += 4 // m.Vni
	return size
}
func (m *GpeAddDelIface) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeBool(m.IsL2)
	buf.EncodeUint32(m.DpTable)
	buf.EncodeUint32(m.Vni)
	return buf.Bytes(), nil
}
func (m *GpeAddDelIface) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.IsL2 = buf.DecodeBool()
	m.DpTable = buf.DecodeUint32()
	m.Vni = buf.DecodeUint32()
	return nil
}

// GpeAddDelIfaceReply defines message 'gpe_add_del_iface_reply'.
type GpeAddDelIfaceReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *GpeAddDelIfaceReply) Reset()               { *m = GpeAddDelIfaceReply{} }
func (*GpeAddDelIfaceReply) GetMessageName() string { return "gpe_add_del_iface_reply" }
func (*GpeAddDelIfaceReply) GetCrcString() string   { return "e8d4e804" }
func (*GpeAddDelIfaceReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GpeAddDelIfaceReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *GpeAddDelIfaceReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *GpeAddDelIfaceReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// Add native fwd rpath
//   - retval - return code
//   - is_add - flag to indicate add or del
//   - table_id - table id for route path
//   - nh_sw_if_index - next-hop sw_if_index (~0 if not set)
//   - is_ip4 - flag to indicate if nh is ip4
//   - nh_addr - next hop ip address
//
// GpeAddDelNativeFwdRpath defines message 'gpe_add_del_native_fwd_rpath'.
type GpeAddDelNativeFwdRpath struct {
	IsAdd       bool                           `binapi:"bool,name=is_add,default=true" json:"is_add,omitempty"`
	TableID     uint32                         `binapi:"u32,name=table_id" json:"table_id,omitempty"`
	NhSwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=nh_sw_if_index" json:"nh_sw_if_index,omitempty"`
	NhAddr      ip_types.Address               `binapi:"address,name=nh_addr" json:"nh_addr,omitempty"`
}

func (m *GpeAddDelNativeFwdRpath) Reset()               { *m = GpeAddDelNativeFwdRpath{} }
func (*GpeAddDelNativeFwdRpath) GetMessageName() string { return "gpe_add_del_native_fwd_rpath" }
func (*GpeAddDelNativeFwdRpath) GetCrcString() string   { return "43fc8b54" }
func (*GpeAddDelNativeFwdRpath) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GpeAddDelNativeFwdRpath) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.IsAdd
	size += 4      // m.TableID
	size += 4      // m.NhSwIfIndex
	size += 1      // m.NhAddr.Af
	size += 1 * 16 // m.NhAddr.Un
	return size
}
func (m *GpeAddDelNativeFwdRpath) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint32(m.TableID)
	buf.EncodeUint32(uint32(m.NhSwIfIndex))
	buf.EncodeUint8(uint8(m.NhAddr.Af))
	buf.EncodeBytes(m.NhAddr.Un.XXX_UnionData[:], 16)
	return buf.Bytes(), nil
}
func (m *GpeAddDelNativeFwdRpath) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.TableID = buf.DecodeUint32()
	m.NhSwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	m.NhAddr.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.NhAddr.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	return nil
}

// GpeAddDelNativeFwdRpathReply defines message 'gpe_add_del_native_fwd_rpath_reply'.
type GpeAddDelNativeFwdRpathReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *GpeAddDelNativeFwdRpathReply) Reset() { *m = GpeAddDelNativeFwdRpathReply{} }
func (*GpeAddDelNativeFwdRpathReply) GetMessageName() string {
	return "gpe_add_del_native_fwd_rpath_reply"
}
func (*GpeAddDelNativeFwdRpathReply) GetCrcString() string { return "e8d4e804" }
func (*GpeAddDelNativeFwdRpathReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GpeAddDelNativeFwdRpathReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *GpeAddDelNativeFwdRpathReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *GpeAddDelNativeFwdRpathReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// enable or disable gpe protocol
//   - is_enable [default=true] - enable protocol if non-zero, else disable
//
// GpeEnableDisable defines message 'gpe_enable_disable'.
type GpeEnableDisable struct {
	IsEnable bool `binapi:"bool,name=is_enable,default=true" json:"is_enable,omitempty"`
}

func (m *GpeEnableDisable) Reset()               { *m = GpeEnableDisable{} }
func (*GpeEnableDisable) GetMessageName() string { return "gpe_enable_disable" }
func (*GpeEnableDisable) GetCrcString() string   { return "c264d7bf" }
func (*GpeEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GpeEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsEnable
	return size
}
func (m *GpeEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsEnable)
	return buf.Bytes(), nil
}
func (m *GpeEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsEnable = buf.DecodeBool()
	return nil
}

// GpeEnableDisableReply defines message 'gpe_enable_disable_reply'.
type GpeEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *GpeEnableDisableReply) Reset()               { *m = GpeEnableDisableReply{} }
func (*GpeEnableDisableReply) GetMessageName() string { return "gpe_enable_disable_reply" }
func (*GpeEnableDisableReply) GetCrcString() string   { return "e8d4e804" }
func (*GpeEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GpeEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *GpeEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *GpeEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// GpeFwdEntriesGet defines message 'gpe_fwd_entries_get'.
type GpeFwdEntriesGet struct {
	Vni uint32 `binapi:"u32,name=vni" json:"vni,omitempty"`
}

func (m *GpeFwdEntriesGet) Reset()               { *m = GpeFwdEntriesGet{} }
func (*GpeFwdEntriesGet) GetMessageName() string { return "gpe_fwd_entries_get" }
func (*GpeFwdEntriesGet) GetCrcString() string   { return "8d1f2fe9" }
func (*GpeFwdEntriesGet) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GpeFwdEntriesGet) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Vni
	return size
}
func (m *GpeFwdEntriesGet) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.Vni)
	return buf.Bytes(), nil
}
func (m *GpeFwdEntriesGet) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Vni = buf.DecodeUint32()
	return nil
}

// GpeFwdEntriesGetReply defines message 'gpe_fwd_entries_get_reply'.
type GpeFwdEntriesGetReply struct {
	Retval  int32         `binapi:"i32,name=retval" json:"retval,omitempty"`
	Count   uint32        `binapi:"u32,name=count" json:"-"`
	Entries []GpeFwdEntry `binapi:"gpe_fwd_entry[count],name=entries" json:"entries,omitempty"`
}

func (m *GpeFwdEntriesGetReply) Reset()               { *m = GpeFwdEntriesGetReply{} }
func (*GpeFwdEntriesGetReply) GetMessageName() string { return "gpe_fwd_entries_get_reply" }
func (*GpeFwdEntriesGetReply) GetCrcString() string   { return "c4844876" }
func (*GpeFwdEntriesGetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GpeFwdEntriesGetReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.Count
	for j1 := 0; j1 < len(m.Entries); j1++ {
		var s1 GpeFwdEntry
		_ = s1
		if j1 < len(m.Entries) {
			s1 = m.Entries[j1]
		}
		size += 4      // s1.FwdEntryIndex
		size += 4      // s1.DpTable
		size += 1      // s1.Leid.Type
		size += 1 * 18 // s1.Leid.Address
		size += 1      // s1.Reid.Type
		size += 1 * 18 // s1.Reid.Address
		size += 4      // s1.Vni
		size += 1      // s1.Action
	}
	return size
}
func (m *GpeFwdEntriesGetReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(len(m.Entries)))
	for j0 := 0; j0 < len(m.Entries); j0++ {
		var v0 GpeFwdEntry // Entries
		if j0 < len(m.Entries) {
			v0 = m.Entries[j0]
		}
		buf.EncodeUint32(v0.FwdEntryIndex)
		buf.EncodeUint32(v0.DpTable)
		buf.EncodeUint8(uint8(v0.Leid.Type))
		buf.EncodeBytes(v0.Leid.Address.XXX_UnionData[:], 18)
		buf.EncodeUint8(uint8(v0.Reid.Type))
		buf.EncodeBytes(v0.Reid.Address.XXX_UnionData[:], 18)
		buf.EncodeUint32(v0.Vni)
		buf.EncodeUint8(v0.Action)
	}
	return buf.Bytes(), nil
}
func (m *GpeFwdEntriesGetReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.Count = buf.DecodeUint32()
	m.Entries = make([]GpeFwdEntry, m.Count)
	for j0 := 0; j0 < len(m.Entries); j0++ {
		m.Entries[j0].FwdEntryIndex = buf.DecodeUint32()
		m.Entries[j0].DpTable = buf.DecodeUint32()
		m.Entries[j0].Leid.Type = lisp_types.EidType(buf.DecodeUint8())
		copy(m.Entries[j0].Leid.Address.XXX_UnionData[:], buf.DecodeBytes(18))
		m.Entries[j0].Reid.Type = lisp_types.EidType(buf.DecodeUint8())
		copy(m.Entries[j0].Reid.Address.XXX_UnionData[:], buf.DecodeBytes(18))
		m.Entries[j0].Vni = buf.DecodeUint32()
		m.Entries[j0].Action = buf.DecodeUint8()
	}
	return nil
}

// GpeFwdEntryPathDetails defines message 'gpe_fwd_entry_path_details'.
type GpeFwdEntryPathDetails struct {
	LclLoc GpeLocator `binapi:"gpe_locator,name=lcl_loc" json:"lcl_loc,omitempty"`
	RmtLoc GpeLocator `binapi:"gpe_locator,name=rmt_loc" json:"rmt_loc,omitempty"`
}

func (m *GpeFwdEntryPathDetails) Reset()               { *m = GpeFwdEntryPathDetails{} }
func (*GpeFwdEntryPathDetails) GetMessageName() string { return "gpe_fwd_entry_path_details" }
func (*GpeFwdEntryPathDetails) GetCrcString() string   { return "483df51a" }
func (*GpeFwdEntryPathDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GpeFwdEntryPathDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1      // m.LclLoc.Weight
	size += 1      // m.LclLoc.Addr.Af
	size += 1 * 16 // m.LclLoc.Addr.Un
	size += 1      // m.RmtLoc.Weight
	size += 1      // m.RmtLoc.Addr.Af
	size += 1 * 16 // m.RmtLoc.Addr.Un
	return size
}
func (m *GpeFwdEntryPathDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint8(m.LclLoc.Weight)
	buf.EncodeUint8(uint8(m.LclLoc.Addr.Af))
	buf.EncodeBytes(m.LclLoc.Addr.Un.XXX_UnionData[:], 16)
	buf.EncodeUint8(m.RmtLoc.Weight)
	buf.EncodeUint8(uint8(m.RmtLoc.Addr.Af))
	buf.EncodeBytes(m.RmtLoc.Addr.Un.XXX_UnionData[:], 16)
	return buf.Bytes(), nil
}
func (m *GpeFwdEntryPathDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.LclLoc.Weight = buf.DecodeUint8()
	m.LclLoc.Addr.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.LclLoc.Addr.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	m.RmtLoc.Weight = buf.DecodeUint8()
	m.RmtLoc.Addr.Af = ip_types.AddressFamily(buf.DecodeUint8())
	copy(m.RmtLoc.Addr.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	return nil
}

// GpeFwdEntryPathDump defines message 'gpe_fwd_entry_path_dump'.
type GpeFwdEntryPathDump struct {
	FwdEntryIndex uint32 `binapi:"u32,name=fwd_entry_index" json:"fwd_entry_index,omitempty"`
}

func (m *GpeFwdEntryPathDump) Reset()               { *m = GpeFwdEntryPathDump{} }
func (*GpeFwdEntryPathDump) GetMessageName() string { return "gpe_fwd_entry_path_dump" }
func (*GpeFwdEntryPathDump) GetCrcString() string   { return "39bce980" }
func (*GpeFwdEntryPathDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GpeFwdEntryPathDump) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.FwdEntryIndex
	return size
}
func (m *GpeFwdEntryPathDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.FwdEntryIndex)
	return buf.Bytes(), nil
}
func (m *GpeFwdEntryPathDump) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.FwdEntryIndex = buf.DecodeUint32()
	return nil
}

// GpeFwdEntryVnisGet defines message 'gpe_fwd_entry_vnis_get'.
type GpeFwdEntryVnisGet struct{}

func (m *GpeFwdEntryVnisGet) Reset()               { *m = GpeFwdEntryVnisGet{} }
func (*GpeFwdEntryVnisGet) GetMessageName() string { return "gpe_fwd_entry_vnis_get" }
func (*GpeFwdEntryVnisGet) GetCrcString() string   { return "51077d14" }
func (*GpeFwdEntryVnisGet) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GpeFwdEntryVnisGet) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *GpeFwdEntryVnisGet) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *GpeFwdEntryVnisGet) Unmarshal(b []byte) error {
	return nil
}

// GpeFwdEntryVnisGetReply defines message 'gpe_fwd_entry_vnis_get_reply'.
type GpeFwdEntryVnisGetReply struct {
	Retval int32    `binapi:"i32,name=retval" json:"retval,omitempty"`
	Count  uint32   `binapi:"u32,name=count" json:"-"`
	Vnis   []uint32 `binapi:"u32[count],name=vnis" json:"vnis,omitempty"`
}

func (m *GpeFwdEntryVnisGetReply) Reset()               { *m = GpeFwdEntryVnisGetReply{} }
func (*GpeFwdEntryVnisGetReply) GetMessageName() string { return "gpe_fwd_entry_vnis_get_reply" }
func (*GpeFwdEntryVnisGetReply) GetCrcString() string   { return "aa70da20" }
func (*GpeFwdEntryVnisGetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GpeFwdEntryVnisGetReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4               // m.Retval
	size += 4               // m.Count
	size += 4 * len(m.Vnis) // m.Vnis
	return size
}
func (m *GpeFwdEntryVnisGetReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(len(m.Vnis)))
	for i := 0; i < len(m.Vnis); i++ {
		var x uint32
		if i < len(m.Vnis) {
			x = uint32(m.Vnis[i])
		}
		buf.EncodeUint32(x)
	}
	return buf.Bytes(), nil
}
func (m *GpeFwdEntryVnisGetReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.Count = buf.DecodeUint32()
	m.Vnis = make([]uint32, m.Count)
	for i := 0; i < len(m.Vnis); i++ {
		m.Vnis[i] = buf.DecodeUint32()
	}
	return nil
}

// get GPE encapsulation mode
//   - mode - LISP (value 0) or VXLAN (value 1)
//
// GpeGetEncapMode defines message 'gpe_get_encap_mode'.
type GpeGetEncapMode struct{}

func (m *GpeGetEncapMode) Reset()               { *m = GpeGetEncapMode{} }
func (*GpeGetEncapMode) GetMessageName() string { return "gpe_get_encap_mode" }
func (*GpeGetEncapMode) GetCrcString() string   { return "51077d14" }
func (*GpeGetEncapMode) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GpeGetEncapMode) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *GpeGetEncapMode) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *GpeGetEncapMode) Unmarshal(b []byte) error {
	return nil
}

// Reply for set_encap_mode
//   - retval - return code
//   - encap_mode - GPE encapsulation mode
//
// GpeGetEncapModeReply defines message 'gpe_get_encap_mode_reply'.
type GpeGetEncapModeReply struct {
	Retval    int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
	EncapMode uint8 `binapi:"u8,name=encap_mode" json:"encap_mode,omitempty"`
}

func (m *GpeGetEncapModeReply) Reset()               { *m = GpeGetEncapModeReply{} }
func (*GpeGetEncapModeReply) GetMessageName() string { return "gpe_get_encap_mode_reply" }
func (*GpeGetEncapModeReply) GetCrcString() string   { return "36e3f7ca" }
func (*GpeGetEncapModeReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GpeGetEncapModeReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 1 // m.EncapMode
	return size
}
func (m *GpeGetEncapModeReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint8(m.EncapMode)
	return buf.Bytes(), nil
}
func (m *GpeGetEncapModeReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.EncapMode = buf.DecodeUint8()
	return nil
}

// get GPE native fwd rpath
// GpeNativeFwdRpathsGet defines message 'gpe_native_fwd_rpaths_get'.
type GpeNativeFwdRpathsGet struct {
	IsIP4 bool `binapi:"bool,name=is_ip4" json:"is_ip4,omitempty"`
}

func (m *GpeNativeFwdRpathsGet) Reset()               { *m = GpeNativeFwdRpathsGet{} }
func (*GpeNativeFwdRpathsGet) GetMessageName() string { return "gpe_native_fwd_rpaths_get" }
func (*GpeNativeFwdRpathsGet) GetCrcString() string   { return "f652ceb4" }
func (*GpeNativeFwdRpathsGet) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GpeNativeFwdRpathsGet) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsIP4
	return size
}
func (m *GpeNativeFwdRpathsGet) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsIP4)
	return buf.Bytes(), nil
}
func (m *GpeNativeFwdRpathsGet) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsIP4 = buf.DecodeBool()
	return nil
}

// Reply for get native fwd rpath
//   - retval - return code
//   - table_id - table id for route path
//   - nh_sw_if_index - next-hop sw_if_index (~0 if not set)
//   - nh_addr - next hop address
//
// GpeNativeFwdRpathsGetReply defines message 'gpe_native_fwd_rpaths_get_reply'.
type GpeNativeFwdRpathsGetReply struct {
	Retval  int32               `binapi:"i32,name=retval" json:"retval,omitempty"`
	Count   uint32              `binapi:"u32,name=count" json:"-"`
	Entries []GpeNativeFwdRpath `binapi:"gpe_native_fwd_rpath[count],name=entries" json:"entries,omitempty"`
}

func (m *GpeNativeFwdRpathsGetReply) Reset()               { *m = GpeNativeFwdRpathsGetReply{} }
func (*GpeNativeFwdRpathsGetReply) GetMessageName() string { return "gpe_native_fwd_rpaths_get_reply" }
func (*GpeNativeFwdRpathsGetReply) GetCrcString() string   { return "7a1ca5a2" }
func (*GpeNativeFwdRpathsGetReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GpeNativeFwdRpathsGetReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.Count
	for j1 := 0; j1 < len(m.Entries); j1++ {
		var s1 GpeNativeFwdRpath
		_ = s1
		if j1 < len(m.Entries) {
			s1 = m.Entries[j1]
		}
		size += 4      // s1.FibIndex
		size += 4      // s1.NhSwIfIndex
		size += 1      // s1.NhAddr.Af
		size += 1 * 16 // s1.NhAddr.Un
	}
	return size
}
func (m *GpeNativeFwdRpathsGetReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(len(m.Entries)))
	for j0 := 0; j0 < len(m.Entries); j0++ {
		var v0 GpeNativeFwdRpath // Entries
		if j0 < len(m.Entries) {
			v0 = m.Entries[j0]
		}
		buf.EncodeUint32(v0.FibIndex)
		buf.EncodeUint32(uint32(v0.NhSwIfIndex))
		buf.EncodeUint8(uint8(v0.NhAddr.Af))
		buf.EncodeBytes(v0.NhAddr.Un.XXX_UnionData[:], 16)
	}
	return buf.Bytes(), nil
}
func (m *GpeNativeFwdRpathsGetReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.Count = buf.DecodeUint32()
	m.Entries = make([]GpeNativeFwdRpath, m.Count)
	for j0 := 0; j0 < len(m.Entries); j0++ {
		m.Entries[j0].FibIndex = buf.DecodeUint32()
		m.Entries[j0].NhSwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
		m.Entries[j0].NhAddr.Af = ip_types.AddressFamily(buf.DecodeUint8())
		copy(m.Entries[j0].NhAddr.Un.XXX_UnionData[:], buf.DecodeBytes(16))
	}
	return nil
}

// Set GPE encapsulation mode
//   - mode - LISP (value 0) or VXLAN (value 1)
//
// GpeSetEncapMode defines message 'gpe_set_encap_mode'.
type GpeSetEncapMode struct {
	IsVxlan bool `binapi:"bool,name=is_vxlan" json:"is_vxlan,omitempty"`
}

func (m *GpeSetEncapMode) Reset()               { *m = GpeSetEncapMode{} }
func (*GpeSetEncapMode) GetMessageName() string { return "gpe_set_encap_mode" }
func (*GpeSetEncapMode) GetCrcString() string   { return "bd819eac" }
func (*GpeSetEncapMode) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *GpeSetEncapMode) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsVxlan
	return size
}
func (m *GpeSetEncapMode) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsVxlan)
	return buf.Bytes(), nil
}
func (m *GpeSetEncapMode) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsVxlan = buf.DecodeBool()
	return nil
}

// GpeSetEncapModeReply defines message 'gpe_set_encap_mode_reply'.
type GpeSetEncapModeReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *GpeSetEncapModeReply) Reset()               { *m = GpeSetEncapModeReply{} }
func (*GpeSetEncapModeReply) GetMessageName() string { return "gpe_set_encap_mode_reply" }
func (*GpeSetEncapModeReply) GetCrcString() string   { return "e8d4e804" }
func (*GpeSetEncapModeReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *GpeSetEncapModeReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *GpeSetEncapModeReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *GpeSetEncapModeReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_lisp_gpe_binapi_init() }
func file_lisp_gpe_binapi_init() {
	api.RegisterMessage((*GpeAddDelFwdEntry)(nil), "gpe_add_del_fwd_entry_f0847644")
	api.RegisterMessage((*GpeAddDelFwdEntryReply)(nil), "gpe_add_del_fwd_entry_reply_efe5f176")
	api.RegisterMessage((*GpeAddDelIface)(nil), "gpe_add_del_iface_3ccff273")
	api.RegisterMessage((*GpeAddDelIfaceReply)(nil), "gpe_add_del_iface_reply_e8d4e804")
	api.RegisterMessage((*GpeAddDelNativeFwdRpath)(nil), "gpe_add_del_native_fwd_rpath_43fc8b54")
	api.RegisterMessage((*GpeAddDelNativeFwdRpathReply)(nil), "gpe_add_del_native_fwd_rpath_reply_e8d4e804")
	api.RegisterMessage((*GpeEnableDisable)(nil), "gpe_enable_disable_c264d7bf")
	api.RegisterMessage((*GpeEnableDisableReply)(nil), "gpe_enable_disable_reply_e8d4e804")
	api.RegisterMessage((*GpeFwdEntriesGet)(nil), "gpe_fwd_entries_get_8d1f2fe9")
	api.RegisterMessage((*GpeFwdEntriesGetReply)(nil), "gpe_fwd_entries_get_reply_c4844876")
	api.RegisterMessage((*GpeFwdEntryPathDetails)(nil), "gpe_fwd_entry_path_details_483df51a")
	api.RegisterMessage((*GpeFwdEntryPathDump)(nil), "gpe_fwd_entry_path_dump_39bce980")
	api.RegisterMessage((*GpeFwdEntryVnisGet)(nil), "gpe_fwd_entry_vnis_get_51077d14")
	api.RegisterMessage((*GpeFwdEntryVnisGetReply)(nil), "gpe_fwd_entry_vnis_get_reply_aa70da20")
	api.RegisterMessage((*GpeGetEncapMode)(nil), "gpe_get_encap_mode_51077d14")
	api.RegisterMessage((*GpeGetEncapModeReply)(nil), "gpe_get_encap_mode_reply_36e3f7ca")
	api.RegisterMessage((*GpeNativeFwdRpathsGet)(nil), "gpe_native_fwd_rpaths_get_f652ceb4")
	api.RegisterMessage((*GpeNativeFwdRpathsGetReply)(nil), "gpe_native_fwd_rpaths_get_reply_7a1ca5a2")
	api.RegisterMessage((*GpeSetEncapMode)(nil), "gpe_set_encap_mode_bd819eac")
	api.RegisterMessage((*GpeSetEncapModeReply)(nil), "gpe_set_encap_mode_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*GpeAddDelFwdEntry)(nil),
		(*GpeAddDelFwdEntryReply)(nil),
		(*GpeAddDelIface)(nil),
		(*GpeAddDelIfaceReply)(nil),
		(*GpeAddDelNativeFwdRpath)(nil),
		(*GpeAddDelNativeFwdRpathReply)(nil),
		(*GpeEnableDisable)(nil),
		(*GpeEnableDisableReply)(nil),
		(*GpeFwdEntriesGet)(nil),
		(*GpeFwdEntriesGetReply)(nil),
		(*GpeFwdEntryPathDetails)(nil),
		(*GpeFwdEntryPathDump)(nil),
		(*GpeFwdEntryVnisGet)(nil),
		(*GpeFwdEntryVnisGetReply)(nil),
		(*GpeGetEncapMode)(nil),
		(*GpeGetEncapModeReply)(nil),
		(*GpeNativeFwdRpathsGet)(nil),
		(*GpeNativeFwdRpathsGetReply)(nil),
		(*GpeSetEncapMode)(nil),
		(*GpeSetEncapModeReply)(nil),
	}
}
