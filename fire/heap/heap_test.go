package heapx

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	h := NewMaxHeap([]int{1, 8, 6})
	heap.Push(h, 2)
	for {
		if h.Len() > 0 {
			fmt.Println(heap.Pop(h))
		} else {
			break
		}
	}
}
func TestMin(t *testing.T) {
	h := NewMinHeap([]int{1, 8, 6})
	heap.Push(h, 2)
	for {
		if h.Len() > 0 {
			fmt.Println(heap.Pop(h))
		} else {
			break
		}
	}
}
func TestHeapx(t *testing.T) {
	less := func(s1, s2 string) bool {
		return len(s1) > len(s2)
	}
	h := NewHeapX([]string{"syzhang42", "1234", "fwefnewhjfew"}, less)
	heap.Push(h, "123456")
	for {
		if h.Len() > 0 {
			fmt.Println(heap.Pop(h))
		} else {
			break
		}
	}
}
