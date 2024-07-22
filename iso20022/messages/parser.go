package messages

import (
	"encoding/xml"
	"fmt"
	"reflect"

	"github.com/moov-io/iso20022/pkg/document"
	"github.com/moov-io/iso20022/pkg/head_v01"
	"github.com/moov-io/iso20022/pkg/head_v02"
	"github.com/moov-io/iso20022/pkg/pacs_v04"
	"github.com/moov-io/iso20022/pkg/pacs_v06"
	"github.com/moov-io/iso20022/pkg/pacs_v07"
	"github.com/moov-io/iso20022/pkg/pacs_v08"
	"github.com/moov-io/iso20022/pkg/pacs_v09"
	"github.com/moov-io/iso20022/pkg/pacs_v10"
	"github.com/moov-io/iso20022/pkg/pacs_v11"
	"github.com/moov-io/iso20022/pkg/utils"
	"github.com/pkg/errors"

	"github.com/CoreumFoundation/iso20022-client/iso20022/addressbook"
	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
)

type Parser struct {
	log logger.Logger
}

func NewParser(log logger.Logger) *Parser {
	return &Parser{
		log: log,
	}
}

type Envelope struct {
	XMLName  xml.Name
	Attrs    []xml.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
	AppHdr   document.Iso20022Message `xml:"AppHdr"`
	Document document.Iso20022Message `xml:",any"`
}

func (doc Envelope) Validate() error {
	if len(doc.NameSpace()) == 0 {
		return utils.Validate(&doc)
	}

	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && doc.NameSpace() == attr.Value {
			return utils.Validate(&doc)
		}
	}

	return utils.NewErrInvalidNameSpace()
}

func (doc Envelope) NameSpace() string {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			return attr.Value
		}
	}
	return ""
}

func (p Parser) parseIsoMessage(msg []byte) (header, doc document.Iso20022Message, err error) {
	dummyDoc := new(documentDummy)

	err = xml.Unmarshal(msg, dummyDoc)
	if err != nil {
		return nil, nil, err
	}

	if dummyDoc.XMLName.Local != "" {
		constructors := extendedMessageConstructor[dummyDoc.XMLName.Local]
		if constructors != nil {
			for _, constructor := range constructors {
				doc = constructor()
				err = xml.Unmarshal(msg, doc)
				if err == nil {
					return
				}
			}
			return
		}
	}

	constructor := messageConstructor[dummyDoc.NameSpace()]
	if constructor != nil {
		actualDoc := &document.Iso20022DocumentObject{
			Message: constructor(),
		}

		err = xml.Unmarshal(msg, actualDoc)
		return nil, actualDoc.Message, err
	}

	var headerConstructor func() document.Iso20022Message

	if dummyDoc.AppHdr != nil {
		headerNamespace := dummyDoc.AppHdr.NameSpace()
		if headerNamespace == "" {
			innerDoc := new(elementDummy)
			err = xml.Unmarshal(dummyDoc.AppHdr.Rest, innerDoc)
			if err != nil {
				return nil, nil, err
			}
			for _, attr := range dummyDoc.Attrs {
				if attr.Name.Local == innerDoc.XMLName.Space {
					headerNamespace = attr.Value
					break
				}
			}
			if headerNamespace == "" {
				return nil, nil, utils.NewErrOmittedNameSpace()
			}
		}

		headerConstructor = messageConstructor[headerNamespace]
		if headerConstructor == nil {
			return nil, nil, utils.NewErrUnsupportedNameSpace()
		}
	} else {
		headerConstructor = func() document.Iso20022Message {
			return nil
		}
	}

	var containedDoc document.Iso20022Message
	documentNamespace := dummyDoc.Document.NameSpace()
	if documentNamespace == "" {
		innerDoc := new(elementDummy)
		err = xml.Unmarshal(dummyDoc.Document.Rest, innerDoc)
		if err != nil {
			return nil, nil, err
		}
		for _, attr := range dummyDoc.Attrs {
			if attr.Name.Local == innerDoc.XMLName.Space {
				documentNamespace = attr.Value
				break
			}
		}
		if documentNamespace == "" {
			return nil, nil, utils.NewErrOmittedNameSpace()
		}
		documentConstructor := messageConstructor[documentNamespace]
		if documentConstructor == nil {
			return nil, nil, utils.NewErrUnsupportedNameSpace()
		}
		containedDoc = documentConstructor()
	} else {
		containedDoc, err = document.NewDocument(documentNamespace)
		if err != nil {
			return nil, nil, err
		}
	}

	envelope := &Envelope{
		AppHdr:   headerConstructor(),
		Document: containedDoc,
	}

	err = xml.Unmarshal(msg, envelope)
	if err != nil {
		return nil, nil, err
	}

	if envelope.AppHdr != nil {
		err = envelope.AppHdr.Validate()
		if err != nil {
			return nil, nil, err
		}
	}

	var resDoc document.Iso20022Message

	if innerDoc, ok := envelope.Document.(*document.Iso20022DocumentObject); ok {
		err = innerDoc.Message.Validate()
		if err != nil {
			return nil, nil, err
		}
		resDoc = innerDoc.Message
	} else {
		resDoc = envelope.Document
	}

	return envelope.AppHdr, resDoc, nil
}

func (p Parser) ExtractMetadataFromIsoMessage(msg []byte) (id string, party *addressbook.Party, err error) {
	headDoc, containedDoc, err := p.parseIsoMessage(msg)
	if err != nil {
		return "", nil, err
	}

	party = new(addressbook.Party)
	emptyParty := new(addressbook.Party)

	if headDoc != nil {
		switch head := headDoc.(type) {
		case *head_v01.BusinessApplicationHeaderV01:
			extractPartyFromHeadV01BranchAndFinancialInstitutionIdentification5(head.To.FIId, party)
			return string(head.BizMsgIdr), party, nil
		case *head_v02.BusinessApplicationHeaderV02:
			extractPartyFromHeadV02BranchAndFinancialInstitutionIdentification6(head.To.FIId, party)
			return string(head.BizMsgIdr), party, nil
		}
	}

	if containedDoc != nil {
		switch doc := containedDoc.(type) {
		case *pacs_v04.FIToFIPaymentStatusRequestV04:
			id = string(doc.GrpHdr.MsgId)
			extractPartyFromPacsV04BranchAndFinancialInstitutionIdentification6(doc.GrpHdr.InstdAgt, party)
			if (party == nil || reflect.DeepEqual(party, emptyParty)) && len(doc.TxInf) > 0 {
				if id == "" && doc.TxInf[0].StsReqId != nil {
					id = string(*doc.TxInf[0].StsReqId)
				}
				extractPartyFromPacsV04BranchAndFinancialInstitutionIdentification6(doc.TxInf[0].InstdAgt, party)
			}
		case *pacs_v04.FinancialInstitutionDirectDebitV04:
			id = string(doc.GrpHdr.MsgId)
			extractPartyFromPacsV04BranchAndFinancialInstitutionIdentification6(doc.GrpHdr.InstdAgt, party)
			if (party == nil || reflect.DeepEqual(party, emptyParty)) && len(doc.CdtInstr) > 0 {
				if id == "" {
					id = string(doc.CdtInstr[0].CdtId)
				}
				extractPartyFromPacsV04BranchAndFinancialInstitutionIdentification6(doc.CdtInstr[0].InstdAgt, party)
			}
		case *pacs_v06.FIToFICustomerCreditTransferV06:
			id = string(doc.GrpHdr.MsgId)
			extractPartyFromPacsV06BranchAndFinancialInstitutionIdentification5(doc.GrpHdr.InstdAgt, party)
			if (party == nil || reflect.DeepEqual(party, emptyParty)) && len(doc.CdtTrfTxInf) > 0 {
				if id == "" && doc.CdtTrfTxInf[0].PmtId.InstrId != nil {
					id = string(*doc.CdtTrfTxInf[0].PmtId.InstrId)
				}
				extractPartyFromPacsV06BranchAndFinancialInstitutionIdentification5(doc.CdtTrfTxInf[0].InstdAgt, party)
			}
		case *pacs_v07.FIToFIPaymentStatusReportV07:
			id = string(doc.GrpHdr.MsgId)
			extractPartyFromPacsV07BranchAndFinancialInstitutionIdentification5(doc.GrpHdr.InstdAgt, party)
			if (party == nil || reflect.DeepEqual(party, emptyParty)) && len(doc.TxInfAndSts) > 0 {
				if id == "" && doc.TxInfAndSts[0].StsId != nil {
					id = string(*doc.TxInfAndSts[0].StsId)
				}
				extractPartyFromPacsV07BranchAndFinancialInstitutionIdentification5(doc.TxInfAndSts[0].InstdAgt, party)
			}
		case *pacs_v08.FIToFICustomerCreditTransferV08:
			id = string(doc.GrpHdr.MsgId)
			extractPartyFromPacsV08BranchAndFinancialInstitutionIdentification6(doc.GrpHdr.InstdAgt, party)
			if (party == nil || reflect.DeepEqual(party, emptyParty)) && len(doc.CdtTrfTxInf) > 0 {
				if id == "" && doc.CdtTrfTxInf[0].PmtId.InstrId != nil {
					id = string(*doc.CdtTrfTxInf[0].PmtId.InstrId)
				}
				extractPartyFromPacsV08BranchAndFinancialInstitutionIdentification6(doc.CdtTrfTxInf[0].InstdAgt, party)
			}
		case *pacs_v08.FIToFICustomerDirectDebitV08:
			id = string(doc.GrpHdr.MsgId)
			extractPartyFromPacsV08BranchAndFinancialInstitutionIdentification6(doc.GrpHdr.InstdAgt, party)
			if (party == nil || reflect.DeepEqual(party, emptyParty)) && len(doc.DrctDbtTxInf) > 0 {
				if id == "" && doc.DrctDbtTxInf[0].PmtId.InstrId != nil {
					id = string(*doc.DrctDbtTxInf[0].PmtId.InstrId)
				}
				extractPartyFromPacsV08BranchAndFinancialInstitutionIdentification6(doc.DrctDbtTxInf[0].InstdAgt, party)
			}
		case *pacs_v08.FIToFIPaymentStatusReportV08:
			id = string(doc.GrpHdr.MsgId)
			extractPartyFromPacsV08BranchAndFinancialInstitutionIdentification5(doc.GrpHdr.InstdAgt, party)
			if (party == nil || reflect.DeepEqual(party, emptyParty)) && len(doc.TxInfAndSts) > 0 {
				if id == "" && doc.TxInfAndSts[0].StsId != nil {
					id = string(*doc.TxInfAndSts[0].StsId)
				}
				extractPartyFromPacsV08BranchAndFinancialInstitutionIdentification5(doc.TxInfAndSts[0].InstdAgt, party)
			}
		case *pacs_v09.FIToFICustomerCreditTransferV09:
			id = string(doc.GrpHdr.MsgId)
			extractPartyFromPacsV09BranchAndFinancialInstitutionIdentification6(doc.GrpHdr.InstdAgt, party)
			if (party == nil || reflect.DeepEqual(party, emptyParty)) && len(doc.CdtTrfTxInf) > 0 {
				if id == "" && doc.CdtTrfTxInf[0].PmtId.InstrId != nil {
					id = string(*doc.CdtTrfTxInf[0].PmtId.InstrId)
				}
				if doc.CdtTrfTxInf[0].InstdAgt != nil {
					extractPartyFromPacsV09BranchAndFinancialInstitutionIdentification6(doc.CdtTrfTxInf[0].InstdAgt, party)
				} else {
					extractPartyFromPacsV09BranchAndFinancialInstitutionIdentification6(&doc.CdtTrfTxInf[0].CdtrAgt, party)
				}
			}
		case *pacs_v09.FinancialInstitutionCreditTransferV09:
			id = string(doc.GrpHdr.MsgId)
			extractPartyFromPacsV09BranchAndFinancialInstitutionIdentification6(doc.GrpHdr.InstdAgt, party)
			if (party == nil || reflect.DeepEqual(party, emptyParty)) && len(doc.CdtTrfTxInf) > 0 {
				if id == "" && doc.CdtTrfTxInf[0].PmtId.InstrId != nil {
					id = string(*doc.CdtTrfTxInf[0].PmtId.InstrId)
				}
				if doc.CdtTrfTxInf[0].InstdAgt != nil {
					extractPartyFromPacsV09BranchAndFinancialInstitutionIdentification6(doc.CdtTrfTxInf[0].InstdAgt, party)
				} else {
					extractPartyFromPacsV09BranchAndFinancialInstitutionIdentification6(doc.CdtTrfTxInf[0].CdtrAgt, party)
				}
			}
		case *pacs_v10.FIToFIPaymentReversalV10:
			id = string(doc.GrpHdr.MsgId)
			extractPartyFromPacsV10BranchAndFinancialInstitutionIdentification6(doc.GrpHdr.InstdAgt, party)
			if (party == nil || reflect.DeepEqual(party, emptyParty)) && len(doc.TxInf) > 0 {
				if id == "" && doc.TxInf[0].RvslId != nil {
					id = string(*doc.TxInf[0].RvslId)
				}
				extractPartyFromPacsV10BranchAndFinancialInstitutionIdentification6(doc.TxInf[0].InstdAgt, party)
			}
		case *pacs_v10.FIToFIPaymentStatusReportV10:
			id = string(doc.GrpHdr.MsgId)
			extractPartyFromPacsV10BranchAndFinancialInstitutionIdentification6(doc.GrpHdr.InstdAgt, party)
			if (party == nil || reflect.DeepEqual(party, emptyParty)) && len(doc.TxInfAndSts) > 0 {
				if id == "" && doc.TxInfAndSts[0].StsId != nil {
					id = string(*doc.TxInfAndSts[0].StsId)
				}
				extractPartyFromPacsV10BranchAndFinancialInstitutionIdentification6(doc.TxInfAndSts[0].InstdAgt, party)
			}
		case *pacs_v10.PaymentReturnV10:
			id = string(doc.GrpHdr.MsgId)
			extractPartyFromPacsV10BranchAndFinancialInstitutionIdentification6(doc.GrpHdr.InstdAgt, party)
			if (party == nil || reflect.DeepEqual(party, emptyParty)) && len(doc.TxInf) > 0 {
				if id == "" && doc.TxInf[0].RtrId != nil {
					id = string(*doc.TxInf[0].RtrId)
				}
				extractPartyFromPacsV10BranchAndFinancialInstitutionIdentification6(doc.TxInf[0].InstdAgt, party)
			}
		case *pacs_v11.FIToFIPaymentStatusReportV11:
			id = string(doc.GrpHdr.MsgId)
			extractPartyFromPacsV11BranchAndFinancialInstitutionIdentification6(doc.GrpHdr.InstdAgt, party)
			if (party == nil || reflect.DeepEqual(party, emptyParty)) && len(doc.TxInfAndSts) > 0 {
				if id == "" && doc.TxInfAndSts[0].StsId != nil {
					id = string(*doc.TxInfAndSts[0].StsId)
				}
				extractPartyFromPacsV11BranchAndFinancialInstitutionIdentification6(doc.TxInfAndSts[0].InstdAgt, party)
			}
		default:
			return "", nil, fmt.Errorf("couldn't find party from %v", reflect.TypeOf(containedDoc).String())
		}
	}

	if party == nil || reflect.DeepEqual(party, emptyParty) {
		return "", nil, errors.New("couldn't find party")
	}

	return
}
