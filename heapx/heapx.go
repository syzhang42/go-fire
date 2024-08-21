package heapx

import "container/heap"

type Heapx[T any] struct {
	elements []T
	less     func(i, j T) bool
}

func (h Heapx[T]) Len() int           { return len(h.elements) }
func (h Heapx[T]) Less(i, j int) bool { return h.less(h.elements[i], h.elements[j]) }
func (h Heapx[T]) Swap(i, j int)      { h.elements[i], h.elements[j] = h.elements[j], h.elements[i] }

func (h *Heapx[T]) Push(x interface{}) {
	h.elements = append(h.elements, x.(T))
}

func (h *Heapx[T]) Pop() interface{} {
	old := h.elements
	n := len(old)
	x := old[n-1]
	h.elements = old[0 : n-1]
	return x
}

func NewMaxHeap(data []int) *Heapx[int] {
	temp := []int{}
	temp = append(temp, data...)

	h := &Heapx[int]{
		elements: temp,
		less:     func(i, j int) bool { return i > j },
	}
	heap.Init(h)
	return h
}
func NewMinHeap(data []int) *Heapx[int] {
	temp := []int{}
	temp = append(temp, data...)
	h := &Heapx[int]{
		elements: temp,
		less:     func(i, j int) bool { return i < j },
	}
	heap.Init(h)
	return h
}

// Max heap: use > instead of <
func NewHeapX[T any](data []T, less func(i, j T) bool) *Heapx[T] {
	temp := []T{}
	temp = append(temp, data...)
	h := &Heapx[T]{
		elements: temp,
		less:     less,
	}
	heap.Init(h)
	return h
}
