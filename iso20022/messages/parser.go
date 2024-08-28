package messages

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/xml"
	"reflect"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/head_001_001_01"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/head_001_001_02"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/messages"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_002_001_07"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_002_001_08"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_002_001_10"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_002_001_11"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_003_001_08"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_004_001_10"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_007_001_10"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_008_001_06"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_008_001_08"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_008_001_09"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_008_001_12"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_009_001_09"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_010_001_04"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_028_001_04"
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/pacs_028_001_06"
	"github.com/CoreumFoundation/iso20022-client/iso20022/addressbook"
	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
	"github.com/CoreumFoundation/iso20022-client/iso20022/processes"
)

type Parser struct {
	log       logger.Logger
	converter Converter
}

func NewParser(log logger.Logger, converter Converter) *Parser {
	return &Parser{
		log:       log,
		converter: converter,
	}
}

type Envelope struct {
	XMLName  xml.Name
	Attrs    []xml.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
	AppHdr   messages.Iso20022Message `xml:"AppHdr"`
	Document messages.Iso20022Message `xml:",any"`
}

func (doc Envelope) Validate() error {
	if len(doc.NameSpace()) == 0 {
		return Validate(&doc)
	}

	for _, attr := range doc.Attrs {
		if attr.Name.Local == XmlDefaultNamespace && doc.NameSpace() == attr.Value {
			return Validate(&doc)
		}
	}

	return errors.New("The namespace of document is invalid")
}

func (doc Envelope) NameSpace() string {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == XmlDefaultNamespace {
			return attr.Value
		}
	}
	return ""
}

func (p Parser) parseIsoMessage(msg []byte) (supl processes.SupplementaryDataParser, header, doc messages.Iso20022Message, err error) {
	dummyDoc := new(documentDummy)
	attrs := make(map[string]string)

	err = xml.Unmarshal(msg, dummyDoc)
	if err != nil {
		return nil, nil, nil, err
	}

	if dummyDoc.XMLName.Space != "" {
		if ns, ok := attrs[dummyDoc.XMLName.Space]; ok {
			constructor := messageConstructor[ns]
			actualDoc := &Iso20022DocumentObject{
				Message: constructor(),
			}
			err = xml.Unmarshal(msg, actualDoc)
			return nil, nil, actualDoc.Message, err
		}
	}

	for _, attr := range dummyDoc.Attrs {
		attrs[attr.Name.Local] = attr.Value
		supl = newSupplementaryDataParser(attrs)
	}

	if dummyDoc.XMLName.Local != "" {
		constructors := extendedMessageConstructor[dummyDoc.XMLName.Local]
		if constructors != nil {
			for _, constructor := range constructors {
				doc = constructor.Constructor()
				aa := xml.NewDecoder(bytes.NewReader(msg))
				aa.DefaultSpace = constructor.Urn
				err = aa.Decode(doc)
				if err == nil {
					err = doc.Validate()
					if err == nil {
						return
					}
				}
			}
			return
		}
	}

	constructor := messageConstructor[dummyDoc.NameSpace()]
	if constructor != nil {
		actualDoc := &Iso20022DocumentObject{
			Message: constructor(),
		}

		err = xml.Unmarshal(msg, actualDoc)
		return supl, nil, actualDoc.Message, err
	}

	var headerConstructor func() messages.Iso20022Message

	headerNamespace := ""
	if dummyDoc.AppHdr != nil {
		headerNamespace = dummyDoc.AppHdr.NameSpace()
		if headerNamespace == "" {
			innerDoc := new(elementDummy)
			err = xml.Unmarshal(dummyDoc.AppHdr.Rest, innerDoc)
			if err != nil {
				return supl, nil, nil, err
			}
			if ns, ok := attrs[innerDoc.XMLName.Space]; ok {
				headerNamespace = ns
			}
			if headerNamespace == "" {
				return supl, nil, nil, errors.New("The namespace of document is omitted")
			}
		}

		headerConstructor = messageConstructor[headerNamespace]
		if headerConstructor == nil {
			return supl, nil, nil, errors.New("The namespace of document is unsupported")
		}
	} else {
		headerConstructor = func() messages.Iso20022Message {
			return nil
		}
	}

	var containedDoc messages.Iso20022Message
	documentNamespace := dummyDoc.Document.NameSpace()
	if documentNamespace == "" {
		innerDoc := new(elementDummy)
		err = xml.Unmarshal(dummyDoc.Document.Rest, innerDoc)
		if err != nil {
			return supl, nil, nil, err
		}
		if ns, ok := attrs[innerDoc.XMLName.Space]; ok {
			documentNamespace = ns
		}
		if documentNamespace == "" {
			return supl, nil, nil, errors.New("The namespace of document is omitted")
		}
		documentConstructor := messageConstructor[documentNamespace]
		if documentConstructor == nil {
			return supl, nil, nil, errors.New("The namespace of document is unsupported")
		}
		containedDoc = documentConstructor()
	} else {
		constructor = messageConstructor[documentNamespace]
		if constructor == nil {
			return supl, nil, nil, errors.New("The namespace of document is unsupported")
		}

		containedDoc = &Iso20022DocumentObject{
			Message: constructor(),
		}
	}

	envelope := &Envelope{
		AppHdr:   headerConstructor(),
		Document: containedDoc,
	}

	err = xml.Unmarshal(msg, envelope)
	if err != nil {
		return supl, nil, nil, err
	}

	if envelope.AppHdr != nil {
		err = envelope.AppHdr.Validate()
		if err != nil {
			return supl, nil, nil, err
		}
	}

	var resDoc messages.Iso20022Message

	if innerDoc, ok := envelope.Document.(*Iso20022DocumentObject); ok {
		err = innerDoc.Message.Validate()
		if err != nil {
			return supl, nil, nil, err
		}
		resDoc = innerDoc.Message
	} else {
		resDoc = envelope.Document
	}

	return supl, envelope.AppHdr, resDoc, nil
}

type SupplementaryDataParser struct {
	attrs map[string]string
}

func newSupplementaryDataParser(attrs map[string]string) processes.SupplementaryDataParser {
	return &SupplementaryDataParser{
		attrs: attrs,
	}
}

func (s *SupplementaryDataParser) Parse(msg []byte) (doc messages.Iso20022Message, err error) {
	dummyDoc := new(documentDummy)

	err = xml.Unmarshal(msg, dummyDoc)
	if err != nil {
		return nil, err
	}

	if dummyDoc.XMLName.Space != "" {
		if ns, ok := s.attrs[dummyDoc.XMLName.Space]; ok {
			constructor := messageConstructor[ns]
			if constructor != nil {
				actualDoc := &Iso20022DocumentObject{
					Message: constructor(),
				}
				err = xml.Unmarshal(msg, actualDoc)
				return actualDoc.Message, err
			}
		}
	}
	return nil, errors.New("could not extract the supplementary data")
}

func (p Parser) ExtractMessageAndMetadataFromIsoMessage(msg []byte) (message messages.Iso20022Message, metadata processes.Metadata, references *processes.Metadata, supplementaryDataParser processes.SupplementaryDataParser, err error) {
	supplementaryDataParser, headDoc, containedDoc, err := p.parseIsoMessage(msg)
	if err != nil {
		return message, metadata, references, supplementaryDataParser, err
	}

	uetr := ""
	endToEndID := ""
	txID := ""
	id := ""
	sender := new(addressbook.Party)
	receiver := new(addressbook.Party)

	origId := ""
	origUetr := ""
	origEndToEndID := ""
	origTxID := ""
	origSender := new(addressbook.Party)
	origReceiver := new(addressbook.Party)

	emptyParty := new(addressbook.Party)

	if containedDoc != nil {
		switch doc := containedDoc.(type) {
		case *pacs_028_001_04.FIToFIPaymentStatusRequestV04:
			id = string(doc.GrpHdr.MsgId)
			receiver = p.converter.ConvertFromPacs02800104(doc.GrpHdr.InstdAgt).ToParty()
			sender = p.converter.ConvertFromPacs02800104(doc.GrpHdr.InstgAgt).ToParty()
			if (receiver == nil || reflect.DeepEqual(receiver, emptyParty)) && len(doc.TxInf) > 0 {
				if id == "" && doc.TxInf[0].StsReqId != nil {
					id = string(*doc.TxInf[0].StsReqId)
				}
				receiver = p.converter.ConvertFromPacs02800104(doc.TxInf[0].InstdAgt).ToParty()
			}
			if (sender == nil || reflect.DeepEqual(sender, emptyParty)) && len(doc.TxInf) > 0 {
				sender = p.converter.ConvertFromPacs02800104(doc.TxInf[0].InstgAgt).ToParty()
			}
			if len(doc.OrgnlGrpInf) > 0 {
				origId = string(doc.OrgnlGrpInf[0].OrgnlMsgId)
			}
			if len(doc.TxInf) > 0 {
				if origId == "" {
					origId = string(doc.TxInf[0].OrgnlGrpInf.OrgnlMsgId)
					if origId == "" && doc.TxInf[0].OrgnlInstrId != nil {
						origId = string(*doc.TxInf[0].OrgnlInstrId)
					}
				}
				if doc.TxInf[0].OrgnlUETR != nil {
					origUetr = string(*doc.TxInf[0].OrgnlUETR)
				}
				if doc.TxInf[0].OrgnlEndToEndId != nil {
					origEndToEndID = string(*doc.TxInf[0].OrgnlEndToEndId)
				}
				if doc.TxInf[0].OrgnlTxId != nil {
					origTxID = string(*doc.TxInf[0].OrgnlTxId)
				}
			}
			origSender = sender
			origReceiver = receiver
		case *pacs_028_001_06.FIToFIPaymentStatusRequestV06:
			id = string(doc.GrpHdr.MsgId)
			receiver = p.converter.ConvertFromPacs02800106(doc.GrpHdr.InstdAgt).ToParty()
			sender = p.converter.ConvertFromPacs02800106(doc.GrpHdr.InstgAgt).ToParty()
			if (receiver == nil || reflect.DeepEqual(receiver, emptyParty)) && len(doc.TxInf) > 0 {
				if id == "" && doc.TxInf[0].StsReqId != nil {
					id = string(*doc.TxInf[0].StsReqId)
				}
				receiver = p.converter.ConvertFromPacs02800106(doc.TxInf[0].InstdAgt).ToParty()
			}
			if (sender == nil || reflect.DeepEqual(sender, emptyParty)) && len(doc.TxInf) > 0 {
				sender = p.converter.ConvertFromPacs02800106(doc.TxInf[0].InstgAgt).ToParty()
			}
			if len(doc.OrgnlGrpInf) > 0 {
				origId = string(doc.OrgnlGrpInf[0].OrgnlMsgId)
			}
			if len(doc.TxInf) > 0 {
				if origId == "" {
					origId = string(doc.TxInf[0].OrgnlGrpInf.OrgnlMsgId)
					if origId == "" && doc.TxInf[0].OrgnlInstrId != nil {
						origId = string(*doc.TxInf[0].OrgnlInstrId)
					}
				}
				if doc.TxInf[0].OrgnlUETR != nil {
					origUetr = string(*doc.TxInf[0].OrgnlUETR)
				}
				if doc.TxInf[0].OrgnlEndToEndId != nil {
					origEndToEndID = string(*doc.TxInf[0].OrgnlEndToEndId)
				}
				if doc.TxInf[0].OrgnlTxId != nil {
					origTxID = string(*doc.TxInf[0].OrgnlTxId)
				}
			}
			origSender = sender
			origReceiver = receiver
		case *pacs_010_001_04.FinancialInstitutionDirectDebitV04:
			id = string(doc.GrpHdr.MsgId)
			receiver = p.converter.ConvertFromPacs01000104(doc.GrpHdr.InstdAgt).ToParty()
			sender = p.converter.ConvertFromPacs01000104(doc.GrpHdr.InstgAgt).ToParty()
			if (receiver == nil || reflect.DeepEqual(receiver, emptyParty)) && len(doc.CdtInstr) > 0 {
				if id == "" {
					id = string(doc.CdtInstr[0].CdtId)
				}
				receiver = p.converter.ConvertFromPacs01000104(doc.CdtInstr[0].InstdAgt).ToParty()
			}
			if (sender == nil || reflect.DeepEqual(sender, emptyParty)) && len(doc.CdtInstr) > 0 {
				sender = p.converter.ConvertFromPacs01000104(doc.CdtInstr[0].InstgAgt).ToParty()
			}
			if len(doc.CdtInstr) > 0 && len(doc.CdtInstr[0].DrctDbtTxInf) > 0 {
				endToEndID = string(doc.CdtInstr[0].DrctDbtTxInf[0].PmtId.EndToEndId)
				if doc.CdtInstr[0].DrctDbtTxInf[0].PmtId.UETR != nil {
					uetr = string(*doc.CdtInstr[0].DrctDbtTxInf[0].PmtId.UETR)
				}
				if doc.CdtInstr[0].DrctDbtTxInf[0].PmtId.TxId != nil {
					txID = string(*doc.CdtInstr[0].DrctDbtTxInf[0].PmtId.TxId)
				}
			}
		case *pacs_008_001_06.FIToFICustomerCreditTransferV06:
			id = string(doc.GrpHdr.MsgId)
			receiver = p.converter.ConvertFromPacs00800106(doc.GrpHdr.InstdAgt).ToParty()
			sender = p.converter.ConvertFromPacs00800106(doc.GrpHdr.InstgAgt).ToParty()
			if (receiver == nil || reflect.DeepEqual(receiver, emptyParty)) && len(doc.CdtTrfTxInf) > 0 {
				if id == "" && doc.CdtTrfTxInf[0].PmtId.InstrId != nil {
					id = string(*doc.CdtTrfTxInf[0].PmtId.InstrId)
				}
				receiver = p.converter.ConvertFromPacs00800106(doc.CdtTrfTxInf[0].InstdAgt).ToParty()
			}
			if (sender == nil || reflect.DeepEqual(sender, emptyParty)) && len(doc.CdtTrfTxInf) > 0 {
				sender = p.converter.ConvertFromPacs00800106(doc.CdtTrfTxInf[0].InstgAgt).ToParty()
			}
			if len(doc.CdtTrfTxInf) > 0 {
				endToEndID = string(doc.CdtTrfTxInf[0].PmtId.EndToEndId)
				if doc.CdtTrfTxInf[0].PmtId.TxId != "" {
					txID = string(doc.CdtTrfTxInf[0].PmtId.TxId)
				}
			}
		case *pacs_002_001_07.FIToFIPaymentStatusReportV07:
			id = string(doc.GrpHdr.MsgId)
			receiver = p.converter.ConvertFromPacs00200107(doc.GrpHdr.InstdAgt).ToParty()
			sender = p.converter.ConvertFromPacs00200107(doc.GrpHdr.InstgAgt).ToParty()
			if (receiver == nil || reflect.DeepEqual(receiver, emptyParty)) && len(doc.TxInfAndSts) > 0 {
				if id == "" && doc.TxInfAndSts[0].StsId != nil {
					id = string(*doc.TxInfAndSts[0].StsId)
				}
				receiver = p.converter.ConvertFromPacs00200107(doc.TxInfAndSts[0].InstdAgt).ToParty()
			}
			if (sender == nil || reflect.DeepEqual(sender, emptyParty)) && len(doc.TxInfAndSts) > 0 {
				sender = p.converter.ConvertFromPacs00200107(doc.TxInfAndSts[0].InstgAgt).ToParty()
			}
			if len(doc.OrgnlGrpInfAndSts) > 0 {
				origId = string(doc.OrgnlGrpInfAndSts[0].OrgnlMsgId)
			}
			if len(doc.TxInfAndSts) > 0 {
				if origId == "" {
					if doc.TxInfAndSts[0].OrgnlGrpInf != nil {
						origId = string(doc.TxInfAndSts[0].OrgnlGrpInf.OrgnlMsgId)
					}
					if origId == "" && doc.TxInfAndSts[0].OrgnlInstrId != nil {
						origId = string(*doc.TxInfAndSts[0].OrgnlInstrId)
					}
				}
				if doc.TxInfAndSts[0].OrgnlEndToEndId != nil {
					origEndToEndID = string(*doc.TxInfAndSts[0].OrgnlEndToEndId)
				}
				if doc.TxInfAndSts[0].OrgnlTxId != nil {
					origTxID = string(*doc.TxInfAndSts[0].OrgnlTxId)
				}
			}
			origSender = receiver
			origReceiver = sender
		case *pacs_008_001_08.FIToFICustomerCreditTransferV08:
			id = string(doc.GrpHdr.MsgId)
			receiver = p.converter.ConvertFromPacs00800108(doc.GrpHdr.InstdAgt).ToParty()
			sender = p.converter.ConvertFromPacs00800108(doc.GrpHdr.InstgAgt).ToParty()
			if (receiver == nil || reflect.DeepEqual(receiver, emptyParty)) && len(doc.CdtTrfTxInf) > 0 {
				if id == "" && doc.CdtTrfTxInf[0].PmtId.InstrId != nil {
					id = string(*doc.CdtTrfTxInf[0].PmtId.InstrId)
				}
				receiver = p.converter.ConvertFromPacs00800108(doc.CdtTrfTxInf[0].InstdAgt).ToParty()
			}
			if (sender == nil || reflect.DeepEqual(sender, emptyParty)) && len(doc.CdtTrfTxInf) > 0 {
				sender = p.converter.ConvertFromPacs00800108(doc.CdtTrfTxInf[0].InstgAgt).ToParty()
			}
			if len(doc.CdtTrfTxInf) > 0 {
				endToEndID = string(doc.CdtTrfTxInf[0].PmtId.EndToEndId)
				if doc.CdtTrfTxInf[0].PmtId.UETR != nil {
					uetr = string(*doc.CdtTrfTxInf[0].PmtId.UETR)
				}
				if doc.CdtTrfTxInf[0].PmtId.TxId != nil {
					txID = string(*doc.CdtTrfTxInf[0].PmtId.TxId)
				}
			}
		case *pacs_003_001_08.FIToFICustomerDirectDebitV08:
			id = string(doc.GrpHdr.MsgId)
			receiver = p.converter.ConvertFromPacs00300108(doc.GrpHdr.InstdAgt).ToParty()
			sender = p.converter.ConvertFromPacs00300108(doc.GrpHdr.InstgAgt).ToParty()
			if (receiver == nil || reflect.DeepEqual(receiver, emptyParty)) && len(doc.DrctDbtTxInf) > 0 {
				if id == "" && doc.DrctDbtTxInf[0].PmtId.InstrId != nil {
					id = string(*doc.DrctDbtTxInf[0].PmtId.InstrId)
				}
				receiver = p.converter.ConvertFromPacs00300108(doc.DrctDbtTxInf[0].InstdAgt).ToParty()
			}
			if (sender == nil || reflect.DeepEqual(sender, emptyParty)) && len(doc.DrctDbtTxInf) > 0 {
				sender = p.converter.ConvertFromPacs00300108(doc.DrctDbtTxInf[0].InstgAgt).ToParty()
			}
			if len(doc.DrctDbtTxInf) > 0 {
				endToEndID = string(doc.DrctDbtTxInf[0].PmtId.EndToEndId)
				if doc.DrctDbtTxInf[0].PmtId.UETR != nil {
					uetr = string(*doc.DrctDbtTxInf[0].PmtId.UETR)
				}
				if doc.DrctDbtTxInf[0].PmtId.TxId != nil {
					txID = string(*doc.DrctDbtTxInf[0].PmtId.TxId)
				}
			}
		case *pacs_002_001_08.FIToFIPaymentStatusReportV08:
			id = string(doc.GrpHdr.MsgId)
			receiver = p.converter.ConvertFromPacs00200108(doc.GrpHdr.InstdAgt).ToParty()
			sender = p.converter.ConvertFromPacs00200108(doc.GrpHdr.InstgAgt).ToParty()
			if (receiver == nil || reflect.DeepEqual(receiver, emptyParty)) && len(doc.TxInfAndSts) > 0 {
				if id == "" && doc.TxInfAndSts[0].StsId != nil {
					id = string(*doc.TxInfAndSts[0].StsId)
				}
				receiver = p.converter.ConvertFromPacs00200108(doc.TxInfAndSts[0].InstdAgt).ToParty()
			}
			if (sender == nil || reflect.DeepEqual(sender, emptyParty)) && len(doc.TxInfAndSts) > 0 {
				sender = p.converter.ConvertFromPacs00200108(doc.TxInfAndSts[0].InstgAgt).ToParty()
			}
			if len(doc.OrgnlGrpInfAndSts) > 0 {
				origId = string(doc.OrgnlGrpInfAndSts[0].OrgnlMsgId)
			}
			if len(doc.TxInfAndSts) > 0 {
				if origId == "" {
					if doc.TxInfAndSts[0].OrgnlGrpInf != nil {
						origId = string(doc.TxInfAndSts[0].OrgnlGrpInf.OrgnlMsgId)
					}
					if origId == "" && doc.TxInfAndSts[0].OrgnlInstrId != nil {
						origId = string(*doc.TxInfAndSts[0].OrgnlInstrId)
					}
				}
				if doc.TxInfAndSts[0].OrgnlEndToEndId != nil {
					origEndToEndID = string(*doc.TxInfAndSts[0].OrgnlEndToEndId)
				}
				if doc.TxInfAndSts[0].OrgnlTxId != nil {
					origTxID = string(*doc.TxInfAndSts[0].OrgnlTxId)
				}
			}
			origSender = receiver
			origReceiver = sender
		case *pacs_008_001_09.FIToFICustomerCreditTransferV09:
			id = string(doc.GrpHdr.MsgId)
			receiver = p.converter.ConvertFromPacs00800109(doc.GrpHdr.InstdAgt).ToParty()
			sender = p.converter.ConvertFromPacs00800109(doc.GrpHdr.InstgAgt).ToParty()
			if (receiver == nil || reflect.DeepEqual(receiver, emptyParty)) && len(doc.CdtTrfTxInf) > 0 {
				if id == "" && doc.CdtTrfTxInf[0].PmtId.InstrId != nil {
					id = string(*doc.CdtTrfTxInf[0].PmtId.InstrId)
				}
				if doc.CdtTrfTxInf[0].InstdAgt != nil {
					receiver = p.converter.ConvertFromPacs00800109(doc.CdtTrfTxInf[0].InstdAgt).ToParty()
				}
			}
			if (sender == nil || reflect.DeepEqual(sender, emptyParty)) && len(doc.CdtTrfTxInf) > 0 {
				if doc.CdtTrfTxInf[0].InstgAgt != nil {
					sender = p.converter.ConvertFromPacs00800109(doc.CdtTrfTxInf[0].InstgAgt).ToParty()
				}
			}
			if len(doc.CdtTrfTxInf) > 0 {
				endToEndID = string(doc.CdtTrfTxInf[0].PmtId.EndToEndId)
				if doc.CdtTrfTxInf[0].PmtId.UETR != nil {
					uetr = string(*doc.CdtTrfTxInf[0].PmtId.UETR)
				}
				if doc.CdtTrfTxInf[0].PmtId.TxId != nil {
					txID = string(*doc.CdtTrfTxInf[0].PmtId.TxId)
				}
			}
		case *pacs_009_001_09.FinancialInstitutionCreditTransferV09:
			id = string(doc.GrpHdr.MsgId)
			receiver = p.converter.ConvertFromPacs00900109(doc.GrpHdr.InstdAgt).ToParty()
			sender = p.converter.ConvertFromPacs00900109(doc.GrpHdr.InstgAgt).ToParty()
			if (receiver == nil || reflect.DeepEqual(receiver, emptyParty)) && len(doc.CdtTrfTxInf) > 0 {
				if id == "" && doc.CdtTrfTxInf[0].PmtId.InstrId != nil {
					id = string(*doc.CdtTrfTxInf[0].PmtId.InstrId)
				}
				if doc.CdtTrfTxInf[0].InstdAgt != nil {
					receiver = p.converter.ConvertFromPacs00900109(doc.CdtTrfTxInf[0].InstdAgt).ToParty()
				}
			}
			if (sender == nil || reflect.DeepEqual(sender, emptyParty)) && len(doc.CdtTrfTxInf) > 0 {
				if doc.CdtTrfTxInf[0].InstdAgt != nil {
					sender = p.converter.ConvertFromPacs00900109(doc.CdtTrfTxInf[0].InstgAgt).ToParty()
				}
			}
			if len(doc.CdtTrfTxInf) > 0 {
				endToEndID = string(doc.CdtTrfTxInf[0].PmtId.EndToEndId)
				if doc.CdtTrfTxInf[0].PmtId.UETR != nil {
					uetr = string(*doc.CdtTrfTxInf[0].PmtId.UETR)
				}
				if doc.CdtTrfTxInf[0].PmtId.TxId != nil {
					txID = string(*doc.CdtTrfTxInf[0].PmtId.TxId)
				}
			}
		case *pacs_007_001_10.FIToFIPaymentReversalV10:
			id = string(doc.GrpHdr.MsgId)
			receiver = p.converter.ConvertFromPacs00700110(doc.GrpHdr.InstdAgt).ToParty()
			sender = p.converter.ConvertFromPacs00700110(doc.GrpHdr.InstgAgt).ToParty()
			if (receiver == nil || reflect.DeepEqual(receiver, emptyParty)) && len(doc.TxInf) > 0 {
				if id == "" && doc.TxInf[0].RvslId != nil {
					id = string(*doc.TxInf[0].RvslId)
				}
				receiver = p.converter.ConvertFromPacs00700110(doc.TxInf[0].InstdAgt).ToParty()
			}
			if (sender == nil || reflect.DeepEqual(sender, emptyParty)) && len(doc.TxInf) > 0 {
				sender = p.converter.ConvertFromPacs00700110(doc.TxInf[0].InstgAgt).ToParty()
			}
			if len(doc.OrgnlGrpInf.OrgnlMsgId) > 0 {
				origId = string(doc.OrgnlGrpInf.OrgnlMsgId)
			}
			if len(doc.TxInf) > 0 {
				if origId == "" {
					if doc.TxInf[0].OrgnlGrpInf != nil {
						origId = string(doc.TxInf[0].OrgnlGrpInf.OrgnlMsgId)
					}
					if origId == "" && doc.TxInf[0].OrgnlInstrId != nil {
						origId = string(*doc.TxInf[0].OrgnlInstrId)
					}
				}
				if doc.TxInf[0].OrgnlEndToEndId != nil {
					origEndToEndID = string(*doc.TxInf[0].OrgnlEndToEndId)
				}
				if doc.TxInf[0].OrgnlTxId != nil {
					origTxID = string(*doc.TxInf[0].OrgnlTxId)
				}
			}
			origSender = receiver
			origReceiver = sender
		case *pacs_002_001_10.FIToFIPaymentStatusReportV10:
			id = string(doc.GrpHdr.MsgId)
			receiver = p.converter.ConvertFromPacs00200110(doc.GrpHdr.InstdAgt).ToParty()
			sender = p.converter.ConvertFromPacs00200110(doc.GrpHdr.InstgAgt).ToParty()
			if (receiver == nil || reflect.DeepEqual(receiver, emptyParty)) && len(doc.TxInfAndSts) > 0 {
				if id == "" && doc.TxInfAndSts[0].StsId != nil {
					id = string(*doc.TxInfAndSts[0].StsId)
				}
				receiver = p.converter.ConvertFromPacs00200110(doc.TxInfAndSts[0].InstdAgt).ToParty()
			}
			if (sender == nil || reflect.DeepEqual(sender, emptyParty)) && len(doc.TxInfAndSts) > 0 {
				sender = p.converter.ConvertFromPacs00200110(doc.TxInfAndSts[0].InstgAgt).ToParty()
			}
			if len(doc.OrgnlGrpInfAndSts) > 0 {
				origId = string(doc.OrgnlGrpInfAndSts[0].OrgnlMsgId)
			}
			if len(doc.TxInfAndSts) > 0 {
				if origId == "" {
					if doc.TxInfAndSts[0].OrgnlGrpInf != nil {
						origId = string(doc.TxInfAndSts[0].OrgnlGrpInf.OrgnlMsgId)
					}
					if origId == "" && doc.TxInfAndSts[0].OrgnlInstrId != nil {
						origId = string(*doc.TxInfAndSts[0].OrgnlInstrId)
					}
				}
				if doc.TxInfAndSts[0].OrgnlEndToEndId != nil {
					origEndToEndID = string(*doc.TxInfAndSts[0].OrgnlEndToEndId)
				}
				if doc.TxInfAndSts[0].OrgnlTxId != nil {
					origTxID = string(*doc.TxInfAndSts[0].OrgnlTxId)
				}
			}
			origSender = receiver
			origReceiver = sender
		case *pacs_004_001_10.PaymentReturnV10:
			id = string(doc.GrpHdr.MsgId)
			receiver = p.converter.ConvertFromPacs00400110(doc.GrpHdr.InstdAgt).ToParty()
			sender = p.converter.ConvertFromPacs00400110(doc.GrpHdr.InstgAgt).ToParty()
			if (receiver == nil || reflect.DeepEqual(receiver, emptyParty)) && len(doc.TxInf) > 0 {
				if id == "" && doc.TxInf[0].RtrId != nil {
					id = string(*doc.TxInf[0].RtrId)
				}
				receiver = p.converter.ConvertFromPacs00400110(doc.TxInf[0].InstdAgt).ToParty()
			}
			if (sender == nil || reflect.DeepEqual(sender, emptyParty)) && len(doc.TxInf) > 0 {
				sender = p.converter.ConvertFromPacs00400110(doc.TxInf[0].InstgAgt).ToParty()
			}
			if len(doc.OrgnlGrpInf.OrgnlMsgId) > 0 {
				origId = string(doc.OrgnlGrpInf.OrgnlMsgId)
			}
			if len(doc.TxInf) > 0 {
				if origId == "" {
					if doc.TxInf[0].OrgnlGrpInf != nil {
						origId = string(doc.TxInf[0].OrgnlGrpInf.OrgnlMsgId)
					}
					if origId == "" && doc.TxInf[0].OrgnlInstrId != nil {
						origId = string(*doc.TxInf[0].OrgnlInstrId)
					}
				}
				if doc.TxInf[0].OrgnlEndToEndId != nil {
					origEndToEndID = string(*doc.TxInf[0].OrgnlEndToEndId)
				}
				if doc.TxInf[0].OrgnlTxId != nil {
					origTxID = string(*doc.TxInf[0].OrgnlTxId)
				}
			}
			origSender = receiver
			origReceiver = sender
		case *pacs_002_001_11.FIToFIPaymentStatusReportV11:
			id = string(doc.GrpHdr.MsgId)
			receiver = p.converter.ConvertFromPacs00200111(doc.GrpHdr.InstdAgt).ToParty()
			sender = p.converter.ConvertFromPacs00200111(doc.GrpHdr.InstgAgt).ToParty()
			if (receiver == nil || reflect.DeepEqual(receiver, emptyParty)) && len(doc.TxInfAndSts) > 0 {
				if id == "" && doc.TxInfAndSts[0].StsId != nil {
					id = string(*doc.TxInfAndSts[0].StsId)
				}
				receiver = p.converter.ConvertFromPacs00200111(doc.TxInfAndSts[0].InstdAgt).ToParty()
			}
			if (sender == nil || reflect.DeepEqual(sender, emptyParty)) && len(doc.TxInfAndSts) > 0 {
				sender = p.converter.ConvertFromPacs00200111(doc.TxInfAndSts[0].InstgAgt).ToParty()
			}
			if len(doc.OrgnlGrpInfAndSts) > 0 {
				origId = string(doc.OrgnlGrpInfAndSts[0].OrgnlMsgId)
			}
			if len(doc.TxInfAndSts) > 0 {
				if origId == "" {
					if doc.TxInfAndSts[0].OrgnlGrpInf != nil {
						origId = string(doc.TxInfAndSts[0].OrgnlGrpInf.OrgnlMsgId)
					}
					if origId == "" && doc.TxInfAndSts[0].OrgnlInstrId != nil {
						origId = string(*doc.TxInfAndSts[0].OrgnlInstrId)
					}
				}
				if doc.TxInfAndSts[0].OrgnlEndToEndId != nil {
					origEndToEndID = string(*doc.TxInfAndSts[0].OrgnlEndToEndId)
				}
				if doc.TxInfAndSts[0].OrgnlTxId != nil {
					origTxID = string(*doc.TxInfAndSts[0].OrgnlTxId)
				}
			}
			origSender = receiver
			origReceiver = sender
		case *pacs_008_001_12.FIToFICustomerCreditTransferV12:
			id = string(doc.GrpHdr.MsgId)
			receiver = p.converter.ConvertFromPacs00800112(doc.GrpHdr.InstdAgt).ToParty()
			sender = p.converter.ConvertFromPacs00800112(doc.GrpHdr.InstgAgt).ToParty()
			if (receiver == nil || reflect.DeepEqual(receiver, emptyParty)) && len(doc.CdtTrfTxInf) > 0 {
				if id == "" && doc.CdtTrfTxInf[0].PmtId.InstrId != nil {
					id = string(*doc.CdtTrfTxInf[0].PmtId.InstrId)
				}
				if doc.CdtTrfTxInf[0].InstdAgt != nil {
					receiver = p.converter.ConvertFromPacs00800112(doc.CdtTrfTxInf[0].InstdAgt).ToParty()
				}
			}
			if (sender == nil || reflect.DeepEqual(sender, emptyParty)) && len(doc.CdtTrfTxInf) > 0 {
				if doc.CdtTrfTxInf[0].InstdAgt != nil {
					sender = p.converter.ConvertFromPacs00800112(doc.CdtTrfTxInf[0].InstgAgt).ToParty()
				}
			}
			if len(doc.CdtTrfTxInf) > 0 {
				endToEndID = string(doc.CdtTrfTxInf[0].PmtId.EndToEndId)
				if doc.CdtTrfTxInf[0].PmtId.UETR != nil {
					uetr = string(*doc.CdtTrfTxInf[0].PmtId.UETR)
				}
				if doc.CdtTrfTxInf[0].PmtId.TxId != nil {
					txID = string(*doc.CdtTrfTxInf[0].PmtId.TxId)
				}
			}
		default:
			return containedDoc, metadata, references, supplementaryDataParser, errors.New("couldn't find receiver from " + reflect.TypeOf(containedDoc).String())
		}
	}

	if origId != "" {
		references = &processes.Metadata{
			Uetr:     MakeUETR(p.log, origUetr, origEndToEndID, origTxID),
			ID:       origId,
			Sender:   origSender,
			Receiver: origReceiver,
		}
	}

	if receiver == nil || reflect.DeepEqual(receiver, emptyParty) {
		return containedDoc, metadata, references, supplementaryDataParser, errors.New("couldn't find receiver")
	}

	metadata.Uetr = MakeUETR(p.log, uetr, endToEndID, txID)
	metadata.ID = id
	metadata.Sender = sender
	metadata.Receiver = receiver

	if headDoc != nil {
		switch head := headDoc.(type) {
		case *head_001_001_01.BusinessApplicationHeaderV01:
			if metadata.ID != "" {
				metadata.ID = string(head.BizMsgIdr)
			}
			if metadata.Receiver == nil || reflect.DeepEqual(metadata.Receiver, emptyParty) {
				metadata.Receiver = p.converter.ConvertFromHead00100101(head.To.FIId).ToParty()
			}
			if metadata.Sender == nil || reflect.DeepEqual(metadata.Sender, emptyParty) {
				metadata.Sender = p.converter.ConvertFromHead00100101(head.Fr.FIId).ToParty()
			}
			return message, metadata, references, supplementaryDataParser, nil
		case *head_001_001_02.BusinessApplicationHeaderV02:
			if metadata.ID != "" {
				metadata.ID = string(head.BizMsgIdr)
			}
			if metadata.Receiver == nil || reflect.DeepEqual(metadata.Receiver, emptyParty) {
				metadata.Receiver = p.converter.ConvertFromHead00100102(head.To.FIId).ToParty()
			}
			if metadata.Sender == nil || reflect.DeepEqual(metadata.Sender, emptyParty) {
				metadata.Sender = p.converter.ConvertFromHead00100102(head.Fr.FIId).ToParty()
			}
			return containedDoc, metadata, references, supplementaryDataParser, nil
		}
	}

	return containedDoc, metadata, references, supplementaryDataParser, nil
}

func (p Parser) GetTransactionStatus(isoMsg messages.Iso20022Message) processes.TransactionStatus {
	switch msg := isoMsg.(type) {
	case *pacs_002_001_07.FIToFIPaymentStatusReportV07:
		if status := msg.TxInfAndSts[0].TxSts; status != nil {
			return processes.ParseTransactionStatus(string(*status))
		}
	case *pacs_002_001_08.FIToFIPaymentStatusReportV08:
		if status := msg.TxInfAndSts[0].TxSts; status != nil {
			return processes.ParseTransactionStatus(string(*status))
		}
	case *pacs_002_001_10.FIToFIPaymentStatusReportV10:
		if status := msg.TxInfAndSts[0].TxSts; status != nil {
			return processes.ParseTransactionStatus(string(*status))
		}
	case *pacs_002_001_11.FIToFIPaymentStatusReportV11:
		if status := msg.TxInfAndSts[0].TxSts; status != nil {
			return processes.ParseTransactionStatus(string(*status))
		}
	}
	return processes.TransactionStatusNone
}

func MakeUETR(log logger.Logger, uetr, endToEndID, txID string) string {
	if uetr != "" {
		return uetr
	}
	if endToEndID != "" {
		uetr = stringToUUID(log, endToEndID)
		if uetr != "" {
			return uetr
		}
	}
	if txID != "" {
		uetr = stringToUUID(log, txID)
		if uetr != "" {
			return uetr
		}
	}
	return ""
}

func stringToUUID(log logger.Logger, text string) string {
	h := md5.New()
	h.Write([]byte(text))
	md5Hash := h.Sum(nil)
	e2eUUID, err := uuid.FromBytes(md5Hash)
	if err != nil {
		log.Warn(
			context.TODO(),
			"could not generate uuid from md5",
			zap.Binary("md5", md5Hash), zap.Error(err),
		)
		return ""
	} else {
		return e2eUUID.String()
	}
}
