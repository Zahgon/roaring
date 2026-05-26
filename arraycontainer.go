package roaring

import (
	"errors"
)

type arrayContainer struct {
	content []uint16
}

var (
	ErrArrayIncorrectSort = errors.New("incorrectly sorted array")
	ErrEmptyArray         = errors.New("empty array")
	ErrArrayInvalidSize   = errors.New("invalid array size")
)

func (ac *arrayContainer) String() string { _ = "STUB: not implemented"; return "" }

func (ac *arrayContainer) fillLeastSignificant16bits(x []uint32, i int, mask uint32) int {
	_ = "STUB: not implemented"
	return 0
}

func (ac *arrayContainer) iterate(cb func(x uint16) bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (ac *arrayContainer) getShortIterator() shortPeekable {
	_ = "STUB: not implemented"
	return *new(shortPeekable)
}

func (ac *arrayContainer) getReverseIterator() shortIterable {
	_ = "STUB: not implemented"
	return *new(shortIterable)
}

func (ac *arrayContainer) getManyIterator() manyIterable {
	_ = "STUB: not implemented"
	return *new(manyIterable)
}

func (ac *arrayContainer) getUnsetIterator() shortPeekable {
	_ = "STUB: not implemented"
	return *new(shortPeekable)
}

func (ac *arrayContainer) minimum() uint16 {
	_ = "STUB: not implemented"
	// assume not empty
	return 0
}

func (ac *arrayContainer) safeMinimum() (uint16, error) { _ = "STUB: not implemented"; return 0, nil }

func (ac *arrayContainer) maximum() uint16 { _ = "STUB: not implemented"; return 0 }

// assume not empty

func (ac *arrayContainer) safeMaximum() (uint16, error) { _ = "STUB: not implemented"; return 0, nil }

func (ac *arrayContainer) getSizeInBytes() int { _ = "STUB: not implemented"; return 0 }

func (ac *arrayContainer) serializedSizeInBytes() int { _ = "STUB: not implemented"; return 0 }

func arrayContainerSizeInBytes(card int) int {
	_ = "STUB: not implemented"

	// add the values in the range [firstOfRange,endx)
	return 0
}

func (ac *arrayContainer) iaddRange(firstOfRange, endx int) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// remove the values in the range [firstOfRange,endx)
func (ac *arrayContainer) iremoveRange(firstOfRange, endx int) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// flip the values in the range [firstOfRange,endx)
func (ac *arrayContainer) not(firstOfRange, endx int) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// remove everything in [firstOfRange,endx-1]

// flip the values in the range [firstOfRange,lastOfRange]
func (ac *arrayContainer) notClose(firstOfRange, lastOfRange int) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// unlike add and remove, not uses an inclusive range [firstOfRange,lastOfRange]

// determine the span of array indices to be affected^M

// a hack for sure

func (ac *arrayContainer) equals(o container) bool { _ = "STUB: not implemented"; return false }

// Check if the containers are the same object.

// use generic comparison

func (ac *arrayContainer) toBitmapContainer() *bitmapContainer {
	_ = "STUB: not implemented"
	return nil
}

func (ac *arrayContainer) iadd(x uint16) (wasNew bool) {
	_ = "STUB: not implemented"
	// Special case adding to the end of the container.
	return false
}

func (ac *arrayContainer) iaddReturnMinimized(x uint16) container {
	_ = "STUB: not implemented"
	// Special case adding to the end of the container.
	return *new(container)
}

// iremoveReturnMinimized is allowed to change the return type to minimize storage.
func (ac *arrayContainer) iremoveReturnMinimized(x uint16) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ac *arrayContainer) iremove(x uint16) bool { _ = "STUB: not implemented"; return false }

func (ac *arrayContainer) remove(x uint16) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ac *arrayContainer) or(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ac *arrayContainer) orCardinality(a container) int { _ = "STUB: not implemented"; return 0 }

func (ac *arrayContainer) ior(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ac *arrayContainer) iorArray(value2 *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// doubling the capacity reduces new slice allocations in the case of
// repeated calls to iorArray().

// the second check is to handle overly large array containers
// and should not occur in normal usage,
// as all array containers should be at most arrayDefaultMaxSize

// reslice to match actual used capacity

// Only converting to a bitmap when arrayDefaultMaxSize
// is actually exceeded minimizes conversions in the case of repeated
// calls to iorArray().

// Note: such code does not make practical sense, except for lazy evaluations
func (ac *arrayContainer) iorBitmap(bc2 *bitmapContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// DO NOT DO THIS:
// *ac = *newArrayContainerFromBitmap(bc1)
// This will create gigantic array containers in the case of repeated calls to iorBitmap.

func (ac *arrayContainer) iorRun16(rc *runContainer16) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// heuristic for if the container should maybe be an
// array container.

func (ac *arrayContainer) lazyIOR(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ac *arrayContainer) lazyIorArray(ac2 *arrayContainer) container {
	_ = "STUB: not implemented"
	// TODO actually make this lazy
	return *new(container)
}

func (ac *arrayContainer) lazyIorBitmap(bc *bitmapContainer) container {
	_ = "STUB: not implemented"
	// TODO actually make this lazy
	return *new(container)
}

func (ac *arrayContainer) lazyIorRun16(rc *runContainer16) container {
	_ = "STUB: not implemented"
	// TODO actually make this lazy
	return *new(container)
}

func (ac *arrayContainer) lazyOR(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ac *arrayContainer) orArray(value2 *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// it could be a bitmap!

// reslice to match actual used capacity

func (ac *arrayContainer) orArrayCardinality(value2 *arrayContainer) int {
	_ = "STUB: not implemented"
	return 0
}

func (ac *arrayContainer) lazyorArray(value2 *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// it could be a bitmap!

// reslice to match actual used capacity

func (ac *arrayContainer) and(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ac *arrayContainer) andCardinality(a container) int { _ = "STUB: not implemented"; return 0 }

func (ac *arrayContainer) intersects(a container) bool { _ = "STUB: not implemented"; return false }

func (ac *arrayContainer) iand(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ac *arrayContainer) iandBitmap(bc *bitmapContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// branchless

func (ac *arrayContainer) xor(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ac *arrayContainer) ixor(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ac *arrayContainer) ixorArray(value2 *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ac *arrayContainer) ixorBitmap(value2 *bitmapContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ac *arrayContainer) ixorRun16(value2 *runContainer16) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ac *arrayContainer) xorArray(value2 *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// it could be a bitmap!

func (ac *arrayContainer) andNot(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ac *arrayContainer) andNotRun16(rc *runContainer16) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ac *arrayContainer) iandNot(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ac *arrayContainer) iandNotRun16(rc *runContainer16) container {
	_ = "STUB: not implemented"
	// Fast path: if either the array container or the run container is empty, the result is the array.
	return *new(container)
}

// Empty

// Fast path: if the run container is full, the result is empty.

// All values in [start_run, end_end] are part of the run

// We are going to read values in the array at index i, and we are
// going to write them at index pos. So we do in-place processing.
// We always have that pos <= i by construction. So we can either
// overwrite a value just read, or a value that was previous read.

// the value in the array appears before the run [start_run, end_end]

// nothing to do, the value is in the array but also in the run.

// We have the value in the array after the run. We cannot tell
// whether we need to keep it or not. So let us move to another run.

// retry with the same i

// We have exhausted the number of runs. We can keep the rest of the values
// from i to len(ac.content) - 1 inclusively.
// We are done, the rest of the array will be kept

// We 'shink' the slice.

func (ac *arrayContainer) andNotArray(value2 *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ac *arrayContainer) iandNotArray(value2 *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ac *arrayContainer) andNotBitmap(value2 *bitmapContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ac *arrayContainer) andBitmap(value2 *bitmapContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ac *arrayContainer) iandNotBitmap(value2 *bitmapContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func copyOf(array []uint16, size int) []uint16 { _ = "STUB: not implemented"; return nil }

// flip the values in the range [firstOfRange,endx)
func (ac *arrayContainer) inot(firstOfRange, endx int) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// remove everything in [firstOfRange,endx-1]

// flip the values in the range [firstOfRange,lastOfRange]
func (ac *arrayContainer) inotClose(firstOfRange, lastOfRange int) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// unlike add and remove, not uses an inclusive range [firstOfRange,lastOfRange]

// determine the span of array indices to be affected

// no expansion needed

func (ac *arrayContainer) negateRange(buffer []uint16, startIndex, lastIndex, startRange, lastRange int) {
	_ = "STUB: not implemented"
	// compute the negation into buffer
	return
}

// value here always >= valInRange,
// until it is exhausted
// n.b., we can start initially exhausted.

// if there are extra items (greater than the biggest
// pre-existing one in range), buffer them

func (ac *arrayContainer) isFull() bool { _ = "STUB: not implemented"; return false }

func (ac *arrayContainer) andArray(value2 *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ac *arrayContainer) andArrayCardinality(value2 *arrayContainer) int {
	_ = "STUB: not implemented"
	return 0
}

func (ac *arrayContainer) intersectsArray(value2 *arrayContainer) bool {
	_ = "STUB: not implemented"
	return false
}

func (ac *arrayContainer) iandArray(value2 *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ac *arrayContainer) getCardinality() int { _ = "STUB: not implemented"; return 0 }

func (ac *arrayContainer) isEmpty() bool { _ = "STUB: not implemented"; return false }

func (ac *arrayContainer) rank(x uint16) int { _ = "STUB: not implemented"; return 0 }

// getCardinalityInRange returns the number of values in the half-open range [start, end).
func (ac *arrayContainer) getCardinalityInRange(start, end uint) int {
	_ = "STUB: not implemented"
	return 0
}

// Find the first index >= start

// end can be up to 65536 (1<<16), which overflows uint16.
// In that case, all elements from loIdx onward are included.

// Find the first index >= end (i.e., past the last included value)

func (ac *arrayContainer) selectInt(x uint16) int { _ = "STUB: not implemented"; return 0 }

func (ac *arrayContainer) clone() container { _ = "STUB: not implemented"; return *new(container) }

func (ac *arrayContainer) contains(x uint16) bool { _ = "STUB: not implemented"; return false }

func (ac *arrayContainer) loadData(bitmapContainer *bitmapContainer) {
	_ = "STUB: not implemented"
	return
}

func (ac *arrayContainer) resetTo(a container) { _ = "STUB: not implemented"; return }

func (ac *arrayContainer) realloc(size int) { _ = "STUB: not implemented"; return }

// previousValue returns either the target if found or the previous smaller present value.
// If the target is out of bounds a -1 is returned.
// Ex: target=4 ac=[2,3,4,6,7] returns 4
// Ex: target=5 ac=[2,3,4,6,7] returns 4
// Ex: target=6 ac=[2,3,4,6,7] returns 6
// Ex: target=8 ac=[2,3,4,6,7] returns 7
// Ex: target=1 ac=[2,3,4,6,7] returns -1
// Ex: target=0 ac=[2,3,4,6,7] returns -1
func (ac *arrayContainer) previousValue(target uint16) int { _ = "STUB: not implemented"; return 0 }

// previousAbsentValue returns either the target if not found or the next larger missing value.
// If the target is out of bounds a -1 is returned
// Ex: target=4 ac=[1,2,3,4,6,7] returns 0
// Ex: target=5 ac=[1,2,3,4,6,7] returns 5
// Ex: target=6 ac=[1,2,3,4,6,7] returns 5
// Ex: target=8 ac=[1,2,3,4,6,7] returns 8
func (ac *arrayContainer) previousAbsentValue(target uint16) int {
	_ = "STUB: not implemented"
	return 0
}

// If the target was found at index 1, then the next value down must be result.value-1

// This uses the pigeon-hole principle.
// the if statement compares the difference in indices vs
// the difference in values. Suppose mid = 10 and result.index = 5
// with ac.content[mid] = 100 and target = 10
// then we have 5 slots for values but we need to fit in 90 values
// so some of the values must be missing

// nextAbsentValue returns either the target if not found or the next larger missing value.
// If the target is out of bounds a -1 is returned
// Ex: target=4 ac=[1,2,3,4,6,7] returns 5
// Ex: target=5 ac=[1,2,3,4,6,7] returns 5
// Ex: target=0 ac=[1,2,3,4,6,7] returns 0
// Ex: target=8 ac=[1,2,3,4,6,7] returns 8
func (ac *arrayContainer) nextAbsentValue(target uint16) int { _ = "STUB: not implemented"; return 0 }

// This uses the pigeon-hole principle.
// the if statement compares the difference in indices vs
// the difference in values. Suppose mid = 10 and result.index = 5
// with ac.content[mid] = 100 and target = 10
// then we have 5 slots for values but we need to fit in 90 values
// so some of the values must be missing

// nextValue returns either the target if found or the next larger value.
// if the target is out of bounds a -1 is returned
//
// Ex: target=4 ac=[1,2,3,4,6,7] returns 4
// Ex: target=5 ac=[1,2,3,4,6,7] returns 6
// Ex: target=6 ac=[1,2,3,4,6,7] returns 6
// Ex: target=0 ac=[1,2,3,4,6,7] returns 1
// Ex: target=100 ac=[1,2,3,4,6,7] returns -1
func (ac *arrayContainer) nextValue(target uint16) int { _ = "STUB: not implemented"; return 0 }

//if target < ac.minimum() {
//	return -1
//}
//if target > ac.maximum() {
//		return -1
//	}

func newArrayContainer() *arrayContainer { _ = "STUB: not implemented"; return nil }

func newArrayContainerFromBitmap(bc *bitmapContainer) *arrayContainer {
	_ = "STUB: not implemented"
	return nil
}

func newArrayContainerCapacity(size int) *arrayContainer { _ = "STUB: not implemented"; return nil }

func newArrayContainerSize(size int) *arrayContainer { _ = "STUB: not implemented"; return nil }

func newArrayContainerRange(firstOfRun, lastOfRun int) *arrayContainer {
	_ = "STUB: not implemented"
	return nil
}

func (ac *arrayContainer) numberOfRuns() (nr int) { _ = "STUB: not implemented"; return 0 }

// convert to run or array *if needed*
func (ac *arrayContainer) toEfficientContainer() container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (ac *arrayContainer) containerType() contype { _ = "STUB: not implemented"; return *new(contype) }

func (ac *arrayContainer) addOffset(x uint16) (container, container) {
	_ = "STUB: not implemented"
	return *new(container), *new(container)
}

// Some elements will fall into low part, allocate a container.
// Checking the first one is enough because they are ordered.

// Some elements will fall into high part, allocate a container.
// Checking the last one is enough because they are ordered.

// OK, if high == nil then highbits(y) == 0 for all y.

// OK, if low == nil then highbits(y) > 0 for all y.

// Ensure proper nil interface.

// validate checks cardinality and sort order of the array container
func (ac *arrayContainer) validate() error { _ = "STUB: not implemented"; return nil }
