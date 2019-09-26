// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

package walletrpc

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

// A BlockID message contains identifiers to select a block: a height or a
// hash. If the hash is present it takes precedence.
type BlockID struct {
	Height               uint64   `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	Hash                 []byte   `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockID) Reset()         { *m = BlockID{} }
func (m *BlockID) String() string { return proto.CompactTextString(m) }
func (*BlockID) ProtoMessage()    {}
func (*BlockID) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{0}
}

func (m *BlockID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockID.Unmarshal(m, b)
}
func (m *BlockID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockID.Marshal(b, m, deterministic)
}
func (m *BlockID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockID.Merge(m, src)
}
func (m *BlockID) XXX_Size() int {
	return xxx_messageInfo_BlockID.Size(m)
}
func (m *BlockID) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockID.DiscardUnknown(m)
}

var xxx_messageInfo_BlockID proto.InternalMessageInfo

func (m *BlockID) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockID) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

// BlockRange technically allows ranging from hash to hash etc but this is not
// currently intended for support, though there is no reason you couldn't do
// it. Further permutations are left as an exercise.
type BlockRange struct {
	Start                *BlockID `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	End                  *BlockID `protobuf:"bytes,2,opt,name=end,proto3" json:"end,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BlockRange) Reset()         { *m = BlockRange{} }
func (m *BlockRange) String() string { return proto.CompactTextString(m) }
func (*BlockRange) ProtoMessage()    {}
func (*BlockRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{1}
}

func (m *BlockRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BlockRange.Unmarshal(m, b)
}
func (m *BlockRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BlockRange.Marshal(b, m, deterministic)
}
func (m *BlockRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockRange.Merge(m, src)
}
func (m *BlockRange) XXX_Size() int {
	return xxx_messageInfo_BlockRange.Size(m)
}
func (m *BlockRange) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockRange.DiscardUnknown(m)
}

var xxx_messageInfo_BlockRange proto.InternalMessageInfo

func (m *BlockRange) GetStart() *BlockID {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *BlockRange) GetEnd() *BlockID {
	if m != nil {
		return m.End
	}
	return nil
}

// A TxFilter contains the information needed to identify a particular
// transaction: either a block and an index, or a direct transaction hash.
type TxFilter struct {
	Block                *BlockID `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Index                uint64   `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Hash                 []byte   `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TxFilter) Reset()         { *m = TxFilter{} }
func (m *TxFilter) String() string { return proto.CompactTextString(m) }
func (*TxFilter) ProtoMessage()    {}
func (*TxFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{2}
}

func (m *TxFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TxFilter.Unmarshal(m, b)
}
func (m *TxFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TxFilter.Marshal(b, m, deterministic)
}
func (m *TxFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TxFilter.Merge(m, src)
}
func (m *TxFilter) XXX_Size() int {
	return xxx_messageInfo_TxFilter.Size(m)
}
func (m *TxFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_TxFilter.DiscardUnknown(m)
}

var xxx_messageInfo_TxFilter proto.InternalMessageInfo

func (m *TxFilter) GetBlock() *BlockID {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *TxFilter) GetIndex() uint64 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *TxFilter) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

// RawTransaction contains the complete transaction data. It also optionally includes
// the block height in which the transaction was included
type RawTransaction struct {
	Data                 []byte   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Height               uint64   `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RawTransaction) Reset()         { *m = RawTransaction{} }
func (m *RawTransaction) String() string { return proto.CompactTextString(m) }
func (*RawTransaction) ProtoMessage()    {}
func (*RawTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{3}
}

func (m *RawTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RawTransaction.Unmarshal(m, b)
}
func (m *RawTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RawTransaction.Marshal(b, m, deterministic)
}
func (m *RawTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RawTransaction.Merge(m, src)
}
func (m *RawTransaction) XXX_Size() int {
	return xxx_messageInfo_RawTransaction.Size(m)
}
func (m *RawTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_RawTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_RawTransaction proto.InternalMessageInfo

func (m *RawTransaction) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *RawTransaction) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type SendResponse struct {
	ErrorCode            int32    `protobuf:"varint,1,opt,name=errorCode,proto3" json:"errorCode,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SendResponse) Reset()         { *m = SendResponse{} }
func (m *SendResponse) String() string { return proto.CompactTextString(m) }
func (*SendResponse) ProtoMessage()    {}
func (*SendResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{4}
}

func (m *SendResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SendResponse.Unmarshal(m, b)
}
func (m *SendResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SendResponse.Marshal(b, m, deterministic)
}
func (m *SendResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendResponse.Merge(m, src)
}
func (m *SendResponse) XXX_Size() int {
	return xxx_messageInfo_SendResponse.Size(m)
}
func (m *SendResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SendResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SendResponse proto.InternalMessageInfo

func (m *SendResponse) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *SendResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

// Empty placeholder. Someday we may want to specify e.g. a particular chain fork.
type ChainSpec struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChainSpec) Reset()         { *m = ChainSpec{} }
func (m *ChainSpec) String() string { return proto.CompactTextString(m) }
func (*ChainSpec) ProtoMessage()    {}
func (*ChainSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{5}
}

func (m *ChainSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChainSpec.Unmarshal(m, b)
}
func (m *ChainSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChainSpec.Marshal(b, m, deterministic)
}
func (m *ChainSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChainSpec.Merge(m, src)
}
func (m *ChainSpec) XXX_Size() int {
	return xxx_messageInfo_ChainSpec.Size(m)
}
func (m *ChainSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ChainSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ChainSpec proto.InternalMessageInfo

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{6}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type LightdInfo struct {
	Version                 string   `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	Vendor                  string   `protobuf:"bytes,2,opt,name=vendor,proto3" json:"vendor,omitempty"`
	TaddrSupport            bool     `protobuf:"varint,3,opt,name=taddrSupport,proto3" json:"taddrSupport,omitempty"`
	ChainName               string   `protobuf:"bytes,4,opt,name=chainName,proto3" json:"chainName,omitempty"`
	SaplingActivationHeight uint64   `protobuf:"varint,5,opt,name=saplingActivationHeight,proto3" json:"saplingActivationHeight,omitempty"`
	ConsensusBranchId       string   `protobuf:"bytes,6,opt,name=consensusBranchId,proto3" json:"consensusBranchId,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *LightdInfo) Reset()         { *m = LightdInfo{} }
func (m *LightdInfo) String() string { return proto.CompactTextString(m) }
func (*LightdInfo) ProtoMessage()    {}
func (*LightdInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{7}
}

func (m *LightdInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LightdInfo.Unmarshal(m, b)
}
func (m *LightdInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LightdInfo.Marshal(b, m, deterministic)
}
func (m *LightdInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LightdInfo.Merge(m, src)
}
func (m *LightdInfo) XXX_Size() int {
	return xxx_messageInfo_LightdInfo.Size(m)
}
func (m *LightdInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LightdInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LightdInfo proto.InternalMessageInfo

func (m *LightdInfo) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *LightdInfo) GetVendor() string {
	if m != nil {
		return m.Vendor
	}
	return ""
}

func (m *LightdInfo) GetTaddrSupport() bool {
	if m != nil {
		return m.TaddrSupport
	}
	return false
}

func (m *LightdInfo) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

func (m *LightdInfo) GetSaplingActivationHeight() uint64 {
	if m != nil {
		return m.SaplingActivationHeight
	}
	return 0
}

func (m *LightdInfo) GetConsensusBranchId() string {
	if m != nil {
		return m.ConsensusBranchId
	}
	return ""
}

type TransparentAddress struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransparentAddress) Reset()         { *m = TransparentAddress{} }
func (m *TransparentAddress) String() string { return proto.CompactTextString(m) }
func (*TransparentAddress) ProtoMessage()    {}
func (*TransparentAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{8}
}

func (m *TransparentAddress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransparentAddress.Unmarshal(m, b)
}
func (m *TransparentAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransparentAddress.Marshal(b, m, deterministic)
}
func (m *TransparentAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransparentAddress.Merge(m, src)
}
func (m *TransparentAddress) XXX_Size() int {
	return xxx_messageInfo_TransparentAddress.Size(m)
}
func (m *TransparentAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_TransparentAddress.DiscardUnknown(m)
}

var xxx_messageInfo_TransparentAddress proto.InternalMessageInfo

func (m *TransparentAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type TransparentAddressBlockFilter struct {
	Address              string      `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Range                *BlockRange `protobuf:"bytes,2,opt,name=range,proto3" json:"range,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TransparentAddressBlockFilter) Reset()         { *m = TransparentAddressBlockFilter{} }
func (m *TransparentAddressBlockFilter) String() string { return proto.CompactTextString(m) }
func (*TransparentAddressBlockFilter) ProtoMessage()    {}
func (*TransparentAddressBlockFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0b84a42fa06f626, []int{9}
}

func (m *TransparentAddressBlockFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransparentAddressBlockFilter.Unmarshal(m, b)
}
func (m *TransparentAddressBlockFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransparentAddressBlockFilter.Marshal(b, m, deterministic)
}
func (m *TransparentAddressBlockFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransparentAddressBlockFilter.Merge(m, src)
}
func (m *TransparentAddressBlockFilter) XXX_Size() int {
	return xxx_messageInfo_TransparentAddressBlockFilter.Size(m)
}
func (m *TransparentAddressBlockFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_TransparentAddressBlockFilter.DiscardUnknown(m)
}

var xxx_messageInfo_TransparentAddressBlockFilter proto.InternalMessageInfo

func (m *TransparentAddressBlockFilter) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *TransparentAddressBlockFilter) GetRange() *BlockRange {
	if m != nil {
		return m.Range
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockID)(nil), "cash.z.wallet.sdk.rpc.BlockID")
	proto.RegisterType((*BlockRange)(nil), "cash.z.wallet.sdk.rpc.BlockRange")
	proto.RegisterType((*TxFilter)(nil), "cash.z.wallet.sdk.rpc.TxFilter")
	proto.RegisterType((*RawTransaction)(nil), "cash.z.wallet.sdk.rpc.RawTransaction")
	proto.RegisterType((*SendResponse)(nil), "cash.z.wallet.sdk.rpc.SendResponse")
	proto.RegisterType((*ChainSpec)(nil), "cash.z.wallet.sdk.rpc.ChainSpec")
	proto.RegisterType((*Empty)(nil), "cash.z.wallet.sdk.rpc.Empty")
	proto.RegisterType((*LightdInfo)(nil), "cash.z.wallet.sdk.rpc.LightdInfo")
	proto.RegisterType((*TransparentAddress)(nil), "cash.z.wallet.sdk.rpc.TransparentAddress")
	proto.RegisterType((*TransparentAddressBlockFilter)(nil), "cash.z.wallet.sdk.rpc.TransparentAddressBlockFilter")
}

func init() { proto.RegisterFile("service.proto", fileDescriptor_a0b84a42fa06f626) }

var fileDescriptor_a0b84a42fa06f626 = []byte{
	// 630 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0x5d, 0x4f, 0x13, 0x41,
	0x14, 0xa5, 0xd0, 0xed, 0xc7, 0x6d, 0x81, 0x30, 0x11, 0x6d, 0x1a, 0x54, 0x1c, 0x63, 0xe2, 0x83,
	0xd9, 0x10, 0xc4, 0xe8, 0x83, 0x2f, 0xb4, 0x7e, 0x35, 0x41, 0xa3, 0xd3, 0x3e, 0xe1, 0x03, 0x19,
	0x76, 0x87, 0xee, 0x4a, 0xbb, 0xb3, 0x99, 0x19, 0x4a, 0xf5, 0xc7, 0xf8, 0xfb, 0xfc, 0x19, 0xce,
	0xc7, 0x96, 0x6e, 0x03, 0x4b, 0xfb, 0xb6, 0x77, 0xe6, 0xde, 0x73, 0xce, 0x3d, 0x73, 0x6f, 0x0b,
	0x9b, 0x92, 0x89, 0x49, 0x1c, 0x30, 0x3f, 0x15, 0x5c, 0x71, 0xb4, 0x1b, 0x50, 0x19, 0xf9, 0x7f,
	0xfc, 0x6b, 0x3a, 0x1a, 0x31, 0xe5, 0xcb, 0xf0, 0xd2, 0x17, 0x69, 0xd0, 0xde, 0x0d, 0xf8, 0x38,
	0xa5, 0x81, 0x3a, 0xbb, 0xe0, 0x62, 0x4c, 0x95, 0x74, 0xd9, 0xf8, 0x0d, 0x54, 0x3b, 0x23, 0x1e,
	0x5c, 0xf6, 0x3e, 0xa0, 0x87, 0x50, 0x89, 0x58, 0x3c, 0x8c, 0x54, 0xab, 0xb4, 0x5f, 0x7a, 0x59,
	0x26, 0x59, 0x84, 0x10, 0x94, 0x23, 0x0d, 0xd9, 0x5a, 0xd7, 0xa7, 0x4d, 0x62, 0xbf, 0xb1, 0x02,
	0xb0, 0x65, 0x84, 0x26, 0x43, 0x86, 0x8e, 0xc0, 0x93, 0x8a, 0x0a, 0x57, 0xd8, 0x38, 0x7c, 0xe2,
	0xdf, 0x29, 0xc1, 0xcf, 0x88, 0x88, 0x4b, 0x46, 0x07, 0xb0, 0xc1, 0x92, 0xd0, 0xc2, 0x2e, 0xaf,
	0x31, 0xa9, 0xf8, 0x17, 0xd4, 0x06, 0xd3, 0x4f, 0xf1, 0x48, 0x31, 0x61, 0x38, 0xcf, 0xcd, 0xdd,
	0xaa, 0x9c, 0x36, 0x19, 0x3d, 0x00, 0x2f, 0x4e, 0x42, 0x36, 0xb5, 0xac, 0x65, 0xe2, 0x82, 0x9b,
	0x0e, 0x37, 0x72, 0x1d, 0xbe, 0x87, 0x2d, 0x42, 0xaf, 0x07, 0x82, 0x26, 0x52, 0xbb, 0x16, 0xf3,
	0xc4, 0x64, 0x85, 0x54, 0x51, 0x4b, 0xa8, 0xb3, 0xcc, 0x77, 0xce, 0xb3, 0xf5, 0xbc, 0x67, 0xf8,
	0x3b, 0x34, 0xfb, 0x5a, 0x31, 0x61, 0x32, 0xe5, 0x89, 0x64, 0x68, 0x0f, 0xea, 0x4c, 0x08, 0x2e,
	0xba, 0x3c, 0x64, 0x16, 0xc0, 0x23, 0xf3, 0x03, 0x84, 0xa1, 0x69, 0x83, 0xaf, 0x4c, 0x4a, 0x3a,
	0x64, 0x16, 0xab, 0x4e, 0x16, 0xce, 0x70, 0x03, 0xea, 0xdd, 0x88, 0xc6, 0x49, 0x3f, 0x65, 0x01,
	0xae, 0x82, 0xf7, 0x71, 0x9c, 0xaa, 0xdf, 0xf8, 0x5f, 0x09, 0xe0, 0xc4, 0x30, 0x86, 0xbd, 0xe4,
	0x82, 0xa3, 0x16, 0x54, 0x27, 0x4c, 0x48, 0xad, 0xd6, 0x92, 0xd4, 0xc9, 0x2c, 0x34, 0x42, 0x27,
	0x5a, 0x10, 0x17, 0x19, 0x78, 0x16, 0x19, 0x6a, 0x45, 0xc3, 0x50, 0xf4, 0xaf, 0xd2, 0x94, 0xeb,
	0x17, 0x34, 0x16, 0xd4, 0xc8, 0xc2, 0x99, 0x11, 0x1f, 0x18, 0xea, 0x6f, 0x74, 0xcc, 0x5a, 0x65,
	0x5b, 0x3e, 0x3f, 0x40, 0xef, 0xe0, 0x91, 0xa4, 0xe9, 0x28, 0x4e, 0x86, 0xc7, 0xda, 0xa7, 0x09,
	0x35, 0x5e, 0x7d, 0x71, 0x9e, 0x78, 0xd6, 0x93, 0xa2, 0x6b, 0xf4, 0x0a, 0x76, 0x02, 0xe3, 0x4e,
	0x22, 0xaf, 0x64, 0x47, 0x1b, 0x1d, 0x44, 0xbd, 0xb0, 0x55, 0xb1, 0xf8, 0xb7, 0x2f, 0xb0, 0x0f,
	0xc8, 0xbe, 0x46, 0x4a, 0x05, 0x4b, 0xd4, 0xb1, 0xd6, 0xa7, 0x9d, 0x31, 0x1d, 0x53, 0xf7, 0x39,
	0xeb, 0x38, 0x0b, 0xb1, 0x80, 0xc7, 0xb7, 0xf3, 0xed, 0x38, 0x64, 0x13, 0x54, 0x58, 0x8a, 0xde,
	0x82, 0x27, 0xcc, 0x60, 0x67, 0xb3, 0xf9, 0xec, 0xbe, 0xd9, 0xb2, 0x1b, 0x40, 0x5c, 0xfe, 0xe1,
	0x5f, 0x0f, 0x76, 0xba, 0x6e, 0xcf, 0x06, 0xd3, 0xbe, 0x12, 0x4c, 0x1b, 0x24, 0xd0, 0x00, 0xb6,
	0x3e, 0x33, 0x75, 0x42, 0x15, 0x93, 0xca, 0xd6, 0xa0, 0xfd, 0x02, 0xc4, 0x9b, 0x17, 0x6e, 0x2f,
	0x99, 0x67, 0xbc, 0x86, 0x7e, 0x40, 0x4d, 0xa3, 0x3a, 0xbc, 0x25, 0xd9, 0xed, 0xe7, 0x45, 0x7c,
	0x4e, 0xab, 0x4d, 0xd3, 0x90, 0x3f, 0x61, 0x73, 0x06, 0xe9, 0x16, 0x7b, 0x79, 0xe7, 0x2b, 0x42,
	0x1f, 0x94, 0xd0, 0xa9, 0x75, 0x21, 0xbf, 0x50, 0x4f, 0x0b, 0x4a, 0x67, 0x3b, 0xde, 0x7e, 0x51,
	0x90, 0xb0, 0xb8, 0x98, 0x5a, 0xf8, 0x19, 0x6c, 0x9b, 0x75, 0xcb, 0x83, 0xaf, 0x56, 0x5b, 0x28,
	0x3f, 0xbf, 0xbd, 0x9a, 0x40, 0xc0, 0xb6, 0x16, 0x9f, 0x0d, 0xd1, 0x60, 0x1a, 0x87, 0x12, 0x1d,
	0x15, 0xa9, 0xbf, 0x6f, 0xe8, 0x56, 0x6e, 0x49, 0x1b, 0x46, 0xec, 0x6b, 0xe4, 0xb6, 0x7b, 0xaf,
	0xa0, 0xd6, 0xfe, 0x14, 0xb4, 0x8b, 0xde, 0x6a, 0x0e, 0x80, 0xd7, 0x3a, 0x8d, 0xd3, 0xba, 0xbb,
	0xd6, 0x37, 0xe7, 0x15, 0xfb, 0x17, 0xf0, 0xfa, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x84, 0x2f,
	0xbf, 0xf7, 0x41, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CompactTxStreamerClient is the client API for CompactTxStreamer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CompactTxStreamerClient interface {
	// Compact Blocks
	GetLatestBlock(ctx context.Context, in *ChainSpec, opts ...grpc.CallOption) (*BlockID, error)
	GetBlock(ctx context.Context, in *BlockID, opts ...grpc.CallOption) (*CompactBlock, error)
	GetBlockRange(ctx context.Context, in *BlockRange, opts ...grpc.CallOption) (CompactTxStreamer_GetBlockRangeClient, error)
	// Transactions
	GetTransaction(ctx context.Context, in *TxFilter, opts ...grpc.CallOption) (*RawTransaction, error)
	SendTransaction(ctx context.Context, in *RawTransaction, opts ...grpc.CallOption) (*SendResponse, error)
	// t-Address support
	GetAddressTxids(ctx context.Context, in *TransparentAddressBlockFilter, opts ...grpc.CallOption) (CompactTxStreamer_GetAddressTxidsClient, error)
	// Misc
	GetLightdInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*LightdInfo, error)
}

type compactTxStreamerClient struct {
	cc *grpc.ClientConn
}

func NewCompactTxStreamerClient(cc *grpc.ClientConn) CompactTxStreamerClient {
	return &compactTxStreamerClient{cc}
}

func (c *compactTxStreamerClient) GetLatestBlock(ctx context.Context, in *ChainSpec, opts ...grpc.CallOption) (*BlockID, error) {
	out := new(BlockID)
	err := c.cc.Invoke(ctx, "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetLatestBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactTxStreamerClient) GetBlock(ctx context.Context, in *BlockID, opts ...grpc.CallOption) (*CompactBlock, error) {
	out := new(CompactBlock)
	err := c.cc.Invoke(ctx, "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactTxStreamerClient) GetBlockRange(ctx context.Context, in *BlockRange, opts ...grpc.CallOption) (CompactTxStreamer_GetBlockRangeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CompactTxStreamer_serviceDesc.Streams[0], "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetBlockRange", opts...)
	if err != nil {
		return nil, err
	}
	x := &compactTxStreamerGetBlockRangeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CompactTxStreamer_GetBlockRangeClient interface {
	Recv() (*CompactBlock, error)
	grpc.ClientStream
}

type compactTxStreamerGetBlockRangeClient struct {
	grpc.ClientStream
}

func (x *compactTxStreamerGetBlockRangeClient) Recv() (*CompactBlock, error) {
	m := new(CompactBlock)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *compactTxStreamerClient) GetTransaction(ctx context.Context, in *TxFilter, opts ...grpc.CallOption) (*RawTransaction, error) {
	out := new(RawTransaction)
	err := c.cc.Invoke(ctx, "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactTxStreamerClient) SendTransaction(ctx context.Context, in *RawTransaction, opts ...grpc.CallOption) (*SendResponse, error) {
	out := new(SendResponse)
	err := c.cc.Invoke(ctx, "/cash.z.wallet.sdk.rpc.CompactTxStreamer/SendTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *compactTxStreamerClient) GetAddressTxids(ctx context.Context, in *TransparentAddressBlockFilter, opts ...grpc.CallOption) (CompactTxStreamer_GetAddressTxidsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_CompactTxStreamer_serviceDesc.Streams[1], "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetAddressTxids", opts...)
	if err != nil {
		return nil, err
	}
	x := &compactTxStreamerGetAddressTxidsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CompactTxStreamer_GetAddressTxidsClient interface {
	Recv() (*RawTransaction, error)
	grpc.ClientStream
}

type compactTxStreamerGetAddressTxidsClient struct {
	grpc.ClientStream
}

func (x *compactTxStreamerGetAddressTxidsClient) Recv() (*RawTransaction, error) {
	m := new(RawTransaction)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *compactTxStreamerClient) GetLightdInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*LightdInfo, error) {
	out := new(LightdInfo)
	err := c.cc.Invoke(ctx, "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetLightdInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompactTxStreamerServer is the server API for CompactTxStreamer service.
type CompactTxStreamerServer interface {
	// Compact Blocks
	GetLatestBlock(context.Context, *ChainSpec) (*BlockID, error)
	GetBlock(context.Context, *BlockID) (*CompactBlock, error)
	GetBlockRange(*BlockRange, CompactTxStreamer_GetBlockRangeServer) error
	// Transactions
	GetTransaction(context.Context, *TxFilter) (*RawTransaction, error)
	SendTransaction(context.Context, *RawTransaction) (*SendResponse, error)
	// t-Address support
	GetAddressTxids(*TransparentAddressBlockFilter, CompactTxStreamer_GetAddressTxidsServer) error
	// Misc
	GetLightdInfo(context.Context, *Empty) (*LightdInfo, error)
}

// UnimplementedCompactTxStreamerServer can be embedded to have forward compatible implementations.
type UnimplementedCompactTxStreamerServer struct {
}

func (*UnimplementedCompactTxStreamerServer) GetLatestBlock(ctx context.Context, req *ChainSpec) (*BlockID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLatestBlock not implemented")
}
func (*UnimplementedCompactTxStreamerServer) GetBlock(ctx context.Context, req *BlockID) (*CompactBlock, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (*UnimplementedCompactTxStreamerServer) GetBlockRange(req *BlockRange, srv CompactTxStreamer_GetBlockRangeServer) error {
	return status.Errorf(codes.Unimplemented, "method GetBlockRange not implemented")
}
func (*UnimplementedCompactTxStreamerServer) GetTransaction(ctx context.Context, req *TxFilter) (*RawTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (*UnimplementedCompactTxStreamerServer) SendTransaction(ctx context.Context, req *RawTransaction) (*SendResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendTransaction not implemented")
}
func (*UnimplementedCompactTxStreamerServer) GetAddressTxids(req *TransparentAddressBlockFilter, srv CompactTxStreamer_GetAddressTxidsServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAddressTxids not implemented")
}
func (*UnimplementedCompactTxStreamerServer) GetLightdInfo(ctx context.Context, req *Empty) (*LightdInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLightdInfo not implemented")
}

func RegisterCompactTxStreamerServer(s *grpc.Server, srv CompactTxStreamerServer) {
	s.RegisterService(&_CompactTxStreamer_serviceDesc, srv)
}

func _CompactTxStreamer_GetLatestBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChainSpec)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).GetLatestBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetLatestBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).GetLatestBlock(ctx, req.(*ChainSpec))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompactTxStreamer_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).GetBlock(ctx, req.(*BlockID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompactTxStreamer_GetBlockRange_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BlockRange)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CompactTxStreamerServer).GetBlockRange(m, &compactTxStreamerGetBlockRangeServer{stream})
}

type CompactTxStreamer_GetBlockRangeServer interface {
	Send(*CompactBlock) error
	grpc.ServerStream
}

type compactTxStreamerGetBlockRangeServer struct {
	grpc.ServerStream
}

func (x *compactTxStreamerGetBlockRangeServer) Send(m *CompactBlock) error {
	return x.ServerStream.SendMsg(m)
}

func _CompactTxStreamer_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxFilter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).GetTransaction(ctx, req.(*TxFilter))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompactTxStreamer_SendTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RawTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).SendTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cash.z.wallet.sdk.rpc.CompactTxStreamer/SendTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).SendTransaction(ctx, req.(*RawTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompactTxStreamer_GetAddressTxids_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TransparentAddressBlockFilter)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CompactTxStreamerServer).GetAddressTxids(m, &compactTxStreamerGetAddressTxidsServer{stream})
}

type CompactTxStreamer_GetAddressTxidsServer interface {
	Send(*RawTransaction) error
	grpc.ServerStream
}

type compactTxStreamerGetAddressTxidsServer struct {
	grpc.ServerStream
}

func (x *compactTxStreamerGetAddressTxidsServer) Send(m *RawTransaction) error {
	return x.ServerStream.SendMsg(m)
}

func _CompactTxStreamer_GetLightdInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompactTxStreamerServer).GetLightdInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cash.z.wallet.sdk.rpc.CompactTxStreamer/GetLightdInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompactTxStreamerServer).GetLightdInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _CompactTxStreamer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cash.z.wallet.sdk.rpc.CompactTxStreamer",
	HandlerType: (*CompactTxStreamerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetLatestBlock",
			Handler:    _CompactTxStreamer_GetLatestBlock_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _CompactTxStreamer_GetBlock_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _CompactTxStreamer_GetTransaction_Handler,
		},
		{
			MethodName: "SendTransaction",
			Handler:    _CompactTxStreamer_SendTransaction_Handler,
		},
		{
			MethodName: "GetLightdInfo",
			Handler:    _CompactTxStreamer_GetLightdInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetBlockRange",
			Handler:       _CompactTxStreamer_GetBlockRange_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAddressTxids",
			Handler:       _CompactTxStreamer_GetAddressTxids_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "service.proto",
}
