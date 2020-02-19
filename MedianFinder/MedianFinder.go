package main

import (
	"container/heap"
	"fmt"
)

type MinHeap []int
type MaxHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	//fmt.Println("heap.pop.before:", *h)
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	//fmt.Println("heap.pop:after :", *h)
	return x
}

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	//fmt.Println("heap.pop.before:", *h)
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	//fmt.Println("heap.pop:after :", *h)
	return x
}

type MedianFinder struct {
	minHeap MinHeap
	maxHeap MaxHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	mf := MedianFinder{}
	heap.Init(&mf.minHeap)
	heap.Init(&mf.maxHeap)
	return mf
}

func (this *MedianFinder) AddNum(num int) {
	// 较小的数据放入大顶堆，较大的数据放入小顶堆。那么大顶和小顶为我们需要的数据
	// 初始放入大顶堆
	lenMax := len(this.maxHeap)
	lenMin := len(this.minHeap)
	if lenMax == 0 {
		heap.Push(&this.maxHeap, num)
	} else if lenMin == lenMax { // 大顶堆和小顶堆长度相等，根据插入数据，优先放入小顶堆
		if num >= this.minHeap[0] {
			v := heap.Pop(&this.minHeap)
			heap.Push(&this.minHeap, num)
			heap.Push(&this.maxHeap, v)
		} else {
			heap.Push(&this.maxHeap, num)
		}
	} else if lenMin > lenMax { // 小顶堆数据多，放入大顶堆 ps因为遵循了优先放入大顶堆，所以这里没有额外的判断
		heap.Push(&this.maxHeap, num)
	} else { // 大顶堆数据多，如果值大于大顶堆，那么放入小顶堆，否则把大顶堆的数据拿出来，在放入大顶堆
		if num >= this.maxHeap[0] {
			heap.Push(&this.minHeap, num)
		} else {
			v := heap.Pop(&this.maxHeap)
			heap.Push(&this.minHeap, v)
			heap.Push(&this.maxHeap, num)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	lenMax := len(this.maxHeap)
	lenMin := len(this.minHeap)
	if lenMax == lenMin {
		return float64(this.maxHeap[0]+this.minHeap[0]) / 2.0
	} else if lenMax > lenMin {
		return float64(this.maxHeap[0])
	} else {
		return float64(this.minHeap[0])
	}
}

func main() {
	//100, 16, 4, 8, 70, 2, 36, 22, 5, 12
	mf := MedianFinder{}
	//mf.AddNum(40)
	//fmt.Println("add:", mf, " mid:", mf.FindMedian())
	//mf.AddNum(12)
	//fmt.Println("add:", mf, " mid:", mf.FindMedian())
	//mf.AddNum(16)
	//fmt.Println("add:", mf, " mid:", mf.FindMedian())
	//mf.AddNum(14)
	//fmt.Println("add:", mf, " mid:", mf.FindMedian())
	//mf.AddNum(35)
	//fmt.Println("add:", mf, " mid:", mf.FindMedian())
	//mf.AddNum(19)
	//fmt.Println("add:", mf, " mid:", mf.FindMedian())
	//mf.AddNum(34)
	//fmt.Println("add:", mf, " mid:", mf.FindMedian())
	//mf.AddNum(35)
	//fmt.Println("add:", mf, " mid:", mf.FindMedian())
	//mf.AddNum(28)
	//fmt.Println("add:", mf, " mid:", mf.FindMedian())
	//mf.AddNum(35)
	//fmt.Println("add:", mf, " mid:", mf.FindMedian())
	mf.AddNum(100)
	fmt.Println("add:", mf, " mid:", mf.FindMedian())
	mf.AddNum(16)
	fmt.Println("add:", mf, " mid:", mf.FindMedian())
	mf.AddNum(4)
	fmt.Println("add:", mf, " mid:", mf.FindMedian())
	mf.AddNum(8)
	fmt.Println("add:", mf, " mid:", mf.FindMedian())
	mf.AddNum(70)
	fmt.Println("add:", mf, " mid:", mf.FindMedian())
	mf.AddNum(2)
	fmt.Println("add:", mf, " mid:", mf.FindMedian())
	mf.AddNum(36)
	fmt.Println("add:", mf, " mid:", mf.FindMedian())
	mf.AddNum(22)
	fmt.Println("add:", mf, " mid:", mf.FindMedian())
	mf.AddNum(5)
	fmt.Println("add:", mf, " mid:", mf.FindMedian())
	mf.AddNum(12)
	fmt.Println("add:", mf, " mid:", mf.FindMedian())
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
