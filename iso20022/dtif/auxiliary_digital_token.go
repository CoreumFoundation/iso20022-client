package dtif

import (
	"fmt"
	"math/big"
)

type AuxiliaryDigitalTokenJson struct {
	Header struct {
		DTI             string `json:"DTI"`
		DTIType         int    `json:"DTIType"`
		TemplateVersion string `json:"templateVersion"`
	} `json:"Header"`
	Informative struct {
		LongName   string `json:"LongName"`
		ShortNames []struct {
			ShortName string `json:"ShortName"`
		} `json:"ShortNames"`
		UnitMultiplier any `json:"UnitMultiplier"`
	} `json:"Informative"`
	Normative struct {
		AuxiliaryDistributedLedger  string `json:"AuxiliaryDistributedLedger"`
		AuxiliaryTechnicalReference string `json:"AuxiliaryTechnicalReference"`
	} `json:"Normative"`
}

// DTI returns token DTI.
func (j *AuxiliaryDigitalTokenJson) DTI() string {
	return j.Header.DTI
}

// Denom returns token denom.
func (j *AuxiliaryDigitalTokenJson) Denom() *string {
	if j.Normative.AuxiliaryTechnicalReference == "<locked>" {
		return nil
	}
	return &j.Normative.AuxiliaryTechnicalReference
}

// PriceMultiplier returns token price multiplier.
func (j *AuxiliaryDigitalTokenJson) PriceMultiplier() *big.Int {
	if j.Normative.AuxiliaryTechnicalReference == "<locked>" {
		return nil
	}
	switch multiplier := j.Informative.UnitMultiplier.(type) {
	case string:
		if multiplier == "<locked>" {
			return nil
		}
		res, _ := new(big.Int).SetString(multiplier, 10)
		return res
	case float64:
		res, _ := new(big.Float).SetFloat64(multiplier).Int(nil)
		return res
	default:
		panic(fmt.Errorf("unknown multiplier type %v", j.Informative.UnitMultiplier))
	}

}
