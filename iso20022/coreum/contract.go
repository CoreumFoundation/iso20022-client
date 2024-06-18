//nolint:tagliatelle // contract spec
package coreum

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	cosmoserrors "github.com/cosmos/cosmos-sdk/types/errors"
	sdktxtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/CoreumFoundation/coreum-tools/pkg/retry"
	"github.com/CoreumFoundation/coreum/v4/pkg/client"
	assetfttypes "github.com/CoreumFoundation/coreum/v4/x/asset/ft/types"
	nfttypes "github.com/CoreumFoundation/coreum/v4/x/asset/nft/types"
	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
)

// UserType is the contract user type.
type UserType string

// UserTypes.
const (
	UserTypeInitiator   UserType = "initiator"
	UserTypeDestination UserType = "destination"
)

// QueryMethod is the contract query method.
type QueryMethod string

// QueryMethods.
const (
	QueryMethodActiveSessions QueryMethod = "active_sessions"
	QueryMethodClosedSessions QueryMethod = "closed_sessions"
)

// NFTInfo is NFT information.
type NFTInfo struct {
	ClassId string `json:"class_id"`
	Id      string `json:"id"`
}

// Sessions is a list of sessions.
type Sessions struct {
	Sessions []Session `json:"sessions"`
}

// Session is session information.
type Session struct {
	Id                     string         `json:"id"`
	Initiator              sdk.AccAddress `json:"initiatior"`
	Destination            sdk.AccAddress `json:"destination"`
	Messages               []Message      `json:"messages"`
	FundsInEscrow          []sdk.Coin     `json:"funds_in_escrow"`
	ConfirmedByInitiator   bool           `json:"confirmed_by_initiatior"`
	ConfirmedByDestination bool           `json:"confirmed_by_destination"`
}

// Message is a single message details.
type Message struct {
	Sender   sdk.AccAddress `json:"sender"`
	Receiver sdk.AccAddress `json:"receiver"`
	Content  NFTInfo        `json:"content"`
}

// ******************** Internal transport object  ********************

type startSessionRequest struct {
	StartSession struct {
		Message     NFTInfo        `json:"message"`
		Destination sdk.AccAddress `json:"destination"`
	} `json:"start_session"`
}

type sendMessageRequest struct {
	SendMessage struct {
		SessionId string  `json:"session_id"`
		Message   NFTInfo `json:"message"`
	} `json:"send_message"`
}

type confirmSessionRequest struct {
	ConfirmSession struct {
		SessionId string `json:"session_id"`
	} `json:"confirm_session"`
}

type cancelSessionRequest struct {
	CancelSession struct {
		SessionId string `json:"session_id"`
	} `json:"cancel_session"`
}

type querySessionsRequest struct {
	StartAfterKey string         `json:"start_after_key,omitempty"`
	Limit         *uint32        `json:"limit,omitempty"`
	Address       sdk.AccAddress `json:"address"`
	UserType      UserType       `json:"user_type"`
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

// ******************** Execute ********************

// StartSession executes `start_session` method with transfer action.
func (c *ContractClient) StartSession(
	ctx context.Context, sender sdk.AccAddress, message NFTInfo, destination sdk.AccAddress,
) (*sdk.TxResponse, error) {
	req := startSessionRequest{}
	req.StartSession.Message = message
	req.StartSession.Destination = destination

	txRes, err := c.execute(ctx, sender, execRequest{
		Body: req,
		//Body: map[ExecMethod]startSessionRequest{
		//	ExecMethodStartSession: req,
		//},
	})
	if err != nil {
		return nil, err
	}

	return txRes, nil
}

// SendMessage executes `send_message` method with transfer action.
func (c *ContractClient) SendMessage(
	ctx context.Context, sender sdk.AccAddress, sessionId string, message NFTInfo,
) (*sdk.TxResponse, error) {
	req := sendMessageRequest{}
	req.SendMessage.SessionId = sessionId
	req.SendMessage.Message = message

	txRes, err := c.execute(ctx, sender, execRequest{
		Body: req,
	})
	if err != nil {
		return nil, err
	}

	return txRes, nil
}

// ConfirmSession executes `start_session` method with transfer action.
func (c *ContractClient) ConfirmSession(
	ctx context.Context, sender sdk.AccAddress, sessionId string,
) (*sdk.TxResponse, error) {
	req := confirmSessionRequest{}
	req.ConfirmSession.SessionId = sessionId

	txRes, err := c.execute(ctx, sender, execRequest{
		Body: req,
	})
	if err != nil {
		return nil, err
	}

	return txRes, nil
}

// CancelSession executes `start_session` method with transfer action.
func (c *ContractClient) CancelSession(
	ctx context.Context, sender sdk.AccAddress, sessionId string,
) (*sdk.TxResponse, error) {
	req := cancelSessionRequest{}
	req.CancelSession.SessionId = sessionId

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
	classId, Id string,
	data *types.Any,
) (*sdk.TxResponse, error) {
	msgIssueClass := &nfttypes.MsgMint{
		Sender:    sender.String(),
		ClassID:   classId,
		ID:        Id,
		Data:      data,
		Recipient: c.cfg.ContractAddress.String(),
	}

	txRes, err := client.BroadcastTx(ctx, c.clientCtx.WithFromAddress(sender), c.getTxFactory(), msgIssueClass)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to mint the NFT, classId:%s, id:%s", classId, Id)
	}
	c.log.Info(ctx, "NFT minted successfully")

	return txRes, nil
}

// ******************** Query ********************

// GetActiveSessions returns a list of all active sessions.
func (c *ContractClient) GetActiveSessions(
	ctx context.Context,
	address sdk.AccAddress,
	userType UserType,
	startAfterKey string,
	limit *uint32,
) ([]Session, error) {
	var response Sessions
	err := c.query(ctx, map[QueryMethod]querySessionsRequest{
		QueryMethodActiveSessions: {
			StartAfterKey: startAfterKey,
			Limit:         limit,
			Address:       address,
			UserType:      userType,
		},
	}, &response)
	if err != nil {
		return nil, err
	}

	return response.Sessions, nil
}

// GetClosedSessions returns a list of all closed sessions.
func (c *ContractClient) GetClosedSessions(
	ctx context.Context,
	address sdk.AccAddress,
	userType UserType,
	startAfterKey string,
	limit *uint32,
) ([]Session, error) {
	var response Sessions
	err := c.query(ctx, map[QueryMethod]querySessionsRequest{
		QueryMethodClosedSessions: {
			StartAfterKey: startAfterKey,
			Limit:         limit,
			Address:       address,
			UserType:      userType,
		},
	}, &response)
	if err != nil {
		return nil, err
	}

	return response.Sessions, nil
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
