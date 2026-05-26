package roaring

//
// Copyright (c) 2016 by the roaring authors.
// Licensed under the Apache License, Version 2.0.
//
// We derive a few lines of code from the sort.Search
// function in the golang standard library. That function
// is Copyright 2009 The Go Authors, and licensed
// under the following BSD-style license.
/*
Copyright (c) 2009 The Go Authors. All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are
met:

   * Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.
   * Redistributions in binary form must reproduce the above
copyright notice, this list of conditions and the following disclaimer
in the documentation and/or other materials provided with the
distribution.
   * Neither the name of Google Inc. nor the names of its
contributors may be used to endorse or promote products derived from
this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

import (
	"errors"
)

// runContainer16 does run-length encoding of sets of
// uint16 integers.
type runContainer16 struct {
	// iv is a slice of sorted, non-overlapping, non-adjacent intervals.
	iv []interval16
}

// interval16 is the internal to runContainer16
// structure that maintains the individual [start, last]
// closed intervals.
type interval16 struct {
	start  uint16
	length uint16 // length minus 1
}

var (
	ErrRunIntervalsEmpty  = errors.New("run contained no interval")
	ErrRunNonSorted       = errors.New("runs were not sorted")
	ErrRunIntervalEqual   = errors.New("intervals were equal")
	ErrRunIntervalOverlap = errors.New("intervals overlapped or were continguous")
	ErrRunIntervalSize    = errors.New("too many intervals relative to data")
	MaxNumIntervals       = 2048
	MaxIntervalsSum       = 2048
)

func newInterval16Range(start, last uint16) interval16 {
	_ = "STUB: not implemented"
	return *new(interval16)
}

// runlen returns the count of integers in the interval.
func (iv interval16) runlen() int { _ = "STUB: not implemented"; return 0 }

func (iv interval16) last() uint16 { _ = "STUB: not implemented"; return 0 }

// String produces a human viewable string of the contents.
func (iv interval16) String() string { _ = "STUB: not implemented"; return "" }

func ivalString16(iv []interval16) string { _ = "STUB: not implemented"; return "" }

// String produces a human viewable string of the contents.
func (rc *runContainer16) String() string { _ = "STUB: not implemented"; return "" }

// addHelper helps build a runContainer16.
type addHelper16 struct {
	runstart      uint16
	runlen        uint16
	actuallyAdded uint16
	m             []interval16
	rc            *runContainer16
}

func (ah *addHelper16) storeIval(runstart, runlen uint16) { _ = "STUB: not implemented"; return }

func (ah *addHelper16) add(cur, prev uint16, i int) { _ = "STUB: not implemented"; return }

// ignore duplicates

// newRunContainerRange makes a new container made of just the specified closed interval [rangestart,rangelast]
func newRunContainer16Range(rangestart uint16, rangelast uint16) *runContainer16 {
	_ = "STUB: not implemented"
	return nil
}

// newRunContainer16FromVals makes a new container from vals.
//
// For efficiency, vals should be sorted in ascending order.
// Ideally vals should not contain duplicates, but we detect and
// ignore them. If vals is already sorted in ascending order, then
// pass alreadySorted = true. Otherwise, for !alreadySorted,
// we will sort vals before creating a runContainer16 of them.
// We sort the original vals, so this will change what the
// caller sees in vals as a side effect.
func newRunContainer16FromVals(alreadySorted bool, vals ...uint16) *runContainer16 {
	_ = "STUB: not implemented"
	// keep this in sync with newRunContainer16FromArray below
	return nil
}

// nothing more

// newRunContainer16FromBitmapContainer makes a new run container from bc,
// somewhat efficiently. For reference, see the Java
// https://github.com/RoaringBitmap/RoaringBitmap/blob/master/src/main/java/org/roaringbitmap/RunContainer.java#L145-L192
func newRunContainer16FromBitmapContainer(bc *bitmapContainer) *runContainer16 {
	_ = "STUB: not implemented"
	return nil
}

// index of current long in bitmap
// its value

// potentially multiword advance to first 1 bit

// wrap up, no more runs

// stuff 1s into number's LSBs

// find the next 0, potentially in a later word

// a final unterminated run of 1s

// now, zero out everything right of runEnd.

// We've lathered and rinsed, so repeat...

// newRunContainer16FromArray populates a new
// runContainer16 from the contents of arr.
func newRunContainer16FromArray(arr *arrayContainer) *runContainer16 {
	_ = "STUB: not implemented"
	// keep this in sync with newRunContainer16FromVals above
	return nil
}

// nothing more

// set adds the integers in vals to the set. Vals
// must be sorted in increasing order; if not, you should set
// alreadySorted to false, and we will sort them in place for you.
// (Be aware of this side effect -- it will affect the callers
// view of vals).
//
// If you have a small number of additions to an already
// big runContainer16, calling Add() may be faster.
func (rc *runContainer16) set(alreadySorted bool, vals ...uint16) {
	_ = "STUB: not implemented"
	return
}

// canMerge returns true iff the intervals
// a and b either overlap or they are
// contiguous and so can be merged into
// a single interval.
func canMerge16(a, b interval16) bool { _ = "STUB: not implemented"; return false }

// haveOverlap differs from canMerge in that
// it tells you if the intersection of a
// and b would contain an element (otherwise
// it would be the empty set, and we return
// false).
func haveOverlap16(a, b interval16) bool { _ = "STUB: not implemented"; return false }

// mergeInterval16s joins a and b into a
// new interval, and panics if it cannot.
func mergeInterval16s(a, b interval16) (res interval16) {
	_ = "STUB: not implemented"
	return *new(interval16)
}

// intersectInterval16s returns the intersection
// of a and b. The isEmpty flag will be true if
// a and b were disjoint.
func intersectInterval16s(a, b interval16) (res interval16, isEmpty bool) {
	_ = "STUB: not implemented"
	return *new(interval16), false
}

// union merges two runContainer16s, producing
// a new runContainer16 with the union of rc and b.
func (rc *runContainer16) union(b *runContainer16) *runContainer16 {
	_ = "STUB: not implemented"
	// rc is also known as 'a' here, but golint insisted we
	// call it rc for consistency with the rest of the methods.
	return nil
}

// next from a
// next from b

// merged holds the current merge output, which might
// get additional merges before being appended to m.

// is merged being used at the moment?

// currently considering this interval16 from a
// currently considering this interval16 from b

// we know that merged is disjoint from cura and curb

// !mergedUsed

// finish by merging anything remaining into merged we can:

// unionCardinality returns the cardinality of the merger of two runContainer16s,  the union of rc and b.
func (rc *runContainer16) unionCardinality(b *runContainer16) uint {
	_ = "STUB: not implemented"
	// rc is also known as 'a' here, but golint insisted we
	// call it rc for consistency with the rest of the methods.
	return 0
}

// next from a
// next from b

// merged holds the current merge output, which might
// get additional merges before being appended to m.

// is merged being used at the moment?

// currently considering this interval16 from a
// currently considering this interval16 from b

// we know that merged is disjoint from cura and curb
// m = append(m, merged)

// !mergedUsed

// m = append(m, cura)

// m = append(m, curb)

// finish by merging anything remaining into merged we can:

// m = append(m, merged)

// indexOfIntervalAtOrAfter is a helper for union.
func (rc *runContainer16) indexOfIntervalAtOrAfter(key int, startIndex int) int {
	_ = "STUB: not implemented"
	return 0
}

// intersect returns a new runContainer16 holding the
// intersection of rc (also known as 'a')  and b.
func (rc *runContainer16) intersect(b *runContainer16) *runContainer16 {
	_ = "STUB: not implemented"
	return nil
}

// isOverlap

// note that we change astart without advancing acuri,
// since we need to capture any 2ndary intersections with a.iv[acuri]

// note that we change bstart without advancing bcuri,
// since we need to capture any 2ndary intersections with b.iv[bcuri]

// neither had leftover, both completely consumed

// advance to next a interval

// advance to next b interval

// end for toploop

// intersectCardinality returns the cardinality of  the
// intersection of rc (also known as 'a')  and b.
func (rc *runContainer16) intersectCardinality(b *runContainer16) int {
	_ = "STUB: not implemented"
	return 0
}

// isOverlap

// note that we change astart without advancing acuri,
// since we need to capture any 2ndary intersections with a.iv[acuri]

// note that we change bstart without advancing bcuri,
// since we need to capture any 2ndary intersections with b.iv[bcuri]

// neither had leftover, both completely consumed

// advance to next a interval

// advance to next b interval

// end for toploop

// get returns true iff key is in the container.
func (rc *runContainer16) contains(key uint16) bool { _ = "STUB: not implemented"; return false }

// numIntervals returns the count of intervals in the container.
func (rc *runContainer16) numIntervals() int {
	_ = "STUB: not implemented"

	// searchRange returns alreadyPresent to indicate if the
	// key is already in one of our interval16s.
	//
	// If key is alreadyPresent, then whichInterval16 tells
	// you where.
	//
	// If key is not already present, then whichInterval16 is
	// set as follows:
	//
	//	a) whichInterval16 == len(rc.iv)-1 if key is beyond our
	//	   last interval16 in rc.iv;
	//
	//	b) whichInterval16 == -1 if key is before our first
	//	   interval16 in rc.iv;
	//
	//	c) whichInterval16 is set to the minimum index of rc.iv
	//	   which comes strictly before the key;
	//	   so  rc.iv[whichInterval16].last < key,
	//	   and  if whichInterval16+1 exists, then key < rc.iv[whichInterval16+1].start
	//	   (Note that whichInterval16+1 won't exist when
	//	   whichInterval16 is the last interval.)
	//
	// runContainer16.search always returns whichInterval16 < len(rc.iv).
	//
	// The search space is from startIndex to endxIndex. If endxIndex is set to zero, then there
	// no upper bound.
	return 0
}

func (rc *runContainer16) searchRange(key int, startIndex int, endxIndex int) (whichInterval16 int, alreadyPresent bool, numCompares int) {
	_ = "STUB: not implemented"
	return 0, false, 0
}

// sort.Search returns the smallest index i
// in [0, n) at which f(i) is true, assuming that on the range [0, n),
// f(i) == true implies f(i+1) == true.
// If there is no such index, Search returns n.

// For correctness, this began as verbatim snippet from
// sort.Search in the Go standard lib.
// We inline our comparison function for speed, and
// annotate with numCompares
// to observe and test that extra bounds are utilized.

// avoid overflow when computing h as the bisector
// i <= h < j

// end std lib snippet.

// The above is a simple in-lining and annotation of:
/*	below := sort.Search(n,
	func(i int) bool {
		return key < rc.iv[i].start
	})
*/

// all falses => key is >= start of all interval16s
// ... so does it belong to the last interval16?

// yes, it belongs to the last interval16

// no, it is beyond the last interval16.
// leave alreadyPreset = false

// INVAR: key is below rc.iv[below]

// key is before the first first interval16.
// leave alreadyPresent = false

// INVAR: key is >= rc.iv[below-1].start and
//        key is <  rc.iv[below].start

// is key in below-1 interval16?

// yes, it is. key is in below-1 interval16.

// INVAR: key >= rc.iv[below-1].endx && key < rc.iv[below].start
// leave alreadyPresent = false

// search returns alreadyPresent to indicate if the
// key is already in one of our interval16s.
//
// If key is alreadyPresent, then whichInterval16 tells
// you where.
//
// If key is not already present, then whichInterval16 is
// set as follows:
//
//	a) whichInterval16 == len(rc.iv)-1 if key is beyond our
//	   last interval16 in rc.iv;
//
//	b) whichInterval16 == -1 if key is before our first
//	   interval16 in rc.iv;
//
//	c) whichInterval16 is set to the maximum index of rc.iv
//	   which comes strictly before the key;
//	   so  rc.iv[whichInterval16].last < key,
//	   and  if whichInterval16+1 exists, then key < rc.iv[whichInterval16+1].start
//	   (Note that whichInterval16+1 won't exist when
//	   whichInterval16 is the last interval.)
//
// runContainer16.search always returns whichInterval16 < len(rc.iv).
func (rc *runContainer16) search(key int) (whichInterval16 int, alreadyPresent bool, numCompares int) {
	_ = "STUB: not implemented"
	return 0, false, 0

	// getCardinality returns the count of the integers stored in the
	// runContainer16. The running complexity depends on the size
	// of the container.
}

func (rc *runContainer16) getCardinality() int {
	_ = "STUB: not implemented"
	// have to compute it
	return 0
}

// isEmpty returns true if the container is empty.
// It runs in constant time.
func (rc *runContainer16) isEmpty() bool { _ = "STUB: not implemented"; return false }

// AsSlice decompresses the contents into a []uint16 slice.
func (rc *runContainer16) AsSlice() []uint16 { _ = "STUB: not implemented"; return nil }

// newRunContainer16 creates an empty run container.
func newRunContainer16() *runContainer16 { _ = "STUB: not implemented"; return nil }

// newRunContainer16CopyIv creates a run container, initializing
// with a copy of the supplied iv slice.
func newRunContainer16CopyIv(iv []interval16) *runContainer16 {
	_ = "STUB: not implemented"
	return nil
}

func (rc *runContainer16) Clone() *runContainer16 { _ = "STUB: not implemented"; return nil }

// newRunContainer16TakeOwnership returns a new runContainer16
// backed by the provided iv slice, which we will
// assume exclusive control over from now on.
func newRunContainer16TakeOwnership(iv []interval16) *runContainer16 {
	_ = "STUB: not implemented"
	return nil
}

const (
	baseRc16Size        = 2
	perIntervalRc16Size = 4
)

// see also runContainer16SerializedSizeInBytes(numRuns int) int

// getSizeInBytes returns the number of bytes of memory
// required by this runContainer16.
func (rc *runContainer16) getSizeInBytes() int { _ = "STUB: not implemented"; return 0 }

// runContainer16SerializedSizeInBytes returns the number of bytes of disk
// required to hold numRuns in a runContainer16.
func runContainer16SerializedSizeInBytes(numRuns int) int { _ = "STUB: not implemented"; return 0 }

// Add adds a single value k to the set.
func (rc *runContainer16) Add(k uint16) (wasNew bool) {
	_ = "STUB: not implemented"
	// TODO comment from runContainer16.java:
	// it might be better and simpler to do return
	// toBitmapOrArrayContainer(getCardinality()).add(k)
	// but note that some unit tests use this method to build up test
	// runcontainers without calling runOptimize
	return false
}

// already there

// we may need to extend the first run

// nope, k stands alone, starting the new first interval16.

// are we off the end? handle both index == n and index == n-1:

// INVAR: index and index+1 both exist, and k goes between them.
//
// Now: add k into the middle,
// possibly fusing with index or index+1 interval16
// and possibly resulting in fusing of two interval16s
// that had a one integer gap.

// are we fusing left and right by adding k?

// fuse into left

// remove redundant right

// are we an addition to left?

// yes

// are we an addition to right?

// yes

// k makes a standalone new interval16, inserted in the middle

// runIterator16 advice: you must call hasNext()
// before calling next()/peekNext() to insure there are contents.
type runIterator16 struct {
	rc            *runContainer16
	curIndex      int
	curPosInIndex uint16
}

// newRunIterator16 returns a new empty run container.
func (rc *runContainer16) newRunIterator16() *runIterator16 { _ = "STUB: not implemented"; return nil }

func (rc *runContainer16) iterate(cb func(x uint16) bool) bool {
	_ = "STUB: not implemented"
	return false
}

// hasNext returns false if calling next will panic. It
// returns true when there is at least one more value
// available in the iteration sequence.
func (ri *runIterator16) hasNext() bool { _ = "STUB: not implemented"; return false }

// next returns the next value in the iteration sequence.
func (ri *runIterator16) next() uint16 { _ = "STUB: not implemented"; return 0 }

// peekNext returns the next value in the iteration sequence without advancing the iterator
func (ri *runIterator16) peekNext() uint16 { _ = "STUB: not implemented"; return 0 }

// advanceIfNeeded advances as long as the next value is smaller than minval
func (ri *runIterator16) advanceIfNeeded(minval uint16) { _ = "STUB: not implemented"; return }

// interval cannot be -1 because of minval > peekNext

// if the minval is present, set the curPosIndex at the right position

// otherwise interval is set to to the minimum index of rc.iv
// which comes strictly before the key, that's why we set the next interval

// runReverseIterator16 advice: you must call hasNext()
// before calling next() to insure there are contents.
type runReverseIterator16 struct {
	rc            *runContainer16
	curIndex      int    // index into rc.iv
	curPosInIndex uint16 // offset in rc.iv[curIndex]
}

// newRunReverseIterator16 returns a new empty run iterator.
func (rc *runContainer16) newRunReverseIterator16() *runReverseIterator16 {
	_ = "STUB: not implemented"
	return nil
}

// hasNext returns false if calling next will panic. It
// returns true when there is at least one more value
// available in the iteration sequence.
func (ri *runReverseIterator16) hasNext() bool { _ = "STUB: not implemented"; return false }

// next returns the next value in the iteration sequence.
func (ri *runReverseIterator16) next() uint16 { _ = "STUB: not implemented"; return 0 }

func (rc *runContainer16) newManyRunIterator16() *runIterator16 {
	_ = "STUB: not implemented"
	return nil
}

// hs are the high bits to include to avoid needing to reiterate over the buffer in NextMany
func (ri *runIterator16) nextMany(hs uint32, buf []uint32) int { _ = "STUB: not implemented"; return 0 }

// start and end are inclusive

// add as many as you can from this seq

// allows BCE

// update values

// moreVals always fits in uint16

func (ri *runIterator16) nextMany64(hs uint64, buf []uint64) int {
	_ = "STUB: not implemented"
	return 0
}

// start and end are inclusive

// add as many as you can from this seq

// allows BCE

// update values

// moreVals always fits in uint16

// remove removes key from the container.
func (rc *runContainer16) removeKey(key uint16) (wasPresent bool) {
	_ = "STUB: not implemented"
	return false
}

// already removed, nothing to do.

// internal helper functions

func (rc *runContainer16) deleteAt(curIndex *int, curPosInIndex *uint16) {
	_ = "STUB: not implemented"
	return
}

// are we first, last, or in the middle of our interval16?

// our interval disappears

// curIndex stays the same, since the delete did
// the advance for us.

// no longer overflowable

// length

// our interval16 cannot disappear, else we would have been pos == 0, case first above.

// if we leave *curIndex alone, then Next() will work properly even after the delete.

// middle
// split into two, adding an interval16

// update curIndex and curPosInIndex

func have4Overlap16(astart, alast, bstart, blast int) bool { _ = "STUB: not implemented"; return false }

func intersectWithLeftover16(astart, alast, bstart, blast int) (isOverlap, isLeftoverA, isLeftoverB bool, leftoverstart int, intersection interval16) {
	_ = "STUB: not implemented"
	return false, false, false, 0, *new(interval16)
}

// do the intersection:

// alast == blast

func (rc *runContainer16) findNextIntervalThatIntersectsStartingFrom(startIndex int, key int) (index int, done bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// rc.search always returns w < len(rc.iv)

// not found and comes before lower bound startIndex,
// so just use the lower bound.

// also this bump up means that we are done

func sliceToString16(m []interval16) string { _ = "STUB: not implemented"; return "" }

// helper for invert
func (rc *runContainer16) invertlastInterval(origin uint16, lastIdx int) []interval16 {
	_ = "STUB: not implemented"
	return nil
}

// empty container

// invert splits

// invert returns a new container (not inplace), that is
// the inversion of rc. For each bit b in rc, the
// returned value has !b
func (rc *runContainer16) invert() *runContainer16 { _ = "STUB: not implemented"; return nil }

// invertlastInteval will add both intervals (b) and (c) in
// diagram below.

// INVAR: i and cur are not the last interval, there is a next at i+1
//
// ........[cur.start, cur.last] ...... [next.start, next.last]....
//    ^                             ^                           ^
//   (a)                           (b)                         (c)
//
// Now: we add interval (a); but if (a) is empty, for cur.start==0, we skip it.

func (iv interval16) equal(b interval16) bool { _ = "STUB: not implemented"; return false }

func (iv interval16) isSuperSetOf(b interval16) bool { _ = "STUB: not implemented"; return false }

func (iv interval16) isNonContiguousDisjoint(b interval16) bool {
	_ = "STUB: not implemented"
	// cover the zero start case
	return false
}

func (iv interval16) subtractInterval(del interval16) (left []interval16, delcount int) {
	_ = "STUB: not implemented"
	return nil, 0
}

func (rc *runContainer16) isubtract(del interval16) { _ = "STUB: not implemented"; return }

// already done.

// done

// INVAR there is some intersection between rc and del

// some intervals will remain

// would overwrite values in iv b/c res0 can have len 2. so
// write to origiv instead.

//	rc.iv = append(pre, caboose...)
//	return

// shrink

// stay the same

// changeSize > 0 is only possible when ilast == istart.
// Hence we now know: changeSize == 1 and len(res0) == 2

// len(rc.iv) is correct now, no need to rc.iv = rc.iv[:newSize]

// copy the tail into place

// copy the new item(s) into place

// we get to discard whole intervals

// from the search() definition:

// if del.start is not present, then istart is
// set as follows:
//
//  a) istart == n-1 if del.start is beyond our
//     last interval16 in rc.iv;
//
//  b) istart == -1 if del.start is before our first
//     interval16 in rc.iv;
//
//  c) istart is set to the minimum index of rc.iv
//     which comes strictly before the del.start;
//     so  del.start > rc.iv[istart].last,
//     and  if istart+1 exists, then del.start < rc.iv[istart+1].startx

// if del.last is not present, then ilast is
// set as follows:
//
//  a) ilast == n-1 if del.last is beyond our
//     last interval16 in rc.iv;
//
//  b) ilast == -1 if del.last is before our first
//     interval16 in rc.iv;
//
//  c) ilast is set to the minimum index of rc.iv
//     which comes strictly before the del.last;
//     so  del.last > rc.iv[ilast].last,
//     and  if ilast+1 exists, then del.last < rc.iv[ilast+1].start

// INVAR: istart >= 0

// INVAR: ilast < n-1

// we can only shrink or stay the same size
// i.e. we either eliminate the whole interval,
// or just cut off the right side.

// len(res) must be 1

// we can only shrink or stay the same size

// move the tail first to make room for res1

// compute rc minus b, and return the result as a new value (not inplace).
// port of run_container_andnot from CRoaring...
// https://github.com/RoaringBitmap/CRoaring/blob/master/src/containers/run.c#L435-L496
func (rc *runContainer16) AndNotRunContainer16(b *runContainer16) *runContainer16 {
	_ = "STUB: not implemented"
	return nil
}

// output the first run

// exit the second run

//   a: [             ]
//   b:            [    ]
// alast >= bstart
// blast >= astart

func (rc *runContainer16) numberOfRuns() (nr int) { _ = "STUB: not implemented"; return 0 }

func (rc *runContainer16) containerType() contype { _ = "STUB: not implemented"; return *new(contype) }

func (rc *runContainer16) equals16(srb *runContainer16) bool {
	_ = "STUB: not implemented"
	// Check if the containers are the same object.
	return false
}

// compile time verify we meet interface requirements
var _ container = &runContainer16{}

func (rc *runContainer16) clone() container { _ = "STUB: not implemented"; return *new(container) }

func (rc *runContainer16) minimum() uint16 {
	_ = "STUB: not implemented"
	// assume not empty
	return 0
}

func (rc *runContainer16) safeMinimum() (uint16, error) { _ = "STUB: not implemented"; return 0, nil }

func (rc *runContainer16) maximum() uint16 { _ = "STUB: not implemented"; return 0 }

// assume not empty

func (rc *runContainer16) safeMaximum() (uint16, error) { _ = "STUB: not implemented"; return 0, nil }

// assume not empty

func (rc *runContainer16) isFull() bool { _ = "STUB: not implemented"; return false }

func (rc *runContainer16) and(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// Important: there is no reason to believe that the
// result of intersecting two run containers is itself
// a run container. Hence we convert to efficient container.
// We only use run containers when they are efficient.

func (rc *runContainer16) andCardinality(a container) int { _ = "STUB: not implemented"; return 0 }

// andBitmapContainer finds the intersection of rc and b.
func (rc *runContainer16) andBitmapContainer(bc *bitmapContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (rc *runContainer16) andArrayCardinality(ac *arrayContainer) int {
	_ = "STUB: not implemented"
	return 0
}

// won't happen in actual code

func (rc *runContainer16) iand(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// Important: there is no reason to believe that the
// result of intersecting two run containers is itself
// a run container. Hence we convert to efficient container.
// We only use run containers when they are efficient.

// inplace intersection with array is not supported
// It is likely not very useful either.

// inplace intersection with bitmap is not supported
// It is very difficult to do this inplace and likely not useful.

func (rc *runContainer16) inplaceIntersect(rc2 *runContainer16) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (rc *runContainer16) andArray(ac *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (rc *runContainer16) andNot(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (rc *runContainer16) fillLeastSignificant16bits(x []uint32, i int, mask uint32) int {
	_ = "STUB: not implemented"
	return 0
}

func (rc *runContainer16) getShortIterator() shortPeekable {
	_ = "STUB: not implemented"
	return *new(shortPeekable)
}

func (rc *runContainer16) getReverseIterator() shortIterable {
	_ = "STUB: not implemented"
	return *new(shortIterable)
}

func (rc *runContainer16) getManyIterator() manyIterable {
	_ = "STUB: not implemented"
	return *new(manyIterable)
}

type runUnsetIterator16 struct {
	rc       *runContainer16
	curIndex int
	nextVal  int
}

func (rc *runContainer16) newRunUnsetIterator16() *runUnsetIterator16 {
	_ = "STUB: not implemented"
	return nil
}

func (rui *runUnsetIterator16) hasNext() bool { _ = "STUB: not implemented"; return false }

func (rui *runUnsetIterator16) next() uint16 { _ = "STUB: not implemented"; return 0 }

func (rui *runUnsetIterator16) peekNext() uint16 { _ = "STUB: not implemented"; return 0 }

func (rui *runUnsetIterator16) advanceIfNeeded(minval uint16) { _ = "STUB: not implemented"; return }

func (rc *runContainer16) getUnsetIterator() shortPeekable {
	_ = "STUB: not implemented"
	return *new(shortPeekable)
}

// add the values in the range [firstOfRange, endx). endx
// is still abe to express 2^16 because it is an int not an uint16.
func (rc *runContainer16) iaddRange(firstOfRange, endx int) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// remove the values in the range [firstOfRange,endx)
func (rc *runContainer16) iremoveRange(firstOfRange, endx int) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// empty removal

// not flip the values in the range [firstOfRange,endx)
func (rc *runContainer16) not(firstOfRange, endx int) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// Not flips the values in the range [firstOfRange,endx).
// This is not inplace. Only the returned value has the flipped bits.
//
// Currently implemented as (!A intersect B) union (A minus B),
// where A is rc, and B is the supplied [firstOfRange, endx) interval.
//
// TODO(time optimization): convert this to a single pass
// algorithm by copying AndNotRunContainer16() and modifying it.
// Current routine is correct but
// makes 2 more passes through the arrays than should be
// strictly necessary. Measure both ways though--this may not matter.
func (rc *runContainer16) Not(firstOfRange, endx int) *runContainer16 {
	_ = "STUB: not implemented"
	return nil
}

// algo:
// (!A intersect B) union (A minus B)

// equals is now logical equals; it does not require the
// same underlying container type.
func (rc *runContainer16) equals(o container) bool { _ = "STUB: not implemented"; return false }

// maybe value instead of pointer

// Check if the containers are the same object.

// use generic comparison

// k := 0

// k++

func (rc *runContainer16) iaddReturnMinimized(x uint16) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (rc *runContainer16) iadd(x uint16) (wasNew bool) { _ = "STUB: not implemented"; return false }

func (rc *runContainer16) iremoveReturnMinimized(x uint16) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (rc *runContainer16) iremove(x uint16) bool { _ = "STUB: not implemented"; return false }

func (rc *runContainer16) or(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (rc *runContainer16) orCardinality(a container) int { _ = "STUB: not implemented"; return 0 }

// orBitmapContainer finds the union of rc and bc.
func (rc *runContainer16) orBitmapContainer(bc *bitmapContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (rc *runContainer16) andBitmapContainerCardinality(bc *bitmapContainer) int {
	_ = "STUB: not implemented"
	return 0
}

// bc.computeCardinality()

func (rc *runContainer16) orBitmapContainerCardinality(bc *bitmapContainer) int {
	_ = "STUB: not implemented"
	return 0
}

// orArray finds the union of rc and ac.
func (rc *runContainer16) orArray(ac *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// orArray finds the union of rc and ac.
func (rc *runContainer16) orArrayCardinality(ac *arrayContainer) int {
	_ = "STUB: not implemented"
	return 0
}

func (rc *runContainer16) ior(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (rc *runContainer16) inplaceUnion(rc2 *runContainer16) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// Such code should not be used as it will not preserve the container invariants:
//func (rc *runContainer16) iorBitmapContainer(bc *bitmapContainer) container {
//	it := bc.getShortIterator()
//	for it.hasNext() {
//		rc.Add(it.next())
//	}
//	return rc
//}

func (rc *runContainer16) iorArray(ac *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// TODO: perform the union algorithm in-place using rc.iv
// this can be done with methods like the in-place array container union
// but maybe lazily moving the remaining elements back.

func runArrayUnionToRuns(rc *runContainer16, ac *arrayContainer) ([]interval16, uint16) {
	_ = "STUB: not implemented"
	return nil, 0
}

// have to find the first range
// options are
// 1. from array container
// 2. from run container

// lazyIOR is described (not yet implemented) in
// this nice note from @lemire on
// https://github.com/RoaringBitmap/roaring/pull/70#issuecomment-263613737
//
// Description of lazyOR and lazyIOR from @lemire:
//
// Lazy functions are optional and can be simply
// wrapper around non-lazy functions.
//
// The idea of "laziness" is as follows. It is
// inspired by the concept of lazy evaluation
// you might be familiar with (functional programming
// and all that). So a roaring bitmap is
// such that all its containers are, in some
// sense, chosen to use as little memory as
// possible. This is nice. Also, all bitsets
// are "cardinality aware" so that you can do
// fast rank/select queries, or query the
// cardinality of the whole bitmap... very fast,
// without latency.
//
// However, imagine that you are aggregating 100
// bitmaps together. So you OR the first two, then OR
// that with the third one and so forth. Clearly,
// intermediate bitmaps don't need to be as
// compressed as possible, right? They can be
// in a "dirty state". You only need the end
// result to be in a nice state... which you
// can achieve by calling repairAfterLazy at the end.
//
// The Java/C code does something special for
// the in-place lazy OR runs. The idea is that
// instead of taking two run containers and
// generating a new one, we actually try to
// do the computation in-place through a
// technique invented by @gssiyankai (pinging him!).
// What you do is you check whether the host
// run container has lots of extra capacity.
// If it does, you move its data at the end of
// the backing array, and then you write
// the answer at the beginning. What this
// trick does is minimize memory allocations.
func (rc *runContainer16) lazyIOR(a container) container {
	_ = "STUB: not implemented"
	// not lazy at the moment
	return *new(container)
}

// lazyOR is described above in lazyIOR.
func (rc *runContainer16) lazyOR(a container) container {
	_ = "STUB: not implemented"
	// not lazy at the moment
	return *new(container)
}

func (rc *runContainer16) intersects(a container) bool {
	_ = "STUB: not implemented"
	// TODO: optimize by doing inplace/less allocation
	return false
}

func (rc *runContainer16) xor(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (rc *runContainer16) ixor(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (rc *runContainer16) ixorArray(value2 *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (rc *runContainer16) ixorBitmap(value2 *bitmapContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (rc *runContainer16) ixorRunContainer16(value2 *runContainer16) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (rc *runContainer16) iandNot(a container) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// flip the values in the range [firstOfRange,endx)
func (rc *runContainer16) inot(firstOfRange, endx int) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// TODO: minimize copies, do it all inplace; not() makes a copy.

func (rc *runContainer16) rank(x uint16) int { _ = "STUB: not implemented"; return 0 }

// getCardinalityInRange returns the number of values in the half-open range [start, end).
func (rc *runContainer16) getCardinalityInRange(start, end uint) int {
	_ = "STUB: not implemented"
	return 0
}

// end is exclusive, so the last included value is end-1.

// Find the interval containing or just before 'start'.

// Find the interval containing or just before 'last'.

// Both are before the first interval → nothing in range.

// Determine the effective first interval index to start counting from.

// start falls between intervals or before the first; next interval is wStart+1.

// Determine the effective last interval index.

// last falls between intervals; the last fully-before interval is wEnd,
// but its values are all < start of next, and we need values <= last.
// All of wEnd's values are < start (if wEnd < firstIdx), we handle below.

// If start and end land in the same interval (or there's only one relevant).

// Clamp

// Handle the first interval (may be partial if start > iv[firstIdx].start).

// start is inside this interval; count from start to end of interval.

// start is before this interval; count the full interval.

// Sum full intervals strictly between firstIdx and lastIdx.

// Handle the last interval (may be partial if last < iv[lastIdx].last()).

// last is inside this interval; count from start of interval to last.

// last is beyond this interval; count the full interval.
// But wait — if !endInside, then last < iv[lastIdx].start,
// which means lastIdx's values are all > last. Don't count it.
// Actually, wEnd == lastIdx means iv[lastIdx].start <= last (from search semantics).
// If !endInside and wEnd == lastIdx, last is between iv[lastIdx].last() and iv[lastIdx+1].start.
// That means last >= iv[lastIdx].last()+1, so we can count the full interval.

func (rc *runContainer16) selectInt(x uint16) int { _ = "STUB: not implemented"; return 0 }

func (rc *runContainer16) andNotRunContainer16(b *runContainer16) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (rc *runContainer16) andNotArray(ac *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (rc *runContainer16) andNotBitmap(bc *bitmapContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (rc *runContainer16) toBitmapContainer() *bitmapContainer {
	_ = "STUB: not implemented"
	return nil
}

func (rc *runContainer16) iandNotRunContainer16(x2 *runContainer16) container {
	_ = "STUB: not implemented"
	// TODO: check size and optimize the return value
	return *new(container)
}

func (rc *runContainer16) iandNotArray(ac *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (rc *runContainer16) iandNotBitmap(bc *bitmapContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (rc *runContainer16) xorRunContainer16(x2 *runContainer16) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (rc *runContainer16) xorArray(ac *arrayContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (rc *runContainer16) xorBitmap(bc *bitmapContainer) container {
	_ = "STUB: not implemented"
	return *new(container)
}

// convert to bitmap or array *if needed*
func (rc *runContainer16) toEfficientContainer() container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (rc *runContainer16) toEfficientContainerFromCardinality(card int) container {
	_ = "STUB: not implemented"
	return *new(container)
}

func (rc *runContainer16) toArrayContainer() *arrayContainer { _ = "STUB: not implemented"; return nil }

func newRunContainer16FromContainer(c container) *runContainer16 {
	_ = "STUB: not implemented"
	return nil
}

// And finds the intersection of rc and b.
func (rc *runContainer16) And(b *Bitmap) *Bitmap { _ = "STUB: not implemented"; return nil }

// Xor returns the exclusive-or of rc and b.
func (rc *runContainer16) Xor(b *Bitmap) *Bitmap { _ = "STUB: not implemented"; return nil }

// Or returns the union of rc and b.
func (rc *runContainer16) Or(b *Bitmap) *Bitmap { _ = "STUB: not implemented"; return nil }

// serializedSizeInBytes returns the number of bytes of memory
// required by this runContainer16. This is for the
// Roaring format, as specified https://github.com/RoaringBitmap/RoaringFormatSpec/
func (rc *runContainer16) serializedSizeInBytes() int {
	_ = "STUB: not implemented"
	// number of runs in one uint16, then each run
	// needs two more uint16
	return 0
}

func (rc *runContainer16) addOffset(x uint16) (container, container) {
	_ = "STUB: not implemented"
	return *new(container), *new(container)
}

// Some elements will fall into low part, allocate a container.
// Checking the first one is enough because they are ordered.

// Some elements will fall into high part, allocate a container.
// Checking the last one is enough because they are ordered.

// Ensure proper nil interface.

// nextValue returns either the `target` if found or the next larger value.
// If the target is in the interior or a run then `target` will be returned
// Ex: If our run structure resmembles [[a,c], [d,f]] with a <= target <= c then `target` will be returned.
// Ex: If c < target < d then d is returned.
// Ex: If target < a then a is returned
// if the target > max, this is out of bounds and -1 is returned
func (rc *runContainer16) nextValue(target uint16) int { _ = "STUB: not implemented"; return 0 }

// The if relies on the non-contiguous nature of runs.
// If we have two runs [a,b] and another run [c,d]
// We can rely on the invariant that b+1 < c
// We will return c

// nextAbsentValue returns the next absent value.
// By construction the next absent value will be located between gaps in runs
//
// Ex: if our runs resemble [[a,b],[c,d]] and a <= target <= b  then b+1 will not be equal to c, b+1 will be returned
// Ex: if target < a then target is returned
// Ex: if target > d then target is returned
func (rc *runContainer16) nextAbsentValue(target uint16) int { _ = "STUB: not implemented"; return 0 }

// previousValue will return the previous present value
// If the target is in the interior of a run  then `target` will be returned
//
// Example:
// If our run structure resmembles [[a,c], [d,f]] with a <= target  <= c then target will be returned.
// If c < target < d then c is returned.
// if target > f then f is returned
// if the target is less than a, this is out of bounds and -1 is returned
func (rc *runContainer16) previousValue(target uint16) int { _ = "STUB: not implemented"; return 0 }

// previousAbsentValue will return the previous absent value
// If the target is in the interior of a run then then the start of the range minus 1 will be returned
//
// Example:
// If our run structure resmembles [[x,z], [a,c], [d,f]] with a <= target  <= c then a-1 will be returned.
// if the target < x then target is returned
// if target > f then target is returned
func (rc *runContainer16) previousAbsentValue(target uint16) int {
	_ = "STUB: not implemented"
	return 0
}

// isNonContiguousDisjoint returns an error if the intervals overlap e.g have non-empty intersection
func isNonContiguousDisjoint(outer interval16, inner interval16) error {
	_ = "STUB: not implemented"
	return nil
}

// validate checks the run container referential integrity
// Ensures runs are not degenerate, non-contiguous and non-overlapping
func (rc *runContainer16) validate() error { _ = "STUB: not implemented"; return nil }

// The length being stored is the actual length - 1.
// So we need to add 1 to get the actual length.
// It is not possible to have a run with length 0.

// only check the start of runs
// if the run length overlap the next check will catch that.

/*
		if number of distinct values in the container >= 2048 then
	    check that the number of runs is no more than 2047
	    (otherwise you could use a bitset container)
		else
	    check that the number of runs < (number of distinct values) / 2
	    (otherwise you could use an array container)
*/

// this is always ok:
