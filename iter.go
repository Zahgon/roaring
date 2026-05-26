package roaring

import "iter"

// Values returns an iterator that yields the elements of the bitmap in
// increasing order. Starting with Go 1.23, users can use a for loop to iterate
// over it.
func Values(b *Bitmap) iter.Seq[uint32] { _ = "STUB: not implemented"; return nil }

// Backward returns an iterator that yields the elements of the bitmap in
// decreasing order. Starting with Go 1.23, users can use a for loop to iterate
// over it.
func Backward(b *Bitmap) iter.Seq[uint32] { _ = "STUB: not implemented"; return nil }

// Unset creates an iterator that yields values in the range [min, max] that are NOT contained in the bitmap.
// The iterator becomes invalid if the bitmap is modified (e.g., with Add or Remove).
func Unset(b *Bitmap, min, max uint32) iter.Seq[uint32] { _ = "STUB: not implemented"; return nil }

// Ranges iterates contiguous ranges of values present in the bitmap as
// half-open [start, endExclusive) pairs. endExclusive is uint64 to represent
// ranges that include MaxUint32. Ranges spanning container boundaries are merged.
func (b *Bitmap) Ranges() iter.Seq2[uint32, uint64] { _ = "STUB: not implemented"; return nil }
