package roaring

import (
	"io"
)

// writeTo for runContainer16 follows this
// spec: https://github.com/RoaringBitmap/RoaringFormatSpec
func (b *runContainer16) writeTo(stream io.Writer) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
