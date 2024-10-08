// Code generated by GoComply XSD2Go for iso20022; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:pacs.028.001.03 with prefix 's8'
package pacs_028_001_03

import (
	"encoding/xml"

	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/pkg/iso"
)

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchAndFinancialInstitutionIdentification6) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.FinInstnId, xml.StartElement{Name: xml.Name{Local: "s8:FinInstnId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v BranchAndFinancialInstitutionIdentification6TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.FinInstnId, xml.StartElement{Name: xml.Name{Local: "s8:FinInstnId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemMemberIdentification2) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MmbId, xml.StartElement{Name: xml.Name{Local: "s8:MmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v ClearingSystemMemberIdentification2TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MmbId, xml.StartElement{Name: xml.Name{Local: "s8:MmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v Document) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.FIToFIPmtStsReq, xml.StartElement{Name: xml.Name{Local: "s8:FIToFIPmtStsReq"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FinancialInstitutionIdentification18) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.ClrSysMmbId, xml.StartElement{Name: xml.Name{Local: "s8:ClrSysMmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FinancialInstitutionIdentification18TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.ClrSysMmbId, xml.StartElement{Name: xml.Name{Local: "s8:ClrSysMmbId"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FIToFIPaymentStatusRequestV03) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.GrpHdr, xml.StartElement{Name: xml.Name{Local: "s8:GrpHdr"}})
	e.EncodeElement(v.OrgnlGrpInf, xml.StartElement{Name: xml.Name{Local: "s8:OrgnlGrpInf"}})
	e.EncodeElement(v.TxInf, xml.StartElement{Name: xml.Name{Local: "s8:TxInf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v FIToFIPaymentStatusRequestV03TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.GrpHdr, xml.StartElement{Name: xml.Name{Local: "s8:GrpHdr"}})
	e.EncodeElement(v.OrgnlGrpInf, xml.StartElement{Name: xml.Name{Local: "s8:OrgnlGrpInf"}})
	e.EncodeElement(v.TxInf, xml.StartElement{Name: xml.Name{Local: "s8:TxInf"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v GroupHeader91) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MsgId, xml.StartElement{Name: xml.Name{Local: "s8:MsgId"}})
	e.EncodeElement(v.CreDtTm, xml.StartElement{Name: xml.Name{Local: "s8:CreDtTm"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v GroupHeader91TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.MsgId, xml.StartElement{Name: xml.Name{Local: "s8:MsgId"}})
	e.EncodeElement(v.CreDtTm, xml.StartElement{Name: xml.Name{Local: "s8:CreDtTm"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v OriginalGroupInformation27) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.OrgnlMsgId, xml.StartElement{Name: xml.Name{Local: "s8:OrgnlMsgId"}})
	e.EncodeElement(v.OrgnlMsgNmId, xml.StartElement{Name: xml.Name{Local: "s8:OrgnlMsgNmId"}})
	e.EncodeElement(v.OrgnlCreDtTm, xml.StartElement{Name: xml.Name{Local: "s8:OrgnlCreDtTm"}})
	e.EncodeElement(v.OrgnlNbOfTxs, xml.StartElement{Name: xml.Name{Local: "s8:OrgnlNbOfTxs"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v OriginalGroupInformation27TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.OrgnlMsgId, xml.StartElement{Name: xml.Name{Local: "s8:OrgnlMsgId"}})
	e.EncodeElement(v.OrgnlMsgNmId, xml.StartElement{Name: xml.Name{Local: "s8:OrgnlMsgNmId"}})
	e.EncodeElement(v.OrgnlCreDtTm, xml.StartElement{Name: xml.Name{Local: "s8:OrgnlCreDtTm"}})
	e.EncodeElement(v.OrgnlNbOfTxs, xml.StartElement{Name: xml.Name{Local: "s8:OrgnlNbOfTxs"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v OriginalTransactionReference28) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.IntrBkSttlmAmt, xml.StartElement{Name: xml.Name{Local: "s8:IntrBkSttlmAmt"}})
	e.EncodeElement(v.IntrBkSttlmDt, xml.StartElement{Name: xml.Name{Local: "s8:IntrBkSttlmDt"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v OriginalTransactionReference28TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.IntrBkSttlmAmt, xml.StartElement{Name: xml.Name{Local: "s8:IntrBkSttlmAmt"}})
	e.EncodeElement(v.IntrBkSttlmDt, xml.StartElement{Name: xml.Name{Local: "s8:IntrBkSttlmDt"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PaymentTransaction113) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.OrgnlInstrId, xml.StartElement{Name: xml.Name{Local: "s8:OrgnlInstrId"}})
	e.EncodeElement(v.OrgnlTxId, xml.StartElement{Name: xml.Name{Local: "s8:OrgnlTxId"}})
	e.EncodeElement(v.AccptncDtTm, xml.StartElement{Name: xml.Name{Local: "s8:AccptncDtTm"}})
	e.EncodeElement(v.InstgAgt, xml.StartElement{Name: xml.Name{Local: "s8:InstgAgt"}})
	e.EncodeElement(v.InstdAgt, xml.StartElement{Name: xml.Name{Local: "s8:InstdAgt"}})
	e.EncodeElement(v.OrgnlTxRef, xml.StartElement{Name: xml.Name{Local: "s8:OrgnlTxRef"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

// MarshalXML is a custom marshaller that allows us to manipulate the XML tag in order to use the proper namespace prefix
func (v PaymentTransaction113TCH) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(xml.StartElement{Name: xml.Name{Local: start.Name.Local}})
	e.EncodeElement(v.OrgnlInstrId, xml.StartElement{Name: xml.Name{Local: "s8:OrgnlInstrId"}})
	e.EncodeElement(v.OrgnlTxId, xml.StartElement{Name: xml.Name{Local: "s8:OrgnlTxId"}})
	e.EncodeElement(v.AccptncDtTm, xml.StartElement{Name: xml.Name{Local: "s8:AccptncDtTm"}})
	e.EncodeElement(v.InstgAgt, xml.StartElement{Name: xml.Name{Local: "s8:InstgAgt"}})
	e.EncodeElement(v.InstdAgt, xml.StartElement{Name: xml.Name{Local: "s8:InstdAgt"}})
	e.EncodeElement(v.OrgnlTxRef, xml.StartElement{Name: xml.Name{Local: "s8:OrgnlTxRef"}})
	e.EncodeToken(xml.EndElement{Name: xml.Name{Local: start.Name.Local}})
	return nil
}

func (a ActiveOrHistoricCurrencyAndAmountSimpleType) MarshalText() ([]byte, error) {
	return iso.Amount(a).MarshalText()
}
