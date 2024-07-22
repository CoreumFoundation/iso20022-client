// Code generated by GoComply XSD2Go for iso20022; DO NOT EDIT.
// Validations for urn:iso
package messages

import (
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/pkg/iso"
	"github.com/moov-io/base"
)

// XSD Element validations

func (v Message) Validate() error {
	var errs base.ErrorList = base.ErrorList{}
	baseName := "Message"
	if v.AppHdr11 != nil {
		iso.AddError(&errs, baseName+".AppHdr11", v.AppHdr11.Validate())
	}
	if v.AppHdr12 != nil {
		iso.AddError(&errs, baseName+".AppHdr12", v.AppHdr12.Validate())
	}
	if v.AppHdr14 != nil {
		iso.AddError(&errs, baseName+".AppHdr14", v.AppHdr14.Validate())
	}
	if v.AppHdr21 != nil {
		iso.AddError(&errs, baseName+".AppHdr21", v.AppHdr21.Validate())
	}
	if v.FIToFICustomerCreditTransferV06 != nil {
		iso.AddError(&errs, baseName+".FIToFICustomerCreditTransferV06", v.FIToFICustomerCreditTransferV06.Validate())
	}
	if v.FIToFICustomerCreditTransferV08 != nil {
		iso.AddError(&errs, baseName+".FIToFICustomerCreditTransferV08", v.FIToFICustomerCreditTransferV08.Validate())
	}
	if v.FIToFICustomerCreditTransferV09 != nil {
		iso.AddError(&errs, baseName+".FIToFICustomerCreditTransferV09", v.FIToFICustomerCreditTransferV09.Validate())
	}
	if v.FIToFICustomerCreditTransferV12 != nil {
		iso.AddError(&errs, baseName+".FIToFICustomerCreditTransferV12", v.FIToFICustomerCreditTransferV12.Validate())
	}
	if v.FIToFIPaymentStatusReportV07 != nil {
		iso.AddError(&errs, baseName+".FIToFIPaymentStatusReportV07", v.FIToFIPaymentStatusReportV07.Validate())
	}
	if v.FIToFIPaymentStatusReportV08 != nil {
		iso.AddError(&errs, baseName+".FIToFIPaymentStatusReportV08", v.FIToFIPaymentStatusReportV08.Validate())
	}
	if v.FIToFIPaymentStatusReportV10 != nil {
		iso.AddError(&errs, baseName+".FIToFIPaymentStatusReportV10", v.FIToFIPaymentStatusReportV10.Validate())
	}
	if v.FIToFIPaymentStatusReportV11 != nil {
		iso.AddError(&errs, baseName+".FIToFIPaymentStatusReportV11", v.FIToFIPaymentStatusReportV11.Validate())
	}
	if v.FIToFIPaymentStatusReportV12 != nil {
		iso.AddError(&errs, baseName+".FIToFIPaymentStatusReportV12", v.FIToFIPaymentStatusReportV12.Validate())
	}
	if v.FIToFIPaymentStatusReportV14 != nil {
		iso.AddError(&errs, baseName+".FIToFIPaymentStatusReportV14", v.FIToFIPaymentStatusReportV14.Validate())
	}
	if v.FIToFICustomerDirectDebitV08 != nil {
		iso.AddError(&errs, baseName+".FIToFICustomerDirectDebitV08", v.FIToFICustomerDirectDebitV08.Validate())
	}
	if v.FIToFICustomerDirectDebitV11 != nil {
		iso.AddError(&errs, baseName+".FIToFICustomerDirectDebitV11", v.FIToFICustomerDirectDebitV11.Validate())
	}
	if v.FinancialInstitutionCreditTransferV08 != nil {
		iso.AddError(&errs, baseName+".FinancialInstitutionCreditTransferV08", v.FinancialInstitutionCreditTransferV08.Validate())
	}
	if v.FinancialInstitutionCreditTransferV09 != nil {
		iso.AddError(&errs, baseName+".FinancialInstitutionCreditTransferV09", v.FinancialInstitutionCreditTransferV09.Validate())
	}
	if v.FinancialInstitutionCreditTransferV11 != nil {
		iso.AddError(&errs, baseName+".FinancialInstitutionCreditTransferV11", v.FinancialInstitutionCreditTransferV11.Validate())
	}
	if v.FinancialInstitutionDirectDebitV04 != nil {
		iso.AddError(&errs, baseName+".FinancialInstitutionDirectDebitV04", v.FinancialInstitutionDirectDebitV04.Validate())
	}
	if v.FinancialInstitutionDirectDebitV06 != nil {
		iso.AddError(&errs, baseName+".FinancialInstitutionDirectDebitV06", v.FinancialInstitutionDirectDebitV06.Validate())
	}
	if v.PaymentReturnV10 != nil {
		iso.AddError(&errs, baseName+".PaymentReturnV10", v.PaymentReturnV10.Validate())
	}
	if v.PaymentReturnV13 != nil {
		iso.AddError(&errs, baseName+".PaymentReturnV13", v.PaymentReturnV13.Validate())
	}
	if v.FIToFIPaymentReversalV10 != nil {
		iso.AddError(&errs, baseName+".FIToFIPaymentReversalV10", v.FIToFIPaymentReversalV10.Validate())
	}
	if v.FIToFIPaymentReversalV13 != nil {
		iso.AddError(&errs, baseName+".FIToFIPaymentReversalV13", v.FIToFIPaymentReversalV13.Validate())
	}
	if v.FIToFIPaymentStatusRequestV03 != nil {
		iso.AddError(&errs, baseName+".FIToFIPaymentStatusRequestV03", v.FIToFIPaymentStatusRequestV03.Validate())
	}
	if v.FIToFIPaymentStatusRequestV04 != nil {
		iso.AddError(&errs, baseName+".FIToFIPaymentStatusRequestV04", v.FIToFIPaymentStatusRequestV04.Validate())
	}
	if v.FIToFIPaymentStatusRequestV06 != nil {
		iso.AddError(&errs, baseName+".FIToFIPaymentStatusRequestV06", v.FIToFIPaymentStatusRequestV06.Validate())
	}
	if v.MultilateralSettlementRequestV02 != nil {
		iso.AddError(&errs, baseName+".MultilateralSettlementRequestV02", v.MultilateralSettlementRequestV02.Validate())
	}
	if errs.Empty() {
		return nil
	}
	return errs
}

// XSD ComplexType validations
