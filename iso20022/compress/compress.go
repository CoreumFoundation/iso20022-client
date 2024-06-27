package compress

import "github.com/klauspost/compress/zstd"

type Compressor struct {
	encoder *zstd.Encoder
	decoder *zstd.Decoder
}

func New() (*Compressor, error) {
	encoder, err := zstd.NewWriter(nil)
	if err != nil {
		return nil, err
	}

	decoder, err := zstd.NewReader(nil)
	if err != nil {
		return nil, err
	}

	return &Compressor{
		encoder,
		decoder,
	}, nil
}

// Compress a buffer
func (c *Compressor) Compress(src []byte) []byte {
	return c.encoder.EncodeAll(src, make([]byte, 0, len(src)))
}

// Decompress a buffer
func (c *Compressor) Decompress(src []byte) ([]byte, error) {
	return c.decoder.DecodeAll(src, make([]byte, 0, len(src)))
}
