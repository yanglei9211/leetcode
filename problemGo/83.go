package main

import "fmt"

//Given a sorted linked list, delete all nodes that have duplicate numbers, leaving only distinct
// numbers from the original list.
//
//For example,
//Given 1->2->3->3->4->4->5, return 1->2->5.
//Given 1->1->1->2->3, return 2->3.
//

type ListNode struct {
	Val  int
	Next *ListNode
}

const nullint = -999999999

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	ttt := head
	head = &ListNode{Val: 0, Next: ttt}
	var v, cnt int
	var pt *ListNode
	pt = head
	v = nullint
	cnt = 0
	pt = pt.Next
	var cur *ListNode
	cur = head
	var tmp *ListNode
	tmp = pt
	for pt != nil {
		if v != pt.Val {
			fmt.Println(v, cnt)
			if cnt >= 1 {
				cur.Next = tmp
				cur = tmp
				cnt = 1
				tmp = pt
				v = pt.Val
			} else {
				tmp = pt
				cnt = 1
				v = pt.Val
			}

		} else {
			cnt++
		}
		pt = pt.Next
	}
	if cnt >= 1 {
		cur.Next = tmp
		cur = tmp
		cur.Next = nil
	} else {
		cur.Next = nil
	}
	return head.Next
}
func show(s *ListNode) {
	for s != nil {
		fmt.Printf("%d->", s.Val)
		s = s.Next
	}
	fmt.Println()
}

func main() {
	a := []int{3, 3, 2, 1, 1}
	head := ListNode{}
	for _, x := range a {
		tmp := ListNode{Val: x, Next: head.Next}
		head.Next = &tmp
	}
	show(head.Next)
	res := deleteDuplicates(&head)
	show(res)
}
