package compress

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompressor(t *testing.T) {
	c, err := New()
	require.NoError(t, err)

	text := []byte("some text")

	compressed := c.Compress(text)

	decompressed, err := c.Decompress(compressed)
	require.NoError(t, err)
	require.Equal(t, text, decompressed)
}
