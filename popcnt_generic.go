//go:build !amd64 || appengine || go1.9
// +build !amd64 appengine go1.9

package roaring

func popcntSlice(s []uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func popcntMaskSlice(s, m []uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func popcntAndSlice(s, m []uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func popcntOrSlice(s, m []uint64) uint64 { _ = "STUB: not implemented"; return 0 }

func popcntXorSlice(s, m []uint64) uint64 { _ = "STUB: not implemented"; return 0 }
