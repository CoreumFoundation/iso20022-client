package iso_test

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/pkg/iso"
)

var cdataText = "<some>XML</some>"

type TestStruct struct {
	AddtlData *iso.Cdata `xml:"AddtlData,omitempty"`
}

type TestStructNamespace struct {
	AddtlData *iso.Cdata `xml:"mr:AddtlData,omitempty"`
}

func TestCdata_write(t *testing.T) {
	input := TestStruct{AddtlData: &iso.Cdata{CDataString: cdataText}}

	output, err := xml.Marshal(input)
	require.NoError(t, err)

	assert.Equal(t, "<TestStruct><AddtlData><![CDATA[<some>XML</some>]]></AddtlData></TestStruct>", string(output))
}

func TestCdata_write_withNamespace(t *testing.T) {
	input := TestStructNamespace{AddtlData: &iso.Cdata{CDataString: cdataText}}

	output, err := xml.Marshal(input)
	require.NoError(t, err)

	assert.Equal(t, "<TestStructNamespace><mr:AddtlData><![CDATA[<some>XML</some>]]></mr:AddtlData></TestStructNamespace>", string(output))
}

func TestCdata_read(t *testing.T) {
	input := "<TestStruct><AddtlData><![CDATA[<some>XML</some>]]></AddtlData></TestStruct>"
	output := &TestStruct{}

	err := xml.Unmarshal([]byte(input), &output)
	require.NoError(t, err)

	assert.Equal(t, &TestStruct{AddtlData: &iso.Cdata{CDataString: cdataText}}, output)
}

func TestCdata_read_withNamespace(t *testing.T) {
	input := "<TestStructNamespace><mr:AddtlData><![CDATA[<some>XML</some>]]></mr:AddtlData></TestStructNamespace>"
	// It sounds counter-intuitive, but go will not match the mr:AddtlData tag with the mr:AddtlData defined structure when
	// reading, but it will match to the AddtlData defined structure. It also doesn't care about the root tag name
	// (TestStruct from the defined struct vs. TestStructNamespace from the XML).
	output := &TestStruct{}

	err := xml.Unmarshal([]byte(input), output)
	require.NoError(t, err)

	assert.Equal(t, &TestStruct{AddtlData: &iso.Cdata{CDataString: cdataText}}, output)
}
