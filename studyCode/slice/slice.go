package main

import (
	"fmt"
	"unsafe"
)

type Slice []int

func (A *Slice) Append(value int) {
	fmt.Println("A1=", unsafe.Pointer(&A))
	*A = append(*A, value)
	fmt.Println("A2=", unsafe.Pointer(&A))
}
func main() {
	var c = make([]int, 2, 5)
	fmt.Println("c1 is ", c)
	d := append(c, 1)
	fmt.Println("c2 is ", c, "ptr=", &(c))
	fmt.Println("d2 is ", d, "ptr=", &d)
	c[0] = 3
	fmt.Println("c3 is ", c)
	fmt.Println("d3 is ", d)
	c = append(c, 99)
	fmt.Println("c4 is ", c)
	fmt.Println("d4 is ", d)

	mSlice := make(Slice, 10, 20)
	fmt.Println("mSlice1=", unsafe.Pointer(&mSlice))
	mSlice.Append(5)
	fmt.Println("mSlice2=", unsafe.Pointer(&mSlice))
}
