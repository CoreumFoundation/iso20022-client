package crypto

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/crypto/xsalsa20symmetric"
	"github.com/stretchr/testify/require"
)

func TestGenerateSecp256k1SharedKey(t *testing.T) {
	t.Parallel()

	privateKey := secp256k1.GenPrivKey()

	sharedKey, err := generateSecp256k1SharedKey(privateKey, privateKey.PubKey().Bytes())
	require.NoError(t, err)
	require.Len(t, sharedKey, 32)
}

func TestEncrypt(t *testing.T) {
	t.Parallel()

	privateKey := secp256k1.GenPrivKey()

	sharedKey, err := generateSecp256k1SharedKey(privateKey, privateKey.PubKey().Bytes())
	require.NoError(t, err)

	txt := []byte("hello world")
	cypherText := xsalsa20symmetric.EncryptSymmetric(txt, sharedKey)

	plainText, err := xsalsa20symmetric.DecryptSymmetric(cypherText, sharedKey)
	require.NoError(t, err)
	require.Equal(t, txt, plainText)
}

func TestGenerateSharedKey(t *testing.T) {
	c := Cryptography{}

	privateKey := secp256k1.GenPrivKey()
	pubKeyBytes := privateKey.PubKey().Bytes()

	testData := []struct {
		name        string
		privateKey  cryptotypes.PrivKey
		pubKeyBytes []byte
		ok          bool
	}{
		{
			name:        "correct secp256k1 pair",
			privateKey:  privateKey,
			pubKeyBytes: pubKeyBytes,
			ok:          true,
		},
		{
			name:        "incorrect secp256k1 pair",
			privateKey:  privateKey,
			pubKeyBytes: append(pubKeyBytes, []byte("invalid")...),
			ok:          false,
		},
	}

	for _, tt := range testData {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.GenerateSharedKey(tt.privateKey, tt.pubKeyBytes)
			if tt.ok {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}

func TestGenerateSameSharedKeyForBothParties(t *testing.T) {
	c := Cryptography{}

	privateKey1 := secp256k1.GenPrivKey()
	pubKey1Bytes := privateKey1.PubKey().Bytes()

	privateKey2 := secp256k1.GenPrivKey()
	pubKey2Bytes := privateKey2.PubKey().Bytes()

	sharedKey1, err := c.GenerateSharedKey(privateKey1, pubKey2Bytes)
	require.NoError(t, err)

	sharedKey2, err := c.GenerateSharedKey(privateKey2, pubKey1Bytes)
	require.NoError(t, err)

	require.Equal(t, sharedKey1, sharedKey2)
}
