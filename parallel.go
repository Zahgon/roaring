package roaring

import (
	"runtime"
)

var defaultWorkerCount = runtime.NumCPU()

type bitmapContainerKey struct {
	key    uint16
	idx    int
	bitmap *Bitmap
}

type multipleContainers struct {
	key        uint16
	containers []container
	idx        int
}

type keyedContainer struct {
	key       uint16
	container container
	idx       int
}

type bitmapContainerHeap []bitmapContainerKey

func (h bitmapContainerHeap) Len() int           { _ = "STUB: not implemented"; return 0 }
func (h bitmapContainerHeap) Less(i, j int) bool { _ = "STUB: not implemented"; return false }
func (h bitmapContainerHeap) Swap(i, j int)      { _ = "STUB: not implemented"; return }

func (h *bitmapContainerHeap) Push(x interface{}) {
	_ = "STUB: not implemented"
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	return
}

func (h *bitmapContainerHeap) Pop() interface{} { _ = "STUB: not implemented"; return nil }

func (h bitmapContainerHeap) Peek() bitmapContainerKey {
	_ = "STUB: not implemented"
	return *new(bitmapContainerKey)
}

func (h *bitmapContainerHeap) popIncrementing() (key uint16, container container) {
	_ = "STUB: not implemented"
	return 0, *new(container)
}

func (h *bitmapContainerHeap) Next(containers []container) multipleContainers {
	_ = "STUB: not implemented"
	return *new(multipleContainers)
}

func newBitmapContainerHeap(bitmaps ...*Bitmap) bitmapContainerHeap {
	_ = "STUB: not implemented"
	// Initialize heap
	return *new(bitmapContainerHeap)
}

func repairAfterLazy(c container) container { _ = "STUB: not implemented"; return *new(container) }

func toBitmapContainer(c container) container { _ = "STUB: not implemented"; return *new(container) }

func appenderRoutine(bitmapChan chan<- *Bitmap, resultChan <-chan keyedContainer, expectedKeysChan <-chan int) {
	_ = "STUB: not implemented"
	return
}

// in case a resulting container was empty, see ParAnd function

// ParHeapOr computes the union (OR) of all provided bitmaps in parallel,
// where the parameter "parallelism" determines how many workers are to be used
// (if it is set to 0, a default number of workers is chosen)
// ParHeapOr uses a heap to compute the union. For rare cases it might be faster than ParOr
func ParHeapOr(parallelism int, bitmaps ...*Bitmap) *Bitmap { _ = "STUB: not implemented"; return nil }

// Assumes only structs with >=2 containers are passed

// ParAnd computes the intersection (AND) of all provided bitmaps in parallel,
// where the parameter "parallelism" determines how many workers are to be used
// (if it is set to 0, a default number of workers is chosen)
func ParAnd(parallelism int, bitmaps ...*Bitmap) *Bitmap { _ = "STUB: not implemented"; return nil }

// Assumes only structs with >=2 containers are passed

// Send a nil explicitly if the result of the intersection is an empty container

// ParOr computes the union (OR) of all provided bitmaps in parallel,
// where the parameter "parallelism" determines how many workers are to be used
// (if it is set to 0, a default number of workers is chosen)
func ParOr(parallelism int, bitmaps ...*Bitmap) *Bitmap { _ = "STUB: not implemented"; return nil }

// revert to FastOr. Since the key range is 0
// no container-level aggregation parallelism is achievable

// it's fine to panic to indicate an implementation error

type parChunkSpec struct {
	start uint16
	end   uint16
	idx   int
}

type parChunk struct {
	ra  *roaringArray
	idx int
}

func (c parChunk) size() int { _ = "STUB: not implemented"; return 0 }

func parNaiveStartAt(ra *roaringArray, start uint16, last uint16) int {
	_ = "STUB: not implemented"
	return 0
}

func lazyOrOnRange(ra1, ra2 *roaringArray, start, last uint16) *roaringArray {
	_ = "STUB: not implemented"
	return nil
}

func lazyIOrOnRange(ra1, ra2 *roaringArray, start, last uint16) *roaringArray {
	_ = "STUB: not implemented"
	return nil
}
