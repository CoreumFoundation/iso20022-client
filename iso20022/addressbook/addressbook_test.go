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
	ab := NewWithRepoAddress("coreum-devnet-1", "file://./testdata/%s/addressbook.json")

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

	ab := NewWithRepoAddress("coreum-devnet-1", "file://./testdata/%s/addressbook.json")

	require.NoError(t, ab.Update(ctx))

	addr, ok := ab.Lookup(BranchAndIdentification{
		Identification: Identification{
			Bic: "6P9YGUDF",
		},
	})
	require.True(t, ok)

	keyBytes, err := base64.StdEncoding.DecodeString(addr.PublicKey)
	require.NoError(t, err)

	privateKey := secp256k1.GenPrivKey()

	_, err = c.GenerateSharedKey(ab.KeyAlgo(), privateKey, keyBytes)
	require.NoError(t, err)
}

func TestLookupByAccountAddress(t *testing.T) {
	ctx := context.Background()

	ab := NewWithRepoAddress("coreum-devnet-1", "file://./testdata/%s/addressbook.json")

	require.NoError(t, ab.Update(ctx))

	addr, ok := ab.LookupByAccountAddress("devcore1kdujjkp8u0j9lww3n7qs7r5fwkljelvecsq43d")

	require.True(t, ok)
	require.Equal(t, "6P9YGUDF", addr.BranchAndIdentification.Identification.Bic)
}

func TestForEach(t *testing.T) {
	ctx := context.Background()

	ab := NewWithRepoAddress("coreum-devnet-1", "file://./testdata/%s/addressbook.json")

	require.NoError(t, ab.Update(ctx))

	called := false

	ab.ForEach(func(address Address) bool {
		called = true
		require.NotEmpty(t, address.Bech32EncodedAddress)
		return false
	})

	require.True(t, called)
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
			ab:   NewWithRepoAddress("wrong-chain-id", "file://./testdata/%s/addressbook.json"),
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
			ab:   NewWithRepoAddress("coreum-devnet-1", "file://./testdata/%s/addressbook.json"),
		},
	}

	for _, tt := range testData {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			require.NoError(t, tt.ab.Update(ctx))
			require.NoError(t, tt.ab.Update(ctx))
		})
	}
}

func TestSerialization(t *testing.T) {
	p := PostalAddress{
		AddressType: &AddressType{
			Code: "l",
			Proprietary: &Proprietary{
				Id:         "m",
				Issuer:     "n",
				SchemeName: "o",
			},
		},
		CareOf:             "a",
		Department:         "b",
		SubDepartment:      "c",
		StreetName:         "d",
		BuildingNumber:     "e",
		BuildingName:       "",
		Floor:              "f",
		UnitNumber:         "",
		PostalBox:          "",
		Room:               "",
		PostalCode:         "g",
		TownName:           "",
		TownLocationName:   "",
		DistrictName:       "h",
		CountrySubDivision: "",
		CountryCode:        "i",
		AddressLine:        []string{"j", "k"},
	}
	require.Equal(t, "AddressType=(Code=l/Proprietary=(Id=m/Issuer=n/SchemeName=o))/BuildingNumber=e/CareOf=a/CountryCode=i/Department=b/DistrictName=h/Floor=f/PostalCode=g/StreetName=d/SubDepartment=c", p.String())

	c := ClearingSystemId{
		Code:        "a",
		Proprietary: "b",
	}
	require.Equal(t, "Code=a/Proprietary=b", c.String())

	o := Other{
		Issuer: "a",
		SchemeName: SchemeName{
			Code:        "b",
			Proprietary: "c",
		},
	}
	require.Equal(t, "Issuer=a/SchemeName=(Code=b/Proprietary=c)", o.String())

	b := Branch{
		Id: "a",
	}
	require.Equal(t, "Id=a", b.String())
}
