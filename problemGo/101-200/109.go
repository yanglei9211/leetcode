package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type ss struct {
	list *ListNode
}

func newSs(head *ListNode) ss {
	res := ss{head}
	return res
}

func (s *ss) count() int {
	//h := &ListNode{-1, s.list}
	h := s.list
	cnt := 0
	for h != nil {
		h = h.Next
		cnt++
	}
	return cnt
}

func (s *ss) do(n int) *TreeNode {
	if n == 0 {
		return nil
	}
	cur := TreeNode{}
	l := s.do(n >> 1)
	cur.Val = s.list.Val
	cur.Left = l
	s.list = s.list.Next
	r := s.do(n - (n >> 1) - 1)
	cur.Right = r
	return &cur
}

func sortedListToBST(head *ListNode) *TreeNode {
	ss := newSs(head)
	res := ss.do(ss.count())
	return res
}

func showList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d --> ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func inorderTraversal(r *TreeNode) {
	if r == nil {
		return
	}
	inorderTraversal(r.Left)
	fmt.Printf("%d ", r.Val)
	inorderTraversal(r.Right)
}

func main() {
	l := make([]ListNode, 8)
	l[7] = ListNode{7, nil}
	for i := 6; i >= 0; i-- {
		l[i] = ListNode{i, &l[i+1]}
	}

	r := sortedListToBST(&l[0])
	inorderTraversal(r)

}
