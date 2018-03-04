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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m == n {
		return head
	}
	if m == 1 {
		res := reversePre(head, n-m+1)
		return res
	}
	var p *ListNode
	p = head
	cnt := 1
	for p != nil && cnt+1 < m {
		p = p.Next
		cnt++
	}
	nh := reversePre(p.Next, n-m+1)
	p.Next = nh
	return head
}

func showList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d --> ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	a := []int{3, 5}
	h := ListNode{Val: 0, Next: nil}
	for i := len(a) - 1; i >= 0; i-- {
		tp := ListNode{Val: a[i], Next: h.Next}
		h.Next = &tp
	}
	showList(h.Next)
	res := reverseBetween(h.Next, 1, 2)
	showList(res)
}
