// Code generated by GoComply XSD2Go for iso20022; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:head.001.001.02
package head_001_001_02

import (
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/xmldsig"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/pkg/iso"
)

// XSD ComplexType declarations

type AddressType3Choice struct {
	Cd    *AddressType2Code        `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Cd,omitempty"`
	Prtry *GenericIdentification30 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Prtry,omitempty"`
}

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 FinInstnId"`
	BrnchId    *BranchData3                         `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 BrnchId,omitempty"`
}

type BranchData3 struct {
	Id      *Max35Text       `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Id,omitempty"`
	LEI     *LEIIdentifier   `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 LEI,omitempty"`
	Nm      *Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Nm,omitempty"`
	PstlAdr *PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 PstlAdr,omitempty"`
}

type BusinessApplicationHeader5 struct {
	CharSet    *UnicodeChartsCode           `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 CharSet,omitempty"`
	Fr         Party44Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Fr"`
	To         Party44Choice                `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 To"`
	BizMsgIdr  Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 BizMsgIdr"`
	MsgDefIdr  Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 MsgDefIdr"`
	BizSvc     *Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 BizSvc,omitempty"`
	CreDt      iso.ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 CreDt"`
	CpyDplct   *CopyDuplicate1Code          `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 CpyDplct,omitempty"`
	PssblDplct *YesNoIndicator              `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 PssblDplct,omitempty"`
	Prty       *BusinessMessagePriorityCode `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Prty,omitempty"`
	Sgntr      *Sgntr                       `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Sgntr,omitempty"`
}

type BusinessApplicationHeaderV02 struct {
	CharSet    *UnicodeChartsCode            `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 CharSet,omitempty"`
	Fr         Party44Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Fr"`
	To         Party44Choice                 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 To"`
	BizMsgIdr  Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 BizMsgIdr"`
	MsgDefIdr  Max35Text                     `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 MsgDefIdr"`
	BizSvc     *Max35Text                    `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 BizSvc,omitempty"`
	MktPrctc   *ImplementationSpecification1 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 MktPrctc,omitempty"`
	CreDt      iso.ISODateTime               `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 CreDt"`
	BizPrcgDt  *iso.ISODateTime              `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 BizPrcgDt,omitempty"`
	CpyDplct   *CopyDuplicate1Code           `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 CpyDplct,omitempty"`
	PssblDplct *YesNoIndicator               `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 PssblDplct,omitempty"`
	Prty       *BusinessMessagePriorityCode  `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Prty,omitempty"`
	Sgntr      *Sgntr                        `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Sgntr,omitempty"`
	Rltd       []*BusinessApplicationHeader5 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Rltd,omitempty"`
}

type ClearingSystemIdentification2Choice struct {
	Cd    *ExternalClearingSystemIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Cd,omitempty"`
	Prtry *Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Prtry,omitempty"`
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 ClrSysId,omitempty"`
	MmbId    Max35Text                            `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 MmbId"`
}

type Contact4 struct {
	NmPrfx    *NamePrefix2Code             `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 NmPrfx,omitempty"`
	Nm        *Max140Text                  `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Nm,omitempty"`
	PhneNb    *PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 PhneNb,omitempty"`
	MobNb     *PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 MobNb,omitempty"`
	FaxNb     *PhoneNumber                 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 FaxNb,omitempty"`
	EmailAdr  *Max2048Text                 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 EmailAdr,omitempty"`
	EmailPurp *Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 EmailPurp,omitempty"`
	JobTitl   *Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 JobTitl,omitempty"`
	Rspnsblty *Max35Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Rspnsblty,omitempty"`
	Dept      *Max70Text                   `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Dept,omitempty"`
	Othr      []*OtherContact1             `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Othr,omitempty"`
	PrefrdMtd *PreferredContactMethod1Code `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 PrefrdMtd,omitempty"`
}

type DateAndPlaceOfBirth1 struct {
	BirthDt     iso.ISODate `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 BirthDt"`
	PrvcOfBirth *Max35Text  `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 PrvcOfBirth,omitempty"`
	CityOfBirth Max35Text   `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 CityOfBirth"`
	CtryOfBirth CountryCode `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 CtryOfBirth"`
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    *ExternalFinancialInstitutionIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Cd,omitempty"`
	Prtry *Max35Text                                       `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Prtry,omitempty"`
}

type FinancialInstitutionIdentification18 struct {
	BICFI       *BICFIDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 BICFI,omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 ClrSysMmbId,omitempty"`
	LEI         *LEIIdentifier                       `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 LEI,omitempty"`
	Nm          *Max140Text                          `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Nm,omitempty"`
	PstlAdr     *PostalAddress24                     `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 PstlAdr,omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Othr,omitempty"`
}

type GenericFinancialIdentification1 struct {
	Id      Max35Text                                 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Id"`
	SchmeNm *FinancialIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 SchmeNm,omitempty"`
	Issr    *Max35Text                                `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Issr,omitempty"`
}

type GenericIdentification30 struct {
	Id      Exact4AlphaNumericText `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Id"`
	Issr    Max35Text              `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Issr"`
	SchmeNm *Max35Text             `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 SchmeNm,omitempty"`
}

type GenericOrganisationIdentification1 struct {
	Id      Max35Text                                    `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Id"`
	SchmeNm *OrganisationIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 SchmeNm,omitempty"`
	Issr    *Max35Text                                   `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Issr,omitempty"`
}

type GenericPersonIdentification1 struct {
	Id      Max35Text                              `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Id"`
	SchmeNm *PersonIdentificationSchemeName1Choice `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 SchmeNm,omitempty"`
	Issr    *Max35Text                             `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Issr,omitempty"`
}

type ImplementationSpecification1 struct {
	Regy Max350Text  `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Regy"`
	Id   Max2048Text `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Id"`
}

type OrganisationIdentification29 struct {
	AnyBIC *AnyBICDec2014Identifier              `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 AnyBIC,omitempty"`
	LEI    *LEIIdentifier                        `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 LEI,omitempty"`
	Othr   []*GenericOrganisationIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Othr,omitempty"`
}

type OrganisationIdentificationSchemeName1Choice struct {
	Cd    *ExternalOrganisationIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Cd,omitempty"`
	Prtry *Max35Text                               `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Prtry,omitempty"`
}

type OtherContact1 struct {
	ChanlTp Max4Text    `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 ChanlTp"`
	Id      *Max128Text `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Id,omitempty"`
}

type Party38Choice struct {
	OrgId  *OrganisationIdentification29 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 OrgId,omitempty"`
	PrvtId *PersonIdentification13       `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 PrvtId,omitempty"`
}

type Party44Choice struct {
	OrgId *PartyIdentification135                       `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 OrgId,omitempty"`
	FIId  *BranchAndFinancialInstitutionIdentification6 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 FIId,omitempty"`
}

type PartyIdentification135 struct {
	Nm        *Max140Text      `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Nm,omitempty"`
	PstlAdr   *PostalAddress24 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 PstlAdr,omitempty"`
	Id        *Party38Choice   `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Id,omitempty"`
	CtryOfRes *CountryCode     `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 CtryOfRes,omitempty"`
	CtctDtls  *Contact4        `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 CtctDtls,omitempty"`
}

type PersonIdentification13 struct {
	DtAndPlcOfBirth *DateAndPlaceOfBirth1           `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 DtAndPlcOfBirth,omitempty"`
	Othr            []*GenericPersonIdentification1 `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Othr,omitempty"`
}

type PersonIdentificationSchemeName1Choice struct {
	Cd    *ExternalPersonIdentification1Code `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Cd,omitempty"`
	Prtry *Max35Text                         `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Prtry,omitempty"`
}

type PostalAddress24 struct {
	AdrTp       *AddressType3Choice `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 AdrTp,omitempty"`
	Dept        *Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Dept,omitempty"`
	SubDept     *Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 SubDept,omitempty"`
	StrtNm      *Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 StrtNm,omitempty"`
	BldgNb      *Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 BldgNb,omitempty"`
	BldgNm      *Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 BldgNm,omitempty"`
	Flr         *Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Flr,omitempty"`
	PstBx       *Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 PstBx,omitempty"`
	Room        *Max70Text          `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Room,omitempty"`
	PstCd       *Max16Text          `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 PstCd,omitempty"`
	TwnNm       *Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 TwnNm,omitempty"`
	TwnLctnNm   *Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 TwnLctnNm,omitempty"`
	DstrctNm    *Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 DstrctNm,omitempty"`
	CtrySubDvsn *Max35Text          `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 CtrySubDvsn,omitempty"`
	Ctry        *CountryCode        `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 Ctry,omitempty"`
	AdrLine     []*Max70Text        `xml:"urn:iso:std:iso:20022:tech:xsd:head.001.001.02 AdrLine,omitempty"`
}

type Sgntr struct {
	Signature *xmldsig.Signature `xml:"http://www.w3.org/2000/09/xmldsig# Signature"`
}

// XSD SimpleType declarations

type AddressType2Code string

const AddressType2CodeAddr AddressType2Code = "ADDR"
const AddressType2CodePbox AddressType2Code = "PBOX"
const AddressType2CodeHome AddressType2Code = "HOME"
const AddressType2CodeBizz AddressType2Code = "BIZZ"
const AddressType2CodeMlto AddressType2Code = "MLTO"
const AddressType2CodeDlvy AddressType2Code = "DLVY"

type AnyBICDec2014Identifier string

type BICFIDec2014Identifier string

type BusinessMessagePriorityCode string

type CopyDuplicate1Code string

const CopyDuplicate1CodeCodu CopyDuplicate1Code = "CODU"
const CopyDuplicate1CodeCopy CopyDuplicate1Code = "COPY"
const CopyDuplicate1CodeDupl CopyDuplicate1Code = "DUPL"

type CountryCode string

type Exact4AlphaNumericText string

type ExternalClearingSystemIdentification1Code string

type ExternalFinancialInstitutionIdentification1Code string

type ExternalOrganisationIdentification1Code string

type ExternalPersonIdentification1Code string

type LEIIdentifier string

type Max128Text string

type Max140Text string

type Max16Text string

type Max2048Text string

type Max350Text string

type Max35Text string

type Max4Text string

type Max70Text string

type NamePrefix2Code string

const NamePrefix2CodeDoct NamePrefix2Code = "DOCT"
const NamePrefix2CodeMadm NamePrefix2Code = "MADM"
const NamePrefix2CodeMiss NamePrefix2Code = "MISS"
const NamePrefix2CodeMist NamePrefix2Code = "MIST"
const NamePrefix2CodeMiks NamePrefix2Code = "MIKS"

type PhoneNumber string

type PreferredContactMethod1Code string

const PreferredContactMethod1CodeLett PreferredContactMethod1Code = "LETT"
const PreferredContactMethod1CodeMail PreferredContactMethod1Code = "MAIL"
const PreferredContactMethod1CodePhon PreferredContactMethod1Code = "PHON"
const PreferredContactMethod1CodeFaxx PreferredContactMethod1Code = "FAXX"
const PreferredContactMethod1CodeCell PreferredContactMethod1Code = "CELL"

type UnicodeChartsCode string

type YesNoIndicator bool
