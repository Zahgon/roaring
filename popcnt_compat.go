//go:build !go1.9
// +build !go1.9

package roaring

// bit population count, take from
// https://code.google.com/p/go/issues/detail?id=4988#c11
// credit: https://code.google.com/u/arnehormann/
// credit: https://play.golang.org/p/U7SogJ7psJ
// credit: http://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetParallel
func popcount(x uint64) uint64 { _ = "STUB: not implemented"; return 0 }
