package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sycInorder(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	l := sycInorder(p.Left, q.Left)
	r := sycInorder(p.Right, q.Right)
	return l && r
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return sycInorder(p, q)
}

func g1() *TreeNode {
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
	return &d4
}

func g2() *TreeNode {
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
	return &d4
}

func main() {
	fmt.Println(isSameTree(g1(), g2()))
}
