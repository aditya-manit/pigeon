// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: paloma/consensus/consensus_queue.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// message for storing the queued signed message in the internal queue
type QueuedSignedMessage struct {
	Id          uint64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Msg         *types.Any  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	BytesToSign []byte      `protobuf:"bytes,3,opt,name=bytesToSign,proto3" json:"bytesToSign,omitempty"`
	SignData    []*SignData `protobuf:"bytes,4,rep,name=signData,proto3" json:"signData,omitempty"`
}

func (m *QueuedSignedMessage) Reset()         { *m = QueuedSignedMessage{} }
func (m *QueuedSignedMessage) String() string { return proto.CompactTextString(m) }
func (*QueuedSignedMessage) ProtoMessage()    {}
func (*QueuedSignedMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_7361456ceb9bcf74, []int{0}
}
func (m *QueuedSignedMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueuedSignedMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueuedSignedMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueuedSignedMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueuedSignedMessage.Merge(m, src)
}
func (m *QueuedSignedMessage) XXX_Size() int {
	return m.Size()
}
func (m *QueuedSignedMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_QueuedSignedMessage.DiscardUnknown(m)
}

var xxx_messageInfo_QueuedSignedMessage proto.InternalMessageInfo

func (m *QueuedSignedMessage) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *QueuedSignedMessage) GetMsg() *types.Any {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *QueuedSignedMessage) GetBytesToSign() []byte {
	if m != nil {
		return m.BytesToSign
	}
	return nil
}

func (m *QueuedSignedMessage) GetSignData() []*SignData {
	if m != nil {
		return m.SignData
	}
	return nil
}

type BatchOfConsensusMessages struct {
	Msg *types.Any `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *BatchOfConsensusMessages) Reset()         { *m = BatchOfConsensusMessages{} }
func (m *BatchOfConsensusMessages) String() string { return proto.CompactTextString(m) }
func (*BatchOfConsensusMessages) ProtoMessage()    {}
func (*BatchOfConsensusMessages) Descriptor() ([]byte, []int) {
	return fileDescriptor_7361456ceb9bcf74, []int{1}
}
func (m *BatchOfConsensusMessages) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BatchOfConsensusMessages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BatchOfConsensusMessages.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BatchOfConsensusMessages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchOfConsensusMessages.Merge(m, src)
}
func (m *BatchOfConsensusMessages) XXX_Size() int {
	return m.Size()
}
func (m *BatchOfConsensusMessages) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchOfConsensusMessages.DiscardUnknown(m)
}

var xxx_messageInfo_BatchOfConsensusMessages proto.InternalMessageInfo

func (m *BatchOfConsensusMessages) GetMsg() *types.Any {
	if m != nil {
		return m.Msg
	}
	return nil
}

type Batch struct {
	Msgs        []*types.Any `protobuf:"bytes,1,rep,name=msgs,proto3" json:"msgs,omitempty"`
	BytesToSign []byte       `protobuf:"bytes,2,opt,name=bytesToSign,proto3" json:"bytesToSign,omitempty"`
}

func (m *Batch) Reset()         { *m = Batch{} }
func (m *Batch) String() string { return proto.CompactTextString(m) }
func (*Batch) ProtoMessage()    {}
func (*Batch) Descriptor() ([]byte, []int) {
	return fileDescriptor_7361456ceb9bcf74, []int{2}
}
func (m *Batch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Batch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Batch.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Batch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Batch.Merge(m, src)
}
func (m *Batch) XXX_Size() int {
	return m.Size()
}
func (m *Batch) XXX_DiscardUnknown() {
	xxx_messageInfo_Batch.DiscardUnknown(m)
}

var xxx_messageInfo_Batch proto.InternalMessageInfo

func (m *Batch) GetMsgs() []*types.Any {
	if m != nil {
		return m.Msgs
	}
	return nil
}

func (m *Batch) GetBytesToSign() []byte {
	if m != nil {
		return m.BytesToSign
	}
	return nil
}

type SignData struct {
	ValAddress             github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,1,opt,name=valAddress,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"valAddress,omitempty"`
	Signature              []byte                                        `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	ExtraData              []byte                                        `protobuf:"bytes,3,opt,name=extraData,proto3" json:"extraData,omitempty"`
	ExternalAccountAddress string                                        `protobuf:"bytes,4,opt,name=externalAccountAddress,proto3" json:"externalAccountAddress,omitempty"`
}

func (m *SignData) Reset()         { *m = SignData{} }
func (m *SignData) String() string { return proto.CompactTextString(m) }
func (*SignData) ProtoMessage()    {}
func (*SignData) Descriptor() ([]byte, []int) {
	return fileDescriptor_7361456ceb9bcf74, []int{3}
}
func (m *SignData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignData.Merge(m, src)
}
func (m *SignData) XXX_Size() int {
	return m.Size()
}
func (m *SignData) XXX_DiscardUnknown() {
	xxx_messageInfo_SignData.DiscardUnknown(m)
}

var xxx_messageInfo_SignData proto.InternalMessageInfo

func (m *SignData) GetValAddress() github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.ValAddress
	}
	return nil
}

func (m *SignData) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *SignData) GetExtraData() []byte {
	if m != nil {
		return m.ExtraData
	}
	return nil
}

func (m *SignData) GetExternalAccountAddress() string {
	if m != nil {
		return m.ExternalAccountAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*QueuedSignedMessage)(nil), "volumefi.paloma.consensus.QueuedSignedMessage")
	proto.RegisterType((*BatchOfConsensusMessages)(nil), "volumefi.paloma.consensus.BatchOfConsensusMessages")
	proto.RegisterType((*Batch)(nil), "volumefi.paloma.consensus.Batch")
	proto.RegisterType((*SignData)(nil), "volumefi.paloma.consensus.SignData")
}

func init() {
	proto.RegisterFile("paloma/consensus/consensus_queue.proto", fileDescriptor_7361456ceb9bcf74)
}

var fileDescriptor_7361456ceb9bcf74 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x8e, 0xd3, 0x30,
	0x10, 0xad, 0xdb, 0x82, 0x76, 0xdd, 0x15, 0x07, 0xb3, 0x42, 0xd9, 0x15, 0x0a, 0x51, 0x91, 0x56,
	0xb9, 0xac, 0x2d, 0x16, 0x89, 0x2b, 0x6a, 0xe1, 0xc2, 0x01, 0xa1, 0x4d, 0x11, 0x07, 0x2e, 0xc8,
	0x4d, 0x5c, 0x37, 0x22, 0xb1, 0x4b, 0xc6, 0x5e, 0x35, 0x7f, 0xc1, 0xa7, 0xf0, 0x19, 0x1c, 0x97,
	0x1b, 0x27, 0x84, 0xda, 0xbf, 0xe0, 0x84, 0xe2, 0x24, 0x6d, 0x04, 0x54, 0x9c, 0x32, 0x9a, 0x79,
	0x6f, 0xe6, 0xbd, 0x17, 0xe3, 0x8b, 0x15, 0xcf, 0x74, 0xce, 0x59, 0xac, 0x15, 0x08, 0x05, 0x16,
	0xf6, 0xd5, 0x87, 0x4f, 0x56, 0x58, 0x41, 0x57, 0x85, 0x36, 0x9a, 0x9c, 0xdd, 0xe8, 0xcc, 0xe6,
	0x62, 0x91, 0xd2, 0x9a, 0x40, 0x77, 0xb0, 0xf3, 0x33, 0xa9, 0xb5, 0xcc, 0x04, 0x73, 0xc0, 0xb9,
	0x5d, 0x30, 0xae, 0xca, 0x9a, 0x75, 0x7e, 0x2a, 0xb5, 0xd4, 0xae, 0x64, 0x55, 0x55, 0x77, 0xc7,
	0x5f, 0x10, 0xbe, 0x7f, 0x5d, 0xed, 0x4e, 0x66, 0xa9, 0x54, 0x22, 0x79, 0x2d, 0x00, 0xb8, 0x14,
	0xe4, 0x1e, 0xee, 0xa7, 0x89, 0x87, 0x02, 0x14, 0x0e, 0xa3, 0x7e, 0x9a, 0x90, 0x0b, 0x3c, 0xc8,
	0x41, 0x7a, 0xfd, 0x00, 0x85, 0xa3, 0xab, 0x53, 0x5a, 0x9f, 0xa1, 0xed, 0x19, 0x3a, 0x51, 0x65,
	0x54, 0x01, 0x48, 0x80, 0x47, 0xf3, 0xd2, 0x08, 0x78, 0xab, 0xab, 0x7d, 0xde, 0x20, 0x40, 0xe1,
	0x49, 0xd4, 0x6d, 0x91, 0xe7, 0xf8, 0x08, 0x52, 0xa9, 0x5e, 0x72, 0xc3, 0xbd, 0x61, 0x30, 0x08,
	0x47, 0x57, 0x8f, 0xe9, 0x41, 0x43, 0x74, 0xd6, 0x40, 0xa3, 0x1d, 0x69, 0x3c, 0xc5, 0xde, 0x94,
	0x9b, 0x78, 0xf9, 0x66, 0xf1, 0xa2, 0x85, 0x35, 0xaa, 0xa1, 0x95, 0x89, 0xfe, 0x23, 0x73, 0x3c,
	0xc3, 0x77, 0xdc, 0x0e, 0x12, 0xe2, 0x61, 0x0e, 0x12, 0x3c, 0xe4, 0x94, 0xfc, 0x9b, 0xe1, 0x10,
	0x7f, 0x3a, 0xeb, 0xff, 0xe5, 0x6c, 0xfc, 0x0d, 0xe1, 0xa3, 0x56, 0x2f, 0xb9, 0xc6, 0xf8, 0x86,
	0x67, 0x93, 0x24, 0x29, 0x04, 0x80, 0x13, 0x74, 0x32, 0x7d, 0xf2, 0xeb, 0xc7, 0xa3, 0x4b, 0x99,
	0x9a, 0xa5, 0x9d, 0xd3, 0x58, 0xe7, 0x2c, 0xd6, 0x90, 0x6b, 0x68, 0x3e, 0x97, 0x90, 0x7c, 0x64,
	0xa6, 0x5c, 0x09, 0xa0, 0xef, 0x76, 0xc4, 0xa8, 0xb3, 0x84, 0x3c, 0xc4, 0xc7, 0x55, 0x08, 0xdc,
	0xd8, 0x42, 0x34, 0xf7, 0xf7, 0x8d, 0x6a, 0x2a, 0xd6, 0xa6, 0xe0, 0x2e, 0xd8, 0x3a, 0xf7, 0x7d,
	0x83, 0x3c, 0xc3, 0x0f, 0xc4, 0xda, 0x88, 0x42, 0xf1, 0x6c, 0x12, 0xc7, 0xda, 0x2a, 0xd3, 0x4a,
	0x1b, 0x06, 0x28, 0x3c, 0x8e, 0x0e, 0x4c, 0xa7, 0xaf, 0xbe, 0x6e, 0x7c, 0x74, 0xbb, 0xf1, 0xd1,
	0xcf, 0x8d, 0x8f, 0x3e, 0x6f, 0xfd, 0xde, 0xed, 0xd6, 0xef, 0x7d, 0xdf, 0xfa, 0xbd, 0xf7, 0xac,
	0x63, 0xa4, 0xfe, 0x6d, 0xf1, 0x92, 0xa7, 0xaa, 0xa9, 0xd9, 0xba, 0xf3, 0x8c, 0x9d, 0xab, 0xf9,
	0x5d, 0x17, 0xea, 0xd3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x67, 0x30, 0x5e, 0xbe, 0xe7, 0x02,
	0x00, 0x00,
}

func (m *QueuedSignedMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueuedSignedMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueuedSignedMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SignData) > 0 {
		for iNdEx := len(m.SignData) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SignData[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintConsensusQueue(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.BytesToSign) > 0 {
		i -= len(m.BytesToSign)
		copy(dAtA[i:], m.BytesToSign)
		i = encodeVarintConsensusQueue(dAtA, i, uint64(len(m.BytesToSign)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Msg != nil {
		{
			size, err := m.Msg.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConsensusQueue(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintConsensusQueue(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BatchOfConsensusMessages) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BatchOfConsensusMessages) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BatchOfConsensusMessages) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Msg != nil {
		{
			size, err := m.Msg.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintConsensusQueue(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Batch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Batch) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Batch) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BytesToSign) > 0 {
		i -= len(m.BytesToSign)
		copy(dAtA[i:], m.BytesToSign)
		i = encodeVarintConsensusQueue(dAtA, i, uint64(len(m.BytesToSign)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Msgs) > 0 {
		for iNdEx := len(m.Msgs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Msgs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintConsensusQueue(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *SignData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ExternalAccountAddress) > 0 {
		i -= len(m.ExternalAccountAddress)
		copy(dAtA[i:], m.ExternalAccountAddress)
		i = encodeVarintConsensusQueue(dAtA, i, uint64(len(m.ExternalAccountAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ExtraData) > 0 {
		i -= len(m.ExtraData)
		copy(dAtA[i:], m.ExtraData)
		i = encodeVarintConsensusQueue(dAtA, i, uint64(len(m.ExtraData)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintConsensusQueue(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ValAddress) > 0 {
		i -= len(m.ValAddress)
		copy(dAtA[i:], m.ValAddress)
		i = encodeVarintConsensusQueue(dAtA, i, uint64(len(m.ValAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConsensusQueue(dAtA []byte, offset int, v uint64) int {
	offset -= sovConsensusQueue(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueuedSignedMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovConsensusQueue(uint64(m.Id))
	}
	if m.Msg != nil {
		l = m.Msg.Size()
		n += 1 + l + sovConsensusQueue(uint64(l))
	}
	l = len(m.BytesToSign)
	if l > 0 {
		n += 1 + l + sovConsensusQueue(uint64(l))
	}
	if len(m.SignData) > 0 {
		for _, e := range m.SignData {
			l = e.Size()
			n += 1 + l + sovConsensusQueue(uint64(l))
		}
	}
	return n
}

func (m *BatchOfConsensusMessages) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Msg != nil {
		l = m.Msg.Size()
		n += 1 + l + sovConsensusQueue(uint64(l))
	}
	return n
}

func (m *Batch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Msgs) > 0 {
		for _, e := range m.Msgs {
			l = e.Size()
			n += 1 + l + sovConsensusQueue(uint64(l))
		}
	}
	l = len(m.BytesToSign)
	if l > 0 {
		n += 1 + l + sovConsensusQueue(uint64(l))
	}
	return n
}

func (m *SignData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValAddress)
	if l > 0 {
		n += 1 + l + sovConsensusQueue(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovConsensusQueue(uint64(l))
	}
	l = len(m.ExtraData)
	if l > 0 {
		n += 1 + l + sovConsensusQueue(uint64(l))
	}
	l = len(m.ExternalAccountAddress)
	if l > 0 {
		n += 1 + l + sovConsensusQueue(uint64(l))
	}
	return n
}

func sovConsensusQueue(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConsensusQueue(x uint64) (n int) {
	return sovConsensusQueue(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueuedSignedMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConsensusQueue
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueuedSignedMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueuedSignedMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensusQueue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensusQueue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Msg == nil {
				m.Msg = &types.Any{}
			}
			if err := m.Msg.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BytesToSign", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensusQueue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BytesToSign = append(m.BytesToSign[:0], dAtA[iNdEx:postIndex]...)
			if m.BytesToSign == nil {
				m.BytesToSign = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignData", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensusQueue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignData = append(m.SignData, &SignData{})
			if err := m.SignData[len(m.SignData)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConsensusQueue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BatchOfConsensusMessages) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConsensusQueue
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BatchOfConsensusMessages: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BatchOfConsensusMessages: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensusQueue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Msg == nil {
				m.Msg = &types.Any{}
			}
			if err := m.Msg.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConsensusQueue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Batch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConsensusQueue
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Batch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Batch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msgs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensusQueue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Msgs = append(m.Msgs, &types.Any{})
			if err := m.Msgs[len(m.Msgs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BytesToSign", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensusQueue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BytesToSign = append(m.BytesToSign[:0], dAtA[iNdEx:postIndex]...)
			if m.BytesToSign == nil {
				m.BytesToSign = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConsensusQueue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *SignData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConsensusQueue
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SignData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensusQueue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValAddress = append(m.ValAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.ValAddress == nil {
				m.ValAddress = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensusQueue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = append(m.Signature[:0], dAtA[iNdEx:postIndex]...)
			if m.Signature == nil {
				m.Signature = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExtraData", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensusQueue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExtraData = append(m.ExtraData[:0], dAtA[iNdEx:postIndex]...)
			if m.ExtraData == nil {
				m.ExtraData = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalAccountAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensusQueue
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExternalAccountAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConsensusQueue(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConsensusQueue
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipConsensusQueue(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConsensusQueue
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConsensusQueue
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowConsensusQueue
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthConsensusQueue
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConsensusQueue
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConsensusQueue
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConsensusQueue        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConsensusQueue          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConsensusQueue = fmt.Errorf("proto: unexpected end of group")
)
