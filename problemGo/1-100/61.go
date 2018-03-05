package main

import "fmt"

//Given a list, rotate the list to the right by k places, where k is non-negative.
//
//
//Example:
//
//Given 1->2->3->4->5->NULL and k = 2,
//
//return 4->5->1->2->3->NULL.

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

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := 0
	pt := head
	var last *ListNode
	for ; pt != nil; pt = pt.Next {
		n++
		if pt.Next == nil {
			last = pt
		}
	}
	k %= n
	if k == 0 {
		return head
	}

	j := n - k
	pt = head
	i := 0
	fmt.Println(n)
	var res *ListNode
	for ; pt != nil; pt = pt.Next {
		i++
		if i == j {
			res = pt.Next
			pt.Next = nil
			last.Next = head
			rr := ListNode{}
			rr.Next = res
			return res
		}
	}
	rr := ListNode{}
	rr.Next = res
	return nil
}

func link_append(r *ListNode, x int) *ListNode {
	nxt := ListNode{Val: x, Next: nil}
	r.Next = &nxt
	return &nxt
}

func show_link(head *ListNode) {
	for pt := head; pt != nil; pt = pt.Next {
		fmt.Printf("%d -> ", pt.Val)
	}
	fmt.Println()
}

func main() {
	a := []int{1}
	head := ListNode{}
	pt := &head
	for _, x := range a {
		pt = link_append(pt, x)
	}
	show_link(head.Next)
	res := rotateRight(head.Next, 0)
	show_link(res)
}
