package crypto

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256r1"
	"github.com/cosmos/cosmos-sdk/crypto/xsalsa20symmetric"
	"github.com/stretchr/testify/require"
)

func TestGenerateSecp256r1SharedKey(t *testing.T) {
	t.Parallel()

	privateKey, err := secp256r1.GenPrivKey()
	require.NoError(t, err)

	sharedKey, err := generateSecp256r1SharedKey(privateKey, privateKey.PubKey().(*secp256r1.PubKey))
	require.NoError(t, err)
	require.Len(t, sharedKey, 32)
}

func TestGenerateSecp256k1SharedKey(t *testing.T) {
	t.Parallel()

	privateKey := secp256k1.GenPrivKey()

	sharedKey, err := generateSecp256k1SharedKey(privateKey, privateKey.PubKey().(*secp256k1.PubKey))
	require.NoError(t, err)
	require.Len(t, sharedKey, 32)
}

func TestEncrypt(t *testing.T) {
	t.Parallel()

	privateKey := secp256k1.GenPrivKey()

	sharedKey, err := generateSecp256k1SharedKey(privateKey, privateKey.PubKey().(*secp256k1.PubKey))
	require.NoError(t, err)

	txt := []byte("hello world")
	cypherText := xsalsa20symmetric.EncryptSymmetric(txt, sharedKey)

	plainText, err := xsalsa20symmetric.DecryptSymmetric(cypherText, sharedKey)
	require.NoError(t, err)
	require.Equal(t, txt, plainText)
}
