package types_test

import (
	"bytes"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/suite"
	"github.com/tendermint/tendermint/crypto"

	"github.com/redactedfury/petri/x/liquidity/types"
)

type keysTestSuite struct {
	suite.Suite
}

func TestKeysTestSuite(t *testing.T) {
	suite.Run(t, new(keysTestSuite))
}

func (s *keysTestSuite) TestGetPairKey() {
	s.Require().Equal([]byte{0xa2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}, types.GetPairKey(1, 0))
	s.Require().Equal([]byte{0xa2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x9}, types.GetPairKey(2, 9))
	s.Require().Equal([]byte{0xa2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xb, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa}, types.GetPairKey(11, 10))
}

func (s *keysTestSuite) TestGetPairIndexKey() {
	s.Require().Equal([]byte{0xa3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x6, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x31, 0x6, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x32}, types.GetPairIndexKey(1, "denom1", "denom2"))
	s.Require().Equal([]byte{0xa3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x6, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x33, 0x6, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x34}, types.GetPairIndexKey(2, "denom3", "denom4"))
}

func (s *keysTestSuite) TestPairsByDenomsIndexKey() {
	testCases := []struct {
		denomA   string
		denomB   string
		pairId   uint64
		expected []byte
	}{
		{
			"denomA",
			"denomB",
			1,
			[]byte{0xa4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x6, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x41, 0x6, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x42, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1},
		},
		{
			"denomC",
			"denomD",
			20,
			[]byte{0xa4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x6, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x43, 0x6, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x44, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x14},
		},
		{
			"denomE",
			"denomF",
			13,
			[]byte{0xa4, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x6, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x45, 0x6, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x46, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xd},
		},
	}

	for _, tc := range testCases {
		key := types.GetPairsByDenomsIndexKey(1, tc.denomA, tc.denomB, tc.pairId)
		s.Require().Equal(tc.expected, key)

		s.Require().True(bytes.HasPrefix(key, types.GetPairsByDenomsIndexKeyPrefix(1, tc.denomA, tc.denomB)))

		denomA, denomB, pairId := types.ParsePairsByDenomsIndexKey(key)
		s.Require().Equal(tc.denomA, denomA)
		s.Require().Equal(tc.denomB, denomB)
		s.Require().Equal(tc.pairId, pairId)
	}
}

func (s *keysTestSuite) TestGetPoolKey() {
	s.Require().Equal([]byte{0xa5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1}, types.GetPoolKey(1, 1))
	s.Require().Equal([]byte{0xa5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5}, types.GetPoolKey(1, 5))
	s.Require().Equal([]byte{0xa5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa}, types.GetPoolKey(1, 10))
}

func (s *keysTestSuite) TestGetPoolByReserveAddressIndexKey() {
	reserveAddr1 := types.PoolReserveAddress(1, 1)
	reserveAddr2 := types.PoolReserveAddress(2, 2)
	reserveAddr3 := types.PoolReserveAddress(3, 3)
	s.Require().Equal([]byte{0xa6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x20, 0x81, 0xe2, 0x85, 0x30, 0xf1, 0xda, 0xf7, 0x8a, 0x92, 0xae, 0xc2, 0x63, 0xb7, 0x1d, 0x4f, 0xaa, 0x7e, 0x35, 0xbd, 0xcd, 0xf8, 0x55, 0xbd, 0xca, 0x64, 0xb4, 0xd5, 0x88, 0x1b, 0x3, 0x84, 0x28}, types.GetPoolByReserveAddressIndexKey(1, reserveAddr1))
	s.Require().Equal([]byte{0xa6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x20, 0x4f, 0xd4, 0xd4, 0x26, 0x1a, 0x57, 0x18, 0x58, 0xe, 0xc7, 0x58, 0x7c, 0xea, 0x9b, 0xa2, 0xf9, 0x8c, 0x27, 0x18, 0xa9, 0x10, 0x2e, 0xa, 0x61, 0x5a, 0x7b, 0x28, 0x9e, 0x42, 0x47, 0x38, 0xd4}, types.GetPoolByReserveAddressIndexKey(2, reserveAddr2))
	s.Require().Equal([]byte{0xa6, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0x20, 0x74, 0xcf, 0x4b, 0xaa, 0xfa, 0x24, 0x6c, 0x4a, 0x9, 0x4f, 0x1d, 0x97, 0x88, 0x2e, 0x91, 0x35, 0x88, 0xc, 0x8a, 0x17, 0xaf, 0xa6, 0x7, 0x43, 0x62, 0x5a, 0x5c, 0x6e, 0xd3, 0xd2, 0xdc, 0x3d}, types.GetPoolByReserveAddressIndexKey(3, reserveAddr3))
}

func (s *keysTestSuite) TestPoolsByPairIndexKey() {
	testCases := []struct {
		pairId   uint64
		poolId   uint64
		expected []byte
	}{
		{
			5,
			10,
			[]byte{0xa7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xa},
		},
		{
			2,
			7,
			[]byte{0xa7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7},
		},
		{
			3,
			5,
			[]byte{0xa7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x5},
		},
	}

	for _, tc := range testCases {
		key := types.GetPoolsByPairIndexKey(1, tc.pairId, tc.poolId)
		s.Require().Equal(tc.expected, key)

		s.Require().True(bytes.HasPrefix(key, types.GetPoolsByPairIndexKeyPrefix(1, tc.pairId)))

		poolId := types.ParsePoolsByPairIndexKey(key)
		s.Require().Equal(tc.poolId, poolId)
	}
}

func (s *keysTestSuite) TestGetDepositRequestKey() {
	s.Require().Equal([]byte{0xa8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1}, types.GetDepositRequestKey(1, 1, 1))
	s.Require().Equal([]byte{0xa8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xe8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xe9}, types.GetDepositRequestKey(1, 1000, 1001))
}

func (s *keysTestSuite) TestDepositRequestIndexKey() {
	depositor := sdk.AccAddress(crypto.AddressHash([]byte("depositor")))
	key := types.GetDepositRequestIndexKey(1, depositor, 1, 2)
	s.Require().Equal([]byte{0xa9, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x14, 0x9a, 0x69, 0x97, 0x1f, 0x1d, 0xb2, 0xe1, 0xd8, 0x77, 0x73, 0x6f, 0x7d, 0x36, 0x96, 0x90, 0xa3, 0xbf, 0x57, 0xcf, 0x22, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2}, key)
	s.Require().True(bytes.HasPrefix(key, types.GetDepositRequestIndexKeyPrefix(1, depositor)))
	depositor2, poolId, reqId := types.ParseDepositRequestIndexKey(key)
	s.Require().Equal(depositor, depositor2)
	s.Require().Equal(uint64(1), poolId)
	s.Require().Equal(uint64(2), reqId)
}

func (s *keysTestSuite) TestGetWithdrawRequestKey() {
	s.Require().Equal([]byte{0xb0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1}, types.GetWithdrawRequestKey(1, 1, 1))
	s.Require().Equal([]byte{0xb0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xe8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xe9}, types.GetWithdrawRequestKey(2, 1000, 1001))
}

func (s *keysTestSuite) TestWithdrawRequestIndexKey() {
	withdrawer := sdk.AccAddress(crypto.AddressHash([]byte("withdrawer")))
	key := types.GetWithdrawRequestIndexKey(1, withdrawer, 1, 2)
	s.Require().Equal([]byte{0xb1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x14, 0x19, 0xcd, 0x70, 0x1f, 0x44, 0xf1, 0xed, 0xe, 0x3, 0xa7, 0xf3, 0xf8, 0x7c, 0xff, 0x84, 0x79, 0x58, 0xc6, 0x56, 0xc2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x2}, key)
	s.Require().True(bytes.HasPrefix(key, types.GetWithdrawRequestIndexKeyPrefix(1, withdrawer)))
	withdrawer2, poolId, reqId := types.ParseWithdrawRequestIndexKey(key)
	s.Require().Equal(withdrawer, withdrawer2)
	s.Require().Equal(uint64(1), poolId)
	s.Require().Equal(uint64(2), reqId)
}

func (s *keysTestSuite) TestGetOrderKey() {
	s.Require().Equal([]byte{0xb2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1}, types.GetOrderKey(1, 1, 1))
	s.Require().Equal([]byte{0xb2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xe8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xe9}, types.GetOrderKey(1, 1000, 1001))
}

func (s *keysTestSuite) TestGetOrdersByPairKeyPrefix() {
	s.Require().Equal([]byte{0xb2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1}, types.GetOrdersByPairKeyPrefix(1, 1))
	s.Require().Equal([]byte{0xb2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x3, 0xe8}, types.GetOrdersByPairKeyPrefix(1, 1000))
}

func (s *keysTestSuite) TestOrderIndexKey() {
	orderer := sdk.AccAddress(crypto.AddressHash([]byte("orderer")))
	key := types.GetOrderIndexKey(1, orderer, 1, 1)
	s.Require().Equal([]byte{0xb3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x14, 0x54, 0x7e, 0xfe, 0x47, 0x8f, 0xc9, 0xf9, 0x52, 0xb2, 0x5c, 0xbc, 0x50, 0xf2, 0x85, 0xf7, 0x7d, 0xff, 0x52, 0x9f, 0x25, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1}, key)
	s.Require().True(bytes.HasPrefix(key, types.GetOrderIndexKeyPrefix(1, orderer)))
	orderer2, pairId, orderId := types.ParseOrderIndexKey(key)
	s.Require().Equal(orderer, orderer2)
	s.Require().Equal(uint64(1), pairId)
	s.Require().Equal(uint64(1), orderId)
}
