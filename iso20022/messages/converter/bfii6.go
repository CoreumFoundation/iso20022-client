package converter

import (
	"github.com/CoreumFoundation/iso20022-client/iso20022/addressbook"
)

type BranchAndFinancialInstitutionIdentification6 struct {
	FinInstnId FinancialInstitutionIdentification18
	BrnchId    *BranchData3
}

func (agent *BranchAndFinancialInstitutionIdentification6) ToParty() *addressbook.Party {
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

type FinancialInstitutionIdentification18 struct {
	BICFI       *string
	ClrSysMmbId *ClearingSystemMemberIdentification2
	LEI         *string
	Nm          *string
	PstlAdr     *PostalAddress24
	Othr        *GenericFinancialIdentification1
}

type ClearingSystemMemberIdentification2 struct {
	ClrSysId *ClearingSystemIdentification2Choice
	MmbId    string
}

type ClearingSystemIdentification2Choice struct {
	Cd    *string
	Prtry *string
}

type GenericFinancialIdentification1 struct {
	Id      string
	SchmeNm *FinancialIdentificationSchemeName1Choice
	Issr    *string
}

type FinancialIdentificationSchemeName1Choice struct {
	Cd    *string
	Prtry *string
}

type BranchData3 struct {
	Id      *string
	LEI     *string
	Nm      *string
	PstlAdr *PostalAddress24
}

type PostalAddress24 struct {
	AdrTp       *AddressType3Choice
	Dept        *string
	SubDept     *string
	StrtNm      *string
	BldgNb      *string
	BldgNm      *string
	Flr         *string
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

func (po24 PostalAddress24) ToPostalAddress() *addressbook.PostalAddress {
	res := &addressbook.PostalAddress{}
	if po24.AdrTp != nil {
		res.AddressType = &addressbook.AddressType{
			Code: addressbook.AddressTypeCode(*po24.AdrTp.Cd),
			Proprietary: &addressbook.Proprietary{
				Id:         string(po24.AdrTp.Prtry.Id),
				Issuer:     string(po24.AdrTp.Prtry.Issr),
				SchemeName: "",
			},
		}
	}
	if po24.AdrTp.Prtry.SchmeNm != nil {
		res.AddressType.Proprietary.SchemeName = string(*po24.AdrTp.Prtry.SchmeNm)
	}
	if po24.Dept != nil {
		res.Department = string(*po24.Dept)
	}
	if po24.SubDept != nil {
		res.SubDepartment = string(*po24.SubDept)
	}
	if po24.StrtNm != nil {
		res.StreetName = string(*po24.StrtNm)
	}
	if po24.BldgNb != nil {
		res.BuildingNumber = string(*po24.BldgNb)
	}
	if po24.BldgNm != nil {
		res.BuildingName = string(*po24.BldgNm)
	}
	if po24.Flr != nil {
		res.Floor = string(*po24.Flr)
	}
	if po24.PstBx != nil {
		res.PostalBox = string(*po24.PstBx)
	}
	if po24.Room != nil {
		res.Room = string(*po24.Room)
	}
	if po24.PstCd != nil {
		res.PostalCode = string(*po24.PstCd)
	}
	if po24.TwnNm != nil {
		res.TownName = string(*po24.TwnNm)
	}
	if po24.TwnLctnNm != nil {
		res.TownLocationName = string(*po24.TwnLctnNm)
	}
	if po24.DstrctNm != nil {
		res.DistrictName = string(*po24.DstrctNm)
	}
	if po24.CtrySubDvsn != nil {
		res.CountrySubDivision = string(*po24.CtrySubDvsn)
	}
	if po24.Ctry != nil {
		res.CountryCode = string(*po24.Ctry)
	}
	if po24.AdrLine != nil {
		res.AddressLine = make([]string, 0, len(po24.AdrLine))
		for _, line := range po24.AdrLine {
			res.AddressLine = append(
				res.AddressLine,
				string(*line),
			)
		}
	}
	return res
}

type AddressType3Choice struct {
	Cd    *string
	Prtry *GenericIdentification30
}

type GenericIdentification30 struct {
	Id      string
	Issr    string
	SchmeNm *string
}
