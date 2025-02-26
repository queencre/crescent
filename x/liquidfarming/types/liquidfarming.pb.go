// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: crescent/liquidfarming/v1beta1/liquidfarming.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/regen-network/cosmos-proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// AuctionStatus enumerates the valid status of an auction.
type AuctionStatus int32

const (
	// AUCTION_STATUS_UNSPECIFIED defines the default auction status
	AuctionStatusNil AuctionStatus = 0
	// AUCTION_STATUS_STARTED defines the started auction status
	AuctionStatusStarted AuctionStatus = 1
	// AUCTION_STATUS_FINISHED defines the finished auction status
	AuctionStatusFinished AuctionStatus = 2
	// AUCTION_STATUS_SKIPPED defines the skipped auction status
	AuctionStatusSkipped AuctionStatus = 3
)

var AuctionStatus_name = map[int32]string{
	0: "AUCTION_STATUS_UNSPECIFIED",
	1: "AUCTION_STATUS_STARTED",
	2: "AUCTION_STATUS_FINISHED",
	3: "AUCTION_STATUS_SKIPPED",
}

var AuctionStatus_value = map[string]int32{
	"AUCTION_STATUS_UNSPECIFIED": 0,
	"AUCTION_STATUS_STARTED":     1,
	"AUCTION_STATUS_FINISHED":    2,
	"AUCTION_STATUS_SKIPPED":     3,
}

func (x AuctionStatus) String() string {
	return proto.EnumName(AuctionStatus_name, int32(x))
}

func (AuctionStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c85a706fbdcf4344, []int{0}
}

// RewardsAuction defines rewards auction that is created by the module
// for every rewards_auction_duration in params.
type RewardsAuction struct {
	// id specifies the unique auction id
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// pool_id specifies the liquidity pool id
	PoolId uint64 `protobuf:"varint,2,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	// bidding_coin_denom specifies the bidding coin denomination
	BiddingCoinDenom string `protobuf:"bytes,3,opt,name=bidding_coin_denom,json=biddingCoinDenom,proto3" json:"bidding_coin_denom,omitempty"`
	// paying_reserve_address specfies the account that reserves bidding amounts placed by bidders
	PayingReserveAddress string `protobuf:"bytes,4,opt,name=paying_reserve_address,json=payingReserveAddress,proto3" json:"paying_reserve_address,omitempty"`
	// start_time specifies the start time of an auction
	StartTime time.Time `protobuf:"bytes,5,opt,name=start_time,json=startTime,proto3,stdtime" json:"start_time"`
	// end_time specifies the end time of an auction
	EndTime time.Time `protobuf:"bytes,6,opt,name=end_time,json=endTime,proto3,stdtime" json:"end_time"`
	// status specifies the status of an auction
	Status AuctionStatus `protobuf:"varint,7,opt,name=status,proto3,enum=crescent.liquidfarming.v1beta1.AuctionStatus" json:"status,omitempty"`
	// winner specifies the bidder who won the auction
	// the value is determined when an auction is finished
	Winner string `protobuf:"bytes,8,opt,name=winner,proto3" json:"winner,omitempty"`
	// winning_amount specifies the winning amount for the uaction
	WinningAmount types.Coin `protobuf:"bytes,9,opt,name=winning_amount,json=winningAmount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"winning_amount"`
	// rewards specifies the farming rewards for are accumulated in the farm module
	// the value is determined when an auction is finished
	Rewards github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,10,rep,name=rewards,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"rewards"`
	Fees    github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,11,rep,name=fees,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"fees"`
	FeeRate github_com_cosmos_cosmos_sdk_types.Dec   `protobuf:"bytes,12,opt,name=fee_rate,json=feeRate,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"fee_rate"`
}

func (m *RewardsAuction) Reset()         { *m = RewardsAuction{} }
func (m *RewardsAuction) String() string { return proto.CompactTextString(m) }
func (*RewardsAuction) ProtoMessage()    {}
func (*RewardsAuction) Descriptor() ([]byte, []int) {
	return fileDescriptor_c85a706fbdcf4344, []int{0}
}
func (m *RewardsAuction) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RewardsAuction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RewardsAuction.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RewardsAuction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardsAuction.Merge(m, src)
}
func (m *RewardsAuction) XXX_Size() int {
	return m.Size()
}
func (m *RewardsAuction) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardsAuction.DiscardUnknown(m)
}

var xxx_messageInfo_RewardsAuction proto.InternalMessageInfo

// CompoundingRewards records the amount of pool coin that is used for a bidder to place a bid
// for an auction. It is used internally to calculate unfarm amount.
type CompoundingRewards struct {
	Amount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,1,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
}

func (m *CompoundingRewards) Reset()         { *m = CompoundingRewards{} }
func (m *CompoundingRewards) String() string { return proto.CompactTextString(m) }
func (*CompoundingRewards) ProtoMessage()    {}
func (*CompoundingRewards) Descriptor() ([]byte, []int) {
	return fileDescriptor_c85a706fbdcf4344, []int{1}
}
func (m *CompoundingRewards) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CompoundingRewards) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CompoundingRewards.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CompoundingRewards) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompoundingRewards.Merge(m, src)
}
func (m *CompoundingRewards) XXX_Size() int {
	return m.Size()
}
func (m *CompoundingRewards) XXX_DiscardUnknown() {
	xxx_messageInfo_CompoundingRewards.DiscardUnknown(m)
}

var xxx_messageInfo_CompoundingRewards proto.InternalMessageInfo

// Bid defines standard bid for a rewards auction.
type Bid struct {
	// pool_id specifies the pool id
	PoolId uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	// bidder specifies the bech32-encoded address that places a bid for the auction
	Bidder string `protobuf:"bytes,2,opt,name=bidder,proto3" json:"bidder,omitempty"`
	// amount specifies the amount to place a bid
	Amount types.Coin `protobuf:"bytes,3,opt,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount"`
}

func (m *Bid) Reset()         { *m = Bid{} }
func (m *Bid) String() string { return proto.CompactTextString(m) }
func (*Bid) ProtoMessage()    {}
func (*Bid) Descriptor() ([]byte, []int) {
	return fileDescriptor_c85a706fbdcf4344, []int{2}
}
func (m *Bid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Bid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Bid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Bid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Bid.Merge(m, src)
}
func (m *Bid) XXX_Size() int {
	return m.Size()
}
func (m *Bid) XXX_DiscardUnknown() {
	xxx_messageInfo_Bid.DiscardUnknown(m)
}

var xxx_messageInfo_Bid proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("crescent.liquidfarming.v1beta1.AuctionStatus", AuctionStatus_name, AuctionStatus_value)
	proto.RegisterType((*RewardsAuction)(nil), "crescent.liquidfarming.v1beta1.RewardsAuction")
	proto.RegisterType((*CompoundingRewards)(nil), "crescent.liquidfarming.v1beta1.CompoundingRewards")
	proto.RegisterType((*Bid)(nil), "crescent.liquidfarming.v1beta1.Bid")
}

func init() {
	proto.RegisterFile("crescent/liquidfarming/v1beta1/liquidfarming.proto", fileDescriptor_c85a706fbdcf4344)
}

var fileDescriptor_c85a706fbdcf4344 = []byte{
	// 754 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xcf, 0x6a, 0xe3, 0x46,
	0x1c, 0x96, 0x62, 0x57, 0x4e, 0x66, 0xbb, 0xc1, 0x0c, 0x6e, 0x56, 0xd1, 0x41, 0x16, 0x39, 0xb4,
	0xa6, 0x34, 0x52, 0x37, 0x0d, 0x3d, 0x14, 0x4a, 0xf1, 0x5f, 0x2a, 0x0a, 0x6e, 0x90, 0x9d, 0x4b,
	0x7b, 0x10, 0x92, 0x67, 0xec, 0x1d, 0x62, 0xcd, 0x68, 0x35, 0xe3, 0xa4, 0xfb, 0x06, 0x21, 0xa7,
	0x7d, 0x81, 0xc0, 0x42, 0x6f, 0x7d, 0x92, 0x1c, 0xf7, 0x58, 0x7a, 0xd8, 0x6d, 0x93, 0x07, 0xe8,
	0x2b, 0x94, 0x19, 0x8d, 0x4d, 0x15, 0x96, 0x65, 0x17, 0x72, 0x8a, 0x7e, 0xf3, 0xcd, 0xf7, 0x7d,
	0xbf, 0x3f, 0xf3, 0x8b, 0xc1, 0xd1, 0xac, 0xc0, 0x7c, 0x86, 0xa9, 0x08, 0x96, 0xe4, 0xf9, 0x8a,
	0xa0, 0x79, 0x52, 0x64, 0x84, 0x2e, 0x82, 0xf3, 0xa7, 0x29, 0x16, 0xc9, 0xd3, 0xea, 0xa9, 0x9f,
	0x17, 0x4c, 0x30, 0xe8, 0xae, 0x39, 0x7e, 0x15, 0xd5, 0x1c, 0xa7, 0xb5, 0x60, 0x0b, 0xa6, 0xae,
	0x06, 0xf2, 0xab, 0x64, 0x39, 0xfb, 0x33, 0xc6, 0x33, 0xc6, 0xe3, 0x12, 0x28, 0x03, 0x0d, 0xb9,
	0x65, 0x14, 0xa4, 0x09, 0xc7, 0x1b, 0xe7, 0x19, 0x23, 0x54, 0xe3, 0xed, 0x05, 0x63, 0x8b, 0x25,
	0x0e, 0x54, 0x94, 0xae, 0xe6, 0x81, 0x20, 0x19, 0xe6, 0x22, 0xc9, 0xf2, 0xf2, 0xc2, 0xc1, 0xa5,
	0x05, 0x76, 0x23, 0x7c, 0x91, 0x14, 0x88, 0x77, 0x57, 0x33, 0x41, 0x18, 0x85, 0xbb, 0x60, 0x8b,
	0x20, 0xdb, 0xf4, 0xcc, 0x4e, 0x3d, 0xda, 0x22, 0x08, 0x3e, 0x01, 0x8d, 0x9c, 0xb1, 0x65, 0x4c,
	0x90, 0xbd, 0xa5, 0x0e, 0x2d, 0x19, 0x86, 0x08, 0x7e, 0x05, 0x60, 0x4a, 0x10, 0x22, 0x74, 0x11,
	0x4b, 0xcb, 0x18, 0x61, 0xca, 0x32, 0xbb, 0xe6, 0x99, 0x9d, 0x9d, 0xa8, 0xa9, 0x91, 0x3e, 0x23,
	0x74, 0x20, 0xcf, 0xe1, 0x31, 0xd8, 0xcb, 0x93, 0x17, 0xf2, 0x72, 0x81, 0x39, 0x2e, 0xce, 0x71,
	0x9c, 0x20, 0x54, 0x60, 0xce, 0xed, 0xba, 0x62, 0xb4, 0x4a, 0x34, 0x2a, 0xc1, 0x6e, 0x89, 0xc1,
	0x3e, 0x00, 0x5c, 0x24, 0x85, 0x88, 0x65, 0xe2, 0xf6, 0x27, 0x9e, 0xd9, 0x79, 0x74, 0xe4, 0xf8,
	0x65, 0x55, 0xfe, 0xba, 0x2a, 0x7f, 0xba, 0xae, 0xaa, 0xb7, 0x7d, 0xf3, 0xa6, 0x6d, 0xbc, 0x7c,
	0xdb, 0x36, 0xa3, 0x1d, 0xc5, 0x93, 0x08, 0xfc, 0x01, 0x6c, 0x63, 0x8a, 0x4a, 0x09, 0xeb, 0x23,
	0x24, 0x1a, 0x98, 0x22, 0x25, 0x30, 0x04, 0x16, 0x17, 0x89, 0x58, 0x71, 0xbb, 0xe1, 0x99, 0x9d,
	0xdd, 0xa3, 0x43, 0xff, 0xfd, 0x83, 0xf4, 0x75, 0x2f, 0x27, 0x8a, 0x14, 0x69, 0x32, 0xdc, 0x03,
	0xd6, 0x05, 0xa1, 0x14, 0x17, 0xf6, 0xb6, 0x2a, 0x59, 0x47, 0xf0, 0x39, 0xd8, 0x95, 0x5f, 0xb2,
	0x37, 0x49, 0xc6, 0x56, 0x54, 0xd8, 0x3b, 0x2a, 0xcb, 0x7d, 0x5f, 0x0f, 0x5b, 0x8e, 0x77, 0xa3,
	0x2d, 0x5b, 0xda, 0x0b, 0x64, 0x92, 0x7f, 0xbc, 0x6d, 0x7f, 0xb1, 0x20, 0xe2, 0xd9, 0x2a, 0xf5,
	0x67, 0x2c, 0xd3, 0x2f, 0x43, 0xff, 0x39, 0xe4, 0xe8, 0x2c, 0x10, 0x2f, 0x72, 0xcc, 0x15, 0x21,
	0x7a, 0xac, 0x1d, 0xba, 0xca, 0x00, 0x62, 0xd0, 0x28, 0xca, 0xb1, 0xdb, 0xc0, 0xab, 0xbd, 0xdf,
	0xeb, 0x6b, 0xed, 0xd5, 0xf9, 0x40, 0x2f, 0x1e, 0xad, 0xb5, 0x61, 0x0c, 0xea, 0x73, 0x8c, 0xb9,
	0xfd, 0xe8, 0xe1, 0x3d, 0x94, 0x30, 0x0c, 0xc1, 0xf6, 0x1c, 0xe3, 0xb8, 0x48, 0x04, 0xb6, 0x3f,
	0x95, 0x4d, 0xed, 0xf9, 0x52, 0xe9, 0xaf, 0x37, 0xed, 0xcf, 0x3f, 0x40, 0x69, 0x80, 0x67, 0x51,
	0x63, 0x8e, 0x71, 0x94, 0x08, 0x7c, 0x90, 0x02, 0xd8, 0x67, 0x59, 0xce, 0x56, 0x14, 0xa9, 0x77,
	0x58, 0x56, 0x30, 0x02, 0x96, 0x9e, 0x89, 0xf9, 0xd1, 0xf2, 0x21, 0x15, 0x91, 0x66, 0x7f, 0x57,
	0xbf, 0x7c, 0xd5, 0x36, 0x0e, 0x5e, 0x99, 0xa0, 0xd6, 0xab, 0xee, 0x94, 0x59, 0xd9, 0xa9, 0x3d,
	0x60, 0xc9, 0xcd, 0xc1, 0x85, 0xda, 0xb5, 0x9d, 0x48, 0x47, 0x30, 0xdd, 0xa4, 0x51, 0x7b, 0xf0,
	0xa7, 0x51, 0x49, 0xf1, 0xcb, 0x7f, 0x4d, 0xf0, 0xb8, 0xf2, 0x7c, 0xe1, 0x31, 0x70, 0xba, 0xa7,
	0xfd, 0x69, 0xf8, 0xf3, 0x38, 0x9e, 0x4c, 0xbb, 0xd3, 0xd3, 0x49, 0x7c, 0x3a, 0x9e, 0x9c, 0x0c,
	0xfb, 0xe1, 0x28, 0x1c, 0x0e, 0x9a, 0x86, 0xd3, 0xba, 0xba, 0xf6, 0x9a, 0x15, 0xca, 0x98, 0x2c,
	0xe5, 0xbe, 0xdf, 0x63, 0x4d, 0xa6, 0xdd, 0x68, 0x3a, 0x1c, 0x34, 0x4d, 0xc7, 0xbe, 0xba, 0xf6,
	0x5a, 0x15, 0xc6, 0x44, 0x2e, 0x2b, 0x46, 0xf0, 0x5b, 0xf0, 0xe4, 0x1e, 0x6b, 0x14, 0x8e, 0xc3,
	0xc9, 0x8f, 0xc3, 0x41, 0x73, 0xcb, 0xd9, 0xbf, 0xba, 0xf6, 0x3e, 0xab, 0xd0, 0x46, 0x84, 0x12,
	0xfe, 0x0c, 0xa3, 0x77, 0xb9, 0xfd, 0x14, 0x9e, 0x9c, 0x0c, 0x07, 0xcd, 0xda, 0xbb, 0xdc, 0xce,
	0x48, 0x9e, 0x63, 0xe4, 0xd4, 0x2f, 0x7f, 0x77, 0x8d, 0xde, 0xaf, 0x37, 0xff, 0xb8, 0xc6, 0xcd,
	0xad, 0x6b, 0xbe, 0xbe, 0x75, 0xcd, 0xbf, 0x6f, 0x5d, 0xf3, 0xe5, 0x9d, 0x6b, 0xbc, 0xbe, 0x73,
	0x8d, 0x3f, 0xef, 0x5c, 0xe3, 0x97, 0xef, 0xff, 0xdf, 0x46, 0xbd, 0xf5, 0x87, 0x14, 0x8b, 0x0b,
	0x56, 0x9c, 0x6d, 0x0e, 0x82, 0xf3, 0xe3, 0xe0, 0xb7, 0x7b, 0x3f, 0x04, 0xaa, 0xc3, 0xa9, 0xa5,
	0xfe, 0xc3, 0x7c, 0xf3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0x6d, 0xfd, 0x41, 0x2f, 0x06,
	0x00, 0x00,
}

func (m *RewardsAuction) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RewardsAuction) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RewardsAuction) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.FeeRate.Size()
		i -= size
		if _, err := m.FeeRate.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLiquidfarming(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x62
	if len(m.Fees) > 0 {
		for iNdEx := len(m.Fees) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fees[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLiquidfarming(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.Rewards) > 0 {
		for iNdEx := len(m.Rewards) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rewards[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintLiquidfarming(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	{
		size, err := m.WinningAmount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLiquidfarming(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	if len(m.Winner) > 0 {
		i -= len(m.Winner)
		copy(dAtA[i:], m.Winner)
		i = encodeVarintLiquidfarming(dAtA, i, uint64(len(m.Winner)))
		i--
		dAtA[i] = 0x42
	}
	if m.Status != 0 {
		i = encodeVarintLiquidfarming(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x38
	}
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.EndTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintLiquidfarming(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x32
	n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.StartTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime):])
	if err3 != nil {
		return 0, err3
	}
	i -= n3
	i = encodeVarintLiquidfarming(dAtA, i, uint64(n3))
	i--
	dAtA[i] = 0x2a
	if len(m.PayingReserveAddress) > 0 {
		i -= len(m.PayingReserveAddress)
		copy(dAtA[i:], m.PayingReserveAddress)
		i = encodeVarintLiquidfarming(dAtA, i, uint64(len(m.PayingReserveAddress)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.BiddingCoinDenom) > 0 {
		i -= len(m.BiddingCoinDenom)
		copy(dAtA[i:], m.BiddingCoinDenom)
		i = encodeVarintLiquidfarming(dAtA, i, uint64(len(m.BiddingCoinDenom)))
		i--
		dAtA[i] = 0x1a
	}
	if m.PoolId != 0 {
		i = encodeVarintLiquidfarming(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintLiquidfarming(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CompoundingRewards) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CompoundingRewards) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CompoundingRewards) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLiquidfarming(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Bid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Bid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Bid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintLiquidfarming(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Bidder) > 0 {
		i -= len(m.Bidder)
		copy(dAtA[i:], m.Bidder)
		i = encodeVarintLiquidfarming(dAtA, i, uint64(len(m.Bidder)))
		i--
		dAtA[i] = 0x12
	}
	if m.PoolId != 0 {
		i = encodeVarintLiquidfarming(dAtA, i, uint64(m.PoolId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLiquidfarming(dAtA []byte, offset int, v uint64) int {
	offset -= sovLiquidfarming(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RewardsAuction) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovLiquidfarming(uint64(m.Id))
	}
	if m.PoolId != 0 {
		n += 1 + sovLiquidfarming(uint64(m.PoolId))
	}
	l = len(m.BiddingCoinDenom)
	if l > 0 {
		n += 1 + l + sovLiquidfarming(uint64(l))
	}
	l = len(m.PayingReserveAddress)
	if l > 0 {
		n += 1 + l + sovLiquidfarming(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.StartTime)
	n += 1 + l + sovLiquidfarming(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.EndTime)
	n += 1 + l + sovLiquidfarming(uint64(l))
	if m.Status != 0 {
		n += 1 + sovLiquidfarming(uint64(m.Status))
	}
	l = len(m.Winner)
	if l > 0 {
		n += 1 + l + sovLiquidfarming(uint64(l))
	}
	l = m.WinningAmount.Size()
	n += 1 + l + sovLiquidfarming(uint64(l))
	if len(m.Rewards) > 0 {
		for _, e := range m.Rewards {
			l = e.Size()
			n += 1 + l + sovLiquidfarming(uint64(l))
		}
	}
	if len(m.Fees) > 0 {
		for _, e := range m.Fees {
			l = e.Size()
			n += 1 + l + sovLiquidfarming(uint64(l))
		}
	}
	l = m.FeeRate.Size()
	n += 1 + l + sovLiquidfarming(uint64(l))
	return n
}

func (m *CompoundingRewards) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Amount.Size()
	n += 1 + l + sovLiquidfarming(uint64(l))
	return n
}

func (m *Bid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PoolId != 0 {
		n += 1 + sovLiquidfarming(uint64(m.PoolId))
	}
	l = len(m.Bidder)
	if l > 0 {
		n += 1 + l + sovLiquidfarming(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovLiquidfarming(uint64(l))
	return n
}

func sovLiquidfarming(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLiquidfarming(x uint64) (n int) {
	return sovLiquidfarming(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RewardsAuction) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidfarming
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
			return fmt.Errorf("proto: RewardsAuction: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RewardsAuction: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BiddingCoinDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
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
				return ErrInvalidLengthLiquidfarming
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BiddingCoinDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PayingReserveAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
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
				return ErrInvalidLengthLiquidfarming
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PayingReserveAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
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
				return ErrInvalidLengthLiquidfarming
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.StartTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
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
				return ErrInvalidLengthLiquidfarming
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.EndTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= AuctionStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Winner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
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
				return ErrInvalidLengthLiquidfarming
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Winner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field WinningAmount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
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
				return ErrInvalidLengthLiquidfarming
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.WinningAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rewards", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
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
				return ErrInvalidLengthLiquidfarming
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rewards = append(m.Rewards, types.Coin{})
			if err := m.Rewards[len(m.Rewards)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fees", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
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
				return ErrInvalidLengthLiquidfarming
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fees = append(m.Fees, types.Coin{})
			if err := m.Fees[len(m.Fees)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeRate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
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
				return ErrInvalidLengthLiquidfarming
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FeeRate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidfarming(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidfarming
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
func (m *CompoundingRewards) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidfarming
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
			return fmt.Errorf("proto: CompoundingRewards: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CompoundingRewards: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
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
				return ErrInvalidLengthLiquidfarming
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidfarming(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidfarming
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
func (m *Bid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLiquidfarming
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
			return fmt.Errorf("proto: Bid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Bid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolId", wireType)
			}
			m.PoolId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PoolId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bidder", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
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
				return ErrInvalidLengthLiquidfarming
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bidder = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLiquidfarming
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
				return ErrInvalidLengthLiquidfarming
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLiquidfarming
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLiquidfarming(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLiquidfarming
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
func skipLiquidfarming(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLiquidfarming
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
					return 0, ErrIntOverflowLiquidfarming
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
					return 0, ErrIntOverflowLiquidfarming
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
				return 0, ErrInvalidLengthLiquidfarming
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLiquidfarming
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLiquidfarming
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLiquidfarming        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLiquidfarming          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLiquidfarming = fmt.Errorf("proto: unexpected end of group")
)
