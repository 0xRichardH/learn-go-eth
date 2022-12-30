package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAddressValid(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc    string
		address string
		result  bool
	}{
		{
			desc:    "Is valid address",
			address: "0x323b5d4c32345ced77393b3530b1eed0f346429d",
			result:  true,
		},
		{
			desc:    "Is invalid address",
			address: "0xZYXb5d4c32345ced77393b3530b1eed0f346429d",
			result:  false,
		},
	}

	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()
			result := IsAddressValid(tC.address)
			assert.Equal(t, tC.result, result)
		})
	}
}

func TestIsContract(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		desc    string
		address string
		result  bool
	}{
		{
			desc:    "It is a smart contract address",
			address: "0xe41d2489571d322189246dafa5ebde1f4699f498",
			result:  true,
		},
		{
			desc:    "It is an account address",
			address: "0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4",
			result:  false,
		},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			result := IsContract(tC.address)
			assert.Equal(t, tC.result, result)
		})
	}
}
