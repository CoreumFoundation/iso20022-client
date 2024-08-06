package messages

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/CoreumFoundation/iso20022-client/iso20022/addressbook"
	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
	"github.com/CoreumFoundation/iso20022-client/iso20022/messages/generated"
)

func TestParseIsoMessage(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	logMock := logger.NewAnyLogMock(ctrl)
	parser := NewParser(logMock, &generated.ConverterImpl{})

	tests := []struct {
		name            string
		messageFilePath string
		id              string
		party           *addressbook.Party
		hasError        bool
	}{
		{
			name:            "pacs008",
			messageFilePath: "testdata/pacs008-1.xml",
			id:              "P5607186 298",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "6P9YGUDF",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs008 within envelope",
			messageFilePath: "testdata/pacs008-2.xml",
			id:              "P5607186 298",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "6P9YGUDF",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs008 within envelope 2",
			messageFilePath: "testdata/pacs008-3.xml",
			id:              "BISSD20220717BLUEUSNY001B0123456789",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "ISSDINTL07X",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs008 within envelope 3",
			messageFilePath: "testdata/pacs008-6.xml",
			id:              "BISSD20220718BLUEUSNY001B8160187564",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "ISSDINTL07X",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs008 within envelope without head 2",
			messageFilePath: "testdata/pacs008-4.xml",
			id:              "20220717USDDSO0123456789BLUEUSNY001",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "GRENCHZZ002",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs008 within envelope without head 3",
			messageFilePath: "testdata/pacs008-7.xml",
			id:              "20220718USDDSA9153934686BLUEUSNY001",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "GRENCHZZ002",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs008 within envelope without header 2",
			messageFilePath: "testdata/pacs008-5.xml",
			id:              "20220717USDDSO0123456789BLUEUSNY001",
			party:           &addressbook.Party{},
			hasError:        true,
		},
		{
			name:            "pacs008 within envelope without header 3",
			messageFilePath: "testdata/pacs008-8.xml",
			id:              "20220718USDDSA9153934686BLUEUSNY001",
			party:           &addressbook.Party{},
			hasError:        true,
		},
		{
			name:            "pacs008 - First FIToFICustomerCreditTransfer",
			messageFilePath: "testdata/pacs008-9.xml",
			id:              "BBBB/150928-CCT/JPY/123",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "AAAAGB2L",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs008 - Second FIToFICustomerCreditTransfer",
			messageFilePath: "testdata/pacs008-10.xml",
			id:              "BBBB/150928-CT/EUR/912",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "EEEEDEFF",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs008 - Third FIToFICustomerCreditTransfer",
			messageFilePath: "testdata/pacs008-11.xml",
			id:              "BBBB/150928-CCT/USD/897",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "BBBBUS66",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs008 - Fourth FIToFICustomerCreditTransfer",
			messageFilePath: "testdata/pacs008-12.xml",
			id:              "EEEE/150929-EUR/059",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "DDDDBEBB",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs009 within envelope",
			messageFilePath: "testdata/pacs009-1.xml",
			id:              "BISSD20220720GRENCHZZ002B9194560468",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "ISSDINTL07X",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs009 within envelope without head",
			messageFilePath: "testdata/pacs009-2.xml",
			id:              "20220720CHFDSA9621795075GRENCHZZ002",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "CBFCCHZZ003",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs009 within envelope without header",
			messageFilePath: "testdata/pacs009-3.xml",
			id:              "20220720CHFDSA9621795075GRENCHZZ002",
			party:           &addressbook.Party{},
			hasError:        true,
		},
		{
			name:            "pacs002 within envelope",
			messageFilePath: "testdata/pacs002-1.xml",
			id:              "BISSD20220720CBFCCHZZ003B3144867881",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "ISSDINTL07X",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs002 within envelope 2",
			messageFilePath: "testdata/pacs002-4.xml",
			id:              "BISSD20220717ISSDINTL07XT0097036156",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "BLUEUSNY001",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs002 within envelope 3",
			messageFilePath: "testdata/pacs002-7.xml",
			id:              "BISSD20220717ISSDINTL07XT0097036156",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "BLUEUSNY001",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs002 within envelope without head",
			messageFilePath: "testdata/pacs002-2.xml",
			id:              "20220720CHFDSA6064961508CBFCCHZZ003",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "GRENCHZZ002",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs002 within envelope without head 2",
			messageFilePath: "testdata/pacs002-5.xml",
			id:              "20220717USDDSO9519156420ISSDINTL07X",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "BLUEUSNY001",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs002 within envelope without head 3",
			messageFilePath: "testdata/pacs002-8.xml",
			id:              "20220717USDDSO9519156420ISSDINTL07X",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "BLUEUSNY001",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs002 within envelope without header",
			messageFilePath: "testdata/pacs002-3.xml",
			id:              "BISSD20220719GRENCHZZ002B7173301669",
			party:           &addressbook.Party{},
			hasError:        true,
		},
		{
			name:            "pacs002 within envelope without header 2",
			messageFilePath: "testdata/pacs002-6.xml",
			id:              "",
			party:           &addressbook.Party{},
			hasError:        true,
		},
		{
			name:            "pacs002 within envelope without header 3",
			messageFilePath: "testdata/pacs002-9.xml",
			id:              "",
			party:           &addressbook.Party{},
			hasError:        true,
		},
		{
			name:            "pacs002 - FIToFIPaymentStatusReport",
			messageFilePath: "testdata/pacs002-10.xml",
			id:              "ABABUS23-STATUS-456/04",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "AAAAUS29",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs003 - FIToFICustomerDirectDebit",
			messageFilePath: "testdata/pacs003-1.xml",
			id:              "AAAA100628-123v",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "ABABUS23",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs004 - First PaymentReturn",
			messageFilePath: "testdata/pacs004-7.xml",
			id:              "CCCC/151122-PR007",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "BBBBIE2D",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs004 - Second PaymentReturn",
			messageFilePath: "testdata/pacs004-8.xml",
			id:              "BBBB/151122-PR05",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "AAAAGB2L",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs004 - First PaymentReturn - 2",
			messageFilePath: "testdata/pacs004-9.xml",
			id:              "BBBBUS39-RETURN-0123",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "ABABUS23",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs004 - Second PaymentReturn - 2",
			messageFilePath: "testdata/pacs004-10.xml",
			id:              "ABABUS23RETURN-546",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "AAAAUS29",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs004 within envelope - return",
			messageFilePath: "testdata/pacs004-1.xml",
			id:              "BISSD20220719GRENCHZZ002B7173301669",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "ISSDINTL07X",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs004 within envelope - cancellation",
			messageFilePath: "testdata/pacs004-4.xml",
			id:              "BISSD20220717GRENCHZZ002B7173301669",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "ISSDINTL07X",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs004 within envelope without head - return",
			messageFilePath: "testdata/pacs004-2.xml",
			id:              "20220719USDDSO0388509871GRENCHZZ002",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "BLUEUSNY001",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs004 within envelope without head - cancellation",
			messageFilePath: "testdata/pacs004-5.xml",
			id:              "20220717USDDSO9314441124GRENCHZZ002",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "BLUEUSNY001",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs004 within envelope without header - return",
			messageFilePath: "testdata/pacs004-3.xml",
			id:              "",
			party:           &addressbook.Party{},
			hasError:        true,
		},
		{
			name:            "pacs004 within envelope without header - cancellation",
			messageFilePath: "testdata/pacs004-6.xml",
			id:              "",
			party:           &addressbook.Party{},
			hasError:        true,
		},
		{
			name:            "pacs007 - FIToFIPaymentReversal",
			messageFilePath: "testdata/pacs007-1.xml",
			id:              "AAAAUS29-REVERSAL/0012",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "ABABUS23",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs010 - FinancialInstitutionDirectDebit",
			messageFilePath: "testdata/pacs010-1.xml",
			id:              "MMMM/151121 PPPP_OOOO_EUR",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "NNNNDEFF",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs028 - FIToFIPaymentStatusRequest",
			messageFilePath: "testdata/pacs028-1.xml",
			id:              "BBBB/150929-CCT/JPY/456",
			party: &addressbook.Party{
				Identification: addressbook.Identification{
					BusinessIdentifiersCode: "AAAAGB2L",
				},
			},
			hasError: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fileContent, err := os.ReadFile(tt.messageFilePath)
			require.NoError(t, err)

			metadata, err := parser.ExtractMetadataFromIsoMessage(fileContent)
			if tt.hasError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, metadata.Receiver)
				require.Equal(t, tt.id, metadata.ID)
				require.True(t, tt.party.Equal(*metadata.Receiver))
			}
		})
	}
}
