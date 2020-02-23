package main

/*
   链接：https://leetcode-cn.com/problems/sliding-window-maximum/solution/hua-dong-chuang-kou-zui-da-zhi-by-leetcode-3/
   来源：力扣（LeetCode）
*/
import (
	"container/heap"
	"fmt"
)

type Heap []int

func (x Heap) Less(i, j int) bool  { return x[i] > x[j] }
func (x Heap) Len() int            { return len(x) }
func (x Heap) Swap(i, j int)       { x[i], x[j] = x[j], x[i] }
func (x *Heap) Push(i interface{}) { *x = append(*x, i.(int)) }
func (x *Heap) Pop() interface{} {
	old := *x
	l := len(old)
	v := old[l-1]
	*x = old[:l-1]
	return v
}

func maxSlidingWindow(nums []int, k int) []int {
	l := len(nums)
	if l <= 1 {
		return nums
	}
	if k > l {
		k = l
	}
	ret := make([]int, 0)
	for i := 0; i <= l-k; i++ {
		h := make(Heap, 0)
		heap.Init(&h)
		for j := i; j < k+i; j++ {
			heap.Push(&h, nums[j])
		}

		ret = append(ret, heap.Pop(&h).(int))
	}
	return ret
}

func main() {
	nums := []int{1, -1}
	ret := maxSlidingWindow(nums, 1)
	fmt.Println(ret)
}
