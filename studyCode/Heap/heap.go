package main

import "fmt"

type IntSlice []int

type Heap struct {
	size int
	data IntSlice
}

func (this *Heap) Heap(data IntSlice) {
	for _, v := range data {
		this.data = append(this.data, v)
		this.size++
	}
}

func (this *Heap) ShiftUp(pos int) {
	for {
		if pos <= 1 {
			break
		}
		// 子节点比父节点大，上浮
		if this.data[pos] < this.data[pos/2] {
			this.data[pos], this.data[pos/2] = this.data[pos/2], this.data[pos]
		}
		pos >>= 1
	}
}

func (this *Heap) ShiftDown(pos int) {
	for {
		if pos<<1 > this.size {
			break
		}
		next := pos << 1
		// 父节点比子节点大，下沉
		if next < this.size {
			if this.data[pos] > this.data[next] {
				if this.data[next] > this.data[next+1] {
					next++
				}
				this.data[pos], this.data[next] = this.data[next], this.data[pos]
				pos = next
			} else {
				return
			}

		} else {
			return
		}

	}
}

func (this *Heap) Push(num int) {
	this.size++
	this.data = append(this.data, num)
	// this.data[this.size] = num
	this.ShiftUp(this.size)
}

func (this *Heap) Pop() {
	this.data[this.size], this.data[1] = this.data[1], this.data[this.size]
	this.size--
	this.data = this.data[:this.size+1]
	this.ShiftDown(1)
}

func (this *Heap) Top() int {
	return this.data[1]
}

// 这个排序似乎是有问题的
// func (this *Heap) HeapSort() {
// 	m := this.size
// 	for i := 1; i <= this.size; i++ {
// 		this.data[m], this.data[i] = this.data[i], this.data[m]
// 		m--
// 		this.ShiftDown(i)
// 	}
// }

func main() {
	heap := Heap{0, make([]int, 1)}
	heap.Push(6)
	fmt.Println("heap=", heap)
	heap.Push(5)

	fmt.Println("heap=", heap)
	heap.Push(4)

	fmt.Println("heap=", heap)
	heap.Push(3)

	fmt.Println("heap=", heap)
	heap.Push(2)

	fmt.Println("heap=", heap)
	heap.Push(1)

	fmt.Println("heap=", heap)

	heap.Pop()
	fmt.Println("pop =", heap)

	heap.Pop()
	fmt.Println("pop =", heap)
}
