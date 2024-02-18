// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

// Package ioam_cache contains generated bindings for API file ioam_cache.api.
//
// Contents:
// -  2 messages
package ioam_cache

import (
	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "ioam_cache"
	APIVersion = "1.0.0"
	VersionCrc = 0xd0a0cf20
)

// /*  API to control ioam caching
// IoamCacheIP6EnableDisable defines message 'ioam_cache_ip6_enable_disable'.
type IoamCacheIP6EnableDisable struct {
	IsDisable bool `binapi:"bool,name=is_disable" json:"is_disable,omitempty"`
}

func (m *IoamCacheIP6EnableDisable) Reset()               { *m = IoamCacheIP6EnableDisable{} }
func (*IoamCacheIP6EnableDisable) GetMessageName() string { return "ioam_cache_ip6_enable_disable" }
func (*IoamCacheIP6EnableDisable) GetCrcString() string   { return "47705c03" }
func (*IoamCacheIP6EnableDisable) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *IoamCacheIP6EnableDisable) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1 // m.IsDisable
	return size
}
func (m *IoamCacheIP6EnableDisable) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsDisable)
	return buf.Bytes(), nil
}
func (m *IoamCacheIP6EnableDisable) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsDisable = buf.DecodeBool()
	return nil
}

// IoamCacheIP6EnableDisableReply defines message 'ioam_cache_ip6_enable_disable_reply'.
type IoamCacheIP6EnableDisableReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *IoamCacheIP6EnableDisableReply) Reset() { *m = IoamCacheIP6EnableDisableReply{} }
func (*IoamCacheIP6EnableDisableReply) GetMessageName() string {
	return "ioam_cache_ip6_enable_disable_reply"
}
func (*IoamCacheIP6EnableDisableReply) GetCrcString() string { return "e8d4e804" }
func (*IoamCacheIP6EnableDisableReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *IoamCacheIP6EnableDisableReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *IoamCacheIP6EnableDisableReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *IoamCacheIP6EnableDisableReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

func init() { file_ioam_cache_binapi_init() }
func file_ioam_cache_binapi_init() {
	api.RegisterMessage((*IoamCacheIP6EnableDisable)(nil), "ioam_cache_ip6_enable_disable_47705c03")
	api.RegisterMessage((*IoamCacheIP6EnableDisableReply)(nil), "ioam_cache_ip6_enable_disable_reply_e8d4e804")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*IoamCacheIP6EnableDisable)(nil),
		(*IoamCacheIP6EnableDisableReply)(nil),
	}
}
