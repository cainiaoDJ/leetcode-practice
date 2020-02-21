package main

import (
	"container/heap"
	"fmt"
)

type MyMap struct {
	key int
	cnt int
}

type MyMapSlice map[int]int
type sortMapSlice []MyMap

func (h sortMapSlice) Less(i, j int) bool { return h[i].cnt > h[j].cnt }
func (h sortMapSlice) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h sortMapSlice) Len() int           { return len(h) }

func (h *sortMapSlice) Push(x interface{}) { *h = append(*h, x.(MyMap)) }
func (h *sortMapSlice) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func topKFrequent(nums []int, k int) []int {
	list := make(MyMapSlice, 0)
	for _, v := range nums {
		list[v]++
	}

	sms := make(sortMapSlice, 0)
	heap.Init(&sms)
	for k, v := range list {
		//sms = append(sms, MyMap{k,v})
		heap.Push(&sms, MyMap{k, v})
	}

	//sort.Stable(sms)
	ret := make([]int, 0)
	i := 0
	for {
		if i >= k {
			return ret
		}
		v := heap.Pop(&sms)
		ret = append(ret, v.(MyMap).key)
		i++
	}
}

func main() {
	nums := []int{1, 2, 1, 3, 3, 3, 6, 7, 8, 4, 5, 6, 7, 4, 3, 5, 6}
	fmt.Println("ret = ", topKFrequent(nums, 2))
}
