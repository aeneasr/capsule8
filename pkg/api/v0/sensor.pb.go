// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sensor.proto

package v0

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Sensor struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Sysname  string `protobuf:"bytes,10,opt,name=sysname" json:"sysname,omitempty"`
	Nodename string `protobuf:"bytes,11,opt,name=nodename" json:"nodename,omitempty"`
	Release  string `protobuf:"bytes,12,opt,name=release" json:"release,omitempty"`
	Version  string `protobuf:"bytes,13,opt,name=version" json:"version,omitempty"`
}

func (m *Sensor) Reset()                    { *m = Sensor{} }
func (m *Sensor) String() string            { return proto.CompactTextString(m) }
func (*Sensor) ProtoMessage()               {}
func (*Sensor) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Sensor) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Sensor) GetSysname() string {
	if m != nil {
		return m.Sysname
	}
	return ""
}

func (m *Sensor) GetNodename() string {
	if m != nil {
		return m.Nodename
	}
	return ""
}

func (m *Sensor) GetRelease() string {
	if m != nil {
		return m.Release
	}
	return ""
}

func (m *Sensor) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*Sensor)(nil), "capsule8.v0.Sensor")
}

func init() { proto.RegisterFile("sensor.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 147 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0xcd, 0x2b,
	0xce, 0x2f, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x4e, 0x4e, 0x2c, 0x28, 0x2e, 0xcd,
	0x49, 0xb5, 0xd0, 0x2b, 0x33, 0x50, 0x6a, 0x62, 0xe4, 0x62, 0x0b, 0x06, 0xcb, 0x0a, 0xf1, 0x71,
	0x31, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x08, 0x49, 0x70,
	0xb1, 0x17, 0x57, 0x16, 0xe7, 0x25, 0xe6, 0xa6, 0x4a, 0x70, 0x81, 0x05, 0x61, 0x5c, 0x21, 0x29,
	0x2e, 0x8e, 0xbc, 0xfc, 0x94, 0x54, 0xb0, 0x14, 0x37, 0x58, 0x0a, 0xce, 0x07, 0xe9, 0x2a, 0x4a,
	0xcd, 0x49, 0x4d, 0x2c, 0x4e, 0x95, 0xe0, 0x81, 0xe8, 0x82, 0x72, 0x41, 0x32, 0x65, 0xa9, 0x45,
	0xc5, 0x99, 0xf9, 0x79, 0x12, 0xbc, 0x10, 0x19, 0x28, 0xd7, 0x89, 0x25, 0x8a, 0xa9, 0xcc, 0x20,
	0x89, 0x0d, 0xec, 0x3c, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0xf3, 0x78, 0x4e, 0xae,
	0x00, 0x00, 0x00,
}