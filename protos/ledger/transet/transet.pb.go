// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ledger/transet/transet.proto

/*
Package transet is a generated protocol buffer package.

It is generated from these files:
	ledger/transet/transet.proto

It has these top-level messages:
	Version
	TranSet
*/
package transet

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Version struct {
	BlockNum uint64 `protobuf:"varint,1,opt,name=block_num,json=blockNum" json:"block_num,omitempty"`
	TxNum    uint64 `protobuf:"varint,2,opt,name=tx_num,json=txNum" json:"tx_num,omitempty"`
}

func (m *Version) Reset()                    { *m = Version{} }
func (m *Version) String() string            { return proto.CompactTextString(m) }
func (*Version) ProtoMessage()               {}
func (*Version) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Version) GetBlockNum() uint64 {
	if m != nil {
		return m.BlockNum
	}
	return 0
}

func (m *Version) GetTxNum() uint64 {
	if m != nil {
		return m.TxNum
	}
	return 0
}

type TranSet struct {
	From    string   `protobuf:"bytes,1,opt,name=from" json:"from,omitempty"`
	FromVer *Version `protobuf:"bytes,2,opt,name=from_ver,json=fromVer" json:"from_ver,omitempty"`
	Transet []byte   `protobuf:"bytes,3,opt,name=transet,proto3" json:"transet,omitempty"`
}

func (m *TranSet) Reset()                    { *m = TranSet{} }
func (m *TranSet) String() string            { return proto.CompactTextString(m) }
func (*TranSet) ProtoMessage()               {}
func (*TranSet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *TranSet) GetFrom() string {
	if m != nil {
		return m.From
	}
	return ""
}

func (m *TranSet) GetFromVer() *Version {
	if m != nil {
		return m.FromVer
	}
	return nil
}

func (m *TranSet) GetTranset() []byte {
	if m != nil {
		return m.Transet
	}
	return nil
}

func init() {
	proto.RegisterType((*Version)(nil), "transet.Version")
	proto.RegisterType((*TranSet)(nil), "transet.TranSet")
}

func init() { proto.RegisterFile("ledger/transet/transet.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x4f, 0x4b, 0x03, 0x31,
	0x10, 0xc5, 0x59, 0xad, 0xdd, 0x76, 0xf4, 0x20, 0x01, 0x61, 0x41, 0x0f, 0xa5, 0xa7, 0x42, 0x21,
	0x01, 0x3d, 0x8a, 0x17, 0x3f, 0x40, 0x0f, 0xab, 0xf4, 0xe0, 0xa5, 0x64, 0x77, 0xd3, 0x6d, 0xdc,
	0xdd, 0x19, 0x99, 0x24, 0xd2, 0x8f, 0x2f, 0xcd, 0x36, 0x48, 0x4f, 0xf3, 0xe7, 0xe5, 0x97, 0xbc,
	0x17, 0x78, 0xea, 0x4d, 0xd3, 0x1a, 0x56, 0x9e, 0x35, 0x3a, 0xe3, 0x53, 0x95, 0x3f, 0x4c, 0x9e,
	0x44, 0x7e, 0x1e, 0x97, 0x6f, 0x90, 0x6f, 0x0d, 0x3b, 0x4b, 0x28, 0x1e, 0x61, 0x5e, 0xf5, 0x54,
	0x77, 0x3b, 0x0c, 0x43, 0x91, 0x2d, 0xb2, 0xd5, 0xa4, 0x9c, 0xc5, 0xc5, 0x26, 0x0c, 0xe2, 0x01,
	0xa6, 0xfe, 0x18, 0x95, 0xab, 0xa8, 0xdc, 0xf8, 0xe3, 0x26, 0x0c, 0xcb, 0x06, 0xf2, 0x4f, 0xd6,
	0xf8, 0x61, 0xbc, 0x10, 0x30, 0xd9, 0x33, 0x8d, 0xe4, 0xbc, 0x8c, 0xbd, 0x58, 0xc3, 0xec, 0x54,
	0x77, 0xbf, 0x86, 0x23, 0x77, 0xfb, 0x7c, 0x2f, 0x93, 0x91, 0xf3, 0xb3, 0x65, 0x7e, 0x3a, 0xb1,
	0x35, 0x2c, 0x0a, 0x48, 0xae, 0x8a, 0xeb, 0x45, 0xb6, 0xba, 0x2b, 0xd3, 0xf8, 0xfe, 0x0d, 0x6b,
	0xe2, 0x56, 0x5a, 0xec, 0xea, 0x83, 0xb6, 0xf8, 0xdf, 0xc4, 0x30, 0x4e, 0x8e, 0x51, 0xd3, 0xcd,
	0x5f, 0xaf, 0xad, 0xf5, 0x87, 0x50, 0xc9, 0x9a, 0x06, 0x65, 0xb1, 0xeb, 0x75, 0xe5, 0xf6, 0x14,
	0xb0, 0xd1, 0xde, 0x12, 0xaa, 0x04, 0xab, 0x11, 0x56, 0x97, 0xff, 0x54, 0x4d, 0xe3, 0xfa, 0xe5,
	0x2f, 0x00, 0x00, 0xff, 0xff, 0x8b, 0x8b, 0x7d, 0xf7, 0x40, 0x01, 0x00, 0x00,
}
