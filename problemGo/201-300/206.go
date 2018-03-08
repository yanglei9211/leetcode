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

func reversePre(head *ListNode, n int) *ListNode {
	var p0, p1 *ListNode
	cnt := 1

	p0 = head
	p1 = head.Next
	for p0 != nil && p1 != nil && cnt < n {
		pNext := p1.Next
		p1.Next = p0
		p0 = p1
		p1 = pNext
		cnt++
	}
	head.Next = p1
	return p0
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	res := reversePre(head, 0x3f3f3f3f)
	return res
}

func showList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d --> ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	a := []int{3, 5, 18, 22, 59, 126}
	h := ListNode{Val: 0, Next: nil}
	for i := len(a) - 1; i >= 0; i-- {
		tp := ListNode{Val: a[i], Next: h.Next}
		h.Next = &tp
	}
	showList(h.Next)
	res := reverseList(h.Next)
	showList(res)
}
