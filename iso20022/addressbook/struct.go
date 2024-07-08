package addressbook

import "reflect"

type Proprietary struct {
	Id         string `json:"id"`
	Issuer     string `json:"issuer"`
	SchemeName string `json:"scheme_name"`
}

type AddressTypeCode string

const (
	AddressTypeCodeAddress   AddressTypeCode = "ADDR"
	AddressTypeCodePostalBox AddressTypeCode = "PBOX"
	AddressTypeCodeHome      AddressTypeCode = "HOME"
	AddressTypeCodeBusiness  AddressTypeCode = "BIZZ"
	AddressTypeCodeMailTo    AddressTypeCode = "MLTO"
	AddressTypeCodeDelivery  AddressTypeCode = "DLVY"
)

func (a AddressTypeCode) String() string {
	return string(a)
}

type AddressType struct {
	Code        AddressTypeCode `json:"code"`
	Proprietary *Proprietary    `json:"proprietary"`
}

type PostalAddress struct {
	AddressType        *AddressType `json:"address_type"`
	CareOf             string       `json:"care_of"`
	Department         string       `json:"department"`
	SubDepartment      string       `json:"sub_department"`
	StreetName         string       `json:"street_name"`
	BuildingNumber     string       `json:"building_number"`
	BuildingName       string       `json:"building_name"`
	Floor              string       `json:"floor"`
	UnitNumber         string       `json:"unit_number"`
	PostalBox          string       `json:"postal_box"`
	Room               string       `json:"room"`
	PostalCode         string       `json:"postal_code"`
	TownName           string       `json:"town_name"`
	TownLocationName   string       `json:"town_location_name"`
	DistrictName       string       `json:"district_name"`
	CountrySubDivision string       `json:"country_sub_division"`
	CountryCode        string       `json:"country_code"`
	AddressLine        []string     `json:"address_line"`
}

func (p *PostalAddress) Equal(other *PostalAddress) bool {
	return reflect.DeepEqual(p, other)
}

type ClearingSystemId struct {
	Code        string `json:"code"`
	Proprietary string `json:"proprietary"`
}

func (c *ClearingSystemId) Equal(other *ClearingSystemId) bool {
	return reflect.DeepEqual(c, other)
}

type ClearingSystemMemberIdentification struct {
	ClearingSystemId *ClearingSystemId `json:"clearing_system_id"`
	MemberId         string            `json:"member_id"`
}

type SchemeName struct {
	Code        string `json:"code"`
	Proprietary string `json:"proprietary"`
}

type Other struct {
	Id         string      `json:"id"`
	Issuer     string      `json:"issuer"`
	SchemeName *SchemeName `json:"scheme_name"`
}

func (o *Other) Equal(other *Other) bool {
	return reflect.DeepEqual(o, other)
}

type Identification struct {
	BusinessIdentifiersCode            string                              `json:"bic"`
	ClearingSystemMemberIdentification *ClearingSystemMemberIdentification `json:"clearing_system_member_identification"`
	LegalEntityIdentifier              string                              `json:"lei"`
	Name                               string                              `json:"name"`
	PostalAddress                      *PostalAddress                      `json:"postal_address"`
	Other                              *Other                              `json:"other"`
}

type Branch struct {
	Id                    string         `json:"id"`
	LegalEntityIdentifier string         `json:"lei"`
	Name                  string         `json:"name"`
	PostalAddress         *PostalAddress `json:"postal_address"`
}

func (b *Branch) Equal(other *Branch) bool {
	return reflect.DeepEqual(b, other)
}

type Party struct {
	Identification Identification `json:"identification"`
	Branch         *Branch        `json:"branch"`
}

// Equal checks if two ISO20022 Party are equal.
// The expected one can have more fields, and it will match if required fields of the actual one matches
func (b Party) Equal(expected Party) bool {
	if b.Branch != nil && expected.Branch != nil {
		if !b.Branch.Equal(expected.Branch) {
			return false
		}
	} else if b.Branch == nil && expected.Branch != nil || b.Branch != nil && expected.Branch == nil {
		return false
	}

	actualId := b.Identification
	expectedId := expected.Identification

	if actualId.BusinessIdentifiersCode != "" {
		return actualId.BusinessIdentifiersCode == expectedId.BusinessIdentifiersCode
	}

	if actualId.LegalEntityIdentifier != "" {
		return actualId.LegalEntityIdentifier == expectedId.LegalEntityIdentifier
	}

	if actualId.ClearingSystemMemberIdentification != nil {
		if expectedId.ClearingSystemMemberIdentification == nil {
			return false
		}

		actualCls := actualId.ClearingSystemMemberIdentification
		expectedCls := expectedId.ClearingSystemMemberIdentification
		if actualCls.MemberId != expectedCls.MemberId || len(actualCls.MemberId) == 0 {
			return false
		}
		if actualCls.ClearingSystemId != nil {
			if expectedCls.ClearingSystemId == nil {
				return false
			}

			if !actualCls.ClearingSystemId.Equal(expectedCls.ClearingSystemId) {
				return false
			}
		}
		return true
	}

	if actualId.Other != nil {
		return actualId.Other.Equal(expectedId.Other)
	}

	if actualId.PostalAddress != nil && expectedId.PostalAddress != nil {
		return actualId.Name == expectedId.Name && actualId.PostalAddress.Equal(expectedId.PostalAddress)
	}

	return false
}

type Address struct {
	Bech32EncodedAddress string `json:"bech32_encoded_address"`
	PublicKey            string `json:"public_key"`
	Party                Party  `json:"party"`
}

type StoredAddressBook struct {
	Addresses []Address `json:"addresses"`
}
