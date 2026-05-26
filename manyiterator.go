package roaring

type manyIterable interface {
	nextMany(hs uint32, buf []uint32) int
	nextMany64(hs uint64, buf []uint64) int
}

func (si *shortIterator) nextMany(hs uint32, buf []uint32) int { _ = "STUB: not implemented"; return 0 }

func (si *shortIterator) nextMany64(hs uint64, buf []uint64) int {
	_ = "STUB: not implemented"
	return 0
}
