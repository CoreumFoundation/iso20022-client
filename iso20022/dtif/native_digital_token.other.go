package dtif

import (
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"

	"github.com/samber/lo"
)

// NativeDigitalTokenOtherJson Native Digital token with non-blockchain distributed ledger technology protocol
type NativeDigitalTokenOtherJson struct {
	// Header corresponds to the JSON schema field "Header".
	Header NativeDigitalTokenOtherJsonHeader `json:"Header" yaml:"Header" mapstructure:"Header"`

	// Informative corresponds to the JSON schema field "Informative".
	Informative NativeDigitalTokenOtherJsonInformative `json:"Informative" yaml:"Informative" mapstructure:"Informative"`

	// Metadata corresponds to the JSON schema field "Metadata".
	Metadata NativeDigitalTokenOtherJsonMetadata `json:"Metadata" yaml:"Metadata" mapstructure:"Metadata"`
}

type NativeDigitalTokenOtherJsonHeader struct {
	// Other
	DLTType NativeDigitalTokenOtherJsonHeaderDLTType `json:"DLTType" yaml:"DLTType" mapstructure:"DLTType"`

	// DTI corresponds to the JSON schema field "DTI".
	DTI string `json:"DTI" yaml:"DTI" mapstructure:"DTI"`

	// Native Digital Token
	DTIType NativeDigitalTokenOtherJsonHeaderDTIType `json:"DTIType" yaml:"DTIType" mapstructure:"DTIType"`

	// The template version (i.e. JSON schema document) for which this record is valid
	TemplateVersion NativeDigitalTokenOtherJsonHeaderTemplateVersion `json:"templateVersion" yaml:"templateVersion" mapstructure:"templateVersion"`
}

type NativeDigitalTokenOtherJsonHeaderDLTType int

var enumValues_NativeDigitalTokenOtherJsonHeaderDLTType = []interface{}{
	1,
	0,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NativeDigitalTokenOtherJsonHeaderDLTType) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_NativeDigitalTokenOtherJsonHeaderDLTType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_NativeDigitalTokenOtherJsonHeaderDLTType, v)
	}
	*j = NativeDigitalTokenOtherJsonHeaderDLTType(v)
	return nil
}

type NativeDigitalTokenOtherJsonHeaderDTIType int

var enumValues_NativeDigitalTokenOtherJsonHeaderDTIType = []interface{}{
	0,
	1,
	2,
	3,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NativeDigitalTokenOtherJsonHeaderDTIType) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_NativeDigitalTokenOtherJsonHeaderDTIType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_NativeDigitalTokenOtherJsonHeaderDTIType, v)
	}
	*j = NativeDigitalTokenOtherJsonHeaderDTIType(v)
	return nil
}

type NativeDigitalTokenOtherJsonHeaderTemplateVersion string

const NativeDigitalTokenOtherJsonHeaderTemplateVersionV100 NativeDigitalTokenOtherJsonHeaderTemplateVersion = "V1.0.0"

var enumValues_NativeDigitalTokenOtherJsonHeaderTemplateVersion = []interface{}{
	"V1.0.0",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NativeDigitalTokenOtherJsonHeaderTemplateVersion) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_NativeDigitalTokenOtherJsonHeaderTemplateVersion {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_NativeDigitalTokenOtherJsonHeaderTemplateVersion, v)
	}
	*j = NativeDigitalTokenOtherJsonHeaderTemplateVersion(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NativeDigitalTokenOtherJsonHeader) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["DLTType"]; raw != nil && !ok {
		return fmt.Errorf("field DLTType in NativeDigitalTokenOtherJsonHeader: required")
	}
	if _, ok := raw["DTI"]; raw != nil && !ok {
		return fmt.Errorf("field DTI in NativeDigitalTokenOtherJsonHeader: required")
	}
	if _, ok := raw["DTIType"]; raw != nil && !ok {
		return fmt.Errorf("field DTIType in NativeDigitalTokenOtherJsonHeader: required")
	}
	if _, ok := raw["templateVersion"]; raw != nil && !ok {
		return fmt.Errorf("field templateVersion in NativeDigitalTokenOtherJsonHeader: required")
	}
	type Plain NativeDigitalTokenOtherJsonHeader
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = NativeDigitalTokenOtherJsonHeader(plain)
	return nil
}

type NativeDigitalTokenOtherJsonInformative struct {
	// DTIExternalIdentifiers corresponds to the JSON schema field
	// "DTIExternalIdentifiers".
	DTIExternalIdentifiers []NativeDigitalTokenOtherJsonInformativeDTIExternalIdentifiersElem `json:"DTIExternalIdentifiers,omitempty" yaml:"DTIExternalIdentifiers,omitempty" mapstructure:"DTIExternalIdentifiers,omitempty"`

	// Digital token long name is a string containing the full name of the digital
	// token
	LongName *string `json:"LongName,omitempty" yaml:"LongName,omitempty" mapstructure:"LongName,omitempty"`

	// OrigLangLongName corresponds to the JSON schema field "OrigLangLongName".
	OrigLangLongName string `json:"OrigLangLongName,omitempty" yaml:"OrigLangLongName,omitempty" mapstructure:"OrigLangLongName,omitempty"`

	// If true, access to reading the distributed ledger is unrestricted and the data
	// elements are accessible for independent verification by the general public
	PublicDistributedLedgerIndication *bool `json:"PublicDistributedLedgerIndication,omitempty" yaml:"PublicDistributedLedgerIndication,omitempty" mapstructure:"PublicDistributedLedgerIndication,omitempty"`

	// ShortNames corresponds to the JSON schema field "ShortNames".
	ShortNames []NativeDigitalTokenOtherJsonInformativeShortNamesElem `json:"ShortNames,omitempty" yaml:"ShortNames,omitempty" mapstructure:"ShortNames,omitempty"`

	// Uniform Resource Locator (URL) pointing to the digital token’s reference
	// implementation or software repository
	URL *string `json:"URL,omitempty" yaml:"URL,omitempty" mapstructure:"URL,omitempty"`

	// UnderlyingAssetExternalIdentifiers corresponds to the JSON schema field
	// "UnderlyingAssetExternalIdentifiers".
	UnderlyingAssetExternalIdentifiers []NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElem `json:"UnderlyingAssetExternalIdentifiers,omitempty" yaml:"UnderlyingAssetExternalIdentifiers,omitempty" mapstructure:"UnderlyingAssetExternalIdentifiers,omitempty"`

	// Multiplier used to map from the unit of value stored on the distributed ledger
	// to the unit of value associated with the digital token long name
	UnitMultiplier *big.Int `json:"UnitMultiplier,omitempty" yaml:"UnitMultiplier,omitempty" mapstructure:"UnitMultiplier,omitempty"`
}

type NativeDigitalTokenOtherJsonInformativeDTIExternalIdentifiersElem struct {
	// External identifier type for the digital token
	DTIExternalIdentifierType NativeDigitalTokenOtherJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType `json:"DTIExternalIdentifierType" yaml:"DTIExternalIdentifierType" mapstructure:"DTIExternalIdentifierType"`

	// External identifier for the digital token
	ExternalIdentifierValue string `json:"ExternalIdentifierValue" yaml:"ExternalIdentifierValue" mapstructure:"ExternalIdentifierValue"`
}

type NativeDigitalTokenOtherJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType string

const NativeDigitalTokenOtherJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierTypeITIN NativeDigitalTokenOtherJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType = "ITIN"

var enumValues_NativeDigitalTokenOtherJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType = []interface{}{
	"ITIN",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NativeDigitalTokenOtherJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_NativeDigitalTokenOtherJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_NativeDigitalTokenOtherJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType, v)
	}
	*j = NativeDigitalTokenOtherJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NativeDigitalTokenOtherJsonInformativeDTIExternalIdentifiersElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["DTIExternalIdentifierType"]; raw != nil && !ok {
		return fmt.Errorf("field DTIExternalIdentifierType in NativeDigitalTokenOtherJsonInformativeDTIExternalIdentifiersElem: required")
	}
	if _, ok := raw["ExternalIdentifierValue"]; raw != nil && !ok {
		return fmt.Errorf("field ExternalIdentifierValue in NativeDigitalTokenOtherJsonInformativeDTIExternalIdentifiersElem: required")
	}
	type Plain NativeDigitalTokenOtherJsonInformativeDTIExternalIdentifiersElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = NativeDigitalTokenOtherJsonInformativeDTIExternalIdentifiersElem(plain)
	return nil
}

type NativeDigitalTokenOtherJsonInformativeShortNamesElem struct {
	// OrigLangShortName corresponds to the JSON schema field "OrigLangShortName".
	OrigLangShortName string `json:"OrigLangShortName,omitempty" yaml:"OrigLangShortName,omitempty" mapstructure:"OrigLangShortName,omitempty"`

	// Short name or ticker symbol of the digital token
	ShortName string `json:"ShortName" yaml:"ShortName" mapstructure:"ShortName"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NativeDigitalTokenOtherJsonInformativeShortNamesElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["ShortName"]; raw != nil && !ok {
		return fmt.Errorf("field ShortName in NativeDigitalTokenOtherJsonInformativeShortNamesElem: required")
	}
	type Plain NativeDigitalTokenOtherJsonInformativeShortNamesElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["OrigLangShortName"]; !ok || v == nil {
		plain.OrigLangShortName = "Basic Latin (Unicode)"
	}
	*j = NativeDigitalTokenOtherJsonInformativeShortNamesElem(plain)
	return nil
}

type NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElem struct {
	// Underlying asset external identifier type for the digital token
	UnderlyingAssetExternalIdentifierType NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType `json:"UnderlyingAssetExternalIdentifierType" yaml:"UnderlyingAssetExternalIdentifierType" mapstructure:"UnderlyingAssetExternalIdentifierType"`

	// Underlying asset external identifier value for the digital token
	UnderlyingAssetExternalIdentifierValue string `json:"UnderlyingAssetExternalIdentifierValue" yaml:"UnderlyingAssetExternalIdentifierValue" mapstructure:"UnderlyingAssetExternalIdentifierValue"`
}

type NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType string

const NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeCCY NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "CCY"
const NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeCountryCode NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "Country code"
const NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeCusip NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "Cusip"
const NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeFIGI NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "FIGI"
const NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeISIN NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "ISIN"
const NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeLEI NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "LEI"
const NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeRIC NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "RIC"
const NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeSEDOL NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "SEDOL"

var enumValues_NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = []interface{}{
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
func (j *NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType, v)
	}
	*j = NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["UnderlyingAssetExternalIdentifierType"]; raw != nil && !ok {
		return fmt.Errorf("field UnderlyingAssetExternalIdentifierType in NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElem: required")
	}
	if _, ok := raw["UnderlyingAssetExternalIdentifierValue"]; raw != nil && !ok {
		return fmt.Errorf("field UnderlyingAssetExternalIdentifierValue in NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElem: required")
	}
	type Plain NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = NativeDigitalTokenOtherJsonInformativeUnderlyingAssetExternalIdentifiersElem(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NativeDigitalTokenOtherJsonInformative) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain NativeDigitalTokenOtherJsonInformative
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["OrigLangLongName"]; !ok || v == nil {
		plain.OrigLangLongName = "Basic Latin (Unicode)"
	}
	*j = NativeDigitalTokenOtherJsonInformative(plain)
	return nil
}

type NativeDigitalTokenOtherJsonMetadata struct {
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
func (j *NativeDigitalTokenOtherJsonMetadata) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["recDateTime"]; raw != nil && !ok {
		return fmt.Errorf("field recDateTime in NativeDigitalTokenOtherJsonMetadata: required")
	}
	if _, ok := raw["recVersion"]; raw != nil && !ok {
		return fmt.Errorf("field recVersion in NativeDigitalTokenOtherJsonMetadata: required")
	}
	type Plain NativeDigitalTokenOtherJsonMetadata
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
	*j = NativeDigitalTokenOtherJsonMetadata(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NativeDigitalTokenOtherJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["Header"]; raw != nil && !ok {
		return fmt.Errorf("field Header in NativeDigitalTokenOtherJson: required")
	}
	if _, ok := raw["Informative"]; raw != nil && !ok {
		return fmt.Errorf("field Informative in NativeDigitalTokenOtherJson: required")
	}
	if _, ok := raw["Metadata"]; raw != nil && !ok {
		return fmt.Errorf("field Metadata in NativeDigitalTokenOtherJson: required")
	}
	type Plain NativeDigitalTokenOtherJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = NativeDigitalTokenOtherJson(plain)
	return nil
}

// DTI returns token DTI.
func (j *NativeDigitalTokenOtherJson) DTI() string {
	return j.Header.DTI
}

// Denom returns token denom.
func (j *NativeDigitalTokenOtherJson) Denom() *string {
	return lo.ToPtr("")
}
