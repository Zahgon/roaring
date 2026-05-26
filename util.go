package roaring

import (
	"math"
)

const (
	arrayDefaultMaxSize        = 4096 // containers with 4096 or fewer integers should be array containers.
	arrayLazyLowerBound        = 1024
	maxCapacity                = 1 << 16
	serialCookieNoRunContainer = 12346 // only arrays and bitmaps
	invalidCardinality         = -1
	serialCookie               = 12347 // runs, arrays, and bitmaps
	noOffsetThreshold          = 4

	// MaxUint32 is the largest uint32 value.
	MaxUint32 = math.MaxUint32

	// MaxRange is One more than the maximum allowed bitmap bit index. For use as an upper
	// bound for ranges.
	MaxRange uint64 = MaxUint32 + 1

	// MaxUint16 is the largest 16 bit unsigned int.
	// This is the largest value an interval16 can store.
	MaxUint16 = math.MaxUint16

	// Compute wordSizeInBytes, the size of a word in bytes.
	_m              = ^uint64(0)
	_logS           = _m>>8&1 + _m>>16&1 + _m>>32&1
	wordSizeInBytes = 1 << _logS

	// other constants used in ctz_generic.go
	wordSizeInBits = wordSizeInBytes << 3 // word size in bits
)

const maxWord = 1<<wordSizeInBits - 1

// doesn't apply to runContainers
func getSizeInBytesFromCardinality(card int) int { _ = "STUB: not implemented"; return 0 }

// bitmapContainer

// arrayContainer

func fill(arr []uint64, val uint64) { _ = "STUB: not implemented"; return }

func fillRange(arr []uint64, start, end int, val uint64) { _ = "STUB: not implemented"; return }

func fillArrayAND(container []uint16, bitmap1, bitmap2 []uint64) { _ = "STUB: not implemented"; return }

// TODO: rewrite in assembly

func fillArrayANDNOT(container []uint16, bitmap1, bitmap2 []uint64) {
	_ = "STUB: not implemented"
	return
}

// TODO: rewrite in assembly

func fillArrayXOR(container []uint16, bitmap1, bitmap2 []uint64) { _ = "STUB: not implemented"; return }

// TODO: rewrite in assembly

func highbits(x uint32) uint16 { _ = "STUB: not implemented"; return 0 }

func lowbits(x uint32) uint16 { _ = "STUB: not implemented"; return 0 }

func combineLoHi16(lob uint16, hob uint16) uint32 { _ = "STUB: not implemented"; return 0 }

func combineLoHi32(lob uint32, hob uint32) uint32 { _ = "STUB: not implemented"; return 0 }

const maxLowBit = 0xFFFF

func flipBitmapRange(bitmap []uint64, start int, end int) { _ = "STUB: not implemented"; return }

func resetBitmapRange(bitmap []uint64, start int, end int) { _ = "STUB: not implemented"; return }

func setBitmapRange(bitmap []uint64, start int, end int) { _ = "STUB: not implemented"; return }

func flipBitmapRangeAndCardinalityChange(bitmap []uint64, start int, end int) int {
	_ = "STUB: not implemented"
	return 0
}

func resetBitmapRangeAndCardinalityChange(bitmap []uint64, start int, end int) int {
	_ = "STUB: not implemented"
	return 0
}

func setBitmapRangeAndCardinalityChange(bitmap []uint64, start int, end int) int {
	_ = "STUB: not implemented"
	return 0
}

func wordCardinalityForBitmapRange(bitmap []uint64, start int, end int) uint64 {
	_ = "STUB: not implemented"
	return 0
}

func selectBitPosition(w uint64, j int) int {
	_ = "STUB: not implemented"

	// Divide 64bit
	return 0
}

// Divide 32bit

// Divide 16bit

// Lookup in final byte

func panicOn(err error) { _ = "STUB: not implemented"; return }

type ph struct {
	orig int
	rand int
}

func getRandomPermutation(n int) []int { _ = "STUB: not implemented"; return nil }

func minOfInt(a, b int) int { _ = "STUB: not implemented"; return 0 }

func maxOfInt(a, b int) int { _ = "STUB: not implemented"; return 0 }

func maxOfUint16(a, b uint16) uint16 { _ = "STUB: not implemented"; return 0 }

func minOfUint16(a, b uint16) uint16 { _ = "STUB: not implemented"; return 0 }
