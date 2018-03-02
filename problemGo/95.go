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

func TreeCopy(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	tp := TreeNode{Val: root.Val}
	tp.Left = TreeCopy(root.Left)
	tp.Right = TreeCopy(root.Right)
	return &tp
}

type Solve struct {
	root *TreeNode
	n    int
}

func (s Solve) Save() {
	fmt.Println("assss")
	inorderTraversal(s.root)
}

func genTree(l, r int) []*TreeNode {
	if l > r {
		return []*TreeNode{nil}
	}
	res := []*TreeNode{}
	for k := l; k <= r; k++ {
		//rt := TreeNode{Val: k}
		lefts := genTree(l, k-1)
		rights := genTree(k+1, r)
		for i := 0; i < len(lefts); i++ {
			for j := 0; j < len(rights); j++ {
				rt := TreeNode{Val: k, Left: lefts[i], Right: rights[j]}
				res = append(res, &rt)
			}
		}
	}

	return res
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	res := genTree(1, n)
	return res
}

func numTrees(n int) int {
	if n == 0 {
		return 0
	}
	res := genTree(1, n)
	return len(res)
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
	generateTrees(3)
	//d1 := TreeNode{Val: 1}
	//d5 := TreeNode{Val: 5}
	//d3 := TreeNode{Val: 3}
	//d2 := TreeNode{Val: 2}
	//d7 := TreeNode{Val: 7}
	//d1.Left = &d5
	//d1.Right = &d3
	//d3.Left = &d2
	//d3.Right = &d7
	//r := TreeCopy(&d1)
	//inorderTraversal(r)
	//fmt.Println()
	//inorderTraversal(&d1)
}
