package processes

import (
	"context"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/CoreumFoundation/coreum/v4/pkg/client"
	"github.com/CoreumFoundation/iso20022-client/iso20022/addressbook"
	"github.com/CoreumFoundation/iso20022-client/iso20022/coreum"
	"github.com/CoreumFoundation/iso20022-client/iso20022/queue"
)

//go:generate mockgen -destination=model_mocks_test.go -package=processes_test . ContractClient,AddressBook,Cryptography,Parser,MessageQueue

type ContractClient interface {
	SetContractAddress(contractAddress types.AccAddress) error
	GetContractAddress() types.AccAddress
	IsInitialized() bool
	BroadcastMessages(
		ctx context.Context,
		sender types.AccAddress,
		messages ...types.Msg,
	) (*types.TxResponse, error)
	SendMessage(
		ctx context.Context, sender, destination types.AccAddress, message coreum.NFTInfo,
	) (*types.TxResponse, error)
	SendMessages(
		ctx context.Context, sender types.AccAddress, messages ...coreum.MessageWithDestination,
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
		data *codectypes.Any,
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
	) (*codectypes.Any, error)
}

type AddressBook interface {
	Update(ctx context.Context) error
	Validate() error
	Lookup(expectedAddress addressbook.Party) (*addressbook.Address, bool)
	LookupByAccountAddress(bech32EncodedAddress string) (*addressbook.Address, bool)
}

type Cryptography interface {
	GenerateSharedKeyByPrivateKeyName(ctx client.Context, privateKeyName string, publicKeyBytes []byte) ([]byte, error)
	GenerateSharedKey(privateKey cryptotypes.PrivKey, publicKeyBytes []byte) ([]byte, error)
	EncryptSymmetric(plaintext []byte, secret []byte) (ciphertext []byte)
	DecryptSymmetric(ciphertext []byte, secret []byte) (plaintext []byte, err error)
}

type Parser interface {
	ExtractMetadataFromIsoMessage(msg []byte) (id string, party *addressbook.Party, err error)
}

type MessageQueue interface {
	Start(ctx context.Context) error
	PushToSend(id string, msg []byte) queue.Status
	PushToReceive(msg []byte)
	PopFromSend(ctx context.Context, n int, dur time.Duration) [][]byte
	PopFromReceive() ([]byte, bool)
	SetStatus(id string, status queue.Status)
	Close()
}
