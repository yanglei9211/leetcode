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
func minDepth(rt *TreeNode) int {
	if rt == nil {
		return 0
	}
	if rt.Left == nil && rt.Right == nil {
		return 1
	} else if rt.Left == nil {
		rs := minDepth(rt.Right)
		return rs + 1
	} else if rt.Right == nil {
		ls := minDepth(rt.Left)
		return ls + 1
	} else {
		ls := minDepth(rt.Left)
		rs := minDepth(rt.Right)
		fmt.Println(rt.Val, ls, rs)
		if ls > rs {
			ls, rs = rs, ls
		}
		return ls + 1
	}

}
