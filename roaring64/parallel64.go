package roaring64

import (
	"runtime"
)

var defaultWorkerCount = runtime.NumCPU()

// ParOr computes the union (OR) of all provided bitmaps in parallel,
// where the parameter "parallelism" determines how many workers are to be used
// (if it is set to 0, a default number of workers is chosen)
func ParOr(parallelism int, bitmaps ...*Bitmap) *Bitmap { _ = "STUB: not implemented"; return nil }

// The following might overflow and we do not want that!
// as it might lead to a channel of size 0 later which,
// on some systems, would block indefinitely.

// All bitmaps have the same key,
// we can merge the 32-bit roaring bitmaps in parallel

// We cannot use int since int is 32-bit on 32-bit systems.

// it's fine to panic to indicate an implementation error

type parChunkSpec struct {
	start uint32
	end   uint32
	idx   int
}

type parChunk struct {
	ra  *roaringArray64
	idx int
}

func (c parChunk) size() int {
	_ = "STUB: not implemented"

	// parNaiveStartAt returns the index of the first key that is inclusive between start and last
	// Returns the size if there is no such key
	return 0
}

func parNaiveStartAt(ra *roaringArray64, start uint32, last uint32) int {
	_ = "STUB: not implemented"
	return 0
}

func orOnRange(ra1, ra2 *roaringArray64, start, last uint32) *roaringArray64 {
	_ = "STUB: not implemented"
	return nil
}

// answer.appendContainer(key1, c1.lazyOR(ra2.getContainerAtIndex(idx2)), false)

func iorOnRange(ra1, ra2 *roaringArray64, start, last uint32) *roaringArray64 {
	_ = "STUB: not implemented"
	return nil
}

// ra1.containers[idx1] = c1.lazyIOR(ra2.getContainerAtIndex(idx2))
