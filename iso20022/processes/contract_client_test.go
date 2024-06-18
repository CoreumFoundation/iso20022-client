package processes_test

import (
	"context"
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	coreumlogger "github.com/CoreumFoundation/coreum-tools/pkg/logger"
	coreumchainclient "github.com/CoreumFoundation/coreum/v4/pkg/client"
	coreumchainconstant "github.com/CoreumFoundation/coreum/v4/pkg/config/constant"
	"github.com/CoreumFoundation/iso20022-client/iso20022/addressbook"
	"github.com/CoreumFoundation/iso20022-client/iso20022/compress"
	"github.com/CoreumFoundation/iso20022-client/iso20022/coreum"
	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
	"github.com/CoreumFoundation/iso20022-client/iso20022/processes"
)

func TestMain(m *testing.M) {
	coreum.SetSDKConfig(coreumchainconstant.AddressPrefixDev)
	m.Run()
}

//nolint:tparallel // the test is parallel, but test cases are not
func TestContractClient_Start(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name                  string
		contractClientBuilder func(ctrl *gomock.Controller) processes.ContractClient
		addressBookBuilder    func(ctrl *gomock.Controller) processes.AddressBook
		cryptographyBuilder   func(ctrl *gomock.Controller) processes.Cryptography
		parserBuilder         func(ctrl *gomock.Controller) processes.Parser
		run                   func(sendCh chan []byte, receiveCh chan []byte) error
	}{
		{
			name: "sending_one_message",
			contractClientBuilder: func(ctrl *gomock.Controller) processes.ContractClient {
				contractClientMock := NewMockContractClient(ctrl)
				contractClientMock.EXPECT().IsInitialized().Return(true)
				contractClientMock.EXPECT().GetUnreadMessages(gomock.Any(), gomock.Any(), gomock.Any()).Return([]coreum.Message{}, nil).AnyTimes()
				contractClientMock.EXPECT().IssueNFTClass(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil)
				contractClientMock.EXPECT().MintNFT(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil)
				contractClientMock.EXPECT().SendMessage(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, nil)
				return contractClientMock
			},
			addressBookBuilder: func(ctrl *gomock.Controller) processes.AddressBook {
				addressBookMock := NewMockAddressBook(ctrl)
				//addressBookMock.EXPECT().Update(gomock.Any()).Return(nil)
				addressBookMock.EXPECT().Lookup(addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "6P9YGUDF",
					},
				}).Return(&addressbook.Address{
					Bech32EncodedAddress: "devcore1kdujjkp8u0j9lww3n7qs7r5fwkljelvecsq43d",
					PublicKey:            "A2nYC1ZLFxVLL3kyGUGF4Hjlpwsd+FS7jmxIWahM0P5V",
					Party: addressbook.Party{
						Identification: addressbook.Identification{
							BusinessIdentifiersCode: "6P9YGUDF",
						},
					},
				}, true)
				return addressBookMock
			},
			cryptographyBuilder: func(ctrl *gomock.Controller) processes.Cryptography {
				cryptographyMock := NewMockCryptography(ctrl)
				cryptographyMock.EXPECT().GenerateSharedKeyByPrivateKeyName(gomock.Any(), gomock.Any(), gomock.Any()).Return([]byte("Thirty-two bytes long shared key"), nil)
				return cryptographyMock
			},
			parserBuilder: func(ctrl *gomock.Controller) processes.Parser {
				parserMock := NewMockParser(ctrl)
				parserMock.EXPECT().ExtractIdentificationFromIsoMessage(gomock.Any()).Return(&addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "6P9YGUDF",
					},
				}, nil)
				return parserMock
			},
			run: func(sendCh chan []byte, receiveCh chan []byte) error {
				sendCh <- []byte("hello world")
				return nil
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx := coreumlogger.WithLogger(context.Background(), coreumlogger.New(coreumlogger.ToolDefaultConfig))
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			t.Cleanup(cancel)

			ctrl := gomock.NewController(t)
			logMock := logger.NewAnyLogMock(ctrl)

			var contractClient processes.ContractClient
			if tt.contractClientBuilder != nil {
				contractClient = tt.contractClientBuilder(ctrl)
			}
			var addressBook processes.AddressBook
			if tt.addressBookBuilder != nil {
				addressBook = tt.addressBookBuilder(ctrl)
			}
			var cryptography processes.Cryptography
			if tt.cryptographyBuilder != nil {
				cryptography = tt.cryptographyBuilder(ctrl)
			}
			var parser processes.Parser
			if tt.parserBuilder != nil {
				parser = tt.parserBuilder(ctrl)
			}
			sendCh := make(chan []byte, 2)
			receiveCh := make(chan []byte, 2)
			go func() {
				go func() {
					runRrr := tt.run(sendCh, receiveCh)
					require.NoError(t, runRrr)
				}()
				<-ctx.Done()
				close(receiveCh)
				close(sendCh)
			}()

			cfg := processes.ContractClientProcessConfig{
				CoreumContractAddress: genAccount(),
				ClientAddress:         genAccount(),
				ClientKeyName:         "abc",
				PollInterval:          time.Second,
			}
			comp, err := compress.New()
			require.NoError(t, err)

			client, err := processes.NewContractClientProcess(cfg, logMock, comp, coreumchainclient.Context{}, addressBook, contractClient, cryptography, parser, sendCh, receiveCh)
			require.NoError(t, err)

			err = client.Start(ctx)
			if err == nil || !errors.Is(err, context.DeadlineExceeded) {
				require.NoError(t, err)
			}
		})
	}
}

func genAccount() sdk.AccAddress {
	return sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
}
