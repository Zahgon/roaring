//go:build !arm64 || gccgo || appengine
// +build !arm64 gccgo appengine

package roaring

func union2by2(set1 []uint16, set2 []uint16, buffer []uint16) int {
	_ = "STUB: not implemented"
	return 0
}

// if (set1[k1]>set2[k2])
