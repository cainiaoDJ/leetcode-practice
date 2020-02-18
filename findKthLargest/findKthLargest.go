package main

import (
	"fmt"
	"sort"
)

type array []int

func (s array) Len() int           { return len(s) }
func (s array) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s array) Less(i, j int) bool { return s[i] < s[j] }

func findKthLargest(nums []int, k int) int {
	if len(nums) < k {
		return 0
	}
	sort.Sort(array(nums))
	fmt.Println(nums)
	return nums[len(nums)-k]
}
func main() {
	// fmt.Println("hello")
	n := []int{1}
	ret := findKthLargest(n, 1)
	fmt.Println(ret)
}
