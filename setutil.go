package roaring

func difference(set1 []uint16, set2 []uint16, buffer []uint16) int {
	_ = "STUB: not implemented"
	return 0
}

// if (val1>val2)

func exclusiveUnion2by2(set1 []uint16, set2 []uint16, buffer []uint16) int {
	_ = "STUB: not implemented"
	return 0
}

// if (val1>val2)

// union2by2Cardinality computes the cardinality of the union
func union2by2Cardinality(set1 []uint16, set2 []uint16) int { _ = "STUB: not implemented"; return 0 }

// if (set1[k1]>set2[k2])

func intersection2by2(
	set1 []uint16,
	set2 []uint16,
	buffer []uint16,
) int {
	_ = "STUB: not implemented"
	return 0
}

// intersection2by2Cardinality computes the cardinality of the intersection
func intersection2by2Cardinality(
	set1 []uint16,
	set2 []uint16,
) int {
	_ = "STUB: not implemented"
	return 0
}

// intersects2by2 computes whether the two sets intersect
func intersects2by2(
	set1 []uint16,
	set2 []uint16,
) bool {
	_ = "STUB: not implemented"
	return false
}

func intersects2by2Bool(
	set1 []uint16,
	set2 []uint16,
) bool {
	_ = "STUB: not implemented"
	return false
}

// (set2[k2] == set1[k1])

func onesidedgallopingintersect2by2Bool(
	smallset []uint16,
	largeset []uint16,
) bool {
	_ = "STUB: not implemented"
	return false
}

// (set2[k2] == set1[k1])

func localintersect2by2(
	set1 []uint16,
	set2 []uint16,
	buffer []uint16,
) int {
	_ = "STUB: not implemented"
	return 0
}

// (set2[k2] == set1[k1])

// / localintersect2by2Cardinality computes the cardinality of the intersection
func localintersect2by2Cardinality(
	set1 []uint16,
	set2 []uint16,
) int {
	_ = "STUB: not implemented"
	return 0
}

// (set2[k2] == set1[k1])

func advanceUntil(
	array []uint16,
	pos int,
	length int,
	min uint16,
) int {
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

func onesidedgallopingintersect2by2(
	smallset []uint16,
	largeset []uint16,
	buffer []uint16,
) int {
	_ = "STUB: not implemented"
	return 0
}

func onesidedgallopingintersect2by2Cardinality(
	smallset []uint16,
	largeset []uint16,
) int {
	_ = "STUB: not implemented"
	return 0
}

func binarySearch(array []uint16, ikey uint16) int { _ = "STUB: not implemented"; return 0 }

// searchResult provides information about a search request.
// The values will depend on the context of the search
type searchResult struct {
	value      uint16
	index      int
	exactMatch bool
}

// notFound returns a bool depending the search context
// For cases `previousValue` and `nextValue` if target is present in the slice
// this function will return `true` otherwise `false`
// For `nextAbsentValue` and `previousAbsentValue` this will only return `False`
func (sr *searchResult) notFound() bool { _ = "STUB: not implemented"; return false }

// outOfBounds indicates whether the target was outside the lower and upper bounds of the container
func (sr *searchResult) outOfBounds() bool { _ = "STUB: not implemented"; return false }

// binarySearchUntil is a helper function around binarySearchUntilWithBounds
// The user does not have to pass in the lower and upper bound
// The lower bound is taken to be `0` and the upper bound `len(array)-1`
func binarySearchUntil(array []uint16, target uint16) searchResult {
	_ = "STUB: not implemented"
	return *new(searchResult)
}

// binarySearchUntilWithBounds returns a `searchResult`.
// If an exact match is found the `searchResult{target, <index>, true}` will be returned, where `<index>` is
// `target`s index in `array`, and `result.notFound()` evaluates to `false`.
// If a match is not found, but `target` was in-bounds then the result.index will be the closest smaller value
// Example: [ 8,9,11,12] if the target was 10, then `searchResult{9, 1, false}` will be returned.
// If `target` was out of bounds `searchResult{0, -1, false}` will be returned.
func binarySearchUntilWithBounds(array []uint16, target uint16, lowIndex int, maxIndex int) searchResult {
	_ = "STUB: not implemented"
	return *new(searchResult)
}

// binarySearchPast is a wrapper around binarySearchPastWithBounds
// The user does not have to pass in the lower and upper bound
// The lower bound is taken to be `0` and the upper bound `len(array)-1`
func binarySearchPast(array []uint16, target uint16) searchResult {
	_ = "STUB: not implemented"
	return *new(searchResult)
}

// binarySearchPastWithBounds looks for the smallest value larger than or equal to `target`
// If `target` is out of bounds a `searchResult` indicating out of bounds is returned
// `target` does not have to exist in the slice.
//
// Example:
// Suppose the slice is [...10,13...] with `target` equal to 11
// The searchResult will have searchResult.value = 13
func binarySearchPastWithBounds(array []uint16, target uint16, lowIndex int, maxIndex int) searchResult {
	_ = "STUB: not implemented"
	return *new(searchResult)
}
