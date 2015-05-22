package main

import "fmt"

type heap struct {
	size     int
	heap_arr []int
}

func newheap(size int) *heap {
	h := &heap{size: 0, heap_arr: make([]int, size)}
	return h
}

func (h *heap) insert(x int) {
	h.heap_arr[h.size] = x
	h.size += 1
	return
}

func main() {
	h := newheap(10)
	h.insert(3)
	fmt.Println(h.heap_arr[0])
}
