package main

import (
	"fmt"
)

//func build(preorder []int, inorder []int) *TreeNode {
//
//}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	inIndex := map[int]int{}
	for i := 0; i < len(preorder); i++ {
		inIndex[inorder[i]] = i
	}
	vrt := preorder[0]
	idx, _ := inIndex[vrt]
	left := buildTree(preorder[1:idx+1], inorder[:idx])
	right := buildTree(preorder[idx+1:], inorder[idx+1:])
	rt := TreeNode{Val: vrt, Left: left, Right: right}
	return &rt
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
	pre := []int{3, 9, 20, 15, 7}
	in := []int{9, 3, 15, 20, 7}
	res := buildTree(pre, in)
	inorderTraversal(res)
}
