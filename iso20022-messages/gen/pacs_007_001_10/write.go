// Code generated by GoComply XSD2Go for iso20022; DO NOT EDIT.
// Models for urn:iso:std:iso:20022:tech:xsd:pacs.007.001.10
package pacs_007_001_10

import (
	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/pkg/iso"
)

func (a ActiveCurrencyAndAmountSimpleType) MarshalText() ([]byte, error) {
	return iso.Amount(a).MarshalText()
}

func (a ActiveOrHistoricCurrencyAndAmountSimpleType) MarshalText() ([]byte, error) {
	return iso.Amount(a).MarshalText()
}
