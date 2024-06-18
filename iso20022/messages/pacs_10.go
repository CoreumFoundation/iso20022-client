package messages

import (
	"github.com/moov-io/iso20022/pkg/pacs_v10"

	"github.com/CoreumFoundation/iso20022-client/iso20022/addressbook"
)

func extractPartyFromPacsV10BranchAndFinancialInstitutionIdentification6(agent *pacs_v10.BranchAndFinancialInstitutionIdentification6, res *addressbook.Party) {
	if agent == nil {
		return
	}

	if agent.BrnchId != nil {
		res.Branch = &addressbook.Branch{}
		if agent.BrnchId.Id != nil {
			res.Branch.Id = string(*agent.BrnchId.Id)
		}
		if agent.BrnchId.LEI != nil {
			res.Branch.LegalEntityIdentifier = string(*agent.BrnchId.LEI)
		}
		if agent.BrnchId.Nm != nil {
			res.Branch.Name = string(*agent.BrnchId.Nm)
		}
		res.Branch.PostalAddress = postalAddressFromPacsV10PostalAddress24(agent.BrnchId.PstlAdr)
	}

	if agent.FinInstnId.BICFI != nil {
		res.Identification.BusinessIdentifiersCode = string(*agent.FinInstnId.BICFI)
	}

	if agent.FinInstnId.BICFI != nil {
		res.Identification.BusinessIdentifiersCode = string(*agent.FinInstnId.BICFI)
	}
	if agent.FinInstnId.LEI != nil {
		res.Identification.BusinessIdentifiersCode = string(*agent.FinInstnId.LEI)
	}
	if agent.FinInstnId.ClrSysMmbId != nil {
		res.Identification.ClearingSystemMemberIdentification = &addressbook.ClearingSystemMemberIdentification{
			MemberId: string(agent.FinInstnId.ClrSysMmbId.MmbId),
		}
		if res.Identification.ClearingSystemMemberIdentification.ClearingSystemId != nil {
			res.Identification.ClearingSystemMemberIdentification.ClearingSystemId = &addressbook.ClearingSystemId{
				Code:        string(agent.FinInstnId.ClrSysMmbId.ClrSysId.Cd),
				Proprietary: string(agent.FinInstnId.ClrSysMmbId.ClrSysId.Prtry),
			}
		}
	}
	if agent.FinInstnId.Nm != nil && agent.FinInstnId.PstlAdr != nil {
		res.Identification.Name = string(*agent.FinInstnId.Nm)
		res.Identification.PostalAddress = postalAddressFromPacsV10PostalAddress24(agent.FinInstnId.PstlAdr)
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
				Code:        string(agent.FinInstnId.Othr.SchmeNm.Cd),
				Proprietary: string(agent.FinInstnId.Othr.SchmeNm.Prtry),
			}
		}
	}
}

func postalAddressFromPacsV10PostalAddress24(po24 *pacs_v10.PostalAddress24) (res *addressbook.PostalAddress) {
	if po24 == nil {
		return nil
	}

	res = &addressbook.PostalAddress{}
	if po24.AdrTp != nil {
		res.AddressType = &addressbook.AddressType{
			Code: addressbook.AddressTypeCode(po24.AdrTp.Cd),
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
				string(line),
			)
		}
	}
	return res
}
