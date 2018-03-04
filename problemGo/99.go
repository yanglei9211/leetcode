package main

import (
	"fmt"
	"sort"
)

//Two elements of a binary search tree (BST) are swapped by mistake.
//
//Recover the tree without changing its structure.
//
//Note:
//A solution using O(n) space is pretty straight forward. Could you devise a constant space solution?

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

func inorderTraversal(r *TreeNode) {
	if r == nil {
		return
	}
	inorderTraversal(r.Left)
	fmt.Printf("%d ", r.Val)
	inorderTraversal(r.Right)
}

func recoverTree(root *TreeNode) {
	s := newSolve()
	s.inorderTraversal(root)
	fmt.Println(s.valList)
	sort.Ints(s.valList)
	for i := 0; i < len(s.valList); i++ {
		s.nodeList[i].Val = s.valList[i]
	}
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
	d2.Right = &d5
	d7.Left = &d3
	d7.Right = &d8
	inorderTraversal(&d4)
	fmt.Println()
	recoverTree(&d4)
	inorderTraversal(&d4)
}
