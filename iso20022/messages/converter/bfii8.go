package converter

import (
	"github.com/CoreumFoundation/iso20022-client/iso20022/addressbook"
)

//type BICFIDec2014Identifier string
//type ExternalClearingSystemIdentification1Code string
//type Max35Text string
//type LEIIdentifier string
//type Max140Text string
//type ExternalFinancialInstitutionIdentification1Code string
//type CountryCode string
//type Exact4AlphaNumericText string
//type Max70Text string
//type Max16Text string
//type AddressType2Code string
//
//const AddressType2CodeAddr AddressType2Code = "ADDR"
//const AddressType2CodePbox AddressType2Code = "PBOX"
//const AddressType2CodeHome AddressType2Code = "HOME"
//const AddressType2CodeBizz AddressType2Code = "BIZZ"
//const AddressType2CodeMlto AddressType2Code = "MLTO"
//const AddressType2CodeDlvy AddressType2Code = "DLVY"

type BranchAndFinancialInstitutionIdentification8 struct {
	FinInstnId FinancialInstitutionIdentification23
	BrnchId    *BranchData5
}

func (agent *BranchAndFinancialInstitutionIdentification8) ToParty() *addressbook.Party {
	res := new(addressbook.Party)
	if agent == nil {
		return res
	}

	if agent.BrnchId != nil {
		res.Branch = &addressbook.Branch{}
		if agent.BrnchId.Id != nil {
			res.Branch.Id = *agent.BrnchId.Id
		}
		if agent.BrnchId.LEI != nil {
			res.Branch.LegalEntityIdentifier = *agent.BrnchId.LEI
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
	if agent.FinInstnId.LEI != nil {
		res.Identification.BusinessIdentifiersCode = *agent.FinInstnId.LEI
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

type FinancialInstitutionIdentification23 struct {
	BICFI       *string
	ClrSysMmbId *ClearingSystemMemberIdentification2
	LEI         *string
	Nm          *string
	PstlAdr     *PostalAddress27
	Othr        *GenericFinancialIdentification1
}

type BranchData5 struct {
	Id      *string
	LEI     *string
	Nm      *string
	PstlAdr *PostalAddress27
}

type PostalAddress27 struct {
	AdrTp       *AddressType3Choice
	CareOf      *string
	Dept        *string
	SubDept     *string
	StrtNm      *string
	BldgNb      *string
	BldgNm      *string
	Flr         *string
	UnitNb      *string
	PstBx       *string
	Room        *string
	PstCd       *string
	TwnNm       *string
	TwnLctnNm   *string
	DstrctNm    *string
	CtrySubDvsn *string
	Ctry        *string
	AdrLine     []*string
}

func (po27 PostalAddress27) ToPostalAddress() *addressbook.PostalAddress {
	res := &addressbook.PostalAddress{}
	if po27.AdrTp != nil {
		res.AddressType = &addressbook.AddressType{
			Code: addressbook.AddressTypeCode(*po27.AdrTp.Cd),
			Proprietary: &addressbook.Proprietary{
				Id:         po27.AdrTp.Prtry.Id,
				Issuer:     po27.AdrTp.Prtry.Issr,
				SchemeName: "",
			},
		}
	}
	if po27.AdrTp.Prtry.SchmeNm != nil {
		res.AddressType.Proprietary.SchemeName = *po27.AdrTp.Prtry.SchmeNm
	}
	if po27.Dept != nil {
		res.Department = *po27.Dept
	}
	if po27.SubDept != nil {
		res.SubDepartment = *po27.SubDept
	}
	if po27.StrtNm != nil {
		res.StreetName = *po27.StrtNm
	}
	if po27.BldgNb != nil {
		res.BuildingNumber = *po27.BldgNb
	}
	if po27.BldgNm != nil {
		res.BuildingName = *po27.BldgNm
	}
	if po27.Flr != nil {
		res.Floor = *po27.Flr
	}
	if po27.PstBx != nil {
		res.PostalBox = *po27.PstBx
	}
	if po27.Room != nil {
		res.Room = *po27.Room
	}
	if po27.PstCd != nil {
		res.PostalCode = *po27.PstCd
	}
	if po27.TwnNm != nil {
		res.TownName = *po27.TwnNm
	}
	if po27.TwnLctnNm != nil {
		res.TownLocationName = *po27.TwnLctnNm
	}
	if po27.DstrctNm != nil {
		res.DistrictName = *po27.DstrctNm
	}
	if po27.CtrySubDvsn != nil {
		res.CountrySubDivision = *po27.CtrySubDvsn
	}
	if po27.Ctry != nil {
		res.CountryCode = *po27.Ctry
	}
	if po27.AdrLine != nil {
		res.AddressLine = make([]string, 0, len(po27.AdrLine))
		for _, line := range po27.AdrLine {
			res.AddressLine = append(
				res.AddressLine,
				*line,
			)
		}
	}
	return res
}
