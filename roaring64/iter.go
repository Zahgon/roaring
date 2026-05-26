package roaring64

import "iter"

// Values returns an iterator that yields the elements of the bitmap in
// increasing order. Starting with Go 1.23, users can use a for loop to iterate
// over it.
func Values(b *Bitmap) iter.Seq[uint64] { _ = "STUB: not implemented"; return nil }

// Backward returns an iterator that yields the elements of the bitmap in
// decreasing order. Starting with Go 1.23, users can use a for loop to iterate
// over it.
func Backward(b *Bitmap) iter.Seq[uint64] { _ = "STUB: not implemented"; return nil }
