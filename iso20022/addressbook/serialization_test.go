package addressbook

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSerializeStruct(t *testing.T) {
	p := PostalAddress{
		AddressType: &AddressType{
			Code: "l",
			Proprietary: &Proprietary{
				Id:         "m",
				Issuer:     "n",
				SchemeName: "o",
			},
		},
		CareOf:             "a",
		Department:         "b",
		SubDepartment:      "c",
		StreetName:         "d",
		BuildingNumber:     "e",
		BuildingName:       "",
		Floor:              "f",
		UnitNumber:         "",
		PostalBox:          "",
		Room:               "",
		PostalCode:         "g",
		TownName:           "",
		TownLocationName:   "",
		DistrictName:       "h",
		CountrySubDivision: "",
		CountryCode:        "i",
		AddressLine:        []string{"j", "k"},
	}
	require.Equal(t, "AddressType=(Code=l/Proprietary=(Id=m/Issuer=n/SchemeName=o))/BuildingNumber=e/CareOf=a/CountryCode=i/Department=b/DistrictName=h/Floor=f/PostalCode=g/StreetName=d/SubDepartment=c", SerializeStruct(p))
}
