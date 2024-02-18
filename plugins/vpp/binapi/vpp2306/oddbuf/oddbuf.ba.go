// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

// Package oddbuf contains generated bindings for API file oddbuf.api.
//
// Contents:
// -  2 messages
package oddbuf

import (
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
	APIFile    = "oddbuf"
	APIVersion = "0.1.0"
	VersionCrc = 0x3a1a2c50
)

// @brief API to enable / disable oddbuf on an interface
//   - enable_disable - 1 to enable, 0 to disable the feature
//   - sw_if_index - interface handle
//
// OddbufEnableDisable defines message 'oddbuf_enable_disable'.
type OddbufEnableDisable struct {
	EnableDisable bool                           `binapi:"bool,name=enable_disable" json:"enable_disable,omitempty"`
	SwIfIndex     interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *OddbufEnableDisable) Reset()               { *m = OddbufEnableDisable{} }
func (*OddbufEnableDisable) GetMessageName() string { return "oddbuf_enable_disable" }
func (*OddbufEnableDisable) GetCrcString() string   { return "3865946c" }
func (*OddbufEnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *OddbufEnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.EnableDisable
	size += 4 // m.SwIfIndex
	return size
}
func (m *OddbufEnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.EnableDisable)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *OddbufEnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.EnableDisable = buf.DecodeBool()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// OddbufEnableDisableReply defines message 'oddbuf_enable_disable_reply'.
type OddbufEnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *OddbufEnableDisableReply) Reset()               { *m = OddbufEnableDisableReply{} }
func (*OddbufEnableDisableReply) GetMessageName() string { return "oddbuf_enable_disable_reply" }
func (*OddbufEnableDisableReply) GetCrcString() string   { return "e8d4e804" }
func (*OddbufEnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *OddbufEnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *OddbufEnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *OddbufEnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_oddbuf_binapi_init() }
func file_oddbuf_binapi_init() {
	api.RegisterMessage((*OddbufEnableDisable)(nil), "oddbuf_enable_disable_3865946c")
	api.RegisterMessage((*OddbufEnableDisableReply)(nil), "oddbuf_enable_disable_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*OddbufEnableDisable)(nil),
		(*OddbufEnableDisableReply)(nil),
	}
}
