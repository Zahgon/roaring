package roaring

type shortIterable interface {
	hasNext() bool
	next() uint16
}

type shortPeekable interface {
	shortIterable
	peekNext() uint16
	advanceIfNeeded(minval uint16)
}

type shortIterator struct {
	slice []uint16
	loc   int
}

func (si *shortIterator) hasNext() bool { _ = "STUB: not implemented"; return false }

func (si *shortIterator) next() uint16 { _ = "STUB: not implemented"; return 0 }

func (si *shortIterator) peekNext() uint16 { _ = "STUB: not implemented"; return 0 }

func (si *shortIterator) advanceIfNeeded(minval uint16) { _ = "STUB: not implemented"; return }

type reverseIterator struct {
	slice []uint16
	loc   int
}

func (si *reverseIterator) hasNext() bool { _ = "STUB: not implemented"; return false }

func (si *reverseIterator) next() uint16 { _ = "STUB: not implemented"; return 0 }

type arrayContainerUnsetIterator struct {
	content []uint16
	// pos is the index of the next set bit that is >= nextVal.
	// When nextVal reaches content[pos], pos is incremented.
	pos     int
	nextVal int
}

func (acui *arrayContainerUnsetIterator) next() uint16 { _ = "STUB: not implemented"; return 0 }

func (acui *arrayContainerUnsetIterator) hasNext() bool { _ = "STUB: not implemented"; return false }

func (acui *arrayContainerUnsetIterator) peekNext() uint16 { _ = "STUB: not implemented"; return 0 }

func (acui *arrayContainerUnsetIterator) advanceIfNeeded(minval uint16) {
	_ = "STUB: not implemented"
	return
}

func newArrayContainerUnsetIterator(content []uint16) *arrayContainerUnsetIterator {
	_ = "STUB: not implemented"
	return nil
}
