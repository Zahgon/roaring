//go:build !go1.9
// +build !go1.9

package roaring

// LeadingZeroBits returns the number of consecutive most significant zero
// bits of x.
func countLeadingZeros(i uint64) int { _ = "STUB: not implemented"; return 0 }
