package dtif

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/CoreumFoundation/iso20022-client/iso20022/logger"
)

func TestEmptyDtif(t *testing.T) {
	ctrl := gomock.NewController(t)
	logMock := logger.NewAnyLogMock(ctrl)
	d := NewWithSourceAddress(logMock, "S87NJRT7T", "file://./testdata/data.json")

	denom, _, ok := d.LookupByDTI("KNNT25FGR")
	require.False(t, ok)
	require.Empty(t, denom)
}

func TestLookup(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	logMock := logger.NewAnyLogMock(ctrl)

	d := NewWithSourceAddress(logMock, "S87NJRT7T", "file://./testdata/data.json")

	require.NoError(t, d.Update(ctx))

	denom, _, ok := d.LookupByDTI("KNNT25FGR")
	require.True(t, ok)
	require.Equal(t, "ibc/71F11BC0AF8E526B80E44172EBA9D3F0A8E03950BB882325435691EBC9450B1D", denom)
}

func TestLookupByDenom(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	logMock := logger.NewAnyLogMock(ctrl)

	d := NewWithSourceAddress(logMock, "S87NJRT7T", "file://./testdata/data.json")

	require.NoError(t, d.Update(ctx))

	dti, ok := d.LookupByDenom("ibc/71F11BC0AF8E526B80E44172EBA9D3F0A8E03950BB882325435691EBC9450B1D")

	require.True(t, ok)
	require.Equal(t, "KNNT25FGR", dti)
}

func TestUpdate(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)
	logMock := logger.NewAnyLogMock(ctrl)

	testData := []struct {
		name string
		d    *Dtif
		err  bool
	}{
		{
			name: "wrong path",
			d:    NewWithSourceAddress(logMock, "S87NJRT7T", "file://./testdata/non-existing.json"),
			err:  true,
		},
	}

	for _, tt := range testData {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			res := tt.d.Update(ctx)
			if tt.err {
				require.Error(t, res)
			} else {
				require.NoError(t, res)
			}
		})
	}
}

func TestCache(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	testData := []struct {
		name        string
		dtifBuilder func(ctrl *gomock.Controller) *Dtif
	}{
		{
			name: "actual repo",
			dtifBuilder: func(ctrl *gomock.Controller) *Dtif {
				logMock := logger.NewMockLogger(ctrl)
				logMock.EXPECT().Debug(gomock.Any(), "DTIF data updated")
				logMock.EXPECT().Debug(gomock.Any(), "DTIF data is not changed, no need update")
				return New(logMock, "S87NJRT7T")
			},
		},
		{
			name: "local file",
			dtifBuilder: func(ctrl *gomock.Controller) *Dtif {
				logMock := logger.NewMockLogger(ctrl)
				logMock.EXPECT().Debug(gomock.Any(), "DTIF data updated")
				logMock.EXPECT().Debug(gomock.Any(), "DTIF data is not changed, no need update")
				return NewWithSourceAddress(
					logMock,
					"S87NJRT7T",
					"file://./testdata/data.json",
				)
			},
		},
	}

	for _, tt := range testData {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			d := tt.dtifBuilder(ctrl)
			require.NoError(t, d.Update(ctx))
			require.NoError(t, d.Update(ctx))
		})
	}
}
