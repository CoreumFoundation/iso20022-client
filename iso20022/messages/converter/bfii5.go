package converter

import "github.com/CoreumFoundation/iso20022-client/iso20022/addressbook"

type BranchAndFinancialInstitutionIdentification5 struct {
	FinInstnId FinancialInstitutionIdentification8 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08 FinInstnId"`
	BrnchId    *BranchData2                        `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08 BrnchId,omitempty"`
}

func (agent *BranchAndFinancialInstitutionIdentification5) ToParty() *addressbook.Party {
	res := new(addressbook.Party)
	if agent == nil {
		return res
	}

	if agent.BrnchId != nil {
		res.Branch = &addressbook.Branch{}
		if agent.BrnchId.Id != nil {
			res.Branch.Id = *agent.BrnchId.Id
		}
		if agent.BrnchId.Nm != nil {
			res.Branch.Name = *agent.BrnchId.Nm
		}
		if agent.BrnchId.PstlAdr != nil {
			res.Branch.PostalAddress = agent.BrnchId.PstlAdr.ToPostalAddress()
		}
	}

	if agent.FinInstnId.BICFI != nil {
		res.Identification.BusinessIdentifiersCode = *agent.FinInstnId.BICFI
	}
	if agent.FinInstnId.ClrSysMmbId != nil {
		res.Identification.ClearingSystemMemberIdentification = &addressbook.ClearingSystemMemberIdentification{
			MemberId: agent.FinInstnId.ClrSysMmbId.MmbId,
		}
		if res.Identification.ClearingSystemMemberIdentification.ClearingSystemId != nil {
			res.Identification.ClearingSystemMemberIdentification.ClearingSystemId = &addressbook.ClearingSystemId{
				Code: *agent.FinInstnId.ClrSysMmbId.ClrSysId.Cd,
			}
			if agent.FinInstnId.ClrSysMmbId.ClrSysId.Prtry != nil {
				res.Identification.ClearingSystemMemberIdentification.ClearingSystemId.Proprietary = *agent.FinInstnId.ClrSysMmbId.ClrSysId.Prtry
			}
		}
	}
	if agent.FinInstnId.Nm != nil && agent.FinInstnId.PstlAdr != nil {
		res.Identification.Name = *agent.FinInstnId.Nm
		if agent.FinInstnId.PstlAdr != nil {
			res.Identification.PostalAddress = agent.FinInstnId.PstlAdr.ToPostalAddress()
		}

	}
	if agent.FinInstnId.Othr != nil {
		res.Identification.Other = &addressbook.Other{
			Id: agent.FinInstnId.Othr.Id,
		}
		if agent.FinInstnId.Othr.Issr != nil {
			res.Identification.Other.Issuer = *agent.FinInstnId.Othr.Issr
		}
		if agent.FinInstnId.Othr.SchmeNm != nil {
			res.Identification.Other.SchemeName = &addressbook.SchemeName{
				Code:        *agent.FinInstnId.Othr.SchmeNm.Cd,
				Proprietary: *agent.FinInstnId.Othr.SchmeNm.Prtry,
			}
		}
	}

	return res
}

type FinancialInstitutionIdentification8 struct {
	BICFI       *string                              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08 BICFI,omitempty"`
	ClrSysMmbId *ClearingSystemMemberIdentification2 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08 ClrSysMmbId,omitempty"`
	Nm          *string                              `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08 Nm,omitempty"`
	PstlAdr     *PostalAddress6                      `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08 PstlAdr,omitempty"`
	Othr        *GenericFinancialIdentification1     `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08 Othr,omitempty"`
}

type BranchData2 struct {
	Id      *string         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08 Id,omitempty"`
	Nm      *string         `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08 Nm,omitempty"`
	PstlAdr *PostalAddress6 `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08 PstlAdr,omitempty"`
}

type PostalAddress6 struct {
	AdrTp       *string   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08 AdrTp,omitempty"`
	Dept        *string   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08 Dept,omitempty"`
	SubDept     *string   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08 SubDept,omitempty"`
	StrtNm      *string   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08 StrtNm,omitempty"`
	BldgNb      *string   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08 BldgNb,omitempty"`
	PstCd       *string   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08 PstCd,omitempty"`
	TwnNm       *string   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08 TwnNm,omitempty"`
	CtrySubDvsn *string   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08 CtrySubDvsn,omitempty"`
	Ctry        *string   `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08 Ctry,omitempty"`
	AdrLine     []*string `xml:"urn:iso:std:iso:20022:tech:xsd:pacs.002.001.08 AdrLine,omitempty"`
}

func (po6 PostalAddress6) ToPostalAddress() *addressbook.PostalAddress {
	res := &addressbook.PostalAddress{}
	if po6.AdrTp != nil {
		res.AddressType = &addressbook.AddressType{
			Code: addressbook.AddressTypeCode(*po6.AdrTp),
		}
	}
	if po6.Dept != nil {
		res.Department = string(*po6.Dept)
	}
	if po6.SubDept != nil {
		res.SubDepartment = string(*po6.SubDept)
	}
	if po6.StrtNm != nil {
		res.StreetName = string(*po6.StrtNm)
	}
	if po6.BldgNb != nil {
		res.BuildingNumber = string(*po6.BldgNb)
	}
	if po6.PstCd != nil {
		res.PostalCode = string(*po6.PstCd)
	}
	if po6.TwnNm != nil {
		res.TownName = string(*po6.TwnNm)
	}
	if po6.CtrySubDvsn != nil {
		res.CountrySubDivision = string(*po6.CtrySubDvsn)
	}
	if po6.Ctry != nil {
		res.CountryCode = string(*po6.Ctry)
	}
	if po6.AdrLine != nil {
		res.AddressLine = make([]string, 0, len(po6.AdrLine))
		for _, line := range po6.AdrLine {
			res.AddressLine = append(
				res.AddressLine,
				string(*line),
			)
		}
	}
	return res
}
