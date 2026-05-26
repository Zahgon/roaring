package internal

import (
	"io"
)

// ByteInput typed interface around io.Reader or raw bytes
type ByteInput interface {
	// Next returns a slice containing the next n bytes from the buffer,
	// advancing the buffer as if the bytes had been returned by Read.
	Next(n int) ([]byte, error)
	// NextReturnsSafeSlice returns true if Next() returns a safe slice as opposed
	// to a slice that points to an underlying buffer possibly owned by another system.
	// When NextReturnsSafeSlice returns false, the result from Next() should be copied
	// before it is modified (i.e., it is immutable).
	NextReturnsSafeSlice() bool
	// ReadUInt32 reads uint32 with LittleEndian order
	ReadUInt32() (uint32, error)
	// ReadUInt16 reads uint16 with LittleEndian order
	ReadUInt16() (uint16, error)
	// GetReadBytes returns read bytes
	GetReadBytes() int64
	// SkipBytes skips exactly n bytes
	SkipBytes(n int) error
}

// NewByteInputFromReader creates reader wrapper
func NewByteInputFromReader(reader io.Reader) ByteInput {
	_ = "STUB: not implemented"
	return *new(ByteInput)
}

// NewByteInput creates raw bytes wrapper
func NewByteInput(buf []byte) ByteInput { _ = "STUB: not implemented"; return *new(ByteInput) }

// ByteBuffer raw bytes wrapper
type ByteBuffer struct {
	buf []byte
	off int
}

// NewByteBuffer creates a new ByteBuffer.
func NewByteBuffer(buf []byte) *ByteBuffer { _ = "STUB: not implemented"; return nil }

var _ io.Reader = (*ByteBuffer)(nil)

// Read implements io.Reader.
func (b *ByteBuffer) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Next returns a slice containing the next n bytes from the reader
// If there are fewer bytes than the given n, io.ErrUnexpectedEOF will be returned
func (b *ByteBuffer) Next(n int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// NextReturnsSafeSlice returns false since ByteBuffer might hold
// an array owned by some other systems.
func (b *ByteBuffer) NextReturnsSafeSlice() bool {
	_ = "STUB: not implemented"

	// ReadUInt32 reads uint32 with LittleEndian order
	return false
}

func (b *ByteBuffer) ReadUInt32() (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadUInt16 reads uint16 with LittleEndian order
func (b *ByteBuffer) ReadUInt16() (uint16, error) { _ = "STUB: not implemented"; return 0, nil }

// GetReadBytes returns read bytes
func (b *ByteBuffer) GetReadBytes() int64 { _ = "STUB: not implemented"; return 0 }

// SkipBytes skips exactly n bytes
func (b *ByteBuffer) SkipBytes(n int) error { _ = "STUB: not implemented"; return nil }

// Reset resets the given buffer with a new byte slice
func (b *ByteBuffer) Reset(buf []byte) { _ = "STUB: not implemented"; return }

// ByteInputAdapter reader wrapper
type ByteInputAdapter struct {
	r         io.Reader
	readBytes int
	buf       [4]byte
}

var _ io.Reader = (*ByteInputAdapter)(nil)

// Read implements io.Reader.
func (b *ByteInputAdapter) Read(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Next returns a slice containing the next n bytes from the buffer,
// advancing the buffer as if the bytes had been returned by Read.
func (b *ByteInputAdapter) Next(n int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// NextReturnsSafeSlice returns true since ByteInputAdapter always returns a slice
// allocated with make([]byte, ...)
func (b *ByteInputAdapter) NextReturnsSafeSlice() bool {
	_ = "STUB: not implemented"

	// ReadUInt32 reads uint32 with LittleEndian order
	return false
}

func (b *ByteInputAdapter) ReadUInt32() (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadUInt16 reads uint16 with LittleEndian order
func (b *ByteInputAdapter) ReadUInt16() (uint16, error) { _ = "STUB: not implemented"; return 0, nil }

// GetReadBytes returns read bytes
func (b *ByteInputAdapter) GetReadBytes() int64 { _ = "STUB: not implemented"; return 0 }

// SkipBytes skips exactly n bytes
func (b *ByteInputAdapter) SkipBytes(n int) error { _ = "STUB: not implemented"; return nil }

// Reset resets the given buffer with a new stream
func (b *ByteInputAdapter) Reset(stream io.Reader) { _ = "STUB: not implemented"; return }
