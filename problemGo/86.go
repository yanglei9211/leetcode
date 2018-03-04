package main

import "fmt"

//Given a linked list and a value x, partition it such that all nodes less than x come before nodes greater than or equal to x.
//
//You should preserve the original relative order of the nodes in each of the two partitions.
//
//For example,
//Given 1->4->3->2->5->2 and x = 3,
//return 1->2->2->4->3->5.

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

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	gt := ListNode{Val: 0, Next: nil}
	gpt := &gt

	nh := &ListNode{Val: 0, Next: head}
	head = nh
	pt := head.Next
	last := head
	for pt != nil {
		nt := pt.Next
		if pt.Val >= x {
			last.Next = pt.Next
			pt.Next = nil
			gpt.Next = pt
			gpt = gpt.Next
			pt = nt
		} else {
			last = pt
			pt = nt
		}
	}
	last.Next = gt.Next
	return head.Next
}

func main() {
	//a := []int{1,4,3,2,5,2}
	a := []int{1}
	head := ListNode{}
	for _, x := range a {
		tmp := ListNode{Val: x, Next: head.Next}
		head.Next = &tmp
	}
	partition(head.Next, 0)
}

func show(s *ListNode) {
	for s != nil {
		fmt.Printf("%d->", s.Val)
		s = s.Next
	}
	fmt.Println()
}
