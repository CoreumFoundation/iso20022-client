//nolint:tagliatelle // contract spec
package coreum

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	cosmoserrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdktxtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/CoreumFoundation/coreum-tools/pkg/retry"
	"github.com/CoreumFoundation/coreum/v4/pkg/client"
	"github.com/CoreumFoundation/coreum/v4/testutil/event"
	assetfttypes "github.com/CoreumFoundation/coreum/v4/x/asset/ft/types"
	nfttypes "github.com/CoreumFoundation/coreum/v4/x/asset/nft/types"
	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
)

const (
	contractLabel = "iso20022"
)

// QueryMethod is the contract query method.
type QueryMethod string

// QueryMethods.
const (
	QueryMethodUnreadMessages QueryMethod = "unread_messages"
	QueryMethodReadMessages   QueryMethod = "read_messages"
)

// NFTInfo is NFT information.
type NFTInfo struct {
	ClassId string `json:"class_id"`
	Id      string `json:"id"`
}

// MessageWithDestination is a message with destination.
type MessageWithDestination struct {
	Destination sdk.AccAddress
	NFT         NFTInfo
	Funds       sdk.Coins
}

// Messages is a list of messages.
type Messages struct {
	Messages []Message `json:"messages"`
}

// Message is a single message details.
type Message struct {
	Sender        sdk.AccAddress `json:"sender"`
	Receiver      sdk.AccAddress `json:"receiver"`
	Time          uint64         `json:"time"`
	Content       NFTInfo        `json:"content"`
	FundsInEscrow []sdk.Coin     `json:"funds_in_escrow"`
}

// ******************** Internal transport object  ********************

type instantiateRequest struct{}

type sendMessageRequest struct {
	SendMessage struct {
		Destination sdk.AccAddress `json:"destination"`
		Message     NFTInfo        `json:"message"`
	} `json:"send_message"`
}

type markAsReadRequest struct {
	MarkAsRead struct {
		Until uint64 `json:"until"`
	} `json:"mark_as_read"`
}

type querySessionsRequest struct {
	StartAfterKey *string        `json:"start_after_key,omitempty"`
	Limit         *uint32        `json:"limit,omitempty"`
	Address       sdk.AccAddress `json:"address"`
}

type execRequest struct {
	Body  any
	Funds sdk.Coins
}

// ******************** Client ********************

// ContractClientConfig represent the ContractClient config.
type ContractClientConfig struct {
	ContractAddress       sdk.AccAddress
	GasAdjustment         float64
	GasPriceAdjustment    sdk.Dec
	PageLimit             uint32
	OutOfGasRetryDelay    time.Duration
	OutOfGasRetryAttempts uint32
	TxsQueryPageLimit     uint32
}

// DefaultContractClientConfig returns default ContractClient config.
func DefaultContractClientConfig(contractAddress sdk.AccAddress) ContractClientConfig {
	return ContractClientConfig{
		ContractAddress:       contractAddress,
		GasAdjustment:         1.4,
		GasPriceAdjustment:    sdk.MustNewDecFromStr("1.2"),
		PageLimit:             50,
		OutOfGasRetryDelay:    500 * time.Millisecond,
		OutOfGasRetryAttempts: 5,
		TxsQueryPageLimit:     1000,
	}
}

// ContractClient is the bridge contract client.
type ContractClient struct {
	cfg                ContractClientConfig
	log                logger.Logger
	clientCtx          client.Context
	wasmClient         wasmtypes.QueryClient
	assetftClient      assetfttypes.QueryClient
	cometServiceClient sdktxtypes.ServiceClient

	execMu sync.Mutex
}

// NewContractClient returns a new instance of the ContractClient.
func NewContractClient(cfg ContractClientConfig, log logger.Logger, clientCtx client.Context) *ContractClient {
	return &ContractClient{
		cfg: cfg,
		log: log,
		clientCtx: clientCtx.
			WithBroadcastMode(flags.BroadcastSync).
			WithAwaitTx(true).WithGasPriceAdjustment(cfg.GasPriceAdjustment).
			WithGasAdjustment(cfg.GasAdjustment),
		wasmClient:         wasmtypes.NewQueryClient(clientCtx),
		assetftClient:      assetfttypes.NewQueryClient(clientCtx),
		cometServiceClient: sdktxtypes.NewServiceClient(clientCtx),

		execMu: sync.Mutex{},
	}
}

// DeployAndInstantiate instantiates the contract.
func (c *ContractClient) DeployAndInstantiate(
	ctx context.Context,
	sender sdk.AccAddress,
	contractByteCodePath string,
) (sdk.AccAddress, error) {
	_, codeID, err := c.DeployContract(ctx, sender, contractByteCodePath)
	if err != nil {
		return nil, err
	}

	reqPayload, err := json.Marshal(instantiateRequest{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to marshal instantiate payload")
	}

	msg := &wasmtypes.MsgInstantiateContract{
		Sender: sender.String(),
		Admin:  sender.String(),
		CodeID: codeID,
		Label:  contractLabel,
		Msg:    reqPayload,
	}

	c.log.Info(ctx, "Instantiating contract.", zap.Any("msg", msg))
	res, err := client.BroadcastTx(ctx, c.clientCtx.WithFromAddress(sender), c.getTxFactory(), msg)
	if err != nil {
		return nil, errors.Wrap(err, "failed to deploy bytecode")
	}

	contractAddr, err := event.FindStringEventAttribute(
		res.Events, wasmtypes.EventTypeInstantiate, wasmtypes.AttributeKeyContractAddr,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to find contract address in the tx result")
	}

	sdkContractAddr, err := sdk.AccAddressFromBech32(contractAddr)
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert contract address to sdk.AccAddress")
	}
	c.log.Info(ctx, "The contract is instantiated.", zap.String("address", sdkContractAddr.String()))

	return sdkContractAddr, nil
}

// DeployContract deploys the contract bytecode.
func (c *ContractClient) DeployContract(
	ctx context.Context,
	sender sdk.AccAddress,
	contractByteCodePath string,
) (*sdk.TxResponse, uint64, error) {
	c.log.Info(
		ctx,
		"Deploying contract",
		zap.String("path", contractByteCodePath),
	)

	contactByteCode, err := os.ReadFile(contractByteCodePath)
	if err != nil {
		return nil, 0, errors.Wrapf(err, "failed to get contract bytecode by path:%s", contractByteCodePath)
	}

	msgStoreCode := &wasmtypes.MsgStoreCode{
		Sender:       sender.String(),
		WASMByteCode: contactByteCode,
	}
	c.log.Info(ctx, "Deploying contract bytecode.")

	txRes, err := client.BroadcastTx(ctx, c.clientCtx.WithFromAddress(sender), c.getTxFactory(), msgStoreCode)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to deploy wasm bytecode")
	}
	// handle the genereate only case
	if txRes == nil {
		return nil, 0, nil
	}
	codeID, err := event.FindUint64EventAttribute(txRes.Events, wasmtypes.EventTypeStoreCode, wasmtypes.AttributeKeyCodeID)
	if err != nil {
		return nil, 0, errors.Wrap(err, "failed to find code ID in the tx result")
	}
	c.log.Info(ctx, "The contract bytecode is deployed.", zap.Uint64("codeID", codeID))

	return txRes, codeID, nil
}

// MigrateContract calls the executes the contract migration.
func (c *ContractClient) MigrateContract(
	ctx context.Context,
	sender sdk.AccAddress,
	codeID uint64,
) (*sdk.TxResponse, error) {
	msgMigrate := &wasmtypes.MsgMigrateContract{
		Sender:   sender.String(),
		Contract: c.GetContractAddress().String(),
		CodeID:   codeID,
		Msg:      []byte("{}"),
	}

	txRes, err := client.BroadcastTx(ctx, c.clientCtx.WithFromAddress(sender), c.getTxFactory(), msgMigrate)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to migrate contract, codeID:%d", codeID)
	}
	c.log.Info(ctx, "Contract migrated successfully")

	return txRes, nil
}

// SetContractAddress sets the client contract address if it was not set before.
func (c *ContractClient) SetContractAddress(contractAddress sdk.AccAddress) error {
	if c.cfg.ContractAddress != nil {
		return errors.New("contract address is already set")
	}

	c.cfg.ContractAddress = contractAddress

	return nil
}

// GetContractAddress returns contract address used by the client.
func (c *ContractClient) GetContractAddress() sdk.AccAddress {
	return c.cfg.ContractAddress
}

// IsInitialized returns true if address used by the client is set.
func (c *ContractClient) IsInitialized() bool {
	return !c.cfg.ContractAddress.Empty()
}

// BroadcastMessages broadcasts messages.
func (c *ContractClient) BroadcastMessages(
	ctx context.Context,
	sender sdk.AccAddress,
	messages ...sdk.Msg,
) (*sdk.TxResponse, error) {
	return client.BroadcastTx(ctx, c.clientCtx.WithFromAddress(sender), c.getTxFactory(), messages...)
}

// ******************** Execute ********************

// SendMessage executes `send_message` method with transfer action.
func (c *ContractClient) SendMessage(
	ctx context.Context, sender, destination sdk.AccAddress, message NFTInfo,
) (*sdk.TxResponse, error) {
	req := sendMessageRequest{}
	req.SendMessage.Destination = destination
	req.SendMessage.Message = message

	txRes, err := c.execute(ctx, sender, execRequest{
		Body: req,
	})
	if err != nil {
		return nil, err
	}

	return txRes, nil
}

// SendMessages executes multiple `send_message` method with transfer action.
func (c *ContractClient) SendMessages(
	ctx context.Context, sender sdk.AccAddress, messages ...MessageWithDestination,
) (*sdk.TxResponse, error) {
	reqs := make([]execRequest, len(messages))
	for i, msg := range messages {
		req := sendMessageRequest{}
		req.SendMessage.Destination = msg.Destination
		req.SendMessage.Message = msg.NFT
		reqs[i] = execRequest{
			Body:  req,
			Funds: msg.Funds,
		}
	}

	txRes, err := c.execute(ctx, sender, reqs...)
	if err != nil {
		return nil, err
	}

	return txRes, nil
}

// MarkAsRead executes `mark_as_read` method with transfer action.
func (c *ContractClient) MarkAsRead(
	ctx context.Context, sender sdk.AccAddress, until uint64,
) (*sdk.TxResponse, error) {
	req := markAsReadRequest{}
	req.MarkAsRead.Until = until

	txRes, err := c.execute(ctx, sender, execRequest{
		Body: req,
	})
	if err != nil {
		return nil, err
	}

	return txRes, nil
}

// IssueNFTClass issues the nft class.
func (c *ContractClient) IssueNFTClass(
	ctx context.Context,
	sender sdk.AccAddress,
	symbol, name, description string,
) (*sdk.TxResponse, error) {
	msgIssueClass := &nfttypes.MsgIssueClass{
		Issuer:      sender.String(),
		Symbol:      symbol,
		Name:        name,
		Description: description,
		RoyaltyRate: sdk.ZeroDec(),
	}

	txRes, err := client.BroadcastTx(ctx, c.clientCtx.WithFromAddress(sender), c.getTxFactory(), msgIssueClass)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to issue NFT class, symbol:%s, name:%s", symbol, name)
	}
	c.log.Info(ctx, "NFT class issued successfully")

	return txRes, nil
}

// MintNFT mints the nft.
func (c *ContractClient) MintNFT(
	ctx context.Context,
	sender sdk.AccAddress,
	classId, id string,
	data *types.Any,
) (*sdk.TxResponse, error) {
	msgIssueClass := &nfttypes.MsgMint{
		Sender:    sender.String(),
		ClassID:   classId,
		ID:        id,
		Data:      data,
		Recipient: c.cfg.ContractAddress.String(),
	}

	txRes, err := client.BroadcastTx(ctx, c.clientCtx.WithFromAddress(sender), c.getTxFactory(), msgIssueClass)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to mint the NFT, classId:%s, id:%s", classId, id)
	}
	c.log.Info(ctx, "NFT minted successfully")

	return txRes, nil
}

// ******************** Query ********************

// GetUnreadMessages returns a list of all unread messages.
func (c *ContractClient) GetUnreadMessages(
	ctx context.Context,
	address sdk.AccAddress,
	limit *uint32,
) ([]Message, error) {
	var response Messages
	err := c.query(ctx, map[QueryMethod]querySessionsRequest{
		QueryMethodUnreadMessages: {
			Limit:   limit,
			Address: address,
		},
	}, &response)
	if err != nil {
		return nil, err
	}

	return response.Messages, nil
}

// GetReadMessages returns a list of all read messages.
func (c *ContractClient) GetReadMessages(
	ctx context.Context,
	address sdk.AccAddress,
	startAfterKey string,
	limit *uint32,
) ([]Message, error) {
	var response Messages
	err := c.query(ctx, map[QueryMethod]querySessionsRequest{
		QueryMethodReadMessages: {
			StartAfterKey: &startAfterKey,
			Limit:         limit,
			Address:       address,
		},
	}, &response)
	if err != nil {
		return nil, err
	}

	return response.Messages, nil
}

// QueryNFT queries the nft.
func (c *ContractClient) QueryNFT(
	ctx context.Context,
	classId, id string,
) (*types.Any, error) {
	nftClient := nft.NewQueryClient(c.clientCtx)
	resp, err := nftClient.NFT(ctx, &nft.QueryNFTRequest{
		ClassId: classId,
		Id:      id,
	})
	if err != nil {
		return nil, err
	}
	return resp.Nft.Data, nil
}

func (c *ContractClient) execute(
	ctx context.Context,
	sender sdk.AccAddress,
	requests ...execRequest,
) (*sdk.TxResponse, error) {
	if c.cfg.ContractAddress == nil {
		return nil, errors.New("failed to execute with empty contract address")
	}
	c.execMu.Lock()
	defer c.execMu.Unlock()

	msgs := make([]sdk.Msg, 0, len(requests))
	for _, req := range requests {
		payload, err := json.Marshal(req.Body)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to marshal payload, request:%+v", req.Body)
		}
		c.log.Debug(ctx, "Executing contract", zap.String("payload", string(payload)))
		msg := &wasmtypes.MsgExecuteContract{
			Sender:   sender.String(),
			Contract: c.cfg.ContractAddress.String(),
			Msg:      payload,
			Funds:    req.Funds,
		}
		msgs = append(msgs, msg)
	}

	clientCtx := c.clientCtx.WithFromAddress(sender)
	if clientCtx.GenerateOnly() {
		unsignedTx, err := client.GenerateUnsignedTx(ctx, clientCtx, c.getTxFactory(), msgs...)
		if err != nil {
			return nil, err
		}

		txData, err := clientCtx.TxConfig().TxJSONEncoder()(unsignedTx.GetTx())
		if err != nil {
			return nil, err
		}

		return nil, clientCtx.PrintString(fmt.Sprintf("%s\n", txData))
	}

	var res *sdk.TxResponse
	outOfGasRetryAttempt := uint32(1)
	err := retry.Do(ctx, c.cfg.OutOfGasRetryDelay, func() error {
		var err error
		res, err = client.BroadcastTx(ctx, clientCtx.WithFromAddress(sender), c.getTxFactory(), msgs...)
		if err == nil {
			return nil
		}
		// stop if we have reached the max retries
		if outOfGasRetryAttempt >= c.cfg.OutOfGasRetryAttempts {
			return err
		}
		if cosmoserrors.ErrOutOfGas.Is(err) {
			outOfGasRetryAttempt++
			c.log.Info(ctx, "Out of gas, retrying Coreum tx execution")
			return retry.Retryable(errors.Wrapf(err, "retry tx execution, out of gas"))
		}

		return errors.Wrapf(err, "failed to execute transaction, message:%+v", msgs)
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ContractClient) query(ctx context.Context, request, response any) error {
	if c.cfg.ContractAddress == nil {
		return errors.New("failed to execute with empty contract address")
	}

	payload, err := json.Marshal(request)
	if err != nil {
		return errors.Wrapf(err, "failed to marshal query request")
	}
	c.log.Debug(ctx, "Querying contract", zap.String("payload", string(payload)))

	query := &wasmtypes.QuerySmartContractStateRequest{
		Address:   c.cfg.ContractAddress.String(),
		QueryData: payload,
	}
	resp, err := c.wasmClient.SmartContractState(ctx, query)
	if err != nil {
		return errors.Wrapf(err, "query failed, request:%+v", request)
	}

	c.log.Debug(ctx, "Query is succeeded", zap.String("data", string(resp.Data)))
	if err := json.Unmarshal(resp.Data, response); err != nil {
		return errors.Wrapf(
			err,
			"failed to unmarshal wasm contract response, request:%s, response:%s",
			string(payload),
			string(resp.Data),
		)
	}

	return nil
}

func (c *ContractClient) getTxFactory() client.Factory {
	return client.Factory{}.
		WithKeybase(c.clientCtx.Keyring()).
		WithChainID(c.clientCtx.ChainID()).
		WithTxConfig(c.clientCtx.TxConfig()).
		WithMemo("client: iso20022").
		WithSimulateAndExecute(true)
}

// ******************** Contract error ********************

// IsPaymentErrorError returns true if error is `PaymentError`.
func IsPaymentErrorError(err error) bool {
	return isError(err, "Payment error")
}

// IsInvalidDestinationError returns true if error is `InvalidDestination`.
func IsInvalidDestinationError(err error) bool {
	return isError(err, "InvalidDestination")
}

// IsSessionNotFoundError returns true if error is `SessionNotFound`.
func IsSessionNotFoundError(err error) bool {
	return isError(err, "SessionNotFound")
}

// IsUnauthorizedError returns true if error is `Unauthorized`.
func IsUnauthorizedError(err error) bool {
	return isError(err, "Unauthorized")
}

// IsNFTClassNotFoundError returns true if error is `NFTClassNotFound`.
func IsNFTClassNotFoundError(err error) bool {
	return isError(err, "NFTClassNotFound")
}

// IsSenderIsNotNFTIssuerError returns true if error is `SenderIsNotNFTIssuer`.
func IsSenderIsNotNFTIssuerError(err error) bool {
	return isError(err, "SenderIsNotNFTIssuer")
}

// IsNFTIdNotFoundError returns true if error is `NFTIdNotFound`.
func IsNFTIdNotFoundError(err error) bool {
	return isError(err, "NFTIdNotFound")
}

// IsContractIsNotOwnerOfNFTError returns true if error is `ContractIsNotOwnerOfNFT`.
func IsContractIsNotOwnerOfNFTError(err error) bool {
	return isError(err, "ContractIsNotOwnerOfNFT")
}

// IsSessionAlreadyConfirmedError returns true if error is `SessionAlreadyConfirmed`.
func IsSessionAlreadyConfirmedError(err error) bool {
	return isError(err, "SessionAlreadyConfirmed")
}

func isError(err error, errorString string) bool {
	return err != nil && strings.Contains(err.Error(), errorString)
}
