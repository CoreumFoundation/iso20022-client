package dtif

import (
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"

	"github.com/samber/lo"
)

// NativeDigitalTokenBlockchainJson Digital token with a privileged
// position in the blockchain distributed ledger technology protocol
type NativeDigitalTokenBlockchainJson struct {
	// Header corresponds to the JSON schema field "Header".
	Header NativeDigitalTokenBlockchainJsonHeader `json:"Header" yaml:"Header" mapstructure:"Header"`

	// Informative corresponds to the JSON schema field "Informative".
	Informative NativeDigitalTokenBlockchainJsonInformative `json:"Informative" yaml:"Informative" mapstructure:"Informative"`

	// Metadata corresponds to the JSON schema field "Metadata".
	Metadata NativeDigitalTokenBlockchainJsonMetadata `json:"Metadata" yaml:"Metadata" mapstructure:"Metadata"`

	// Normative corresponds to the JSON schema field "Normative".
	Normative NativeDigitalTokenBlockchainJsonNormative `json:"Normative" yaml:"Normative" mapstructure:"Normative"`
}

type NativeDigitalTokenBlockchainJsonHeader struct {
	// Blockchain
	DLTType NativeDigitalTokenBlockchainJsonHeaderDLTType `json:"DLTType" yaml:"DLTType" mapstructure:"DLTType"`

	// DTI corresponds to the JSON schema field "DTI".
	DTI string `json:"DTI" yaml:"DTI" mapstructure:"DTI"`

	// Native Digital Token
	DTIType NativeDigitalTokenBlockchainJsonHeaderDTIType `json:"DTIType" yaml:"DTIType" mapstructure:"DTIType"`

	// The template version (i.e. JSON schema document) for which this record is valid
	TemplateVersion NativeDigitalTokenBlockchainJsonHeaderTemplateVersion `json:"templateVersion" yaml:"templateVersion" mapstructure:"templateVersion"`
}

type NativeDigitalTokenBlockchainJsonHeaderDLTType int

var enumValues_NativeDigitalTokenBlockchainJsonHeaderDLTType = []interface{}{
	1,
	0,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NativeDigitalTokenBlockchainJsonHeaderDLTType) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_NativeDigitalTokenBlockchainJsonHeaderDLTType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_NativeDigitalTokenBlockchainJsonHeaderDLTType, v)
	}
	*j = NativeDigitalTokenBlockchainJsonHeaderDLTType(v)
	return nil
}

type NativeDigitalTokenBlockchainJsonHeaderDTIType int

var enumValues_NativeDigitalTokenBlockchainJsonHeaderDTIType = []interface{}{
	0,
	1,
	2,
	3,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NativeDigitalTokenBlockchainJsonHeaderDTIType) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_NativeDigitalTokenBlockchainJsonHeaderDTIType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_NativeDigitalTokenBlockchainJsonHeaderDTIType, v)
	}
	*j = NativeDigitalTokenBlockchainJsonHeaderDTIType(v)
	return nil
}

type NativeDigitalTokenBlockchainJsonHeaderTemplateVersion string

const NativeDigitalTokenBlockchainJsonHeaderTemplateVersionV100 NativeDigitalTokenBlockchainJsonHeaderTemplateVersion = "V1.0.0"

var enumValues_NativeDigitalTokenBlockchainJsonHeaderTemplateVersion = []interface{}{
	"V1.0.0",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NativeDigitalTokenBlockchainJsonHeaderTemplateVersion) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_NativeDigitalTokenBlockchainJsonHeaderTemplateVersion {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_NativeDigitalTokenBlockchainJsonHeaderTemplateVersion, v)
	}
	*j = NativeDigitalTokenBlockchainJsonHeaderTemplateVersion(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NativeDigitalTokenBlockchainJsonHeader) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["DLTType"]; raw != nil && !ok {
		return fmt.Errorf("field DLTType in NativeDigitalTokenBlockchainJsonHeader: required")
	}
	if _, ok := raw["DTI"]; raw != nil && !ok {
		return fmt.Errorf("field DTI in NativeDigitalTokenBlockchainJsonHeader: required")
	}
	if _, ok := raw["DTIType"]; raw != nil && !ok {
		return fmt.Errorf("field DTIType in NativeDigitalTokenBlockchainJsonHeader: required")
	}
	if _, ok := raw["templateVersion"]; raw != nil && !ok {
		return fmt.Errorf("field templateVersion in NativeDigitalTokenBlockchainJsonHeader: required")
	}
	type Plain NativeDigitalTokenBlockchainJsonHeader
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = NativeDigitalTokenBlockchainJsonHeader(plain)
	return nil
}

type NativeDigitalTokenBlockchainJsonInformative struct {
	// DTIExternalIdentifiers corresponds to the JSON schema field
	// "DTIExternalIdentifiers".
	DTIExternalIdentifiers []NativeDigitalTokenBlockchainJsonInformativeDTIExternalIdentifiersElem `json:"DTIExternalIdentifiers,omitempty" yaml:"DTIExternalIdentifiers,omitempty" mapstructure:"DTIExternalIdentifiers,omitempty"`

	// Digital token long name is a string containing the full name of the digital
	// token
	LongName *string `json:"LongName,omitempty" yaml:"LongName,omitempty" mapstructure:"LongName,omitempty"`

	// OrigLangLongName corresponds to the JSON schema field "OrigLangLongName".
	OrigLangLongName string `json:"OrigLangLongName,omitempty" yaml:"OrigLangLongName,omitempty" mapstructure:"OrigLangLongName,omitempty"`

	// If true, access to reading the distributed ledger is unrestricted and the data
	// elements are accessible for independent verification by the general public
	PublicDistributedLedgerIndication *bool `json:"PublicDistributedLedgerIndication,omitempty" yaml:"PublicDistributedLedgerIndication,omitempty" mapstructure:"PublicDistributedLedgerIndication,omitempty"`

	// ShortNames corresponds to the JSON schema field "ShortNames".
	ShortNames []NativeDigitalTokenBlockchainJsonInformativeShortNamesElem `json:"ShortNames,omitempty" yaml:"ShortNames,omitempty" mapstructure:"ShortNames,omitempty"`

	// Uniform Resource Locator (URL) pointing to the digital token’s reference
	// implementation or software repository
	URL *string `json:"URL,omitempty" yaml:"URL,omitempty" mapstructure:"URL,omitempty"`

	// UnderlyingAssetExternalIdentifiers corresponds to the JSON schema field
	// "UnderlyingAssetExternalIdentifiers".
	UnderlyingAssetExternalIdentifiers []NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElem `json:"UnderlyingAssetExternalIdentifiers,omitempty" yaml:"UnderlyingAssetExternalIdentifiers,omitempty" mapstructure:"UnderlyingAssetExternalIdentifiers,omitempty"`

	// Multiplier used to map from the unit of value stored on the distributed ledger
	// to the unit of value associated with the digital token long name
	UnitMultiplier *big.Int `json:"UnitMultiplier,omitempty" yaml:"UnitMultiplier,omitempty" mapstructure:"UnitMultiplier,omitempty"`
}

type NativeDigitalTokenBlockchainJsonInformativeDTIExternalIdentifiersElem struct {
	// External identifier type for the digital token
	DTIExternalIdentifierType NativeDigitalTokenBlockchainJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType `json:"DTIExternalIdentifierType" yaml:"DTIExternalIdentifierType" mapstructure:"DTIExternalIdentifierType"`

	// External identifier for the digital token
	ExternalIdentifierValue string `json:"ExternalIdentifierValue" yaml:"ExternalIdentifierValue" mapstructure:"ExternalIdentifierValue"`
}

type NativeDigitalTokenBlockchainJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType string

const NativeDigitalTokenBlockchainJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierTypeITIN NativeDigitalTokenBlockchainJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType = "ITIN"

var enumValues_NativeDigitalTokenBlockchainJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType = []interface{}{
	"ITIN",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NativeDigitalTokenBlockchainJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_NativeDigitalTokenBlockchainJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_NativeDigitalTokenBlockchainJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType, v)
	}
	*j = NativeDigitalTokenBlockchainJsonInformativeDTIExternalIdentifiersElemDTIExternalIdentifierType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NativeDigitalTokenBlockchainJsonInformativeDTIExternalIdentifiersElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["DTIExternalIdentifierType"]; raw != nil && !ok {
		return fmt.Errorf("field DTIExternalIdentifierType in NativeDigitalTokenBlockchainJsonInformativeDTIExternalIdentifiersElem: required")
	}
	if _, ok := raw["ExternalIdentifierValue"]; raw != nil && !ok {
		return fmt.Errorf("field ExternalIdentifierValue in NativeDigitalTokenBlockchainJsonInformativeDTIExternalIdentifiersElem: required")
	}
	type Plain NativeDigitalTokenBlockchainJsonInformativeDTIExternalIdentifiersElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = NativeDigitalTokenBlockchainJsonInformativeDTIExternalIdentifiersElem(plain)
	return nil
}

type NativeDigitalTokenBlockchainJsonInformativeShortNamesElem struct {
	// OrigLangShortName corresponds to the JSON schema field "OrigLangShortName".
	OrigLangShortName string `json:"OrigLangShortName,omitempty" yaml:"OrigLangShortName,omitempty" mapstructure:"OrigLangShortName,omitempty"`

	// Short name or ticker symbol of the digital token
	ShortName string `json:"ShortName" yaml:"ShortName" mapstructure:"ShortName"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NativeDigitalTokenBlockchainJsonInformativeShortNamesElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["ShortName"]; raw != nil && !ok {
		return fmt.Errorf("field ShortName in NativeDigitalTokenBlockchainJsonInformativeShortNamesElem: required")
	}
	type Plain NativeDigitalTokenBlockchainJsonInformativeShortNamesElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["OrigLangShortName"]; !ok || v == nil {
		plain.OrigLangShortName = "Basic Latin (Unicode)"
	}
	*j = NativeDigitalTokenBlockchainJsonInformativeShortNamesElem(plain)
	return nil
}

type NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElem struct {
	// Underlying asset external identifier type for the digital token
	UnderlyingAssetExternalIdentifierType NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType `json:"UnderlyingAssetExternalIdentifierType" yaml:"UnderlyingAssetExternalIdentifierType" mapstructure:"UnderlyingAssetExternalIdentifierType"`

	// Underlying asset external identifier value for the digital token
	UnderlyingAssetExternalIdentifierValue string `json:"UnderlyingAssetExternalIdentifierValue" yaml:"UnderlyingAssetExternalIdentifierValue" mapstructure:"UnderlyingAssetExternalIdentifierValue"`
}

type NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType string

const NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeCCY NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "CCY"
const NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeCountryCode NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "Country code"
const NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeCusip NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "Cusip"
const NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeFIGI NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "FIGI"
const NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeISIN NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "ISIN"
const NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeLEI NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "LEI"
const NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeRIC NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "RIC"
const NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierTypeSEDOL NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = "SEDOL"

var enumValues_NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType = []interface{}{
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
func (j *NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType, v)
	}
	*j = NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElemUnderlyingAssetExternalIdentifierType(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["UnderlyingAssetExternalIdentifierType"]; raw != nil && !ok {
		return fmt.Errorf("field UnderlyingAssetExternalIdentifierType in NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElem: required")
	}
	if _, ok := raw["UnderlyingAssetExternalIdentifierValue"]; raw != nil && !ok {
		return fmt.Errorf("field UnderlyingAssetExternalIdentifierValue in NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElem: required")
	}
	type Plain NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = NativeDigitalTokenBlockchainJsonInformativeUnderlyingAssetExternalIdentifiersElem(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NativeDigitalTokenBlockchainJsonInformative) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain NativeDigitalTokenBlockchainJsonInformative
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["OrigLangLongName"]; !ok || v == nil {
		plain.OrigLangLongName = "Basic Latin (Unicode)"
	}
	*j = NativeDigitalTokenBlockchainJsonInformative(plain)
	return nil
}

type NativeDigitalTokenBlockchainJsonMetadata struct {
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
func (j *NativeDigitalTokenBlockchainJsonMetadata) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["recDateTime"]; raw != nil && !ok {
		return fmt.Errorf("field recDateTime in NativeDigitalTokenBlockchainJsonMetadata: required")
	}
	if _, ok := raw["recVersion"]; raw != nil && !ok {
		return fmt.Errorf("field recVersion in NativeDigitalTokenBlockchainJsonMetadata: required")
	}
	type Plain NativeDigitalTokenBlockchainJsonMetadata
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
	*j = NativeDigitalTokenBlockchainJsonMetadata(plain)
	return nil
}

type NativeDigitalTokenBlockchainJsonNormative struct {
	// Forks corresponds to the JSON schema field "Forks".
	Forks []NativeDigitalTokenBlockchainJsonNormativeForksElem `json:"Forks,omitempty" yaml:"Forks,omitempty" mapstructure:"Forks,omitempty"`

	// Block hash of the genesis block
	GenesisBlockHash *string `json:"GenesisBlockHash,omitempty" yaml:"GenesisBlockHash,omitempty" mapstructure:"GenesisBlockHash,omitempty"`

	// Block hash algorithm used to produce the block hash of the genesis block
	GenesisBlockHashAlgorithm *NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm `json:"GenesisBlockHashAlgorithm,omitempty" yaml:"GenesisBlockHashAlgorithm,omitempty" mapstructure:"GenesisBlockHashAlgorithm,omitempty"`

	// Timestamp, expressed in Coordinated Universal Time, recorded in the genesis
	// block
	GenesisBlockUTCTimestamp *JsonTime `json:"GenesisBlockUTCTimestamp,omitempty" yaml:"GenesisBlockUTCTimestamp,omitempty" mapstructure:"GenesisBlockUTCTimestamp,omitempty"`
}

// Each fork is a creation of two or more different versions of a distributed
// ledger originating from a common starting point with a single history; A fork
// may or may not result in the creation of a new digital token
type NativeDigitalTokenBlockchainJsonNormativeForksElem struct {
	// If true, the consensus mechanism used to validate the block at the fork block
	// height identified in the fork record differ from the consensus mechanism used
	// to validate the block immediately prior to the block at the fork block height
	// identified in the fork record
	ConsensusMechanismChangeResponse *bool `json:"ConsensusMechanismChangeResponse,omitempty" yaml:"ConsensusMechanismChangeResponse,omitempty" mapstructure:"ConsensusMechanismChangeResponse,omitempty"`

	// If true, the fork result in the creation of a new digital token
	DigitalTokenCreationResponse *bool `json:"DigitalTokenCreationResponse,omitempty" yaml:"DigitalTokenCreationResponse,omitempty" mapstructure:"DigitalTokenCreationResponse,omitempty"`

	// Block hash of the block with a block height equal to the fork block height
	ForkBlockHash *string `json:"ForkBlockHash,omitempty" yaml:"ForkBlockHash,omitempty" mapstructure:"ForkBlockHash,omitempty"`

	// Block hash algorithm of the fork block
	ForkBlockHashAlgorithm *string `json:"ForkBlockHashAlgorithm,omitempty" yaml:"ForkBlockHashAlgorithm,omitempty" mapstructure:"ForkBlockHashAlgorithm,omitempty"`

	// Block height of the first block after the fork
	ForkBlockHeight *int `json:"ForkBlockHeight,omitempty" yaml:"ForkBlockHeight,omitempty" mapstructure:"ForkBlockHeight,omitempty"`

	// Timestamp, expressed in Coordinated Universal Time, recorded in the fork block
	ForkBlockUTCTimestamp *JsonTime `json:"ForkBlockUTCTimestamp,omitempty" yaml:"ForkBlockUTCTimestamp,omitempty" mapstructure:"ForkBlockUTCTimestamp,omitempty"`

	// A reference to the base record the fork record modifies
	ForkReferenceDTI *string `json:"ForkReferenceDTI,omitempty" yaml:"ForkReferenceDTI,omitempty" mapstructure:"ForkReferenceDTI,omitempty"`
}

type NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm string

const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmBLAKE256 NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "BLAKE-256"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmBLAKE2B256 NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "BLAKE2b-256"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmBLAKE2B2560X NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "BLAKE2b-256 (0x)"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmBLAKE2Base58 NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "BLAKE2 (Base58)"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmBLAKE2S256 NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "BLAKE2s-256"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmDoubleSHA256 NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "Double SHA-256"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmKeccak256 NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "Keccak-256"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmLISKSPEC NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "LISK SPEC"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmLOST NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "LOST"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmPoseidonHash NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "Poseidon Hash"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA224160Bits NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA-224 (160Bits)"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA256 NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA-256"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA256Base32 NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA-256 (Base32)"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA256Base58 NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA-256 (Base58)"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA256Base64Url NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA-256 (Base64url)"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA256Decimals NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA-256 (Decimals)"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA3256 NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA3-256"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA3256Base58 NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA3-256 (Base58)"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA384 NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA-384"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA512256Base16 NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA512-256 (Base16)"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA512256Base32 NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA512/256 (Base32)"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA512Half NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA-512Half"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHAKE256 NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHAKE-256"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSteemHash NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SteemHash"
const NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmVRFSignature NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "VRF Signature"

var enumValues_NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = []interface{}{
	"BLAKE2 (Base58)",
	"BLAKE-256",
	"BLAKE2b-256",
	"BLAKE2b-256 (0x)",
	"BLAKE2s-256",
	"Double SHA-256",
	"Keccak-256",
	"LISK SPEC",
	"LOST",
	"Poseidon Hash",
	"SHAKE-256",
	"SHA-224 (160Bits)",
	"SHA3-256",
	"SHA3-256 (Base58)",
	"SHA-256 (Decimals)",
	"SHA-256",
	"SHA-256 (Base32)",
	"SHA-256 (Base58)",
	"SHA-256 (Base64url)",
	"SHA-384",
	"SHA-512Half",
	"SHA512-256 (Base16)",
	"SHA512/256 (Base32)",
	"SteemHash",
	"VRF Signature",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm, v)
	}
	*j = NativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *NativeDigitalTokenBlockchainJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["Header"]; raw != nil && !ok {
		return fmt.Errorf("field Header in NativeDigitalTokenBlockchainJson: required")
	}
	if _, ok := raw["Informative"]; raw != nil && !ok {
		return fmt.Errorf("field Informative in NativeDigitalTokenBlockchainJson: required")
	}
	if _, ok := raw["Metadata"]; raw != nil && !ok {
		return fmt.Errorf("field Metadata in NativeDigitalTokenBlockchainJson: required")
	}
	if _, ok := raw["Normative"]; raw != nil && !ok {
		return fmt.Errorf("field Normative in NativeDigitalTokenBlockchainJson: required")
	}
	type Plain NativeDigitalTokenBlockchainJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = NativeDigitalTokenBlockchainJson(plain)
	return nil
}

// DTI returns token DTI.
func (j *NativeDigitalTokenBlockchainJson) DTI() string {
	return j.Header.DTI
}

// Denom returns token denom.
func (j *NativeDigitalTokenBlockchainJson) Denom() *string {
	return lo.ToPtr("")
}

// PriceMultiplier returns token price multiplier.
func (j *NativeDigitalTokenBlockchainJson) PriceMultiplier() *big.Int {
	return j.Informative.UnitMultiplier
}
