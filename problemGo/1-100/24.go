package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var arr = []int{}
	var head *ListNode
	var last *ListNode

	for _, x := range arr {
		var tp *ListNode
		tp = new(ListNode)
		tp.Val = x
		tp.Next = nil
		if last != nil {
			last.Next = tp
		}
		if head == nil {
			head = tp
		}
		last = tp
	}
	//showList(head)
	swapPairs(head)
}

func swapPairs(head *ListNode) *ListNode {
	var new_head *ListNode
	var last *ListNode
	if head == nil || head.Next == nil {
		return head
	}
	left := head
	right := head.Next
	new_head = right
	for left != nil && right != nil {
		if last != nil {
			last.Next = right
		}
		left.Next = right.Next
		right.Next = left
		last = left
		left = left.Next
		if left == nil {
			break
		} else {
			right = left.Next
		}
	}
	return new_head
}

func showList(head *ListNode) {
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
