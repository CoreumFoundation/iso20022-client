package coreum

import (
	"github.com/cosmos/cosmos-sdk/client"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

// GetPubKey retrieves public key of an account by its address.
func GetPubKey(clientCtx client.Context, address sdk.AccAddress) (cryptotypes.PubKey, error) {
	ar := types.AccountRetriever{}
	acc, err := ar.GetAccount(clientCtx, address)
	if err != nil {
		return nil, err
	}
	return acc.GetPubKey(), nil
}
