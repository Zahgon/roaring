package roaring64

import (
	"github.com/RoaringBitmap/roaring/v2"
)

// IntIterable64 allows you to iterate over the values in a Bitmap
type IntIterable64 interface {
	HasNext() bool
	Next() uint64
}

// IntPeekable64 allows you to look at the next value without advancing and
// advance as long as the next value is smaller than minval
type IntPeekable64 interface {
	IntIterable64
	// PeekNext peeks the next value without advancing the iterator
	PeekNext() uint64
	// AdvanceIfNeeded advances as long as the next value is smaller than minval
	AdvanceIfNeeded(minval uint64)
}

type intIterator struct {
	pos              int
	hs               uint64
	iter             roaring.IntPeekable
	highlowcontainer *roaringArray64

	// These embedded iterators per container type help reduce load in the GC.
	// This way, instead of making up-to 4 billion allocations per full iteration
	// we get a single allocation and simply reinitialize the embedded iterator
	// and point to it in the generic `iter` member on each key bound.
	bitmapIter roaring.IntIterator
}

// HasNext returns true if there are more integers to iterate over
func (ii *intIterator) HasNext() bool { _ = "STUB: not implemented"; return false }

func (ii *intIterator) init() {
	if ii.highlowcontainer.size() > ii.pos {
		ii.hs = uint64(ii.highlowcontainer.getKeyAtIndex(ii.pos)) << 32
		ii.bitmapIter.Initialize(ii.highlowcontainer.getContainerAtIndex(ii.pos))
		ii.iter = &ii.bitmapIter
	}
}

// Next returns the next integer
func (ii *intIterator) Next() uint64 { _ = "STUB: not implemented"; return 0 }

// PeekNext peeks the next value without advancing the iterator
func (ii *intIterator) PeekNext() uint64 { _ = "STUB: not implemented"; return 0 }

// AdvanceIfNeeded advances as long as the next value is smaller than minval
func (ii *intIterator) AdvanceIfNeeded(minval uint64) { _ = "STUB: not implemented"; return }

// IntIterator64 is meant to allow you to iterate through the values of a bitmap, see Initialize(a *Bitmap)
type IntIterator64 = intIterator

// Initialize configures the existing iterator so that it can iterate through the values of
// the provided bitmap.
// The iteration results are undefined if the bitmap is modified (e.g., with Add or Remove).
func (ii *intIterator) Initialize(a *Bitmap) { _ = "STUB: not implemented"; return }

func newIntIterator(a *Bitmap) *intIterator { _ = "STUB: not implemented"; return nil }

type intReverseIterator struct {
	pos              int
	hs               uint64
	iter             roaring.IntIterable
	highlowcontainer *roaringArray64

	// Stack-allocated embedded iterator to reduce GC pressure.
	bitmapIter roaring.IntReverseIterator
}

// HasNext returns true if there are more integers to iterate over
func (ii *intReverseIterator) HasNext() bool { _ = "STUB: not implemented"; return false }

func (ii *intReverseIterator) init() {
	if ii.pos >= 0 {
		ii.hs = uint64(ii.highlowcontainer.getKeyAtIndex(ii.pos)) << 32
		ii.bitmapIter.Initialize(ii.highlowcontainer.getContainerAtIndex(ii.pos))
		ii.iter = &ii.bitmapIter
	} else {
		ii.iter = nil
	}
}

// Next returns the next integer
func (ii *intReverseIterator) Next() uint64 { _ = "STUB: not implemented"; return 0 }

// IntReverseIterator64 is meant to allow you to iterate through the values of a bitmap in reverse, see Initialize(a *Bitmap)
type IntReverseIterator64 = intReverseIterator

// Initialize configures the existing iterator so that it can iterate through the values of
// the provided bitmap in reverse.
// The iteration results are undefined if the bitmap is modified (e.g., with Add or Remove).
func (ii *intReverseIterator) Initialize(a *Bitmap) { _ = "STUB: not implemented"; return }

func newIntReverseIterator(a *Bitmap) *intReverseIterator { _ = "STUB: not implemented"; return nil }

// ManyIntIterable64 allows you to iterate over the values in a Bitmap
type ManyIntIterable64 interface {
	// pass in a buffer to fill up with values, returns how many values were returned
	NextMany([]uint64) int
}

type manyIntIterator struct {
	pos              int
	hs               uint64
	iter             roaring.ManyIntIterable
	highlowcontainer *roaringArray64

	// Stack-allocated embedded iterator to reduce GC pressure.
	bitmapIter roaring.ManyIntIterator
}

func (ii *manyIntIterator) init() {
	if ii.highlowcontainer.size() > ii.pos {
		ii.hs = uint64(ii.highlowcontainer.getKeyAtIndex(ii.pos)) << 32
		ii.bitmapIter.Initialize(ii.highlowcontainer.getContainerAtIndex(ii.pos))
		ii.iter = &ii.bitmapIter
	} else {
		ii.iter = nil
	}
}

func (ii *manyIntIterator) NextMany(buf []uint64) int { _ = "STUB: not implemented"; return 0 }

// ManyIntIterator64 is meant to allow you to iterate through the values of a bitmap, see Initialize(a *Bitmap)
type ManyIntIterator64 = manyIntIterator

// Initialize configures the existing iterator so that it can iterate through the values of
// the provided bitmap.
// The iteration results are undefined if the bitmap is modified (e.g., with Add or Remove).
func (ii *manyIntIterator) Initialize(a *Bitmap) { _ = "STUB: not implemented"; return }

func newManyIntIterator(a *Bitmap) *manyIntIterator { _ = "STUB: not implemented"; return nil }
