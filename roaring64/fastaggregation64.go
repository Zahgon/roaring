package roaring64

// FastAnd computes the intersection between many bitmaps quickly
// Compared to the And function, it can take many bitmaps as input, thus saving the trouble
// of manually calling "And" many times.
func FastAnd(bitmaps ...*Bitmap) *Bitmap { _ = "STUB: not implemented"; return nil }

// FastOr computes the union between many bitmaps quickly, as opposed to having to call Or repeatedly.
func FastOr(bitmaps ...*Bitmap) *Bitmap { _ = "STUB: not implemented"; return nil }
