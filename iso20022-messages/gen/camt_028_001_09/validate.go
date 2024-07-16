// Code generated by GoComply XSD2Go for iso20022; DO NOT EDIT.
// Validations for urn:iso:std:iso:20022:tech:xsd:camt.028.001.09
package camt_028_001_09

import (
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/pkg/iso"
	"github.com/moov-io/base"
)

// XSD ComplexType validations

func (v ActiveOrHistoricCurrencyAndAmount) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ActiveOrHistoricCurrencyAndAmount"

	iso.AddError(&errs, baseName+".Ccy", v.Ccy.Validate())

	if errs.Empty() {
		return nil
	}
	return errs
}

func (v AdditionalPaymentInformationV09) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "AdditionalPaymentInformationV09"
	iso.AddError(&errs, baseName+".Assgnmt", v.Assgnmt.Validate())
	iso.AddError(&errs, baseName+".Case", v.Case.Validate())
	iso.AddError(&errs, baseName+".Undrlyg", v.Undrlyg.Validate())
	iso.AddError(&errs, baseName+".Inf", v.Inf.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v AdditionalPaymentInformationV09TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "AdditionalPaymentInformationV09TCH"
	iso.AddError(&errs, baseName+".Assgnmt", v.Assgnmt.Validate())
	iso.AddError(&errs, baseName+".Case", v.Case.Validate())
	iso.AddError(&errs, baseName+".Undrlyg", v.Undrlyg.Validate())
	iso.AddError(&errs, baseName+".Inf", v.Inf.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v BranchAndFinancialInstitutionIdentification6) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "BranchAndFinancialInstitutionIdentification6"
	iso.AddError(&errs, baseName+".FinInstnId", v.FinInstnId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v BranchAndFinancialInstitutionIdentification6TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "BranchAndFinancialInstitutionIdentification6TCH"
	iso.AddError(&errs, baseName+".FinInstnId", v.FinInstnId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Case5) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Case5"
	iso.AddError(&errs, baseName+".Id", v.Id.Validate())
	iso.AddError(&errs, baseName+".Cretr", v.Cretr.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Case5TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Case5TCH"
	iso.AddError(&errs, baseName+".Id", v.Id.Validate())
	iso.AddError(&errs, baseName+".Cretr", v.Cretr.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v CaseAssignment5) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "CaseAssignment5"
	iso.AddError(&errs, baseName+".Id", v.Id.Validate())
	iso.AddError(&errs, baseName+".Assgnr", v.Assgnr.Validate())
	iso.AddError(&errs, baseName+".Assgne", v.Assgne.Validate())
	iso.AddError(&errs, baseName+".CreDtTm", v.CreDtTm.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v CaseAssignment5TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "CaseAssignment5TCH"
	iso.AddError(&errs, baseName+".Id", v.Id.Validate())
	iso.AddError(&errs, baseName+".Assgnr", v.Assgnr.Validate())
	iso.AddError(&errs, baseName+".Assgne", v.Assgne.Validate())
	iso.AddError(&errs, baseName+".CreDtTm", v.CreDtTm.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ClearingSystemMemberIdentification2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ClearingSystemMemberIdentification2"
	iso.AddError(&errs, baseName+".MmbId", v.MmbId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v ClearingSystemMemberIdentification2TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "ClearingSystemMemberIdentification2TCH"
	iso.AddError(&errs, baseName+".MmbId", v.MmbId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Contact4) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Contact4"
	if v.PhneNb != nil {
		iso.AddError(&errs, baseName+".PhneNb", v.PhneNb.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v DateAndDateTime2Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "DateAndDateTime2Choice"
	if v.Dt != nil {
		iso.AddError(&errs, baseName+".Dt", v.Dt.Validate())
	}
	if v.DtTm != nil {
		iso.AddError(&errs, baseName+".DtTm", v.DtTm.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v DateAndPlaceOfBirth1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "DateAndPlaceOfBirth1"
	iso.AddError(&errs, baseName+".BirthDt", v.BirthDt.Validate())
	iso.AddError(&errs, baseName+".CityOfBirth", v.CityOfBirth.Validate())
	iso.AddError(&errs, baseName+".CtryOfBirth", v.CtryOfBirth.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v DiscountAmountAndType1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "DiscountAmountAndType1"
	iso.AddError(&errs, baseName+".Tp", v.Tp.Validate())
	iso.AddError(&errs, baseName+".Amt", v.Amt.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v DiscountAmountAndType1TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "DiscountAmountAndType1TCH"
	iso.AddError(&errs, baseName+".Tp", v.Tp.Validate())
	iso.AddError(&errs, baseName+".Amt", v.Amt.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v DiscountAmountType1Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "DiscountAmountType1Choice"
	if v.Prtry != nil {
		iso.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v DiscountAmountType1ChoiceTCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "DiscountAmountType1ChoiceTCH"
	if v.Prtry != nil {
		iso.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Document) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Document"
	iso.AddError(&errs, baseName+".AddtlPmtInf", v.AddtlPmtInf.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v FinancialInstitutionIdentification18) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "FinancialInstitutionIdentification18"
	iso.AddError(&errs, baseName+".ClrSysMmbId", v.ClrSysMmbId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v FinancialInstitutionIdentification18TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "FinancialInstitutionIdentification18TCH"
	iso.AddError(&errs, baseName+".ClrSysMmbId", v.ClrSysMmbId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GenericPersonIdentification1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GenericPersonIdentification1"
	iso.AddError(&errs, baseName+".Id", v.Id.Validate())
	if v.SchmeNm != nil {
		iso.AddError(&errs, baseName+".SchmeNm", v.SchmeNm.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OrganisationIdentification29) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OrganisationIdentification29"
	if v.LEI != nil {
		iso.AddError(&errs, baseName+".LEI", v.LEI.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OrganisationIdentification29TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OrganisationIdentification29TCH"
	iso.AddError(&errs, baseName+".LEI", v.LEI.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OrganisationIdentification30) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OrganisationIdentification30"
	if v.LEI != nil {
		iso.AddError(&errs, baseName+".LEI", v.LEI.Validate())
	}
	if v.Othr != nil {
		for indx := range v.Othr {
			iso.AddError(&errs, baseName+".Othr", v.Othr[indx].Validate())
		}
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OrganisationIdentification30TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OrganisationIdentification30TCH"
	if v.LEI != nil {
		iso.AddError(&errs, baseName+".LEI", v.LEI.Validate())
	}
	if v.Othr != nil {
		for indx := range v.Othr {
			iso.AddError(&errs, baseName+".Othr", v.Othr[indx].Validate())
		}
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GenericOrganisationIdentification1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GenericOrganisationIdentification1"
	iso.AddError(&errs, baseName+".Id", v.Id.Validate())
	if v.SchmeNm != nil {
		iso.AddError(&errs, baseName+".SchmeNm", v.SchmeNm.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GenericOrganisationIdentification1TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GenericOrganisationIdentification1TCH"
	iso.AddError(&errs, baseName+".Id", v.Id.Validate())
	if v.SchmeNm != nil {
		iso.AddError(&errs, baseName+".SchmeNm", v.SchmeNm.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OrganisationIdentificationSchemeName1Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OrganisationIdentificationSchemeName1Choice"
	if v.Prtry != nil {
		iso.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OrganisationIdentificationSchemeName1ChoiceTCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OrganisationIdentificationSchemeName1ChoiceTCH"
	if v.Prtry != nil {
		iso.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PersonIdentificationSchemeName1Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PersonIdentificationSchemeName1Choice"
	if v.Prtry != nil {
		iso.AddError(&errs, baseName+".Prtry", v.Prtry.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Party38Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Party38Choice"
	if v.OrgId != nil {
		iso.AddError(&errs, baseName+".OrgId", v.OrgId.Validate())
	}
	if v.PrvtId != nil {
		iso.AddError(&errs, baseName+".PrvtId", v.PrvtId.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Party38ChoiceTCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Party38ChoiceTCH"
	if v.OrgId != nil {
		iso.AddError(&errs, baseName+".OrgId", v.OrgId.Validate())
	}
	if v.PrvtId != nil {
		iso.AddError(&errs, baseName+".PrvtId", v.PrvtId.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Party38ChoiceTCH2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Party38ChoiceTCH2"
	if v.OrgId != nil {
		iso.AddError(&errs, baseName+".OrgId", v.OrgId.Validate())
	}
	if v.PrvtId != nil {
		iso.AddError(&errs, baseName+".PrvtId", v.PrvtId.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Party38ChoiceTCH3) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Party38ChoiceTCH3"
	if v.OrgId != nil {
		iso.AddError(&errs, baseName+".OrgId", v.OrgId.Validate())
	}
	if v.PrvtId != nil {
		iso.AddError(&errs, baseName+".PrvtId", v.PrvtId.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Party38ChoiceTCH4) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Party38ChoiceTCH4"
	if v.OrgId != nil {
		iso.AddError(&errs, baseName+".OrgId", v.OrgId.Validate())
	}
	if v.PrvtId != nil {
		iso.AddError(&errs, baseName+".PrvtId", v.PrvtId.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Party40Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Party40Choice"
	if v.Pty != nil {
		iso.AddError(&errs, baseName+".Pty", v.Pty.Validate())
	}
	if v.Agt != nil {
		iso.AddError(&errs, baseName+".Agt", v.Agt.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Party40ChoiceTCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Party40ChoiceTCH"
	if v.Agt != nil {
		iso.AddError(&errs, baseName+".Agt", v.Agt.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v Party40ChoiceTCH2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Party40ChoiceTCH2"
	if v.Pty != nil {
		iso.AddError(&errs, baseName+".Pty", v.Pty.Validate())
	}
	if v.Agt != nil {
		iso.AddError(&errs, baseName+".Agt", v.Agt.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PartyIdentification135) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PartyIdentification135"
	if v.Nm != nil {
		iso.AddError(&errs, baseName+".Nm", v.Nm.Validate())
	}
	if v.PstlAdr != nil {
		iso.AddError(&errs, baseName+".PstlAdr", v.PstlAdr.Validate())
	}
	if v.Id != nil {
		iso.AddError(&errs, baseName+".Id", v.Id.Validate())
	}
	if v.CtctDtls != nil {
		iso.AddError(&errs, baseName+".CtctDtls", v.CtctDtls.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PartyIdentification135TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PartyIdentification135TCH"
	iso.AddError(&errs, baseName+".Nm", v.Nm.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PartyIdentification135TCH2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PartyIdentification135TCH2"
	if v.Nm != nil {
		iso.AddError(&errs, baseName+".Nm", v.Nm.Validate())
	}
	if v.PstlAdr != nil {
		iso.AddError(&errs, baseName+".PstlAdr", v.PstlAdr.Validate())
	}
	if v.Id != nil {
		iso.AddError(&errs, baseName+".Id", v.Id.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PartyIdentification135TCH3) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PartyIdentification135TCH3"
	if v.Nm != nil {
		iso.AddError(&errs, baseName+".Nm", v.Nm.Validate())
	}
	if v.PstlAdr != nil {
		iso.AddError(&errs, baseName+".PstlAdr", v.PstlAdr.Validate())
	}
	if v.Id != nil {
		iso.AddError(&errs, baseName+".Id", v.Id.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PartyIdentification135TCH4) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PartyIdentification135TCH4"
	if v.Nm != nil {
		iso.AddError(&errs, baseName+".Nm", v.Nm.Validate())
	}
	if v.PstlAdr != nil {
		iso.AddError(&errs, baseName+".PstlAdr", v.PstlAdr.Validate())
	}
	if v.Id != nil {
		iso.AddError(&errs, baseName+".Id", v.Id.Validate())
	}
	if v.CtctDtls != nil {
		iso.AddError(&errs, baseName+".CtctDtls", v.CtctDtls.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PartyIdentification135TCH5) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PartyIdentification135TCH5"
	if v.Nm != nil {
		iso.AddError(&errs, baseName+".Nm", v.Nm.Validate())
	}
	if v.PstlAdr != nil {
		iso.AddError(&errs, baseName+".PstlAdr", v.PstlAdr.Validate())
	}
	if v.Id != nil {
		iso.AddError(&errs, baseName+".Id", v.Id.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PaymentComplementaryInformation8) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PaymentComplementaryInformation8"
	if v.InstrId != nil {
		iso.AddError(&errs, baseName+".InstrId", v.InstrId.Validate())
	}
	if v.EndToEndId != nil {
		iso.AddError(&errs, baseName+".EndToEndId", v.EndToEndId.Validate())
	}
	if v.TxId != nil {
		iso.AddError(&errs, baseName+".TxId", v.TxId.Validate())
	}
	if v.UltmtDbtr != nil {
		iso.AddError(&errs, baseName+".UltmtDbtr", v.UltmtDbtr.Validate())
	}
	if v.Dbtr != nil {
		iso.AddError(&errs, baseName+".Dbtr", v.Dbtr.Validate())
	}
	if v.Cdtr != nil {
		iso.AddError(&errs, baseName+".Cdtr", v.Cdtr.Validate())
	}
	if v.UltmtCdtr != nil {
		iso.AddError(&errs, baseName+".UltmtCdtr", v.UltmtCdtr.Validate())
	}
	if v.RmtInf != nil {
		iso.AddError(&errs, baseName+".RmtInf", v.RmtInf.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PaymentComplementaryInformation8TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PaymentComplementaryInformation8TCH"
	if v.InstrId != nil {
		iso.AddError(&errs, baseName+".InstrId", v.InstrId.Validate())
	}
	if v.EndToEndId != nil {
		iso.AddError(&errs, baseName+".EndToEndId", v.EndToEndId.Validate())
	}
	if v.TxId != nil {
		iso.AddError(&errs, baseName+".TxId", v.TxId.Validate())
	}
	if v.UltmtDbtr != nil {
		iso.AddError(&errs, baseName+".UltmtDbtr", v.UltmtDbtr.Validate())
	}
	if v.Dbtr != nil {
		iso.AddError(&errs, baseName+".Dbtr", v.Dbtr.Validate())
	}
	if v.Cdtr != nil {
		iso.AddError(&errs, baseName+".Cdtr", v.Cdtr.Validate())
	}
	if v.UltmtCdtr != nil {
		iso.AddError(&errs, baseName+".UltmtCdtr", v.UltmtCdtr.Validate())
	}
	if v.RmtInf != nil {
		iso.AddError(&errs, baseName+".RmtInf", v.RmtInf.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PersonIdentification13) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PersonIdentification13"
	if v.DtAndPlcOfBirth != nil {
		iso.AddError(&errs, baseName+".DtAndPlcOfBirth", v.DtAndPlcOfBirth.Validate())
	}
	if v.Othr != nil {
		for indx := range v.Othr {
			iso.AddError(&errs, baseName+".Othr", v.Othr[indx].Validate())
		}
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PersonIdentification13TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PersonIdentification13TCH"
	if v.DtAndPlcOfBirth != nil {
		iso.AddError(&errs, baseName+".DtAndPlcOfBirth", v.DtAndPlcOfBirth.Validate())
	}
	if v.Othr != nil {
		for indx := range v.Othr {
			iso.AddError(&errs, baseName+".Othr", v.Othr[indx].Validate())
		}
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PersonIdentification13TCH2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PersonIdentification13TCH2"
	iso.AddError(&errs, baseName+".DtAndPlcOfBirth", v.DtAndPlcOfBirth.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PersonIdentification13TCH3) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PersonIdentification13TCH3"
	iso.AddError(&errs, baseName+".Othr", v.Othr.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PostalAddress24) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PostalAddress24"
	iso.AddError(&errs, baseName+".StrtNm", v.StrtNm.Validate())
	if v.BldgNb != nil {
		iso.AddError(&errs, baseName+".BldgNb", v.BldgNb.Validate())
	}
	iso.AddError(&errs, baseName+".PstCd", v.PstCd.Validate())
	iso.AddError(&errs, baseName+".TwnNm", v.TwnNm.Validate())
	iso.AddError(&errs, baseName+".CtrySubDvsn", v.CtrySubDvsn.Validate())
	iso.AddError(&errs, baseName+".Ctry", v.Ctry.Validate())
	if v.AdrLine != nil {
		iso.AddError(&errs, baseName+".AdrLine", v.AdrLine.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PostalAddress24TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PostalAddress24TCH"
	iso.AddError(&errs, baseName+".StrtNm", v.StrtNm.Validate())
	if v.BldgNb != nil {
		iso.AddError(&errs, baseName+".BldgNb", v.BldgNb.Validate())
	}
	iso.AddError(&errs, baseName+".PstCd", v.PstCd.Validate())
	iso.AddError(&errs, baseName+".TwnNm", v.TwnNm.Validate())
	iso.AddError(&errs, baseName+".CtrySubDvsn", v.CtrySubDvsn.Validate())
	iso.AddError(&errs, baseName+".Ctry", v.Ctry.Validate())
	if v.AdrLine != nil {
		iso.AddError(&errs, baseName+".AdrLine", v.AdrLine.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v RemittanceAmount2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "RemittanceAmount2"
	if v.DscntApldAmt != nil {
		for indx := range v.DscntApldAmt {
			iso.AddError(&errs, baseName+".DscntApldAmt", v.DscntApldAmt[indx].Validate())
		}
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v RemittanceAmount2TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "RemittanceAmount2TCH"
	if v.DscntApldAmt != nil {
		for indx := range v.DscntApldAmt {
			iso.AddError(&errs, baseName+".DscntApldAmt", v.DscntApldAmt[indx].Validate())
		}
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v RemittanceInformation16) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "RemittanceInformation16"
	if v.Ustrd != nil {
		for indx := range v.Ustrd {
			iso.AddError(&errs, baseName+".Ustrd", v.Ustrd[indx].Validate())
		}
	}
	if v.Strd != nil {
		iso.AddError(&errs, baseName+".Strd", v.Strd.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v RemittanceInformation16TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "RemittanceInformation16TCH"
	if v.Ustrd != nil {
		for indx := range v.Ustrd {
			iso.AddError(&errs, baseName+".Ustrd", v.Ustrd[indx].Validate())
		}
	}
	if v.Strd != nil {
		iso.AddError(&errs, baseName+".Strd", v.Strd.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v StructuredRemittanceInformation16) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "StructuredRemittanceInformation16"
	if v.RfrdDocAmt != nil {
		iso.AddError(&errs, baseName+".RfrdDocAmt", v.RfrdDocAmt.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v StructuredRemittanceInformation16TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "StructuredRemittanceInformation16TCH"
	if v.RfrdDocAmt != nil {
		iso.AddError(&errs, baseName+".RfrdDocAmt", v.RfrdDocAmt.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v UnderlyingGroupInformation1) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "UnderlyingGroupInformation1"
	iso.AddError(&errs, baseName+".OrgnlMsgId", v.OrgnlMsgId.Validate())
	iso.AddError(&errs, baseName+".OrgnlMsgNmId", v.OrgnlMsgNmId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v UnderlyingGroupInformation1TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "UnderlyingGroupInformation1TCH"
	iso.AddError(&errs, baseName+".OrgnlMsgId", v.OrgnlMsgId.Validate())
	iso.AddError(&errs, baseName+".OrgnlMsgNmId", v.OrgnlMsgNmId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v UnderlyingGroupInformation1TCH2) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "UnderlyingGroupInformation1TCH2"
	iso.AddError(&errs, baseName+".OrgnlMsgId", v.OrgnlMsgId.Validate())
	iso.AddError(&errs, baseName+".OrgnlMsgNmId", v.OrgnlMsgNmId.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v UnderlyingPaymentInstruction5) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "UnderlyingPaymentInstruction5"
	iso.AddError(&errs, baseName+".OrgnlGrpInf", v.OrgnlGrpInf.Validate())
	iso.AddError(&errs, baseName+".OrgnlPmtInfId", v.OrgnlPmtInfId.Validate())
	if v.OrgnlEndToEndId != nil {
		iso.AddError(&errs, baseName+".OrgnlEndToEndId", v.OrgnlEndToEndId.Validate())
	}
	if v.OrgnlUETR != nil {
		iso.AddError(&errs, baseName+".OrgnlUETR", v.OrgnlUETR.Validate())
	}
	iso.AddError(&errs, baseName+".OrgnlInstdAmt", v.OrgnlInstdAmt.Validate())
	if v.ReqdExctnDt != nil {
		iso.AddError(&errs, baseName+".ReqdExctnDt", v.ReqdExctnDt.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v UnderlyingPaymentInstruction5TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "UnderlyingPaymentInstruction5TCH"
	iso.AddError(&errs, baseName+".OrgnlGrpInf", v.OrgnlGrpInf.Validate())
	iso.AddError(&errs, baseName+".OrgnlPmtInfId", v.OrgnlPmtInfId.Validate())
	if v.OrgnlEndToEndId != nil {
		iso.AddError(&errs, baseName+".OrgnlEndToEndId", v.OrgnlEndToEndId.Validate())
	}
	if v.OrgnlUETR != nil {
		iso.AddError(&errs, baseName+".OrgnlUETR", v.OrgnlUETR.Validate())
	}
	iso.AddError(&errs, baseName+".OrgnlInstdAmt", v.OrgnlInstdAmt.Validate())
	if v.ReqdExctnDt != nil {
		iso.AddError(&errs, baseName+".ReqdExctnDt", v.ReqdExctnDt.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v UnderlyingPaymentTransaction4) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "UnderlyingPaymentTransaction4"
	iso.AddError(&errs, baseName+".OrgnlGrpInf", v.OrgnlGrpInf.Validate())
	iso.AddError(&errs, baseName+".OrgnlInstrId", v.OrgnlInstrId.Validate())
	if v.OrgnlEndToEndId != nil {
		iso.AddError(&errs, baseName+".OrgnlEndToEndId", v.OrgnlEndToEndId.Validate())
	}
	iso.AddError(&errs, baseName+".OrgnlTxId", v.OrgnlTxId.Validate())
	if v.OrgnlUETR != nil {
		iso.AddError(&errs, baseName+".OrgnlUETR", v.OrgnlUETR.Validate())
	}
	iso.AddError(&errs, baseName+".OrgnlIntrBkSttlmAmt", v.OrgnlIntrBkSttlmAmt.Validate())
	iso.AddError(&errs, baseName+".OrgnlIntrBkSttlmDt", v.OrgnlIntrBkSttlmDt.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v UnderlyingPaymentTransaction4TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "UnderlyingPaymentTransaction4TCH"
	iso.AddError(&errs, baseName+".OrgnlGrpInf", v.OrgnlGrpInf.Validate())
	iso.AddError(&errs, baseName+".OrgnlInstrId", v.OrgnlInstrId.Validate())
	if v.OrgnlEndToEndId != nil {
		iso.AddError(&errs, baseName+".OrgnlEndToEndId", v.OrgnlEndToEndId.Validate())
	}
	iso.AddError(&errs, baseName+".OrgnlTxId", v.OrgnlTxId.Validate())
	if v.OrgnlUETR != nil {
		iso.AddError(&errs, baseName+".OrgnlUETR", v.OrgnlUETR.Validate())
	}
	iso.AddError(&errs, baseName+".OrgnlIntrBkSttlmAmt", v.OrgnlIntrBkSttlmAmt.Validate())
	iso.AddError(&errs, baseName+".OrgnlIntrBkSttlmDt", v.OrgnlIntrBkSttlmDt.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v UnderlyingTransaction5Choice) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "UnderlyingTransaction5Choice"
	if v.Initn != nil {
		iso.AddError(&errs, baseName+".Initn", v.Initn.Validate())
	}
	if v.IntrBk != nil {
		iso.AddError(&errs, baseName+".IntrBk", v.IntrBk.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v UnderlyingTransaction5ChoiceTCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "UnderlyingTransaction5ChoiceTCH"
	if v.Initn != nil {
		iso.AddError(&errs, baseName+".Initn", v.Initn.Validate())
	}
	if v.IntrBk != nil {
		iso.AddError(&errs, baseName+".IntrBk", v.IntrBk.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

// XSD SimpleType validations

func (v ActiveOrHistoricCurrencyCode) Validate() error {
	if err := iso.ValidatePattern(string(v), `[A-Z]{3,3}`); err != nil {
		return err
	}
	if err := iso.ValidateEnumeration(string(v), "USD"); err != nil {
		return err
	}
	return nil
}

func (v CountryCode) Validate() error {
	if err := iso.ValidatePattern(string(v), `[A-Z]{2,2}`); err != nil {
		return err
	}
	if err := iso.ValidateEnumeration(string(v), "AD", "AE", "AF", "AG", "AI", "AL", "AM", "AO", "AQ", "AR", "AS", "AT", "AU", "AW", "AX", "AZ", "BA", "BB", "BD", "BE", "BF", "BG", "BH", "BI", "BJ", "BL", "BM", "BN", "BO", "BQ", "BR", "BS", "BT", "BV", "BW", "BY", "BZ", "CA", "CC", "CD", "CF", "CG", "CH", "CI", "CK", "CL", "CM", "CN", "CO", "CR", "CU", "CV", "CW", "CX", "CY", "CZ", "DE", "DJ", "DK", "DM", "DO", "DZ", "EC", "EE", "EG", "EH", "ER", "ES", "ET", "FI", "FJ", "FK", "FM", "FO", "FR", "GA", "GB", "GD", "GE", "GF", "GG", "GH", "GI", "GL", "GM", "GN", "GP", "GQ", "GR", "GS", "GT", "GU", "GW", "GY", "HK", "HM", "HN", "HR", "HT", "HU", "ID", "IE", "IL", "IM", "IN", "IO", "IQ", "IR", "IS", "IT", "JE", "JM", "JO", "JP", "KE", "KG", "KH", "KI", "KM", "KN", "KP", "KR", "KW", "KY", "KZ", "LA", "LB", "LC", "LI", "LK", "LR", "LS", "LT", "LU", "LV", "LY", "MA", "MC", "MD", "ME", "MF", "MG", "MH", "MK", "ML", "MM", "MN", "MO", "MP", "MQ", "MR", "MS", "MT", "MU", "MV", "MW", "MX", "MY", "MZ", "NA", "NC", "NE", "NF", "NG", "NI", "NL", "NO", "NP", "NR", "NU", "NZ", "OM", "PA", "PE", "PF", "PG", "PH", "PK", "PL", "PM", "PN", "PR", "PS", "PT", "PW", "PY", "QA", "RE", "RO", "RS", "RU", "RW", "SA", "SB", "SC", "SD", "SE", "SG", "SH", "SI", "SJ", "SK", "SL", "SM", "SN", "SO", "SR", "SS", "ST", "SV", "SX", "SY", "SZ", "TC", "TD", "TF", "TG", "TH", "TJ", "TK", "TL", "TM", "TN", "TO", "TR", "TT", "TV", "TW", "TZ", "UA", "UG", "UM", "US", "UY", "UZ", "VA", "VC", "VE", "VG", "VI", "VN", "VU", "WF", "WS", "YE", "YT", "ZA", "ZM", "ZW"); err != nil {
		return err
	}
	return nil
}

func (v LEIIdentifier) Validate() error {
	if err := iso.ValidatePattern(string(v), `[A-Z0-9]{18,18}[0-9]{2,2}`); err != nil {
		return err
	}
	return nil
}

func (v Max140Text) Validate() error {
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 140); err != nil {
		return err
	}
	return nil
}

func (v Max16Text) Validate() error {
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 16); err != nil {
		return err
	}
	return nil
}

func (v Max35Text) Validate() error {
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 35); err != nil {
		return err
	}
	return nil
}

func (v Max35TextTCH) Validate() error {
	if err := iso.ValidatePattern(string(v), `M[0-9]{4}(((01|03|05|07|08|10|12)((0[1-9])|([1-2][0-9])|(3[0-1])))|((04|06|09|11)((0[1-9])|([1-2][0-9])|30))|((02)((0[1-9])|([1-2][0-9]))))[A-Z0-9]{11}.*`); err != nil {
		return err
	}
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 35); err != nil {
		return err
	}
	return nil
}

func (v Max35TextTCH2) Validate() error {
	if err := iso.ValidateMinLength(string(v), 9); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 9); err != nil {
		return err
	}
	return nil
}

func (v Max35TextTCH3) Validate() error {
	if err := iso.ValidateEnumeration(string(v), "DSCT", "FULL", "ORIG"); err != nil {
		return err
	}
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 4); err != nil {
		return err
	}
	return nil
}

func (v Max70Text) Validate() error {
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 70); err != nil {
		return err
	}
	return nil
}

func (v OrigMsgName) Validate() error {
	if err := iso.ValidateEnumeration(string(v), "pacs.008.001.06", "pacs.008.001.08", "pain.013.001.07"); err != nil {
		return err
	}
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 35); err != nil {
		return err
	}
	return nil
}

func (v OrigMsgNameTCH) Validate() error {
	if err := iso.ValidateEnumeration(string(v), "pain.013.001.07"); err != nil {
		return err
	}
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 35); err != nil {
		return err
	}
	return nil
}

func (v OrigMsgNameTCH2) Validate() error {
	if err := iso.ValidateEnumeration(string(v), "pacs.008.001.06", "pacs.008.001.08"); err != nil {
		return err
	}
	if err := iso.ValidateMinLength(string(v), 1); err != nil {
		return err
	}
	if err := iso.ValidateMaxLength(string(v), 35); err != nil {
		return err
	}
	return nil
}

func (v PhoneNumber) Validate() error {
	if err := iso.ValidatePattern(string(v), `\+[0-9]{1,3}-[0-9()+\-]{1,30}`); err != nil {
		return err
	}
	return nil
}

func (v UUIDv4Identifier) Validate() error {
	if err := iso.ValidatePattern(string(v), `[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89ab][a-f0-9]{3}-[a-f0-9]{12}`); err != nil {
		return err
	}
	return nil
}
