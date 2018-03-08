package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	rt := postorder[len(postorder)-1]
	var pos int
	for pos = 0; pos < len(inorder); pos++ {
		if inorder[pos] == rt {
			break
		}
	}
	left := buildTree(inorder[:pos], postorder[:pos])
	right := buildTree(inorder[pos+1:], postorder[pos:len(postorder)-1])
	res := TreeNode{Val: rt, Left: left, Right: right}
	return &res
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
	post := []int{9, 15, 7, 20, 3}
	in := []int{9, 3, 15, 20, 7}
	res := buildTree(in, post)
	inorderTraversal(res)
}
