package processes_test

import (
	"context"
	"testing"
	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	coreumlogger "github.com/CoreumFoundation/coreum-tools/pkg/logger"
	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
	"github.com/CoreumFoundation/iso20022-client/iso20022/processes"
)

//nolint:tparallel // the test is parallel, but test cases are not
func TestDtifUpdater_Start(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name        string
		dtifBuilder func(ctrl *gomock.Controller) processes.Dtif
	}{
		{
			name: "update_once",
			dtifBuilder: func(ctrl *gomock.Controller) processes.Dtif {
				dtifMock := NewMockDtif(ctrl)
				dtifMock.EXPECT().Update(gomock.Any()).Return(nil)
				return dtifMock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ctx := coreumlogger.WithLogger(context.Background(), coreumlogger.New(coreumlogger.ToolDefaultConfig))
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			t.Cleanup(cancel)

			ctrl := gomock.NewController(t)
			logMock := logger.NewAnyLogMock(ctrl)

			var dtifProcess processes.Dtif
			if tt.dtifBuilder != nil {
				dtifProcess = tt.dtifBuilder(ctrl)
			}
			client, err := processes.NewDtifUpdaterProcess(3*time.Second, logMock, dtifProcess)
			require.NoError(t, err)

			err = client.Start(ctx)
			if err == nil || !errors.Is(err, context.DeadlineExceeded) {
				require.NoError(t, err)
			}
		})
	}
}
