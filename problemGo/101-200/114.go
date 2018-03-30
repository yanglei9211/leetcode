package main

import (
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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zag(rt *TreeNode) *TreeNode{
    if rt.Left == nil && rt.Right == nil {
        return rt
    }
	var ll, rr *TreeNode
	if rt.Left != nil {
		ll = zag(rt.Left)
		ll.Right = rt.Right
	}

    if rt.Right != nil {
		rr = zag(rt.Right)
	}
	if rt.Left != nil {
		rt.Right = rt.Left
		rt.Left = nil
	}
	if rr != nil {
		return rr
	} else {
		return ll
	}
}

func flatten(root *TreeNode)  {
    if root != nil {
        zag(root)
    }
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
	r := make([]TreeNode, 3)
	//r[2] = TreeNode{3, nil, nil}
	r[1] = TreeNode{2, nil, nil}
	r[0] = TreeNode{1, nil, &r[1]}
	flatten(&r[0])
	inorderTraversal(&r[0])
}