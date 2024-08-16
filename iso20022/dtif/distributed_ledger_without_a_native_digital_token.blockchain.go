package dtif

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// DistributedLedgerWithoutANativeDigitalTokenBlockchainJson Distributed
// Ledger without a Native Digital Token exists solely to support
// registration and identifier assignment of auxiliary digital token
type DistributedLedgerWithoutANativeDigitalTokenBlockchainJson struct {
	// Header corresponds to the JSON schema field "Header".
	Header DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeader `json:"Header" yaml:"Header" mapstructure:"Header"`

	// Informative corresponds to the JSON schema field "Informative".
	Informative DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonInformative `json:"Informative" yaml:"Informative" mapstructure:"Informative"`

	// Metadata corresponds to the JSON schema field "Metadata".
	Metadata DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonMetadata `json:"Metadata" yaml:"Metadata" mapstructure:"Metadata"`

	// Normative corresponds to the JSON schema field "Normative".
	Normative DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormative `json:"Normative" yaml:"Normative" mapstructure:"Normative"`
}

type DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeader struct {
	// Blockchain
	DLTType DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeaderDLTType `json:"DLTType" yaml:"DLTType" mapstructure:"DLTType"`

	// DTI corresponds to the JSON schema field "DTI".
	DTI string `json:"DTI" yaml:"DTI" mapstructure:"DTI"`

	// Distributed Ledger without a Native Digital Token
	DTIType DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeaderDTIType `json:"DTIType" yaml:"DTIType" mapstructure:"DTIType"`

	// The template version (i.e. JSON schema document) for which this record is valid
	TemplateVersion DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeaderTemplateVersion `json:"templateVersion" yaml:"templateVersion" mapstructure:"templateVersion"`
}

type DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeaderDLTType int

var enumValues_DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeaderDLTType = []interface{}{
	1,
	0,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeaderDLTType) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeaderDLTType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeaderDLTType, v)
	}
	*j = DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeaderDLTType(v)
	return nil
}

type DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeaderDTIType int

var enumValues_DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeaderDTIType = []interface{}{
	0,
	1,
	2,
	3,
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeaderDTIType) UnmarshalJSON(b []byte) error {
	var v int
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeaderDTIType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeaderDTIType, v)
	}
	*j = DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeaderDTIType(v)
	return nil
}

type DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeaderTemplateVersion string

const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeaderTemplateVersionV100 DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeaderTemplateVersion = "V1.0.0"

var enumValues_DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeaderTemplateVersion = []interface{}{
	"V1.0.0",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeaderTemplateVersion) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeaderTemplateVersion {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeaderTemplateVersion, v)
	}
	*j = DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeaderTemplateVersion(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeader) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["DLTType"]; raw != nil && !ok {
		return fmt.Errorf("field DLTType in DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeader: required")
	}
	if _, ok := raw["DTI"]; raw != nil && !ok {
		return fmt.Errorf("field DTI in DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeader: required")
	}
	if _, ok := raw["DTIType"]; raw != nil && !ok {
		return fmt.Errorf("field DTIType in DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeader: required")
	}
	if _, ok := raw["templateVersion"]; raw != nil && !ok {
		return fmt.Errorf("field templateVersion in DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeader: required")
	}
	type Plain DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeader
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonHeader(plain)
	return nil
}

type DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonInformative struct {
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
func (j *DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonInformative) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	type Plain DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonInformative
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["OrigLangLongName"]; !ok || v == nil {
		plain.OrigLangLongName = "Basic Latin (Unicode)"
	}
	*j = DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonInformative(plain)
	return nil
}

type DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonMetadata struct {
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
func (j *DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonMetadata) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["recDateTime"]; raw != nil && !ok {
		return fmt.Errorf("field recDateTime in DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonMetadata: required")
	}
	if _, ok := raw["recVersion"]; raw != nil && !ok {
		return fmt.Errorf("field recVersion in DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonMetadata: required")
	}
	type Plain DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonMetadata
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
	*j = DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonMetadata(plain)
	return nil
}

type DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormative struct {
	// Forks corresponds to the JSON schema field "Forks".
	Forks []DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeForksElem `json:"Forks,omitempty" yaml:"Forks,omitempty" mapstructure:"Forks,omitempty"`

	// Block hash of the genesis block
	GenesisBlockHash *string `json:"GenesisBlockHash,omitempty" yaml:"GenesisBlockHash,omitempty" mapstructure:"GenesisBlockHash,omitempty"`

	// Block hash algorithm used to produce the block hash of the genesis block
	GenesisBlockHashAlgorithm *DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm `json:"GenesisBlockHashAlgorithm,omitempty" yaml:"GenesisBlockHashAlgorithm,omitempty" mapstructure:"GenesisBlockHashAlgorithm,omitempty"`

	// Timestamp, expressed in Coordinated Universal Time, recorded in the genesis
	// block
	GenesisBlockUTCTimestamp *JsonTime `json:"GenesisBlockUTCTimestamp,omitempty" yaml:"GenesisBlockUTCTimestamp,omitempty" mapstructure:"GenesisBlockUTCTimestamp,omitempty"`
}

// Each fork is a creation of two or more different versions of a distributed
// ledger originating from a common starting point with a single history; A fork
// may or may not result in the creation of a new digital token
type DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeForksElem struct {
	// If true, the consensus mechanism used to validate the block at the fork block
	// height identified in the fork record differ from the consensus mechanism used
	// to validate the block immediately prior to the block at the fork block height
	// identified in the fork record
	ConsensusMechanismChangeResponse *bool `json:"ConsensusMechanismChangeResponse,omitempty" yaml:"ConsensusMechanismChangeResponse,omitempty" mapstructure:"ConsensusMechanismChangeResponse,omitempty"`

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

type DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm string

const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmBLAKE256 DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "BLAKE-256"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmBLAKE2B256 DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "BLAKE2b-256"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmBLAKE2B2560X DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "BLAKE2b-256 (0x)"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmBLAKE2Base58 DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "BLAKE2 (Base58)"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmBLAKE2S256 DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "BLAKE2s-256"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmDoubleSHA256 DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "Double SHA-256"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmKeccak256 DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "Keccak-256"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmLISKSPEC DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "LISK SPEC"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmLOST DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "LOST"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmPoseidonHash DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "Poseidon Hash"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA224160Bits DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA-224 (160Bits)"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA256 DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA-256"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA256Base32 DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA-256 (Base32)"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA256Base58 DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA-256 (Base58)"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA256Base64Url DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA-256 (Base64url)"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA256Decimals DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA-256 (Decimals)"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA3256 DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA3-256"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA3256Base58 DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA3-256 (Base58)"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA384 DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA-384"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA512256Base16 DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA512-256 (Base16)"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA512256Base32 DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA512/256 (Base32)"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHA512Half DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHA-512Half"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSHAKE256 DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SHAKE-256"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmSteemHash DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "SteemHash"
const DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithmVRFSignature DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = "VRF Signature"

var enumValues_DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm = []interface{}{
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
func (j *DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm, v)
	}
	*j = DistributedLedgerWithoutANativeDigitalTokenBlockchainJsonNormativeGenesisBlockHashAlgorithm(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DistributedLedgerWithoutANativeDigitalTokenBlockchainJson) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["Header"]; raw != nil && !ok {
		return fmt.Errorf("field Header in DistributedLedgerWithoutANativeDigitalTokenBlockchainJson: required")
	}
	if _, ok := raw["Informative"]; raw != nil && !ok {
		return fmt.Errorf("field Informative in DistributedLedgerWithoutANativeDigitalTokenBlockchainJson: required")
	}
	if _, ok := raw["Metadata"]; raw != nil && !ok {
		return fmt.Errorf("field Metadata in DistributedLedgerWithoutANativeDigitalTokenBlockchainJson: required")
	}
	if _, ok := raw["Normative"]; raw != nil && !ok {
		return fmt.Errorf("field Normative in DistributedLedgerWithoutANativeDigitalTokenBlockchainJson: required")
	}
	type Plain DistributedLedgerWithoutANativeDigitalTokenBlockchainJson
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DistributedLedgerWithoutANativeDigitalTokenBlockchainJson(plain)
	return nil
}

// DTI returns token DTI.
func (j *DistributedLedgerWithoutANativeDigitalTokenBlockchainJson) DTI() string {
	return j.Header.DTI
}

// Denom returns token denom.
func (j *DistributedLedgerWithoutANativeDigitalTokenBlockchainJson) Denom() *string {
	// TODO: Make sure this is the right field to extract denom from
	return j.Informative.LongName
}
