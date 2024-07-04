package addressbook

import (
	"context"
	"encoding/base64"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/stretchr/testify/require"

	"github.com/CoreumFoundation/iso20022-client/iso20022/crypto"
)

// TODO: Mock file/web reads

func TestEmptyAddressBook(t *testing.T) {
	ab := NewWithRepoAddress("file://./testdata/coreum-devnet-1/addressbook.json")

	addr, ok := ab.Lookup(BranchAndIdentification{
		Identification: Identification{
			Bic: "6P9YGUDF",
		},
	})
	require.False(t, ok)
	require.Nil(t, addr)
}

func TestLookup(t *testing.T) {
	ctx := context.Background()
	c := &crypto.Cryptography{}

	ab := NewWithRepoAddress("file://./testdata/coreum-devnet-1/addressbook.json")

	require.NoError(t, ab.Update(ctx))

	require.NoError(t, ab.Validate())

	addr, ok := ab.Lookup(BranchAndIdentification{
		Identification: Identification{
			Bic: "6P9YGUDF",
		},
	})
	require.True(t, ok)

	keyBytes, err := base64.StdEncoding.DecodeString(addr.PublicKey)
	require.NoError(t, err)

	privateKey := secp256k1.GenPrivKey()

	_, err = c.GenerateSharedKey(privateKey, keyBytes)
	require.NoError(t, err)
}

func TestLookupByAccountAddress(t *testing.T) {
	ctx := context.Background()

	ab := NewWithRepoAddress("file://./testdata/coreum-devnet-1/addressbook.json")

	require.NoError(t, ab.Update(ctx))

	require.NoError(t, ab.Validate())

	addr, ok := ab.LookupByAccountAddress("devcore1kdujjkp8u0j9lww3n7qs7r5fwkljelvecsq43d")

	require.True(t, ok)
	require.Equal(t, "6P9YGUDF", addr.BranchAndIdentification.Identification.Bic)
}

func TestUpdate(t *testing.T) {
	ctx := context.Background()

	testData := []struct {
		name string
		ab   *AddressBook
		err  bool
	}{
		{
			name: "wrong path",
			ab:   NewWithRepoAddress("file://./testdata/wrong-chain-id/addressbook.json"),
			err:  true,
		},
		{
			name: "wrong url",
			ab:   New("wrong-chain-id"),
			err:  true,
		},
	}

	for _, tt := range testData {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			res := tt.ab.Update(ctx)
			if tt.err {
				require.Error(t, res)
			} else {
				require.NoError(t, res)
			}
			err := tt.ab.Validate()
			require.NoError(t, err)
		})
	}
}

func TestCache(t *testing.T) {
	ctx := context.Background()

	testData := []struct {
		name string
		ab   *AddressBook
	}{
		{
			name: "actual repo",
			ab:   New("coreum-devnet-1"),
		},
		{
			name: "local file",
			ab:   NewWithRepoAddress("file://./testdata/coreum-devnet-1/addressbook.json"),
		},
	}

	for _, tt := range testData {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			require.NoError(t, tt.ab.Update(ctx))
			require.NoError(t, tt.ab.Update(ctx))
			require.NoError(t, tt.ab.Validate())
		})
	}
}
