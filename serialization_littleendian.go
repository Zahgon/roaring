//go:build (386 && !appengine) || (amd64 && !appengine) || (arm && !appengine) || (arm64 && !appengine) || (ppc64le && !appengine) || (mipsle && !appengine) || (mips64le && !appengine) || (mips64p32le && !appengine) || (wasm && !appengine)
// +build 386,!appengine amd64,!appengine arm,!appengine arm64,!appengine ppc64le,!appengine mipsle,!appengine mips64le,!appengine mips64p32le,!appengine wasm,!appengine

package roaring

import (
	"errors"
	"io"
)

func (ac *arrayContainer) writeTo(stream io.Writer) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (bc *bitmapContainer) writeTo(stream io.Writer) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func uint64SliceAsByteSlice(slice []uint64) []byte { _ = "STUB: not implemented"; return nil }

func uint16SliceAsByteSlice(slice []uint16) []byte { _ = "STUB: not implemented"; return nil }

func interval16SliceAsByteSlice(slice []interval16) []byte { _ = "STUB: not implemented"; return nil }

func (bc *bitmapContainer) asLittleEndianByteSlice() []byte { _ = "STUB: not implemented"; return nil }

// Deserialization code follows

// //
// These methods (byteSliceAsUint16Slice,...) do not make copies,
// they are pointer-based (unsafe). The caller is responsible to
// ensure that the input slice does not get garbage collected, deleted
// or modified while you hold the returned slince.
// //
func byteSliceAsUint16Slice(slice []byte) (result []uint16) {
	_ = "STUB: not implemented" // here we create a new slice holder
	return nil
}

func byteSliceAsUint64Slice(slice []byte) (result []uint64) { _ = "STUB: not implemented"; return nil }

func byteSliceAsInterval16Slice(slice []byte) (result []interval16) {
	_ = "STUB: not implemented"
	return nil
}

func byteSliceAsContainerSlice(slice []byte) (result []container) {
	_ = "STUB: not implemented"
	return nil
}

func byteSliceAsBitsetSlice(slice []byte) (result []bitmapContainer) {
	_ = "STUB: not implemented"
	return nil
}

func byteSliceAsArraySlice(slice []byte) (result []arrayContainer) {
	_ = "STUB: not implemented"
	return nil
}

func byteSliceAsRun16Slice(slice []byte) (result []runContainer16) {
	_ = "STUB: not implemented"
	return nil
}

func byteSliceAsBoolSlice(slice []byte) (result []bool) { _ = "STUB: not implemented"; return nil }

// FrozenView creates a static view of a serialized bitmap stored in buf.
// It uses CRoaring's frozen bitmap format.
//
// The format specification is available here:
// https://github.com/RoaringBitmap/CRoaring/blob/2c867e9f9c9e2a3a7032791f94c4c7ae3013f6e0/src/roaring.c#L2756-L2783
//
// The provided byte array (buf) is expected to be a constant.
// The function makes the best effort attempt not to copy data.
// Only little endian is supported. The function will err if it detects a big
// endian serialized file.
// You should take care not to modify buff as it will likely result in
// unexpected program behavior.
// If said buffer comes from a memory map, it's advisable to give it read
// only permissions, either at creation or by calling Mprotect from the
// golang.org/x/sys/unix package.
//
// Resulting bitmaps are effectively immutable in the following sense:
// a copy-on-write marker is used so that when you modify the resulting
// bitmap, copies of selected data (containers) are made.
// You should *not* change the copy-on-write status of the resulting
// bitmaps (SetCopyOnWrite).
//
// If buf becomes unavailable, then a bitmap created with
// FromBuffer would be effectively broken. Furthermore, any
// bitmap derived from this bitmap (e.g., via Or, And) might
// also be broken. Thus, before making buf unavailable, you should
// call CloneCopyOnWriteContainers on all such bitmaps.
func (rb *Bitmap) FrozenView(buf []byte) error { _ = "STUB: not implemented"; return nil }

func (rb *Bitmap) MustFrozenView(buf []byte) error { _ = "STUB: not implemented"; return nil }

/* Verbatim specification from CRoaring.
 *
 * FROZEN SERIALIZATION FORMAT DESCRIPTION
 *
 * -- (beginning must be aligned by 32 bytes) --
 * <bitset_data> uint64_t[BITSET_CONTAINER_SIZE_IN_WORDS * num_bitset_containers]
 * <run_data>    rle16_t[total number of rle elements in all run containers]
 * <array_data>  uint16_t[total number of array elements in all array containers]
 * <keys>        uint16_t[num_containers]
 * <counts>      uint16_t[num_containers]
 * <typecodes>   uint8_t[num_containers]
 * <header>      uint32_t
 *
 * <header> is a 4-byte value which is a bit union of frozenCookie (15 bits)
 * and the number of containers (17 bits).
 *
 * <counts> stores number of elements for every container.
 * Its meaning depends on container type.
 * For array and bitset containers, this value is the container cardinality minus one.
 * For run container, it is the number of rle_t elements (n_runs).
 *
 * <bitset_data>,<array_data>,<run_data> are flat arrays of elements of
 * all containers of respective type.
 *
 * <*_data> and <keys> are kept close together because they are not accessed
 * during deserilization. This may reduce IO in case of large mmaped bitmaps.
 * All members have their native alignments during deserilization except <header>,
 * which is not guaranteed to be aligned by 4 bytes.
 */
const frozenCookie = 13766

var (
	// ErrFrozenBitmapInvalidCookie is returned when the header does not contain the frozenCookie.
	ErrFrozenBitmapInvalidCookie = errors.New("header does not contain the frozenCookie")
	// ErrFrozenBitmapBigEndian is returned when the header is big endian.
	ErrFrozenBitmapBigEndian = errors.New("loading big endian frozen bitmaps is not supported")
	// ErrFrozenBitmapIncomplete is returned when the buffer is too small to contain a frozen bitmap.
	ErrFrozenBitmapIncomplete = errors.New("input buffer too small to contain a frozen bitmap")
	// ErrFrozenBitmapOverpopulated is returned when the number of containers is too large.
	ErrFrozenBitmapOverpopulated = errors.New("too many containers")
	// ErrFrozenBitmapUnexpectedData is returned when the buffer contains unexpected data.
	ErrFrozenBitmapUnexpectedData = errors.New("spurious data in input")
	// ErrFrozenBitmapInvalidTypecode is returned when the typecode is invalid.
	ErrFrozenBitmapInvalidTypecode = errors.New("unrecognized typecode")
	// ErrFrozenBitmapBufferTooSmall is returned when the buffer is too small.
	ErrFrozenBitmapBufferTooSmall = errors.New("buffer too small")
)

func (ra *roaringArray) frozenView(buf []byte) error { _ = "STUB: not implemented"; return nil }

// 1 byte per type, 2 bytes per key, 2 bytes per count.

// Not consuming the full input is a bug.

// GetFrozenSizeInBytes returns the size in bytes of the frozen bitmap.
func (rb *Bitmap) GetFrozenSizeInBytes() uint64 { _ = "STUB: not implemented"; return 0 }

// Freeze serializes the bitmap in the CRoaring's frozen format.
func (rb *Bitmap) Freeze() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// FreezeTo serializes the bitmap in the CRoaring's frozen format.
func (rb *Bitmap) FreezeTo(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// WriteFrozenTo serializes the bitmap in the CRoaring's frozen format.
func (rb *Bitmap) WriteFrozenTo(wr io.Writer) (int, error) {
	_ = "STUB: not implemented"
	// FIXME: this is a naive version that iterates 4 times through the
	// containers and allocates 3*len(containers) bytes; it's quite likely
	// it can be done more efficiently.
	return 0, nil
}
