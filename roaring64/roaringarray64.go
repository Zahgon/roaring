package roaring64

import (
	"errors"

	"github.com/RoaringBitmap/roaring/v2"
)

type roaringArray64 struct {
	keys            []uint32
	containers      []*roaring.Bitmap
	needCopyOnWrite []bool
	copyOnWrite     bool
}

var (
	ErrKeySortOrder          = errors.New("keys were out of order")
	ErrCardinalityConstraint = errors.New("size of arrays was not coherent")
)

// runOptimize compresses the element containers to minimize space consumed.
// Q: how does this interact with copyOnWrite and needCopyOnWrite?
// A: since we aren't changing the logical content, just the representation,
//
//	we don't bother to check the needCopyOnWrite bits. We replace
//	(possibly all) elements of ra.containers in-place with space
//	optimized versions.
func (ra *roaringArray64) runOptimize() { _ = "STUB: not implemented"; return }

func (ra *roaringArray64) appendContainer(key uint32, value *roaring.Bitmap, mustCopyOnWrite bool) {
	_ = "STUB: not implemented"
	return
}

func (ra *roaringArray64) appendWithoutCopy(sa roaringArray64, startingindex int) {
	_ = "STUB: not implemented"
	return
}

func (ra *roaringArray64) appendCopy(sa roaringArray64, startingindex int) {
	_ = "STUB: not implemented"
	// cow only if the two request it, or if we already have a lightweight copy
	return
}

// since there is no copy-on-write, we need to clone the container (this is important)

func (ra *roaringArray64) appendWithoutCopyMany(sa roaringArray64, startingindex, end int) {
	_ = "STUB: not implemented"
	return
}

func (ra *roaringArray64) appendCopyMany(sa roaringArray64, startingindex, end int) {
	_ = "STUB: not implemented"
	return
}

func (ra *roaringArray64) appendCopiesUntil(sa roaringArray64, stoppingKey uint32) {
	_ = "STUB: not implemented"
	// cow only if the two request it, or if we already have a lightweight copy
	return
}

// since there is no copy-on-write, we need to clone the container (this is important)

func (ra *roaringArray64) appendCopiesAfter(sa roaringArray64, beforeStart uint32) {
	_ = "STUB: not implemented"
	// cow only if the two request it, or if we already have a lightweight copy
	return
}

// since there is no copy-on-write, we need to clone the container (this is important)

func (ra *roaringArray64) removeIndexRange(begin, end int) { _ = "STUB: not implemented"; return }

func (ra *roaringArray64) resize(newsize int) { _ = "STUB: not implemented"; return }

func (ra *roaringArray64) clear() { _ = "STUB: not implemented"; return }

func (ra *roaringArray64) clone() *roaringArray64 { _ = "STUB: not implemented"; return nil }

// this is where copyOnWrite is used.

// sa.needCopyOnWrite is shared

// make a full copy

// clone all containers which have needCopyOnWrite set to true
// This can be used to make sure it is safe to munmap a []byte
// that the roaring array may still have a reference to.
func (ra *roaringArray64) cloneCopyOnWriteContainers() { _ = "STUB: not implemented"; return }

// unused function:
// func (ra *roaringArray64) containsKey(x uint32) bool {
//	return (ra.binarySearch(0, int64(len(ra.keys)), x) >= 0)
// }

func (ra *roaringArray64) getContainer(x uint32) *roaring.Bitmap {
	_ = "STUB: not implemented"
	return nil
}

func (ra *roaringArray64) getContainerAtIndex(i int) *roaring.Bitmap {
	_ = "STUB: not implemented"
	return nil
}

func (ra *roaringArray64) getWritableContainerAtIndex(i int) *roaring.Bitmap {
	_ = "STUB: not implemented"
	return nil
}

func (ra *roaringArray64) getIndex(x uint32) int {
	_ = "STUB: not implemented"
	// before the binary search, we optimize for frequent cases
	return 0
}

func (ra *roaringArray64) getKeyAtIndex(i int) uint32 { _ = "STUB: not implemented"; return 0 }

func (ra *roaringArray64) insertNewKeyValueAt(i int, key uint32, value *roaring.Bitmap) {
	_ = "STUB: not implemented"
	return
}

func (ra *roaringArray64) remove(key uint32) bool { _ = "STUB: not implemented"; return false }

// if a new key

func (ra *roaringArray64) removeAtIndex(i int) { _ = "STUB: not implemented"; return }

func (ra *roaringArray64) setContainerAtIndex(i int, c *roaring.Bitmap) {
	_ = "STUB: not implemented"
	return
}

func (ra *roaringArray64) replaceKeyAndContainerAtIndex(i int, key uint32, c *roaring.Bitmap, mustCopyOnWrite bool) {
	_ = "STUB: not implemented"
	return
}

func (ra *roaringArray64) size() int { _ = "STUB: not implemented"; return 0 }

func (ra *roaringArray64) binarySearch(begin, end int64, ikey uint32) int {
	_ = "STUB: not implemented"
	return 0
}

// avoid overflow

func (ra *roaringArray64) equals(o interface{}) bool { _ = "STUB: not implemented"; return false }

func (ra *roaringArray64) hasRunCompression() bool { _ = "STUB: not implemented"; return false }

/**
 * Find the smallest integer index strictly larger than pos such that array[index].key&gt;=min. If none can
 * be found, return size. Based on code by O. Kaser.
 *
 * @param min minimal value
 * @param pos index to exceed
 * @return the smallest index greater than pos such that array[index].key is at least as large as
 *         min, or size if it is not possible.
 */
func (ra *roaringArray64) advanceUntil(min uint32, pos int) int {
	_ = "STUB: not implemented"
	return 0
}

// means
// array
// has no
// item
// >= min
// pos = array.length;

// we know that the next-smallest span was too small

func (ra *roaringArray64) markAllAsNeedingCopyOnWrite() { _ = "STUB: not implemented"; return }

func (ra *roaringArray64) needsCopyOnWrite(i int) bool { _ = "STUB: not implemented"; return false }

func (ra *roaringArray64) setNeedsCopyOnWrite(i int) { _ = "STUB: not implemented"; return }

// should be dirt cheap
func (ra *roaringArray64) serializedSizeInBytes() uint64 { _ = "STUB: not implemented"; return 0 }

func (ra *roaringArray64) checkKeysSorted() bool { _ = "STUB: not implemented"; return false }

// validate checks the referential integrity
// ensures len(keys) == len(containers), recurses and checks each container type
func (ra *roaringArray64) validate() error { _ = "STUB: not implemented"; return nil }
