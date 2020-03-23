package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	i := 0
	slow, fast := head, head
	for slow.Next != nil && fast.Next != nil {
		if i%2 == 0 {
			fast = fast.Next
		} else {
			slow = slow.Next
			fast = fast.Next
		}
		i++
	}
	// 这里反着来,是因为上面i多加1了
	if i%2 == 1 {
		return slow.Next
	} else {
		return slow
	}
}

func (l *ListNode) initListNode(list []int) {
	head := l
	head.Val = list[0]
	for _, v := range list[1:] {
		head.Next = &ListNode{v, nil}
		head = head.Next
	}
}

func main() {
	var a = []int{1, 2, 3, 4}
	var ln = &ListNode{}
	ln.initListNode(a)
	fmt.Println(ln)
	fmt.Println(middleNode(ln))
}
