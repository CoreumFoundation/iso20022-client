package messages

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/document"
	"github.com/moov-io/iso20022/pkg/head_v01"
	"github.com/moov-io/iso20022/pkg/head_v02"
	"github.com/moov-io/iso20022/pkg/pacs_v04"
	"github.com/moov-io/iso20022/pkg/pacs_v08"
	"github.com/moov-io/iso20022/pkg/pacs_v09"
	"github.com/moov-io/iso20022/pkg/pacs_v10"
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
	Document document.Iso20022Message `xml:"Document"`
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

	constructor := messageConstructor[dummyDoc.NameSpace()]
	if constructor != nil {
		actualDoc := &document.Iso20022DocumentObject{
			Message: constructor(),
		}

		err = xml.Unmarshal(msg, actualDoc)
		return nil, actualDoc.Message, err
	}

	headerNamespace := dummyDoc.AppHdr.NameSpace()
	if headerNamespace == "" {
		return nil, nil, utils.NewErrOmittedNameSpace()
	}

	headerConstructor := messageConstructor[headerNamespace]
	if headerConstructor == nil {
		return nil, nil, utils.NewErrUnsupportedNameSpace()
	}

	documentNamespace := dummyDoc.Document.NameSpace()
	if documentNamespace == "" {
		return nil, nil, utils.NewErrOmittedNameSpace()
	}

	containedDoc, err := document.NewDocument(documentNamespace)
	if err != nil {
		return nil, nil, err
	}

	envelope := &Envelope{
		AppHdr:   headerConstructor(),
		Document: containedDoc,
	}

	err = xml.Unmarshal(msg, envelope)
	if err != nil {
		return nil, nil, err
	}

	err = envelope.Validate()
	if err != nil {
		return nil, nil, err
	}

	return envelope.AppHdr, envelope.Document.(*document.Iso20022DocumentObject).Message, nil
}

func (p Parser) ExtractIdentificationFromIsoMessage(msg []byte) (*addressbook.Party, error) {
	headerDoc, containedDoc, err := p.parseIsoMessage(msg)
	if err != nil {
		return nil, err
	}

	res := new(addressbook.Party)

	if headerDoc != nil {
		switch header := headerDoc.(type) {
		case *head_v01.BusinessApplicationHeaderV01:
			extractPartyFromHeadV01BranchAndFinancialInstitutionIdentification5(header.To.FIId, res)
			return res, nil
		case *head_v02.BusinessApplicationHeaderV02:
			extractPartyFromHeadV02BranchAndFinancialInstitutionIdentification6(header.To.FIId, res)
			return res, nil
		}
	}

	if containedDoc != nil {
		switch doc := containedDoc.(type) {
		case pacs_v04.FIToFIPaymentStatusRequestV04:
			extractPartyFromPacsV04BranchAndFinancialInstitutionIdentification6(doc.GrpHdr.InstdAgt, res)
			if len(doc.TxInf) > 0 {
				extractPartyFromPacsV04BranchAndFinancialInstitutionIdentification6(doc.TxInf[0].InstdAgt, res)
			}
			return res, nil
		case *pacs_v04.FinancialInstitutionDirectDebitV04:
			extractPartyFromPacsV04BranchAndFinancialInstitutionIdentification6(doc.GrpHdr.InstdAgt, res)
			if len(doc.CdtInstr) > 0 {
				extractPartyFromPacsV04BranchAndFinancialInstitutionIdentification6(doc.CdtInstr[0].InstdAgt, res)
			}
			return res, nil
		case *pacs_v08.FIToFICustomerCreditTransferV08:
			extractPartyFromPacsV08BranchAndFinancialInstitutionIdentification6(doc.GrpHdr.InstdAgt, res)
			if len(doc.CdtTrfTxInf) > 0 {
				extractPartyFromPacsV08BranchAndFinancialInstitutionIdentification6(doc.CdtTrfTxInf[0].InstdAgt, res)
			}
			return res, nil
		case *pacs_v08.FIToFICustomerDirectDebitV08:
			extractPartyFromPacsV08BranchAndFinancialInstitutionIdentification6(doc.GrpHdr.InstdAgt, res)
			if len(doc.DrctDbtTxInf) > 0 {
				extractPartyFromPacsV08BranchAndFinancialInstitutionIdentification6(doc.DrctDbtTxInf[0].InstdAgt, res)
			}
			return res, nil
		case *pacs_v08.FIToFIPaymentStatusReportV08:
			extractPartyFromPacsV08BranchAndFinancialInstitutionIdentification5(doc.GrpHdr.InstdAgt, res)
			if len(doc.TxInfAndSts) > 0 {
				extractPartyFromPacsV08BranchAndFinancialInstitutionIdentification5(doc.TxInfAndSts[0].InstdAgt, res)
			}
			return res, nil
		case *pacs_v09.FIToFICustomerCreditTransferV09:
			extractPartyFromPacsV09BranchAndFinancialInstitutionIdentification6(doc.GrpHdr.InstdAgt, res)
			if len(doc.CdtTrfTxInf) > 0 {
				extractPartyFromPacsV09BranchAndFinancialInstitutionIdentification6(doc.CdtTrfTxInf[0].InstdAgt, res)
			}
			return res, nil
		case *pacs_v09.FinancialInstitutionCreditTransferV09:
			extractPartyFromPacsV09BranchAndFinancialInstitutionIdentification6(doc.GrpHdr.InstdAgt, res)
			if len(doc.CdtTrfTxInf) > 0 {
				extractPartyFromPacsV09BranchAndFinancialInstitutionIdentification6(doc.CdtTrfTxInf[0].InstdAgt, res)
			}
			return res, nil
		case pacs_v10.FIToFIPaymentReversalV10:
			extractPartyFromPacsV10BranchAndFinancialInstitutionIdentification6(doc.GrpHdr.InstdAgt, res)
			if len(doc.TxInf) > 0 {
				extractPartyFromPacsV10BranchAndFinancialInstitutionIdentification6(doc.TxInf[0].InstdAgt, res)
			}
			return res, nil
		case pacs_v10.PaymentReturnV10:
			extractPartyFromPacsV10BranchAndFinancialInstitutionIdentification6(doc.GrpHdr.InstdAgt, res)
			if len(doc.TxInf) > 0 {
				extractPartyFromPacsV10BranchAndFinancialInstitutionIdentification6(doc.TxInf[0].InstdAgt, res)
			}
			return res, nil
		}
	}

	return nil, errors.New("couldn't find identification")
}
