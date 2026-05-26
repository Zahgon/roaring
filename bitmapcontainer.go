package roaring

import (
	"unsafe"
)

type bitmapContainer struct {
	cardinality int
	bitmap      []uint64
}

func (bc bitmapContainer) String() string { _ = "STUB: not implemented"; return "" }

func newBitmapContainer() *bitmapContainer { _ = "STUB: not implemented"; return nil }

func newBitmapContainerwithRange(firstOfRun, lastOfRun int) *bitmapContainer {
	_ = "STUB: not implemented"
	return nil
}

func (bc *bitmapContainer) minimum() uint16 { _ = "STUB: not implemented"; return 0 }

func (bc *bitmapContainer) safeMinimum() (uint16, error) { _ = "STUB: not implemented"; return 0, nil }

// i should be non-zero
func clz(i uint64) int { _ = "STUB: not implemented"; return 0 }

func (bc *bitmapContainer) maximum() uint16 { _ = "STUB: not implemented"; return 0 }

func (bc *bitmapContainer) safeMaximum() (uint16, error) { _ = "STUB: not implemented"; return 0, nil }

func (bc *bitmapContainer) iterate(cb func(x uint16) bool) bool {
	_ = "STUB: not implemented"
	return false
}

type bitmapContainerShortIterator struct {
	ptr *bitmapContainer
	i   int
}

func (bcsi *bitmapContainerShortIterator) next() uint16 { _ = "STUB: not implemented"; return 0 }

func (bcsi *bitmapContainerShortIterator) hasNext() bool { _ = "STUB: not implemented"; return false }

func (bcsi *bitmapContainerShortIterator) peekNext() uint16 { _ = "STUB: not implemented"; return 0 }

func (bcsi *bitmapContainerShortIterator) advanceIfNeeded(minval uint16) {
	_ = "STUB: not implemented"
	return
}

func newBitmapContainerShortIterator(a *bitmapContainer) *bitmapContainerShortIterator {
	_ = "STUB: not implemented"
	return nil
}

func (bc *bitmapContainer) getShortIterator() shortPeekable {
	_ = "STUB: not implemented"
	return *new(shortPeekable)
}

type reverseBitmapContainerShortIterator struct {
	ptr *bitmapContainer
	i   int
}

func (bcsi *reverseBitmapContainerShortIterator) next() uint16 { _ = "STUB: not implemented"; return 0 }

func (bcsi *reverseBitmapContainerShortIterator) hasNext() bool {
	_ = "STUB: not implemented"
	return false
}

func newReverseBitmapContainerShortIterator(a *bitmapContainer) *reverseBitmapContainerShortIterator {
	_ = "STUB: not implemented"
	return nil
}

func (bc *bitmapContainer) getReverseIterator() shortIterable {
	_ = "STUB: not implemented"
	return *new(shortIterable)
}

type bitmapContainerManyIterator struct {
	ptr    *bitmapContainer
	base   int
	bitset uint64
}

func (bcmi *bitmapContainerManyIterator) nextMany(hs uint32, buf []uint32) int {
	_ = "STUB: not implemented"
	return 0
}

// nextMany64 returns the number of values added to the buffer
func (bcmi *bitmapContainerManyIterator) nextMany64(hs uint64, buf []uint64) int {
	_ = "STUB: not implemented"
	return 0
}

func newBitmapContainerManyIterator(a *bitmapContainer) *bitmapContainerManyIterator {
	_ = "STUB: not implemented"
	return nil
}

func (bc *bitmapContainer) getManyIterator() manyIterable {
	_ = "STUB: not implemented"
	return *new(manyIterable)
}

type bitmapContainerUnsetIterator struct {
	ptr *bitmapContainer
	i   int
}

func (bcui *bitmapContainerUnsetIterator) next() uint16 { _ = "STUB: not implemented"; return 0 }

func (bcui *bitmapContainerUnsetIterator) hasNext() bool { _ = "STUB: not implemented"; return false }

func (bcui *bitmapContainerUnsetIterator) peekNext() uint16 { _ = "STUB: not implemented"; return 0 }

func (bcui *bitmapContainerUnsetIterator) advanceIfNeeded(minval uint16) {
	_ = "STUB: not implemented"
	return
}

func newBitmapContainerUnsetIterator(a *bitmapContainer) *bitmapContainerUnsetIterator {
	_ = "STUB: not implemented"
	return nil
}

func (bc *bitmapContainer) getUnsetIterator() shortPeekable {
	_ = "STUB: not implemented"
	return *new(shortPeekable)
}

func (bc *bitmapContainer) getSizeInBytes() int { _ = "STUB: not implemented"; return 0 }

func (bc *bitmapContainer) serializedSizeInBytes() int { _ = "STUB: not implemented"; return 0 }

const bcBaseBytes = int(unsafe.Sizeof(bitmapContainer{}))

// bitmapContainer doesn't depend on card, always fully allocated
func bitmapContainerSizeInBytes() int { _ = "STUB: not implemented"; return 0 }

func bitmapEquals(a, b []uint64) bool { _ = "STUB: not implemented"; return false }

func (bc *bitmapContainer) fillLeastSignificant16bits(x []uint32, i int, mask uint32) int {
	_ = "STUB: not implemented"
	// TODO: should be written as optimized assembly
	return 0
}

func (bc *bitmapContainer) equals(o container) bool { _ = "STUB: not implemented"; return false }

// use generic comparison

func (bc *bitmapContainer) iaddReturnMinimized(i uint16) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// iadd adds the arg i, returning true if not already present
func (bc *bitmapContainer) iadd(i uint16) bool { _ = "STUB: not implemented"; return false }

func (bc *bitmapContainer) iremoveReturnMinimized(i uint16) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// iremove returns true if i was found.
func (bc *bitmapContainer) iremove(i uint16) bool { _ = "STUB: not implemented"; return false }

func (bc *bitmapContainer) isFull() bool { _ = "STUB: not implemented"; return false }

func (bc *bitmapContainer) getCardinality() int { _ = "STUB: not implemented"; return 0 }

func (bc *bitmapContainer) isEmpty() bool { _ = "STUB: not implemented"; return false }

func (bc *bitmapContainer) clone() container { _ = "STUB: not implemented"; return *new(container) }

// add all values in range [firstOfRange,lastOfRange)
func (bc *bitmapContainer) iaddRange(firstOfRange, lastOfRange int) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// remove all values in range [firstOfRange,lastOfRange)
func (bc *bitmapContainer) iremoveRange(firstOfRange, lastOfRange int) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// flip all values in range [firstOfRange,endx)
func (bc *bitmapContainer) inot(firstOfRange, endx int) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// flip all values in range [firstOfRange,endx)
func (bc *bitmapContainer) not(firstOfRange, endx int) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) or(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) orCardinality(a container) int { _ = "STUB: not implemented"; return 0 }

func (bc *bitmapContainer) ior(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// bc.computeCardinality()

func (bc *bitmapContainer) lazyIOR(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// Manually inlined setBitmapRange function

func (bc *bitmapContainer) lazyOR(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// TODO: implement lazy OR

func (bc *bitmapContainer) orArray(value2 *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) orArrayCardinality(value2 *arrayContainer) int {
	_ = "STUB: not implemented"
	return 0
}

// branchless:

func (bc *bitmapContainer) orBitmap(value2 *bitmapContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) orBitmapCardinality(value2 *bitmapContainer) int {
	_ = "STUB: not implemented"
	return 0
}

func (bc *bitmapContainer) andBitmapCardinality(value2 *bitmapContainer) int {
	_ = "STUB: not implemented"
	return 0
}

func (bc *bitmapContainer) computeCardinality() { _ = "STUB: not implemented"; return }

func (bc *bitmapContainer) iorArray(ac *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) iorBitmap(value2 *bitmapContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) lazyIORArray(value2 *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) lazyORArray(value2 *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) lazyIORBitmap(value2 *bitmapContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) lazyORBitmap(value2 *bitmapContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) xor(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) xorArray(value2 *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) rank(x uint16) int {
	_ = "STUB: not implemented"
	// TODO: rewrite in assembly
	return 0
}

func (bc *bitmapContainer) selectInt(x uint16) int { _ = "STUB: not implemented"; return 0 }

func (bc *bitmapContainer) xorBitmap(value2 *bitmapContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) and(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) andCardinality(a container) int { _ = "STUB: not implemented"; return 0 }

func (bc *bitmapContainer) intersects(a container) bool { _ = "STUB: not implemented"; return false }

func (bc *bitmapContainer) iand(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) iandRun16(rc *runContainer16) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) iandArray(ac *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) andArray(value2 *arrayContainer) *arrayContainer {
	_ = "STUB: not implemented"
	return nil
}

func (bc *bitmapContainer) andArrayCardinality(value2 *arrayContainer) int {
	_ = "STUB: not implemented"
	return 0
}

func (bc *bitmapContainer) getCardinalityInRange(start, end uint) int {
	_ = "STUB: not implemented"
	return 0
}

func (bc *bitmapContainer) andBitmap(value2 *bitmapContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) intersectsArray(value2 *arrayContainer) bool {
	_ = "STUB: not implemented"
	return false
}

func (bc *bitmapContainer) intersectsBitmap(value2 *bitmapContainer) bool {
	_ = "STUB: not implemented"
	return false
}

func (bc *bitmapContainer) iandBitmap(value2 *bitmapContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) ixor(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) ixorArray(value2 *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) ixorRun16(value2 *runContainer16) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) ixorBitmap(value2 *bitmapContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) andNot(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) andNotRun16(rc *runContainer16) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) iandNot(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) iandNotArray(ac *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// Nothing to do.

// Word by word, we remove the elements in ac from bc. The approach is to build
// a mask of the elements to remove, and then apply it to the bitmap.

// Flush the current word.

// We're removing bits that are set in the mask and in the current word.
// To figure out the cardinality change, we count the number of bits that
// are set in the mask and in the current word.

// Flush the last word.

func (bc *bitmapContainer) iandNotRun16(rc *runContainer16) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// Nothing to do.

// inclusive

// before cardinality - after cardinality (for word range)

func (bc *bitmapContainer) andNotArray(value2 *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) andNotBitmap(value2 *bitmapContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) iandNotBitmapSurely(value2 *bitmapContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (bc *bitmapContainer) contains(i uint16) bool {
	_ = "STUB: not implemented" // testbit
	return false
}

func (bc *bitmapContainer) bitValue(i uint16) uint64 { _ = "STUB: not implemented"; return 0 }

func (bc *bitmapContainer) loadData(arrayContainer *arrayContainer) {
	_ = "STUB: not implemented"
	return
}

func (bc *bitmapContainer) resetTo(a container) { _ = "STUB: not implemented"; return }

func (bc *bitmapContainer) toArrayContainer() *arrayContainer {
	_ = "STUB: not implemented"
	return nil
}

func (bc *bitmapContainer) fillArray(container []uint16) {
	_ = "STUB: not implemented"
	// TODO: rewrite in assembly
	return
}

// NextSetBit returns the next set bit e.g the next int packed into the bitmaparray
func (bc *bitmapContainer) NextSetBit(i uint) int { _ = "STUB: not implemented"; return 0 }

func (bc *bitmapContainer) NextUnsetBit(i uint) int { _ = "STUB: not implemented"; return 0 }

// PrevSetBit returns the previous set bit e.g the previous int packed into the bitmaparray
func (bc *bitmapContainer) PrevSetBit(i int) int { _ = "STUB: not implemented"; return 0 }

func (bc *bitmapContainer) uPrevSetBit(i uint) int { _ = "STUB: not implemented"; return 0 }

// reference the java implementation
// https://github.com/RoaringBitmap/RoaringBitmap/blob/master/src/main/java/org/roaringbitmap/BitmapContainer.java#L875-L892
func (bc *bitmapContainer) numberOfRuns() int { _ = "STUB: not implemented"; return 0 }

// convert to run or array *if needed*
func (bc *bitmapContainer) toEfficientContainer() container {
	_ = "STUB: not implemented"
	return *new(container)
}

func newBitmapContainerFromRun(rc *runContainer16) *bitmapContainer {
	_ = "STUB: not implemented"
	return nil
}

// bc.computeCardinality()

func (bc *bitmapContainer) containerType() contype { _ = "STUB: not implemented"; return *new(contype) }

func (bc *bitmapContainer) addOffset(x uint16) (container, container) {
	_ = "STUB: not implemented"
	return *new(container), *new(container)
}

// All elements from bc ended up in low, meaning high will be empty.

// low is empty, let's reuse the container for high.

// None of the containers will be empty, so allocate both.

// Ensure proper nil interface.

// nextValue returns either the `target` if found or the next largest value.
// if the target is out of bounds a -1 is returned
//
// Example :
// Suppose the bitmap container represents the following slice
// [1,2,10,11,100]
// target=0 returns 1
// target=1 returns 1
// target=10 returns 10
// target=90 returns 100
func (bc *bitmapContainer) nextValue(target uint16) int { _ = "STUB: not implemented"; return 0 }

// nextAbsentValue returns the next absent value.
// if the target is out of bounds a -1 is returned
func (bc *bitmapContainer) nextAbsentValue(target uint16) int { _ = "STUB: not implemented"; return 0 }

// Check if all 1's
// if statement - we skip the if we have all ones [1,1,1,1...1]

// we have something like [X,Y,Z, 0,0,0]. This means the target bit is zero

// other wise something like [X,Y,0,1,1,1..1], where x and y can be either 1 or 0.

// previousValue returns either the `target` if found or the previous largest value.
// if the target is out of bounds a -1 is returned

// Example :
// Suppose the bitmap container represents the following slice
// [1,2,10,11,100]
// target=0 returns -1
// target=1 returns -1
// target=2 returns -1
// target=10 returns 9
// target=50 returns 10
// target=100 returns 99
func (bc *bitmapContainer) previousValue(target uint16) int { _ = "STUB: not implemented"; return 0 }

// previousAbsentValue returns the next absent value.
func (bc *bitmapContainer) previousAbsentValue(target uint16) int {
	_ = "STUB: not implemented"
	return 0
}

// Check if all 1's
// if statement - we skip if we have all ones [1,1,1,1...1] as no value is absent

// we have something like shifted=[X,Y,Z,..., 0,0,0]. This means the target bit is zero

// The rotate will rotate the target bit into the leading position.
// We know the target bit is not zero because of the countTrailingZero check above
// We then shift the target bit out of the way.
// Assume a structure like an original structure like [X,Y,Z,..., Target, A, B,C...]
// shifted will be [X,Y,Z...Target]
// shiftedRotated will be [A,B,C....]
// If countLeadingZeros > 0 then A is zero, if not at least A is 1 return
// Else count the number of ones's until a 0

// validate checks that the container size is non-negative
func (bc *bitmapContainer) validate() error { _ = "STUB: not implemented"; return nil }
