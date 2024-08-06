package messages

import (
	"bytes"
	"encoding/xml"
	"reflect"

	"github.com/pkg/errors"

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

func (p Parser) parseIsoMessage(msg []byte) (header, doc messages.Iso20022Message, err error) {
	dummyDoc := new(documentDummy)

	err = xml.Unmarshal(msg, dummyDoc)
	if err != nil {
		return nil, nil, err
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
		return nil, actualDoc.Message, err
	}

	var headerConstructor func() messages.Iso20022Message

	headerNamespace := ""
	if dummyDoc.AppHdr != nil {
		headerNamespace = dummyDoc.AppHdr.NameSpace()
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
				return nil, nil, errors.New("The namespace of document is omitted")
			}
		}

		headerConstructor = messageConstructor[headerNamespace]
		if headerConstructor == nil {
			return nil, nil, errors.New("The namespace of document is unsupported")
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
			return nil, nil, err
		}
		for _, attr := range dummyDoc.Attrs {
			if attr.Name.Local == innerDoc.XMLName.Space {
				documentNamespace = attr.Value
				break
			}
		}
		if documentNamespace == "" {
			return nil, nil, errors.New("The namespace of document is omitted")
		}
		documentConstructor := messageConstructor[documentNamespace]
		if documentConstructor == nil {
			return nil, nil, errors.New("The namespace of document is unsupported")
		}
		containedDoc = documentConstructor()
	} else {
		constructor = messageConstructor[documentNamespace]
		if constructor == nil {
			return nil, nil, errors.New("The namespace of document is unsupported")
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
		return nil, nil, err
	}

	if envelope.AppHdr != nil {
		err = envelope.AppHdr.Validate()
		if err != nil {
			return nil, nil, err
		}
	}

	var resDoc messages.Iso20022Message

	if innerDoc, ok := envelope.Document.(*Iso20022DocumentObject); ok {
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

func (p Parser) ExtractMetadataFromIsoMessage(msg []byte) (data processes.Metadata, err error) {
	headDoc, containedDoc, err := p.parseIsoMessage(msg)
	if err != nil {
		return data, err
	}

	id := ""
	sender := new(addressbook.Party)
	receiver := new(addressbook.Party)
	emptyParty := new(addressbook.Party)

	if headDoc != nil {
		switch head := headDoc.(type) {
		case *head_001_001_01.BusinessApplicationHeaderV01:
			data.ID = string(head.BizMsgIdr)
			data.Receiver = p.converter.ConvertFromHead00100101(head.To.FIId).ToParty()
			data.Sender = p.converter.ConvertFromHead00100101(head.Fr.FIId).ToParty()
			return data, nil
		case *head_001_001_02.BusinessApplicationHeaderV02:
			data.ID = string(head.BizMsgIdr)
			data.Receiver = p.converter.ConvertFromHead00100102(head.To.FIId).ToParty()
			data.Sender = p.converter.ConvertFromHead00100102(head.Fr.FIId).ToParty()
			return data, nil
		}
	}

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
		default:
			return data, errors.New("couldn't find receiver from " + reflect.TypeOf(containedDoc).String())
		}
	}

	if receiver == nil || reflect.DeepEqual(receiver, emptyParty) {
		return data, errors.New("couldn't find receiver")
	}

	data.ID = id
	data.Sender = sender
	data.Receiver = receiver

	return data, nil
}
