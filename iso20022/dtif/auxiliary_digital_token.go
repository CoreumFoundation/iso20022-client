package dtif

import (
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
)

// AuxiliaryDigitalTokenJson Non-native digital token created as an application
// on an existing blockchain or other distributed ledger technology for its
// issuance, storage, or transaction record
type AuxiliaryDigitalTokenJson struct {
	// Header corresponds to the JSON schema field "Header".
	Header AuxiliaryDigitalTokenJsonHeader `json:"Header" yaml:"Header" mapstructure:"Header"`

	// Informative corresponds to the JSON schema field "Informative".
	Informative AuxiliaryDigitalTokenJsonInformative `json:"Informative" yaml:"Informative" mapstructure:"Informative"`

	// Metadata corresponds to the JSON schema field "Metadata".
	Metadata AuxiliaryDigitalTokenJsonMetadata `json:"Metadata" yaml:"Metadata" mapstructure:"Metadata"`

	// Normative corresponds to the JSON schema field "Normative".
	Normative AuxiliaryDigitalTokenJsonNormative `json:"Normative" yaml:"Normative" mapstructure:"Normative"`
}

type AuxiliaryDigitalTokenJsonHeader struct {
	// DTI corresponds to the JSON schema field "DTI".
	DTI string `json:"DTI" yaml:"DTI" mapstructure:"DTI"`

	// Category of the digital token identifier
	DTIType AuxiliaryDigitalTokenJsonHeaderDTIType `json:"DTIType" yaml:"DTIType" mapstructure:"DTIType"`

	// The template version (i.e. JSON schema document) for which this record is valid
	TemplateVersion AuxiliaryDigitalTokenJsonHeaderTemplateVersion `json:"templateVersion" yaml:"templateVersion" mapstructure:"templateVersion"`
}

type AuxiliaryDigitalTokenJsonHeaderDTIType int

var enumValues_AuxiliaryDigitalTokenJsonHeaderDTIType = []interface{}{
	0,
	1,
	2,
	3,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuxiliaryDigitalTokenJsonHeaderDTIType) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AuxiliaryDigitalTokenJsonHeaderDTIType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AuxiliaryDigitalTokenJsonHeaderDTIType, v)
	}
	*j = AuxiliaryDigitalTokenJsonHeaderDTIType(v)
	return nil
}

type AuxiliaryDigitalTokenJsonHeaderTemplateVersion string

const AuxiliaryDigitalTokenJsonHeaderTemplateVersionV100 AuxiliaryDigitalTokenJsonHeaderTemplateVersion = "V1.0.0"

var enumValues_AuxiliaryDigitalTokenJsonHeaderTemplateVersion = []interface{}{
	"V1.0.0",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuxiliaryDigitalTokenJsonHeaderTemplateVersion) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AuxiliaryDigitalTokenJsonHeaderTemplateVersion {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AuxiliaryDigitalTokenJsonHeaderTemplateVersion, v)
	}
	*j = AuxiliaryDigitalTokenJsonHeaderTemplateVersion(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuxiliaryDigitalTokenJsonHeader) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["DTI"]; raw != nil && !ok {
		return fmt.Errorf("field DTI in AuxiliaryDigitalTokenJsonHeader: required")
	}
	if _, ok := raw["DTIType"]; raw != nil && !ok {
		return fmt.Errorf("field DTIType in AuxiliaryDigitalTokenJsonHeader: required")
	}
	if _, ok := raw["templateVersion"]; raw != nil && !ok {
		return fmt.Errorf("field templateVersion in AuxiliaryDigitalTokenJsonHeader: required")
	}
	type Plain AuxiliaryDigitalTokenJsonHeader
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AuxiliaryDigitalTokenJsonHeader(plain)
	return nil
}

type AuxiliaryDigitalTokenJsonInformative struct {
	// DTIExternalIdentifiers corresponds to the JSON schema field
	// "DTIExternalIdentifiers".
	DTIExternalIdentifiers []AuxiliaryDigitalTokenJsonInformativeDTIExternalIdentifiersElem `json:"DTIExternalIdentifiers,omitempty" yaml:"DTIExternalIdentifiers,omitempty" mapstructure:"DTIExternalIdentifiers,omitempty"`

	// Digital token long name is a string containing the full name of the digital
	// token
	LongName *string `json:"LongName,omitempty" yaml:"LongName,omitempty" mapstructure:"LongName,omitempty"`

	// OrigLangLongName corresponds to the JSON schema field "OrigLangLongName".
	OrigLangLongName string `json:"OrigLangLongName,omitempty" yaml:"OrigLangLongName,omitempty" mapstructure:"OrigLangLongName,omitempty"`

	// ShortNames corresponds to the JSON schema field "ShortNames".
	ShortNames []AuxiliaryDigitalTokenJsonInformativeShortNamesElem `json:"ShortNames,omitempty" yaml:"ShortNames,omitempty" mapstructure:"ShortNames,omitempty"`

	// UnderlyingAssetExternalIdentifiers corresponds to the JSON schema field
	// "UnderlyingAssetExternalIdentifiers".
	UnderlyingAssetExternalIdentifiers []AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElem `json:"UnderlyingAssetExternalIdentifiers,omitempty" yaml:"UnderlyingAssetExternalIdentifiers,omitempty" mapstructure:"UnderlyingAssetExternalIdentifiers,omitempty"`

	// Multiplier used to map from the unit of value stored on the distributed ledger
	// to the unit of value associated with the digital token long name
	UnitMultiplier *big.Int `json:"UnitMultiplier,omitempty" yaml:"UnitMultiplier,omitempty" mapstructure:"UnitMultiplier,omitempty"`
}

type AuxiliaryDigitalTokenJsonInformativeDTIExternalIdentifiersElem struct {
	// External identifier type for the digital token
	DTIExternalIdentifierType AuxiliaryDigitalTokenJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType `json:"DTIExternalIdentifierType" yaml:"DTIExternalIdentifierType" mapstructure:"DTIExternalIdentifierType"`

	// External identifier for the digital token
	ExternalIdentifierValue string `json:"ExternalIdentifierValue" yaml:"ExternalIdentifierValue" mapstructure:"ExternalIdentifierValue"`
}

type AuxiliaryDigitalTokenJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType string

const AuxiliaryDigitalTokenJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierTypeITIN AuxiliaryDigitalTokenJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType = "ITIN"

var enumValues_AuxiliaryDigitalTokenJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType = []interface{}{
	"ITIN",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuxiliaryDigitalTokenJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AuxiliaryDigitalTokenJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AuxiliaryDigitalTokenJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType, v)
	}
	*j = AuxiliaryDigitalTokenJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuxiliaryDigitalTokenJsonInformativeDTIExternalIdentifiersElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["DTIExternalIdentifierType"]; raw != nil && !ok {
		return fmt.Errorf("field DTIExternalIdentifierType in AuxiliaryDigitalTokenJsonInformativeDTIExternalIdentifiersElem: required")
	}
	if _, ok := raw["ExternalIdentifierValue"]; raw != nil && !ok {
		return fmt.Errorf("field ExternalIdentifierValue in AuxiliaryDigitalTokenJsonInformativeDTIExternalIdentifiersElem: required")
	}
	type Plain AuxiliaryDigitalTokenJsonInformativeDTIExternalIdentifiersElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AuxiliaryDigitalTokenJsonInformativeDTIExternalIdentifiersElem(plain)
	return nil
}

type AuxiliaryDigitalTokenJsonInformativeShortNamesElem struct {
	// OrigLangShortName corresponds to the JSON schema field "OrigLangShortName".
	OrigLangShortName string `json:"OrigLangShortName,omitempty" yaml:"OrigLangShortName,omitempty" mapstructure:"OrigLangShortName,omitempty"`

	// Short name or ticker symbol of the digital token
	ShortName string `json:"ShortName" yaml:"ShortName" mapstructure:"ShortName"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuxiliaryDigitalTokenJsonInformativeShortNamesElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["ShortName"]; raw != nil && !ok {
		return fmt.Errorf("field ShortName in AuxiliaryDigitalTokenJsonInformativeShortNamesElem: required")
	}
	type Plain AuxiliaryDigitalTokenJsonInformativeShortNamesElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["OrigLangShortName"]; !ok || v == nil {
		plain.OrigLangShortName = "Basic Latin (Unicode)"
	}
	*j = AuxiliaryDigitalTokenJsonInformativeShortNamesElem(plain)
	return nil
}

type AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElem struct {
	// Underlying asset external identifier type for the digital token
	UnderlyingAssetExternalIdentifierType AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType `json:"UnderlyingAssetExternalIdentifierType" yaml:"UnderlyingAssetExternalIdentifierType" mapstructure:"UnderlyingAssetExternalIdentifierType"`

	// Underlying asset external identifier value for the digital token
	UnderlyingAssetExternalIdentifierValue string `json:"UnderlyingAssetExternalIdentifierValue" yaml:"UnderlyingAssetExternalIdentifierValue" mapstructure:"UnderlyingAssetExternalIdentifierValue"`
}

type AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType string

const AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeCCY AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "CCY"
const AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeCountryCode AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "Country code"
const AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeCusip AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "Cusip"
const AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeFIGI AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "FIGI"
const AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeISIN AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "ISIN"
const AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeLEI AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "LEI"
const AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeRIC AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "RIC"
const AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeSEDOL AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "SEDOL"

var enumValues_AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = []interface{}{
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
func (j *AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType, v)
	}
	*j = AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["UnderlyingAssetExternalIdentifierType"]; raw != nil && !ok {
		return fmt.Errorf("field UnderlyingAssetExternalIdentifierType in AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElem: required")
	}
	if _, ok := raw["UnderlyingAssetExternalIdentifierValue"]; raw != nil && !ok {
		return fmt.Errorf("field UnderlyingAssetExternalIdentifierValue in AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElem: required")
	}
	type Plain AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AuxiliaryDigitalTokenJsonInformativeUnderlyingAssetExternalIdentifiersElem(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuxiliaryDigitalTokenJsonInformative) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain AuxiliaryDigitalTokenJsonInformative
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["OrigLangLongName"]; !ok || v == nil {
		plain.OrigLangLongName = "Basic Latin (Unicode)"
	}
	*j = AuxiliaryDigitalTokenJsonInformative(plain)
	return nil
}

type AuxiliaryDigitalTokenJsonMetadata struct {
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
func (j *AuxiliaryDigitalTokenJsonMetadata) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["recDateTime"]; raw != nil && !ok {
		return fmt.Errorf("field recDateTime in AuxiliaryDigitalTokenJsonMetadata: required")
	}
	if _, ok := raw["recVersion"]; raw != nil && !ok {
		return fmt.Errorf("field recVersion in AuxiliaryDigitalTokenJsonMetadata: required")
	}
	type Plain AuxiliaryDigitalTokenJsonMetadata
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
	*j = AuxiliaryDigitalTokenJsonMetadata(plain)
	return nil
}

type AuxiliaryDigitalTokenJsonNormative struct {
	// A DTI of a native digital token [DTIType=1] or of a distributed ledger without
	// a distributed ledger [DTIType=2]. The distributed ledger of the DTI is used as
	// the platform for the auxiliary digital token
	AuxiliaryDistributedLedger *string `json:"AuxiliaryDistributedLedger,omitempty" yaml:"AuxiliaryDistributedLedger,omitempty" mapstructure:"AuxiliaryDistributedLedger,omitempty"`

	// AuxiliaryMechanism corresponds to the JSON schema field "AuxiliaryMechanism".
	AuxiliaryMechanism *AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism `json:"AuxiliaryMechanism,omitempty" yaml:"AuxiliaryMechanism,omitempty" mapstructure:"AuxiliaryMechanism,omitempty"`

	// Element, such as a smart contract address, used to uniquely identify an
	// auxiliary digital token’s origin on a distributed ledger technology platform
	AuxiliaryTechnicalReference *string `json:"AuxiliaryTechnicalReference,omitempty" yaml:"AuxiliaryTechnicalReference,omitempty" mapstructure:"AuxiliaryTechnicalReference,omitempty"`
}

type AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism string

const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismARC20 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "ARC-20"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismASA AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "ASA"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismATS AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "ATS"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismBEP2 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "BEP-2"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismBEP20 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "BEP-20"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismBIP32 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "BIP-32"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismBRC20 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "BRC-20"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismCIS2 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "CIS-2"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismCRC20 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "CRC20"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismCW20 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "CW20"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismCardanoSmartContract AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "Cardano Smart Contract"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismDOG20 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "DOG-20"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismEOSIOTOKEN AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "EOSIO.TOKEN"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismERC20 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "ERC-20"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismESDT AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "ESDT"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismFA2 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "FA2"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismFT AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "FT"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismFungibleCashToken AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "Fungible CashToken"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismGSDAP AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "GSDAP"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismHRC20 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "HRC20"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismHRC_20 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "HRC-20"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismHTS AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "HTS"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismIBCCoin AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "IBC Coin"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismICRC1 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "ICRC-1"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismJetton AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "Jetton"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismKIP20 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "KIP-20"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismKIP7 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "KIP-7"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismLightning AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "Lightning"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismLiquidAsset AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "LiquidAsset"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismNEP141 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "NEP-141"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismNEP17 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "NEP-17"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismNEP5 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "NEP-5"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismNativeAttribute AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "Native Attribute"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismNativeCoin AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "Native Coin"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismOEP4 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "OEP-4"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismOMNI AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "OMNI"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismPolkadotAsset AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "Polkadot Asset"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismSDX AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "SDX"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismSEP1 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "SEP-1"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismSEP20 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "SEP20"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismSIP10 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "SIP-10"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismSLPToken AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "SLP Token"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismSORAToken AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "SORA Token"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismSPLToken AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "SPL-Token"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismSUICoin AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "SUI Coin"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismStatemineAsset AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "StatemineAsset"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismStatemintAsset AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "StatemintAsset"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismTAI2 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "TAI2"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismTIP3 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "TIP-3"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismTRC10 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "TRC-10"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismTRC20 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "TRC-20"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismVRC20 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "VRC20"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismVRC21 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "VRC21"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismVRC25 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "VRC25"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismViteTokens AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "Vite Tokens"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismWRC20 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "WRC-20"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismWSC AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "WSC"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismXRC20 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "XRC20"
const AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanismZRC2 AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = "ZRC-2"

var enumValues_AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism = []interface{}{
	"ATS",
	"ASA",
	"ARC-20",
	"BEP-2",
	"BEP-20",
	"BIP-32",
	"BRC-20",
	"Cardano Smart Contract",
	"CIS-2",
	"CRC20",
	"CW20",
	"DOG-20",
	"ERC-20",
	"ESDT",
	"EOSIO.TOKEN",
	"FA2",
	"FT",
	"Fungible CashToken",
	"GSDAP",
	"HRC20",
	"HRC-20",
	"HTS",
	"IBC Coin",
	"ICRC-1",
	"Jetton",
	"KIP-7",
	"KIP-20",
	"Lightning",
	"LiquidAsset",
	"Native Attribute",
	"Native Coin",
	"NEP-5",
	"NEP-17",
	"NEP-141",
	"OEP-4",
	"OMNI",
	"Polkadot Asset",
	"SEP-1",
	"SEP20",
	"SIP-10",
	"SORA Token",
	"SPL-Token",
	"SUI Coin",
	"StatemineAsset",
	"StatemintAsset",
	"TAI2",
	"TIP-3",
	"TRC-10",
	"TRC-20",
	"SDX",
	"SLP Token",
	"WRC-20",
	"WSC",
	"Vite Tokens",
	"VRC20",
	"VRC21",
	"VRC25",
	"XRC20",
	"ZRC-2",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism, v)
	}
	*j = AuxiliaryDigitalTokenJsonNormativeAuxiliaryMechanism(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *AuxiliaryDigitalTokenJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["Header"]; raw != nil && !ok {
		return fmt.Errorf("field Header in AuxiliaryDigitalTokenJson: required")
	}
	if _, ok := raw["Informative"]; raw != nil && !ok {
		return fmt.Errorf("field Informative in AuxiliaryDigitalTokenJson: required")
	}
	if _, ok := raw["Metadata"]; raw != nil && !ok {
		return fmt.Errorf("field Metadata in AuxiliaryDigitalTokenJson: required")
	}
	if _, ok := raw["Normative"]; raw != nil && !ok {
		return fmt.Errorf("field Normative in AuxiliaryDigitalTokenJson: required")
	}
	type Plain AuxiliaryDigitalTokenJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = AuxiliaryDigitalTokenJson(plain)
	return nil
}

// DTI returns token DTI.
func (j *AuxiliaryDigitalTokenJson) DTI() string {
	return j.Header.DTI
}

// Denom returns token denom.
func (j *AuxiliaryDigitalTokenJson) Denom() *string {
	return j.Normative.AuxiliaryTechnicalReference
}
