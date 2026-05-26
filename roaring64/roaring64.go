package roaring64

import (
	"io"

	"github.com/RoaringBitmap/roaring/v2"
)

const (
	serialCookieNoRunContainer = 12346 // only arrays and bitmaps
	serialCookie               = 12347 // runs, arrays, and bitmaps
)

// Bitmap represents a compressed bitmap where you can add integers.
type Bitmap struct {
	highlowcontainer roaringArray64
}

// ToBase64 serializes a bitmap as Base64
func (rb *Bitmap) ToBase64() (string, error) { _ = "STUB: not implemented"; return "", nil }

// FromBase64 deserializes a bitmap from Base64
func (rb *Bitmap) FromBase64(str string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// ToBytes returns an array of bytes corresponding to what is written
// when calling WriteTo
func (rb *Bitmap) ToBytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// WriteTo writes a serialized version of this bitmap to stream.
// The format is compatible with other 64-bit RoaringBitmap
// implementations (Java, Go, C++) and it has a specification :
// https://github.com/RoaringBitmap/RoaringFormatSpec#extention-for-64-bit-implementations
func (rb *Bitmap) WriteTo(stream io.Writer) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// FromUnsafeBytes reads a serialized version of this bitmap from the byte buffer without copy.
// It is the caller's responsibility to ensure that the input data is not modified and remains valid for the entire lifetime of this bitmap.
// This method avoids small allocations but holds references to the input data buffer. It is GC-friendly, but it may consume more memory eventually.
func (rb *Bitmap) FromUnsafeBytes(data []byte) (p int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadFrom reads a serialized version of this bitmap from stream.
// The format is compatible with other 64-bit RoaringBitmap
// implementations (Java, Go, C++) and it has a specification :
// https://github.com/RoaringBitmap/RoaringFormatSpec#extention-for-64-bit-implementations
func (rb *Bitmap) ReadFrom(stream io.Reader) (p int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// MarshalBinary implements the encoding.BinaryMarshaler interface for the bitmap
// (same as ToBytes)
func (rb *Bitmap) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for the bitmap
		nil
}

func (rb *Bitmap) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

// RunOptimize attempts to further compress the runs of consecutive values found in the bitmap
func (rb *Bitmap) RunOptimize() { _ = "STUB: not implemented"; return }

// HasRunCompression returns true if the bitmap benefits from run compression
func (rb *Bitmap) HasRunCompression() bool { _ = "STUB: not implemented"; return false }

// NewBitmap creates a new empty Bitmap (see also New)
func NewBitmap() *Bitmap {
	_ = "STUB: not implemented"

	// New creates a new empty Bitmap (same as NewBitmap)
	return nil
}

func New() *Bitmap {
	_ = "STUB: not implemented"

	// Clear resets the Bitmap to be logically empty, but may retain
	// some memory allocations that may speed up future operations
	return nil
}

func (rb *Bitmap) Clear() { _ = "STUB: not implemented"; return }

// ToArray creates a new slice containing all of the integers stored in the Bitmap in sorted order
func (rb *Bitmap) ToArray() []uint64 { _ = "STUB: not implemented"; return nil }

// GetSizeInBytes estimates the memory usage of the Bitmap. Note that this
// might differ slightly from the amount of bytes required for persistent storage
func (rb *Bitmap) GetSizeInBytes() uint64 { _ = "STUB: not implemented"; return 0 }

// String creates a string representation of the Bitmap
func (rb *Bitmap) String() string {
	_ = "STUB: not implemented"
	// inspired by https://github.com/fzandona/goroar/
	return ""
}

// to avoid exhausting the memory

// Iterator creates a new IntPeekable to iterate over the integers contained in the bitmap, in sorted order;
// the iterator becomes invalid if the bitmap is modified (e.g., with Add or Remove).
func (rb *Bitmap) Iterator() IntPeekable64 {
	_ = "STUB: not implemented"
	return *

	// ReverseIterator creates a new IntIterable to iterate over the integers contained in the bitmap, in sorted order;
	// the iterator becomes invalid if the bitmap is modified (e.g., with Add or Remove).
	new(IntPeekable64)
}

func (rb *Bitmap) ReverseIterator() IntIterable64 {
	_ = "STUB: not implemented"
	return *new(IntIterable64)
}

// ManyIterator creates a new ManyIntIterable to iterate over the integers contained in the bitmap, in sorted order;
// the iterator becomes invalid if the bitmap is modified (e.g., with Add or Remove).
func (rb *Bitmap) ManyIterator() ManyIntIterable64 {
	_ = "STUB: not implemented"
	return *new(ManyIntIterable64)
}

// Clone creates a copy of the Bitmap
func (rb *Bitmap) Clone() *Bitmap { _ = "STUB: not implemented"; return nil }

// Minimum get the smallest value stored in this roaring bitmap, assumes that it is not empty
func (rb *Bitmap) Minimum() uint64 { _ = "STUB: not implemented"; return 0 }

// Maximum get the largest value stored in this roaring bitmap, assumes that it is not empty
func (rb *Bitmap) Maximum() uint64 { _ = "STUB: not implemented"; return 0 }

// Contains returns true if the integer is contained in the bitmap
func (rb *Bitmap) Contains(x uint64) bool { _ = "STUB: not implemented"; return false }

// ContainsInt returns true if the integer is contained in the bitmap (this is a convenience method, the parameter is casted to uint64 and Contains is called)
func (rb *Bitmap) ContainsInt(x int) bool { _ = "STUB: not implemented"; return false }

// Equals returns true if the two bitmaps contain the same integers
func (rb *Bitmap) Equals(srb *Bitmap) bool { _ = "STUB: not implemented"; return false }

// Add the integer x to the bitmap
func (rb *Bitmap) Add(x uint64) { _ = "STUB: not implemented"; return }

// CheckedAdd adds the integer x to the bitmap and return true  if it was added (false if the integer was already present)
func (rb *Bitmap) CheckedAdd(x uint64) bool { _ = "STUB: not implemented"; return false }

// AddInt adds the integer x to the bitmap (convenience method: the parameter is casted to uint64 and we call Add)
func (rb *Bitmap) AddInt(x int) {
	_ = "STUB: not implemented"

	// Remove the integer x from the bitmap
	return
}

func (rb *Bitmap) Remove(x uint64) { _ = "STUB: not implemented"; return }

// CheckedRemove removes the integer x from the bitmap and return true if the integer was effectively remove (and false if the integer was not present)
func (rb *Bitmap) CheckedRemove(x uint64) bool { _ = "STUB: not implemented"; return false }

// IsEmpty returns true if the Bitmap is empty (it is faster than doing (GetCardinality() == 0))
func (rb *Bitmap) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// GetCardinality returns the number of integers contained in the bitmap
func (rb *Bitmap) GetCardinality() uint64 { _ = "STUB: not implemented"; return 0 }

// Rank returns the number of integers that are smaller or equal to x (Rank(infinity) would be GetCardinality())
func (rb *Bitmap) Rank(x uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// Select returns the xth integer in the bitmap
func (rb *Bitmap) Select(x uint64) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// And computes the intersection between two bitmaps and stores the result in the current bitmap
func (rb *Bitmap) And(x2 *Bitmap) { _ = "STUB: not implemented"; return }

// s1 > s2

// OrCardinality returns the cardinality of the union between two bitmaps, bitmaps are not modified
func (rb *Bitmap) OrCardinality(x2 *Bitmap) uint64 { _ = "STUB: not implemented"; return 0 }

// TODO: could be faster if we did not have to materialize the container

// AndCardinality returns the cardinality of the intersection between two bitmaps, bitmaps are not modified
func (rb *Bitmap) AndCardinality(x2 *Bitmap) uint64 { _ = "STUB: not implemented"; return 0 }

// s1 > s2

// Intersects checks whether two bitmap intersects, bitmaps are not modified
func (rb *Bitmap) Intersects(x2 *Bitmap) bool { _ = "STUB: not implemented"; return false }

// s1 > s2

// Xor computes the symmetric difference between two bitmaps and stores the result in the current bitmap
func (rb *Bitmap) Xor(x2 *Bitmap) { _ = "STUB: not implemented"; return }

// TODO: couple be computed in-place for reduced memory usage

// Or computes the union between two bitmaps and stores the result in the current bitmap
func (rb *Bitmap) Or(x2 *Bitmap) { _ = "STUB: not implemented"; return }

// AndNot computes the difference between two bitmaps and stores the result in the current bitmap
func (rb *Bitmap) AndNot(x2 *Bitmap) { _ = "STUB: not implemented"; return }

// s1 > s2

// TODO:implement as a copy

// Or computes the union between two bitmaps and returns the result
func Or(x1, x2 *Bitmap) *Bitmap { _ = "STUB: not implemented"; return nil }

// And computes the intersection between two bitmaps and returns the result
func And(x1, x2 *Bitmap) *Bitmap { _ = "STUB: not implemented"; return nil }

// s1 > s2

// Xor computes the symmetric difference between two bitmaps and returns the result
func Xor(x1, x2 *Bitmap) *Bitmap { _ = "STUB: not implemented"; return nil }

// AndNot computes the difference between two bitmaps and returns the result
func AndNot(x1, x2 *Bitmap) *Bitmap { _ = "STUB: not implemented"; return nil }

// s1 > s2

// AddMany add all of the values in dat
func (rb *Bitmap) AddMany(dat []uint64) { _ = "STUB: not implemented"; return }

// getOrCreateContainer gets the roaring.Bitmap for key hb,
// or creates an *empty* roaring.Bitmap, inserts it to rb.highlowcontainer, and returns the new roaring.Bitmap.
func (rb *Bitmap) getOrCreateContainer(hb uint32) *roaring.Bitmap {
	_ = "STUB: not implemented"
	return nil
}

// BitmapOf generates a new bitmap filled with the specified integers
func BitmapOf(dat ...uint64) *Bitmap { _ = "STUB: not implemented"; return nil }

// Flip negates the bits in the given range (i.e., [rangeStart,rangeEnd)), any integer present in this range and in the bitmap is removed,
// and any integer present in the range and not in the bitmap is added.
func (rb *Bitmap) Flip(rangeStart, rangeEnd uint64) { _ = "STUB: not implemented"; return }

// *think* the range of ones must never be empty.

// FlipInt calls Flip after casting the parameters  (convenience method)
func (rb *Bitmap) FlipInt(rangeStart, rangeEnd int) { _ = "STUB: not implemented"; return }

// AddRange adds the integers in [rangeStart, rangeEnd) to the bitmap.
func (rb *Bitmap) AddRange(rangeStart, rangeEnd uint64) { _ = "STUB: not implemented"; return }

// RemoveRange removes the integers in [rangeStart, rangeEnd) from the bitmap.
func (rb *Bitmap) RemoveRange(rangeStart, rangeEnd uint64) { _ = "STUB: not implemented"; return }

// Flip negates the bits in the given range  (i.e., [rangeStart,rangeEnd)), any integer present in this range and in the bitmap is removed,
// and any integer present in the range and not in the bitmap is added, a new bitmap is returned leaving
// the current bitmap unchanged.
func Flip(rb *Bitmap, rangeStart, rangeEnd uint64) *Bitmap { _ = "STUB: not implemented"; return nil }

// copy the containers before the active area

// *think* the range of ones must never be empty.

// copy the containers after the active area.

// SetCopyOnWrite sets this bitmap to use copy-on-write so that copies are fast and memory conscious
// if the parameter is true, otherwise we leave the default where hard copies are made
// (copy-on-write requires extra care in a threaded context).
// Calling SetCopyOnWrite(true) on a bitmap created with FromBuffer is unsafe.
func (rb *Bitmap) SetCopyOnWrite(val bool) { _ = "STUB: not implemented"; return }

// GetCopyOnWrite gets this bitmap's copy-on-write property
func (rb *Bitmap) GetCopyOnWrite() (val bool) { _ = "STUB: not implemented"; return false }

// CloneCopyOnWriteContainers clones all containers which have
// needCopyOnWrite set to true.
// This can be used to make sure it is safe to munmap a []byte
// that the roaring array may still have a reference to, after
// calling FromBuffer.
// More generally this function is useful if you call FromBuffer
// to construct a bitmap with a backing array buf
// and then later discard the buf array. Note that you should call
// CloneCopyOnWriteContainers on all bitmaps that were derived
// from the 'FromBuffer' bitmap since they map have dependencies
// on the buf array as well.
func (rb *Bitmap) CloneCopyOnWriteContainers() { _ = "STUB: not implemented"; return }

// FlipInt calls Flip after casting the parameters (convenience method)
func FlipInt(bm *Bitmap, rangeStart, rangeEnd int) *Bitmap { _ = "STUB: not implemented"; return nil }

// Stats returns details on container type usage in a Statistics struct.
func (rb *Bitmap) Stats() roaring.Statistics {
	_ = "STUB: not implemented"
	return *new(roaring.Statistics)
}

// GetSerializedSizeInBytes computes the serialized size in bytes
// of the Bitmap. It should correspond to the number
// of bytes written when invoking WriteTo. You can expect
// that this function is much cheaper computationally than WriteTo.
func (rb *Bitmap) GetSerializedSizeInBytes() uint64 { _ = "STUB: not implemented"; return 0 }

func (rb *Bitmap) Validate() error { _ = "STUB: not implemented"; return nil }

// Roaring32AsRoaring64 inserts a 32-bit roaring bitmap into
// a 64-bit roaring bitmap. No copy is made.
func Roaring32AsRoaring64(bm32 *roaring.Bitmap) *Bitmap { _ = "STUB: not implemented"; return nil }

func roaring32AsRoaring64(bm32 *roaring.Bitmap, key uint32) *Bitmap {
	_ = "STUB: not implemented"
	return nil
}
