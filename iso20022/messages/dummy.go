package messages

import (
	"encoding/xml"

	"github.com/pkg/errors"
)

type elementDummy struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:",any,attr,omitempty" json:",omitempty"`
	Rest    []byte     `xml:",innerxml"`
}

func (dummy elementDummy) NameSpace() string {
	for _, attr := range dummy.Attrs {
		if attr.Name.Local == XmlDefaultNamespace {
			return attr.Value
		}
	}
	return ""
}

func (dummy elementDummy) Validate() error {
	if len(dummy.NameSpace()) == 0 {
		return Validate(&dummy)
	}

	for _, attr := range dummy.Attrs {
		if attr.Name.Local == XmlDefaultNamespace && dummy.NameSpace() == attr.Value {
			return Validate(&dummy)
		}
	}

	return errors.New("The namespace of document is invalid")
}

type documentDummy struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr,omitempty" json:",omitempty"`
	AppHdr   *elementDummy
	Document *elementDummy `xml:",any"`
}

func (dummy documentDummy) NameSpace() string {
	for _, attr := range dummy.Attrs {
		if attr.Name.Local == XmlDefaultNamespace {
			return attr.Value
		}
	}
	return ""
}

func (dummy documentDummy) Validate() error {
	if len(dummy.NameSpace()) == 0 {
		return Validate(&dummy)
	}

	for _, attr := range dummy.Attrs {
		if attr.Name.Local == XmlDefaultNamespace && dummy.NameSpace() == attr.Value {
			return Validate(&dummy)
		}
	}

	return errors.New("The namespace of document is invalid")
}
