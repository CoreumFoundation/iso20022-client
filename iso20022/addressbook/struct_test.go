package addressbook

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEqual(t *testing.T) {
	fi := Party{
		Identification: Identification{
			BusinessIdentifiersCode: "abc",
			LegalEntityIdentifier:   "abc",
			Name:                    "John Doe",
			PostalAddress: &PostalAddress{
				AddressType: &AddressType{
					Code: AddressTypeCodeBusiness,
				},
				CareOf:         "Someone",
				Department:     "Software department",
				StreetName:     "Something",
				BuildingNumber: "Something Building",
				Floor:          "4",
				UnitNumber:     "12",
				PostalCode:     "1234567890",
				TownName:       "LA",
				CountryCode:    "US",
				AddressLine:    []string{"1234 Something"},
			},
			Other: &Other{
				Id: "abc",
			},
		},
		Branch: &Branch{
			Id: "branch",
		},
	}

	testData := []struct {
		name       string
		actualFI   Party
		expectedFI Party
		equal      bool
	}{
		{
			name:       "empty address",
			actualFI:   Party{},
			expectedFI: Party{},
			equal:      false,
		},
		{
			name: "equal bic",
			actualFI: Party{
				Identification: Identification{
					BusinessIdentifiersCode: "abc",
				},
			},
			expectedFI: Party{
				Identification: Identification{
					BusinessIdentifiersCode: "abc",
				},
			},
			equal: true,
		},
		{
			name: "equal bic different branch",
			actualFI: Party{
				Identification: Identification{
					BusinessIdentifiersCode: "abc",
				},
				Branch: &Branch{
					Id: "b1",
				},
			},
			expectedFI: Party{
				Identification: Identification{
					BusinessIdentifiersCode: "abc",
				},
				Branch: &Branch{
					Id: "b2",
				},
			},
			equal: false,
		},
		{
			name: "equal clearing system different member id",
			actualFI: Party{
				Identification: Identification{
					ClearingSystemMemberIdentification: &ClearingSystemMemberIdentification{
						ClearingSystemId: &ClearingSystemId{
							Code:        "a",
							Proprietary: "b",
						},
						MemberId: "1",
					},
				},
			},
			expectedFI: Party{
				Identification: Identification{
					ClearingSystemMemberIdentification: &ClearingSystemMemberIdentification{
						ClearingSystemId: &ClearingSystemId{
							Code:        "a",
							Proprietary: "b",
						},
						MemberId: "2",
					},
				},
			},
			equal: false,
		},
		{
			name: "equal lei",
			actualFI: Party{
				Identification: Identification{
					LegalEntityIdentifier: "abc",
				},
			},
			expectedFI: Party{
				Identification: Identification{
					LegalEntityIdentifier: "abc",
				},
			},
			equal: true,
		},
		{
			name: "equal other",
			actualFI: Party{
				Identification: Identification{
					Other: &Other{
						Id: "abc",
					},
				},
			},
			expectedFI: Party{
				Identification: Identification{
					Other: &Other{
						Id: "abc",
					},
				},
			},
			equal: true,
		},
		{
			name: "equal address",
			actualFI: Party{
				Identification: Identification{
					Name: "John Doe",
					PostalAddress: &PostalAddress{
						AddressType: &AddressType{
							Code: AddressTypeCodeBusiness,
						},
						CareOf:         "Someone",
						Department:     "Software department",
						StreetName:     "Something",
						BuildingNumber: "Something Building",
						Floor:          "4",
						UnitNumber:     "12",
						PostalCode:     "1234567890",
						TownName:       "LA",
						CountryCode:    "US",
						AddressLine:    []string{"1234 Something"},
					},
				},
			},
			expectedFI: Party{
				Identification: Identification{
					Name: "John Doe",
					PostalAddress: &PostalAddress{
						AddressType: &AddressType{
							Code: AddressTypeCodeBusiness,
						},
						CareOf:         "Someone",
						Department:     "Software department",
						StreetName:     "Something",
						BuildingNumber: "Something Building",
						Floor:          "4",
						UnitNumber:     "12",
						PostalCode:     "1234567890",
						TownName:       "LA",
						CountryCode:    "US",
						AddressLine:    []string{"1234 Something"},
					},
				},
			},
			equal: true,
		},

		{
			name:     "equal bic",
			actualFI: fi,
			expectedFI: Party{
				Identification: Identification{
					BusinessIdentifiersCode: "abc",
				},
				Branch: &Branch{
					Id: "branch",
				},
			},
			equal: true,
		},
		{
			name: "equal bic different branch",
			actualFI: Party{
				Identification: Identification{
					BusinessIdentifiersCode: "abc",
				},
				Branch: &Branch{
					Id: "b2",
				},
			},
			expectedFI: fi,
			equal:      false,
		},
		{
			name: "equal clearing system different member id",
			actualFI: Party{
				Identification: Identification{
					ClearingSystemMemberIdentification: &ClearingSystemMemberIdentification{
						ClearingSystemId: &ClearingSystemId{
							Code:        "a",
							Proprietary: "b",
						},
						MemberId: "2",
					},
				},
			},
			expectedFI: fi,
			equal:      false,
		},
		{
			name: "equal lei",
			actualFI: Party{
				Identification: Identification{
					LegalEntityIdentifier: "abc",
				},
				Branch: &Branch{
					Id: "branch",
				},
			},
			expectedFI: fi,
			equal:      true,
		},
		{
			name: "equal other",
			actualFI: Party{
				Identification: Identification{
					Other: &Other{
						Id: "abc",
					},
				},
				Branch: &Branch{
					Id: "branch",
				},
			},
			expectedFI: fi,
			equal:      true,
		},
		{
			name: "equal address",
			actualFI: Party{
				Identification: Identification{
					Name: "John Doe",
					PostalAddress: &PostalAddress{
						AddressType: &AddressType{
							Code: AddressTypeCodeBusiness,
						},
						CareOf:         "Someone",
						Department:     "Software department",
						StreetName:     "Something",
						BuildingNumber: "Something Building",
						Floor:          "4",
						UnitNumber:     "12",
						PostalCode:     "1234567890",
						TownName:       "LA",
						CountryCode:    "US",
						AddressLine:    []string{"1234 Something"},
					},
				},
				Branch: &Branch{
					Id: "branch",
				},
			},
			expectedFI: fi,
			equal:      true,
		},

		// TODO: Add more cases
	}

	for _, tt := range testData {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.equal {
				require.True(t, tt.actualFI.Equal(tt.expectedFI))
			} else {
				require.False(t, tt.actualFI.Equal(tt.expectedFI))
			}
		})
	}
}
