package processes_test

import (
	"context"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	logger2 "github.com/CoreumFoundation/coreum-tools/pkg/logger"
	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
	"github.com/CoreumFoundation/iso20022-client/iso20022/processes"
)

//nolint:tparallel // the test is parallel, but test cases are not
func TestAddressBookUpdater_Start(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name               string
		addressBookBuilder func(ctrl *gomock.Controller) processes.AddressBook
	}{
		{
			name: "update_once",
			addressBookBuilder: func(ctrl *gomock.Controller) processes.AddressBook {
				addressBookMock := NewMockAddressBook(ctrl)
				addressBookMock.EXPECT().Update(gomock.Any()).Return(nil)
				return addressBookMock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx := logger2.WithLogger(context.Background(), logger2.New(logger2.ToolDefaultConfig))
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			t.Cleanup(cancel)

			ctrl := gomock.NewController(t)
			logMock := logger.NewAnyLogMock(ctrl)

			var addressBook processes.AddressBook
			if tt.addressBookBuilder != nil {
				addressBook = tt.addressBookBuilder(ctrl)
			}
			client, err := processes.NewAddressBookUpdaterProcess(3*time.Second, logMock, addressBook)
			require.NoError(t, err)

			err = client.Start(ctx)
			if err == nil || !errors.Is(err, context.DeadlineExceeded) {
				require.NoError(t, err)
			}
		})
	}
}
