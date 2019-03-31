package sort

// Heap data structure
type Heap struct {
	slice    []int
	heapSize int
}

// Len returns length of Heap instance
func (h *Heap) Len() int {
	return len(h.slice)
}

// MaxHeapify float down element i if it's smaller than its children nodes
func MaxHeapify(h *Heap, i int) {
	l := i*2 + 1
	r := i*2 + 2

	largest := i

	if l < h.heapSize && h.slice[l] > h.slice[i] {
		largest = l
	}

	if r < h.heapSize && h.slice[r] > h.slice[largest] {
		largest = r
	}

	if largest != i {
		h.slice[i], h.slice[largest] = h.slice[largest], h.slice[i]
		MaxHeapify(h, largest)
	}
}

// BuildMaxHeap initialsize a max heap from an array
func BuildMaxHeap(a []int) *Heap {
	h := &Heap{slice: a, heapSize: len(a)}
	for i := h.Len() / 2; i >= 0; i-- {
		MaxHeapify(h, i)
	}

	return h
}

// HeapSort sort an array
func HeapSort(a []int) {
	h := BuildMaxHeap(a)
	for i := h.Len() - 1; i > 0; i-- {
		h.slice[0], h.slice[i] = h.slice[i], h.slice[0]
		h.heapSize--
		MaxHeapify(h, 0)
	}
}
