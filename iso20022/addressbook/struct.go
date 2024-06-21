package addressbook

type Proprietary struct {
	Id         string `json:"id"`
	Issuer     string `json:"issuer"`
	SchemeName string `json:"scheme_name"`
}

type AddressType struct {
	Code        string       `json:"code"`
	Proprietary *Proprietary `json:"proprietary"`
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

func (p PostalAddress) String() string {
	return SerializeStruct(p)
}

type ClearingSystemId struct {
	Code        string `json:"code"`
	Proprietary string `json:"proprietary"`
}

func (c ClearingSystemId) String() string {
	return SerializeStruct(c)
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
	Issuer     string     `json:"issuer"`
	SchemeName SchemeName `json:"scheme_name"`
}

func (o Other) String() string {
	return SerializeStruct(o)
}

type Identification struct {
	Bic                                string                              `json:"bic"`
	ClearingSystemMemberIdentification *ClearingSystemMemberIdentification `json:"clearing_system_member_identification"`
	Lei                                string                              `json:"lei"`
	Name                               string                              `json:"name"`
	PostalAddress                      *PostalAddress                      `json:"postal_address"`
	Other                              *Other                              `json:"other"`
}

type Branch struct {
	Id            string         `json:"id"`
	Lei           string         `json:"lei"`
	Name          string         `json:"name"`
	PostalAddress *PostalAddress `json:"postal_address"`
}

func (b Branch) String() string {
	return SerializeStruct(b)
}

type BranchAndIdentification struct {
	Identification Identification `json:"identification"`
	Branch         *Branch        `json:"branch"`
}

// Equal checks if two ISO20022 BranchAndIdentification are equal
func (b BranchAndIdentification) Equal(other BranchAndIdentification) bool {
	if b.Branch != nil && other.Branch != nil {
		if !AreStringsEqual(b.Branch.String(), other.Branch.String()) {
			return false
		}
	} else if b.Branch == nil && other.Branch != nil || b.Branch != nil && other.Branch == nil {
		return false
	}

	actualId := b.Identification
	expectedId := other.Identification

	if AreStringsEqual(actualId.Bic, expectedId.Bic) || AreStringsEqual(actualId.Lei, expectedId.Lei) {
		return true
	}

	if actualId.ClearingSystemMemberIdentification != nil && expectedId.ClearingSystemMemberIdentification != nil {
		actualCls := actualId.ClearingSystemMemberIdentification
		expectedCls := expectedId.ClearingSystemMemberIdentification
		if AreStringsEqual(actualCls.MemberId, expectedCls.MemberId) {
			return true
		}
		if actualCls.ClearingSystemId != nil && expectedCls.ClearingSystemId != nil {
			if AreStringsEqual(actualCls.ClearingSystemId.String(), expectedCls.ClearingSystemId.String()) {
				return true
			}
		}
	}

	if actualId.PostalAddress != nil && expectedId.PostalAddress != nil {
		if AreStringsEqual(actualId.Name, expectedId.Name) && AreStringsEqual(actualId.PostalAddress.String(), expectedId.PostalAddress.String()) {
			return true
		}
	}

	if (actualId.Other != nil && expectedId.Other != nil) && AreStringsEqual(actualId.Other.String(), expectedId.Other.String()) {
		return true
	}

	return false
}

type Address struct {
	Bech32EncodedAddress    string                  `json:"bech32_encoded_address"`
	PublicKey               string                  `json:"public_key"`
	BranchAndIdentification BranchAndIdentification `json:"branch_and_identification"`
}

type StoredAddressBook struct {
	Schema       string    `json:"$schema"`
	ChainId      string    `json:"chain_id"`
	NetworkType  string    `json:"network_type"`
	Bech32Prefix string    `json:"bech32_prefix"`
	KeyAlgo      string    `json:"key_algo"`
	Addresses    []Address `json:"addresses"`
}
