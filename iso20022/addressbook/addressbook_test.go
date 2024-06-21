package addressbook

import (
	"context"
	"encoding/base64"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/stretchr/testify/require"

	"github.com/CoreumFoundation/iso20022-client/iso20022/crypto"
)

func TestAddressBook(t *testing.T) {
	ctx := context.Background()

	ab := New("coreum-testnet-1")

	require.NoError(t, ab.Update(ctx))

	addr, ok := ab.Lookup(BranchAndIdentification{
		Identification: Identification{
			Bic: "6P9YGUDF",
		},
	})
	require.True(t, ok)

	t.Logf("Address: %s", addr.Bech32EncodedAddress)
	t.Logf("PublicKey: %s", addr.PublicKey)

	keyBytes, err := base64.StdEncoding.DecodeString(addr.PublicKey)
	require.NoError(t, err)

	privateKey := secp256k1.GenPrivKey()

	sharedKey, err := crypto.GenerateSharedKey(ab.KeyAlgo(), privateKey, keyBytes)
	require.NoError(t, err)

	t.Logf("SharedKey: %x", sharedKey)
}

func TestPostalAddress(t *testing.T) {
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
}
