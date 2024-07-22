package messages

import (
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/head_001_001_01"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/head_001_001_02"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_002_001_07"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_002_001_08"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_002_001_10"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_002_001_11"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_003_001_08"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_004_001_10"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_007_001_10"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_008_001_06"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_008_001_08"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_008_001_09"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_008_001_12"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_009_001_09"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_010_001_04"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_028_001_04"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_028_001_06"
	"github.com/CoreumFoundation/iso20022-client/iso20022/messages/converter"
)

//go:generate go run github.com/jmattheis/goverter/cmd/goverter gen ./...

// goverter:converter
// goverter:skipCopySameType
type Converter interface {
	ConvertFromHead00100101(source *head_001_001_01.BranchAndFinancialInstitutionIdentification5) *converter.BranchAndFinancialInstitutionIdentification5
	ConvertFromHead00100102(source *head_001_001_02.BranchAndFinancialInstitutionIdentification6) *converter.BranchAndFinancialInstitutionIdentification6
	ConvertFromPacs00200107(source *pacs_002_001_07.BranchAndFinancialInstitutionIdentification5) *converter.BranchAndFinancialInstitutionIdentification5
	ConvertFromPacs00200108(source *pacs_002_001_08.BranchAndFinancialInstitutionIdentification5) *converter.BranchAndFinancialInstitutionIdentification5
	ConvertFromPacs00200110(source *pacs_002_001_10.BranchAndFinancialInstitutionIdentification6) *converter.BranchAndFinancialInstitutionIdentification6
	ConvertFromPacs00200111(source *pacs_002_001_11.BranchAndFinancialInstitutionIdentification6) *converter.BranchAndFinancialInstitutionIdentification6
	ConvertFromPacs00300108(source *pacs_003_001_08.BranchAndFinancialInstitutionIdentification6) *converter.BranchAndFinancialInstitutionIdentification6
	ConvertFromPacs00400110(source *pacs_004_001_10.BranchAndFinancialInstitutionIdentification6) *converter.BranchAndFinancialInstitutionIdentification6
	ConvertFromPacs00700110(source *pacs_007_001_10.BranchAndFinancialInstitutionIdentification6) *converter.BranchAndFinancialInstitutionIdentification6
	ConvertFromPacs00800106(source *pacs_008_001_06.BranchAndFinancialInstitutionIdentification5) *converter.BranchAndFinancialInstitutionIdentification5
	ConvertFromPacs00800108(source *pacs_008_001_08.BranchAndFinancialInstitutionIdentification6) *converter.BranchAndFinancialInstitutionIdentification6
	ConvertFromPacs00800109(source *pacs_008_001_09.BranchAndFinancialInstitutionIdentification6) *converter.BranchAndFinancialInstitutionIdentification6
	ConvertFromPacs00800112(source *pacs_008_001_12.BranchAndFinancialInstitutionIdentification8) *converter.BranchAndFinancialInstitutionIdentification8
	ConvertFromPacs00900109(source *pacs_009_001_09.BranchAndFinancialInstitutionIdentification6) *converter.BranchAndFinancialInstitutionIdentification6
	ConvertFromPacs01000104(source *pacs_010_001_04.BranchAndFinancialInstitutionIdentification6) *converter.BranchAndFinancialInstitutionIdentification6
	ConvertFromPacs02800104(source *pacs_028_001_04.BranchAndFinancialInstitutionIdentification6) *converter.BranchAndFinancialInstitutionIdentification6
	ConvertFromPacs02800106(source *pacs_028_001_06.BranchAndFinancialInstitutionIdentification8) *converter.BranchAndFinancialInstitutionIdentification8
}
