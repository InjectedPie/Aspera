// Code generated by protoc-gen-go. DO NOT EDIT.
// source: p2p.proto

package p2p

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetCumulativeDifficultyResponse struct {
	BlockchainHeight     int32    `protobuf:"varint,1,opt,name=blockchainHeight" json:"blockchainHeight,omitempty"`
	CumulativeDifficulty string   `protobuf:"bytes,2,opt,name=cumulativeDifficulty" json:"cumulativeDifficulty,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetCumulativeDifficultyResponse) Reset()         { *m = GetCumulativeDifficultyResponse{} }
func (m *GetCumulativeDifficultyResponse) String() string { return proto.CompactTextString(m) }
func (*GetCumulativeDifficultyResponse) ProtoMessage()    {}
func (*GetCumulativeDifficultyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_5339ae67d5193418, []int{0}
}
func (m *GetCumulativeDifficultyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetCumulativeDifficultyResponse.Unmarshal(m, b)
}
func (m *GetCumulativeDifficultyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetCumulativeDifficultyResponse.Marshal(b, m, deterministic)
}
func (dst *GetCumulativeDifficultyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetCumulativeDifficultyResponse.Merge(dst, src)
}
func (m *GetCumulativeDifficultyResponse) XXX_Size() int {
	return xxx_messageInfo_GetCumulativeDifficultyResponse.Size(m)
}
func (m *GetCumulativeDifficultyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetCumulativeDifficultyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetCumulativeDifficultyResponse proto.InternalMessageInfo

func (m *GetCumulativeDifficultyResponse) GetBlockchainHeight() int32 {
	if m != nil {
		return m.BlockchainHeight
	}
	return 0
}

func (m *GetCumulativeDifficultyResponse) GetCumulativeDifficulty() string {
	if m != nil {
		return m.CumulativeDifficulty
	}
	return ""
}

type GetNextBlocksResponse struct {
	NextBlocks           []*Block `protobuf:"bytes,1,rep,name=nextBlocks" json:"nextBlocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNextBlocksResponse) Reset()         { *m = GetNextBlocksResponse{} }
func (m *GetNextBlocksResponse) String() string { return proto.CompactTextString(m) }
func (*GetNextBlocksResponse) ProtoMessage()    {}
func (*GetNextBlocksResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_5339ae67d5193418, []int{1}
}
func (m *GetNextBlocksResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNextBlocksResponse.Unmarshal(m, b)
}
func (m *GetNextBlocksResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNextBlocksResponse.Marshal(b, m, deterministic)
}
func (dst *GetNextBlocksResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNextBlocksResponse.Merge(dst, src)
}
func (m *GetNextBlocksResponse) XXX_Size() int {
	return xxx_messageInfo_GetNextBlocksResponse.Size(m)
}
func (m *GetNextBlocksResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNextBlocksResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNextBlocksResponse proto.InternalMessageInfo

func (m *GetNextBlocksResponse) GetNextBlocks() []*Block {
	if m != nil {
		return m.NextBlocks
	}
	return nil
}

type GetNextBlockIdsResponse struct {
	NextBlockIds         []uint64 `protobuf:"varint,1,rep,packed,name=nextBlockIds" json:"nextBlockIds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetNextBlockIdsResponse) Reset()         { *m = GetNextBlockIdsResponse{} }
func (m *GetNextBlockIdsResponse) String() string { return proto.CompactTextString(m) }
func (*GetNextBlockIdsResponse) ProtoMessage()    {}
func (*GetNextBlockIdsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_5339ae67d5193418, []int{2}
}
func (m *GetNextBlockIdsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetNextBlockIdsResponse.Unmarshal(m, b)
}
func (m *GetNextBlockIdsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetNextBlockIdsResponse.Marshal(b, m, deterministic)
}
func (dst *GetNextBlockIdsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetNextBlockIdsResponse.Merge(dst, src)
}
func (m *GetNextBlockIdsResponse) XXX_Size() int {
	return xxx_messageInfo_GetNextBlockIdsResponse.Size(m)
}
func (m *GetNextBlockIdsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetNextBlockIdsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetNextBlockIdsResponse proto.InternalMessageInfo

func (m *GetNextBlockIdsResponse) GetNextBlockIds() []uint64 {
	if m != nil {
		return m.NextBlockIds
	}
	return nil
}

type GetPeers struct {
	Peers                []string `protobuf:"bytes,1,rep,name=peers" json:"peers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPeers) Reset()         { *m = GetPeers{} }
func (m *GetPeers) String() string { return proto.CompactTextString(m) }
func (*GetPeers) ProtoMessage()    {}
func (*GetPeers) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_5339ae67d5193418, []int{3}
}
func (m *GetPeers) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPeers.Unmarshal(m, b)
}
func (m *GetPeers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPeers.Marshal(b, m, deterministic)
}
func (dst *GetPeers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPeers.Merge(dst, src)
}
func (m *GetPeers) XXX_Size() int {
	return xxx_messageInfo_GetPeers.Size(m)
}
func (m *GetPeers) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPeers.DiscardUnknown(m)
}

var xxx_messageInfo_GetPeers proto.InternalMessageInfo

func (m *GetPeers) GetPeers() []string {
	if m != nil {
		return m.Peers
	}
	return nil
}

type Block struct {
	PayloadLength        uint32     `protobuf:"varint,1,opt,name=payloadLength" json:"payloadLength,omitempty"`
	TotalAmount          int64      `protobuf:"varint,2,opt,name=totalAmount" json:"totalAmount,omitempty"`
	GenerationSignature  []byte     `protobuf:"bytes,3,opt,name=generationSignature,proto3" json:"generationSignature,omitempty"`
	GeneratorPublicKey   []byte     `protobuf:"bytes,4,opt,name=generatorPublicKey,proto3" json:"generatorPublicKey,omitempty"`
	PayloadHash          []byte     `protobuf:"bytes,5,opt,name=payloadHash,proto3" json:"payloadHash,omitempty"`
	BlockSignature       []byte     `protobuf:"bytes,6,opt,name=blockSignature,proto3" json:"blockSignature,omitempty"`
	Transactions         []*any.Any `protobuf:"bytes,7,rep,name=transactions" json:"transactions,omitempty"`
	Version              int32      `protobuf:"varint,8,opt,name=version" json:"version,omitempty"`
	Nonce                uint64     `protobuf:"varint,9,opt,name=nonce" json:"nonce,omitempty"`
	TotalFee             int64      `protobuf:"varint,10,opt,name=totalFee" json:"totalFee,omitempty"`
	BlockATs             []byte     `protobuf:"bytes,11,opt,name=blockATs,proto3" json:"blockATs,omitempty"`
	PreviousBlock        uint64     `protobuf:"varint,12,opt,name=previousBlock" json:"previousBlock,omitempty"`
	Timestamp            uint32     `protobuf:"varint,13,opt,name=timestamp" json:"timestamp,omitempty"`
	Id                   uint64     `protobuf:"varint,14,opt,name=id" json:"id,omitempty"`
	Height               int32      `protobuf:"varint,15,opt,name=height" json:"height,omitempty"`
	PreviousBlockHash    []byte     `protobuf:"bytes,16,opt,name=previousBlockHash,proto3" json:"previousBlockHash,omitempty"`
	BaseTarget           uint64     `protobuf:"varint,17,opt,name=baseTarget" json:"baseTarget,omitempty"`
	CumulativeDifficulty []byte     `protobuf:"bytes,18,opt,name=cumulativeDifficulty,proto3" json:"cumulativeDifficulty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_p2p_5339ae67d5193418, []int{4}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (dst *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(dst, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetPayloadLength() uint32 {
	if m != nil {
		return m.PayloadLength
	}
	return 0
}

func (m *Block) GetTotalAmount() int64 {
	if m != nil {
		return m.TotalAmount
	}
	return 0
}

func (m *Block) GetGenerationSignature() []byte {
	if m != nil {
		return m.GenerationSignature
	}
	return nil
}

func (m *Block) GetGeneratorPublicKey() []byte {
	if m != nil {
		return m.GeneratorPublicKey
	}
	return nil
}

func (m *Block) GetPayloadHash() []byte {
	if m != nil {
		return m.PayloadHash
	}
	return nil
}

func (m *Block) GetBlockSignature() []byte {
	if m != nil {
		return m.BlockSignature
	}
	return nil
}

func (m *Block) GetTransactions() []*any.Any {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func (m *Block) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Block) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *Block) GetTotalFee() int64 {
	if m != nil {
		return m.TotalFee
	}
	return 0
}

func (m *Block) GetBlockATs() []byte {
	if m != nil {
		return m.BlockATs
	}
	return nil
}

func (m *Block) GetPreviousBlock() uint64 {
	if m != nil {
		return m.PreviousBlock
	}
	return 0
}

func (m *Block) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Block) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Block) GetHeight() int32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Block) GetPreviousBlockHash() []byte {
	if m != nil {
		return m.PreviousBlockHash
	}
	return nil
}

func (m *Block) GetBaseTarget() uint64 {
	if m != nil {
		return m.BaseTarget
	}
	return 0
}

func (m *Block) GetCumulativeDifficulty() []byte {
	if m != nil {
		return m.CumulativeDifficulty
	}
	return nil
}

func init() {
	proto.RegisterType((*GetCumulativeDifficultyResponse)(nil), "p2p.GetCumulativeDifficultyResponse")
	proto.RegisterType((*GetNextBlocksResponse)(nil), "p2p.GetNextBlocksResponse")
	proto.RegisterType((*GetNextBlockIdsResponse)(nil), "p2p.GetNextBlockIdsResponse")
	proto.RegisterType((*GetPeers)(nil), "p2p.GetPeers")
	proto.RegisterType((*Block)(nil), "p2p.Block")
}

func init() { proto.RegisterFile("p2p.proto", fileDescriptor_p2p_5339ae67d5193418) }

var fileDescriptor_p2p_5339ae67d5193418 = []byte{
	// 511 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0xdf, 0x8f, 0xd2, 0x40,
	0x10, 0xc7, 0xd3, 0xe3, 0xc7, 0xc1, 0xf0, 0xc3, 0x63, 0x45, 0x5d, 0x89, 0xd1, 0xa6, 0x31, 0x86,
	0x18, 0xd3, 0x33, 0xf8, 0xe2, 0x8b, 0x0f, 0x78, 0x46, 0xce, 0x68, 0xcc, 0x65, 0xbd, 0x7f, 0x60,
	0x29, 0x43, 0xd9, 0x58, 0x76, 0x9b, 0xee, 0x96, 0xc8, 0xa3, 0xcf, 0xfe, 0xd3, 0xa6, 0x53, 0x0e,
	0x8a, 0x87, 0x6f, 0xcc, 0xe7, 0x3b, 0xfd, 0x76, 0xf8, 0xce, 0x14, 0xda, 0xe9, 0x24, 0x0d, 0xd3,
	0xcc, 0x38, 0xc3, 0x6a, 0xe9, 0x24, 0x1d, 0x3d, 0x8d, 0x8d, 0x89, 0x13, 0xbc, 0x24, 0x34, 0xcf,
	0x97, 0x97, 0x52, 0x6f, 0x4b, 0x7d, 0x34, 0x70, 0x99, 0xd4, 0x56, 0x46, 0x4e, 0x19, 0x5d, 0xa2,
	0xe0, 0xb7, 0x07, 0x2f, 0x66, 0xe8, 0xae, 0xf2, 0x75, 0x9e, 0x48, 0xa7, 0x36, 0xf8, 0x49, 0x2d,
	0x97, 0x2a, 0xca, 0x13, 0xb7, 0x15, 0x68, 0x53, 0xa3, 0x2d, 0xb2, 0xd7, 0x70, 0x31, 0x4f, 0x4c,
	0xf4, 0x33, 0x5a, 0x49, 0xa5, 0xaf, 0x51, 0xc5, 0x2b, 0xc7, 0x3d, 0xdf, 0x1b, 0x37, 0xc4, 0x3d,
	0xce, 0x26, 0x30, 0x8c, 0x4e, 0x78, 0xf1, 0x33, 0xdf, 0x1b, 0xb7, 0xc5, 0x49, 0x2d, 0xb8, 0x82,
	0x47, 0x33, 0x74, 0xdf, 0xf1, 0x97, 0xfb, 0x58, 0xd8, 0xd9, 0xca, 0x8b, 0x41, 0xef, 0x29, 0xf7,
	0xfc, 0xda, 0xb8, 0x33, 0x81, 0xb0, 0xf8, 0xbf, 0x84, 0x44, 0x45, 0x0d, 0x3e, 0xc0, 0x93, 0xaa,
	0xc9, 0x97, 0xc5, 0xc1, 0x26, 0x80, 0xae, 0xae, 0x70, 0x32, 0xaa, 0x8b, 0x23, 0x16, 0xf8, 0xd0,
	0x9a, 0xa1, 0xbb, 0x41, 0xcc, 0x2c, 0x1b, 0x42, 0x23, 0x2d, 0x7e, 0x50, 0x63, 0x5b, 0x94, 0x45,
	0xf0, 0xa7, 0x01, 0x0d, 0x6a, 0x67, 0x2f, 0xa1, 0x97, 0xca, 0x6d, 0x62, 0xe4, 0xe2, 0x1b, 0xea,
	0xd8, 0xad, 0x28, 0x8c, 0x9e, 0x38, 0x86, 0xcc, 0x87, 0x8e, 0x33, 0x4e, 0x26, 0xd3, 0xb5, 0xc9,
	0xb5, 0xa3, 0x00, 0x6a, 0xa2, 0x8a, 0xd8, 0x5b, 0x78, 0x18, 0xa3, 0xc6, 0x4c, 0x16, 0xfb, 0xf8,
	0xa1, 0x62, 0x2d, 0x5d, 0x9e, 0x21, 0xaf, 0xf9, 0xde, 0xb8, 0x2b, 0x4e, 0x49, 0x2c, 0x04, 0xb6,
	0xc3, 0x26, 0xbb, 0xc9, 0xe7, 0x89, 0x8a, 0xbe, 0xe2, 0x96, 0xd7, 0xe9, 0x81, 0x13, 0x4a, 0x31,
	0xc3, 0x6e, 0xa8, 0x6b, 0x69, 0x57, 0xbc, 0x41, 0x8d, 0x55, 0xc4, 0x5e, 0x41, 0x9f, 0x76, 0x78,
	0x78, 0x7d, 0x93, 0x9a, 0xfe, 0xa1, 0xec, 0x3d, 0x74, 0x2b, 0xc7, 0x63, 0xf9, 0x39, 0x2d, 0x63,
	0x18, 0x96, 0xc7, 0x16, 0xde, 0x1d, 0x5b, 0x38, 0xd5, 0x5b, 0x71, 0xd4, 0xc9, 0x38, 0x9c, 0x6f,
	0x30, 0xb3, 0xca, 0x68, 0xde, 0xa2, 0xa3, 0xb9, 0x2b, 0x8b, 0x9c, 0xb5, 0xd1, 0x11, 0xf2, 0xb6,
	0xef, 0x8d, 0xeb, 0xa2, 0x2c, 0xd8, 0x08, 0x5a, 0x14, 0xd2, 0x67, 0x44, 0x0e, 0x14, 0xda, 0xbe,
	0x2e, 0x34, 0x9a, 0x6b, 0x7a, 0x6b, 0x79, 0x87, 0xe6, 0xdc, 0xd7, 0xb4, 0x95, 0x0c, 0x37, 0xca,
	0xe4, 0x96, 0xd6, 0xc4, 0xbb, 0xe4, 0x7a, 0x0c, 0xd9, 0x33, 0x68, 0x3b, 0xb5, 0x46, 0xeb, 0xe4,
	0x3a, 0xe5, 0x3d, 0xda, 0xdb, 0x01, 0xb0, 0x3e, 0x9c, 0xa9, 0x05, 0xef, 0xd3, 0x83, 0x67, 0x6a,
	0xc1, 0x1e, 0x43, 0x73, 0x55, 0xde, 0xfb, 0x03, 0x1a, 0x7d, 0x57, 0xb1, 0x37, 0x30, 0x38, 0xb2,
	0xa5, 0x74, 0x2f, 0x68, 0xa0, 0xfb, 0x02, 0x7b, 0x0e, 0x30, 0x97, 0x16, 0x6f, 0x65, 0x16, 0xa3,
	0xe3, 0x03, 0x72, 0xaf, 0x90, 0xff, 0x7e, 0x33, 0x8c, 0x0c, 0x4f, 0x6a, 0xf3, 0x26, 0x25, 0xfe,
	0xee, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x73, 0xec, 0x94, 0x85, 0xfe, 0x03, 0x00, 0x00,
}
