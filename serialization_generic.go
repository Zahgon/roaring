//go:build (!amd64 && !386 && !arm && !arm64 && !ppc64le && !mipsle && !mips64le && !mips64p32le && !wasm) || appengine
// +build !amd64,!386,!arm,!arm64,!ppc64le,!mipsle,!mips64le,!mips64p32le,!wasm appengine

package roaring

import (
	"io"
)

func (b *arrayContainer) writeTo(stream io.Writer) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *arrayContainer) readFrom(stream io.Reader) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (b *bitmapContainer) writeTo(stream io.Writer) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Write set

func (b *bitmapContainer) readFrom(stream io.Reader) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (bc *bitmapContainer) asLittleEndianByteSlice() []byte { _ = "STUB: not implemented"; return nil }

func uint64SliceAsByteSlice(slice []uint64) []byte { _ = "STUB: not implemented"; return nil }

func uint16SliceAsByteSlice(slice []uint16) []byte { _ = "STUB: not implemented"; return nil }

func interval16SliceAsByteSlice(slice []interval16) []byte { _ = "STUB: not implemented"; return nil }

func byteSliceAsUint16Slice(slice []byte) []uint16 { _ = "STUB: not implemented"; return nil }

func byteSliceAsUint64Slice(slice []byte) []uint64 { _ = "STUB: not implemented"; return nil }

// Converts a byte slice to a interval16 slice.
// The function assumes that the slice byte buffer is run container data
// encoded according to Roaring Format Spec
func byteSliceAsInterval16Slice(byteSlice []byte) []interval16 {
	_ = "STUB: not implemented"
	return nil
}
