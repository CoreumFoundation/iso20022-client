package processes

import (
	"context"
	"encoding/base64"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/gogoproto/proto"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/CoreumFoundation/coreum-tools/pkg/parallel"
	coreumchainclient "github.com/CoreumFoundation/coreum/v4/pkg/client"
	nfttypes "github.com/CoreumFoundation/coreum/v4/x/asset/nft/types"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/supl_xxx_001_01"
	"github.com/CoreumFoundation/iso20022-client/iso20022/compress"
	"github.com/CoreumFoundation/iso20022-client/iso20022/coreum"
	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
	"github.com/CoreumFoundation/iso20022-client/iso20022/queue"
)

const collectionID = "isopayment"

// ContractClientProcessConfig is ContractClientProcess config.
type ContractClientProcessConfig struct {
	CoreumContractAddress sdk.AccAddress
	ClientAddress         sdk.AccAddress
	ClientKeyName         string
	PollInterval          time.Duration
	Denom                 string
}

// ContractClientProcess is the process that sends and receives messages to and from the contract.
type ContractClientProcess struct {
	cfg            ContractClientProcessConfig
	log            logger.Logger
	compressor     *compress.Compressor
	clientContext  coreumchainclient.Context
	addressBook    AddressBook
	contractClient ContractClient
	cryptography   Cryptography
	parser         Parser
	messageQueue   MessageQueue
	dtif           Dtif
	nftClassId     string
}

// NewContractClientProcess returns a new instance of the ContractClientProcess.
func NewContractClientProcess(cfg ContractClientProcessConfig, log logger.Logger, compressor *compress.Compressor, clientContext coreumchainclient.Context, addressBook AddressBook, contractClient ContractClient, cryptography Cryptography, parser Parser, messageQueue MessageQueue, dtif Dtif) (*ContractClientProcess, error) {
	if cfg.CoreumContractAddress.Empty() {
		return nil, errors.Errorf("failed to init the process, the contract address is nil or empty")
	}
	if !contractClient.IsInitialized() {
		return nil, errors.Errorf("failed to init the process, the contract client is not initialized")
	}

	return &ContractClientProcess{
		cfg:            cfg,
		log:            log,
		compressor:     compressor,
		clientContext:  clientContext,
		addressBook:    addressBook,
		contractClient: contractClient,
		cryptography:   cryptography,
		parser:         parser,
		messageQueue:   messageQueue,
		dtif:           dtif,
	}, nil
}

// Start starts the process.
func (p *ContractClientProcess) Start(ctx context.Context) error {
	p.log.Info(ctx, "Starting the contract client process")
	return parallel.Run(ctx, func(ctx context.Context, spawn parallel.SpawnFn) error {
		spawn("msg-receiver", parallel.Continue, func(ctx context.Context) error {
			expiredSessionsTicker := time.NewTicker(time.Hour)
			messagesTicker := time.NewTicker(p.cfg.PollInterval)
			for {
				select {
				case <-expiredSessionsTicker.C:
					err := p.cancelExpiredSessions(ctx)
					if err != nil {
						p.log.Error(
							ctx,
							"Failed to cancel expired sessions",
							zap.Error(err),
						)
						continue
					}
				case <-messagesTicker.C:
					err := p.receiveMessages(ctx)
					if err != nil {
						if errors.Is(err, context.Canceled) {
							p.log.Warn(
								ctx,
								"Context canceled during receiving messages",
								zap.String("error", err.Error()),
							)
						} else {
							p.log.Error(
								ctx,
								"Failed to receive messages",
								zap.Error(err),
							)
							continue
						}
					}
				case <-ctx.Done():
					messagesTicker.Stop()
					expiredSessionsTicker.Stop()
					return errors.WithStack(ctx.Err())
				}
			}
		})
		spawn("msg-sender", parallel.Continue, func(ctx context.Context) error {
			for {
				msgs := p.messageQueue.PopFromSend(ctx, 10, time.Second)
				if len(msgs) == 0 {
					return errors.WithStack(ctx.Err())
				}

				messages := make([]*MessageWithMetadata, 0, len(msgs))
				for _, msg := range msgs {
					message, err := p.ExtractMetadata(msg)
					if err != nil {
						if message != nil {
							p.messageQueue.SetStatus(message.Id, queue.StatusError)
							p.log.Error(
								ctx,
								"Failed to process the message",
								zap.Error(err),
								zap.Any("uetr", message.Uetr),
								zap.Any("id", message.Id),
								zap.Any("source", message.Source),
								zap.Any("destination", message.Destination),
								zap.Any("msg", msg),
							)
						} else {
							p.log.Error(
								ctx,
								"Failed to process the message",
								zap.Error(err),
								zap.Any("msg", msg),
							)
						}
						continue
					}
					messages = append(messages, message)
				}

				if len(messages) > 0 {
					if err := p.sendMessages(ctx, messages); err != nil {
						if errors.Is(err, context.Canceled) || strings.Contains(err.Error(), "context canceled") {
							p.log.Warn(
								ctx,
								"Context canceled during the message processing",
								zap.String("error", err.Error()),
							)
							return nil
						} else {
							for _, message := range messages {
								p.messageQueue.SetStatus(message.Id, queue.StatusError)
								p.log.Error(
									ctx,
									"Failed to send the message",
									zap.Error(err),
									zap.Any("id", message.Id),
									zap.Any("destination", message.Destination),
									zap.Any("msg", message.Message),
								)
							}
							continue
						}
					} else {
						for _, message := range messages {
							p.messageQueue.SetStatus(message.Id, queue.StatusSent)
						}
					}
				}
			}
		})

		return nil
	})
}

func (p *ContractClientProcess) cancelExpiredSessions(ctx context.Context) error {
	limit := uint32(100)

	sessions, err := p.contractClient.GetActiveSessions(
		ctx,
		p.cfg.ClientAddress,
		coreum.UserTypeInitiator,
		nil,
		&limit,
	)
	if err != nil {
		return err
	}

	expireTime := uint64(time.Now().Add(-24 * time.Hour).UnixNano())

	cancelSessions := make([]coreum.CancelSession, 0)

	for _, session := range sessions {
		if !session.ConfirmedByDestination && !session.ConfirmedByInitiator && session.StartTime < expireTime {
			cancelSessions = append(cancelSessions, coreum.CancelSession{
				Uetr:        session.Uetr,
				Initiator:   session.Initiator,
				Destination: session.Destination,
			})
		}
	}

	if len(cancelSessions) > 0 {
		_, err = p.contractClient.CancelSessions(ctx, p.cfg.ClientAddress, cancelSessions...)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *ContractClientProcess) receiveMessages(ctx context.Context) error {
	limit := uint32(10)

	messages, err := p.contractClient.GetNewMessages(
		ctx,
		p.cfg.ClientAddress,
		&limit,
	)
	if err != nil {
		return err
	}

	lastReadTime := uint64(0)

	for _, msg := range messages {
		nft, err := p.contractClient.QueryNFT(ctx, msg.Content.ClassId, msg.Content.Id)
		if err != nil {
			p.log.Error(ctx, "could not get the NFT") // TODO
			continue
		}

		var data nfttypes.DataBytes

		err = proto.Unmarshal(nft.Value, &data)
		if err != nil {
			return err
		}

		entry, found := p.addressBook.LookupByAccountAddress(msg.Sender.String())
		if !found {
			p.log.Error(ctx, "could not find sender institute in the address book") // TODO
			continue
		}

		publicKeyBytes, err := base64.StdEncoding.DecodeString(entry.PublicKey)
		if err != nil {
			p.log.Error(ctx, "could not decode the sender public key", zap.Error(err)) // TODO
			continue
		}

		sharedKey, err := p.cryptography.GenerateSharedKeyByPrivateKeyName(p.clientContext, p.cfg.ClientKeyName, publicKeyBytes)
		if err != nil {
			p.log.Error(ctx, "could not calculate the shared key", zap.Error(err)) // TODO
			continue
		}

		data.Data, err = p.cryptography.DecryptSymmetric(data.Data, sharedKey)
		if err != nil {
			p.log.Error(ctx, "could not decrypt the message", zap.Error(err)) // TODO
			continue
		}

		data.Data, err = p.compressor.Decompress(data.Data)
		if err != nil {
			p.log.Error(ctx, "could not decompress the message", zap.Error(err)) // TODO
			continue
		}

		_, metadata, _, _, err := p.parser.ExtractMessageAndMetadataFromIsoMessage(data.Data)
		if err != nil {
			p.log.Error(ctx, "could not extract metadata the message", zap.Error(err)) // TODO
			continue
		}

		if msg.Time > lastReadTime {
			lastReadTime = msg.Time
		}

		if !metadata.Sender.Equal(entry.Party) {
			p.log.Error(ctx, "message sender is not verified") // TODO
			continue
		}

		p.log.Info(ctx, "Message received successfully", zap.String("sender", msg.Sender.String()))

		p.messageQueue.PushToReceive(data.Data)
	}

	if lastReadTime > 0 {
		_, err = p.contractClient.MarkAsRead(
			ctx,
			p.cfg.ClientAddress,
			lastReadTime,
		)
	}

	return err
}

func (p *ContractClientProcess) sendMessages(ctx context.Context, messages []*MessageWithMetadata) error {
	classId, id := p.generateNftId()

	if p.nftClassId == "" {
		_, err := p.contractClient.IssueNFTClass(
			ctx,
			p.cfg.ClientAddress,
			collectionID,
			classId,
			"ISOPayment-Messages",
		)
		if err != nil && !strings.Contains(err.Error(), "already used") {
			return err
		}

		p.nftClassId = classId
	}

	mintMsgs := make([]sdk.Msg, 0, len(messages))
	startSessions := make([]coreum.StartSession, 0, len(messages))
	confirmSessions := make([]coreum.ConfirmSession, 0, len(messages))
	cancelSessions := make([]coreum.CancelSession, 0, len(messages))
	sendMessages := make([]coreum.SendMessage, 0, len(messages))

	for _, message := range messages {
		sharedKey, err := p.cryptography.GenerateSharedKeyByPrivateKeyName(p.clientContext, p.cfg.ClientKeyName, message.PublicKeyBytes)
		if err != nil {
			return err
		}

		msg := p.compressor.Compress(message.Message)

		msg = p.cryptography.EncryptSymmetric(msg, sharedKey)

		data, err := types.NewAnyWithValue(&nfttypes.DataBytes{Data: msg})
		if err != nil {
			return err
		}

		mintMsgs = append(mintMsgs, &nfttypes.MsgMint{
			Sender:    p.cfg.ClientAddress.String(),
			ClassID:   classId,
			ID:        id,
			Data:      data,
			Recipient: p.cfg.CoreumContractAddress.String(),
		})

		nft := coreum.NFTInfo{
			ClassId: strings.ToLower(classId),
			Id:      id,
		}
		if !message.AttachedFunds.IsZero() {
			startSessions = append(startSessions, coreum.StartSession{
				Uetr:        message.Uetr,
				Message:     nft,
				Destination: message.Destination,
				Funds:       message.AttachedFunds,
			})
		}

		switch p.parser.GetTransactionStatus(message.ParsedMessage) {
		case TransactionStatusCreditorAcceptedSettlementCompleted, TransactionStatusAcceptedCustomerProfile,
			TransactionStatusAcceptedSettlementCompleted, TransactionStatusAcceptedSettlementInProcess,
			TransactionStatusAcceptedTechnicalValidation, TransactionStatusAcceptedWithChange,
			TransactionStatusAcceptedWithoutPosting, TransactionStatusAcceptedFundsChecked,
			TransactionStatusPartiallyAcceptedTechnical, TransactionStatusPartiallyAccepted:
			if message.references == nil {
				return errors.New("Could not find the referenced message")
			}
			if message.references.Receiver == nil {
				return errors.New("Could not find the original receiver")
			}
			origReceiverParty, ok := p.addressBook.Lookup(*message.references.Receiver)
			if !ok {
				return errors.New("Could not find the original receiver")
			}
			origReceiver, err := sdk.AccAddressFromBech32(origReceiverParty.Bech32EncodedAddress)
			if err != nil {
				return errors.Wrap(err, "Could not find the original receiver")
			}
			if message.references.Sender == nil {
				return errors.New("Could not find the original sender")
			}
			origSenderParty, ok := p.addressBook.Lookup(*message.references.Sender)
			if !ok {
				return errors.New("Could not find the original sender")
			}
			origSender, err := sdk.AccAddressFromBech32(origSenderParty.Bech32EncodedAddress)
			if err != nil {
				return errors.Wrap(err, "Could not find the original sender")
			}
			confirmSessions = append(confirmSessions, coreum.ConfirmSession{
				Uetr:        message.references.Uetr,
				Initiator:   origReceiver,
				Destination: origSender,
			})
		case TransactionStatusCancelled, TransactionStatusRejected:
			if message.references == nil {
				return errors.New("Could not find the referenced message")
			}
			if message.references.Receiver == nil {
				return errors.New("Could not find the original receiver")
			}
			origReceiverParty, ok := p.addressBook.Lookup(*message.references.Receiver)
			if !ok {
				return errors.New("Could not find the original receiver")
			}
			origReceiver, err := sdk.AccAddressFromBech32(origReceiverParty.Bech32EncodedAddress)
			if err != nil {
				return errors.Wrap(err, "Could not find the original receiver")
			}
			if message.references.Sender == nil {
				return errors.New("Could not find the original sender")
			}
			origSenderParty, ok := p.addressBook.Lookup(*message.references.Sender)
			if !ok {
				return errors.New("Could not find the original sender")
			}
			origSender, err := sdk.AccAddressFromBech32(origSenderParty.Bech32EncodedAddress)
			if err != nil {
				return errors.Wrap(err, "Could not find the original sender")
			}
			cancelSessions = append(cancelSessions, coreum.CancelSession{
				Uetr:        message.references.Uetr,
				Initiator:   origReceiver,
				Destination: origSender,
			})
		}

		sendMessages = append(sendMessages, coreum.SendMessage{
			Uetr:        message.Uetr,
			ID:          message.Id,
			Destination: message.Destination,
			Message:     nft,
		})
	}

	_, err := p.contractClient.BroadcastMessages(ctx, p.cfg.ClientAddress, mintMsgs...)
	if err != nil {
		return err
	}

	if len(startSessions) > 0 {
		_, err = p.contractClient.StartSessions(ctx, p.cfg.ClientAddress, startSessions...)
		if err != nil {
			return err
		}
	}

	if len(cancelSessions) > 0 {
		_, err = p.contractClient.CancelSessions(ctx, p.cfg.ClientAddress, cancelSessions...)
		if err != nil {
			return err
		}
	}

	if len(confirmSessions) > 0 {
		_, err = p.contractClient.ConfirmSessions(ctx, p.cfg.ClientAddress, confirmSessions...)
		if err != nil {
			return err
		}
	}

	if len(sendMessages) > 0 {
		_, err = p.contractClient.SendMessages(ctx, p.cfg.ClientAddress, sendMessages...)
		if err != nil {
			return err
		}

		p.log.Info(ctx, "Messages sent successfully", zap.Int("count", len(messages)))
	}

	return nil
}

func (p *ContractClientProcess) ExtractMetadata(rawMessage []byte) (*MessageWithMetadata, error) {
	message, metadata, references, suplParser, err := p.parser.ExtractMessageAndMetadataFromIsoMessage(rawMessage)
	if err != nil {
		return nil, err
	}

	result := MessageWithMetadata{
		metadata.Uetr,
		metadata.ID,
		nil,
		nil,
		nil,
		rawMessage,
		message,
		references,
		nil,
	}

	receiver, found := p.addressBook.Lookup(*metadata.Receiver)
	if !found {
		return &result, errors.New("could not find the receiver party in the address book")
	}

	result.Destination, err = sdk.AccAddressFromBech32(receiver.Bech32EncodedAddress)
	if err != nil {
		return &result, err
	}

	sender, found := p.addressBook.Lookup(*metadata.Receiver)
	if found {
		result.Source, err = sdk.AccAddressFromBech32(sender.Bech32EncodedAddress)
		if err != nil {
			return &result, err
		}
	}

	result.PublicKeyBytes, err = base64.StdEncoding.DecodeString(receiver.PublicKey)
	if err != nil {
		return &result, err
	}

	attachedFunds := sdk.NewCoins()
	suplData, found := p.parser.GetSupplementaryDataWithCorrectClearingSystem(message, "COREUM")
	if found {
		suplMsg, err := suplParser.Parse(suplData)
		if err != nil {
			return &result, err
		}
		supl, ok := suplMsg.(*supl_xxx_001_01.CryptoCurrencyAndAmountType)
		if ok {
			if supl.Cccy != "" {
				attachedFunds = attachedFunds.Add(sdk.NewCoin(string(supl.Cccy), sdk.NewInt(int64(supl.Value))))
			} else if supl.Dti != "" {
				denom, priceMultiplier, found := p.dtif.LookupByDTI(string(supl.Dti))
				if found {
					if priceMultiplier == nil {
						priceMultiplier = big.NewInt(1)
					}
					value := sdk.MustNewDecFromStr(strconv.FormatFloat(float64(supl.Value), 'f', -1, 64)).
						Mul(sdk.NewDecFromBigInt(priceMultiplier))
					if !value.IsInteger() {
						return nil, errors.New("The amount needs more precision than what the token supports")
					}
					attachedFunds = attachedFunds.Add(
						sdk.NewCoin(denom, value.TruncateInt()),
					)
				}
			}
		}
	}

	result.AttachedFunds = attachedFunds

	return &result, nil
}

func (p *ContractClientProcess) generateNftId() (string, string) {
	addr := p.cfg.ClientAddress.String()
	classId := fmt.Sprintf("%s-%s", collectionID, addr)
	NftPrefix := fmt.Sprintf("nft_%s", classId)
	id := fmt.Sprintf("%s_%d", NftPrefix, time.Now().UnixNano())
	return classId, id
}
