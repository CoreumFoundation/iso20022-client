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

//go:generate mockgen -destination=model_mocks_test.go -package=processes_test . ContractClient,AddressBook,Cryptography,Parser,MessageQueue,Dtif

type ContractClient interface {
	DeployAndInstantiate(
		ctx context.Context,
		sender types.AccAddress,
		contractByteCodePath string,
	) (types.AccAddress, error)
	DeployContract(
		ctx context.Context,
		sender types.AccAddress,
		contractByteCodePath string,
	) (*types.TxResponse, uint64, error)
	MigrateContract(
		ctx context.Context,
		sender types.AccAddress,
		codeID uint64,
	) (*types.TxResponse, error)
	SetContractAddress(contractAddress types.AccAddress) error
	GetContractAddress() types.AccAddress
	IsInitialized() bool
	BroadcastMessages(
		ctx context.Context,
		sender types.AccAddress,
		messages ...types.Msg,
	) (*types.TxResponse, error)
	StartSession(
		ctx context.Context, eutr string, sender types.AccAddress, message coreum.NFTInfo, destination types.AccAddress, funds types.Coins,
	) (*types.TxResponse, error)
	StartSessions(
		ctx context.Context, sender types.AccAddress, sessions ...coreum.StartSession,
	) (*types.TxResponse, error)
	SendMessage(
		ctx context.Context, sender, destination types.AccAddress, eutr, ID string, message coreum.NFTInfo,
	) (*types.TxResponse, error)
	SendMessages(
		ctx context.Context, sender types.AccAddress, messages ...coreum.SendMessage,
	) (*types.TxResponse, error)
	ConfirmSession(
		ctx context.Context, eutr string, sender, initiator, destination types.AccAddress,
	) (*types.TxResponse, error)
	ConfirmSessions(
		ctx context.Context, sender types.AccAddress, messages ...coreum.ConfirmSession,
	) (*types.TxResponse, error)
	CancelSession(
		ctx context.Context, eutr string, sender, initiator, destination types.AccAddress,
	) (*types.TxResponse, error)
	CancelSessions(
		ctx context.Context, sender types.AccAddress, messages ...coreum.CancelSession,
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
		classId, Id string,
		data *codectypes.Any,
	) (*types.TxResponse, error)
	GetActiveSessions(
		ctx context.Context,
		address types.AccAddress,
		userType coreum.UserType,
		startAfterKey *string,
		limit *uint32,
	) ([]coreum.Session, error)
	GetClosedSessions(
		ctx context.Context,
		address types.AccAddress,
		userType coreum.UserType,
		startAfterKey *string,
		limit *uint32,
	) ([]coreum.Session, error)
	GetNewMessages(
		ctx context.Context,
		address types.AccAddress,
		limit *uint32,
	) ([]coreum.Message, error)
	GetMessages(
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

type Metadata struct {
	Eutr     string
	ID       string
	Sender   *addressbook.Party
	Receiver *addressbook.Party
}

type Parser interface {
	ExtractMetadataFromIsoMessage(msg []byte) (data Metadata, err error)
}

type MessageQueue interface {
	Start(ctx context.Context) error
	PushToSend(id string, msg []byte) queue.Status
	PushToReceive(msg []byte)
	PopFromSend(ctx context.Context, n int, dur time.Duration) [][]byte
	PopFromReceive() ([]byte, bool)
	GetStatus(id string) *queue.Status
	SetStatus(id string, status queue.Status)
	Close()
}

type Dtif interface {
	Update(ctx context.Context) error
	LookupByDTI(dti string) (string, bool)
	LookupByDenom(denom string) (string, bool)
}
