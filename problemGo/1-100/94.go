package main

import (
	"container/list"
	"errors"
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Stack struct {
	list *list.List
}

func NewStack() *Stack {
	ls := list.New()
	return &Stack{ls}
}

func (s *Stack) Push(value interface{}) {
	s.list.PushBack(value)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("stack empty")
	} else {
		e := s.list.Back()
		if e != nil {
			s.list.Remove(e)
			return e.Value, nil
		}
		return nil, errors.New("inter error")
	}
}

func (s *Stack) Top() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("stack empty")
	} else {
		e := s.list.Back()
		if e != nil {
			return e.Value, nil
		}
		return nil, errors.New("inter error")
	}
}

func (s *Stack) Len() int {
	return s.list.Len()
}

func (s *Stack) Empty() bool {
	return s.list.Len() == 0
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func doLeft(st *Stack, node *TreeNode) {
	for node != nil {
		st.Push(node)
		node = node.Left
	}
}

func inorderTraversal(root *TreeNode) []int {
	st := NewStack()
	ans := []int{}
	doLeft(st, root)
	for st.Len() > 0 {
		//doLeft(st, root)
		for !st.Empty() {
			f, _ := st.Pop()
			node := f.(*TreeNode)
			ans = append(ans, node.Val)
			if node.Right != nil {
				root = node.Right
				doLeft(st, root)
				break
			}
		}
	}
	return ans
}

func main() {
	d1 := TreeNode{Val: 1}
	d5 := TreeNode{Val: 5}
	d3 := TreeNode{Val: 3}
	d2 := TreeNode{Val: 2}
	d7 := TreeNode{Val: 7}
	d1.Left = &d5
	d1.Right = &d3
	d3.Left = &d2
	d3.Right = &d7
	res := inorderTraversal(&d1)
	fmt.Println(res)
}
