package dtif

import (
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"

	"github.com/samber/lo"
)

// FunctionallyFungibleGroupOfDigitalTokensJson Some or all of the digital tokens
// within a functionally fungible group digital tokens may be technically distinct;
// Software updates may render digital token(s) obsolete resulting in the creation
// of new digital tokens. Typically, either ownership of the new reflects equivalent
// value of the old, or a method to convert old digital tokens to new digital tokens
// exists. The old and the new digital tokens of the same set are technically incompatible.
// The set of these digital tokens may also be defined as a functionally fungible group
type FunctionallyFungibleGroupOfDigitalTokensJson struct {
	// Header corresponds to the JSON schema field "Header".
	Header FunctionallyFungibleGroupOfDigitalTokensJsonHeader `json:"Header" yaml:"Header" mapstructure:"Header"`

	// Informative corresponds to the JSON schema field "Informative".
	Informative FunctionallyFungibleGroupOfDigitalTokensJsonInformative `json:"Informative" yaml:"Informative" mapstructure:"Informative"`

	// Metadata corresponds to the JSON schema field "Metadata".
	Metadata FunctionallyFungibleGroupOfDigitalTokensJsonMetadata `json:"Metadata" yaml:"Metadata" mapstructure:"Metadata"`

	// Set of digital tokens which are functionally capable of mutual substitution
	// between the individual units of digital assets
	Normative FunctionallyFungibleGroupOfDigitalTokensJsonNormative `json:"Normative" yaml:"Normative" mapstructure:"Normative"`
}

type FunctionallyFungibleGroupOfDigitalTokensJsonHeader struct {
	// DTI corresponds to the JSON schema field "DTI".
	DTI string `json:"DTI" yaml:"DTI" mapstructure:"DTI"`

	// Functionally Fungible group of Digital TokensFunctionally Fungible group of
	// Digital Tokens
	DTIType FunctionallyFungibleGroupOfDigitalTokensJsonHeaderDTIType `json:"DTIType" yaml:"DTIType" mapstructure:"DTIType"`

	// The template version (i.e. JSON schema document) for which this record is valid
	TemplateVersion FunctionallyFungibleGroupOfDigitalTokensJsonHeaderTemplateVersion `json:"templateVersion" yaml:"templateVersion" mapstructure:"templateVersion"`
}

type FunctionallyFungibleGroupOfDigitalTokensJsonHeaderDTIType int

var enumValues_FunctionallyFungibleGroupOfDigitalTokensJsonHeaderDTIType = []interface{}{
	0,
	1,
	2,
	3,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *FunctionallyFungibleGroupOfDigitalTokensJsonHeaderDTIType) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_FunctionallyFungibleGroupOfDigitalTokensJsonHeaderDTIType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_FunctionallyFungibleGroupOfDigitalTokensJsonHeaderDTIType, v)
	}
	*j = FunctionallyFungibleGroupOfDigitalTokensJsonHeaderDTIType(v)
	return nil
}

type FunctionallyFungibleGroupOfDigitalTokensJsonHeaderTemplateVersion string

const FunctionallyFungibleGroupOfDigitalTokensJsonHeaderTemplateVersionV100 FunctionallyFungibleGroupOfDigitalTokensJsonHeaderTemplateVersion = "V1.0.0"

var enumValues_FunctionallyFungibleGroupOfDigitalTokensJsonHeaderTemplateVersion = []interface{}{
	"V1.0.0",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *FunctionallyFungibleGroupOfDigitalTokensJsonHeaderTemplateVersion) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_FunctionallyFungibleGroupOfDigitalTokensJsonHeaderTemplateVersion {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_FunctionallyFungibleGroupOfDigitalTokensJsonHeaderTemplateVersion, v)
	}
	*j = FunctionallyFungibleGroupOfDigitalTokensJsonHeaderTemplateVersion(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *FunctionallyFungibleGroupOfDigitalTokensJsonHeader) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["DTI"]; raw != nil && !ok {
		return fmt.Errorf("field DTI in FunctionallyFungibleGroupOfDigitalTokensJsonHeader: required")
	}
	if _, ok := raw["DTIType"]; raw != nil && !ok {
		return fmt.Errorf("field DTIType in FunctionallyFungibleGroupOfDigitalTokensJsonHeader: required")
	}
	if _, ok := raw["templateVersion"]; raw != nil && !ok {
		return fmt.Errorf("field templateVersion in FunctionallyFungibleGroupOfDigitalTokensJsonHeader: required")
	}
	type Plain FunctionallyFungibleGroupOfDigitalTokensJsonHeader
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = FunctionallyFungibleGroupOfDigitalTokensJsonHeader(plain)
	return nil
}

type FunctionallyFungibleGroupOfDigitalTokensJsonInformative struct {
	// DTIExternalIdentifiers corresponds to the JSON schema field
	// "DTIExternalIdentifiers".
	DTIExternalIdentifiers []FunctionallyFungibleGroupOfDigitalTokensJsonInformativeDTIExternalIdentifiersElem `json:"DTIExternalIdentifiers,omitempty" yaml:"DTIExternalIdentifiers,omitempty" mapstructure:"DTIExternalIdentifiers,omitempty"`

	// Digital token long name is a string containing the full name of the digital
	// token
	LongName *string `json:"LongName,omitempty" yaml:"LongName,omitempty" mapstructure:"LongName,omitempty"`

	// OrigLangLongName corresponds to the JSON schema field "OrigLangLongName".
	OrigLangLongName string `json:"OrigLangLongName,omitempty" yaml:"OrigLangLongName,omitempty" mapstructure:"OrigLangLongName,omitempty"`

	// ShortNames corresponds to the JSON schema field "ShortNames".
	ShortNames []FunctionallyFungibleGroupOfDigitalTokensJsonInformativeShortNamesElem `json:"ShortNames,omitempty" yaml:"ShortNames,omitempty" mapstructure:"ShortNames,omitempty"`

	// UnderlyingAssetExternalIdentifiers corresponds to the JSON schema field
	// "UnderlyingAssetExternalIdentifiers".
	UnderlyingAssetExternalIdentifiers []FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElem `json:"UnderlyingAssetExternalIdentifiers,omitempty" yaml:"UnderlyingAssetExternalIdentifiers,omitempty" mapstructure:"UnderlyingAssetExternalIdentifiers,omitempty"`
}

type FunctionallyFungibleGroupOfDigitalTokensJsonInformativeDTIExternalIdentifiersElem struct {
	// External identifier type for the digital token
	DTIExternalIdentifierType FunctionallyFungibleGroupOfDigitalTokensJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType `json:"DTIExternalIdentifierType" yaml:"DTIExternalIdentifierType" mapstructure:"DTIExternalIdentifierType"`

	// External identifier for the digital token
	ExternalIdentifierValue string `json:"ExternalIdentifierValue" yaml:"ExternalIdentifierValue" mapstructure:"ExternalIdentifierValue"`
}

type FunctionallyFungibleGroupOfDigitalTokensJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType string

const FunctionallyFungibleGroupOfDigitalTokensJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierTypeITIN FunctionallyFungibleGroupOfDigitalTokensJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType = "ITIN"

var enumValues_FunctionallyFungibleGroupOfDigitalTokensJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType = []interface{}{
	"ITIN",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *FunctionallyFungibleGroupOfDigitalTokensJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_FunctionallyFungibleGroupOfDigitalTokensJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_FunctionallyFungibleGroupOfDigitalTokensJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType, v)
	}
	*j = FunctionallyFungibleGroupOfDigitalTokensJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *FunctionallyFungibleGroupOfDigitalTokensJsonInformativeDTIExternalIdentifiersElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["DTIExternalIdentifierType"]; raw != nil && !ok {
		return fmt.Errorf("field DTIExternalIdentifierType in FunctionallyFungibleGroupOfDigitalTokensJsonInformativeDTIExternalIdentifiersElem: required")
	}
	if _, ok := raw["ExternalIdentifierValue"]; raw != nil && !ok {
		return fmt.Errorf("field ExternalIdentifierValue in FunctionallyFungibleGroupOfDigitalTokensJsonInformativeDTIExternalIdentifiersElem: required")
	}
	type Plain FunctionallyFungibleGroupOfDigitalTokensJsonInformativeDTIExternalIdentifiersElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = FunctionallyFungibleGroupOfDigitalTokensJsonInformativeDTIExternalIdentifiersElem(plain)
	return nil
}

type FunctionallyFungibleGroupOfDigitalTokensJsonInformativeShortNamesElem struct {
	// OrigLangShortName corresponds to the JSON schema field "OrigLangShortName".
	OrigLangShortName string `json:"OrigLangShortName,omitempty" yaml:"OrigLangShortName,omitempty" mapstructure:"OrigLangShortName,omitempty"`

	// Short name or ticker symbol of the digital token
	ShortName string `json:"ShortName" yaml:"ShortName" mapstructure:"ShortName"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *FunctionallyFungibleGroupOfDigitalTokensJsonInformativeShortNamesElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["ShortName"]; raw != nil && !ok {
		return fmt.Errorf("field ShortName in FunctionallyFungibleGroupOfDigitalTokensJsonInformativeShortNamesElem: required")
	}
	type Plain FunctionallyFungibleGroupOfDigitalTokensJsonInformativeShortNamesElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["OrigLangShortName"]; !ok || v == nil {
		plain.OrigLangShortName = "Basic Latin (Unicode)"
	}
	*j = FunctionallyFungibleGroupOfDigitalTokensJsonInformativeShortNamesElem(plain)
	return nil
}

type FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElem struct {
	// Underlying asset external identifier type for the digital token
	UnderlyingAssetExternalIdentifierType FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType `json:"UnderlyingAssetExternalIdentifierType" yaml:"UnderlyingAssetExternalIdentifierType" mapstructure:"UnderlyingAssetExternalIdentifierType"`

	// Underlying asset external identifier value for the digital token
	UnderlyingAssetExternalIdentifierValue string `json:"UnderlyingAssetExternalIdentifierValue" yaml:"UnderlyingAssetExternalIdentifierValue" mapstructure:"UnderlyingAssetExternalIdentifierValue"`
}

type FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType string

const FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeCCY FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "CCY"
const FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeCountryCode FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "Country code"
const FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeCusip FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "Cusip"
const FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeFIGI FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "FIGI"
const FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeISIN FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "ISIN"
const FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeLEI FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "LEI"
const FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeRIC FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "RIC"
const FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeSEDOL FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "SEDOL"

var enumValues_FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = []interface{}{
	"ISIN",
	"Cusip",
	"SEDOL",
	"RIC",
	"CCY",
	"Country code",
	"FIGI",
	"LEI",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType, v)
	}
	*j = FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["UnderlyingAssetExternalIdentifierType"]; raw != nil && !ok {
		return fmt.Errorf("field UnderlyingAssetExternalIdentifierType in FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElem: required")
	}
	if _, ok := raw["UnderlyingAssetExternalIdentifierValue"]; raw != nil && !ok {
		return fmt.Errorf("field UnderlyingAssetExternalIdentifierValue in FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElem: required")
	}
	type Plain FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = FunctionallyFungibleGroupOfDigitalTokensJsonInformativeUnderlyingAssetExternalIdentifiersElem(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *FunctionallyFungibleGroupOfDigitalTokensJsonInformative) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain FunctionallyFungibleGroupOfDigitalTokensJsonInformative
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["OrigLangLongName"]; !ok || v == nil {
		plain.OrigLangLongName = "Basic Latin (Unicode)"
	}
	*j = FunctionallyFungibleGroupOfDigitalTokensJsonInformative(plain)
	return nil
}

type FunctionallyFungibleGroupOfDigitalTokensJsonMetadata struct {
	// Record is marked Deleted due to technical reason(s) (e.g. invalid or duplicate
	// token) or due to business reason(s) (e.g. a provisional token becomes invalid).
	// Users are advised not to use this token.
	Deleted bool `json:"Deleted,omitempty" yaml:"Deleted,omitempty" mapstructure:"Deleted,omitempty"`

	// If true, the record contains normative or informative data elements that are
	// disputed by one or more DTIF users
	Disputed bool `json:"Disputed,omitempty" yaml:"Disputed,omitempty" mapstructure:"Disputed,omitempty"`

	// If true, the record contains normative or informative data elements that are
	// verified according to the DTIF guidelines but may not be independently
	// verifiable by the general public
	Private bool `json:"Private,omitempty" yaml:"Private,omitempty" mapstructure:"Private,omitempty"`

	// If true, the record contains normative or informative data elements in a
	// provisional request for a DTI when one or more data elements (which may include
	// mandatory data elements) are not yet available
	Provisional bool `json:"Provisional,omitempty" yaml:"Provisional,omitempty" mapstructure:"Provisional,omitempty"`

	// The time at which this record version has been created
	RecDateTime JsonTime `json:"recDateTime" yaml:"recDateTime" mapstructure:"recDateTime"`

	// The version of the record, incremental integer starting at 1 and incrementing
	// each time the record is updated (e.g. with a new fork)
	RecVersion int `json:"recVersion" yaml:"recVersion" mapstructure:"recVersion"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *FunctionallyFungibleGroupOfDigitalTokensJsonMetadata) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["recDateTime"]; raw != nil && !ok {
		return fmt.Errorf("field recDateTime in FunctionallyFungibleGroupOfDigitalTokensJsonMetadata: required")
	}
	if _, ok := raw["recVersion"]; raw != nil && !ok {
		return fmt.Errorf("field recVersion in FunctionallyFungibleGroupOfDigitalTokensJsonMetadata: required")
	}
	type Plain FunctionallyFungibleGroupOfDigitalTokensJsonMetadata
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["Deleted"]; !ok || v == nil {
		plain.Deleted = false
	}
	if v, ok := raw["Disputed"]; !ok || v == nil {
		plain.Disputed = false
	}
	if v, ok := raw["Private"]; !ok || v == nil {
		plain.Private = false
	}
	if v, ok := raw["Provisional"]; !ok || v == nil {
		plain.Provisional = false
	}
	*j = FunctionallyFungibleGroupOfDigitalTokensJsonMetadata(plain)
	return nil
}

// Set of digital tokens which are functionally capable of mutual substitution
// between the individual units of digital assets
type FunctionallyFungibleGroupOfDigitalTokensJsonNormative struct {
	// FunctionallyFungibleDTI corresponds to the JSON schema field "Functionally
	// fungible DTI".
	FunctionallyFungibleDTI []string `json:"Functionally fungible DTI" yaml:"Functionally fungible DTI" mapstructure:"Functionally fungible DTI"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *FunctionallyFungibleGroupOfDigitalTokensJsonNormative) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["Functionally fungible DTI"]; raw != nil && !ok {
		return fmt.Errorf("field Functionally fungible DTI in FunctionallyFungibleGroupOfDigitalTokensJsonNormative: required")
	}
	type Plain FunctionallyFungibleGroupOfDigitalTokensJsonNormative
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = FunctionallyFungibleGroupOfDigitalTokensJsonNormative(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *FunctionallyFungibleGroupOfDigitalTokensJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["Header"]; raw != nil && !ok {
		return fmt.Errorf("field Header in FunctionallyFungibleGroupOfDigitalTokensJson: required")
	}
	if _, ok := raw["Informative"]; raw != nil && !ok {
		return fmt.Errorf("field Informative in FunctionallyFungibleGroupOfDigitalTokensJson: required")
	}
	if _, ok := raw["Metadata"]; raw != nil && !ok {
		return fmt.Errorf("field Metadata in FunctionallyFungibleGroupOfDigitalTokensJson: required")
	}
	if _, ok := raw["Normative"]; raw != nil && !ok {
		return fmt.Errorf("field Normative in FunctionallyFungibleGroupOfDigitalTokensJson: required")
	}
	type Plain FunctionallyFungibleGroupOfDigitalTokensJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = FunctionallyFungibleGroupOfDigitalTokensJson(plain)
	return nil
}

// DTI returns token DTI.
func (j *FunctionallyFungibleGroupOfDigitalTokensJson) DTI() string {
	return j.Header.DTI
}

// Denom returns token denom.
func (j *FunctionallyFungibleGroupOfDigitalTokensJson) Denom() *string {
	return lo.ToPtr("")
}

// PriceMultiplier returns token price multiplier.
func (j *FunctionallyFungibleGroupOfDigitalTokensJson) PriceMultiplier() *big.Int {
	return big.NewInt(1)
}
