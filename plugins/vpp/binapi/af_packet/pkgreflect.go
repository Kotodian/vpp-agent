// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.

package af_packet

import "reflect"

var Types = map[string]reflect.Type{
	"AfPacketCreate": reflect.TypeOf((*AfPacketCreate)(nil)).Elem(),
	"AfPacketCreateReply": reflect.TypeOf((*AfPacketCreateReply)(nil)).Elem(),
	"AfPacketDelete": reflect.TypeOf((*AfPacketDelete)(nil)).Elem(),
	"AfPacketDeleteReply": reflect.TypeOf((*AfPacketDeleteReply)(nil)).Elem(),
	"AfPacketDetails": reflect.TypeOf((*AfPacketDetails)(nil)).Elem(),
	"AfPacketDump": reflect.TypeOf((*AfPacketDump)(nil)).Elem(),
	"AfPacketSetL4CksumOffload": reflect.TypeOf((*AfPacketSetL4CksumOffload)(nil)).Elem(),
	"AfPacketSetL4CksumOffloadReply": reflect.TypeOf((*AfPacketSetL4CksumOffloadReply)(nil)).Elem(),
	"Services": reflect.TypeOf((*Services)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"NewAfPacketCreate": reflect.ValueOf(NewAfPacketCreate),
	"NewAfPacketCreateReply": reflect.ValueOf(NewAfPacketCreateReply),
	"NewAfPacketDelete": reflect.ValueOf(NewAfPacketDelete),
	"NewAfPacketDeleteReply": reflect.ValueOf(NewAfPacketDeleteReply),
	"NewAfPacketDetails": reflect.ValueOf(NewAfPacketDetails),
	"NewAfPacketDump": reflect.ValueOf(NewAfPacketDump),
	"NewAfPacketSetL4CksumOffload": reflect.ValueOf(NewAfPacketSetL4CksumOffload),
	"NewAfPacketSetL4CksumOffloadReply": reflect.ValueOf(NewAfPacketSetL4CksumOffloadReply),
}

var Variables = map[string]reflect.Value{
}

var Consts = map[string]reflect.Value{
}
