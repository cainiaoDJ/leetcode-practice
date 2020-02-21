package main

import (
	"fmt"
	"sort"
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
func topKFrequent(nums []int, k int) []int {
	list := make(MyMapSlice, 0)
	for _, v := range nums {
		_, ok := list[v]
		if ok == true {
			list[v]++
		} else {
			list[v] = 1
		}

	}
	list2 := make(sortMapSlice, 0)
	for k, v := range list {
		list2 = append(list2, MyMap{k, v})
	}
	sort.Stable(list2)
	ret := make([]int, 0)
	i := 0
	for {
		if i >= k {
			return ret
		}
		ret = append(ret, list2[i].key)
		i++
	}
}

func main() {
	nums := []int{1, 2, 1, 3, 3, 3, 6, 7, 8, 4, 5, 6, 7, 4, 3, 5, 6}
	fmt.Println("ret = ", topKFrequent(nums, 2))
}
