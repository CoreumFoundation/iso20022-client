package crypto

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256r1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	secp256k1v4 "github.com/decred/dcrd/dcrec/secp256k1/v4"
	"github.com/pkg/errors"

	coreumchainclient "github.com/CoreumFoundation/coreum/v4/pkg/client"
)

func GenerateSharedKeyByPrivateKeyName(ctx coreumchainclient.Context, algo string, privateKeyName string, publicKeyBytes []byte) ([]byte, error) {
	// getting private key
	key, err := ctx.Keyring().Key(privateKeyName)
	if err != nil {
		return nil, err
	}

	recordLocal := key.GetLocal()
	if recordLocal == nil {
		return nil, errors.New("private key can only be local") // TODO
	}

	privKey, err := extractPrivateKeyFromLocal(recordLocal)
	if err != nil {
		return nil, err
	}

	return GenerateSharedKey(algo, privKey, publicKeyBytes)
}

func GenerateSharedKey(algo string, privateKey cryptotypes.PrivKey, publicKeyBytes []byte) ([]byte, error) {
	var err error
	var sharedKey []byte

	switch privateKey.(type) {
	case *secp256r1.PrivKey:
		switch algo {
		case "secp256r1":
			sharedKey, err = generateSecp256r1SharedKey(privateKey.(*secp256r1.PrivKey), publicKeyBytes)
			if err != nil {
				return nil, err
			}
		default:
			return nil, errors.New("keys are not from the same curve") // TODO
		}
	case *secp256k1.PrivKey:
		switch algo {
		case "secp256k1":
			sharedKey, err = generateSecp256k1SharedKey(privateKey.(*secp256k1.PrivKey), publicKeyBytes)
			if err != nil {
				return nil, err
			}
		default:
			return nil, errors.New("keys are not from the same curve") // TODO
		}
	default:
		return nil, errors.New("unsupported key type") // TODO
	}

	// TODO: It is recommended to securely hash the result before using as a cryptographic key.
	h := sha256.Sum256(sharedKey)
	return h[:], nil
}

//func GenerateSharedKeyByAccount(ctx coreumchainclient.Context, localAccount, remoteAccount sdk.AccAddress) ([]byte, error) {
//	// getting private key
//	key, err := ctx.Keyring().KeyByAddress(localAccount)
//	if err != nil {
//		return nil, err
//	}
//
//	recordLocal := key.GetLocal()
//	if recordLocal == nil {
//		return nil, errors.New("private key can only be local")
//	}
//
//	privKey, err := extractPrivateKeyFromLocal(recordLocal)
//	if err != nil {
//		return nil, err
//	}
//
//	// getting public key
//	pubKey, err := coreum.GetPubKey(ctx.SDKContext(), remoteAccount)
//	if err != nil {
//		return nil, err
//	}
//
//	var sharedKey []byte
//
//	switch privKey.(type) {
//	case *secp256r1.PrivKey:
//		switch pubKey.(type) {
//		case *secp256r1.PubKey:
//			sharedKey, err = generateSecp256r1SharedKey(privKey.(*secp256r1.PrivKey), pubKey.(*secp256r1.PubKey))
//			if err != nil {
//				return nil, err
//			}
//		default:
//			return nil, errors.New("keys are not from the same curve")
//		}
//	case *secp256k1.PrivKey:
//		switch pubKey.(type) {
//		case *secp256k1.PubKey:
//			sharedKey, err = generateSecp256k1SharedKey(privKey.(*secp256k1.PrivKey), pubKey.(*secp256k1.PubKey))
//			if err != nil {
//				return nil, err
//			}
//		default:
//			return nil, errors.New("keys are not from the same curve")
//		}
//	default:
//		return nil, errors.New("unsupported key type")
//	}
//
//	// TODO: It is recommended to securely hash the result before using as a cryptographic key.
//	h := sha256.Sum256(sharedKey)
//	return h[:], nil
//}

func extractPrivateKeyFromLocal(rl *keyring.Record_Local) (cryptotypes.PrivKey, error) {
	if rl.PrivKey == nil {
		return nil, errors.New("private key is not available")
	}

	privateKey, ok := rl.PrivKey.GetCachedValue().(cryptotypes.PrivKey)
	if !ok {
		return nil, errors.New("unable to cast any to cryptotypes.PrivKey") // TODO
	}

	return privateKey, nil
}

func generateSecp256k1SharedKey(privateKey *secp256k1.PrivKey, publicKeyBytes []byte) ([]byte, error) {
	// preparing private key
	pvKey := secp256k1v4.PrivKeyFromBytes(privateKey.Bytes())

	// preparing public key
	pbKey, err := secp256k1v4.ParsePubKey(publicKeyBytes)
	if err != nil {
		return nil, err
	}

	// generating shared key
	return secp256k1v4.GenerateSharedSecret(pvKey, pbKey), nil
}

func generateSecp256r1SharedKey(privateKey *secp256r1.PrivKey, publicKeyBytes []byte) ([]byte, error) {
	// preparing private key
	ecdhPriv, err := privateKey.Secret.PrivKey.PrivateKey.ECDH()
	if err != nil {
		return nil, err
	}

	// preparing public key
	pbKey, err := x509.ParsePKIXPublicKey(publicKeyBytes)
	publicKey, ok := pbKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("unable to cast any to ecdsa.PublicKey") // TODO
	}
	ecdhPub, err := publicKey.ECDH()
	if err != nil {
		return nil, err
	}

	// generating shared key
	return ecdhPriv.ECDH(ecdhPub)
}
