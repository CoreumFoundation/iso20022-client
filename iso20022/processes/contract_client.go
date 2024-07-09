package processes

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/xsalsa20symmetric"
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
)

const collectionID = "ISOPayment"

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
				msgs := p.messageQueue.PopFromSend(ctx, 1, time.Second)
				if len(msgs) == 0 {
					return errors.WithStack(ctx.Err())
				}

				// TODO: Batch
				msg := msgs[0]
				id, destination, publicKey, err := p.extractMetadata(msg)
				if err != nil {
					p.log.Error(
						ctx,
						"Failed to process the message",
						zap.Error(err),
						zap.Any("id", id),
						zap.Any("destination", destination),
						zap.Any("msg", msg),
					)
					continue
				}

				if err = p.sendMessages(ctx, destination, publicKey, msg); err != nil {
					if errors.Is(err, context.Canceled) {
						p.log.Warn(
							ctx,
							"Context canceled during the message processing",
							zap.String("error", err.Error()),
						)
					} else {
						p.log.Error(
							ctx,
							"Failed to send the message",
							zap.Error(err),
							zap.Any("id", id),
							zap.Any("destination", destination),
							zap.Any("msg", msg),
						)
						continue
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

		data.Data, err = xsalsa20symmetric.DecryptSymmetric(data.Data, sharedKey)
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

func (p *ContractClientProcess) sendMessages(ctx context.Context, destination sdk.AccAddress, destinationPubKey, msg []byte) error {
	classId, id := p.generateNftId()

	sharedKey, err := p.cryptography.GenerateSharedKeyByPrivateKeyName(p.clientContext, p.cfg.ClientKeyName, destinationPubKey)
	if err != nil {
		return err
	}

	msg = p.compressor.Compress(msg)

	msg = xsalsa20symmetric.EncryptSymmetric(msg, sharedKey)

	data, err := types.NewAnyWithValue(&nfttypes.DataBytes{Data: msg})
	if err != nil {
		return err
	}

	if p.nftClassId == "" {
		_, err = p.contractClient.IssueNFTClass(
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

	_, err = p.contractClient.MintNFT(
		ctx,
		p.cfg.ClientAddress,
		strings.ToLower(classId),
		id,
		data,
	)
	if err != nil {
		return err
	}

	nft := coreum.NFTInfo{
		ClassId: strings.ToLower(classId),
		Id:      id,
	}

	_, err = p.contractClient.SendMessage(ctx, p.cfg.ClientAddress, destination, nft)
	if err != nil {
		return err
	}

	p.log.Info(ctx, "Message sent successfully", zap.String("receiver", destination.String()))

	return nil
}

func (p *ContractClientProcess) extractMetadata(msg []byte) (string, sdk.AccAddress, []byte, error) {
	messageId, parsedTarget, err := p.parser.ExtractMetadataFromIsoMessage(msg)
	if err != nil {
		return "", nil, nil, err
	}

	entry, found := p.addressBook.Lookup(*parsedTarget)
	if !found {
		return "", nil, nil, errors.New("could not find the target party in the address book")
	}

	address, err := sdk.AccAddressFromBech32(entry.Bech32EncodedAddress)
	if err != nil {
		return "", nil, nil, err
	}

	publicKeyBytes, err := base64.StdEncoding.DecodeString(entry.PublicKey)
	if err != nil {
		return "", nil, nil, err
	}

	return messageId, address, publicKeyBytes, nil
}

func (p *ContractClientProcess) generateNftId() (string, string) {
	// FIXME
	addr := p.cfg.ClientAddress.String()
	classId := fmt.Sprintf("%s-%s", collectionID, addr)
	NftPrefix := fmt.Sprintf("NFT_%s", classId)
	id := fmt.Sprintf("%s_%d", NftPrefix, time.Now().UnixNano())
	return classId, id
}
