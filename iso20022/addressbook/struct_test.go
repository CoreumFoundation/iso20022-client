package addressbook

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEqual(t *testing.T) {
	testData := []struct {
		name  string
		addr1 BranchAndIdentification
		addr2 BranchAndIdentification
		equal bool
	}{
		{
			name:  "empty address",
			addr1: BranchAndIdentification{},
			addr2: BranchAndIdentification{},
			equal: false,
		},
		{
			name: "equal bic",
			addr1: BranchAndIdentification{
				Identification: Identification{
					Bic: "abc",
				},
			},
			addr2: BranchAndIdentification{
				Identification: Identification{
					Bic: "abc",
				},
			},
			equal: true,
		},
		{
			name: "equal bic different branch",
			addr1: BranchAndIdentification{
				Identification: Identification{
					Bic: "abc",
				},
				Branch: &Branch{
					Id: "b1",
				},
			},
			addr2: BranchAndIdentification{
				Identification: Identification{
					Bic: "abc",
				},
				Branch: &Branch{
					Id: "b2",
				},
			},
			equal: false,
		},
		{
			name: "equal clearing system different member id",
			addr1: BranchAndIdentification{
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
			addr2: BranchAndIdentification{
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
		// TODO: Add more cases
	}

	for _, tt := range testData {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			if tt.equal {
				require.True(t, tt.addr1.Equal(tt.addr2))
			} else {
				require.False(t, tt.addr1.Equal(tt.addr2))
			}
		})
	}
}
