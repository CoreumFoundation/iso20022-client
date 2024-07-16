package iso_test

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/CoreumFoundation/iso20022-client/iso20022-messages/pkg/iso"
)

func TestAmountFormat(t *testing.T) {
	var amt = iso.Amount(634)
	var amtTag = "<Amount>634.00</Amount>"

	out, err := amt.MarshalText()
	require.NoError(t, err)
	require.Equal(t, "634.00", string(out))

	out, err = xml.Marshal(amt)
	require.NoError(t, err)
	require.Equal(t, amtTag, string(out))

	var read iso.Amount
	err = xml.Unmarshal([]byte(amtTag), &read)
	require.NoError(t, err)
	require.Equal(t, amt, read)

	t.Run("large", func(t *testing.T) {
		amt = iso.Amount(1_252_363.25)
		bs, err := xml.Marshal(amt)
		require.NoError(t, err)
		require.Equal(t, "<Amount>1252363.25</Amount>", string(bs))
	})
}

func TestAmountValidate(t *testing.T) {
	var amt = iso.Amount(634)
	require.NoError(t, amt.Validate())
}
