package roaring

// Or function that requires repairAfterLazy
func lazyOR(x1, x2 *Bitmap) *Bitmap { _ = "STUB: not implemented"; return nil }

// In-place Or function that requires repairAfterLazy
func (x1 *Bitmap) lazyOR(x2 *Bitmap) *Bitmap { _ = "STUB: not implemented"; return nil }

// runContainer16.lazyIOR falls back to a slow ior path
// (O(N log R) per merged element); promote to bitmapContainer
// first, whose lazy union is O(1024) regardless of cardinality.
// See https://github.com/RoaringBitmap/roaring/issues/81.

// to be called after lazy aggregates
func (x1 *Bitmap) repairAfterLazy() { _ = "STUB: not implemented"; return }

// FastAnd computes the intersection between many bitmaps quickly
// Compared to the And function, it can take many bitmaps as input, thus saving the trouble
// of manually calling "And" many times.
//
// Performance hints: if you have very large and tiny bitmaps,
// it may be beneficial performance-wise to put a tiny bitmap
// in first position.
func FastAnd(bitmaps ...*Bitmap) *Bitmap { _ = "STUB: not implemented"; return nil }

// FastOr computes the union between many bitmaps quickly, as opposed to having to call Or repeatedly.
// It might also be faster than calling Or repeatedly.
func FastOr(bitmaps ...*Bitmap) *Bitmap { _ = "STUB: not implemented"; return nil }

// here is where repairAfterLazy is called.

// HeapOr computes the union between many bitmaps quickly using a heap.
// It might be faster than calling Or repeatedly.
func HeapOr(bitmaps ...*Bitmap) *Bitmap { _ = "STUB: not implemented"; return nil }

// TODO:  for better speed, we could do the operation lazily, see Java implementation

// HeapXor computes the symmetric difference between many bitmaps quickly (as opposed to calling Xor repeated).
// Internally, this function uses a heap.
// It might be faster than calling Xor repeatedly.
func HeapXor(bitmaps ...*Bitmap) *Bitmap { _ = "STUB: not implemented"; return nil }

// AndAny provides a result equivalent to x1.And(FastOr(bitmaps)).
// It's optimized to minimize allocations. It also might be faster than separate calls.
func (x1 *Bitmap) AndAny(bitmaps ...*Bitmap) { _ = "STUB: not implemented"; return }

// accumulate containers for current key, find next minimal key in filters
// and exclude filters that do not have related values anymore

//TODO: special case for run containers?
