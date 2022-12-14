// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nova/ibcstaking/v1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/types"
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

type GenesisState struct {
	// params defines all the parameters of module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3ed08ccf74c1df4, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

type RegisteredZone struct {
	ZoneId            string             `protobuf:"bytes,1,opt,name=zone_id,json=zoneId,proto3" json:"zone_id,omitempty"`
	IcaConnectionInfo *IcaConnectionInfo `protobuf:"bytes,2,opt,name=ica_connection_info,json=icaConnectionInfo,proto3" json:"ica_connection_info,omitempty"`
	IcaAccount        *IcaAccount        `protobuf:"bytes,3,opt,name=ica_account,json=icaAccount,proto3" json:"ica_account,omitempty"`
	ValidatorAddress  string             `protobuf:"bytes,4,opt,name=validator_address,json=validatorAddress,proto3" json:"validator_address,omitempty"`
	BaseDenom         string             `protobuf:"bytes,5,opt,name=base_denom,json=baseDenom,proto3" json:"base_denom,omitempty"`
	SnDenom           string             `protobuf:"bytes,6,opt,name=sn_denom,json=snDenom,proto3" json:"sn_denom,omitempty"`
}

func (m *RegisteredZone) Reset()         { *m = RegisteredZone{} }
func (m *RegisteredZone) String() string { return proto.CompactTextString(m) }
func (*RegisteredZone) ProtoMessage()    {}
func (*RegisteredZone) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3ed08ccf74c1df4, []int{1}
}
func (m *RegisteredZone) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RegisteredZone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RegisteredZone.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RegisteredZone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegisteredZone.Merge(m, src)
}
func (m *RegisteredZone) XXX_Size() int {
	return m.Size()
}
func (m *RegisteredZone) XXX_DiscardUnknown() {
	xxx_messageInfo_RegisteredZone.DiscardUnknown(m)
}

var xxx_messageInfo_RegisteredZone proto.InternalMessageInfo

func (m *RegisteredZone) GetZoneId() string {
	if m != nil {
		return m.ZoneId
	}
	return ""
}

func (m *RegisteredZone) GetIcaConnectionInfo() *IcaConnectionInfo {
	if m != nil {
		return m.IcaConnectionInfo
	}
	return nil
}

func (m *RegisteredZone) GetIcaAccount() *IcaAccount {
	if m != nil {
		return m.IcaAccount
	}
	return nil
}

func (m *RegisteredZone) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func (m *RegisteredZone) GetBaseDenom() string {
	if m != nil {
		return m.BaseDenom
	}
	return ""
}

func (m *RegisteredZone) GetSnDenom() string {
	if m != nil {
		return m.SnDenom
	}
	return ""
}

type IcaAccount struct {
	DaomodifierAddress string     `protobuf:"bytes,1,opt,name=daomodifier_address,json=daomodifierAddress,proto3" json:"daomodifier_address,omitempty"`
	HostAddress        string     `protobuf:"bytes,2,opt,name=host_address,json=hostAddress,proto3" json:"host_address,omitempty"`
	Balance            types.Coin `protobuf:"bytes,3,opt,name=balance,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"balance"`
}

func (m *IcaAccount) Reset()         { *m = IcaAccount{} }
func (m *IcaAccount) String() string { return proto.CompactTextString(m) }
func (*IcaAccount) ProtoMessage()    {}
func (*IcaAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3ed08ccf74c1df4, []int{2}
}
func (m *IcaAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IcaAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IcaAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IcaAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IcaAccount.Merge(m, src)
}
func (m *IcaAccount) XXX_Size() int {
	return m.Size()
}
func (m *IcaAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_IcaAccount.DiscardUnknown(m)
}

var xxx_messageInfo_IcaAccount proto.InternalMessageInfo

func (m *IcaAccount) GetDaomodifierAddress() string {
	if m != nil {
		return m.DaomodifierAddress
	}
	return ""
}

func (m *IcaAccount) GetHostAddress() string {
	if m != nil {
		return m.HostAddress
	}
	return ""
}

func (m *IcaAccount) GetBalance() types.Coin {
	if m != nil {
		return m.Balance
	}
	return types.Coin{}
}

//zone name, connection id, portID(owner address)
type IcaConnectionInfo struct {
	ConnectionId string `protobuf:"bytes,1,opt,name=connection_id,json=connectionId,proto3" json:"connection_id,omitempty"`
	PortId       string `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
}

func (m *IcaConnectionInfo) Reset()         { *m = IcaConnectionInfo{} }
func (m *IcaConnectionInfo) String() string { return proto.CompactTextString(m) }
func (*IcaConnectionInfo) ProtoMessage()    {}
func (*IcaConnectionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_d3ed08ccf74c1df4, []int{3}
}
func (m *IcaConnectionInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IcaConnectionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IcaConnectionInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IcaConnectionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IcaConnectionInfo.Merge(m, src)
}
func (m *IcaConnectionInfo) XXX_Size() int {
	return m.Size()
}
func (m *IcaConnectionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_IcaConnectionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_IcaConnectionInfo proto.InternalMessageInfo

func (m *IcaConnectionInfo) GetConnectionId() string {
	if m != nil {
		return m.ConnectionId
	}
	return ""
}

func (m *IcaConnectionInfo) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "nova.ibcstaking.v1.GenesisState")
	proto.RegisterType((*RegisteredZone)(nil), "nova.ibcstaking.v1.RegisteredZone")
	proto.RegisterType((*IcaAccount)(nil), "nova.ibcstaking.v1.IcaAccount")
	proto.RegisterType((*IcaConnectionInfo)(nil), "nova.ibcstaking.v1.IcaConnectionInfo")
}

func init() { proto.RegisterFile("nova/ibcstaking/v1/genesis.proto", fileDescriptor_d3ed08ccf74c1df4) }

var fileDescriptor_d3ed08ccf74c1df4 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0x4d, 0x6f, 0x13, 0x31,
	0x10, 0xcd, 0x96, 0x92, 0x50, 0x27, 0x20, 0xe2, 0x22, 0x91, 0x44, 0x62, 0x53, 0x82, 0x90, 0x2a,
	0xa1, 0xae, 0x09, 0x5c, 0xb8, 0xa1, 0x26, 0x48, 0x34, 0x37, 0x58, 0xc4, 0xa5, 0x97, 0x95, 0xd7,
	0x76, 0xb6, 0x56, 0x13, 0x4f, 0xb4, 0x76, 0x23, 0xe0, 0x57, 0xf0, 0x33, 0x10, 0x7f, 0x84, 0x1e,
	0x7b, 0xe4, 0x04, 0x28, 0xf9, 0x23, 0xc8, 0x1f, 0xf9, 0x10, 0x2d, 0xa7, 0xf5, 0xcc, 0x7b, 0xfb,
	0xe6, 0x69, 0xe6, 0xa1, 0x03, 0x05, 0x73, 0x4a, 0x64, 0xce, 0xb4, 0xa1, 0xe7, 0x52, 0x15, 0x64,
	0xde, 0x27, 0x85, 0x50, 0x42, 0x4b, 0x9d, 0xcc, 0x4a, 0x30, 0x80, 0xb1, 0x65, 0x24, 0x1b, 0x46,
	0x32, 0xef, 0x77, 0x1e, 0x14, 0x50, 0x80, 0x83, 0x89, 0x7d, 0x79, 0x66, 0xa7, 0xcd, 0x40, 0x4f,
	0x41, 0x67, 0x1e, 0xf0, 0x45, 0x80, 0x62, 0x5f, 0x91, 0x9c, 0x6a, 0x41, 0xe6, 0xfd, 0x5c, 0x18,
	0xda, 0x27, 0x0c, 0xa4, 0x0a, 0x78, 0xf7, 0x06, 0x1b, 0x33, 0x5a, 0xd2, 0x69, 0x10, 0xe8, 0x9d,
	0xa0, 0xc6, 0x5b, 0x6f, 0xeb, 0x83, 0xa1, 0x46, 0xe0, 0x57, 0xa8, 0xea, 0xf1, 0x56, 0x74, 0x10,
	0x1d, 0xd6, 0x5f, 0x74, 0x92, 0xeb, 0x36, 0x93, 0x77, 0x8e, 0x31, 0xd8, 0xbd, 0xfc, 0xd5, 0xad,
	0xa4, 0x81, 0xdf, 0xfb, 0xb6, 0x83, 0xee, 0xa5, 0xa2, 0x90, 0xda, 0x88, 0x52, 0xf0, 0x53, 0x50,
	0x02, 0x3f, 0x44, 0xb5, 0x2f, 0xa0, 0x44, 0x26, 0xb9, 0x53, 0xdb, 0x4b, 0xab, 0xb6, 0x1c, 0x71,
	0xfc, 0x11, 0xed, 0x4b, 0x46, 0x33, 0x06, 0x4a, 0x09, 0x66, 0x24, 0xa8, 0x4c, 0xaa, 0x31, 0xb4,
	0x76, 0xdc, 0xc8, 0xa7, 0x37, 0x8d, 0x1c, 0x31, 0x3a, 0x5c, 0xb3, 0x47, 0x6a, 0x0c, 0x69, 0x53,
	0xfe, 0xdb, 0xc2, 0xaf, 0x51, 0xdd, 0xca, 0x52, 0xc6, 0xe0, 0x42, 0x99, 0xd6, 0x2d, 0x27, 0x17,
	0xff, 0x47, 0xee, 0xd8, 0xb3, 0x52, 0x24, 0xd7, 0x6f, 0xfc, 0x0c, 0x35, 0xe7, 0x74, 0x22, 0x39,
	0x35, 0x50, 0x66, 0x94, 0xf3, 0x52, 0x68, 0xdd, 0xda, 0x75, 0xd6, 0xef, 0xaf, 0x81, 0x63, 0xdf,
	0xc7, 0x8f, 0x10, 0xb2, 0x6b, 0xcf, 0xb8, 0x50, 0x30, 0x6d, 0xdd, 0x76, 0xac, 0x3d, 0xdb, 0x79,
	0x63, 0x1b, 0xb8, 0x8d, 0xee, 0x68, 0x15, 0xc0, 0xaa, 0x03, 0x6b, 0x5a, 0x39, 0xa8, 0xf7, 0x23,
	0x42, 0x68, 0xe3, 0x00, 0x13, 0xb4, 0xcf, 0x29, 0x4c, 0x81, 0xcb, 0xb1, 0x14, 0x9b, 0xb9, 0x7e,
	0x65, 0x78, 0x0b, 0x5a, 0x4d, 0x7e, 0x8c, 0x1a, 0x67, 0xa0, 0xcd, 0x9a, 0xb9, 0xe3, 0x98, 0x75,
	0xdb, 0x5b, 0x51, 0x04, 0xaa, 0xe5, 0x74, 0x42, 0x15, 0x13, 0x61, 0x0d, 0xed, 0x24, 0x04, 0xc7,
	0x3a, 0x4c, 0x42, 0x54, 0x92, 0x21, 0x48, 0x35, 0x78, 0x6e, 0xef, 0xf8, 0xfd, 0x77, 0xf7, 0xb0,
	0x90, 0xe6, 0xec, 0x22, 0x4f, 0x18, 0x4c, 0x43, 0xca, 0xc2, 0xe7, 0x48, 0xf3, 0x73, 0x62, 0x3e,
	0xcf, 0x84, 0x76, 0x3f, 0xe8, 0x74, 0xa5, 0xdd, 0x7b, 0x8f, 0x9a, 0xd7, 0x2e, 0x83, 0x9f, 0xa0,
	0xbb, 0xdb, 0x97, 0x5d, 0x1d, 0xbf, 0xb1, 0x69, 0x8e, 0xb8, 0xcd, 0xc6, 0x0c, 0x4a, 0x63, 0x61,
	0x6f, 0xbf, 0x6a, 0xcb, 0x11, 0x1f, 0x9c, 0x5c, 0x2e, 0xe2, 0xe8, 0x6a, 0x11, 0x47, 0x7f, 0x16,
	0x71, 0xf4, 0x75, 0x19, 0x57, 0xae, 0x96, 0x71, 0xe5, 0xe7, 0x32, 0xae, 0x9c, 0x26, 0x5b, 0xfe,
	0x86, 0xb4, 0x94, 0x8a, 0x1e, 0x4d, 0x68, 0xae, 0x89, 0xcb, 0xf8, 0xa7, 0xed, 0x94, 0x3b, 0xaf,
	0x79, 0xd5, 0x45, 0xfc, 0xe5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5c, 0xe7, 0xbb, 0xbc, 0x8c,
	0x03, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *RegisteredZone) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RegisteredZone) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RegisteredZone) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SnDenom) > 0 {
		i -= len(m.SnDenom)
		copy(dAtA[i:], m.SnDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.SnDenom)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.BaseDenom) > 0 {
		i -= len(m.BaseDenom)
		copy(dAtA[i:], m.BaseDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.BaseDenom)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x22
	}
	if m.IcaAccount != nil {
		{
			size, err := m.IcaAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.IcaConnectionInfo != nil {
		{
			size, err := m.IcaConnectionInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.ZoneId) > 0 {
		i -= len(m.ZoneId)
		copy(dAtA[i:], m.ZoneId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ZoneId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IcaAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IcaAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IcaAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Balance.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.HostAddress) > 0 {
		i -= len(m.HostAddress)
		copy(dAtA[i:], m.HostAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.HostAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.DaomodifierAddress) > 0 {
		i -= len(m.DaomodifierAddress)
		copy(dAtA[i:], m.DaomodifierAddress)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.DaomodifierAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IcaConnectionInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IcaConnectionInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IcaConnectionInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PortId) > 0 {
		i -= len(m.PortId)
		copy(dAtA[i:], m.PortId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.PortId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ConnectionId) > 0 {
		i -= len(m.ConnectionId)
		copy(dAtA[i:], m.ConnectionId)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ConnectionId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *RegisteredZone) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ZoneId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.IcaConnectionInfo != nil {
		l = m.IcaConnectionInfo.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.IcaAccount != nil {
		l = m.IcaAccount.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.BaseDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.SnDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *IcaAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DaomodifierAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.HostAddress)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = m.Balance.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *IcaConnectionInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ConnectionId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = len(m.PortId)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *RegisteredZone) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: RegisteredZone: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RegisteredZone: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ZoneId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ZoneId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IcaConnectionInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.IcaConnectionInfo == nil {
				m.IcaConnectionInfo = &IcaConnectionInfo{}
			}
			if err := m.IcaConnectionInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IcaAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.IcaAccount == nil {
				m.IcaAccount = &IcaAccount{}
			}
			if err := m.IcaAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SnDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SnDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *IcaAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: IcaAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IcaAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DaomodifierAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DaomodifierAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HostAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HostAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Balance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *IcaConnectionInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: IcaConnectionInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IcaConnectionInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConnectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConnectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PortId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PortId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
