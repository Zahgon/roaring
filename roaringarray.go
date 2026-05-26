package roaring

import (
	"errors"
	"io"

	"github.com/RoaringBitmap/roaring/v2/internal"
)

type container interface {
	// addOffset returns the (low, high) parts of the shifted container.
	// Whenever one of them would be empty, nil will be returned instead to
	// avoid unnecessary allocations.
	addOffset(uint16) (container, container)

	clone() container
	and(container) container
	andCardinality(container) int
	iand(container) container // i stands for inplace
	andNot(container) container
	iandNot(container) container // i stands for inplace
	isEmpty() bool
	getCardinality() int
	// rank returns the number of integers that are
	// smaller or equal to x. rank(infinity) would be getCardinality().
	rank(uint16) int

	// getCardinalityInRange returns the number of integers that are
	// within the half-open range [start, end). It is equivalent to
	// rank(end-1) - rank(start-1) but may be faster.
	getCardinalityInRange(start, end uint) int

	iadd(x uint16) bool                   // inplace, returns true if x was new.
	iaddReturnMinimized(uint16) container // may change return type to minimize storage.

	iaddRange(start, endx int) container // i stands for inplace, range is [firstOfRange,endx)

	iremove(x uint16) bool                   // inplace, returns true if x was present.
	iremoveReturnMinimized(uint16) container // may change return type to minimize storage.

	not(start, final int) container        // range is [firstOfRange,lastOfRange)
	inot(firstOfRange, endx int) container // i stands for inplace, range is [firstOfRange,endx)
	xor(r container) container
	ixor(r container) container // i stands for inplace
	getShortIterator() shortPeekable
	getUnsetIterator() shortPeekable
	iterate(cb func(x uint16) bool) bool
	getReverseIterator() shortIterable
	getManyIterator() manyIterable
	contains(i uint16) bool
	maximum() uint16
	minimum() uint16

	// equals is now logical equals; it does not require the
	// same underlying container types, but compares across
	// any of the implementations.
	equals(r container) bool

	fillLeastSignificant16bits(array []uint32, i int, mask uint32) int
	or(r container) container
	orCardinality(r container) int
	isFull() bool
	ior(r container) container   // i stands for inplace
	intersects(r container) bool // whether the two containers intersect
	lazyOR(r container) container
	lazyIOR(r container) container
	getSizeInBytes() int
	iremoveRange(start, final int) container // i stands for inplace, range is [firstOfRange,lastOfRange)
	selectInt(x uint16) int                  // selectInt returns the xth integer in the container
	serializedSizeInBytes() int
	writeTo(io.Writer) (int, error)

	numberOfRuns() int
	toEfficientContainer() container
	String() string
	containerType() contype

	safeMinimum() (uint16, error)
	safeMaximum() (uint16, error)
	nextValue(x uint16) int
	previousValue(x uint16) int
	nextAbsentValue(x uint16) int
	previousAbsentValue(x uint16) int
	validate() error
}

type contype uint8

const (
	bitmapContype contype = iota
	arrayContype
	run16Contype
	run32Contype
)

var (
	ErrKeySortOrder          = errors.New("keys were out of order")
	ErrCardinalityConstraint = errors.New("size of arrays was not coherent")
)

// careful: range is [firstOfRange,lastOfRange]
func rangeOfOnes(start, last int) container { _ = "STUB: not implemented"; return *new(container) }

type roaringArray struct {
	keys            []uint16
	containers      []container `msg:"-"` // don't try to serialize directly.
	needCopyOnWrite []bool
	copyOnWrite     bool
}

func newRoaringArray() *roaringArray { _ = "STUB: not implemented"; return nil }

// runOptimize compresses the element containers to minimize space consumed.
// Q: how does this interact with copyOnWrite and needCopyOnWrite?
// A: since we aren't changing the logical content, just the representation,
//
//	we don't bother to check the needCopyOnWrite bits. We replace
//	(possibly all) elements of ra.containers in-place with space
//	optimized versions.
func (ra *roaringArray) runOptimize() { _ = "STUB: not implemented"; return }

func (ra *roaringArray) appendContainer(key uint16, value container, mustCopyOnWrite bool) {
	_ = "STUB: not implemented"
	return
}

func (ra *roaringArray) appendWithoutCopy(sa roaringArray, startingindex int) {
	_ = "STUB: not implemented"
	return
}

func (ra *roaringArray) appendCopy(sa roaringArray, startingindex int) {
	_ = "STUB: not implemented"
	// cow only if the two request it, or if we already have a lightweight copy
	return
}

// since there is no copy-on-write, we need to clone the container (this is important)

func (ra *roaringArray) appendWithoutCopyMany(sa roaringArray, startingindex, end int) {
	_ = "STUB: not implemented"
	return
}

func (ra *roaringArray) appendCopyMany(sa roaringArray, startingindex, end int) {
	_ = "STUB: not implemented"
	return
}

func (ra *roaringArray) appendCopiesUntil(sa roaringArray, stoppingKey uint16) {
	_ = "STUB: not implemented"
	// cow only if the two request it, or if we already have a lightweight copy
	return
}

// since there is no copy-on-write, we need to clone the container (this is important)

func (ra *roaringArray) appendCopiesAfter(sa roaringArray, beforeStart uint16) {
	_ = "STUB: not implemented"
	// cow only if the two request it, or if we already have a lightweight copy
	return
}

// since there is no copy-on-write, we need to clone the container (this is important)

func (ra *roaringArray) removeIndexRange(begin, end int) { _ = "STUB: not implemented"; return }

func (ra *roaringArray) resize(newsize int) { _ = "STUB: not implemented"; return }

func (ra *roaringArray) clear() { _ = "STUB: not implemented"; return }

func (ra *roaringArray) clone() *roaringArray { _ = "STUB: not implemented"; return nil }

// this is where copyOnWrite is used.

// sa.needCopyOnWrite is shared

// make a full copy

// clone all containers which have needCopyOnWrite set to true
// This can be used to make sure it is safe to munmap a []byte
// that the roaring array may still have a reference to.
func (ra *roaringArray) cloneCopyOnWriteContainers() { _ = "STUB: not implemented"; return }

// unused function:
//func (ra *roaringArray) containsKey(x uint16) bool {
//	return (ra.binarySearch(0, int64(len(ra.keys)), x) >= 0)
//}

// getContainer returns the container with key `x`
// if no such container exists `nil` is returned
func (ra *roaringArray) getContainer(x uint16) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ra *roaringArray) getContainerAtIndex(i int) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ra *roaringArray) getFastContainerAtIndex(i int, needsWriteable bool) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// getUnionedWritableContainer switches behavior for in-place Or
// depending on whether the container requires a copy on write.
// If it does using the non-inplace or() method leads to fewer allocations.
func (ra *roaringArray) getUnionedWritableContainer(pos int, other container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ra *roaringArray) getWritableContainerAtIndex(i int) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// getIndex returns the index of the container with key `x`
// if no such container exists a negative value is returned
func (ra *roaringArray) getIndex(x uint16) int {
	_ = "STUB: not implemented"
	// Todo : test
	// before the binary search, we optimize for frequent cases
	return 0
}

func (ra *roaringArray) getKeyAtIndex(i int) uint16 { _ = "STUB: not implemented"; return 0 }

func (ra *roaringArray) insertNewKeyValueAt(i int, key uint16, value container) {
	_ = "STUB: not implemented"
	return
}

func (ra *roaringArray) remove(key uint16) bool { _ = "STUB: not implemented"; return false }

// if a new key

func (ra *roaringArray) removeAtIndex(i int) { _ = "STUB: not implemented"; return }

func (ra *roaringArray) setContainerAtIndex(i int, c container) { _ = "STUB: not implemented"; return }

func (ra *roaringArray) replaceKeyAndContainerAtIndex(i int, key uint16, c container, mustCopyOnWrite bool) {
	_ = "STUB: not implemented"
	return
}

func (ra *roaringArray) size() int { _ = "STUB: not implemented"; return 0 }

// binarySearch returns the index of the key.
// negative value returned if not found
func (ra *roaringArray) binarySearch(begin, end int64, ikey uint16) int {
	_ = "STUB: not implemented"
	// TODO: add unit tests
	return 0
}

// avoid overflow

func (ra *roaringArray) equals(o interface{}) bool { _ = "STUB: not implemented"; return false }

func (ra *roaringArray) headerSize() uint64 { _ = "STUB: not implemented"; return 0 }

// for small bitmaps, we omit the offsets

// - 4 because we pack the size with the cookie

// should be dirt cheap
func (ra *roaringArray) serializedSizeInBytes() uint64 { _ = "STUB: not implemented"; return 0 }

// spec: https://github.com/RoaringBitmap/RoaringFormatSpec
func (ra *roaringArray) writeTo(w io.Writer) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// compute isRun bitmap without temporary allocation

// descriptive header

// offset header

// spec: https://github.com/RoaringBitmap/RoaringFormatSpec
func (ra *roaringArray) toBytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Reads a serialized roaringArray from a byte slice.
func (ra *roaringArray) readFrom(stream internal.ByteInput, cookieHeader ...byte) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// If NextReturnsSafeSlice is false, then willNeedCopyOnWrite should be true

// create is-run-container bitmap

// descriptive header

// Allocate slices upfront as number of containers is known

// run container

// bitmap container

// array container

func (ra *roaringArray) hasRunCompression() bool { _ = "STUB: not implemented"; return false }

/**
 * Find the smallest integer index larger than pos such that array[index].key&gt;=min. If none can
 * be found, return size. Based on code by O. Kaser.
 *
 * @param min minimal value
 * @param pos index to exceed
 * @return the smallest index greater than pos such that array[index].key is at least as large as
 *         min, or size if it is not possible.
 */
func (ra *roaringArray) advanceUntil(min uint16, pos int) int { _ = "STUB: not implemented"; return 0 }

// means
// array
// has no
// item
// >= min
// pos = array.length;

// we know that the next-smallest span was too small

func (ra *roaringArray) markAllAsNeedingCopyOnWrite() { _ = "STUB: not implemented"; return }

func (ra *roaringArray) needsCopyOnWrite(i int) bool { _ = "STUB: not implemented"; return false }

func (ra *roaringArray) setNeedsCopyOnWrite(i int) { _ = "STUB: not implemented"; return }

func (ra *roaringArray) checkKeysSorted() bool { _ = "STUB: not implemented"; return false }

// validate checks the referential integrity
// ensures len(keys) == len(containers), recurses and checks each container type
func (ra *roaringArray) validate() error { _ = "STUB: not implemented"; return nil }
