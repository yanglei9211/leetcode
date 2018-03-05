package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Solve struct {
	valList  []int
	nodeList []*TreeNode
	idx      int
}

func newSolve() Solve {
	s := Solve{}
	s.valList = make([]int, 0)
	s.nodeList = make([]*TreeNode, 0)
	s.idx = 0
	return s
}

func (s *Solve) inorderTraversal(r *TreeNode) {
	if r == nil {
		return
	}
	s.inorderTraversal(r.Left)
	s.valList = append(s.valList, r.Val)
	//s.nodeIndex[r] = s.idx
	s.nodeList = append(s.nodeList, r)
	s.idx++
	s.inorderTraversal(r.Right)
}

func isValidBST(root *TreeNode) bool {
	s := newSolve()
	s.inorderTraversal(root)
	for i := 1; i < len(s.valList); i++ {
		if s.valList[i-1] >= s.valList[i] {
			return false
		}
	}
	return true
}

func main() {
	d4 := TreeNode{Val: 4}
	d2 := TreeNode{Val: 2}
	d3 := TreeNode{Val: 3}
	d5 := TreeNode{Val: 5}
	d7 := TreeNode{Val: 7}
	d8 := TreeNode{Val: 8}
	d4.Left = &d2
	d4.Right = &d7
	d2.Right = &d3
	d7.Left = &d5
	d7.Right = &d8
	fmt.Println(isValidBST(&d4))
}
