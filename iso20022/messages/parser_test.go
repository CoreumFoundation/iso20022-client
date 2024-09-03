package messages

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_008_001_08"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/supl_xxx_001_01"
	"github.com/CoreumFoundation/iso20022-client/iso20022/addressbook"
	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
	"github.com/CoreumFoundation/iso20022-client/iso20022/messages/generated"
	"github.com/CoreumFoundation/iso20022-client/iso20022/processes"
)

type metadata struct {
	uetr     string
	id       string
	receiver *addressbook.Party
	sender   *addressbook.Party
}

func TestParseIsoMessage(t *testing.T) {
	t.Parallel()

	requireT := require.New(t)
	ctrl := gomock.NewController(t)
	logMock := logger.NewAnyLogMock(ctrl)
	parser := NewParser(logMock, &generated.ConverterImpl{})

	tests := []struct {
		name            string
		messageFilePath string
		metadata        metadata
		references      *metadata
		status          processes.TransactionStatus
		hasError        bool
	}{
		{
			name:            "pacs008",
			messageFilePath: "testdata/pacs008-1.xml",
			metadata: metadata{
				uetr: "00000000-0000-4000-8000-000000000000",
				id:   "P5607186 298",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "6P9YGUDF",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "B61NZT4Y",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs008 within envelope",
			messageFilePath: "testdata/pacs008-2.xml",
			metadata: metadata{
				uetr: "00000000-0000-4000-8000-000000000000",
				id:   "P5607186 299",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "6P9YGUDF",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "B61NZT4Y",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs008 within envelope 2",
			messageFilePath: "testdata/pacs008-3.xml",
			metadata: metadata{
				uetr: "c0bf145e-3858-1f95-999c-a35994c1223a",
				id:   "BISSD20220717BLUEUSNY001B0123456789",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "GRENCHZZ002",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BLUEUSNY001",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs008 within envelope 3",
			messageFilePath: "testdata/pacs008-6.xml",
			metadata: metadata{
				uetr: "6e2a125b-a8ec-6c16-ca9e-b28937eff09b",
				id:   "BISSD20220718BLUEUSNY001B8160187564",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "GRENCHZZ002",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BLUEUSNY001",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs008 within envelope without head 2",
			messageFilePath: "testdata/pacs008-4.xml",
			metadata: metadata{
				uetr: "c0bf145e-3858-1f95-999c-a35994c1223a",
				id:   "20220717USDDSO0123456789BLUEUSNY001",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "GRENCHZZ002",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BLUEUSNY001",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs008 within envelope without head 3",
			messageFilePath: "testdata/pacs008-7.xml",
			metadata: metadata{
				uetr: "6e2a125b-a8ec-6c16-ca9e-b28937eff09b",
				id:   "20220718USDDSA9153934686BLUEUSNY001",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "GRENCHZZ002",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BLUEUSNY001",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs008 within envelope without header 2",
			messageFilePath: "testdata/pacs008-5.xml",
			metadata: metadata{
				id:       "20220717USDDSO0123456789BLUEUSNY001",
				receiver: &addressbook.Party{},
				sender:   &addressbook.Party{},
			},
			hasError: true,
		},
		{
			name:            "pacs008 within envelope without header 3",
			messageFilePath: "testdata/pacs008-8.xml",
			metadata: metadata{
				id:       "20220718USDDSA9153934686BLUEUSNY001",
				receiver: &addressbook.Party{},
				sender:   &addressbook.Party{},
			},
			hasError: true,
		},
		{
			name:            "pacs008 - First FIToFICustomerCreditTransfer",
			messageFilePath: "testdata/pacs008-9.xml",
			metadata: metadata{
				uetr: "6606f642-8337-d475-6759-6b578d118669",
				id:   "BBBB/150928-CCT/JPY/123",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "AAAAGB2L",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BBBBUS33",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs008 - Second FIToFICustomerCreditTransfer",
			messageFilePath: "testdata/pacs008-10.xml",
			metadata: metadata{
				uetr: "07534004-4700-b40c-20b2-6c83971b8709",
				id:   "BBBB/150928-CT/EUR/912",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "EEEEDEFF",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BBBBUS33",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs008 - Third FIToFICustomerCreditTransfer",
			messageFilePath: "testdata/pacs008-11.xml",
			metadata: metadata{
				uetr: "b6eeed18-6047-be88-cd95-833985ebde7d",
				id:   "BBBB/150928-CCT/USD/897",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BBBBUS66",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BBBBUS33",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs008 - Fourth FIToFICustomerCreditTransfer",
			messageFilePath: "testdata/pacs008-12.xml",
			metadata: metadata{
				uetr: "07534004-4700-b40c-20b2-6c83971b8709",
				id:   "EEEE/150929-EUR/059",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "DDDDBEBB",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "EEEEDEFF",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs008 - Supplementary Data",
			messageFilePath: "testdata/pacs008-13.xml",
			metadata: metadata{
				uetr: "00000000-0000-4000-8000-000000000000",
				id:   "P5607186 298",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "6P9YGUDF",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "B61NZT4Y",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs009 within envelope",
			messageFilePath: "testdata/pacs009-1.xml",
			metadata: metadata{
				uetr: "0751b2f8-99f5-b945-5461-4149b822ddfa",
				id:   "BISSD20220720GRENCHZZ002B9194560468",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "CBFCCHZZ003",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "GRENCHZZ002",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs009 within envelope without head",
			messageFilePath: "testdata/pacs009-2.xml",
			metadata: metadata{
				uetr: "0751b2f8-99f5-b945-5461-4149b822ddfa",
				id:   "20220720CHFDSA9621795075GRENCHZZ002",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "CBFCCHZZ003",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "GRENCHZZ002",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs009 within envelope without header",
			messageFilePath: "testdata/pacs009-3.xml",
			metadata: metadata{
				id:       "20220720CHFDSA9621795075GRENCHZZ002",
				receiver: &addressbook.Party{},
				sender:   &addressbook.Party{},
			},
			hasError: true,
		},
		{
			name:            "pacs002 within envelope",
			messageFilePath: "testdata/pacs002-1.xml",
			metadata: metadata{
				id: "BISSD20220720CBFCCHZZ003B3144867881",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "GRENCHZZ002",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "CBFCCHZZ003",
					},
				},
			},
			references: &metadata{
				uetr: "0751b2f8-99f5-b945-5461-4149b822ddfa",
				id:   "CHFDSA20220720GRENCHZZ002B380494409",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "CBFCCHZZ003",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "GRENCHZZ002",
					},
				},
			},
			status:   processes.TransactionStatusAcceptedTechnicalValidation,
			hasError: false,
		},
		{
			name:            "pacs002 within envelope 2",
			messageFilePath: "testdata/pacs002-4.xml",
			metadata: metadata{
				id: "BISSD20220717ISSDINTL07XT0097036156",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BLUEUSNY001",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "ISSDINTL07X",
					},
				},
			},
			status:   processes.TransactionStatusAcceptedTechnicalValidation,
			hasError: false,
		},
		{
			name:            "pacs002 within envelope 3",
			messageFilePath: "testdata/pacs002-7.xml",
			metadata: metadata{
				id: "BISSD20220717ISSDINTL07XT0097036156",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BLUEUSNY001",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "ISSDINTL07X",
					},
				},
			},
			status:   processes.TransactionStatusAcceptedTechnicalValidation,
			hasError: false,
		},
		{
			name:            "pacs002 within envelope without head",
			messageFilePath: "testdata/pacs002-2.xml",
			metadata: metadata{
				id: "20220720CHFDSA6064961508CBFCCHZZ003",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "GRENCHZZ002",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "CBFCCHZZ003",
					},
				},
			},
			references: &metadata{
				uetr: "0751b2f8-99f5-b945-5461-4149b822ddfa",
				id:   "CHFDSA20220720GRENCHZZ002B380494409",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "CBFCCHZZ003",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "GRENCHZZ002",
					},
				},
			},
			status:   processes.TransactionStatusAcceptedTechnicalValidation,
			hasError: false,
		},
		{
			name:            "pacs002 within envelope without head 2",
			messageFilePath: "testdata/pacs002-5.xml",
			metadata: metadata{
				id: "20220717USDDSO9519156420ISSDINTL07X",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BLUEUSNY001",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "ISSDINTL07X",
					},
				},
			},
			status:   processes.TransactionStatusAcceptedTechnicalValidation,
			hasError: false,
		},
		{
			name:            "pacs002 within envelope without head 3",
			messageFilePath: "testdata/pacs002-8.xml",
			metadata: metadata{
				id: "20220717USDDSO9519156420ISSDINTL07X",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BLUEUSNY001",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "ISSDINTL07X",
					},
				},
			},
			status:   processes.TransactionStatusAcceptedTechnicalValidation,
			hasError: false,
		},
		{
			name:            "pacs002 within envelope without header",
			messageFilePath: "testdata/pacs002-3.xml",
			metadata: metadata{
				id:       "BISSD20220719GRENCHZZ002B7173301669",
				receiver: &addressbook.Party{},
				sender:   &addressbook.Party{},
			},
			hasError: true,
		},
		{
			name:            "pacs002 within envelope without header 2",
			messageFilePath: "testdata/pacs002-6.xml",
			metadata: metadata{
				id:       "",
				receiver: &addressbook.Party{},
				sender:   &addressbook.Party{},
			},
			hasError: true,
		},
		{
			name:            "pacs002 within envelope without header 3",
			messageFilePath: "testdata/pacs002-9.xml",
			metadata: metadata{
				id:       "",
				receiver: &addressbook.Party{},
				sender:   &addressbook.Party{},
			},
			hasError: true,
		},
		{
			name:            "pacs002 - FIToFIPaymentStatusReport",
			messageFilePath: "testdata/pacs002-10.xml",
			metadata: metadata{
				id: "ABABUS23-STATUS-456/04",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "AAAAUS29",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "ABABUS23",
					},
				},
			},
			references: &metadata{
				uetr: "4ceba488-ed0d-8bd7-c184-e0ee5c3c119c",
				id:   "AAAA100628-123v",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "ABABUS23",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "AAAAUS29",
					},
				},
			},
			status:   processes.TransactionStatusRejected,
			hasError: false,
		},
		{
			name:            "pacs003 - FIToFICustomerDirectDebit",
			messageFilePath: "testdata/pacs003-1.xml",
			metadata: metadata{
				uetr: "79b431e1-3f57-4cf4-708a-bf2697eb1ae7",
				id:   "AAAA100628-123v",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "ABABUS23",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "AAAAUS29",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs004 - First PaymentReturn",
			messageFilePath: "testdata/pacs004-7.xml",
			metadata: metadata{
				id: "CCCC/151122-PR007",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BBBBIE2D",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "CCCCIE2D",
					},
				},
			},
			references: &metadata{
				uetr: "f30b990f-928d-1629-8004-723039de15d6",
				id:   "BBBB/151109-CBJ056",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "CCCCIE2D",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BBBBIE2D",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs004 - Second PaymentReturn",
			messageFilePath: "testdata/pacs004-8.xml",
			metadata: metadata{
				id: "BBBB/151122-PR05",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "AAAAGB2L",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BBBBIE2D",
					},
				},
			},
			references: &metadata{
				uetr: "f30b990f-928d-1629-8004-723039de15d6",
				id:   "AAAA/151109-CCT/EUR443",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BBBBIE2D",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "AAAAGB2L",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs004 - First PaymentReturn - 2",
			messageFilePath: "testdata/pacs004-9.xml",
			metadata: metadata{
				id: "BBBBUS39-RETURN-0123",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "ABABUS23",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BBBBUS39",
					},
				},
			},
			references: &metadata{
				id: "ABABUS23-589cd",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BBBBUS39",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "ABABUS23",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs004 - Second PaymentReturn - 2",
			messageFilePath: "testdata/pacs004-10.xml",
			metadata: metadata{
				id: "ABABUS23RETURN-546",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "AAAAUS29",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "ABABUS23",
					},
				},
			},
			references: &metadata{
				id: "AAAA060327-123v",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "ABABUS23",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "AAAAUS29",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs004 within envelope - return",
			messageFilePath: "testdata/pacs004-1.xml",
			metadata: metadata{
				id: "BISSD20220719GRENCHZZ002B7173301669",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BLUEUSNY001",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "GRENCHZZ002",
					},
				},
			},
			references: &metadata{
				uetr: "c0bf145e-3858-1f95-999c-a35994c1223a",
				id:   "20220717USDDSO0123456789BLUEUSNY001",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "GRENCHZZ002",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BLUEUSNY001",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs004 within envelope - cancellation",
			messageFilePath: "testdata/pacs004-4.xml",
			metadata: metadata{
				id: "BISSD20220717GRENCHZZ002B7173301669",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BLUEUSNY001",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "GRENCHZZ002",
					},
				},
			},
			references: &metadata{
				uetr: "c0bf145e-3858-1f95-999c-a35994c1223a",
				id:   "20220717USDDSO0123456789BLUEUSNY001",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "GRENCHZZ002",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BLUEUSNY001",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs004 within envelope without head - return",
			messageFilePath: "testdata/pacs004-2.xml",
			metadata: metadata{
				id: "20220719USDDSO0388509871GRENCHZZ002",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BLUEUSNY001",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "GRENCHZZ002",
					},
				},
			},
			references: &metadata{
				uetr: "c0bf145e-3858-1f95-999c-a35994c1223a",
				id:   "20220717USDDSO0123456789BLUEUSNY001",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "GRENCHZZ002",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BLUEUSNY001",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs004 within envelope without head - cancellation",
			messageFilePath: "testdata/pacs004-5.xml",
			metadata: metadata{
				id: "20220717USDDSO9314441124GRENCHZZ002",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BLUEUSNY001",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "GRENCHZZ002",
					},
				},
			},
			references: &metadata{
				uetr: "c0bf145e-3858-1f95-999c-a35994c1223a",
				id:   "20220717USDDSO0123456789BLUEUSNY001",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "GRENCHZZ002",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BLUEUSNY001",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs004 within envelope without header - return",
			messageFilePath: "testdata/pacs004-3.xml",
			metadata: metadata{
				id:       "",
				receiver: &addressbook.Party{},
				sender:   &addressbook.Party{},
			},
			hasError: true,
		},
		{
			name:            "pacs004 within envelope without header - cancellation",
			messageFilePath: "testdata/pacs004-6.xml",
			metadata: metadata{
				id:       "",
				receiver: &addressbook.Party{},
				sender:   &addressbook.Party{},
			},
			hasError: true,
		},
		{
			name:            "pacs007 - FIToFIPaymentReversal",
			messageFilePath: "testdata/pacs007-1.xml",
			metadata: metadata{
				id: "AAAAUS29-REVERSAL/0012",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "ABABUS23",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "AAAAUS29",
					},
				},
			},
			references: &metadata{
				uetr: "4ceba488-ed0d-8bd7-c184-e0ee5c3c119c",
				id:   "AAAA120628-123v",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "AAAAUS29",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "ABABUS23",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs010 - FinancialInstitutionDirectDebit",
			messageFilePath: "testdata/pacs010-1.xml",
			metadata: metadata{
				uetr: "0d1af344-e77a-5ff6-044e-de04da466367",
				id:   "MMMM/151121 PPPP_OOOO_EUR",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "NNNNDEFF",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "MMMMBEBB",
					},
				},
			},
			hasError: false,
		},
		{
			name:            "pacs028 - FIToFIPaymentStatusRequest",
			messageFilePath: "testdata/pacs028-1.xml",
			metadata: metadata{
				id: "BBBB/150929-CCT/JPY/456",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "AAAAGB2L",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BBBBUS33",
					},
				},
			},
			references: &metadata{
				uetr: "6606f642-8337-d475-6759-6b578d118669",
				id:   "BBBB/150928-CCT/JPY/123",
				receiver: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "AAAAGB2L",
					},
				},
				sender: &addressbook.Party{
					Identification: addressbook.Identification{
						BusinessIdentifiersCode: "BBBBUS33",
					},
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
			requireT.NoError(err)

			msg, md, ref, _, err := parser.ExtractMessageAndMetadataFromIsoMessage(fileContent)
			if tt.hasError {
				requireT.Error(err)
			} else {
				requireT.NoError(err)
				requireT.NotNil(md.Receiver)
				requireT.Equal(tt.metadata.uetr, md.Uetr)
				requireT.Equal(tt.metadata.id, md.ID)
				requireT.True(tt.metadata.receiver.Equal(*md.Receiver))
				requireT.True(tt.metadata.sender.Equal(*md.Sender))
				if tt.references != nil || ref != nil {
					requireT.NotNil(md.Receiver)
					requireT.NotNil(tt.references)
					requireT.NotNil(ref)
					requireT.Equal(tt.references.uetr, ref.Uetr)
					requireT.Equal(tt.references.id, ref.ID)
					requireT.True(tt.references.receiver.Equal(*ref.Receiver))
					requireT.True(tt.references.sender.Equal(*ref.Sender))
				}
				transactionStatus := parser.GetTransactionStatus(msg)
				if (tt.status != "" && tt.status != processes.TransactionStatusNone) || transactionStatus != processes.TransactionStatusNone {
					requireT.Equal(tt.status, transactionStatus)
				}
			}
		})
	}
}

func TestParseSupplementaryDataFromIsoMessage(t *testing.T) {
	t.Parallel()

	requireT := require.New(t)
	ctrl := gomock.NewController(t)
	logMock := logger.NewAnyLogMock(ctrl)
	parser := NewParser(logMock, &generated.ConverterImpl{})

	fileContent, err := os.ReadFile("testdata/pacs008-13.xml")
	requireT.NoError(err)

	msg, _, _, suplParser, err := parser.ExtractMessageAndMetadataFromIsoMessage(fileContent)
	requireT.NoError(err)

	message, ok := msg.(*pacs_008_001_08.FIToFICustomerCreditTransferV08)
	requireT.True(ok)
	requireT.NotEmpty(message.CdtTrfTxInf)

	supl, err := suplParser.Parse(message.CdtTrfTxInf[0].SplmtryData[0].Envlp.Doc)
	requireT.NoError(err)

	sup, ok := supl.(*supl_xxx_001_01.CryptoCurrencyAndAmountType)
	requireT.True(ok)
	requireT.Equal(&supl_xxx_001_01.CryptoCurrencyAndAmountType{
		Value: 100000,
		Cccy:  "ABC-testcore1adst6w4e79tddzhcgaru2l2gms8jjep6a4caa7",
	}, sup)
}

func TestGetSupplementaryDataWithCorrectClearingSystem(t *testing.T) {
	t.Parallel()

	requireT := require.New(t)
	ctrl := gomock.NewController(t)
	logMock := logger.NewAnyLogMock(ctrl)
	parser := NewParser(logMock, &generated.ConverterImpl{})

	tests := []struct {
		name              string
		messageFilePath   string
		currencyAndAmount *supl_xxx_001_01.CryptoCurrencyAndAmountType
		hasError          bool
	}{
		{
			name:            "with_cryptocurrency_attribute",
			messageFilePath: "testdata/pacs008-14.xml",
			currencyAndAmount: &supl_xxx_001_01.CryptoCurrencyAndAmountType{
				Value: 499250,
				Cccy:  "ibc/E1E3674A0E4E1EF9C69646F9AF8D9497173821826074622D831BAB73CCB99A2D",
			},
			hasError: false,
		},
		{
			name:            "with_dti_attribute",
			messageFilePath: "testdata/pacs008-15.xml",
			currencyAndAmount: &supl_xxx_001_01.CryptoCurrencyAndAmountType{
				Value: 499250,
				Dti:   "KNNT25FGR",
			},
			hasError: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			fileContent, err := os.ReadFile(tt.messageFilePath)
			requireT.NoError(err)

			msg, _, _, suplParser, err := parser.ExtractMessageAndMetadataFromIsoMessage(fileContent)
			requireT.NoError(err)

			suplData, found := parser.GetSupplementaryDataWithCorrectClearingSystem(msg, "COREUM")
			requireT.True(found)
			requireT.NotEmpty(suplData)

			supl, err := suplParser.Parse(suplData)
			if tt.hasError {
				requireT.Error(err)
			} else {
				requireT.NoError(err)
			}

			sup, ok := supl.(*supl_xxx_001_01.CryptoCurrencyAndAmountType)
			requireT.True(ok)
			requireT.Equal(tt.currencyAndAmount, sup)
		})
	}
}
