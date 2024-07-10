package processes

import (
	"context"
	"encoding/base64"
	"fmt"
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
	nftClassId     string
}

// NewContractClientProcess returns a new instance of the ContractClientProcess.
func NewContractClientProcess(cfg ContractClientProcessConfig, log logger.Logger, compressor *compress.Compressor, clientContext coreumchainclient.Context, addressBook AddressBook, contractClient ContractClient, cryptography Cryptography, parser Parser, messageQueue MessageQueue) (*ContractClientProcess, error) {
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
	}, nil
}

// Start starts the process.
func (p *ContractClientProcess) Start(ctx context.Context) error {
	p.log.Info(ctx, "Starting the contract client process")
	return parallel.Run(ctx, func(ctx context.Context, spawn parallel.SpawnFn) error {
		spawn("msg-receiver", parallel.Fail, func(ctx context.Context) error {
			ticker := time.NewTicker(p.cfg.PollInterval)
			for {
				select {
				case <-ticker.C:
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
					ticker.Stop()
					return errors.WithStack(ctx.Err())
				}
			}
		})
		spawn("msg-sender", parallel.Fail, func(ctx context.Context) error {
			for {
				msgs := p.messageQueue.PopFromSend(ctx, 10, time.Second)
				if len(msgs) == 0 {
					return errors.WithStack(ctx.Err())
				}

				messages := make([]*messageWithMetadata, 0, len(msgs))
				for _, msg := range msgs {
					message, err := p.extractMetadata(msg)
					if err != nil {
						p.messageQueue.SetStatus(message.Id, queue.StatusError)
						p.log.Error(
							ctx,
							"Failed to process the message",
							zap.Error(err),
							zap.Any("id", message.Id),
							zap.Any("destination", message.Destination),
							zap.Any("msg", msg),
						)
						continue
					}
					messages = append(messages, message)
				}

				if err := p.sendMessages(ctx, messages); err != nil {
					if errors.Is(err, context.Canceled) {
						p.log.Warn(
							ctx,
							"Context canceled during the message processing",
							zap.String("error", err.Error()),
						)
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
		})

		return nil
	})
}

func (p *ContractClientProcess) receiveMessages(ctx context.Context) error {
	limit := uint32(10)

	messages, err := p.contractClient.GetUnreadMessages(
		ctx,
		p.cfg.ClientAddress,
		&limit,
	)
	if err != nil {
		return err
	}

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

		p.log.Info(ctx, "Message received successfully", zap.String("sender", msg.Sender.String()))

		_, err = p.contractClient.MarkAsRead(
			ctx,
			p.cfg.ClientAddress,
			msg.Time,
		)
		if err != nil {
			return err
		}

		p.messageQueue.PushToReceive(data.Data)
	}

	return nil
}

func (p *ContractClientProcess) sendMessages(ctx context.Context, messages []*messageWithMetadata) error {
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
	sendMessages := make([]coreum.MessageWithDestination, 0, len(messages))

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
		sendMessages = append(sendMessages, coreum.MessageWithDestination{
			Destination: message.Destination,
			NFT:         nft,
		})
	}

	_, err := p.contractClient.BroadcastMessages(ctx, p.cfg.ClientAddress, mintMsgs...)
	if err != nil {
		return err
	}

	p.log.Info(ctx, "Messages sent successfully", zap.Int("count", len(messages)))
	_, err = p.contractClient.SendMessages(ctx, p.cfg.ClientAddress, sendMessages...)
	if err != nil {
		return err
	}

	p.log.Info(ctx, "Messages sent successfully", zap.Int("count", len(messages)))

	return nil
}

type messageWithMetadata struct {
	Id             string
	Destination    sdk.AccAddress
	PublicKeyBytes []byte
	Message        []byte
}

func (p *ContractClientProcess) extractMetadata(msg []byte) (*messageWithMetadata, error) {
	messageId, parsedTarget, err := p.parser.ExtractMetadataFromIsoMessage(msg)
	if err != nil {
		return nil, err
	}

	entry, found := p.addressBook.Lookup(*parsedTarget)
	if !found {
		return nil, errors.New("could not find the target party in the address book")
	}

	address, err := sdk.AccAddressFromBech32(entry.Bech32EncodedAddress)
	if err != nil {
		return nil, err
	}

	publicKeyBytes, err := base64.StdEncoding.DecodeString(entry.PublicKey)
	if err != nil {
		return nil, err
	}

	return &messageWithMetadata{messageId, address, publicKeyBytes, msg}, nil
}

func (p *ContractClientProcess) generateNftId() (string, string) {
	addr := p.cfg.ClientAddress.String()
	classId := fmt.Sprintf("%s-%s", collectionID, addr)
	NftPrefix := fmt.Sprintf("nft_%s", classId)
	id := fmt.Sprintf("%s_%d", NftPrefix, time.Now().UnixNano())
	return classId, id
}
