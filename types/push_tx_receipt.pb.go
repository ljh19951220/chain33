// Code generated by protoc-gen-go. DO NOT EDIT.
// source: push_tx_receipt.proto

package types

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type TxReceipts4SubscribePerBlk struct {
	Tx          []*Transaction `protobuf:"bytes,1,rep,name=tx,proto3" json:"tx,omitempty"`
	ReceiptData []*ReceiptData `protobuf:"bytes,2,rep,name=receiptData,proto3" json:"receiptData,omitempty"`
	// repeated KeyValue    KV          = 3;
	Height               int64    `protobuf:"varint,4,opt,name=height,proto3" json:"height,omitempty"`
	BlockHash            []byte   `protobuf:"bytes,5,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
	ParentHash           []byte   `protobuf:"bytes,6,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	PreviousHash         []byte   `protobuf:"bytes,7,opt,name=previousHash,proto3" json:"previousHash,omitempty"`
	AddDelType           int32    `protobuf:"varint,8,opt,name=addDelType,proto3" json:"addDelType,omitempty"`
	SeqNum               int64    `protobuf:"varint,9,opt,name=seqNum,proto3" json:"seqNum,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxReceipts4SubscribePerBlk) Reset()         { *m = TxReceipts4SubscribePerBlk{} }
func (m *TxReceipts4SubscribePerBlk) String() string { return proto.CompactTextString(m) }
func (*TxReceipts4SubscribePerBlk) ProtoMessage()    {}
func (*TxReceipts4SubscribePerBlk) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5e438adec79672e, []int{0}
}

func (m *TxReceipts4SubscribePerBlk) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxReceipts4SubscribePerBlk.Unmarshal(m, b)
}
func (m *TxReceipts4SubscribePerBlk) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxReceipts4SubscribePerBlk.Marshal(b, m, deterministic)
}
func (m *TxReceipts4SubscribePerBlk) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxReceipts4SubscribePerBlk.Merge(m, src)
}
func (m *TxReceipts4SubscribePerBlk) XXX_Size() int {
	return xxx_messageInfo_TxReceipts4SubscribePerBlk.Size(m)
}
func (m *TxReceipts4SubscribePerBlk) XXX_DiscardUnknown() {
	xxx_messageInfo_TxReceipts4SubscribePerBlk.DiscardUnknown(m)
}

var xxx_messageInfo_TxReceipts4SubscribePerBlk proto.InternalMessageInfo

func (m *TxReceipts4SubscribePerBlk) GetTx() []*Transaction {
	if m != nil {
		return m.Tx
	}
	return nil
}

func (m *TxReceipts4SubscribePerBlk) GetReceiptData() []*ReceiptData {
	if m != nil {
		return m.ReceiptData
	}
	return nil
}

func (m *TxReceipts4SubscribePerBlk) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *TxReceipts4SubscribePerBlk) GetBlockHash() []byte {
	if m != nil {
		return m.BlockHash
	}
	return nil
}

func (m *TxReceipts4SubscribePerBlk) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *TxReceipts4SubscribePerBlk) GetPreviousHash() []byte {
	if m != nil {
		return m.PreviousHash
	}
	return nil
}

func (m *TxReceipts4SubscribePerBlk) GetAddDelType() int32 {
	if m != nil {
		return m.AddDelType
	}
	return 0
}

func (m *TxReceipts4SubscribePerBlk) GetSeqNum() int64 {
	if m != nil {
		return m.SeqNum
	}
	return 0
}

type TxReceipts4Subscribe struct {
	TxReceipts           []*TxReceipts4SubscribePerBlk `protobuf:"bytes,1,rep,name=txReceipts,proto3" json:"txReceipts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *TxReceipts4Subscribe) Reset()         { *m = TxReceipts4Subscribe{} }
func (m *TxReceipts4Subscribe) String() string { return proto.CompactTextString(m) }
func (*TxReceipts4Subscribe) ProtoMessage()    {}
func (*TxReceipts4Subscribe) Descriptor() ([]byte, []int) {
	return fileDescriptor_d5e438adec79672e, []int{1}
}

func (m *TxReceipts4Subscribe) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxReceipts4Subscribe.Unmarshal(m, b)
}
func (m *TxReceipts4Subscribe) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxReceipts4Subscribe.Marshal(b, m, deterministic)
}
func (m *TxReceipts4Subscribe) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxReceipts4Subscribe.Merge(m, src)
}
func (m *TxReceipts4Subscribe) XXX_Size() int {
	return xxx_messageInfo_TxReceipts4Subscribe.Size(m)
}
func (m *TxReceipts4Subscribe) XXX_DiscardUnknown() {
	xxx_messageInfo_TxReceipts4Subscribe.DiscardUnknown(m)
}

var xxx_messageInfo_TxReceipts4Subscribe proto.InternalMessageInfo

func (m *TxReceipts4Subscribe) GetTxReceipts() []*TxReceipts4SubscribePerBlk {
	if m != nil {
		return m.TxReceipts
	}
	return nil
}

func init() {
	proto.RegisterType((*TxReceipts4SubscribePerBlk)(nil), "types.TxReceipts4SubscribePerBlk")
	proto.RegisterType((*TxReceipts4Subscribe)(nil), "types.TxReceipts4Subscribe")
}

func init() {
	proto.RegisterFile("push_tx_receipt.proto", fileDescriptor_d5e438adec79672e)
}

var fileDescriptor_d5e438adec79672e = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x5f, 0x4f, 0xc2, 0x30,
	0x14, 0xc5, 0xb3, 0x21, 0x28, 0x17, 0x5e, 0x6c, 0xd4, 0x2c, 0xc4, 0x3f, 0x73, 0x4f, 0x7b, 0x1a,
	0x89, 0xe3, 0x0b, 0x48, 0x78, 0xf0, 0xc9, 0x98, 0xca, 0x8b, 0xbe, 0x90, 0xae, 0x34, 0xac, 0x01,
	0xd6, 0xda, 0xde, 0x99, 0xf1, 0x8d, 0xfc, 0x98, 0x86, 0x0e, 0xdd, 0x48, 0xf4, 0xf1, 0x9e, 0xf3,
	0xeb, 0xe9, 0xcd, 0xb9, 0x70, 0xa9, 0x4b, 0x9b, 0x2f, 0xb0, 0x5a, 0x18, 0xc1, 0x85, 0xd4, 0x98,
	0x68, 0xa3, 0x50, 0x91, 0x2e, 0xee, 0xb4, 0xb0, 0xa3, 0x73, 0x34, 0xac, 0xb0, 0x8c, 0xa3, 0x54,
	0x45, 0xed, 0x44, 0x5f, 0x3e, 0x8c, 0xe6, 0x15, 0xad, 0x69, 0x3b, 0x79, 0x2d, 0x33, 0xcb, 0x8d,
	0xcc, 0xc4, 0x8b, 0x30, 0xd3, 0xcd, 0x9a, 0x44, 0xe0, 0x63, 0x15, 0x78, 0x61, 0x27, 0x1e, 0x3c,
	0x90, 0xc4, 0xa5, 0x24, 0xf3, 0x26, 0x84, 0xfa, 0x58, 0x91, 0x09, 0x0c, 0x0e, 0xbf, 0xcd, 0x18,
	0xb2, 0xc0, 0x3f, 0x82, 0x69, 0xe3, 0xd0, 0x36, 0x46, 0xae, 0xa0, 0x97, 0x0b, 0xb9, 0xca, 0x31,
	0x38, 0x09, 0xbd, 0xb8, 0x43, 0x0f, 0x13, 0xb9, 0x86, 0x7e, 0xb6, 0x51, 0x7c, 0xfd, 0xc4, 0x6c,
	0x1e, 0x74, 0x43, 0x2f, 0x1e, 0xd2, 0x46, 0x20, 0xb7, 0x00, 0x9a, 0x19, 0x51, 0xa0, 0xb3, 0x7b,
	0xce, 0x6e, 0x29, 0x24, 0x82, 0xa1, 0x36, 0xe2, 0x53, 0xaa, 0xd2, 0x3a, 0xe2, 0xd4, 0x11, 0x47,
	0xda, 0x3e, 0x83, 0x2d, 0x97, 0x33, 0xb1, 0x99, 0xef, 0xb4, 0x08, 0xce, 0x42, 0x2f, 0xee, 0xd2,
	0x96, 0xb2, 0xdf, 0xcc, 0x8a, 0x8f, 0xe7, 0x72, 0x1b, 0xf4, 0xeb, 0xcd, 0xea, 0x29, 0x7a, 0x83,
	0x8b, 0xbf, 0x9a, 0x22, 0x8f, 0x00, 0xf8, 0xab, 0x1f, 0xba, 0xba, 0xff, 0xe9, 0xea, 0xdf, 0x6a,
	0x69, 0xeb, 0xd1, 0xf4, 0xee, 0xfd, 0x66, 0x25, 0x31, 0x2f, 0xb3, 0x84, 0xab, 0xed, 0x38, 0x4d,
	0x79, 0x31, 0xe6, 0x39, 0x93, 0x45, 0x9a, 0x8e, 0x5d, 0x4e, 0xd6, 0x73, 0xd7, 0x4a, 0xbf, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xd6, 0xd0, 0x6e, 0x3c, 0xe0, 0x01, 0x00, 0x00,
}
