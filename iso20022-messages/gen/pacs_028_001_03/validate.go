// Code generated by GoComply XSD2Go for iso20022; DO NOT EDIT.
// Validations for urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03
package pacs_028_001_03

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

func (v Document) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Document"
	iso.AddError(&errs, baseName+".FIToFIPmtStsReq", v.FIToFIPmtStsReq.Validate())
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

func (v FIToFIPaymentStatusRequestV03) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "FIToFIPaymentStatusRequestV03"
	iso.AddError(&errs, baseName+".GrpHdr", v.GrpHdr.Validate())
	iso.AddError(&errs, baseName+".OrgnlGrpInf", v.OrgnlGrpInf.Validate())
	iso.AddError(&errs, baseName+".TxInf", v.TxInf.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v FIToFIPaymentStatusRequestV03TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "FIToFIPaymentStatusRequestV03TCH"
	iso.AddError(&errs, baseName+".GrpHdr", v.GrpHdr.Validate())
	iso.AddError(&errs, baseName+".OrgnlGrpInf", v.OrgnlGrpInf.Validate())
	iso.AddError(&errs, baseName+".TxInf", v.TxInf.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GroupHeader91) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GroupHeader91"
	iso.AddError(&errs, baseName+".MsgId", v.MsgId.Validate())
	iso.AddError(&errs, baseName+".CreDtTm", v.CreDtTm.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v GroupHeader91TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "GroupHeader91TCH"
	iso.AddError(&errs, baseName+".MsgId", v.MsgId.Validate())
	iso.AddError(&errs, baseName+".CreDtTm", v.CreDtTm.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OriginalGroupInformation27) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OriginalGroupInformation27"
	iso.AddError(&errs, baseName+".OrgnlMsgId", v.OrgnlMsgId.Validate())
	iso.AddError(&errs, baseName+".OrgnlMsgNmId", v.OrgnlMsgNmId.Validate())
	iso.AddError(&errs, baseName+".OrgnlCreDtTm", v.OrgnlCreDtTm.Validate())
	iso.AddError(&errs, baseName+".OrgnlNbOfTxs", v.OrgnlNbOfTxs.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OriginalGroupInformation27TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OriginalGroupInformation27TCH"
	iso.AddError(&errs, baseName+".OrgnlMsgId", v.OrgnlMsgId.Validate())
	iso.AddError(&errs, baseName+".OrgnlMsgNmId", v.OrgnlMsgNmId.Validate())
	iso.AddError(&errs, baseName+".OrgnlCreDtTm", v.OrgnlCreDtTm.Validate())
	iso.AddError(&errs, baseName+".OrgnlNbOfTxs", v.OrgnlNbOfTxs.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OriginalTransactionReference28) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OriginalTransactionReference28"
	iso.AddError(&errs, baseName+".IntrBkSttlmAmt", v.IntrBkSttlmAmt.Validate())
	iso.AddError(&errs, baseName+".IntrBkSttlmDt", v.IntrBkSttlmDt.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v OriginalTransactionReference28TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "OriginalTransactionReference28TCH"
	iso.AddError(&errs, baseName+".IntrBkSttlmAmt", v.IntrBkSttlmAmt.Validate())
	iso.AddError(&errs, baseName+".IntrBkSttlmDt", v.IntrBkSttlmDt.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PaymentTransaction113) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PaymentTransaction113"
	iso.AddError(&errs, baseName+".OrgnlInstrId", v.OrgnlInstrId.Validate())
	iso.AddError(&errs, baseName+".OrgnlTxId", v.OrgnlTxId.Validate())
	iso.AddError(&errs, baseName+".AccptncDtTm", v.AccptncDtTm.Validate())
	iso.AddError(&errs, baseName+".InstgAgt", v.InstgAgt.Validate())
	iso.AddError(&errs, baseName+".InstdAgt", v.InstdAgt.Validate())
	iso.AddError(&errs, baseName+".OrgnlTxRef", v.OrgnlTxRef.Validate())
	if errs.Empty() {
		return nil
	}
	return errs
}

func (v PaymentTransaction113TCH) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "PaymentTransaction113TCH"
	iso.AddError(&errs, baseName+".OrgnlInstrId", v.OrgnlInstrId.Validate())
	iso.AddError(&errs, baseName+".OrgnlTxId", v.OrgnlTxId.Validate())
	iso.AddError(&errs, baseName+".AccptncDtTm", v.AccptncDtTm.Validate())
	iso.AddError(&errs, baseName+".InstgAgt", v.InstgAgt.Validate())
	iso.AddError(&errs, baseName+".InstdAgt", v.InstdAgt.Validate())
	iso.AddError(&errs, baseName+".OrgnlTxRef", v.OrgnlTxRef.Validate())
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

func (v Max1NumericText) Validate() error {
	if err := iso.ValidatePattern(string(v), `[1]{1,1}`); err != nil {
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

func (v OrigMsgName) Validate() error {
	if err := iso.ValidateEnumeration(string(v), "pacs.008.001.06", "pacs.008.001.08", "pacs.009.001.08", "pain.013.001.07"); err != nil {
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
