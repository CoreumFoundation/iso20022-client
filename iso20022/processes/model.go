package processes

import (
	"context"

	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	types3 "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/CoreumFoundation/coreum/v4/pkg/client"
	"github.com/CoreumFoundation/iso20022-client/iso20022/addressbook"
	"github.com/CoreumFoundation/iso20022-client/iso20022/coreum"
)

//go:generate mockgen -destination=model_mocks_test.go -package=processes_test . ContractClient,AddressBook,Cryptography

type ContractClient interface {
	SetContractAddress(contractAddress types.AccAddress) error
	GetContractAddress() types.AccAddress
	IsInitialized() bool
	SendMessage(
		ctx context.Context, sender, destination types.AccAddress, message coreum.NFTInfo,
	) (*types.TxResponse, error)
	MarkAsRead(
		ctx context.Context, sender types.AccAddress, until uint64,
	) (*types.TxResponse, error)
	IssueNFTClass(
		ctx context.Context,
		sender types.AccAddress,
		symbol, name, description string,
	) (*types.TxResponse, error)
	MintNFT(
		ctx context.Context,
		sender types.AccAddress,
		classId, id string,
		data *types2.Any,
	) (*types.TxResponse, error)
	GetUnreadMessages(
		ctx context.Context,
		address types.AccAddress,
		limit *uint32,
	) ([]coreum.Message, error)
	GetReadMessages(
		ctx context.Context,
		address types.AccAddress,
		startAfterKey string,
		limit *uint32,
	) ([]coreum.Message, error)
	QueryNFT(
		ctx context.Context,
		classId, id string,
	) (*types2.Any, error)
}

type AddressBook interface {
	Update(ctx context.Context) error
	KeyAlgo() string
	ForEach(f func(address addressbook.Address) bool)
	Lookup(expectedAddress addressbook.BranchAndIdentification) (*addressbook.Address, bool)
	LookupByAccountAddress(bech32EncodedAddress string) (*addressbook.Address, bool)
}

type Cryptography interface {
	GenerateSharedKeyByPrivateKeyName(ctx client.Context, algo string, privateKeyName string, publicKeyBytes []byte) ([]byte, error)
	GenerateSharedKey(algo string, privateKey types3.PrivKey, publicKeyBytes []byte) ([]byte, error)
}
