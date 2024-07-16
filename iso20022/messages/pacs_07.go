package messages

import (
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_002_001_07"
	"github.com/CoreumFoundation/iso20022-client/iso20022/addressbook"
)

func extractPartyFromPacs00200107BranchAndFinancialInstitutionIdentification5(agent *pacs_002_001_07.BranchAndFinancialInstitutionIdentification5, res *addressbook.Party) {
	if agent == nil {
		return
	}

	if agent.BrnchId != nil {
		res.Branch = &addressbook.Branch{}
		if agent.BrnchId.Id != nil {
			res.Branch.Id = string(*agent.BrnchId.Id)
		}
		if agent.BrnchId.Nm != nil {
			res.Branch.Name = string(*agent.BrnchId.Nm)
		}
		res.Branch.PostalAddress = postalAddressFromPacs00200107PostalAddress6(agent.BrnchId.PstlAdr)
	}

	if agent.FinInstnId.BICFI != nil {
		res.Identification.BusinessIdentifiersCode = string(*agent.FinInstnId.BICFI)
	}
	if agent.FinInstnId.ClrSysMmbId != nil {
		res.Identification.ClearingSystemMemberIdentification = &addressbook.ClearingSystemMemberIdentification{
			MemberId: string(agent.FinInstnId.ClrSysMmbId.MmbId),
		}
		if res.Identification.ClearingSystemMemberIdentification.ClearingSystemId != nil {
			res.Identification.ClearingSystemMemberIdentification.ClearingSystemId = &addressbook.ClearingSystemId{
				Code:        string(*agent.FinInstnId.ClrSysMmbId.ClrSysId.Cd),
				Proprietary: string(*agent.FinInstnId.ClrSysMmbId.ClrSysId.Prtry),
			}
		}
	}
	if agent.FinInstnId.Nm != nil && agent.FinInstnId.PstlAdr != nil {
		res.Identification.Name = string(*agent.FinInstnId.Nm)
		res.Identification.PostalAddress = postalAddressFromPacs00200107PostalAddress6(agent.FinInstnId.PstlAdr)
	}
	if agent.FinInstnId.Othr != nil {
		res.Identification.Other = &addressbook.Other{
			Id: string(agent.FinInstnId.Othr.Id),
		}
		if agent.FinInstnId.Othr.Issr != nil {
			res.Identification.Other.Issuer = string(*agent.FinInstnId.Othr.Issr)
		}
		if agent.FinInstnId.Othr.SchmeNm != nil {
			res.Identification.Other.SchemeName = &addressbook.SchemeName{
				Code:        string(*agent.FinInstnId.Othr.SchmeNm.Cd),
				Proprietary: string(*agent.FinInstnId.Othr.SchmeNm.Prtry),
			}
		}
	}
}

func postalAddressFromPacs00200107PostalAddress6(po6 *pacs_002_001_07.PostalAddress6) (res *addressbook.PostalAddress) {
	if po6 == nil {
		return nil
	}

	res = &addressbook.PostalAddress{}
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
