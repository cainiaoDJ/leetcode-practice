package main

import (
	"container/heap"
	"fmt"
)

type Heap []int

func (h Heap) Less(i, j int) bool  { return h[i] < h[j] }
func (h Heap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h Heap) Len() int            { return len(h) }
func (h *Heap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *Heap) Pop() interface{} {
	old := *h
	l := len(*h)
	x := old[l-1]
	*h = old[:l-1]
	return x
}
func kthSmallest(matrix [][]int, k int) int {
	h := Heap{}
	heap.Init(&h)
	for _, r := range matrix {
		for _, v := range r {
			heap.Push(&h, v)
		}
	}
	num := 0
	for i := 0; i < k; i++ {
		num = heap.Pop(&h).(int)
	}

	return num
}

func main() {
	input := [][]int{
		{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}
	fmt.Println(kthSmallest(input, 8))
}
