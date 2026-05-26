// Package roaring is an implementation of Roaring Bitmaps in Go.
// They provide fast compressed bitmap data structures (also called bitset).
// They are ideally suited to represent sets of integers over
// relatively small ranges.
// See http://roaringbitmap.org for details.
package roaring

import (
	"io"

	"github.com/bits-and-blooms/bitset"
)

// Bitmap represents a compressed bitmap where you can add integers.
type Bitmap struct {
	highlowcontainer roaringArray
}

// ToBase64 serializes a bitmap as Base64
func (rb *Bitmap) ToBase64() (string, error) { _ = "STUB: not implemented"; return "", nil }

// FromBase64 deserializes a bitmap from Base64
func (rb *Bitmap) FromBase64(str string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// WriteTo writes a serialized version of this bitmap to stream.
// The format is compatible with other RoaringBitmap
// implementations (Java, C) and is documented here:
// https://github.com/RoaringBitmap/RoaringFormatSpec
func (rb *Bitmap) WriteTo(stream io.Writer) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ToBytes returns an array of bytes corresponding to what is written
// when calling WriteTo
func (rb *Bitmap) ToBytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

const (
	wordSize            = uint64(64)
	log2WordSize        = uint64(6)
	capacity            = ^uint64(0)
	bitmapContainerSize = (1 << 16) / 64 // bitmap size in words
)

// DenseSize returns the size of the bitmap when stored as a dense bitmap.
func (rb *Bitmap) DenseSize() uint64 { _ = "STUB: not implemented"; return 0 }

// ToDense returns a slice of uint64s representing the bitmap as a dense bitmap.
// Useful to convert a roaring bitmap to a format that can be used by other libraries
// like https://github.com/bits-and-blooms/bitset or https://github.com/kelindar/bitmap
func (rb *Bitmap) ToDense() []uint64 { _ = "STUB: not implemented"; return nil }

// FromDense creates a bitmap from a slice of uint64s representing the bitmap as a dense bitmap.
// Useful to convert bitmaps from libraries like https://github.com/bits-and-blooms/bitset or
// https://github.com/kelindar/bitmap into roaring bitmaps fast and with convenience.
//
// This function will not create any run containers, only array and bitmap containers. It's up to
// the caller to call RunOptimize if they want to further compress the runs of consecutive values.
//
// When doCopy is true, the bitmap is copied into a new slice for each bitmap container.
// This is useful when the bitmap is going to be modified after this function returns or if it's
// undesirable to hold references to large bitmaps which the GC would not be able to collect.
// One copy can still happen even when doCopy is false if the bitmap length is not divisible
// by bitmapContainerSize.
//
// See also FromBitSet.
func FromDense(bitmap []uint64, doCopy bool) *Bitmap { _ = "STUB: not implemented"; return nil }

// round up

// FromDense unmarshalls from a slice of uint64s representing the bitmap as a dense bitmap.
// Useful to convert bitmaps from libraries like https://github.com/bits-and-blooms/bitset or
// https://github.com/kelindar/bitmap into roaring bitmaps fast and with convenience.
// Callers are responsible for ensuring that the bitmap is empty before calling this function.
//
// This function will not create any run containers, only array and bitmap containers. It is up to
// the caller to call RunOptimize if they want to further compress the runs of consecutive values.
//
// When doCopy is true, the bitmap is copied into a new slice for each bitmap container.
// This is useful when the bitmap is going to be modified after this function returns or if it's
// undesirable to hold references to large bitmaps which the GC would not be able to collect.
// One copy can still happen even when doCopy is false if the bitmap length is not divisible
// by bitmapContainerSize.
//
// See FromBitSet.
func (rb *Bitmap) FromDense(bitmap []uint64, doCopy bool) { _ = "STUB: not implemented"; return }

// WriteDenseTo writes to a slice of uint64s representing the bitmap as a dense bitmap.
// Callers are responsible for allocating enough space in the bitmap using DenseSize.
// Useful to convert a roaring bitmap to a format that can be used by other libraries
// like https://github.com/bits-and-blooms/bitset or https://github.com/kelindar/bitmap
func (rb *Bitmap) WriteDenseTo(bitmap []uint64) { _ = "STUB: not implemented"; return }

// Checksum computes a hash (FNV-1a) for a bitmap that is suitable for
// using bitmaps as elements in hash sets or as keys in hash maps, as well as
// generally quick comparisons.
func (rb *Bitmap) Checksum() uint64 { _ = "STUB: not implemented"; return 0 }

// Hash the keys (uint16 slice) directly

// Hash low byte first (little endian)

// Hash high byte

// 0 separator

// Hash in little-endian byte order (unrolled loop)

// Hash low byte first (little endian)

// Hash high byte

// Hash start (uint16)

// Hash length (uint16)

// FromUnsafeBytes reads a serialized version of this bitmap from the byte buffer without copy
// (for advanced users only, you must be an expert Go programmer!).
// E.g., you can use this method to read a serialized bitmap from a memory-mapped file written out
// with the WriteTo method.
// The format specification is
// https://github.com/RoaringBitmap/RoaringFormatSpec
// It is the caller's responsibility to ensure that the input data is not modified and remains valid for the entire lifetime of this bitmap.
// This method avoids small allocations but holds references to the input data buffer. It is GC-friendly, but it may consume more memory eventually.
// The containers in the resulting bitmap are immutable containers tied to the provided byte array and they rely on
// copy-on-write which means that modifying them creates copies. Thus FromUnsafeBytes is more likely to be appropriate for read-only use cases,
// when the resulting bitmap can be considered immutable.
//
// See also the FromBuffer function. We recommend benchmarking both functions to determine which one is more suitable for your use case.
// See https://github.com/RoaringBitmap/roaring/pull/395 for more details.
func (rb *Bitmap) FromUnsafeBytes(data []byte, cookieHeader ...byte) (p int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadFrom reads a serialized version of this bitmap from stream.
// E.g., you can use this method to read a serialized bitmap from a file written
// with the WriteTo method.
// The format is compatible with other RoaringBitmap
// implementations (Java, C) and is documented here:
// https://github.com/RoaringBitmap/RoaringFormatSpec
// Since io.Reader is regarded as a stream and cannot be read twice,
// we add cookieHeader to accept the 4-byte data that has been read in roaring64.ReadFrom.
// It is not necessary to pass cookieHeader when call roaring.ReadFrom to read the roaring32 data directly.
func (rb *Bitmap) ReadFrom(reader io.Reader, cookieHeader ...byte) (p int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// MustReadFrom calls ReadFrom internally.
// After deserialization Validate will be called.
// If the Bitmap fails to validate, a panic with the validation error will be thrown
func (rb *Bitmap) MustReadFrom(reader io.Reader, cookieHeader ...byte) (p int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// FromBuffer creates a bitmap from its serialized version stored in buffer (E.g., as written by WriteTo).
//
// The format specification is available here:
// https://github.com/RoaringBitmap/RoaringFormatSpec
//
// The provided byte array (buf) is expected to be a constant.
// The function makes the best effort attempt not to copy data.
// You should take care not to modify buff as it will
// likely result in unexpected program behavior.
//
// Resulting bitmaps are effectively immutable in the following sense:
// a copy-on-write marker is used so that when you modify the resulting
// bitmap, copies of selected data (containers) are made.
// You should *not* change the copy-on-write status of the resulting
// bitmaps (SetCopyOnWrite).
//
// Thus FromBuffer is more likely to be appropriate for read-only use cases,
// when the resulting bitmap can be considered immutable.
//
// If buf becomes unavailable, then a bitmap created with
// FromBuffer would be effectively broken. Furthermore, any
// bitmap derived from this bitmap (e.g., via Or, And) might
// also be broken. Thus, before making buf unavailable, you should
// call CloneCopyOnWriteContainers on all such bitmaps.
//
// See also the FromUnsafeBytes function which can have better performance
// in some cases.
func (rb *Bitmap) FromBuffer(buf []byte) (p int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// RunOptimize attempts to further compress the runs of consecutive values found in the bitmap
func (rb *Bitmap) RunOptimize() { _ = "STUB: not implemented"; return }

// HasRunCompression returns true if the bitmap benefits from run compression
func (rb *Bitmap) HasRunCompression() bool { _ = "STUB: not implemented"; return false }

// MarshalBinary implements the encoding.BinaryMarshaler interface for the bitmap
// (same as ToBytes)
func (rb *Bitmap) MarshalBinary() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// UnmarshalBinary implements the encoding.BinaryUnmarshaler interface for the bitmap
		nil
}

func (rb *Bitmap) UnmarshalBinary(data []byte) error { _ = "STUB: not implemented"; return nil }

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

// ToBitSet copies the content of the RoaringBitmap into a bitset.BitSet instance
func (rb *Bitmap) ToBitSet() *bitset.BitSet { _ = "STUB: not implemented"; return nil }

// FromBitSet creates a new RoaringBitmap from a bitset.BitSet instance
func FromBitSet(bitset *bitset.BitSet) *Bitmap { _ = "STUB: not implemented"; return nil }

// ToArray creates a new slice containing all of the integers stored in the Bitmap in sorted order
func (rb *Bitmap) ToArray() []uint32 { _ = "STUB: not implemented"; return nil }

func (rb *Bitmap) toArray(array *[]uint32) *[]uint32 { _ = "STUB: not implemented"; return nil }

// ToExistingArray stores all of the integers stored in the Bitmap in sorted order in the
// slice that is given to ToExistingArray. It is the callers duty to make sure the slice
// has the right size.
func (rb *Bitmap) ToExistingArray(array *[]uint32) *[]uint32 { _ = "STUB: not implemented"; return nil }

// GetSizeInBytes estimates the memory usage of the Bitmap. Note that this
// might differ slightly from the amount of bytes required for persistent storage
func (rb *Bitmap) GetSizeInBytes() uint64 { _ = "STUB: not implemented"; return 0 }

// GetSerializedSizeInBytes computes the serialized size in bytes
// of the Bitmap. It should correspond to the
// number of bytes written when invoking WriteTo. You can expect
// that this function is much cheaper computationally than WriteTo.
func (rb *Bitmap) GetSerializedSizeInBytes() uint64 { _ = "STUB: not implemented"; return 0 }

// BoundSerializedSizeInBytes returns an upper bound on the serialized size in bytes
// assuming that one wants to store "cardinality" integers in [0, universe_size)
func BoundSerializedSizeInBytes(cardinality uint64, universeSize uint64) uint64 {
	_ = "STUB: not implemented"
	return 0
}

// we cannot have more containers than we have values

// IntIterable allows you to iterate over the values in a Bitmap
type IntIterable interface {
	HasNext() bool
	Next() uint32
}

// IntPeekable allows you to look at the next value without advancing and
// advance as long as the next value is smaller than minval
type IntPeekable interface {
	IntIterable
	// PeekNext peeks the next value without advancing the iterator
	PeekNext() uint32
	// AdvanceIfNeeded advances as long as the next value is smaller than minval
	AdvanceIfNeeded(minval uint32)
}

type intIterator struct {
	pos              int
	hs               uint32
	iter             shortPeekable
	highlowcontainer *roaringArray

	// These embedded iterators per container type help reduce load in the GC.
	// This way, instead of making up-to 64k allocations per full iteration
	// we get a single allocation and simply reinitialize the appropriate
	// iterator and point to it in the generic `iter` member on each key bound.
	shortIter  shortIterator
	runIter    runIterator16
	bitmapIter bitmapContainerShortIterator
}

// HasNext returns true if there are more integers to iterate over
func (ii *intIterator) HasNext() bool { _ = "STUB: not implemented"; return false }

func (ii *intIterator) init() {
	if ii.highlowcontainer.size() > ii.pos {
		ii.hs = uint32(ii.highlowcontainer.getKeyAtIndex(ii.pos)) << 16
		c := ii.highlowcontainer.getContainerAtIndex(ii.pos)
		switch t := c.(type) {
		case *arrayContainer:
			ii.shortIter = shortIterator{t.content, 0}
			ii.iter = &ii.shortIter
		case *runContainer16:
			ii.runIter = runIterator16{rc: t, curIndex: 0, curPosInIndex: 0}
			ii.iter = &ii.runIter
		case *bitmapContainer:
			ii.bitmapIter = bitmapContainerShortIterator{t, t.NextSetBit(0)}
			ii.iter = &ii.bitmapIter
		}
	}
}

// Next returns the next integer
func (ii *intIterator) Next() uint32 { _ = "STUB: not implemented"; return 0 }

// PeekNext peeks the next value without advancing the iterator
func (ii *intIterator) PeekNext() uint32 { _ = "STUB: not implemented"; return 0 }

// AdvanceIfNeeded advances as long as the next value is smaller than minval
func (ii *intIterator) AdvanceIfNeeded(minval uint32) { _ = "STUB: not implemented"; return }

// IntIterator is meant to allow you to iterate through the values of a bitmap, see Initialize(a *Bitmap)
type IntIterator = intIterator

// Initialize configures the existing iterator so that it can iterate through the values of
// the provided bitmap.
// The iteration results are undefined if the bitmap is modified (e.g., with Add or Remove).
func (ii *intIterator) Initialize(a *Bitmap) { _ = "STUB: not implemented"; return }

type intReverseIterator struct {
	pos              int
	hs               uint32
	iter             shortIterable
	highlowcontainer *roaringArray

	shortIter  reverseIterator
	runIter    runReverseIterator16
	bitmapIter reverseBitmapContainerShortIterator
}

// HasNext returns true if there are more integers to iterate over
func (ii *intReverseIterator) HasNext() bool { _ = "STUB: not implemented"; return false }

func (ii *intReverseIterator) init() {
	if ii.pos >= 0 {
		ii.hs = uint32(ii.highlowcontainer.getKeyAtIndex(ii.pos)) << 16
		c := ii.highlowcontainer.getContainerAtIndex(ii.pos)
		switch t := c.(type) {
		case *arrayContainer:
			ii.shortIter = reverseIterator{t.content, len(t.content) - 1}
			ii.iter = &ii.shortIter
		case *runContainer16:
			index := len(t.iv) - 1
			pos := uint16(0)

			if index >= 0 {
				pos = t.iv[index].length
			}

			ii.runIter = runReverseIterator16{rc: t, curIndex: index, curPosInIndex: pos}
			ii.iter = &ii.runIter
		case *bitmapContainer:
			pos := -1
			if t.cardinality > 0 {
				pos = int(t.maximum())
			}
			ii.bitmapIter = reverseBitmapContainerShortIterator{t, pos}
			ii.iter = &ii.bitmapIter
		}
	} else {
		ii.iter = nil
	}
}

// Next returns the next integer
func (ii *intReverseIterator) Next() uint32 { _ = "STUB: not implemented"; return 0 }

// IntReverseIterator is meant to allow you to iterate through the values of a bitmap, see Initialize(a *Bitmap)
type IntReverseIterator = intReverseIterator

// Initialize configures the existing iterator so that it can iterate through the values of
// the provided bitmap.
// The iteration results are undefined if the bitmap is modified (e.g., with Add or Remove).
func (ii *intReverseIterator) Initialize(a *Bitmap) { _ = "STUB: not implemented"; return }

// ManyIntIterable allows you to iterate over the values in a Bitmap
type ManyIntIterable interface {
	// NextMany fills buf up with values, returns how many values were returned
	NextMany(buf []uint32) int
	// NextMany64 fills up buf with 64 bit values, uses hs as a mask (OR), returns how many values were returned
	NextMany64(hs uint64, buf []uint64) int
}

type manyIntIterator struct {
	pos              int
	hs               uint32
	iter             manyIterable
	highlowcontainer *roaringArray

	shortIter  shortIterator
	runIter    runIterator16
	bitmapIter bitmapContainerManyIterator
}

func (ii *manyIntIterator) init() {
	if ii.highlowcontainer.size() > ii.pos {
		ii.hs = uint32(ii.highlowcontainer.getKeyAtIndex(ii.pos)) << 16
		c := ii.highlowcontainer.getContainerAtIndex(ii.pos)
		switch t := c.(type) {
		case *arrayContainer:
			ii.shortIter = shortIterator{t.content, 0}
			ii.iter = &ii.shortIter
		case *runContainer16:
			ii.runIter = runIterator16{rc: t, curIndex: 0, curPosInIndex: 0}
			ii.iter = &ii.runIter
		case *bitmapContainer:
			ii.bitmapIter = bitmapContainerManyIterator{t, -1, 0}
			ii.iter = &ii.bitmapIter
		}
	} else {
		ii.iter = nil
	}
}

func (ii *manyIntIterator) NextMany(buf []uint32) int { _ = "STUB: not implemented"; return 0 }

func (ii *manyIntIterator) NextMany64(hs64 uint64, buf []uint64) int {
	_ = "STUB: not implemented"
	return 0
}

// ManyIntIterator is meant to allow you to iterate through the values of a bitmap, see Initialize(a *Bitmap)
type ManyIntIterator = manyIntIterator

// Initialize configures the existing iterator so that it can iterate through the values of
// the provided bitmap.
// The iteration results are undefined if the bitmap is modified (e.g., with Add or Remove).
func (ii *manyIntIterator) Initialize(a *Bitmap) { _ = "STUB: not implemented"; return }

type unsetIterator struct {
	containerIndex   int
	nextKey          int
	hs               uint32
	iter             shortPeekable
	highlowcontainer *roaringArray

	arrayUnsetIter    arrayContainerUnsetIterator
	runUnsetIter      runUnsetIterator16
	bitmapUnsetIter   bitmapContainerUnsetIterator
	emptyContainerVal uint16

	start, end uint64
}

// HasNext returns true if there are more integers to iterate over
func (iui *unsetIterator) HasNext() bool {
	_ = "STUB: not implemented"
	// Skip containers that have no unset bits in our range
	return false
}

// We're in an empty container gap, which has unset bits

// Move to next container

// Check if next value is within range

// Current container has no more unset bits in range, move to next

func (iui *unsetIterator) init() {
	// Check if we've gone past the end range
	if uint64(iui.nextKey)<<16 >= iui.end {
		iui.iter = nil
		return
	}

	// Check if we're in an empty container gap
	if iui.containerIndex >= iui.highlowcontainer.size() ||
		iui.highlowcontainer.getKeyAtIndex(iui.containerIndex) > uint16(iui.nextKey) {
		// We're in a gap - iterate through empty container
		iui.emptyContainerVal = 0
		// If this container overlaps with start, advance to start
		if uint64(iui.nextKey)<<16 < iui.start && iui.start < uint64(iui.nextKey+1)<<16 {
			iui.emptyContainerVal = uint16(iui.start)
		}
		iui.iter = nil
		return
	}

	// We're in an actual container
	iui.hs = uint32(iui.nextKey) << 16
	c := iui.highlowcontainer.getContainerAtIndex(iui.containerIndex)
	switch t := c.(type) {
	case *arrayContainer:
		iui.arrayUnsetIter = *newArrayContainerUnsetIterator(t.content)
		iui.iter = &iui.arrayUnsetIter
	case *runContainer16:
		iui.runUnsetIter = *t.newRunUnsetIterator16()
		iui.iter = &iui.runUnsetIter
	case *bitmapContainer:
		iui.bitmapUnsetIter = *newBitmapContainerUnsetIterator(t)
		iui.iter = &iui.bitmapUnsetIter
	}

	// If this container overlaps with start, advance to the low bits of start
	if uint64(iui.nextKey)<<16 < iui.start && iui.start < uint64(iui.nextKey+1)<<16 {
		iui.iter.advanceIfNeeded(uint16(iui.start))
	}
}

// Next returns the next integer
func (iui *unsetIterator) Next() uint32 { _ = "STUB: not implemented"; return 0 }

// We're in an empty container gap

// Wrapped around or reached end, move to next container

// PeekNext peeks the next value without advancing the iterator
func (iui *unsetIterator) PeekNext() uint32 { _ = "STUB: not implemented"; return 0 }

// AdvanceIfNeeded advances as long as the next value is smaller than minval
func (iui *unsetIterator) AdvanceIfNeeded(minval uint32) { _ = "STUB: not implemented"; return }

// Find the next container that matches or exceeds nextKey

// Initialize configures the unset iterator to iterate over values in [start, end) that are not in the bitmap
func (iui *unsetIterator) Initialize(a *Bitmap, start, end uint64) {
	_ = "STUB: not implemented"
	return
}

// Find the first container that matches or exceeds the start key

// String creates a string representation of the Bitmap
func (rb *Bitmap) String() string {
	_ = "STUB: not implemented"
	// inspired by https://github.com/fzandona/goroar/
	return ""
}

// to avoid exhausting the memory

// Iterate iterates over the bitmap, calling the given callback with each value in the bitmap.  If the callback returns
// false, the iteration is halted.
// The iteration results are undefined if the bitmap is modified (e.g., with Add or Remove).
// There is no guarantee as to what order the values will be iterated.
func (rb *Bitmap) Iterate(cb func(x uint32) bool) { _ = "STUB: not implemented"; return }

// This is hacky but it avoids allocations from invoking an interface method with a closure

// Iterator creates a new IntPeekable to iterate over the integers contained in the bitmap, in sorted order;
// the iterator becomes invalid if the bitmap is modified (e.g., with Add or Remove).
func (rb *Bitmap) Iterator() IntPeekable { _ = "STUB: not implemented"; return *new(IntPeekable) }

// ReverseIterator creates a new IntIterable to iterate over the integers contained in the bitmap, in sorted order;
// the iterator becomes invalid if the bitmap is modified (e.g., with Add or Remove).
func (rb *Bitmap) ReverseIterator() IntIterable {
	_ = "STUB: not implemented"
	return *new(IntIterable)
}

// ManyIterator creates a new ManyIntIterable to iterate over the integers contained in the bitmap, in sorted order;
// the iterator becomes invalid if the bitmap is modified (e.g., with Add or Remove).
func (rb *Bitmap) ManyIterator() ManyIntIterable {
	_ = "STUB: not implemented"
	return *new(ManyIntIterable)
}

// UnsetIterator creates a new IntPeekable to iterate over values in the range [start, end) that are NOT contained in the bitmap.
// The iterator becomes invalid if the bitmap is modified (e.g., with Add or Remove).
func (rb *Bitmap) UnsetIterator(start, end uint64) IntPeekable {
	_ = "STUB: not implemented"
	return *new(IntPeekable)
}

// Clone creates a copy of the Bitmap
func (rb *Bitmap) Clone() *Bitmap { _ = "STUB: not implemented"; return nil }

// Minimum get the smallest value stored in this roaring bitmap, assumes that it is not empty
func (rb *Bitmap) Minimum() uint32 { _ = "STUB: not implemented"; return 0 }

// Maximum get the largest value stored in this roaring bitmap, assumes that it is not empty
func (rb *Bitmap) Maximum() uint32 { _ = "STUB: not implemented"; return 0 }

// Contains returns true if the integer is contained in the bitmap
func (rb *Bitmap) Contains(x uint32) bool { _ = "STUB: not implemented"; return false }

// ContainsInt returns true if the integer is contained in the bitmap (this is a convenience method, the parameter is casted to uint32 and Contains is called)
func (rb *Bitmap) ContainsInt(x int) bool { _ = "STUB: not implemented"; return false }

// Equals returns true if the two bitmaps contain the same integers
func (rb *Bitmap) Equals(o interface{}) bool { _ = "STUB: not implemented"; return false }

// AddOffset adds the value 'offset' to each and every value in a bitmap, generating a new bitmap in the process
func AddOffset(x *Bitmap, offset uint32) (answer *Bitmap) { _ = "STUB: not implemented"; return nil }

// AddOffset64 adds the value 'offset' to each and every value in a bitmap, generating a new bitmap in the process
// If offset + element is outside of the range [0,2^32), that the element will be dropped
func AddOffset64(x *Bitmap, offset int64) (answer *Bitmap) {
	_ = "STUB: not implemented"
	// we need "offset" to be a long because we want to support values
	// between -0xFFFFFFFF up to +-0xFFFFFFFF
	return nil
}

// Add the integer x to the bitmap
func (rb *Bitmap) Add(x uint32) { _ = "STUB: not implemented"; return }

// add the integer x to the bitmap, return the container and its index
func (rb *Bitmap) addwithptr(x uint32) (int, container) {
	_ = "STUB: not implemented"
	return 0, *new(container)
}

// CheckedAdd adds the integer x to the bitmap and return true  if it was added (false if the integer was already present)
func (rb *Bitmap) CheckedAdd(x uint32) bool {
	_ = "STUB: not implemented"
	// TODO: add unit tests for this method
	return false
}

// AddInt adds the integer x to the bitmap (convenience method: the parameter is casted to uint32 and we call Add)
func (rb *Bitmap) AddInt(x int) {
	_ = "STUB: not implemented"

	// Remove the integer x from the bitmap
	return
}

func (rb *Bitmap) Remove(x uint32) { _ = "STUB: not implemented"; return }

// CheckedRemove removes the integer x from the bitmap and return true if the integer was effectively removed (and false if the integer was not present)
func (rb *Bitmap) CheckedRemove(x uint32) bool {
	_ = "STUB: not implemented"
	// TODO: add unit tests for this method
	return false
}

// IsEmpty returns true if the Bitmap is empty (it is faster than doing (GetCardinality() == 0))
func (rb *Bitmap) IsEmpty() bool { _ = "STUB: not implemented"; return false }

// GetCardinality returns the number of integers contained in the bitmap
func (rb *Bitmap) GetCardinality() uint64 { _ = "STUB: not implemented"; return 0 }

// Rank returns the number of integers that are smaller or equal to x (Rank(infinity) would be GetCardinality()).
// If you pass the smallest value, you get the value 1. If you pass a value that is smaller than the smallest
// value, you get 0. Note that this function differs in convention from the Select function since it
// return 1 and not 0 on the smallest value.
func (rb *Bitmap) Rank(x uint32) uint64 { _ = "STUB: not implemented"; return 0 }

// CardinalityInRange returns the number of integers that are in the half-open range [start, end).
// It is equivalent to Rank(uint32(end-1)) - Rank(uint32(start-1)) for start > 0,
// but is optimized to only scan containers that overlap the range, making it
// O(k) in the number of containers spanned by [start, end) rather than O(n)
// in total containers. The parameter type is uint64 to allow end = 1<<32
// (the full 32-bit range).
func (rb *Bitmap) CardinalityInRange(start, end uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// end-1 is the last included value

// Binary-search to find the first container index >= hbStart.

// insertion point

// Handle the case where start and end are in the same container.

// Handle the first container (may be partial).

// Binary-search to find the last container index <= hbEnd.

// index of the last container with key < hbEnd

// Tight loop over middle containers — no per-iteration key comparisons.

// this is the end container, handled below

// Handle the last container (may be partial).

// Select returns the xth integer in the bitmap. If you pass 0, you get
// the smallest element. Note that this function differs in convention from
// the Rank function which returns 1 on the smallest value.
func (rb *Bitmap) Select(x uint32) (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

// And computes the intersection between two bitmaps and stores the result in the current bitmap
func (rb *Bitmap) And(x2 *Bitmap) { _ = "STUB: not implemented"; return }

// s1 > s2

// OrCardinality  returns the cardinality of the union between two bitmaps, bitmaps are not modified
func (rb *Bitmap) OrCardinality(x2 *Bitmap) uint64 { _ = "STUB: not implemented"; return 0 }

// TODO: could be faster if we did not have to materialize the container

// AndCardinality returns the cardinality of the intersection between two bitmaps, bitmaps are not modified
func (rb *Bitmap) AndCardinality(x2 *Bitmap) uint64 { _ = "STUB: not implemented"; return 0 }

// s1 > s2

// IntersectsWithInterval checks whether a bitmap 'rb' and an open interval '[x,y)' intersect.
func (rb *Bitmap) IntersectsWithInterval(x, y uint64) bool { _ = "STUB: not implemented"; return false }

// Intersects checks whether two bitmap intersects, bitmaps are not modified
func (rb *Bitmap) Intersects(x2 *Bitmap) bool { _ = "STUB: not implemented"; return false }

// s1 > s2

// Xor computes the symmetric difference between two bitmaps and stores the result in the current bitmap
func (rb *Bitmap) Xor(x2 *Bitmap) { _ = "STUB: not implemented"; return }

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
func (rb *Bitmap) AddMany(dat []uint32) { _ = "STUB: not implemented"; return }

// BitmapOf generates a new bitmap filled with the specified integers
func BitmapOf(dat ...uint32) *Bitmap { _ = "STUB: not implemented"; return nil }

// Flip negates the bits in the given range (i.e., [rangeStart,rangeEnd)), any integer present in this range and in the bitmap is removed,
// and any integer present in the range and not in the bitmap is added.
// The function uses 64-bit parameters even though a Bitmap stores 32-bit values because it is allowed and meaningful to use [0,uint64(0x100000000)) as a range
// while uint64(0x100000000) cannot be represented as a 32-bit value.
func (rb *Bitmap) Flip(rangeStart, rangeEnd uint64) { _ = "STUB: not implemented"; return }

// *think* the range of ones must never be
// empty.

// FlipInt calls Flip after casting the parameters  (convenience method)
func (rb *Bitmap) FlipInt(rangeStart, rangeEnd int) { _ = "STUB: not implemented"; return }

// AddRange adds the integers in [rangeStart, rangeEnd) to the bitmap.
// The function uses 64-bit parameters even though a Bitmap stores 32-bit values because it is allowed and meaningful to use [0,uint64(0x100000000)) as a range
// while uint64(0x100000000) cannot be represented as a 32-bit value.
func (rb *Bitmap) AddRange(rangeStart, rangeEnd uint64) { _ = "STUB: not implemented"; return }

// *think* the range of ones must never be
// empty.

// RemoveRange removes the integers in [rangeStart, rangeEnd) from the bitmap.
// The function uses 64-bit parameters even though a Bitmap stores 32-bit values because it is allowed and meaningful to use [0,uint64(0x100000000)) as a range
// while uint64(0x100000000) cannot be represented as a 32-bit value.
func (rb *Bitmap) RemoveRange(rangeStart, rangeEnd uint64) { _ = "STUB: not implemented"; return }

// logically, we should assume that the user wants to
// remove all values from rangeStart to infinity
// see https://github.com/RoaringBitmap/roaring/issues/141

// Flip negates the bits in the given range  (i.e., [rangeStart,rangeEnd)), any integer present in this range and in the bitmap is removed,
// and any integer present in the range and not in the bitmap is added, a new bitmap is returned leaving
// the current bitmap unchanged.
// The function uses 64-bit parameters even though a Bitmap stores 32-bit values because it is allowed and meaningful to use [0,uint64(0x100000000)) as a range
// while uint64(0x100000000) cannot be represented as a 32-bit value.
func Flip(bm *Bitmap, rangeStart, rangeEnd uint64) *Bitmap { _ = "STUB: not implemented"; return nil }

// copy the containers before the active area

// *think* the range of ones must never be
// empty.

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

// NextValue returns the next largest value in the bitmap, or -1
// if none is present. This function should not be used inside
// a performance-sensitive loop: prefer iterators if
// performance is a concern.
func (rb *Bitmap) NextValue(target uint32) int64 { _ = "STUB: not implemented"; return 0 }

// if containerKey > orginalKey then we are past the container which mapped to the orignal key
// in that case we can just return the minimum from that container

// PreviousValue returns the previous largest value in the bitmap, or -1
// if none is present. This function should not be used inside
// a performance-sensitive loop: prefer iterators if
// performance is a concern.
func (rb *Bitmap) PreviousValue(target uint32) int64 { _ = "STUB: not implemented"; return 0 }

// target absent, key of first container after target too high

// if containerKey > originalKey then we are past the container which mapped to the original key
// in that case we can just return the minimum from that container

// NextAbsentValue returns the next largest missing value in the bitmap, or -1
// if none is present. This function should not be used inside
// a performance-sensitive loop: prefer iterators if
// performance is a concern.
func (rb *Bitmap) NextAbsentValue(target uint32) int64 { _ = "STUB: not implemented"; return 0 }

// if we are here it means no container found, just return the target

// target is less than the start of the keyspace start
// that means target cannot be in the keyspace

// There is a gap between keys
// Just increment the current key and shift to get HoB

// PreviousAbsentValue returns the previous largest missing value in the bitmap, or -1
// if none is present. This function should not be used inside
// a performance-sensitive loop: prefer iterators if
// performance is a concern.
func (rb *Bitmap) PreviousAbsentValue(target uint32) int64 { _ = "STUB: not implemented"; return 0 }

// if we are here it means no container found, just return the target

// if we are here it means no container found, just return the target

// target is less than the start of the keyspace start
// that means target cannot be in the keyspace

// OR panic, Java panics

// There is a gap between keys, eg missing container
// Just decrement the current key and shift to get HoB of the missing container

// FlipInt calls Flip after casting the parameters (convenience method)
func FlipInt(bm *Bitmap, rangeStart, rangeEnd int) *Bitmap { _ = "STUB: not implemented"; return nil }

// Statistics provides details on the container types in use.
type Statistics struct {
	Cardinality uint64
	Containers  uint64

	ArrayContainers      uint64
	ArrayContainerBytes  uint64
	ArrayContainerValues uint64

	BitmapContainers      uint64
	BitmapContainerBytes  uint64
	BitmapContainerValues uint64

	RunContainers      uint64
	RunContainerBytes  uint64
	RunContainerValues uint64
}

// Stats returns details on container type usage in a Statistics struct.
func (rb *Bitmap) Stats() Statistics { _ = "STUB: not implemented"; return *new(Statistics) }

// Describe prints a description of the bitmap's containers to stdout
func (rb *Bitmap) Describe() { _ = "STUB: not implemented"; return }

// Validate checks if the bitmap is internally consistent.
// You may call it after deserialization to check that the bitmap is valid.
// This function returns an error if the bitmap is invalid, nil otherwise.
func (rb *Bitmap) Validate() error { _ = "STUB: not implemented"; return nil }
