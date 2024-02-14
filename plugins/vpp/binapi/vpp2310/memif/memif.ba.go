// Code generated by GoVPP's binapi-generator. DO NOT EDIT.

// Package memif contains generated bindings for API file memif.api.
//
// Contents:
// -  2 enums
// - 14 messages
package memif

import (
	"strconv"

	api "go.fd.io/govpp/api"
	codec "go.fd.io/govpp/codec"
	ethernet_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/ethernet_types"
	interface_types "go.ligato.io/vpp-agent/v3/plugins/vpp/binapi/vpp2310/interface_types"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the GoVPP api package it is being compiled against.
// A compilation error at this line likely means your copy of the
// GoVPP api package needs to be updated.
const _ = api.GoVppAPIPackageIsVersion2

const (
	APIFile    = "memif"
	APIVersion = "3.1.0"
	VersionCrc = 0xd48ac702
)

// MemifMode defines enum 'memif_mode'.
type MemifMode uint32

const (
	MEMIF_MODE_API_ETHERNET    MemifMode = 0
	MEMIF_MODE_API_IP          MemifMode = 1
	MEMIF_MODE_API_PUNT_INJECT MemifMode = 2
)

var (
	MemifMode_name = map[uint32]string{
		0: "MEMIF_MODE_API_ETHERNET",
		1: "MEMIF_MODE_API_IP",
		2: "MEMIF_MODE_API_PUNT_INJECT",
	}
	MemifMode_value = map[string]uint32{
		"MEMIF_MODE_API_ETHERNET":    0,
		"MEMIF_MODE_API_IP":          1,
		"MEMIF_MODE_API_PUNT_INJECT": 2,
	}
)

func (x MemifMode) String() string {
	s, ok := MemifMode_name[uint32(x)]
	if ok {
		return s
	}
	return "MemifMode(" + strconv.Itoa(int(x)) + ")"
}

// MemifRole defines enum 'memif_role'.
type MemifRole uint32

const (
	MEMIF_ROLE_API_MASTER MemifRole = 0
	MEMIF_ROLE_API_SLAVE  MemifRole = 1
)

var (
	MemifRole_name = map[uint32]string{
		0: "MEMIF_ROLE_API_MASTER",
		1: "MEMIF_ROLE_API_SLAVE",
	}
	MemifRole_value = map[string]uint32{
		"MEMIF_ROLE_API_MASTER": 0,
		"MEMIF_ROLE_API_SLAVE":  1,
	}
)

func (x MemifRole) String() string {
	s, ok := MemifRole_name[uint32(x)]
	if ok {
		return s
	}
	return "MemifRole(" + strconv.Itoa(int(x)) + ")"
}

// Create memory interface
//   - role - role of the interface in the connection (master/slave)
//   - mode - interface mode
//   - rx_queues - number of rx queues (only valid for slave)
//   - tx_queues - number of tx queues (only valid for slave)
//   - id - 32bit integer used to authenticate and match opposite sides
//     of the connection
//   - socket_id - socket filename id to be used for connection
//     establishment
//   - ring_size - the number of entries of RX/TX rings
//   - buffer_size - size of the buffer allocated for each ring entry
//   - no_zero_copy - if true, disable zero copy
//   - hw_addr - interface MAC address
//   - secret - optional, default is "", max length 24
//
// MemifCreate defines message 'memif_create'.
// Deprecated: the message will be removed in the future versions
type MemifCreate struct {
	Role       MemifRole                 `binapi:"memif_role,name=role" json:"role,omitempty"`
	Mode       MemifMode                 `binapi:"memif_mode,name=mode" json:"mode,omitempty"`
	RxQueues   uint8                     `binapi:"u8,name=rx_queues" json:"rx_queues,omitempty"`
	TxQueues   uint8                     `binapi:"u8,name=tx_queues" json:"tx_queues,omitempty"`
	ID         uint32                    `binapi:"u32,name=id" json:"id,omitempty"`
	SocketID   uint32                    `binapi:"u32,name=socket_id" json:"socket_id,omitempty"`
	RingSize   uint32                    `binapi:"u32,name=ring_size" json:"ring_size,omitempty"`
	BufferSize uint16                    `binapi:"u16,name=buffer_size" json:"buffer_size,omitempty"`
	NoZeroCopy bool                      `binapi:"bool,name=no_zero_copy" json:"no_zero_copy,omitempty"`
	HwAddr     ethernet_types.MacAddress `binapi:"mac_address,name=hw_addr" json:"hw_addr,omitempty"`
	Secret     string                    `binapi:"string[24],name=secret" json:"secret,omitempty"`
}

func (m *MemifCreate) Reset()               { *m = MemifCreate{} }
func (*MemifCreate) GetMessageName() string { return "memif_create" }
func (*MemifCreate) GetCrcString() string   { return "b1b25061" }
func (*MemifCreate) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *MemifCreate) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4     // m.Role
	size += 4     // m.Mode
	size += 1     // m.RxQueues
	size += 1     // m.TxQueues
	size += 4     // m.ID
	size += 4     // m.SocketID
	size += 4     // m.RingSize
	size += 2     // m.BufferSize
	size += 1     // m.NoZeroCopy
	size += 1 * 6 // m.HwAddr
	size += 24    // m.Secret
	return size
}
func (m *MemifCreate) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.Role))
	buf.EncodeUint32(uint32(m.Mode))
	buf.EncodeUint8(m.RxQueues)
	buf.EncodeUint8(m.TxQueues)
	buf.EncodeUint32(m.ID)
	buf.EncodeUint32(m.SocketID)
	buf.EncodeUint32(m.RingSize)
	buf.EncodeUint16(m.BufferSize)
	buf.EncodeBool(m.NoZeroCopy)
	buf.EncodeBytes(m.HwAddr[:], 6)
	buf.EncodeString(m.Secret, 24)
	return buf.Bytes(), nil
}
func (m *MemifCreate) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Role = MemifRole(buf.DecodeUint32())
	m.Mode = MemifMode(buf.DecodeUint32())
	m.RxQueues = buf.DecodeUint8()
	m.TxQueues = buf.DecodeUint8()
	m.ID = buf.DecodeUint32()
	m.SocketID = buf.DecodeUint32()
	m.RingSize = buf.DecodeUint32()
	m.BufferSize = buf.DecodeUint16()
	m.NoZeroCopy = buf.DecodeBool()
	copy(m.HwAddr[:], buf.DecodeBytes(6))
	m.Secret = buf.DecodeString(24)
	return nil
}

// Create memory interface response
//   - retval - return value for request
//   - sw_if_index - software index of the newly created interface
//
// MemifCreateReply defines message 'memif_create_reply'.
// Deprecated: the message will be removed in the future versions
type MemifCreateReply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *MemifCreateReply) Reset()               { *m = MemifCreateReply{} }
func (*MemifCreateReply) GetMessageName() string { return "memif_create_reply" }
func (*MemifCreateReply) GetCrcString() string   { return "5383d31f" }
func (*MemifCreateReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *MemifCreateReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *MemifCreateReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *MemifCreateReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// Create memory interface
//   - role - role of the interface in the connection (master/slave)
//   - mode - interface mode
//   - rx_queues - number of rx queues (only valid for slave)
//   - tx_queues - number of tx queues (only valid for slave)
//   - id - 32bit integer used to authenticate and match opposite sides
//     of the connection
//   - socket_id - socket filename id to be used for connection
//     establishment
//   - ring_size - the number of entries of RX/TX rings
//   - buffer_size - size of the buffer allocated for each ring entry
//   - no_zero_copy - if true, disable zero copy
//   - use_dma - if true, use dma accelerate memory copy
//   - hw_addr - interface MAC address
//   - secret - optional, default is "", max length 24
//
// MemifCreateV2 defines message 'memif_create_v2'.
type MemifCreateV2 struct {
	Role       MemifRole                 `binapi:"memif_role,name=role" json:"role,omitempty"`
	Mode       MemifMode                 `binapi:"memif_mode,name=mode" json:"mode,omitempty"`
	RxQueues   uint8                     `binapi:"u8,name=rx_queues" json:"rx_queues,omitempty"`
	TxQueues   uint8                     `binapi:"u8,name=tx_queues" json:"tx_queues,omitempty"`
	ID         uint32                    `binapi:"u32,name=id" json:"id,omitempty"`
	SocketID   uint32                    `binapi:"u32,name=socket_id" json:"socket_id,omitempty"`
	RingSize   uint32                    `binapi:"u32,name=ring_size" json:"ring_size,omitempty"`
	BufferSize uint16                    `binapi:"u16,name=buffer_size" json:"buffer_size,omitempty"`
	NoZeroCopy bool                      `binapi:"bool,name=no_zero_copy" json:"no_zero_copy,omitempty"`
	UseDma     bool                      `binapi:"bool,name=use_dma" json:"use_dma,omitempty"`
	HwAddr     ethernet_types.MacAddress `binapi:"mac_address,name=hw_addr" json:"hw_addr,omitempty"`
	Secret     string                    `binapi:"string[24],name=secret" json:"secret,omitempty"`
}

func (m *MemifCreateV2) Reset()               { *m = MemifCreateV2{} }
func (*MemifCreateV2) GetMessageName() string { return "memif_create_v2" }
func (*MemifCreateV2) GetCrcString() string   { return "8c7de5f7" }
func (*MemifCreateV2) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *MemifCreateV2) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4     // m.Role
	size += 4     // m.Mode
	size += 1     // m.RxQueues
	size += 1     // m.TxQueues
	size += 4     // m.ID
	size += 4     // m.SocketID
	size += 4     // m.RingSize
	size += 2     // m.BufferSize
	size += 1     // m.NoZeroCopy
	size += 1     // m.UseDma
	size += 1 * 6 // m.HwAddr
	size += 24    // m.Secret
	return size
}
func (m *MemifCreateV2) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.Role))
	buf.EncodeUint32(uint32(m.Mode))
	buf.EncodeUint8(m.RxQueues)
	buf.EncodeUint8(m.TxQueues)
	buf.EncodeUint32(m.ID)
	buf.EncodeUint32(m.SocketID)
	buf.EncodeUint32(m.RingSize)
	buf.EncodeUint16(m.BufferSize)
	buf.EncodeBool(m.NoZeroCopy)
	buf.EncodeBool(m.UseDma)
	buf.EncodeBytes(m.HwAddr[:], 6)
	buf.EncodeString(m.Secret, 24)
	return buf.Bytes(), nil
}
func (m *MemifCreateV2) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Role = MemifRole(buf.DecodeUint32())
	m.Mode = MemifMode(buf.DecodeUint32())
	m.RxQueues = buf.DecodeUint8()
	m.TxQueues = buf.DecodeUint8()
	m.ID = buf.DecodeUint32()
	m.SocketID = buf.DecodeUint32()
	m.RingSize = buf.DecodeUint32()
	m.BufferSize = buf.DecodeUint16()
	m.NoZeroCopy = buf.DecodeBool()
	m.UseDma = buf.DecodeBool()
	copy(m.HwAddr[:], buf.DecodeBytes(6))
	m.Secret = buf.DecodeString(24)
	return nil
}

// Create memory interface response
//   - retval - return value for request
//   - sw_if_index - software index of the newly created interface
//
// MemifCreateV2Reply defines message 'memif_create_v2_reply'.
type MemifCreateV2Reply struct {
	Retval    int32                          `binapi:"i32,name=retval" json:"retval,omitempty"`
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *MemifCreateV2Reply) Reset()               { *m = MemifCreateV2Reply{} }
func (*MemifCreateV2Reply) GetMessageName() string { return "memif_create_v2_reply" }
func (*MemifCreateV2Reply) GetCrcString() string   { return "5383d31f" }
func (*MemifCreateV2Reply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *MemifCreateV2Reply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SwIfIndex
	return size
}
func (m *MemifCreateV2Reply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *MemifCreateV2Reply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// Delete memory interface
//   - sw_if_index - software index of the interface to delete
//
// MemifDelete defines message 'memif_delete'.
type MemifDelete struct {
	SwIfIndex interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
}

func (m *MemifDelete) Reset()               { *m = MemifDelete{} }
func (*MemifDelete) GetMessageName() string { return "memif_delete" }
func (*MemifDelete) GetCrcString() string   { return "f9e6675e" }
func (*MemifDelete) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *MemifDelete) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.SwIfIndex
	return size
}
func (m *MemifDelete) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	return buf.Bytes(), nil
}
func (m *MemifDelete) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	return nil
}

// MemifDeleteReply defines message 'memif_delete_reply'.
type MemifDeleteReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *MemifDeleteReply) Reset()               { *m = MemifDeleteReply{} }
func (*MemifDeleteReply) GetMessageName() string { return "memif_delete_reply" }
func (*MemifDeleteReply) GetCrcString() string   { return "e8d4e804" }
func (*MemifDeleteReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *MemifDeleteReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *MemifDeleteReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *MemifDeleteReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// Memory interface details structure
//   - sw_if_index - index of the interface
//   - hw_addr - interface MAC address
//   - id - id associated with the interface
//   - role - role of the interface in the connection (master/slave)
//   - mode - interface mode
//   - zero_copy - zero copy flag present
//   - socket_id - id of the socket filename used by this interface
//     to establish new connections
//   - ring_size - the number of entries of RX/TX rings
//   - buffer_size - size of the buffer allocated for each ring entry
//   - flags - interface_status flags
//   - if_name - name of the interface
//
// MemifDetails defines message 'memif_details'.
type MemifDetails struct {
	SwIfIndex  interface_types.InterfaceIndex `binapi:"interface_index,name=sw_if_index" json:"sw_if_index,omitempty"`
	HwAddr     ethernet_types.MacAddress      `binapi:"mac_address,name=hw_addr" json:"hw_addr,omitempty"`
	ID         uint32                         `binapi:"u32,name=id" json:"id,omitempty"`
	Role       MemifRole                      `binapi:"memif_role,name=role" json:"role,omitempty"`
	Mode       MemifMode                      `binapi:"memif_mode,name=mode" json:"mode,omitempty"`
	ZeroCopy   bool                           `binapi:"bool,name=zero_copy" json:"zero_copy,omitempty"`
	SocketID   uint32                         `binapi:"u32,name=socket_id" json:"socket_id,omitempty"`
	RingSize   uint32                         `binapi:"u32,name=ring_size" json:"ring_size,omitempty"`
	BufferSize uint16                         `binapi:"u16,name=buffer_size" json:"buffer_size,omitempty"`
	Flags      interface_types.IfStatusFlags  `binapi:"if_status_flags,name=flags" json:"flags,omitempty"`
	IfName     string                         `binapi:"string[64],name=if_name" json:"if_name,omitempty"`
}

func (m *MemifDetails) Reset()               { *m = MemifDetails{} }
func (*MemifDetails) GetMessageName() string { return "memif_details" }
func (*MemifDetails) GetCrcString() string   { return "da34feb9" }
func (*MemifDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *MemifDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4     // m.SwIfIndex
	size += 1 * 6 // m.HwAddr
	size += 4     // m.ID
	size += 4     // m.Role
	size += 4     // m.Mode
	size += 1     // m.ZeroCopy
	size += 4     // m.SocketID
	size += 4     // m.RingSize
	size += 2     // m.BufferSize
	size += 4     // m.Flags
	size += 64    // m.IfName
	return size
}
func (m *MemifDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(uint32(m.SwIfIndex))
	buf.EncodeBytes(m.HwAddr[:], 6)
	buf.EncodeUint32(m.ID)
	buf.EncodeUint32(uint32(m.Role))
	buf.EncodeUint32(uint32(m.Mode))
	buf.EncodeBool(m.ZeroCopy)
	buf.EncodeUint32(m.SocketID)
	buf.EncodeUint32(m.RingSize)
	buf.EncodeUint16(m.BufferSize)
	buf.EncodeUint32(uint32(m.Flags))
	buf.EncodeString(m.IfName, 64)
	return buf.Bytes(), nil
}
func (m *MemifDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SwIfIndex = interface_types.InterfaceIndex(buf.DecodeUint32())
	copy(m.HwAddr[:], buf.DecodeBytes(6))
	m.ID = buf.DecodeUint32()
	m.Role = MemifRole(buf.DecodeUint32())
	m.Mode = MemifMode(buf.DecodeUint32())
	m.ZeroCopy = buf.DecodeBool()
	m.SocketID = buf.DecodeUint32()
	m.RingSize = buf.DecodeUint32()
	m.BufferSize = buf.DecodeUint16()
	m.Flags = interface_types.IfStatusFlags(buf.DecodeUint32())
	m.IfName = buf.DecodeString(64)
	return nil
}

// Dump all memory interfaces
// MemifDump defines message 'memif_dump'.
type MemifDump struct{}

func (m *MemifDump) Reset()               { *m = MemifDump{} }
func (*MemifDump) GetMessageName() string { return "memif_dump" }
func (*MemifDump) GetCrcString() string   { return "51077d14" }
func (*MemifDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *MemifDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *MemifDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *MemifDump) Unmarshal(b []byte) error {
	return nil
}

// Create or remove named socket file for memif interfaces
//   - is_add - 0 = remove, 1 = add association
//   - socket_id - non-0 32-bit integer used to identify a socket file
//   - socket_filename - filename of the socket to be used for connection
//     establishment; id 0 always maps to default "/var/vpp/memif.sock";
//     no socket filename needed when is_add == 0.
//
// MemifSocketFilenameAddDel defines message 'memif_socket_filename_add_del'.
// Deprecated: the message will be removed in the future versions
type MemifSocketFilenameAddDel struct {
	IsAdd          bool   `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	SocketID       uint32 `binapi:"u32,name=socket_id" json:"socket_id,omitempty"`
	SocketFilename string `binapi:"string[108],name=socket_filename" json:"socket_filename,omitempty"`
}

func (m *MemifSocketFilenameAddDel) Reset()               { *m = MemifSocketFilenameAddDel{} }
func (*MemifSocketFilenameAddDel) GetMessageName() string { return "memif_socket_filename_add_del" }
func (*MemifSocketFilenameAddDel) GetCrcString() string   { return "a2ce1a10" }
func (*MemifSocketFilenameAddDel) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *MemifSocketFilenameAddDel) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1   // m.IsAdd
	size += 4   // m.SocketID
	size += 108 // m.SocketFilename
	return size
}
func (m *MemifSocketFilenameAddDel) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint32(m.SocketID)
	buf.EncodeString(m.SocketFilename, 108)
	return buf.Bytes(), nil
}
func (m *MemifSocketFilenameAddDel) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.SocketID = buf.DecodeUint32()
	m.SocketFilename = buf.DecodeString(108)
	return nil
}

// MemifSocketFilenameAddDelReply defines message 'memif_socket_filename_add_del_reply'.
// Deprecated: the message will be removed in the future versions
type MemifSocketFilenameAddDelReply struct {
	Retval int32 `binapi:"i32,name=retval" json:"retval,omitempty"`
}

func (m *MemifSocketFilenameAddDelReply) Reset() { *m = MemifSocketFilenameAddDelReply{} }
func (*MemifSocketFilenameAddDelReply) GetMessageName() string {
	return "memif_socket_filename_add_del_reply"
}
func (*MemifSocketFilenameAddDelReply) GetCrcString() string { return "e8d4e804" }
func (*MemifSocketFilenameAddDelReply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *MemifSocketFilenameAddDelReply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	return size
}
func (m *MemifSocketFilenameAddDelReply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	return buf.Bytes(), nil
}
func (m *MemifSocketFilenameAddDelReply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	return nil
}

// Create or remove named socket file for memif interfaces
//   - is_add - 0 = remove, 1 = add association
//   - socket_id - non-0 32-bit integer used to identify a socket file
//     ~0 means autogenerate
//   - socket_filename - filename of the socket to be used for connection
//     establishment; id 0 always maps to default "/var/vpp/memif.sock";
//     no socket filename needed when is_add == 0.
//     socket_filename starting with '@' will create an abstract socket
//     in the given namespace
//
// MemifSocketFilenameAddDelV2 defines message 'memif_socket_filename_add_del_v2'.
type MemifSocketFilenameAddDelV2 struct {
	IsAdd          bool   `binapi:"bool,name=is_add" json:"is_add,omitempty"`
	SocketID       uint32 `binapi:"u32,name=socket_id,default=4294967295" json:"socket_id,omitempty"`
	SocketFilename string `binapi:"string[],name=socket_filename" json:"socket_filename,omitempty"`
}

func (m *MemifSocketFilenameAddDelV2) Reset() { *m = MemifSocketFilenameAddDelV2{} }
func (*MemifSocketFilenameAddDelV2) GetMessageName() string {
	return "memif_socket_filename_add_del_v2"
}
func (*MemifSocketFilenameAddDelV2) GetCrcString() string { return "34223bdf" }
func (*MemifSocketFilenameAddDelV2) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *MemifSocketFilenameAddDelV2) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 1                         // m.IsAdd
	size += 4                         // m.SocketID
	size += 4 + len(m.SocketFilename) // m.SocketFilename
	return size
}
func (m *MemifSocketFilenameAddDelV2) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeBool(m.IsAdd)
	buf.EncodeUint32(m.SocketID)
	buf.EncodeString(m.SocketFilename, 0)
	return buf.Bytes(), nil
}
func (m *MemifSocketFilenameAddDelV2) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.IsAdd = buf.DecodeBool()
	m.SocketID = buf.DecodeUint32()
	m.SocketFilename = buf.DecodeString(0)
	return nil
}

// Create memory interface socket file response
//   - retval - return value for request
//   - socket_id - non-0 32-bit integer used to identify a socket file
//
// MemifSocketFilenameAddDelV2Reply defines message 'memif_socket_filename_add_del_v2_reply'.
type MemifSocketFilenameAddDelV2Reply struct {
	Retval   int32  `binapi:"i32,name=retval" json:"retval,omitempty"`
	SocketID uint32 `binapi:"u32,name=socket_id" json:"socket_id,omitempty"`
}

func (m *MemifSocketFilenameAddDelV2Reply) Reset() { *m = MemifSocketFilenameAddDelV2Reply{} }
func (*MemifSocketFilenameAddDelV2Reply) GetMessageName() string {
	return "memif_socket_filename_add_del_v2_reply"
}
func (*MemifSocketFilenameAddDelV2Reply) GetCrcString() string { return "9f29bdb9" }
func (*MemifSocketFilenameAddDelV2Reply) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *MemifSocketFilenameAddDelV2Reply) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4 // m.Retval
	size += 4 // m.SocketID
	return size
}
func (m *MemifSocketFilenameAddDelV2Reply) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeInt32(m.Retval)
	buf.EncodeUint32(m.SocketID)
	return buf.Bytes(), nil
}
func (m *MemifSocketFilenameAddDelV2Reply) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.Retval = buf.DecodeInt32()
	m.SocketID = buf.DecodeUint32()
	return nil
}

// Memory interface details structure
//   - socket_id - u32 used to identify the given socket filename
//   - socket_filename - corresponding NUL terminated socket filename
//
// MemifSocketFilenameDetails defines message 'memif_socket_filename_details'.
type MemifSocketFilenameDetails struct {
	SocketID       uint32 `binapi:"u32,name=socket_id" json:"socket_id,omitempty"`
	SocketFilename string `binapi:"string[108],name=socket_filename" json:"socket_filename,omitempty"`
}

func (m *MemifSocketFilenameDetails) Reset()               { *m = MemifSocketFilenameDetails{} }
func (*MemifSocketFilenameDetails) GetMessageName() string { return "memif_socket_filename_details" }
func (*MemifSocketFilenameDetails) GetCrcString() string   { return "7ff326f7" }
func (*MemifSocketFilenameDetails) GetMessageType() api.MessageType {
	return api.ReplyMessage
}

func (m *MemifSocketFilenameDetails) Size() (size int) {
	if m == nil {
		return 0
	}
	size += 4   // m.SocketID
	size += 108 // m.SocketFilename
	return size
}
func (m *MemifSocketFilenameDetails) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	buf.EncodeUint32(m.SocketID)
	buf.EncodeString(m.SocketFilename, 108)
	return buf.Bytes(), nil
}
func (m *MemifSocketFilenameDetails) Unmarshal(b []byte) error {
	buf := codec.NewBuffer(b)
	m.SocketID = buf.DecodeUint32()
	m.SocketFilename = buf.DecodeString(108)
	return nil
}

// Dump the table of socket ids and corresponding filenames
// MemifSocketFilenameDump defines message 'memif_socket_filename_dump'.
type MemifSocketFilenameDump struct{}

func (m *MemifSocketFilenameDump) Reset()               { *m = MemifSocketFilenameDump{} }
func (*MemifSocketFilenameDump) GetMessageName() string { return "memif_socket_filename_dump" }
func (*MemifSocketFilenameDump) GetCrcString() string   { return "51077d14" }
func (*MemifSocketFilenameDump) GetMessageType() api.MessageType {
	return api.RequestMessage
}

func (m *MemifSocketFilenameDump) Size() (size int) {
	if m == nil {
		return 0
	}
	return size
}
func (m *MemifSocketFilenameDump) Marshal(b []byte) ([]byte, error) {
	if b == nil {
		b = make([]byte, m.Size())
	}
	buf := codec.NewBuffer(b)
	return buf.Bytes(), nil
}
func (m *MemifSocketFilenameDump) Unmarshal(b []byte) error {
	return nil
}

func init() { file_memif_binapi_init() }
func file_memif_binapi_init() {
	api.RegisterMessage((*MemifCreate)(nil), "memif_create_b1b25061")
	api.RegisterMessage((*MemifCreateReply)(nil), "memif_create_reply_5383d31f")
	api.RegisterMessage((*MemifCreateV2)(nil), "memif_create_v2_8c7de5f7")
	api.RegisterMessage((*MemifCreateV2Reply)(nil), "memif_create_v2_reply_5383d31f")
	api.RegisterMessage((*MemifDelete)(nil), "memif_delete_f9e6675e")
	api.RegisterMessage((*MemifDeleteReply)(nil), "memif_delete_reply_e8d4e804")
	api.RegisterMessage((*MemifDetails)(nil), "memif_details_da34feb9")
	api.RegisterMessage((*MemifDump)(nil), "memif_dump_51077d14")
	api.RegisterMessage((*MemifSocketFilenameAddDel)(nil), "memif_socket_filename_add_del_a2ce1a10")
	api.RegisterMessage((*MemifSocketFilenameAddDelReply)(nil), "memif_socket_filename_add_del_reply_e8d4e804")
	api.RegisterMessage((*MemifSocketFilenameAddDelV2)(nil), "memif_socket_filename_add_del_v2_34223bdf")
	api.RegisterMessage((*MemifSocketFilenameAddDelV2Reply)(nil), "memif_socket_filename_add_del_v2_reply_9f29bdb9")
	api.RegisterMessage((*MemifSocketFilenameDetails)(nil), "memif_socket_filename_details_7ff326f7")
	api.RegisterMessage((*MemifSocketFilenameDump)(nil), "memif_socket_filename_dump_51077d14")
}

// Messages returns list of all messages in this module.
func AllMessages() []api.Message {
	return []api.Message{
		(*MemifCreate)(nil),
		(*MemifCreateReply)(nil),
		(*MemifCreateV2)(nil),
		(*MemifCreateV2Reply)(nil),
		(*MemifDelete)(nil),
		(*MemifDeleteReply)(nil),
		(*MemifDetails)(nil),
		(*MemifDump)(nil),
		(*MemifSocketFilenameAddDel)(nil),
		(*MemifSocketFilenameAddDelReply)(nil),
		(*MemifSocketFilenameAddDelV2)(nil),
		(*MemifSocketFilenameAddDelV2Reply)(nil),
		(*MemifSocketFilenameDetails)(nil),
		(*MemifSocketFilenameDump)(nil),
	}
}
