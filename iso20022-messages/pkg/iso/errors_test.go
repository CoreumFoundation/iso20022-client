package iso_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/pkg/iso"
)

func TestErrorCodes(t *testing.T) {
	ec := iso.IsError("9910")
	require.NotNil(t, ec)
	require.Equal(t, "9910", ec.Code)
	require.Equal(t, iso.ErrTemporary, ec.Level)

	ec = iso.IsError("be07")
	require.NotNil(t, ec)
	require.Equal(t, "BE07", ec.Code)
	require.Equal(t, iso.ErrFatal, ec.Level)

	ec = iso.IsError("ac03")
	require.NotNil(t, ec)
	require.Equal(t, "AC03", ec.Code)
	require.Equal(t, iso.ErrAccountFatal, ec.Level)

	ec = iso.IsError("AM13")
	require.NotNil(t, ec)
	require.Equal(t, "AM13", ec.Code)
	require.Equal(t, iso.ErrTemporary, ec.Level)

	ec = iso.IsError("AM09")
	require.NotNil(t, ec)
	require.Equal(t, "AM09", ec.Code)
	require.Equal(t, iso.ErrLogic, ec.Level)

	ec = iso.IsError("")
	require.Nil(t, ec)

	ec = iso.IsError("9999")
	require.Nil(t, ec)
}

func TestErrorLevel(t *testing.T) {
	// just make sure .Error() doesn't recusively panic
	level := iso.ErrNetwork
	require.Equal(t, "Network issue", level.Error())
	require.Equal(t, "Network issue", fmt.Sprintf("%v", level))
	require.Equal(t, "Network issue", fmt.Errorf("%w", level).Error())
}
