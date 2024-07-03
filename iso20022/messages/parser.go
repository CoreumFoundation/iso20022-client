package messages

import (
	"context"
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/document"
	"github.com/moov-io/iso20022/pkg/head_v01"
	"github.com/moov-io/iso20022/pkg/head_v02"
	"github.com/moov-io/iso20022/pkg/pacs_v08"
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

func (p Parser) ParseIsoMessage(msg []byte) (header, doc document.Iso20022Message, err error) {
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

func (p Parser) ExtractIdentificationFromIsoMessage(ctx context.Context, msg []byte) (*addressbook.BranchAndIdentification, error) {
	header, containedDoc, err := p.ParseIsoMessage(msg)
	if err != nil {
		return nil, err
	}

	res := new(addressbook.BranchAndIdentification)

	if header != nil {
		switch header := header.(type) {
		case *head_v01.BusinessApplicationHeaderV01:
			//header.To.OrgId
			//header.To.FIId
			if header.To.FIId.FinInstnId.BICFI != nil {
				res.Identification.Bic = string(*header.To.FIId.FinInstnId.BICFI)
			}
			return res, nil
		case *head_v02.BusinessApplicationHeaderV02:
			//header.To.OrgId
			//header.To.FIId
			if header.To.FIId.FinInstnId.BICFI != nil {
				res.Identification.Bic = string(*header.To.FIId.FinInstnId.BICFI)
			}
			return res, nil
		}
	}
	if containedDoc != nil {
		switch doc := containedDoc.(type) {
		case *pacs_v08.FIToFICustomerCreditTransferV08:
			extractFromBranchAndFinancialInstitutionIdentification6(doc.GrpHdr.InstdAgt, res)
			if len(doc.CdtTrfTxInf) > 0 {
				extractFromBranchAndFinancialInstitutionIdentification6(doc.CdtTrfTxInf[0].InstdAgt, res)
			}
			return res, nil
		}
	}

	return nil, errors.New("couldn't find identification")
}

func extractFromBranchAndFinancialInstitutionIdentification6(agent *pacs_v08.BranchAndFinancialInstitutionIdentification6, res *addressbook.BranchAndIdentification) {
	if agent == nil {
		return
	}

	// agent.BrnchId
	if agent.FinInstnId.BICFI != nil {
		res.Identification.Bic = string(*agent.FinInstnId.BICFI)
	}
}
