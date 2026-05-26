package roaring

/////////////
// The priorityQueue is used to keep Bitmaps sorted.
////////////

type item struct {
	value *Bitmap
	index int
}

type priorityQueue []*item

func (pq priorityQueue) Len() int { _ = "STUB: not implemented"; return 0 }

func (pq priorityQueue) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (pq priorityQueue) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (pq *priorityQueue) Push(x interface{}) { _ = "STUB: not implemented"; return }

func (pq *priorityQueue) Pop() interface{} { _ = "STUB: not implemented"; return nil }

// for safety

func (pq *priorityQueue) update(item *item, value *Bitmap) { _ = "STUB: not implemented"; return }

/////////////
// The containerPriorityQueue is used to keep the containers of various Bitmaps sorted.
////////////

type containeritem struct {
	value    *Bitmap
	keyindex int
	index    int
}

type containerPriorityQueue []*containeritem

func (pq containerPriorityQueue) Len() int { _ = "STUB: not implemented"; return 0 }

func (pq containerPriorityQueue) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (pq containerPriorityQueue) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (pq *containerPriorityQueue) Push(x interface{}) { _ = "STUB: not implemented"; return }

func (pq *containerPriorityQueue) Pop() interface{} { _ = "STUB: not implemented"; return nil }

// for safety

//func (pq *containerPriorityQueue) update(item *containeritem, value *Bitmap, keyindex int) {
//	item.value = value
//	item.keyindex = keyindex
//	heap.Fix(pq, item.index)
//}
