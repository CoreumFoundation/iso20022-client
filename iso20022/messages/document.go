package messages

import (
	"encoding/xml"
	"errors"

	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/gen/messages"
)

const XmlDefaultNamespace = "xmlns"

type Iso20022DocumentObject struct {
	XMLName xml.Name
	Attrs   []xml.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
	Message messages.Iso20022Message `xml:",any"`
}

func (doc Iso20022DocumentObject) Validate() error {
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

func (doc Iso20022DocumentObject) NameSpace() string {
	for _, attr := range doc.Attrs {
		if attr.Name.Local == XmlDefaultNamespace {
			return attr.Value
		}
	}
	return ""
}

func (doc *Iso20022DocumentObject) GetXmlName() *xml.Name {
	return &doc.XMLName
}

func (doc *Iso20022DocumentObject) GetAttrs() []xml.Attr {
	return doc.Attrs
}

func (doc *Iso20022DocumentObject) InspectMessage() messages.Iso20022Message {
	return doc.Message
}

func (doc Iso20022DocumentObject) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	a := struct {
		XMLName xml.Name
		Attrs   []xml.Attr               `xml:",any,attr,omitempty" json:",omitempty"`
		Message messages.Iso20022Message `xml:",any"`
	}(doc)

	updatingStartElement(&start, doc.Attrs, doc.XMLName)
	return e.EncodeElement(&a, start)
}

func updatingStartElement(start *xml.StartElement, attrs []xml.Attr, name xml.Name) {
	for _, attr := range attrs {
		if attr.Name.Local == XmlDefaultNamespace {
			name.Space = ""
		}
	}
	if len(name.Local) > 0 {
		start.Name.Local = name.Local
	}
	start.Name.Space = name.Space
}
