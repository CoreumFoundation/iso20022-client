// Code generated by GoComply XSD2Go for iso20022; DO NOT EDIT.
// Models for urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 with prefix 'tu'
package admn_008_001_01

import (
	"encoding/xml"

	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/pkg/iso"
)

// XSD ComplexType declarations

type AvailabilityParticipant struct {
	PtcptSgnOff *ParticipantSignOff   `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 PtcptSgnOff,omitempty"`
	PtcptSspd   *ParticipantSuspended `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 PtcptSspd,omitempty"`
}

type AvailabilityParticipantTCH struct {
	PtcptSgnOff *ParticipantSignOffTCH   `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 PtcptSgnOff,omitempty"`
	PtcptSspd   *ParticipantSuspendedTCH `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 PtcptSspd,omitempty"`
}

type AvailabilityReport struct {
	Cnnctn      *Connection              `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 Cnnctn,omitempty"`
	AvlbtyPtcpt *AvailabilityParticipant `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 AvlbtyPtcpt,omitempty"`
}

type AvailabilityReportTCH struct {
	Cnnctn      *ConnectionTCH              `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 Cnnctn,omitempty"`
	AvlbtyPtcpt *AvailabilityParticipantTCH `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 AvlbtyPtcpt,omitempty"`
}

type BranchAndFinancialInstitutionIdentification4ADMN struct {
	FinInstnId FinancialInstitutionIdentification7ADMN `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 FinInstnId"`
}

type ClearingSystemMemberIdentification2ADMN struct {
	MmbId Min11Max11Text `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 MmbId"`
}

type Connection struct {
	CnnctnId []Max20AlphaNumericText `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 CnnctnId"`
}

type ConnectionTCH struct {
	CnnctnId []Max20AlphaNumericText `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 CnnctnId"`
}

type DatabaseAvailabilityReport struct {
	GrpHdr    GrpHdr                `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 GrpHdr"`
	DBRptRspn DatabaseReportReponse `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 DBRptRspn"`
	AvlbtyRpt AvailabilityReport    `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 AvlbtyRpt"`
}

type DatabaseAvailabilityReportTCH struct {
	GrpHdr    GrpHdr                   `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 GrpHdr"`
	DBRptRspn DatabaseReportReponseTCH `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 DBRptRspn"`
	AvlbtyRpt AvailabilityReportTCH    `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 AvlbtyRpt"`
}

type DatabaseReportReponse struct {
	OrgnlInstrId Max35Text                                        `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 OrgnlInstrId"`
	RptCd        ReportCode                                       `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 RptCd"`
	InstgAgt     BranchAndFinancialInstitutionIdentification4ADMN `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 InstgAgt"`
	InstdAgt     BranchAndFinancialInstitutionIdentification4ADMN `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 InstdAgt"`
	TxSts        TransactionIndividualStatus3CodeEcho             `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 TxSts"`
}

type DatabaseReportReponseTCH struct {
	OrgnlInstrId Max35Text                                        `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 OrgnlInstrId"`
	RptCd        ReportCodeTCH                                    `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 RptCd"`
	InstgAgt     BranchAndFinancialInstitutionIdentification4ADMN `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 InstgAgt"`
	InstdAgt     BranchAndFinancialInstitutionIdentification4ADMN `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 InstdAgt"`
	TxSts        TransactionIndividualStatus3CodeEcho             `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 TxSts"`
}

type Document struct {
	XMLName     xml.Name
	DBAvlbtyRpt DatabaseAvailabilityReport `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 DBAvlbtyRpt"`
}

type FinancialInstitutionIdentification7ADMN struct {
	ClrSysMmbId ClearingSystemMemberIdentification2ADMN `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 ClrSysMmbId"`
}

type GrpHdr struct {
	MsgId   Max35Text       `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 MsgId"`
	CreDtTm iso.ISODateTime `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 CreDtTm"`
}

type ParticipantSignOff struct {
	PtcptId []Min11Max11Text `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 PtcptId"`
}

type ParticipantSignOffTCH struct {
	PtcptId []Min11Max11Text `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 PtcptId"`
}

type ParticipantSuspended struct {
	PtcptId []Min11Max11Text `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 PtcptId"`
}

type ParticipantSuspendedTCH struct {
	PtcptId []Min11Max11Text `xml:"urn:iso:std:ma:20022:tech:xsd:admn.008.001.01 PtcptId"`
}

// XSD SimpleType declarations

type Max20AlphaNumericText string

type Max35Text string

type Min11Max11Text string

type ReportCode string

const ReportCodeAvlbty ReportCode = "AVLBTY"

type ReportCodeTCH string

const ReportCodeTCHAvlbty ReportCodeTCH = "AVLBTY"

type TransactionIndividualStatus3CodeEcho string

const TransactionIndividualStatus3CodeEchoActc TransactionIndividualStatus3CodeEcho = "ACTC"
