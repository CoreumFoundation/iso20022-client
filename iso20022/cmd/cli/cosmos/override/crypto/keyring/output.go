package keyring

import (
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AddressFormatter defines the address formatter.
type AddressFormatter func(publicKey cryptotypes.PubKey) sdk.Address

// SelectedAddressFormatter stores the address formatter used by the key commands.
var SelectedAddressFormatter AddressFormatter = CoreumAddressFormatter

// CoreumAddressFormatter formats coreum addresses.
func CoreumAddressFormatter(publicKey cryptotypes.PubKey) sdk.Address {
	return sdk.AccAddress(publicKey.Address())
}

// MkAccKeyOutput create a KeyOutput in with "acc" Bech32 prefixes. If the
// public key is a multisig public key, then the threshold and constituent
// public keys will be added.
func MkAccKeyOutput(k *keyring.Record) (keyring.KeyOutput, error) {
	pk, err := k.GetPubKey()
	if err != nil {
		return keyring.KeyOutput{}, err
	}
	return keyring.NewKeyOutput(k.Name, k.GetType(), SelectedAddressFormatter(pk), pk)
}

// MkAccKeysOutput returns a slice of KeyOutput objects, each with the "acc"
// Bech32 prefixes, given a slice of Record objects. It returns an error if any
// call to MkKeyOutput fails.
func MkAccKeysOutput(records []*keyring.Record) ([]keyring.KeyOutput, error) {
	kos := make([]keyring.KeyOutput, len(records))
	var err error
	for i, r := range records {
		kos[i], err = MkAccKeyOutput(r)
		if err != nil {
			return nil, err
		}
	}

	return kos, nil
}
