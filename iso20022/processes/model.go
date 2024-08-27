package processes

import (
	"context"
	"time"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types"

	"github.com/CoreumFoundation/coreum/v4/pkg/client"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/messages"
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
		ctx context.Context, uetr string, sender types.AccAddress, message coreum.NFTInfo, destination types.AccAddress, funds types.Coins,
	) (*types.TxResponse, error)
	StartSessions(
		ctx context.Context, sender types.AccAddress, sessions ...coreum.StartSession,
	) (*types.TxResponse, error)
	SendMessage(
		ctx context.Context, sender, destination types.AccAddress, uetr, ID string, message coreum.NFTInfo,
	) (*types.TxResponse, error)
	SendMessages(
		ctx context.Context, sender types.AccAddress, messages ...coreum.SendMessage,
	) (*types.TxResponse, error)
	ConfirmSession(
		ctx context.Context, uetr string, sender, initiator, destination types.AccAddress,
	) (*types.TxResponse, error)
	ConfirmSessions(
		ctx context.Context, sender types.AccAddress, messages ...coreum.ConfirmSession,
	) (*types.TxResponse, error)
	CancelSession(
		ctx context.Context, uetr string, sender, initiator, destination types.AccAddress,
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
	Uetr     string
	ID       string
	Sender   *addressbook.Party
	Receiver *addressbook.Party
}

type TransactionStatus string

const (
	// TransactionStatusNone None
	TransactionStatusNone TransactionStatus = "NONE"
	// TransactionStatusCreditorAcceptedSettlementCompleted This status is only applicable for instant payments.
	// Settlement on the creditor's account has been completed.
	TransactionStatusCreditorAcceptedSettlementCompleted TransactionStatus = "ACCC"
	// TransactionStatusAcceptedCustomerProfile Preceding check of technical validation was successful.
	// Customer profile check was also successful.
	TransactionStatusAcceptedCustomerProfile TransactionStatus = "ACCP"
	// TransactionStatusAcceptedSettlementCompleted Settlement on the debtor's account has been completed.
	// The status ACSC is also applicable when:
	//
	// When a recurring payment has passed the end date.
	// When a recurring payment has been withdrawn.
	TransactionStatusAcceptedSettlementCompleted TransactionStatus = "ACSC"
	// TransactionStatusAcceptedSettlementInProcess All preceding checks such as technical validation and customer profile were successful.
	// The payment initiation was successfully signed.
	// The payment initiation has been accepted for execution, but before settlement on the debtor’s account.
	TransactionStatusAcceptedSettlementInProcess TransactionStatus = "ACSP"
	// TransactionStatusAcceptedTechnicalValidation Authentication and syntactical and semantical validation
	// (Technical validation) are successful.
	TransactionStatusAcceptedTechnicalValidation TransactionStatus = "ACTC"
	// TransactionStatusAcceptedWithChange Instruction is accepted, but a change will be made,
	// such as date or remittance not sent.
	TransactionStatusAcceptedWithChange TransactionStatus = "ACWC"
	// TransactionStatusAcceptedWithoutPosting Payment instruction included in the credit transfer is accepted
	// without being posted to the creditor customer's account.
	TransactionStatusAcceptedWithoutPosting TransactionStatus = "ACWP"
	// TransactionStatusReceived Payment initiation has been received by the receiving agent.
	// Technical validation has started.
	TransactionStatusReceived TransactionStatus = "RCVD"
	// TransactionStatusPending Payment initiation or individual transaction
	// included in the payment initiation is pending and in progress for signing.
	// Further checks (and status update) are performed, the payment can still be Rejected or Cancelled if one of the following scenarios occur:
	//
	// PSU cancels the payment at login page > ACTC →  The payment is not signed within 30 minutes →  RJCT
	// PSU closes the web browser at login page > ACTC > The payment is not signed within 30 minutes > RJCT
	// PSU logs in > PDNG > PSU cancels the payment at Overview page > CANC
	// PSU logs in > PDNG > PSU cancels the payment at Overview page > PDNG > The payment is not signed within 30 minutes > RJCT
	// *ACTC - AcceptedTechnicalValidation, RJCT - Rejected, CANC - Cancelled
	TransactionStatusPending TransactionStatus = "PDNG"
	// TransactionStatusRejected Payment initiation or individual transaction
	// included in the payment initiation has been rejected.
	TransactionStatusRejected TransactionStatus = "RJCT"
	// TransactionStatusCancelled Payment initiation has been cancelled before execution.
	// This status is only applicable for future dated payments that have been successfully cancelled.
	TransactionStatusCancelled TransactionStatus = "CANC"
	// TransactionStatusAcceptedFundsChecked Preceding check of technical validation and customer profile was successful,
	// and an automatic funds check was positive.
	TransactionStatusAcceptedFundsChecked TransactionStatus = "ACFC"
	// TransactionStatusPartiallyAcceptedTechnical The payment initiation needs multiple authentications,
	// where some but not yet all have been performed.
	// Syntactical and semantical validations are successful.
	TransactionStatusPartiallyAcceptedTechnical TransactionStatus = "PATC"
	// TransactionStatusPartiallyAccepted A number of transactions have been accepted,
	// whereas another number of transactions have not yet achieved accepted status.
	TransactionStatusPartiallyAccepted TransactionStatus = "PART"
)

func ParseTransactionStatus(status string) TransactionStatus {
	switch status {
	case "ACCC":
		return TransactionStatusCreditorAcceptedSettlementCompleted
	case "ACCP":
		return TransactionStatusAcceptedCustomerProfile
	case "ACSC":
		return TransactionStatusAcceptedSettlementCompleted
	case "ACSP":
		return TransactionStatusAcceptedSettlementInProcess
	case "ACTC":
		return TransactionStatusAcceptedTechnicalValidation
	case "ACWC":
		return TransactionStatusAcceptedWithChange
	case "ACWP":
		return TransactionStatusAcceptedWithoutPosting
	case "RCVD":
		return TransactionStatusReceived
	case "PDNG":
		return TransactionStatusPending
	case "RJCT":
		return TransactionStatusRejected
	case "CANC":
		return TransactionStatusCancelled
	case "ACFC":
		return TransactionStatusAcceptedFundsChecked
	case "PATC":
		return TransactionStatusPartiallyAcceptedTechnical
	case "PART":
		return TransactionStatusPartiallyAccepted
	}
	return TransactionStatusNone
}

type Parser interface {
	ExtractMessageAndMetadataFromIsoMessage(msg []byte) (message messages.Iso20022Message, metadata Metadata, references *Metadata, err error)
	GetTransactionStatus(isoMsg messages.Iso20022Message) TransactionStatus
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
