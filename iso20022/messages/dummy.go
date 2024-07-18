package messages

import (
	"encoding/xml"

	"github.com/moov-io/iso20022/pkg/utils"
)

type elementDummy struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:",any,attr,omitempty" json:",omitempty"`
}

func (dummy elementDummy) NameSpace() string {
	for _, attr := range dummy.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			return attr.Value
		}
	}
	return ""
}

func (dummy elementDummy) Validate() error {
	if len(dummy.NameSpace()) == 0 {
		return utils.Validate(&dummy)
	}

	for _, attr := range dummy.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && dummy.NameSpace() == attr.Value {
			return utils.Validate(&dummy)
		}
	}

	return utils.NewErrInvalidNameSpace()
}

type documentDummy struct {
	XMLName  xml.Name
	Attrs    []xml.Attr `xml:",any,attr,omitempty" json:",omitempty"`
	AppHdr   *elementDummy
	Document *elementDummy
}

func (dummy documentDummy) NameSpace() string {
	for _, attr := range dummy.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace {
			return attr.Value
		}
	}
	return ""
}

func (dummy documentDummy) Validate() error {
	if len(dummy.NameSpace()) == 0 {
		return utils.Validate(&dummy)
	}

	for _, attr := range dummy.Attrs {
		if attr.Name.Local == utils.XmlDefaultNamespace && dummy.NameSpace() == attr.Value {
			return utils.Validate(&dummy)
		}
	}

	return utils.NewErrInvalidNameSpace()
}
