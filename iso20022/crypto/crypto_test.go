package crypto

import (
	"crypto/x509"
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256r1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/crypto/xsalsa20symmetric"
	"github.com/stretchr/testify/require"
)

func TestGenerateSecp256r1SharedKey(t *testing.T) {
	t.Parallel()

	privateKey, err := secp256r1.GenPrivKey()
	require.NoError(t, err)

	publicKey := &privateKey.PubKey().(*secp256r1.PubKey).Key.PublicKey

	pubKeyBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	require.NoError(t, err)

	sharedKey, err := generateSecp256r1SharedKey(privateKey, pubKeyBytes)
	require.NoError(t, err)
	require.Len(t, sharedKey, 32)
}

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

	privateKeyK1 := secp256k1.GenPrivKey()
	pubKeyK1Bytes := privateKeyK1.PubKey().Bytes()

	privateKeyR1, err := secp256r1.GenPrivKey()
	require.NoError(t, err)

	publicKey := &privateKeyR1.PubKey().(*secp256r1.PubKey).Key.PublicKey

	pubKeyR1Bytes, err := x509.MarshalPKIXPublicKey(publicKey)
	require.NoError(t, err)

	testData := []struct {
		name        string
		algo        string
		privateKey  cryptotypes.PrivKey
		pubKeyBytes []byte
		ok          bool
	}{
		{
			name:        "correct secp256k1 pair",
			algo:        "secp256k1",
			privateKey:  privateKeyK1,
			pubKeyBytes: pubKeyK1Bytes,
			ok:          true,
		},
		{
			name:        "correct secp256r1 pair",
			algo:        "secp256r1",
			privateKey:  privateKeyR1,
			pubKeyBytes: pubKeyR1Bytes,
			ok:          true,
		},
		{
			name:        "incorrect secp256k1 pair",
			algo:        "secp256k1",
			privateKey:  privateKeyR1,
			pubKeyBytes: pubKeyK1Bytes,
			ok:          false,
		},
		{
			name:        "correct secp256r1 pair",
			algo:        "secp256r1",
			privateKey:  privateKeyK1,
			pubKeyBytes: pubKeyR1Bytes,
			ok:          false,
		},
	}

	for _, tt := range testData {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			_, err = c.GenerateSharedKey(tt.algo, tt.privateKey, tt.pubKeyBytes)
			if tt.ok {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
