package dtif

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/samber/lo"
)

// DistributedLedgerWithoutANativeDigitalTokenOtherJson Distributed
// Ledger without a Native Digital Token exists solely to support
// registration and identifier assignment of auxiliary digital token
type DistributedLedgerWithoutANativeDigitalTokenOtherJson struct {
	// Header corresponds to the JSON schema field "Header".
	Header DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeader `json:"Header" yaml:"Header" mapstructure:"Header"`

	// Informative corresponds to the JSON schema field "Informative".
	Informative DistributedLedgerWithoutANativeDigitalTokenOtherJsonInformative `json:"Informative" yaml:"Informative" mapstructure:"Informative"`

	// Metadata corresponds to the JSON schema field "Metadata".
	Metadata DistributedLedgerWithoutANativeDigitalTokenOtherJsonMetadata `json:"Metadata" yaml:"Metadata" mapstructure:"Metadata"`
}

type DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeader struct {
	// Other
	DLTType DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeaderDLTType `json:"DLTType" yaml:"DLTType" mapstructure:"DLTType"`

	// DTI corresponds to the JSON schema field "DTI".
	DTI string `json:"DTI" yaml:"DTI" mapstructure:"DTI"`

	// Distributed Ledger without a Native Digital Token
	DTIType DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeaderDTIType `json:"DTIType" yaml:"DTIType" mapstructure:"DTIType"`

	// The template version (i.e. JSON schema document) for which this record is valid
	TemplateVersion DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeaderTemplateVersion `json:"templateVersion" yaml:"templateVersion" mapstructure:"templateVersion"`
}

type DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeaderDLTType int

var enumValues_DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeaderDLTType = []interface{}{
	1,
	0,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeaderDLTType) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeaderDLTType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeaderDLTType, v)
	}
	*j = DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeaderDLTType(v)
	return nil
}

type DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeaderDTIType int

var enumValues_DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeaderDTIType = []interface{}{
	0,
	1,
	2,
	3,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeaderDTIType) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeaderDTIType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeaderDTIType, v)
	}
	*j = DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeaderDTIType(v)
	return nil
}

type DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeaderTemplateVersion string

const DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeaderTemplateVersionV100 DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeaderTemplateVersion = "V1.0.0"

var enumValues_DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeaderTemplateVersion = []interface{}{
	"V1.0.0",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeaderTemplateVersion) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeaderTemplateVersion {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeaderTemplateVersion, v)
	}
	*j = DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeaderTemplateVersion(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeader) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["DLTType"]; raw != nil && !ok {
		return fmt.Errorf("field DLTType in DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeader: required")
	}
	if _, ok := raw["DTI"]; raw != nil && !ok {
		return fmt.Errorf("field DTI in DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeader: required")
	}
	if _, ok := raw["DTIType"]; raw != nil && !ok {
		return fmt.Errorf("field DTIType in DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeader: required")
	}
	if _, ok := raw["templateVersion"]; raw != nil && !ok {
		return fmt.Errorf("field templateVersion in DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeader: required")
	}
	type Plain DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeader
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DistributedLedgerWithoutANativeDigitalTokenOtherJsonHeader(plain)
	return nil
}

type DistributedLedgerWithoutANativeDigitalTokenOtherJsonInformative struct {
	// Refers to the long name of the distributed ledger
	LongName *string `json:"LongName,omitempty" yaml:"LongName,omitempty" mapstructure:"LongName,omitempty"`

	// Refers to the original language long name of the distributed ledger
	OrigLangLongName string `json:"OrigLangLongName,omitempty" yaml:"OrigLangLongName,omitempty" mapstructure:"OrigLangLongName,omitempty"`

	// If true, access to reading the distributed ledger is unrestricted and the data
	// elements are accessible for independent verification by the general public
	PublicDistributedLedgerIndication *bool `json:"PublicDistributedLedgerIndication,omitempty" yaml:"PublicDistributedLedgerIndication,omitempty" mapstructure:"PublicDistributedLedgerIndication,omitempty"`

	// Refers to the reference implementation of the distributed ledger
	URL *string `json:"URL,omitempty" yaml:"URL,omitempty" mapstructure:"URL,omitempty"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DistributedLedgerWithoutANativeDigitalTokenOtherJsonInformative) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain DistributedLedgerWithoutANativeDigitalTokenOtherJsonInformative
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["OrigLangLongName"]; !ok || v == nil {
		plain.OrigLangLongName = "Basic Latin (Unicode)"
	}
	*j = DistributedLedgerWithoutANativeDigitalTokenOtherJsonInformative(plain)
	return nil
}

type DistributedLedgerWithoutANativeDigitalTokenOtherJsonMetadata struct {
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
func (j *DistributedLedgerWithoutANativeDigitalTokenOtherJsonMetadata) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["recDateTime"]; raw != nil && !ok {
		return fmt.Errorf("field recDateTime in DistributedLedgerWithoutANativeDigitalTokenOtherJsonMetadata: required")
	}
	if _, ok := raw["recVersion"]; raw != nil && !ok {
		return fmt.Errorf("field recVersion in DistributedLedgerWithoutANativeDigitalTokenOtherJsonMetadata: required")
	}
	type Plain DistributedLedgerWithoutANativeDigitalTokenOtherJsonMetadata
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
	*j = DistributedLedgerWithoutANativeDigitalTokenOtherJsonMetadata(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DistributedLedgerWithoutANativeDigitalTokenOtherJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["Header"]; raw != nil && !ok {
		return fmt.Errorf("field Header in DistributedLedgerWithoutANativeDigitalTokenOtherJson: required")
	}
	if _, ok := raw["Informative"]; raw != nil && !ok {
		return fmt.Errorf("field Informative in DistributedLedgerWithoutANativeDigitalTokenOtherJson: required")
	}
	if _, ok := raw["Metadata"]; raw != nil && !ok {
		return fmt.Errorf("field Metadata in DistributedLedgerWithoutANativeDigitalTokenOtherJson: required")
	}
	type Plain DistributedLedgerWithoutANativeDigitalTokenOtherJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DistributedLedgerWithoutANativeDigitalTokenOtherJson(plain)
	return nil
}

// DTI returns token DTI.
func (j *DistributedLedgerWithoutANativeDigitalTokenOtherJson) DTI() string {
	return j.Header.DTI
}

// Denom returns token denom.
func (j *DistributedLedgerWithoutANativeDigitalTokenOtherJson) Denom() *string {
	return lo.ToPtr("")
}
