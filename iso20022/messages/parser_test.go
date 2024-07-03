package messages

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/CoreumFoundation/iso20022-client/iso20022/addressbook"
	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
)

func TestParseIsoMessage(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	ctrl := gomock.NewController(t)
	logMock := logger.NewAnyLogMock(ctrl)
	parser := NewParser(logMock)

	tests := []struct {
		name            string
		messageFilePath string
		identification  *addressbook.BranchAndIdentification
		hasError        bool
	}{
		{
			name:            "pacs008",
			messageFilePath: "testdata/pacs008-1.xml",
			identification: &addressbook.BranchAndIdentification{
				Identification: addressbook.Identification{
					Bic: "6P9YGUDF",
				},
			},
			hasError: false,
		},
		{
			name:            "pacs008 within envelope",
			messageFilePath: "testdata/pacs008-2.xml",
			identification: &addressbook.BranchAndIdentification{
				Identification: addressbook.Identification{
					Bic: "6P9YGUDF",
				},
			},
			hasError: false,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			fileContent, err := os.ReadFile(tt.messageFilePath)
			require.NoError(t, err)

			identification, err := parser.ExtractIdentificationFromIsoMessage(ctx, fileContent)
			require.NoError(t, err)
			require.NotNil(t, identification)
			require.True(t, tt.identification.Equal(*identification))
		})
	}
}
