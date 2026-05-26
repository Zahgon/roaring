/*
# Instructions for smat testing for roaring

[smat](https://github.com/mschoch/smat) is a framework that provides
state machine assisted fuzz testing.

To run the smat tests for roaring...

## Prerequisites

Go 1.18 or later (for native fuzzing support).

## Steps

1. Generate initial smat corpus:
```
go test -tags=gofuzz -run=TestGenerateSmatCorpus
```
You should see a directory `workdir` created with initial corpus files.

2. Run the fuzz test:
```
go test -run='^$' -fuzz=FuzzSmat -fuzztime=300s -timeout=60s
```

Adjust `-fuzztime` as needed for longer or shorter runs. If crashes are found,
check the test output and the reproducer files in the `workdir` directory.
You may copy the reproducers to roaring_tests.go
*/

package roaring

import (
	"slices"

	"github.com/bits-and-blooms/bitset"
	"github.com/mschoch/smat"
)

// The native fuzz entry point lives in a _test.go file so the go test
// fuzz engine discovers it. See smat_fuzz_test.go for the fuzz wrapper.

var smatDebug = true

const max_value = 1048576
const max_pairs = 10

func smatLog(prefix, format string, args ...interface{}) { _ = "STUB: not implemented"; return }

type smatContext struct {
	pairs []*smatPair

	// Two registers, x & y.
	x int
	y int

	actions int
	// per-context last action for this fuzz worker
	lastAction *actionRecord
}

// actionRecord stores a snapshot of the state just before an action runs.
type actionRecord struct {
	Name          string
	X, Y          int
	PairSnapshots []string // base64-encoded MarshalBinary of each pair's Bitmap
}

type smatPair struct {
	bm *Bitmap
	bs *bitset.BitSet
	// parent context (nil if unknown)
	ctx *smatContext
}

// ------------------------------------------------------------------

var smatActionMap = smat.ActionMap{
	smat.ActionID('X'): smatAction("x++", smatWrap(func(c *smatContext) { c.x = (c.x + 1) % max_value })),
	smat.ActionID('x'): smatAction("x--", smatWrap(func(c *smatContext) { c.x = (c.x - 1 + max_value) % max_value })),
	smat.ActionID('Y'): smatAction("y++", smatWrap(func(c *smatContext) { c.y = (c.y + 1) % max_value })),
	smat.ActionID('y'): smatAction("y--", smatWrap(func(c *smatContext) { c.y = (c.y - 1 + max_value) % max_value })),
	smat.ActionID('*'): smatAction("x*y", smatWrap(func(c *smatContext) { c.x = (c.x * c.y) % max_value })),
	smat.ActionID('<'): smatAction("x<<", smatWrap(func(c *smatContext) { c.x = (c.x << 1) % max_value })),

	smat.ActionID('^'): smatAction("swap", smatWrap(func(c *smatContext) { c.x, c.y = c.y, c.x })),

	smat.ActionID('['): smatAction(" pushPair", smatWrap(smatPushPair)),
	smat.ActionID(']'): smatAction(" popPair", smatWrap(smatPopPair)),

	smat.ActionID('B'): smatAction(" setBit", smatWrap(smatSetBit)),
	smat.ActionID('b'): smatAction(" removeBit", smatWrap(smatRemoveBit)),

	smat.ActionID('o'): smatAction(" or", smatWrap(smatOr)),
	smat.ActionID('a'): smatAction(" and", smatWrap(smatAnd)),
	smat.ActionID('z'): smatAction(" xor", smatWrap(smatXor)),

	smat.ActionID('#'): smatAction(" cardinality", smatWrap(smatCardinality)),

	smat.ActionID('O'): smatAction(" orCardinality", smatWrap(smatOrCardinality)),
	smat.ActionID('A'): smatAction(" andCardinality", smatWrap(smatAndCardinality)),
	smat.ActionID('Z'): smatAction(" xorCardinality", smatWrap(smatXorCardinality)),

	smat.ActionID('c'): smatAction(" clear", smatWrap(smatClear)),
	smat.ActionID('r'): smatAction(" runOptimize", smatWrap(smatRunOptimize)),

	smat.ActionID('e'): smatAction(" isEmpty", smatWrap(smatIsEmpty)),

	smat.ActionID('i'): smatAction(" intersects", smatWrap(smatIntersects)),

	smat.ActionID('f'): smatAction(" flip", smatWrap(smatFlip)),

	smat.ActionID('-'): smatAction(" difference", smatWrap(smatDifference)),
}

var smatRunningPercentActions []smat.PercentAction

func init() {
	var ids []int
	for actionId := range smatActionMap {
		ids = append(ids, int(actionId))
	}
	slices.Sort(ids)

	pct := 100 / len(smatActionMap)
	for _, actionId := range ids {
		smatRunningPercentActions = append(smatRunningPercentActions,
			smat.PercentAction{Percent: pct, Action: smat.ActionID(actionId)})
	}

	smatActionMap[smat.ActionID('S')] = smatAction("SETUP", smatSetupFunc)
	smatActionMap[smat.ActionID('T')] = smatAction("TEARDOWN", smatTeardownFunc)
}

// We only have one smat state: running.
func smatRunning(next byte) smat.ActionID { _ = "STUB: not implemented"; return *new(smat.ActionID) }

func smatAction(name string, f func(ctx smat.Context) (smat.State, error)) func(smat.Context) (smat.State, error) {
	_ = "STUB: not implemented"
	return nil
}

// Snapshot all pairs' bitmaps (base64 of MarshalBinary) before action

// record per-context last action (no global mutex required)

// catch panics inside action to dump a repro and stack before re-panicking

// best-effort: write quick repro with lastAction from context

// similar to checkEquals repro

// perform the action that caused panic

// saveReproFile writes the given repro content to workdir/<prefix>_<ts>_test.go
// or falls back to the OS temp dir. Returns full path or error.
func saveReproFile(prefix string, ts int64, content string) (string, error) {
	_ = "STUB: not implemented"
	// try workdir
	return "", nil
}

// fallback to temp

// Creates an smat action func based on a simple callback.
func smatWrap(cb func(c *smatContext)) func(smat.Context) (next smat.State, err error) {
	_ = "STUB: not implemented"
	return nil
}

// Invokes a callback function with the input v bounded to len(c.pairs).
func (c *smatContext) withPair(v int, cb func(*smatPair)) { _ = "STUB: not implemented"; return }

// ------------------------------------------------------------------

func smatSetupFunc(ctx smat.Context) (next smat.State, err error) {
	_ = "STUB: not implemented"
	return *new(smat.State), nil
}

func smatTeardownFunc(ctx smat.Context) (next smat.State, err error) {
	_ = "STUB: not implemented"

	// ------------------------------------------------------------------
	return *new(smat.State), nil
}

func smatPushPair(c *smatContext) { _ = "STUB: not implemented"; return }

func smatPopPair(c *smatContext) { _ = "STUB: not implemented"; return }

func smatSetBit(c *smatContext) { _ = "STUB: not implemented"; return }

func smatRemoveBit(c *smatContext) { _ = "STUB: not implemented"; return }

func smatAnd(c *smatContext) { _ = "STUB: not implemented"; return }

func smatOr(c *smatContext) { _ = "STUB: not implemented"; return }

func smatXor(c *smatContext) { _ = "STUB: not implemented"; return }

func smatAndCardinality(c *smatContext) { _ = "STUB: not implemented"; return }

func smatOrCardinality(c *smatContext) { _ = "STUB: not implemented"; return }

func smatXorCardinality(c *smatContext) { _ = "STUB: not implemented"; return }

func smatRunOptimize(c *smatContext) { _ = "STUB: not implemented"; return }

func smatClear(c *smatContext) { _ = "STUB: not implemented"; return }

func smatCardinality(c *smatContext) { _ = "STUB: not implemented"; return }

func smatIsEmpty(c *smatContext) { _ = "STUB: not implemented"; return }

func smatIntersects(c *smatContext) { _ = "STUB: not implemented"; return }

func smatFlip(c *smatContext) { _ = "STUB: not implemented"; return }

func smatDifference(c *smatContext) { _ = "STUB: not implemented"; return }

func (p *smatPair) checkEquals() { _ = "STUB: not implemented"; return }

// marshal current bitmap

// collect last action summary from context (per-worker)

// If debugging enabled, log extra info

// build a reproducible test snippet that reconstructs the bitmap and replays the failing action

// use the snapshot of the modified pair

// assume the modified pair is x % len(pairs), but since pairs are in order, and x is lastAction.X

// perform the action

// print the repro snippet for the developer

// also write the repro snippet to a timestamped file in workdir/

func (p *smatPair) Validate() { _ = "STUB: not implemented"; return }

func (p *smatPair) equalsBitSet(a *bitset.BitSet, b *Bitmap) bool {
	_ = "STUB: not implemented"
	return false
}
